<template>
  <div class="notice-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="标题">
          <el-input v-model="searchForm.query" placeholder="请输入标题" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="草稿" :value="1" />
            <el-option label="已发布" :value="2" />
            <el-option label="已撤回" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('system:notice:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>发布公告
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="title" label="标题" min-width="200" show-overflow-tooltip />
        <el-table-column prop="notice_type" label="类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.notice_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="priority" label="优先级" width="80">
          <template #default="{ row }">
            <el-tag :type="row.priority === 3 ? 'danger' : row.priority === 2 ? 'warning' : 'info'">
              {{ row.priority === 3 ? '紧急' : row.priority === 2 ? '重要' : '普通' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="publisher_name" label="发布人" width="100" />
        <el-table-column prop="publish_time" label="发布时间" width="160" />
        <el-table-column prop="is_top" label="置顶" width="70">
          <template #default="{ row }">
            <el-tag :type="row.is_top === 1 ? 'danger' : 'info'" size="small">
              {{ row.is_top === 1 ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="view_count" label="阅读" width="70" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusTag(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('system:notice:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" v-if="row.status === 1 && hasPermission('system:notice:publish')" @click="handlePublish(row)">发布</el-button>
            <el-button link type="danger" v-if="hasPermission('system:notice:delete')" @click="handleDelete(row)">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="650px" @close="handleDialogClose">
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="90px">
        <el-form-item label="标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入标题" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="类型" prop="notice_type">
              <el-select v-model="form.notice_type" placeholder="请选择" style="width: 100%">
                <el-option label="系统通知" value="SYSTEM" />
                <el-option label="运行通知" value="OPERATION" />
                <el-option label="维护通知" value="MAINTENANCE" />
                <el-option label="其他" value="OTHER" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="优先级" prop="priority">
              <el-select v-model="form.priority" placeholder="请选择" style="width: 100%">
                <el-option label="普通" :value="1" />
                <el-option label="重要" :value="2" />
                <el-option label="紧急" :value="3" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="发布部门" prop="publish_dept">
              <el-input v-model="form.publish_dept" placeholder="发布部门" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="置顶" prop="is_top">
              <el-switch v-model="form.is_top" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="内容" prop="content">
          <el-input v-model="form.content" type="textarea" :rows="4" placeholder="请输入内容" />
        </el-form-item>
        <el-form-item label="发布范围" prop="target_type">
          <el-radio-group v-model="form.target_type">
            <el-radio value="ALL">全部</el-radio>
            <el-radio value="DEPT">部门</el-radio>
            <el-radio value="ROLE">角色</el-radio>
            <el-radio value="USER">用户</el-radio>
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
import { getNoticeList, getNotice, createNotice, updateNotice, deleteNotice, publishNotice } from '@/api/system'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const searchForm = reactive({ query: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const dialogVisible = ref(false)
const dialogTitle = ref('')
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const form = reactive({
  id: undefined as number | undefined,
  title: '',
  content: '',
  notice_type: 'SYSTEM',
  priority: 1,
  publish_dept: '',
  publisher_id: undefined as number | undefined,
  publisher_name: '',
  target_type: 'ALL',
  target_ids: '',
  is_top: 0,
  status: 1,
  remark: ''
})

const formRules: FormRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  notice_type: [{ required: true, message: '请选择类型', trigger: 'change' }]
}

const getTypeText = (type: string) => {
  const map: Record<string, string> = { SYSTEM: '系统', OPERATION: '运行', MAINTENANCE: '维护', OTHER: '其他' }
  return map[type] || type
}

const getStatusTag = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'success', 3: 'warning' }
  return map[status] || 'info'
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '草稿', 2: '已发布', 3: '已撤回' }
  return map[status] || '未知'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => params[k] === '' && delete params[k])
    const res = await getNoticeList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.query = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '发布公告'
  resetForm()
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑公告'
  try {
    const res = await getNotice(row.id)
    Object.assign(form, res.data)
    dialogVisible.value = true
  } catch {
    ElMessage.error('获取公告信息失败')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    submitLoading.value = true
    try {
      if (isEdit.value) {
        await updateNotice(form.id!, form)
        ElMessage.success('更新成功')
      } else {
        await createNotice(form)
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

const handlePublish = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定发布该公告吗？', '提示', { type: 'warning' })
    await publishNotice(row.id)
    ElMessage.success('发布成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '发布失败')
    }
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该公告吗？', '提示', { type: 'warning' })
    await deleteNotice(row.id)
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
  form.title = ''
  form.content = ''
  form.notice_type = 'SYSTEM'
  form.priority = 1
  form.publish_dept = ''
  form.publisher_id = undefined
  form.publisher_name = ''
  form.target_type = 'ALL'
  form.target_ids = ''
  form.is_top = 0
  form.status = 1
  form.remark = ''
}

const handleDialogClose = () => {
  formRef.value?.resetFields()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.notice-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
