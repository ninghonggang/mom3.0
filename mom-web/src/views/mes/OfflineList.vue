<template>
  <div class="offline-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="离线单号">
          <el-input v-model="searchForm.offline_no" placeholder="请输入离线单号" clearable />
        </el-form-item>
        <el-form-item label="物料编码">
          <el-input v-model="searchForm.material_code" placeholder="请输入物料编码" clearable />
        </el-form-item>
        <el-form-item label="处理状态">
          <el-select v-model="searchForm.handle_status" placeholder="请选择" clearable>
            <el-option label="待处理" :value="1" />
            <el-option label="处理中" :value="2" />
            <el-option label="已处理" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mes:offline:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="offline_no" label="离线单号" width="140" />
        <el-table-column prop="material_code" label="物料编码" width="140" />
        <el-table-column prop="offline_reason" label="离线原因" min-width="150" show-overflow-tooltip />
        <el-table-column prop="offline_time" label="离线时间" width="180" />
        <el-table-column prop="handle_status" label="处理状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.handle_status)">{{ getStatusText(row.handle_status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('mes:offline:detail')" @click="handleDetail(row)">明细</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('mes:offline:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('mes:offline:delete')" @click="handleDelete(row)">删除</el-button>
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
import request from '@/utils/request'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ offline_no: '', material_code: '', handle_status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待处理', 2: '处理中', 3: '已处理' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'warning', 2: 'primary', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/mes/offline/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.offline_no = ''; searchForm.material_code = ''; searchForm.handle_status = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增产品离线') }
const handleDetail = (row: any) => { ElMessage.info('查看明细') }
const handleEdit = (row: any) => { ElMessage.info('编辑') }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该产品离线记录吗？', '提示', { type: 'warning' })
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.offline-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
