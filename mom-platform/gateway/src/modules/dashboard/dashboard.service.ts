import { Injectable, Logger } from '@nestjs/common';
import { MesService } from '../mes/mes.service';
import { QmsService } from '../qms/qms.service';
import { EamService } from '../eam/eam.service';
import { WmsService } from '../wms/wms.service';
import { AndonService } from '../andon/andon.service';

/** 趋势/聚合类查询一次抓取的上限，避免看板把整库拉穿。 */
const AGG_PAGE_SIZE = 500;

/** 趋势默认回看天数。 */
const DEFAULT_TREND_DAYS = 7;

/**
 * Dashboard BFF service — 跨域聚合层。
 *
 * 复用各域 BFF service（它们再委托给对应 gRPC 微服务），把多域指标
 * 汇总成看板所需的 KPI。设计要点：
 *
 * 1. **容错**：任一域不可用时只让该域指标降级为 0/空，不让整个看板 500。
 * 2. **契约差异内化**：各域分页字段不统一（MES 用 `pagination`，
 *    QMS/WMS/ANDON 用 `page`），时间戳格式也不同（MES 是
 *    `{seconds,nanos}`，QMS 是 unix 秒字符串），统一在此消化。
 * 3. **数值来源真实**：proto 里 decimal 一律走 string，这里显式转数字。
 */
@Injectable()
export class DashboardService {
  private readonly logger = new Logger(DashboardService.name);

  constructor(
    private readonly mesService: MesService,
    private readonly qmsService: QmsService,
    private readonly eamService: EamService,
    private readonly wmsService: WmsService,
    private readonly andonService: AndonService,
  ) {}

  // ------------------------------------------------------------------ utils

  /** 单域失败不拖垮整个看板：记一条 warn，返回兜底值。 */
  private async safe<T>(label: string, fn: () => Promise<any>, fallback: T): Promise<T> {
    try {
      return (await fn()) as T;
    } catch (err) {
      this.logger.warn(`dashboard: ${label} 聚合失败，已降级 — ${(err as Error).message}`);
      return fallback;
    }
  }

  /** proto 的 decimal/int64 会序列化成字符串，统一转数字。 */
  private num(v: unknown): number {
    const n = Number(v ?? 0);
    return Number.isFinite(n) ? n : 0;
  }

  /** 兼容 `pagination.total`（MES）与 `page.total`（QMS/WMS/ANDON）。 */
  private total(resp: any): number {
    return this.num(resp?.pagination?.total ?? resp?.page?.total ?? 0);
  }

  private items(resp: any): any[] {
    return Array.isArray(resp?.items) ? resp.items : [];
  }

  /**
   * 归一化时间戳到 `YYYY-MM-DD`。
   * 支持 `{seconds,nanos}`（protobuf Timestamp）、unix 秒（数字或字符串）、ISO 字符串。
   */
  private toDateKey(v: any): string | null {
    if (v == null) return null;
    let ms: number;
    if (typeof v === 'object' && v.seconds != null) {
      ms = this.num(v.seconds) * 1000;
    } else if (typeof v === 'string' && v.includes('-')) {
      ms = Date.parse(v);
    } else {
      ms = this.num(v) * 1000;
    }
    if (!ms || Number.isNaN(ms)) return null;
    return new Date(ms).toISOString().slice(0, 10);
  }

  /** 生成最近 n 天的日期键，保证趋势图坐标轴连续（没有数据的日子补 0）。 */
  private recentDateKeys(days: number): string[] {
    const keys: string[] = [];
    const today = new Date();
    for (let i = days - 1; i >= 0; i--) {
      const d = new Date(today);
      d.setDate(today.getDate() - i);
      keys.push(d.toISOString().slice(0, 10));
    }
    return keys;
  }

  private trendDays(query: any): number {
    const n = Number(query?.days);
    if (Number.isFinite(n) && n > 0 && n <= 90) return Math.floor(n);
    return DEFAULT_TREND_DAYS;
  }

  // --------------------------------------------------------------- overview

