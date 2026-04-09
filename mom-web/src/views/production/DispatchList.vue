<template>
  <div class="dispatch-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:dispatch:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增派工
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="order_no" label="工单号" width="150" />
        <el-table-column prop="process_name" label="工序" width="120" />
        <el-table-column prop="station_name" label="工位" width="100" />
        <el-table-column prop="assign_user_name" label="派工人" width="100" />
        <el-table-column prop="quantity" label="派工数量" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="statusType[row.status]">{{ statusText[row.status] }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="success" v-if="hasPermission('production:dispatch:edit') && row.status === 1" @click="handleStart(row)">开始</el-button>
            <el-button link type="warning" v-if="hasPermission('production:dispatch:edit') && row.status === 2" @click="handleComplete(row)">完成</el-button>
            <el-button link type="primary" v-if="hasPermission('production:dispatch:edit')" @click="handleEdit(row)">编辑</el-button>
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

    <el-dialog v-model="dialogVisible" title="派工" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="工单号" prop="order_no">
          <el-input v-model="formData.order_no" />
        </el-form-item>
        <el-form-item label="工序">
          <el-input v-model="formData.process_name" />
        </el-form-item>
        <el-form-item label="工位">
          <el-input v-model="formData.station_name" />
        </el-form-item>
        <el-form-item label="派工数量" prop="quantity">
          <el-input-number v-model="formData.quantity" :min="0" />
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
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, FormInstance, FormRules } from 'element-plus'
import { getDispatchList, createDispatch, updateDispatch, startDispatch, completeDispatch } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ order_no: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, order_no: '', process_name: '', station_name: '', quantity: 0
})

const statusText: Record<number, string> = { 1: '待开始', 2: '进行中', 3: '已完成' }
const statusType: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }

const rules: FormRules = {
  order_no: [{ required: true, message: '请输入工单号', trigger: 'blur' }],
  quantity: [{ required: true, message: '请输入派工数量', trigger: 'blur' }]
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDispatchList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.order_no = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, order_no: '', process_name: '', station_name: '', quantity: 0 })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleStart = async (row: any) => {
  await startDispatch(row.id)
  ElMessage.success('已开始')
  loadData()
}

const handleComplete = async (row: any) => {
  await completeDispatch(row.id)
  ElMessage.success('已完成')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateDispatch(formData.id, formData) : await createDispatch(formData)
    ElMessage.success(formData.id ? '更新成功' : '派工成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.dispatch-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
