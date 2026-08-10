import { Module } from '@nestjs/common';
import { DashboardController } from './dashboard.controller';
import { DashboardService } from './dashboard.service';
import { MesModule } from '../mes/mes.module';
import { QmsModule } from '../qms/qms.module';
import { EamModule } from '../eam/eam.module';
import { WmsModule } from '../wms/wms.module';
import { AndonModule } from '../andon/andon.module';

/**
 * Dashboard module — aggregation layer.
 *
 * Imports the domain modules whose services it needs to aggregate KPIs
 * for the dashboard. Each imported module exports its service.
 */
@Module({
  imports: [MesModule, QmsModule, EamModule, WmsModule, AndonModule],
  controllers: [DashboardController],
  providers: [DashboardService],
})
export class DashboardModule {}
