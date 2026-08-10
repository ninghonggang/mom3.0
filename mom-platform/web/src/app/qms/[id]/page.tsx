'use client';

import { useParams, useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent } from '@/components/ui/card';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatDate, formatNumber } from '@/lib/utils';

const mockSheet = {
  id: '1',
  sheetNo: 'IQC-2025-0001',
  type: 'IQC',
  orderNo: 'PO-240101',
  materialCode: 'MAT-001',
  materialName: '主轴轴承座',
  spec: 'Φ120×80',
  supplier: '供应商A',
  lotNo: 'LOT-20250101',
  batchQty: 500,
  sampleQty: 50,
  status: 'COMPLETED',
  result: 'PASS',
  aqlLevel: 'AQL 1.0 (正常检验)',
  inspector: '质检员A',
  startTime: '2025-01-06 09:00',
  endTime: '2025-01-06 16:30',
  createdAt: '2025-01-05',
};

const mockCharacteristics = [
  { id: 'c1', seq: 1, itemName: '外径', spec: 'Φ120 +0.02/-0', upperLimit: 120.02, lowerLimit: 120.0, unit: 'mm', inspectionMethod: '千分尺', sampleSize: 50, result: '合格', isQualified: true },
  { id: 'c2', seq: 2, itemName: '内孔直径', spec: 'Φ80 +0.03/0', upperLimit: 80.03, lowerLimit: 80.0, unit: 'mm', inspectionMethod: '内径千分尺', sampleSize: 50, result: '合格', isQualified: true },
  { id: 'c3', seq: 3, itemName: '长度', spec: '80 ±0.1', upperLimit: 80.1, lowerLimit: 79.9, unit: 'mm', inspectionMethod: '游标卡尺', sampleSize: 50, result: '合格', isQualified: true },
  { id: 'c4', seq: 4, itemName: '表面粗糙度', spec: 'Ra ≤1.6', upperLimit: 1.6, unit: 'μm', inspectionMethod: '粗糙度仪', sampleSize: 50, result: '不合格', isQualified: false },
  { id: 'c5', seq: 5, itemName: '硬度', spec: 'HRC 28-32', upperLimit: 32, lowerLimit: 28, unit: 'HRC', inspectionMethod: '硬度计', sampleSize: 10, result: '合格', isQualified: true },
];

const mockRecords = [
  { id: 'r1', seq: 1, itemName: '外径', sampleSize: 50, qualified: 50, unqualified: 0, rate: 100 },
  { id: 'r2', seq: 2, itemName: '内孔直径', sampleSize: 50, qualified: 50, unqualified: 0, rate: 100 },
  { id: 'r3', seq: 3, itemName: '长度', sampleSize: 50, qualified: 50, unqualified: 0, rate: 100 },
  { id: 'r4', seq: 4, itemName: '表面粗糙度', sampleSize: 50, qualified: 48, unqualified: 2, rate: 96 },
  { id: 'r5', seq: 5, itemName: '硬度', sampleSize: 10, qualified: 10, unqualified: 0, rate: 100 },
];

