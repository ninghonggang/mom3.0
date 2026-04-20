<template>
  <div class="customer-inquiry-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="询价单号">
          <el-input v-model="searchForm.inquiry_no" placeholder="请输入询价单号" clearable />
        </el-form-item>
        <el-form-item label="客户">
          <el-select v-model="searchForm.customer_id" placeholder="请选择客户" clearable filterable>
            <el-option v-for="c in customerOptions" :key="c.id" :label="c.customer_name" :value="c.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" value="DRAFT" />
            <el-option label="已发送" value="SENT" />
            <el-option label="已报价" value="QUOTED" />
            <el-option label="已赢单" value="WON" />
            <el-option label="已丢单" value="LOST" />
            <el-option label="已取消" value="CANCELLED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('scp:inquiry:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新建询价
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="inquiry_no" label="询价单号" width="150" />
        <el-table-column prop="customer_name" label="客户" min-width="150" />
        <el-table-column prop="contact_person" label="联系人" width="100" />
        <el-table-column prop="contact_phone" label="联系电话" width="130" />
        <el-table-column prop="inquiry_date" label="询价日期" width="120" />
        <el-table-column prop="expected_date" label="期望交期" width="120" />
        <el-table-column prop="quoted_amount" label="报价金额" width="120" align="right">
          <template #default="{ row }">
            {{ row.quoted_amount ? formatCurrency(row.quoted_amount) : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160" />
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('scp:inquiry:view')" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" v-if="hasPermission('scp:inquiry:edit') && row.status === 'DRAFT'" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" v-if="hasPermission('scp:inquiry:send') && row.status === 'DRAFT'" @click="handleSend(row)">发送</el-button>
            <el-button link type="warning" v-if="hasPermission('scp:inquiry:quote') && row.status === 'SENT'" @click="handleQuote(row)">报价</el-button>
            <el-button link type="danger" v-if="hasPermission('scp:inquiry:delete') && row.status === 'DRAFT'" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新建/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="800px" :close-on-click-modal="false">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客户" prop="customer_id">
              <el-select v-model="formData.customer_id" placeholder="请选择客户" filterable @change="handleCustomerChange">
                <el-option v-for="c in customerOptions" :key="c.id" :label="c.customer_name" :value="c.id" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="询价日期" prop="inquiry_date">
              <el-date-picker v-model="formData.inquiry_date" type="date" value-format="YYYY-MM-DD" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="联系人">
              <el-input v-model="formData.contact_person" placeholder="联系人" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话">
              <el-input v-model="formData.contact_phone" placeholder="电话" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="期望交期">
              <el-date-picker v-model="formData.expected_date" type="date" value-format="YYYY-MM-DD" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="有效期至">
              <el-date-picker v-model="formData.valid_until" type="date" value-format="YYYY-MM-DD" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="币种">
              <el-select v-model="formData.currency" style="width:100%">
                <el-option label="人民币" value="CNY" />
                <el-option label="美元" value="USD" />
                <el-option label="欧元" value="EUR" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-col :span="24">
          <el-form-item label="备注">
            <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="备注" />
          </el-form-item>
        </el-col>

        <!-- 询价明细 -->
        <el-divider content-position="left">询价明细</el-divider>
        <div class="items-toolbar">
          <el-button size="small" type="primary" @click="handleAddItem">
            <el-icon><Plus /></el-icon>添加物料
          </el-button>
        </div>
        <el-table :data="formData.items" border size="small" class="items-table">
          <el-table-column label="行号" width="60" align="center">
            <template #default="{ $index }">{{ $index + 1 }}</template>
          </el-table-column>
          <el-table-column label="物料编码" width="140">
            <template #default="{ row }">
              <el-input v-model="row.material_code" placeholder="物料编码" />
            </template>
          </el-table-column>
          <el-table-column label="物料名称" width="160">
            <template #default="{ row }">
              <el-input v-model="row.material_name" placeholder="物料名称" />
            </template>
          </el-table-column>
          <el-table-column label="规格型号" width="140">
            <template #default="{ row }">
              <el-input v-model="row.specification" placeholder="规格" />
            </template>
          </el-table-column>
          <el-table-column label="单位" width="70">
            <template #default="{ row }">
              <el-input v-model="row.unit" placeholder="单位" />
            </template>
          </el-table-column>
          <el-table-column label="数量" width="100">
            <template #default="{ row }">
              <el-input-number v-model="row.required_qty" :min="0" controls-position="right" style="width:100%" />
            </template>
          </el-table-column>
          <el-table-column label="目标价" width="100">
            <template #default="{ row }">
              <el-input-number v-model="row.target_price" :min="0" :precision="2" controls-position="right" style="width:100%" />
            </template>
          </el-table-column>
          <el-table-column label="操作" width="60" align="center">
            <template #default="{ $index }">
              <el-button link type="danger" size="small" @click="handleRemoveItem($index)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="询价详情" width="900px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="询价单号">{{ detailData.inquiry_no }}</el-descriptions-item>
        <el-descriptions-item label="客户">{{ detailData.customer_name }}</el-descriptions-item>
        <el-descriptions-item label="联系人">{{ detailData.contact_person || '-' }}</el-descriptions-item>
        <el-descriptions-item label="联系电话">{{ detailData.contact_phone || '-' }}</el-descriptions-item>
        <el-descriptions-item label="询价日期">{{ detailData.inquiry_date }}</el-descriptions-item>
        <el-descriptions-item label="期望交期">{{ detailData.expected_date || '-' }}</el-descriptions-item>
        <el-descriptions-item label="有效期至">{{ detailData.valid_until || '-' }}</el-descriptions-item>
        <el-descriptions-item label="币种">{{ detailData.currency }}</el-descriptions-item>
        <el-descriptions-item label="报价金额">{{ detailData.quoted_amount ? formatCurrency(detailData.quoted_amount) : '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(detailData.status)">{{ getStatusText(detailData.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>询价明细</el-divider>
      <el-table :data="detailData.items" border size="small">
        <el-table-column prop="line_no" label="行号" width="60" />
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="specification" label="规格" width="120" />
        <el-table-column prop="unit" label="单位" width="70" />
        <el-table-column prop="required_qty" label="数量" width="100" align="right" />
        <el-table-column prop="target_price" label="目标价" width="100" align="right">
          <template #default="{ row }">
            {{ row.target_price ? formatCurrency(row.target_price) : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="quoted_price" label="报价" width="100" align="right">
          <template #default="{ row }">
            {{ row.quoted_price ? formatCurrency(row.quoted_price) : '-' }}
          </template>
        </el-table-column>
      </el-table>

      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
        <el-button type="success" v-if="detailData.status === 'SENT' && hasPermission('scp:inquiry:win')" @click="handleWin(detailData)">标记赢单</el-button>
        <el-button type="danger" v-if="detailData.status === 'SENT' && hasPermission('scp:inquiry:lose')" @click="handleLose(detailData)">标记丢单</el-button>
        <el-button type="warning" v-if="['DRAFT', 'SENT', 'QUOTED'].includes(detailData.status) && hasPermission('scp:inquiry:cancel')" @click="handleCancel(detailData)">取消</el-button>
      </template>
    </el-dialog>

    <!-- 报价对话框 -->
    <el-dialog v-model="quoteDialogVisible" title="填写报价" width="400px">
      <el-form ref="quoteFormRef" :model="quoteForm" label-width="100px">
        <el-form-item label="报价金额" prop="quoted_amount">
          <el-input-number v-model="quoteForm.quoted_amount" :min="0" :precision="2" style="width:100%" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="quoteDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleQuoteSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import {
  getCustomerInquiryList,
  getCustomerInquiryById,
  createCustomerInquiry,
  updateCustomerInquiry,
  deleteCustomerInquiry,
  sendCustomerInquiry,
  quoteCustomerInquiry,
  winCustomerInquiry,
  loseCustomerInquiry,
  cancelCustomerInquiry
} from '@/api/scp'
import { getCustomerList } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const detailVisible = ref(false)
const quoteDialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const quoteFormRef = ref<FormInstance>()

const customerOptions = ref<any[]>([])

const searchForm = reactive({
  inquiry_no: '',
  customer_id: undefined as number | undefined,
  status: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  id: 0,
  customer_id: undefined as number | undefined,
  customer_name: '',
  contact_person: '',
  contact_phone: '',
  inquiry_date: '',
  expected_date: '',
  valid_until: '',
  currency: 'CNY',
  remark: '',
  items: [] as any[]
})

const detailData = reactive<any>({ items: [] })

const quoteForm = reactive({
  quoted_amount: 0
})

const rules: FormRules = {
  customer_id: [{ required: true, message: '请选择客户', trigger: 'change' }],
  inquiry_date: [{ required: true, message: '请选择询价日期', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑询价' : '新建询价')

const formatCurrency = (value: number) => {
  if (!value) return '¥0.00'
  return `¥${value.toFixed(2).replace(/\B(?=(\d{3})+(?!\d))/g, ',')}`
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: '草稿', SENT: '已发送', QUOTED: '已报价', WON: '已赢单', LOST: '已丢单', CANCELLED: '已取消'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: 'info', SENT: 'primary', QUOTED: 'warning', WON: 'success', LOST: 'danger', CANCELLED: 'info'
  }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.inquiry_no) params.inquiry_no = searchForm.inquiry_no
    if (searchForm.customer_id) params.customer_id = searchForm.customer_id
    if (searchForm.status) params.status = searchForm.status
    const res = await getCustomerInquiryList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadCustomers = async () => {
  try {
    const res = await getCustomerList({ page: 1, page_size: 1000 })
    customerOptions.value = res.data.list || []
  } catch (error) {
    console.error('Failed to load customers')
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }

const handleReset = () => {
  searchForm.inquiry_no = ''
  searchForm.customer_id = undefined
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0, customer_id: undefined, customer_name: '',
    contact_person: '', contact_phone: '',
    inquiry_date: new Date().toISOString().split('T')[0],
    expected_date: '', valid_until: '', currency: 'CNY', remark: '', items: []
  })
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  try {
    const res = await getCustomerInquiryById(row.id)
    const data = res.data
    Object.assign(formData, {
      id: data.id,
      customer_id: data.customer_id,
      customer_name: data.customer_name,
      contact_person: data.contact_person || '',
      contact_phone: data.contact_phone || '',
      inquiry_date: data.inquiry_date,
      expected_date: data.expected_date || '',
      valid_until: data.valid_until || '',
      currency: data.currency || 'CNY',
      remark: data.remark || '',
      items: (data.items || []).map((item: any) => ({
        id: item.id,
        material_id: item.material_id,
        material_code: item.material_code,
        material_name: item.material_name,
        specification: item.specification,
        unit: item.unit,
        required_qty: item.required_qty,
        target_price: item.target_price
      }))
    })
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取详情失败')
  }
}

const handleView = async (row: any) => {
  try {
    const res = await getCustomerInquiryById(row.id)
    Object.assign(detailData, res.data)
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取详情失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该询价吗？', '提示', { type: 'warning' })
    await deleteCustomerInquiry(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}

const handleSend = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定发送该询价给客户吗？', '提示', { type: 'warning' })
    await sendCustomerInquiry(row.id)
    ElMessage.success('发送成功')
    loadData()
  } catch {}
}

const handleQuote = (row: any) => {
  Object.assign(quoteForm, { quoted_amount: 0 })
  quoteDialogVisible.value = true
}

const handleQuoteSubmit = async () => {
  try {
    await quoteFormRef.value?.validate()
    await quoteCustomerInquiry(detailData.id, quoteForm.quoted_amount)
    ElMessage.success('报价成功')
    quoteDialogVisible.value = false
    detailVisible.value = false
    loadData()
  } catch {}
}

const handleWin = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定标记为赢单吗？', '提示', { type: 'warning' })
    await winCustomerInquiry(row.id, 0)
    ElMessage.success('操作成功')
    detailVisible.value = false
    loadData()
  } catch {}
}

const handleLose = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定标记为丢单吗？', '提示', { type: 'warning' })
    await loseCustomerInquiry(row.id)
    ElMessage.success('操作成功')
    detailVisible.value = false
    loadData()
  } catch {}
}

const handleCancel = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定取消该询价吗？', '提示', { type: 'warning' })
    await cancelCustomerInquiry(row.id)
    ElMessage.success('取消成功')
    detailVisible.value = false
    loadData()
  } catch {}
}

const handleCustomerChange = (customerId: number) => {
  const customer = customerOptions.value.find(c => c.id === customerId)
  if (customer) {
    formData.customer_name = customer.customer_name
  }
}

const handleAddItem = () => {
  formData.items.push({
    material_code: '', material_name: '', specification: '', unit: 'PCS', required_qty: 1, target_price: 0
  })
}

const handleRemoveItem = (index: number) => {
  formData.items.splice(index, 1)
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (formData.id) {
      await updateCustomerInquiry(formData.id, formData)
    } else {
      await createCustomerInquiry(formData)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally {
    submitLoading.value = false
  }
}

onMounted(() => {
  loadData()
  loadCustomers()
})
</script>

<style scoped lang="scss">
.customer-inquiry-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .items-toolbar { margin-bottom: 8px; }
  .items-table { margin-bottom: 12px; }
}
</style>