  async getOverview() {
    // OEE 只看最近 DEFAULT_TREND_DAYS 天，避免历史久远的记录把当前稼动水平拉平。
    const oeeWindowStart = this.recentDateKeys(DEFAULT_TREND_DAYS)[0];

    const [orders, sheets, oeeResp, balances, alarms] = await Promise.all([
      this.safe('mes.orders', () => this.mesService.listOrders({ page: 1, pageSize: 1 }), {} as any),
      this.safe(
        'qms.inspectionSheets',
        () => this.qmsService.listInspectionSheets({ page: 1, pageSize: AGG_PAGE_SIZE }),
        {} as any,
      ),
      this.safe('eam.oee', () => this.eamService.getOee({ beginDate: oeeWindowStart }), {} as any),
      this.safe(
        'wms.balances',
        () => this.wmsService.listBalances({ page: 1, pageSize: AGG_PAGE_SIZE }),
        {} as any,
      ),
      this.safe(
        'andon.calls',
        () => this.andonService.listCalls({ status: 'OPEN', page: 1, pageSize: 1 }),
        {} as any,
      ),
    ]);

    // 质量合格率：只统计已判定（PASSED/FAILED）的检验单，未判定的不计入分母。
    const sheetItems = this.items(sheets);
    let passed = 0;
    let judged = 0;
    for (const s of sheetItems) {
      const st = String(s?.status ?? '');
      if (st.endsWith('_PASSED')) {
        passed++;
        judged++;
      } else if (st.endsWith('_FAILED')) {
        judged++;
      }
    }
    const qualityRate = judged > 0 ? Number(((passed / judged) * 100).toFixed(2)) : 0;

    // OEE：ListOee 返回逐设备明细，看板取算术平均。
    const oeeItems = this.items(oeeResp);
    const oeeAvg =
      oeeItems.length > 0
        ? Number(
            (
              oeeItems.reduce((sum, o) => sum + this.num(o?.oee), 0) / oeeItems.length
            ).toFixed(4),
          )
        : 0;

    // 库存金额：sum(数量 × 单位成本)。单位成本未维护时该行计 0。
    const balanceItems = this.items(balances);
    const inventoryValue = Number(
      balanceItems
        .reduce((sum, b) => sum + this.num(b?.quantity) * this.num(b?.unitCost), 0)
        .toFixed(2),
    );
    const inventoryQty = balanceItems.reduce((sum, b) => sum + this.num(b?.quantity), 0);

    return {
      totalOrders: this.total(orders),
      qualityRate,
      inspectionJudged: judged,
      oeeAvg,
      inventoryValue,
      inventoryQty,
      inventorySkus: balanceItems.length,
      activeAlarms: this.total(alarms),
      timestamp: new Date().toISOString(),
    };
  }

  // ------------------------------------------------------------ 生产趋势

  async getProductionTrend(query: any) {
    const days = this.trendDays(query);
    const resp = await this.safe(
      'mes.orders(trend)',
      () => this.mesService.listOrders({ page: 1, pageSize: AGG_PAGE_SIZE }),
      {} as any,
    );

    const buckets = new Map<
      string,
      { date: string; orders: number; planned: number; completed: number; rejected: number }
    >();
    for (const key of this.recentDateKeys(days)) {
      buckets.set(key, { date: key, orders: 0, planned: 0, completed: 0, rejected: 0 });
    }

    for (const o of this.items(resp)) {
      const key = this.toDateKey(o?.base?.createdAt ?? o?.createdAt);
      const bucket = key ? buckets.get(key) : undefined;
      if (!bucket) continue; // 落在窗口外的直接跳过
      bucket.orders += 1;
      bucket.planned += this.num(o?.quantity);
      bucket.completed += this.num(o?.completedQty);
      bucket.rejected += this.num(o?.rejectedQty);
    }

    const points = [...buckets.values()].map((b) => ({
      ...b,
      yieldRate:
        b.completed + b.rejected > 0
          ? Number(((b.completed / (b.completed + b.rejected)) * 100).toFixed(2))
          : 0,
    }));

    return {
      period: query?.period ?? 'day',
      days,
      points,
      totals: {
        orders: points.reduce((s, p) => s + p.orders, 0),
        planned: points.reduce((s, p) => s + p.planned, 0),
        completed: points.reduce((s, p) => s + p.completed, 0),
        rejected: points.reduce((s, p) => s + p.rejected, 0),
      },
      timestamp: new Date().toISOString(),
    };
  }

