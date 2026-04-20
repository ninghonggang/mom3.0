<template>
  <div class="execution-log-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="配置ID">
          <el-input v-model="searchForm.config_id" placeholder="请输入配置ID" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="成功" :value="1" />
            <el-option label="失败" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="执行时间">
          <el-date-picker
            v-model="searchForm.execution_time_range"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
            clearable
          />
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
        <el-table-column prop="config_id" label="配置ID" width="100" />
        <el-table-column prop="execution_time" label="执行时间" width="180" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="request_data" label="请求数据" min-width="180" show-overflow-tooltip />
        <el-table-column prop="response_data" label="响应数据" min-width="180" show-overflow-tooltip />
        <el-table-column prop="error_msg" label="错误信息" min-width="180" show-overflow-tooltip>
          <template #default="{ row }">
            <span v-if="row.error_msg" style="color: var(--el-color-danger)">{{ row.error_msg }}</span>
            <span v-else style="color: var(--el-color-success)">-</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              size="small"
              v-if="hasPermission('integration:execution-log:detail')"
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

    <!-- 详情弹窗 -->
    <el-dialog v-model="detailVisible" title="执行日志详情" width="700px">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="ID">{{ currentRow?.id }}</el-descriptions-item>
        <el-descriptions-item label="配置ID">{{ currentRow?.config_id }}</el-descriptions-item>
        <el-descriptions-item label="执行时间">{{ currentRow?.execution_time }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="currentRow?.status === 1 ? 'success' : 'danger'">
            {{ currentRow?.status === 1 ? '成功' : '失败' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="请求数据">
          <pre class="json-pre">{{ formatJson(currentRow?.request_data) }}</pre>
        </el-descriptions-item>
        <el-descriptions-item label="响应数据">
          <pre class="json-pre">{{ formatJson(currentRow?.response_data) }}</pre>
        </el-descriptions-item>
        <el-descriptions-item label="错误信息" v-if="currentRow?.error_msg">
          <span style="color: var(--el-color-danger)">{{ currentRow?.error_msg }}</span>
        </el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>
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
const detailVisible = ref(false)
const currentRow = ref<any>(null)

const searchForm = reactive({
  config_id: '',
  status: '',
  execution_time_range: [] as string[]
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formatJson = (str: string) => {
  if (!str) return '-'
  try {
    return JSON.stringify(JSON.parse(str), null, 2)
  } catch {
    return str
  }
}

const loadData = async () => {
  loading.value = true
  try {
    const params: Record<string, any> = {
      config_id: searchForm.config_id,
      status: searchForm.status,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.execution_time_range?.length === 2) {
      params.start_time = searchForm.execution_time_range[0]
      params.end_time = searchForm.execution_time_range[1]
    }
    const res = await request.get('/integration/execution-log/list', { params })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch {
    ElMessage.error('加载执行日志列表失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.config_id = ''
  searchForm.status = ''
  searchForm.execution_time_range = []
  handleSearch()
}
const handleDetail = (row: any) => {
  currentRow.value = row
  detailVisible.value = true
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.execution-log-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .json-pre {
    margin: 0;
    white-space: pre-wrap;
    word-break: break-all;
    font-size: 12px;
    max-height: 200px;
    overflow-y: auto;
    background: var(--el-fill-color-light);
    padding: 8px;
    border-radius: 4px;
  }
}
</style>
