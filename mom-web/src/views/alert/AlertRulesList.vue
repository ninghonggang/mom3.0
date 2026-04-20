<template>
  <div class="alert-rules-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="规则名称">
          <el-input v-model="searchForm.rule_name" placeholder="请输入规则名称" clearable />
        </el-form-item>
        <el-form-item label="告警类型">
          <el-select v-model="searchForm.alert_type" placeholder="请选择" clearable>
            <el-option label="设备告警" value="EQUIPMENT" />
            <el-option label="生产告警" value="PRODUCTION" />
            <el-option label="质量告警" value="QUALITY" />
            <el-option label="库存告警" value="INVENTORY" />
            <el-option label="系统告警" value="SYSTEM" />
          </el-select>
        </el-form-item>
        <el-form-item label="严重级别">
          <el-select v-model="searchForm.severity_level" placeholder="请选择" clearable>
            <el-option label="Critical" value="CRITICAL" />
            <el-option label="High" value="HIGH" />
            <el-option label="Medium" value="MEDIUM" />
            <el-option label="Low" value="LOW" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.is_enabled" placeholder="请选择" clearable>
            <el-option label="已启用" :value="1" />
            <el-option label="已禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('alert:rules:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增规则
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="rule_code" label="规则编码" width="140" />
        <el-table-column prop="rule_name" label="规则名称" min-width="180" />
        <el-table-column prop="alert_type" label="告警类型" width="110">
          <template #default="{ row }">
            <el-tag>{{ getAlertTypeText(row.alert_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="severity_level" label="严重级别" width="100">
          <template #default="{ row }">
            <el-tag :type="getSeverityType(row.severity_level)">{{ row.severity_level }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="biz_module" label="业务模块" width="110" />
        <el-table-column prop="check_interval" label="检查间隔(秒)" width="110" />
        <el-table-column prop="trigger_count" label="触发次数" width="90" />
        <el-table-column prop="last_trigger_time" label="最近触发" width="160">
          <template #default="{ row }">
            {{ row.last_trigger_time ? formatDateTime(row.last_trigger_time) : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="is_enabled" label="状态" width="90">
          <template #default="{ row }">
            <el-switch
              v-model="row.is_enabled"
              :active-value="1"
              :inactive-value="0"
              @change="handleToggleStatus(row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('alert:rules:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('alert:rules:delete')" @click="handleDelete(row)">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="110px">
        <el-form-item label="规则编码" prop="rule_code">
          <el-input v-model="formData.rule_code" :disabled="!!formData.id" placeholder="请输入规则编码" />
        </el-form-item>
        <el-form-item label="规则名称" prop="rule_name">
          <el-input v-model="formData.rule_name" placeholder="请输入规则名称" />
        </el-form-item>
        <el-form-item label="告警类型" prop="alert_type">
          <el-select v-model="formData.alert_type" placeholder="请选择告警类型">
            <el-option label="设备告警" value="EQUIPMENT" />
            <el-option label="生产告警" value="PRODUCTION" />
            <el-option label="质量告警" value="QUALITY" />
            <el-option label="库存告警" value="INVENTORY" />
            <el-option label="系统告警" value="SYSTEM" />
          </el-select>
        </el-form-item>
        <el-form-item label="业务模块">
          <el-input v-model="formData.biz_module" placeholder="请输入业务模块" />
        </el-form-item>
        <el-form-item label="条件表达式" prop="condition_expression">
          <el-input v-model="formData.condition_expression" type="textarea" :rows="3" placeholder="如: temperature > 100" />
        </el-form-item>
        <el-form-item label="严重级别" prop="severity_level">
          <el-select v-model="formData.severity_level" placeholder="请选择">
            <el-option label="Critical" value="CRITICAL" />
            <el-option label="High" value="HIGH" />
            <el-option label="Medium" value="MEDIUM" />
            <el-option label="Low" value="LOW" />
          </el-select>
        </el-form-item>
        <el-form-item label="检查间隔(秒)">
          <el-input-number v-model="formData.check_interval" :min="10" :step="10" />
        </el-form-item>
        <el-form-item label="最大触发次数">
          <el-input-number v-model="formData.max_trigger_count" :min="0" />
        </el-form-item>
        <el-form-item label="通知方式">
          <el-checkbox-group v-model="notificationChannels">
            <el-checkbox label="EMAIL">邮件</el-checkbox>
            <el-checkbox label="SMS">短信</el-checkbox>
            <el-checkbox label="WECHAT">企业微信</el-checkbox>
            <el-checkbox label="DINGTALK">钉钉</el-checkbox>
          </el-checkbox-group>
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
import {
  getAlertRulesList,
  createAlertRule,
  updateAlertRule,
  deleteAlertRule,
  enableAlertRule,
  disableAlertRule
} from '@/api/alert'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const notificationChannels = ref<string[]>([])

const searchForm = reactive({
  rule_name: '',
  alert_type: '',
  severity_level: '',
  is_enabled: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  id: 0,
  rule_code: '',
  rule_name: '',
  alert_type: '',
  biz_module: '',
  condition_expression: '',
  condition_params: {},
  severity_level: 'MEDIUM',
  notification_channels: [],
  notify_templates: {},
  escalation_rule_id: null,
  check_interval: 60,
  max_trigger_count: 0,
  remark: ''
})

const rules: FormRules = {
  rule_code: [{ required: true, message: '请输入规则编码', trigger: 'blur' }],
  rule_name: [{ required: true, message: '请输入规则名称', trigger: 'blur' }],
  alert_type: [{ required: true, message: '请选择告警类型', trigger: 'change' }],
  condition_expression: [{ required: true, message: '请输入条件表达式', trigger: 'blur' }],
  severity_level: [{ required: true, message: '请选择严重级别', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑规则' : '新增规则')

const getAlertTypeText = (type: string) => {
  const map: Record<string, string> = {
    EQUIPMENT: '设备告警',
    PRODUCTION: '生产告警',
    QUALITY: '质量告警',
    INVENTORY: '库存告警',
    SYSTEM: '系统告警'
  }
  return map[type] || type
}

const getSeverityType = (level: string) => {
  const map: Record<string, string> = {
    CRITICAL: 'danger',
    HIGH: 'warning',
    MEDIUM: 'info',
    LOW: 'success'
  }
  return map[level] || 'info'
}

const formatDateTime = (datetime: string) => {
  if (!datetime) return '-'
  const d = new Date(datetime)
  return d.toLocaleString()
}

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getAlertRulesList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }

const handleReset = () => {
  searchForm.rule_name = ''
  searchForm.alert_type = ''
  searchForm.severity_level = ''
  searchForm.is_enabled = ''
  handleSearch()
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0,
    rule_code: '',
    rule_name: '',
    alert_type: '',
    biz_module: '',
    condition_expression: '',
    condition_params: {},
    severity_level: 'MEDIUM',
    notification_channels: [],
    notify_templates: {},
    escalation_rule_id: null,
    check_interval: 60,
    max_trigger_count: 0,
    remark: ''
  })
  notificationChannels.value = []
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  // 解析 notification_channels
  if (row.notification_channels && Array.isArray(row.notification_channels)) {
    notificationChannels.value = row.notification_channels
  } else {
    notificationChannels.value = []
  }
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该规则吗？', '提示', { type: 'warning' })
    await deleteAlertRule(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleToggleStatus = async (row: any) => {
  try {
    if (row.is_enabled) {
      await enableAlertRule(row.id)
      ElMessage.success('已启用')
    } else {
      await disableAlertRule(row.id)
      ElMessage.success('已禁用')
    }
  } catch (error) {
    row.is_enabled = row.is_enabled ? 0 : 1
    ElMessage.error('操作失败')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.notification_channels = notificationChannels.value
    if (formData.id) {
      await updateAlertRule(formData.id, formData)
    } else {
      await createAlertRule(formData)
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
.alert-rules-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
