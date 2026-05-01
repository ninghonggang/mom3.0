<template>
  <div class="capacity-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="车间">
          <el-input v-model="searchForm.workshop_name" placeholder="请输入车间名称" clearable />
        </el-form-item>
        <el-form-item label="产线">
          <el-input v-model="searchForm.line_name" placeholder="请输入产线名称" clearable />
        </el-form-item>
        <el-form-item label="日期范围">
          <el-date-picker v-model="searchForm.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" value-format="YYYY-MM-DD" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:capacity:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('aps:capacity:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="workshop_name" label="车间" width="120" />
        <el-table-column prop="line_name" label="产线" width="120" />
        <el-table-column prop="work_center_name" label="工作中心" width="140" />
        <el-table-column prop="date" label="日期" width="120" />
        <el-table-column prop="capacity" label="额定产能" width="100" />
        <el-table-column prop="actual_capacity" label="实际产能" width="100" />
        <el-table-column prop="utilization_rate" label="利用率" width="100">
          <template #default="{ row }">
            <span :class="getUtilizationClass(row.utilization_rate)">{{ row.utilization_rate }}%</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:capacity:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:capacity:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="车间" prop="workshop_name">
          <el-input v-model="formData.workshop_name" placeholder="请输入车间名称" />
        </el-form-item>
        <el-form-item label="产线" prop="line_name">
          <el-input v-model="formData.line_name" placeholder="请输入产线名称" />
        </el-form-item>
        <el-form-item label="工作中心" prop="work_center_name">
          <el-input v-model="formData.work_center_name" placeholder="请输入工作中心" />
        </el-form-item>
        <el-form-item label="日期" prop="date">
          <el-date-picker v-model="formData.date" type="date" placeholder="请选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="额定产能" prop="capacity">
          <el-input-number v-model="formData.capacity" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="实际产能" prop="actual_capacity">
          <el-input-number v-model="formData.actual_capacity" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width: 100%">
            <el-option label="正常" :value="1" />
            <el-option label="异常" :value="2" />
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

const searchForm = reactive({ workshop_name: '', line_name: '', dateRange: [] as string[] })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive({
  id: '',
  workshop_name: '',
  line_name: '',
  work_center_name: '',
  date: '',
  capacity: 0,
  actual_capacity: 0,
  status: 1
})

const formRules: FormRules = {
  workshop_name: [{ required: true, message: '请输入车间名称', trigger: 'blur' }],
  line_name: [{ required: true, message: '请输入产线名称', trigger: 'blur' }],
  work_center_name: [{ required: true, message: '请输入工作中心', trigger: 'blur' }],
  date: [{ required: true, message: '请选择日期', trigger: 'change' }],
  capacity: [{ required: true, message: '请输入额定产能', trigger: 'blur' }],
  actual_capacity: [{ required: true, message: '请输入实际产能', trigger: 'blur' }]
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '正常', 2: '异常' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'success', 2: 'danger' }
  return map[status] || 'info'
}

const getUtilizationClass = (rate: number) => {
  if (rate >= 80) return 'text-success'
  if (rate >= 50) return 'text-warning'
  return 'text-danger'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.workshop_name) params.workshop_name = searchForm.workshop_name
    if (searchForm.line_name) params.line_name = searchForm.line_name
    if (searchForm.dateRange?.length === 2) {
      params.start_date = searchForm.dateRange[0]
      params.end_date = searchForm.dateRange[1]
    }
    const res = await request.get('/aps/capacity/list', { params })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.workshop_name = ''; searchForm.line_name = ''; searchForm.dateRange = []; handleSearch() }
const handleAdd = () => {
  dialogTitle.value = '新增产能记录'
  isEdit.value = false
  Object.assign(formData, { id: '', workshop_name: '', line_name: '', work_center_name: '', date: '', capacity: 0, actual_capacity: 0, status: 1 })
  dialogVisible.value = true
}
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑产能记录'
  isEdit.value = true
  Object.assign(formData, row)
  dialogVisible.value = true
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该记录吗？', '提示', { type: 'warning' })
    await request.delete(`/aps/capacity/${row.id}`)
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
          await request.put(`/aps/capacity/${formData.id}`, formData)
          ElMessage.success('更新成功')
        } else {
          await request.post('/aps/capacity', formData)
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
.capacity-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .text-success { color: #67c23a; }
  .text-warning { color: #e6a23c; }
  .text-danger { color: #f56c6c; }
}
</style>
