import { Body, Controller, Get, Param, Patch, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { EamService } from './eam.service';
import { CreateEquipmentDto, CreateRepairOrderDto, UpdateRepairOrderDto, StartDowntimeDto, ResolveDowntimeDto, SaveOeeDto } from './dto/eam.dto';

@ApiTags('EAM - 设备管理')
@Controller('api/eam')
export class EamController {
  constructor(private readonly eamService: EamService) {}

  @Get('equipment')
  @ApiOperation({ summary: 'List equipment' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated equipment list' })
  async listEquipment(@Query() query: any) {
    return this.eamService.listEquipment(query);
  }

  @Post('equipment')
  @ApiOperation({ summary: 'Create equipment' })
  @ApiResponse({ status: 201, description: 'Equipment created' })
  async createEquipment(@Body() dto: CreateEquipmentDto) {
    return this.eamService.createEquipment(dto);
  }

  @Get('equipment/:id')
  @ApiOperation({ summary: 'Get equipment detail' })
  @ApiResponse({ status: 200, description: 'Equipment detail' })
  async getEquipment(@Param('id') id: string) {
    return this.eamService.getEquipment(id);
  }

  @Get('repair-orders')
  @ApiOperation({ summary: 'List repair orders' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated repair orders' })
  async listRepairOrders(@Query() query: any) {
    return this.eamService.listRepairOrders(query);
  }

  @Post('repair-orders')
  @ApiOperation({ summary: 'Create a repair order' })
  @ApiResponse({ status: 201, description: 'Repair order created' })
  async createRepairOrder(@Body() dto: CreateRepairOrderDto) {
    return this.eamService.createRepairOrder(dto);
  }

  @Patch('repair-orders/:id')
  @ApiOperation({ summary: 'Update a repair order' })
  @ApiResponse({ status: 200, description: 'Repair order updated' })
  async updateRepairOrder(@Param('id') id: string, @Body() dto: UpdateRepairOrderDto) {
    return this.eamService.updateRepairOrder(id, dto);
  }

  @Post('downtime/start')
  @ApiOperation({ summary: 'Start an equipment downtime event' })
  @ApiResponse({ status: 201, description: 'Downtime started' })
  async startDowntime(@Body() dto: StartDowntimeDto) {
    return this.eamService.startDowntime(dto);
  }

  @Post('downtime/:id/resolve')
  @ApiOperation({ summary: 'Resolve a downtime event' })
  @ApiResponse({ status: 200, description: 'Downtime resolved' })
  async resolveDowntime(@Param('id') id: string, @Body() dto: ResolveDowntimeDto) {
    return this.eamService.resolveDowntime(id, dto);
  }

  @Get('oee')
  @ApiOperation({ summary: 'Get OEE (Overall Equipment Effectiveness) metrics' })
  @ApiQuery({ name: 'equipmentId', required: false, type: String })
  @ApiQuery({ name: 'beginDate', required: false, type: String, description: 'YYYY-MM-DD' })
  @ApiQuery({ name: 'endDate', required: false, type: String, description: 'YYYY-MM-DD' })
  @ApiResponse({ status: 200, description: 'OEE metrics' })
  async getOee(@Query() query: any) {
    return this.eamService.getOee(query);
  }

  @Post('oee')
  @ApiOperation({ summary: 'Report OEE factors (A/P/Q) for one equipment on one day' })
  @ApiResponse({ status: 201, description: 'OEE saved (idempotent per equipment+date)' })
  async saveOee(@Body() dto: SaveOeeDto) {
    return this.eamService.saveOee(dto);
  }
}
