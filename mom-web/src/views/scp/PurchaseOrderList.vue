<template>
  <div class="purchase-order-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="订单号">
          <el-input v-model="searchForm.po_no" placeholder="请输入订单号" clearable />
        </el-form-item>
        <el-form-item label="供应商">
          <el-select v-model="searchForm.supplier_id" placeholder="请选择供应商" clearable filterable>
            <el-option v-for="s in supplierOptions" :key="s.id" :label="s.name" :value="s.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="草稿" value="DRAFT" />
            <el-option label="待审批" value="PENDING" />
            <el-option label="已审批" value="APPROVED" />
            <el-option label="已发布" value="ISSUED" />
            <el-option label="部分收货" value="PARTIAL" />
            <el-option label="已收货" value="RECEIVED" />
            <el-option label="已关闭" value="CLOSED" />
            <el-option label="已取消" value="CANCELLED" />
          </el-select>
        </el-form-item>
        <el-form-item label="订单日期">
          <el-date-picker v-model="searchForm.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('scp:purchase:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="po_no" label="订单号" width="150" />
        <el-table-column prop="supplier_name" label="供应商" min-width="150" />
        <el-table-column prop="order_date" label="订单日期" width="120" />
        <el-table-column prop="promised_date" label="承诺交期" width="120" />
        <el-table-column prop="total_qty" label="总数量" width="100" />
        <el-table-column prop="total_amount" label="总金额" width="120">
          <template #default="{ row }">
            {{ formatCurrency(row.total_amount) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="approval_status" label="审批状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getApprovalType(row.approval_status)">{{ getApprovalText(row.approval_status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="220" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('scp:purchase:view')" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" v-if="hasPermission('scp:purchase:edit') && row.status === 'DRAFT'" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" v-if="hasPermission('scp:purchase:submit') && row.status === 'DRAFT'" @click="handleSubmit(row)">提交</el-button>
            <el-button link type="danger" v-if="hasPermission('scp:purchase:delete') && row.status === 'DRAFT'" @click="handleDelete(row)">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="900px" :close-on-click-modal="false">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="供应商" prop="supplier_id">
              <el-select v-model="formData.supplier_id" placeholder="请选择供应商" filterable @change="handleSupplierChange">
                <el-option v-for="s in supplierOptions" :key="s.id" :label="s.name" :value="s.id" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="订单日期" prop="order_date">
              <el-date-picker v-model="formData.order_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="承诺交期" prop="promised_date">
              <el-date-picker v-model="formData.promised_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="订单类型" prop="po_type">
              <el-select v-model="formData.po_type">
                <el-option label="标准订单" value="STANDARD" />
                <el-option label="紧急订单" value="URGENT" />
                <el-option label="长期订单" value="LONG_TERM" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种">
              <el-select v-model="formData.currency">
                <el-option label="人民币" value="CNY" />
                <el-option label="美元" value="USD" />
                <el-option label="欧元" value="EUR" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税率">
              <el-input-number v-model="formData.tax_rate" :min="0" :max="100" :precision="2" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="付款条款">
              <el-input v-model="formData.payment_terms" placeholder="请输入付款条款" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="备注">
              <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入备注" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-divider>订单明细</el-divider>
        <div class="items-section">
          <el-table :data="formData.items" border size="small">
            <el-table-column label="行号" width="60" type="index" />
            <el-table-column label="物料" width="200">
              <template #default="{ row, $index }">
                <el-select v-model="row.material_id" placeholder="选择物料" filterable @change="handleMaterialChange($index)">
                  <el-option v-for="m in materialOptions" :key="m.id" :label="`${m.material_code} - ${m.material_name}`" :value="m.id" />
                </el-select>
              </template>
            </el-table-column>
            <el-table-column prop="material_code" label="物料编码" width="120" />
            <el-table-column prop="material_name" label="物料名称" width="150" />
            <el-table-column prop="specification" label="规格型号" width="120" />
            <el-table-column label="单位" width="80">
              <template #default="{ row }">
                <el-input v-model="row.unit" size="small" />
              </template>
            </el-table-column>
            <el-table-column label="数量" width="120">
              <template #default="{ row, $index }">
                <el-input-number v-model="row.order_qty" :min="0" size="small" @blur="calcLineAmount($index)" />
              </template>
            </el-table-column>
            <el-table-column label="单价" width="120">
              <template #default="{ row, $index }">
                <el-input-number v-model="row.unit_price" :min="0" :precision="4" size="small" @blur="calcLineAmount($index)" />
              </template>
            </el-table-column>
            <el-table-column label="交期" width="120">
              <template #default="{ row }">
                <el-date-picker v-model="row.promised_date" type="date" value-format="YYYY-MM-DD" size="small" placeholder="交期" />
              </template>
            </el-table-column>
            <el-table-column label="金额" width="120">
              <template #default="{ row }">
                {{ formatCurrency(row.line_amount || 0) }}
              </template>
            </el-table-column>
            <el-table-column label="操作" width="60">
              <template #default="{ $index }">
                <el-button type="danger" link size="small" @click="removeItem($index)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-button type="primary" link size="small" @click="addItem" class="add-item-btn">+ 添加物料</el-button>
        </div>

        <el-divider />
        <div class="amount-summary">
          <el-row>
            <el-col :span="6">
              <span class="label">总数量：</span>
              <span class="value">{{ formData.total_qty || 0 }}</span>
            </el-col>
            <el-col :span="6">
              <span class="label">总金额：</span>
              <span class="value">{{ formatCurrency(formData.total_amount || 0) }}</span>
            </el-col>
            <el-col :span="6">
              <span class="label">总税额：</span>
              <span class="value">{{ formatCurrency(formData.total_tax_amount || 0) }}</span>
            </el-col>
            <el-col :span="6">
              <span class="label">含税总金额：</span>
              <span class="value highlight">{{ formatCurrency(formData.total_amount_with_tax || 0) }}</span>
            </el-col>
          </el-row>
        </div>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmitForm">确定</el-button>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="采购订单详情" width="1000px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="订单号">{{ detailData.po_no }}</el-descriptions-item>
        <el-descriptions-item label="供应商">{{ detailData.supplier_name }}</el-descriptions-item>
        <el-descriptions-item label="订单类型">{{ getPOTypeText(detailData.po_type) }}</el-descriptions-item>
        <el-descriptions-item label="订单日期">{{ detailData.order_date }}</el-descriptions-item>
        <el-descriptions-item label="承诺交期">{{ detailData.promised_date }}</el-descriptions-item>
        <el-descriptions-item label="币种">{{ detailData.currency }}</el-descriptions-item>
        <el-descriptions-item label="税率">{{ detailData.tax_rate }}%</el-descriptions-item>
        <el-descriptions-item label="总金额">{{ formatCurrency(detailData.total_amount) }}</el-descriptions-item>
        <el-descriptions-item label="订单状态">
          <el-tag :type="getStatusType(detailData.status)">{{ getStatusText(detailData.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="审批状态">
          <el-tag :type="getApprovalType(detailData.approval_status)">{{ getApprovalText(detailData.approval_status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="付款条款" :span="2">{{ detailData.payment_terms || '-' }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>订单明细</el-divider>
      <el-table :data="detailData.items" border size="small">
        <el-table-column prop="line_no" label="行号" width="60" />
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="specification" label="规格" width="120" />
        <el-table-column prop="unit" label="单位" width="80" />
        <el-table-column prop="order_qty" label="订单数量" width="100" />
        <el-table-column prop="delivered_qty" label="已交货" width="100" />
        <el-table-column prop="received_qty" label="已收货" width="100" />
        <el-table-column prop="unit_price" label="单价" width="100">
          <template #default="{ row }">
            {{ formatCurrency(row.unit_price) }}
          </template>
        </el-table-column>
        <el-table-column prop="line_amount" label="金额" width="120">
          <template #default="{ row }">
            {{ formatCurrency(row.line_amount) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag size="small">{{ getItemStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="success"
              size="small"
              v-if="hasPermission('scp:purchase:receive') && ['ISSUED', 'PARTIAL'].includes(detailData.status) && row.status !== 'COMPLETED'"
              @click="handleItemReceive(row)"
            >
              收货
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-divider>订单进度</el-divider>
      <el-steps :active="getStepActive(detailData.status)" align-center>
        <el-step title="创建" :description="detailData.created_at" />
        <el-step title="提交审批" />
        <el-step title="审批通过" />
        <el-step title="发布执行" />
        <el-step title="完成" />
      </el-steps>

      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
        <el-button type="primary" v-if="detailData.status === 'DRAFT'" @click="handleSubmit(detailData)">提交审批</el-button>
        <el-button type="success" v-if="hasPermission('scp:purchase:approve') && detailData.approval_status === 'PENDING'" @click="handleApprove(detailData)">审批通过</el-button>
        <el-button type="danger" v-if="hasPermission('scp:purchase:reject') && detailData.approval_status === 'PENDING'" @click="handleReject(detailData)">拒绝</el-button>
        <el-button type="warning" v-if="hasPermission('scp:purchase:cancel') && ['DRAFT', 'PENDING', 'APPROVED'].includes(detailData.status)" @click="handleCancel(detailData)">取消</el-button>
        <el-button type="primary" v-if="detailData.status === 'APPROVED'" @click="handleIssue(detailData)">发布执行</el-button>
        <el-button type="danger" v-if="hasPermission('scp:purchase:close') && ['ISSUED', 'PARTIAL', 'RECEIVED'].includes(detailData.status)" @click="handleClose(detailData)">关闭订单</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import {
  getPurchaseOrderList,
  getPurchaseOrderById,
  createPurchaseOrder,
  updatePurchaseOrder,
  deletePurchaseOrder,
  submitPurchaseOrder,
  approvePurchaseOrder,
  rejectPurchaseOrder,
  cancelPurchaseOrder,
  closePurchaseOrder,
  issuePurchaseOrder,
  receivePurchaseOrder,
  getSupplierList,
  getMaterialList
} from '@/api/scp'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const detailVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const supplierOptions = ref<any[]>([])
const materialOptions = ref<any[]>([])

const searchForm = reactive({
  po_no: '',
  supplier_id: undefined as number | undefined,
  status: '',
  dateRange: [] as string[]
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

interface OrderItem {
  id?: number
  material_id?: number
  material_code: string
  material_name: string
  specification?: string
  unit: string
  order_qty: number
  unit_price: number
  promised_date?: string
  line_amount: number
  tax_amount: number
  status: string
}

const formData = reactive<any>({
  id: 0,
  supplier_id: undefined as number | undefined,
  supplier_name: '',
  supplier_code: '',
  order_date: '',
  promised_date: '',
  po_type: 'STANDARD',
  currency: 'CNY',
  tax_rate: 13,
  payment_terms: '',
  remark: '',
  items: [] as OrderItem[]
})

const detailData = reactive<any>({ items: [] })

const rules: FormRules = {
  supplier_id: [{ required: true, message: '请选择供应商', trigger: 'change' }],
  order_date: [{ required: true, message: '请选择订单日期', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑采购订单' : '新增采购订单')

const formatCurrency = (value: number) => {
  if (!value) return '¥0.00'
  return `¥${value.toFixed(2).replace(/\B(?=(\d{3})+(?!\d))/g, ',')}`
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: '草稿', PENDING: '待审批', APPROVED: '已审批', ISSUED: '已发布',
    PARTIAL: '部分收货', RECEIVED: '已收货', CLOSED: '已关闭', CANCELLED: '已取消'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: 'info', PENDING: 'warning', APPROVED: 'success', ISSUED: 'primary',
    PARTIAL: 'warning', RECEIVED: 'success', CLOSED: 'info', CANCELLED: 'danger'
  }
  return map[status] || 'info'
}

const getApprovalText = (status: string) => {
  const map: Record<string, string> = { PENDING: '待审批', APPROVED: '已通过', REJECTED: '已拒绝' }
  return map[status] || status
}

const getApprovalType = (status: string) => {
  const map: Record<string, string> = { PENDING: 'warning', APPROVED: 'success', REJECTED: 'danger' }
  return map[status] || 'info'
}

const getPOTypeText = (type: string) => {
  const map: Record<string, string> = { STANDARD: '标准订单', URGENT: '紧急订单', LONG_TERM: '长期订单' }
  return map[type] || type
}

const getItemStatusText = (status: string) => {
  const map: Record<string, string> = { PENDING: '待交货', PARTIAL: '部分交货', COMPLETED: '已完成', CANCELLED: '已取消' }
  return map[status] || status
}

const getStepActive = (status: string) => {
  const map: Record<string, number> = { DRAFT: 0, PENDING: 1, APPROVED: 2, ISSUED: 3, PARTIAL: 4, RECEIVED: 4, CLOSED: 4 }
  return map[status] || 0
}

const calcLineAmount = (index: number) => {
  const item = formData.items[index]
  if (item && item.order_qty && item.unit_price) {
    const amount = item.order_qty * item.unit_price
    const taxAmount = amount * (formData.tax_rate / 100)
    item.line_amount = amount
    item.tax_amount = taxAmount
  } else {
    item.line_amount = 0
    item.tax_amount = 0
  }
  calcTotalAmount()
}

const calcTotalAmount = () => {
  let totalQty = 0
  let totalAmount = 0
  let totalTaxAmount = 0
  formData.items.forEach((item: OrderItem) => {
    totalQty += item.order_qty || 0
    totalAmount += item.line_amount || 0
    totalTaxAmount += item.tax_amount || 0
  })
  formData.total_qty = totalQty
  formData.total_amount = totalAmount
  formData.total_tax_amount = totalTaxAmount
  formData.total_amount_with_tax = totalAmount + totalTaxAmount
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.po_no) params.po_no = searchForm.po_no
    if (searchForm.supplier_id) params.supplier_id = searchForm.supplier_id
    if (searchForm.status) params.status = searchForm.status
    if (searchForm.dateRange && searchForm.dateRange.length === 2) {
      params.start_date = searchForm.dateRange[0]
      params.end_date = searchForm.dateRange[1]
    }
    const res = await getPurchaseOrderList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadSuppliers = async () => {
  try {
    const res = await getSupplierList({ page: 1, page_size: 1000 })
    supplierOptions.value = res.data.list || []
  } catch (error) {
    console.error('Failed to load suppliers')
  }
}

const loadMaterials = async () => {
  try {
    const res = await getMaterialList({ page: 1, page_size: 1000 })
    materialOptions.value = res.data.list || []
  } catch (error) {
    console.error('Failed to load materials')
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }

const handleReset = () => {
  searchForm.po_no = ''
  searchForm.supplier_id = undefined
  searchForm.status = ''
  searchForm.dateRange = []
  handleSearch()
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0, supplier_id: undefined, supplier_name: '', supplier_code: '',
    order_date: new Date().toISOString().split('T')[0], promised_date: '',
    po_type: 'STANDARD', currency: 'CNY', tax_rate: 13, payment_terms: '',
    remark: '', items: [], total_qty: 0, total_amount: 0, total_tax_amount: 0, total_amount_with_tax: 0
  })
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  try {
    const res = await getPurchaseOrderById(row.id)
    const data = res.data
    const supplier = supplierOptions.value.find(s => s.id === data.supplier_id)
    Object.assign(formData, {
      id: data.id,
      supplier_id: data.supplier_id,
      supplier_name: data.supplier_name,
      supplier_code: data.supplier_code,
      order_date: data.order_date,
      promised_date: data.promised_date,
      po_type: data.po_type,
      currency: data.currency,
      tax_rate: data.tax_rate,
      payment_terms: data.payment_terms || '',
      remark: data.remark || '',
      items: (data.items || []).map((item: any) => ({
        id: item.id,
        material_id: item.material_id,
        material_code: item.material_code,
        material_name: item.material_name,
        specification: item.specification,
        unit: item.unit,
        order_qty: item.order_qty,
        unit_price: item.unit_price,
        promised_date: item.promised_date,
        line_amount: item.line_amount,
        tax_amount: item.tax_amount,
        status: item.status
      })),
      total_qty: data.total_qty,
      total_amount: data.total_amount,
      total_tax_amount: data.total_amount * (data.tax_rate / (100 + data.tax_rate)),
      total_amount_with_tax: data.total_amount
    })
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取订单详情失败')
  }
}

const handleView = async (row: any) => {
  try {
    const res = await getPurchaseOrderById(row.id)
    Object.assign(detailData, res.data)
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取订单详情失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该采购订单吗？', '提示', { type: 'warning' })
    await deletePurchaseOrder(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleSubmit = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定提交该采购订单进行审批吗？', '提示', { type: 'warning' })
    await submitPurchaseOrder(row.id)
    ElMessage.success('提交成功')
    loadData()
  } catch (error) {
    ElMessage.error('提交失败')
  }
}

const handleCancel = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定取消该采购订单吗？', '提示', { type: 'warning' })
    await cancelPurchaseOrder(row.id)
    ElMessage.success('取消成功')
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('取消失败')
  }
}

const handleClose = async (row: any) => {
  try {
    const { value: closeReason } = await ElMessageBox.prompt('请输入关闭原因', '关闭订单', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputType: 'textarea',
      inputPlaceholder: '请填写关闭原因',
      inputValidator: (val) => {
        if (!val || val.trim() === '') {
          return '关闭原因不能为空'
        }
        return true
      }
    })
    await closePurchaseOrder(row.id, closeReason)
    ElMessage.success('关闭成功')
    detailVisible.value = false
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '关闭失败')
    }
  }
}

// 发布执行
const handleIssue = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定发布执行该采购订单吗？', '提示', { type: 'warning' })
    await issuePurchaseOrder(row.id)
    ElMessage.success('发布执行成功')
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('发布执行失败')
  }
}

// 行项目收货
const handleItemReceive = async (row: any) => {
  try {
    const { value: receivedQty } = await ElMessageBox.prompt('确认收货', '请输入收货数量', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputType: 'number',
      inputPlaceholder: `最大可收货: ${row.order_qty - (row.received_qty || 0)}`,
      inputValidator: (val) => {
        const qty = parseFloat(val)
        if (isNaN(qty) || qty <= 0) {
          return '请输入有效的收货数量'
        }
        const maxQty = row.order_qty - (row.received_qty || 0)
        if (qty > maxQty) {
          return `收货数量不能超过最大可收货数量 ${maxQty}`
        }
        return true
      }
    })
    const { value: batchNo } = await ElMessageBox.prompt('确认收货', '请输入批号（可选）', {
      confirmButtonText: '确定',
      cancelButtonText: '跳过',
      inputPlaceholder: '批号'
    })
    await receivePurchaseOrder(row.id, parseFloat(receivedQty), batchNo || undefined)
    ElMessage.success('收货成功')
    // 刷新详情数据
    const res = await getPurchaseOrderById(detailData.value.id)
    detailData.value = res.data
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '收货失败')
    }
  }
}

// 审批通过
const handleApprove = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定审批通过该采购订单吗？', '提示', { type: 'warning' })
    await approvePurchaseOrder(row.id)
    ElMessage.success('审批成功')
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('审批失败')
  }
}

// 拒绝
const handleReject = async (row: any) => {
  try {
    const { value } = await ElMessageBox.prompt('请输入拒绝原因', '拒绝审批', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputType: 'textarea'
    })
    await rejectPurchaseOrder(row.id)
    ElMessage.success('已拒绝')
    detailVisible.value = false
    loadData()
  } catch (error) {
    // 用户取消不做处理
  }
}

const handleSupplierChange = (supplierId: number) => {
  const supplier = supplierOptions.value.find(s => s.id === supplierId)
  if (supplier) {
    formData.supplier_name = supplier.supplier_name
    formData.supplier_code = supplier.supplier_code
  }
}

const handleMaterialChange = (index: number) => {
  const item = formData.items[index]
  const material = materialOptions.value.find(m => m.id === item.material_id)
  if (material) {
    item.material_code = material.material_code
    item.material_name = material.material_name
    item.specification = material.specification
    item.unit = material.unit || 'PCS'
  }
  calcLineAmount(index)
}

const addItem = () => {
  formData.items.push({
    material_id: undefined, material_code: '', material_name: '', specification: '',
    unit: 'PCS', order_qty: 0, unit_price: 0, promised_date: '', line_amount: 0, tax_amount: 0, status: 'PENDING'
  })
}

const removeItem = (index: number) => {
  formData.items.splice(index, 1)
  calcTotalAmount()
}

const handleSubmitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  if (formData.items.length === 0) {
    ElMessage.warning('请添加订单明细')
    return
  }
  submitLoading.value = true
  try {
    calcTotalAmount()
    if (formData.id) {
      await updatePurchaseOrder(formData.id, formData)
    } else {
      await createPurchaseOrder(formData)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error(formData.id ? '更新失败' : '创建失败')
  } finally {
    submitLoading.value = false
  }
}

onMounted(() => {
  loadData()
  loadSuppliers()
  loadMaterials()
})
</script>

<style scoped lang="scss">
.purchase-order-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .items-section {
    margin-bottom: 16px;
    .add-item-btn { margin-top: 12px; }
  }
  .amount-summary {
    padding: 16px;
    background: #f5f7fa;
    border-radius: 4px;
    .label { color: #606266; }
    .value { font-weight: bold; color: #303133; }
    .value.highlight { color: #f56c6c; font-size: 16px; }
  }
}
</style>
