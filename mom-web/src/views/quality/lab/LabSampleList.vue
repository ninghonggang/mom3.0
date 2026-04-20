<template>
  <div class="lab-sample-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="样品编号">
          <el-input v-model="searchForm.sample_no" placeholder="请输入样品编号" clearable />
        </el-form-item>
        <el-form-item label="检验类型">
          <el-select v-model="searchForm.inspection_type" placeholder="请选择" clearable>
            <el-option label="来料检验" value="IQC" />
            <el-option label="过程检验" value="IPQC" />
            <el-option label="成品检验" value="FQC" />
            <el-option label="出货检验" value="OQC" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:lab:sample:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="sample_no" label="样品编号" width="150" />
        <el-table-column prop="product_name" label="产品名称" min-width="150" />
        <el-table-column prop="inspection_type" label="检验类型" width="100" />
        <el-table-column prop="quantity" label="样品数量" width="100" />
        <el-table-column prop="receive_time" label="接收时间" width="160" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
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
  sample_no: '',
  inspection_type: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待检验', 2: '检验中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'warning', 2: 'primary', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.sample_no) params.sample_no = searchForm.sample_no
    if (searchForm.inspection_type) params.inspection_type = searchForm.inspection_type

    const res = await request.get('/quality/lab/samples/list', { params })
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
  searchForm.sample_no = ''
  searchForm.inspection_type = ''
  handleSearch()
}

const handleAdd = () => {
  ElMessage.info('新增功能开发中')
}

const handleDetail = (row: any) => {
  ElMessage.info('详情功能开发中')
}

onMounted(() => {
  loadData()
})
</script>
