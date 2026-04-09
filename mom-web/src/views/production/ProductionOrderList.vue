<template>
  <div class="production-order-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待生产" :value="1" />
            <el-option label="生产中" :value="2" />
            <el-option label="已完成" :value="3" />
            <el-option label="已取消" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:order:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="order_no" label="工单号" width="140" />
        <el-table-column prop="material_code" label="物料编码" width="100" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="quantity" label="计划数量" width="100" />
        <el-table-column prop="completed_qty" label="已完成" width="80" />
        <el-table-column prop="workshop_name" label="车间" width="100" />
        <el-table-column prop="plan_start_date" label="计划开始" width="120" />
        <el-table-column prop="plan_end_date" label="计划结束" width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('production:order:release') && row.status === 1" @click="handleStart(row)">开始</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('production:order:release') && row.status === 2" @click="handleComplete(row)">完成</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('production:order:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('production:order:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="工单号" prop="order_no">
              <el-input v-model="formData.order_no" :disabled="!!formData.id" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="物料编码" prop="material_code">
              <el-input v-model="formData.material_code" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="物料名称">
              <el-input v-model="formData.material_name" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="计划数量" prop="quantity">
              <el-input-number v-model="formData.quantity" :min="1" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="车间">
              <el-input v-model="formData.workshop_name" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="优先级">
              <el-select v-model="formData.priority">
                <el-option label="普通" :value="1" />
                <el-option label="紧急" :value="2" />
                <el-option label="加急" :value="3" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="计划开始">
              <el-date-picker v-model="formData.plan_start_date" type="date" value-format="YYYY-MM-DD" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="计划结束">
              <el-date-picker v-model="formData.plan_end_date" type="date" value-format="YYYY-MM-DD" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" />
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
import { getProductionOrderList, createProductionOrder, updateProductionOrder, deleteProductionOrder, startProductionOrder, completeProductionOrder } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ order_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, order_no: '', material_code: '', material_name: '', quantity: 1, workshop_name: '',
  priority: 1, plan_start_date: '', plan_end_date: '', remark: ''
})

const rules: FormRules = {
  order_no: [{ required: true, message: '请输入工单号', trigger: 'blur' }],
  material_code: [{ required: true, message: '请输入物料编码', trigger: 'blur' }],
  quantity: [{ required: true, message: '请输入计划数量', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑工单' : '新增工单')

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待生产', 2: '生产中', 3: '已完成', 4: '已取消' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success', 4: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getProductionOrderList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.order_no = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, order_no: '', material_code: '', material_name: '', quantity: 1, workshop_name: '', priority: 1, plan_start_date: '', plan_end_date: '', remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该工单吗？', '提示', { type: 'warning' })
    await deleteProductionOrder(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleStart = async (row: any) => {
  await startProductionOrder(row.id)
  ElMessage.success('已开始生产')
  loadData()
}

const handleComplete = async (row: any) => {
  await completeProductionOrder(row.id)
  ElMessage.success('已完成生产')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateProductionOrder(formData.id, formData) : await createProductionOrder(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.production-order-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
