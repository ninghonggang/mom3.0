<template>
  <div class="agv-task-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="任务编号">
          <el-input v-model="searchForm.task_no" placeholder="请输入任务编号" clearable />
        </el-form-item>
        <el-form-item label="AGV编号">
          <el-input v-model="searchForm.agv_id" placeholder="请输入AGV编号" clearable />
        </el-form-item>
        <el-form-item label="任务类型">
          <el-select v-model="searchForm.task_type" placeholder="请选择" clearable>
            <el-option label="搬运" value="transport" />
            <el-option label="充电" value="charge" />
            <el-option label="待机" value="idle" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待执行" :value="0" />
            <el-option label="执行中" :value="1" />
            <el-option label="已完成" :value="2" />
            <el-option label="已取消" :value="3" />
            <el-option label="异常" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" border stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="task_no" label="任务编号" width="160" />
        <el-table-column prop="agv_id" label="AGV编号" width="120" />
        <el-table-column prop="task_type" label="任务类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTaskTypeText(row.task_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="start_location" label="起始库位" min-width="120" />
        <el-table-column prop="end_location" label="目标库位" min-width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              size="small"
              v-if="hasPermission('agv:task:detail')"
              @click="handleDetail(row)"
            >
              详情
            </el-button>
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/stores/auth'
import request from '@/utils/request'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({
  task_no: '',
  agv_id: '',
  task_type: '',
  status: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getTaskTypeText = (type: string) => {
  const map: Record<string, string> = {
    transport: '搬运',
    charge: '充电',
    idle: '待机'
  }
  return map[type] || type || '未知'
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 0: '待执行', 1: '执行中', 2: '已完成', 3: '已取消', 4: '异常' }
  return map[status] ?? '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = {
    0: 'info',
    1: 'primary',
    2: 'success',
    3: 'warning',
    4: 'danger'
  }
  return map[status] ?? 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/agv/task/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch {
    ElMessage.error('加载AGV任务列表失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.task_no = ''
  searchForm.agv_id = ''
  searchForm.task_type = ''
  searchForm.status = ''
  handleSearch()
}
const handleDetail = (row: any) => { ElMessage.info(`任务编号：${row.task_no}`) }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.agv-task-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
