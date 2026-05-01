<template>
  <div class="production-daily-report">
    <!-- Search Card -->
    <el-card class="search-card">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="日期范围">
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
        <el-form-item label="车间名称">
          <el-input
            v-model="searchForm.workshop_name"
            placeholder="请输入车间名称"
            clearable
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Toolbar -->
    <div class="toolbar">
      <el-button type="primary" @click="handleGenerate">生成报表</el-button>
    </div>

    <!-- Table -->
    <el-table :data="tableData" border stripe v-loading="loading">
      <el-table-column prop="report_date" label="报表日期" width="120" />
      <el-table-column prop="workshop_name" label="车间" width="120" />
      <el-table-column prop="production_order_count" label="工单数" width="80" />
      <el-table-column prop="completed_order_count" label="完成数" width="80" />
      <el-table-column prop="total_output_qty" label="总产出" width="100" />
      <el-table-column prop="qualified_qty" label="合格数" width="100" />
      <el-table-column prop="defect_qty" label="不良数" width="80" />
      <el-table-column label="合格率%" width="100">
        <template #default="{ row }">
          <span :style="{ color: getPassRateColor(row.pass_rate) }">
            {{ row.pass_rate }}%
          </span>
        </template>
      </el-table-column>
      <el-table-column label="一次合格率%" width="110">
        <template #default="{ row }">
          <span>{{ row.first_pass_rate }}%</span>
        </template>
      </el-table-column>
      <el-table-column label="OEE%" width="100">
        <template #default="{ row }">
          <span :style="{ color: getOEEColor(row.oee) }">
            {{ row.oee }}%
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="output_per_hour" label="小时产量" width="100" />
      <el-table-column prop="worker_count" label="作业人数" width="90" />
      <el-table-column label="操作" width="100" fixed="right">
        <template #default="{ row }">
          <el-button type="primary" link @click="handleDetail(row)">详情</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- Pagination -->
    <div class="pagination">
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.page_size"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <!-- Detail Dialog -->
    <el-dialog
      v-model="detailDialogVisible"
      title="报表详情"
      width="800px"
      destroy-on-close
    >
      <el-descriptions :column="2" border>
        <el-descriptions-item label="报表日期">{{ detailData.report_date }}</el-descriptions-item>
        <el-descriptions-item label="车间">{{ detailData.workshop_name }}</el-descriptions-item>
        <el-descriptions-item label="工单数">{{ detailData.production_order_count }}</el-descriptions-item>
        <el-descriptions-item label="完成数">{{ detailData.completed_order_count }}</el-descriptions-item>
        <el-descriptions-item label="总产出">{{ detailData.total_output_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格数">{{ detailData.qualified_qty }}</el-descriptions-item>
        <el-descriptions-item label="不良数">{{ detailData.defect_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格率%">
          <span :style="{ color: getPassRateColor(detailData.pass_rate) }">
            {{ detailData.pass_rate }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="一次合格率%">{{ detailData.first_pass_rate }}%</el-descriptions-item>
        <el-descriptions-item label="OEE%">
          <span :style="{ color: getOEEColor(detailData.oee) }">
            {{ detailData.oee }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="小时产量">{{ detailData.output_per_hour }}</el-descriptions-item>
        <el-descriptions-item label="作业人数">{{ detailData.worker_count }}</el-descriptions-item>
        <el-descriptions-item label="作业时长">{{ detailData.working_hours }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>

    <!-- Generate Dialog -->
    <el-dialog
      v-model="generateDialogVisible"
      title="生成报表"
      width="500px"
      destroy-on-close
    >
      <el-form :model="generateForm" :rules="generateRules" ref="generateFormRef" label-width="120px">
        <el-form-item label="报表日期" prop="report_date">
          <el-date-picker
            v-model="generateForm.report_date"
            type="date"
            value-format="YYYY-MM-DD"
            placeholder="请选择日期"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="车间ID" prop="workshop_id">
          <el-input-number
            v-model="generateForm.workshop_id"
            :min="1"
            placeholder="请输入车间ID"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="车间名称" prop="workshop_name">
          <el-input
            v-model="generateForm.workshop_name"
            placeholder="请输入车间名称"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="generateDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmitGenerate" :loading="submitting">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import request from '@/utils/request'

// Search form
const searchForm = reactive({
  workshop_name: ''
})

// Date range
const dateRange = ref([])

// Pagination
const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

// Table data
const tableData = ref([])
const loading = ref(false)

// Detail dialog
const detailDialogVisible = ref(false)
const detailData = ref({})

// Generate dialog
const generateDialogVisible = ref(false)
const generateForm = reactive({
  report_date: '',
  workshop_id: null,
  workshop_name: ''
})
const generateRules = {
  report_date: [{ required: true, message: '请选择报表日期', trigger: 'change' }],
  workshop_id: [{ required: true, message: '请输入车间ID', trigger: 'blur' }]
}
const generateFormRef = ref(null)
const submitting = ref(false)

// Get pass rate color
const getPassRateColor = (rate) => {
  if (rate >= 98) return '#67C23A'
  if (rate >= 95) return '#E6A23C'
  return '#F56C6C'
}

// Get OEE color
const getOEEColor = (oee) => {
  if (oee >= 85) return '#67C23A'
  if (oee >= 70) return '#E6A23C'
  return '#F56C6C'
}

// Handle date change
const handleDateChange = () => {
  if (dateRange.value && dateRange.value.length === 2) {
    searchForm.start_date = dateRange.value[0]
    searchForm.end_date = dateRange.value[1]
  } else {
    searchForm.start_date = ''
    searchForm.end_date = ''
  }
}

// Load data
const loadData = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.page_size,
      workshop_name: searchForm.workshop_name || undefined,
      start_date: searchForm.start_date || undefined,
      end_date: searchForm.end_date || undefined
    }
    const res = await request.get('/report/production-daily/list', { params })
    if (res.code === 0 || res.code === 200) {
      tableData.value = res.data.list || res.data.rows || []
      pagination.total = res.data.total || 0
    } else {
      ElMessage.error(res.message || '加载数据失败')
    }
  } catch (error) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

// Handle search
const handleSearch = () => {
  pagination.page = 1
  loadData()
}

// Handle reset
const handleReset = () => {
  searchForm.workshop_name = ''
  dateRange.value = []
  searchForm.start_date = ''
  searchForm.end_date = ''
  pagination.page = 1
  loadData()
}

// Handle size change
const handleSizeChange = () => {
  pagination.page = 1
  loadData()
}

// Handle current change
const handleCurrentChange = () => {
  loadData()
}

// Handle detail
const handleDetail = async (row) => {
  try {
    const res = await request.get(`/report/production-daily/${row.id}`)
    if (res.code === 0 || res.code === 200) {
      detailData.value = res.data
      detailDialogVisible.value = true
    } else {
      ElMessage.error(res.message || '获取详情失败')
    }
  } catch (error) {
    ElMessage.error(error.message || '获取详情失败')
  }
}

// Handle generate
const handleGenerate = () => {
  generateForm.report_date = ''
  generateForm.workshop_id = null
  generateForm.workshop_name = ''
  generateDialogVisible.value = true
}

// Handle submit generate
const handleSubmitGenerate = async () => {
  try {
    await generateFormRef.value.validate()
    submitting.value = true
    const res = await request.post('/report/production-daily/generate', {
      report_date: generateForm.report_date,
      workshop_id: generateForm.workshop_id,
      workshop_name: generateForm.workshop_name
    })
    if (res.code === 0 || res.code === 200) {
      ElMessage.success('生成报表成功')
      generateDialogVisible.value = false
      loadData()
    } else {
      ElMessage.error(res.message || '生成报表失败')
    }
  } catch (error) {
    if (error !== false) {
      ElMessage.error(error.message || '生成报表失败')
    }
  } finally {
    submitting.value = false
  }
}

// Initial load
onMounted(() => {
  loadData()
})
</script>

<style scoped>
.production-daily-report {
  padding: 16px;
}

.search-card {
  margin-bottom: 16px;
}

.toolbar {
  margin-bottom: 16px;
}

.pagination {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}
</style>
