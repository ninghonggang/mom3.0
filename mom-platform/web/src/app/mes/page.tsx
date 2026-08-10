'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select } from '@/components/ui/select';
import { Badge } from '@/components/ui/badge';
import { Pagination } from '@/components/ui/pagination';
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from '@/components/ui/table';
import { formatDate, formatNumber, getStatusLabel, getStatusColor, ORDER_STATUS_MAP } from '@/lib/utils';
import type { ProductionOrder } from '@/types';

const mockOrders: ProductionOrder[] = [
  { id: '1', orderNo: 'MO-2024-0001', materialCode: 'MAT-001', materialName: '主轴轴承座', spec: 'Φ120×80', quantity: 500, completedQty: 320, unit: '件', status: 'IN_PROGRESS', planStartDate: '2025-01-06', planEndDate: '2025-01-10', workCenter: 'CNC-01', priority: 1, bomVersion: 'BOM-V2.1', routingVersion: 'R-V1.3', createdAt: '2025-01-05' },
  { id: '2', orderNo: 'MO-2024-0002', materialCode: 'MAT-002', materialName: '电机法兰盘', spec: 'Φ200×15', quantity: 300, completedQty: 300, unit: '件', status: 'COMPLETED', planStartDate: '2025-01-05', planEndDate: '2025-01-08', actualStartDate: '2025-01-05', actualEndDate: '2025-01-08', workCenter: 'CNC-02', priority: 2, bomVersion: 'BOM-V2.0', routingVersion: 'R-V1.2', createdAt: '2025-01-04' },
  { id: '3', orderNo: 'MO-2024-0003', materialCode: 'MAT-003', materialName: '气缸端盖', spec: 'M80×2', quantity: 800, completedQty: 0, unit: '件', status: 'RELEASED', planStartDate: '2025-01-12', planEndDate: '2025-01-16', workCenter: 'CNC-03', priority: 1, bomVersion: 'BOM-V3.0', routingVersion: 'R-V2.1', createdAt: '2025-01-07' },
  { id: '4', orderNo: 'MO-2024-0004', materialCode: 'MAT-004', materialName: '传动齿轮轴', spec: 'Φ45×200', quantity: 200, completedQty: 200, unit: '件', status: 'CLOSED', planStartDate: '2025-01-02', planEndDate: '2025-01-06', actualStartDate: '2025-01-02', actualEndDate: '2025-01-06', workCenter: 'CNC-01', priority: 3, bomVersion: 'BOM-V1.5', routingVersion: 'R-V1.0', createdAt: '2025-01-01' },
  { id: '5', orderNo: 'MO-2024-0005', materialCode: 'MAT-005', materialName: '液压阀体', spec: '120×80×60', quantity: 150, completedQty: 0, unit: '件', status: 'DRAFT', planStartDate: '2025-01-15', planEndDate: '2025-01-20', workCenter: 'CNC-04', priority: 2, bomVersion: 'BOM-V4.0', routingVersion: 'R-V3.0', createdAt: '2025-01-08' },
];

export default function MesPage() {
  const router = useRouter();
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [page, setPage] = useState(1);

  const filteredOrders = mockOrders.filter((order) => {
    if (search && !order.orderNo.includes(search) && !order.materialName.includes(search)) return false;
    if (statusFilter && order.status !== statusFilter) return false;
    return true;
  });

  const pageSize = 10;
  const totalPages = Math.ceil(filteredOrders.length / pageSize);

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">生产订单管理</h1>
        <Button onClick={() => router.push('/mes')}>
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" />
          </svg>
          新建订单
        </Button>
      </div>

      {/* Filters */}
      <div className="filter-bar">
        <div className="flex items-center gap-2">
          <span className="filter-label">搜索</span>
          <Input
            placeholder="订单号 / 物料名称"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-56"
          />
        </div>
        <div className="flex items-center gap-2">
          <span className="filter-label">状态</span>
          <Select
            value={statusFilter}
            onChange={(e) => setStatusFilter(e.target.value)}
            className="w-32"
          >
            <option value="">全部</option>
            <option value="DRAFT">草稿</option>
            <option value="RELEASED">已下达</option>
            <option value="IN_PROGRESS">生产中</option>
            <option value="COMPLETED">已完成</option>
            <option value="CLOSED">已关闭</option>
            <option value="ON_HOLD">暂停</option>
          </Select>
        </div>
        <Button variant="outline" size="sm" onClick={() => { setSearch(''); setStatusFilter(''); }}>
          重置
        </Button>
      </div>

      {/* Table */}
      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-40">订单号</TableHead>
              <TableHead>物料编码</TableHead>
              <TableHead>物料名称</TableHead>
              <TableHead className="text-right">订单数量</TableHead>
              <TableHead className="text-right">已完成</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>计划开始</TableHead>
              <TableHead>计划结束</TableHead>
              <TableHead>工作中心</TableHead>
              <TableHead className="text-right">操作</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredOrders.map((order) => (
              <TableRow
                key={order.id}
                className="cursor-pointer"
                onClick={() => router.push(`/mes/${order.id}`)}
              >
                <TableCell className="font-medium text-blue-600">
                  {order.orderNo}
                </TableCell>
                <TableCell className="text-gray-500">{order.materialCode}</TableCell>
                <TableCell>{order.materialName}</TableCell>
                <TableCell className="text-right">{formatNumber(order.quantity)}</TableCell>
                <TableCell className="text-right">
                  <div className="flex items-center justify-end gap-2">
                    <div className="h-1.5 w-16 rounded-full bg-gray-200">
                      <div
                        className="h-1.5 rounded-full bg-blue-500"
                        style={{ width: `${(order.completedQty / order.quantity) * 100}%` }}
                      />
                    </div>
                    <span className="text-xs text-gray-500">{order.completedQty}</span>
                  </div>
                </TableCell>
                <TableCell>
                  <span
                    className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${
                      getStatusColor(ORDER_STATUS_MAP, order.status)
                    }`}
                  >
                    {getStatusLabel(ORDER_STATUS_MAP, order.status)}
                  </span>
                </TableCell>
                <TableCell className="text-gray-500">{formatDate(order.planStartDate)}</TableCell>
                <TableCell className="text-gray-500">{formatDate(order.planEndDate)}</TableCell>
                <TableCell className="text-gray-500">{order.workCenter}</TableCell>
                <TableCell className="text-right" onClick={(e) => e.stopPropagation()}>
                  <div className="flex items-center justify-end gap-1">
                    <Button variant="ghost" size="sm">编辑</Button>
                    <Button variant="ghost" size="sm">派工</Button>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <Pagination page={page} totalPages={totalPages} onPageChange={setPage} />
      </div>
    </div>
  );
}
