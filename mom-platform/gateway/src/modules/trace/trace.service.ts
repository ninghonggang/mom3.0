import { Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { buildListRequest, toProtoEnum, toInt } from '../../grpc/case.util';
import {
  CreateTraceRecordDto,
  GenerateSerialDto,
  DataPointDto,
  CollectDto,
  CreateScanLogDto,
} from './dto/trace.dto';

const PKG = 'mom.trace';
const URL = process.env.TRACE_SERVICE_URL || 'localhost:50057';

/**
 * Trace BFF service — 委托给 Trace gRPC 微服务 (trace-service:50057)。
 */
@Injectable()
export class TraceService {
  constructor(private readonly grpc: GrpcService) {}

  createTraceRecord(dto: CreateTraceRecordDto) {
    return this.grpc.call(PKG, 'TraceService', 'CreateTraceRecord', dto, URL);
  }
  forwardTrace(query: any) {
    return this.grpc.call(PKG, 'TraceService', 'ForwardTrace', query ?? {}, URL);
  }
  backwardTrace(query: any) {
    return this.grpc.call(PKG, 'TraceService', 'BackwardTrace', query ?? {}, URL);
  }
  generateSerials(dto: GenerateSerialDto) {
    return this.grpc.call(PKG, 'TraceService', 'GenerateSerials', dto, URL);
  }
  createDataPoint(dto: DataPointDto) {
    return this.grpc.call(
      PKG,
      'TraceService',
      'CreateDataPoint',
      {
        point_code: dto.pointCode,
        point_name: dto.pointName,
        equipment_id: toInt(dto.equipmentId),
        data_type: toProtoEnum('DATA_TYPE', dto.dataType ?? 'NUMBER'),
        upper_limit: dto.upperLimit ?? '',
        lower_limit: dto.lowerLimit ?? '',
        collect_interval_seconds: dto.collectIntervalSeconds ?? 60,
      },
      URL,
    );
  }
  collect(dto: CollectDto) {
    return this.grpc.call(
      PKG,
      'TraceService',
      'CollectData',
      {
        data_point_id: toInt(dto.dataPointId),
        value: dto.value,
        quality: toProtoEnum('DATA_QUALITY', dto.quality ?? 'GOOD'),
      },
      URL,
    );
  }
  listDataPoints(query: any) {
    const { equipmentId, status, ...rest } = query ?? {};
    return this.grpc.call(
      PKG,
      'TraceService',
      'ListDataPoints',
      {
        ...buildListRequest(rest, 'page'),
        equipment_id: toInt(equipmentId),
        status: toProtoEnum('DATA_POINT_STATUS', status),
      },
      URL,
    );
  }
  createScanLog(dto: CreateScanLogDto) {
    return this.grpc.call(PKG, 'TraceService', 'CreateScanLog', dto, URL);
  }
}
