<template>
  <div class="stock-check-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="盘点单号">
          <el-input v-model="searchForm.query" placeholder="请输入盘点单号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="草稿" value="DRAFT" />
            <el-option label="已提交" value="IN_PROGRESS" />
            <el-option label="盘点中" value="COUNTING" />
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
      <el-button type="primary" v-if="hasPermission('wms:stockcheck:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增盘点单
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="check_no" label="盘点单号" width="150" />
        <el-table-column prop="check_type" label="盘点类型" width="100">
          <template #default="{ row }">
            <el-tag :type="getCheckTypeTag(row.check_type)">{{ getCheckTypeText(row.check_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="warehouse_name" label="仓库" min-width="120" />
        <el-table-column prop="plan_start_date" label="计划开始" width="110" />
        <el-table-column prop="plan_end_date" label="计划结束" width="110" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="counted_locations" label="已盘库位" width="90" />
        <el-table-column prop="total_locations" label="总库位" width="90" />
        <el-table-column prop="variance_count" label="差异数" width="80" />
        <el-table-column prop="variance_rate" label="差异率" width="80">
          <template #default="{ row }">{{ row.variance_rate ? row.variance_rate + '%' : '-' }}</template>
        </el-table-column>
        <el-table-column prop="approval_status" label="审核状态" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.approval_status" :type="getApprovalType(row.approval_status)">{{ row.approval_status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('wms:stockcheck:view')" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" v-if="hasPermission('wms:stockcheck:edit') && row.status === 'DRAFT'" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" v-if="hasPermission('wms:stockcheck:submit') && row.status === 'DRAFT'" @click="handleSubmit(row)">提交</el-button>
            <el-button link type="warning" v-if="hasPermission('wms:stockcheck:start') && row.status === 'IN_PROGRESS'" @click="handleStart(row)">开始盘点</el-button>
            <el-button link type="success" v-if="hasPermission('wms:stockcheck:complete') && row.status === 'COUNTING'" @click="handleComplete(row)">完成盘点</el-button>
            <el-button link type="danger" v-if="hasPermission('wms:stockcheck:delete') && row.status === 'DRAFT'" @click="handleDelete(row)">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px" @close="handleDialogClose">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="仓库" prop="warehouse_id">
          <el-select v-model="formData.warehouse_id" placeholder="请选择仓库" @change="handleWarehouseChange">
            <el-option v-for="wh in warehouseList" :key="wh.id" :label="wh.warehouse_name" :value="wh.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="盘点类型" prop="check_type">
          <el-select v-model="formData.check_type" placeholder="请选择盘点类型">
            <el-option label="全盘" value="FULL" />
            <el-option label="抽盘" value="CYCLE" />
            <el-option label="盲盘" value="BLIND" />
            <el-option label="调整盘" value="ADJUSTMENT" />
          </el-select>
        </el-form-item>
        <el-form-item label="计划开始日期" prop="plan_start_date">
          <el-date-picker v-model="formData.plan_start_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="计划结束日期" prop="plan_end_date">
          <el-date-picker v-model="formData.plan_end_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="包含零库存">
          <el-switch v-model="formData.include_zero_stock" :active-value="1" :inactive-value="0" />
        </el-form-item>
        <el-form-item label="包含过期物料">
          <el-switch v-model="formData.include_expired_stock" :active-value="1" :inactive-value="0" />
        </el-form-item>
        <el-form-item label="盲盘模式">
          <el-switch v-model="formData.is_blind_mode" :active-value="1" :inactive-value="0" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmitForm">确定</el-button>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="盘点单详情" width="1000px">
      <el-descriptions :column="3" border>
        <el-descriptions-item label="盘点单号">{{ detailData.check_no }}</el-descriptions-item>
        <el-descriptions-item label="盘点类型">{{ getCheckTypeText(detailData.check_type) }}</el-descriptions-item>
        <el-descriptions-item label="仓库">{{ detailData.warehouse_name }}</el-descriptions-item>
        <el-descriptions-item label="计划开始">{{ detailData.plan_start_date }}</el-descriptions-item>
        <el-descriptions-item label="计划结束">{{ detailData.plan_end_date }}</el-descriptions-item>
        <el-descriptions-item label="实际开始">{{ detailData.actual_start_date || '-' }}</el-descriptions-item>
        <el-descriptions-item label="实际结束">{{ detailData.actual_end_date || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(detailData.status)">{{ getStatusText(detailData.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="审核状态">
          <el-tag v-if="detailData.approval_status" :type="getApprovalType(detailData.approval_status)">{{ detailData.approval_status }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="已盘库位">{{ detailData.counted_locations }} / {{ detailData.total_locations }}</el-descriptions-item>
        <el-descriptions-item label="差异数">{{ detailData.variance_count }}</el-descriptions-item>
        <el-descriptions-item label="差异率">{{ detailData.variance_rate ? detailData.variance_rate + '%' : '-' }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="3">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>盘点明细</el-divider>

      <div class="detail-toolbar">
        <el-button type="primary" size="small" v-if="detailData.status === 'COUNTING'" @click="handleCountAll">批量录入</el-button>
        <el-button type="success" size="small" v-if="hasPermission('wms:stockcheck:approve') && detailData.status === 'COMPLETED' && detailData.approval_status === 'PENDING'" @click="handleApproveDialog(detailData)">审核差异</el-button>
      </div>

      <el-table :data="detailData.items" border max-height="400">
        <el-table-column prop="line_no" label="行号" width="60" />
        <el-table-column prop="location_name" label="库位" width="120" />
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="batch_no" label="批次" width="100" />
        <el-table-column prop="system_qty" label="系统数量" width="100" />
        <el-table-column prop="counted_qty" label="盘点数量" width="100">
          <template #default="{ row }">
            <span v-if="detailData.is_blind_mode && detailData.status !== 'COMPLETED' && row.count_status === 'PENDING'">***</span>
            <span v-else>{{ row.counted_qty || '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="variance_qty" label="差异数量" width="100">
          <template #default="{ row }">
            <span :class="row.variance_qty && row.variance_qty !== 0 ? 'variance-text' : ''">{{ row.variance_qty || '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="count_status" label="盘点状态" width="100">
          <template #default="{ row }">
            <el-tag size="small" :type="getCountStatusType(row.count_status)">{{ row.count_status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="handle_status" label="处理状态" width="100">
          <template #default="{ row }">
            <el-tag size="small" :type="getHandleStatusType(row.handle_status)">{{ row.handle_status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('wms:stockcheck:count') && detailData.status === 'COUNTING' && row.count_status === 'PENDING'" @click="handleCountItem(row)">录入</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('wms:stockcheck:recount') && detailData.status === 'COUNTING' && row.count_status === 'COUNTED'" @click="handleRecount(row)">复盘</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('wms:stockcheck:handle') && row.variance_qty !== 0 && row.variance_qty && row.handle_status === 'PENDING'" @click="handleVariance(row)">处理</el-button>
          </template>
        </el-table-column>
      </el-table>

      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>

    <!-- 盘点数量录入对话框 -->
    <el-dialog v-model="countDialogVisible" title="录入盘点数量" width="400px">
      <el-form ref="countFormRef" :model="countForm" label-width="100px">
        <el-form-item label="物料编码">{{ countForm.material_code }}</el-form-item>
        <el-form-item label="物料名称">{{ countForm.material_name }}</el-form-item>
        <el-form-item label="库位">{{ countForm.location_name }}</el-form-item>
        <el-form-item label="系统数量">{{ countForm.system_qty }}</el-form-item>
        <el-form-item label="盘点数量" prop="counted_qty">
          <el-input-number v-model="countForm.counted_qty" :min="0" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="countForm.remark" type="textarea" :rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="countDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleCountSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 复盘对话框 -->
    <el-dialog v-model="recountDialogVisible" title="复盘" width="400px">
      <el-form ref="recountFormRef" :model="recountForm" label-width="100px">
        <el-form-item label="物料编码">{{ recountForm.material_code }}</el-form-item>
        <el-form-item label="原盘点数量">{{ recountForm.counted_qty }}</el-form-item>
        <el-form-item label="复盘数量" prop="recount_qty">
          <el-input-number v-model="recountForm.recount_qty" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="recountDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleRecountSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 差异处理对话框 -->
    <el-dialog v-model="varianceDialogVisible" title="处理差异" width="500px">
      <el-form ref="varianceFormRef" :model="varianceForm" label-width="100px">
        <el-form-item label="物料编码">{{ varianceForm.material_code }}</el-form-item>
        <el-form-item label="物料名称">{{ varianceForm.material_name }}</el-form-item>
        <el-form-item label="系统数量">{{ varianceForm.system_qty }}</el-form-item>
        <el-form-item label="盘点数量">{{ varianceForm.counted_qty }}</el-form-item>
        <el-form-item label="差异数量">
          <span class="variance-text">{{ varianceForm.variance_qty }}</span>
        </el-form-item>
        <el-form-item label="处理方式" prop="handle_method">
          <el-select v-model="varianceForm.handle_method" placeholder="请选择处理方式">
            <el-option label="调整库存" value="ADJUST" />
            <el-option label="报损" value="WRITE_OFF" />
            <el-option label="报溢" value="WRITE_IN" />
          </el-select>
        </el-form-item>
        <el-form-item label="处理数量" prop="handle_qty">
          <el-input-number v-model="varianceForm.handle_qty" :min="0" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="varianceForm.remark" type="textarea" :rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="varianceDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleVarianceSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 审核对话框 -->
    <el-dialog v-model="approveDialogVisible" title="审核盘点差异" width="500px">
      <el-form ref="approveFormRef" :model="approveForm" label-width="100px">
        <el-form-item label="盘点单号">{{ approveForm.check_no }}</el-form-item>
        <el-form-item label="差异数量">{{ approveForm.variance_count }} 条</el-form-item>
        <el-form-item label="审核意见">
          <el-input v-model="approveForm.comment" type="textarea" :rows="3" placeholder="请输入审核意见" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="approveDialogVisible = false">取消</el-button>
        <el-button type="success" :loading="submitLoading" @click="handleApprove(true)">通过</el-button>
        <el-button type="danger" :loading="submitLoading" @click="handleApprove(false)">驳回</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import {
  getStockCheckList,
  getStockCheckById,
  createStockCheck,
  updateStockCheck,
  deleteStockCheck,
  submitStockCheck,
  startStockCheck,
  completeStockCheck,
  approveStockCheck,
  countStockCheckItem,
  handleStockCheckVariance,
  recountStockCheckItem,
  getStockCheckVariance,
  getWarehouseList
} from '@/api/wms'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const detailVisible = ref(false)
const countDialogVisible = ref(false)
const recountDialogVisible = ref(false)
const varianceDialogVisible = ref(false)
const approveDialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const countFormRef = ref<FormInstance>()
const recountFormRef = ref<FormInstance>()
const varianceFormRef = ref<FormInstance>()
const approveFormRef = ref<FormInstance>()

const warehouseList = ref<any[]>([])

const searchForm = reactive({ query: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  id: 0,
  warehouse_id: undefined,
  warehouse_name: '',
  check_type: 'FULL',
  plan_start_date: '',
  plan_end_date: '',
  include_zero_stock: 1,
  include_expired_stock: 0,
  is_blind_mode: 0,
  remark: ''
})

const detailData = reactive<any>({ items: [] })

const countForm = reactive<any>({
  id: 0,
  material_code: '',
  material_name: '',
  location_name: '',
  system_qty: 0,
  counted_qty: 0,
  remark: ''
})

const recountForm = reactive<any>({
  id: 0,
  material_code: '',
  counted_qty: 0,
  recount_qty: 0
})

const varianceForm = reactive<any>({
  id: 0,
  material_code: '',
  material_name: '',
  system_qty: 0,
  counted_qty: 0,
  variance_qty: 0,
  handle_method: 'ADJUST',
  handle_qty: 0,
  remark: ''
})

const approveForm = reactive<any>({
  id: 0,
  check_no: '',
  variance_count: 0,
  comment: ''
})

const rules: FormRules = {
  warehouse_id: [{ required: true, message: '请选择仓库', trigger: 'change' }],
  check_type: [{ required: true, message: '请选择盘点类型', trigger: 'change' }],
  plan_start_date: [{ required: true, message: '请选择计划开始日期', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑盘点单' : '新增盘点单')

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: '草稿',
    IN_PROGRESS: '已提交',
    COUNTING: '盘点中',
    COMPLETED: '已完成',
    CANCELLED: '已取消'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    DRAFT: 'info',
    IN_PROGRESS: 'warning',
    COUNTING: 'primary',
    COMPLETED: 'success',
    CANCELLED: 'danger'
  }
  return map[status] || 'info'
}

const getCheckTypeText = (type: string) => {
  const map: Record<string, string> = {
    FULL: '全盘',
    CYCLE: '抽盘',
    BLIND: '盲盘',
    ADJUSTMENT: '调整盘'
  }
  return map[type] || type
}

const getCheckTypeTag = (type: string) => {
  const map: Record<string, string> = {
    FULL: 'success',
    CYCLE: 'warning',
    BLIND: 'info',
    ADJUSTMENT: 'primary'
  }
  return map[type] || 'info'
}

const getApprovalType = (status: string) => {
  const map: Record<string, string> = {
    PENDING: 'warning',
    APPROVED: 'success',
    REJECTED: 'danger'
  }
  return map[status] || 'info'
}

const getCountStatusType = (status: string) => {
  const map: Record<string, string> = {
    PENDING: 'info',
    COUNTED: 'success',
    CONFIRMED: 'primary'
  }
  return map[status] || 'info'
}

const getHandleStatusType = (status: string) => {
  const map: Record<string, string> = {
    PENDING: 'info',
    APPROVED: 'warning',
    PROCESSED: 'success'
  }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getStockCheckList({ query: searchForm.query, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadWarehouses = async () => {
  try {
    const res = await getWarehouseList()
    warehouseList.value = res.data.list || []
  } catch (error) {
    console.error('加载仓库列表失败', error)
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.query = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, {
    id: 0, warehouse_id: undefined, warehouse_name: '', check_type: 'FULL',
    plan_start_date: '', plan_end_date: '', include_zero_stock: 1,
    include_expired_stock: 0, is_blind_mode: 0, remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  try {
    const res = await getStockCheckById(row.id)
    const data = res.data
    Object.assign(formData, {
      id: data.check.id,
      warehouse_id: data.check.warehouse_id,
      warehouse_name: data.check.warehouse_name,
      check_type: data.check.check_type,
      plan_start_date: data.check.plan_start_date,
      plan_end_date: data.check.plan_end_date,
      include_zero_stock: data.check.include_zero_stock,
      include_expired_stock: data.check.include_expired_stock,
      is_blind_mode: data.check.is_blind_mode,
      remark: data.check.remark || ''
    })
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取盘点单详情失败')
  }
}

const handleView = async (row: any) => {
  try {
    const res = await getStockCheckById(row.id)
    Object.assign(detailData, res.data)
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取盘点单详情失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该盘点单吗？', '提示', { type: 'warning' })
    await deleteStockCheck(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    if (error !== 'cancel') ElMessage.error('删除失败')
  }
}

const handleSubmit = async (row: any) => {
  try {
    await ElMessageBox.confirm('提交后盘点单将进入执行阶段，确定提交吗？', '提示', { type: 'warning' })
    await submitStockCheck(row.id)
    ElMessage.success('提交成功')
    loadData()
  } catch (error) {
    if (error !== 'cancel') ElMessage.error('提交失败')
  }
}

const handleStart = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定开始盘点吗？', '提示', { type: 'warning' })
    await startStockCheck(row.id)
    ElMessage.success('已开始盘点')
    loadData()
  } catch (error) {
    if (error !== 'cancel') ElMessage.error('操作失败')
  }
}

const handleComplete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定完成盘点吗？', '提示', { type: 'warning' })
    await completeStockCheck(row.id)
    ElMessage.success('盘点已完成')
    loadData()
  } catch (error) {
    if (error !== 'cancel') ElMessage.error('操作失败')
  }
}

const handleWarehouseChange = (warehouseId: number) => {
  const wh = warehouseList.value.find(w => w.id === warehouseId)
  if (wh) {
    formData.warehouse_name = wh.warehouse_name
  }
}

const handleDialogClose = () => {
  formRef.value?.resetFields()
}

const handleSubmitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (formData.id) {
      await updateStockCheck(formData.id, formData)
    } else {
      await createStockCheck(formData)
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

const handleCountItem = (row: any) => {
  Object.assign(countForm, {
    id: row.id,
    material_code: row.material_code,
    material_name: row.material_name,
    location_name: row.location_name,
    system_qty: row.system_qty,
    counted_qty: row.counted_qty || 0,
    remark: ''
  })
  countDialogVisible.value = true
}

const handleCountSubmit = async () => {
  submitLoading.value = true
  try {
    await countStockCheckItem(detailData.check.id, countForm.id, {
      counted_qty: countForm.counted_qty,
      counter_id: 0,
      counter_name: '',
      remark: countForm.remark
    })
    ElMessage.success('录入成功')
    countDialogVisible.value = false
    handleView({ id: detailData.check.id })
    loadData()
  } catch (error) {
    ElMessage.error('录入失败')
  } finally {
    submitLoading.value = false
  }
}

const handleRecount = (row: any) => {
  Object.assign(recountForm, {
    id: row.id,
    material_code: row.material_code,
    counted_qty: row.counted_qty,
    recount_qty: 0
  })
  recountDialogVisible.value = true
}

const handleRecountSubmit = async () => {
  submitLoading.value = true
  try {
    await recountStockCheckItem(detailData.check.id, recountForm.id, {
      recount_qty: recountForm.recount_qty,
      recount_by: 0,
      recount_name: ''
    })
    ElMessage.success('复盘成功')
    recountDialogVisible.value = false
    handleView({ id: detailData.check.id })
    loadData()
  } catch (error) {
    ElMessage.error('复盘失败')
  } finally {
    submitLoading.value = false
  }
}

const handleVariance = (row: any) => {
  Object.assign(varianceForm, {
    id: row.id,
    material_code: row.material_code,
    material_name: row.material_name,
    system_qty: row.system_qty,
    counted_qty: row.counted_qty,
    variance_qty: row.variance_qty,
    handle_method: 'ADJUST',
    handle_qty: Math.abs(row.variance_qty),
    remark: ''
  })
  varianceDialogVisible.value = true
}

const handleVarianceSubmit = async () => {
  submitLoading.value = true
  try {
    await handleStockCheckVariance(detailData.check.id, varianceForm.id, {
      handle_method: varianceForm.handle_method,
      handle_qty: varianceForm.handle_qty,
      handler_id: 0,
      handler_name: '',
      remark: varianceForm.remark
    })
    ElMessage.success('处理成功')
    varianceDialogVisible.value = false
    handleView({ id: detailData.check.id })
    loadData()
  } catch (error) {
    ElMessage.error('处理失败')
  } finally {
    submitLoading.value = false
  }
}

const handleApproveDialog = (row: any) => {
  Object.assign(approveForm, {
    id: row.id,
    check_no: row.check_no,
    variance_count: row.variance_count,
    comment: ''
  })
  approveDialogVisible.value = true
}

const handleApprove = async (approved: boolean) => {
  submitLoading.value = true
  try {
    await approveStockCheck(approveForm.id, {
      approved,
      comment: approveForm.comment
    })
    ElMessage.success(approved ? '审核通过' : '审核驳回')
    approveDialogVisible.value = false
    handleView({ id: approveForm.id })
    loadData()
  } catch (error) {
    ElMessage.error('审核操作失败')
  } finally {
    submitLoading.value = false
  }
}

const handleCountAll = () => {
  ElMessage.info('请在明细中逐条录入盘点数量')
}

onMounted(() => {
  loadData()
  loadWarehouses()
})
</script>

<style scoped lang="scss">
.stock-check-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .variance-text { color: #f56c6c; font-weight: bold; }
  .detail-toolbar { margin-bottom: 12px; }
}
</style>
