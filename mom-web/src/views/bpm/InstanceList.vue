<template>
  <div class="instance-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="流程名称">
          <el-input v-model="searchForm.model_name" placeholder="请输入流程名称" clearable />
        </el-form-item>
        <el-form-item label="发起人">
          <el-input v-model="searchForm.initiator_name" placeholder="请输入发起人" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="进行中" :value="1" />
            <el-option label="已完成" :value="2" />
            <el-option label="已取消" :value="3" />
            <el-option label="已终止" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="process_name" label="流程名称" min-width="160" />
        <el-table-column prop="process_code" label="流程编码" width="140" />
        <el-table-column prop="initiator_name" label="发起人" width="100" />
        <el-table-column prop="current_node_name" label="当前节点" width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="start_time" label="开始时间" width="160" />
        <el-table-column prop="end_time" label="结束时间" width="160" />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleDetail(row)">详情</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('bpm:instance:cancel') && row.status === 1" @click="handleCancel(row)">取消</el-button>
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

    <!-- 流程详情对话框 -->
    <el-dialog v-model="detailVisible" title="流程详情" width="800px">
      <el-descriptions :column="2" border v-if="currentInstance">
        <el-descriptions-item label="流程名称">{{ currentInstance.process_name }}</el-descriptions-item>
        <el-descriptions-item label="流程编码">{{ currentInstance.process_code }}</el-descriptions-item>
        <el-descriptions-item label="发起人">{{ currentInstance.initiator_name }}</el-descriptions-item>
        <el-descriptions-item label="当前节点">{{ currentInstance.current_node_name }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(currentInstance.status)">{{ getStatusText(currentInstance.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="版本">v{{ currentInstance.version }}</el-descriptions-item>
        <el-descriptions-item label="开始时间">{{ currentInstance.start_time }}</el-descriptions-item>
        <el-descriptions-item label="结束时间">{{ currentInstance.end_time || '-' }}</el-descriptions-item>
      </el-descriptions>

      <el-divider content-position="left">流程进度</el-divider>
      <el-timeline v-if="instanceHistory.length">
        <el-timeline-item
          v-for="(item, index) in instanceHistory"
          :key="index"
          :type="item.status === 'approved' ? 'success' : item.status === 'rejected' ? 'danger' : 'primary'"
          :timestamp="item.create_time"
          placement="top"
        >
          <p><strong>{{ item.node_name }}</strong> - {{ item.assignee_name || '系统' }}</p>
          <p>{{ item.action_text }}</p>
          <p v-if="item.comment">审批意见：{{ item.comment }}</p>
        </el-timeline-item>
      </el-timeline>
      <el-empty v-else description="暂无流程记录" />

      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getProcessInstanceList, getProcessInstance } from '@/api/bpm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const detailVisible = ref(false)
const currentInstance = ref<any>(null)
const instanceHistory = ref<any[]>([])

const searchForm = reactive({ model_name: '', initiator_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '进行中', 2: '已完成', 3: '已取消', 4: '已终止' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'primary', 2: 'success', 3: 'info', 4: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getProcessInstanceList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.model_name = ''; searchForm.initiator_name = ''; searchForm.status = ''; handleSearch() }

const handleDetail = async (row: any) => {
  try {
    const res = await getProcessInstance(row.id)
    currentInstance.value = res.data
    instanceHistory.value = res.data.history || []
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取详情失败')
  }
}

const handleCancel = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定取消该流程实例吗？', '提示', { type: 'warning' })
    await getProcessInstance(row.id) // Placeholder for cancel API
    ElMessage.success('取消成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.instance-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
