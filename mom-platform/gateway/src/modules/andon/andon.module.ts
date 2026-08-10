import { Module } from '@nestjs/common';
import { AndonController } from './andon.controller';
import { AndonService } from './andon.service';

/**
 * ANDON — 安灯呼叫与告警
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50058），
 * 无需在此注册静态 client。同时导出 Service 供 Dashboard 聚合层复用。
 */
@Module({
  controllers: [AndonController],
  providers: [AndonService],
  exports: [AndonService],
})
export class AndonModule {}
