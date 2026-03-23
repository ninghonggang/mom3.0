<template>
  <div class="operation-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工序编码">
          <el-input v-model="searchForm.operation_code" placeholder="请输入工序编码" clearable />
        </el-form-item>
        <el-form-item label="工序名称">
          <el-input v-model="searchForm.operation_name" placeholder="请输入工序名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增工序
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="operation_code" label="工序编码" width="120" />
        <el-table-column prop="operation_name" label="工序名称" min-width="150" />
        <el-table-column prop="workcenter_name" label="工作中心" width="120" />
        <el-table-column prop="standard_worktime" label="标准工时(分钟)" width="130" />
        <el-table-column prop="is_key_process" label="关键工序" width="100">
          <template #default="{ row }">
            <el-tag :type="row.is_key_process === 1 ? 'danger' : 'info'" size="small">
              {{ row.is_key_process === 1 ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="is_qc_point" label="质检点" width="80">
          <template #default="{ row }">
            <el-tag :type="row.is_qc_point === 1 ? 'warning' : 'info'" size="small">
              {{ row.is_qc_point === 1 ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sequence" label="顺序号" width="80" />
        <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
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

    <!-- 编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="工序编码" prop="operation_code">
          <el-input v-model="formData.operation_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="工序名称" prop="operation_name">
          <el-input v-model="formData.operation_name" />
        </el-form-item>
        <el-form-item label="工作中心">
          <el-input v-model="formData.workcenter_name" />
        </el-form-item>
        <el-form-item label="标准工时">
          <el-input-number v-model="formData.standard_worktime" :min="0" style="width: 200px" /> 分钟
        </el-form-item>
        <el-form-item label="质量标准">
          <el-input v-model="formData.quality_std" type="textarea" :rows="2" />
        </el-form-item>
        <el-form-item label="关键工序">
          <el-radio-group v-model="formData.is_key_process">
            <el-radio :value="1">是</el-radio>
            <el-radio :value="0">否</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="质检点">
          <el-radio-group v-model="formData.is_qc_point">
            <el-radio :value="1">是</el-radio>
            <el-radio :value="0">否</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="顺序号">
          <el-input-number v-model="formData.sequence" :min="0" style="width: 200px" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" />
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
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getOperationList, createOperation, updateOperation, deleteOperation } from '@/api/mdm'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const searchForm = reactive({ operation_code: '', operation_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0,
  operation_code: '',
  operation_name: '',
  workcenter_id: null as number | null,
  workcenter_name: '',
  standard_worktime: 0,
  quality_std: '',
  is_key_process: 0,
  is_qc_point: 0,
  sequence: 0,
  remark: ''
})

const rules: FormRules = {
  operation_code: [{ required: true, message: '请输入工序编码', trigger: 'blur' }],
  operation_name: [{ required: true, message: '请输入工序名称', trigger: 'blur' }]
}

const dialogTitle = computed(() => isEdit.value ? '编辑工序' : '新增工序')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getOperationList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.operation_code = ''; searchForm.operation_name = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  Object.assign(formData, {
    id: 0, operation_code: '', operation_name: '', workcenter_id: null, workcenter_name: '',
    standard_worktime: 0, quality_std: '', is_key_process: 0, is_qc_point: 0, sequence: 0, remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该工序吗？', '提示', { type: 'warning' })
  await deleteOperation(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (isEdit.value) {
      await updateOperation(formData.id, formData)
      ElMessage.success('更新成功')
    } else {
      await createOperation(formData)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.operation-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
