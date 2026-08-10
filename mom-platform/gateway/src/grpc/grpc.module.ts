import { Global, Module } from '@nestjs/common';
import { GrpcService } from './grpc.service';
import { ResolverService } from './resolver.service';

/**
 * 全局 gRPC 客户端模块。
 * 所有业务域 service 均可注入 GrpcService（原始调用）
 * 与 ResolverService（业务编码 → id 解析）。
 */
@Global()
@Module({
  providers: [GrpcService, ResolverService],
  exports: [GrpcService, ResolverService],
})
export class GrpcModule {}
