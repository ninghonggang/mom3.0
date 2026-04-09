<template>
  <div class="workshop-list">
    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:workshop:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="workshop_code" label="车间编码" width="120" />
        <el-table-column prop="workshop_name" label="车间名称" min-width="150" />
        <el-table-column prop="workshop_type" label="车间类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.workshop_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="manager" label="负责人" width="100" />
        <el-table-column prop="phone" label="联系电话" width="130" />
        <el-table-column prop="address" label="地址" min-width="150" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('mdm:workshop:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:workshop:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="车间编码" prop="workshop_code">
          <el-input v-model="formData.workshop_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="车间名称" prop="workshop_name">
          <el-input v-model="formData.workshop_name" />
        </el-form-item>
        <el-form-item label="车间类型" prop="workshop_type">
          <el-select v-model="formData.workshop_type">
            <el-option label="加工" value="加工" />
            <el-option label="装配" value="装配" />
            <el-option label="检验" value="检验" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人">
          <el-input v-model="formData.manager" />
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="formData.phone" />
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="formData.address" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">禁用</el-radio>
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
import { getWorkshopList, createWorkshop, updateWorkshop, deleteWorkshop } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const formData = reactive({
  id: 0, workshop_code: '', workshop_name: '', workshop_type: '', manager: '', phone: '', address: '', status: 1
})

const rules: FormRules = {
  workshop_code: [{ required: true, message: '请输入车间编码', trigger: 'blur' }],
  workshop_name: [{ required: true, message: '请输入车间名称', trigger: 'blur' }],
  workshop_type: [{ required: true, message: '请选择车间类型', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑车间' : '新增车间')

const getTypeText = (type: string) => {
  const map: Record<string, string> = { '加工': '加工', '装配': '装配', '检验': '检验' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getWorkshopList()
    tableData.value = res.data.list || []
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  Object.assign(formData, { id: 0, workshop_code: '', workshop_name: '', workshop_type: '', manager: '', phone: '', address: '', status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该车间吗？', '提示', { type: 'warning' })
    await deleteWorkshop(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateWorkshop(formData.id, formData) : await createWorkshop(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.workshop-list {
  .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
}
</style>
