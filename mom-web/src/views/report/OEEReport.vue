<template>
  <div class="oee-report-container">
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
        <el-form-item label="车间">
          <el-input
            v-model="searchForm.workshop_id"
            placeholder="车间ID"
            clearable
            style="width: 150px"
          />
        </el-form-item>
        <el-form-item label="产线">
          <el-input
            v-model="searchForm.line_id"
            placeholder="产线ID"
            clearable
            style="width: 150px"
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
    <el-table :data="tableData" border stripe style="width: 100%">
      <el-table-column prop="report_date" label="报表日期" width="120" />
      <el-table-column prop="workshop_name" label="车间" width="120" />
      <el-table-column prop="line_name" label="产线" width="120" />
      <el-table-column prop="availability" label="可用率%" width="100">
        <template #default="{ row }">
          {{ row.availability?.toFixed(2) }}%
        </template>
      </el-table-column>
      <el-table-column prop="performance" label="性能率%" width="100">
        <template #default="{ row }">
          {{ row.performance?.toFixed(2) }}%
        </template>
      </el-table-column>
      <el-table-column prop="quality" label="质量率%" width="100">
        <template #default="{ row }">
          {{ row.quality?.toFixed(2) }}%
        </template>
      </el-table-column>
      <el-table-column prop="oee" label="OEE%" width="100">
        <template #default="{ row }">
          <span
            :style="{
              color: getOEEColor(row.oee),
              fontWeight: 'bold'
            }"
          >
            {{ row.oee?.toFixed(2) }}%
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="planned_production_time" label="计划时间" width="100">
        <template #default="{ row }">
          {{ row.planned_production_time }} min
        </template>
      </el-table-column>
      <el-table-column prop="actual_production_time" label="实际时间" width="100">
        <template #default="{ row }">
          {{ row.actual_production_time }} min
        </template>
      </el-table-column>
      <el-table-column prop="down_time" label="停机时间" width="100">
        <template #default="{ row }">
          {{ row.down_time }} min
        </template>
      </el-table-column>
      <el-table-column label="操作" fixed="right" width="120">
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
      title="OEE报表详情"
      width="800px"
      destroy-on-close
    >
      <el-descriptions :column="2" border>
        <el-descriptions-item label="报表日期">{{ detailData.report_date }}</el-descriptions-item>
        <el-descriptions-item label="车间">{{ detailData.workshop_name }}</el-descriptions-item>
        <el-descriptions-item label="车间ID">{{ detailData.workshop_id }}</el-descriptions-item>
        <el-descriptions-item label="产线">{{ detailData.line_name }}</el-descriptions-item>
        <el-descriptions-item label="产线ID">{{ detailData.line_id }}</el-descriptions-item>
        <el-descriptions-item label="可用率">
          <span :style="{ color: getOEEColor(detailData.availability) }">
            {{ detailData.availability?.toFixed(2) }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="性能率">
          <span :style="{ color: getOEEColor(detailData.performance) }">
            {{ detailData.performance?.toFixed(2) }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="质量率">
          <span :style="{ color: getOEEColor(detailData.quality) }">
            {{ detailData.quality?.toFixed(2) }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="OEE">
          <span
            :style="{
              color: getOEEColor(detailData.oee),
              fontWeight: 'bold',
              fontSize: '16px'
            }"
          >
            {{ detailData.oee?.toFixed(2) }}%
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="计划时间">{{ detailData.planned_production_time }} min</el-descriptions-item>
        <el-descriptions-item label="实际时间">{{ detailData.actual_production_time }} min</el-descriptions-item>
        <el-descriptions-item label="停机时间">{{ detailData.down_time }} min</el-descriptions-item>
        <el-descriptions-item label="速度损失">{{ detailData.speed_loss }}</el-descriptions-item>
        <el-descriptions-item label="缺陷损失">{{ detailData.defect_loss }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>

    <!-- Generate Dialog -->
    <el-dialog
      v-model="generateDialogVisible"
      title="生成OEE报表"
      width="500px"
      destroy-on-close
    >
      <el-form
        ref="generateFormRef"
        :model="generateForm"
        :rules="generateRules"
        label-width="120px"
      >
        <el-form-item label="报表日期" prop="report_date">
          <el-date-picker
            v-model="generateForm.report_date"
            type="date"
            value-format="YYYY-MM-DD"
            placeholder="选择日期"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="车间ID" prop="workshop_id">
          <el-input
            v-model.number="generateForm.workshop_id"
            type="number"
            placeholder="请输入车间ID"
          />
        </el-form-item>
        <el-form-item label="车间名称" prop="workshop_name">
          <el-input
            v-model="generateForm.workshop_name"
            placeholder="请输入车间名称"
          />
        </el-form-item>
        <el-form-item label="产线ID" prop="line_id">
          <el-input
            v-model.number="generateForm.line_id"
            type="number"
            placeholder="请输入产线ID"
          />
        </el-form-item>
        <el-form-item label="产线名称" prop="line_name">
          <el-input
            v-model="generateForm.line_name"
            placeholder="请输入产线名称"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="generateDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="generateLoading" @click="handleGenerateSubmit">确定</el-button>
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
  workshop_id: '',
  line_id: ''
})
const dateRange = ref([])
const start_date = ref('')
const end_date = ref('')

// Table data
const tableData = ref([])

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
const generateLoading = ref(false)
const generateFormRef = ref(null)
const generateForm = reactive({
  report_date: '',
  workshop_id: null,
  workshop_name: '',
  line_id: null,
  line_name: ''
})
const generateRules = {
  report_date: [{ required: true, message: '请选择报表日期', trigger: 'change' }],
  workshop_id: [{ required: true, message: '请输入车间ID', trigger: 'blur' }],
  line_id: [{ required: true, message: '请输入产线ID', trigger: 'blur' }]
}

// Methods
const handleDateChange = (val) => {
  if (val && val.length === 2) {
    start_date.value = val[0]
    end_date.value = val[1]
  } else {
    start_date.value = ''
    end_date.value = ''
  }
}

const getOEEColor = (value) => {
  if (value === null || value === undefined) return '#333'
  if (value >= 85) return '#67C23A'
  if (value >= 70) return '#E6A23C'
  return '#F56C6C'
}

const loadData = async () => {
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.page_size,
      start_date: start_date.value,
      end_date: end_date.value,
      workshop_id: searchForm.workshop_id || undefined,
      line_id: searchForm.line_id || undefined
    }
    const res = await request.get('/report/oee/list', { params })
    if (res.code === 0 || res.code === 200) {
      tableData.value = res.data?.list || []
      pagination.total = res.data?.total || 0
    } else {
      ElMessage.error(res.message || '加载数据失败')
    }
  } catch (error) {
    ElMessage.error('加载数据失败')
    console.error(error)
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.workshop_id = ''
  searchForm.line_id = ''
  dateRange.value = []
  start_date.value = ''
  end_date.value = ''
  pagination.page = 1
  loadData()
}

