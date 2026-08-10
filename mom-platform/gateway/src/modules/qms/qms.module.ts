import { Module } from '@nestjs/common';
import { QmsController } from './qms.controller';
import { QmsService } from './qms.service';

/**
 * QMS — 质量管理 (Quality Management)
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50053），
 * 无需在此注册静态 client。同时导出 Service 供 Dashboard 聚合层复用。
 */
@Module({
  controllers: [QmsController],
  providers: [QmsService],
  exports: [QmsService],
})
export class QmsModule {}
