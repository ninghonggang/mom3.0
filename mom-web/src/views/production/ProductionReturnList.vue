<template>
  <div class="production-return-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="退料单号">
          <el-input v-model="searchForm.return_no" placeholder="请输入退料单号" clearable />
        </el-form-item>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.work_order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待退料" value="PENDING" />
            <el-option label="已退料" value="RETURNED" />
            <el-option label="已取消" value="CANCELLED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:return:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新建退料
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="return_no" label="退料单号" width="150" />
        <el-table-column prop="work_order_no" label="工单号" width="150" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="qty" label="退料数量" width="100" />
        <el-table-column prop="reason" label="退料原因" min-width="150" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
            <el-button link type="success" v-if="hasPermission('production:return:confirm')" @click="handleConfirm(row)">确认退料</el-button>
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
const searchForm = reactive({ return_no: '', work_order_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: string) => ({ PENDING: '待退料', RETURNED: '已退料', CANCELLED: '已取消' })[status] || status
const getStatusType = (status: string) => ({ PENDING: 'warning', RETURNED: 'success', CANCELLED: 'danger' })[status] || 'info'

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.return_no) params.return_no = searchForm.return_no
    if (searchForm.work_order_no) params.work_order_no = searchForm.work_order_no
    if (searchForm.status) params.status = searchForm.status
    const res = await request.get('/production/return/list', { params })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) { ElMessage.error(error.message || '加载数据失败') }
  finally { loading.value = false }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.return_no = ''; searchForm.work_order_no = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => ElMessage.info('新建退料功能开发中')
const handleDetail = () => ElMessage.info('详情功能开发中')
const handleConfirm = () => ElMessage.info('确认退料功能开发中')

onMounted(() => { loadData() })
</script>
