import { IsString, IsOptional, IsEnum } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export enum AndonCallType {
  QUALITY = 'QUALITY',
  EQUIPMENT = 'EQUIPMENT',
  MATERIAL = 'MATERIAL',
  SAFETY = 'SAFETY',
}

export class CreateAndonCallDto {
  @ApiProperty({ description: 'Call point label, e.g. WS-1' })
  @IsString()
  callPoint: string;

  @ApiPropertyOptional({ description: 'Workstation ID (numeric). Falls back to digits in callPoint.' })
  @IsOptional()
  @IsString()
  workstationId?: string;

  @ApiProperty({ enum: AndonCallType, description: 'Andon call type' })
  @IsEnum(AndonCallType)
  type: AndonCallType;

  @ApiProperty({ description: 'Description of the issue' })
  @IsString()
  description: string;

  @ApiPropertyOptional({ description: 'Caller ID' })
  @IsOptional()
  @IsString()
  callerId?: string;

  @ApiPropertyOptional({ description: 'Priority' })
  @IsOptional()
  @IsString()
  priority?: string;
}

export class AcknowledgeCallDto {
  @ApiPropertyOptional({ description: 'Responder ID' })
  @IsOptional()
  @IsString()
  responderId?: string;

  @ApiPropertyOptional({ description: 'Acknowledgement note' })
  @IsOptional()
  @IsString()
  remark?: string;
}

export class ResolveCallDto {
  @ApiPropertyOptional({ description: 'Resolution description' })
  @IsOptional()
  @IsString()
  resolution?: string;

  @ApiPropertyOptional({ description: 'Resolver ID' })
  @IsOptional()
  @IsString()
  resolverId?: string;
}

export enum AlertTriggerType {
  THRESHOLD = 'THRESHOLD',
  EVENT = 'EVENT',
  SCHEDULE = 'SCHEDULE',
}

export enum AlertSeverity {
  P0 = 'P0',
  P1 = 'P1',
  P2 = 'P2',
  P3 = 'P3',
}

/** 告警规则配置 — 告警必须先有配置才能触发 */
export class CreateAlertConfigDto {
  @ApiProperty({ description: 'Unique alert config code, e.g. TEMP_HIGH' })
  @IsString()
  configCode: string;

  @ApiProperty({ description: 'Human readable config name' })
  @IsString()
  configName: string;

  @ApiPropertyOptional({ enum: AlertTriggerType, description: 'Trigger type' })
  @IsOptional()
  @IsEnum(AlertTriggerType)
  triggerType?: AlertTriggerType;

  @ApiPropertyOptional({ enum: AlertSeverity, description: 'Severity level' })
  @IsOptional()
  @IsEnum(AlertSeverity)
  severity?: AlertSeverity;

  @ApiPropertyOptional({ description: 'Trigger condition expression, e.g. value > 90' })
  @IsOptional()
  @IsString()
  triggerCondition?: string;

  @ApiPropertyOptional({ description: 'Notify channels, comma separated (sms,email,wecom)' })
  @IsOptional()
  @IsString()
  notifyChannels?: string;
}

/** 触发一条告警 — 引用已存在的告警配置 */
export class CreateAlertDto {
  @ApiProperty({ description: 'Alert config ID (from POST /andon/alert-configs)' })
  @IsString()
  configId: string;

  @ApiPropertyOptional({ description: 'Target entity ID (equipment / workstation ...)' })
  @IsOptional()
  @IsString()
  targetId?: string;

  @ApiPropertyOptional({ description: 'Target entity type, e.g. EQUIPMENT / WORKSTATION' })
  @IsOptional()
  @IsString()
  targetType?: string;
}
