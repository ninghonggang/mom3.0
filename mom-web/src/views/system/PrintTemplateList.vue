<template>
  <div class="print-template-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="模板编码">
          <el-input v-model="searchForm.template_code" placeholder="请输入模板编码" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('system:print-template:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增模板
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="template_code" label="模板编码" min-width="120" />
        <el-table-column prop="template_name" label="模板名称" min-width="150" />
        <el-table-column prop="template_type" label="模板类型" width="130">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.template_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="entity_type" label="适用单据" width="130" />
        <el-table-column prop="paper_type" label="纸张类型" width="100" />
        <el-table-column prop="paper_width" label="宽度(mm)" width="90" />
        <el-table-column prop="paper_height" label="高度(mm)" width="90" />
        <el-table-column prop="is_default" label="默认" width="70">
          <template #default="{ row }">
            <el-tag :type="row.is_default === 1 ? 'success' : 'info'" size="small">
              {{ row.is_default === 1 ? '是' : '否' }}
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
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('system:print-template:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('system:print-template:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px" @close="handleDialogClose">
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="模板编码" prop="template_code">
              <el-input v-model="form.template_code" placeholder="请输入模板编码" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="模板名称" prop="template_name">
              <el-input v-model="form.template_name" placeholder="请输入模板名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="模板类型" prop="template_type">
              <el-select v-model="form.template_type" placeholder="请选择" style="width: 100%">
                <el-option label="生产工单" value="PRODUCTION_ORDER" />
                <el-option label="包装" value="PACKAGE" />
                <el-option label="标签" value="LABEL" />
                <el-option label="报表" value="REPORT" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="适用单据" prop="entity_type">
              <el-input v-model="form.entity_type" placeholder="如: pro_production_order" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="纸张类型" prop="paper_type">
              <el-select v-model="form.paper_type" placeholder="请选择" style="width: 100%">
                <el-option label="A4" value="A4" />
                <el-option label="A5" value="A5" />
                <el-option label="热敏纸" value="THERMAL" />
                <el-option label="自定义" value="CUSTOM" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="设为默认" prop="is_default">
              <el-switch v-model="form.is_default" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="宽度(mm)" prop="paper_width">
              <el-input-number v-model="form.paper_width" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="高度(mm)" prop="paper_height">
              <el-input-number v-model="form.paper_height" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="模板内容" prop="content">
          <el-input v-model="form.content" type="textarea" :rows="6" placeholder="模板内容(JSON/ZPL/HTML)" />
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getPrintTemplateList, getPrintTemplate, createPrintTemplate, updatePrintTemplate, deletePrintTemplate } from '@/api/system'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const searchForm = reactive({ template_code: '' })

const dialogVisible = ref(false)
const dialogTitle = ref('')
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const form = reactive({
  id: undefined as number | undefined,
  template_code: '',
  template_name: '',
  template_type: '',
  entity_type: '',
  content: '',
  paper_type: '',
  paper_width: 0,
  paper_height: 0,
  is_default: 0,
  status: 1,
  remark: ''
})

const formRules: FormRules = {
  template_code: [{ required: true, message: '请输入模板编码', trigger: 'blur' }],
  template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }],
  template_type: [{ required: true, message: '请选择模板类型', trigger: 'change' }]
}

const getTypeText = (type: string) => {
  const map: Record<string, string> = { PRODUCTION_ORDER: '生产工单', PACKAGE: '包装', LABEL: '标签', REPORT: '报表' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getPrintTemplateList()
    tableData.value = res.data.list || []
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { loadData() }
const handleReset = () => { searchForm.template_code = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增模板'
  resetForm()
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑模板'
  try {
    const res = await getPrintTemplate(row.id)
    Object.assign(form, res.data)
    dialogVisible.value = true
  } catch {
    ElMessage.error('获取模板信息失败')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    submitLoading.value = true
    try {
      if (isEdit.value) {
        await updatePrintTemplate(form.id!, form)
        ElMessage.success('更新成功')
      } else {
        await createPrintTemplate(form)
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
    await ElMessageBox.confirm('确定删除该模板吗？', '提示', { type: 'warning' })
    await deletePrintTemplate(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

const resetForm = () => {
  form.id = undefined
  form.template_code = ''
  form.template_name = ''
  form.template_type = ''
  form.entity_type = ''
  form.content = ''
  form.paper_type = ''
  form.paper_width = 0
  form.paper_height = 0
  form.is_default = 0
  form.status = 1
  form.remark = ''
}

const handleDialogClose = () => {
  formRef.value?.resetFields()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.print-template-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
}
</style>
