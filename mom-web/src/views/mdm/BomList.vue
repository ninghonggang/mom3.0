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
      <el-button type="primary" v-if="hasPermission('mdm:bom:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增BOM
      </el-button>
      <el-button type="success" v-if="hasPermission('mdm:bom:import')" @click="handleImport">
        <el-icon><Upload /></el-icon>导入BOM
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
            <el-button link type="primary" v-if="hasPermission('mdm:bom:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" @click="handleViewItems(row)">明细</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:bom:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 导入对话框 -->
    <el-dialog v-model="importDialogVisible" title="导入BOM" width="600px">
      <el-form ref="importFormRef" :model="importForm" label-width="100px">
        <el-form-item label="选择文件">
          <el-upload
            ref="uploadRef"
            :auto-upload="false"
            :limit="1"
            :on-change="handleFileChange"
            :file-list="fileList"
            accept=".xlsx,.xls"
          >
            <el-button type="primary">选择Excel文件</el-button>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="importDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="importLoading" @click="handleDoImport">开始导入</el-button>
      </template>
    </el-dialog>

    <!-- 导入结果对话框 -->
    <el-dialog v-model="resultDialogVisible" title="导入结果" width="600px">
      <el-alert v-if="importResult?.status === 'SUCCESS'" type="success" :title="`导入成功: ${importResult.success_rows}/${importResult.total_rows} 条`" :closable="false" />
      <el-alert v-else type="warning" :title="`导入完成: 成功${importResult?.success_rows || 0}条, 失败${importResult?.fail_rows || 0}条`" :closable="false" />
      <el-table v-if="failData.length > 0" :data="failData" max-height="300" style="margin-top: 16px">
        <el-table-column prop="bom_code" label="BOM编码" />
        <el-table-column prop="bom_name" label="BOM名称" />
        <el-table-column prop="error" label="失败原因" />
      </el-table>
      <template #footer>
        <el-button @click="resultDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getBOMList, deleteBOM, importBOMs, downloadBOMTemplate, getImportTaskResult } from '@/api/mdm'
import BomItemEditor from './BomItemEditor.vue'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const editorVisible = ref(false)
const currentBomId = ref<number | null>(null)
const editorMode = ref<'create' | 'edit'>('create')

// 导入相关
const importDialogVisible = ref(false)
const resultDialogVisible = ref(false)
const importLoading = ref(false)
const importFormRef = ref()
const importForm = ref<Record<string, any>>({})
const uploadRef = ref()
const fileList = ref<any[]>([])
const importFile = ref<File | null>(null)
const importResult = ref<any>(null)
const failData = ref<any[]>([])
const currentTaskId = ref<number | null>(null)

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
  try {
    await ElMessageBox.confirm('确定删除该BOM吗？', '提示', { type: 'warning' })
    await deleteBOM(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

// 导入相关方法
const handleImport = () => {
  importFile.value = null
  fileList.value = []
  importResult.value = null
  failData.value = []
  importDialogVisible.value = true
}

const handleFileChange = (file: any) => {
  importFile.value = file.raw
}

const handleDoImport = async () => {
  if (!importFile.value) {
    ElMessage.warning('请选择要导入的文件')
    return
  }
  importLoading.value = true
  try {
    const fd = new FormData()
    fd.append('file', importFile.value)
    const res = await importBOMs(fd)
    currentTaskId.value = res.data.task_id
    ElMessage.success(res.data.message || '导入任务已创建')
    importDialogVisible.value = false
    await pollImportResult(currentTaskId.value)
  } finally {
    importLoading.value = false
  }
}

const pollImportResult = async (taskId: number) => {
  const maxAttempts = 30
  let attempts = 0
  const poll = async () => {
    if (attempts >= maxAttempts) {
      ElMessage.warning('导入时间过长，请稍后手动查询结果')
      return
    }
    try {
      const res = await getImportTaskResult(taskId)
      const task = res.data
      if (task.status === 'PROCESSING') {
        attempts++
        setTimeout(poll, 1000)
      } else {
        importResult.value = task
        if (task.fail_data) {
          try {
            failData.value = JSON.parse(task.fail_data)
          } catch { failData.value = [] }
        }
        resultDialogVisible.value = true
        loadData()
      }
    } catch { ElMessage.error('查询导入结果失败') }
  }
  await poll()
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
