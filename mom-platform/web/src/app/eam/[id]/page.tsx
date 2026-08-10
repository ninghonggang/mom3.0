'use client';

import { useParams, useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { OeeChart } from '@/components/charts/oee-chart';
import { BarChart } from '@/components/charts/bar-chart';
import {
  Table, TableHeader, TableBody, TableRow, TableHead, TableCell,
} from '@/components/ui/table';
import { formatDateTime, formatPercent, getStatusLabel, getStatusColor, EQUIPMENT_STATUS_MAP } from '@/lib/utils';

const mockEq = {
  id: '1', equipmentNo: 'EQ-001', name: 'CNC数控加工中心 #01', model: 'DMG MORI CMX 1100V',
  type: 'CNC', workCenter: 'CNC-01', status: 'RUNNING' as const,
  oee: 0.88, availability: 0.95, performance: 0.93, quality: 0.99,
  totalRunTime: 12500, totalDowntime: 320,
  lastMaintenance: '2025-01-02', nextMaintenance: '2025-01-15',
  installDate: '2020-03-15', manufacturer: 'DMG MORI',
  serialNo: 'SN-20200315-001', power: '30kW', weight: '7500kg',
};

const oeeTrend = [
  { label: '周一', planned: 0, actual: 0.90 },
  { label: '周二', planned: 0, actual: 0.87 },
  { label: '周三', planned: 0, actual: 0.92 },
  { label: '周四', planned: 0, actual: 0.85 },
  { label: '周五', planned: 0, actual: 0.88 },
  { label: '周六', planned: 0, actual: 0.91 },
  { label: '周日', planned: 0, actual: 0.89 },
].map(d => ({ ...d, planned: d.actual * 0.95 }));

const mockMaintenancePlans = [
  { id: 'm1', planNo: 'MP-2025-001', type: 'ROUTINE', description: '日常点检保养', plannedDate: '2025-01-15', status: 'PLANNED', technician: '维护员A' },
  { id: 'm2', planNo: 'MP-2024-030', type: 'INSPECTION', description: '主轴精度检测', plannedDate: '2025-01-02', actualDate: '2025-01-02', status: 'COMPLETED', technician: '维护员B' },
  { id: 'm3', planNo: 'MP-2024-028', type: 'OVERHAUL', description: '年度大修', plannedDate: '2024-12-20', actualDate: '2024-12-22', status: 'COMPLETED', technician: '维护团队' },
];

const mockRepairHistory = [
  { id: 'r1', repairNo: 'REP-2024-056', faultDesc: '主轴异响', severity: 'MAJOR', status: 'CONFIRMED', reportedAt: '2024-11-10', endTime: '2024-11-11', downtime: 8 },
  { id: 'r2', repairNo: 'REP-2024-032', faultDesc: '冷却液泵故障', severity: 'MINOR', status: 'CONFIRMED', reportedAt: '2024-08-05', endTime: '2024-08-05', downtime: 2 },
];

const mockDowntimeLogs = [
  { id: 'd1', startTime: '2025-01-06 14:00', duration: 2.5, reason: '刀具磨损更换', category: '换刀', reportedBy: '张操作' },
  { id: 'd2', startTime: '2025-01-04 09:30', duration: 1.0, reason: '程序传输异常', category: '系统', reportedBy: '李操作' },
  { id: 'd3', startTime: '2024-12-28 15:00', duration: 0.5, reason: '夹紧气压不足', category: '机械', reportedBy: '王操作' },
];

const planTypeLabels: Record<string, string> = { ROUTINE: '日常保养', INSPECTION: '精度检测', OVERHAUL: '大修', PREDICTIVE: '预测维护' };
const planStatusLabels: Record<string, string> = { PLANNED: '已计划', IN_PROGRESS: '进行中', COMPLETED: '已完成', OVERDUE: '超期' };

const severityLabels: Record<string, string> = { CRITICAL: '严重', MAJOR: '主要', MINOR: '次要' };
const severityColors: Record<string, string> = { CRITICAL: 'text-red-600', MAJOR: 'text-orange-600', MINOR: 'text-yellow-600' };

export default function EquipmentDetailPage() {
  const params = useParams();
  const router = useRouter();

  return (
    <div className="space-y-6">
      <div className="page-header">
        <div className="flex items-center gap-3">
          <Button variant="ghost" size="sm" onClick={() => router.back()}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="15 18 9 12 15 6" /></svg>
          </Button>
          <h1 className="page-title">{mockEq.name}</h1>
          <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${getStatusColor(EQUIPMENT_STATUS_MAP, mockEq.status)}`}>
            {getStatusLabel(EQUIPMENT_STATUS_MAP, mockEq.status)}
          </span>
        </div>
        <div className="flex items-center gap-2">
          <Button variant="outline" size="sm">编辑</Button>
          <Button size="sm">报修</Button>
        </div>
      </div>

      {/* Equipment Info */}
      <Card>
        <CardContent className="p-5">
          <div className="grid grid-cols-2 gap-x-8 gap-y-3 md:grid-cols-4">
            <div><span className="text-xs text-gray-500">设备编号</span><p className="font-medium">{mockEq.equipmentNo}</p></div>
            <div><span className="text-xs text-gray-500">型号</span><p className="font-medium">{mockEq.model}</p></div>
            <div><span className="text-xs text-gray-500">类型</span><p className="font-medium">{mockEq.type}</p></div>
            <div><span className="text-xs text-gray-500">工作中心</span><p className="font-medium">{mockEq.workCenter}</p></div>
            <div><span className="text-xs text-gray-500">制造商</span><p className="font-medium">{mockEq.manufacturer}</p></div>
            <div><span className="text-xs text-gray-500">序列号</span><p className="font-medium">{mockEq.serialNo}</p></div>
            <div><span className="text-xs text-gray-500">功率</span><p className="font-medium">{mockEq.power}</p></div>
            <div><span className="text-xs text-gray-500">重量</span><p className="font-medium">{mockEq.weight}</p></div>
            <div><span className="text-xs text-gray-500">安装日期</span><p className="font-medium">{mockEq.installDate}</p></div>
            <div><span className="text-xs text-gray-500">累计运行(h)</span><p className="font-medium">{mockEq.totalRunTime.toLocaleString()}</p></div>
            <div><span className="text-xs text-gray-500">累计停机(h)</span><p className="font-medium text-red-600">{mockEq.totalDowntime.toLocaleString()}</p></div>
          </div>
        </CardContent>
      </Card>

      <Tabs defaultValue="oee">
        <TabsList>
          <TabsTrigger value="oee">OEE 指标</TabsTrigger>
          <TabsTrigger value="maintenance">保养计划</TabsTrigger>
          <TabsTrigger value="repairs">维修历史</TabsTrigger>
          <TabsTrigger value="downtime">停机记录</TabsTrigger>
        </TabsList>

        <TabsContent value="oee">
          <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
            <Card>
              <CardContent className="p-6">
                <h3 className="mb-4 text-sm font-medium text-gray-700">当前 OEE</h3>
                <OeeChart
                  availability={mockEq.availability}
                  performance={mockEq.performance}
                  quality={mockEq.quality}
                  oee={mockEq.oee}
                />
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                <h3 className="mb-4 text-sm font-medium text-gray-700">本周 OEE 趋势</h3>
                <BarChart data={oeeTrend} height={200} />
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="maintenance">
          <div className="rounded-lg border bg-white">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>计划编号</TableHead>
                  <TableHead>类型</TableHead>
                  <TableHead>描述</TableHead>
                  <TableHead>计划日期</TableHead>
                  <TableHead>实际日期</TableHead>
                  <TableHead>状态</TableHead>
                  <TableHead>负责人</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {mockMaintenancePlans.map((p) => (
                  <TableRow key={p.id}>
                    <TableCell className="font-medium">{p.planNo}</TableCell>
                    <TableCell>{planTypeLabels[p.type] || p.type}</TableCell>
                    <TableCell className="text-gray-600">{p.description}</TableCell>
                    <TableCell className="text-gray-500">{p.plannedDate}</TableCell>
                    <TableCell className="text-gray-500">{p.actualDate || '-'}</TableCell>
                    <TableCell>
                      <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${
                        p.status === 'COMPLETED' ? 'bg-green-100 text-green-700' : 'bg-blue-100 text-blue-700'
                      }`}>
                        {planStatusLabels[p.status]}
                      </span>
                    </TableCell>
                    <TableCell className="text-gray-500">{p.technician || '-'}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </TabsContent>

        <TabsContent value="repairs">
          <div className="rounded-lg border bg-white">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>维修编号</TableHead>
                  <TableHead>故障描述</TableHead>
                  <TableHead>严重级别</TableHead>
                  <TableHead>报修时间</TableHead>
                  <TableHead>修复时间</TableHead>
                  <TableHead className="text-right">停机(h)</TableHead>
                  <TableHead>状态</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {mockRepairHistory.map((r) => (
                  <TableRow key={r.id} className="cursor-pointer hover:bg-gray-50" onClick={() => router.push(`/eam/repairs/${r.id}`)}>
                    <TableCell className="font-medium text-blue-600">{r.repairNo}</TableCell>
                    <TableCell className="text-gray-600">{r.faultDesc}</TableCell>
                    <TableCell>
                      <span className={severityColors[r.severity]}>{severityLabels[r.severity]}</span>
                    </TableCell>
                    <TableCell className="text-gray-500">{formatDateTime(r.reportedAt)}</TableCell>
                    <TableCell className="text-gray-500">{formatDateTime(r.endTime)}</TableCell>
                    <TableCell className="text-right text-red-600">{r.downtime}</TableCell>
                    <TableCell>
                      <span className="inline-flex items-center rounded-full bg-green-100 px-2 py-0.5 text-xs font-medium text-green-700">已确认</span>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </TabsContent>

        <TabsContent value="downtime">
          <div className="rounded-lg border bg-white">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>开始时间</TableHead>
                  <TableHead className="text-right">持续(h)</TableHead>
                  <TableHead>原因</TableHead>
                  <TableHead>类别</TableHead>
                  <TableHead>报告人</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {mockDowntimeLogs.map((d) => (
                  <TableRow key={d.id}>
                    <TableCell className="text-gray-500">{d.startTime}</TableCell>
                    <TableCell className="text-right text-red-600">{d.duration}</TableCell>
                    <TableCell className="text-gray-600">{d.reason}</TableCell>
                    <TableCell className="text-gray-500">{d.category}</TableCell>
                    <TableCell className="text-gray-500">{d.reportedBy}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
