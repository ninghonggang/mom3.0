import { Body, Controller, Get, Param, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { AndonService } from './andon.service';
import {
  CreateAndonCallDto,
  AcknowledgeCallDto,
  ResolveCallDto,
  CreateAlertDto,
  CreateAlertConfigDto,
} from './dto/andon.dto';

@ApiTags('Andon - 安灯管理')
@Controller('api/andon')
export class AndonController {
  constructor(private readonly andonService: AndonService) {}

  @Post('calls')
  @ApiOperation({ summary: 'Create an andon call' })
  @ApiResponse({ status: 201, description: 'Andon call created' })
  async createCall(@Body() dto: CreateAndonCallDto) {
    return this.andonService.createCall(dto);
  }

  @Post('calls/:id/acknowledge')
  @ApiOperation({ summary: 'Acknowledge an andon call' })
  @ApiResponse({ status: 200, description: 'Call acknowledged' })
  async acknowledgeCall(@Param('id') id: string, @Body() dto: AcknowledgeCallDto) {
    return this.andonService.acknowledgeCall(id, dto);
  }

  @Post('calls/:id/resolve')
  @ApiOperation({ summary: 'Resolve an andon call' })
  @ApiResponse({ status: 200, description: 'Call resolved' })
  async resolveCall(@Param('id') id: string, @Body() dto: ResolveCallDto) {
    return this.andonService.resolveCall(id, dto);
  }

  @Get('calls')
  @ApiOperation({ summary: 'List andon calls' })
  @ApiQuery({ name: 'status', required: false, type: String })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated andon calls' })
  async listCalls(@Query() query: any) {
    return this.andonService.listCalls(query);
  }

  @Post('alert-configs')
  @ApiOperation({ summary: 'Create an alert rule config (prerequisite for triggering alerts)' })
  @ApiResponse({ status: 201, description: 'Alert config created' })
  async createAlertConfig(@Body() dto: CreateAlertConfigDto) {
    return this.andonService.createAlertConfig(dto);
  }

  @Post('alerts')
  @ApiOperation({ summary: 'Trigger an alert against an existing alert config' })
  @ApiResponse({ status: 201, description: 'Alert triggered' })
  async createAlert(@Body() dto: CreateAlertDto) {
    return this.andonService.createAlert(dto);
  }

  @Get('alerts')
  @ApiOperation({ summary: 'List alerts' })
  @ApiQuery({ name: 'severity', required: false, type: String })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated alerts' })
  async listAlerts(@Query() query: any) {
    return this.andonService.listAlerts(query);
  }
}
