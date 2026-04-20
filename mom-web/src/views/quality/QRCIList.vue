<template>
  <div class="qrci-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="QRCI编号"><el-input v-model="searchForm.qrci_no" placeholder="请输入" clearable /></el-form-item>
        <el-form-item label="严重程度">
          <el-select v-model="searchForm.severity_level" placeholder="请选择" clearable>
            <el-option label="CRITICAL" value="CRITICAL" />
            <el-option label="HIGH" value="HIGH" />
            <el-option label="MEDIUM" value="MEDIUM" />
            <el-option label="LOW" value="LOW" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="开放" value="OPEN" />
            <el-option label="进行中" value="IN_PROGRESS" />
            <el-option label="验证中" value="VERIFICATION" />
            <el-option label="已关闭" value="CLOSED" />
          </el-select>
        </el-form-item>
        <el-form-item><el-button type="primary" @click="handleSearch">查询</el-button><el-button @click="handleReset">重置</el-button></el-form-item>
      </el-form>
    </el-card>
    <el-card class="toolbar-card"><el-button type="primary" v-if="hasPermission('quality:qrci:add')" @click="handleAdd"><el-icon><Plus /></el-icon>新增</el-button></el-card>
    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="qrci_no" label="QRCI编号" width="140" />
        <el-table-column prop="severity_level" label="严重程度" width="100">
          <template #default="{ row }"><el-tag :type="getSeverityType(row.severity_level)">{{ row.severity_level }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="defect_description" label="问题描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="responsible_dept_name" label="责任部门" width="120" />
        <el-table-column prop="owner_name" label="责任人" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }"><el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="target_close_date" label="目标关闭日期" width="130" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('quality:qrci:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="success" size="small" v-if="row.status !== 'CLOSED' && hasPermission('quality:qrci:close')" @click="handleClose(row)">关闭</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination"><el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" /></div>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-form-item label="问题描述" prop="defect_description"><el-input v-model="formData.defect_description" type="textarea" /></el-form-item>
        <el-form-item label="严重程度" prop="severity_level">
          <el-select v-model="formData.severity_level" placeholder="请选择">
            <el-option label="CRITICAL" value="CRITICAL" />
            <el-option label="HIGH" value="HIGH" />
            <el-option label="MEDIUM" value="MEDIUM" />
            <el-option label="LOW" value="LOW" />
          </el-select>
        </el-form-item>
        <el-form-item label="发现日期" prop="discovery_date"><el-date-picker v-model="formData.discovery_date" type="date" value-format="YYYY-MM-DD" /></el-form-item>
        <el-form-item label="发现地点"><el-input v-model="formData.discovery_location" /></el-form-item>
        <el-form-item label="责任部门"><el-input v-model="formData.responsible_dept_name" /></el-form-item>
        <el-form-item label="责任人"><el-input v-model="formData.owner_name" /></el-form-item>
        <el-form-item label="目标关闭日期"><el-date-picker v-model="formData.target_close_date" type="date" value-format="YYYY-MM-DD" /></el-form-item>
        <el-form-item label="状态">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option label="开放" value="OPEN" />
            <el-option label="进行中" value="IN_PROGRESS" />
            <el-option label="验证中" value="VERIFICATION" />
            <el-option label="已关闭" value="CLOSED" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer><el-button @click="dialogVisible = false">取消</el-button><el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button></template>
    </el-dialog>
    <el-dialog v-model="detailVisible" title="QRCI详情" width="900px">
      <el-descriptions :column="2" border v-if="currentRow">
        <el-descriptions-item label="QRCI编号">{{ currentRow.qrci_no }}</el-descriptions-item>
        <el-descriptions-item label="严重程度"><el-tag :type="getSeverityType(currentRow.severity_level)">{{ currentRow.severity_level }}</el-tag></el-descriptions-item>
        <el-descriptions-item label="问题描述" :span="2">{{ currentRow.defect_description }}</el-descriptions-item>
        <el-descriptions-item label="发现日期">{{ currentRow.discovery_date }}</el-descriptions-item>
        <el-descriptions-item label="发现地点">{{ currentRow.discovery_location }}</el-descriptions-item>
        <el-descriptions-item label="责任部门">{{ currentRow.responsible_dept_name }}</el-descriptions-item>
        <el-descriptions-item label="责任人">{{ currentRow.owner_name }}</el-descriptions-item>
        <el-descriptions-item label="目标关闭日期">{{ currentRow.target_close_date }}</el-descriptions-item>
        <el-descriptions-item label="实际关闭日期">{{ currentRow.actual_close_date }}</el-descriptions-item>
        <el-descriptions-item label="状态"><el-tag :type="getStatusType(currentRow.status)">{{ getStatusText(currentRow.status) }}</el-tag></el-descriptions-item>
        <el-descriptions-item label="验证结果">{{ currentRow.verification_result }}</el-descriptions-item>
      </el-descriptions>
      <template #footer><el-button @click="detailVisible = false">关闭</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getQRCIList, createQRCI, updateQRCI, closeQRCI } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()
