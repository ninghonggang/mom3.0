<template>
  <div class="warehouse-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="仓库编码">
          <el-input v-model="searchForm.warehouse_code" placeholder="请输入仓库编码" clearable />
        </el-form-item>
        <el-form-item label="仓库名称">
          <el-input v-model="searchForm.warehouse_name" placeholder="请输入仓库名称" clearable />
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
        <el-table-column prop="warehouse_code" label="仓库编码" width="120" />
        <el-table-column prop="warehouse_name" label="仓库名称" min-width="150" />
        <el-table-column prop="warehouse_type" label="仓库类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.warehouse_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="address" label="地址" min-width="180" />
        <el-table-column prop="manager" label="负责人" width="100" />
        <el-table-column prop="phone" label="联系电话" width="130" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
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
        <el-form-item label="仓库编码" prop="warehouse_code">
          <el-input v-model="formData.warehouse_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="仓库名称" prop="warehouse_name">
          <el-input v-model="formData.warehouse_name" />
        </el-form-item>
        <el-form-item label="仓库类型" prop="warehouse_type">
          <el-select v-model="formData.warehouse_type">
            <el-option label="原料仓" value="raw" />
            <el-option label="成品仓" value="finished" />
            <el-option label="半成品仓" value="semi" />
            <el-option label="工具仓" value="tool" />
            <el-option label="废品仓" value="waste" />
          </el-select>
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="formData.address" />
        </el-form-item>
        <el-form-item label="负责人">
          <el-input v-model="formData.manager" />
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="formData.phone" />
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
import { getWarehouseList, createWarehouse, updateWarehouse, deleteWarehouse } from '@/api/wms'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ warehouse_code: '', warehouse_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, warehouse_code: '', warehouse_name: '', warehouse_type: '', address: '', manager: '', phone: '', status: 1
})

const rules: FormRules = {
  warehouse_code: [{ required: true, message: '请输入仓库编码', trigger: 'blur' }],
  warehouse_name: [{ required: true, message: '请输入仓库名称', trigger: 'blur' }],
  warehouse_type: [{ required: true, message: '请选择仓库类型', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑仓库' : '新增仓库')

const getTypeText = (type: string) => {
  const map: Record<string, string> = { raw: '原料仓', finished: '成品仓', semi: '半成品仓', tool: '工具仓', waste: '废品仓' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getWarehouseList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.warehouse_code = ''; searchForm.warehouse_name = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, warehouse_code: '', warehouse_name: '', warehouse_type: '', address: '', manager: '', phone: '', status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该仓库吗？', '提示', { type: 'warning' })
  await deleteWarehouse(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateWarehouse(formData.id, formData) : await createWarehouse(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.warehouse-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
