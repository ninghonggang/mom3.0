<template>
  <div class="first-last-inspect-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="检验类型">
          <el-select v-model="searchForm.inspect_type" placeholder="请选择" clearable>
            <el-option label="首件" value="FIRST" />
            <el-option label="末件" value="LAST" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待检" value="PENDING" />
            <el-option label="已完成" value="COMPLETED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:first-last-inspect:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增检验
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="inspect_no" label="检验单号" min-width="150" />
        <el-table-column prop="inspect_type" label="类型" width="80">
          <template #default="{ row }">
            <el-tag :type="row.inspect_type === 'FIRST' ? 'success' : 'warning'">
              {{ row.inspect_type === 'FIRST' ? '首件' : '末件' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="production_order_id" label="工单ID" width="100" />
        <el-table-column prop="process_id" label="工序ID" width="80" />
        <el-table-column prop="workstation_id" label="工位ID" width="80" />
        <el-table-column prop="inspector_name" label="检验员" width="100" />
        <el-table-column prop="overall_result" label="结果" width="80">
          <template #default="{ row }">
            <el-tag :type="row.overall_result === 'OK' ? 'success' : row.overall_result === 'NG' ? 'danger' : 'info'">
              {{ row.overall_result === 'OK' ? '合格' : row.overall_result === 'NG' ? '不合格' : '待检' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'COMPLETED' ? 'success' : 'info'">
              {{ row.status === 'COMPLETED' ? '已完成' : '待检' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('production:first-last-inspect:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('production:first-last-inspect:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 检验对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px" @close="handleDialogClose">
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="100px">
        <el-form-item label="工单ID" prop="production_order_id">
          <el-input-number v-model="form.production_order_id" :min="1" placeholder="请输入工单ID" style="width: 100%" />
        </el-form-item>
        <el-form-item label="工序ID" prop="process_id">
          <el-input-number v-model="form.process_id" :min="1" placeholder="请输入工序ID" style="width: 100%" />
        </el-form-item>
        <el-form-item label="工位ID" prop="workstation_id">
          <el-input-number v-model="form.workstation_id" :min="1" placeholder="请输入工位ID" style="width: 100%" />
        </el-form-item>
        <el-form-item label="检验类型" prop="inspect_type">
          <el-select v-model="form.inspect_type" placeholder="请选择检验类型" style="width: 100%">
            <el-option label="首件" value="FIRST" />
            <el-option label="末件" value="LAST" />
          </el-select>
        </el-form-item>
        <el-form-item label="检验员" prop="inspector_name">
          <el-input v-model="form.inspector_name" placeholder="请输入检验员姓名" />
        </el-form-item>
        <el-form-item label="总体结果" prop="overall_result">
          <el-select v-model="form.overall_result" placeholder="请选择结果" style="width: 100%">
            <el-option label="合格" value="OK" />
            <el-option label="不合格" value="NG" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" :rows="3" placeholder="请输入备注" />
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
import { getFirstLastInspectList, getFirstLastInspect, createFirstLastInspect, updateFirstLastInspect, deleteFirstLastInspect } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const searchForm = reactive({ order_no: '', inspect_type: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('')
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const form = reactive({
  id: undefined as number | undefined,
  production_order_id: undefined as number | undefined,
  process_id: undefined as number | undefined,
  workstation_id: undefined as number | undefined,
  inspect_type: '',
  inspector_name: '',
  overall_result: '',
  remark: ''
})

const formRules: FormRules = {
  production_order_id: [{ required: true, message: '请输入工单ID', trigger: 'blur' }],
  inspect_type: [{ required: true, message: '请选择检验类型', trigger: 'change' }],
  inspector_name: [{ required: true, message: '请输入检验员', trigger: 'blur' }],
  overall_result: [{ required: true, message: '请选择检验结果', trigger: 'change' }]
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => params[k] === '' && delete params[k])
    const res = await getFirstLastInspectList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.order_no = ''; searchForm.inspect_type = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增检验'
  resetForm()
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑检验'
  try {
    const res = await getFirstLastInspect(row.id)
    Object.assign(form, res.data)
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取检验信息失败')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    submitLoading.value = true
    try {
      if (isEdit.value) {
        await updateFirstLastInspect(form.id!, form)
        ElMessage.success('更新成功')
      } else {
        await createFirstLastInspect(form)
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
    await ElMessageBox.confirm('确定删除该检验单吗？', '提示', { type: 'warning' })
    await deleteFirstLastInspect(row.id)
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
  form.production_order_id = undefined
  form.process_id = undefined
  form.workstation_id = undefined
  form.inspect_type = ''
  form.inspector_name = ''
  form.overall_result = ''
  form.remark = ''
}

const handleDialogClose = () => {
  formRef.value?.resetFields()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.first-last-inspect-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
