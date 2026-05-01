<template>
  <div class="product-unit-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="单位编码">
          <el-input v-model="searchForm.unit_code" placeholder="请输入单位编码" clearable />
        </el-form-item>
        <el-form-item label="单位名称">
          <el-input v-model="searchForm.unit_name" placeholder="请输入单位名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:product-unit:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="unit_code" label="单位编码" width="120" />
        <el-table-column prop="unit_name" label="单位名称" min-width="150" />
        <el-table-column prop="symbol" label="单位符号" width="100">
          <template #default="{ row }">
            <el-tag type="info">{{ row.symbol }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="conversion_rate" label="换算率" width="100">
          <template #default="{ row }">
            {{ row.conversion_rate }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('mdm:product-unit:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:product-unit:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="单位编码" prop="unit_code">
          <el-input v-model="formData.unit_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="单位名称" prop="unit_name">
          <el-input v-model="formData.unit_name" />
        </el-form-item>
        <el-form-item label="单位符号" prop="symbol">
          <el-input v-model="formData.symbol" placeholder="如：个/箱/米" />
        </el-form-item>
        <el-form-item label="换算率" prop="conversion_rate">
          <el-input-number v-model="formData.conversion_rate" :min="0" :precision="2" style="width: 100%" />
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
import { getProductUnitList, createProductUnit, updateProductUnit, deleteProductUnit } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ unit_code: '', unit_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, unit_code: '', unit_name: '', symbol: '', conversion_rate: 1, status: 1 })

const rules: FormRules = {
  unit_code: [{ required: true, message: '请输入单位编码', trigger: 'blur' }],
  unit_name: [{ required: true, message: '请输入单位名称', trigger: 'blur' }],
  symbol: [{ required: true, message: '请输入单位符号', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑计量单位' : '新增计量单位')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getProductUnitList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { Object.assign(searchForm, { unit_code: '', unit_name: '' }); handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, unit_code: '', unit_name: '', symbol: '', conversion_rate: 1, status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除吗？', '提示', { type: 'warning' })
    await deleteProductUnit(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateProductUnit(formData.id, formData) : await createProductUnit(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.product-unit-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