export default function InspectionDetailPage() {
  const params = useParams();
  const router = useRouter();

  return (
    <div className="space-y-6">
      <div className="page-header">
        <div className="flex items-center gap-3">
          <Button variant="ghost" size="sm" onClick={() => router.back()}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="15 18 9 12 15 6" /></svg>
          </Button>
          <h1 className="page-title">{mockSheet.sheetNo}</h1>
          <Badge variant={mockSheet.result === 'PASS' ? 'success' : 'destructive'}>
            {mockSheet.result === 'PASS' ? '合格' : '不合格'}
          </Badge>
        </div>
      </div>

      {/* Sheet Info */}
      <Card>
        <CardContent className="p-5">
          <div className="grid grid-cols-2 gap-x-8 gap-y-3 md:grid-cols-4">
            <div><span className="text-xs text-gray-500">检验类型</span><p className="font-medium">{mockSheet.type}</p></div>
            <div><span className="text-xs text-gray-500">订单号</span><p className="font-medium">{mockSheet.orderNo}</p></div>
            <div><span className="text-xs text-gray-500">物料</span><p className="font-medium">{mockSheet.materialName}</p></div>
            <div><span className="text-xs text-gray-500">规格</span><p className="font-medium">{mockSheet.spec}</p></div>
            <div><span className="text-xs text-gray-500">供应商</span><p className="font-medium">{mockSheet.supplier}</p></div>
            <div><span className="text-xs text-gray-500">批号</span><p className="font-medium">{mockSheet.lotNo}</p></div>
            <div><span className="text-xs text-gray-500">批量</span><p className="font-medium">{mockSheet.batchQty}</p></div>
            <div><span className="text-xs text-gray-500">抽检数</span><p className="font-medium">{mockSheet.sampleQty}</p></div>
            <div><span className="text-xs text-gray-500">AQL</span><p className="font-medium">{mockSheet.aqlLevel}</p></div>
            <div><span className="text-xs text-gray-500">检验员</span><p className="font-medium">{mockSheet.inspector}</p></div>
            <div><span className="text-xs text-gray-500">检验日期</span><p className="font-medium">{mockSheet.startTime}</p></div>
          </div>
        </CardContent>
      </Card>

      <Tabs defaultValue="characteristics">
        <TabsList>
          <TabsTrigger value="characteristics">检验特性</TabsTrigger>
          <TabsTrigger value="results">检验结果</TabsTrigger>
          <TabsTrigger value="judgment">AQL 判定</TabsTrigger>
        </TabsList>

        <TabsContent value="characteristics">
          <div className="rounded-lg border bg-white">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-12">序号</TableHead>
                  <TableHead>检验项目</TableHead>
                  <TableHead>规格要求</TableHead>
                  <TableHead className="text-right">上限</TableHead>
                  <TableHead className="text-right">下限</TableHead>
                  <TableHead>单位</TableHead>
                  <TableHead>检验方法</TableHead>
                  <TableHead className="text-right">抽样数</TableHead>
                  <TableHead>结果</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {mockCharacteristics.map((c) => (
                  <TableRow key={c.id}>
                    <TableCell>{c.seq}</TableCell>
                    <TableCell>{c.itemName}</TableCell>
                    <TableCell className="text-gray-600">{c.spec}</TableCell>
                    <TableCell className="text-right text-gray-600">{c.upperLimit ?? '-'}</TableCell>
                    <TableCell className="text-right text-gray-600">{c.lowerLimit ?? '-'}</TableCell>
                    <TableCell>{c.unit}</TableCell>
                    <TableCell className="text-gray-600">{c.inspectionMethod}</TableCell>
                    <TableCell className="text-right">{c.sampleSize}</TableCell>
                    <TableCell>
                      <Badge variant={c.isQualified ? 'success' : 'destructive'}>{c.result}</Badge>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </TabsContent>

        <TabsContent value="results">
          <div className="rounded-lg border bg-white">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-12">序号</TableHead>
                  <TableHead>检验项目</TableHead>
                  <TableHead className="text-right">抽样数</TableHead>
                  <TableHead className="text-right">合格数</TableHead>
                  <TableHead className="text-right">不合格数</TableHead>
                  <TableHead className="text-right">合格率</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {mockRecords.map((r) => (
                  <TableRow key={r.id}>
                    <TableCell>{r.seq}</TableCell>
                    <TableCell>{r.itemName}</TableCell>
                    <TableCell className="text-right">{r.sampleSize}</TableCell>
                    <TableCell className="text-right text-green-600">{r.qualified}</TableCell>
                    <TableCell className="text-right text-red-600">{r.unqualified}</TableCell>
                    <TableCell className="text-right">{r.rate}%</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </TabsContent>

        <TabsContent value="judgment">
          <Card>
            <CardContent className="p-6">
              <div className="grid grid-cols-2 gap-8">
                <div className="space-y-4">
                  <h3 className="font-medium text-gray-700">AQL 验收标准</h3>
                  <div className="rounded-lg border bg-gray-50 p-4">
                    <div className="grid grid-cols-2 gap-3 text-sm">
                      <div><span className="text-gray-500">检验水平：</span>AQL 1.0 正常检验</div>
                      <div><span className="text-gray-500">样本量字码：</span>J</div>
                      <div><span className="text-gray-500">样本量：</span>50</div>
                      <div><span className="text-gray-500">Ac (接收)：</span>1</div>
                      <div><span className="text-gray-500">Re (拒收)：</span>2</div>
                    </div>
                  </div>
                </div>
                <div className="space-y-4">
                  <h3 className="font-medium text-gray-700">判定结果</h3>
                  <div className="rounded-lg border p-4">
                    <div className="grid grid-cols-2 gap-3 text-sm">
                      <div><span className="text-gray-500">不合格总数：</span>{mockRecords.reduce((s, r) => s + r.unqualified, 0)}</div>
                      <div><span className="text-gray-500">判定标准：</span>Ac=1, Re=2</div>
                      <div className="col-span-2 mt-2 pt-3 border-t">
                        <p className="text-sm text-gray-500">最终判定：</p>
                        <Badge variant={mockSheet.result === 'PASS' ? 'success' : 'destructive'} className="mt-1 text-base px-4 py-1">
                          {mockSheet.result === 'PASS' ? '批次合格，予以接收' : '批次不合格，予以拒收'}
                        </Badge>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
