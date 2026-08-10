'use client';

import { useSearchParams, useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

interface TraceNodeView {
  id: string;
  label: string;
  type: string;
  detail: string;
  time: string;
  children?: TraceNodeView[];
}

const forwardData: TraceNodeView = {
  id: 'root',
  label: '来料',
  type: 'IQC',
  detail: 'LOT-20250101 检验合格',
  time: '2025-01-06',
  children: [
    {
      id: 'op1',
      label: '粗车外圆',
      type: '工序',
      detail: '工作中心: CNC-01, 操作员: 张操作',
      time: '2025-01-06 08:00',
      children: [
        {
          id: 'op2',
          label: '精车内孔',
          type: '工序',
          detail: '工作中心: CNC-02, 操作员: 李操作',
          time: '2025-01-07 13:00',
          children: [
            {
              id: 'op3',
              label: '钻孔攻牙',
              type: '工序',
              detail: '工作中心: DRILL-01',
              time: '2025-01-09 (计划)',
            },
            {
              id: 'q1',
              label: '过程检验',
              type: 'IPQC',
              detail: 'IPQC-2025-0001 抽检合格',
              time: '2025-01-07 17:00',
            },
          ],
        },
      ],
    },
    {
      id: 'q2',
      label: '完工检验',
      type: 'FQC',
      detail: '待检验',
      time: '-',
    },
    {
      id: 'ship',
      label: '发货',
      type: '发货',
      detail: '客户A',
      time: '-',
    },
  ],
};

const backwardData: TraceNodeView = {
  id: 'root-b',
  label: '产品 SN-20250106-001',
  type: '成品',
  detail: '主轴轴承座 Φ120×80',
  time: '2025-01-08',
  children: [
    {
      id: 'bom1',
      label: '物料: 轴承座毛坯',
      type: 'BOM',
      detail: 'MAT-001 供应商A',
      time: '2025-01-06 收货',
      children: [
        {
          id: 'bom1-lot',
          label: '批次 LOT-20250101',
          type: '批次',
          detail: 'IQC 检验合格, AQL 1.0',
          time: '2025-01-06',
          children: [
            {
              id: 'supplier',
              label: '供应商A',
              type: '供应商',
              detail: '原料批号: SUP-LOT-20241220',
              time: '2024-12-20',
            },
          ],
        },
      ],
    },
    {
      id: 'bom2',
      label: '物料: 轴承 6208',
      type: 'BOM',
      detail: 'MAT-007 供应商B',
      time: '2025-01-05 收货',
      children: [
        {
          id: 'bom2-lot',
          label: '批次 LOT-20250100',
          type: '批次',
          detail: 'IQC 检验合格',
          time: '2025-01-05',
        },
      ],
    },
  ],
};

function TraceTree({ node, level = 0 }: { node: TraceNodeView; level?: number }) {
  return (
    <div className="ml-4">
      <div className="relative pb-3">
        {level > 0 && (
          <div className="absolute left-[-16px] top-0 h-full w-px bg-gray-200" />
        )}
        <div className="relative flex items-start gap-3">
          {level > 0 && (
            <div className="absolute left-[-20px] top-3 h-px w-4 bg-gray-200" />
          )}
          <div className="flex-1 rounded-lg border p-3">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2">
                <Badge variant={node.type === 'IQC' ? 'info' : node.type === '工序' ? 'success' : 'default'}>
                  {node.type}
                </Badge>
                <span className="font-medium text-sm">{node.label}</span>
              </div>
              <span className="text-xs text-gray-400">{node.time}</span>
            </div>
            <p className="mt-1 text-xs text-gray-500">{node.detail}</p>
          </div>
        </div>
      </div>
      {node.children?.map((child) => (
        <TraceTree key={child.id} node={child} level={level + 1} />
      ))}
    </div>
  );
}

export function TraceResultContent() {
  const searchParams = useSearchParams();
  const router = useRouter();
  const type = searchParams.get('type') || 'serial_no';
  const value = searchParams.get('value') || '';

  return (
    <div className="space-y-6">
      <div className="page-header">
        <div className="flex items-center gap-3">
          <Button variant="ghost" size="sm" onClick={() => router.back()}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="15 18 9 12 15 6" /></svg>
          </Button>
          <h1 className="page-title">追溯结果</h1>
        </div>
      </div>

      {/* Search Summary */}
      <Card>
        <CardContent className="p-4">
          <div className="flex items-center gap-4">
            <div className="rounded-lg bg-gray-50 px-3 py-1.5">
              <span className="text-xs text-gray-500">
                {type === 'serial_no' ? '序列号' : type === 'batch_no' ? '批号' : '物料'}
              </span>
              <span className="ml-2 text-sm font-medium">{value}</span>
            </div>
            <div className="text-sm text-gray-500">
              向前追溯: 5级 / 向后追溯: 3级
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Forward/Backward Tabs */}
      <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle className="text-base flex items-center gap-2">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#3b82f6" strokeWidth="2"><polyline points="13 17 18 12 13 7" /><polyline points="6 17 11 12 6 7" /></svg>
              向前追溯（产品流向）
            </CardTitle>
          </CardHeader>
          <CardContent>
            <TraceTree node={forwardData} />
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="text-base flex items-center gap-2">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#22c55e" strokeWidth="2"><polyline points="11 17 6 12 11 7" /><polyline points="18 17 13 12 18 7" /></svg>
              向后追溯（物料来源）
            </CardTitle>
          </CardHeader>
          <CardContent>
            <TraceTree node={backwardData} />
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
