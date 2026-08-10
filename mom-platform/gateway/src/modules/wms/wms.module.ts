import { Module } from '@nestjs/common';
import { WmsController } from './wms.controller';
import { WmsService } from './wms.service';

/**
 * WMS — 仓储管理 (Warehouse Management)
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50055），
 * 无需在此注册静态 client。同时导出 Service 供 Dashboard 聚合层复用。
 */
@Module({
  controllers: [WmsController],
  providers: [WmsService],
  exports: [WmsService],
})
export class WmsModule {}
