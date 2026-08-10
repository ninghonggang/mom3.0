interface BarChartData {
  label: string;
  planned: number;
  actual: number;
}

interface BarChartProps {
  data: BarChartData[];
  height?: number;
}

export function BarChart({ data, height = 200 }: BarChartProps) {
  const maxVal = Math.max(...data.map((d) => Math.max(d.planned, d.actual)), 1);
  const barWidth = 100 / data.length / 3;

  return (
    <div className="w-full" style={{ height }}>
      <svg
        viewBox={`0 0 ${data.length * 60} 200`}
        className="w-full h-full"
        preserveAspectRatio="xMidYMid meet"
      >
        {/* Grid lines */}
        {[0, 0.25, 0.5, 0.75, 1].map((pct) => {
          const y = 200 - pct * 160;
          return (
            <g key={pct}>
              <line
                x1={0}
                y1={y}
                x2={data.length * 60}
                y2={y}
                stroke="#e5e7eb"
                strokeWidth="1"
              />
              <text x={-5} y={y + 4} textAnchor="end" fontSize="10" fill="#9ca3af">
                {Math.round(maxVal * pct)}
              </text>
            </g>
          );
        })}

        {/* Bars */}
        {data.map((d, i) => {
          const x = i * 60 + 20;
          const plannedH = (d.planned / maxVal) * 160;
          const actualH = (d.actual / maxVal) * 160;

          return (
            <g key={i}>
              {/* Planned bar */}
              <rect
                x={x - barWidth / 2}
                y={200 - plannedH}
                width={barWidth * 0.8}
                height={plannedH}
                fill="#93c5fd"
                rx="2"
              />
              {/* Actual bar */}
              <rect
                x={x + barWidth / 2 + 2}
                y={200 - actualH}
                width={barWidth * 0.8}
                height={actualH}
                fill="#3b82f6"
                rx="2"
              />
              {/* Label */}
              <text
                x={x + barWidth / 2}
                y={195}
                textAnchor="middle"
                fontSize="9"
                fill="#6b7280"
              >
                {d.label}
              </text>
            </g>
          );
        })}

        {/* Legend */}
        <g transform="translate(10, 10)">
          <rect width="10" height="10" fill="#93c5fd" rx="2" />
          <text x="14" y="9" fontSize="10" fill="#6b7280">计划</text>
          <rect x="50" width="10" height="10" fill="#3b82f6" rx="2" />
          <text x="64" y="9" fontSize="10" fill="#6b7280">实际</text>
        </g>
      </svg>
    </div>
  );
}
