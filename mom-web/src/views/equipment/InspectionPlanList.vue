<template>
  <div class="inspection-plan-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="计划编号">
          <el-input v-model="searchForm.plan_code" placeholder="请输入计划编号" clearable />
        </el-form-item>
        <el-form-item label="计划名称">
          <el-input v-model="searchForm.plan_name" placeholder="请输入计划名称" clearable />
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.equipment_name" placeholder="请输入设备名称" clearable />
        </el-form-item>
        <el-form-item label="点检类型">
          <el-select v-model="searchForm.inspection_type" placeholder="请选择" clearable>
            <el-option label="日常点检" value="ROUTINE" />
            <el-option label="定期点检" value="PERIODIC" />
            <el-option label="专项点检" value="SPECIAL" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" value="DRAFT" />
            <el-option label="已发布" value="PUBLISHED" />
            <el-option label="进行中" value="IN_PROGRESS" />
            <el-option label="已完成" value="COMPLETED" />
            <el-option label="已取消" value="CANCELLED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('equipment:inspection:plan:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="plan_code" label="计划编号" width="140" />
        <el-table-column prop="plan_name" label="计划名称" min-width="150" />
        <el-table-column prop="equipment_name" label="设备名称" width="120" />
        <el-table-column prop="inspection_type" label="点检类型" width="100">
          <template #default="{ row }">
            {{ getTypeText(row.inspection_type) }}
          </template>
        </el-table-column>
        <el-table-column prop="frequency" label="点检频率" width="100">
          <template #default="{ row }">
            {{ getFrequencyText(row.frequency) }}
          </template>
        </el-table-column>
        <el-table-column prop="start_date" label="开始日期" width="120" />
        <el-table-column prop="end_date" label="结束日期" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('equipment:inspection:plan:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('equipment:inspection:plan:assign') && row.status === 'PUBLISHED'" @click="handleAssign(row)">指派</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('equipment:inspection:plan:delete')" @click="handleDelete(row)">删除</el-button>
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
      <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="计划编号" prop="plan_code">
          <el-input v-model="form.plan_code" placeholder="请输入计划编号" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="计划名称" prop="plan_name">
          <el-input v-model="form.plan_name" placeholder="请输入计划名称" />
        </el-form-item>
        <el-form-item label="设备" prop="equipment_id">
          <el-select v-model="form.equipment_id" placeholder="请选择设备" filterable>
            <el-option label="设备A" :value="1" />
            <el-option label="设备B" :value="2" />
            <el-option label="设备C" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="点检类型" prop="inspection_type">
          <el-select v-model="form.inspection_type" placeholder="请选择点检类型">
            <el-option label="日常点检" value="ROUTINE" />
            <el-option label="定期点检" value="PERIODIC" />
            <el-option label="专项点检" value="SPECIAL" />
          </el-select>
        </el-form-item>
        <el-form-item label="点检频率" prop="frequency">
          <el-select v-model="form.frequency" placeholder="请选择点检频率">
            <el-option label="每天" value="DAILY" />
            <el-option label="每周" value="WEEKLY" />
            <el-option label="每月" value="MONTHLY" />
            <el-option label="每季" value="QUARTERLY" />
            <el-option label="每年" value="YEARLY" />
          </el-select>
        </el-form-item>
        <el-form-item label="开始日期" prop="start_date">
          <el-date-picker v-model="form.start_date" type="date" placeholder="请选择开始日期" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item label="结束日期" prop="end_date">
          <el-date-picker v-model="form.end_date" type="date" placeholder="请选择结束日期" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio label="DRAFT">草稿</el-radio>
            <el-radio label="PUBLISHED">已发布</el-radio>
            <el-radio label="IN_PROGRESS">进行中</el-radio>
            <el-radio label="COMPLETED">已完成</el-radio>
            <el-radio label="CANCELLED">已取消</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 指派对话框 -->
    <el-dialog v-model="assignDialogVisible" title="指派点检计划" width="500px">
      <el-form :model="assignForm" label-width="100px">
        <el-form-item label="计划名称">
          <span>{{ assignForm.plan_name }}</span>
        </el-form-item>
        <el-form-item label="指派给">
          <el-select v-model="assignForm.assignee_id" placeholder="请选择执行人" filterable>
            <el-option label="张三" :value="1" />
            <el-option label="李四" :value="2" />
            <el-option label="王五" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="执行时间">
          <el-date-picker v-model="assignForm.execute_time" type="datetime" placeholder="请选择执行时间" value-format="YYYY-MM-DD HH:mm:ss" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="assignDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleAssignSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getInspectionPlanList, createInspectionPlan, updateInspectionPlan, deleteInspectionPlan, assignInspectionPlan } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const assignDialogVisible = ref(false)
const dialogTitle = ref('新增点检计划')
const isEdit = ref(false)
const formRef = ref()

const searchForm = reactive({ plan_code: '', plan_name: '', equipment_name: '', inspection_type: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const form = reactive({
  id: 0,
  plan_code: '',
  plan_name: '',
  equipment_id: 0,
  equipment_name: '',
  inspection_type: 'ROUTINE',
  frequency: 'DAILY',
  start_date: '',
  end_date: '',
  status: 'DRAFT',
  description: ''
})

const assignForm = reactive({
  plan_id: 0,
  plan_name: '',
  assignee_id: 0,
  execute_time: ''
})

const rules = {
  plan_code: [{ required: true, message: '请输入计划编号', trigger: 'blur' }],
  plan_name: [{ required: true, message: '请输入计划名称', trigger: 'blur' }],
  equipment_id: [{ required: true, message: '请选择设备', trigger: 'change' }],
  inspection_type: [{ required: true, message: '请选择点检类型', trigger: 'change' }]
}

const getTypeText = (type: string) => {
  const map: Record<string, string> = { ROUTINE: '日常点检', PERIODIC: '定期点检', SPECIAL: '专项点检' }
  return map[type] || type
}

const getFrequencyText = (frequency: string) => {
  const map: Record<string, string> = { DAILY: '每天', WEEKLY: '每周', MONTHLY: '每月', QUARTERLY: '每季', YEARLY: '每年' }
  return map[frequency] || frequency
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = { DRAFT: '草稿', PUBLISHED: '已发布', IN_PROGRESS: '进行中', COMPLETED: '已完成', CANCELLED: '已取消' }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = { DRAFT: 'info', PUBLISHED: 'success', IN_PROGRESS: 'warning', COMPLETED: 'success', CANCELLED: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getInspectionPlanList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
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
  searchForm.plan_code = ''
  searchForm.plan_name = ''
  searchForm.equipment_name = ''
  searchForm.inspection_type = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增点检计划'
  Object.assign(form, {
    id: 0,
    plan_code: '',
    plan_name: '',
    equipment_id: 0,
    equipment_name: '',
    inspection_type: 'ROUTINE',
    frequency: 'DAILY',
    start_date: '',
    end_date: '',
    status: 'DRAFT',
    description: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑点检计划'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleAssign = (row: any) => {
  assignForm.plan_id = row.id
  assignForm.plan_name = row.plan_name
  assignForm.assignee_id = 0
  assignForm.execute_time = ''
  assignDialogVisible.value = true
}

const handleAssignSubmit = async () => {
  try {
    await assignInspectionPlan(assignForm.plan_id, {
      assignee_id: assignForm.assignee_id,
      execute_time: assignForm.execute_time
    })
    ElMessage.success('指派成功')
    assignDialogVisible.value = false
    loadData()
  } catch (error: any) {
    ElMessage.error(error.message || '指派失败')
  }
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateInspectionPlan(form.id, form)
      ElMessage.success('更新成功')
    } else {
      await createInspectionPlan(form)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } catch (error: any) {
    ElMessage.error(error.message || '操作失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该点检计划吗？', '提示', { type: 'warning' })
    await deleteInspectionPlan(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.inspection-plan-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
