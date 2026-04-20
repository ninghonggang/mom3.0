<template>
  <div class="inspection-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="模板编号"><el-input v-model="searchForm.template_code" placeholder="请输入" clearable /></el-form-item>
        <el-form-item label="模板类型">
          <el-select v-model="searchForm.template_type" placeholder="请选择" clearable>
            <el-option label="日常点检" value="DAILY" />
            <el-option label="周点检" value="WEEKLY" />
            <el-option label="月点检" value="MONTHLY" />
            <el-option label="专项点检" value="SPECIAL" />
          </el-select>
        </el-form-item>
        <el-form-item><el-button type="primary" @click="handleSearch">查询</el-button><el-button @click="handleReset">重置</el-button></el-form-item>
      </el-form>
    </el-card>
    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('equipment:inspection:add')" @click="handleAdd"><el-icon><Plus /></el-icon>新增标准</el-button>
    </el-card>
    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="template_code" label="模板编号" width="140" />
        <el-table-column prop="template_name" label="模板名称" min-width="180" />
        <el-table-column prop="template_type" label="类型" width="100">
          <template #default="{ row }"><el-tag>{{ getTypeText(row.template_type) }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="frequency_type" label="周期" width="80" />
        <el-table-column prop="estimated_minutes" label="预计耗时(分钟)" width="120" />
        <el-table-column prop="is_active" label="状态" width="80">
          <template #default="{ row }"><el-tag :type="row.is_active === 1 ? 'success' : 'info'">{{ row.is_active === 1 ? '启用' : '停用' }}</el-tag></template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('equipment:inspection:edit')" @click="handleEdit(row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination"><el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" /></div>
    </el-card>

    <!-- 模板详情/编辑Dialog -->
    <el-dialog v-model="templateDialogVisible" :title="templateDialogTitle" width="900px">
      <el-form ref="templateFormRef" :model="templateFormData" :rules="templateRules" label-width="110px">
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="模板编号" prop="template_code">
              <el-input v-model="templateFormData.template_code" :disabled="isViewMode" placeholder="自动生成" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="模板名称" prop="template_name">
              <el-input v-model="templateFormData.template_name" :disabled="isViewMode" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="模板类型" prop="template_type">
              <el-select v-model="templateFormData.template_type" :disabled="isViewMode" placeholder="请选择" style="width:100%">
                <el-option label="日常点检" value="DAILY" />
                <el-option label="周点检" value="WEEKLY" />
                <el-option label="月点检" value="MONTHLY" />
                <el-option label="专项点检" value="SPECIAL" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="周期类型" prop="frequency_type">
              <el-select v-model="templateFormData.frequency_type" :disabled="isViewMode" placeholder="请选择" style="width:100%">
                <el-option label="每日" value="DAILY" />
                <el-option label="每周" value="WEEKLY" />
                <el-option label="每月" value="MONTHLY" />
                <el-option label="一次性" value="ONCE" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="周期值" prop="frequency_value">
              <el-input-number v-model="templateFormData.frequency_value" :min="1" :disabled="isViewMode" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="预计耗时(分钟)" prop="estimated_minutes">
              <el-input-number v-model="templateFormData.estimated_minutes" :min="1" :disabled="isViewMode" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="执行时间">
              <el-input v-model="templateFormData.execution_time" :disabled="isViewMode" placeholder="如: 08:00" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态">
              <el-radio-group v-model="templateFormData.is_active" :disabled="isViewMode">
                <el-radio :value="1">启用</el-radio>
                <el-radio :value="0">停用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="生效日期">
              <el-date-picker v-model="templateFormData.effective_date" type="date" value-format="YYYY-MM-DD" :disabled="isViewMode" placeholder="选择日期" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="失效日期">
              <el-date-picker v-model="templateFormData.expiry_date" type="date" value-format="YYYY-MM-DD" :disabled="isViewMode" placeholder="选择日期" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注">
          <el-input v-model="templateFormData.remark" type="textarea" :rows="2" :disabled="isViewMode" />
        </el-form-item>
      </el-form>

      <!-- 点检项目列表 -->
      <div class="items-section">
        <div class="items-header">
          <span class="items-title">点检项目</span>
          <el-button v-if="!isViewMode" type="primary" size="small" @click="handleAddItem"><el-icon><Plus /></el-icon>添加项目</el-button>
        </div>
        <el-table :data="templateFormData.items" border size="small" max-height="300">
          <el-table-column prop="item_name" label="点检项目名称" min-width="150" show-overflow-tooltip />
          <el-table-column prop="check_standard" label="点检标准" min-width="150" show-overflow-tooltip />
          <el-table-column prop="check_method" label="点检方法" width="100">
            <template #default="{ row }">{{ getMethodText(row.check_method) }}</template>
          </el-table-column>
          <el-table-column prop="result_type" label="判定类型" width="100">
            <template #default="{ row }">{{ getResultTypeText(row.result_type) }}</template>
          </el-table-column>
          <el-table-column label="判定标准" width="180" show-overflow-tooltip>
            <template #default="{ row }">
              <span v-if="row.result_type === 'NUMERIC'">{{ row.lower_limit }} ~ {{ row.upper_limit }} {{ row.unit }}</span>
              <span v-else-if="row.result_type === 'OK_NG'">{{ row.standard_value === 1 ? 'OK/NG' : '-' }}</span>
              <span v-else>-</span>
            </template>
          </el-table-column>
          <el-table-column label="关键点" width="80">
            <template #default="{ row }"><el-tag size="small" :type="row.is_critical_point === 1 ? 'danger' : 'info'">{{ row.is_critical_point === 1 ? '是' : '否' }}</el-tag></template>
          </el-table-column>
          <el-table-column label="必检" width="80">
            <template #default="{ row }"><el-tag size="small" :type="row.is_mandatory === 1 ? 'success' : 'info'">{{ row.is_mandatory === 1 ? '是' : '否' }}</el-tag></template>
          </el-table-column>
          <el-table-column label="操作" width="120" fixed="right">
            <template #default="{ row, $index }">
              <el-button link type="primary" size="small" @click="handleViewItem(row)">查看</el-button>
              <el-button v-if="!isViewMode" link type="primary" size="small" @click="handleEditItem(row)">编辑</el-button>
              <el-button v-if="!isViewMode" link type="danger" size="small" @click="handleDeleteItem($index)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <template #footer>
        <el-button @click="templateDialogVisible = false">{{ isViewMode ? '关闭' : '取消' }}</el-button>
        <el-button v-if="!isViewMode" type="primary" :loading="submitLoading" @click="handleTemplateSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 点检项添加/编辑/查看Dialog -->
    <el-dialog v-model="itemDialogVisible" :title="itemDialogTitle" width="550px">
      <el-form ref="itemFormRef" :model="itemFormData" :rules="itemRules" label-width="110px">
        <el-form-item label="点检项目名称" prop="item_name">
          <el-input v-model="itemFormData.item_name" :disabled="isItemViewMode" />
        </el-form-item>
        <el-form-item label="点检标准" prop="check_standard">
          <el-input v-model="itemFormData.check_standard" :disabled="isItemViewMode" />
        </el-form-item>
        <el-form-item label="点检方法" prop="check_method">
          <el-select v-model="itemFormData.check_method" :disabled="isItemViewMode" placeholder="请选择" style="width:100%">
            <el-option label="目视" value="VISUAL" />
            <el-option label="测量" value="MEASURE" />
            <el-option label="敲击" value="TAP" />
            <el-option label="听音" value="AUDIO" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item label="判定类型" prop="result_type">
          <el-select v-model="itemFormData.result_type" :disabled="isItemViewMode" placeholder="请选择" style="width:100%">
            <el-option label="OK/NG" value="OK_NG" />
            <el-option label="数值" value="NUMERIC" />
            <el-option label="文本" value="TEXT" />
          </el-select>
        </el-form-item>
        <template v-if="itemFormData.result_type === 'NUMERIC'">
          <el-row :gutter="16">
            <el-col :span="12">
              <el-form-item label="下限值" prop="lower_limit">
                <el-input-number v-model="itemFormData.lower_limit" :disabled="isItemViewMode" style="width:100%" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="上限值" prop="upper_limit">
                <el-input-number v-model="itemFormData.upper_limit" :disabled="isItemViewMode" style="width:100%" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="单位">
            <el-input v-model="itemFormData.unit" :disabled="isItemViewMode" />
          </el-form-item>
        </template>
        <el-form-item label="关键点">
          <el-switch v-model="itemFormData.is_critical_point" :disabled="isItemViewMode" :active-value="1" :inactive-value="0" />
        </el-form-item>
        <el-form-item label="必检项">
          <el-switch v-model="itemFormData.is_mandatory" :disabled="isItemViewMode" :active-value="1" :inactive-value="0" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="itemFormData.remark" type="textarea" :rows="2" :disabled="isItemViewMode" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="itemDialogVisible = false">{{ isItemViewMode ? '关闭' : '取消' }}</el-button>
        <el-button v-if="!isItemViewMode" type="primary" :loading="itemSubmitLoading" @click="handleItemSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getInspectionTemplateList, getInspectionTemplate, createInspectionTemplate, updateInspectionTemplate } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()
