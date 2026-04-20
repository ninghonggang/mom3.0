<template>
  <div class="asn-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="ASN编号">
          <el-input v-model="searchForm.asn_no" placeholder="请输入ASN编号" clearable />
        </el-form-item>
        <el-form-item label="供应商">
          <el-input v-model="searchForm.supplier_name" placeholder="请输入供应商" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待发货" value="PENDING" />
            <el-option label="已发货" value="SHIPPED" />
            <el-option label="已到货" value="ARRIVED" />
            <el-option label="已入库" value="WAREHOUSED" />
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
      <el-button type="primary" v-if="hasPermission('supplier:asn:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新建ASN
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="asn_no" label="ASN编号" width="150" />
        <el-table-column prop="supplier_name" label="供应商" min-width="150" />
        <el-table-column prop="expected_date" label="预计到货日期" width="120" />
        <el-table-column prop="actual_date" label="实际到货日期" width="120" />
        <el-table-column prop="total_qty" label="物料数量" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
            <el-button link type="success" v-if="row.status === 'SHIPPED' && hasPermission('supplier:asn:receive')" @click="handleReceive(row)">到货</el-button>
            <el-button link type="warning" v-if="row.status === 'ARRIVED' && hasPermission('supplier:asn:confirm')" @click="handleConfirm(row)">确认</el-button>
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
  asn_no: '',
  supplier_name: '',
  status: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    'PENDING': '待发货',
    'SHIPPED': '已发货',
    'ARRIVED': '已到货',
    'WAREHOUSED': '已入库',
    'CANCELLED': '已取消'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    'PENDING': 'info',
    'SHIPPED': 'primary',
    'ARRIVED': 'warning',
    'WAREHOUSED': 'success',
    'CANCELLED': 'danger'
  }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.asn_no) params.asn_no = searchForm.asn_no
    if (searchForm.supplier_name) params.supplier_name = searchForm.supplier_name
    if (searchForm.status) params.status = searchForm.status

    const res = await request.get('/supplier/asn/list', { params })
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
  searchForm.asn_no = ''
  searchForm.supplier_name = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  ElMessage.info('新建ASN功能开发中')
}

const handleDetail = (row: any) => {
  ElMessage.info('详情功能开发中')
}

const handleReceive = (row: any) => {
  ElMessage.info('到货确认功能开发中')
}

const handleConfirm = (row: any) => {
  ElMessage.info('确认入库功能开发中')
}

onMounted(() => {
  loadData()
})
</script>
