<template>
  <div class="delivery-report">
    <!-- Search Card -->
    <el-card class="search-card">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="报表月份">
          <el-date-picker
            v-model="searchForm.monthRange"
            type="monthrange"
            range-separator="至"
            start-placeholder="开始月份"
            end-placeholder="结束月份"
            value-format="YYYY-MM"
            placeholder="选择月份范围"
            @change="handleMonthRangeChange"
          />
        </el-form-item>
        <el-form-item label="客户名称">
          <el-input
            v-model="searchForm.customer_name"
            placeholder="请输入客户名称"
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
    <el-table :data="tableData" border style="width: 100%" v-loading="loading">
      <el-table-column prop="report_month" label="报表月份" width="120" />
      <el-table-column prop="customer_name" label="客户名称" min-width="150" />
      <el-table-column prop="order_count" label="订单数" width="100" />
      <el-table-column prop="total_order_qty" label="订单总量" width="120">
        <template #default="{ row }">
          {{ row.total_order_qty?.toFixed(2) }}
        </template>
      </el-table-column>
      <el-table-column prop="delivered_qty" label="已交付量" width="120">
        <template #default="{ row }">
          {{ row.delivered_qty?.toFixed(2) }}
        </template>
      </el-table-column>
      <el-table-column prop="on_time_deliver_qty" label="准时交付量" width="120">
        <template #default="{ row }">
          {{ row.on_time_deliver_qty?.toFixed(2) }}
        </template>
      </el-table-column>
      <el-table-column prop="delivery_rate" label="交付率%" width="120">
        <template #default="{ row }">
          <span :style="{ color: row.delivery_rate >= 95 ? '#67c23a' : '#e6a23c' }">
            {{ row.delivery_rate?.toFixed(2) }}%
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="on_time_rate" label="准时率%" width="120">
        <template #default="{ row }">
          <span :style="{ color: row.on_time_rate >= 95 ? '#67c23a' : '#e6a23c' }">
            {{ row.on_time_rate?.toFixed(2) }}%
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="late_deliver_qty" label="延期交付量" width="120">
        <template #default="{ row }">
          {{ row.late_deliver_qty?.toFixed(2) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="120" fixed="right">
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
      width="600px"
      destroy-on-close
    >
      <el-descriptions :column="2" border>
        <el-descriptions-item label="报表月份">{{ detailData.report_month }}</el-descriptions-item>
        <el-descriptions-item label="客户名称">{{ detailData.customer_name }}</el-descriptions-item>
        <el-descriptions-item label="客户ID">{{ detailData.customer_id }}</el-descriptions-item>
        <el-descriptions-item label="订单数">{{ detailData.order_count }}</el-descriptions-item>
        <el-descriptions-item label="订单总量">{{ detailData.total_order_qty?.toFixed(2) }}</el-descriptions-item>
        <el-descriptions-item label="已交付量">{{ detailData.delivered_qty?.toFixed(2) }}</el-descriptions-item>
        <el-descriptions-item label="准时交付量">{{ detailData.on_time_deliver_qty?.toFixed(2) }}</el-descriptions-item>
        <el-descriptions-item label="交付率%">
          <span :style="{ color: detailData.delivery_rate >= 95 ? '#67c23a' : '#e6a23c' }">
            {{ detailData.delivery_rate?.toFixed(2) }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="准时率%">
          <span :style="{ color: detailData.on_time_rate >= 95 ? '#67c23a' : '#e6a23c' }">
            {{ detailData.on_time_rate?.toFixed(2) }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="延期交付量">{{ detailData.late_deliver_qty?.toFixed(2) }}</el-descriptions-item>
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
      <el-form :model="generateForm" :rules="generateRules" ref="generateFormRef" label-width="100px">
        <el-form-item label="报表月份" prop="report_month">
          <el-date-picker
            v-model="generateForm.report_month"
            type="month"
            value-format="YYYY-MM"
            placeholder="选择月份"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="客户ID" prop="customer_id">
          <el-input-number
            v-model="generateForm.customer_id"
            :min="1"
            placeholder="请输入客户ID"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="客户名称" prop="customer_name">
          <el-input
            v-model="generateForm.customer_name"
            placeholder="请输入客户名称"
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
  monthRange: [],
  customer_name: ''
})

// Table data
const tableData = ref([])
const loading = ref(false)

// Pagination
const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

// Detail dialog
const detailDialogVisible = ref(false)
const detailData = ref({})

// Generate dialog
const generateDialogVisible = ref(false)
const submitting = ref(false)
const generateForm = reactive({
  report_month: '',
  customer_id: null,
  customer_name: ''
})

const generateRules = {
  report_month: [{ required: true, message: '请选择报表月份', trigger: 'change' }],
  customer_id: [{ required: true, message: '请输入客户ID', trigger: 'blur' }]
}

const generateFormRef = ref(null)

// Handle month range change
const handleMonthRangeChange = () => {
  // Reset to page 1 when search criteria change
  pagination.page = 1
}

// Handle search
const handleSearch = () => {
  pagination.page = 1
  loadData()
}

// Handle reset
const handleReset = () => {
  searchForm.monthRange = []
  searchForm.customer_name = ''
  pagination.page = 1
  pagination.page_size = 10
  loadData()
}

// Load data
const loadData = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.page_size,
      customer_name: searchForm.customer_name || undefined
    }

    if (searchForm.monthRange && searchForm.monthRange.length === 2) {
      params.start_month = searchForm.monthRange[0]
      params.end_month = searchForm.monthRange[1]
    }

    const response = await request.get('/report/delivery/list', { params })
    const result = response.data || response

    tableData.value = result.list || result.data || []
    pagination.total = result.total || 0
  } catch (error) {
    ElMessage.error('加载数据失败: ' + (error.message || '未知错误'))
  } finally {
    loading.value = false
  }
}

// Handle detail
const handleDetail = async (row) => {
  try {
    const response = await request.get(`/report/delivery/${row.id}`)
    const result = response.data || response
    detailData.value = result
    detailDialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取详情失败: ' + (error.message || '未知错误'))
  }
}

// Handle generate
const handleGenerate = () => {
  generateForm.report_month = ''
  generateForm.customer_id = null
  generateForm.customer_name = ''
  generateDialogVisible.value = true
}

// Handle submit generate
const handleSubmitGenerate = async () => {
  if (!generateFormRef.value) return

  try {
    await generateFormRef.value.validate()
  } catch (error) {
    return
  }

  submitting.value = true
  try {
    await request.post('/report/delivery/generate', {
      report_month: generateForm.report_month,
      customer_id: generateForm.customer_id,
      customer_name: generateForm.customer_name
    })
    ElMessage.success('生成报表成功')
    generateDialogVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('生成报表失败: ' + (error.message || '未知错误'))
  } finally {
    submitting.value = false
  }
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

// Initial load
onMounted(() => {
  loadData()
})
</script>

<style scoped>
.delivery-report {
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
