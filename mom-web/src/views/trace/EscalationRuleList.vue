<template>
  <div class="escalation-rule-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="规则名称">
          <el-input v-model="searchForm.rule_name" placeholder="请输入规则名称" clearable />
        </el-form-item>
        <el-form-item label="呼叫类型">
          <el-select v-model="searchForm.andon_type" placeholder="请选择" clearable>
            <el-option label="设备故障" value="EQUIPMENT" />
            <el-option label="物料缺料" value="MATERIAL" />
            <el-option label="品质异常" value="QUALITY" />
            <el-option label="工艺异常" value="TECHNICAL" />
            <el-option label="工装夹具" value="TOOLING" />
            <el-option label="安全警示" value="SAFETY" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.is_enabled" placeholder="请选择" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="停用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增规则
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" stripe>
        <el-table-column prop="rule_code" label="规则编码" width="130" />
        <el-table-column prop="rule_name" label="规则名称" min-width="150" />
        <el-table-column prop="andon_type_name" label="适用类型" width="100">
          <template #default="{ row }">
            <el-tag :type="getTypeTag(row.andon_type)">{{ row.andon_type_name || row.andon_type || '全部' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="升级模式" width="100">
          <template #default="{ row }">
            <el-tag :type="row.escalation_mode === 'TIMEOUT' ? 'warning' : 'info'">
              {{ row.escalation_mode === 'TIMEOUT' ? '超时升级' : row.escalation_mode }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="L1超时" width="80">
          <template #default="{ row }">{{ row.level1_timeout }}分钟</template>
        </el-table-column>
        <el-table-column label="L2超时" width="80">
          <template #default="{ row }">{{ row.level2_timeout || '-' }}分钟</template>
        </el-table-column>
        <el-table-column label="L3超时" width="80">
          <template #default="{ row }">{{ row.level3_timeout || '-' }}分钟</template>
        </el-table-column>
        <el-table-column label="最大等级" width="90">
          <template #default="{ row }">
            <el-tag type="danger">L{{ row.max_escalation_level || 4 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="80">
          <template #default="{ row }">
            <el-switch
              v-model="row.is_enabled"
              :active-value="1"
              :inactive-value="0"
              @change="handleToggleStatus(row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleDetail(row)">详情</el-button>
            <el-button link type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新增/编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px" destroy-on-close>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="110px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="规则编码" prop="rule_code">
              <el-input v-model="formData.rule_code" :disabled="isEdit" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="规则名称" prop="rule_name">
              <el-input v-model="formData.rule_name" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="适用类型">
              <el-select v-model="formData.andon_type" placeholder="为空则适用全部">
                <el-option label="全部适用" value="" />
                <el-option label="设备故障" value="EQUIPMENT" />
                <el-option label="物料缺料" value="MATERIAL" />
                <el-option label="品质异常" value="QUALITY" />
                <el-option label="工艺异常" value="TECHNICAL" />
                <el-option label="工装夹具" value="TOOLING" />
                <el-option label="安全警示" value="SAFETY" />
                <el-option label="其他" value="OTHER" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="升级模式">
              <el-select v-model="formData.escalation_mode">
                <el-option label="超时升级" value="TIMEOUT" />
                <el-option label="手动升级" value="MANUAL" />
                <el-option label="自动升级" value="AUTO" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="最大升级等级">
              <el-input-number v-model="formData.max_escalation_level" :min="1" :max="4" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="是否默认">
              <el-switch v-model="formData.is_default" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-divider content-position="left">L1 第一级</el-divider>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="超时时间" prop="level1_timeout">
              <el-input-number v-model="formData.level1_timeout" :min="1" :max="9999" /> 分钟
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="通知方式">
              <el-select v-model="formData.level1_notify_type">
                <el-option label="工位屏" value="WORKSTATION" />
                <el-option label="大屏" value="SCREEN" />
                <el-option label="推送" value="PUSH" />
                <el-option label="全部" value="ALL" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-divider content-position="left">L2 第二级</el-divider>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="超时时间">
              <el-input-number v-model="formData.level2_timeout" :min="0" :max="9999" /> 分钟
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="通知方式">
              <el-select v-model="formData.level2_notify_type">
                <el-option label="工位屏" value="WORKSTATION" />
                <el-option label="大屏" value="SCREEN" />
                <el-option label="推送" value="PUSH" />
                <el-option label="全部" value="ALL" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-divider content-position="left">L3 第三级</el-divider>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="超时时间">
              <el-input-number v-model="formData.level3_timeout" :min="0" :max="9999" /> 分钟
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="通知方式">
              <el-select v-model="formData.level3_notify_type">
                <el-option label="工位屏" value="WORKSTATION" />
                <el-option label="大屏" value="SCREEN" />
                <el-option label="推送" value="PUSH" />
                <el-option label="全部" value="ALL" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-divider content-position="left">L4 最高级</el-divider>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="超时时间">
              <el-input-number v-model="formData.level4_timeout" :min="0" :max="9999" /> 分钟
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="通知方式">
              <el-select v-model="formData.level4_notify_type">
                <el-option label="工位屏" value="WORKSTATION" />
                <el-option label="大屏" value="SCREEN" />
                <el-option label="推送" value="PUSH" />
                <el-option label="全部" value="ALL" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 详情弹窗 -->
    <el-dialog v-model="detailVisible" title="规则详情" width="600px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="规则编码">{{ detailData.rule_code }}</el-descriptions-item>
        <el-descriptions-item label="规则名称">{{ detailData.rule_name }}</el-descriptions-item>
        <el-descriptions-item label="适用类型">{{ detailData.andon_type_name || '全部' }}</el-descriptions-item>
        <el-descriptions-item label="升级模式">{{ detailData.escalation_mode }}</el-descriptions-item>
        <el-descriptions-item label="最大等级">L{{ detailData.max_escalation_level }}</el-descriptions-item>
        <el-descriptions-item label="是否默认">{{ detailData.is_default === 1 ? '是' : '否' }}</el-descriptions-item>
        <el-descriptions-item label="L1超时">{{ detailData.level1_timeout }}分钟</el-descriptions-item>
        <el-descriptions-item label="L1通知">{{ detailData.level1_notify_type }}</el-descriptions-item>
        <el-descriptions-item label="L2超时">{{ detailData.level2_timeout || '-' }}分钟</el-descriptions-item>
        <el-descriptions-item label="L2通知">{{ detailData.level2_notify_type || '-' }}</el-descriptions-item>
        <el-descriptions-item label="L3超时">{{ detailData.level3_timeout || '-' }}分钟</el-descriptions-item>
        <el-descriptions-item label="L3通知">{{ detailData.level3_notify_type || '-' }}</el-descriptions-item>
        <el-descriptions-item label="L4超时">{{ detailData.level4_timeout || '-' }}分钟</el-descriptions-item>
        <el-descriptions-item label="L4通知">{{ detailData.level4_notify_type || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ detailData.is_enabled === 1 ? '启用' : '停用' }}</el-descriptions-item>
        <el-descriptions-item label="创建人">{{ detailData.created_by }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ detailData.remark }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import {
  getEscalationRuleList,
  getEscalationRuleDetail,
  createEscalationRule,
  updateEscalationRule,
  deleteEscalationRule
} from '@/api/trace'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const detailVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)
const dialogTitle = ref('')

const searchForm = reactive({ rule_name: '', andon_type: '', is_enabled: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const defaultForm = () => ({
  id: 0, rule_code: '', rule_name: '', andon_type: '', escalation_mode: 'TIMEOUT',
  max_escalation_level: 4, is_default: 0, is_enabled: 1,
  level1_timeout: 5, level1_notify_type: 'WORKSTATION',
  level2_timeout: 0, level2_notify_type: '',
  level3_timeout: 0, level3_notify_type: '',
  level4_timeout: 0, level4_notify_type: '',
  remark: ''
})
const formData = reactive(defaultForm())
const detailData = reactive<any>({})

const rules: FormRules = {
  rule_code: [{ required: true, message: '请输入规则编码', trigger: 'blur' }],
  rule_name: [{ required: true, message: '请输入规则名称', trigger: 'blur' }],
  level1_timeout: [{ required: true, message: '请设置L1超时时间', trigger: 'blur' }]
}

const getTypeTag = (type: string) => {
  const map: Record<string, string> = {
    EQUIPMENT: 'danger', MATERIAL: 'warning', QUALITY: 'danger',
    TECHNICAL: 'info', TOOLING: 'info', SAFETY: 'warning', OTHER: 'info'
  }
  return map[type] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => { if (!params[k]) delete params[k] })
    const res = await getEscalationRuleList(params)
    tableData.value = res.data?.list || res.data || []
    pagination.total = res.data?.total || 0
  } catch (e: any) {
    ElMessage.error(e.message || '加载失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { Object.assign(searchForm, { rule_name: '', andon_type: '', is_enabled: '' }); handleSearch() }

const handleAdd = () => {
  Object.assign(formData, defaultForm())
  isEdit.value = false
  dialogTitle.value = '新增升级规则'
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  try {
    const res = await getEscalationRuleDetail(row.id)
    Object.assign(formData, res.data || res)
    isEdit.value = true
    dialogTitle.value = '编辑升级规则'
    dialogVisible.value = true
  } catch (e: any) {
    ElMessage.error(e.message || '加载详情失败')
  }
}

const handleDetail = async (row: any) => {
  try {
    const res = await getEscalationRuleDetail(row.id)
    Object.assign(detailData, res.data || res)
    detailVisible.value = true
  } catch (e: any) {
    ElMessage.error(e.message || '加载详情失败')
  }
}

const handleSubmit = async () => {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    if (isEdit.value) {
      await updateEscalationRule(formData.id, formData)
      ElMessage.success('更新成功')
    } else {
      await createEscalationRule(formData)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '操作失败')
  } finally {
    submitLoading.value = false
  }
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm(`确定删除规则「${row.rule_name}」？`, '删除确认', { type: 'warning' })
  try {
    await deleteEscalationRule(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '删除失败')
  }
}

const handleToggleStatus = async (row: any) => {
  try {
    await updateEscalationRule(row.id, { is_enabled: row.is_enabled })
    ElMessage.success(row.is_enabled ? '已启用' : '已停用')
  } catch (e: any) {
    row.is_enabled = row.is_enabled ? 0 : 1
    ElMessage.error(e.message || '操作失败')
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.escalation-rule-list {
  .search-card { margin-bottom: 16px; }
  .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
