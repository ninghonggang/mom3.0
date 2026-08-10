'use client';

import { useParams, useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from '@/components/ui/table';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import {
  formatDate,
  formatDateTime,
  formatNumber,
  getStatusLabel,
  getStatusColor,
  ORDER_STATUS_MAP,
} from '@/lib/utils';
import type { ProductionOrder, DispatchRecord, JobReport, OrderCompletion, OrderTimeline } from '@/types';

const mockOrder: ProductionOrder = {
  id: '1',
  orderNo: 'MO-2024-0001',
  materialCode: 'MAT-001',
  materialName: '主轴轴承座',
  spec: 'Φ120×80 / 45#钢',
  quantity: 500,
  completedQty: 320,
  unit: '件',
  status: 'IN_PROGRESS',
  planStartDate: '2025-01-06',
  planEndDate: '2025-01-10',
  actualStartDate: '2025-01-06',
  workCenter: 'CNC-01',
  priority: 1,
  bomVersion: 'BOM-V2.1',
  routingVersion: 'R-V1.3',
  createdAt: '2025-01-05',
};

const mockDispatches: DispatchRecord[] = [
  { id: 'd1', orderId: '1', operationSeq: 10, operationName: '粗车外圆', workCenter: 'CNC-01', machineId: 'MC-101', workerName: '张操作', planStart: '2025-01-06 08:00', planEnd: '2025-01-07 12:00', actualStart: '2025-01-06 08:15', actualEnd: '2025-01-07 11:30', quantity: 500, goodQty: 495, defectQty: 5, status: 'COMPLETED' },
  { id: 'd2', orderId: '1', operationSeq: 20, operationName: '精车内孔', workCenter: 'CNC-02', machineId: 'MC-202', workerName: '李操作', planStart: '2025-01-07 13:00', planEnd: '2025-01-08 17:00', actualStart: '2025-01-07 13:30', quantity: 495, goodQty: 320, defectQty: 0, status: 'IN_PROGRESS' },
  { id: 'd3', orderId: '1', operationSeq: 30, operationName: '钻孔攻牙', workCenter: 'DRILL-01', workerName: '王操作', planStart: '2025-01-09 08:00', planEnd: '2025-01-10 12:00', quantity: 495, goodQty: 0, defectQty: 0, status: 'PENDING' },
];

const mockReports: JobReport[] = [
  { id: 'r1', dispatchId: 'd1', operationName: '粗车外圆', workerName: '张操作', machineId: 'MC-101', startTime: '2025-01-06 08:15', endTime: '2025-01-07 11:30', goodQty: 495, defectQty: 5, defectCodes: ['D01-尺寸超差'], remark: '5件外径偏大0.02mm' },
];

const mockCompletion: OrderCompletion = {
  id: 'c1',
  orderId: '1',
  totalQty: 500,
  goodQty: 320,
  defectQty: 5,
  scrapQty: 0,
  passRate: 0.984,
  completedAt: '',
  inspectorName: '',
};

const mockTimeline: OrderTimeline[] = [
  { id: 't1', orderId: '1', event: '订单创建', description: '生产计划员创建工单', operator: '计划员', timestamp: '2025-01-05 09:00' },
  { id: 't2', orderId: '1', event: '订单下达', description: '订单审核通过并下达车间', operator: '计划主管', timestamp: '2025-01-05 14:30' },
  { id: 't3', orderId: '1', event: '开始生产', description: 'CNC-01 开始粗车外圆', operator: '张操作', timestamp: '2025-01-06 08:15' },
  { id: 't4', orderId: '1', event: '工序完成', description: '粗车外圆工序完成,合格495件', operator: '张操作', timestamp: '2025-01-07 11:30' },
  { id: 't5', orderId: '1', event: '开始精车', description: 'CNC-02 开始精车内孔', operator: '李操作', timestamp: '2025-01-07 13:30' },
];

export default function OrderDetailPage() {
  const params = useParams();
  const router = useRouter();

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="page-header">
        <div className="flex items-center gap-3">
          <Button variant="ghost" size="sm" onClick={() => router.back()}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="15 18 9 12 15 6" /></svg>
          </Button>
          <h1 className="page-title">{mockOrder.orderNo}</h1>
          <span
            className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${
              getStatusColor(ORDER_STATUS_MAP, mockOrder.status)
            }`}
          >
            {getStatusLabel(ORDER_STATUS_MAP, mockOrder.status)}
          </span>
        </div>
        <div className="flex items-center gap-2">
          <Button variant="outline" size="sm">编辑</Button>
          <Button variant="outline" size="sm">暂停</Button>
          <Button size="sm">报工</Button>
        </div>
      </div>

      {/* Order Info */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">订单信息</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-2 gap-x-8 gap-y-3 md:grid-cols-4">
            <div><span className="text-sm text-gray-500">物料编码</span><p className="font-medium">{mockOrder.materialCode}</p></div>
            <div><span className="text-sm text-gray-500">物料名称</span><p className="font-medium">{mockOrder.materialName}</p></div>
            <div><span className="text-sm text-gray-500">规格</span><p className="font-medium">{mockOrder.spec}</p></div>
            <div><span className="text-sm text-gray-500">BOM版本</span><p className="font-medium">{mockOrder.bomVersion}</p></div>
            <div><span className="text-sm text-gray-500">订单数量</span><p className="font-medium">{formatNumber(mockOrder.quantity)} {mockOrder.unit}</p></div>
            <div><span className="text-sm text-gray-500">已完成</span><p className="font-medium text-blue-600">{formatNumber(mockOrder.completedQty)} {mockOrder.unit}</p></div>
            <div><span className="text-sm text-gray-500">工作中心</span><p className="font-medium">{mockOrder.workCenter}</p></div>
            <div><span className="text-sm text-gray-500">工艺版本</span><p className="font-medium">{mockOrder.routingVersion}</p></div>
            <div><span className="text-sm text-gray-500">计划开始</span><p className="font-medium">{formatDate(mockOrder.planStartDate)}</p></div>
            <div><span className="text-sm text-gray-500">计划结束</span><p className="font-medium">{formatDate(mockOrder.planEndDate)}</p></div>
            <div><span className="text-sm text-gray-500">实际开始</span><p className="font-medium">{formatDate(mockOrder.actualStartDate)}</p></div>
            <div><span className="text-sm text-gray-500">优先级</span><p className="font-medium">{mockOrder.priority}</p></div>
          </div>
        </CardContent>
      </Card>

      {/* Tabs */}
      <Tabs defaultValue="dispatch">
        <TabsList>
          <TabsTrigger value="dispatch">派工记录</TabsTrigger>
          <TabsTrigger value="reports">作业报表</TabsTrigger>
          <TabsTrigger value="completion">完工信息</TabsTrigger>
          <TabsTrigger value="timeline">时间线</TabsTrigger>
        </TabsList>

        <TabsContent value="dispatch">
          <div className="rounded-lg border bg-white">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-16">工序</TableHead>
                  <TableHead>工序名称</TableHead>
                  <TableHead>工作中心</TableHead>
                  <TableHead>设备</TableHead>
                  <TableHead>操作员</TableHead>
                  <TableHead>计划开始</TableHead>
                  <TableHead>计划结束</TableHead>
                  <TableHead className="text-right">数量</TableHead>
                  <TableHead className="text-right">合格</TableHead>
                  <TableHead>状态</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {mockDispatches.map((d) => (
                  <TableRow key={d.id}>
                    <TableCell className="font-medium">{d.operationSeq}</TableCell>
                    <TableCell>{d.operationName}</TableCell>
                    <TableCell className="text-gray-500">{d.workCenter}</TableCell>
                    <TableCell className="text-gray-500">{d.machineId || '-'}</TableCell>
                    <TableCell>{d.workerName}</TableCell>
                    <TableCell className="text-gray-500">{d.planStart}</TableCell>
                    <TableCell className="text-gray-500">{d.planEnd}</TableCell>
                    <TableCell className="text-right">{d.quantity}</TableCell>
                    <TableCell className="text-right">{d.goodQty}</TableCell>
                    <TableCell>
                      <Badge variant={d.status === 'COMPLETED' ? 'success' : d.status === 'IN_PROGRESS' ? 'warning' : 'default'}>
                        {d.status === 'COMPLETED' ? '已完成' : d.status === 'IN_PROGRESS' ? '进行中' : '待执行'}
                      </Badge>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </TabsContent>

        <TabsContent value="reports">
          <div className="rounded-lg border bg-white">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>工序名称</TableHead>
                  <TableHead>操作员</TableHead>
                  <TableHead>设备</TableHead>
                  <TableHead>开始时间</TableHead>
                  <TableHead>结束时间</TableHead>
                  <TableHead className="text-right">合格数</TableHead>
                  <TableHead className="text-right">不良数</TableHead>
                  <TableHead>不良代码</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {mockReports.map((r) => (
                  <TableRow key={r.id}>
                    <TableCell>{r.operationName}</TableCell>
                    <TableCell>{r.workerName}</TableCell>
                    <TableCell className="text-gray-500">{r.machineId || '-'}</TableCell>
                    <TableCell className="text-gray-500">{r.startTime}</TableCell>
                    <TableCell className="text-gray-500">{r.endTime || '-'}</TableCell>
                    <TableCell className="text-right text-green-600">{r.goodQty}</TableCell>
                    <TableCell className="text-right text-red-600">{r.defectQty}</TableCell>
                    <TableCell className="text-gray-500">{r.defectCodes.join(', ') || '-'}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </TabsContent>

        <TabsContent value="completion">
          <Card>
            <CardContent className="p-6">
              <div className="grid grid-cols-2 gap-y-4 md:grid-cols-4">
                <div><span className="text-sm text-gray-500">总订单数量</span><p className="text-lg font-bold">{mockCompletion.totalQty}</p></div>
                <div><span className="text-sm text-gray-500">合格数量</span><p className="text-lg font-bold text-green-600">{mockCompletion.goodQty}</p></div>
                <div><span className="text-sm text-gray-500">不良数量</span><p className="text-lg font-bold text-red-600">{mockCompletion.defectQty}</p></div>
                <div><span className="text-sm text-gray-500">合格率</span><p className="text-lg font-bold text-blue-600">{(mockCompletion.passRate * 100).toFixed(1)}%</p></div>
              </div>
              <div className="mt-8 text-center text-sm text-gray-400">生产进行中，待全部完工后更新最终数据</div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="timeline">
          <div className="relative pl-8">
            {mockTimeline.map((item, i) => (
              <div key={item.id} className="relative pb-8 last:pb-0">
                {/* Line */}
                {i < mockTimeline.length - 1 && (
                  <div className="absolute left-[-19px] top-4 h-full w-0.5 bg-gray-200" />
                )}
                {/* Dot */}
                <div
                  className={`absolute left-[-25px] top-1 h-3 w-3 rounded-full border-2 ${
                    i === 0 ? 'border-blue-500 bg-blue-100' : 'border-gray-300 bg-white'
                  }`}
                />
                <div>
                  <div className="flex items-center gap-2">
                    <span className="font-medium text-sm">{item.event}</span>
                    <span className="text-xs text-gray-400">{item.timestamp}</span>
                  </div>
                  <p className="mt-1 text-sm text-gray-500">{item.description}</p>
                  <p className="text-xs text-gray-400">操作人: {item.operator}</p>
                </div>
              </div>
            ))}
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
