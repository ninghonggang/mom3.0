import { Injectable } from '@nestjs/common';
import { GrpcService } from '../../grpc/grpc.service';
import { buildListRequest } from '../../grpc/case.util';
import { CreateMaterialDto, CreateBomDto } from './dto/mdm.dto';

const PKG = 'mom.mdm';
const URL = process.env.MDM_SERVICE_URL || 'localhost:50051';

/**
 * MDM BFF service — 委托给 MDM gRPC 微服务 (mdm-service:50051)。
 */
@Injectable()
export class MdmService {
  constructor(private readonly grpc: GrpcService) {}

  listMaterials(query: any) {
    return this.grpc.call(PKG, 'MaterialService', 'ListMaterials', buildListRequest(query, 'pagination'), URL);
  }
  createMaterial(dto: CreateMaterialDto) {
    return this.grpc.call(
      PKG,
      'MaterialService',
      'CreateMaterial',
      {
        material: {
          material_code: dto.materialCode,
          material_name: dto.name,
          spec: dto.specification,
          unit_name: dto.unit,
          category_name: dto.category,
          material_type: dto.materialType,
        },
      },
      URL,
    );
  }
  listBoms(query: any) {
    return this.grpc.call(PKG, 'BOMService', 'ListBOMs', buildListRequest(query, 'pagination'), URL);
  }
  createBom(dto: CreateBomDto) {
    return this.grpc.call(PKG, 'BOMService', 'CreateBOM', dto, URL);
  }
  listWorkshops(query: any) {
    return this.grpc.call(PKG, 'FactoryAssetService', 'ListWorkshops', buildListRequest(query, 'pagination'), URL);
  }
  listProductionLines(query: any) {
    return this.grpc.call(PKG, 'FactoryAssetService', 'ListProductionLines', buildListRequest(query, 'pagination'), URL);
  }
  listWorkstations(query: any) {
    return this.grpc.call(PKG, 'FactoryAssetService', 'ListWorkstations', buildListRequest(query, 'pagination'), URL);
  }
  listCustomers(query: any) {
    return this.grpc.call(PKG, 'PartnerService', 'ListCustomers', buildListRequest(query, 'pagination'), URL);
  }
  listSuppliers(query: any) {
    return this.grpc.call(PKG, 'PartnerService', 'ListSuppliers', buildListRequest(query, 'pagination'), URL);
  }
}
