import { Suspense } from 'react';
import { TraceResultContent } from './trace-result-content';

export default function TraceResultPage() {
  return (
    <Suspense fallback={
      <div className="space-y-6">
        <div className="animate-pulse">
          <div className="h-8 w-64 bg-gray-200 rounded mb-6" />
          <div className="h-20 bg-gray-200 rounded mb-6" />
          <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
            <div className="h-96 bg-gray-200 rounded" />
            <div className="h-96 bg-gray-200 rounded" />
          </div>
        </div>
      </div>
    }>
      <TraceResultContent />
    </Suspense>
  );
}
