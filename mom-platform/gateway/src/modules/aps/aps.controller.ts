import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { ApsService } from './aps.service';
import { CreateMpsPlanDto, GenerateMrpDto, CreateScheduleJobDto } from './dto/aps.dto';

@ApiTags('APS - 高级排程')
@Controller('api/aps')
export class ApsController {
  constructor(private readonly apsService: ApsService) {}

  @Get('mps-plans')
  @ApiOperation({ summary: 'List MPS (master production schedule) plans' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated MPS plans' })
  async listMpsPlans(@Query() query: any) {
    return this.apsService.listMpsPlans(query);
  }

  @Post('mps-plans')
  @ApiOperation({ summary: 'Create an MPS plan' })
  @ApiResponse({ status: 201, description: 'MPS plan created' })
  async createMpsPlan(@Body() dto: CreateMpsPlanDto) {
    return this.apsService.createMpsPlan(dto);
  }

  @Post('mrp/generate')
  @ApiOperation({ summary: 'Generate MRP (material requirements planning) result' })
  @ApiResponse({ status: 200, description: 'MRP generated' })
  async generateMrp(@Body() dto: GenerateMrpDto) {
    return this.apsService.generateMrp(dto);
  }

  @Get('schedule-jobs')
  @ApiOperation({ summary: 'List schedule jobs' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated schedule jobs' })
  async listScheduleJobs(@Query() query: any) {
    return this.apsService.listScheduleJobs(query);
  }

  @Post('schedule-jobs')
  @ApiOperation({ summary: 'Create a schedule job' })
  @ApiResponse({ status: 201, description: 'Schedule job created' })
  async createScheduleJob(@Body() dto: CreateScheduleJobDto) {
    return this.apsService.createScheduleJob(dto);
  }
}
