<template>
  <div class="andon-call">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="呼叫编号">
          <el-input v-model="searchForm.call_no" placeholder="请输入呼叫编号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待响应" :value="1" />
            <el-option label="响应中" :value="2" />
            <el-option label="已解决" :value="3" />
            <el-option label="已关闭" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="call_no" label="呼叫编号" width="140" />
        <el-table-column prop="call_type" label="呼叫类型" width="100">
          <template #default="{ row }">
            <el-tag :type="getCallTypeTag(row.call_type)">{{ getCallTypeText(row.call_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="equipment_code" label="设备编号" width="100" />
        <el-table-column prop="equipment_name" label="设备名称" min-width="120" />
        <el-table-column prop="station_name" label="工位" width="100" />
        <el-table-column prop="call_user" label="呼叫人" width="80" />
        <el-table-column prop="call_time" label="呼叫时间" width="160" />
        <el-table-column prop="response_user" label="响应人" width="80" />
        <el-table-column prop="resolve_time" label="解决时间" width="160" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleResponse(row)" v-if="row.status === 1">响应</el-button>
            <el-button link type="success" size="small" @click="handleResolve(row)" v-if="row.status === 2">解决</el-button>
            <el-button link type="info" size="small" @click="handleDetail(row)">详情</el-button>
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
import { getAndonCallList, responseAndonCall, resolveAndonCall } from '@/api/trace'

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ call_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getCallTypeText = (type: string) => {
  const map: Record<string, string> = { '设备故障': '设备故障', '物料缺料': '物料缺料', '品质异常': '品质异常', '安全警示': '安全警示', '其他': '其他' }
  return map[type] || type
}

const getCallTypeTag = (type: string) => {
  const map: Record<string, string> = { '设备故障': 'danger', '物料缺料': 'warning', '品质异常': 'danger', '安全警示': 'warning', '其他': 'info' }
  return map[type] || 'info'
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待响应', 2: '响应中', 3: '已解决', 4: '已关闭' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'danger', 2: 'warning', 3: 'success', 4: 'info' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getAndonCallList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.call_no = ''; searchForm.status = ''; handleSearch() }
const handleResponse = async (row: any) => {
  await responseAndonCall(row.id)
  ElMessage.success('已响应')
  loadData()
}
const handleResolve = async (row: any) => {
  await resolveAndonCall(row.id)
  ElMessage.success('已解决')
  loadData()
}
const handleDetail = (row: any) => { ElMessage.info('查看详情') }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.andon-call {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
