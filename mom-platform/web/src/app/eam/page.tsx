'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Card, CardContent } from '@/components/ui/card';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { getStatusLabel, getStatusColor, EQUIPMENT_STATUS_MAP, formatPercent } from '@/lib/utils';
import type { Equipment } from '@/types';

const mockEquipment: Equipment[] = [
  { id: '1', equipmentNo: 'EQ-001', name: 'CNC数控加工中心 #01', model: 'DMG MORI CMX 1100V', type: 'CNC', workCenter: 'CNC-01', status: 'RUNNING', oee: 0.88, availability: 0.95, performance: 0.93, quality: 0.99, totalRunTime: 12500, totalDowntime: 320, lastMaintenance: '2025-01-02', nextMaintenance: '2025-01-15', installDate: '2020-03-15', manufacturer: 'DMG MORI' },
  { id: '2', equipmentNo: 'EQ-002', name: 'CNC数控加工中心 #02', model: 'HAAS VF-2SS', type: 'CNC', workCenter: 'CNC-02', status: 'IDLE', oee: 0.72, availability: 0.88, performance: 0.85, quality: 0.97, totalRunTime: 8900, totalDowntime: 1500, lastMaintenance: '2025-01-05', nextMaintenance: '2025-01-20', installDate: '2021-06-20', manufacturer: 'HAAS' },
  { id: '3', equipmentNo: 'EQ-003', name: '全自动钻孔攻牙机', model: 'BROTHER TC-S2Dn', type: 'DRILL', workCenter: 'DRILL-01', status: 'MAINTENANCE', oee: 0.65, availability: 0.78, performance: 0.86, quality: 0.98, totalRunTime: 6000, totalDowntime: 2400, lastMaintenance: '2025-01-01', nextMaintenance: '2025-01-10', installDate: '2019-09-01', manufacturer: 'BROTHER' },
  { id: '4', equipmentNo: 'EQ-004', name: '三坐标测量机', model: 'ZEISS CONTURA', type: 'CMM', workCenter: 'QC-01', status: 'REPAIR', oee: 0.45, availability: 0.60, performance: 0.80, quality: 0.95, totalRunTime: 3200, totalDowntime: 4800, lastMaintenance: '2024-12-20', nextMaintenance: '2025-01-08', installDate: '2018-04-12', manufacturer: 'ZEISS' },
  { id: '5', equipmentNo: 'EQ-005', name: '液压机 100T', model: 'YH-100T', type: 'PRESS', workCenter: 'PRESS-01', status: 'STOPPED', oee: 0.30, availability: 0.50, performance: 0.70, quality: 0.85, totalRunTime: 4500, totalDowntime: 5500, lastMaintenance: '2024-11-15', nextMaintenance: '2025-01-08', installDate: '2017-07-20', manufacturer: '扬力' },
];

export default function EamPage() {
  const router = useRouter();
  const [search, setSearch] = useState('');

  const filtered = mockEquipment.filter(
    (e) => !search || e.name.includes(search) || e.equipmentNo.includes(search)
  );

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">设备台账</h1>
        <Button>新增设备</Button>
      </div>

      <div className="filter-bar">
        <Input placeholder="搜索设备编号 / 名称" value={search} onChange={(e) => setSearch(e.target.value)} className="w-56" />
        <Button variant="outline" size="sm" onClick={() => setSearch('')}>重置</Button>
      </div>

      {/* Summary Cards */}
      <div className="mb-6 grid grid-cols-2 gap-4 lg:grid-cols-5">
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">总设备</p><p className="text-xl font-bold">{mockEquipment.length}</p></CardContent></Card>
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">运行中</p><p className="text-xl font-bold text-green-600">{mockEquipment.filter(e => e.status === 'RUNNING').length}</p></CardContent></Card>
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">待机</p><p className="text-xl font-bold text-gray-600">{mockEquipment.filter(e => e.status === 'IDLE').length}</p></CardContent></Card>
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">保养/维修</p><p className="text-xl font-bold text-orange-600">{mockEquipment.filter(e => e.status === 'MAINTENANCE' || e.status === 'REPAIR').length}</p></CardContent></Card>
        <Card><CardContent className="p-4 text-center"><p className="text-xs text-gray-500">停机</p><p className="text-xl font-bold text-red-600">{mockEquipment.filter(e => e.status === 'STOPPED').length}</p></CardContent></Card>
      </div>

      {/* Equipment Table */}
      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>设备编号</TableHead>
              <TableHead>设备名称</TableHead>
              <TableHead>型号</TableHead>
              <TableHead>类型</TableHead>
              <TableHead>工作中心</TableHead>
              <TableHead>状态</TableHead>
              <TableHead className="text-right">OEE</TableHead>
              <TableHead className="text-right">可用率</TableHead>
              <TableHead className="text-right">运行时间(h)</TableHead>
              <TableHead>下次保养</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((eq) => (
              <TableRow
                key={eq.id}
                className="cursor-pointer hover:bg-gray-50"
                onClick={() => router.push(`/eam/${eq.id}`)}
              >
                <TableCell className="font-medium text-blue-600">{eq.equipmentNo}</TableCell>
                <TableCell>{eq.name}</TableCell>
                <TableCell className="text-gray-500">{eq.model}</TableCell>
                <TableCell className="text-gray-500">{eq.type}</TableCell>
                <TableCell>{eq.workCenter}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${getStatusColor(EQUIPMENT_STATUS_MAP, eq.status)}`}>
                    {getStatusLabel(EQUIPMENT_STATUS_MAP, eq.status)}
                  </span>
                </TableCell>
                <TableCell className="text-right">
                  <span className={`font-medium ${eq.oee >= 0.85 ? 'text-green-600' : eq.oee >= 0.7 ? 'text-yellow-600' : 'text-red-600'}`}>
                    {formatPercent(eq.oee)}
                  </span>
                </TableCell>
                <TableCell className="text-right text-gray-600">{formatPercent(eq.availability)}</TableCell>
                <TableCell className="text-right text-gray-600">{eq.totalRunTime.toLocaleString()}</TableCell>
                <TableCell className="text-gray-500">{eq.nextMaintenance || '-'}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
