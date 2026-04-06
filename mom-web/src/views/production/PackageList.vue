<template>
  <div class="package-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="箱条码">
          <el-input v-model="searchForm.package_no" placeholder="请输入箱条码" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="开箱" value="OPEN" />
            <el-option label="已封箱" value="SEALED" />
            <el-option label="已发货" value="SHIPPED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:package:create')" @click="handleCreate">
        <el-icon><Plus /></el-icon>新建装箱
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="package_no" label="箱条码" min-width="180" />
        <el-table-column prop="package_type" label="箱型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getTypeText(row.package_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="product_code" label="产品编码" width="120" />
        <el-table-column prop="qty" label="数量" width="80" />
        <el-table-column prop="serial_nos" label="序列号" min-width="150" show-overflow-tooltip>
          <template #default="{ row }">
            {{ (row.serial_nos || []).join(', ') }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'OPEN' ? 'info' : row.status === 'SEALED' ? 'success' : 'warning'">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="seal_by" label="封箱人" width="100" />
        <el-table-column prop="seal_time" label="封箱时间" width="160" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('production:package:add-item') && row.status === 'OPEN'" @click="handleAddItem(row)">添加序列号</el-button>
            <el-button link type="success" v-if="hasPermission('production:package:seal') && row.status === 'OPEN'" @click="handleSeal(row)">封箱</el-button>
            <el-button link type="danger" v-if="hasPermission('production:package:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新建装箱对话框 -->
    <el-dialog v-model="createDialogVisible" title="新建装箱" width="500px">
      <el-form ref="createFormRef" :model="createForm" :rules="createRules" label-width="100px">
        <el-form-item label="产品ID" prop="product_id">
          <el-input-number v-model="createForm.product_id" :min="1" />
        </el-form-item>
        <el-form-item label="产品编码" prop="product_code">
          <el-input v-model="createForm.product_code" />
        </el-form-item>
        <el-form-item label="箱型" prop="package_type">
          <el-select v-model="createForm.package_type">
            <el-option label="小箱" value="SMALL_BOX" />
            <el-option label="大箱" value="BIG_BOX" />
            <el-option label="托盘" value="PALLET" />
          </el-select>
        </el-form-item>
        <el-form-item label="工单ID">
          <el-input-number v-model="createForm.production_order_id" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleDoCreate">确定</el-button>
      </template>
    </el-dialog>

    <!-- 添加序列号对话框 -->
    <el-dialog v-model="addItemDialogVisible" title="添加序列号" width="400px">
      <el-form :model="addItemForm" label-width="80px">
        <el-form-item label="箱条码">
          <el-input v-model="addItemForm.package_no" disabled />
        </el-form-item>
        <el-form-item label="序列号">
          <el-input v-model="addItemForm.serial_no" placeholder="请扫描或输入序列号" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="addItemDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleDoAddItem">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getPackageList, createPackage, addPackageItem, sealPackage, deletePackage } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const createDialogVisible = ref(false)
const addItemDialogVisible = ref(false)
const submitLoading = ref(false)
const createFormRef = ref<FormInstance>()

const searchForm = reactive({ package_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const createForm = reactive({ product_id: 0, product_code: '', package_type: 'SMALL_BOX', production_order_id: 0 })
const addItemForm = reactive({ package_no: '', serial_no: '' })

const createRules: FormRules = {
  product_id: [{ required: true, message: '请输入产品ID', trigger: 'blur' }],
  package_type: [{ required: true, message: '请选择箱型', trigger: 'change' }],
}

const getTypeText = (val: string) => ({ SMALL_BOX: '小箱', BIG_BOX: '大箱', PALLET: '托盘' }[val] || val)
const getStatusText = (val: string) => ({ OPEN: '开箱', SEALED: '已封箱', SHIPPED: '已发货' }[val] || val)

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => params[k] === '' && delete params[k])
    const res = await getPackageList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.package_no = ''; searchForm.status = ''; handleSearch() }

const handleCreate = () => {
  Object.assign(createForm, { product_id: 0, product_code: '', package_type: 'SMALL_BOX', production_order_id: 0 })
  createDialogVisible.value = true
}

const handleDoCreate = async () => {
  if (!createFormRef.value) return
  await createFormRef.value.validate()
  submitLoading.value = true
  try {
    await createPackage(createForm)
    ElMessage.success('创建成功')
    createDialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

const handleAddItem = (row: any) => {
  Object.assign(addItemForm, { package_no: row.package_no, serial_no: '' })
  addItemDialogVisible.value = true
}

const handleDoAddItem = async () => {
  if (!addItemForm.serial_no) {
    ElMessage.warning('请输入序列号')
    return
  }
  submitLoading.value = true
  try {
    await addPackageItem(addItemForm)
    ElMessage.success('添加成功')
    addItemDialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

const handleSeal = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定封箱吗？', '提示', { type: 'warning' })
    await sealPackage({ package_no: row.package_no, seal_by: '' })
    ElMessage.success('封箱成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该装箱记录吗？', '提示', { type: 'warning' })
    await deletePackage(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.package-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
