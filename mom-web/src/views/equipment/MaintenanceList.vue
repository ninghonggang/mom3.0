<template>
  <div class="maintenance-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="设备编码"><el-input v-model="searchForm.equipment_code" placeholder="请输入" clearable /></el-form-item>
        <el-form-item label="保养日期"><el-date-picker v-model="searchForm.maintenance_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" /></el-form-item>
        <el-form-item><el-button type="primary" @click="handleSearch">查询</el-button><el-button @click="handleReset">重置</el-button></el-form-item>
      </el-form>
    </el-card>
    <el-card class="toolbar-card"><el-button type="primary" v-if="hasPermission('equipment:maintenance:add')" @click="handleAdd"><el-icon><Plus /></el-icon>新增</el-button></el-card>
    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="equipment_code" label="设备编码" width="120" />
        <el-table-column prop="equipment_name" label="设备名称" min-width="150" />
        <el-table-column prop="maintenance_item" label="保养项目" min-width="150" />
        <el-table-column prop="maintenance_type" label="保养类型" width="100" />
        <el-table-column prop="maintenance_user" label="保养人" width="100" />
        <el-table-column prop="maintenance_date" label="保养日期" width="120" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }"><el-tag :type="row.status === 1 ? 'success' : 'warning'">{{ row.status === 1 ? '已完成' : '进行中' }}</el-tag></template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }"><el-button link type="primary" v-if="hasPermission('equipment:maintenance:edit')" @click="handleEdit(row)">编辑</el-button><el-button link type="danger" v-if="hasPermission('equipment:maintenance:delete')" @click="handleDelete(row)">删除</el-button></template>
        </el-table-column>
      </el-table>
      <div class="pagination"><el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" /></div>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="设备" prop="equipment_id"><el-select v-model="formData.equipment_id" placeholder="请选择"><el-option label="设备A" :value="1" /><el-option label="设备B" :value="2" /></el-select></el-form-item>
        <el-form-item label="保养项目" prop="maintenance_item"><el-input v-model="formData.maintenance_item" /></el-form-item>
        <el-form-item label="保养类型" prop="maintenance_type"><el-select v-model="formData.maintenance_type"><el-option label="日常保养" value="日常保养" /><el-option label="一级保养" value="一级保养" /></el-select></el-form-item>
        <el-form-item label="保养人"><el-input v-model="formData.maintenance_user" /></el-form-item>
        <el-form-item label="保养日期" prop="maintenance_date"><el-date-picker v-model="formData.maintenance_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" /></el-form-item>
        <el-form-item label="状态" prop="status"><el-radio-group v-model="formData.status"><el-radio :value="1">已完成</el-radio><el-radio :value="0">进行中</el-radio></el-radio-group></el-form-item>
      </el-form>
      <template #footer><el-button @click="dialogVisible = false">取消</el-button><el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getEquipmentMaintenanceList, createEquipmentMaintenance, updateEquipmentMaintenance, deleteEquipmentMaintenance } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false), tableData = ref<any[]>([]), dialogVisible = ref(false), submitLoading = ref(false), formRef = ref<FormInstance>()
const searchForm = reactive({ equipment_code: '', maintenance_date: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, equipment_id: 0, maintenance_item: '', maintenance_type: '', maintenance_user: '', maintenance_date: '', status: 1 })
const rules: FormRules = { equipment_id: [{ required: true, message: '请选择设备', trigger: 'change' }], maintenance_item: [{ required: true, message: '请输入保养项目', trigger: 'blur' }], maintenance_date: [{ required: true, message: '请选择日期', trigger: 'change' }] }
const dialogTitle = computed(() => formData.id ? '编辑保养' : '新增保养')
const loadData = async () => { loading.value = true; try { const res = await getEquipmentMaintenanceList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize }); tableData.value = res.data.list || []; pagination.total = res.data.total || 0 } finally { loading.value = false } }
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.equipment_code = ''; searchForm.maintenance_date = ''; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, equipment_id: 0, maintenance_item: '', maintenance_type: '', maintenance_user: '', maintenance_date: '', status: 1 }); dialogVisible.value = true }
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除吗？', '提示', { type: 'warning' })
    await deleteEquipmentMaintenance(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}
const handleSubmit = async () => { if (!formRef.value) return; await formRef.value.validate(); submitLoading.value = true; try { formData.id ? await updateEquipmentMaintenance(formData.id, formData) : await createEquipmentMaintenance(formData); ElMessage.success(formData.id ? '更新成功' : '创建成功'); dialogVisible.value = false; loadData() } finally { submitLoading.value = false } }
onMounted(() => { loadData() })
</script>
<style scoped lang="scss">.maintenance-list { .search-card, .toolbar-card { margin-bottom: 16px; } .toolbar-card :deep(.el-card__body) { padding: 12px 16px; } .pagination { margin-top: 16px; display: flex; justify-content: flex-end; } }</style>
