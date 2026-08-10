import { Module } from '@nestjs/common';
import { EamController } from './eam.controller';
import { EamService } from './eam.service';

/**
 * EAM — 设备资产管理 (Enterprise Asset Management)
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50054），
 * 无需在此注册静态 client。同时导出 Service 供 Dashboard 聚合层复用。
 */
@Module({
  controllers: [EamController],
  providers: [EamService],
  exports: [EamService],
})
export class EamModule {}
