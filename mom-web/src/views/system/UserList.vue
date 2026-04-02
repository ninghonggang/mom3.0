<template>
  <div class="user-list">
    <!-- 搜索区域 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="用户名">
          <el-input v-model="searchForm.username" placeholder="请输入用户名" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="正常" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 工具栏 -->
    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="danger" :disabled="!selectedRows.length" @click="handleBatchDelete">
        <el-icon><Delete /></el-icon>批量删除
      </el-button>
    </el-card>

    <!-- 表格 -->
    <el-card>
      <el-table
        v-loading="loading"
        :data="tableData"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="username" label="用户名" min-width="120" />
        <el-table-column prop="nickname" label="昵称" min-width="120" />
        <el-table-column prop="email" label="邮箱" min-width="180" />
        <el-table-column prop="phone" label="手机号" width="130" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="240" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="primary" @click="handleAssignRoles(row)">分配角色</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="loadData"
          @current-change="loadData"
        />
      </div>
    </el-card>

    <!-- 编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="formData.username" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="昵称" prop="nickname">
          <el-input v-model="formData.nickname" />
        </el-form-item>
        <el-form-item v-if="!formData.id" label="密码" prop="password">
          <el-input v-model="formData.password" type="password" show-password />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="formData.email" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="formData.phone" />
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

    <!-- 分配角色弹窗 -->
    <el-dialog v-model="roleDialogVisible" title="分配角色" width="500px">
      <el-form label-width="80px">
        <el-form-item label="用户名">
          <span>{{ currentUser?.username }}</span>
        </el-form-item>
        <el-form-item label="角色">
          <el-checkbox-group v-model="selectedRoleIds">
            <el-checkbox v-for="role in allRoles" :key="role.id" :value="role.id">
              {{ role.role_name }}
            </el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="roleDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="roleLoading" @click="handleSaveRoles">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getUserList, createUser, updateUser, deleteUser, getAllRoles, assignUserRoles, getUserById } from '@/api/system'

const loading = ref(false)
const tableData = ref<any[]>([])
const selectedRows = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

// 角色分配相关
const roleDialogVisible = ref(false)
const roleLoading = ref(false)
const currentUser = ref<any>(null)
const allRoles = ref<any[]>([])
const selectedRoleIds = ref<number[]>([])

const searchForm = reactive({
  username: '',
  status: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

const formData = reactive({
  id: 0,
  username: '',
  nickname: '',
  password: '',
  email: '',
  phone: '',
  status: 1
})

const rules: FormRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  email: [{ type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑用户' : '新增用户')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getUserList({
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.username = ''
  searchForm.status = ''
  handleSearch()
}

const handleSelectionChange = (rows: any[]) => {
  selectedRows.value = rows
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0,
    username: '',
    nickname: '',
    password: '',
    email: '',
    phone: '',
    status: 1
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定要删除该用户吗？', '提示', { type: 'warning' })
  await deleteUser(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleBatchDelete = async () => {
  const ids = selectedRows.value.map(r => r.id)
  await ElMessageBox.confirm(`确定要删除选中的 ${ids.length} 个用户吗？`, '提示', { type: 'warning' })
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()

  submitLoading.value = true
  try {
    if (formData.id) {
      await updateUser(formData.id, formData)
      ElMessage.success('更新成功')
    } else {
      await createUser(formData)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } finally {
    submitLoading.value = false
  }
}

// 分配角色
const handleAssignRoles = async (row: any) => {
  currentUser.value = row
  roleLoading.value = true
  roleDialogVisible.value = true
  try {
    // 获取所有角色
    const rolesRes = await getAllRoles()
    allRoles.value = rolesRes.data.list || []

    // 获取用户已有角色
    const userRes = await getUserById(row.id)
    selectedRoleIds.value = userRes.data.role_ids || []
  } finally {
    roleLoading.value = false
  }
}

// 保存角色分配
const handleSaveRoles = async () => {
  if (!currentUser.value) return
  roleLoading.value = true
  try {
    await assignUserRoles(currentUser.value.id, selectedRoleIds.value)
    ElMessage.success('角色分配成功')
    roleDialogVisible.value = false
  } finally {
    roleLoading.value = false
  }
}

onMounted(() => {
  loadData()
})
</script>

<script lang="ts">
export default { name: 'UserList' }
</script>

<style scoped lang="scss">
.user-list {
  .search-card, .toolbar-card {
    margin-bottom: 16px;
  }

  .toolbar-card :deep(.el-card__body) {
    padding: 12px 16px;
  }

  .pagination {
    margin-top: 16px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>
