<template>
  <div class="material-category-list">
    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:materialcategory:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" row-key="id" :tree-props="{ children: 'children' }" default-expand-all>
        <el-table-column prop="category_name" label="分类名称" min-width="180" />
        <el-table-column prop="category_code" label="分类编码" width="150" />
        <el-table-column prop="sort" label="排序" width="80" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('mdm:materialcategory:add')" @click="handleAddChild(row)">新增子级</el-button>
            <el-button link type="primary" v-if="hasPermission('mdm:materialcategory:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:materialcategory:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="上级分类" prop="parent_id">
          <el-tree-select
            v-model="formData.parent_id"
            :data="treeData"
            :props="{ label: 'category_name', value: 'id', children: 'children' }"
            placeholder="请选择上级分类"
            check-strictly
            clearable
            class="w-full"
          />
        </el-form-item>
        <el-form-item label="分类名称" prop="category_name">
          <el-input v-model="formData.category_name" />
        </el-form-item>
        <el-form-item label="分类编码" prop="category_code">
          <el-input v-model="formData.category_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="formData.sort" :min="0" />
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
import { getMaterialCategoryTree, createMaterialCategory, updateMaterialCategory, deleteMaterialCategory } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const treeData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const formData = reactive({
  id: 0, parent_id: 0, category_name: '', category_code: '', sort: 0, status: 1
})

const rules: FormRules = {
  category_name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
  category_code: [{ required: true, message: '请输入分类编码', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑分类' : '新增分类')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMaterialCategoryTree()
    tableData.value = res.data || []
    treeData.value = [{ id: 0, category_name: '顶级分类', children: res.data || [] }]
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  Object.assign(formData, { id: 0, parent_id: 0, category_name: '', category_code: '', sort: 0, status: 1 })
  dialogVisible.value = true
}

const handleAddChild = (row: any) => {
  Object.assign(formData, { id: 0, parent_id: row.id, category_name: '', category_code: '', sort: 0, status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该分类吗？', '提示', { type: 'warning' })
    await deleteMaterialCategory(row.id)
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
    formData.id ? await updateMaterialCategory(formData.id, formData) : await createMaterialCategory(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.material-category-list {
  .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .w-full { width: 100%; }
}
</style>
