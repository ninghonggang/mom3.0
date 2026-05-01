<template>
  <div class="asset-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="资产编码">
          <el-input v-model="searchForm.asset_code" placeholder="请输入资产编码" clearable />
        </el-form-item>
        <el-form-item label="资产名称">
          <el-input v-model="searchForm.asset_name" placeholder="请输入资产名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable style="width: 150px">
            <el-option label="启用" value="1" />
            <el-option label="禁用" value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('eam:asset:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="asset_code" label="资产编码" width="150" />
        <el-table-column prop="asset_name" label="资产名称" min-width="150" />
        <el-table-column prop="asset_type" label="资产类型" width="120">
          <template #default="{ row }">
            <el-tag>{{ getAssetTypeText(row.asset_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="specification" label="规格型号" min-width="150" show-overflow-tooltip />
        <el-table-column prop="workshop_name" label="使用车间" width="120" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('eam:asset:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('eam:asset:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="资产编码" prop="asset_code">
          <el-input v-model="formData.asset_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="资产名称" prop="asset_name">
          <el-input v-model="formData.asset_name" />
        </el-form-item>
        <el-form-item label="资产类型" prop="asset_type">
          <el-select v-model="formData.asset_type">
            <el-option label="生产设备" value="production" />
            <el-option label="检测设备" value="testing" />
            <el-option label="动力设备" value="power" />
            <el-option label="办公设备" value="office" />
          </el-select>
        </el-form-item>
        <el-form-item label="规格型号" prop="specification">
          <el-input v-model="formData.specification" />
        </el-form-item>
        <el-form-item label="使用车间" prop="workshop_name">
          <el-input v-model="formData.workshop_name" />
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
import { getAssetList, createAsset, updateAsset, deleteAsset } from '@/api/eam'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ asset_code: '', asset_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, asset_code: '', asset_name: '', asset_type: '', specification: '', workshop_name: '', status: 1 })

const rules: FormRules = {
  asset_code: [{ required: true, message: '请输入资产编码', trigger: 'blur' }],
  asset_name: [{ required: true, message: '请输入资产名称', trigger: 'blur' }],
  asset_type: [{ required: true, message: '请选择资产类型', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑设备资产' : '新增设备资产')

const getAssetTypeText = (type: string) => {
  const map: Record<string, string> = { production: '生产设备', testing: '检测设备', power: '动力设备', office: '办公设备' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getAssetList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { Object.assign(searchForm, { asset_code: '', asset_name: '', status: '' }); handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, asset_code: '', asset_name: '', asset_type: '', specification: '', workshop_name: '', status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除吗？', '提示', { type: 'warning' })
    await deleteAsset(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateAsset(formData.id, formData) : await createAsset(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.asset-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
