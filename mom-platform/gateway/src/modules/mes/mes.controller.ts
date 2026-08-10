import { Body, Controller, Get, Param, Patch, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { MesService } from './mes.service';
import {
  CreateOrderDto,
  UpdateOrderStatusDto,
  DispatchDto,
  ReportDto,
  CompleteDto,
} from './dto/mes.dto';

@ApiTags('MES - 制造执行')
@Controller('api/mes')
export class MesController {
  constructor(private readonly mesService: MesService) {}

  @Get('orders')
  @ApiOperation({ summary: 'List production orders' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiQuery({ name: 'status', required: false, type: String })
  @ApiResponse({ status: 200, description: 'Paginated list of production orders' })
  async listOrders(@Query() query: any) {
    return this.mesService.listOrders(query);
  }

  @Post('orders')
  @ApiOperation({ summary: 'Create a production order' })
  @ApiResponse({ status: 201, description: 'Production order created' })
  async createOrder(@Body() dto: CreateOrderDto) {
    return this.mesService.createOrder(dto);
  }

  @Get('orders/:id')
  @ApiOperation({ summary: 'Get production order detail' })
  @ApiResponse({ status: 200, description: 'Production order detail' })
  @ApiResponse({ status: 404, description: 'Order not found' })
  async getOrder(@Param('id') id: string) {
    return this.mesService.getOrder(id);
  }

  @Patch('orders/:id/status')
  @ApiOperation({ summary: 'Update production order status' })
  @ApiResponse({ status: 200, description: 'Order status updated' })
  async updateOrderStatus(@Param('id') id: string, @Body() dto: UpdateOrderStatusDto) {
    return this.mesService.updateOrderStatus(id, dto);
  }

  @Post('orders/:id/dispatch')
  @ApiOperation({ summary: 'Dispatch a production order to a line/workstation' })
  @ApiResponse({ status: 200, description: 'Order dispatched' })
  async dispatchOrder(@Param('id') id: string, @Body() dto: DispatchDto) {
    return this.mesService.dispatchOrder(id, dto);
  }

  @Post('orders/:id/report')
  @ApiOperation({ summary: 'Submit a job report (production output) for an order' })
  @ApiResponse({ status: 200, description: 'Report accepted' })
  async reportOrder(@Param('id') id: string, @Body() dto: ReportDto) {
    return this.mesService.reportOrder(id, dto);
  }

  @Post('orders/:id/complete')
  @ApiOperation({ summary: 'Complete a production order' })
  @ApiResponse({ status: 200, description: 'Order completed' })
  async completeOrder(@Param('id') id: string, @Body() dto: CompleteDto) {
    return this.mesService.completeOrder(id, dto);
  }
}
