<template>
  <div class="process-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="流程名称">
          <el-input v-model="searchForm.name" placeholder="请输入流程名称" clearable />
        </el-form-item>
        <el-form-item label="分类">
          <el-select v-model="searchForm.category" placeholder="请选择" clearable>
            <el-option label="审批流程" value="approval" />
            <el-option label="执行流程" value="execution" />
            <el-option label="记录流程" value="record" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" :value="0" />
            <el-option label="已发布" :value="1" />
            <el-option label="已禁用" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('bpm:model:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新建流程
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="name" label="流程名称" min-width="160" />
        <el-table-column prop="code" label="流程编码" width="140" />
        <el-table-column prop="category" label="分类" width="100">
          <template #default="{ row }">
            <span>{{ getCategoryText(row.category) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="version" label="版本" width="80" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="create_time" label="创建时间" width="160" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('bpm:model:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('bpm:model:publish') && row.status === 0" @click="handlePublish(row)">发布</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('bpm:model:disable') && row.status === 1" @click="handleDisable(row)">禁用</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('bpm:model:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新建/编辑流程对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="流程名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入流程名称" />
        </el-form-item>
        <el-form-item label="流程编码" prop="code">
          <el-input v-model="formData.code" placeholder="请输入流程编码" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="分类" prop="category">
          <el-select v-model="formData.category" placeholder="请选择分类">
            <el-option label="审批流程" value="approval" />
            <el-option label="执行流程" value="execution" />
            <el-option label="记录流程" value="record" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="formData.description" type="textarea" :rows="3" placeholder="请输入描述" />
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
import { getProcessModelList, createProcessModel, updateProcessModel, deleteProcessModel, publishProcessModel } from '@/api/bpm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ name: '', category: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, name: '', code: '', category: '', description: ''
})

const rules: FormRules = {
  name: [{ required: true, message: '请输入流程名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入流程编码', trigger: 'blur' }],
  category: [{ required: true, message: '请选择分类', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑流程' : '新建流程')

const getCategoryText = (val: string) => {
  const map: Record<string, string> = { approval: '审批流程', execution: '执行流程', record: '记录流程' }
  return map[val] || val
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 0: '草稿', 1: '已发布', 2: '已禁用' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 0: 'info', 1: 'success', 2: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getProcessModelList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.name = ''; searchForm.category = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, name: '', code: '', category: '', description: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该流程吗？', '提示', { type: 'warning' })
    await deleteProcessModel(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handlePublish = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定发布该流程吗？', '提示', { type: 'warning' })
    await publishProcessModel(row.id)
    ElMessage.success('发布成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleDisable = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定禁用该流程吗？', '提示', { type: 'warning' })
    await updateProcessModel(row.id, { ...row, status: 2 })
    ElMessage.success('禁用成功')
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
    formData.id
      ? await updateProcessModel(formData.id, formData)
      : await createProcessModel(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.process-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
