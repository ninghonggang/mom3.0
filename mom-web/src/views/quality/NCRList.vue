<template>
  <div class="ncr-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="NCR单号">
          <el-input v-model="searchForm.ncr_no" placeholder="请输入单号" clearable />
        </el-form-item>
        <el-form-item label="来源">
          <el-select v-model="searchForm.source_type" placeholder="请选择" clearable>
            <el-option label="IQC" value="IQC" />
            <el-option label="IPQC" value="IPQC" />
            <el-option label="FQC" value="FQC" />
            <el-option label="OQC" value="OQC" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待处理" :value="1" />
            <el-option label="处理中" :value="2" />
            <el-option label="已完成" :value="3" />
            <el-option label="已关闭" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:ncr:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="ncr_no" label="NCR单号" width="140" />
        <el-table-column prop="source_type" label="来源" width="80" />
        <el-table-column prop="issue_desc" label="问题描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="root_cause" label="根本原因" min-width="150" show-overflow-tooltip />
        <el-table-column prop="corrective_action" label="纠正措施" min-width="150" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('quality:ncr:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('quality:ncr:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="110px">
        <el-form-item label="NCR单号" prop="ncr_no">
          <el-input v-model="formData.ncr_no" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="来源" prop="source_type">
          <el-select v-model="formData.source_type" placeholder="请选择">
            <el-option label="IQC" value="IQC" />
            <el-option label="IPQC" value="IPQC" />
            <el-option label="FQC" value="FQC" />
            <el-option label="OQC" value="OQC" />
          </el-select>
        </el-form-item>
        <el-form-item label="问题描述" prop="issue_desc">
          <el-input v-model="formData.issue_desc" type="textarea" />
        </el-form-item>
        <el-form-item label="根本原因">
          <el-input v-model="formData.root_cause" type="textarea" />
        </el-form-item>
        <el-form-item label="纠正措施">
          <el-input v-model="formData.corrective_action" type="textarea" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option label="待处理" :value="1" />
            <el-option label="处理中" :value="2" />
            <el-option label="已完成" :value="3" />
            <el-option label="已关闭" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="处理人">
          <el-input v-model="formData.handle_user" />
        </el-form-item>
        <el-form-item label="处理日期">
          <el-date-picker v-model="formData.handle_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
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
import { getNCRList, createNCR, updateNCR, deleteNCR } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ ncr_no: '', source_type: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive<any>({
  id: 0, ncr_no: '', source_type: '', issue_desc: '', root_cause: '', corrective_action: '',
  status: 1, handle_user: '', handle_date: ''
})

const rules: FormRules = {
  ncr_no: [{ required: true, message: '请输入NCR单号', trigger: 'blur' }],
  source_type: [{ required: true, message: '请选择来源', trigger: 'change' }],
  issue_desc: [{ required: true, message: '请输入问题描述', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑NCR' : '新增NCR')
const getStatusText = (s: number) => ({ 1: '待处理', 2: '处理中', 3: '已完成', 4: '已关闭' })[s] || '未知'
const getStatusType = (s: number) => ({ 1: 'info', 2: 'warning', 3: 'success', 4: '' })[s] || 'info'

const loadData = async () => {
  loading.value = true
  try {
    const res = await getNCRList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally { loading.value = false }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.ncr_no = ''; searchForm.source_type = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => {
  Object.assign(formData, { id: 0, ncr_no: '', source_type: '', issue_desc: '', root_cause: '', corrective_action: '', status: 1, handle_user: '', handle_date: '' })
  dialogVisible.value = true
}
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该NCR记录吗？', '提示', { type: 'warning' })
    await deleteNCR(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateNCR(formData.id, formData) : await createNCR(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.ncr-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
