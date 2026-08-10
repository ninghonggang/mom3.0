import { Module } from '@nestjs/common';
import { GrpcModule } from './grpc/grpc.module';
import { MesModule } from './modules/mes/mes.module';
import { QmsModule } from './modules/qms/qms.module';
import { EamModule } from './modules/eam/eam.module';
import { WmsModule } from './modules/wms/wms.module';
import { MdmModule } from './modules/mdm/mdm.module';
import { ApsModule } from './modules/aps/aps.module';
import { TraceModule } from './modules/trace/trace.module';
import { AndonModule } from './modules/andon/andon.module';
import { DashboardModule } from './modules/dashboard/dashboard.module';

@Module({
  imports: [
    // Shared gRPC client factory (global)
    GrpcModule,
    // Core MOM domain modules — each delegates to its gRPC microservice
    MesModule,
    QmsModule,
    EamModule,
    WmsModule,
    MdmModule,
    ApsModule,
    TraceModule,
    AndonModule,
    // Aggregation layer
    DashboardModule,
  ],
})
export class AppModule {}
