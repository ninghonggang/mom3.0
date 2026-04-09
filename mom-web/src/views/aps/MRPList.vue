<template>
  <div class="mrp-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="MRP编号">
          <el-input v-model="searchForm.mrp_no" placeholder="请输入MRP编号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待计算" :value="1" />
            <el-option label="计算中" :value="2" />
            <el-option label="已完成" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:mrp:add')" @click="dialogVisible = true; formData = {}">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('aps:mrp:run')" @click="showRunDialog">
        <el-icon><Grid /></el-icon>MRP计算
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="mrp_no" label="MRP编号" width="140" />
        <el-table-column prop="mrp_type" label="类型" width="80" />
        <el-table-column prop="plan_date" label="计划日期" width="110" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="120" show-overflow-tooltip />
        <el-table-column prop="created_by" label="创建人" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="220" fixed="right">
          <template #default="{ row }">
            <el-button link type="success" size="small" v-if="hasPermission('aps:mrp:run') && row.status === 1" @click="handleRunMRP(row)">MRP计算</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('aps:mrp:shortage') && row.status === 3" @click="handleShortage(row)">缺料分析</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('aps:mrp:results')" @click="handleResults(row)">计算结果</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:mrp:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新增对话框 -->
    <el-dialog v-model="dialogVisible" title="新增MRP" width="500px">
      <el-form :model="formData" label-width="100px">
        <el-form-item label="MRP编号">
          <el-input v-model="formData.mrp_no" placeholder="自动生成" disabled />
        </el-form-item>
        <el-form-item label="类型" required>
          <el-select v-model="formData.mrp_type" placeholder="请选择" style="width: 100%">
            <el-option label="MPS" value="MPS" />
            <el-option label="MRP" value="MRP" />
          </el-select>
        </el-form-item>
        <el-form-item label="计划日期">
          <el-date-picker v-model="formData.plan_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>

    <!-- MRP计算对话框 -->
    <el-dialog v-model="runDialogVisible" title="MRP计算" width="400px">
      <el-form :model="runForm" label-width="100px">
        <el-form-item label="MRP编号">
          <el-input v-model="runForm.mrp_no" placeholder="自动生成" disabled />
        </el-form-item>
        <el-form-item label="计划月份" required>
          <el-date-picker v-model="runForm.plan_month" type="month" value-format="YYYY-MM" placeholder="选择月份" style="width: 100%" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="runDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="runLoading" @click="handleRun">执行MRP计算</el-button>
      </template>
    </el-dialog>

    <!-- MRP计算结果对话框 -->
    <el-dialog v-model="resultsDialogVisible" title="MRP计算结果" width="900px">
      <el-table v-loading="resultsLoading" :data="resultsData" max-height="400">
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="quantity" label="毛需求" width="100" align="right" />
        <el-table-column prop="stock_qty" label="库存数量" width="100" align="right" />
        <el-table-column prop="allocated_qty" label="已分配" width="100" align="right" />
        <el-table-column prop="net_qty" label="净需求" width="100" align="right">
          <template #default="{ row }">
            <span :class="{ 'text-danger': row.net_qty > 0 }">{{ row.net_qty }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="source_type" label="来源" width="80" />
        <el-table-column prop="source_no" label="来源单号" width="120" />
      </el-table>
      <template #footer>
        <el-button @click="resultsDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>

    <!-- 缺料分析对话框 -->
    <el-dialog v-model="shortageDialogVisible" title="缺料分析" width="800px">
      <el-alert v-if="shortageData.length === 0" title="暂无缺料" type="success" :closable="false" />
      <el-table v-else v-loading="shortageLoading" :data="shortageData" max-height="400">
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="quantity" label="需求数量" width="100" align="right" />
        <el-table-column prop="stock_qty" label="库存数量" width="100" align="right" />
        <el-table-column prop="net_qty" label="缺料数量" width="100" align="right">
          <template #default="{ row }">
            <span class="text-danger">{{ row.net_qty }}</span>
          </template>
        </el-table-column>
      </el-table>

      <el-divider v-if="suggestionData.length > 0" content-position="left">采购建议</el-divider>
      <el-table v-if="suggestionData.length > 0" :data="suggestionData" max-height="300">
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="demand_qty" label="需求数量" width="100" align="right" />
        <el-table-column prop="urgent_level" label="紧急程度" width="100">
          <template #default="{ row }">
            <el-tag :type="getUrgentType(row.urgent_level)" size="small">{{ row.urgent_level }}</el-tag>
          </template>
        </el-table-column>
      </el-table>

      <template #footer>
        <el-button @click="shortageDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getMRPList, createMRP, runMRP, getMRPResults, getMRPShortage, getMRPPurchaseSuggestion, deleteMRP } from '@/api/aps'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const saveLoading = ref(false)
const runLoading = ref(false)
const resultsLoading = ref(false)
const shortageLoading = ref(false)

const dialogVisible = ref(false)
const runDialogVisible = ref(false)
const resultsDialogVisible = ref(false)
const shortageDialogVisible = ref(false)

const tableData = ref<any[]>([])
const resultsData = ref<any[]>([])
const shortageData = ref<any[]>([])
const suggestionData = ref<any[]>([])

const formData = ref<any>({})
const runForm = reactive<{ mrp_id?: number; mrp_no?: string; plan_month: string }>({ plan_month: '' })

const searchForm = reactive({ mrp_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待计算', 2: '计算中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }
  return map[status] || 'info'
}

const getUrgentType = (level: string) => {
  const map: Record<string, string> = { URGENT: 'danger', HIGH: 'warning', NORMAL: 'info' }
  return map[level] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMRPList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.mrp_no = ''; searchForm.status = ''; handleSearch() }

const handleSave = async () => {
  if (!formData.value.mrp_type) {
    ElMessage.warning('请选择MRP类型')
    return
  }
  saveLoading.value = true
  try {
    await createMRP(formData.value)
    ElMessage.success('创建成功')
    dialogVisible.value = false
    loadData()
  } catch (error: any) {
    ElMessage.error(error.message || '创建失败')
  } finally {
    saveLoading.value = false
  }
}

const showRunDialog = () => {
  runForm.mrp_no = ''
  runForm.plan_month = ''
  runDialogVisible.value = true
}

const handleRunMRP = async (row: any) => {
  try {
    await ElMessageBox.confirm('将根据MPS计划和BOM展开进行MRP计算，是否继续？', 'MRP计算', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'info'
    })
    runLoading.value = true
    await runMRP({ id: row.id, plan_month: row.plan_month || '' })
    ElMessage.success('MRP计算完成')
    runDialogVisible.value = false
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || 'MRP计算失败')
    }
  } finally {
    runLoading.value = false
  }
}

