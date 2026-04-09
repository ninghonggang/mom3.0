<template>
  <div class="spare-part-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="备件编码">
          <el-input v-model="searchForm.spare_part_code" placeholder="请输入备件编码" clearable />
        </el-form-item>
        <el-form-item label="备件名称">
          <el-input v-model="searchForm.spare_part_name" placeholder="请输入备件名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('equipment:spare:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="spare_part_code" label="备件编码" width="120" />
        <el-table-column prop="spare_part_name" label="备件名称" min-width="150" />
        <el-table-column prop="spare_part_type" label="备件类型" width="120" />
        <el-table-column prop="specification" label="规格型号" width="120" />
        <el-table-column prop="unit" label="单位" width="80" />
        <el-table-column prop="stock_quantity" label="库存数量" width="100" />
        <el-table-column prop="min_stock" label="最小库存" width="100" />
        <el-table-column prop="max_stock" label="最大库存" width="100" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('equipment:spare:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('equipment:spare:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="备件编码" prop="spare_part_code">
          <el-input v-model="formData.spare_part_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="备件名称" prop="spare_part_name">
          <el-input v-model="formData.spare_part_name" />
        </el-form-item>
        <el-form-item label="备件类型">
          <el-input v-model="formData.spare_part_type" />
        </el-form-item>
        <el-form-item label="规格型号">
          <el-input v-model="formData.specification" />
        </el-form-item>
        <el-form-item label="单位">
          <el-input v-model="formData.unit" />
        </el-form-item>
        <el-form-item label="库存数量">
          <el-input-number v-model="formData.stock_quantity" :min="0" />
        </el-form-item>
        <el-form-item label="最小库存">
          <el-input-number v-model="formData.min_stock" :min="0" />
        </el-form-item>
        <el-form-item label="最大库存">
          <el-input-number v-model="formData.max_stock" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" />
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
import { getSparePartList, createSparePart, updateSparePart, deleteSparePart } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ spare_part_code: '', spare_part_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive({
  id: 0, spare_part_code: '', spare_part_name: '', spare_part_type: '', specification: '', unit: '', stock_quantity: 0, min_stock: 0, max_stock: 0, status: 1, remark: ''
})

const rules: FormRules = {
  spare_part_code: [{ required: true, message: '请输入备件编码', trigger: 'blur' }],
  spare_part_name: [{ required: true, message: '请输入备件名称', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑备件' : '新增备件')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getSparePartList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.spare_part_code = ''; searchForm.spare_part_name = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, spare_part_code: '', spare_part_name: '', spare_part_type: '', specification: '', unit: '', stock_quantity: 0, min_stock: 0, max_stock: 0, status: 1, remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该备件吗？', '提示', { type: 'warning' })
    await deleteSparePart(row.id)
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
    formData.id ? await updateSparePart(formData.id, formData) : await createSparePart(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.spare-part-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
