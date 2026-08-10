'use client';

import { useParams, useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent } from '@/components/ui/card';
import { formatDateTime, getStatusLabel, getStatusColor, REPAIR_STATUS_MAP } from '@/lib/utils';

const mockRepair = {
  id: '1',
  repairNo: 'REP-2025-0001',
  equipmentId: 'EQ-004',
  equipmentName: '三坐标测量机 (ZEISS CONTURA)',
  faultDesc: 'Z轴测量精度超差，重复测量同一工件时Z轴数据偏差超过0.005mm，标准要求≤0.002mm。',
  faultType: '精度',
  severity: 'CRITICAL' as const,
  status: 'IN_PROGRESS' as const,
  reporter: '质检员A',
  reportedAt: '2025-01-07 09:00',
  technician: '维修员A',
  startTime: '2025-01-07 10:00',
  downtime: 4,
  rootCause: '',
  solution: '',
};

const statusFlow = ['REPORTED', 'ASSIGNED', 'IN_PROGRESS', 'COMPLETED', 'CONFIRMED'];
const activeStep = statusFlow.indexOf(mockRepair.status);

export default function RepairDetailPage() {
  const params = useParams();
  const router = useRouter();

  return (
    <div className="space-y-6">
      <div className="page-header">
        <div className="flex items-center gap-3">
          <Button variant="ghost" size="sm" onClick={() => router.back()}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="15 18 9 12 15 6" /></svg>
          </Button>
          <h1 className="page-title">{mockRepair.repairNo}</h1>
          <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${getStatusColor(REPAIR_STATUS_MAP, mockRepair.status)}`}>
            {getStatusLabel(REPAIR_STATUS_MAP, mockRepair.status)}
          </span>
        </div>
        <div className="flex items-center gap-2">
          <Button variant="outline" size="sm">编辑</Button>
          <Button size="sm">完成维修</Button>
        </div>
      </div>

      {/* Status Flow */}
      <Card>
        <CardContent className="p-6">
          <h3 className="mb-4 text-sm font-medium text-gray-700">处理流程</h3>
          <div className="flex items-center">
            {statusFlow.map((step, i) => (
              <div key={step} className="flex items-center flex-1">
                <div className="flex flex-col items-center">
                  <div className={`flex h-8 w-8 items-center justify-center rounded-full text-xs font-medium ${
                    i <= activeStep ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-400'
                  }`}>
                    {i + 1}
                  </div>
                  <span className={`mt-1 text-xs ${i <= activeStep ? 'text-blue-600 font-medium' : 'text-gray-400'}`}>
                    {getStatusLabel(REPAIR_STATUS_MAP, step)}
                  </span>
                </div>
                {i < statusFlow.length - 1 && (
                  <div className={`flex-1 h-0.5 mx-2 ${i < activeStep ? 'bg-blue-600' : 'bg-gray-200'}`} />
                )}
              </div>
            ))}
          </div>
        </CardContent>
      </Card>

      {/* Repair Info */}
      <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
        <Card>
          <CardContent className="p-5">
            <h3 className="mb-4 text-sm font-medium text-gray-700">报修信息</h3>
            <div className="space-y-3">
              <div><span className="text-xs text-gray-500">设备</span><p className="font-medium">{mockRepair.equipmentName}</p></div>
              <div><span className="text-xs text-gray-500">故障描述</span><p className="text-sm text-gray-700">{mockRepair.faultDesc}</p></div>
              <div className="grid grid-cols-2 gap-3">
                <div><span className="text-xs text-gray-500">故障类型</span><p className="font-medium">{mockRepair.faultType}</p></div>
                <div><span className="text-xs text-gray-500">严重级别</span><p>
                  <Badge variant="destructive">严重</Badge>
                </p></div>
                <div><span className="text-xs text-gray-500">报修人</span><p className="font-medium">{mockRepair.reporter}</p></div>
                <div><span className="text-xs text-gray-500">报修时间</span><p className="font-medium">{mockRepair.reportedAt}</p></div>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-5">
            <h3 className="mb-4 text-sm font-medium text-gray-700">维修信息</h3>
            <div className="space-y-3">
              <div className="grid grid-cols-2 gap-3">
                <div><span className="text-xs text-gray-500">维修员</span><p className="font-medium">{mockRepair.technician || '-'}</p></div>
                <div><span className="text-xs text-gray-500">开始时间</span><p className="font-medium">{mockRepair.startTime || '-'}</p></div>
                <div><span className="text-xs text-gray-500">停机时长</span><p className="font-medium text-red-600">{mockRepair.downtime} 小时</p></div>
              </div>
              {mockRepair.rootCause && (
                <>
                  <div><span className="text-xs text-gray-500">根本原因</span><p className="text-sm text-gray-700">{mockRepair.rootCause}</p></div>
                  <div><span className="text-xs text-gray-500">解决方案</span><p className="text-sm text-gray-700">{mockRepair.solution}</p></div>
                </>
              )}
              {!mockRepair.rootCause && (
                <div className="rounded-lg bg-gray-50 p-4 text-center text-sm text-gray-400">
                  维修进行中，待填写根本原因及解决方案
                </div>
              )}
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
