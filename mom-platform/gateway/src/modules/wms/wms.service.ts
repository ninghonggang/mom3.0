import { BadRequestException, Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { ResolverService } from '../../grpc/resolver.service';
import { buildListRequest, toDecimalString, toInt, toProtoEnum } from '../../grpc/case.util';
import {
  CreateWarehouseDto,
  CreateLocationDto,
  CreateReceiveOrderDto,
  ConfirmReceiveDto,
  PutawayDto,
  CreateDeliveryOrderDto,
  PickDto,
  ShipDto,
  CreateCountPlanDto,
  CreateCountRecordDto,
} from './dto/wms.dto';

const PKG = 'mom.wms';
const URL = process.env.WMS_SERVICE_URL || 'localhost:50055';

/**
 * WMS BFF service — 委托给 WMS gRPC 微服务 (wms-service:50055)。
 *
 * REST 层用「物料编码 + 数量」描述单据，gRPC 契约用「material_id + 字符串数量」，
 * 且确认/上架/拣货是**行级**操作（需要 line_id）。本 service 负责这层翻译：
 * 编码经 ResolverService 解析成 id；行级动作先回查单据拿到 line_id 再下发，
 * 让前端不必感知内部主键。
 */
@Injectable()
export class WmsService {
  constructor(
    private readonly grpc: GrpcService,
    private readonly resolver: ResolverService,
  ) {}

  listWarehouses(query: any) {
    return this.grpc.call(PKG, 'WmsService', 'ListWarehouses', buildListRequest(query, 'page'), URL);
  }

  async createWarehouse(dto: CreateWarehouseDto) {
    const warehouse = await this.grpc.call(
      PKG,
      'WmsService',
      'CreateWarehouse',
      {
        warehouse_code: dto.warehouseCode,
        warehouse_name: dto.name,
        warehouse_type: dto.type || 'RAW',
      },
      URL,
    );
    this.resolver.invalidateWarehouses();
    return warehouse;
  }

  listLocations(query: any) {
    return this.grpc.call(PKG, 'WmsService', 'ListLocations', buildListRequest(query, 'page'), URL);
  }

  async createLocation(dto: CreateLocationDto) {
    const location = await this.grpc.call(
      PKG,
      'WmsService',
      'CreateLocation',
      {
        warehouse_id: await this.resolver.warehouseId(dto.warehouseId),
        area_id: toInt(dto.areaId),
        location_code: dto.locationCode,
        location_type: toProtoEnum('LOCATION_TYPE', dto.locationType ?? 'STORAGE'),
        capacity: toDecimalString(dto.capacity ?? 0),
      },
      URL,
    );
    // 新建库位后清缓存，避免后续按编码解析时命中旧的“不存在”结果
    this.resolver.invalidateLocations();
    return location;
  }

  listBalances(query: any) {
    return this.grpc.call(PKG, 'WmsService', 'ListBalances', buildListRequest(query, 'page'), URL);
  }

  // ---------- 入库 ----------

  async createReceiveOrder(dto: CreateReceiveOrderDto) {
    if (!dto.lines?.length) {
      throw new BadRequestException('lines must contain at least one item');
    }
    const lines = await Promise.all(
      dto.lines.map(async (l) => ({
        material_id: await this.resolver.materialId(l.materialCode),
        expected_qty: toDecimalString(l.quantity),
        unit_price: toDecimalString(l.unitPrice ?? 0),
      })),
    );
    return this.grpc.call(
      PKG,
      'WmsService',
      'CreateReceiveOrder',
      {
        po_id: toInt(dto.poNo?.replace(/\D/g, '')),
        supplier_id: toInt(dto.supplierId),
        lines,
      },
      URL,
    );
  }

  /**
   * 确认收货。未显式给出实收数量时，按订单行的应收数量全量确认。
   * `actualQuantities` 支持以 lineId 或物料编码为键。
   */
  async confirmReceive(id: string, dto: ConfirmReceiveDto) {
    const order: any = await this.grpc.call(PKG, 'WmsService', 'GetReceiveOrder', { id: toInt(id) }, URL);
    const orderLines: any[] = order?.lines ?? [];
    if (!orderLines.length) {
      throw new BadRequestException(`receive order ${id} has no lines`);
    }
    const actual = dto.actualQuantities ?? {};
    const lines = orderLines.map((l) => {
      const byId = actual[String(l.id)];
      const qty = byId !== undefined ? byId : Number(l.expectedQty ?? 0);
      return {
        line_id: toInt(l.id),
        received_qty: toDecimalString(qty),
        batch_no: l.batchNo || '',
        expire_date: l.expireDate || '',
      };
    });
    return this.grpc.call(PKG, 'WmsService', 'ReceiveConfirm', { id: toInt(id), lines }, URL);
  }

  /** 上架。把订单每一行的实收数量放进指定库位。 */
  async putaway(id: string, dto: PutawayDto) {
    const order: any = await this.grpc.call(PKG, 'WmsService', 'GetReceiveOrder', { id: toInt(id) }, URL);
    const orderLines: any[] = order?.lines ?? [];
    if (!orderLines.length) {
      throw new BadRequestException(`receive order ${id} has no lines`);
    }
    const locationId = await this.resolver.locationId(dto.locationCode, undefined, { required: true });
    const lines = orderLines.map((l) => ({
      line_id: toInt(l.id),
      location_id: locationId,
      quantity: toDecimalString(l.receivedQty ?? l.expectedQty ?? 0),
    }));
    return this.grpc.call(PKG, 'WmsService', 'Putaway', { receive_order_id: toInt(id), lines }, URL);
  }

  // ---------- 出库 ----------

  async createDeliveryOrder(dto: CreateDeliveryOrderDto) {
    if (!dto.lines?.length) {
      throw new BadRequestException('lines must contain at least one item');
    }
    const lines = await Promise.all(
      dto.lines.map(async (l) => ({
        material_id: await this.resolver.materialId(l.materialCode),
        ordered_qty: toDecimalString(l.quantity),
      })),
    );
    return this.grpc.call(
      PKG,
      'WmsService',
      'CreateDeliveryOrder',
      {
        so_id: toInt(dto.soNo?.replace(/\D/g, '')),
        customer_id: toInt(dto.customerId),
        lines,
      },
      URL,
    );
  }

  /** 拣货。按订单行的需求量从指定库位全量拣出。 */
  async pickDelivery(id: string, dto: PickDto) {
    const order: any = await this.grpc.call(PKG, 'WmsService', 'GetDeliveryOrder', { id: toInt(id) }, URL);
    const orderLines: any[] = order?.lines ?? [];
    if (!orderLines.length) {
      throw new BadRequestException(`delivery order ${id} has no lines`);
    }
    const codes = dto.locationCodes ?? [];
    const pickerId = toInt(dto.operatorId);
    const lines = await Promise.all(
      orderLines.map(async (l, i) => ({
        line_id: toInt(l.id),
        location_id: await this.resolver.locationId(codes[i] ?? codes[0], undefined, { required: true }),
        picked_qty: toDecimalString(l.orderedQty ?? 0),
        picker_id: pickerId,
      })),
    );
    return this.grpc.call(PKG, 'WmsService', 'PickItems', { delivery_order_id: toInt(id), lines }, URL);
  }

  shipDelivery(id: string, _dto: ShipDto) {
    return this.grpc.call(PKG, 'WmsService', 'ShipOrder', { id: toInt(id) }, URL);
  }

  // ---------- 盘点 ----------

  async createCountPlan(dto: CreateCountPlanDto) {
    return this.grpc.call(
      PKG,
      'WmsService',
      'CreateCountPlan',
      {
        warehouse_id: await this.resolver.warehouseId(dto.warehouseId),
        plan_type: dto.countType || 'FULL',
      },
      URL,
    );
  }

  async createCountRecord(dto: CreateCountRecordDto) {
    return this.grpc.call(
      PKG,
      'WmsService',
      'SubmitCountRecord',
      {
        plan_id: toInt(dto.planId),
        items: [
          {
            material_id: await this.resolver.materialId(dto.materialCode),
            location_id: await this.resolver.locationId(dto.locationCode),
            actual_qty: toDecimalString(dto.countedQuantity),
          },
        ],
      },
      URL,
    );
  }
}
