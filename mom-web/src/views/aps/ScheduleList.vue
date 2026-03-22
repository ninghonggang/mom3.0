<template>
  <div class="schedule-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="计划编号">
          <el-input v-model="searchForm.plan_no" placeholder="请输入计划编号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待执行" :value="1" />
            <el-option label="执行中" :value="2" />
            <el-option label="已完成" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="dialogVisible = true; formData = {}">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="plan_no" label="计划编号" width="140" />
        <el-table-column prop="plan_type" label="计划类型" width="90" />
        <el-table-column prop="algorithm" label="算法" width="90" />
        <el-table-column prop="start_date" label="开始日期" width="110" />
        <el-table-column prop="end_date" label="结束日期" width="110" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_by" label="创建人" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="210" fixed="right">
          <template #default="{ row }">
            <el-button link type="success" size="small" @click="handleExecute(row)" v-if="row.status === 1">执行</el-button>
            <el-button link type="primary" size="small" @click="handleResults(row)">结果</el-button>
            <el-button link type="danger" size="small" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新增对话框 -->
    <el-dialog v-model="dialogVisible" title="新增排程计划" width="500px">
      <el-form :model="formData" label-width="100px">
        <el-form-item label="计划编号">
          <el-input v-model="formData.plan_no" placeholder="自动生成" disabled />
        </el-form-item>
        <el-form-item label="计划类型" required>
          <el-select v-model="formData.plan_type" placeholder="请选择" style="width: 100%">
            <el-option label="粗排" value="粗排" />
            <el-option label="细排" value="细排" />
          </el-select>
        </el-form-item>
        <el-form-item label="算法">
          <el-select v-model="formData.algorithm" placeholder="请选择" style="width: 100%">
            <el-option label="遗传算法" value="遗传" />
            <el-option label="粒子群" value="粒子群" />
            <el-option label="启发式" value="启发式" />
          </el-select>
        </el-form-item>
        <el-form-item label="开始日期">
          <el-date-picker v-model="formData.start_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
        </el-form-item>
        <el-form-item label="结束日期">
          <el-date-picker v-model="formData.end_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getScheduleList, createSchedule, executeSchedule, deleteSchedule } from '@/api/aps'

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const tableData = ref<any[]>([])
const formData = ref<any>({})

const searchForm = reactive({ plan_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待执行', 2: '执行中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getScheduleList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.plan_no = ''; searchForm.status = ''; handleSearch() }

const handleSave = async () => {
  saveLoading.value = true
  try {
    await createSchedule(formData.value)
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleExecute = async (row: any) => {
  await executeSchedule(row.id)
  ElMessage.success('执行完成')
  loadData()
}

const handleResults = (row: any) => { ElMessage.info('查看排程结果') }

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该计划吗？', '提示', { type: 'warning' })
  await deleteSchedule(row.id)
  ElMessage.success('删除成功')
  loadData()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.schedule-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
