<template>
  <div class="oqc-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="检验单号">
          <el-input v-model="searchForm.oqc_no" placeholder="请输入检验单号" clearable />
        </el-form-item>
        <el-form-item label="发货单号">
          <el-input v-model="searchForm.shipping_no" placeholder="请输入发货单号" clearable />
        </el-form-item>
        <el-form-item label="客户">
          <el-input v-model="searchForm.customer_name" placeholder="请输入客户" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:oqc:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="oqc_no" label="检验单号" width="140" />
        <el-table-column prop="shipping_no" label="发货单号" width="120" />
        <el-table-column prop="customer_name" label="客户" min-width="120" />
        <el-table-column prop="quantity" label="数量" width="80" />
        <el-table-column prop="check_user_name" label="检验人" width="80" />
        <el-table-column prop="check_date" label="检验日期" width="120" />
        <el-table-column prop="result" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.result)">{{ getStatusText(row.result) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('quality:oqc:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('quality:oqc:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="检验单号" prop="oqc_no">
          <el-input v-model="formData.oqc_no" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="发货单号" prop="shipping_no">
          <el-input v-model="formData.shipping_no" />
        </el-form-item>
        <el-form-item label="客户">
          <el-input v-model="formData.customer_name" />
        </el-form-item>
        <el-form-item label="数量" prop="quantity">
          <el-input-number v-model="formData.quantity" :min="0" />
        </el-form-item>
        <el-form-item label="合格数">
          <el-input-number v-model="formData.qualified_qty" :min="0" />
        </el-form-item>
        <el-form-item label="不合格数">
          <el-input-number v-model="formData.unqualified_qty" :min="0" />
        </el-form-item>
        <el-form-item label="检验人">
          <el-input v-model="formData.check_user_name" />
        </el-form-item>
        <el-form-item label="检验日期">
          <el-date-picker v-model="formData.check_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="结果" prop="result">
          <el-select v-model="formData.result" placeholder="请选择">
            <el-option label="待检验" :value="1" />
            <el-option label="合格" :value="2" />
            <el-option label="不合格" :value="3" />
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
import { getOQCList, createOQC, updateOQC, deleteOQC } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ oqc_no: '', shipping_no: '', customer_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive<any>({
  id: 0, oqc_no: '', shipping_no: '', customer_name: '', quantity: 0,
  qualified_qty: 0, unqualified_qty: 0, check_user_name: '', check_date: '', result: 1, remark: ''
})

const rules: FormRules = {
  oqc_no: [{ required: true, message: '请输入检验单号', trigger: 'blur' }],
  shipping_no: [{ required: true, message: '请输入发货单号', trigger: 'blur' }],
  quantity: [{ required: true, message: '请输入数量', trigger: 'blur' }],
  result: [{ required: true, message: '请选择结果', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑OQC' : '新增OQC')
const getStatusText = (s: number) => ({ 1: '待检验', 2: '合格', 3: '不合格' })[s] || '未知'
const getStatusType = (s: number) => ({ 1: 'info', 2: 'success', 3: 'danger' })[s] || 'info'

const loadData = async () => {
  loading.value = true
  try {
    const res = await getOQCList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally { loading.value = false }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.oqc_no = ''; searchForm.shipping_no = ''; searchForm.customer_name = ''; handleSearch() }
const handleAdd = () => {
  Object.assign(formData, { id: 0, oqc_no: '', shipping_no: '', customer_name: '', quantity: 0, qualified_qty: 0, unqualified_qty: 0, check_user_name: '', check_date: '', result: 1, remark: '' })
  dialogVisible.value = true
}
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该OQC记录吗？', '提示', { type: 'warning' })
    await deleteOQC(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateOQC(formData.id, formData) : await createOQC(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.oqc-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
