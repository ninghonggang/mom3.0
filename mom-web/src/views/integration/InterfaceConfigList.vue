<template>
  <div class="interface-config-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="配置名称">
          <el-input v-model="searchForm.config_name" placeholder="请输入配置名称" clearable />
        </el-form-item>
        <el-form-item label="系统类型">
          <el-select v-model="searchForm.system_type" placeholder="请选择" clearable>
            <el-option label="ERP" value="erp" />
            <el-option label="WMS" value="wms" />
            <el-option label="MES" value="mes" />
            <el-option label="AGV" value="agv" />
            <el-option label="其他" value="other" />
          </el-select>
        </el-form-item>
        <el-form-item label="认证方式">
          <el-select v-model="searchForm.auth_type" placeholder="请选择" clearable>
            <el-option label="无认证" value="none" />
            <el-option label="Basic" value="basic" />
            <el-option label="Bearer Token" value="bearer" />
            <el-option label="API Key" value="apikey" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否启用">
          <el-select v-model="searchForm.is_enabled" placeholder="请选择" clearable>
            <el-option label="启用" :value="true" />
            <el-option label="禁用" :value="false" />
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
        <el-table-column prop="config_name" label="配置名称" min-width="150" />
        <el-table-column prop="system_type" label="系统类型" width="100">
          <template #default="{ row }">
            <el-tag type="info">{{ (row.system_type || '').toUpperCase() }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="endpoint" label="接口地址" min-width="200" show-overflow-tooltip />
        <el-table-column prop="auth_type" label="认证方式" width="120">
          <template #default="{ row }">
            {{ getAuthTypeText(row.auth_type) }}
          </template>
        </el-table-column>
        <el-table-column prop="is_enabled" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="row.is_enabled ? 'success' : 'info'">
              {{ row.is_enabled ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              size="small"
              v-if="hasPermission('integration:interface-config:detail')"
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
  config_name: '',
  system_type: '',
  auth_type: '',
  is_enabled: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getAuthTypeText = (type: string) => {
  const map: Record<string, string> = {
    none: '无认证',
    basic: 'Basic',
    bearer: 'Bearer Token',
    apikey: 'API Key'
  }
  return map[type] || type || '-'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/integration/interface-config/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch {
    ElMessage.error('加载接口配置列表失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.config_name = ''
  searchForm.system_type = ''
  searchForm.auth_type = ''
  searchForm.is_enabled = ''
  handleSearch()
}
const handleDetail = (row: any) => { ElMessage.info(`配置名称：${row.config_name}`) }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.interface-config-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
