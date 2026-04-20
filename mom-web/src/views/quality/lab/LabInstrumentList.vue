<template>
  <div class="lab-instrument-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="仪器名称">
          <el-input v-model="searchForm.name" placeholder="请输入仪器名称" clearable />
        </el-form-item>
        <el-form-item label="仪器编号">
          <el-input v-model="searchForm.code" placeholder="请输入仪器编号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="正常" :value="1" />
            <el-option label="校准中" :value="2" />
            <el-option label="维修中" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:lab:instrument:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="instrument_code" label="仪器编号" width="120" />
        <el-table-column prop="instrument_name" label="仪器名称" min-width="150" />
        <el-table-column prop="model" label="型号" width="120" />
        <el-table-column prop="calibration_date" label="最近校准日期" width="130" />
        <el-table-column prop="next_calibration_date" label="下次校准日期" width="130" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
            <el-button link type="warning" v-if="hasPermission('quality:lab:instrument:calibrate')" @click="handleCalibrate(row)">校准</el-button>
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
  code: '',
  status: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '正常', 2: '校准中', 3: '维修中' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'success', 2: 'warning', 3: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.name) params.name = searchForm.name
    if (searchForm.code) params.code = searchForm.code
    if (searchForm.status) params.status = searchForm.status

    const res = await request.get('/quality/lab-instrument/list', { params })
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
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  ElMessage.info('新增功能开发中')
}

const handleDetail = (row: any) => {
  ElMessage.info('详情功能开发中')
}

const handleCalibrate = (row: any) => {
  ElMessage.info('校准功能开发中')
}

onMounted(() => {
  loadData()
})
</script>
