import { Module } from '@nestjs/common';
import { ApsController } from './aps.controller';
import { ApsService } from './aps.service';

/**
 * APS — 高级计划排程 (Advanced Planning & Scheduling)
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50056），
 * 无需在此注册静态 client。
 */
@Module({
  controllers: [ApsController],
  providers: [ApsService],
})
export class ApsModule {}
