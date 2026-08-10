import { IsString, IsNumber, IsOptional, IsEnum, IsDateString, Min, Max } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export enum EquipmentStatus {
  RUNNING = 'RUNNING',
  IDLE = 'IDLE',
  DOWN = 'DOWN',
  MAINTENANCE = 'MAINTENANCE',
}

export class CreateEquipmentDto {
  @ApiProperty({ description: 'Equipment code (unique)' })
  @IsString()
  equipmentCode: string;

  @ApiProperty({ description: 'Equipment name' })
  @IsString()
  name: string;

  @ApiPropertyOptional({ description: 'Equipment category / type' })
  @IsOptional()
  @IsString()
  category?: string;

  @ApiPropertyOptional({ description: 'Workshop ID' })
  @IsOptional()
  @IsString()
  workshopId?: string;

  @ApiPropertyOptional({ description: 'Production line ID' })
  @IsOptional()
  @IsString()
  lineId?: string;

  @ApiPropertyOptional({ description: 'Equipment model' })
  @IsOptional()
  @IsString()
  model?: string;

  @ApiPropertyOptional({ description: 'Specification' })
  @IsOptional()
  @IsString()
  specification?: string;

  @ApiPropertyOptional({ description: 'Target OEE, e.g. "0.85"' })
  @IsOptional()
  @IsString()
  targetOee?: string;
}

export class CreateRepairOrderDto {
  @ApiProperty({ description: 'Equipment ID' })
  @IsString()
  equipmentId: string;

  @ApiProperty({ description: 'Fault description' })
  @IsString()
  faultDescription: string;

  @ApiPropertyOptional({ description: 'Fault category' })
  @IsOptional()
  @IsString()
  faultCategory?: string;

  @ApiPropertyOptional({ description: 'Reporter ID' })
  @IsOptional()
  @IsString()
  reporterId?: string;

  @ApiPropertyOptional({ description: 'Priority' })
  @IsOptional()
  @IsString()
  priority?: string;
}

export class UpdateRepairOrderDto {
  @ApiPropertyOptional({ description: 'Repair order status' })
  @IsOptional()
  @IsString()
  status?: string;

  @ApiPropertyOptional({ description: 'Repair description / root cause' })
  @IsOptional()
  @IsString()
  repairDescription?: string;

  @ApiPropertyOptional({ description: 'Repair technician ID' })
  @IsOptional()
  @IsString()
  technicianId?: string;
}

export class StartDowntimeDto {
  @ApiProperty({ description: 'Equipment ID' })
  @IsString()
  equipmentId: string;

  @ApiProperty({ description: 'Downtime reason / category' })
  @IsString()
  reason: string;

  @ApiPropertyOptional({ description: 'Downtime type (planned / unplanned)' })
  @IsOptional()
  @IsString()
  type?: string;

  @ApiPropertyOptional({ description: 'Reporter ID' })
  @IsOptional()
  @IsString()
  reporterId?: string;
}

export class ResolveDowntimeDto {
  @ApiPropertyOptional({ description: 'Resolution description' })
  @IsOptional()
  @IsString()
  resolution?: string;

  @ApiPropertyOptional({ description: 'Resolver ID' })
  @IsOptional()
  @IsString()
  resolverId?: string;
}

/**
 * 上报某设备某日的 OEE 三要素。服务端按 OEE = A × P × Q 计算并落库，
 * 同一 (设备, 日期) 重复上报为幂等覆盖，供看板 oeeAvg 聚合使用。
 */
export class SaveOeeDto {
  @ApiProperty({ description: 'Equipment ID' })
  @IsString()
  equipmentId: string;

  @ApiPropertyOptional({ description: 'Calculation date (YYYY-MM-DD), defaults to today' })
  @IsOptional()
  @IsDateString()
  calcDate?: string;

  @ApiProperty({ description: 'Availability ratio, 0~1', example: 0.92 })
  @IsNumber()
  @Min(0)
  @Max(1)
  availability: number;

  @ApiProperty({ description: 'Performance ratio, 0~1', example: 0.88 })
  @IsNumber()
  @Min(0)
  @Max(1)
  performance: number;

  @ApiProperty({ description: 'Quality ratio, 0~1', example: 0.97 })
  @IsNumber()
  @Min(0)
  @Max(1)
  quality: number;
}
