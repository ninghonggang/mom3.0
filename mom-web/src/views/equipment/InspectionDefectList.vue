<template>
  <div class="inspection-defect-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="缺陷编码">
          <el-input v-model="searchForm.defect_code" placeholder="请输入编码" clearable />
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.equipment_name" placeholder="请输入设备名称" clearable />
        </el-form-item>
        <el-form-item label="严重程度">
          <el-select v-model="searchForm.severity" placeholder="请选择" clearable>
            <el-option label="轻微" value="MINOR" />
            <el-option label="一般" value="NORMAL" />
            <el-option label="严重" value="SERIOUS" />
            <el-option label="紧急" value="URGENT" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待处理" value="PENDING" />
            <el-option label="已指派" value="ASSIGNED" />
            <el-option label="处理中" value="PROCESSING" />
            <el-option label="已解决" value="RESOLVED" />
            <el-option label="已关闭" value="CLOSED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('equipment:inspection:defect:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="defect_code" label="缺陷编码" width="140" />
        <el-table-column prop="equipment_name" label="设备名称" min-width="120" />
        <el-table-column prop="equipment_code" label="设备编码" width="120" />
        <el-table-column prop="inspection_record_code" label="点检记录" width="140" />
        <el-table-column prop="defect_description" label="缺陷描述" min-width="180" show-overflow-tooltip />
        <el-table-column prop="severity" label="严重程度" width="100">
          <template #default="{ row }">
            <el-tag :type="getSeverityType(row.severity)">
              {{ getSeverityText(row.severity) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="reporter_name" label="上报人" width="100" />
        <el-table-column prop="assignee_name" label="指派给" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="160">
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('equipment:inspection:defect:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('equipment:inspection:defect:assign') && row.status === 'PENDING'" @click="handleAssign(row)">指派</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('equipment:inspection:defect:resolve') && row.status === 'ASSIGNED'" @click="handleResolve(row)">处理</el-button>
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

    <!-- 新增/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="缺陷编码" prop="defect_code">
          <el-input v-model="form.defect_code" placeholder="请输入编码" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="设备名称" prop="equipment_name">
          <el-input v-model="form.equipment_name" placeholder="请输入设备名称" />
        </el-form-item>
        <el-form-item label="设备编码" prop="equipment_code">
          <el-input v-model="form.equipment_code" placeholder="请输入设备编码" />
        </el-form-item>
        <el-form-item label="点检记录" prop="inspection_record_code">
          <el-input v-model="form.inspection_record_code" placeholder="请输入点检记录编码" />
        </el-form-item>
        <el-form-item label="缺陷描述" prop="defect_description">
          <el-input v-model="form.defect_description" type="textarea" :rows="3" placeholder="请输入缺陷描述" />
        </el-form-item>
        <el-form-item label="严重程度" prop="severity">
          <el-select v-model="form.severity" placeholder="请选择严重程度">
            <el-option label="轻微" value="MINOR" />
            <el-option label="一般" value="NORMAL" />
            <el-option label="严重" value="SERIOUS" />
            <el-option label="紧急" value="URGENT" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="form.status" placeholder="请选择状态">
            <el-option label="待处理" value="PENDING" />
            <el-option label="已指派" value="ASSIGNED" />
            <el-option label="处理中" value="PROCESSING" />
            <el-option label="已解决" value="RESOLVED" />
            <el-option label="已关闭" value="CLOSED" />
          </el-select>
        </el-form-item>
        <el-form-item label="指派给" prop="assignee_id">
          <el-input-number v-model="form.assignee_id" :min="0" placeholder="请输入指派人ID" />
        </el-form-item>
        <el-form-item label="处理结果" prop="resolution">
          <el-input v-model="form.resolution" type="textarea" :rows="3" placeholder="请输入处理结果" />
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
import { ElMessage, ElMessageBox } from 'element-plus'
import { getInspectionDefectList, createInspectionDefect, updateInspectionDefect, assignInspectionDefect, resolveInspectionDefect } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('新增缺陷')
const isEdit = ref(false)
const formRef = ref()

const searchForm = reactive({
  defect_code: '',
  equipment_name: '',
  severity: '',
  status: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const form = reactive({
  id: 0,
  defect_code: '',
  equipment_name: '',
  equipment_code: '',
  inspection_record_code: '',
  defect_description: '',
  severity: 'NORMAL',
  status: 'PENDING',
  assignee_id: 0,
  resolution: ''
})

const rules = {
  defect_code: [{ required: true, message: '请输入缺陷编码', trigger: 'blur' }],
  defect_description: [{ required: true, message: '请输入缺陷描述', trigger: 'blur' }]
}

const getSeverityType = (severity: string) => {
  const map: Record<string, string> = {
    MINOR: 'info',
    NORMAL: '',
    SERIOUS: 'warning',
    URGENT: 'danger'
  }
  return map[severity] || ''
}

const getSeverityText = (severity: string) => {
  const map: Record<string, string> = { MINOR: '轻微', NORMAL: '一般', SERIOUS: '严重', URGENT: '紧急' }
  return map[severity] || severity
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    PENDING: 'info',
    ASSIGNED: 'primary',
    PROCESSING: 'warning',
    RESOLVED: 'success',
    CLOSED: 'info'
  }
  return map[status] || ''
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    PENDING: '待处理',
    ASSIGNED: '已指派',
    PROCESSING: '处理中',
    RESOLVED: '已解决',
    CLOSED: '已关闭'
  }
  return map[status] || status
}

const formatDateTime = (datetime: string) => {
  if (!datetime) return ''
  return datetime.replace('T', ' ').substring(0, 19)
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getInspectionDefectList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.defect_code = ''
  searchForm.equipment_name = ''
  searchForm.severity = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增缺陷'
  Object.assign(form, {
    id: 0,
    defect_code: '',
    equipment_name: '',
    equipment_code: '',
    inspection_record_code: '',
    defect_description: '',
    severity: 'NORMAL',
    status: 'PENDING',
    assignee_id: 0,
    resolution: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑缺陷'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateInspectionDefect(form.id, form)
      ElMessage.success('更新成功')
    } else {
      await createInspectionDefect(form)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } catch (error: any) {
    ElMessage.error(error.message || '操作失败')
  }
}

const handleAssign = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定指派该缺陷吗？', '提示', { type: 'warning' })
    await assignInspectionDefect(row.id, { assignee_id: row.assignee_id || 1 })
    ElMessage.success('指派成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '指派失败')
    }
  }
}

const handleResolve = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定处理该缺陷吗？', '提示', { type: 'warning' })
    await resolveInspectionDefect(row.id, { resolution: row.resolution || '已处理' })
    ElMessage.success('处理成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '处理失败')
    }
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.inspection-defect-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
