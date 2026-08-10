import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { BarChart } from '@/components/charts/bar-chart';
import { formatNumber, formatPercent, formatDateTime } from '@/lib/utils';

// Mock data
const stats = {
  totalOrders: 128,
  qualityPassRate: 0.967,
  oeeAverage: 0.854,
  activeAlarms: 7,
};

const productionTrend = [
  { label: '周一', planned: 420, actual: 405 },
  { label: '周二', planned: 450, actual: 438 },
  { label: '周三', planned: 435, actual: 430 },
  { label: '周四', planned: 460, actual: 455 },
  { label: '周五', planned: 445, actual: 442 },
  { label: '周六', planned: 250, actual: 248 },
  { label: '周日', planned: 120, actual: 118 },
];

const recentAlarms = [
  { id: '1', title: 'CNC-03 主轴温度过高', severity: 'CRITICAL', source: 'EAM', status: 'ACTIVE', createdAt: new Date().toISOString(), message: '' },
  { id: '2', title: '产线A 物料短缺预警', severity: 'WARNING', source: 'WMS', status: 'ACTIVE', createdAt: new Date().toISOString(), message: '' },
  { id: '3', title: '质检工位 IPQC-02 超时未处理', severity: 'WARNING', source: 'QMS', status: 'ACKNOWLEDGED', createdAt: new Date().toISOString(), message: '' },
  { id: '4', title: 'SC-01 焊接质量异常', severity: 'CRITICAL', source: 'QMS', status: 'ACTIVE', createdAt: new Date().toISOString(), message: '' },
  { id: '5', title: '设备 ER-104 保养到期提醒', severity: 'INFO', source: 'EAM', status: 'ACTIVE', createdAt: new Date().toISOString(), message: '' },
];

const orderDistribution = [
  { status: '待下达', count: 12, color: 'bg-gray-400' },
  { status: '已下达', count: 18, color: 'bg-blue-500' },
  { status: '生产中', count: 45, color: 'bg-yellow-500' },
  { status: '已完成', count: 38, color: 'bg-green-500' },
  { status: '已关闭', count: 12, color: 'bg-gray-300' },
  { status: '暂停', count: 3, color: 'bg-orange-500' },
];

const severityColors: Record<string, string> = {
  CRITICAL: 'bg-red-100 text-red-700',
  WARNING: 'bg-yellow-100 text-yellow-700',
  INFO: 'bg-blue-100 text-blue-700',
};

const severityLabels: Record<string, string> = {
  CRITICAL: '严重',
  WARNING: '警告',
  INFO: '信息',
};

