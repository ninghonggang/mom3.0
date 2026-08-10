'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select } from '@/components/ui/select';
import { Pagination } from '@/components/ui/pagination';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatDateTime, getStatusLabel, getStatusColor, REPAIR_STATUS_MAP } from '@/lib/utils';
import type { RepairOrder } from '@/types';

const mockRepairs: RepairOrder[] = [
  { id: '1', repairNo: 'REP-2025-0001', equipmentId: 'EQ-004', equipmentName: '三坐标测量机', faultDesc: 'Z轴测量精度超差', faultType: '精度', severity: 'CRITICAL', status: 'IN_PROGRESS', reporter: '质检员A', reportedAt: '2025-01-07 09:00', technician: '维修员A', startTime: '2025-01-07 10:00', downtime: 4 },
  { id: '2', repairNo: 'REP-2025-0002', equipmentId: 'EQ-002', equipmentName: 'CNC数控加工中心 #02', faultDesc: '换刀臂卡滞', faultType: '机械', severity: 'MAJOR', status: 'ASSIGNED', reporter: '张操作', reportedAt: '2025-01-06 15:00', technician: '维修员B', downtime: 0 },
  { id: '3', repairNo: 'REP-2025-0003', equipmentId: 'EQ-003', equipmentName: '全自动钻孔攻牙机', faultDesc: '主轴转速异常波动', faultType: '电气', severity: 'MAJOR', status: 'COMPLETED', reporter: '李操作', reportedAt: '2025-01-04 08:00', technician: '维修员A', startTime: '2025-01-04 08:30', endTime: '2025-01-04 14:00', rootCause: '变频器参数偏移', solution: '重新校准变频器参数', downtime: 5.5 },
  { id: '4', repairNo: 'REP-2025-0004', equipmentId: 'EQ-005', equipmentName: '液压机 100T', faultDesc: '液压油泄漏', faultType: '液压', severity: 'CRITICAL', status: 'REPORTED', reporter: '王操作', reportedAt: '2025-01-08 07:30', downtime: 0 },
  { id: '5', repairNo: 'REP-2024-056', equipmentId: 'EQ-001', equipmentName: 'CNC数控加工中心 #01', faultDesc: '主轴异响', faultType: '机械', severity: 'MAJOR', status: 'CONFIRMED', reporter: '张操作', reportedAt: '2024-11-10 14:00', technician: '维修员B', startTime: '2024-11-10 15:00', endTime: '2024-11-11 07:00', rootCause: '主轴轴承磨损', solution: '更换主轴轴承', downtime: 16 },
];

const severityLabels: Record<string, string> = { CRITICAL: '严重', MAJOR: '主要', MINOR: '次要' };
const severityColors: Record<string, string> = { CRITICAL: 'bg-red-100 text-red-700', MAJOR: 'bg-orange-100 text-orange-700', MINOR: 'bg-yellow-100 text-yellow-700' };

export default function RepairsPage() {
  const router = useRouter();
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [page, setPage] = useState(1);

  const filtered = mockRepairs.filter((r) => {
    if (search && !r.repairNo.includes(search) && !r.equipmentName.includes(search)) return false;
    if (statusFilter && r.status !== statusFilter) return false;
    return true;
  });

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">维修工单</h1>
        <Button>新建维修</Button>
      </div>

      <div className="filter-bar">
        <Input placeholder="搜索维修编号 / 设备名称" value={search} onChange={(e) => setSearch(e.target.value)} className="w-56" />
        <Select value={statusFilter} onChange={(e) => setStatusFilter(e.target.value)} className="w-28">
          <option value="">全部状态</option>
          <option value="REPORTED">已报修</option>
          <option value="ASSIGNED">已分配</option>
          <option value="IN_PROGRESS">维修中</option>
          <option value="COMPLETED">已修复</option>
          <option value="CONFIRMED">已确认</option>
        </Select>
        <Button variant="outline" size="sm" onClick={() => { setSearch(''); setStatusFilter(''); }}>重置</Button>
      </div>

      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>维修编号</TableHead>
              <TableHead>设备名称</TableHead>
              <TableHead>故障描述</TableHead>
              <TableHead>故障类型</TableHead>
              <TableHead>严重级别</TableHead>
              <TableHead>报修人</TableHead>
              <TableHead>报修时间</TableHead>
              <TableHead>维修员</TableHead>
              <TableHead className="text-right">停机(h)</TableHead>
              <TableHead>状态</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((r) => (
              <TableRow
                key={r.id}
                className="cursor-pointer hover:bg-gray-50"
                onClick={() => router.push(`/eam/repairs/${r.id}`)}
              >
                <TableCell className="font-medium text-blue-600">{r.repairNo}</TableCell>
                <TableCell>{r.equipmentName}</TableCell>
                <TableCell className="max-w-xs truncate text-gray-600">{r.faultDesc}</TableCell>
                <TableCell className="text-gray-500">{r.faultType}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${severityColors[r.severity]}`}>
                    {severityLabels[r.severity]}
                  </span>
                </TableCell>
                <TableCell className="text-gray-500">{r.reporter}</TableCell>
                <TableCell className="text-gray-500">{formatDateTime(r.reportedAt)}</TableCell>
                <TableCell className="text-gray-500">{r.technician || '-'}</TableCell>
                <TableCell className="text-right text-red-600">{r.downtime}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${getStatusColor(REPAIR_STATUS_MAP, r.status)}`}>
                    {getStatusLabel(REPAIR_STATUS_MAP, r.status)}
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
