'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

export default function TracePage() {
  const router = useRouter();
  const [searchType, setSearchType] = useState<'serial_no' | 'batch_no' | 'material'>('serial_no');
  const [searchValue, setSearchValue] = useState('');

  const handleSearch = () => {
    if (searchValue.trim()) {
      router.push(`/trace/result?type=${searchType}&value=${encodeURIComponent(searchValue)}`);
    }
  };

  return (
    <div className="space-y-6">
      <div className="page-header">
        <h1 className="page-title">质量追溯</h1>
      </div>

      <Card className="max-w-2xl mx-auto mt-12">
        <CardHeader>
          <CardTitle>追溯查询</CardTitle>
          <p className="text-sm text-gray-500">输入编号查询产品完整追溯链</p>
        </CardHeader>
        <CardContent>
          {/* Search Type Tabs */}
          <div className="mb-6 inline-flex rounded-lg bg-gray-100 p-1">
            {([
              { value: 'serial_no', label: '序列号' },
              { value: 'batch_no', label: '批号' },
              { value: 'material', label: '物料编码' },
            ] as const).map((type) => (
              <button
                key={type.value}
                onClick={() => setSearchType(type.value)}
                className={`rounded-md px-4 py-1.5 text-sm font-medium transition-all ${
                  searchType === type.value
                    ? 'bg-white text-gray-900 shadow-sm'
                    : 'text-gray-500 hover:text-gray-700'
                }`}
              >
                {type.label}
              </button>
            ))}
          </div>

          {/* Search Input */}
          <div className="flex gap-3">
            <Input
              placeholder={
                searchType === 'serial_no'
                  ? '输入产品序列号，如 SN-20250106-001'
                  : searchType === 'batch_no'
                  ? '输入批号，如 LOT-20250101'
                  : '输入物料编码，如 MAT-001'
              }
              value={searchValue}
              onChange={(e) => setSearchValue(e.target.value)}
              onKeyDown={(e) => e.key === 'Enter' && handleSearch()}
              className="flex-1"
            />
            <Button onClick={handleSearch} disabled={!searchValue.trim()}>
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><circle cx="11" cy="11" r="8" /><path d="m21 21-4.3-4.3" /></svg>
              追溯
            </Button>
          </div>

          {/* Quick Access */}
          <div className="mt-8 border-t pt-6">
            <h4 className="mb-3 text-sm font-medium text-gray-600">最近追溯</h4>
            <div className="space-y-2">
              {[
                { type: 'SN', value: 'SN-20250106-001', label: '主轴轴承座' },
                { type: 'LOT', value: 'LOT-20250101', label: '主轴轴承座批次' },
                { type: 'MAT', value: 'MAT-002', label: '电机法兰盘' },
              ].map((item) => (
                <div
                  key={item.value}
                  className="flex items-center justify-between rounded-lg border p-3 cursor-pointer hover:bg-gray-50"
                  onClick={() => {
                    setSearchValue(item.value);
                    const t = item.type === 'SN' ? 'serial_no' : item.type === 'LOT' ? 'batch_no' : 'material';
                    setSearchType(t);
                  }}
                >
                  <div className="flex items-center gap-3">
                    <span className="inline-flex h-6 w-6 items-center justify-center rounded bg-blue-50 text-xs font-medium text-blue-600">
                      {item.type}
                    </span>
                    <span className="text-sm font-medium text-gray-700">{item.value}</span>
                  </div>
                  <span className="text-xs text-gray-400">{item.label}</span>
                </div>
              ))}
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
