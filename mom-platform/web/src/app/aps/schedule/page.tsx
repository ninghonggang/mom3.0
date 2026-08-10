import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

const mockScheduleTasks = [
  { id: '1', orderNo: 'MO-2024-0001', opName: '粗车外圆', workCenter: 'CNC-01', start: 0, duration: 12, color: 'bg-blue-500', progress: 100 },
  { id: '2', orderNo: 'MO-2024-0001', opName: '精车内孔', workCenter: 'CNC-02', start: 12, duration: 16, color: 'bg-blue-500', progress: 65 },
  { id: '3', orderNo: 'MO-2024-0001', opName: '钻孔攻牙', workCenter: 'DRILL-01', start: 28, duration: 8, color: 'bg-blue-500', progress: 0 },
  { id: '4', orderNo: 'MO-2024-0002', opName: '铣面', workCenter: 'CNC-02', start: 0, duration: 8, color: 'bg-green-500', progress: 100 },
  { id: '5', orderNo: 'MO-2024-0002', opName: '钻孔', workCenter: 'DRILL-01', start: 8, duration: 6, color: 'bg-green-500', progress: 100 },
  { id: '6', orderNo: 'MO-2024-0003', opName: '粗车', workCenter: 'CNC-03', start: 20, duration: 10, color: 'bg-purple-500', progress: 0 },
];

const workCenters = ['CNC-01', 'CNC-02', 'CNC-03', 'DRILL-01'];

export default function SchedulePage() {
  const totalHours = 40;
  const hourWidth = 100 / totalHours;

  return (
    <div className="space-y-6">
      <div className="page-header">
        <h1 className="page-title">排程甘特图</h1>
        <div className="flex items-center gap-2">
          <div className="flex items-center gap-1">
            <div className="h-3 w-3 rounded bg-blue-500" />
            <span className="text-xs text-gray-500">MO-0001</span>
          </div>
          <div className="flex items-center gap-1">
            <div className="h-3 w-3 rounded bg-green-500" />
            <span className="text-xs text-gray-500">MO-0002</span>
          </div>
          <div className="flex items-center gap-1">
            <div className="h-3 w-3 rounded bg-purple-500" />
            <span className="text-xs text-gray-500">MO-0003</span>
          </div>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle className="text-base">本周排程计划 (2025-01-06 ~ 2025-01-10)</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-x-auto">
            <div className="min-w-[800px]">
              {/* Time header */}
              <div className="mb-2 flex border-b pb-2">
                <div className="w-24 flex-shrink-0 text-xs font-medium text-gray-500">工作中心</div>
                <div className="flex flex-1">
                  {Array.from({ length: totalHours / 8 }, (_, day) => (
                    <div key={day} className="flex" style={{ width: `${hourWidth * 8}%` }}>
                      <div className="flex-1 text-center text-xs text-gray-400">
                        {['一','二','三','四','五'][day]}
                      </div>
                    </div>
                  ))}
                </div>
              </div>

              {/* Work center rows */}
              {workCenters.map((wc) => (
                <div key={wc} className="mb-1 flex items-center">
                  <div className="w-24 flex-shrink-0 py-2 text-xs font-medium text-gray-600">{wc}</div>
                  <div className="relative flex-1" style={{ height: 36 }}>
                    {mockScheduleTasks
                      .filter((t) => t.workCenter === wc)
                      .map((task) => (
                        <div
                          key={task.id}
                          className={`absolute top-1 h-7 rounded ${task.color} flex items-center px-2`}
                          style={{
                            left: `${task.start * hourWidth}%`,
                            width: `${task.duration * hourWidth}%`,
                          }}
                        >
                          <div className="flex w-full items-center justify-between">
                            <span className="text-[10px] text-white font-medium truncate">
                              {task.opName}
                            </span>
                            {task.progress > 0 && task.progress < 100 && (
                              <span className="text-[10px] text-white/80">{task.progress}%</span>
                            )}
                          </div>
                        </div>
                      ))}
                  </div>
                </div>
              ))}

              {/* Time grid lines */}
              <div className="mt-2 flex border-t pt-2">
                <div className="w-24 flex-shrink-0" />
                <div className="flex flex-1">
                  {Array.from({ length: totalHours / 2 }, (_, i) => (
                    <div key={i} className="text-[9px] text-gray-300" style={{ width: `${hourWidth * 2}%` }}>
                      {`${Math.floor(i / 1) * 2}h`}
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>
          <div className="mt-6 rounded-lg bg-gray-50 p-4 text-center text-sm text-gray-400">
            完整甘特图组件将在集成阶段引入专业排程可视化库（如 dhtmlx-gantt 或 frappe-gantt）
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
