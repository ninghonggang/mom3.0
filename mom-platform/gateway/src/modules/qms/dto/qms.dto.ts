import { IsString, IsNumber, IsOptional, IsEnum, IsDateString, Min, IsArray, ValidateNested } from 'class-validator';
import { Type } from 'class-transformer';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export enum InspectionResult {
  PASS = 'PASS',
  FAIL = 'FAIL',
  CONDITIONAL = 'CONDITIONAL',
}

export class CreateInspectionSheetDto {
  @ApiProperty({ description: 'Inspection sheet number' })
  @IsString()
  sheetNo: string;

  @ApiProperty({ description: 'Inspection type (first-article / process / final)' })
  @IsString()
  inspectionType: string;

  @ApiPropertyOptional({ description: 'Related production order ID' })
  @IsOptional()
  @IsString()
  orderId?: string;

  @ApiPropertyOptional({ description: 'Material / product code' })
  @IsOptional()
  @IsString()
  materialCode?: string;

  @ApiPropertyOptional({ description: 'Sample size (defaults to 1)' })
  @IsOptional()
  @IsNumber()
  @Min(1)
  sampleSize?: number;

  @ApiPropertyOptional({ description: 'Inspector ID' })
  @IsOptional()
  @IsString()
  inspectorId?: string;
}

export class UpdateInspectionSheetDto {
  @ApiPropertyOptional({
    description: 'Sheet status: PENDING / IN_PROGRESS / PASSED / FAILED / WAIVED / CANCELLED',
  })
  @IsOptional()
  @IsString()
  status?: string;

  @ApiPropertyOptional({ description: 'Defect count' })
  @IsOptional()
  @IsNumber()
  @Min(0)
  defectCount?: number;

  @ApiPropertyOptional({ description: 'Inspector ID' })
  @IsOptional()
  @IsString()
  inspectorId?: string;
}

export enum CharDataType {
  NUMBER = 'NUMBER',
  TEXT = 'TEXT',
  PASS_FAIL = 'PASS_FAIL',
  BOOLEAN = 'BOOLEAN',
}

/** 检验特性（质量特性主数据）：检验结果必须挂在已存在的特性上。 */
export class CreateCharacteristicDto {
  @ApiProperty({ description: 'Characteristic code (unique)' })
  @IsString()
  charCode: string;

  @ApiProperty({ description: 'Characteristic name' })
  @IsString()
  charName: string;

  @ApiPropertyOptional({ enum: CharDataType, description: 'Data type (defaults to NUMBER)' })
  @IsOptional()
  @IsEnum(CharDataType)
  dataType?: CharDataType;

  @ApiPropertyOptional({ description: 'Upper spec limit (USL)' })
  @IsOptional()
  @IsString()
  usl?: string;

  @ApiPropertyOptional({ description: 'Lower spec limit (LSL)' })
  @IsOptional()
  @IsString()
  lsl?: string;

  @ApiPropertyOptional({ description: 'Target value' })
  @IsOptional()
  @IsString()
  target?: string;

  @ApiPropertyOptional({ description: 'Unit of measure' })
  @IsOptional()
  @IsString()
  unit?: string;
}

export class InspectionResultItemDto {
  @ApiProperty({ description: 'Inspection item / parameter code' })
  @IsString()
  itemCode: string;

  @ApiProperty({ enum: InspectionResult, description: 'Inspection result' })
  @IsEnum(InspectionResult)
  result: InspectionResult;

  @ApiPropertyOptional({ description: 'Measured value' })
  @IsOptional()
  @IsString()
  measuredValue?: string;

  @ApiPropertyOptional({ description: 'Remark' })
  @IsOptional()
  @IsString()
  remark?: string;
}

export class SubmitResultsDto {
  @ApiProperty({ type: [InspectionResultItemDto], description: 'Inspection result items' })
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => InspectionResultItemDto)
  results: InspectionResultItemDto[];

  @ApiPropertyOptional({ description: 'Inspector ID' })
  @IsOptional()
  @IsString()
  inspectorId?: string;
}

export class CreateNcrDto {
  @ApiProperty({ description: 'NCR (non-conformance report) number' })
  @IsString()
  ncrNo: string;

  @ApiPropertyOptional({ description: 'Related inspection sheet ID' })
  @IsOptional()
  @IsString()
  sheetId?: string;

  @ApiProperty({ description: 'Non-conformance description' })
  @IsString()
  description: string;

  @ApiPropertyOptional({ description: 'Material / product code' })
  @IsOptional()
  @IsString()
  materialCode?: string;

  @ApiPropertyOptional({ description: 'Non-conforming quantity (defaults to 1)' })
  @IsOptional()
  @IsNumber()
  @Min(0)
  quantity?: number;

  @ApiPropertyOptional({ description: 'Defect category' })
  @IsOptional()
  @IsString()
  defectCategory?: string;

  @ApiPropertyOptional({ description: 'Severity: CRITICAL / MAJOR / MINOR' })
  @IsOptional()
  @IsString()
  severity?: string;
}

export class UpdateNcrDto {
  @ApiPropertyOptional({ description: 'NCR status' })
  @IsOptional()
  @IsString()
  status?: string;

  @ApiPropertyOptional({ description: 'Updated description' })
  @IsOptional()
  @IsString()
  description?: string;
}

export class NcrActionDto {
  @ApiProperty({ description: 'Action type (rework / scrap / concession / re-inspect)' })
  @IsString()
  actionType: string;

  @ApiPropertyOptional({ description: 'Action description' })
  @IsOptional()
  @IsString()
  description?: string;

  @ApiPropertyOptional({ description: 'Responsible person ID' })
  @IsOptional()
  @IsString()
  handlerId?: string;
}
