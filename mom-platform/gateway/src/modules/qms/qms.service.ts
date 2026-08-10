import { Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { ResolverService } from '../../grpc/resolver.service';
import {
  buildListRequest,
  deepCamelToSnake,
  toDecimalString,
  toInt,
  toProtoEnum,
} from '../../grpc/case.util';
import {
  CreateInspectionSheetDto,
  UpdateInspectionSheetDto,
  CreateCharacteristicDto,
  SubmitResultsDto,
  CreateNcrDto,
  UpdateNcrDto,
  NcrActionDto,
} from './dto/qms.dto';

const PKG = 'mom.qms';
const URL = process.env.QMS_SERVICE_URL || 'localhost:50053';

/**
 * QMS BFF service — 委托给 QMS gRPC 微服务 (qms-service:50053)。
 *
 * proto 侧检验类型/NCR 严重度/状态都是带前缀的枚举，物料以 id 引用；
 * REST 侧沿用裸枚举值和物料编码。此处集中做这层映射。
 */
@Injectable()
export class QmsService {
  constructor(
    private readonly grpc: GrpcService,
    private readonly resolver: ResolverService,
  ) {}

  listInspectionSheets(query: any) {
    return this.grpc.call(PKG, 'QmsService', 'ListInspectionSheets', buildListRequest(query, 'page'), URL);
  }

  async createInspectionSheet(dto: CreateInspectionSheetDto) {
    return this.grpc.call(
      PKG,
      'QmsService',
      'CreateInspectionSheet',
      {
        inspection_type: toProtoEnum('INSPECTION_TYPE', dto.inspectionType),
        material_id: dto.materialCode ? await this.resolver.materialId(dto.materialCode) : 0,
        batch_id: toInt(dto.orderId),
        sample_size: toDecimalString(dto.sampleSize ?? 1),
        inspector_id: toInt(dto.inspectorId),
      },
      URL,
    );
  }

  getInspectionSheet(id: string) {
    return this.grpc.call(PKG, 'QmsService', 'GetInspectionSheet', { id: toInt(id) }, URL);
  }

  /** 变更检验单状态（判定合格/不合格）。NCR 只能建立在 FAILED 的检验单上。 */
  updateInspectionSheet(id: string, dto: UpdateInspectionSheetDto) {
    return this.grpc.call(
      PKG,
      'QmsService',
      'UpdateInspectionSheet',
      {
        id: toInt(id),
        status: toProtoEnum('INSPECTION_SHEET_STATUS', dto.status),
        defect_count: toDecimalString(dto.defectCount ?? 0),
        inspector_id: toInt(dto.inspectorId),
      },
      URL,
    );
  }

  listCharacteristics(query: any) {
    return this.grpc.call(PKG, 'QmsService', 'ListCharacteristics', buildListRequest(query, 'page'), URL);
  }

  /** 建立检验特性（USL/LSL 决定 RecordInspectionResult 的自动判定）。 */
  createCharacteristic(dto: CreateCharacteristicDto) {
    return this.grpc.call(
      PKG,
      'QmsService',
      'CreateCharacteristic',
      {
        char_code: dto.charCode,
        char_name: dto.charName,
        data_type: toProtoEnum('CHAR_DATA_TYPE', dto.dataType ?? 'NUMBER'),
        usl: dto.usl ?? '',
        lsl: dto.lsl ?? '',
        target: dto.target ?? '',
        unit: dto.unit ?? '',
      },
      URL,
    );
  }

  /**
   * 提交检验结果。REST 用 itemCode 标识检验项，proto 用 char_id（检验特性主键）；
   * 数字编码直接透传，非数字按提交顺序退化为序号，保证不丢数据。
   */
  submitResults(id: string, dto: SubmitResultsDto) {
    const entries = (dto.results ?? []).map((r, i) => ({
      char_id: /^\d+$/.test(r.itemCode) ? Number(r.itemCode) : i + 1,
      value: r.measuredValue ?? '',
      pass: r.result === 'PASS',
    }));
    return this.grpc.call(
      PKG,
      'QmsService',
      'RecordInspectionResult',
      { sheet_id: toInt(id), entries },
      URL,
    );
  }

  listNcrs(query: any) {
    return this.grpc.call(PKG, 'QmsService', 'ListNcrs', buildListRequest(query, 'page'), URL);
  }

  async createNcr(dto: CreateNcrDto) {
    return this.grpc.call(
      PKG,
      'QmsService',
      'CreateNcr',
      {
        inspection_sheet_id: toInt(dto.sheetId),
        material_id: dto.materialCode ? await this.resolver.materialId(dto.materialCode) : 0,
        batch_id: 0,
        quantity: toDecimalString(dto.quantity ?? 1),
        severity: toProtoEnum('NCR_SEVERITY', dto.severity),
      },
      URL,
    );
  }

  updateNcr(id: string, dto: UpdateNcrDto) {
    return this.grpc.call(
      PKG,
      'QmsService',
      'UpdateNcr',
      { id: toInt(id), status: toProtoEnum('NCR_STATUS', dto.status) },
      URL,
    );
  }

  addAction(id: string, dto: NcrActionDto) {
    return this.grpc.call(
      PKG,
      'QmsService',
      'AddNcrAction',
      {
        ncr_id: toInt(id),
        action_type: dto.actionType,
        action_desc: dto.description ?? '',
        responsible_id: toInt(dto.handlerId),
      },
      URL,
    );
  }

  getSpcData(query: any) {
    return this.grpc.call(PKG, 'QmsService', 'ListSpcData', deepCamelToSnake(query ?? {}), URL);
  }
}
