<template>
  <div class="electronic-sop-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="SOP编号">
          <el-input v-model="searchForm.query" placeholder="请输入SOP编号/名称" clearable />
        </el-form-item>
        <el-form-item label="内容类型">
          <el-select v-model="searchForm.content_type" placeholder="请选择" clearable>
            <el-option label="PDF" value="PDF" />
            <el-option label="视频" value="VIDEO" />
            <el-option label="图片" value="IMAGE" />
            <el-option label="HTML" value="HTML" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" value="1" />
            <el-option label="已发布" value="2" />
            <el-option label="已作废" value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:electronic-sop:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增SOP
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="sop_no" label="SOP编号" min-width="120" />
        <el-table-column prop="sop_name" label="SOP名称" min-width="180" />
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="process_name" label="工序" width="100" />
        <el-table-column prop="version" label="版本" width="80" />
        <el-table-column prop="content_type" label="内容类型" width="100">
          <template #default="{ row }">
            <el-tag :type="getContentTypeTag(row.content_type)">
              {{ getContentTypeText(row.content_type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusTag(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('production:electronic-sop:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('production:electronic-sop:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px" @close="handleDialogClose">
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="SOP编号" prop="sop_no">
              <el-input v-model="form.sop_no" placeholder="请输入SOP编号" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SOP名称" prop="sop_name">
              <el-input v-model="form.sop_name" placeholder="请输入SOP名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="物料" prop="material_code">
              <el-input v-model="form.material_code" placeholder="物料编码" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="版本" prop="version">
              <el-input v-model="form.version" placeholder="如V1.0" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="工序" prop="process_name">
              <el-input v-model="form.process_name" placeholder="工序名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="内容类型" prop="content_type">
              <el-select v-model="form.content_type" placeholder="请选择" style="width: 100%">
                <el-option label="PDF" value="PDF" />
                <el-option label="视频" value="VIDEO" />
                <el-option label="图片" value="IMAGE" />
                <el-option label="HTML" value="HTML" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="内容URL" prop="content_url">
          <el-input v-model="form.content_url" placeholder="请输入内容URL" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="工位" prop="workstation_name">
              <el-input v-model="form.workstation_name" placeholder="工位名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="车间" prop="workshop_name">
              <el-input v-model="form.workshop_name" placeholder="车间名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-select v-model="form.status" placeholder="请选择" style="width: 100%">
                <el-option label="草稿" :value="1" />
                <el-option label="已发布" :value="2" />
                <el-option label="已作废" :value="3" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
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
import { getElectronicSOPList, getElectronicSOP, createElectronicSOP, updateElectronicSOP, deleteElectronicSOP } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const searchForm = reactive({ query: '', content_type: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const dialogVisible = ref(false)
const dialogTitle = ref('')
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const form = reactive({
  id: undefined as number | undefined,
  sop_no: '',
  sop_name: '',
  material_id: undefined as number | undefined,
  material_code: '',
  material_name: '',
  version: '',
  process_id: undefined as number | undefined,
  process_name: '',
  content_type: '',
  content_url: '',
  thumbnail_url: '',
  workstation_id: undefined as number | undefined,
  workstation_name: '',
  workshop_id: undefined as number | undefined,
  workshop_name: '',
  status: 1,
  remark: ''
})

const formRules: FormRules = {
  sop_no: [{ required: true, message: '请输入SOP编号', trigger: 'blur' }],
  sop_name: [{ required: true, message: '请输入SOP名称', trigger: 'blur' }],
  content_type: [{ required: true, message: '请选择内容类型', trigger: 'change' }]
}

const getContentTypeTag = (type: string) => {
  const map: Record<string, string> = { PDF: 'danger', VIDEO: 'warning', IMAGE: 'success', HTML: 'info' }
  return map[type] || 'info'
}

const getContentTypeText = (type: string) => {
  const map: Record<string, string> = { PDF: 'PDF', VIDEO: '视频', IMAGE: '图片', HTML: 'HTML' }
  return map[type] || type
}

const getStatusTag = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'success', 3: 'danger' }
  return map[status] || 'info'
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '草稿', 2: '已发布', 3: '已作废' }
  return map[status] || '未知'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => params[k] === '' && delete params[k])
    const res = await getElectronicSOPList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.query = ''; searchForm.content_type = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增SOP'
  resetForm()
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑SOP'
  try {
    const res = await getElectronicSOP(row.id)
    Object.assign(form, res.data)
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取SOP信息失败')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    submitLoading.value = true
    try {
      if (isEdit.value) {
        await updateElectronicSOP(form.id!, form)
        ElMessage.success('更新成功')
      } else {
        await createElectronicSOP(form)
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
    await ElMessageBox.confirm('确定删除该SOP吗？', '提示', { type: 'warning' })
    await deleteElectronicSOP(row.id)
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
  form.sop_no = ''
  form.sop_name = ''
  form.material_id = undefined
  form.material_code = ''
  form.material_name = ''
  form.version = ''
  form.process_id = undefined
  form.process_name = ''
  form.content_type = ''
  form.content_url = ''
  form.thumbnail_url = ''
  form.workstation_id = undefined
  form.workstation_name = ''
  form.workshop_id = undefined
  form.workshop_name = ''
  form.status = 1
  form.remark = ''
}

const handleDialogClose = () => {
  formRef.value?.resetFields()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.electronic-sop-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
