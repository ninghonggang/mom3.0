<template>
  <div class="oper-log-list">
    <!-- 搜索区域 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="操作人员">
          <el-input v-model="searchForm.operator" placeholder="请输入操作人员" clearable />
        </el-form-item>
        <el-form-item label="操作类型">
          <el-select v-model="searchForm.business_type" placeholder="请选择" clearable>
            <el-option label="新增" :value="1" />
            <el-option label="修改" :value="2" />
            <el-option label="删除" :value="3" />
            <el-option label="查询" :value="4" />
            <el-option label="导出" :value="5" />
            <el-option label="导入" :value="6" />
          </el-select>
        </el-form-item>
        <el-form-item label="操作状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="成功" :value="1" />
            <el-option label="失败" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="操作时间">
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
      <el-button type="danger" v-if="hasPermission('system:operlog:clean')" @click="handleClean">
        <el-icon><Delete /></el-icon>清理日志
      </el-button>
      <el-button type="warning" v-if="hasPermission('system:operlog:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <!-- 表格 -->
    <el-card>
      <el-table
        v-loading="loading"
        :data="tableData"
        @row-click="handleRowClick"
      >
        <el-table-column prop="operator" label="操作人员" min-width="120" />
        <el-table-column prop="business_type" label="操作类型" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.business_type === 1" type="success">新增</el-tag>
            <el-tag v-else-if="row.business_type === 2" type="warning">修改</el-tag>
            <el-tag v-else-if="row.business_type === 3" type="danger">删除</el-tag>
            <el-tag v-else-if="row.business_type === 4" type="info">查询</el-tag>
            <el-tag v-else-if="row.business_type === 5" type="info">导出</el-tag>
            <el-tag v-else-if="row.business_type === 6" type="info">导入</el-tag>
            <span v-else>{{ row.business_type }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="method" label="请求方法" min-width="150" show-overflow-tooltip />
        <el-table-column prop="url" label="请求地址" min-width="200" show-overflow-tooltip />
        <el-table-column prop="status" label="操作状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="ip" label="操作IP" min-width="140" />
        <el-table-column prop="oper_time" label="操作时间" width="180" />
        <el-table-column label="操作" width="80" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click.stop="handleDetail(row)">详情</el-button>
          </template>
        </el-table-column>
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

    <!-- 详情弹窗 -->
    <el-dialog v-model="detailVisible" title="操作日志详情" width="700px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="操作人员">{{ currentRow?.operator }}</el-descriptions-item>
        <el-descriptions-item label="操作类型">
          <el-tag v-if="currentRow?.business_type === 1" type="success">新增</el-tag>
          <el-tag v-else-if="currentRow?.business_type === 2" type="warning">修改</el-tag>
          <el-tag v-else-if="currentRow?.business_type === 3" type="danger">删除</el-tag>
          <el-tag v-else-if="currentRow?.business_type === 4" type="info">查询</el-tag>
          <el-tag v-else-if="currentRow?.business_type === 5" type="info">导出</el-tag>
          <el-tag v-else-if="currentRow?.business_type === 6" type="info">导入</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="请求方法">{{ currentRow?.method }}</el-descriptions-item>
        <el-descriptions-item label="请求地址">{{ currentRow?.url }}</el-descriptions-item>
        <el-descriptions-item label="操作IP">{{ currentRow?.ip }}</el-descriptions-item>
        <el-descriptions-item label="操作状态">
          <el-tag :type="currentRow?.status === 1 ? 'success' : 'danger'">
            {{ currentRow?.status === 1 ? '成功' : '失败' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="操作时间" :span="2">{{ currentRow?.oper_time }}</el-descriptions-item>
        <el-descriptions-item label="请求参数" :span="2">
          <pre style="max-height: 200px; overflow: auto; margin: 0;">{{ currentRow?.param }}</pre>
        </el-descriptions-item>
        <el-descriptions-item label="返回结果" :span="2">
          <pre style="max-height: 200px; overflow: auto; margin: 0;">{{ currentRow?.result }}</pre>
        </el-descriptions-item>
        <el-descriptions-item label="错误信息" :span="2" v-if="currentRow?.error">
          <span style="color: #f56c6c;">{{ currentRow?.error }}</span>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getOperLogList, cleanOperLog, exportOperLog } from '@/api/system'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dateRange = ref<[string, string] | null>(null)
const detailVisible = ref(false)
const currentRow = ref<any>(null)

const searchForm = reactive({
  operator: '',
  business_type: '',
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
    const res = await getOperLogList(params)
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
  searchForm.operator = ''
  searchForm.business_type = ''
  searchForm.status = ''
  dateRange.value = null
  handleSearch()
}

const handleRowClick = (row: any) => {
  currentRow.value = row
}

const handleDetail = (row: any) => {
  currentRow.value = row
  detailVisible.value = true
}

const handleClean = async () => {
  await ElMessageBox.confirm('确定要清理所有操作日志吗？此操作不可恢复！', '警告', { type: 'warning' })
  await cleanOperLog()
  ElMessage.success('清理成功')
  loadData()
}

const handleExport = async () => {
  await exportOperLog(searchForm)
  ElMessage.success('导出成功')
}

onMounted(() => {
  loadData()
})
</script>

<script lang="ts">
export default { name: 'OperLogList' }
</script>

<style scoped lang="scss">
.oper-log-list {
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
