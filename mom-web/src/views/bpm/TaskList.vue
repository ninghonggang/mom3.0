<template>
  <div class="task-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="流程名称">
          <el-input v-model="searchForm.process_name" placeholder="请输入流程名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-tabs v-model="activeTab" @tab-change="handleTabChange">
        <el-tab-pane label="待办任务" name="todo" />
        <el-tab-pane label="已办任务" name="done" />
      </el-tabs>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="process_name" label="流程名称" min-width="160" />
        <el-table-column prop="process_code" label="流程编码" width="140" />
        <el-table-column prop="node_name" label="当前节点" width="120" />
        <el-table-column prop="initiator_name" label="发起人" width="100" />
        <el-table-column prop="priority" label="优先级" width="80">
          <template #default="{ row }">
            <el-tag :type="getPriorityType(row.priority)" size="small">{{ getPriorityText(row.priority) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间" width="160" />
        <el-table-column prop="due_date" label="截止时间" width="120" />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleDetail(row)">详情</el-button>
            <el-button link type="success" size="small" v-if="activeTab === 'todo'" @click="handleApprove(row)">通过</el-button>
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

    <!-- 任务处理对话框 -->
    <el-dialog v-model="taskDialogVisible" title="任务处理" width="600px">
      <el-form ref="taskFormRef" :model="taskForm" label-width="100px">
        <el-form-item label="流程名称">
          <el-input :model-value="currentTask?.process_name" disabled />
        </el-form-item>
        <el-form-item label="当前节点">
          <el-input :model-value="currentTask?.node_name" disabled />
        </el-form-item>
        <el-form-item label="发起人">
          <el-input :model-value="currentTask?.initiator_name" disabled />
        </el-form-item>
        <el-form-item label="审批意见" prop="comment">
          <el-input v-model="taskForm.comment" type="textarea" :rows="4" placeholder="请输入审批意见" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="taskDialogVisible = false">取消</el-button>
        <el-button type="success" :loading="approveLoading" @click="handleSubmitApprove">通过</el-button>
        <el-button type="danger" :loading="rejectLoading" @click="handleSubmitReject">驳回</el-button>
      </template>
    </el-dialog>

    <!-- 任务详情对话框 -->
    <el-dialog v-model="detailVisible" title="任务详情" width="700px">
      <el-descriptions :column="2" border v-if="currentTask">
        <el-descriptions-item label="流程名称">{{ currentTask.process_name }}</el-descriptions-item>
        <el-descriptions-item label="流程编码">{{ currentTask.process_code }}</el-descriptions-item>
        <el-descriptions-item label="当前节点">{{ currentTask.node_name }}</el-descriptions-item>
        <el-descriptions-item label="发起人">{{ currentTask.initiator_name }}</el-descriptions-item>
        <el-descriptions-item label="优先级">
          <el-tag :type="getPriorityType(currentTask.priority)">{{ getPriorityText(currentTask.priority) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ currentTask.create_time }}</el-descriptions-item>
        <el-descriptions-item label="截止时间">{{ currentTask.due_date || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="currentTask.status === 1 ? 'primary' : 'success'">
            {{ currentTask.status === 1 ? '待处理' : '已处理' }}
          </el-tag>
        </el-descriptions-item>
      </el-descriptions>

      <el-divider content-position="left">表单数据</el-divider>
      <el-descriptions v-if="currentTask && currentTask.form_data" :column="1" border>
        <el-descriptions-item v-for="(val, key) in currentTask.form_data" :key="key" :label="String(key)">
          {{ val }}
        </el-descriptions-item>
      </el-descriptions>
      <el-empty v-else description="暂无表单数据" />

      <el-divider content-position="left">审批记录</el-divider>
      <el-timeline v-if="taskHistory.length">
        <el-timeline-item
          v-for="(item, index) in taskHistory"
          :key="index"
          :type="item.action === 'approve' ? 'success' : 'danger'"
          :timestamp="item.create_time"
          placement="top"
        >
          <p><strong>{{ item.node_name }}</strong> - {{ item.assignee_name }}</p>
          <p>{{ item.action_text }}</p>
          <p v-if="item.comment">审批意见：{{ item.comment }}</p>
        </el-timeline-item>
      </el-timeline>
      <el-empty v-else description="暂无审批记录" />

      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getTaskList, getTask, approveTask, rejectTask } from '@/api/bpm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const activeTab = ref('todo')
const taskDialogVisible = ref(false)
const detailVisible = ref(false)
const approveLoading = ref(false)
const rejectLoading = ref(false)
const currentTask = ref<any>(null)
const taskHistory = ref<any[]>([])

const searchForm = reactive({ process_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const taskForm = reactive({ comment: '' })

const getPriorityText = (priority: number) => {
  const map: Record<number, string> = { 1: '普通', 2: '紧急', 3: '加急' }
  return map[priority] || '普通'
}

const getPriorityType = (priority: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'danger' }
  return map[priority] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params = activeTab.value === 'todo'
      ? { ...searchForm, page: pagination.page, page_size: pagination.pageSize, status: 1 }
      : { ...searchForm, page: pagination.page, page_size: pagination.pageSize, status: 2 }
    const res = await getTaskList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.process_name = ''; handleSearch() }

const handleTabChange = () => {
  pagination.page = 1
  loadData()
}

const handleDetail = async (row: any) => {
  try {
    const res = await getTask(row.id)
    currentTask.value = res.data
    taskHistory.value = res.data.history || []
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取详情失败')
  }
}

const handleApprove = (row: any) => {
  currentTask.value = row
  taskForm.comment = ''
  taskDialogVisible.value = true
}

const handleSubmitApprove = async () => {
  if (!currentTask.value) return
  approveLoading.value = true
  try {
    await approveTask(currentTask.value.id, taskForm.comment)
    ElMessage.success('审批通过')
    taskDialogVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    approveLoading.value = false
  }
}

const handleSubmitReject = async () => {
  if (!currentTask.value) return
  rejectLoading.value = true
  try {
    await rejectTask(currentTask.value.id, taskForm.comment)
    ElMessage.success('已驳回')
    taskDialogVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    rejectLoading.value = false
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.task-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 0 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
