interface OeeChartProps {
  availability: number;
  performance: number;
  quality: number;
  oee: number;
}

export function OeeChart({ availability, performance, quality, oee }: OeeChartProps) {
  const radius = 70;
  const circumference = 2 * Math.PI * radius;
  const strokeDasharray = (val: number) =>
    `${(val * circumference).toFixed(0)} ${circumference}`;

  return (
    <div className="flex items-center justify-center gap-4">
      <div className="relative">
        <svg width="180" height="180" viewBox="0 0 180 180">
          <circle cx="90" cy="90" r={radius} fill="none" stroke="#e5e7eb" strokeWidth="12" />
          <circle
            cx="90"
            cy="90"
            r={radius}
            fill="none"
            stroke="#3b82f6"
            strokeWidth="12"
            strokeLinecap="round"
            strokeDasharray={strokeDasharray(oee)}
            transform="rotate(-90 90 90)"
            className="transition-all duration-1000"
          />
        </svg>
        <div className="absolute inset-0 flex flex-col items-center justify-center">
          <span className="text-2xl font-bold text-gray-800">{(oee * 100).toFixed(1)}%</span>
          <span className="text-xs text-gray-500">OEE</span>
        </div>
      </div>
      <div className="space-y-3">
        <div className="flex items-center gap-2">
          <div className="h-3 w-3 rounded-full bg-green-500" />
          <span className="text-xs text-gray-600">可用率 {(availability * 100).toFixed(1)}%</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="h-3 w-3 rounded-full bg-blue-500" />
          <span className="text-xs text-gray-600">性能率 {(performance * 100).toFixed(1)}%</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="h-3 w-3 rounded-full bg-yellow-500" />
          <span className="text-xs text-gray-600">质量率 {(quality * 100).toFixed(1)}%</span>
        </div>
      </div>
    </div>
  );
}
