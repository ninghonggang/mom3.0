<template>
  <div class="production-cost">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="成本类型">
          <el-select v-model="searchForm.cost_type" placeholder="请选择" clearable>
            <el-option label="材料成本" value="material" />
            <el-option label="人工成本" value="labor" />
            <el-option label="制造费用" value="overhead" />
          </el-select>
        </el-form-item>
        <el-form-item label="日期范围">
          <el-date-picker v-model="dateRange" type="daterange" value-format="YYYY-MM-DD" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon>新增成本
      </el-button>
      <el-button type="success" @click="handleViewSummary" :disabled="!selectedOrderId">
        <el-icon><Histogram /></el-icon>工单汇总
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" @row-click="handleRowClick">
        <el-table-column prop="order_no" label="工单号" width="140" />
        <el-table-column prop="cost_type" label="成本类型" width="100">
          <template #default="{ row }">
            <el-tag :type="getCostTypeColor(row.cost_type)">{{ getCostTypeText(row.cost_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="cost_item" label="成本项目" width="150" />
        <el-table-column prop="quantity" label="数量" width="100" />
        <el-table-column prop="unit_price" label="单价" width="100" />
        <el-table-column prop="amount" label="金额" width="120">
          <template #default="{ row }">
            <span class="amount-text">¥{{ row.amount.toFixed(2) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="process_name" label="工序" width="120" />
        <el-table-column prop="department_name" label="部门" width="120" />
        <el-table-column prop="worker_name" label="操作人" width="100" />
        <el-table-column prop="cost_date" label="成本日期" width="120" />
        <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
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
    <el-dialog v-model="dialogVisible" title="新增成本记录" width="600px">
      <el-form :model="formData" label-width="100px" :rules="formRules" ref="formRef">
        <el-form-item label="工单号" prop="order_id">
          <el-select v-model="formData.order_id" placeholder="请选择工单" filterable style="width: 100%" @change="handleOrderChange">
            <el-option v-for="order in orderList" :key="order.id" :label="order.order_no" :value="order.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="成本类型" prop="cost_type">
          <el-select v-model="formData.cost_type" placeholder="请选择" style="width: 100%">
            <el-option label="材料成本" value="material" />
            <el-option label="人工成本" value="labor" />
            <el-option label="制造费用" value="overhead" />
          </el-select>
        </el-form-item>
        <el-form-item label="成本项目" prop="cost_item">
          <el-input v-model="formData.cost_item" placeholder="请输入成本项目" />
        </el-form-item>
        <el-form-item label="数量">
          <el-input-number v-model="formData.quantity" :min="0" :precision="2" style="width: 100%" />
        </el-form-item>
        <el-form-item label="单价">
          <el-input-number v-model="formData.unit_price" :min="0" :precision="4" style="width: 100%" />
        </el-form-item>
        <el-form-item label="金额" prop="amount">
          <el-input-number v-model="formData.amount" :min="0" :precision="2" style="width: 100%" />
        </el-form-item>
        <el-form-item label="工序">
          <el-input v-model="formData.process_name" placeholder="请输入工序" />
        </el-form-item>
        <el-form-item label="部门">
          <el-input v-model="formData.department_name" placeholder="请输入部门" />
        </el-form-item>
        <el-form-item label="操作人">
          <el-input v-model="formData.worker_name" placeholder="请输入操作人" />
        </el-form-item>
        <el-form-item label="成本日期">
          <el-date-picker v-model="formData.cost_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>

    <!-- 工单成本汇总对话框 -->
    <el-dialog v-model="summaryDialogVisible" title="工单成本汇总" width="500px">
      <el-descriptions v-if="summaryData" :column="2" border>
        <el-descriptions-item label="工单号">{{ summaryData.order_no }}</el-descriptions-item>
        <el-descriptions-item label="已完成数量">{{ summaryData.completed_qty }}</el-descriptions-item>
        <el-descriptions-item label="材料成本">
          <span class="amount-text">¥{{ summaryData.material_cost.toFixed(2) }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="人工成本">
          <span class="amount-text">¥{{ summaryData.labor_cost.toFixed(2) }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="制造费用">
          <span class="amount-text">¥{{ summaryData.overhead_cost.toFixed(2) }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="单位成本">
          <span class="amount-text">¥{{ summaryData.unit_cost.toFixed(4) }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="总成本" :span="2">
          <span class="total-cost">¥{{ summaryData.total_cost.toFixed(2) }}</span>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getProductionCostList, createProductionCost, getProductionCostSummary, deleteProductionCost } from '@/api/production_cost'
import { getProductionOrderList } from '@/api/production'

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const summaryDialogVisible = ref(false)
const tableData = ref<any[]>([])
const orderList = ref<any[]>([])
const selectedOrderId = ref<number | null>(null)
const summaryData = ref<any>(null)
const dateRange = ref<[string, string] | null>(null)

const searchForm = reactive({
  order_no: '',
  cost_type: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  order_id: null,
  order_no: '',
  cost_type: '',
  cost_item: '',
  quantity: 0,
  unit_price: 0,
  amount: 0,
  process_name: '',
  department_name: '',
  worker_name: '',
  cost_date: '',
  remark: ''
})

const formRules = {
  order_id: [{ required: true, message: '请选择工单', trigger: 'change' }],
  cost_type: [{ required: true, message: '请选择成本类型', trigger: 'change' }],
  cost_item: [{ required: true, message: '请输入成本项目', trigger: 'blur' }],
  amount: [{ required: true, message: '请输入金额', trigger: 'blur' }]
}

const getCostTypeText = (type: string) => {
  const map: Record<string, string> = { material: '材料', labor: '人工', overhead: '制造' }
  return map[type] || type
}

const getCostTypeColor = (type: string) => {
  const map: Record<string, string> = { material: 'warning', labor: 'success', overhead: 'primary' }
  return map[type] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize,
      order_no: searchForm.order_no,
      cost_type: searchForm.cost_type
    }
    if (dateRange.value) {
      params.start_date = dateRange.value[0]
      params.end_date = dateRange.value[1]
    }
    const res = await getProductionCostList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadOrders = async () => {
  try {
    const res = await getProductionOrderList({ page: 1, page_size: 100 })
    orderList.value = res.data.list || []
  } catch (e) {
    console.error(e)
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.order_no = ''; searchForm.cost_type = ''; dateRange.value = null; handleSearch() }

const handleRowClick = (row: any) => {
  selectedOrderId.value = row.order_id
}

const handleCreate = () => {
  formData.order_id = null
  formData.order_no = ''
  formData.cost_type = ''
  formData.cost_item = ''
  formData.quantity = 0
  formData.unit_price = 0
  formData.amount = 0
  formData.process_name = ''
  formData.department_name = ''
  formData.worker_name = ''
  formData.cost_date = ''
  formData.remark = ''
  dialogVisible.value = true
}

const handleOrderChange = (orderId: number) => {
  const order = orderList.value.find(o => o.id === orderId)
  if (order) {
    formData.order_no = order.order_no
  }
}

const handleSave = async () => {
  if (!formData.order_id || !formData.cost_type || !formData.cost_item || !formData.amount) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    await createProductionCost(formData)
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleViewSummary = async () => {
  if (!selectedOrderId.value) {
    ElMessage.warning('请先选择一条记录')
    return
  }
  try {
    const res = await getProductionCostSummary(selectedOrderId.value)
    summaryData.value = res.data
    summaryDialogVisible.value = true
  } catch (e) {
    console.error(e)
    ElMessage.error('加载汇总失败')
  }
}

onMounted(() => { loadData(); loadOrders() })
</script>

<style scoped lang="scss">
.production-cost {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .amount-text { color: #E6A23C; font-weight: 500; }
  .total-cost { color: #F56C6C; font-weight: bold; font-size: 18px; }
}
</style>