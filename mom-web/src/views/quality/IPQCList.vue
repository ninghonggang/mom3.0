<template>
  <div class="ipqc-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="检验单号">
          <el-input v-model="searchForm.ipqc_no" placeholder="请输入检验单号" clearable />
        </el-form-item>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.result" placeholder="请选择" clearable>
            <el-option label="待检验" :value="1" />
            <el-option label="合格" :value="2" />
            <el-option label="不合格" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="ipqc_no" label="检验单号" width="140" />
        <el-table-column prop="order_no" label="工单号" width="120" />
        <el-table-column prop="process_name" label="工序" min-width="100" />
        <el-table-column prop="quantity" label="数量" width="80" />
        <el-table-column prop="sample_size" label="抽样数" width="70" />
        <el-table-column prop="check_user_name" label="检验人" width="80" />
        <el-table-column prop="check_date" label="检验日期" width="120" />
        <el-table-column prop="result" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.result)">{{ getStatusText(row.result) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" @click="handleDelete(row)">删除</el-button>
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
import { getIPQCList } from '@/api/quality'

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ ipqc_no: '', order_no: '', result: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待检验', 2: '合格', 3: '不合格' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'success', 3: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getIPQCList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.ipqc_no = ''; searchForm.order_no = ''; searchForm.result = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增功能') }
const handleEdit = (row: any) => { ElMessage.info('编辑功能') }
const handleDelete = (row: any) => { ElMessage.info('删除功能') }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.ipqc-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
