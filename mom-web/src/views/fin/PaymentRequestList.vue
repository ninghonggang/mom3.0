<template>
  <div class="payment-request-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="申请单号">
          <el-input v-model="searchForm.request_no" placeholder="请输入申请单号" clearable />
        </el-form-item>
        <el-form-item label="申请类型">
          <el-select v-model="searchForm.request_type" placeholder="请选择" clearable>
            <el-option label="采购付款" value="PURCHASE" />
            <el-option label="费用报销" value="EXPENSE" />
            <el-option label="其他付款" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待审批" value="PENDING" />
            <el-option label="已审批" value="APPROVED" />
            <el-option label="已付款" value="PAID" />
            <el-option label="已驳回" value="REJECTED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('fin:payment:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新建申请
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="request_no" label="申请单号" width="150" />
        <el-table-column prop="request_type" label="申请类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.request_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="applicant" label="申请人" width="100" />
        <el-table-column prop="amount" label="申请金额" width="120">
          <template #default="{ row }">
            {{ formatCurrency(row.amount) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
            <el-button link type="success" v-if="row.status === 'PENDING' && hasPermission('fin:payment:approve')" @click="handleApprove(row)">审批</el-button>
            <el-button link type="danger" v-if="row.status === 'PENDING' && hasPermission('fin:payment:reject')" @click="handleReject(row)">驳回</el-button>
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
const searchForm = reactive({ request_no: '', request_type: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formatCurrency = (amount: number) => new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY' }).format(amount)
const getTypeText = (type: string) => ({ PURCHASE: '采购付款', EXPENSE: '费用报销', OTHER: '其他付款' })[type] || type
const getStatusText = (status: string) => ({ PENDING: '待审批', APPROVED: '已审批', PAID: '已付款', REJECTED: '已驳回' })[status] || status
const getStatusType = (status: string) => ({ PENDING: 'warning', APPROVED: 'primary', PAID: 'success', REJECTED: 'danger' })[status] || 'info'

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.request_no) params.request_no = searchForm.request_no
    if (searchForm.request_type) params.request_type = searchForm.request_type
    if (searchForm.status) params.status = searchForm.status
    const res = await request.get('/fin/payment-request/list', { params })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) { ElMessage.error(error.message || '加载数据失败') }
  finally { loading.value = false }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.request_no = ''; searchForm.request_type = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => ElMessage.info('新建申请功能开发中')
const handleDetail = () => ElMessage.info('详情功能开发中')
const handleApprove = () => ElMessage.info('审批功能开发中')
const handleReject = () => ElMessage.info('驳回功能开发中')

onMounted(() => { loadData() })
</script>
