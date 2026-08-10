import { IsString, IsOptional, IsArray, ValidateNested, IsNumber, Min } from 'class-validator';
import { Type } from 'class-transformer';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class CreateMaterialDto {
  @ApiProperty({ description: 'Material code (unique)' })
  @IsString()
  materialCode: string;

  @ApiProperty({ description: 'Material name' })
  @IsString()
  name: string;

  @ApiPropertyOptional({ description: 'Specification' })
  @IsOptional()
  @IsString()
  specification?: string;

  @ApiPropertyOptional({ description: 'Unit of measure' })
  @IsOptional()
  @IsString()
  unit?: string;

  @ApiPropertyOptional({ description: 'Material category' })
  @IsOptional()
  @IsString()
  category?: string;

  @ApiPropertyOptional({ description: 'Material type (raw / semi / finished)' })
  @IsOptional()
  @IsString()
  materialType?: string;
}

export class BomLineDto {
  @ApiProperty({ description: 'Component material code' })
  @IsString()
  componentCode: string;

  @ApiProperty({ description: 'Component quantity per unit' })
  @IsNumber()
  @Min(0)
  quantity: number;

  @ApiPropertyOptional({ description: 'Unit' })
  @IsOptional()
  @IsString()
  unit?: string;
}

export class CreateBomDto {
  @ApiProperty({ description: 'Finished / parent material code' })
  @IsString()
  productCode: string;

  @ApiPropertyOptional({ description: 'BOM version' })
  @IsOptional()
  @IsString()
  version?: string;

  @ApiProperty({ type: () => BomLineDto, description: 'BOM lines', isArray: true })
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => BomLineDto)
  lines: BomLineDto[];
}