const handleRun = async () => {
  if (!runForm.plan_month) {
    ElMessage.warning('请选择计划月份')
    return
  }
  try {
    await ElMessageBox.confirm('将根据MPS计划和BOM展开进行MRP计算，是否继续？', 'MRP计算', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'info'
    })
    runLoading.value = true
    await runMRP({ id: runForm.mrp_id || 0, plan_month: runForm.plan_month })
    ElMessage.success('MRP计算完成')
    runDialogVisible.value = false
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || 'MRP计算失败')
    }
  } finally {
    runLoading.value = false
  }
}

const handleResults = async (row: any) => {
  resultsDialogVisible.value = true
  resultsLoading.value = true
  try {
    const res = await getMRPResults(row.id)
    resultsData.value = res.data.list || []
  } catch (error: any) {
    ElMessage.error(error.message || '获取计算结果失败')
  } finally {
    resultsLoading.value = false
  }
}

const handleShortage = async (row: any) => {
  shortageDialogVisible.value = true
  shortageLoading.value = true
  shortageData.value = []
  suggestionData.value = []
  try {
    const [shortageRes, suggestionRes] = await Promise.all([
      getMRPShortage(row.id),
      getMRPPurchaseSuggestion(row.id)
    ])
    shortageData.value = shortageRes.data.list || []
    suggestionData.value = suggestionRes.data.list || []
  } catch (error: any) {
    ElMessage.error(error.message || '获取缺料分析失败')
  } finally {
    shortageLoading.value = false
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定要删除该MRP计划吗？', '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteMRP(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.mrp-list {
  .search-card { margin-bottom: 16px; }
  .toolbar-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .text-danger { color: #f56c6c; font-weight: bold; }
}
</style>
