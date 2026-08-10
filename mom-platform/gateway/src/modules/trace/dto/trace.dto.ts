import { IsString, IsNumber, IsOptional, IsEnum, Min } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class CreateTraceRecordDto {
  @ApiProperty({ description: 'Serial number / unit identifier' })
  @IsString()
  serialNo: string;

  @ApiProperty({ description: 'Material / product code' })
  @IsString()
  productCode: string;

  @ApiPropertyOptional({ description: 'Production order ID' })
  @IsOptional()
  @IsString()
  orderId?: string;

  @ApiPropertyOptional({ description: 'Workstation ID' })
  @IsOptional()
  @IsString()
  workstationId?: string;

  @ApiPropertyOptional({ description: 'Batch / lot number' })
  @IsOptional()
  @IsString()
  batchNo?: string;
}

export class GenerateSerialDto {
  @ApiProperty({ description: 'Product / material code' })
  @IsString()
  productCode: string;

  @ApiProperty({ description: 'Number of serials to generate' })
  @IsNumber()
  @Min(1)
  count: number;

  @ApiPropertyOptional({ description: 'Prefix for serial numbers' })
  @IsOptional()
  @IsString()
  prefix?: string;
}

export enum TraceDataType {
  NUMBER = 'NUMBER',
  STRING = 'STRING',
  BOOLEAN = 'BOOLEAN',
}

export enum DataQuality {
  GOOD = 'GOOD',
  BAD = 'BAD',
  UNCERTAIN = 'UNCERTAIN',
}

/**
 * 数据点「定义」— 描述一个可采集的测点（不是采集值本身）。
 * 采集值请使用 POST /trace/collect。
 */
export class DataPointDto {
  @ApiProperty({ description: 'Unique data point code, e.g. TEMP_01' })
  @IsString()
  pointCode: string;

  @ApiProperty({ description: 'Data point display name' })
  @IsString()
  pointName: string;

  @ApiPropertyOptional({ description: 'Bound equipment ID' })
  @IsOptional()
  @IsString()
  equipmentId?: string;

  @ApiPropertyOptional({ enum: TraceDataType, description: 'Value data type' })
  @IsOptional()
  @IsEnum(TraceDataType)
  dataType?: TraceDataType;

  @ApiPropertyOptional({ description: 'Upper control limit' })
  @IsOptional()
  @IsString()
  upperLimit?: string;

  @ApiPropertyOptional({ description: 'Lower control limit' })
  @IsOptional()
  @IsString()
  lowerLimit?: string;

  @ApiPropertyOptional({ description: 'Collect interval in seconds' })
  @IsOptional()
  @IsNumber()
  @Min(1)
  collectIntervalSeconds?: number;
}

/** 采集一条测点数据 — dataPointId 来自 POST /trace/data-points 的返回 */
export class CollectDto {
  @ApiProperty({ description: 'Data point ID (from POST /trace/data-points)' })
  @IsString()
  dataPointId: string;

  @ApiProperty({ description: 'Collected value (as string, keeps decimal precision)' })
  @IsString()
  value: string;

  @ApiPropertyOptional({ enum: DataQuality, description: 'Data quality flag' })
  @IsOptional()
  @IsEnum(DataQuality)
  quality?: DataQuality;
}

export class CreateScanLogDto {
  @ApiProperty({ description: 'Serial number scanned' })
  @IsString()
  serialNo: string;

  @ApiProperty({ description: 'Scan point / workstation ID' })
  @IsString()
  scanPoint: string;

  @ApiPropertyOptional({ description: 'Operator ID' })
  @IsOptional()
  @IsString()
  operatorId?: string;

  @ApiPropertyOptional({ description: 'Scan result (OK / NG)' })
  @IsOptional()
  @IsString()
  result?: string;
}
