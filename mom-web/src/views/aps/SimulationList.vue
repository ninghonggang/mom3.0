<template>
  <div class="simulation-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="模拟编号">
          <el-input v-model="searchForm.sim_no" placeholder="请输入模拟编号" clearable />
        </el-form-item>
        <el-form-item label="模拟名称">
          <el-input v-model="searchForm.sim_name" placeholder="请输入模拟名称" clearable />
        </el-form-item>
        <el-form-item label="产品族">
          <el-input v-model="searchForm.product_family" placeholder="请输入产品族" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" :value="1" />
            <el-option label="已模拟" :value="2" />
            <el-option label="已确认" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:simulation:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('aps:simulation:run')" @click="handleRun()">
        <el-icon><VideoPlay /></el-icon>运行模拟
      </el-button>
      <el-button type="warning" v-if="hasPermission('aps:simulation:confirm')" @click="handleConfirm()">
        <el-icon><Check /></el-icon>确认
      </el-button>
      <el-button type="info" v-if="hasPermission('aps:simulation:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="sim_no" label="模拟编号" width="140" />
        <el-table-column prop="sim_name" label="模拟名称" min-width="160" />
        <el-table-column prop="product_family" label="产品族" width="120" />
        <el-table-column prop="sim_date" label="模拟日期" width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="160" show-overflow-tooltip />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:simulation:detail')" @click="handleDetail(row)">明细</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('aps:simulation:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:simulation:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="模拟编号" prop="sim_no">
          <el-input v-model="formData.sim_no" placeholder="请输入模拟编号" />
        </el-form-item>
        <el-form-item label="模拟名称" prop="sim_name">
          <el-input v-model="formData.sim_name" placeholder="请输入模拟名称" />
        </el-form-item>
        <el-form-item label="产品族" prop="product_family">
          <el-input v-model="formData.product_family" placeholder="请输入产品族" />
        </el-form-item>
        <el-form-item label="模拟日期" prop="sim_date">
          <el-date-picker v-model="formData.sim_date" type="date" placeholder="请选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width: 100%">
            <el-option label="草稿" :value="1" />
            <el-option label="已模拟" :value="2" />
            <el-option label="已确认" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="formData.remark" type="textarea" :rows="3" placeholder="请输入备注" />
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
import { useAuthStore } from '@/stores/auth'
import request from '@/utils/request'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const searchForm = reactive({ sim_no: '', sim_name: '', product_family: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive({
  id: '',
  sim_no: '',
  sim_name: '',
  product_family: '',
  sim_date: '',
  status: 1,
  remark: ''
})

const formRules: FormRules = {
  sim_no: [{ required: true, message: '请输入模拟编号', trigger: 'blur' }],
  sim_name: [{ required: true, message: '请输入模拟名称', trigger: 'blur' }],
  product_family: [{ required: true, message: '请输入产品族', trigger: 'blur' }],
  sim_date: [{ required: true, message: '请选择模拟日期', trigger: 'change' }]
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '草稿', 2: '已模拟', 3: '已确认' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/aps/simulation/list', { params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize } })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.sim_no = ''; searchForm.sim_name = ''; searchForm.product_family = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => {
  dialogTitle.value = '新增模拟'
  isEdit.value = false
  Object.assign(formData, { id: '', sim_no: '', sim_name: '', product_family: '', sim_date: '', status: 1, remark: '' })
  dialogVisible.value = true
}
const handleRun = async () => {
  try {
    await request.post('/aps/simulation/run')
    ElMessage.success('模拟运行成功')
    loadData()
  } catch { ElMessage.error('模拟运行失败') }
}
const handleConfirm = async () => {
  try {
    await request.post('/aps/simulation/confirm')
    ElMessage.success('确认成功')
    loadData()
  } catch { ElMessage.error('确认失败') }
}
const handleExport = () => { ElMessage.info('导出功能') }
const handleDetail = (row: any) => { ElMessage.info('查看明细') }
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑模拟'
  isEdit.value = true
  Object.assign(formData, row)
  dialogVisible.value = true
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该模拟记录吗？', '提示', { type: 'warning' })
    await request.delete(`/aps/simulation/${row.id}`)
    ElMessage.success('删除成功')
    loadData()
  } catch { /* cancelled or error */ }
}
const handleDialogClose = () => { formRef.value?.resetFields() }
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await request.put(`/aps/simulation/${formData.id}`, formData)
        } else {
          await request.post('/aps/simulation', formData)
        }
        ElMessage.success('操作成功')
        dialogVisible.value = false
        loadData()
      } catch { ElMessage.error('操作失败') }
    }
  })
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.simulation-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