export default function DashboardPage() {
  return (
    <div className="space-y-6">
      {/* Page Header */}
      <div className="page-header">
        <h1 className="page-title">仪表盘</h1>
        <span className="text-sm text-gray-500">
          数据更新时间：{formatDateTime(new Date().toISOString())}
        </span>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
        {/* Total Orders */}
        <Card>
          <CardContent className="p-5">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-500">今日生产订单</p>
                <p className="mt-1 text-3xl font-bold text-gray-900">{stats.totalOrders}</p>
              </div>
              <div className="flex h-12 w-12 items-center justify-center rounded-full bg-blue-50">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#3b82f6" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                  <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z" />
                  <polyline points="14 2 14 8 20 8" />
                  <line x1="16" y1="13" x2="8" y2="13" />
                  <line x1="16" y1="17" x2="8" y2="17" />
                </svg>
              </div>
            </div>
            <div className="mt-3 flex items-center gap-1 text-xs text-green-600">
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="18 15 12 9 6 15" /></svg>
              <span>较昨日 +5.2%</span>
            </div>
          </CardContent>
        </Card>

        {/* Quality Pass Rate */}
        <Card>
          <CardContent className="p-5">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-500">质量合格率</p>
                <p className="mt-1 text-3xl font-bold text-gray-900">{formatPercent(stats.qualityPassRate)}</p>
              </div>
              <div className="flex h-12 w-12 items-center justify-center rounded-full bg-green-50">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#22c55e" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                  <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" />
                  <polyline points="22 4 12 14.01 9 11.01" />
                </svg>
              </div>
            </div>
            <div className="mt-3 flex items-center gap-1 text-xs text-green-600">
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="18 15 12 9 6 15" /></svg>
              <span>较上周 +0.3%</span>
            </div>
          </CardContent>
        </Card>

        {/* OEE */}
        <Card>
          <CardContent className="p-5">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-500">OEE 平均值</p>
                <p className="mt-1 text-3xl font-bold text-gray-900">{formatPercent(stats.oeeAverage)}</p>
              </div>
              <div className="flex h-12 w-12 items-center justify-center rounded-full bg-purple-50">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#a855f7" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                  <path d="M12 20V10" />
                  <path d="M18 20V4" />
                  <path d="M6 20v-4" />
                </svg>
              </div>
            </div>
            <div className="mt-3 flex items-center gap-1 text-xs text-green-600">
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="18 15 12 9 6 15" /></svg>
              <span>目标达成率 98%</span>
            </div>
          </CardContent>
        </Card>

        {/* Active Alarms */}
        <Card>
          <CardContent className="p-5">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-500">活跃告警</p>
                <p className="mt-1 text-3xl font-bold text-red-600">{stats.activeAlarms}</p>
              </div>
              <div className="flex h-12 w-12 items-center justify-center rounded-full bg-red-50">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#ef4444" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                  <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9" />
                  <path d="M13.73 21a2 2 0 0 1-3.46 0" />
                </svg>
              </div>
            </div>
            <div className="mt-3 flex items-center gap-1 text-xs text-red-600">
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="18 6 12 12 6 6" /></svg>
              <span>其中严重级别 2 项</span>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Charts and Detail Row */}
      <div className="grid grid-cols-1 gap-6 lg:grid-cols-3">
        {/* Production Trend */}
        <Card className="lg:col-span-2">
          <CardHeader>
            <CardTitle>生产趋势（本周）</CardTitle>
          </CardHeader>
          <CardContent>
            <BarChart data={productionTrend} height={220} />
          </CardContent>
        </Card>

        {/* Order Status Distribution */}
        <Card>
          <CardHeader>
            <CardTitle>订单状态分布</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              {orderDistribution.map((item) => {
                const maxCount = Math.max(...orderDistribution.map((d) => d.count));
                return (
                  <div key={item.status}>
                    <div className="mb-1 flex items-center justify-between text-sm">
                      <span className="text-gray-600">{item.status}</span>
                      <span className="font-medium text-gray-900">{item.count}</span>
                    </div>
                    <div className="h-2 w-full rounded-full bg-gray-100">
                      <div
                        className={`h-2 rounded-full ${item.color} transition-all`}
                        style={{ width: `${(item.count / maxCount) * 100}%` }}
                      />
                    </div>
                  </div>
                );
              })}
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Recent Alarms */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between">
          <CardTitle>最近告警</CardTitle>
          <button className="text-sm text-blue-600 hover:text-blue-700">查看全部 &rarr;</button>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            {recentAlarms.map((alarm) => (
              <div
                key={alarm.id}
                className="flex items-center justify-between rounded-lg border border-gray-100 bg-gray-50 p-3 transition-colors hover:bg-gray-100"
              >
                <div className="flex items-center gap-3">
                  <span
                    className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${
                      severityColors[alarm.severity]
                    }`}
                  >
                    {severityLabels[alarm.severity]}
                  </span>
                  <div>
                    <p className="text-sm font-medium text-gray-900">{alarm.title}</p>
                    <p className="text-xs text-gray-500">
                      {alarm.source} · {formatDateTime(alarm.createdAt)}
                    </p>
                  </div>
                </div>
                <Badge
                  variant={alarm.status === 'ACTIVE' ? 'destructive' : 'warning'}
                >
                  {alarm.status === 'ACTIVE' ? '待处理' : '已确认'}
                </Badge>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
