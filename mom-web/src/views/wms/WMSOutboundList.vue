<template>
  <div class="wms-outbound-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="出库单号">
          <el-input v-model="searchForm.outbound_no" placeholder="请输入出库单号" clearable />
        </el-form-item>
        <el-form-item label="客户">
          <el-input v-model="searchForm.customer_name" placeholder="请输入客户" clearable />
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
      <el-button type="primary" v-if="hasPermission('wms:outbound:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="outbound_no" label="出库单号" width="150" />
        <el-table-column prop="customer_name" label="客户名称" min-width="150" />
        <el-table-column prop="outbound_type" label="出库类型" width="120">
          <template #default="{ row }">
            <el-tag>{{ getOutboundTypeText(row.outbound_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="outbound_date" label="出库日期" width="120" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('wms:outbound:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('wms:outbound:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="出库单号" prop="outbound_no">
          <el-input v-model="formData.outbound_no" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="客户名称" prop="customer_name">
          <el-input v-model="formData.customer_name" />
        </el-form-item>
        <el-form-item label="出库类型" prop="outbound_type">
          <el-select v-model="formData.outbound_type">
            <el-option label="生产领料" value="production" />
            <el-option label="销售出库" value="sales" />
            <el-option label="调拨出库" value="transfer" />
          </el-select>
        </el-form-item>
        <el-form-item label="出库日期" prop="outbound_date">
          <el-date-picker v-model="formData.outbound_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" rows="3" />
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
import { getOutboundList, createOutbound, updateOutbound, deleteOutbound } from '@/api/wms'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ outbound_no: '', customer_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, outbound_no: '', customer_name: '', outbound_type: '', outbound_date: '', status: 1, remark: '' })

const rules: FormRules = {
  outbound_no: [{ required: true, message: '请输入出库单号', trigger: 'blur' }],
  customer_name: [{ required: true, message: '请输入客户名称', trigger: 'blur' }],
  outbound_type: [{ required: true, message: '请选择出库类型', trigger: 'change' }]
}

const dialogTitle = computed(() => formData.id ? '编辑出库单' : '新增出库单')

const getOutboundTypeText = (type: string) => {
  const map: Record<string, string> = { production: '生产领料', sales: '销售出库', transfer: '调拨出库' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getOutboundList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { Object.assign(searchForm, { outbound_no: '', customer_name: '', status: '' }); handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, outbound_no: '', customer_name: '', outbound_type: '', outbound_date: '', status: 1, remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除吗？', '提示', { type: 'warning' })
    await deleteOutbound(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateOutbound(formData.id, formData) : await createOutbound(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.wms-outbound-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
