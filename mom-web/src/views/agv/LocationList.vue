<template>
  <div class="agv-location-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="库位编码">
          <el-input v-model="searchForm.location_code" placeholder="请输入库位编码" clearable />
        </el-form-item>
        <el-form-item label="仓库编码">
          <el-input v-model="searchForm.warehouse_code" placeholder="请输入仓库编码" clearable />
        </el-form-item>
        <el-form-item label="库位类型">
          <el-select v-model="searchForm.location_type" placeholder="请选择" clearable>
            <el-option label="存储位" value="storage" />
            <el-option label="拣货位" value="pick" />
            <el-option label="暂存位" value="buffer" />
            <el-option label="充电位" value="charge" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否可用">
          <el-select v-model="searchForm.is_available" placeholder="请选择" clearable>
            <el-option label="可用" :value="true" />
            <el-option label="不可用" :value="false" />
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
        <el-table-column prop="location_code" label="库位编码" width="150" />
        <el-table-column prop="warehouse_code" label="仓库编码" width="130" />
        <el-table-column prop="location_type" label="库位类型" width="110">
          <template #default="{ row }">
            <el-tag type="info">{{ getLocationTypeText(row.location_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="is_available" label="是否可用" width="100">
          <template #default="{ row }">
            <el-tag :type="row.is_available ? 'success' : 'danger'">
              {{ row.is_available ? '可用' : '不可用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              size="small"
              v-if="hasPermission('agv:location:detail')"
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
  location_code: '',
  warehouse_code: '',
  location_type: '',
  is_available: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getLocationTypeText = (type: string) => {
  const map: Record<string, string> = {
    storage: '存储位',
    pick: '拣货位',
    buffer: '暂存位',
    charge: '充电位'
  }
  return map[type] || type || '未知'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/agv/location/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch {
    ElMessage.error('加载库位映射列表失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.location_code = ''
  searchForm.warehouse_code = ''
  searchForm.location_type = ''
  searchForm.is_available = ''
  handleSearch()
}
const handleDetail = (row: any) => { ElMessage.info(`库位编码：${row.location_code}`) }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.agv-location-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
