<template>
  <div class="transfer-order-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="调拨单号">
          <el-input v-model="searchForm.query" placeholder="请输入调拨单号" clearable />
        </el-form-item>
        <el-form-item label="调拨类型">
          <el-select v-model="searchForm.transfer_type" placeholder="请选择类型" clearable>
            <el-option label="调拨" value="TRANSFER" />
            <el-option label="调整" value="ADJUSTMENT" />
            <el-option label="调入" value="TRANSFER_IN" />
            <el-option label="调出" value="TRANSFER_OUT" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="草稿" value="DRAFT" />
            <el-option label="已提交" value="SUBMITTED" />
            <el-option label="已审批" value="APPROVED" />
            <el-option label="在途" value="IN_TRANSIT" />
            <el-option label="已完成" value="COMPLETED" />
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
      <el-button type="primary" v-if="hasPermission('wms:transfer:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="transfer_no" label="调拨单号" width="150" />
        <el-table-column prop="transfer_type" label="调拨类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTransferTypeText(row.transfer_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="from_warehouse_name" label="源仓库" min-width="120" />
        <el-table-column prop="to_warehouse_name" label="目标仓库" min-width="120" />
        <el-table-column prop="transfer_reason" label="调拨原因" min-width="100" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="requester_name" label="申请人" width="100" />
        <el-table-column prop="request_time" label="申请时间" width="160" />
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('wms:transfer:view')" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" v-if="hasPermission('wms:transfer:edit')" @click="handleEdit(row)" :disabled="row.status !== 'DRAFT'">编辑</el-button>
            <el-button link type="success" v-if="hasPermission('wms:transfer:submit')" @click="handleSubmit(row)" :disabled="row.status !== 'DRAFT'">提交</el-button>
            <el-button link type="danger" v-if="hasPermission('wms:transfer:delete')" @click="handleDelete(row)" :disabled="row.status !== 'DRAFT'">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="800px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="调拨类型" prop="transfer_type">
          <el-select v-model="formData.transfer_type" placeholder="请选择调拨类型">
            <el-option label="调拨" value="TRANSFER" />
            <el-option label="调整" value="ADJUSTMENT" />
            <el-option label="调入" value="TRANSFER_IN" />
            <el-option label="调出" value="TRANSFER_OUT" />
          </el-select>
        </el-form-item>
        <el-form-item label="源仓库" prop="from_warehouse_id">
          <el-input-number v-model="formData.from_warehouse_id" :min="1" />
        </el-form-item>
        <el-form-item label="源仓库名称" prop="from_warehouse_name">
          <el-input v-model="formData.from_warehouse_name" />
        </el-form-item>
        <el-form-item label="目标仓库" prop="to_warehouse_id">
          <el-input-number v-model="formData.to_warehouse_id" :min="1" />
        </el-form-item>
        <el-form-item label="目标仓库名称" prop="to_warehouse_name">
          <el-input v-model="formData.to_warehouse_name" />
        </el-form-item>
        <el-form-item label="调拨原因" prop="transfer_reason">
          <el-select v-model="formData.transfer_reason" placeholder="请选择调拨原因">
            <el-option label="生产领料" value="生产领料" />
            <el-option label="退料" value="退料" />
            <el-option label="调拨申请" value="调拨申请" />
            <el-option label="盘盈调账" value="盘盈调账" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" />
        </el-form-item>

        <el-divider>调拨明细</el-divider>
        <div v-for="(item, index) in formData.items" :key="index" class="item-row">
          <el-form-item label="物料ID" prop="material_id">
            <el-input-number v-model="item.material_id" :min="1" />
          </el-form-item>
          <el-form-item label="物料编码">
            <el-input v-model="item.material_code" />
          </el-form-item>
          <el-form-item label="物料名称">
            <el-input v-model="item.material_name" />
          </el-form-item>
          <el-form-item label="规格">
            <el-input v-model="item.specification" />
          </el-form-item>
          <el-form-item label="单位">
            <el-input v-model="item.unit" />
          </el-form-item>
          <el-form-item label="申请数量">
            <el-input-number v-model="item.request_qty" :min="0" />
          </el-form-item>
          <el-form-item label="批次号">
            <el-input v-model="item.batch_no" />
          </el-form-item>
          <el-button type="danger" link @click="removeItem(index)">删除</el-button>
        </div>
        <el-button type="primary" link @click="addItem">+ 添加物料</el-button>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmitForm">确定</el-button>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="调拨单详情" width="900px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="调拨单号">{{ detailData.order?.transfer_no }}</el-descriptions-item>
        <el-descriptions-item label="调拨类型">{{ getTransferTypeText(detailData.order?.transfer_type) }}</el-descriptions-item>
        <el-descriptions-item label="源仓库">{{ detailData.order?.from_warehouse_name }}</el-descriptions-item>
        <el-descriptions-item label="目标仓库">{{ detailData.order?.to_warehouse_name }}</el-descriptions-item>
        <el-descriptions-item label="调拨原因">{{ detailData.order?.transfer_reason }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(detailData.order?.status)">{{ getStatusText(detailData.order?.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="申请人">{{ detailData.order?.requester_name }}</el-descriptions-item>
        <el-descriptions-item label="申请时间">{{ detailData.order?.request_time }}</el-descriptions-item>
        <el-descriptions-item label="审批人">{{ detailData.order?.approver_id }}</el-descriptions-item>
        <el-descriptions-item label="审批时间">{{ detailData.order?.approved_time }}</el-descriptions-item>
        <el-descriptions-item label="备注">{{ detailData.order?.remark }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>调拨明细</el-divider>
      <el-table :data="detailData.items" border>
        <el-table-column prop="line_no" label="行号" width="60" />
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="specification" label="规格" width="120" />
        <el-table-column prop="unit" label="单位" width="60" />
        <el-table-column prop="request_qty" label="申请数量" width="100" />
        <el-table-column prop="approved_qty" label="审批数量" width="100" />
        <el-table-column prop="transfer_qty" label="调拨数量" width="100" />
        <el-table-column prop="batch_no" label="批次号" width="120" />
      </el-table>

      <el-divider>状态流转</el-divider>
      <el-steps :active="getStepActive(detailData.order?.status)" finish-status="success" align-center>
        <el-step title="草稿" />
        <el-step title="已提交" />
        <el-step title="已审批" />
        <el-step title="在途" />
        <el-step title="已完成" />
      </el-steps>
      <div class="action-buttons" v-if="detailData.order">
        <el-button type="success" v-if="detailData.order.status === 'SUBMITTED' && hasPermission('wms:transfer:approve')" @click="handleApprove(detailData.order)">审批</el-button>
        <el-button type="primary" v-if="detailData.order.status === 'APPROVED' && hasPermission('wms:transfer:start')" @click="handleStart(detailData.order)">开始调拨</el-button>
        <el-button type="warning" v-if="detailData.order.status === 'IN_TRANSIT' && hasPermission('wms:transfer:ship')" @click="handleShip(detailData.order)">发货</el-button>
        <el-button type="warning" v-if="detailData.order.status === 'IN_TRANSIT' && hasPermission('wms:transfer:receive')" @click="handleReceive(detailData.order)">收货</el-button>
        <el-button type="success" v-if="detailData.order.status === 'IN_TRANSIT' && hasPermission('wms:transfer:complete')" @click="handleComplete(detailData.order)">完成</el-button>
        <el-button type="danger" v-if="['DRAFT', 'SUBMITTED'].includes(detailData.order.status) && hasPermission('wms:transfer:cancel')" @click="handleCancel(detailData.order)">取消</el-button>
        <el-button type="info" @click="handleViewTrace(detailData.order)">跟踪记录</el-button>
      </div>
    </el-dialog>

    <!-- 审批对话框 -->
    <el-dialog v-model="approveDialogVisible" title="审批调拨单" width="400px">
      <el-form :model="approveForm" label-width="80px">
        <el-form-item label="审批结果">
          <el-radio-group v-model="approveForm.approved">
            <el-radio :label="true">通过</el-radio>
            <el-radio :label="false">拒绝</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="审批意见">
          <el-input v-model="approveForm.comment" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="approveDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleApproveSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 发货/收货对话框 -->
    <el-dialog v-model="operatorDialogVisible" :title="operatorDialogTitle" width="400px">
      <el-form :model="operatorForm" label-width="80px">
        <el-form-item label="操作人员ID">
          <el-input-number v-model="operatorForm.operator_id" :min="1" />
        </el-form-item>
        <el-form-item label="操作人姓名">
          <el-input v-model="operatorForm.operator_name" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="operatorDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleOperatorSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 取消对话框 -->
    <el-dialog v-model="cancelDialogVisible" title="取消调拨" width="400px">
      <el-form :model="cancelForm" label-width="80px">
        <el-form-item label="取消原因">
          <el-input v-model="cancelForm.reason" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="cancelDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleCancelSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 跟踪记录对话框 -->
    <el-dialog v-model="traceDialogVisible" title="调拨跟踪记录" width="700px">
      <el-timeline>
        <el-timeline-item v-for="trace in traceData" :key="trace.id" :timestamp="trace.operate_time" placement="top">
          <p><strong>{{ trace.action }}</strong> - {{ trace.operator_name }}</p>
          <p v-if="trace.remark">备注: {{ trace.remark }}</p>
        </el-timeline-item>
      </el-timeline>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import {
  getTransferOrderList,
  createTransferOrder,
  updateTransferOrder,
  deleteTransferOrder,
  getTransferOrderById,
  submitTransferOrder,
  approveTransferOrder,
  startTransferOrder,
  shipTransferOrder,
  receiveTransferOrder,
  completeTransferOrder,
  cancelTransferOrder,
  getTransferOrderTrace
} from '@/api/wms'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const detailVisible = ref(false)
const approveDialogVisible = ref(false)
const operatorDialogVisible = ref(false)
const cancelDialogVisible = ref(false)
const traceDialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({
  query: '',
  transfer_type: '',
  status: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const detailData = reactive<any>({ order: {}, items: [] })
const traceData = ref<any[]>([])

interface TransferItem {
  material_id: number
  material_code: string
  material_name: string
  specification: string
  unit: string
  request_qty: number
  batch_no: string
}

const formData = reactive<any>({
  id: 0,
  transfer_type: 'TRANSFER',
  from_warehouse_id: 0,
  from_warehouse_name: '',
  to_warehouse_id: 0,
  to_warehouse_name: '',
  transfer_reason: '',
  remark: '',
  items: [] as TransferItem[]
})

const approveForm = reactive({ approved: true, comment: '' })
const operatorForm = reactive({ operator_id: 0, operator_name: '' })
const cancelForm = reactive({ reason: '' })
let currentOrder = ref<any>(null)

const rules: FormRules = {
  transfer_type: [{ required: true, message: '请选择调拨类型', trigger: 'change' }],
  from_warehouse_id: [{ required: true, message: '请输入源仓库ID', trigger: 'blur' }],
  to_warehouse_id: [{ required: true, message: '请输入目标仓库ID', trigger: 'blur' }],
  transfer_reason: [{ required: true, message: '请选择调拨原因', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑调拨单' : '新增调拨单')
const operatorDialogTitle = computed(() => currentOrder.value?.action === 'ship' ? '发货确认' : '收货确认')

const getTransferTypeText = (type: string) => {
  const map: Record<string, string> = {
    TRANSFER: '调拨', ADJUSTMENT: '调整', TRANSFER_IN: '调入', TRANSFER_OUT: '调出'
  }
  return map[type] || type
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: '草稿', SUBMITTED: '已提交', APPROVED: '已审批',
    IN_TRANSIT: '在途', COMPLETED: '已完成', CANCELLED: '已取消', REJECTED: '已拒绝'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: 'info', SUBMITTED: 'warning', APPROVED: 'primary',
    IN_TRANSIT: 'warning', COMPLETED: 'success', CANCELLED: 'danger', REJECTED: 'danger'
  }
  return map[status] || 'info'
}

const getStepActive = (status: string) => {
  const map: Record<string, number> = {
    DRAFT: 0, SUBMITTED: 1, APPROVED: 2, IN_TRANSIT: 3, COMPLETED: 4
  }
  return map[status] ?? 0
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getTransferOrderList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.query = ''
  searchForm.transfer_type = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0, transfer_type: 'TRANSFER', from_warehouse_id: 0, from_warehouse_name: '',
    to_warehouse_id: 0, to_warehouse_name: '', transfer_reason: '', remark: '', items: []
  })
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  try {
    const res = await getTransferOrderById(row.id)
    const data = res.data
    Object.assign(formData, {
      id: data.order.id,
      transfer_type: data.order.transfer_type,
      from_warehouse_id: data.order.from_warehouse_id,
      from_warehouse_name: data.order.from_warehouse_name,
      to_warehouse_id: data.order.to_warehouse_id,
      to_warehouse_name: data.order.to_warehouse_name,
      transfer_reason: data.order.transfer_reason,
      remark: data.order.remark || '',
      items: (data.items || []).map((item: any) => ({
        material_id: item.material_id,
        material_code: item.material_code,
        material_name: item.material_name,
        specification: item.specification || '',
        unit: item.unit || '',
        request_qty: item.request_qty,
        batch_no: item.batch_no || ''
      }))
    })
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取调拨单详情失败')
  }
}

const handleView = async (row: any) => {
  try {
    const res = await getTransferOrderById(row.id)
    Object.assign(detailData, res.data)
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取调拨单详情失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该调拨单吗？', '提示', { type: 'warning' })
    await deleteTransferOrder(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    ElMessage.error('删除失败')
  }
}

const handleSubmit = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定提交该调拨单吗？', '提示', { type: 'warning' })
    await submitTransferOrder(row.id)
    ElMessage.success('提交成功')
    loadData()
  } catch (error) {
    ElMessage.error('提交失败')
  }
}

const handleApprove = (row: any) => {
  currentOrder.value = { ...row, action: 'approve' }
  approveForm.approved = true
  approveForm.comment = ''
  approveDialogVisible.value = true
}

const handleApproveSubmit = async () => {
  try {
    await approveTransferOrder(currentOrder.value.id, approveForm)
    ElMessage.success('审批成功')
    approveDialogVisible.value = false
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('审批失败')
  }
}

const handleStart = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定开始调拨吗？', '提示', { type: 'warning' })
    await startTransferOrder(row.id)
    ElMessage.success('操作成功')
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

const handleShip = (row: any) => {
  currentOrder.value = { ...row, action: 'ship' }
  operatorForm.operator_id = 0
  operatorForm.operator_name = ''
  operatorDialogVisible.value = true
}

const handleReceive = (row: any) => {
  currentOrder.value = { ...row, action: 'receive' }
  operatorForm.operator_id = 0
  operatorForm.operator_name = ''
  operatorDialogVisible.value = true
}

const handleOperatorSubmit = async () => {
  try {
    if (currentOrder.value.action === 'ship') {
      await shipTransferOrder(currentOrder.value.id, operatorForm)
    } else {
      await receiveTransferOrder(currentOrder.value.id, operatorForm)
    }
    ElMessage.success('操作成功')
    operatorDialogVisible.value = false
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

const handleComplete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定完成调拨吗？', '提示', { type: 'warning' })
    await completeTransferOrder(row.id)
    ElMessage.success('操作成功')
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

const handleCancel = (row: any) => {
  currentOrder.value = row
  cancelForm.reason = ''
  cancelDialogVisible.value = true
}

const handleCancelSubmit = async () => {
  try {
    await cancelTransferOrder(currentOrder.value.id, cancelForm.reason)
    ElMessage.success('取消成功')
    cancelDialogVisible.value = false
    detailVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('取消失败')
  }
}

const handleViewTrace = async (row: any) => {
  try {
    const res = await getTransferOrderTrace(row.id)
    traceData.value = res.data.traces || []
    traceDialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取跟踪记录失败')
  }
}

const addItem = () => {
  formData.items.push({ material_id: 0, material_code: '', material_name: '', specification: '', unit: '', request_qty: 0, batch_no: '' })
}

const removeItem = (index: number) => {
  formData.items.splice(index, 1)
}

const handleSubmitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (formData.id) {
      await updateTransferOrder(formData.id, formData)
    } else {
      await createTransferOrder(formData)
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

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.transfer-order-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .item-row {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    align-items: center;
    margin-bottom: 12px;
    padding: 12px;
    background: #f5f7fa;
    border-radius: 4px;
  }
  .action-buttons {
    margin-top: 20px;
    display: flex;
    gap: 10px;
    justify-content: center;
  }
}
</style>
