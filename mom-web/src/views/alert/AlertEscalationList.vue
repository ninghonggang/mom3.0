<template>
  <div class="alert-escalation-list">
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
      <el-button type="primary" v-if="hasPermission('alert:escalation:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增升级规则
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="rule_code" label="规则编码" width="140" />
        <el-table-column prop="rule_name" label="规则名称" min-width="180" />
        <el-table-column prop="alert_type" label="告警类型" width="110">
          <template #default="{ row }">
            {{ row.alert_type ? getAlertTypeText(row.alert_type) : '通用' }}
          </template>
        </el-table-column>
        <el-table-column prop="severity_level" label="严重级别" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.severity_level" :type="getSeverityType(row.severity_level)">{{ row.severity_level }}</el-tag>
            <span v-else>通用</span>
          </template>
        </el-table-column>
        <el-table-column prop="escalation_levels" label="升级路径" min-width="200">
          <template #default="{ row }">
            <template v-if="row.escalation_levels && Array.isArray(row.escalation_levels)">
              <span v-for="(level, index) in row.escalation_levels" :key="index" class="escalation-level">
                <el-tag size="small">{{ level.level || index + 1 }}</el-tag>
                <span class="level-text">{{ level.name || '层级' + (index + 1) }}</span>
                <span v-if="index < row.escalation_levels.length - 1" class="arrow">→</span>
              </span>
            </template>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="is_enabled" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="row.is_enabled ? 'success' : 'info'">{{ row.is_enabled ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('alert:escalation:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('alert:escalation:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="告警类型">
          <el-select v-model="formData.alert_type" placeholder="请选择（为空表示通用）" clearable>
            <el-option label="设备告警" value="EQUIPMENT" />
            <el-option label="生产告警" value="PRODUCTION" />
            <el-option label="质量告警" value="QUALITY" />
            <el-option label="库存告警" value="INVENTORY" />
            <el-option label="系统告警" value="SYSTEM" />
          </el-select>
        </el-form-item>
        <el-form-item label="严重级别">
          <el-select v-model="formData.severity_level" placeholder="请选择（为空表示通用）" clearable>
            <el-option label="Critical" value="CRITICAL" />
            <el-option label="High" value="HIGH" />
            <el-option label="Medium" value="MEDIUM" />
            <el-option label="Low" value="LOW" />
          </el-select>
        </el-form-item>
        <el-form-item label="升级路径" prop="escalation_levels">
          <div class="escalation-config">
            <div v-for="(level, index) in escalationLevels" :key="index" class="escalation-row">
              <span class="level-num">第{{ index + 1 }}级</span>
              <el-input v-model="level.name" placeholder="层级名称" style="width: 120px" />
              <el-input v-model="level.timeout" placeholder="超时时间" style="width: 100px">
                <template #append>分钟</template>
              </el-input>
              <el-input v-model="level.notify_channels" placeholder="通知方式" style="width: 140px" />
              <el-button link type="danger" @click="removeLevel(index)" :disabled="escalationLevels.length <= 1">删除</el-button>
            </div>
            <el-button type="primary" link @click="addLevel">+ 添加层级</el-button>
          </div>
        </el-form-item>
        <el-form-item label="是否启用">
          <el-switch v-model="formData.is_enabled" :active-value="1" :inactive-value="0" />
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
  getAlertEscalationList,
  createAlertEscalation,
  updateAlertEscalation,
  deleteAlertEscalation
} from '@/api/alert'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({
  rule_name: '',
  alert_type: '',
  is_enabled: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  id: 0,
  rule_code: '',
  rule_name: '',
  alert_type: '',
  severity_level: '',
  escalation_levels: [],
  is_enabled: 1,
  remark: ''
})

const escalationLevels = ref<any[]>([
  { name: '', timeout: '', notify_channels: '' }
])

const rules: FormRules = {
  rule_code: [{ required: true, message: '请输入规则编码', trigger: 'blur' }],
  rule_name: [{ required: true, message: '请输入规则名称', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑升级规则' : '新增升级规则')

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

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getAlertEscalationList(params)
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
  searchForm.is_enabled = ''
  handleSearch()
}

const addLevel = () => {
  escalationLevels.value.push({ name: '', timeout: '', notify_channels: '' })
}

const removeLevel = (index: number) => {
  escalationLevels.value.splice(index, 1)
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0,
    rule_code: '',
    rule_name: '',
    alert_type: '',
    severity_level: '',
    escalation_levels: [],
    is_enabled: 1,
    remark: ''
  })
  escalationLevels.value = [{ name: '', timeout: '', notify_channels: '' }]
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  // 解析 escalation_levels
  if (row.escalation_levels && Array.isArray(row.escalation_levels)) {
    escalationLevels.value = row.escalation_levels.map((level: any) => ({
      name: level.name || '',
      timeout: level.timeout || '',
      notify_channels: level.notify_channels || ''
    }))
  } else {
    escalationLevels.value = [{ name: '', timeout: '', notify_channels: '' }]
  }
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该升级规则吗？', '提示', { type: 'warning' })
    await deleteAlertEscalation(row.id)
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
    formData.escalation_levels = escalationLevels.value.filter(l => l.name || l.timeout || l.notify_channels)
    if (formData.id) {
      await updateAlertEscalation(formData.id, formData)
    } else {
      await createAlertEscalation(formData)
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
.alert-escalation-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }

  .escalation-config {
    width: 100%;
  }

  .escalation-row {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 8px;

    .level-num {
      width: 50px;
      color: #909399;
      font-size: 12px;
    }
  }

  .escalation-level {
    display: inline-flex;
    align-items: center;
    gap: 4px;

    .level-text {
      margin: 0 4px;
    }

    .arrow {
      color: #409eff;
      margin: 0 4px;
    }
  }
}
</style>
