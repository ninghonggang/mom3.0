import { IsString, IsNumber, IsOptional, IsEnum, IsDateString, Min, IsInt, IsArray, ValidateNested } from 'class-validator';
import { Type } from 'class-transformer';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export enum OrderStatus {
  PENDING = 'PENDING',
  DISPATCHED = 'DISPATCHED',
  IN_PROGRESS = 'IN_PROGRESS',
  COMPLETED = 'COMPLETED',
  SUSPENDED = 'SUSPENDED',
  CLOSED = 'CLOSED',
}

export class OrderItemDto {
  @ApiProperty({ description: 'Material code' })
  @IsString()
  materialCode: string;

  @ApiProperty({ description: 'Planned quantity' })
  @IsNumber()
  @Min(0)
  quantity: number;

  @ApiPropertyOptional({ description: 'Unit of measure' })
  @IsOptional()
  @IsString()
  unit?: string;
}

export class CreateOrderDto {
  @ApiProperty({ description: 'Production order number (unique)' })
  @IsString()
  orderNo: string;

  @ApiProperty({ description: 'Product / material code' })
  @IsString()
  productCode: string;

  @ApiProperty({ description: 'Planned production quantity' })
  @IsNumber()
  @Min(0)
  plannedQuantity: number;

  @ApiPropertyOptional({ description: 'Workshop ID' })
  @IsOptional()
  @IsString()
  workshopId?: string;

  @ApiPropertyOptional({ description: 'Production line ID' })
  @IsOptional()
  @IsString()
  lineId?: string;

  @ApiPropertyOptional({ description: 'Workstation ID' })
  @IsOptional()
  @IsString()
  workstationId?: string;

  @ApiPropertyOptional({ description: 'Planned start time (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  plannedStartTime?: string;

  @ApiPropertyOptional({ description: 'Planned end time (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  plannedEndTime?: string;

  @ApiPropertyOptional({ description: 'BOM ID', type: [OrderItemDto] })
  @IsOptional()
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => OrderItemDto)
  items?: OrderItemDto[];
}

export class UpdateOrderStatusDto {
  @ApiProperty({ enum: OrderStatus, description: 'Target order status' })
  @IsEnum(OrderStatus)
  status: OrderStatus;

  @ApiPropertyOptional({ description: 'Reason for status change' })
  @IsOptional()
  @IsString()
  reason?: string;
}

export class DispatchDto {
  @ApiPropertyOptional({ description: 'Workshop ID to dispatch to' })
  @IsOptional()
  @IsString()
  workshopId?: string;

  @ApiPropertyOptional({ description: 'Production line ID' })
  @IsOptional()
  @IsString()
  lineId?: string;

  @ApiPropertyOptional({ description: 'Workstation ID' })
  @IsOptional()
  @IsString()
  workstationId?: string;

  @ApiPropertyOptional({ description: 'Operator / crew ID' })
  @IsOptional()
  @IsString()
  operatorId?: string;

  @ApiPropertyOptional({ description: 'Dispatch notes' })
  @IsOptional()
  @IsString()
  remark?: string;
}

export class ReportDto {
  @ApiProperty({ description: 'Reported (good) quantity' })
  @IsNumber()
  @Min(0)
  goodQuantity: number;

  @ApiPropertyOptional({ description: 'Defective quantity' })
  @IsOptional()
  @IsNumber()
  @Min(0)
  scrapQuantity?: number;

  @ApiPropertyOptional({ description: 'Operator ID' })
  @IsOptional()
  @IsString()
  operatorId?: string;

  @ApiPropertyOptional({ description: 'Workstation ID' })
  @IsOptional()
  @IsString()
  workstationId?: string;

  @ApiPropertyOptional({ description: 'Shift' })
  @IsOptional()
  @IsString()
  shift?: string;

  @ApiPropertyOptional({ description: 'Remark' })
  @IsOptional()
  @IsString()
  remark?: string;
}

export class CompleteDto {
  @ApiProperty({ description: 'Completed (good) quantity to receive into stock' })
  @IsNumber()
  @Min(0)
  actualQuantity: number;

  @ApiPropertyOptional({ description: 'Target warehouse ID for finished goods' })
  @IsOptional()
  @IsString()
  warehouseId?: string;

  @ApiPropertyOptional({ description: 'Target location code/ID for finished goods' })
  @IsOptional()
  @IsString()
  locationCode?: string;

  @ApiPropertyOptional({ description: 'Finished-goods batch number' })
  @IsOptional()
  @IsString()
  batchNo?: string;

  @ApiPropertyOptional({ description: 'Actual completion time (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  actualEndTime?: string;

  @ApiPropertyOptional({ description: 'Completion remark' })
  @IsOptional()
  @IsString()
  remark?: string;
}
