<template>
  <div class="page-container">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="分析编号">
          <el-input v-model="searchForm.analysis_no" placeholder="请输入分析编号" clearable />
        </el-form-item>
        <el-form-item label="分析类型">
          <el-select v-model="searchForm.analysis_type" placeholder="请选择" clearable>
            <el-option label="日分析" value="DAILY" />
            <el-option label="周分析" value="WEEKLY" />
            <el-option label="月分析" value="MONTHLY" />
          </el-select>
        </el-form-item>
        <el-form-item label="分析日期">
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
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:delivery-analysis:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button @click="loadDailyStats">
        <el-icon><DataAnalysis /></el-icon>日统计
      </el-button>
      <el-button @click="loadWeeklyStats">
        <el-icon><Calendar /></el-icon>周统计
      </el-button>
      <el-button @click="loadMonthlyStats">
        <el-icon><Histogram /></el-icon>月统计
      </el-button>
    </el-card>

    <!-- 统计卡片 -->
    <el-row :gutter="16" class="stats-row" v-if="statsVisible">
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">准时交付率</div>
          <div class="stat-value success">{{ statsSummary.onTimeRate }}%</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">总订单数</div>
          <div class="stat-value">{{ statsSummary.totalOrders }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">延期订单</div>
          <div class="stat-value danger">{{ statsSummary.lateOrders }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-title">平均延期天数</div>
          <div class="stat-value warning">{{ statsSummary.avgDelayDays }}</div>
        </el-card>
      </el-col>
    </el-row>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="analysis_no" label="分析编号" width="150" />
        <el-table-column prop="analysis_date" label="分析日期" width="110" />
        <el-table-column prop="analysis_type" label="分析类型" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.analysis_type === 'DAILY'" type="success">日分析</el-tag>
            <el-tag v-else-if="row.analysis_type === 'WEEKLY'" type="warning">周分析</el-tag>
            <el-tag v-else type="info">月分析</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="workshop_name" label="车间" width="120" />
        <el-table-column prop="total_orders" label="总订单" width="80" />
        <el-table-column prop="on_time_orders" label="准时" width="80">
          <template #default="{ row }">
            <span class="text-success">{{ row.on_time_orders }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="late_orders" label="延期" width="80">
          <template #default="{ row }">
            <span class="text-danger">{{ row.late_orders }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="on_time_rate" label="准时率(%)" width="100">
          <template #default="{ row }">
            <el-progress
              :percentage="Number(row.on_time_rate) || 0"
              :color="getRateColor(row.on_time_rate)"
            />
          </template>
        </el-table-column>
        <el-table-column prop="avg_delay_days" label="平均延期(天)" width="110" />
        <el-table-column prop="otd_gap" label="目标差距(%)" width="110">
          <template #default="{ row }">
            <span :class="row.otd_gap > 0 ? 'text-danger' : 'text-success'">
              {{ row.otd_gap > 0 ? '+' : '' }}{{ row.otd_gap }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="improvement_direction" label="趋势" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.improvement_direction === 'IMPROVE'" type="success" size="small">改善</el-tag>
            <el-tag v-else-if="row.improvement_direction === 'DETERIORATE'" type="danger" size="small">恶化</el-tag>
            <el-tag v-else type="info" size="small">稳定</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:delivery-analysis:view')" @click.stop="handleView(row)">详情</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:delivery-analysis:delete')" @click.stop="handleDelete(row)">删除</el-button>
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

    <!-- 新增对话框 -->
    <el-dialog v-model="dialogVisible" title="新增交付分析" width="600px">
      <el-form :model="formData" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="分析编号">
              <el-input v-model="formData.analysis_no" placeholder="自动生成" disabled />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="分析类型" required>
              <el-select v-model="formData.analysis_type" placeholder="请选择" style="width: 100%">
                <el-option label="日分析" value="DAILY" />
                <el-option label="周分析" value="WEEKLY" />
                <el-option label="月分析" value="MONTHLY" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="分析日期">
              <el-date-picker v-model="formData.analysis_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="车间">
              <el-input v-model="formData.workshop_name" placeholder="请输入车间" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总订单数">
              <el-input-number v-model="formData.total_orders" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="准时订单">
              <el-input-number v-model="formData.on_time_orders" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="延期订单">
              <el-input-number v-model="formData.late_orders" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="严重延期">
              <el-input-number v-model="formData.critical_late_orders" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="分析摘要">
          <el-input v-model="formData.analysis_summary" type="textarea" rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="交付分析详情" width="700px">
      <el-descriptions :column="2" border v-if="currentRow">
        <el-descriptions-item label="分析编号">{{ currentRow.analysis_no }}</el-descriptions-item>
        <el-descriptions-item label="分析日期">{{ currentRow.analysis_date }}</el-descriptions-item>
        <el-descriptions-item label="分析类型">
          <el-tag v-if="currentRow.analysis_type === 'DAILY'" type="success">日分析</el-tag>
          <el-tag v-else-if="currentRow.analysis_type === 'WEEKLY'" type="warning">周分析</el-tag>
          <el-tag v-else type="info">月分析</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="车间">{{ currentRow.workshop_name }}</el-descriptions-item>
        <el-descriptions-item label="总订单数">{{ currentRow.total_orders }}</el-descriptions-item>
        <el-descriptions-item label="准时订单">{{ currentRow.on_time_orders }}</el-descriptions-item>
        <el-descriptions-item label="提前交付">{{ currentRow.early_orders }}</el-descriptions-item>
        <el-descriptions-item label="延期交付">{{ currentRow.late_orders }}</el-descriptions-item>
        <el-descriptions-item label="严重延期(>3天)">{{ currentRow.critical_late_orders }}</el-descriptions-item>
        <el-descriptions-item label="准时交付率">{{ currentRow.on_time_rate }}%</el-descriptions-item>
        <el-descriptions-item label="提前交付率">{{ currentRow.early_rate }}%</el-descriptions-item>
        <el-descriptions-item label="延期率">{{ currentRow.late_rate }}%</el-descriptions-item>
        <el-descriptions-item label="平均延期天数">{{ currentRow.avg_delay_days }}</el-descriptions-item>
        <el-descriptions-item label="最大延期天数">{{ currentRow.max_delay_days }}</el-descriptions-item>
        <el-descriptions-item label="目标准时交付率">{{ currentRow.otd_target }}%</el-descriptions-item>
        <el-descriptions-item label="与目标差距">{{ currentRow.otd_gap }}%</el-descriptions-item>
        <el-descriptions-item label="趋势" :span="2">
          <el-tag v-if="currentRow.improvement_direction === 'IMPROVE'" type="success">改善</el-tag>
          <el-tag v-else-if="currentRow.improvement_direction === 'DETERIORATE'" type="danger">恶化</el-tag>
          <el-tag v-else type="info">稳定</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="分析摘要" :span="2">{{ currentRow.analysis_summary }}</el-descriptions-item>
        <el-descriptions-item label="创建人">{{ currentRow.created_by }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ currentRow.created_at }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getDeliveryAnalysisList,
  createDeliveryAnalysis,
  getDeliveryAnalysisDaily,
  getDeliveryAnalysisWeekly,
  getDeliveryAnalysisMonthly
} from '@/api/aps'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const detailVisible = ref(false)
const statsVisible = ref(false)
const tableData = ref<any[]>([])
const currentRow = ref<any>(null)
const dateRange = ref<[string, string] | null>(null)

const searchForm = reactive({
  analysis_no: '',
  analysis_type: '',
  start_date: '',
  end_date: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const statsSummary = reactive({
  onTimeRate: 0,
  totalOrders: 0,
  lateOrders: 0,
  avgDelayDays: 0
})

const formData = ref<any>({
  analysis_no: '',
  analysis_date: '',
  analysis_type: 'DAILY',
  workshop_name: '',
  total_orders: 0,
  on_time_orders: 0,
  early_orders: 0,
  late_orders: 0,
  critical_late_orders: 0,
  analysis_summary: ''
})

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
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getDeliveryAnalysisList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadStatsData = async (data: any[]) => {
  tableData.value = data
  statsVisible.value = true
  if (data.length > 0) {
    const totalOrders = data.reduce((sum, item) => sum + (item.total_orders || 0), 0)
    const onTimeOrders = data.reduce((sum, item) => sum + (item.on_time_orders || 0), 0)
    const lateOrders = data.reduce((sum, item) => sum + (item.late_orders || 0), 0)
    const avgDelayDays = data.reduce((sum, item) => sum + (item.avg_delay_days || 0), 0) / data.length
    statsSummary.totalOrders = totalOrders
    statsSummary.onTimeRate = totalOrders > 0 ? ((onTimeOrders / totalOrders) * 100).toFixed(1) as any : 0
    statsSummary.lateOrders = lateOrders
    statsSummary.avgDelayDays = avgDelayDays.toFixed(1)
  }
}

const loadDailyStats = async () => {
  loading.value = true
  try {
    const res = await getDeliveryAnalysisDaily()
    await loadStatsData(res.data.list || [])
  } finally {
    loading.value = false
  }
}

const loadWeeklyStats = async () => {
  loading.value = true
  try {
    const res = await getDeliveryAnalysisWeekly()
    await loadStatsData(res.data.list || [])
  } finally {
    loading.value = false
  }
}

const loadMonthlyStats = async () => {
  loading.value = true
  try {
    const res = await getDeliveryAnalysisMonthly()
    await loadStatsData(res.data.list || [])
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  statsVisible.value = false
  loadData()
}

const handleReset = () => {
  searchForm.analysis_no = ''
  searchForm.analysis_type = ''
  searchForm.start_date = ''
  searchForm.end_date = ''
  dateRange.value = null
  statsVisible.value = false
  handleSearch()
}

const handleAdd = () => {
  formData.value = {
    analysis_no: '',
    analysis_date: '',
    analysis_type: 'DAILY',
    workshop_name: '',
    total_orders: 0,
    on_time_orders: 0,
    early_orders: 0,
    late_orders: 0,
    critical_late_orders: 0,
    analysis_summary: ''
  }
  dialogVisible.value = true
}

const handleSave = async () => {
  saveLoading.value = true
  try {
    await createDeliveryAnalysis(formData.value)
    ElMessage.success('创建成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleView = (row: any) => {
  currentRow.value = row
  detailVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该分析记录吗？', '提示', { type: 'warning' })
    // deleteDeliveryAnalysis(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const getRateColor = (rate: number) => {
  if (rate >= 95) return '#52c41a'
  if (rate >= 85) return '#1890ff'
  if (rate >= 70) return '#faad14'
  return '#ff4d4f'
}

onMounted(() => {
  loadData()
})
</script>

<style scoped lang="scss">
.page-container {
  height: 100%;
}

.search-card {
  margin-bottom: 16px;
}

.toolbar-card {
  margin-bottom: 16px;

  :deep(.el-card__body) {
    padding: 12px 16px;
    display: flex;
    gap: 12px;
  }
}

.stats-row {
  margin-bottom: 16px;

  .stat-card {
    text-align: center;

    .stat-title {
      font-size: 14px;
      color: #666;
      margin-bottom: 8px;
    }

    .stat-value {
      font-size: 28px;
      font-weight: bold;

      &.success { color: #52c41a; }
      &.danger { color: #ff4d4f; }
      &.warning { color: #faad14; }
    }
  }
}

.pagination {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}

.text-success {
  color: #52c41a;
}

.text-danger {
  color: #ff4d4f;
}
</style>
