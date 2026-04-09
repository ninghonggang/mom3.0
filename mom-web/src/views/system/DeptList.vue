<template>
  <div class="dept-list">
    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('system:dept:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" row-key="id" :tree-props="{ children: 'children' }">
        <el-table-column prop="dept_name" label="部门名称" min-width="180" />
        <el-table-column prop="dept_code" label="部门编码" width="120" />
        <el-table-column prop="leader" label="负责人" width="100" />
        <el-table-column prop="phone" label="联系电话" width="130" />
        <el-table-column prop="sort" label="排序" width="80" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('system:dept:add')" @click="handleAddChild(row)">新增子级</el-button>
            <el-button link type="primary" v-if="hasPermission('system:dept:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('system:dept:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="上级部门" prop="parent_id">
          <el-tree-select
            v-model="formData.parent_id"
            :data="treeData"
            :props="{ label: 'dept_name', value: 'id', children: 'children' }"
            placeholder="请选择上级部门"
            check-strictly
            clearable
            class="w-full"
          />
        </el-form-item>
        <el-form-item label="部门名称" prop="dept_name">
          <el-input v-model="formData.dept_name" />
        </el-form-item>
        <el-form-item label="部门编码" prop="dept_code">
          <el-input v-model="formData.dept_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="负责人" prop="leader">
          <el-input v-model="formData.leader" />
        </el-form-item>
        <el-form-item label="联系电话" prop="phone">
          <el-input v-model="formData.phone" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="formData.sort" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="1">正常</el-radio>
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
import { getDeptTree, createDept, updateDept, deleteDept } from '@/api/system'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const treeData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const formData = reactive({
  id: 0, parent_id: 0, dept_name: '', dept_code: '', leader: '', phone: '', sort: 0, status: 1
})

const rules: FormRules = {
  dept_name: [{ required: true, message: '请输入部门名称', trigger: 'blur' }],
  dept_code: [{ required: true, message: '请输入部门编码', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑部门' : '新增部门')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDeptTree()
    tableData.value = res.data || []
    treeData.value = [{ id: 0, dept_name: '顶级部门', children: res.data || [] }]
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  Object.assign(formData, { id: 0, parent_id: 0, dept_name: '', dept_code: '', leader: '', phone: '', sort: 0, status: 1 })
  dialogVisible.value = true
}

const handleAddChild = (row: any) => {
  Object.assign(formData, { id: 0, parent_id: row.id, dept_name: '', dept_code: '', leader: '', phone: '', sort: 0, status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该部门吗？', '提示', { type: 'warning' })
    await deleteDept(row.id)
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
    formData.id ? await updateDept(formData.id, formData) : await createDept(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.dept-list {
  .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .w-full { width: 100%; }
}
</style>
