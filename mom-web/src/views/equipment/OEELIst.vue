<template>
  <div class="oee-container">
    <!-- OEE Dashboard Cards -->
    <el-row :gutter="16" class="dashboard-row">
      <el-col :span="6">
        <div class="dashboard-card availability">
          <div class="card-title">可用率</div>
          <div class="card-value">{{ dashboardData.availability }}%</div>
          <div class="card-sub">目标: ≥85%</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="dashboard-card performance">
          <div class="card-title">性能率</div>
          <div class="card-value">{{ dashboardData.performance }}%</div>
          <div class="card-sub">目标: ≥90%</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="dashboard-card quality">
          <div class="card-title">质量率</div>
          <div class="card-value">{{ dashboardData.quality }}%</div>
          <div class="card-sub">目标: ≥99%</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="dashboard-card oee">
          <div class="card-title">OEE综合效率</div>
          <div class="card-value">{{ dashboardData.oee }}%</div>
          <div class="card-sub">目标: ≥77%</div>
        </div>
      </el-col>
    </el-row>

    <!-- Search Form -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="设备">
          <el-select v-model="searchForm.equipment_id" placeholder="请选择设备" clearable filterable>
            <el-option v-for="eq in equipmentList" :key="eq.id" :label="eq.equipment_name" :value="eq.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="日期范围">
          <el-date-picker v-model="dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" value-format="YYYY-MM-DD" @change="handleDateChange" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Toolbar -->
    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('equipment:oee:add')" @click="handleCalculate">
        <el-icon><Plus /></el-icon>计算OEE
      </el-button>
    </el-card>

    <!-- OEE Table -->
    <el-card>
      <el-table v-loading="loading" :data="tableData" border stripe>
        <el-table-column prop="equipment_name" label="设备名称" width="150" />
        <el-table-column prop="record_date" label="日期" width="120" />
        <el-table-column prop="plan_time" label="计划时间(分钟)" width="120" />
        <el-table-column prop="run_time" label="运行时间" width="100" />
        <el-table-column prop="down_time" label="停机时间" width="100" />
        <el-table-column prop="output_qty" label="产出数量" width="100" />
        <el-table-column prop="qualified_qty" label="合格数量" width="100" />
        <el-table-column prop="availability" label="可用率%" width="100">
          <template #default="{ row }">
            <span :class="getRateClass(row.availability, 85)">{{ row.availability }}%</span>
          </template>
        </el-table-column>
        <el-table-column prop="performance" label="性能率%" width="100">
          <template #default="{ row }">
            <span :class="getRateClass(row.performance, 90)">{{ row.performance }}%</span>
          </template>
        </el-table-column>
        <el-table-column prop="quality" label="质量率%" width="100">
          <template #default="{ row }">
            <span :class="getRateClass(row.quality, 99)">{{ row.quality }}%</span>
          </template>
        </el-table-column>
        <el-table-column prop="oee" label="OEE%" width="100">
          <template #default="{ row }">
            <span :class="getOEEClass(row.oee)">{{ row.oee }}%</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('equipment:oee:view')" @click="handleView(row)">详情</el-button>
            <el-button link type="danger" v-if="hasPermission('equipment:oee:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="loadData"
          @current-change="loadData"
        />
      </div>
    </el-card>

    <!-- Trend Chart -->
    <el-card class="chart-card">
      <template #header>
        <span>OEE趋势图</span>
      </template>
      <div ref="chartRef" class="chart-container"></div>
    </el-card>

    <!-- Calculate Dialog -->
    <el-dialog v-model="dialogVisible" title="计算OEE" width="600px">
      <el-form ref="formRef" :model="formData" label-width="120px">
        <el-form-item label="设备" prop="equipment_id">
          <el-select v-model="formData.equipment_id" placeholder="请选择设备" filterable @change="handleEquipmentChange">
            <el-option v-for="eq in equipmentList" :key="eq.id" :label="eq.equipment_name" :value="eq.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="记录日期" prop="record_date">
          <el-date-picker v-model="formData.record_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="计划时间(分钟)" prop="plan_time">
          <el-input-number v-model="formData.plan_time" :min="0" />
        </el-form-item>
        <el-form-item label="运行时间(分钟)" prop="run_time">
          <el-input-number v-model="formData.run_time" :min="0" />
        </el-form-item>
        <el-form-item label="停机时间(分钟)" prop="down_time">
          <el-input-number v-model="formData.down_time" :min="0" />
        </el-form-item>
        <el-form-item label="空闲时间(分钟)" prop="idle_time">
          <el-input-number v-model="formData.idle_time" :min="0" />
        </el-form-item>
        <el-form-item label="计划停机时间" prop="plan_stop_time">
          <el-input-number v-model="formData.plan_stop_time" :min="0" />
        </el-form-item>
        <el-form-item label="产出数量" prop="output_qty">
          <el-input-number v-model="formData.output_qty" :min="0" />
        </el-form-item>
        <el-form-item label="合格数量" prop="qualified_qty">
          <el-input-number v-model="formData.qualified_qty" :min="0" />
        </el-form-item>
        <el-form-item label="理论产量" prop="theoretical_output">
          <el-input-number v-model="formData.theoretical_output" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- Detail Dialog -->
    <el-dialog v-model="detailVisible" title="OEE详情" width="700px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="设备名称">{{ detailData.equipment_name }}</el-descriptions-item>
        <el-descriptions-item label="设备编码">{{ detailData.equipment_code }}</el-descriptions-item>
        <el-descriptions-item label="记录日期">{{ detailData.record_date }}</el-descriptions-item>
        <el-descriptions-item label="计划时间">{{ detailData.plan_time }} 分钟</el-descriptions-item>
        <el-descriptions-item label="运行时间">{{ detailData.run_time }} 分钟</el-descriptions-item>
        <el-descriptions-item label="停机时间">{{ detailData.down_time }} 分钟</el-descriptions-item>
        <el-descriptions-item label="空闲时间">{{ detailData.idle_time }} 分钟</el-descriptions-item>
        <el-descriptions-item label="计划停机">{{ detailData.plan_stop_time }} 分钟</el-descriptions-item>
        <el-descriptions-item label="产出数量">{{ detailData.output_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格数量">{{ detailData.qualified_qty }}</el-descriptions-item>
        <el-descriptions-item label="可用率">{{ detailData.availability }}%</el-descriptions-item>
        <el-descriptions-item label="性能率">{{ detailData.performance }}%</el-descriptions-item>
        <el-descriptions-item label="质量率">{{ detailData.quality }}%</el-descriptions-item>
        <el-descriptions-item label="OEE">
          <el-tag type="success">{{ detailData.oee }}%</el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import * as echarts from 'echarts'
import { getOEEList, calculateOEE, getOEEChart, deleteOEE, getEquipmentList } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const detailVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref()
const chartRef = ref()
let chartInstance: echarts.ECharts | null = null

const searchForm = reactive({
  equipment_id: null as number | null,
  start_date: '',
  end_date: ''
})

const dateRange = ref<[string, string] | null>(null)
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive({
  equipment_id: null as number | null,
  equipment_code: '',
  equipment_name: '',
  record_date: '',
  plan_time: 480,
  run_time: 0,
  down_time: 0,
  idle_time: 0,
  plan_stop_time: 0,
  output_qty: 0,
  qualified_qty: 0,
  theoretical_output: 0
})

const detailData = reactive<any>({})
const equipmentList = ref<any[]>([])

const dashboardData = computed(() => {
  if (tableData.value.length === 0) {
    return { availability: 0, performance: 0, quality: 0, oee: 0 }
  }
  const avg = (key: string) => {
    const sum = tableData.value.reduce((acc, row) => acc + (parseFloat(row[key]) || 0), 0)
    return Math.round((sum / tableData.value.length) * 100) / 100
  }
  return {
    availability: avg('availability'),
    performance: avg('performance'),
    quality: avg('quality'),
    oee: avg('oee')
  }
})

const getRateClass = (value: number, target: number) => {
  if (value >= target) return 'rate-good'
  if (value >= target * 0.8) return 'rate-warning'
  return 'rate-bad'
}

const getOEEClass = (value: number) => {
  if (value >= 85) return 'rate-good'
  if (value >= 70) return 'rate-warning'
  return 'rate-bad'
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

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.pageSize,
      ...searchForm
    }
    const res = await getOEEList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadEquipmentList = async () => {
  try {
    const res = await getEquipmentList({ page: 1, page_size: 1000 })
    equipmentList.value = res.data.list || []
  } catch (e) {
    console.error('Failed to load equipment list', e)
  }
}

const loadChartData = async () => {
  if (!chartRef.value) return
  try {
    const params = {
      ...searchForm
    }
    const res = await getOEEChart(params)
    const list = res.data.list || []

    if (chartInstance) {
      chartInstance.dispose()
    }
    chartInstance = echarts.init(chartRef.value)

    const dates = list.map((item: any) => item.record_date)
    const oeeData = list.map((item: any) => item.oee)
    const availData = list.map((item: any) => item.availability)
    const perfData = list.map((item: any) => item.performance)
    const qualData = list.map((item: any) => item.quality)

    const option = {
      tooltip: { trigger: 'axis' },
      legend: { data: ['OEE', '可用率', '性能率', '质量率'] },
      xAxis: { type: 'category', data: dates },
      yAxis: { type: 'value', max: 100, axisLabel: { formatter: '{value}%' } },
      series: [
        { name: 'OEE', type: 'line', data: oeeData, smooth: true },
        { name: '可用率', type: 'line', data: availData, smooth: true },
        { name: '性能率', type: 'line', data: perfData, smooth: true },
        { name: '质量率', type: 'line', data: qualData, smooth: true }
      ]
    }
    chartInstance.setOption(option)
  } catch (e) {
    console.error('Failed to load chart', e)
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
  loadChartData()
}

const handleReset = () => {
  searchForm.equipment_id = null
  searchForm.start_date = ''
  searchForm.end_date = ''
  dateRange.value = null
  handleSearch()
}

const handleCalculate = () => {
  Object.assign(formData, {
    equipment_id: null,
    equipment_code: '',
    equipment_name: '',
    record_date: '',
    plan_time: 480,
    run_time: 0,
    down_time: 0,
    idle_time: 0,
    plan_stop_time: 0,
    output_qty: 0,
    qualified_qty: 0,
    theoretical_output: 0
  })
  dialogVisible.value = true
}

const handleEquipmentChange = (id: number) => {
  const eq = equipmentList.value.find(e => e.id === id)
  if (eq) {
    formData.equipment_code = eq.equipment_code
    formData.equipment_name = eq.equipment_name
  }
}

const handleSubmit = async () => {
  if (!formData.equipment_id) {
    ElMessage.warning('请选择设备')
    return
  }
  if (!formData.record_date) {
    ElMessage.warning('请选择日期')
    return
  }
  submitLoading.value = true
  try {
    await calculateOEE(formData)
    ElMessage.success('OEE计算成功')
    dialogVisible.value = false
    loadData()
    loadChartData()
  } finally {
    submitLoading.value = false
  }
}

const handleView = (row: any) => {
  Object.assign(detailData, row)
  detailVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该OEE记录吗？', '提示', { type: 'warning' })
    await deleteOEE(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => {
  loadEquipmentList()
  loadData()
  window.addEventListener('resize', () => chartInstance?.resize())
})

onUnmounted(() => {
  window.removeEventListener('resize', () => chartInstance?.resize())
  chartInstance?.dispose()
})
</script>

<style scoped lang="scss">
.oee-container {
  .dashboard-row { margin-bottom: 16px; }
  .dashboard-card {
    padding: 20px;
    border-radius: 8px;
    color: #fff;
    .card-title { font-size: 14px; opacity: 0.9; }
    .card-value { font-size: 32px; font-weight: bold; margin: 8px 0; }
    .card-sub { font-size: 12px; opacity: 0.7; }
    &.availability { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); }
    &.performance { background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); }
    &.quality { background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); }
    &.oee { background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%); }
  }
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .chart-card { margin-top: 16px; }
  .chart-container { height: 350px; }
  .rate-good { color: #67c23a; font-weight: bold; }
  .rate-warning { color: #e6a23c; font-weight: bold; }
  .rate-bad { color: #f56c6c; font-weight: bold; }
}
</style>
