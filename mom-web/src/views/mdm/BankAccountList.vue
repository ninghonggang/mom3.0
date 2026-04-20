<template>
  <div class="bank-account-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="开户银行">
          <el-input v-model="searchForm.bank_name" placeholder="请输入开户银行" clearable />
        </el-form-item>
        <el-form-item label="所属">
          <el-select v-model="searchForm.owner_type" placeholder="请选择" clearable>
            <el-option label="客户" value="CUSTOMER" />
            <el-option label="供应商" value="SUPPLIER" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:bank_account:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="bank_name" label="开户银行" width="150" />
        <el-table-column prop="account_name" label="账户名称" width="150" />
        <el-table-column prop="bank_account" label="银行账号" width="180" />
        <el-table-column prop="branch_name" label="支行名称" width="150" />
        <el-table-column prop="currency" label="币种" width="80" />
        <el-table-column prop="owner_type" label="所属类型" width="100">
          <template #default="{ row }">
            {{ row.owner_type === 'CUSTOMER' ? '客户' : row.owner_type === 'SUPPLIER' ? '供应商' : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="is_primary" label="主账户" width="80">
          <template #default="{ row }">
            <el-tag :type="row.is_primary ? 'success' : 'info'" size="small">{{ row.is_primary ? '是' : '否' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 'ACTIVE' ? 'success' : 'danger'">{{ row.status === 'ACTIVE' ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('mdm:bank_account:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:bank_account:delete')" @click="handleDelete(row)">删除</el-button>
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
          </el-select>
        </el-form-item>
        <el-form-item label="所属ID" prop="owner_id">
          <el-input-number v-model="formData.owner_id" :min="1" />
        </el-form-item>
        <el-form-item label="开户银行" prop="bank_name"><el-input v-model="formData.bank_name" /></el-form-item>
        <el-form-item label="账户名称" prop="account_name"><el-input v-model="formData.account_name" /></el-form-item>
        <el-form-item label="银行账号" prop="bank_account"><el-input v-model="formData.bank_account" /></el-form-item>
        <el-form-item label="支行名称"><el-input v-model="formData.branch_name" /></el-form-item>
        <el-form-item label="币种">
          <el-select v-model="formData.currency" placeholder="请选择">
            <el-option label="人民币" value="CNY" />
            <el-option label="美元" value="USD" />
            <el-option label="欧元" value="EUR" />
            <el-option label="日元" value="JPY" />
          </el-select>
        </el-form-item>
        <el-form-item label="主账户">
          <el-switch v-model="formData.is_primary" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="formData.status">
            <el-radio value="ACTIVE">启用</el-radio>
            <el-radio value="INACTIVE">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注"><el-input v-model="formData.remark" type="textarea" /></el-form-item>
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
import { getBankAccountList, getBankAccountById, createBankAccount, updateBankAccount, deleteBankAccount } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const searchForm = reactive({ bank_name: '', owner_type: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, owner_type: '', owner_id: 0, bank_name: '', account_name: '', bank_account: '', branch_name: '', currency: 'CNY', is_primary: false, status: 'ACTIVE', remark: '' })
const rules: FormRules = { owner_type: [{ required: true, message: '请选择所属类型', trigger: 'change' }], owner_id: [{ required: true, message: '请输入所属ID', trigger: 'blur' }], bank_name: [{ required: true, message: '请输入开户银行', trigger: 'blur' }], account_name: [{ required: true, message: '请输入账户名称', trigger: 'blur' }], bank_account: [{ required: true, message: '请输入银行账号', trigger: 'blur' }] }
const dialogTitle = computed(() => formData.id ? '编辑银行账户' : '新增银行账户')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getBankAccountList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.bank_name = ''; searchForm.owner_type = ''; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, owner_type: '', owner_id: 0, bank_name: '', account_name: '', bank_account: '', branch_name: '', currency: 'CNY', is_primary: false, status: 'ACTIVE', remark: '' }); dialogVisible.value = true }
const handleEdit = async (row: any) => {
  try {
    const res = await getBankAccountById(row.id)
    Object.assign(formData, res.data)
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取银行账户详情失败')
  }
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该银行账户吗？', '提示', { type: 'warning' })
    await deleteBankAccount(row.id)
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
      await updateBankAccount(formData.id, formData)
    } else {
      await createBankAccount(formData)
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
.bank-account-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
