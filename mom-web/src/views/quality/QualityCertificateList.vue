<template>
  <div class="quality-certificate-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="订单编号">
          <el-input v-model="searchForm.order_code" placeholder="请输入订单编号" clearable />
        </el-form-item>
        <el-form-item label="产品编码">
          <el-input v-model="searchForm.product_code" placeholder="请输入产品编码" clearable />
        </el-form-item>
        <el-form-item label="产品名称">
          <el-input v-model="searchForm.product_name" placeholder="请输入产品名称" clearable />
        </el-form-item>
        <el-form-item label="批次号">
          <el-input v-model="searchForm.batch_no" placeholder="请输入批次号" clearable />
        </el-form-item>
        <el-form-item label="证书类型">
          <el-select v-model="searchForm.cert_type" placeholder="请选择" clearable style="width: 120px">
            <el-option label="COC" value="COC" />
            <el-option label="质检报告" value="质检报告" />
          </el-select>
        </el-form-item>
        <el-form-item label="检验结果">
          <el-select v-model="searchForm.result" placeholder="请选择" clearable style="width: 100px">
            <el-option label="合格" :value="1" />
            <el-option label="不合格" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon>新增证书
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="cert_code" label="证书编号" width="140" />
        <el-table-column prop="cert_type" label="证书类型" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.cert_type === 'COC'" type="success">COC</el-tag>
            <el-tag v-else type="warning">质检报告</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="order_code" label="订单编号" width="130" />
        <el-table-column prop="product_code" label="产品编码" width="110" />
        <el-table-column prop="product_name" label="产品名称" min-width="150" />
        <el-table-column prop="batch_no" label="批次号" width="110" />
        <el-table-column prop="quantity" label="数量" width="80" />
        <el-table-column prop="unit" label="单位" width="60" />
        <el-table-column prop="inspector" label="检验员" width="80" />
        <el-table-column prop="inspect_date" label="检验日期" width="110">
          <template #default="{ row }">
            {{ row.inspect_date ? row.inspect_date.slice(0, 10) : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="result" label="结果" width="80">
          <template #default="{ row }">
            <el-tag :type="row.result === 1 ? 'success' : 'danger'">
              {{ row.result === 1 ? '合格' : '不合格' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">
              {{ row.status === 1 ? '有效' : '无效' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新增/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form :model="formData" label-width="100px" :rules="formRules" ref="formRef">
        <el-form-item label="证书类型" prop="cert_type">
          <el-select v-model="formData.cert_type" placeholder="请选择证书类型" style="width: 100%">
            <el-option label="COC" value="COC" />
            <el-option label="质检报告" value="质检报告" />
          </el-select>
        </el-form-item>
        <el-form-item label="订单编号">
          <el-input v-model="formData.order_code" placeholder="请输入订单编号" />
        </el-form-item>
        <el-form-item label="产品编码" prop="product_code">
          <el-input v-model="formData.product_code" placeholder="请输入产品编码" />
        </el-form-item>
        <el-form-item label="产品名称" prop="product_name">
          <el-input v-model="formData.product_name" placeholder="请输入产品名称" />
        </el-form-item>
        <el-form-item label="批次号">
          <el-input v-model="formData.batch_no" placeholder="请输入批次号" />
        </el-form-item>
        <el-form-item label="数量">
          <el-input-number v-model="formData.quantity" :min="0" :precision="2" style="width: 100%" />
        </el-form-item>
        <el-form-item label="单位">
          <el-input v-model="formData.unit" placeholder="请输入单位" />
        </el-form-item>
        <el-form-item label="检验员">
          <el-input v-model="formData.inspector" placeholder="请输入检验员" />
        </el-form-item>
        <el-form-item label="检验日期">
          <el-date-picker v-model="formData.inspect_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="检验结果">
          <el-radio-group v-model="formData.result">
            <el-radio :label="1">合格</el-radio>
            <el-radio :label="0">不合格</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="发证日期">
          <el-date-picker v-model="formData.issue_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="有效期">
          <el-date-picker v-model="formData.expiry_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remarks" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getQualityCertificateList, createQualityCertificate, updateQualityCertificate, deleteQualityCertificate } from '@/api/quality_certificate'

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const tableData = ref<any[]>([])
const dialogTitle = ref('新增证书')

const searchForm = reactive({
  order_code: '',
  product_code: '',
  product_name: '',
  batch_no: '',
  cert_type: '',
  result: null as number | null
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  id: null,
  cert_type: 'COC',
  order_code: '',
  product_id: 0,
  product_code: '',
  product_name: '',
  batch_no: '',
  quantity: 1,
  unit: '',
  inspector: '',
  inspect_date: '',
  result: 1,
  issue_date: '',
  expiry_date: '',
  remarks: '',
  attachments: ''
})

const formRules = {
  cert_type: [{ required: true, message: '请选择证书类型', trigger: 'change' }],
  product_code: [{ required: true, message: '请输入产品编码', trigger: 'blur' }],
  product_name: [{ required: true, message: '请输入产品名称', trigger: 'blur' }]
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getQualityCertificateList({
      page: pagination.page,
      page_size: pagination.pageSize,
      order_code: searchForm.order_code,
      product_code: searchForm.product_code,
      product_name: searchForm.product_name,
      batch_no: searchForm.batch_no,
      cert_type: searchForm.cert_type,
      result: searchForm.result
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.order_code = ''
  searchForm.product_code = ''
  searchForm.product_name = ''
  searchForm.batch_no = ''
  searchForm.cert_type = ''
  searchForm.result = null
  handleSearch()
}

const handleCreate = () => {
  dialogTitle.value = '新增证书'
  formData.id = null
  formData.cert_type = 'COC'
  formData.order_code = ''
  formData.product_id = 0
  formData.product_code = ''
  formData.product_name = ''
  formData.batch_no = ''
  formData.quantity = 1
  formData.unit = ''
  formData.inspector = ''
  formData.inspect_date = ''
  formData.result = 1
  formData.issue_date = ''
  formData.expiry_date = ''
  formData.remarks = ''
  formData.attachments = ''
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  dialogTitle.value = '编辑证书'
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleSave = async () => {
  if (!formData.cert_type || !formData.product_code || !formData.product_name) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    if (formData.id) {
      await updateQualityCertificate(formData.id, formData)
    } else {
      await createQualityCertificate(formData)
    }
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定删除该证书吗？`, '提示', { type: 'warning' })
    await deleteQualityCertificate(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.quality-certificate-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>