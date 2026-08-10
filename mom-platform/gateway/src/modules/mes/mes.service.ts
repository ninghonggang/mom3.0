import { Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { ResolverService } from '../../grpc/resolver.service';
import { buildListRequest, toProtoTimestamp } from '../../grpc/case.util';
import {
  CreateOrderDto,
  UpdateOrderStatusDto,
  DispatchDto,
  ReportDto,
  CompleteDto,
} from './dto/mes.dto';

const PKG = 'mom.mes';
const URL = process.env.MES_SERVICE_URL || 'localhost:50052';

/**
 * MES BFF service — 委托给 MES gRPC 微服务 (mes-service:50052)。
 */
@Injectable()
export class MesService {
  constructor(
    private readonly grpc: GrpcService,
    private readonly resolver: ResolverService,
  ) {}

  listOrders(query: any) {
    return this.grpc.call(PKG, 'ProductionOrderService', 'ListOrders', buildListRequest(query, 'pagination'), URL);
  }
  async createOrder(dto: CreateOrderDto) {
    // 工单必须挂在真实物料上：先把 productCode 解析成 MDM material_id
    const mat = await this.resolver.material(dto.productCode);
    return this.grpc.call(
      PKG,
      'ProductionOrderService',
      'CreateOrder',
      {
        order: {
          order_no: dto.orderNo,
          material_id: mat.id,
          material_code: mat.materialCode,
          material_name: mat.materialName,
          material_spec: mat.spec,
          quantity: dto.plannedQuantity,
          workshop_id: Number(dto.workshopId) || 0,
          line_id: Number(dto.lineId) || 0,
          plan_start_date: toProtoTimestamp(dto.plannedStartTime),
          plan_end_date: toProtoTimestamp(dto.plannedEndTime),
        },
      },
      URL,
    );
  }

  getOrder(id: string) {
    return this.grpc.call(PKG, 'ProductionOrderService', 'GetOrder', { id }, URL);
  }
  updateOrderStatus(id: string, dto: UpdateOrderStatusDto) {
    return this.grpc.call(PKG, 'ProductionOrderService', 'UpdateOrderStatus', { id, ...dto }, URL);
  }
  dispatchOrder(id: string, dto: DispatchDto) {
    // proto: CreateDispatchRequest { order_id, repeated Dispatch dispatches }
    return this.grpc.call(
      PKG,
      'DispatchService',
      'CreateDispatch',
      {
        order_id: Number(id) || 0,
        dispatches: [
          {
            order_id: Number(id) || 0,
            line_id: Number(dto.lineId) || 0,
            workstation_id: Number(dto.workstationId) || 0,
            employee_id: Number(dto.operatorId) || 0,
          },
        ],
      },
      URL,
    );
  }
  reportOrder(id: string, dto: ReportDto) {
    // proto: CreateJobReportRequest { MobileJobReport report }
    const good = Number(dto.goodQuantity) || 0;
    const scrap = Number(dto.scrapQuantity) || 0;
    return this.grpc.call(
      PKG,
      'JobReportService',
      'CreateReport',
      {
        report: {
          order_id: Number(id) || 0,
          employee_id: Number(dto.operatorId) || 0,
          workstation_id: Number(dto.workstationId) || 0,
          reported_qty: good + scrap,
          qualified_qty: good,
          defective_qty: scrap,
          remark: dto.remark,
        },
      },
      URL,
    );
  }
  async completeOrder(id: string, dto: CompleteDto) {
    // proto: CreateCompleteRequest { ProductionComplete complete }
    return this.grpc.call(
      PKG,
      'ProductionCompleteService',
      'CreateComplete',
      {
        complete: {
          order_id: Number(id) || 0,
          warehouse_id: await this.resolver.warehouseId(dto.warehouseId ?? ''),
          location_id: await this.resolver.locationId(dto.locationCode),
          quantity: dto.actualQuantity,
          batch_no: dto.batchNo ?? '',
          complete_time: toProtoTimestamp(dto.actualEndTime),
        },
      },
      URL,
    );
  }
}
