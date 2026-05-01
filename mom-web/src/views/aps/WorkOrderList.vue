<template>
  <div class="work-order-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单编号">
          <el-input v-model="searchForm.work_order_no" placeholder="请输入工单编号" clearable />
        </el-form-item>
        <el-form-item label="产品名称">
          <el-input v-model="searchForm.product_name" placeholder="请输入产品名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待生产" :value="1" />
            <el-option label="生产中" :value="2" />
            <el-option label="已完成" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="优先级">
          <el-select v-model="searchForm.priority" placeholder="请选择" clearable>
            <el-option label="低" :value="1" />
            <el-option label="中" :value="2" />
            <el-option label="高" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:work-order:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('aps:work-order:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="work_order_no" label="工单编号" width="140" />
        <el-table-column prop="product_code" label="产品编码" width="120" />
        <el-table-column prop="product_name" label="产品名称" min-width="160" />
        <el-table-column prop="plan_qty" label="计划数量" width="100" />
        <el-table-column prop="produced_qty" label="已生产数量" width="110" />
        <el-table-column prop="work_center_name" label="工作中心" width="140" />
        <el-table-column prop="priority" label="优先级" width="80">
          <template #default="{ row }">
            <el-tag :type="getPriorityType(row.priority)">{{ getPriorityText(row.priority) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:work-order:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:work-order:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px" @close="handleDialogClose">
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-form-item label="工单编号" prop="work_order_no">
          <el-input v-model="formData.work_order_no" placeholder="请输入工单编号" />
        </el-form-item>
        <el-form-item label="产品编码" prop="product_code">
          <el-input v-model="formData.product_code" placeholder="请输入产品编码" />
        </el-form-item>
        <el-form-item label="产品名称" prop="product_name">
          <el-input v-model="formData.product_name" placeholder="请输入产品名称" />
        </el-form-item>
        <el-form-item label="计划数量" prop="plan_qty">
          <el-input-number v-model="formData.plan_qty" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="工作中心" prop="work_center_name">
          <el-input v-model="formData.work_center_name" placeholder="请输入工作中心" />
        </el-form-item>
        <el-form-item label="优先级" prop="priority">
          <el-select v-model="formData.priority" placeholder="请选择优先级" style="width: 100%">
            <el-option label="低" :value="1" />
            <el-option label="中" :value="2" />
            <el-option label="高" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width: 100%">
            <el-option label="待生产" :value="1" />
            <el-option label="生产中" :value="2" />
            <el-option label="已完成" :value="3" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import request from '@/utils/request'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const searchForm = reactive({ work_order_no: '', product_name: '', status: '', priority: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive({
  id: '',
  work_order_no: '',
  product_code: '',
  product_name: '',
  plan_qty: 0,
  produced_qty: 0,
  work_center_name: '',
  priority: 2,
  status: 1
})

const formRules: FormRules = {
  work_order_no: [{ required: true, message: '请输入工单编号', trigger: 'blur' }],
  product_code: [{ required: true, message: '请输入产品编码', trigger: 'blur' }],
  product_name: [{ required: true, message: '请输入产品名称', trigger: 'blur' }],
  plan_qty: [{ required: true, message: '请输入计划数量', trigger: 'blur' }],
  work_center_name: [{ required: true, message: '请输入工作中心', trigger: 'blur' }]
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待生产', 2: '生产中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }
  return map[status] || 'info'
}

const getPriorityText = (priority: number) => {
  const map: Record<number, string> = { 1: '低', 2: '中', 3: '高' }
  return map[priority] || '中'
}

const getPriorityType = (priority: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'danger' }
  return map[priority] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    // TODO: 后端接口开发完成后替换为实际 API /aps/work-order/list
    const res = await request.get('/aps/workorder/list', { params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize } })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.work_order_no = ''; searchForm.product_name = ''; searchForm.status = ''; searchForm.priority = ''; handleSearch() }
const handleAdd = () => {
  dialogTitle.value = '新增工单'
  isEdit.value = false
  Object.assign(formData, { id: '', work_order_no: '', product_code: '', product_name: '', plan_qty: 0, produced_qty: 0, work_center_name: '', priority: 2, status: 1 })
  dialogVisible.value = true
}
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑工单'
  isEdit.value = true
  Object.assign(formData, row)
  dialogVisible.value = true
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该工单吗？', '提示', { type: 'warning' })
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}
const handleExport = () => { ElMessage.info('导出功能') }
const handleDialogClose = () => { formRef.value?.resetFields() }
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      ElMessage.success('操作成功')
      dialogVisible.value = false
      loadData()
    }
  })
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.work-order-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
