<template>
  <div class="code-rule-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="规则编码">
          <el-input v-model="searchForm.rule_code" placeholder="请输入规则编码" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:code-rule:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增规则
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="rule_code" label="规则编码" min-width="120" />
        <el-table-column prop="rule_name" label="规则名称" min-width="150" />
        <el-table-column prop="entity_type" label="实体类型" width="150" />
        <el-table-column prop="prefix" label="前缀" width="80" />
        <el-table-column prop="date_format" label="日期格式" width="100" />
        <el-table-column prop="seq_length" label="序号长度" width="80" />
        <el-table-column prop="reset_type" label="重置类型" width="100">
          <template #default="{ row }">
            <el-tag :type="row.reset_type === 'NONE' ? 'info' : 'warning'">
              {{ getResetTypeText(row.reset_type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="example" label="示例" min-width="150" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('production:code-rule:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" v-if="hasPermission('production:code-rule:generate')" @click="handleGenerate(row)">生成编码</el-button>
            <el-button link type="danger" v-if="hasPermission('production:code-rule:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px" @close="handleDialogClose">
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="100px">
        <el-form-item label="规则编码" prop="rule_code">
          <el-input v-model="form.rule_code" placeholder="请输入规则编码" />
        </el-form-item>
        <el-form-item label="规则名称" prop="rule_name">
          <el-input v-model="form.rule_name" placeholder="请输入规则名称" />
        </el-form-item>
        <el-form-item label="实体类型" prop="entity_type">
          <el-select v-model="form.entity_type" placeholder="请选择" style="width: 100%">
            <el-option label="生产工单" value="PRODUCTION_ORDER" />
            <el-option label="销售订单" value="SALES_ORDER" />
            <el-option label="包装" value="PACKAGE" />
            <el-option label="检验单" value="INSPECT" />
            <el-option label="流程卡" value="FLOW_CARD" />
          </el-select>
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="前缀" prop="prefix">
              <el-input v-model="form.prefix" placeholder="如PO" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="日期格式" prop="date_format">
              <el-select v-model="form.date_format" placeholder="请选择" style="width: 100%">
                <el-option label="无" value="" />
                <el-option label="YYYYMMDD" value="YYYYMMDD" />
                <el-option label="YYYYMM" value="YYYYMM" />
                <el-option label="YYYY" value="YYYY" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="序号长度" prop="seq_length">
              <el-input-number v-model="form.seq_length" :min="1" :max="10" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="重置类型" prop="reset_type">
              <el-select v-model="form.reset_type" placeholder="请选择" style="width: 100%">
                <el-option label="不重置" value="NONE" />
                <el-option label="每日" value="DAILY" />
                <el-option label="每月" value="MONTHLY" />
                <el-option label="每年" value="YEARLY" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="示例" prop="example">
          <el-input v-model="form.example" placeholder="如PO202604070001" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" :rows="2" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitLoading">确定</el-button>
      </template>
    </el-dialog>

    <!-- 生成编码对话框 -->
    <el-dialog v-model="genDialogVisible" title="生成编码" width="400px">
      <el-form label-width="100px">
        <el-form-item label="规则编码">
          {{ currentRule?.rule_code }}
        </el-form-item>
        <el-form-item label="生成编码">
          <el-input v-model="generatedCode" readonly />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="genDialogVisible = false">关闭</el-button>
        <el-button type="primary" @click="copyCode">复制编码</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getCodeRuleList, getCodeRule, createCodeRule, updateCodeRule, deleteCodeRule, generateCode } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const searchForm = reactive({ rule_code: '' })

const dialogVisible = ref(false)
const dialogTitle = ref('')
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const genDialogVisible = ref(false)
const currentRule = ref<any>(null)
const generatedCode = ref('')

const form = reactive({
  id: undefined as number | undefined,
  rule_code: '',
  rule_name: '',
  entity_type: '',
  prefix: '',
  date_format: '',
  seq_length: 4,
  reset_type: 'NONE',
  status: 1,
  remark: ''
})

const formRules: FormRules = {
  rule_code: [{ required: true, message: '请输入规则编码', trigger: 'blur' }],
  rule_name: [{ required: true, message: '请输入规则名称', trigger: 'blur' }],
  entity_type: [{ required: true, message: '请选择实体类型', trigger: 'change' }]
}

const getResetTypeText = (type: string) => {
  const map: Record<string, string> = { NONE: '不重置', DAILY: '每日', MONTHLY: '每月', YEARLY: '每年' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getCodeRuleList()
    tableData.value = res.data.list || []
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { loadData() }
const handleReset = () => { searchForm.rule_code = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增规则'
  resetForm()
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑规则'
  try {
    const res = await getCodeRule(row.id)
    Object.assign(form, res.data)
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取规则信息失败')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    submitLoading.value = true
    try {
      if (isEdit.value) {
        await updateCodeRule(form.id!, form)
        ElMessage.success('更新成功')
      } else {
        await createCodeRule(form)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadData()
    } catch (error: any) {
      ElMessage.error(error.message || '操作失败')
    } finally {
      submitLoading.value = false
    }
  })
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该规则吗？', '提示', { type: 'warning' })
    await deleteCodeRule(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

const handleGenerate = async (row: any) => {
  currentRule.value = row
  try {
    const res = await generateCode(row.rule_code)
    generatedCode.value = res.data.code
    genDialogVisible.value = true
  } catch (error: any) {
    ElMessage.error(error.message || '生成失败')
  }
}

const copyCode = () => {
  navigator.clipboard.writeText(generatedCode.value)
  ElMessage.success('已复制到剪贴板')
}

const resetForm = () => {
  form.id = undefined
  form.rule_code = ''
  form.rule_name = ''
  form.entity_type = ''
  form.prefix = ''
  form.date_format = ''
  form.seq_length = 4
  form.reset_type = 'NONE'
  form.status = 1
  form.remark = ''
}

const handleDialogClose = () => {
  formRef.value?.resetFields()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.code-rule-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
}
</style>
