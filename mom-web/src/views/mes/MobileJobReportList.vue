<template>
  <div class="mobile-job-report-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_code" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="报工类型">
          <el-select v-model="searchForm.report_type" placeholder="请选择" clearable style="width: 100px">
            <el-option label="正常" :value="1" />
            <el-option label="补报" :value="2" />
            <el-option label="异常" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable style="width: 100px">
            <el-option label="已提交" :value="1" />
            <el-option label="已确认" :value="2" />
            <el-option label="已审核" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleReport">
        <el-icon><Edit /></el-icon>报工
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="order_code" label="工单号" width="140" />
        <el-table-column prop="workstation_name" label="工位" width="100" />
        <el-table-column prop="employee_name" label="员工" width="80" />
        <el-table-column prop="reported_quantity" label="报工数量" width="90" />
        <el-table-column prop="qualified_quantity" label="合格数量" width="90" />
        <el-table-column prop="defective_quantity" label="不良数量" width="90" />
        <el-table-column prop="report_type" label="类型" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.report_type === 1" type="success">正常</el-tag>
            <el-tag v-else-if="row.report_type === 2" type="warning">补报</el-tag>
            <el-tag v-else type="danger">异常</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.status === 1" type="info">已提交</el-tag>
            <el-tag v-else-if="row.status === 2" type="warning">已确认</el-tag>
            <el-tag v-else type="success">已审核</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="报工时间" width="160">
          <template #default="{ row }">
            {{ row.created_at ? row.created_at.slice(0, 19) : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleView(row)">查看</el-button>
            <el-button link type="success" size="small" @click="handleConfirm(row)" v-if="row.status === 1">确认</el-button>
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

    <!-- 报工对话框 -->
    <el-dialog v-model="reportVisible" title="移动端报工" width="600px">
      <el-form :model="reportForm" label-width="100px">
        <el-form-item label="工单" required>
          <el-select v-model="reportForm.order_id" placeholder="请选择工单" style="width: 100%" @change="handleOrderChange">
            <el-option v-for="o in pendingOrders" :key="o.order_id" :label="o.order_code" :value="o.order_id">
              <span>{{ o.order_code }}</span>
              <span style="float:right;color:#999;font-size:12px">{{ o.product_name }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="报工数量" required>
          <el-input-number v-model="reportForm.reported_quantity" :min="1" style="width: 100%" />
        </el-form-item>
        <el-form-item label="合格数量">
          <el-input-number v-model="reportForm.qualified_quantity" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="工作时长(分)">
          <el-input-number v-model="reportForm.work_minutes" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="报工类型">
          <el-radio-group v-model="reportForm.report_type">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="2">补报</el-radio>
            <el-radio :label="3">异常</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="reportForm.remarks" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="reportVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleReportSubmit">提交报工</el-button>
      </template>
    </el-dialog>

    <!-- 查看对话框 -->
    <el-dialog v-model="viewVisible" title="报工详情" width="600px">
      <el-descriptions :column="2" border v-if="currentRecord">
        <el-descriptions-item label="工单号">{{ currentRecord.order_code }}</el-descriptions-item>
        <el-descriptions-item label="工位">{{ currentRecord.workstation_name }}</el-descriptions-item>
        <el-descriptions-item label="工序">{{ currentRecord.process_name }}</el-descriptions-item>
        <el-descriptions-item label="员工">{{ currentRecord.employee_name }}</el-descriptions-item>
        <el-descriptions-item label="报工数量">{{ currentRecord.reported_quantity }}</el-descriptions-item>
        <el-descriptions-item label="合格数量">{{ currentRecord.qualified_quantity }}</el-descriptions-item>
        <el-descriptions-item label="不良数量">{{ currentRecord.defective_quantity }}</el-descriptions-item>
        <el-descriptions-item label="工作时长">{{ currentRecord.work_minutes }}分钟</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag v-if="currentRecord.status === 1" type="info">已提交</el-tag>
          <el-tag v-else-if="currentRecord.status === 2" type="warning">已确认</el-tag>
          <el-tag v-else type="success">已审核</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="报工时间" :span="2">{{ currentRecord.created_at }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ currentRecord.remarks || '-' }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getMobileJobReportList, getMobileJobReport, createMobileJobReport, confirmMobileJobReport, getMobilePendingOrders } from '@/api/mobile_job_report'

const loading = ref(false)
const saveLoading = ref(false)
const reportVisible = ref(false)
const viewVisible = ref(false)
const tableData = ref<any[]>([])
const pendingOrders = ref<any[]>([])
const currentRecord = ref<any>(null)

const searchForm = reactive({
  order_code: '',
  report_type: null as number | null,
  status: null as number | null
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const reportForm = reactive({
  order_id: null as number | null,
  reported_quantity: 1,
  qualified_quantity: 1,
  work_minutes: 0,
  report_type: 1,
  remarks: ''
})

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMobileJobReportList({
      page: pagination.page,
      page_size: pagination.pageSize,
      order_code: searchForm.order_code,
      report_type: searchForm.report_type,
      status: searchForm.status
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadPendingOrders = async () => {
  try {
    const res = await getMobilePendingOrders({})
    pendingOrders.value = res.data.list || []
  } catch (e) {
    console.error(e)
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.order_code = ''
  searchForm.report_type = null
  searchForm.status = null
  handleSearch()
}

const handleOrderChange = (orderId: number) => {
  const order = pendingOrders.value.find(o => o.order_id === orderId)
  if (order) {
    reportForm.reported_quantity = Math.min(order.remaining_quantity, 1)
  }
}

const handleReport = () => {
  reportForm.order_id = null
  reportForm.reported_quantity = 1
  reportForm.qualified_quantity = 1
  reportForm.work_minutes = 0
  reportForm.report_type = 1
  reportForm.remarks = ''
  reportVisible.value = true
}

const handleReportSubmit = async () => {
  if (!reportForm.order_id) {
    ElMessage.warning('请选择工单')
    return
  }
  if (reportForm.reported_quantity <= 0) {
    ElMessage.warning('报工数量必须大于0')
    return
  }
  saveLoading.value = true
  try {
    await createMobileJobReport(reportForm)
    ElMessage.success('报工成功')
    reportVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleView = async (row: any) => {
  try {
    const res = await getMobileJobReport(row.id)
    currentRecord.value = res.data
    viewVisible.value = true
  } catch (e) {
    console.error(e)
  }
}

const handleConfirm = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定确认该报工吗？', '提示', { type: 'warning' })
    await confirmMobileJobReport(row.id)
    ElMessage.success('确认成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

onMounted(() => { loadData(); loadPendingOrders() })
</script>

<style scoped lang="scss">
.mobile-job-report-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>