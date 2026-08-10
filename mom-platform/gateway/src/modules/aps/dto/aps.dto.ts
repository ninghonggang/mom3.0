import { IsString, IsNumber, IsOptional, IsDateString, Min } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class CreateMpsPlanDto {
  @ApiProperty({ description: 'MPS plan number' })
  @IsString()
  planNo: string;

  @ApiProperty({ description: 'Product / material code' })
  @IsString()
  productCode: string;

  @ApiProperty({ description: 'Planned quantity' })
  @IsNumber()
  @Min(0)
  plannedQuantity: number;

  @ApiPropertyOptional({ description: 'Plan month (YYYY-MM); derived from startDate when omitted' })
  @IsOptional()
  @IsString()
  planMonth?: string;

  @ApiPropertyOptional({ description: 'Planned start date (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  startDate?: string;

  @ApiPropertyOptional({ description: 'Planned end date (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  endDate?: string;

  @ApiPropertyOptional({ description: 'Remark' })
  @IsOptional()
  @IsString()
  remark?: string;
}

export class GenerateMrpDto {
  @ApiProperty({ description: 'MPS plan ID to explode requirements from' })
  @IsString()
  mpsId: string;

  @ApiPropertyOptional({ description: 'Operator name recorded on the MRP run' })
  @IsOptional()
  @IsString()
  operator?: string;
}

export class CreateScheduleJobDto {
  @ApiProperty({ description: 'Schedule plan number' })
  @IsString()
  planNo: string;

  @ApiPropertyOptional({ description: 'Source MPS plan ID' })
  @IsOptional()
  @IsString()
  mpsId?: string;

  @ApiPropertyOptional({ description: 'Plan type: ROUGH / FINE' })
  @IsOptional()
  @IsString()
  planType?: string;

  @ApiPropertyOptional({ description: 'Scheduling algorithm: FIFO / EDD / SPT / LPT' })
  @IsOptional()
  @IsString()
  algorithm?: string;

  @ApiPropertyOptional({ description: 'Work center ID' })
  @IsOptional()
  @IsString()
  workcenterId?: string;

  @ApiPropertyOptional({ description: 'Planned start (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  plannedStart?: string;

  @ApiPropertyOptional({ description: 'Planned end (ISO-8601)' })
  @IsOptional()
  @IsDateString()
  plannedEnd?: string;
}
