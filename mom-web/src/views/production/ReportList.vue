<template>
  <div class="report-list">
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
      <el-button type="primary" v-if="hasPermission('production:report:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增报工
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="order_no" label="工单号" width="150" />
        <el-table-column prop="process_name" label="工序" width="120" />
        <el-table-column prop="station_name" label="工位" width="100" />
        <el-table-column prop="report_user_name" label="报工人" width="100" />
        <el-table-column prop="report_date" label="报工日期" width="120" />
        <el-table-column prop="quantity" label="报工数量" width="100" />
        <el-table-column prop="qualified_qty" label="合格数量" width="100" />
        <el-table-column prop="rejected_qty" label="不良数量" width="100" />
        <el-table-column prop="work_time" label="工时(分钟)" width="100" />
        <el-table-column prop="remark" label="备注" min-width="120" show-overflow-tooltip />
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

    <el-dialog v-model="dialogVisible" title="生产报工" width="600px">
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
        <el-form-item label="报工数量" prop="quantity">
          <el-input-number v-model="formData.quantity" :min="0" />
        </el-form-item>
        <el-form-item label="合格数量" prop="qualified_qty">
          <el-input-number v-model="formData.qualified_qty" :min="0" />
        </el-form-item>
        <el-form-item label="不良数量">
          <el-input-number v-model="formData.rejected_qty" :min="0" />
        </el-form-item>
        <el-form-item label="工时(分钟)">
          <el-input-number v-model="formData.work_time" :min="0" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" rows="2" />
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
import { getProductionReportList, createProductionReport } from '@/api/production'
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
  order_no: '', process_name: '', station_name: '', quantity: 0,
  qualified_qty: 0, rejected_qty: 0, work_time: 0, remark: ''
})

const rules: FormRules = {
  order_no: [{ required: true, message: '请输入工单号', trigger: 'blur' }],
  quantity: [{ required: true, message: '请输入报工数量', trigger: 'blur' }],
  qualified_qty: [{ required: true, message: '请输入合格数量', trigger: 'blur' }]
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getProductionReportList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.order_no = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { order_no: '', process_name: '', station_name: '', quantity: 0, qualified_qty: 0, rejected_qty: 0, work_time: 0, remark: '' })
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    await createProductionReport(formData)
    ElMessage.success('报工成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.report-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
