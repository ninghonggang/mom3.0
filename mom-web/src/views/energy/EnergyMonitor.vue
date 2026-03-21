<template>
  <div class="energy-monitor">
    <el-row :gutter="16" class="stats-row">
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background: #409eff">
              <el-icon :size="30"><Monitor /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ stats.today_power }}</div>
              <div class="stat-label">今日用电(kWh)</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background: #67c23a">
              <el-icon :size="30"><TrendCharts /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ stats.month_power }}</div>
              <div class="stat-label">本月用电(kWh)</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background: #e6a23c">
              <el-icon :size="30"><Coin /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ stats.today_cost }}</div>
              <div class="stat-label">今日成本(元)</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background: #f56c6c">
              <el-icon :size="30"><Warning /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ stats.alarm_count }}</div>
              <div class="stat-label">告警数量</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="16" style="margin-top: 16px">
      <el-col :span="16">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>用电趋势</span>
              <el-radio-group v-model="trendType" size="small">
                <el-radio-button value="day">日</el-radio-button>
                <el-radio-button value="month">月</el-radio-button>
                <el-radio-button value="year">年</el-radio-button>
              </el-radio-group>
            </div>
          </template>
          <div v-loading="chartLoading" style="height: 300px">
            <div ref="trendChartRef" style="height: 100%"></div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <span>能耗占比</span>
          </template>
          <div v-loading="pieLoading" style="height: 300px">
            <div ref="pieChartRef" style="height: 100%"></div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card style="margin-top: 16px">
      <template #header>
        <div class="card-header">
          <span>实时数据</span>
          <el-button type="primary" size="small" @click="loadData">
            <el-icon><Refresh /></el-icon>刷新
          </el-button>
        </div>
      </template>
      <el-table :data="realtimeData" v-loading="tableLoading">
        <el-table-column prop="equipment_name" label="设备名称" min-width="150" />
        <el-table-column prop="meter_no" label="电表编号" width="120" />
        <el-table-column prop="voltage" label="电压(V)" width="100" />
        <el-table-column prop="current" label="电流(A)" width="100" />
        <el-table-column prop="power" label="功率(kW)" width="100" />
        <el-table-column prop="energy" label="累计电量(kWh)" width="140" />
        <el-table-column prop="update_time" label="更新时间" width="180" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '在线' : '离线' }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import * as echarts from 'echarts'
import { getEnergyStats, getEnergyTrend } from '@/api/trace'

const stats = reactive({
  today_power: '0',
  month_power: '0',
  today_cost: '0',
  alarm_count: '0'
})

const trendType = ref('day')
const chartLoading = ref(false)
const pieLoading = ref(false)
const tableLoading = ref(false)
const trendChartRef = ref<HTMLElement>()
const pieChartRef = ref<HTMLElement>()
const realtimeData = ref<any[]>([])

let trendChart: echarts.ECharts | null = null
let pieChart: echarts.ECharts | null = null
let refreshTimer: number | null = null

const initCharts = () => {
  if (trendChartRef.value) {
    trendChart = echarts.init(trendChartRef.value)
    trendChart.setOption({
      tooltip: { trigger: 'axis' },
      xAxis: { type: 'category', data: ['00:00', '04:00', '08:00', '12:00', '16:00', '20:00', '24:00'] },
      yAxis: { type: 'value' },
      series: [{ data: [120, 80, 150, 280, 220, 190, 130], type: 'line', smooth: true, areaStyle: {} }]
    })
  }

  if (pieChartRef.value) {
    pieChart = echarts.init(pieChartRef.value)
    pieChart.setOption({
      tooltip: { trigger: 'item' },
      series: [{
        type: 'pie',
        radius: ['40%', '70%'],
        data: [
          { value: 1048, name: '照明' },
          { value: 735, name: '设备' },
          { value: 580, name: '空调' },
          { value: 484, name: '其他' }
        ]
      }]
    })
  }
}

const loadData = async () => {
  tableLoading.value = true
  try {
    const res = await getEnergyStats({})
    Object.assign(stats, res.data)
    realtimeData.value = res.data.realtime || []
  } finally {
    tableLoading.value = false
  }
}

onMounted(() => {
  loadData()
  initCharts()
  refreshTimer = window.setInterval(loadData, 30000)
})

onUnmounted(() => {
  if (refreshTimer) clearInterval(refreshTimer)
  trendChart?.dispose()
  pieChart?.dispose()
})
</script>

<style scoped lang="scss">
.energy-monitor {
  .stats-row { margin-top: 0 }
  .stat-card {
    display: flex;
    align-items: center;
    gap: 16px;
  }
  .stat-icon {
    width: 60px;
    height: 60px;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
  }
  .stat-content { flex: 1; }
  .stat-value { font-size: 24px; font-weight: bold; }
  .stat-label { font-size: 14px; color: #999; margin-top: 4px; }
  .card-header { display: flex; justify-content: space-between; align-items: center; }
}
</style>
