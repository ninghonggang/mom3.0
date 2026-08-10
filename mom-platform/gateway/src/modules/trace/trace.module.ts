import { Module } from '@nestjs/common';
import { TraceController } from './trace.controller';
import { TraceService } from './trace.service';

/**
 * TRACE — 产品追溯 (Traceability)
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50057），
 * 无需在此注册静态 client。
 */
@Module({
  controllers: [TraceController],
  providers: [TraceService],
})
export class TraceModule {}
