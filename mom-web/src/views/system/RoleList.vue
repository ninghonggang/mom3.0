<template>
  <div class="role-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="角色名称">
          <el-input v-model="searchForm.role_name" placeholder="请输入角色名称" clearable />
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
        <el-table-column prop="role_name" label="角色名称" min-width="120" />
        <el-table-column prop="role_key" label="角色标识" min-width="120" />
        <el-table-column prop="role_sort" label="排序" width="80" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="primary" @click="handlePerms(row)">权限</el-button>
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
        <el-form-item label="角色名称" prop="role_name">
          <el-input v-model="formData.role_name" />
        </el-form-item>
        <el-form-item label="角色标识" prop="role_key">
          <el-input v-model="formData.role_key" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="排序" prop="role_sort">
          <el-input-number v-model="formData.role_sort" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="1">正常</el-radio>
            <el-radio :value="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 权限配置对话框 -->
    <el-dialog v-model="permsDialogVisible" :title="'权限配置 - ' + currentRoleName" width="900px">
      <el-tabs>
        <el-tab-pane label="菜单权限">
          <el-scrollbar v-loading="permsLoading" height="400px">
            <el-tree
              ref="menuTreeRef"
              :data="menuTree"
              :props="{ children: 'children', label: 'menu_name' }"
              node-key="id"
              show-checkbox
              :default-expand-all="true"
              check-strictly
              @check="handleTreeCheck"
            />
          </el-scrollbar>
        </el-tab-pane>
        <el-tab-pane label="按钮权限">
          <el-scrollbar v-loading="permsLoading" height="400px">
            <div v-if="selectedMenusWithPerms.length === 0" class="no-perms-tip">
              请先在"菜单权限"中选择一个菜单，这里将显示该菜单的按钮权限
            </div>
            <div v-else class="perms-container">
              <div v-for="menu in selectedMenusWithPerms" :key="menu.id" class="menu-perms">
                <div class="menu-name">{{ menu.menu_name }}</div>
                <el-checkbox-group v-model="selectedPerms">
                  <el-checkbox
                    v-for="perm in menu.perms.split(',')"
                    :key="perm"
                    :value="perm.trim()"
                    :label="perm.trim()"
                  >
                    {{ getPermLabel(perm.trim()) }}
                  </el-checkbox>
                </el-checkbox-group>
              </div>
            </div>
          </el-scrollbar>
        </el-tab-pane>
      </el-tabs>
      <template #footer>
        <el-button @click="permsDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSavePerms">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getRoleList, createRole, updateRole, deleteRole, getRoleMenus, assignRoleMenus, getMenuTree, getRolePerms, assignRolePerms } from '@/api/system'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

// 权限对话框
const permsDialogVisible = ref(false)
const permsLoading = ref(false)
const menuTree = ref<any[]>([])
const checkedMenuKeys = ref<string[]>([])
const currentRoleId = ref<number>(0)
const currentRoleName = ref('')
const menuTreeRef = ref()
const selectedPerms = ref<string[]>([])
const selectedMenuIds = ref<number[]>([])

const searchForm = reactive({ role_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, role_name: '', role_key: '', role_sort: 0, status: 1, remark: '' })

const rules: FormRules = {
  role_name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
  role_key: [{ required: true, message: '请输入角色标识', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑角色' : '新增角色')

// 菜单权限名称映射
const permLabelMap: Record<string, string> = {
  'add': '新增',
  'create': '新增',
  'edit': '编辑',
  'update': '编辑',
  'delete': '删除',
  'del': '删除',
  'query': '查询',
  'list': '查询',
  'get': '查看',
  'view': '查看',
  'export': '导出',
  'import': '导入',
  'submit': '提交',
  'approve': '审批',
  'reject': '拒绝',
  'start': '启动',
  'stop': '停止',
  'complete': '完成',
  'cancel': '取消',
  'reset': '重置',
  'assign': '分配',
  'enable': '启用',
  'disable': '禁用'
}

const getPermLabel = (perm: string) => {
  // 如果是中文直接返回
  if (/[\u4e00-\u9fa5]/.test(perm)) return perm
  // 尝试从映射获取
  const key = perm.split(':')[1] || perm
  return permLabelMap[key.toLowerCase()] || perm
}

// 获取选中菜单的权限信息
const selectedMenusWithPerms = computed(() => {
  const result: any[] = []
  const allMenus = flattenMenuTree(menuTree.value)

  for (const menuId of selectedMenuIds.value) {
    const menu = allMenus.find(m => m.id === menuId)
    if (menu && menu.perms && menu.perms.trim()) {
      result.push(menu)
    }
  }
  return result
})

// 将菜单树扁平化
const flattenMenuTree = (menus: any[]): any[] => {
  const result: any[] = []
  for (const menu of menus) {
    result.push(menu)
    if (menu.children && menu.children.length > 0) {
      result.push(...flattenMenuTree(menu.children))
    }
  }
  return result
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getRoleList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.role_name = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, role_name: '', role_key: '', role_sort: 0, status: 1, remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该角色吗？', '提示', { type: 'warning' })
  await deleteRole(row.id)
  ElMessage.success('删除成功')
  loadData()
}

// 权限配置
const handlePerms = async (row: any) => {
  currentRoleId.value = row.id
  currentRoleName.value = row.role_name
  permsLoading.value = true
  permsDialogVisible.value = true
  selectedPerms.value = []
  selectedMenuIds.value = []
  try {
    // 获取所有菜单树
    const treeRes = await getMenuTree()
    menuTree.value = treeRes.data || []

    // 获取角色已有的菜单
    const roleMenusRes = await getRoleMenus(row.id)
    const roleMenus = roleMenusRes.data || []
    const keys = roleMenus.map((m: any) => String(m.id))
    selectedMenuIds.value = roleMenus.map((m: any) => m.id)

    // 获取角色已有的权限
    const permsRes = await getRolePerms(row.id)
    selectedPerms.value = permsRes.data || []

    // 设置选中的菜单
    setTimeout(() => {
      if (menuTreeRef.value) {
        menuTreeRef.value.setCheckedKeys(keys)
      }
    }, 100)
  } finally {
    permsLoading.value = false
  }
}

// 树节点选中变化
const handleTreeCheck = (node: any, checked: any) => {
  checkedMenuKeys.value = checked.checkedKeys || []
  selectedMenuIds.value = (checked.checkedKeys || []).map((k: any) => parseInt(k)).filter((k: any) => !isNaN(k))
}

// 保存权限
const handleSavePerms = async () => {
  // 获取选中的菜单ID
  const menuIds = menuTreeRef.value?.getCheckedKeys() || []
  const ids = menuIds.map((id: any) => parseInt(id)).filter((id: any) => !isNaN(id))

  // 保存菜单权限
  await assignRoleMenus(currentRoleId.value, ids)

  // 保存按钮权限
  await assignRolePerms(currentRoleId.value, selectedPerms.value)

  ElMessage.success('权限分配成功')
  permsDialogVisible.value = false
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateRole(formData.id, formData) : await createRole(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.role-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}

.no-perms-tip {
  text-align: center;
  color: #999;
  padding: 40px;
}

.perms-container {
  padding: 0 16px;
}

.menu-perms {
  margin-bottom: 20px;
  border-bottom: 1px solid #eee;
  padding-bottom: 16px;

  .menu-name {
    font-weight: bold;
    margin-bottom: 12px;
    color: var(--el-color-primary);
  }

  .el-checkbox {
    margin-right: 16px;
    margin-bottom: 8px;
  }
}
</style>
