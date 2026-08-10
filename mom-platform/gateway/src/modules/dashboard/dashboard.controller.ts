import { Controller, Get, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { DashboardService } from './dashboard.service';

@ApiTags('Dashboard - 综合看板')
@Controller('api/dashboard')
export class DashboardController {
  constructor(private readonly dashboardService: DashboardService) {}

  @Get('overview')
  @ApiOperation({ summary: 'Aggregate overview KPIs (orders, quality rate, OEE, inventory value)' })
  @ApiResponse({ status: 200, description: 'Overview KPIs' })
  async getOverview() {
    return this.dashboardService.getOverview();
  }

  @Get('production-trend')
  @ApiOperation({ summary: 'Production output trend' })
  @ApiQuery({ name: 'period', required: false, type: String, description: 'day | week | month' })
  @ApiResponse({ status: 200, description: 'Production trend data' })
  async getProductionTrend(@Query() query: any) {
    return this.dashboardService.getProductionTrend(query);
  }

  @Get('quality-trend')
  @ApiOperation({ summary: 'Quality (pass-rate) trend' })
  @ApiQuery({ name: 'period', required: false, type: String, description: 'day | week | month' })
  @ApiResponse({ status: 200, description: 'Quality trend data' })
  async getQualityTrend(@Query() query: any) {
    return this.dashboardService.getQualityTrend(query);
  }

  @Get('alarms')
  @ApiOperation({ summary: 'Active alarms (open andon calls + alerts)' })
  @ApiResponse({ status: 200, description: 'Active alarms' })
  async getAlarms() {
    return this.dashboardService.getAlarms();
  }
}
