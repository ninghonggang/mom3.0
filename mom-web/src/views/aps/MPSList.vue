<template>
  <div class="mps-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="计划编号">
          <el-input v-model="searchForm.plan_no" placeholder="请输入计划编号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待下达" :value="1" />
            <el-option label="已下达" :value="2" />
            <el-option label="执行中" :value="3" />
            <el-option label="已完成" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:mps:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('aps:mps:run')" @click="handleRun">
        <el-icon><VideoPlay /></el-icon>运行MPS
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="plan_no" label="计划编号" width="140" />
        <el-table-column prop="plan_name" label="计划名称" min-width="150" />
        <el-table-column prop="plan_year" label="计划年份" width="100" />
        <el-table-column prop="plan_month" label="计划月份" width="100" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_by" label="创建人" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:mps:detail')" @click="handleDetail(row)">明细</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('aps:mps:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:mps:delete')" @click="handleDelete(row)">删除</el-button>
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
import { ElMessage, ElMessageBox } from 'element-plus'
import { getMPSList } from '@/api/aps'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ plan_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待下达', 2: '已下达', 3: '执行中', 4: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'primary', 4: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMPSList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.plan_no = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增MPS计划') }
const handleRun = () => { ElMessage.info('运行MPS') }
const handleDetail = (row: any) => { ElMessage.info('查看明细') }
const handleEdit = (row: any) => { ElMessage.info('编辑') }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该计划吗？', '提示', { type: 'warning' })
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.mps-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
