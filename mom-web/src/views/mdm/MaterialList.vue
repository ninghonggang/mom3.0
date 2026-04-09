<template>
  <div class="material-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="物料编码">
          <el-input v-model="searchForm.material_code" placeholder="请输入物料编码" clearable />
        </el-form-item>
        <el-form-item label="物料名称">
          <el-input v-model="searchForm.material_name" placeholder="请输入物料名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:material:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('mdm:material:import')" @click="handleImport">
        <el-icon><Upload /></el-icon>批量导入
      </el-button>
      <el-button type="info" v-if="hasPermission('mdm:material:export')" @click="handleDownloadTemplate">
        <el-icon><Download /></el-icon>下载模板
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="material_type" label="物料类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.material_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="spec" label="规格" width="120" />
        <el-table-column prop="unit" label="单位" width="80" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('mdm:material:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:material:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="物料编码" prop="material_code">
          <el-input v-model="formData.material_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="物料名称" prop="material_name">
          <el-input v-model="formData.material_name" />
        </el-form-item>
        <el-form-item label="物料类型" prop="material_type">
          <el-select v-model="formData.material_type" placeholder="请选择">
            <el-option label="原材料" value="raw" />
            <el-option label="半成品" value="semi" />
            <el-option label="成品" value="finished" />
          </el-select>
        </el-form-item>
        <el-form-item label="规格">
          <el-input v-model="formData.spec" />
        </el-form-item>
        <el-form-item label="单位" prop="unit">
          <el-input v-model="formData.unit" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 导入对话框 -->
    <el-dialog v-model="importDialogVisible" title="批量导入物料" width="600px">
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
            <template #tip>
              <div class="el-upload__tip">支持 .xlsx 或 .xls 格式的Excel文件</div>
            </template>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="importDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="importLoading" @click="handleDoImport">开始导入</el-button>
      </template>
    </el-dialog>

    <!-- 导入结果对话框 -->
    <el-dialog v-model="resultDialogVisible" title="导入结果" width="700px">
      <div v-if="importResult" class="import-result">
        <el-row :gutter="20" class="result-summary">
          <el-col :span="6">
            <div class="result-item">
              <div class="result-label">总行数</div>
              <div class="result-value">{{ importResult.total_rows }}</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="result-item success">
              <div class="result-label">成功</div>
              <div class="result-value">{{ importResult.success_rows }}</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="result-item danger">
              <div class="result-label">失败</div>
              <div class="result-value">{{ importResult.fail_rows }}</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="result-item">
              <div class="result-label">状态</div>
              <div class="result-value">
                <el-tag :type="importResult.status === 'SUCCESS' ? 'success' : 'danger'">
                  {{ importResult.status === 'SUCCESS' ? '成功' : '部分失败' }}
                </el-tag>
              </div>
            </div>
          </el-col>
        </el-row>

        <el-divider v-if="importResult.fail_rows > 0" content-position="left">失败数据</el-divider>
        <el-table v-if="importResult.fail_rows > 0" :data="failData" border max-height="300">
          <el-table-column prop="row_num" label="行号" width="60" />
          <el-table-column prop="material_code" label="物料编码" width="120" />
          <el-table-column prop="material_name" label="物料名称" min-width="120" />
          <el-table-column prop="error" label="错误原因" min-width="150" />
        </el-table>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getMaterialList, createMaterial, updateMaterial, deleteMaterial, importMaterials, downloadMaterialTemplate, getImportTaskResult } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

// 导入相关
const importDialogVisible = ref(false)
const resultDialogVisible = ref(false)
const importLoading = ref(false)
const importFormRef = ref<FormInstance>()
const uploadRef = ref()
const fileList = ref<any[]>([])
const importFile = ref<File | null>(null)
const importResult = ref<any>(null)
const failData = ref<any[]>([])
const currentTaskId = ref<number | null>(null)

const searchForm = reactive({ material_code: '', material_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, material_code: '', material_name: '', material_type: '', spec: '', unit: '', status: 1 })

const importForm = reactive({})

const rules: FormRules = {
  material_code: [{ required: true, message: '请输入物料编码', trigger: 'blur' }],
  material_name: [{ required: true, message: '请输入物料名称', trigger: 'blur' }],
  material_type: [{ required: true, message: '请选择物料类型', trigger: 'change' }],
  unit: [{ required: true, message: '请输入单位', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑物料' : '新增物料')

const getTypeText = (type: string) => {
  const map: Record<string, string> = { raw: '原材料', semi: '半成品', finished: '成品' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMaterialList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.material_code = ''; searchForm.material_name = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, material_code: '', material_name: '', material_type: '', spec: '', unit: '', status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该物料吗？', '提示', { type: 'warning' })
    await deleteMaterial(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateMaterial(formData.id, formData) : await createMaterial(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

// 导入相关方法
const handleImport = () => {
  importFile.value = null
  fileList.value = []
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
    const formData = new FormData()
    formData.append('file', importFile.value)

    const res = await importMaterials(formData)
    currentTaskId.value = res.data.task_id
    ElMessage.success(res.data.message || '导入任务已创建')

    importDialogVisible.value = false

    // 轮询查询导入结果
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
          } catch {
            failData.value = []
          }
        }
        resultDialogVisible.value = true
        loadData() // 刷新列表
      }
    } catch (error) {
      ElMessage.error('查询导入结果失败')
    }
  }

  await poll()
}

const handleDownloadTemplate = async () => {
  try {
    const res = await downloadMaterialTemplate()
    const blob = new Blob([res as unknown as BlobPart], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = '物料导入模板.xlsx'
    link.click()
    window.URL.revokeObjectURL(url)
    ElMessage.success('模板下载成功')
  } catch (error) {
    ElMessage.error('模板下载失败')
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.material-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}

.import-result {
  .result-summary {
    margin-bottom: 20px;
  }
  .result-item {
    text-align: center;
    padding: 15px;
    background: #f5f7fa;
    border-radius: 4px;
    .result-label {
      font-size: 14px;
      color: #909399;
      margin-bottom: 8px;
    }
    .result-value {
      font-size: 24px;
      font-weight: bold;
    }
    &.success .result-value { color: #67c23a; }
    &.danger .result-value { color: #f56c6c; }
  }
}
</style>