const handleSizeChange = () => {
  pagination.page = 1
  loadData()
}

const handleCurrentChange = () => {
  loadData()
}

const handleDetail = async (row) => {
  try {
    const res = await request.get(`/report/oee/${row.id}`)
    if (res.code === 0 || res.code === 200) {
      detailData.value = res.data || {}
      detailDialogVisible.value = true
    } else {
      ElMessage.error(res.message || '获取详情失败')
    }
  } catch (error) {
    ElMessage.error('获取详情失败')
    console.error(error)
  }
}

const handleGenerate = () => {
  generateForm.report_date = ''
  generateForm.workshop_id = null
  generateForm.workshop_name = ''
  generateForm.line_id = null
  generateForm.line_name = ''
  generateDialogVisible.value = true
}

const handleGenerateSubmit = async () => {
  if (!generateFormRef.value) return
  try {
    const valid = await generateFormRef.value.validate()
    if (!valid) return
    
    generateLoading.value = true
    const res = await request.post('/report/oee/generate', {
      report_date: generateForm.report_date,
      workshop_id: generateForm.workshop_id,
      workshop_name: generateForm.workshop_name,
      line_id: generateForm.line_id,
      line_name: generateForm.line_name
    })
    
    if (res.code === 0 || res.code === 200) {
      ElMessage.success('生成报表成功')
      generateDialogVisible.value = false
      loadData()
    } else {
      ElMessage.error(res.message || '生成报表失败')
    }
  } catch (error) {
    ElMessage.error('生成报表失败')
    console.error(error)
  } finally {
    generateLoading.value = false
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.oee-report-container {
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
