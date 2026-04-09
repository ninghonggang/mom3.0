<template>
  <div class="flow-card-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="指示单号">
          <el-input v-model="searchForm.query" placeholder="请输入指示单号/工单" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待生产" :value="1" />
            <el-option label="生产中" :value="2" />
            <el-option label="已完成" :value="3" />
            <el-option label="已取消" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:flow-card:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增指示单
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="card_no" label="指示单号" min-width="130" />
        <el-table-column prop="order_no" label="工单编号" width="130" />
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="workshop_name" label="车间" width="100" />
        <el-table-column prop="line_name" label="产线" width="100" />
        <el-table-column prop="process_name" label="当前工序" width="100" />
        <el-table-column prop="station_name" label="工位" width="100" />
        <el-table-column prop="plan_qty" label="计划数量" width="100" />
        <el-table-column prop="completed_qty" label="已完成" width="100" />
        <el-table-column prop="priority" label="优先级" width="80">
          <template #default="{ row }">
            <el-tag :type="row.priority === 3 ? 'danger' : row.priority === 2 ? 'warning' : 'info'">
              {{ row.priority === 3 ? '加急' : row.priority === 2 ? '紧急' : '普通' }}
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
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('production:flow-card:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('production:flow-card:delete')" @click="handleDelete(row)">删除</el-button>
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
            <el-form-item label="指示单号" prop="card_no">
              <el-input v-model="form.card_no" placeholder="请输入指示单号" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="工单编号" prop="order_no">
              <el-input v-model="form.order_no" placeholder="工单编号" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="物料编码" prop="material_code">
              <el-input v-model="form.material_code" placeholder="物料编码" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="物料名称" prop="material_name">
              <el-input v-model="form.material_name" placeholder="物料名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="车间" prop="workshop_name">
              <el-input v-model="form.workshop_name" placeholder="车间名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="产线" prop="line_name">
              <el-input v-model="form.line_name" placeholder="产线名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="当前工序" prop="process_name">
              <el-input v-model="form.process_name" placeholder="工序名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="工位" prop="station_name">
              <el-input v-model="form.station_name" placeholder="工位名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="计划数量" prop="plan_qty">
              <el-input-number v-model="form.plan_qty" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="优先级" prop="priority">
              <el-select v-model="form.priority" placeholder="请选择" style="width: 100%">
                <el-option label="普通" :value="1" />
                <el-option label="紧急" :value="2" />
                <el-option label="加急" :value="3" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-select v-model="form.status" placeholder="请选择" style="width: 100%">
                <el-option label="待生产" :value="1" />
                <el-option label="生产中" :value="2" />
                <el-option label="已完成" :value="3" />
                <el-option label="已取消" :value="4" />
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
import { getFlowCardList, getFlowCard, createFlowCard, updateFlowCard, deleteFlowCard } from '@/api/production'
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
  card_no: '',
  order_id: undefined as number | undefined,
  order_no: '',
  material_id: undefined as number | undefined,
  material_code: '',
  material_name: '',
  workshop_id: undefined as number | undefined,
  workshop_name: '',
  line_id: undefined as number | undefined,
  line_name: '',
  process_id: undefined as number | undefined,
  process_name: '',
  station_id: undefined as number | undefined,
  station_name: '',
  plan_qty: 0,
  completed_qty: 0,
  status: 1,
  priority: 1,
  remark: ''
})

const formRules: FormRules = {
  card_no: [{ required: true, message: '请输入指示单号', trigger: 'blur' }]
}

const getStatusTag = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success', 4: 'danger' }
  return map[status] || 'info'
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待生产', 2: '生产中', 3: '已完成', 4: '已取消' }
  return map[status] || '未知'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => params[k] === '' && delete params[k])
    const res = await getFlowCardList(params)
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
  dialogTitle.value = '新增指示单'
  resetForm()
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑指示单'
  try {
    const res = await getFlowCard(row.id)
    Object.assign(form, res.data.card || res.data)
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取指示单信息失败')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    submitLoading.value = true
    try {
      if (isEdit.value) {
        await updateFlowCard(form.id!, form)
        ElMessage.success('更新成功')
      } else {
        await createFlowCard(form)
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
    await ElMessageBox.confirm('确定删除该指示单吗？', '提示', { type: 'warning' })
    await deleteFlowCard(row.id)
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
  form.card_no = ''
  form.order_id = undefined
  form.order_no = ''
  form.material_id = undefined
  form.material_code = ''
  form.material_name = ''
  form.workshop_id = undefined
  form.workshop_name = ''
  form.line_id = undefined
  form.line_name = ''
  form.process_id = undefined
  form.process_name = ''
  form.station_id = undefined
  form.station_name = ''
  form.plan_qty = 0
  form.completed_qty = 0
  form.status = 1
  form.priority = 1
  form.remark = ''
}

const handleDialogClose = () => {
  formRef.value?.resetFields()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.flow-card-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