  // ------------------------------------------------------------ 质量趋势

  async getQualityTrend(query: any) {
    const days = this.trendDays(query);
    const resp = await this.safe(
      'qms.inspectionSheets(trend)',
      () => this.qmsService.listInspectionSheets({ page: 1, pageSize: AGG_PAGE_SIZE }),
      {} as any,
    );

    const buckets = new Map<
      string,
      { date: string; total: number; passed: number; failed: number; defects: number }
    >();
    for (const key of this.recentDateKeys(days)) {
      buckets.set(key, { date: key, total: 0, passed: 0, failed: 0, defects: 0 });
    }

    for (const s of this.items(resp)) {
      const key = this.toDateKey(s?.createdAt ?? s?.base?.createdAt);
      const bucket = key ? buckets.get(key) : undefined;
      if (!bucket) continue;
      bucket.total += 1;
      const st = String(s?.status ?? '');
      if (st.endsWith('_PASSED')) bucket.passed += 1;
      else if (st.endsWith('_FAILED')) bucket.failed += 1;
      bucket.defects += this.num(s?.defectCount);
    }

    const points = [...buckets.values()].map((b) => ({
      ...b,
      passRate:
        b.passed + b.failed > 0
          ? Number(((b.passed / (b.passed + b.failed)) * 100).toFixed(2))
          : 0,
    }));

    return {
      period: query?.period ?? 'day',
      days,
      points,
      totals: {
        sheets: points.reduce((s, p) => s + p.total, 0),
        passed: points.reduce((s, p) => s + p.passed, 0),
        failed: points.reduce((s, p) => s + p.failed, 0),
        defects: points.reduce((s, p) => s + p.defects, 0),
      },
      timestamp: new Date().toISOString(),
    };
  }

  // ---------------------------------------------------------------- 告警

  async getAlarms() {
    const [callsResp, alertsResp] = await Promise.all([
      this.safe(
        'andon.calls',
        () => this.andonService.listCalls({ status: 'OPEN', page: 1, pageSize: 50 }),
        {} as any,
      ),
      this.safe(
        'andon.alerts',
        () => this.andonService.listAlerts({ page: 1, pageSize: 50 }),
        {} as any,
      ),
    ]);

    const calls = this.items(callsResp);
    const alerts = this.items(alertsResp);

    // 合并成统一告警流，按发生时间倒序，便于前端直接渲染一个列表。
    const merged = [
      ...calls.map((c) => ({
        source: 'ANDON_CALL' as const,
        id: String(c?.id ?? c?.base?.id ?? ''),
        type: c?.andonType ?? c?.type ?? '',
        status: c?.status ?? '',
        description: c?.description ?? '',
        occurredAt: this.num(c?.triggeredAt ?? c?.createdAt ?? c?.base?.createdAt?.seconds),
      })),
      ...alerts.map((a) => ({
        source: 'ALERT' as const,
        id: String(a?.id ?? a?.base?.id ?? ''),
        type: a?.alertType ?? a?.type ?? '',
        status: a?.status ?? '',
        description: a?.message ?? a?.description ?? '',
        occurredAt: this.num(a?.triggeredAt ?? a?.createdAt ?? a?.base?.createdAt?.seconds),
      })),
    ].sort((x, y) => y.occurredAt - x.occurredAt);

    return {
      calls,
      alerts,
      merged,
      summary: {
        openCalls: calls.length,
        activeAlerts: alerts.length,
        total: merged.length,
      },
      timestamp: new Date().toISOString(),
    };
  }
}
