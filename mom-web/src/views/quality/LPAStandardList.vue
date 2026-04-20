<template>
  <div class="lpa-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="标准编号"><el-input v-model="searchForm.standard_code" placeholder="请输入" clearable /></el-form-item>
        <el-form-item label="审核频率">
          <el-select v-model="searchForm.audit_frequency" placeholder="请选择" clearable>
            <el-option label="每日" value="DAILY" />
            <el-option label="每周" value="WEEKLY" />
            <el-option label="每月" value="MONTHLY" />
            <el-option label="每季度" value="QUARTERLY" />
          </el-select>
        </el-form-item>
        <el-form-item><el-button type="primary" @click="handleSearch">查询</el-button><el-button @click="handleReset">重置</el-button></el-form-item>
      </el-form>
    </el-card>
    <el-card class="toolbar-card"><el-button type="primary" v-if="hasPermission('quality:lpa:standard:manage')" @click="handleAdd"><el-icon><Plus /></el-icon>新增标准</el-button></el-card>
    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="standard_code" label="标准编号" width="140" />
        <el-table-column prop="standard_name" label="标准名称" min-width="180" />
        <el-table-column prop="dept_name" label="适用部门" width="120" />
        <el-table-column prop="audit_frequency" label="审核频率" width="100">
          <template #default="{ row }"><el-tag>{{ getFrequencyText(row.audit_frequency) }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="question_count" label="问题项数量" width="100" />
        <el-table-column prop="passing_score" label="及格分数" width="90" />
        <el-table-column prop="is_active" label="状态" width="80">
          <template #default="{ row }"><el-tag :type="row.is_active === 1 ? 'success' : 'info'">{{ row.is_active === 1 ? '启用' : '停用' }}</el-tag></template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('quality:lpa:standard:manage')" @click="handleEdit(row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination"><el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" /></div>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="110px">
        <el-form-item label="标准编号" prop="standard_code"><el-input v-model="formData.standard_code" :disabled="!!formData.id" /></el-form-item>
        <el-form-item label="标准名称" prop="standard_name"><el-input v-model="formData.standard_name" /></el-form-item>
        <el-form-item label="版本"><el-input v-model="formData.version" /></el-form-item>
        <el-form-item label="适用部门"><el-input v-model="formData.dept_name" /></el-form-item>
        <el-form-item label="审核频率" prop="audit_frequency">
          <el-select v-model="formData.audit_frequency" placeholder="请选择">
            <el-option label="每日" value="DAILY" />
            <el-option label="每周" value="WEEKLY" />
            <el-option label="每月" value="MONTHLY" />
            <el-option label="每季度" value="QUARTERLY" />
          </el-select>
        </el-form-item>
        <el-form-item label="及格分数"><el-input-number v-model="formData.passing_score" :min="0" :max="100" /></el-form-item>
        <el-form-item label="是否启用">
          <el-switch v-model="formData.is_active" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <template #footer><el-button @click="dialogVisible = false">取消</el-button><el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button></template>
    </el-dialog>
    <el-dialog v-model="detailVisible" title="LPA审核标准详情" width="800px">
      <el-descriptions :column="2" border v-if="currentRow">
        <el-descriptions-item label="标准编号">{{ currentRow.standard_code }}</el-descriptions-item>
        <el-descriptions-item label="标准名称">{{ currentRow.standard_name }}</el-descriptions-item>
        <el-descriptions-item label="版本">{{ currentRow.version }}</el-descriptions-item>
        <el-descriptions-item label="适用部门">{{ currentRow.dept_name }}</el-descriptions-item>
        <el-descriptions-item label="审核频率">{{ getFrequencyText(currentRow.audit_frequency) }}</el-descriptions-item>
        <el-descriptions-item label="问题项数量">{{ currentRow.question_count }}</el-descriptions-item>
        <el-descriptions-item label="及格分数">{{ currentRow.passing_score }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ currentRow.is_active === 1 ? '启用' : '停用' }}</el-descriptions-item>
      </el-descriptions>
      <template #footer><el-button @click="detailVisible = false">关闭</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, FormInstance, FormRules } from 'element-plus'
import { getLPAStandardList, createLPAStandard, updateLPAStandard } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()
const loading = ref(false), tableData = ref<any[]>([]), dialogVisible = ref(false), detailVisible = ref(false), submitLoading = ref(false), formRef = ref<FormInstance>(), currentRow = ref<any>(null)
const searchForm = reactive({ standard_code: '', audit_frequency: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive<any>({ id: 0, standard_code: '', standard_name: '', version: '1.0', dept_name: '', audit_frequency: 'DAILY', passing_score: 80, is_active: 1 })
const rules: FormRules = { standard_code: [{ required: true, message: '请输入标准编号', trigger: 'blur' }], standard_name: [{ required: true, message: '请输入标准名称', trigger: 'blur' }], audit_frequency: [{ required: true, message: '请选择审核频率', trigger: 'change' }] }
const dialogTitle = computed(() => formData.id ? '编辑LPA标准' : '新增LPA标准')
const getFrequencyText = (s: string) => ({ DAILY: '每日', WEEKLY: '每周', MONTHLY: '每月', QUARTERLY: '每季度' }[s] || s)
const loadData = async () => { loading.value = true; try { const res = await getLPAStandardList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize }); tableData.value = res.data.list || []; pagination.total = res.data.total || 0 } finally { loading.value = false } }
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.standard_code = ''; searchForm.audit_frequency = ''; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, standard_code: '', standard_name: '', version: '1.0', dept_name: '', audit_frequency: 'DAILY', passing_score: 80, is_active: 1 }); dialogVisible.value = true }
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleView = (row: any) => { currentRow.value = row; detailVisible.value = true }
const handleSubmit = async () => { if (!formRef.value) return; await formRef.value.validate(); submitLoading.value = true; try { formData.id ? await updateLPAStandard(formData.id, formData) : await createLPAStandard(formData); ElMessage.success(formData.id ? '更新成功' : '创建成功'); dialogVisible.value = false; loadData() } finally { submitLoading.value = false } }
onMounted(() => { loadData() })
</script>
<style scoped lang="scss">.lpa-list { .search-card, .toolbar-card { margin-bottom: 16px; } .toolbar-card :deep(.el-card__body) { padding: 12px 16px; } .pagination { margin-top: 16px; display: flex; justify-content: flex-end; } }</style>
