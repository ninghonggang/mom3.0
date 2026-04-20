<template>
  <div class="sales-settlement-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="结算单号">
          <el-input v-model="searchForm.settlement_no" placeholder="请输入结算单号" clearable />
        </el-form-item>
        <el-form-item label="客户">
          <el-input v-model="searchForm.customer_name" placeholder="请输入客户" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待审批" value="PENDING" />
            <el-option label="已审批" value="APPROVED" />
            <el-option label="已收款" value="RECEIVED" />
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
      <el-button type="primary" v-if="hasPermission('fin:sales:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新建结算
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="settlement_no" label="结算单号" width="150" />
        <el-table-column prop="customer_name" label="客户" min-width="150" />
        <el-table-column prop="so_no" label="销售单号" width="150" />
        <el-table-column prop="amount" label="结算金额" width="120">
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
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
            <el-button link type="success" v-if="row.status === 'PENDING' && hasPermission('fin:sales:approve')" @click="handleApprove(row)">审批</el-button>
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
const searchForm = reactive({ settlement_no: '', customer_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formatCurrency = (amount: number) => new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY' }).format(amount)

const getStatusText = (status: string) => ({ PENDING: '待审批', APPROVED: '已审批', RECEIVED: '已收款', CANCELLED: '已取消' })[status] || status
const getStatusType = (status: string) => ({ PENDING: 'warning', APPROVED: 'primary', RECEIVED: 'success', CANCELLED: 'danger' })[status] || 'info'

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.settlement_no) params.settlement_no = searchForm.settlement_no
    if (searchForm.customer_name) params.customer_name = searchForm.customer_name
    if (searchForm.status) params.status = searchForm.status
    const res = await request.get('/fin/sales-settlement/list', { params })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) { ElMessage.error(error.message || '加载数据失败') }
  finally { loading.value = false }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.settlement_no = ''; searchForm.customer_name = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => ElMessage.info('新建结算功能开发中')
const handleDetail = () => ElMessage.info('详情功能开发中')
const handleApprove = () => ElMessage.info('审批功能开发中')

onMounted(() => { loadData() })
</script>
