<template>
  <div class="attachment-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="文件名称">
          <el-input v-model="searchForm.file_name" placeholder="请输入文件名称" clearable />
        </el-form-item>
        <el-form-item label="所属类型">
          <el-select v-model="searchForm.owner_type" placeholder="请选择" clearable>
            <el-option label="客户" value="CUSTOMER" />
            <el-option label="供应商" value="SUPPLIER" />
            <el-option label="产品" value="PRODUCT" />
            <el-option label="订单" value="ORDER" />
          </el-select>
        </el-form-item>
        <el-form-item label="附件分类">
          <el-select v-model="searchForm.category" placeholder="请选择" clearable>
            <el-option label="证照" value="LICENSE" />
            <el-option label="证书" value="CERTIFICATE" />
            <el-option label="合同" value="CONTRACT" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:attachment:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="file_name" label="文件名称" min-width="200" show-overflow-tooltip />
        <el-table-column prop="file_type" label="文件类型" width="100" />
        <el-table-column prop="file_size" label="文件大小" width="100">
          <template #default="{ row }">
            {{ formatFileSize(row.file_size) }}
          </template>
        </el-table-column>
        <el-table-column prop="owner_type" label="所属类型" width="100">
          <template #default="{ row }">
            {{ getOwnerTypeName(row.owner_type) }}
          </template>
        </el-table-column>
        <el-table-column prop="category" label="分类" width="100">
          <template #default="{ row }">
            {{ getCategoryName(row.category) }}
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" min-width="150" show-overflow-tooltip />
        <el-table-column prop="created_by" label="上传人" width="100" />
        <el-table-column prop="created_at" label="上传时间" width="160" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDownload(row)">下载</el-button>
            <el-button link type="primary" v-if="hasPermission('mdm:attachment:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:attachment:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" />
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="所属类型" prop="owner_type">
          <el-select v-model="formData.owner_type" placeholder="请选择">
            <el-option label="客户" value="CUSTOMER" />
            <el-option label="供应商" value="SUPPLIER" />
            <el-option label="产品" value="PRODUCT" />
            <el-option label="订单" value="ORDER" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属ID" prop="owner_id">
          <el-input-number v-model="formData.owner_id" :min="1" />
        </el-form-item>
        <el-form-item label="文件名称" prop="file_name"><el-input v-model="formData.file_name" /></el-form-item>
        <el-form-item label="文件路径" prop="file_path"><el-input v-model="formData.file_path" /></el-form-item>
        <el-form-item label="文件大小"><el-input-number v-model="formData.file_size" :min="0" /></el-form-item>
        <el-form-item label="文件类型"><el-input v-model="formData.file_type" /></el-form-item>
        <el-form-item label="附件分类">
          <el-select v-model="formData.category" placeholder="请选择">
            <el-option label="证照" value="LICENSE" />
            <el-option label="证书" value="CERTIFICATE" />
            <el-option label="合同" value="CONTRACT" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述"><el-input v-model="formData.description" type="textarea" /></el-form-item>
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
import { getAttachmentList, getAttachmentById, createAttachment, updateAttachment, deleteAttachment } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const searchForm = reactive({ file_name: '', owner_type: '', category: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, owner_type: '', owner_id: 0, file_name: '', file_path: '', file_size: 0, file_type: '', category: '', description: '' })
const rules: FormRules = { owner_type: [{ required: true, message: '请选择所属类型', trigger: 'change' }], owner_id: [{ required: true, message: '请输入所属ID', trigger: 'blur' }], file_name: [{ required: true, message: '请输入文件名称', trigger: 'blur' }], file_path: [{ required: true, message: '请输入文件路径', trigger: 'blur' }] }
const dialogTitle = computed(() => formData.id ? '编辑附件' : '新增附件')

const formatFileSize = (size: number) => {
  if (size < 1024) return size + ' B'
  if (size < 1024 * 1024) return (size / 1024).toFixed(2) + ' KB'
  if (size < 1024 * 1024 * 1024) return (size / (1024 * 1024)).toFixed(2) + ' MB'
  return (size / (1024 * 1024 * 1024)).toFixed(2) + ' GB'
}

const getOwnerTypeName = (type: string) => {
  const map: Record<string, string> = { CUSTOMER: '客户', SUPPLIER: '供应商', PRODUCT: '产品', ORDER: '订单' }
  return map[type] || type
}

const getCategoryName = (category: string) => {
  const map: Record<string, string> = { LICENSE: '证照', CERTIFICATE: '证书', CONTRACT: '合同', OTHER: '其他' }
  return map[category] || category
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getAttachmentList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.file_name = ''; searchForm.owner_type = ''; searchForm.category = ''; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, owner_type: '', owner_id: 0, file_name: '', file_path: '', file_size: 0, file_type: '', category: '', description: '' }); dialogVisible.value = true }
const handleEdit = async (row: any) => {
  try {
    const res = await getAttachmentById(row.id)
    Object.assign(formData, res.data)
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取附件详情失败')
  }
}
const handleDownload = (row: any) => {
  window.open(row.file_path, '_blank')
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该附件吗？', '提示', { type: 'warning' })
    await deleteAttachment(row.id)
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
    if (formData.id) {
      await updateAttachment(formData.id, formData)
    } else {
      await createAttachment(formData)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally {
    submitLoading.value = false
  }
}
onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.attachment-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
