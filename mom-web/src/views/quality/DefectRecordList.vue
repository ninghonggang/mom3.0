<template>
  <div class="defect-record-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="记录单号">
          <el-input v-model="searchForm.record_no" placeholder="请输入单号" clearable />
        </el-form-item>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待处理" :value="1" />
            <el-option label="处理中" :value="2" />
            <el-option label="已处理" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:defectrecord:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="record_no" label="记录单号" width="140" />
        <el-table-column prop="order_no" label="工单号" width="120" />
        <el-table-column prop="process_name" label="工序" width="100" />
        <el-table-column prop="defect_code" label="不良代码" width="100" />
        <el-table-column prop="defect_name" label="不良名称" min-width="120" />
        <el-table-column prop="quantity" label="数量" width="70" />
        <el-table-column prop="handle_method" label="处理方式" width="90">
          <template #default="{ row }">
            <el-tag>{{ getHandleMethodText(row.handle_method) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('quality:defectrecord:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('quality:defectrecord:process')" @click="handleProcess(row)">处理</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('quality:defectrecord:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新增/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="记录单号" prop="record_no">
          <el-input v-model="formData.record_no" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="工单号" prop="order_no">
          <el-input v-model="formData.order_no" />
        </el-form-item>
        <el-form-item label="工序">
          <el-input v-model="formData.process_name" />
        </el-form-item>
        <el-form-item label="不良代码" prop="defect_code">
          <el-input v-model="formData.defect_code" />
        </el-form-item>
        <el-form-item label="不良名称">
          <el-input v-model="formData.defect_name" />
        </el-form-item>
        <el-form-item label="数量" prop="quantity">
          <el-input-number v-model="formData.quantity" :min="0" />
        </el-form-item>
        <el-form-item label="处理方式">
          <el-select v-model="formData.handle_method" placeholder="请选择">
            <el-option label="返工" :value="1" />
            <el-option label="返修" :value="2" />
            <el-option label="报废" :value="3" />
            <el-option label="特采" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option label="待处理" :value="1" />
            <el-option label="处理中" :value="2" />
            <el-option label="已处理" :value="3" />
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

    <!-- 处理对话框 -->
    <el-dialog v-model="processVisible" title="处理不良品" width="500px">
      <el-form ref="processFormRef" :model="processData" :rules="processRules" label-width="100px">
        <el-form-item label="处理方式" prop="handle_method">
          <el-select v-model="processData.handle_method" placeholder="请选择">
            <el-option label="返工" :value="1" />
            <el-option label="返修" :value="2" />
            <el-option label="报废" :value="3" />
            <el-option label="特采" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="处理说明">
          <el-input v-model="processData.handle_remark" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="processVisible = false">取消</el-button>
        <el-button type="primary" :loading="processLoading" @click="handleDoProcess">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getDefectRecordList, createDefectRecord, updateDefectRecord, handleDefect, deleteDefectRecord } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const processVisible = ref(false)
const submitLoading = ref(false)
const processLoading = ref(false)
const formRef = ref<FormInstance>()
const processFormRef = ref<FormInstance>()

const searchForm = reactive({ record_no: '', order_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive<any>({
  id: 0, record_no: '', order_no: '', process_name: '', defect_code: '', defect_name: '',
  quantity: 0, handle_method: 0, status: 1, remark: ''
})
const processData = reactive<any>({ id: 0, handle_method: 0, handle_remark: '' })

const rules: FormRules = {
  record_no: [{ required: true, message: '请输入记录单号', trigger: 'blur' }],
  order_no: [{ required: true, message: '请输入工单号', trigger: 'blur' }],
  defect_code: [{ required: true, message: '请输入不良代码', trigger: 'blur' }],
  quantity: [{ required: true, message: '请输入数量', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}

const processRules: FormRules = {
  handle_method: [{ required: true, message: '请选择处理方式', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑不良品记录' : '新增不良品记录')

const getHandleMethodText = (m: number) => ({ 1: '返工', 2: '返修', 3: '报废', 4: '特采' })[m] || '未知'
const getStatusText = (s: number) => ({ 1: '待处理', 2: '处理中', 3: '已处理' })[s] || '未知'
const getStatusType = (s: number) => ({ 1: 'info', 2: 'warning', 3: 'success' })[s] || 'info'

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDefectRecordList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally { loading.value = false }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.record_no = ''; searchForm.order_no = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => {
  Object.assign(formData, { id: 0, record_no: '', order_no: '', process_name: '', defect_code: '', defect_name: '', quantity: 0, handle_method: 0, status: 1, remark: '' })
  dialogVisible.value = true
}
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleProcess = (row: any) => {
  processData.id = row.id
  processData.handle_method = row.handle_method || 0
  processData.handle_remark = ''
  processVisible.value = true
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该记录吗？', '提示', { type: 'warning' })
    await deleteDefectRecord(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateDefectRecord(formData.id, formData) : await createDefectRecord(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}
const handleDoProcess = async () => {
  if (!processFormRef.value) return
  await processFormRef.value.validate()
  processLoading.value = true
  try {
    await handleDefect(processData.id, { handle_method: processData.handle_method, handle_remark: processData.handle_remark })
    ElMessage.success('处理成功')
    processVisible.value = false
    loadData()
  } finally { processLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.defect-record-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
