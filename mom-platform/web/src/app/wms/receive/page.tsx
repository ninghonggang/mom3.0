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
import type { ReceiveOrder } from '@/types';

const mockReceives: ReceiveOrder[] = [
  { id: '1', receiveNo: 'REC-2025-0101', orderNo: 'PO-240101', supplier: '供应商A', materialCode: 'MAT-001', materialName: '主轴轴承座', expectedQty: 500, receivedQty: 500, unit: '件', status: 'COMPLETED', receiveDate: '2025-01-06', inspector: '质检员A', inspectionResult: 'PASS' },
  { id: '2', receiveNo: 'REC-2025-0102', orderNo: 'PO-240102', supplier: '供应商B', materialCode: 'MAT-002', materialName: '电机法兰盘', expectedQty: 300, receivedQty: 280, unit: '件', status: 'PARTIAL', receiveDate: '2025-01-07', inspector: '质检员B' },
  { id: '3', receiveNo: 'REC-2025-0103', orderNo: 'PO-240103', supplier: '供应商A', materialCode: 'MAT-005', materialName: '液压阀体', expectedQty: 200, receivedQty: 0, unit: '件', status: 'PENDING', receiveDate: '2025-01-10' },
  { id: '4', receiveNo: 'REC-2025-0104', orderNo: 'PO-240104', supplier: '供应商C', materialCode: 'MAT-006', materialName: '滚珠丝杠', expectedQty: 100, receivedQty: 100, unit: '根', status: 'COMPLETED', receiveDate: '2025-01-05', inspector: '质检员A', inspectionResult: 'PASS' },
];

const statusLabels: Record<string, string> = {
  PENDING: '待收货',
  PARTIAL: '部分收货',
  COMPLETED: '已完成',
};

const statusColors: Record<string, string> = {
  PENDING: 'bg-gray-100 text-gray-600',
  PARTIAL: 'bg-yellow-100 text-yellow-700',
  COMPLETED: 'bg-green-100 text-green-700',
};

export default function ReceivePage() {
  const [search, setSearch] = useState('');
  const [page, setPage] = useState(1);

  const filtered = mockReceives.filter(
    (r) => !search || r.receiveNo.includes(search) || r.supplier.includes(search) || r.materialName.includes(search)
  );

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">收货单</h1>
        <Button>新建收货</Button>
      </div>

      <div className="filter-bar">
        <Input placeholder="搜索收货单号 / 供应商 / 物料" value={search} onChange={(e) => setSearch(e.target.value)} className="w-64" />
        <Button variant="outline" size="sm" onClick={() => setSearch('')}>重置</Button>
      </div>

      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>收货单号</TableHead>
              <TableHead>采购订单</TableHead>
              <TableHead>供应商</TableHead>
              <TableHead>物料名称</TableHead>
              <TableHead className="text-right">预计数量</TableHead>
              <TableHead className="text-right">实收数量</TableHead>
              <TableHead>单位</TableHead>
              <TableHead>收货日期</TableHead>
              <TableHead>检验员</TableHead>
              <TableHead>检验结果</TableHead>
              <TableHead>状态</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((r) => (
              <TableRow key={r.id}>
                <TableCell className="font-medium text-blue-600">{r.receiveNo}</TableCell>
                <TableCell className="text-gray-500">{r.orderNo}</TableCell>
                <TableCell>{r.supplier}</TableCell>
                <TableCell>{r.materialName}</TableCell>
                <TableCell className="text-right">{formatNumber(r.expectedQty)}</TableCell>
                <TableCell className="text-right font-medium">{formatNumber(r.receivedQty)}</TableCell>
                <TableCell>{r.unit}</TableCell>
                <TableCell className="text-gray-500">{formatDate(r.receiveDate)}</TableCell>
                <TableCell className="text-gray-500">{r.inspector || '-'}</TableCell>
                <TableCell>
                  {r.inspectionResult ? (
                    <Badge variant={r.inspectionResult === 'PASS' ? 'success' : 'destructive'}>
                      {r.inspectionResult === 'PASS' ? '合格' : '不合格'}
                    </Badge>
                  ) : '-'}
                </TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${statusColors[r.status]}`}>
                    {statusLabels[r.status]}
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
