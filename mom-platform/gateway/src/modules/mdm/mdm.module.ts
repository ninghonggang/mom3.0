import { Module } from '@nestjs/common';
import { MdmController } from './mdm.controller';
import { MdmService } from './mdm.service';

/**
 * MDM — 主数据管理 (Master Data Management)
 *
 * 通过 GrpcService 以 @grpc/proto-loader 动态调用 gRPC 微服务（默认 :50051），
 * 无需在此注册静态 client。
 */
@Module({
  controllers: [MdmController],
  providers: [MdmService],
})
export class MdmModule {}
