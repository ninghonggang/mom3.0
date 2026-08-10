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
import { formatDate, formatNumber } from '@/lib/utils';
import type { DeliveryOrder } from '@/types';

const mockDeliveries: DeliveryOrder[] = [
  { id: '1', deliveryNo: 'DEL-2025-0101', orderNo: 'SO-240101', customer: '客户A', materialCode: 'MAT-003', materialName: '气缸端盖', quantity: 200, unit: '件', status: 'SHIPPED', planDate: '2025-01-08', actualDate: '2025-01-08' },
  { id: '2', deliveryNo: 'DEL-2025-0102', orderNo: 'SO-240102', customer: '客户B', materialCode: 'MAT-004', materialName: '传动齿轮轴', quantity: 150, unit: '件', status: 'PICKING', planDate: '2025-01-10' },
  { id: '3', deliveryNo: 'DEL-2025-0103', orderNo: 'SO-240103', customer: '客户A', materialCode: 'MAT-002', materialName: '电机法兰盘', quantity: 300, unit: '件', status: 'PENDING', planDate: '2025-01-15' },
  { id: '4', deliveryNo: 'DEL-2025-0104', orderNo: 'SO-240104', customer: '客户C', materialCode: 'MAT-001', materialName: '主轴轴承座', quantity: 500, unit: '件', status: 'SHIPPED', planDate: '2025-01-06', actualDate: '2025-01-06' },
];

const statusLabels: Record<string, string> = {
  PENDING: '待发货',
  PICKING: '拣货中',
  PACKED: '已打包',
  SHIPPED: '已发货',
  DELIVERED: '已签收',
};

const statusColors: Record<string, string> = {
  PENDING: 'bg-gray-100 text-gray-600',
  PICKING: 'bg-blue-100 text-blue-700',
  PACKED: 'bg-purple-100 text-purple-700',
  SHIPPED: 'bg-green-100 text-green-700',
  DELIVERED: 'bg-gray-200 text-gray-500',
};

export default function DeliveryPage() {
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [page, setPage] = useState(1);

  const filtered = mockDeliveries.filter((d) => {
    if (search && !d.deliveryNo.includes(search) && !d.customer.includes(search) && !d.materialName.includes(search)) return false;
    if (statusFilter && d.status !== statusFilter) return false;
    return true;
  });

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">发货单</h1>
        <Button>新建发货</Button>
      </div>

      <div className="filter-bar">
        <Input placeholder="搜索发货单号 / 客户 / 物料" value={search} onChange={(e) => setSearch(e.target.value)} className="w-64" />
        <Select value={statusFilter} onChange={(e) => setStatusFilter(e.target.value)} className="w-28">
          <option value="">全部状态</option>
          <option value="PENDING">待发货</option>
          <option value="PICKING">拣货中</option>
          <option value="PACKED">已打包</option>
          <option value="SHIPPED">已发货</option>
          <option value="DELIVERED">已签收</option>
        </Select>
        <Button variant="outline" size="sm" onClick={() => { setSearch(''); setStatusFilter(''); }}>重置</Button>
      </div>

      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>发货单号</TableHead>
              <TableHead>销售订单</TableHead>
              <TableHead>客户</TableHead>
              <TableHead>物料名称</TableHead>
              <TableHead className="text-right">数量</TableHead>
              <TableHead>单位</TableHead>
              <TableHead>计划日期</TableHead>
              <TableHead>实际日期</TableHead>
              <TableHead>状态</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((d) => (
              <TableRow key={d.id}>
                <TableCell className="font-medium text-blue-600">{d.deliveryNo}</TableCell>
                <TableCell className="text-gray-500">{d.orderNo}</TableCell>
                <TableCell>{d.customer}</TableCell>
                <TableCell>{d.materialName}</TableCell>
                <TableCell className="text-right">{formatNumber(d.quantity)}</TableCell>
                <TableCell>{d.unit}</TableCell>
                <TableCell className="text-gray-500">{formatDate(d.planDate)}</TableCell>
                <TableCell className="text-gray-500">{formatDate(d.actualDate)}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${statusColors[d.status]}`}>
                    {statusLabels[d.status]}
                  </span>
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