const loading = ref(false), tableData = ref<any[]>([])
const searchForm = reactive({ template_code: '', template_type: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const submitLoading = ref(false)
const templateDialogVisible = ref(false)
const templateFormRef = ref<FormInstance>()
const isViewMode = ref(false)
const currentTemplateId = ref(0)

const templateFormData = reactive<any>({
  id: 0,
  template_code: '',
  template_name: '',
  template_type: '',
  frequency_type: '',
  frequency_value: 1,
  execution_time: '',
  estimated_minutes: 30,
  is_active: 1,
  effective_date: '',
  expiry_date: '',
  remark: '',
  items: [] as any[],
})

const templateRules: FormRules = {
  template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }],
  template_type: [{ required: true, message: '请选择模板类型', trigger: 'change' }],
  frequency_type: [{ required: true, message: '请选择周期类型', trigger: 'change' }],
}

const templateDialogTitle = computed(() => {
  if (isViewMode.value) return '查看点检标准'
  return currentTemplateId.value ? '编辑点检标准' : '新增点检标准'
})

const getTypeText = (s: string) => ({ DAILY: '日常', WEEKLY: '周', MONTHLY: '月', SPECIAL: '专项' }[s] || s)
const getMethodText = (s: string) => ({ VISUAL: '目视', MEASURE: '测量', TAP: '敲击', AUDIO: '听音', OTHER: '其他' }[s] || s)
const getResultTypeText = (s: string) => ({ OK_NG: 'OK/NG', NUMERIC: '数值', TEXT: '文本' }[s] || s)

