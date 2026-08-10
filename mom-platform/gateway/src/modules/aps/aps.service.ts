import { Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { ResolverService } from '../../grpc/resolver.service';
import { buildListRequest, toInt, toProtoEnum, toProtoTimestamp } from '../../grpc/case.util';
import { CreateMpsPlanDto, GenerateMrpDto, CreateScheduleJobDto } from './dto/aps.dto';

const PKG = 'mom.aps';
const URL = process.env.APS_SERVICE_URL || 'localhost:50056';

/**
 * APS BFF service — 委托给 APS gRPC 微服务 (aps-service:50056)。
 */
@Injectable()
export class ApsService {
  constructor(
    private readonly grpc: GrpcService,
    private readonly resolver: ResolverService,
  ) {}

  listMpsPlans(query: any) {
    return this.grpc.call(PKG, 'MPSService', 'ListMPS', buildListRequest(query, 'pagination'), URL);
  }

  async createMpsPlan(dto: CreateMpsPlanDto) {
    return this.grpc.call(
      PKG,
      'MPSService',
      'CreateMPS',
      {
        mps_no: dto.planNo,
        plan_month: dto.planMonth || (dto.startDate ?? new Date().toISOString()).slice(0, 7),
        material_id: await this.resolver.materialId(dto.productCode),
        quantity: dto.plannedQuantity,
        remark: dto.remark ?? '',
      },
      URL,
    );
  }

  releaseMpsPlan(id: string, operator?: string) {
    return this.grpc.call(
      PKG,
      'MPSService',
      'ReleaseMPS',
      { id: toInt(id), operator: operator ?? 'system' },
      URL,
    );
  }

  /** MRP 基于一份已存在的 MPS 展开物料需求。 */
  generateMrp(dto: GenerateMrpDto) {
    return this.grpc.call(
      PKG,
      'MRPService',
      'RunMRP',
      { mps_id: toInt(dto.mpsId), operator: dto.operator ?? 'system' },
      URL,
    );
  }

  listScheduleJobs(query: any) {
    return this.grpc.call(
      PKG,
      'ScheduleService',
      'ListSchedules',
      buildListRequest(query, 'pagination'),
      URL,
    );
  }

  createScheduleJob(dto: CreateScheduleJobDto) {
    return this.grpc.call(
      PKG,
      'ScheduleService',
      'CreateSchedule',
      {
        plan: {
          plan_no: dto.planNo,
          mps_id: toInt(dto.mpsId),
          plan_type: toProtoEnum('PLAN_TYPE', dto.planType || 'FINE'),
          algorithm: toProtoEnum('SCHEDULING_ALGORITHM', dto.algorithm || 'FIFO'),
          start_date: toProtoTimestamp(dto.plannedStart),
          end_date: toProtoTimestamp(dto.plannedEnd),
          workcenter_id: toInt(dto.workcenterId),
        },
      },
      URL,
    );
  }
}
