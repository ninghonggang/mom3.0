<template>
  <div class="downtime-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.equipment_name" placeholder="请输入设备名称" clearable />
        </el-form-item>
        <el-form-item label="停机类型">
          <el-select v-model="searchForm.downtime_type" placeholder="请选择" clearable>
            <el-option label="计划停机" value="PLANNED" />
            <el-option label="故障停机" value="BREAKDOWN" />
            <el-option label="换型停机" value="CHANGEOVER" />
            <el-option label="待料停机" value="MATERIAL_WAITING" />
          </el-select>
        </el-form-item>
        <el-form-item label="日期">
          <el-date-picker v-model="searchForm.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('eam:downtime:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="equipment_name" label="设备名称" min-width="150" />
        <el-table-column prop="downtime_type" label="停机类型" width="120">
          <template #default="{ row }">
            <el-tag>{{ getDowntimeTypeText(row.downtime_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="start_time" label="开始时间" width="160" />
        <el-table-column prop="end_time" label="结束时间" width="160" />
        <el-table-column prop="duration" label="持续时间(分钟)" width="130" />
        <el-table-column prop="reason" label="停机原因" min-width="150" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'warning'">{{ row.status === 1 ? '已完成' : '进行中' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
            <el-button link type="success" v-if="row.status === 0 && hasPermission('eam:downtime:end')" @click="handleEnd(row)">结束</el-button>
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
  equipment_name: '',
  downtime_type: '',
  dateRange: []
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getDowntimeTypeText = (type: string) => {
  const map: Record<string, string> = {
    'PLANNED': '计划停机',
    'BREAKDOWN': '故障停机',
    'CHANGEOVER': '换型停机',
    'MATERIAL_WAITING': '待料停机'
  }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.equipment_name) params.equipment_name = searchForm.equipment_name
    if (searchForm.downtime_type) params.downtime_type = searchForm.downtime_type
    if (searchForm.dateRange && searchForm.dateRange.length === 2) {
      params.start_date = searchForm.dateRange[0]
      params.end_date = searchForm.dateRange[1]
    }

    const res = await request.get('/eam/downtime/list', { params })
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
  searchForm.equipment_name = ''
  searchForm.downtime_type = ''
  searchForm.dateRange = []
  handleSearch()
}

const handleAdd = () => {
  ElMessage.info('新增功能开发中')
}

const handleDetail = (row: any) => {
  ElMessage.info('详情功能开发中')
}

const handleEnd = (row: any) => {
  ElMessage.info('结束停机功能开发中')
}

onMounted(() => {
  loadData()
})
</script>
