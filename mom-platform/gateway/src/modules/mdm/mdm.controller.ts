import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { MdmService } from './mdm.service';
import { CreateMaterialDto, CreateBomDto } from './dto/mdm.dto';

@ApiTags('MDM - 主数据管理')
@Controller('api/mdm')
export class MdmController {
  constructor(private readonly mdmService: MdmService) {}

  @Get('materials')
  @ApiOperation({ summary: 'List materials' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiQuery({ name: 'category', required: false, type: String })
  @ApiResponse({ status: 200, description: 'Paginated materials' })
  async listMaterials(@Query() query: any) {
    return this.mdmService.listMaterials(query);
  }

  @Post('materials')
  @ApiOperation({ summary: 'Create a material' })
  @ApiResponse({ status: 201, description: 'Material created' })
  async createMaterial(@Body() dto: CreateMaterialDto) {
    return this.mdmService.createMaterial(dto);
  }

  @Get('boms')
  @ApiOperation({ summary: 'List BOMs' })
  @ApiQuery({ name: 'productCode', required: false, type: String })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated BOMs' })
  async listBoms(@Query() query: any) {
    return this.mdmService.listBoms(query);
  }

  @Post('boms')
  @ApiOperation({ summary: 'Create a BOM' })
  @ApiResponse({ status: 201, description: 'BOM created' })
  async createBom(@Body() dto: CreateBomDto) {
    return this.mdmService.createBom(dto);
  }

  @Get('workshops')
  @ApiOperation({ summary: 'List workshops' })
  @ApiResponse({ status: 200, description: 'Workshop list' })
  async listWorkshops(@Query() query: any) {
    return this.mdmService.listWorkshops(query);
  }

  @Get('production-lines')
  @ApiOperation({ summary: 'List production lines' })
  @ApiQuery({ name: 'workshopId', required: false, type: String })
  @ApiResponse({ status: 200, description: 'Production line list' })
  async listProductionLines(@Query() query: any) {
    return this.mdmService.listProductionLines(query);
  }

  @Get('workstations')
  @ApiOperation({ summary: 'List workstations' })
  @ApiQuery({ name: 'lineId', required: false, type: String })
  @ApiResponse({ status: 200, description: 'Workstation list' })
  async listWorkstations(@Query() query: any) {
    return this.mdmService.listWorkstations(query);
  }

  @Get('customers')
  @ApiOperation({ summary: 'List customers' })
  @ApiResponse({ status: 200, description: 'Customer list' })
  async listCustomers(@Query() query: any) {
    return this.mdmService.listCustomers(query);
  }

  @Get('suppliers')
  @ApiOperation({ summary: 'List suppliers' })
  @ApiResponse({ status: 200, description: 'Supplier list' })
  async listSuppliers(@Query() query: any) {
    return this.mdmService.listSuppliers(query);
  }
}
