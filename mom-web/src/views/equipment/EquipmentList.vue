<template>
  <div class="equipment-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="设备编码">
          <el-input v-model="searchForm.equipment_code" placeholder="请输入设备编码" clearable />
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.equipment_name" placeholder="请输入设备名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="equipment_code" label="设备编码" width="120" />
        <el-table-column prop="equipment_name" label="设备名称" min-width="150" />
        <el-table-column prop="equipment_type" label="设备类型" width="100" />
        <el-table-column prop="brand" label="品牌" width="100" />
        <el-table-column prop="model" label="型号" width="100" />
        <el-table-column prop="workshop_name" label="所属车间" width="100" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="设备编码" prop="equipment_code">
          <el-input v-model="formData.equipment_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="设备名称" prop="equipment_name">
          <el-input v-model="formData.equipment_name" />
        </el-form-item>
        <el-form-item label="设备类型" prop="equipment_type">
          <el-input v-model="formData.equipment_type" />
        </el-form-item>
        <el-form-item label="品牌">
          <el-input v-model="formData.brand" />
        </el-form-item>
        <el-form-item label="型号">
          <el-input v-model="formData.model" />
        </el-form-item>
        <el-form-item label="序列号">
          <el-input v-model="formData.serial_number" />
        </el-form-item>
        <el-form-item label="供应商">
          <el-input v-model="formData.supplier" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="1">运行</el-radio>
            <el-radio :value="2">待机</el-radio>
            <el-radio :value="3">故障</el-radio>
            <el-radio :value="5">报废</el-radio>
          </el-radio-group>
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
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getEquipmentList, createEquipment, updateEquipment, deleteEquipment } from '@/api/equipment'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ equipment_code: '', equipment_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, equipment_code: '', equipment_name: '', equipment_type: '', brand: '', model: '', serial_number: '', supplier: '', status: 1
})

const rules: FormRules = {
  equipment_code: [{ required: true, message: '请输入设备编码', trigger: 'blur' }],
  equipment_name: [{ required: true, message: '请输入设备名称', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑设备' : '新增设备')

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '运行', 2: '待机', 3: '故障', 4: '维修', 5: '报废' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'success', 2: 'info', 3: 'danger', 4: 'warning', 5: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getEquipmentList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.equipment_code = ''; searchForm.equipment_name = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, equipment_code: '', equipment_name: '', equipment_type: '', brand: '', model: '', serial_number: '', supplier: '', status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该设备吗？', '提示', { type: 'warning' })
  await deleteEquipment(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateEquipment(formData.id, formData) : await createEquipment(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.equipment-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
