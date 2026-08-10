import { Body, Controller, Get, Param, Patch, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiQuery, ApiResponse, ApiTags } from '@nestjs/swagger';
import { QmsService } from './qms.service';
import {
  CreateInspectionSheetDto,
  UpdateInspectionSheetDto,
  CreateCharacteristicDto,
  SubmitResultsDto,
  CreateNcrDto,
  UpdateNcrDto,
  NcrActionDto,
} from './dto/qms.dto';

@ApiTags('QMS - 质量管理')
@Controller('api/qms')
export class QmsController {
  constructor(private readonly qmsService: QmsService) {}

  @Get('inspection-sheets')
  @ApiOperation({ summary: 'List inspection sheets' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated inspection sheets' })
  async listInspectionSheets(@Query() query: any) {
    return this.qmsService.listInspectionSheets(query);
  }

  @Post('inspection-sheets')
  @ApiOperation({ summary: 'Create an inspection sheet' })
  @ApiResponse({ status: 201, description: 'Inspection sheet created' })
  async createInspectionSheet(@Body() dto: CreateInspectionSheetDto) {
    return this.qmsService.createInspectionSheet(dto);
  }

  @Get('inspection-sheets/:id')
  @ApiOperation({ summary: 'Get inspection sheet detail' })
  @ApiResponse({ status: 200, description: 'Inspection sheet detail' })
  async getInspectionSheet(@Param('id') id: string) {
    return this.qmsService.getInspectionSheet(id);
  }

  @Patch('inspection-sheets/:id')
  @ApiOperation({ summary: 'Update an inspection sheet (status / defect count)' })
  @ApiResponse({ status: 200, description: 'Inspection sheet updated' })
  async updateInspectionSheet(@Param('id') id: string, @Body() dto: UpdateInspectionSheetDto) {
    return this.qmsService.updateInspectionSheet(id, dto);
  }

  @Get('characteristics')
  @ApiOperation({ summary: 'List inspection characteristics' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated inspection characteristics' })
  async listCharacteristics(@Query() query: any) {
    return this.qmsService.listCharacteristics(query);
  }

  @Post('characteristics')
  @ApiOperation({ summary: 'Create an inspection characteristic' })
  @ApiResponse({ status: 201, description: 'Characteristic created' })
  async createCharacteristic(@Body() dto: CreateCharacteristicDto) {
    return this.qmsService.createCharacteristic(dto);
  }

  @Post('inspection-sheets/:id/results')
  @ApiOperation({ summary: 'Submit inspection results for a sheet' })
  @ApiResponse({ status: 200, description: 'Results submitted' })
  async submitResults(@Param('id') id: string, @Body() dto: SubmitResultsDto) {
    return this.qmsService.submitResults(id, dto);
  }

  @Get('ncrs')
  @ApiOperation({ summary: 'List non-conformance reports (NCRs)' })
  @ApiQuery({ name: 'page', required: false, type: Number })
  @ApiQuery({ name: 'pageSize', required: false, type: Number })
  @ApiResponse({ status: 200, description: 'Paginated NCRs' })
  async listNcrs(@Query() query: any) {
    return this.qmsService.listNcrs(query);
  }

  @Post('ncrs')
  @ApiOperation({ summary: 'Create an NCR' })
  @ApiResponse({ status: 201, description: 'NCR created' })
  async createNcr(@Body() dto: CreateNcrDto) {
    return this.qmsService.createNcr(dto);
  }

  @Patch('ncrs/:id')
  @ApiOperation({ summary: 'Update an NCR' })
  @ApiResponse({ status: 200, description: 'NCR updated' })
  async updateNcr(@Param('id') id: string, @Body() dto: UpdateNcrDto) {
    return this.qmsService.updateNcr(id, dto);
  }

  @Post('ncrs/:id/actions')
  @ApiOperation({ summary: 'Add a disposition action to an NCR' })
  @ApiResponse({ status: 200, description: 'Action added' })
  async addNcrAction(@Param('id') id: string, @Body() dto: NcrActionDto) {
    return this.qmsService.addAction(id, dto);
  }

  @Get('spc-data')
  @ApiOperation({ summary: 'Get SPC (statistical process control) data' })
  @ApiQuery({ name: 'itemCode', required: false, type: String })
  @ApiQuery({ name: 'from', required: false, type: String })
  @ApiQuery({ name: 'to', required: false, type: String })
  @ApiResponse({ status: 200, description: 'SPC data points' })
  async getSpcData(@Query() query: any) {
    return this.qmsService.getSpcData(query);
  }
}
