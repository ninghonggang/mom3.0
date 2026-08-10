import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { TraceService } from './trace.service';
import { CreateTraceRecordDto, GenerateSerialDto, DataPointDto, CollectDto, CreateScanLogDto } from './dto/trace.dto';

@ApiTags('Trace - 追溯管理')
@Controller('api/trace')
export class TraceController {
  constructor(private readonly traceService: TraceService) {}

  @Post('records')
  @ApiOperation({ summary: 'Create a traceability record' })
  @ApiResponse({ status: 201, description: 'Trace record created' })
  async createTraceRecord(@Body() dto: CreateTraceRecordDto) {
    return this.traceService.createTraceRecord(dto);
  }

  @Get('forward')
  @ApiOperation({ summary: 'Forward traceability (downstream genealogy)' })
  @ApiQuery({ name: 'serialNo', required: true, type: String })
  @ApiResponse({ status: 200, description: 'Forward trace result' })
  async forwardTrace(@Query('serialNo') serialNo: string, @Query() query: any) {
    return this.traceService.forwardTrace({ serialNo, ...query });
  }

  @Get('backward')
  @ApiOperation({ summary: 'Backward traceability (upstream genealogy)' })
  @ApiQuery({ name: 'serialNo', required: true, type: String })
  @ApiResponse({ status: 200, description: 'Backward trace result' })
  async backwardTrace(@Query('serialNo') serialNo: string, @Query() query: any) {
    return this.traceService.backwardTrace({ serialNo, ...query });
  }

  @Post('serials/generate')
  @ApiOperation({ summary: 'Generate serial numbers' })
  @ApiResponse({ status: 200, description: 'Serials generated' })
  async generateSerials(@Body() dto: GenerateSerialDto) {
    return this.traceService.generateSerials(dto);
  }

  @Get('data-points')
  @ApiOperation({ summary: 'List traceability data points' })
  @ApiQuery({ name: 'serialNo', required: false, type: String })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated data points' })
  async listDataPoints(@Query() query: any) {
    return this.traceService.listDataPoints(query);
  }

  @Post('data-points')
  @ApiOperation({ summary: 'Create a single traceability data point' })
  @ApiResponse({ status: 201, description: 'Data point created' })
  async createDataPoint(@Body() dto: DataPointDto) {
    return this.traceService.createDataPoint(dto);
  }

  @Post('collect')
  @ApiOperation({ summary: 'Batch collect traceability data points' })
  @ApiResponse({ status: 200, description: 'Batch accepted' })
  async collect(@Body() dto: CollectDto) {
    return this.traceService.collect(dto);
  }

  @Post('scan-logs')
  @ApiOperation({ summary: 'Record a scan log' })
  @ApiResponse({ status: 201, description: 'Scan log created' })
  async createScanLog(@Body() dto: CreateScanLogDto) {
    return this.traceService.createScanLog(dto);
  }
}
