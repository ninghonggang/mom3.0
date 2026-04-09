<template>
  <div class="login-log-list">
    <!-- 搜索区域 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="用户名">
          <el-input v-model="searchForm.username" placeholder="请输入用户名" clearable />
        </el-form-item>
        <el-form-item label="登录状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="成功" :value="1" />
            <el-option label="失败" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="登录时间">
          <el-date-picker
            v-model="dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 工具栏 -->
    <el-card class="toolbar-card">
      <el-button type="danger" v-if="hasPermission('system:loginlog:clean')" @click="handleClean">
        <el-icon><Delete /></el-icon>清理日志
      </el-button>
      <el-button type="warning" v-if="hasPermission('system:loginlog:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <!-- 表格 -->
    <el-card>
      <el-table
        v-loading="loading"
        :data="tableData"
      >
        <el-table-column prop="username" label="用户名" min-width="120" />
        <el-table-column prop="ip" label="登录IP" min-width="140" />
        <el-table-column prop="location" label="登录地点" min-width="150" />
        <el-table-column prop="browser" label="浏览器" min-width="120" />
        <el-table-column prop="os" label="操作系统" min-width="120" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="msg" label="提示消息" min-width="150" show-overflow-tooltip />
        <el-table-column prop="login_time" label="登录时间" width="180" />
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
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
import { ElMessage, ElMessageBox } from 'element-plus'
import { getLoginLogList, cleanLoginLog, exportLoginLog } from '@/api/system'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dateRange = ref<[string, string] | null>(null)

const searchForm = reactive({
  username: '',
  status: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (dateRange.value) {
      params.start_date = dateRange.value[0]
      params.end_date = dateRange.value[1]
    }
    const res = await getLoginLogList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.username = ''
  searchForm.status = ''
  dateRange.value = null
  handleSearch()
}

const handleClean = async () => {
  await ElMessageBox.confirm('确定要清理所有登录日志吗？此操作不可恢复！', '警告', { type: 'warning' })
  await cleanLoginLog()
  ElMessage.success('清理成功')
  loadData()
}

const handleExport = async () => {
  await exportLoginLog(searchForm)
  ElMessage.success('导出成功')
}

onMounted(() => {
  loadData()
})
</script>

<script lang="ts">
export default { name: 'LoginLogList' }
</script>

<style scoped lang="scss">
.login-log-list {
  .search-card, .toolbar-card {
    margin-bottom: 16px;
  }

  .toolbar-card :deep(.el-card__body) {
    padding: 12px 16px;
  }

  .pagination {
    margin-top: 16px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>
