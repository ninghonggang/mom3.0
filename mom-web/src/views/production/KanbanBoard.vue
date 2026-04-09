<template>
  <div class="kanban-board">
    <el-row :gutter="20" class="stat-row">
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%)">
            <el-icon :size="28"><TrendCharts /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ formatNumber(dashboard.output_stats.today_output) }}</div>
            <div class="stat-label">今日产量</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%)">
            <el-icon :size="28"><Calendar /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ formatNumber(dashboard.output_stats.week_output) }}</div>
            <div class="stat-label">本周产量</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)">
            <el-icon :size="28"><DataAnalysis /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ formatNumber(dashboard.output_stats.month_output) }}</div>
            <div class="stat-label">本月产量</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)">
            <el-icon :size="28"><CircleCheck /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ dashboard.output_stats.qualified_rate.toFixed(1) }}%</div>
            <div class="stat-label">合格率</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="chart-row">
      <el-col :span="16">
        <el-card shadow="hover" class="chart-card">
          <template #header>
            <div class="card-header">
              <span>每小时产量趋势</span>
              <el-tag type="info" size="small">实时刷新</el-tag>
            </div>
          </template>
          <div ref="hourlyChartRef" style="height: 320px"></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover" class="chart-card">
          <template #header>
            <div class="card-header">
              <span>工单进度</span>
            </div>
          </template>
          <div ref="orderProgressChartRef" style="height: 320px"></div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="line-row">
      <el-col :span="24">
        <el-card shadow="hover" class="line-card">
          <template #header>
            <div class="card-header">
              <span>产线状态监控</span>
              <el-button type="primary" size="small" @click="loadData">
                <el-icon><Refresh /></el-icon> 刷新
              </el-button>
            </div>
          </template>
          <div class="line-grid">
            <div v-for="line in dashboard.production_lines" :key="line.line_id" class="line-item">
              <div class="line-header">
                <span class="line-name">{{ line.line_name }}</span>
                <el-tag :type="getStatusType(line.status)" size="small">{{ line.status_text }}</el-tag>
              </div>
              <div class="line-info">
                <span class="workshop">{{ line.workshop_name }}</span>
              </div>
              <div class="line-output">
                <div class="output-value">{{ formatNumber(line.output) }}</div>
                <div class="output-target">/ {{ formatNumber(line.target_output) }}</div>
              </div>
              <el-progress
                :percentage="Math.min(line.completion, 100)"
                :color="getProgressColor(line.completion)"
                :stroke-width="10"
              />
              <div class="line-footer">
                <span>完成率 {{ line.completion.toFixed(1) }}%</span>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import * as echarts from 'echarts'
import { getKanbanDashboard } from '@/api/production'
import { ElMessage } from 'element-plus'
import { TrendCharts, Calendar, DataAnalysis, CircleCheck, Refresh } from '@element-plus/icons-vue'

interface HourlyOutput {
  hour: string
  output: number
}

interface LineStatus {
  line_id: number
  line_name: string
  workshop_name: string
  status: string
  status_text: string
  output: number
  target_output: number
  completion: number
}

interface OutputStats {
  today_output: number
  week_output: number
  month_output: number
  qualified_rate: number
}

interface OrderProgress {
  total: number
  pending: number
  in_process: number
  completed: number
  cancelled: number
}

interface KanbanData {
  production_lines: LineStatus[]
  output_stats: OutputStats
  order_progress: OrderProgress
  hourly_output: HourlyOutput[]
}

const hourlyChartRef = ref<HTMLElement>()
const orderProgressChartRef = ref<HTMLElement>()
let hourlyChart: echarts.ECharts | null = null
let orderProgressChart: echarts.ECharts | null = null
let refreshTimer: ReturnType<typeof setInterval> | null = null

const dashboard = reactive<KanbanData>({
  production_lines: [],
  output_stats: { today_output: 0, week_output: 0, month_output: 0, qualified_rate: 0 },
  order_progress: { total: 0, pending: 0, in_process: 0, completed: 0, cancelled: 0 },
  hourly_output: []
})

const formatNumber = (num: number): string => {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + '万'
  }
  return num.toFixed(0)
}

const getStatusType = (status: string): string => {
  const map: Record<string, string> = {
    running: 'success',
    idle: 'warning',
    fault: 'danger'
  }
  return map[status] || 'info'
}

const getProgressColor = (percentage: number): string => {
  if (percentage >= 90) return '#67c23a'
  if (percentage >= 70) return '#409eff'
  if (percentage >= 50) return '#e6a23c'
  return '#f56c6c'
}

const initHourlyChart = () => {
  if (!hourlyChartRef.value) return
  hourlyChart = echarts.init(hourlyChartRef.value)
  updateHourlyChart()
}

