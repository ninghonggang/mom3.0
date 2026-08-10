'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select } from '@/components/ui/select';
import { Pagination } from '@/components/ui/pagination';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatNumber, getStatusLabel, getStatusColor, INVENTORY_STATUS_MAP } from '@/lib/utils';
import type { InventoryBalance } from '@/types';

const mockInventory: InventoryBalance[] = [
  { id: '1', materialCode: 'MAT-001', materialName: '主轴轴承座', spec: 'Φ120×80', unit: '件', warehouse: '原材料仓A', location: 'A-01-02-03', batchNo: 'LOT-20250101', quantity: 500, lockedQty: 50, availableQty: 450, status: 'NORMAL', lastUpdate: '2025-01-08' },
  { id: '2', materialCode: 'MAT-002', materialName: '电机法兰盘', spec: 'Φ200×15', unit: '件', warehouse: '原材料仓A', location: 'A-02-01-05', batchNo: 'LOT-20250102', quantity: 300, lockedQty: 0, availableQty: 300, status: 'NORMAL', lastUpdate: '2025-01-07' },
  { id: '3', materialCode: 'MAT-003', materialName: '气缸端盖', spec: 'M80×2', unit: '件', warehouse: '半成品仓B', location: 'B-01-04-01', batchNo: 'LOT-20250103', quantity: 200, lockedQty: 200, availableQty: 0, status: 'LOCKED', lastUpdate: '2025-01-08' },
  { id: '4', materialCode: 'MAT-004', materialName: '传动齿轮轴', spec: 'Φ45×200', unit: '件', warehouse: '成品仓C', location: 'C-03-02-08', batchNo: 'LOT-20250104', quantity: 150, lockedQty: 0, availableQty: 150, status: 'NORMAL', lastUpdate: '2025-01-06' },
  { id: '5', materialCode: 'MAT-005', materialName: '液压阀体', spec: '120×80×60', unit: '件', warehouse: '不合格品仓', location: 'NG-01-01', batchNo: 'LOT-20250105', quantity: 20, lockedQty: 20, availableQty: 0, status: 'ON_HOLD', lastUpdate: '2025-01-05' },
  { id: '6', materialCode: 'MAT-006', materialName: '滚珠丝杠', spec: '25×1000', unit: '根', warehouse: '辅料仓D', location: 'D-02-05-02', batchNo: 'LOT-20250106', quantity: 50, lockedQty: 5, availableQty: 45, status: 'NORMAL', lastUpdate: '2025-01-04' },
];

export default function InventoryPage() {
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [page, setPage] = useState(1);

  const filtered = mockInventory.filter((inv) => {
    if (search && !inv.materialCode.includes(search) && !inv.materialName.includes(search)) return false;
    if (statusFilter && inv.status !== statusFilter) return false;
    return true;
  });

  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">库存余额</h1>
        <Button>导出</Button>
      </div>

      <div className="filter-bar">
        <Input placeholder="搜索物料编码 / 名称" value={search} onChange={(e) => setSearch(e.target.value)} className="w-56" />
        <Select value={statusFilter} onChange={(e) => setStatusFilter(e.target.value)} className="w-28">
          <option value="">全部状态</option>
          <option value="NORMAL">正常</option>
          <option value="LOCKED">锁定</option>
          <option value="ON_HOLD">冻结</option>
          <option value="SCRAP">报废</option>
        </Select>
        <Button variant="outline" size="sm" onClick={() => { setSearch(''); setStatusFilter(''); }}>重置</Button>
      </div>

      <div className="rounded-lg border bg-white">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>物料编码</TableHead>
              <TableHead>物料名称</TableHead>
              <TableHead>规格</TableHead>
              <TableHead>单位</TableHead>
              <TableHead>仓库</TableHead>
              <TableHead>库位</TableHead>
              <TableHead>批号</TableHead>
              <TableHead className="text-right">库存数量</TableHead>
              <TableHead className="text-right">锁定数量</TableHead>
              <TableHead className="text-right">可用数量</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>最后更新</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filtered.map((inv) => (
              <TableRow key={inv.id}>
                <TableCell className="font-medium text-blue-600">{inv.materialCode}</TableCell>
                <TableCell>{inv.materialName}</TableCell>
                <TableCell className="text-gray-500">{inv.spec}</TableCell>
                <TableCell>{inv.unit}</TableCell>
                <TableCell className="text-gray-500">{inv.warehouse}</TableCell>
                <TableCell className="text-gray-500">{inv.location}</TableCell>
                <TableCell className="text-gray-500">{inv.batchNo}</TableCell>
                <TableCell className="text-right font-medium">{formatNumber(inv.quantity)}</TableCell>
                <TableCell className="text-right text-orange-600">{inv.lockedQty > 0 ? formatNumber(inv.lockedQty) : '-'}</TableCell>
                <TableCell className="text-right font-medium text-green-600">{formatNumber(inv.availableQty)}</TableCell>
                <TableCell>
                  <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${getStatusColor(INVENTORY_STATUS_MAP, inv.status)}`}>
                    {getStatusLabel(INVENTORY_STATUS_MAP, inv.status)}
                  </span>
                </TableCell>
                <TableCell className="text-gray-500">{inv.lastUpdate}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <Pagination page={page} totalPages={Math.ceil(filtered.length / 10)} onPageChange={setPage} />
      </div>
    </div>
  );
}
