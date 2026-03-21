<template>
  <div class="sales-order-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="订单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入订单号" clearable />
        </el-form-item>
        <el-form-item label="客户名称">
          <el-input v-model="searchForm.customer_name" placeholder="请输入客户名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="全部" clearable style="width: 120px">
            <el-option label="待确认" :value="1" />
            <el-option label="已确认" :value="2" />
            <el-option label="生产中" :value="3" />
            <el-option label="已完成" :value="4" />
            <el-option label="已关闭" :value="5" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="order_no" label="订单号" width="160" />
        <el-table-column prop="customer_name" label="客户名称" min-width="150" />
        <el-table-column prop="order_date" label="订单日期" width="120" />
        <el-table-column prop="delivery_date" label="交货日期" width="120" />
        <el-table-column prop="order_type" label="订单类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ row.order_type === 'standard' ? '标准' : row.order_type === 'custom' ? '定制' : row.order_type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="priority" label="优先级" width="80">
          <template #default="{ row }">
            <el-tag :type="row.priority === 3 ? 'danger' : row.priority === 2 ? 'warning' : 'info'">
              {{ row.priority === 3 ? '加急' : row.priority === 2 ? '紧急' : '普通' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="statusType[row.status]">
              {{ statusText[row.status] || row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="120" show-overflow-tooltip />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" @click="handleConfirm(row)" v-if="row.status === 1">确认</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="客户名称" prop="customer_name">
          <el-input v-model="formData.customer_name" />
        </el-form-item>
        <el-form-item label="订单日期">
          <el-date-picker v-model="formData.order_date" type="date" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="交货日期">
          <el-date-picker v-model="formData.delivery_date" type="date" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="订单类型" prop="order_type">
          <el-select v-model="formData.order_type" style="width: 100%">
            <el-option label="标准" value="standard" />
            <el-option label="定制" value="custom" />
          </el-select>
        </el-form-item>
        <el-form-item label="优先级">
          <el-radio-group v-model="formData.priority">
            <el-radio :value="1">普通</el-radio>
            <el-radio :value="2">紧急</el-radio>
            <el-radio :value="3">加急</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getSalesOrderList, createSalesOrder, updateSalesOrder, deleteSalesOrder, confirmSalesOrder } from '@/api/production'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ order_no: '', customer_name: '', status: '' as any })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, customer_name: '', order_date: '', delivery_date: '',
  order_type: 'standard', priority: 1, status: 1, remark: ''
})

const statusText: Record<number, string> = { 1: '待确认', 2: '已确认', 3: '生产中', 4: '已完成', 5: '已关闭' }
const statusType: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'primary', 4: 'success', 5: 'info' }

const rules: FormRules = {
  customer_name: [{ required: true, message: '请输入客户名称', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑销售订单' : '新增销售订单')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getSalesOrderList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.order_no = ''; searchForm.customer_name = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, customer_name: '', order_date: '', delivery_date: '', order_type: 'standard', priority: 1, status: 1, remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleConfirm = async (row: any) => {
  await ElMessageBox.confirm('确认该订单？', '提示', { type: 'warning' })
  await confirmSalesOrder(row.id)
  ElMessage.success('确认成功')
  loadData()
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该订单吗？', '提示', { type: 'warning' })
  await deleteSalesOrder(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateSalesOrder(formData.id, formData) : await createSalesOrder(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.sales-order-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