const updateHourlyChart = () => {
  if (!hourlyChart) return
  const hours = dashboard.hourly_output.map(item => item.hour)
  const outputs = dashboard.hourly_output.map(item => item.output)

  hourlyChart.setOption({
    tooltip: {
      trigger: 'axis',
      formatter: '{b}: {c} 件'
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: hours,
      axisLabel: { interval: Math.floor(hours.length / 8) }
    },
    yAxis: {
      type: 'value',
      name: '产量'
    },
    series: [
      {
        name: '产量',
        type: 'bar',
        data: outputs,
        itemStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: '#409eff' },
            { offset: 1, color: '#79bbff' }
          ]),
          borderRadius: [4, 4, 0, 0]
        },
        barWidth: '60%'
      }
    ]
  })
}

const initOrderProgressChart = () => {
  if (!orderProgressChartRef.value) return
  orderProgressChart = echarts.init(orderProgressChartRef.value)
  updateOrderProgressChart()
}

const updateOrderProgressChart = () => {
  if (!orderProgressChart) return
  const { pending, in_process, completed, cancelled } = dashboard.order_progress

  orderProgressChart.setOption({
    tooltip: {
      trigger: 'item',
      formatter: '{b}: {c} ({d}%)'
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      top: 'middle'
    },
    series: [
      {
        name: '工单状态',
        type: 'pie',
        radius: ['45%', '70%'],
        center: ['60%', '50%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 8,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: true,
          formatter: '{b}\n{c}'
        },
        data: [
          { value: pending, name: '待生产', itemStyle: { color: '#909399' } },
          { value: in_process, name: '生产中', itemStyle: { color: '#409eff' } },
          { value: completed, name: '已完成', itemStyle: { color: '#67c23a' } },
          { value: cancelled, name: '已取消', itemStyle: { color: '#f56c6c' } }
        ]
      }
    ]
  })
}

const loadData = async () => {
  try {
    const res = await getKanbanDashboard()
    Object.assign(dashboard, res.data)
    updateHourlyChart()
    updateOrderProgressChart()
  } catch (error) {
    ElMessage.error('加载看板数据失败')
  }
}

const handleResize = () => {
  hourlyChart?.resize()
  orderProgressChart?.resize()
}

onMounted(() => {
  loadData()
  initHourlyChart()
  initOrderProgressChart()
  window.addEventListener('resize', handleResize)
  // 每30秒自动刷新
  refreshTimer = setInterval(loadData, 30000)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  if (refreshTimer) {
    clearInterval(refreshTimer)
  }
  hourlyChart?.dispose()
  orderProgressChart?.dispose()
})
</script>

<style scoped lang="scss">
.kanban-board {
  padding: 16px;

  .stat-row {
    margin-bottom: 20px;
  }

  .stat-card {
    display: flex;
    align-items: center;
    padding: 8px;
    border-radius: 12px;
    transition: transform 0.3s;

    &:hover {
      transform: translateY(-4px);
    }

    .stat-icon {
      width: 56px;
      height: 56px;
      border-radius: 12px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: #fff;
      flex-shrink: 0;
    }

    .stat-content {
      margin-left: 16px;

      .stat-value {
        font-size: 26px;
        font-weight: 700;
        color: #303133;
        line-height: 1.2;
      }

      .stat-label {
        font-size: 13px;
        color: #909399;
        margin-top: 4px;
      }
    }
  }

  .chart-row {
    margin-bottom: 20px;
  }

  .chart-card {
    border-radius: 12px;

    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      font-size: 16px;
      font-weight: 500;
    }
  }

  .line-card {
    border-radius: 12px;

    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      font-size: 16px;
      font-weight: 500;
    }
  }

  .line-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 16px;

    @media (max-width: 1200px) {
      grid-template-columns: repeat(3, 1fr);
    }

    @media (max-width: 768px) {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  .line-item {
    background: #f5f7fa;
    border-radius: 10px;
    padding: 16px;
    transition: all 0.3s;

    &:hover {
      background: #ecf5ff;
      transform: scale(1.02);
    }

    .line-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 8px;

      .line-name {
        font-size: 15px;
        font-weight: 600;
        color: #303133;
      }
    }

    .line-info {
      .workshop {
        font-size: 12px;
        color: #909399;
      }
    }

    .line-output {
      display: flex;
      align-items: baseline;
      margin: 12px 0 8px;

      .output-value {
        font-size: 22px;
        font-weight: 700;
        color: #409eff;
      }

      .output-target {
        font-size: 12px;
        color: #909399;
        margin-left: 4px;
      }
    }

    .line-footer {
      margin-top: 8px;
      font-size: 12px;
      color: #909399;
      text-align: right;
    }
  }
}
</style>
