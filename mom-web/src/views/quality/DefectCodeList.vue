<template>
  <div class="defect-code-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="代码">
          <el-input v-model="searchForm.defect_code" placeholder="请输入代码" clearable />
        </el-form-item>
        <el-form-item label="名称">
          <el-input v-model="searchForm.defect_name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="searchForm.defect_type" placeholder="请选择" clearable>
            <el-option label="尺寸" value="尺寸" />
            <el-option label="外观" value="外观" />
            <el-option label="功能" value="功能" />
            <el-option label="其他" value="其他" />
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
        <el-table-column prop="defect_code" label="代码" width="100" />
        <el-table-column prop="defect_name" label="名称" min-width="150" />
        <el-table-column prop="defect_type" label="类型" width="80" />
        <el-table-column prop="severity" label="严重程度" width="100">
          <template #default="{ row }">
            <el-tag :type="getSeverityType(row.severity)">{{ getSeverityText(row.severity) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
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
import { getDefectCodeList, deleteDefectCode } from '@/api/quality'

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ defect_code: '', defect_name: '', defect_type: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getSeverityText = (severity: number) => {
  const map: Record<number, string> = { 1: '轻微', 2: '一般', 3: '严重' }
  return map[severity] || '未知'
}

const getSeverityType = (severity: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'danger' }
  return map[severity] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDefectCodeList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.defect_code = ''; searchForm.defect_name = ''; searchForm.defect_type = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增功能') }
const handleEdit = (row: any) => { ElMessage.info('编辑功能') }
const handleDelete = (row: any) => {
  ElMessageBox.confirm('确定删除该不良代码?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    await deleteDefectCode(row.id)
    ElMessage.success('删除成功')
    loadData()
  }).catch(() => {})
}

onMounted(() => { loadData() })
</script>

<script lang="ts">
import { ElMessageBox } from 'element-plus'
export default { name: 'DefectCodeList' }
</script>

<style scoped lang="scss">
.defect-code-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
