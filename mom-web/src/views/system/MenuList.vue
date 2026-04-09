<template>
  <div class="menu-list">
    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('system:menu:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" row-key="id" :tree-props="{ children: 'children' }">
        <el-table-column prop="menu_name" label="菜单名称" min-width="150" />
        <el-table-column prop="icon" label="图标" width="80">
          <template #default="{ row }">
            <el-icon v-if="getIcon(row.icon)"><component :is="getIcon(row.icon)" /></el-icon>
          </template>
        </el-table-column>
        <el-table-column prop="menu_type" label="类型" width="80">
          <template #default="{ row }">
            <el-tag :type="getMenuTypeTag(row.menu_type)">{{ getMenuTypeText(row.menu_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="path" label="路由" min-width="150" />
        <el-table-column prop="component" label="组件" min-width="150" />
        <el-table-column prop="perms" label="权限标识" min-width="150" />
        <el-table-column prop="sort" label="排序" width="80" />
        <el-table-column prop="visible" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.visible === 1 ? 'success' : 'danger'">
              {{ row.visible === 1 ? '显示' : '隐藏' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('system:menu:add')" @click="handleAddChild(row)">新增子级</el-button>
            <el-button link type="primary" v-if="hasPermission('system:menu:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('system:menu:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="上级菜单" prop="parent_id">
          <el-tree-select
            v-model="formData.parent_id"
            :data="treeData"
            :props="{ label: 'menu_name', value: 'id', children: 'children' }"
            placeholder="请选择上级菜单"
            check-strictly
            clearable
            class="w-full"
          />
        </el-form-item>
        <el-form-item label="菜单类型" prop="menu_type">
          <el-radio-group v-model="formData.menu_type">
            <el-radio value="M">目录</el-radio>
            <el-radio value="C">菜单</el-radio>
            <el-radio value="F">按钮</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="菜单名称" prop="menu_name">
          <el-input v-model="formData.menu_name" />
        </el-form-item>
        <el-form-item label="图标" v-if="formData.menu_type !== 'F'">
          <el-input v-model="formData.icon" placeholder="如: User, Home" />
        </el-form-item>
        <el-form-item label="路由" v-if="formData.menu_type !== 'F'" prop="path">
          <el-input v-model="formData.path" />
        </el-form-item>
        <el-form-item label="组件" v-if="formData.menu_type === 'C'" prop="component">
          <el-input v-model="formData.component" placeholder="如: system/UserList" />
        </el-form-item>
        <el-form-item label="权限标识" v-if="formData.menu_type === 'F'" prop="perms">
          <el-input v-model="formData.perms" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="formData.sort" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="visible">
          <el-radio-group v-model="formData.visible">
            <el-radio :value="1">显示</el-radio>
            <el-radio :value="0">隐藏</el-radio>
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
import { getMenuTree, createMenu, updateMenu, deleteMenu } from '@/api/system'
import { useAuthStore } from '@/stores/auth'
import {
  House, Setting, Box, List, Monitor, CircleCheck, Calendar, Search, Lightning,
  Bell, Grid, Folder, Key, Document, Connection, DataBoard, DataLine
} from '@element-plus/icons-vue'

const iconMap: Record<string, any> = {
  House, Setting, Box, List, Monitor, CircleCheck, Calendar, Search, Lightning,
  Bell, Grid, Folder, Key, Document, Connection, DataBoard, DataLine
}
const getIcon = (iconName?: string) => iconName ? (iconMap[iconName] || null) : null

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const treeData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const formData = reactive({
  id: 0, parent_id: 0, menu_name: '', menu_type: 'M', path: '',
  component: '', perms: '', icon: '', sort: 0, visible: 1
})

const rules: FormRules = {
  menu_name: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
  menu_type: [{ required: true, message: '请选择菜单类型', trigger: 'change' }],
  path: [{ required: true, message: '请输入路由路径', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑菜单' : '新增菜单')

const getMenuTypeText = (type: string) => {
  const map: Record<string, string> = { M: '目录', C: '菜单', F: '按钮' }
  return map[type] || type
}

const getMenuTypeTag = (type: string) => {
  const map: Record<string, string> = { M: 'warning', C: 'primary', F: 'info' }
  return map[type] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMenuTree()
    tableData.value = res.data || []
    treeData.value = [{ id: 0, menu_name: '顶级菜单', children: res.data || [] }]
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  Object.assign(formData, { id: 0, parent_id: 0, menu_name: '', menu_type: 'M', path: '', component: '', perms: '', icon: '', sort: 0, visible: 1 })
  dialogVisible.value = true
}

const handleAddChild = (row: any) => {
  Object.assign(formData, { id: 0, parent_id: row.id, menu_name: '', menu_type: 'C', path: '', component: '', perms: '', icon: '', sort: 0, visible: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该菜单吗？', '提示', { type: 'warning' })
    await deleteMenu(row.id)
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
    formData.id ? await updateMenu(formData.id, formData) : await createMenu(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.menu-list {
  .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .w-full { width: 100%; }
}
</style>