const loadData = async () => {
  loading.value = true
  try {
    const res = await getInspectionTemplateList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.template_code = ''; searchForm.template_type = ''; handleSearch() }

const handleAdd = () => {
  currentTemplateId.value = 0
  isViewMode.value = false
  Object.assign(templateFormData, {
    id: 0, template_code: '', template_name: '', template_type: '', frequency_type: '',
    frequency_value: 1, execution_time: '', estimated_minutes: 30, is_active: 1,
    effective_date: '', expiry_date: '', remark: '', items: [],
  })
  templateDialogVisible.value = true
}

const handleEdit = async (row: any) => {
  currentTemplateId.value = row.id
  isViewMode.value = false
  try {
    const res = await getInspectionTemplate(row.id)
    const tpl = res.data.template || res.data
    const items = res.data.items || []
    Object.assign(templateFormData, {
      id: tpl.id, template_code: tpl.template_code, template_name: tpl.template_name,
      template_type: tpl.template_type, frequency_type: tpl.frequency_type,
      frequency_value: tpl.frequency_value, execution_time: tpl.execution_time,
      estimated_minutes: tpl.estimated_minutes, is_active: tpl.is_active,
      effective_date: tpl.effective_date ? String(tpl.effective_date).slice(0, 10) : '',
      expiry_date: tpl.expiry_date ? String(tpl.expiry_date).slice(0, 10) : '',
      remark: tpl.remark || '', items,
    })
    templateDialogVisible.value = true
  } catch {
    ElMessage.error('加载模板详情失败')
  }
}

const handleView = async (row: any) => {
  currentTemplateId.value = row.id
  isViewMode.value = true
  try {
    const res = await getInspectionTemplate(row.id)
    const tpl = res.data.template || res.data
    const items = res.data.items || []
    Object.assign(templateFormData, {
      id: tpl.id, template_code: tpl.template_code, template_name: tpl.template_name,
      template_type: tpl.template_type, frequency_type: tpl.frequency_type,
      frequency_value: tpl.frequency_value, execution_time: tpl.execution_time,
      estimated_minutes: tpl.estimated_minutes, is_active: tpl.is_active,
      effective_date: tpl.effective_date ? String(tpl.effective_date).slice(0, 10) : '',
      expiry_date: tpl.expiry_date ? String(tpl.expiry_date).slice(0, 10) : '',
      remark: tpl.remark || '', items,
    })
    templateDialogVisible.value = true
  } catch {
    ElMessage.error('加载模板详情失败')
  }
}

const handleTemplateSubmit = async () => {
  if (!templateFormRef.value) return
  await templateFormRef.value.validate()
  submitLoading.value = true
  try {
    const payload = {
      template: {
        template_name: templateFormData.template_name,
        template_type: templateFormData.template_type,
        frequency_type: templateFormData.frequency_type,
        frequency_value: templateFormData.frequency_value,
        execution_time: templateFormData.execution_time,
        estimated_minutes: templateFormData.estimated_minutes,
        is_active: templateFormData.is_active,
        effective_date: templateFormData.effective_date,
        expiry_date: templateFormData.expiry_date,
        remark: templateFormData.remark,
      },
      items: templateFormData.items.map((item: any, idx: number) => ({
        ...item,
        item_code: item.item_code || `ITEM-${Date.now()}-${idx}`,
        sort_order: item.sort_order || idx + 1,
      })),
    }
    if (currentTemplateId.value) {
      await updateInspectionTemplate(currentTemplateId.value, payload)
      ElMessage.success('更新成功')
    } else {
      await createInspectionTemplate(payload)
      ElMessage.success('创建成功')
    }
    templateDialogVisible.value = false
    loadData()
  } catch (e: any) {
    ElMessage.error(e?.message || e?.error || '操作失败')
  } finally {
    submitLoading.value = false
  }
}

// ---- 点检项目操作 ----
const itemDialogVisible = ref(false)
const itemSubmitLoading = ref(false)
const itemFormRef = ref<FormInstance>()
const isItemViewMode = ref(false)
const editingItemIndex = ref(-1)

const itemFormData = reactive<any>({
  item_name: '', check_standard: '', check_method: '', result_type: 'OK_NG',
  upper_limit: 0, lower_limit: 0, unit: '', is_critical_point: 0, is_mandatory: 1, remark: '',
})

const itemRules: FormRules = {
  item_name: [{ required: true, message: '请输入点检项目名称', trigger: 'blur' }],
  check_standard: [{ required: true, message: '请输入点检标准', trigger: 'blur' }],
  check_method: [{ required: true, message: '请选择点检方法', trigger: 'change' }],
  result_type: [{ required: true, message: '请选择判定类型', trigger: 'change' }],
}

const itemDialogTitle = computed(() => {
  if (isItemViewMode.value) return '查看点检项'
  return editingItemIndex.value >= 0 ? '编辑点检项' : '添加点检项'
})

const handleAddItem = () => {
  editingItemIndex.value = -1
  isItemViewMode.value = false
  Object.assign(itemFormData, {
    item_name: '', check_standard: '', check_method: '', result_type: 'OK_NG',
    upper_limit: 0, lower_limit: 0, unit: '', is_critical_point: 0, is_mandatory: 1, remark: '',
  })
  itemDialogVisible.value = true
}

const handleEditItem = (row: any) => {
  editingItemIndex.value = templateFormData.items.indexOf(row)
  isItemViewMode.value = false
  Object.assign(itemFormData, { ...row })
  itemDialogVisible.value = true
}

const handleViewItem = (row: any) => {
  editingItemIndex.value = templateFormData.items.indexOf(row)
  isItemViewMode.value = true
  Object.assign(itemFormData, { ...row })
  itemDialogVisible.value = true
}

const handleDeleteItem = (index: number) => {
  templateFormData.items.splice(index, 1)
}

const handleItemSubmit = async () => {
  if (!itemFormRef.value) return
  await itemFormRef.value.validate()
  itemSubmitLoading.value = true
  try {
    if (editingItemIndex.value >= 0) {
      templateFormData.items[editingItemIndex.value] = { ...itemFormData }
    } else {
      templateFormData.items.push({ ...itemFormData })
    }
    itemDialogVisible.value = false
  } finally {
    itemSubmitLoading.value = false
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.inspection-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}

.items-section {
  border: 1px solid #ebeef5;
  border-radius: 4px;
  padding: 12px;
  margin-top: 16px;
  .items-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
    .items-title { font-size: 14px; font-weight: 500; color: #303133; }
  }
}
</style>
