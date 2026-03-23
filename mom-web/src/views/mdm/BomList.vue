<template>
  <div class="bom-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="BOM编码">
          <el-input v-model="searchForm.bom_code" placeholder="请输入BOM编码" clearable />
        </el-form-item>
        <el-form-item label="BOM名称">
          <el-input v-model="searchForm.bom_name" placeholder="请输入BOM名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" value="DRAFT" />
            <el-option label="生效" value="ACTIVE" />
            <el-option label="失效" value="EXPIRED" />
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
        <el-icon><Plus /></el-icon>新增BOM
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" row-key="id" default-expand-all>
        <el-table-column prop="bom_code" label="BOM编码" width="150" />
        <el-table-column prop="bom_name" label="BOM名称" min-width="180" />
        <el-table-column prop="material_code" label="产品编码" width="120" />
        <el-table-column prop="material_name" label="产品名称" width="150" />
        <el-table-column prop="version" label="版本" width="80" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="eff_date" label="生效日期" width="120" />
        <el-table-column prop="exp_date" label="失效日期" width="120" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" @click="handleViewItems(row)">明细</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
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

    <!-- BOM编辑弹窗 -->
    <BomItemEditor
      v-model="editorVisible"
      :bom-id="currentBomId"
      :mode="editorMode"
      @refresh="loadData"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getBOMList, deleteBOM } from '@/api/mdm'
import BomItemEditor from './BomItemEditor.vue'

const loading = ref(false)
const tableData = ref<any[]>([])
const editorVisible = ref(false)
const currentBomId = ref<number | null>(null)
const editorMode = ref<'create' | 'edit'>('create')

const searchForm = reactive({ bom_code: '', bom_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusType = (status: string) => {
  const map: Record<string, string> = { DRAFT: 'info', ACTIVE: 'success', EXPIRED: 'warning' }
  return map[status] || 'info'
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = { DRAFT: '草稿', ACTIVE: '生效', EXPIRED: '失效' }
  return map[status] || status
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getBOMList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.bom_code = ''; searchForm.bom_name = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  currentBomId.value = null
  editorMode.value = 'create'
  editorVisible.value = true
}

const handleEdit = (row: any) => {
  currentBomId.value = row.id
  editorMode.value = 'edit'
  editorVisible.value = true
}

const handleViewItems = (row: any) => {
  currentBomId.value = row.id
  editorMode.value = 'edit'
  editorVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该BOM吗？', '提示', { type: 'warning' })
  await deleteBOM(row.id)
  ElMessage.success('删除成功')
  loadData()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.bom-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
