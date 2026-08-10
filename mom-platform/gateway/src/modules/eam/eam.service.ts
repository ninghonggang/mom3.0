import { Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { buildListRequest, toProtoEnum, toInt, toDecimalString } from '../../grpc/case.util';
import {
  CreateEquipmentDto,
  CreateRepairOrderDto,
  UpdateRepairOrderDto,
  StartDowntimeDto,
  ResolveDowntimeDto,
  SaveOeeDto,
} from './dto/eam.dto';

const PKG = 'mom.eam';
const URL = process.env.EAM_SERVICE_URL || 'localhost:50054';

/**
 * EAM BFF service — 委托给 EAM gRPC 微服务 (eam-service:50054)。
 *
 * 注意：proto 中所有枚举都带域前缀（EQUIPMENT_TYPE_ / URGENCY_ /
 * REPAIR_STATUS_ / DOWNTIME_TYPE_），REST 层接受裸值，这里统一用
 * toProtoEnum 补前缀，否则 proto-loader 会退化成 *_UNSPECIFIED(0)，
 * 导致服务端状态机校验失败（FAILED_PRECONDITION）。
 */
@Injectable()
export class EamService {
  constructor(private readonly grpc: GrpcService) {}

  listEquipment(query: any) {
    return this.grpc.call(PKG, 'EamService', 'ListEquipment', buildListRequest(query, 'page'), URL);
  }

  createEquipment(dto: CreateEquipmentDto) {
    return this.grpc.call(
      PKG,
      'EamService',
      'CreateEquipment',
      {
        equipment_code: dto.equipmentCode,
        equipment_name: dto.name,
        equipment_type: toProtoEnum('EQUIPMENT_TYPE', dto.category ?? 'MACHINE'),
        model: dto.model ?? '',
        specification: dto.specification ?? '',
        workshop_id: toInt(dto.workshopId),
        line_id: toInt(dto.lineId),
        target_oee: dto.targetOee ?? '',
      },
      URL,
    );
  }

  getEquipment(id: string) {
    return this.grpc.call(PKG, 'EamService', 'GetEquipment', { id: toInt(id) }, URL);
  }

  listRepairOrders(query: any) {
    return this.grpc.call(PKG, 'EamService', 'ListRepairOrders', buildListRequest(query, 'page'), URL);
  }

  createRepairOrder(dto: CreateRepairOrderDto) {
    return this.grpc.call(
      PKG,
      'EamService',
      'CreateRepairOrder',
      {
        equipment_id: toInt(dto.equipmentId),
        fault_type: dto.faultCategory ?? '',
        fault_desc: dto.faultDescription,
        urgency: toProtoEnum('URGENCY', dto.priority ?? 'NORMAL'),
        reporter_id: toInt(dto.reporterId),
      },
      URL,
    );
  }

  updateRepairOrder(id: string, dto: UpdateRepairOrderDto) {
    return this.grpc.call(
      PKG,
      'EamService',
      'UpdateRepairOrder',
      {
        id: toInt(id),
        status: toProtoEnum('REPAIR_STATUS', dto.status),
        repairman_id: toInt(dto.technicianId),
      },
      URL,
    );
  }

  startDowntime(dto: StartDowntimeDto) {
    return this.grpc.call(
      PKG,
      'EamService',
      'StartDowntime',
      {
        equipment_id: toInt(dto.equipmentId),
        downtime_type: toProtoEnum('DOWNTIME_TYPE', dto.type ?? 'UNPLANNED'),
        reason: dto.reason,
      },
      URL,
    );
  }

  // 解决方案与处理人一并落库，停机记录才有复盘价值。
  resolveDowntime(id: string, dto: ResolveDowntimeDto) {
    return this.grpc.call(
      PKG,
      'EamService',
      'ResolveDowntime',
      {
        id: toInt(id),
        resolution: dto?.resolution ?? '',
        resolver_id: dto?.resolverId ?? '',
      },
      URL,
    );
  }

  /**
   * ListOeeRequest 没有分页字段，只有 equipment_id / begin_date / end_date，
   * 因此这里显式构造请求，而不是套用 buildListRequest（会塞入服务端忽略的 page）。
   */
  getOee(query: any) {
    return this.grpc.call(
      PKG,
      'EamService',
      'ListOee',
      {
        equipment_id: query?.equipmentId ? toInt(query.equipmentId) : 0,
        begin_date: query?.beginDate ?? '',
        end_date: query?.endDate ?? '',
      },
      URL,
    );
  }

  // 上报 OEE 三要素；OEE 由服务端计算，同一 (设备, 日期) 幂等覆盖。
  saveOee(dto: SaveOeeDto) {
    return this.grpc.call(
      PKG,
      'EamService',
      'SaveOee',
      {
        equipment_id: toInt(dto.equipmentId),
        calc_date: dto.calcDate ?? '',
        availability: toDecimalString(dto.availability),
        performance: toDecimalString(dto.performance),
        quality: toDecimalString(dto.quality),
      },
      URL,
    );
  }
}
