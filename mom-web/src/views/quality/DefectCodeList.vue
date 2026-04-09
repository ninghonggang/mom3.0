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
      <el-button type="primary" v-if="hasPermission('quality:defectcode:add')" @click="handleAdd">
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
            <el-button link type="primary" size="small" v-if="hasPermission('quality:defectcode:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('quality:defectcode:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="代码" prop="defect_code">
          <el-input v-model="formData.defect_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="名称" prop="defect_name">
          <el-input v-model="formData.defect_name" />
        </el-form-item>
        <el-form-item label="类型" prop="defect_type">
          <el-select v-model="formData.defect_type" placeholder="请选择">
            <el-option label="尺寸" value="尺寸" />
            <el-option label="外观" value="外观" />
            <el-option label="功能" value="功能" />
            <el-option label="其他" value="其他" />
          </el-select>
        </el-form-item>
        <el-form-item label="严重程度" prop="severity">
          <el-select v-model="formData.severity" placeholder="请选择">
            <el-option label="轻微" :value="1" />
            <el-option label="一般" :value="2" />
            <el-option label="严重" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getDefectCodeList, createDefectCode, updateDefectCode, deleteDefectCode } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ defect_code: '', defect_name: '', defect_type: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive<any>({ id: 0, defect_code: '', defect_name: '', defect_type: '', severity: 1, status: 1, remark: '' })

const rules: FormRules = {
  defect_code: [{ required: true, message: '请输入代码', trigger: 'blur' }],
  defect_name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
  defect_type: [{ required: true, message: '请选择类型', trigger: 'change' }],
  severity: [{ required: true, message: '请选择严重程度', trigger: 'change' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑不良代码' : '新增不良代码')
const getSeverityText = (s: number) => ({ 1: '轻微', 2: '一般', 3: '严重' })[s] || '未知'
const getSeverityType = (s: number) => ({ 1: 'info', 2: 'warning', 3: 'danger' })[s] || 'info'

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDefectCodeList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally { loading.value = false }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.defect_code = ''; searchForm.defect_name = ''; searchForm.defect_type = ''; handleSearch() }
const handleAdd = () => {
  Object.assign(formData, { id: 0, defect_code: '', defect_name: '', defect_type: '', severity: 1, status: 1, remark: '' })
  dialogVisible.value = true
}
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该不良代码吗？', '提示', { type: 'warning' })
    await deleteDefectCode(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateDefectCode(formData.id, formData) : await createDefectCode(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.defect-code-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
