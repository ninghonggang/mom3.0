'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatNumber, formatDate } from '@/lib/utils';
import type { MpsPlan } from '@/types';

const mockPlans: MpsPlan[] = [
  { id: '1', planNo: 'MPS-2025-W02', materialCode: 'MAT-001', materialName: '主轴轴承座', quantity: 500, startDate: '2025-01-06', endDate: '2025-01-10', status: 'IN_PROGRESS', frozen: true },
  { id: '2', planNo: 'MPS-2025-W02', materialCode: 'MAT-002', materialName: '电机法兰盘', quantity: 300, startDate: '2025-01-05', endDate: '2025-01-08', status: 'COMPLETED', frozen: true },
  { id: '3', planNo: 'MPS-2025-W03', materialCode: 'MAT-003', materialName: '气缸端盖', quantity: 800, startDate: '2025-01-12', endDate: '2025-01-16', status: 'PLANNED', frozen: false },
  { id: '4', planNo: 'MPS-2025-W03', materialCode: 'MAT-005', materialName: '液压阀体', quantity: 150, startDate: '2025-01-15', endDate: '2025-01-20', status: 'PLANNED', frozen: false },
  { id: '5', planNo: 'MPS-2025-W04', materialCode: 'MAT-006', materialName: '滚珠丝杠', quantity: 100, startDate: '2025-01-22', endDate: '2025-01-28', status: 'DRAFT', frozen: false },
];

const statusLabels: Record<string, string> = {
  DRAFT: '草稿',
  PLANNED: '已排程',
  RELEASED: '已下达',
  IN_PROGRESS: '执行中',
  COMPLETED: '已完成',
};

const statusColors: Record<string, string> = {
  DRAFT: 'bg-gray-100 text-gray-600',
  PLANNED: 'bg-blue-100 text-blue-700',
  RELEASED: 'bg-purple-100 text-purple-700',
  IN_PROGRESS: 'bg-yellow-100 text-yellow-700',
  COMPLETED: 'bg-green-100 text-green-700',
};

export default function ApsPage() {
  const [search, setSearch] = useState('');

  const filtered = mockPlans.filter(
    (p) => !search || p.planNo.includes(search) || p.materialName.includes(search)
  );

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">MPS 主生产计划</h1>
        <div className="flex items-center gap-2">
          <Button variant="outline" size="sm">冻结/解冻</Button>
          <Button size="sm">新建计划</Button>
        </div>
      </div>

      <div className="filter-bar">
        <Input placeholder="搜索计划编号 / 物料名称" value={search} onChange={(e) => setSearch(e.target.value)} className="w-56" />
        <Button variant="outline" size="sm" onClick={() => setSearch('')}>重置</Button>
      </div>

      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>计划编号</TableHead>
              <TableHead>物料编码</TableHead>
              <TableHead>物料名称</TableHead>
              <TableHead className="text-right">数量</TableHead>
              <TableHead>开始日期</TableHead>
              <TableHead>结束日期</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>冻结</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((plan) => (
              <TableRow key={plan.id}>
                <TableCell className="font-medium text-blue-600">{plan.planNo}</TableCell>
                <TableCell className="text-gray-500">{plan.materialCode}</TableCell>
                <TableCell>{plan.materialName}</TableCell>
                <TableCell className="text-right">{formatNumber(plan.quantity)}</TableCell>
                <TableCell className="text-gray-500">{formatDate(plan.startDate)}</TableCell>
                <TableCell className="text-gray-500">{formatDate(plan.endDate)}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${statusColors[plan.status]}`}>
                    {statusLabels[plan.status]}
                  </span>
                </TableCell>
                <TableCell>
                  {plan.frozen ? (
                    <Badge variant="warning">已冻结</Badge>
                  ) : (
                    <span className="text-gray-400">-</span>
                  )}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
