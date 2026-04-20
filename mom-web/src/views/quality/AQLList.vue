<template>
  <div class="aql-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="AQL方案名称">
          <el-input v-model="searchForm.name" placeholder="请输入方案名称" clearable />
        </el-form-item>
        <el-form-item label="检验级别">
          <el-select v-model="searchForm.inspection_level" placeholder="请选择" clearable>
            <el-option label="I级" value="I" />
            <el-option label="II级" value="II" />
            <el-option label="III级" value="III" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:aql:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="name" label="AQL方案名称" min-width="150" />
        <el-table-column prop="inspection_level" label="检验级别" width="100" />
        <el-table-column prop="aql_value" label="AQL值" width="100" />
        <el-table-column prop="sample_size" label="样本数" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('quality:aql:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('quality:aql:delete')" @click="handleDelete(row)">删除</el-button>
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
import request from '@/utils/request'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({
  name: '',
  inspection_level: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.name) params.name = searchForm.name
    if (searchForm.inspection_level) params.inspection_level = searchForm.inspection_level

    const res = await request.get('/quality/aql/list', { params })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.name = ''
  searchForm.inspection_level = ''
  handleSearch()
}

const handleAdd = () => {
  ElMessage.info('新增功能开发中')
}

const handleEdit = (row: any) => {
  ElMessage.info('编辑功能开发中')
}

const handleDelete = (row: any) => {
  ElMessage.info('删除功能开发中')
}

onMounted(() => {
  loadData()
})
</script>
