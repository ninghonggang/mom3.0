<template>
  <div class="shift-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="班次编码">
          <el-input v-model="searchForm.shift_code" placeholder="请输入班次编码" clearable />
        </el-form-item>
        <el-form-item label="班次名称">
          <el-input v-model="searchForm.shift_name" placeholder="请输入班次名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增班次
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="shift_code" label="班次编码" width="120" />
        <el-table-column prop="shift_name" label="班次名称" min-width="120" />
        <el-table-column prop="start_time" label="开始时间" width="100" />
        <el-table-column prop="end_time" label="结束时间" width="100" />
        <el-table-column prop="work_hours" label="工作时长" width="100">
          <template #default="{ row }">
            {{ row.work_hours }}小时
          </template>
        </el-table-column>
        <el-table-column prop="is_night" label="夜班" width="80">
          <template #default="{ row }">
            <el-tag :type="row.is_night === 1 ? 'warning' : 'info'" size="small">
              {{ row.is_night === 1 ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="班次编码" prop="shift_code">
          <el-input v-model="formData.shift_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="班次名称" prop="shift_name">
          <el-input v-model="formData.shift_name" />
        </el-form-item>
        <el-form-item label="开始时间" prop="start_time">
          <el-time-picker v-model="formData.start_time" format="HH:mm" value-format="HH:mm" style="width: 100%" />
        </el-form-item>
        <el-form-item label="结束时间" prop="end_time">
          <el-time-picker v-model="formData.end_time" format="HH:mm" value-format="HH:mm" style="width: 100%" />
        </el-form-item>
        <el-form-item label="工作时长">
          <el-input-number v-model="formData.work_hours" :min="0" :max="24" :precision="1" style="width: 200px" /> 小时
        </el-form-item>
        <el-form-item label="夜班">
          <el-radio-group v-model="formData.is_night">
            <el-radio :value="1">是</el-radio>
            <el-radio :value="0">否</el-radio>
          </el-radio-group>
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
import { getMdmShiftList, createMdmShift, updateMdmShift, deleteMdmShift } from '@/api/mdm'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const searchForm = reactive({ shift_code: '', shift_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0,
  shift_code: '',
  shift_name: '',
  start_time: '08:00',
  end_time: '17:00',
  work_hours: 8.0,
  is_night: 0,
  remark: ''
})

const rules: FormRules = {
  shift_code: [{ required: true, message: '请输入班次编码', trigger: 'blur' }],
  shift_name: [{ required: true, message: '请输入班次名称', trigger: 'blur' }],
  start_time: [{ required: true, message: '请选择开始时间', trigger: 'change' }],
  end_time: [{ required: true, message: '请选择结束时间', trigger: 'change' }]
}

const dialogTitle = computed(() => isEdit.value ? '编辑班次' : '新增班次')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMdmShiftList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.shift_code = ''; searchForm.shift_name = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  Object.assign(formData, {
    id: 0, shift_code: '', shift_name: '', start_time: '08:00', end_time: '17:00', work_hours: 8.0, is_night: 0, remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该班次吗？', '提示', { type: 'warning' })
  await deleteMdmShift(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (isEdit.value) {
      await updateMdmShift(formData.id, formData)
      ElMessage.success('更新成功')
    } else {
      await createMdmShift(formData)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.shift-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
