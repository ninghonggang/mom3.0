import { Card, CardContent } from '@/components/ui/card';
import { formatNumber } from '@/lib/utils';

const summaryCards = [
  { label: '仓库总数', value: 6, icon: 'warehouse', color: 'blue' },
  { label: '库存物料种类', value: 1248, icon: 'package', color: 'green' },
  { label: '今日收货(批)', value: 15, icon: 'receive', color: 'purple' },
  { label: '今日发货(批)', value: 22, icon: 'delivery', color: 'orange' },
];

const warehouseList = [
  { name: '原材料仓A', code: 'WH-RAW-01', capacity: 5000, used: 3850, status: '正常' },
  { name: '半成品仓B', code: 'WH-WIP-01', capacity: 3000, used: 2200, status: '正常' },
  { name: '成品仓C', code: 'WH-FG-01', capacity: 8000, used: 5600, status: '正常' },
  { name: '辅料仓D', code: 'WH-AUX-01', capacity: 2000, used: 1800, status: '库存偏高' },
  { name: '待检仓', code: 'WH-QC-01', capacity: 1000, used: 350, status: '正常' },
  { name: '不合格品仓', code: 'WH-NG-01', capacity: 500, used: 120, status: '正常' },
];

const recentReceives = [
  { no: 'REC-2025-0101', supplier: '供应商A', material: '主轴轴承座', qty: 500, date: '2025-01-06' },
  { no: 'REC-2025-0102', supplier: '供应商B', material: '电机法兰盘', qty: 300, date: '2025-01-07' },
  { no: 'REC-2025-0103', supplier: '供应商A', material: '液压阀体', qty: 150, date: '2025-01-08' },
];

export default function WmsPage() {
  return (
    <div className="space-y-6">
      <div className="page-header">
        <h1 className="page-title">仓管总览</h1>
      </div>

      {/* Summary Cards */}
      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
        {summaryCards.map((card) => (
          <Card key={card.label}>
            <CardContent className="p-5">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm text-gray-500">{card.label}</p>
                  <p className="mt-1 text-2xl font-bold text-gray-900">
                    {typeof card.value === 'number' ? formatNumber(card.value) : card.value}
                  </p>
                </div>
                <div className={`flex h-10 w-10 items-center justify-center rounded-full bg-${card.color}-50`}>
                  {card.icon === 'warehouse' && (
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#3b82f6" strokeWidth="2"><rect x="2" y="7" width="20" height="14" rx="2" /><path d="M16 7V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2" /></svg>
                  )}
                  {card.icon === 'package' && (
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#22c55e" strokeWidth="2"><path d="m16.5 9.4-9-5.19" /><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" /><polyline points="3.29 7 12 12 20.71 7" /><line x1="12" y1="22" x2="12" y2="12" /></svg>
                  )}
                  {card.icon === 'receive' && (
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#a855f7" strokeWidth="2"><polyline points="23 6 13.5 15.5 8.5 10.5 1 18" /><polyline points="17 6 23 6 23 12" /></svg>
                  )}
                  {card.icon === 'delivery' && (
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#f97316" strokeWidth="2"><polyline points="23 18 13.5 8.5 8.5 13.5 1 6" /><polyline points="17 18 23 18 23 12" /></svg>
                  )}
                </div>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      {/* Warehouse Overview */}
      <Card>
        <CardContent className="p-6">
          <h3 className="mb-4 text-sm font-medium text-gray-700">仓库概览</h3>
          <div className="space-y-4">
            {warehouseList.map((wh) => (
              <div key={wh.code} className="flex items-center gap-4">
                <div className="w-32 text-sm font-medium text-gray-700">{wh.name}</div>
                <div className="flex-1">
                  <div className="mb-1 flex items-center justify-between text-xs">
                    <span className="text-gray-500">{wh.code}</span>
                    <span className="text-gray-500">
                      {formatNumber(wh.used)} / {formatNumber(wh.capacity)}
                    </span>
                  </div>
                  <div className="h-2 w-full rounded-full bg-gray-100">
                    <div
                      className={`h-2 rounded-full transition-all ${
                        wh.used / wh.capacity > 0.8
                          ? 'bg-red-500'
                          : wh.used / wh.capacity > 0.6
                          ? 'bg-yellow-500'
                          : 'bg-green-500'
                      }`}
                      style={{ width: `${(wh.used / wh.capacity) * 100}%` }}
                    />
                  </div>
                </div>
                <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${
                  wh.status === '库存偏高' ? 'bg-yellow-100 text-yellow-700' : 'bg-green-100 text-green-700'
                }`}>
                  {wh.status}
                </span>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>

      {/* Recent Receives */}
      <Card>
        <CardContent className="p-6">
          <h3 className="mb-4 text-sm font-medium text-gray-700">最近收货</h3>
          <div className="space-y-2">
            {recentReceives.map((r) => (
              <div key={r.no} className="flex items-center justify-between rounded-lg border p-3">
                <div className="flex items-center gap-4">
                  <span className="text-sm font-medium text-blue-600">{r.no}</span>
                  <span className="text-sm text-gray-500">{r.supplier}</span>
                  <span className="text-sm text-gray-700">{r.material}</span>
                </div>
                <div className="flex items-center gap-4">
                  <span className="text-sm font-medium">{r.qty} 件</span>
                  <span className="text-xs text-gray-400">{r.date}</span>
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