const loading = ref(false), tableData = ref<any[]>([]), dialogVisible = ref(false), detailVisible = ref(false), submitLoading = ref(false), formRef = ref<FormInstance>(), currentRow = ref<any>(null)
const searchForm = reactive({ qrci_no: '', severity_level: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive<any>({ id: 0, defect_description: '', severity_level: '', discovery_date: '', discovery_location: '', responsible_dept_name: '', owner_name: '', target_close_date: '', status: 'OPEN' })
const rules: FormRules = { defect_description: [{ required: true, message: '请输入问题描述', trigger: 'blur' }], severity_level: [{ required: true, message: '请选择严重程度', trigger: 'change' }], discovery_date: [{ required: true, message: '请选择发现日期', trigger: 'change' }] }
const dialogTitle = computed(() => formData.id ? '编辑QRCI' : '新增QRCI')
const getSeverityType = (s: string) => ({ CRITICAL: 'danger', HIGH: 'warning', MEDIUM: 'info', LOW: 'success' }[s] || 'info')
const getStatusType = (s: string) => ({ OPEN: 'info', IN_PROGRESS: 'warning', VERIFICATION: 'warning', CLOSED: 'success', CANCELLED: 'info' }[s] || 'info')
const getStatusText = (s: string) => ({ OPEN: '开放', IN_PROGRESS: '进行中', VERIFICATION: '验证中', CLOSED: '已关闭', CANCELLED: '已取消' }[s] || s)
const loadData = async () => { loading.value = true; try { const res = await getQRCIList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize }); tableData.value = res.data.list || []; pagination.total = res.data.total || 0 } finally { loading.value = false } }
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.qrci_no = ''; searchForm.severity_level = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, defect_description: '', severity_level: '', discovery_date: '', discovery_location: '', responsible_dept_name: '', owner_name: '', target_close_date: '', status: 'OPEN' }); dialogVisible.value = true }
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleView = (row: any) => { currentRow.value = row; detailVisible.value = true }
const handleClose = async (row: any) => { try { await ElMessageBox.confirm('确定关闭该QRCI吗？', '提示', { type: 'warning' }); await closeQRCI(row.id); ElMessage.success('关闭成功'); loadData() } catch {} }
const handleSubmit = async () => { if (!formRef.value) return; await formRef.value.validate(); submitLoading.value = true; try { formData.id ? await updateQRCI(formData.id, formData) : await createQRCI(formData); ElMessage.success(formData.id ? '更新成功' : '创建成功'); dialogVisible.value = false; loadData() } finally { submitLoading.value = false } }
onMounted(() => { loadData() })
</script>
<style scoped lang="scss">.qrci-list { .search-card, .toolbar-card { margin-bottom: 16px; } .toolbar-card :deep(.el-card__body) { padding: 12px 16px; } .pagination { margin-top: 16px; display: flex; justify-content: flex-end; } }</style>
