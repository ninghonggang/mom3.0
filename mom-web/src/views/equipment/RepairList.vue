<template>
  <div class="repair-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="设备编码">
          <el-input v-model="searchForm.equipment_code" placeholder="请输入设备编码" clearable />
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.equipment_name" placeholder="请输入设备名称" clearable />
        </el-form-item>
        <el-form-item label="维修状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待派单" :value="1" />
            <el-option label="处理中" :value="2" />
            <el-option label="已完成" :value="3" />
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
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="equipment_code" label="设备编码" width="120" />
        <el-table-column prop="equipment_name" label="设备名称" min-width="150" />
        <el-table-column prop="fault_description" label="故障描述" min-width="200" />
        <el-table-column prop="repair_type" label="维修类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.repair_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="repair_user" label="维修人" width="100" />
        <el-table-column prop="create_time" label="报修时间" width="160" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="warning" v-if="row.status === 1" @click="handleStart(row)">开始维修</el-button>
            <el-button link type="success" v-if="row.status === 2" @click="handleComplete(row)">完成维修</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="设备" prop="equipment_id">
          <el-select v-model="formData.equipment_id" placeholder="请选择设备">
            <el-option label="设备A" :value="1" />
            <el-option label="设备B" :value="2" />
            <el-option label="设备C" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="故障描述" prop="fault_description">
          <el-input v-model="formData.fault_description" type="textarea" />
        </el-form-item>
        <el-form-item label="维修类型" prop="repair_type">
          <el-select v-model="formData.repair_type">
            <el-option label="紧急维修" value="紧急维修" />
            <el-option label="计划维修" value="计划维修" />
            <el-option label="预防性维修" value="预防性维修" />
          </el-select>
        </el-form-item>
        <el-form-item label="维修人">
          <el-input v-model="formData.repair_user" />
        </el-form-item>
        <el-form-item label="维修结果">
          <el-input v-model="formData.repair_result" type="textarea" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" />
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
import { getEquipmentRepairList, createEquipmentRepair, updateEquipmentRepair, deleteEquipmentRepair, startRepair, completeRepair } from '@/api/equipment'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ equipment_code: '', equipment_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive({
  id: 0, equipment_id: 0, equipment_code: '', equipment_name: '', fault_description: '', repair_type: '', repair_user: '', repair_result: '', remark: ''
})

const rules: FormRules = {
  equipment_id: [{ required: true, message: '请选择设备', trigger: 'change' }],
  fault_description: [{ required: true, message: '请输入故障描述', trigger: 'blur' }],
  repair_type: [{ required: true, message: '请选择维修类型', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑维修' : '新增维修')

const getTypeText = (type: string) => {
  const map: Record<string, string> = { '紧急维修': '紧急维修', '计划维修': '计划维修', '预防性维修': '预防性维修' }
  return map[type] || type
}

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待派单', 2: '处理中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getEquipmentRepairList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.equipment_code = ''; searchForm.equipment_name = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, equipment_id: 0, equipment_code: '', equipment_name: '', fault_description: '', repair_type: '', repair_user: '', repair_result: '', remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleStart = async (row: any) => {
  await ElMessageBox.confirm('确定开始维修吗？', '提示', { type: 'warning' })
  await startRepair(row.id)
  ElMessage.success('已开始维修')
  loadData()
}

const handleComplete = async (row: any) => {
  await ElMessageBox.confirm('确定完成维修吗？', '提示', { type: 'warning' })
  await completeRepair(row.id)
  ElMessage.success('维修已完成')
  loadData()
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该维修记录吗？', '提示', { type: 'warning' })
  await deleteEquipmentRepair(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateEquipmentRepair(formData.id, formData) : await createEquipmentRepair(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.repair-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
