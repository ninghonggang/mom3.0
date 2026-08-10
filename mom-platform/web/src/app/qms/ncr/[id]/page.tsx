'use client';

import { useParams, useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent } from '@/components/ui/card';
import { formatDateTime, getStatusLabel, getStatusColor, NCR_SEVERITY_MAP, NCR_STATUS_MAP } from '@/lib/utils';

const mockNcr = {
  id: '1',
  ncrNo: 'NCR-2025-0001',
  source: 'IQC-2025-0001',
  orderNo: 'PO-240101',
  materialCode: 'MAT-001',
  materialName: '主轴轴承座',
  spec: 'Φ120×80',
  supplier: '供应商A',
  lotNo: 'LOT-20250101',
  defectDescription: '表面粗糙度超差，Ra实测值1.8μm，标准要求Ra≤1.6μm。共计2件。',
  severity: 'MAJOR' as const,
  status: 'DISPOSITION' as const,
  quantity: 2,
  inspector: '质检员A',
  foundAt: '2025-01-06 15:30',
  disposition: '返工处理',
};

const mockActions = [
  { id: 'a1', seq: 1, action: 'NCR 开立', operator: '质检员A', result: '发现来料表面粗糙度超差2件，开立NCR', createdAt: '2025-01-06 15:30' },
  { id: 'a2', seq: 2, action: '评审处理', operator: '质量工程师', result: '经评审决定返工处理，将表面重新打磨至Ra≤1.6μm', createdAt: '2025-01-06 17:00' },
  { id: 'a3', seq: 3, action: '返工执行', operator: '操作员B', result: '正在进行返工中...', createdAt: '2025-01-07 08:30' },
];

const statusFlow = ['OPEN', 'DISPOSITION', 'IN_REWORK', 'VERIFIED', 'CLOSED'];
const activeStep = statusFlow.indexOf(mockNcr.status);

export default function NcrDetailPage() {
  const params = useParams();
  const router = useRouter();

  return (
    <div className="space-y-6">
      <div className="page-header">
        <div className="flex items-center gap-3">
          <Button variant="ghost" size="sm" onClick={() => router.back()}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="15 18 9 12 15 6" /></svg>
          </Button>
          <h1 className="page-title">{mockNcr.ncrNo}</h1>
          <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${getStatusColor(NCR_STATUS_MAP, mockNcr.status)}`}>
            {getStatusLabel(NCR_STATUS_MAP, mockNcr.status)}
          </span>
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
                  <div
                    className={`flex h-8 w-8 items-center justify-center rounded-full text-xs font-medium ${
                      i <= activeStep
                        ? 'bg-blue-600 text-white'
                        : 'bg-gray-200 text-gray-400'
                    }`}
                  >
                    {i + 1}
                  </div>
                  <span className={`mt-1 text-xs ${i <= activeStep ? 'text-blue-600 font-medium' : 'text-gray-400'}`}>
                    {getStatusLabel(NCR_STATUS_MAP, step)}
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

      {/* NCR Info */}
      <Card>
        <CardContent className="p-5">
          <h3 className="mb-4 text-sm font-medium text-gray-700">不合格信息</h3>
          <div className="grid grid-cols-2 gap-x-8 gap-y-3 md:grid-cols-4">
            <div><span className="text-xs text-gray-500">来源检验单</span><p className="font-medium">{mockNcr.source}</p></div>
            <div><span className="text-xs text-gray-500">物料编码</span><p className="font-medium">{mockNcr.materialCode}</p></div>
            <div><span className="text-xs text-gray-500">物料名称</span><p className="font-medium">{mockNcr.materialName}</p></div>
            <div><span className="text-xs text-gray-500">规格</span><p className="font-medium">{mockNcr.spec}</p></div>
            <div><span className="text-xs text-gray-500">供应商</span><p className="font-medium">{mockNcr.supplier}</p></div>
            <div><span className="text-xs text-gray-500">批号</span><p className="font-medium">{mockNcr.lotNo}</p></div>
            <div><span className="text-xs text-gray-500">不合格数量</span><p className="font-medium text-red-600">{mockNcr.quantity} 件</p></div>
            <div><span className="text-xs text-gray-500">严重级别</span><p>
              <span className={`inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ${getStatusColor(NCR_SEVERITY_MAP, mockNcr.severity)}`}>
                {getStatusLabel(NCR_SEVERITY_MAP, mockNcr.severity)}
              </span>
            </p></div>
            <div className="col-span-2 md:col-span-4">
              <span className="text-xs text-gray-500">缺陷描述</span>
              <p className="mt-1 text-sm text-gray-700">{mockNcr.defectDescription}</p>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Action History */}
      <Card>
        <CardContent className="p-6">
          <h3 className="mb-4 text-sm font-medium text-gray-700">处理历史</h3>
          <div className="relative pl-8">
            {mockActions.map((action, i) => (
              <div key={action.id} className="relative pb-8 last:pb-0">
                {i < mockActions.length - 1 && (
                  <div className="absolute left-[-19px] top-4 h-full w-0.5 bg-gray-200" />
                )}
                <div className={`absolute left-[-25px] top-1 h-3 w-3 rounded-full border-2 ${i === 0 ? 'border-blue-500 bg-blue-100' : 'border-gray-300 bg-white'}`} />
                <div>
                  <div className="flex items-center gap-2">
                    <span className="font-medium text-sm">{action.action}</span>
                    <span className="text-xs text-gray-400">{action.createdAt}</span>
                  </div>
                  <p className="mt-1 text-sm text-gray-500">{action.result}</p>
                  <p className="text-xs text-gray-400">操作人: {action.operator}</p>
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
