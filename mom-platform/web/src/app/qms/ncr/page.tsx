'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select } from '@/components/ui/select';
import { Badge } from '@/components/ui/badge';
import { Pagination } from '@/components/ui/pagination';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatDateTime, getStatusLabel, getStatusColor, NCR_SEVERITY_MAP, NCR_STATUS_MAP } from '@/lib/utils';
import type { Ncr } from '@/types';

const mockNcrs: Ncr[] = [
  { id: '1', ncrNo: 'NCR-2025-0001', source: 'IQC-2025-0001', orderNo: 'PO-240101', materialCode: 'MAT-001', materialName: '主轴轴承座', defectDescription: '表面粗糙度超差 Ra实测1.8μm', severity: 'MAJOR', status: 'DISPOSITION', quantity: 2, inspector: '质检员A', foundAt: '2025-01-06', disposition: '返工处理' },
  { id: '2', ncrNo: 'NCR-2025-0002', source: 'IPQC-2025-0001', orderNo: 'MO-2024-0001', materialCode: 'MAT-001', materialName: '主轴轴承座', defectDescription: '外径尺寸超差 Φ120.04', severity: 'CRITICAL', status: 'OPEN', quantity: 5, inspector: '质检员B', foundAt: '2025-01-07', disposition: '' },
  { id: '3', ncrNo: 'NCR-2025-0003', source: 'FQC-2025-0001', orderNo: 'MO-2024-0002', materialCode: 'MAT-002', materialName: '电机法兰盘', defectDescription: '表面划痕', severity: 'MINOR', status: 'CLOSED', quantity: 1, inspector: '质检员A', foundAt: '2025-01-08', disposition: '让步接收', closedAt: '2025-01-08' },
  { id: '4', ncrNo: 'NCR-2025-0004', source: 'OQC-2025-0001', orderNo: 'SO-240101', materialCode: 'MAT-003', materialName: '气缸端盖', defectDescription: '硬度检测不合格', severity: 'MAJOR', status: 'IN_REWORK', quantity: 3, inspector: '质检员C', foundAt: '2025-01-09', disposition: '返工' },
];

export default function NcrListPage() {
  const router = useRouter();
  const [search, setSearch] = useState('');
  const [severityFilter, setSeverityFilter] = useState('');
  const [page, setPage] = useState(1);

  const filtered = mockNcrs.filter((n) => {
    if (search && !n.ncrNo.includes(search) && !n.materialName.includes(search)) return false;
    if (severityFilter && n.severity !== severityFilter) return false;
    return true;
  });

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">不合格品报告 (NCR)</h1>
      </div>

      <div className="filter-bar">
        <Input
          placeholder="搜索 NCR号 / 物料名称"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-56"
        />
        <Select value={severityFilter} onChange={(e) => setSeverityFilter(e.target.value)} className="w-28">
          <option value="">全部级别</option>
          <option value="CRITICAL">严重</option>
          <option value="MAJOR">主要</option>
          <option value="MINOR">次要</option>
        </Select>
        <Button variant="outline" size="sm" onClick={() => { setSearch(''); setSeverityFilter(''); }}>重置</Button>
      </div>

      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>NCR 编号</TableHead>
              <TableHead>来源</TableHead>
              <TableHead>物料名称</TableHead>
              <TableHead>缺陷描述</TableHead>
              <TableHead>严重级别</TableHead>
              <TableHead className="text-right">数量</TableHead>
              <TableHead>处置</TableHead>
              <TableHead>检验员</TableHead>
              <TableHead>发现日期</TableHead>
              <TableHead>状态</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((ncr) => (
              <TableRow
                key={ncr.id}
                className="cursor-pointer hover:bg-gray-50"
                onClick={() => router.push(`/qms/ncr/${ncr.id}`)}
              >
                <TableCell className="font-medium text-blue-600">{ncr.ncrNo}</TableCell>
                <TableCell className="text-gray-500">{ncr.source}</TableCell>
                <TableCell>{ncr.materialName}</TableCell>
                <TableCell className="max-w-xs truncate text-gray-600">{ncr.defectDescription}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${getStatusColor(NCR_SEVERITY_MAP, ncr.severity)}`}>
                    {getStatusLabel(NCR_SEVERITY_MAP, ncr.severity)}
                  </span>
                </TableCell>
                <TableCell className="text-right">{ncr.quantity}</TableCell>
                <TableCell className="text-gray-500">{ncr.disposition || '-'}</TableCell>
                <TableCell className="text-gray-500">{ncr.inspector}</TableCell>
                <TableCell className="text-gray-500">{formatDateTime(ncr.foundAt)}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${getStatusColor(NCR_STATUS_MAP, ncr.status)}`}>
                    {getStatusLabel(NCR_STATUS_MAP, ncr.status)}
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
