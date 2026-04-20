<template>
  <div class="process-route-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="路线名称">
          <el-input v-model="searchForm.route_name" placeholder="请输入路线名称" clearable />
        </el-form-item>
        <el-form-item label="物料编码">
          <el-input v-model="searchForm.material_code" placeholder="请输入物料编码" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mes:processRoute:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="route_name" label="路线名称" min-width="150" />
        <el-table-column prop="material_code" label="物料编码" width="140" />
        <el-table-column prop="version" label="版本" width="80" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('mes:processRoute:detail')" @click="handleDetail(row)">明细</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('mes:processRoute:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('mes:processRoute:delete')" @click="handleDelete(row)">删除</el-button>
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

const searchForm = reactive({ route_name: '', material_code: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/mes/process-routes/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.route_name = ''; searchForm.material_code = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增工艺路线') }
const handleDetail = (row: any) => { ElMessage.info('查看明细') }
const handleEdit = (row: any) => { ElMessage.info('编辑') }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该工艺路线吗？', '提示', { type: 'warning' })
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.process-route-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
