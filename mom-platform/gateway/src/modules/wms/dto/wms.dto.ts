import { IsString, IsNumber, IsOptional, IsDateString, IsEnum, Min, IsArray, ValidateNested } from 'class-validator';
import { Type } from 'class-transformer';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class CreateWarehouseDto {
  @ApiProperty({ description: 'Warehouse code (unique)' })
  @IsString()
  warehouseCode: string;

  @ApiProperty({ description: 'Warehouse name' })
  @IsString()
  name: string;

  @ApiPropertyOptional({ description: 'Warehouse type' })
  @IsOptional()
  @IsString()
  type?: string;

  @ApiPropertyOptional({ description: 'Address / location' })
  @IsOptional()
  @IsString()
  address?: string;
}

export enum LocationType {
  PICK = 'PICK',
  STORAGE = 'STORAGE',
  INBOUND = 'INBOUND',
  OUTBOUND = 'OUTBOUND',
}

export class CreateLocationDto {
  @ApiProperty({ description: 'Warehouse code or ID this location belongs to' })
  @IsString()
  warehouseId: string;

  @ApiProperty({ description: 'Location / bin code (unique), e.g. A-01-01' })
  @IsString()
  locationCode: string;

  @ApiPropertyOptional({ enum: LocationType, description: 'Location type' })
  @IsOptional()
  @IsEnum(LocationType)
  locationType?: LocationType;

  @ApiPropertyOptional({ description: 'Storage area ID' })
  @IsOptional()
  @IsString()
  areaId?: string;

  @ApiPropertyOptional({ description: 'Capacity of the bin' })
  @IsOptional()
  @IsNumber()
  @Min(0)
  capacity?: number;
}

export class CreateReceiveOrderDto {
  @ApiProperty({ description: 'Receive order number' })
  @IsString()
  receiveNo: string;

  @ApiProperty({ description: 'Warehouse ID' })
  @IsString()
  warehouseId: string;

  @ApiPropertyOptional({ description: 'Supplier ID' })
  @IsOptional()
  @IsString()
  supplierId?: string;

  @ApiPropertyOptional({ description: 'Related purchase order number' })
  @IsOptional()
  @IsString()
  poNo?: string;

  @ApiProperty({ type: () => ReceiveLineDto, description: 'Receive lines', isArray: true })
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => ReceiveLineDto)
  lines: ReceiveLineDto[];
}

export class ReceiveLineDto {
  @ApiProperty({ description: 'Material code' })
  @IsString()
  materialCode: string;

  @ApiProperty({ description: 'Quantity to receive' })
  @IsNumber()
  @Min(0)
  quantity: number;

  @ApiPropertyOptional({ description: 'Unit' })
  @IsOptional()
  @IsString()
  unit?: string;

  @ApiPropertyOptional({
    description:
      'Purchase unit price. Feeds the moving weighted average that maintains inventory unit cost, which drives inventory valuation.',
  })
  @IsOptional()
  @IsNumber()
  @Min(0)
  unitPrice?: number;

  @ApiPropertyOptional({ description: 'Batch / lot number' })
  @IsOptional()
  @IsString()
  batchNo?: string;
}

export class ConfirmReceiveDto {
  @ApiPropertyOptional({ description: 'Actual received quantity per line (map)' })
  @IsOptional()
  actualQuantities?: Record<string, number>;

  @ApiPropertyOptional({ description: 'Remark' })
  @IsOptional()
  @IsString()
  remark?: string;
}

export class PutawayDto {
  @ApiPropertyOptional({ description: 'Target location code' })
  @IsOptional()
  @IsString()
  locationCode?: string;

  @ApiPropertyOptional({ description: 'Operator ID' })
  @IsOptional()
  @IsString()
  operatorId?: string;
}

export class CreateDeliveryOrderDto {
  @ApiProperty({ description: 'Delivery order number' })
  @IsString()
  deliveryNo: string;

  @ApiProperty({ description: 'Warehouse ID' })
  @IsString()
  warehouseId: string;

  @ApiPropertyOptional({ description: 'Customer ID' })
  @IsOptional()
  @IsString()
  customerId?: string;

  @ApiPropertyOptional({ description: 'Related sales order number' })
  @IsOptional()
  @IsString()
  soNo?: string;

  @ApiProperty({ type: () => DeliveryLineDto, description: 'Delivery lines', isArray: true })
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => DeliveryLineDto)
  lines: DeliveryLineDto[];
}

export class DeliveryLineDto {
  @ApiProperty({ description: 'Material code' })
  @IsString()
  materialCode: string;

  @ApiProperty({ description: 'Quantity to deliver' })
  @IsNumber()
  @Min(0)
  quantity: number;

  @ApiPropertyOptional({ description: 'Unit' })
  @IsOptional()
  @IsString()
  unit?: string;

  @ApiPropertyOptional({ description: 'Batch / lot number' })
  @IsOptional()
  @IsString()
  batchNo?: string;
}

export class PickDto {
  @ApiPropertyOptional({ description: 'Picked location codes' })
  @IsOptional()
  locationCodes?: string[];

  @ApiPropertyOptional({ description: 'Operator ID' })
  @IsOptional()
  @IsString()
  operatorId?: string;
}

export class ShipDto {
  @ApiPropertyOptional({ description: 'Carrier' })
  @IsOptional()
  @IsString()
  carrier?: string;

  @ApiPropertyOptional({ description: 'Tracking number' })
  @IsOptional()
  @IsString()
  trackingNo?: string;

  @ApiPropertyOptional({ description: 'Operator ID' })
  @IsOptional()
  @IsString()
  operatorId?: string;
}

export class CreateCountPlanDto {
  @ApiProperty({ description: 'Count plan number' })
  @IsString()
  planNo: string;

  @ApiProperty({ description: 'Warehouse ID' })
  @IsString()
  warehouseId: string;

  @ApiPropertyOptional({ description: 'Count type (full / cycle)' })
  @IsOptional()
  @IsString()
  countType?: string;

  @ApiPropertyOptional({ description: 'Planned count date (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  plannedDate?: string;
}

export class CreateCountRecordDto {
  @ApiProperty({ description: 'Count plan ID' })
  @IsString()
  planId: string;

  @ApiProperty({ description: 'Location code' })
  @IsString()
  locationCode: string;

  @ApiProperty({ description: 'Material code' })
  @IsString()
  materialCode: string;

  @ApiProperty({ description: 'Counted quantity' })
  @IsNumber()
  @Min(0)
  countedQuantity: number;

  @ApiPropertyOptional({ description: 'Operator ID' })
  @IsOptional()
  @IsString()
  operatorId?: string;
}
