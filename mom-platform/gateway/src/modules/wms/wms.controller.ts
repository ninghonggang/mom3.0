import { Body, Controller, Get, Param, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { WmsService } from './wms.service';
import {
  CreateWarehouseDto,
  CreateLocationDto,
  CreateReceiveOrderDto,
  ConfirmReceiveDto,
  PutawayDto,
  CreateDeliveryOrderDto,
  PickDto,
  ShipDto,
  CreateCountPlanDto,
  CreateCountRecordDto,
} from './dto/wms.dto';

@ApiTags('WMS - 仓储管理')
@Controller('api/wms')
export class WmsController {
  constructor(private readonly wmsService: WmsService) {}

  @Get('warehouses')
  @ApiOperation({ summary: 'List warehouses' })
  @ApiResponse({ status: 200, description: 'Warehouse list' })
  async listWarehouses(@Query() query: any) {
    return this.wmsService.listWarehouses(query);
  }

  @Post('warehouses')
  @ApiOperation({ summary: 'Create a warehouse' })
  @ApiResponse({ status: 201, description: 'Warehouse created' })
  async createWarehouse(@Body() dto: CreateWarehouseDto) {
    return this.wmsService.createWarehouse(dto);
  }

  @Get('locations')
  @ApiOperation({ summary: 'List storage locations' })
  @ApiQuery({ name: 'warehouseId', required: false, type: String })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated locations' })
  async listLocations(@Query() query: any) {
    return this.wmsService.listLocations(query);
  }

  @Post('locations')
  @ApiOperation({ summary: 'Create a storage location (bin)' })
  @ApiResponse({ status: 201, description: 'Location created' })
  async createLocation(@Body() dto: CreateLocationDto) {
    return this.wmsService.createLocation(dto);
  }

  @Get('balances')
  @ApiOperation({ summary: 'List inventory balances' })
  @ApiQuery({ name: 'warehouseId', required: false, type: String })
  @ApiQuery({ name: 'materialCode', required: false, type: String })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated inventory balances' })
  async listBalances(@Query() query: any) {
    return this.wmsService.listBalances(query);
  }

  @Post('receive-orders')
  @ApiOperation({ summary: 'Create a receive order (inbound)' })
  @ApiResponse({ status: 201, description: 'Receive order created' })
  async createReceiveOrder(@Body() dto: CreateReceiveOrderDto) {
    return this.wmsService.createReceiveOrder(dto);
  }

  @Post('receive-orders/:id/confirm')
  @ApiOperation({ summary: 'Confirm a receive order (actual quantities)' })
  @ApiResponse({ status: 200, description: 'Receive order confirmed' })
  async confirmReceive(@Param('id') id: string, @Body() dto: ConfirmReceiveDto) {
    return this.wmsService.confirmReceive(id, dto);
  }

  @Post('receive-orders/:id/putaway')
  @ApiOperation({ summary: 'Put away received goods to a storage location' })
  @ApiResponse({ status: 200, description: 'Putaway completed' })
  async putaway(@Param('id') id: string, @Body() dto: PutawayDto) {
    return this.wmsService.putaway(id, dto);
  }

  @Post('delivery-orders')
  @ApiOperation({ summary: 'Create a delivery order (outbound)' })
  @ApiResponse({ status: 201, description: 'Delivery order created' })
  async createDeliveryOrder(@Body() dto: CreateDeliveryOrderDto) {
    return this.wmsService.createDeliveryOrder(dto);
  }

  @Post('delivery-orders/:id/pick')
  @ApiOperation({ summary: 'Pick goods for a delivery order' })
  @ApiResponse({ status: 200, description: 'Pick completed' })
  async pickDelivery(@Param('id') id: string, @Body() dto: PickDto) {
    return this.wmsService.pickDelivery(id, dto);
  }

  @Post('delivery-orders/:id/ship')
  @ApiOperation({ summary: 'Ship a delivery order' })
  @ApiResponse({ status: 200, description: 'Delivery shipped' })
  async shipDelivery(@Param('id') id: string, @Body() dto: ShipDto) {
    return this.wmsService.shipDelivery(id, dto);
  }

  @Post('count-plans')
  @ApiOperation({ summary: 'Create an inventory count plan' })
  @ApiResponse({ status: 201, description: 'Count plan created' })
  async createCountPlan(@Body() dto: CreateCountPlanDto) {
    return this.wmsService.createCountPlan(dto);
  }

  @Post('count-records')
  @ApiOperation({ summary: 'Submit an inventory count record' })
  @ApiResponse({ status: 201, description: 'Count record created' })
  async createCountRecord(@Body() dto: CreateCountRecordDto) {
    return this.wmsService.createCountRecord(dto);
  }
}
