<template>
  <div class="dynamic-rule-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="规则编码">
          <el-input v-model="searchForm.rule_code" placeholder="请输入规则编码" clearable />
        </el-form-item>
        <el-form-item label="检验类型">
          <el-select v-model="searchForm.inspection_type" placeholder="请选择" clearable>
            <el-option label="IQC来料检验" value="IQC" />
            <el-option label="IPQC过程检验" value="IPQC" />
            <el-option label="FQC最终检验" value="FQC" />
            <el-option label="OOC出货检验" value="OOC" />
          </el-select>
        </el-form-item>
        <el-form-item label="触发类型">
          <el-select v-model="searchForm.trigger_type" placeholder="请选择" clearable>
            <el-option label="检验前" value="BEFORE" />
            <el-option label="检验后" value="AFTER" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.enabled" placeholder="请选择" clearable>
            <el-option label="启用" value="true" />
            <el-option label="禁用" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:dynamicRule:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="info" @click="handleEvaluate">
        <el-icon><MagicStick /></el-icon>测试规则
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="rule_code" label="规则编码" width="140" />
        <el-table-column prop="rule_name" label="规则名称" min-width="150" />
        <el-table-column prop="inspection_type" label="检验类型" width="100">
          <template #default="{ row }">
            <el-tag :type="getInspectionTypeTag(row.inspection_type)">{{ getInspectionTypeText(row.inspection_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="trigger_type" label="触发类型" width="90">
          <template #default="{ row }">
            <el-tag :type="row.trigger_type === 'BEFORE' ? 'warning' : 'success'">
              {{ row.trigger_type === 'BEFORE' ? '检验前' : '检验后' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="action_type" label="动作类型" width="120">
          <template #default="{ row }">
            <el-tag :type="getActionTypeTag(row.action_type)">{{ getActionTypeText(row.action_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="priority" label="优先级" width="80" />
        <el-table-column prop="enabled" label="状态" width="80">
          <template #default="{ row }">
            <el-switch
              v-model="row.enabled"
              @change="handleEnabledChange(row)"
              :disabled="!hasPermission('quality:dynamicRule:edit')"
            />
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="120" show-overflow-tooltip />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('quality:dynamicRule:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('quality:dynamicRule:delete')" @click="handleDelete(row)">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px" class="rule-dialog">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="110px">
        <el-form-item label="规则编码" prop="rule_code">
          <el-input v-model="formData.rule_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="规则名称" prop="rule_name">
          <el-input v-model="formData.rule_name" />
        </el-form-item>
        <el-form-item label="检验类型" prop="inspection_type">
          <el-select v-model="formData.inspection_type" placeholder="请选择">
            <el-option label="IQC来料检验" value="IQC" />
            <el-option label="IPQC过程检验" value="IPQC" />
            <el-option label="FQC最终检验" value="FQC" />
            <el-option label="OOC出货检验" value="OOC" />
          </el-select>
        </el-form-item>
        <el-form-item label="触发类型" prop="trigger_type">
          <el-radio-group v-model="formData.trigger_type">
            <el-radio label="BEFORE">检验前</el-radio>
            <el-radio label="AFTER">检验后</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="条件表达式" prop="condition_str">
          <el-input
            v-model="formData.condition_str"
            type="textarea"
            :rows="3"
            placeholder="例如: quality_level == &quot;A&quot; && defect_count > 5"
          />
          <div class="form-tip">
            支持操作符: ==, !=, >, <, >=, <=, in, not_in, contains
            多条件用 && 分隔，例如: supplier_rating == &quot;C&quot; && defect_count > 0
          </div>
        </el-form-item>
        <el-form-item label="动作类型" prop="action_type">
          <el-select v-model="formData.action_type" placeholder="请选择">
            <el-option label="增加检验项" value="ADD_ITEM" />
            <el-option label="跳过检验项" value="SKIP_ITEM" />
            <el-option label="升级检验级别" value="UPGRADE_LEVEL" />
            <el-option label="拒收" value="REJECT" />
          </el-select>
        </el-form-item>
        <el-form-item label="动作配置" prop="action_config">
          <template v-if="formData.action_type === 'ADD_ITEM'">
            <el-input v-model="formData.action_config.additional_items_str" placeholder="增加的检验项编码，多个用逗号分隔" />
          </template>
          <template v-else-if="formData.action_type === 'SKIP_ITEM'">
            <el-input v-model="formData.action_config.skip_items_str" placeholder="跳过的检验项编码，多个用逗号分隔" />
          </template>
          <template v-else-if="formData.action_type === 'UPGRADE_LEVEL'">
            <el-select v-model="formData.action_config.upgrade_to_level" placeholder="请选择">
              <el-option label="普通" value="NORMAL" />
              <el-option label="加严" value="STRICT" />
              <el-option label="特采" value="SPECIAL" />
            </el-select>
          </template>
          <template v-else-if="formData.action_type === 'REJECT'">
            <el-input v-model="formData.action_config.reject_reason" placeholder="拒收原因" />
          </template>
        </el-form-item>
        <el-form-item label="优先级">
          <el-input-number v-model="formData.priority" :min="0" :max="100" />
        </el-form-item>
        <el-form-item label="启用">
          <el-switch v-model="formData.enabled" />
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

    <!-- 测试规则对话框 -->
    <el-dialog v-model="evaluateDialogVisible" title="测试规则" width="600px">
      <el-form :model="evaluateForm" label-width="120px">
        <el-form-item label="检验类型">
          <el-select v-model="evaluateForm.inspection_type" placeholder="请选择">
            <el-option label="IQC来料检验" value="IQC" />
            <el-option label="IPQC过程检验" value="IPQC" />
            <el-option label="FQC最终检验" value="FQC" />
            <el-option label="OOC出货检验" value="OOC" />
          </el-select>
        </el-form-item>
        <el-form-item label="供应商等级">
          <el-input v-model="evaluateForm.supplier_rating" placeholder="例如: A, B, C" />
        </el-form-item>
        <el-form-item label="质量等级">
          <el-input v-model="evaluateForm.quality_level" placeholder="例如: A, B, C" />
        </el-form-item>
        <el-form-item label="不良数量">
          <el-input-number v-model="evaluateForm.defect_count" :min="0" />
        </el-form-item>
        <el-form-item label="不良类型">
          <el-input v-model="evaluateForm.defect_types_str" placeholder="多个用逗号分隔，例如: 焊接不良,尺寸超差" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="evaluateDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="evaluateLoading" @click="handleEvaluateSubmit">评估</el-button>
      </template>
    </el-dialog>

    <!-- 评估结果对话框 -->
    <el-dialog v-model="resultDialogVisible" title="规则评估结果" width="600px">
      <el-alert v-if="evaluateResults.length === 0" type="info" :closable="false">
        没有匹配的规则
      </el-alert>
      <el-table v-else :data="evaluateResults">
        <el-table-column prop="rule_name" label="规则名称" />
        <el-table-column prop="action_type" label="动作类型">
          <template #default="{ row }">
            <el-tag :type="getActionTypeTag(row.action_type)">{{ getActionTypeText(row.action_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="动作配置" prop="action_config_str" />
      </el-table>
      <template #footer>
        <el-button @click="resultDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getDynamicRuleList, createDynamicRule, updateDynamicRule, deleteDynamicRule, evaluateDynamicRules } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const evaluateDialogVisible = ref(false)
const evaluateLoading = ref(false)
const resultDialogVisible = ref(false)
const evaluateResults = ref<any[]>([])

const searchForm = reactive({
  rule_code: '',
  inspection_type: '',
  trigger_type: '',
  enabled: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  id: 0,
  rule_code: '',
  rule_name: '',
  inspection_type: 'IQC',
  trigger_type: 'BEFORE',
  condition_str: '',
  action_type: 'ADD_ITEM',
  action_config: {
    additional_items_str: '',
    skip_items_str: '',
    upgrade_to_level: '',
    reject_reason: ''
  },
  priority: 0,
  enabled: true,
  remark: ''
})

const evaluateForm = reactive({
  inspection_type: 'IQC',
  supplier_rating: '',
  quality_level: '',
  defect_count: 0,
  defect_types_str: ''
})

const rules: FormRules = {
  rule_code: [{ required: true, message: '请输入规则编码', trigger: 'blur' }],
  rule_name: [{ required: true, message: '请输入规则名称', trigger: 'blur' }],
  inspection_type: [{ required: true, message: '请选择检验类型', trigger: 'change' }],
  trigger_type: [{ required: true, message: '请选择触发类型', trigger: 'change' }],
  action_type: [{ required: true, message: '请选择动作类型', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑动态规则' : '新增动态规则')

const getInspectionTypeText = (type: string) => {
  const map: Record<string, string> = { IQC: 'IQC', IPQC: 'IPQC', FQC: 'FQC', OOC: 'OOC' }
  return map[type] || type
}

const getInspectionTypeTag = (type: string) => {
  const map: Record<string, string> = { IQC: 'warning', IPQC: 'success', FQC: 'info', OOC: 'danger' }
  return map[type] || 'info'
}

const getActionTypeText = (type: string) => {
  const map: Record<string, string> = {
    ADD_ITEM: '增加检验项',
    SKIP_ITEM: '跳过检验项',
    UPGRADE_LEVEL: '升级级别',
    REJECT: '拒收'
  }
  return map[type] || type
}

const getActionTypeTag = (type: string) => {
  const map: Record<string, string> = {
    ADD_ITEM: 'success',
    SKIP_ITEM: 'warning',
    UPGRADE_LEVEL: 'info',
    REJECT: 'danger'
  }
  return map[type] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.rule_code) params.rule_code = searchForm.rule_code
    if (searchForm.inspection_type) params.inspection_type = searchForm.inspection_type
    if (searchForm.trigger_type) params.trigger_type = searchForm.trigger_type
    if (searchForm.enabled) params.enabled = searchForm.enabled

    const res = await getDynamicRuleList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.rule_code = ''
  searchForm.inspection_type = ''
  searchForm.trigger_type = ''
  searchForm.enabled = ''
  handleSearch()
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0,
    rule_code: '',
    rule_name: '',
    inspection_type: 'IQC',
    trigger_type: 'BEFORE',
    condition_str: '',
    action_type: 'ADD_ITEM',
    action_config: {
      additional_items_str: '',
      skip_items_str: '',
      upgrade_to_level: '',
      reject_reason: ''
    },
    priority: 0,
    enabled: true,
    remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  // 解析 condition_expression
  try {
    const conditions = JSON.parse(row.condition_expression || '[]')
    if (Array.isArray(conditions) && conditions.length > 0) {
      formData.condition_str = conditions.map((c: any) => {
        if (c.operator === 'in' || c.operator === 'not_in') {
          return `${c.field} ${c.operator} [${c.value}]`
        }
        return `${c.field} ${c.operator} ${typeof c.value === 'string' ? `"${c.value}"` : c.value}`
      }).join(' && ')
    } else {
      formData.condition_str = ''
    }
  } catch {
    formData.condition_str = ''
  }
  // 解析 action_config
  try {
    const config = JSON.parse(row.action_config || '{}')
    if (config.additional_items) {
      formData.action_config.additional_items_str = config.additional_items.join(',')
    }
    if (config.skip_items) {
      formData.action_config.skip_items_str = config.skip_items.join(',')
    }
    if (config.upgrade_to_level) {
      formData.action_config.upgrade_to_level = config.upgrade_to_level
    }
    if (config.reject_reason) {
      formData.action_config.reject_reason = config.reject_reason
    }
  } catch {
    formData.action_config = { additional_items_str: '', skip_items_str: '', upgrade_to_level: '', reject_reason: '' }
  }
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该规则吗？', '提示', { type: 'warning' })
    await deleteDynamicRule(row.id)
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
    const data: any = {
      rule_code: formData.rule_code,
      rule_name: formData.rule_name,
      inspection_type: formData.inspection_type,
      trigger_type: formData.trigger_type,
      action_type: formData.action_type,
      priority: formData.priority,
      enabled: formData.enabled,
      remark: formData.remark,
      condition_str: formData.condition_str
    }

    // 构建 action_config
    const actionConfig: any = {}
    if (formData.action_type === 'ADD_ITEM' && formData.action_config.additional_items_str) {
      actionConfig.additional_items = formData.action_config.additional_items_str.split(',').map((s: string) => s.trim()).filter(Boolean)
    } else if (formData.action_type === 'SKIP_ITEM' && formData.action_config.skip_items_str) {
      actionConfig.skip_items = formData.action_config.skip_items_str.split(',').map((s: string) => s.trim()).filter(Boolean)
    } else if (formData.action_type === 'UPGRADE_LEVEL' && formData.action_config.upgrade_to_level) {
      actionConfig.upgrade_to_level = formData.action_config.upgrade_to_level
    } else if (formData.action_type === 'REJECT' && formData.action_config.reject_reason) {
      actionConfig.reject_reason = formData.action_config.reject_reason
    }
    data.action_config = actionConfig

    if (formData.id) {
      await updateDynamicRule(formData.id, data)
    } else {
      await createDynamicRule(data)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally {
    submitLoading.value = false
  }
}

const handleEnabledChange = async (row: any) => {
  try {
    await updateDynamicRule(row.id, { enabled: row.enabled })
    ElMessage.success('状态更新成功')
  } catch (error) {
    row.enabled = !row.enabled // revert
  }
}

const handleEvaluate = () => {
  evaluateResults.value = []
  evaluateDialogVisible.value = true
}

const handleEvaluateSubmit = async () => {
  evaluateLoading.value = true
  try {
    const defectTypes = evaluateForm.defect_types_str
      ? evaluateForm.defect_types_str.split(',').map(s => s.trim()).filter(Boolean)
      : []

    const res = await evaluateDynamicRules({
      inspection_type: evaluateForm.inspection_type,
      supplier_rating: evaluateForm.supplier_rating,
      quality_level: evaluateForm.quality_level,
      defect_count: evaluateForm.defect_count,
      defect_types: defectTypes
    })

    evaluateResults.value = (res.data.results || []).filter((r: any) => r.matched).map((r: any) => {
      let actionConfigStr = ''
      if (r.action_type === 'ADD_ITEM' && r.action_config.additional_items) {
        actionConfigStr = '增加: ' + r.action_config.additional_items.join(', ')
      } else if (r.action_type === 'SKIP_ITEM' && r.action_config.skip_items) {
        actionConfigStr = '跳过: ' + r.action_config.skip_items.join(', ')
      } else if (r.action_type === 'UPGRADE_LEVEL' && r.action_config.upgrade_to_level) {
        actionConfigStr = '升级到: ' + r.action_config.upgrade_to_level
      } else if (r.action_type === 'REJECT' && r.action_config.reject_reason) {
        actionConfigStr = '原因: ' + r.action_config.reject_reason
      }
      return { ...r, action_config_str: actionConfigStr }
    })

    resultDialogVisible.value = true
  } catch (error) {
    ElMessage.error('评估失败')
  } finally {
    evaluateLoading.value = false
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.dynamic-rule-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .form-tip { font-size: 12px; color: #999; margin-top: 4px; line-height: 1.4; }
}
</style>
