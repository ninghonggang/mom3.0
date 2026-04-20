<template>
  <div class="rfq-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="询价单号">
          <el-input v-model="searchForm.rfq_no" placeholder="请输入询价单号" clearable />
        </el-form-item>
        <el-form-item label="询价名称">
          <el-input v-model="searchForm.rfq_name" placeholder="请输入询价名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" value="DRAFT" />
            <el-option label="已发布" value="PUBLISHED" />
            <el-option label="已截止" value="CLOSED" />
            <el-option label="已授标" value="AWARDED" />
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
      <el-button type="primary" v-if="hasPermission('scp:rfq:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新建询价单
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="rfq_no" label="询价单号" width="150" />
        <el-table-column prop="rfq_name" label="询价名称" min-width="180" />
        <el-table-column prop="rfq_type" label="类型" width="100">
          <template #default="{ row }">
            <el-tag size="small">{{ getTypeText(row.rfq_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="inquiry_date" label="询价日期" width="120" />
        <el-table-column prop="deadline_date" label="报价截止日期" width="120" />
        <el-table-column prop="invited_suppliers" label="邀请供应商" width="100" align="center" />
        <el-table-column prop="quoted_suppliers" label="已报价" width="90" align="center" />
        <el-table-column prop="currency" label="币种" width="70" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_by" label="创建人" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="160" />
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('scp:rfq:edit') && row.status === 'DRAFT'" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('scp:rfq:publish') && row.status === 'DRAFT'" @click="handlePublish(row)">发布</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('scp:rfq:close') && row.status === 'PUBLISHED'" @click="handleClose(row)">截止</el-button>
            <el-button link type="info" size="small" v-if="hasPermission('scp:rfq:viewquotes') && row.status !== 'DRAFT'" @click="handleViewQuotes(row)">查看报价</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('scp:rfq:delete') && row.status === 'DRAFT'" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新建/编辑Dialog -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="800px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="110px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="询价单号" prop="rfq_no">
              <el-input v-model="formData.rfq_no" :disabled="!!formData.id" placeholder="系统自动生成" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="询价名称" prop="rfq_name">
              <el-input v-model="formData.rfq_name" placeholder="请输入询价名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="询价类型" prop="rfq_type">
              <el-select v-model="formData.rfq_type" style="width:100%">
                <el-option label="标准询价" value="STANDARD" />
                <el-option label="快速询价" value="QUICK" />
                <el-option label="年度询价" value="ANNUAL" />
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
            <el-form-item label="报价截止日期" prop="deadline_date">
              <el-date-picker v-model="formData.deadline_date" type="date" value-format="YYYY-MM-DD" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种">
              <el-select v-model="formData.currency" style="width:100%">
                <el-option label="人民币" value="CNY" />
                <el-option label="美元" value="USD" />
                <el-option label="欧元" value="EUR" />
                <el-option label="日元" value="JPY" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款条款">
              <el-input v-model="formData.payment_terms" placeholder="如: 30%预付,70%到付" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="交货条款">
              <el-input v-model="formData.delivery_terms" placeholder="如: DAP, EXW, FOB" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入备注信息" />
        </el-form-item>

        <!-- 询价物料明细 -->
        <el-divider content-position="left">询价物料</el-divider>
        <div class="items-toolbar">
          <el-button size="small" type="primary" @click="handleAddItem">
            <el-icon><Plus /></el-icon>添加物料
          </el-button>
        </div>
        <el-table :data="formData.items" border size="small" class="items-table">
          <el-table-column label="行号" width="60" align="center">
            <template #default="{ $index }">{{ $index + 1 }}</template>
          </el-table-column>
          <el-table-column label="物料编码" width="150">
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
              <el-input v-model="row.specification" placeholder="规格型号" />
            </template>
          </el-table-column>
          <el-table-column label="单位" width="70">
            <template #default="{ row }">
              <el-input v-model="row.unit" placeholder="单位" />
            </template>
          </el-table-column>
          <el-table-column label="询价数量" width="100">
            <template #default="{ row }">
              <el-input-number v-model="row.quantity" :min="0" controls-position="right" style="width:100%" />
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

        <!-- 邀请供应商 -->
        <el-divider content-position="left">邀请供应商</el-divider>
        <div class="items-toolbar">
          <el-button size="small" type="primary" @click="handleAddSupplier">
            <el-icon><Plus /></el-icon>添加供应商
          </el-button>
        </div>
        <el-table :data="formData.suppliers" border size="small" class="items-table">
          <el-table-column label="供应商编码" width="140">
            <template #default="{ row }">
              <el-input v-model="row.supplier_code" placeholder="供应商编码" />
            </template>
          </el-table-column>
          <el-table-column label="供应商名称" width="200">
            <template #default="{ row }">
              <el-input v-model="row.supplier_name" placeholder="供应商名称" />
            </template>
          </el-table-column>
          <el-table-column label="联系人">
            <template #default="{ row }">
              <el-input v-model="row.contact_person" placeholder="联系人" />
            </template>
          </el-table-column>
          <el-table-column label="邮箱" width="180">
            <template #default="{ row }">
              <el-input v-model="row.contact_email" placeholder="邮箱" />
            </template>
          </el-table-column>
          <el-table-column label="电话" width="130">
            <template #default="{ row }">
              <el-input v-model="row.contact_phone" placeholder="电话" />
            </template>
          </el-table-column>
          <el-table-column label="操作" width="60" align="center">
            <template #default="{ $index }">
              <el-button link type="danger" size="small" @click="handleRemoveSupplier($index)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 报价查看Dialog -->
    <el-dialog v-model="quotesDialogVisible" title="供应商报价" width="900px">
      <el-table v-loading="quotesLoading" :data="quotesData">
        <el-table-column prop="quote_no" label="报价单号" width="150" />
        <el-table-column prop="supplier_name" label="供应商名称" min-width="160" />
        <el-table-column prop="contact_person" label="联系人" width="100" />
        <el-table-column prop="contact_phone" label="电话" width="120" />
        <el-table-column prop="quote_date" label="报价日期" width="120" />
        <el-table-column prop="total_amount" label="报价金额" width="120" align="right">
          <template #default="{ row }">
            {{ row.total_amount ? Number(row.total_amount).toFixed(2) : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag size="small" :type="row.status === 'SUBMITTED' ? 'success' : 'info'">
              {{ row.status === 'SUBMITTED' ? '已提交' : row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="currentRFQStatus === 'CLOSED' && hasPermission('scp:rfq:award')" @click="handleAward(row)">授标</el-button>
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <el-button @click="quotesDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { getRFQList, createRFQ, updateRFQ, deleteRFQ, publishRFQ, closeRFQ, getRFQQuotes, awardRFQ } from '@/api/scp'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const quotesDialogVisible = ref(false)
const submitLoading = ref(false)
const quotesLoading = ref(false)
const formRef = ref<FormInstance>()
const currentRFQStatus = ref<string>('')

const searchForm = reactive({ rfq_no: '', rfq_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const quotesData = ref<any[]>([])

const formData = reactive<any>({
  id: 0, rfq_no: '', rfq_name: '', rfq_type: 'STANDARD',
  inquiry_date: '', deadline_date: '', currency: 'CNY',
  payment_terms: '', delivery_terms: '', remark: '',
  items: [] as any[], suppliers: [] as any[]
})

const rules: FormRules = {
  rfq_name: [{ required: true, message: '请输入询价名称', trigger: 'blur' }],
  inquiry_date: [{ required: true, message: '请选择询价日期', trigger: 'change' }],
  deadline_date: [{ required: true, message: '请选择报价截止日期', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑询价单' : '新建询价单')

const getStatusText = (status: string) => {
  const map: Record<string, string> = { DRAFT: '草稿', PUBLISHED: '已发布', CLOSED: '已截止', AWARDED: '已授标', CANCELLED: '已取消' }
  return map[status] || '未知'
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = { DRAFT: 'info', PUBLISHED: 'primary', CLOSED: 'warning', AWARDED: 'success', CANCELLED: 'danger' }
  return map[status] || 'info'
}

const getTypeText = (type: string) => {
  const map: Record<string, string> = { STANDARD: '标准', QUICK: '快速', ANNUAL: '年度' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getRFQList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.rfq_no = ''; searchForm.rfq_name = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, {
    id: 0, rfq_no: '', rfq_name: '', rfq_type: 'STANDARD',
    inquiry_date: '', deadline_date: '', currency: 'CNY',
    payment_terms: '', delivery_terms: '', remark: '',
    items: [], suppliers: []
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, { ...row, items: row.items || [], suppliers: row.suppliers || [] })
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该询价单吗？', '提示', { type: 'warning' })
    await deleteRFQ(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}

const handlePublish = async (row: any) => {
  try {
    await ElMessageBox.confirm('发布后将被邀请供应商看到，确定发布吗？', '提示', { type: 'warning' })
    await publishRFQ(row.id)
    ElMessage.success('发布成功')
    loadData()
  } catch {}
}

const handleClose = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定截止该询价吗？截止后将无法再接收报价。', '提示', { type: 'warning' })
    await closeRFQ(row.id)
    ElMessage.success('已截止')
    loadData()
  } catch {}
}

const handleViewQuotes = async (row: any) => {
  currentRFQStatus.value = row.status
  quotesDialogVisible.value = true
  quotesLoading.value = true
  try {
    const res = await getRFQQuotes(row.id)
    quotesData.value = res.data || []
  } finally {
    quotesLoading.value = false
  }
}

const handleAward = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定授标给供应商「${row.supplier_name}」吗？`, '确认授标', { type: 'warning' })
    const currentRFQ = tableData.value.find(t => t.id === (quotesData.value[0] as any)?.rfq_id)
    if (currentRFQ) {
      await awardRFQ(currentRFQ.id, { supplier_id: row.supplier_id, quote_id: row.id })
      ElMessage.success('授标成功')
      quotesDialogVisible.value = false
      loadData()
    }
  } catch {}
}

const handleAddItem = () => {
  formData.items.push({ material_code: '', material_name: '', specification: '', unit: 'PCS', quantity: 1, target_price: 0 })
}

const handleRemoveItem = (index: number) => {
  formData.items.splice(index, 1)
}

const handleAddSupplier = () => {
  formData.suppliers.push({ supplier_code: '', supplier_name: '', contact_person: '', contact_email: '', contact_phone: '' })
}

const handleRemoveSupplier = (index: number) => {
  formData.suppliers.splice(index, 1)
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (formData.id) {
      await updateRFQ(formData.id, formData)
    } else {
      await createRFQ(formData)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.rfq-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .items-toolbar { margin-bottom: 8px; }
  .items-table { margin-bottom: 12px; }
}
</style>
