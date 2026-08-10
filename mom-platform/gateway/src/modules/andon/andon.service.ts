import { Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { buildListRequest, toProtoEnum, toInt } from '../../grpc/case.util';
import {
  CreateAndonCallDto,
  AcknowledgeCallDto,
  ResolveCallDto,
  CreateAlertDto,
  CreateAlertConfigDto,
} from './dto/andon.dto';

const PKG = 'mom.andon';
const URL = process.env.ANDON_SERVICE_URL || 'localhost:50058';

/**
 * Andon BFF service — 委托给 Andon gRPC 微服务 (andon-service:50058)。
 */
@Injectable()
export class AndonService {
  constructor(private readonly grpc: GrpcService) {}

  createCall(dto: CreateAndonCallDto) {
    // callPoint 是人读标签（WS-1），proto 需要数字工位 id：优先取 workstationId，
    // 否则从标签里抽取数字，避免整条呼叫落成 workstation_id=0。
    const workstationId = dto.workstationId ?? (dto.callPoint.match(/\d+/)?.[0] ?? '0');
    return this.grpc.call(
      PKG,
      'AndonService',
      'TriggerAndon',
      {
        workstation_id: toInt(workstationId),
        reporter_id: toInt(dto.callerId),
        andon_type: toProtoEnum('ANDON_TYPE', dto.type),
        description: dto.description,
      },
      URL,
    );
  }
  acknowledgeCall(id: string, dto: AcknowledgeCallDto) {
    return this.grpc.call(
      PKG,
      'AndonService',
      'AcknowledgeAndon',
      { id: toInt(id), operator_id: toInt(dto.responderId) },
      URL,
    );
  }
  resolveCall(id: string, dto: ResolveCallDto) {
    return this.grpc.call(
      PKG,
      'AndonService',
      'ResolveAndon',
      { id: toInt(id), operator_id: toInt(dto.resolverId), resolution: dto.resolution ?? '' },
      URL,
    );
  }
  listCalls(query: any) {
    return this.grpc.call(PKG, 'AndonService', 'ListAndonCalls', buildListRequest(query, 'page'), URL);
  }
  listAlerts(query: any) {
    return this.grpc.call(PKG, 'AndonService', 'ListAlerts', buildListRequest(query, 'page'), URL);
  }
  createAlertConfig(dto: CreateAlertConfigDto) {
    return this.grpc.call(
      PKG,
      'AndonService',
      'CreateAlertConfig',
      {
        config_code: dto.configCode,
        config_name: dto.configName,
        trigger_type: toProtoEnum('TRIGGER_TYPE', dto.triggerType ?? 'THRESHOLD'),
        severity: toProtoEnum('ALERT_SEVERITY', dto.severity ?? 'P2'),
        trigger_condition: dto.triggerCondition ?? '',
        notify_channels: dto.notifyChannels ?? '',
      },
      URL,
    );
  }
  createAlert(dto: CreateAlertDto) {
    return this.grpc.call(
      PKG,
      'AndonService',
      'TriggerAlert',
      {
        config_id: toInt(dto.configId),
        target_id: toInt(dto.targetId),
        target_type: dto.targetType ?? '',
      },
      URL,
    );
  }
}
