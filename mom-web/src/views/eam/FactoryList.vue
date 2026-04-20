<template>
  <div class="factory-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="厂区名称">
          <el-input v-model="searchForm.name" placeholder="请输入厂区名称" clearable />
        </el-form-item>
        <el-form-item label="厂区编码">
          <el-input v-model="searchForm.code" placeholder="请输入厂区编码" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('eam:factory:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="factory_code" label="厂区编码" width="120" />
        <el-table-column prop="factory_name" label="厂区名称" min-width="150" />
        <el-table-column prop="location" label="位置" min-width="150" />
        <el-table-column prop="area" label="面积(m²)" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('eam:factory:delete')" @click="handleDelete(row)">删除</el-button>
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
  code: ''
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
    if (searchForm.code) params.code = searchForm.code

    const res = await request.get('/eam/factory/list', { params })
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
  searchForm.code = ''
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
