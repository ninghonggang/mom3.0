<template>
  <div class="iqc-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="检验单号">
          <el-input v-model="searchForm.iqc_no" placeholder="请输入检验单号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待检验" :value="1" />
            <el-option label="检验中" :value="2" />
            <el-option label="合格" :value="3" />
            <el-option label="不合格" :value="4" />
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
        <el-table-column prop="iqc_no" label="检验单号" width="140" />
        <el-table-column prop="material_code" label="物料编码" width="100" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="supplier_name" label="供应商" min-width="120" />
        <el-table-column prop="quantity" label="送检数量" width="90" />
        <el-table-column prop="qualified_qty" label="合格数量" width="90" />
        <el-table-column prop="unqualified_qty" label="不合格数" width="90" />
        <el-table-column prop="check_user" label="检验人" width="80" />
        <el-table-column prop="check_date" label="检验日期" width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
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
import { getIQCList } from '@/api/quality'

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ iqc_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待检验', 2: '检验中', 3: '合格', 4: '不合格' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success', 4: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getIQCList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.iqc_no = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增功能') }
const handleEdit = (row: any) => { ElMessage.info('编辑功能') }
const handleDelete = (row: any) => { ElMessage.info('删除功能') }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.iqc-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
