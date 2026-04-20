<template>
  <div class="agv-device-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="AGV编码">
          <el-input v-model="searchForm.agv_code" placeholder="请输入AGV编码" clearable />
        </el-form-item>
        <el-form-item label="AGV名称">
          <el-input v-model="searchForm.agv_name" placeholder="请输入AGV名称" clearable />
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="searchForm.agv_type" placeholder="请选择" clearable>
            <el-option label="潜伏式" value="latent" />
            <el-option label="叉车式" value="forklift" />
            <el-option label="牵引式" value="traction" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="空闲" :value="0" />
            <el-option label="运行中" :value="1" />
            <el-option label="充电中" :value="2" />
            <el-option label="故障" :value="3" />
            <el-option label="离线" :value="4" />
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
        <el-table-column prop="agv_code" label="AGV编码" width="130" />
        <el-table-column prop="agv_name" label="AGV名称" min-width="130" />
        <el-table-column prop="agv_type" label="类型" width="100">
          <template #default="{ row }">
            <el-tag type="info">{{ getAgvTypeText(row.agv_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="location" label="当前位置" min-width="130" />
        <el-table-column prop="last_heartbeat" label="最后心跳" width="180" />
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              size="small"
              v-if="hasPermission('agv:device:detail')"
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
  agv_code: '',
  agv_name: '',
  agv_type: '',
  status: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getAgvTypeText = (type: string) => {
  const map: Record<string, string> = {
    latent: '潜伏式',
    forklift: '叉车式',
    traction: '牵引式'
  }
  return map[type] || type || '未知'
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 0: '空闲', 1: '运行中', 2: '充电中', 3: '故障', 4: '离线' }
  return map[status] ?? '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = {
    0: 'success',
    1: 'primary',
    2: 'warning',
    3: 'danger',
    4: 'info'
  }
  return map[status] ?? 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/agv/device/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch {
    ElMessage.error('加载AGV设备列表失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.agv_code = ''
  searchForm.agv_name = ''
  searchForm.agv_type = ''
  searchForm.status = ''
  handleSearch()
}
const handleDetail = (row: any) => { ElMessage.info(`AGV编码：${row.agv_code}`) }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.agv-device-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
