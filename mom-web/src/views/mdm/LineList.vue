<template>
  <div class="line-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="生产线编码">
          <el-input v-model="searchForm.line_code" placeholder="请输入生产线编码" clearable />
        </el-form-item>
        <el-form-item label="生产线名称">
          <el-input v-model="searchForm.line_name" placeholder="请输入生产线名称" clearable />
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
        <el-table-column prop="line_code" label="生产线编码" width="120" />
        <el-table-column prop="line_name" label="生产线名称" min-width="150" />
        <el-table-column prop="workshop_name" label="所属车间" width="120" />
        <el-table-column prop="line_type" label="生产线类型" width="100" />
        <el-table-column prop="manager" label="负责人" width="100" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
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
        <el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" />
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="生产线编码" prop="line_code"><el-input v-model="formData.line_code" :disabled="!!formData.id" /></el-form-item>
        <el-form-item label="生产线名称" prop="line_name"><el-input v-model="formData.line_name" /></el-form-item>
        <el-form-item label="所属车间" prop="workshop_id">
          <el-select v-model="formData.workshop_id" placeholder="请选择车间">
            <el-option label="一车间" :value="1" /><el-option label="二车间" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="生产线类型"><el-input v-model="formData.line_type" /></el-form-item>
        <el-form-item label="负责人"><el-input v-model="formData.manager" /></el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status"><el-radio :value="1">启用</el-radio><el-radio :value="0">禁用</el-radio></el-radio-group>
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
import { getProductionLineList, createProductionLine, updateProductionLine, deleteProductionLine } from '@/api/mdm'

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const searchForm = reactive({ line_code: '', line_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, line_code: '', line_name: '', workshop_id: 0, line_type: '', manager: '', status: 1 })
const rules: FormRules = { line_code: [{ required: true, message: '请输入生产线编码', trigger: 'blur' }], line_name: [{ required: true, message: '请输入生产线名称', trigger: 'blur' }] }
const dialogTitle = computed(() => formData.id ? '编辑生产线' : '新增生产线')

const loadData = async () => { loading.value = true; try { const res = await getProductionLineList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize }); tableData.value = res.data.list || []; pagination.total = res.data.total || 0 } finally { loading.value = false } }
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.line_code = ''; searchForm.line_name = ''; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, line_code: '', line_name: '', workshop_id: 0, line_type: '', manager: '', status: 1 }); dialogVisible.value = true }
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleDelete = async (row: any) => { await ElMessageBox.confirm('确定删除该生产线吗？', '提示', { type: 'warning' }); await deleteProductionLine(row.id); ElMessage.success('删除成功'); loadData() }
const handleSubmit = async () => { if (!formRef.value) return; await formRef.value.validate(); submitLoading.value = true; try { formData.id ? await updateProductionLine(formData.id, formData) : await createProductionLine(formData); ElMessage.success(formData.id ? '更新成功' : '创建成功'); dialogVisible.value = false; loadData() } finally { submitLoading.value = false } }
onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.line-list { .search-card, .toolbar-card { margin-bottom: 16px; } .toolbar-card :deep(.el-card__body) { padding: 12px 16px; } .pagination { margin-top: 16px; display: flex; justify-content: flex-end; } }
</style>
