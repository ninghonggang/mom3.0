<template>
  <div class="iqc-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="检验单号">
          <el-input v-model="searchForm.iqc_no" placeholder="请输入检验单号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待检验" :value="1" />
            <el-option label="检验中" :value="2" />
            <el-option label="合格" :value="3" />
            <el-option label="不合格" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:iqc:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="iqc_no" label="检验单号" width="140" />
        <el-table-column prop="material_code" label="物料编码" width="100" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="supplier_name" label="供应商" min-width="120" />
        <el-table-column prop="quantity" label="送检数量" width="90" />
        <el-table-column prop="qualified_qty" label="合格数量" width="90" />
        <el-table-column prop="unqualified_qty" label="不合格数" width="90" />
        <el-table-column prop="check_user" label="检验人" width="80" />
        <el-table-column prop="check_date" label="检验日期" width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('quality:iqc:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('quality:iqc:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="检验单号" prop="iqc_no">
          <el-input v-model="formData.iqc_no" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="物料编码" prop="material_code">
          <el-input v-model="formData.material_code" />
        </el-form-item>
        <el-form-item label="物料名称" prop="material_name">
          <el-input v-model="formData.material_name" />
        </el-form-item>
        <el-form-item label="供应商">
          <el-input v-model="formData.supplier_name" />
        </el-form-item>
        <el-form-item label="送检数量" prop="quantity">
          <el-input-number v-model="formData.quantity" :min="0" />
        </el-form-item>
        <el-form-item label="合格数量">
          <el-input-number v-model="formData.qualified_qty" :min="0" />
        </el-form-item>
        <el-form-item label="不合格数">
          <el-input-number v-model="formData.unqualified_qty" :min="0" />
        </el-form-item>
        <el-form-item label="检验人">
          <el-input v-model="formData.check_user" />
        </el-form-item>
        <el-form-item label="检验日期">
          <el-date-picker v-model="formData.check_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option label="待检验" :value="1" />
            <el-option label="检验中" :value="2" />
            <el-option label="合格" :value="3" />
            <el-option label="不合格" :value="4" />
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
import { getIQCList, createIQC, updateIQC, deleteIQC } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ iqc_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive<any>({
  id: 0, iqc_no: '', material_code: '', material_name: '', supplier_name: '',
  quantity: 0, qualified_qty: 0, unqualified_qty: 0, check_user: '', check_date: '', status: 1, remark: ''
})

const rules: FormRules = {
  iqc_no: [{ required: true, message: '请输入检验单号', trigger: 'blur' }],
  material_code: [{ required: true, message: '请输入物料编码', trigger: 'blur' }],
  material_name: [{ required: true, message: '请输入物料名称', trigger: 'blur' }],
  quantity: [{ required: true, message: '请输入送检数量', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑IQC' : '新增IQC')

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待检验', 2: '检验中', 3: '合格', 4: '不合格' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success', 4: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getIQCList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.iqc_no = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, iqc_no: '', material_code: '', material_name: '', supplier_name: '',
    quantity: 0, qualified_qty: 0, unqualified_qty: 0, check_user: '', check_date: '', status: 1, remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该IQC记录吗？', '提示', { type: 'warning' })
    await deleteIQC(row.id)
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
      await updateIQC(formData.id, formData)
    } else {
      await createIQC(formData)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.iqc-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
