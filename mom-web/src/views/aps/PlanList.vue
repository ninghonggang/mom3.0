<template>
  <div class="plan-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="计划月份">
          <el-date-picker v-model="searchForm.plan_month" type="month" placeholder="请选择月份" clearable value-format="YYYY-MM" />
        </el-form-item>
        <el-form-item label="物料名称">
          <el-input v-model="searchForm.material_name" placeholder="请输入物料名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" :value="0" />
            <el-option label="已下达" :value="1" />
            <el-option label="已排程" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:plan:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('aps:plan:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="plan_month" label="计划月份" width="120" />
        <el-table-column prop="material_code" label="物料编码" width="140" />
        <el-table-column prop="material_name" label="物料名称" min-width="160" />
        <el-table-column prop="quantity" label="计划数量" width="100" />
        <el-table-column prop="scheduled_qty" label="已排数量" width="100" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:plan:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:plan:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="计划月份" prop="plan_month">
          <el-date-picker v-model="formData.plan_month" type="month" placeholder="请选择月份" value-format="YYYY-MM" style="width: 100%" />
        </el-form-item>
        <el-form-item label="物料编码" prop="material_code">
          <el-input v-model="formData.material_code" placeholder="请输入物料编码" />
        </el-form-item>
        <el-form-item label="物料名称" prop="material_name">
          <el-input v-model="formData.material_name" placeholder="请输入物料名称" />
        </el-form-item>
        <el-form-item label="计划数量" prop="quantity">
          <el-input-number v-model="formData.quantity" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width: 100%">
            <el-option label="草稿" :value="0" />
            <el-option label="已下达" :value="1" />
            <el-option label="已排程" :value="2" />
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

const searchForm = reactive({ plan_month: '', material_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive({
  id: '',
  plan_month: '',
  material_code: '',
  material_name: '',
  quantity: 0,
  status: 0
})

const formRules: FormRules = {
  plan_month: [{ required: true, message: '请选择计划月份', trigger: 'change' }],
  material_code: [{ required: true, message: '请输入物料编码', trigger: 'blur' }],
  material_name: [{ required: true, message: '请输入物料名称', trigger: 'blur' }],
  quantity: [{ required: true, message: '请输入计划数量', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 0: '草稿', 1: '已下达', 2: '已排程' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 0: 'info', 1: 'warning', 2: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/aps/mps/list', { params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize } })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.plan_month = ''; searchForm.material_name = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => {
  dialogTitle.value = '新增主生产计划'
  isEdit.value = false
  Object.assign(formData, { id: '', plan_month: '', material_code: '', material_name: '', quantity: 0, status: 0 })
  dialogVisible.value = true
}
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑主生产计划'
  isEdit.value = true
  Object.assign(formData, row)
  dialogVisible.value = true
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该计划吗？', '提示', { type: 'warning' })
    await request.delete(`/aps/mps/${row.id}`)
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
      try {
        if (isEdit.value) {
          await request.put(`/aps/mps/${formData.id}`, formData)
          ElMessage.success('更新成功')
        } else {
          await request.post('/aps/mps', formData)
          ElMessage.success('创建成功')
        }
        dialogVisible.value = false
        loadData()
      } catch (error) {
        ElMessage.error('操作失败')
      }
    }
  })
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.plan-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
