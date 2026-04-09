<template>
  <div class="inventory-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="物料编码">
          <el-input v-model="searchForm.material_code" placeholder="请输入物料编码" clearable />
        </el-form-item>
        <el-form-item label="物料名称">
          <el-input v-model="searchForm.material_name" placeholder="请输入物料名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('wms:inventory:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="warehouse_id" label="仓库ID" width="100" />
        <el-table-column prop="location_id" label="库位ID" width="100" />
        <el-table-column prop="quantity" label="库存数量" width="100" />
        <el-table-column prop="available_qty" label="可用数量" width="100" />
        <el-table-column prop="allocated_qty" label="已分配数量" width="110" />
        <el-table-column prop="locked_qty" label="冻结数量" width="100" />
        <el-table-column prop="batch_no" label="批次号" width="120" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('wms:inventory:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('wms:inventory:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="物料ID" prop="material_id">
          <el-input-number v-model="formData.material_id" :min="1" />
        </el-form-item>
        <el-form-item label="物料编码" prop="material_code">
          <el-input v-model="formData.material_code" />
        </el-form-item>
        <el-form-item label="物料名称" prop="material_name">
          <el-input v-model="formData.material_name" />
        </el-form-item>
        <el-form-item label="仓库ID" prop="warehouse_id">
          <el-input-number v-model="formData.warehouse_id" :min="1" />
        </el-form-item>
        <el-form-item label="库位ID" prop="location_id">
          <el-input-number v-model="formData.location_id" :min="1" />
        </el-form-item>
        <el-form-item label="库存数量" prop="quantity">
          <el-input-number v-model="formData.quantity" :min="0" />
        </el-form-item>
        <el-form-item label="可用数量">
          <el-input-number v-model="formData.available_qty" :min="0" />
        </el-form-item>
        <el-form-item label="已分配数量">
          <el-input-number v-model="formData.allocated_qty" :min="0" />
        </el-form-item>
        <el-form-item label="冻结数量">
          <el-input-number v-model="formData.locked_qty" :min="0" />
        </el-form-item>
        <el-form-item label="批次号">
          <el-input v-model="formData.batch_no" />
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
import { getInventoryList, adjustInventory, deleteInventory } from '@/api/wms'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ material_code: '', material_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, material_id: 0 as number | undefined, material_code: '', material_name: '',
  warehouse_id: 0 as number | undefined, location_id: 0 as number | undefined,
  quantity: 0 as number | undefined, available_qty: 0 as number | undefined,
  allocated_qty: 0 as number | undefined, locked_qty: 0 as number | undefined, batch_no: ''
})

const rules: FormRules = {
  material_id: [{ required: true, message: '请输入物料ID', trigger: 'blur' }],
  material_code: [{ required: true, message: '请输入物料编码', trigger: 'blur' }],
  warehouse_id: [{ required: true, message: '请输入仓库ID', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑库存' : '新增库存')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getInventoryList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.material_code = ''; searchForm.material_name = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, material_id: undefined, material_code: '', material_name: '',
    warehouse_id: undefined, location_id: undefined, quantity: undefined, available_qty: undefined,
    allocated_qty: undefined, locked_qty: undefined, batch_no: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该库存记录吗？', '提示', { type: 'warning' })
    await deleteInventory(row.id)
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
    await adjustInventory(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.inventory-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
