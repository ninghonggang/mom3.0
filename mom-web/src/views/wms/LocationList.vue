<template>
  <div class="location-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="库位编码">
          <el-input v-model="searchForm.location_code" placeholder="请输入库位编码" clearable />
        </el-form-item>
        <el-form-item label="库位名称">
          <el-input v-model="searchForm.location_name" placeholder="请输入库位名称" clearable />
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
        <el-table-column prop="location_code" label="库位编码" width="120" />
        <el-table-column prop="location_name" label="库位名称" min-width="150" />
        <el-table-column prop="warehouse_id" label="仓库ID" width="100" />
        <el-table-column prop="zone_code" label="区域" width="80" />
        <el-table-column prop="row" label="排" width="60" />
        <el-table-column prop="col" label="列" width="60" />
        <el-table-column prop="layer" label="层" width="60" />
        <el-table-column prop="location_type" label="库位类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.location_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="capacity" label="容量" width="80" />
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
        <el-form-item label="库位编码" prop="location_code">
          <el-input v-model="formData.location_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="库位名称" prop="location_name">
          <el-input v-model="formData.location_name" />
        </el-form-item>
        <el-form-item label="所属仓库" prop="warehouse_id">
          <el-input-number v-model="formData.warehouse_id" :min="1" />
        </el-form-item>
        <el-form-item label="区域">
          <el-input v-model="formData.zone_code" />
        </el-form-item>
        <el-form-item label="排/列/层">
          <el-input-number v-model="formData.row" :min="0" placeholder="排" style="width: 100px" />
          <el-input-number v-model="formData.col" :min="0" placeholder="列" style="width: 100px; margin-left: 8px" />
          <el-input-number v-model="formData.layer" :min="0" placeholder="层" style="width: 100px; margin-left: 8px" />
        </el-form-item>
        <el-form-item label="库位类型" prop="location_type">
          <el-select v-model="formData.location_type">
            <el-option label="存储" value="storage" />
            <el-option label="备货" value="stock" />
            <el-option label="发货" value="shipping" />
          </el-select>
        </el-form-item>
        <el-form-item label="容量">
          <el-input-number v-model="formData.capacity" :min="0" />
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
import { getLocationList, createLocation, updateLocation, deleteLocation } from '@/api/wms'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ location_code: '', location_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, location_code: '', location_name: '', warehouse_id: 0 as number | undefined,
  zone_code: '', row: 0 as number | undefined, col: 0 as number | undefined,
  layer: 0 as number | undefined, location_type: '', capacity: 0 as number | undefined, status: 1
})

const rules: FormRules = {
  location_code: [{ required: true, message: '请输入库位编码', trigger: 'blur' }],
  location_name: [{ required: true, message: '请输入库位名称', trigger: 'blur' }],
  warehouse_id: [{ required: true, message: '请输入仓库ID', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑库位' : '新增库位')

const getTypeText = (type: string) => {
  const map: Record<string, string> = { storage: '存储', stock: '备货', shipping: '发货' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getLocationList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.location_code = ''; searchForm.location_name = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, location_code: '', location_name: '', warehouse_id: undefined,
    zone_code: '', row: undefined, col: undefined, layer: undefined, location_type: '', capacity: undefined, status: 1 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定删除该库位吗？', '提示', { type: 'warning' })
  await deleteLocation(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateLocation(formData.id, formData) : await createLocation(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.location-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
