import { Module } from '@nestjs/common';
import { MesController } from './mes.controller';
import { MesService } from './mes.service';

/**
 * MES — 生产执行 (Manufacturing Execution)
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50052），
 * 无需在此注册静态 client。同时导出 Service 供 Dashboard 聚合层复用。
 */
@Module({
  controllers: [MesController],
  providers: [MesService],
  exports: [MesService],
})
export class MesModule {}
