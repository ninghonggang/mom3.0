'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent } from '@/components/ui/card';
import { formatDateTime, getStatusLabel, getStatusColor, ANDON_STATUS_MAP } from '@/lib/utils';
import type { AndonCall } from '@/types';

const mockCalls: AndonCall[] = [
  { id: '1', callNo: 'AND-2025-0001', workCenter: 'CNC-01', machineId: 'MC-101', type: 'EQUIPMENT', severity: 'CRITICAL', description: '主轴温度异常升至75°C，超出正常范围', status: 'ACTIVE', caller: '张操作', calledAt: '2025-01-08 14:30' },
  { id: '2', callNo: 'AND-2025-0002', workCenter: 'QC-01', type: 'QUALITY', severity: 'WARNING', description: 'IPQC 表面粗糙度检测连续3件不合格', status: 'ACTIVE', caller: '质检员B', calledAt: '2025-01-08 14:00' },
  { id: '3', callNo: 'AND-2025-0003', workCenter: 'CNC-02', type: 'MATERIAL', severity: 'WARNING', description: '原料 LOT-20250102 库存不足，仅剩50件', status: 'RESPONDED', caller: '李操作', calledAt: '2025-01-08 13:15', respondedAt: '2025-01-08 13:20', responder: '仓管员' },
  { id: '4', callNo: 'AND-2025-0004', workCenter: 'PRESS-01', type: 'SAFETY', severity: 'CRITICAL', description: '安全光幕被遮挡，设备紧急停止', status: 'RESOLVED', caller: '王操作', calledAt: '2025-01-08 10:00', respondedAt: '2025-01-08 10:05', resolvedAt: '2025-01-08 10:20', responder: '安全员' },
];

const typeLabels: Record<string, string> = {
  MATERIAL: '物料',
  EQUIPMENT: '设备',
  QUALITY: '质量',
  SAFETY: '安全',
  PROCESS: '工艺',
};

const typeColors: Record<string, string> = {
  MATERIAL: 'bg-purple-100 text-purple-700',
  EQUIPMENT: 'bg-blue-100 text-blue-700',
  QUALITY: 'bg-yellow-100 text-yellow-700',
  SAFETY: 'bg-red-100 text-red-700',
  PROCESS: 'bg-green-100 text-green-700',
};

const severityColors: Record<string, string> = {
  CRITICAL: 'bg-red-100 text-red-700',
  WARNING: 'bg-yellow-100 text-yellow-700',
  INFO: 'bg-blue-100 text-blue-700',
};

export default function AndonPage() {
  const [filter, setFilter] = useState('ALL');

  const filtered = filter === 'ALL' ? mockCalls : mockCalls.filter((c) => c.status === filter);

  return (
    <div className="space-y-6">
      <div className="page-header">
        <h1 className="page-title">安灯系统</h1>
        <div className="flex items-center gap-2">
          <span className="text-sm text-gray-500">
            活跃: {mockCalls.filter(c => c.status === 'ACTIVE').length} 个
          </span>
          <Button size="sm">发起呼叫</Button>
        </div>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-2 gap-4 lg:grid-cols-4">
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">活跃</p><p className="text-xl font-bold text-red-600">{mockCalls.filter(c => c.status === 'ACTIVE').length}</p></CardContent></Card>
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">已响应</p><p className="text-xl font-bold text-yellow-600">{mockCalls.filter(c => c.status === 'RESPONDED').length}</p></CardContent></Card>
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">已解决</p><p className="text-xl font-bold text-green-600">{mockCalls.filter(c => c.status === 'RESOLVED').length}</p></CardContent></Card>
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">超时</p><p className="text-xl font-bold text-orange-600">{mockCalls.filter(c => c.status === 'TIMEOUT').length}</p></CardContent></Card>
      </div>

      {/* Filter Tabs */}
      <div className="inline-flex rounded-lg bg-gray-100 p-1">
        {[
          { value: 'ALL', label: '全部' },
          { value: 'ACTIVE', label: '活跃' },
          { value: 'RESPONDED', label: '已响应' },
          { value: 'RESOLVED', label: '已解决' },
        ].map((f) => (
          <button
            key={f.value}
            onClick={() => setFilter(f.value)}
            className={`rounded-md px-3 py-1 text-sm font-medium transition-all ${
              filter === f.value ? 'bg-white text-gray-900 shadow-sm' : 'text-gray-500'
            }`}
          >
            {f.label}
          </button>
        ))}
      </div>

      {/* Andon Cards */}
      <div className="space-y-3">
        {filtered.map((call) => (
          <Card key={call.id} className={call.status === 'ACTIVE' ? 'border-l-4 border-l-red-500' : ''}>
            <CardContent className="p-4">
              <div className="flex items-start justify-between">
                <div className="flex items-start gap-4">
                  {/* Severity Indicator */}
                  <div className={`mt-0.5 h-3 w-3 rounded-full flex-shrink-0 ${
                    call.severity === 'CRITICAL' ? 'bg-red-500 animate-pulse' :
                    call.severity === 'WARNING' ? 'bg-yellow-500' : 'bg-blue-500'
                  }`} />

                  <div>
                    <div className="flex items-center gap-2 mb-1">
                      <span className="font-medium text-sm">{call.callNo}</span>
                      <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${typeColors[call.type]}`}>
                        {typeLabels[call.type]}
                      </span>
                      <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${severityColors[call.severity]}`}>
                        {call.severity === 'CRITICAL' ? '严重' : call.severity === 'WARNING' ? '警告' : '信息'}
                      </span>
                    </div>
                    <p className="text-sm text-gray-700">{call.description}</p>
                    <div className="mt-2 flex items-center gap-4 text-xs text-gray-400">
                      <span>工位: {call.workCenter}</span>
                      {call.machineId && <span>设备: {call.machineId}</span>}
                      <span>呼叫: {call.caller}</span>
                      <span>{formatDateTime(call.calledAt)}</span>
                      {call.responder && <span>响应: {call.responder} ({formatDateTime(call.respondedAt!)})</span>}
                    </div>
                  </div>
                </div>

                <div className="flex items-center gap-2">
                  <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${getStatusColor(ANDON_STATUS_MAP, call.status)}`}>
                    {getStatusLabel(ANDON_STATUS_MAP, call.status)}
                  </span>
                  {call.status === 'ACTIVE' && (
                    <Button size="sm" variant="outline">响应</Button>
                  )}
                </div>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
