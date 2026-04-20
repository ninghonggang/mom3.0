<template>
  <div class="alert-statistics">
    <el-row :gutter="16" class="stat-cards">
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-icon danger">
            <el-icon><WarningFilled /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stats.today_count || 0 }}</div>
            <div class="stat-label">今日告警</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-icon warning">
            <el-icon><Bell /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stats.total_count || 0 }}</div>
            <div class="stat-label">总告警数</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-icon info">
            <el-icon><Clock /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stats.pending_count || 0 }}</div>
            <div class="stat-label">待处理</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-icon success">
            <el-icon><CircleCheckFilled /></el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stats.resolved_count || 0 }}</div>
            <div class="stat-label">已解决</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="16" class="chart-row">
      <el-col :span="16">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>告警趋势</span>
              <el-radio-group v-model="trendType" size="small">
                <el-radio-button value="week">近7天</el-radio-button>
                <el-radio-button value="month">近30天</el-radio-button>
              </el-radio-group>
            </div>
          </template>
          <div class="chart-container" ref="trendChartRef"></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <span>告警级别分布</span>
          </template>
          <div class="chart-container" ref="severityChartRef"></div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="16" class="chart-row">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>告警类型分布</span>
          </template>
          <div class="chart-container" ref="typeChartRef"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>告警来源模块分布</span>
          </template>
          <div class="chart-container" ref="moduleChartRef"></div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted, watch } from 'vue'
import * as echarts from 'echarts'
import { getAlertStatistics } from '@/api/alert'
import { WarningFilled, Bell, Clock, CircleCheckFilled } from '@element-plus/icons-vue'

const trendType = ref('week')
const trendChartRef = ref<HTMLDivElement>()
const severityChartRef = ref<HTMLDivElement>()
const typeChartRef = ref<HTMLDivElement>()
const moduleChartRef = ref<HTMLDivElement>()

let trendChart: echarts.ECharts | null = null
let severityChart: echarts.ECharts | null = null
let typeChart: echarts.ECharts | null = null
let moduleChart: echarts.ECharts | null = null

const stats = ref<any>({
  today_count: 0,
  total_count: 0,
  pending_count: 0,
  resolved_count: 0
})

const trendData = ref<any[]>([])
const severityData = ref<any[]>([])
const typeData = ref<any[]>([])
const moduleData = ref<any[]>([])

const initCharts = () => {
  if (trendChartRef.value) {
    trendChart = echarts.init(trendChartRef.value)
  }
  if (severityChartRef.value) {
    severityChart = echarts.init(severityChartRef.value)
  }
  if (typeChartRef.value) {
    typeChart = echarts.init(typeChartRef.value)
  }
  if (moduleChartRef.value) {
    moduleChart = echarts.init(moduleChartRef.value)
  }
}

const updateTrendChart = () => {
  if (!trendChart) return
  const dates = trendData.value.map((item: any) => item.date)
  const counts = trendData.value.map((item: any) => item.count)

  trendChart.setOption({
    tooltip: { trigger: 'axis' },
    xAxis: { type: 'category', data: dates, boundaryGap: false },
    yAxis: { type: 'value', name: '告警数' },
    series: [
      {
        name: '告警数',
        type: 'line',
        data: counts,
        smooth: true,
        areaStyle: { opacity: 0.3 },
        itemStyle: { color: '#409eff' }
      }
    ],
    grid: { left: 50, right: 20, top: 20, bottom: 30 }
  })
}

const updateSeverityChart = () => {
  if (!severityChart) return
  severityChart.setOption({
    tooltip: { trigger: 'item', formatter: '{b}: {c} ({d}%)' },
    legend: { orient: 'vertical', left: 'left' },
    series: [
      {
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: { borderRadius: 4, borderColor: '#fff', borderWidth: 2 },
        label: { show: false },
        emphasis: { label: { show: true, fontSize: 14, fontWeight: 'bold' } },
        data: severityData.value
      }
    ]
  })
}

const updateTypeChart = () => {
  if (!typeChart) return
  typeChart.setOption({
    tooltip: { trigger: 'item', formatter: '{b}: {c} ({d}%)' },
    legend: { orient: 'vertical', left: 'left' },
    series: [
      {
        type: 'pie',
        radius: '60%',
        data: typeData.value,
        itemStyle: { borderRadius: 4, borderColor: '#fff', borderWidth: 2 },
        emphasis: { label: { show: true, fontSize: 14, fontWeight: 'bold' } }
      }
    ]
  })
}

const updateModuleChart = () => {
  if (!moduleChart) return
  moduleChart.setOption({
    tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
    grid: { left: 100, right: 20, top: 20, bottom: 30 },
    xAxis: { type: 'value' },
    yAxis: { type: 'category', data: moduleData.value.map((item: any) => item.name) },
    series: [
      {
        type: 'bar',
        data: moduleData.value.map((item: any) => item.value),
        itemStyle: { color: '#67c23a', borderRadius: [0, 4, 4, 0] }
      }
    ]
  })
}

const loadStatistics = async () => {
  try {
    const res = await getAlertStatistics()
    stats.value = {
      today_count: res.data.today_count || 0,
      total_count: res.data.total_count || 0,
      pending_count: res.data.pending_count || 0,
      resolved_count: res.data.resolved_count || 0
    }
    // 趋势数据
    trendData.value = res.data.trend_data || []
    // 级别分布
    severityData.value = res.data.severity_distribution || []
    // 类型分布
    typeData.value = res.data.type_distribution || []
    // 模块分布
    moduleData.value = res.data.module_distribution || []

    updateTrendChart()
    updateSeverityChart()
    updateTypeChart()
    updateModuleChart()
  } catch (error) {
    console.error('Failed to load statistics:', error)
  }
}

const handleResize = () => {
  trendChart?.resize()
  severityChart?.resize()
  typeChart?.resize()
  moduleChart?.resize()
}

watch(trendType, () => {
  loadStatistics()
})

onMounted(() => {
  initCharts()
  loadStatistics()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  trendChart?.dispose()
  severityChart?.dispose()
  typeChart?.dispose()
  moduleChart?.dispose()
})
</script>

<style scoped lang="scss">
.alert-statistics {
  .stat-cards { margin-bottom: 16px; }

  .stat-card {
    display: flex;
    align-items: center;
    padding: 8px;

    .stat-icon {
      width: 60px;
      height: 60px;
      border-radius: 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 16px;

      .el-icon {
        font-size: 28px;
        color: #fff;
      }

      &.danger { background: linear-gradient(135deg, #f56c6c, #e64b4b); }
      &.warning { background: linear-gradient(135deg, #e6a23c, #d89030); }
      &.info { background: linear-gradient(135deg, #409eff, #3080e0); }
      &.success { background: linear-gradient(135deg, #67c23a, #56a82a); }
    }

    .stat-content {
      flex: 1;

      .stat-value {
        font-size: 28px;
        font-weight: bold;
        color: #303133;
        line-height: 1.2;
      }

      .stat-label {
        font-size: 14px;
        color: #909399;
        margin-top: 4px;
      }
    }
  }

  .chart-row { margin-bottom: 16px; }

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .chart-container {
    height: 280px;
    width: 100%;
  }
}
</style>
