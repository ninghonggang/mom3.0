<template>
  <div class="andon-stats">
    <el-row :gutter="16" class="stat-cards">
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-item">
            <div class="stat-value danger">{{ stats.today_calls || 0 }}</div>
            <div class="stat-label">今日呼叫</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-item">
            <div class="stat-value warning">{{ stats.pending_calls || 0 }}</div>
            <div class="stat-label">待响应</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-item">
            <div class="stat-value">{{ stats.avg_response_time || 0 }}<span class="unit">min</span></div>
            <div class="stat-label">平均响应时间</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-item">
            <div class="stat-value success">{{ stats.resolution_rate || 0 }}<span class="unit">%</span></div>
            <div class="stat-label">解决率</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="统计周期">
          <el-date-picker
            v-model="dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
            @change="handleDateChange"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadStats">查询</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-row :gutter="16">
      <!-- 呼叫类型分布 -->
      <el-col :span="12">
        <el-card>
          <template #header><span>呼叫类型分布</span></template>
          <div ref="typeChartRef" style="height: 280px"></div>
        </el-card>
      </el-col>
      <!-- 升级趋势 -->
      <el-col :span="12">
        <el-card>
          <template #header><span>每日呼叫趋势</span></template>
          <div ref="trendChartRef" style="height: 280px"></div>
        </el-card>
      </el-col>
    </el-row>

    <el-card style="margin-top: 16px">
      <template #header>
        <span>详细数据</span>
        <el-button type="primary" size="small" style="float: right" @click="loadStats">刷新</el-button>
      </template>
      <el-table v-loading="loading" :data="detailList">
        <el-table-column prop="andon_type" label="呼叫类型" width="120">
          <template #default="{ row }">
            <el-tag :type="getTypeTag(row.andon_type)">{{ row.andon_type_name || row.andon_type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="call_count" label="呼叫次数" width="100" />
        <el-table-column prop="avg_response_time" label="平均响应(min)" width="140" />
        <el-table-column prop="avg_handle_time" label="平均处理(min)" width="140" />
        <el-table-column prop="escalation_count" label="升级次数" width="100" />
        <el-table-column prop="max_level" label="最高等级" width="100">
          <template #default="{ row }">
            <el-tag type="danger" v-if="row.max_level > 0">L{{ row.max_level }}</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="resolved_count" label="已解决" width="100" />
        <el-table-column prop="carry_over_count" label="遗留" width="100" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import * as echarts from 'echarts'
import { getAndonStats } from '@/api/trace'

const loading = ref(false)
const stats = ref<any>({})
const detailList = ref<any[]>([])
const dateRange = ref<[string, string] | null>(null)
const searchForm = reactive({ start_date: '', end_date: '' })
const typeChartRef = ref<HTMLElement>()
const trendChartRef = ref<HTMLElement>()
let typeChart: echarts.ECharts | null = null
let trendChart: echarts.ECharts | null = null

const getTypeTag = (type: string) => {
  const map: Record<string, string> = {
    EQUIPMENT: 'danger', MATERIAL: 'warning', QUALITY: 'danger',
    TECHNICAL: 'info', TOOLING: 'info', SAFETY: 'warning', OTHER: 'info'
  }
  return map[type] || 'info'
}

const handleDateChange = (val: [string, string] | null) => {
  if (val) {
    searchForm.start_date = val[0]
    searchForm.end_date = val[1]
  } else {
    searchForm.start_date = ''
    searchForm.end_date = ''
  }
}

const loadStats = async () => {
  loading.value = true
  try {
    const params: any = {}
    if (searchForm.start_date) params.start_date = searchForm.start_date
    if (searchForm.end_date) params.end_date = searchForm.end_date
    const res = await getAndonStats(params)
    const data = res.data || res
    stats.value = data

    if (data.by_type && Array.isArray(data.by_type)) {
      detailList.value = data.by_type
      renderTypeChart(data.by_type)
    }
    if (data.daily_trend && Array.isArray(data.daily_trend)) {
      renderTrendChart(data.daily_trend)
    }
  } catch (e: any) {
    ElMessage.error(e.message || '加载失败')
  } finally {
    loading.value = false
  }
}

const renderTypeChart = (data: any[]) => {
  if (!typeChartRef.value) return
  if (!typeChart) typeChart = echarts.init(typeChartRef.value)
  const option = {
    tooltip: { trigger: 'item' },
    legend: { bottom: 0 },
    series: [{
      type: 'pie',
      radius: ['40%', '70%'],
      data: data.map((d: any) => ({
        name: d.andon_type_name || d.andon_type,
        value: d.call_count
      }))
    }]
  }
  typeChart.setOption(option)
}

const renderTrendChart = (data: any[]) => {
  if (!trendChartRef.value) return
  if (!trendChart) trendChart = echarts.init(trendChartRef.value)
  const option = {
    tooltip: { trigger: 'axis' },
    legend: { bottom: 0 },
    grid: { top: 10, right: 20, bottom: 40, left: 50 },
    xAxis: { type: 'category', data: data.map((d: any) => d.date) },
    yAxis: { type: 'value' },
    series: [
      { name: '呼叫次数', type: 'bar', data: data.map((d: any) => d.call_count) },
      { name: '升级次数', type: 'line', data: data.map((d: any) => d.escalation_count) }
    ]
  }
  trendChart.setOption(option)
}

onMounted(() => {
  loadStats()
  window.addEventListener('resize', () => {
    typeChart?.resize()
    trendChart?.resize()
  })
})

onUnmounted(() => {
  typeChart?.dispose()
  trendChart?.dispose()
})
</script>

<style scoped lang="scss">
.andon-stats {
  .stat-cards { margin-bottom: 16px; }
  .stat-item { text-align: center; padding: 8px 0; }
  .stat-value {
    font-size: 32px;
    font-weight: bold;
    line-height: 1.2;
    margin-bottom: 4px;
    &.danger { color: #f56c6c; }
    &.warning { color: #e6a23c; }
    &.success { color: #67c23a; }
    .unit { font-size: 14px; font-weight: normal; margin-left: 2px; }
  }
  .stat-label { color: #909399; font-size: 14px; }
  .search-card { margin-bottom: 16px; }
}
</style>
