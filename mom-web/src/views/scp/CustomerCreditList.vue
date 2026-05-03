<template>
  <div class="customer-credit-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="客户编码">
          <el-input v-model="searchForm.customer_code" placeholder="请输入客户编码" clearable />
        </el-form-item>
        <el-form-item label="客户名称">
          <el-input v-model="searchForm.customer_name" placeholder="请输入客户名称" clearable />
        </el-form-item>
        <el-form-item label="信用等级">
          <el-select v-model="searchForm.credit_level" placeholder="请选择" clearable style="width: 100px">
            <el-option label="A" value="A" />
            <el-option label="B" value="B" />
            <el-option label="C" value="C" />
            <el-option label="D" value="D" />
          </el-select>
        </el-form-item>
        <el-form-item label="风险等级">
          <el-select v-model="searchForm.risk_level" placeholder="请选择" clearable style="width: 100px">
            <el-option label="低" :value="1" />
            <el-option label="中" :value="2" />
            <el-option label="高" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="黑名单">
          <el-select v-model="searchForm.blacklist" placeholder="请选择" clearable style="width: 100px">
            <el-option label="是" :value="1" />
            <el-option label="否" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon>新增信用
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="customer_code" label="客户编码" width="120" />
        <el-table-column prop="customer_name" label="客户名称" min-width="150" />
        <el-table-column prop="credit_limit" label="信用额度" width="120" :formatter="formatCurrency" />
        <el-table-column prop="used_credit" label="已用额度" width="120" :formatter="formatCurrency" />
        <el-table-column prop="available_credit" label="可用额度" width="120" :formatter="formatCurrency" />
        <el-table-column prop="credit_level" label="信用等级" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.credit_level === 'A'" type="success">A</el-tag>
            <el-tag v-else-if="row.credit_level === 'B'" type="warning">B</el-tag>
            <el-tag v-else-if="row.credit_level === 'C'" type="info">C</el-tag>
            <el-tag v-else type="danger">D</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="risk_level" label="风险等级" width="100">
          <template #default="{ row }">
            <el-tag :type="row.risk_level === 1 ? 'success' : row.risk_level === 2 ? 'warning' : 'danger'">
              {{ row.risk_level === 1 ? '低' : row.risk_level === 2 ? '中' : '高' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="blacklist" label="黑名单" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.blacklist === 1" type="danger">是</el-tag>
            <span v-else>否</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">
              {{ row.status === 1 ? '正常' : '冻结' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="warning" size="small" @click="handleFreeze(row)" v-if="row.status === 1">冻结</el-button>
            <el-button link type="success" size="small" @click="handleUnfreeze(row)" v-else>解冻</el-button>
            <el-button link type="danger" size="small" @click="handleDelete(row)">删除</el-button>
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
      <el-form :model="formData" label-width="100px" :rules="formRules" ref="formRef">
        <el-form-item label="客户" prop="customer_id">
          <el-select v-model="formData.customer_id" placeholder="请选择客户" filterable style="width: 100%" @change="handleCustomerChange">
            <el-option v-for="c in customerList" :key="c.id" :label="`${c.customer_code} - ${c.customer_name}`" :value="c.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="信用额度" prop="credit_limit">
          <el-input-number v-model="formData.credit_limit" :min="0" :precision="2" style="width: 100%" />
        </el-form-item>
        <el-form-item label="信用等级">
          <el-select v-model="formData.credit_level" placeholder="请选择信用等级" style="width: 100%">
            <el-option label="A" value="A" />
            <el-option label="B" value="B" />
            <el-option label="C" value="C" />
            <el-option label="D" value="D" />
          </el-select>
        </el-form-item>
        <el-form-item label="账期(天)">
          <el-input-number v-model="formData.payment_days" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="风险等级">
          <el-select v-model="formData.risk_level" placeholder="请选择风险等级" style="width: 100%">
            <el-option label="低" :value="1" />
            <el-option label="中" :value="2" />
            <el-option label="高" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="预警阈值">
          <el-input-number v-model="formData.alert_threshold" :min="0" :max="100" :precision="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remarks" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getCustomerCreditList, createCustomerCredit, updateCustomerCredit, deleteCustomerCredit, freezeCustomerCredit, unfreezeCustomerCredit } from '@/api/customer_credit'
import { getCustomerList } from '@/api/mdm'

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const tableData = ref<any[]>([])
const customerList = ref<any[]>([])
const dialogTitle = ref('新增信用')

const searchForm = reactive({
  customer_code: '',
  customer_name: '',
  credit_level: '',
  risk_level: null as number | null,
  blacklist: null as number | null
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  id: null,
  customer_id: null,
  customer_code: '',
  customer_name: '',
  credit_limit: 0,
  credit_level: 'A',
  payment_days: 30,
  risk_level: 1,
  alert_threshold: 80,
  remarks: ''
})

const formRules = {
  customer_id: [{ required: true, message: '请选择客户', trigger: 'change' }],
  credit_limit: [{ required: true, message: '请输入信用额度', trigger: 'blur' }]
}

const formatCurrency = (row: any, col: any, cellValue: number) => {
  if (cellValue == null) return '-'
  return cellValue.toLocaleString('zh-CN', { minimumFractionDigits: 2 })
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getCustomerCreditList({
      page: pagination.page,
      page_size: pagination.pageSize,
      customer_code: searchForm.customer_code,
      customer_name: searchForm.customer_name,
      credit_level: searchForm.credit_level,
      risk_level: searchForm.risk_level,
      blacklist: searchForm.blacklist
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadCustomers = async () => {
  try {
    const res = await getCustomerList({ page: 1, page_size: 500 })
    customerList.value = res.data.list || []
  } catch (e) {
    console.error(e)
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.customer_code = ''
  searchForm.customer_name = ''
  searchForm.credit_level = ''
  searchForm.risk_level = null
  searchForm.blacklist = null
  handleSearch()
}

const handleCustomerChange = (id: number) => {
  const c = customerList.value.find(cust => cust.id === id)
  if (c) {
    formData.customer_code = c.customer_code
    formData.customer_name = c.customer_name
  }
}

const handleCreate = () => {
  dialogTitle.value = '新增信用'
  formData.id = null
  formData.customer_id = null
  formData.customer_code = ''
  formData.customer_name = ''
  formData.credit_limit = 0
  formData.credit_level = 'A'
  formData.payment_days = 30
  formData.risk_level = 1
  formData.alert_threshold = 80
  formData.remarks = ''
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  dialogTitle.value = '编辑信用'
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleSave = async () => {
  if (!formData.customer_id || formData.credit_limit === undefined) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    if (formData.id) {
      await updateCustomerCredit(formData.id, formData)
    } else {
      await createCustomerCredit(formData)
    }
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleFreeze = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定冻结该客户信用吗？', '提示', { type: 'warning' })
    await freezeCustomerCredit(row.id)
    ElMessage.success('冻结成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

const handleUnfreeze = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定解冻该客户信用吗？', '提示', { type: 'warning' })
    await unfreezeCustomerCredit(row.id)
    ElMessage.success('解冻成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该客户信用吗？', '提示', { type: 'warning' })
    await deleteCustomerCredit(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

onMounted(() => { loadData(); loadCustomers() })
</script>

<style scoped lang="scss">
.customer-credit-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>