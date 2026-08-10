'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select } from '@/components/ui/select';
import { Badge } from '@/components/ui/badge';
import { Pagination } from '@/components/ui/pagination';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatDate, formatNumber } from '@/lib/utils';
import type { InspectionSheet, InspectionType } from '@/types';

const types: { value: InspectionType; label: string }[] = [
  { value: 'IQC', label: 'IQC 来料检验' },
  { value: 'IPQC', label: 'IPQC 过程检验' },
  { value: 'FQC', label: 'FQC 完工检验' },
  { value: 'OQC', label: 'OQC 出货检验' },
];

const mockSheets: InspectionSheet[] = [
  { id: '1', sheetNo: 'IQC-2025-0001', type: 'IQC', orderNo: 'PO-240101', materialCode: 'MAT-001', materialName: '主轴轴承座', supplier: '供应商A', lotNo: 'LOT-20250101', batchQty: 500, sampleQty: 50, status: 'COMPLETED', result: 'PASS', aqlLevel: 'AQL 1.0', inspector: '质检员A', startTime: '2025-01-06', endTime: '2025-01-06', createdAt: '2025-01-05' },
  { id: '2', sheetNo: 'IPQC-2025-0001', type: 'IPQC', orderNo: 'MO-2024-0001', materialCode: 'MAT-001', materialName: '主轴轴承座', lotNo: 'LOT-20250102', batchQty: 320, sampleQty: 32, status: 'IN_PROGRESS', result: 'PENDING', aqlLevel: 'AQL 1.5', inspector: '质检员B', startTime: '2025-01-07', createdAt: '2025-01-07' },
  { id: '3', sheetNo: 'FQC-2025-0001', type: 'FQC', orderNo: 'MO-2024-0002', materialCode: 'MAT-002', materialName: '电机法兰盘', lotNo: 'LOT-20250103', batchQty: 300, sampleQty: 30, status: 'PENDING', result: 'PENDING', aqlLevel: 'AQL 1.0', createdAt: '2025-01-08' },
  { id: '4', sheetNo: 'OQC-2025-0001', type: 'OQC', orderNo: 'SO-240101', materialCode: 'MAT-003', materialName: '气缸端盖', supplier: '供应商B', lotNo: 'LOT-20250104', batchQty: 200, sampleQty: 20, status: 'COMPLETED', result: 'FAIL', aqlLevel: 'AQL 1.0', inspector: '质检员A', startTime: '2025-01-09', endTime: '2025-01-09', createdAt: '2025-01-08' },
];

const resultColors: Record<string, string> = {
  PASS: 'bg-green-100 text-green-700',
  FAIL: 'bg-red-100 text-red-700',
  PENDING: 'bg-gray-100 text-gray-500',
};

const statusColors: Record<string, string> = {
  PENDING: 'bg-gray-100 text-gray-600',
  IN_PROGRESS: 'bg-blue-100 text-blue-700',
  COMPLETED: 'bg-green-100 text-green-700',
  CLOSED: 'bg-gray-200 text-gray-500',
};

const statusLabels: Record<string, string> = {
  PENDING: '待检验',
  IN_PROGRESS: '检验中',
  COMPLETED: '已完成',
  CLOSED: '已关闭',
};

export default function QmsPage() {
  const router = useRouter();
  const [search, setSearch] = useState('');

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">检验单管理</h1>
        <Button>
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" /></svg>
          新建检验单
        </Button>
      </div>

      <div className="filter-bar">
        <Input
          placeholder="搜索检验单号 / 物料名称"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-64"
        />
        <Button variant="outline" size="sm" onClick={() => setSearch('')}>重置</Button>
      </div>

      <Tabs defaultValue="IQC">
        <TabsList>
          {types.map((t) => (
            <TabsTrigger key={t.value} value={t.value}>{t.label}</TabsTrigger>
          ))}
        </TabsList>

        {types.map((t) => {
          const filtered = mockSheets.filter(
            (s) => s.type === t.value && (!search || s.sheetNo.includes(search) || s.materialName.includes(search))
          );
          return (
            <TabsContent key={t.value} value={t.value}>
              <div className="rounded-lg border bg-white">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>检验单号</TableHead>
                      <TableHead>订单号</TableHead>
                      <TableHead>物料名称</TableHead>
                      {t.value === 'IQC' || t.value === 'OQC' ? <TableHead>供应商</TableHead> : null}
                      <TableHead>批号</TableHead>
                      <TableHead className="text-right">批量</TableHead>
                      <TableHead className="text-right">抽检数</TableHead>
                      <TableHead>AQL</TableHead>
                      <TableHead>检验员</TableHead>
                      <TableHead>检验日期</TableHead>
                      <TableHead>状态</TableHead>
                      <TableHead>结果</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filtered.map((sheet) => (
                      <TableRow
                        key={sheet.id}
                        className="cursor-pointer hover:bg-gray-50"
                        onClick={() => router.push(`/qms/${sheet.id}`)}
                      >
                        <TableCell className="font-medium text-blue-600">{sheet.sheetNo}</TableCell>
                        <TableCell className="text-gray-500">{sheet.orderNo}</TableCell>
                        <TableCell>{sheet.materialName}</TableCell>
                        {(t.value === 'IQC' || t.value === 'OQC') && (
                          <TableCell className="text-gray-500">{sheet.supplier || '-'}</TableCell>
                        )}
                        <TableCell className="text-gray-500">{sheet.lotNo}</TableCell>
                        <TableCell className="text-right">{formatNumber(sheet.batchQty)}</TableCell>
                        <TableCell className="text-right">{sheet.sampleQty}</TableCell>
                        <TableCell className="text-gray-500">{sheet.aqlLevel || '-'}</TableCell>
                        <TableCell className="text-gray-500">{sheet.inspector || '-'}</TableCell>
                        <TableCell className="text-gray-500">{formatDate(sheet.startTime) || '-'}</TableCell>
                        <TableCell>
                          <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${statusColors[sheet.status]}`}>
                            {statusLabels[sheet.status]}
                          </span>
                        </TableCell>
                        <TableCell>
                          <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${resultColors[sheet.result]}`}>
                            {sheet.result === 'PASS' ? '合格' : sheet.result === 'FAIL' ? '不合格' : '待判定'}
                          </span>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>
            </TabsContent>
          );
        })}
      </Tabs>
    </div>
  );
}
