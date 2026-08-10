'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select } from '@/components/ui/select';
import { Badge } from '@/components/ui/badge';
import { Pagination } from '@/components/ui/pagination';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatDateTime } from '@/lib/utils';
import type { Alert } from '@/types';

const mockAlerts: Alert[] = [
  { id: '1', title: 'CNC-01 主轴温度过高', message: '主轴温度传感器读数75°C，超过阈值70°C。建议立即检查冷却系统。', severity: 'CRITICAL', source: 'EAM', status: 'ACTIVE', createdAt: '2025-01-08 14:30' },
  { id: '2', title: 'IPQC-02 连续3件不合格', message: '检验工位IPQC-02在30分钟内连续发现3件表面粗糙度不合格。', severity: 'WARNING', source: 'QMS', status: 'ACTIVE', createdAt: '2025-01-08 14:00' },
  { id: '3', title: '原料 LOT-20250102 库存不足', message: 'CNC-02所需原料LOT-20250102可用库存仅剩50件，低于安全库存100件。', severity: 'WARNING', source: 'WMS', status: 'ACTIVE', createdAt: '2025-01-08 13:15' },
  { id: '4', title: 'PRESS-01 安全光幕触发', message: '安全光幕被遮挡，设备已紧急停止。需现场确认后恢复运行。', severity: 'CRITICAL', source: 'EAM', status: 'RESOLVED', createdAt: '2025-01-08 10:00', acknowledgedAt: '2025-01-08 10:05' },
  { id: '5', title: 'CNC-02 换刀提醒', message: 'CNC-02当前刀具寿命已使用85%，建议提前准备备用刀具。', severity: 'INFO', source: 'EAM', status: 'ACKNOWLEDGED', createdAt: '2025-01-08 11:30', acknowledgedAt: '2025-01-08 11:35' },
  { id: '6', title: 'OEE 日指标低于目标', message: '昨日OEE值为72%，低于月度目标80%。主要影响：设备EO-003停机维修5.5小时。', severity: 'WARNING', source: 'EAM', status: 'ACKNOWLEDGED', createdAt: '2025-01-08 08:00', acknowledgedAt: '2025-01-08 08:30' },
];

const severityColors: Record<string, string> = {
  CRITICAL: 'bg-red-100 text-red-700',
  WARNING: 'bg-yellow-100 text-yellow-700',
  INFO: 'bg-blue-100 text-blue-700',
};

const statusColors: Record<string, string> = {
  ACTIVE: 'bg-red-100 text-red-700',
  ACKNOWLEDGED: 'bg-yellow-100 text-yellow-700',
  RESOLVED: 'bg-green-100 text-green-700',
};

const severityLabels: Record<string, string> = {
  CRITICAL: '严重',
  WARNING: '警告',
  INFO: '信息',
};

const statusLabels: Record<string, string> = {
  ACTIVE: '活跃',
  ACKNOWLEDGED: '已确认',
  RESOLVED: '已解决',
};

export default function AlertsPage() {
  const [search, setSearch] = useState('');
  const [severityFilter, setSeverityFilter] = useState('');
  const [page, setPage] = useState(1);

  const filtered = mockAlerts.filter((a) => {
    if (search && !a.title.includes(search) && !a.source.includes(search)) return false;
    if (severityFilter && a.severity !== severityFilter) return false;
    return true;
  });

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">告警记录</h1>
      </div>

      <div className="filter-bar">
        <Input placeholder="搜索告警标题 / 来源" value={search} onChange={(e) => setSearch(e.target.value)} className="w-56" />
        <Select value={severityFilter} onChange={(e) => setSeverityFilter(e.target.value)} className="w-24">
          <option value="">全部级别</option>
          <option value="CRITICAL">严重</option>
          <option value="WARNING">警告</option>
          <option value="INFO">信息</option>
        </Select>
        <Button variant="outline" size="sm" onClick={() => { setSearch(''); setSeverityFilter(''); }}>重置</Button>
      </div>

      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-12">级别</TableHead>
              <TableHead>标题</TableHead>
              <TableHead>消息</TableHead>
              <TableHead>来源</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>创建时间</TableHead>
              <TableHead className="text-right">操作</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((alert) => (
              <TableRow key={alert.id}>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${severityColors[alert.severity]}`}>
                    {severityLabels[alert.severity]}
                  </span>
                </TableCell>
                <TableCell className="font-medium">{alert.title}</TableCell>
                <TableCell className="max-w-md truncate text-gray-600">{alert.message}</TableCell>
                <TableCell className="text-gray-500">{alert.source}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${statusColors[alert.status]}`}>
                    {statusLabels[alert.status]}
                  </span>
                </TableCell>
                <TableCell className="text-gray-500">{formatDateTime(alert.createdAt)}</TableCell>
                <TableCell className="text-right">
                  {alert.status === 'ACTIVE' && (
                    <Button variant="ghost" size="sm">确认</Button>
                  )}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <Pagination page={page} totalPages={Math.ceil(filtered.length / 10)} onPageChange={setPage} />
      </div>
    </div>
  );
}
