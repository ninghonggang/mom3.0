<template>
  <div class="customer-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="客户编码">
          <el-input v-model="searchForm.code" placeholder="请输入客户编码" clearable />
        </el-form-item>
        <el-form-item label="客户名称">
          <el-input v-model="searchForm.name" placeholder="请输入客户名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:customer:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="code" label="客户编码" width="120" />
        <el-table-column prop="name" label="客户名称" min-width="150" />
        <el-table-column prop="type" label="客户类型" width="100" />
        <el-table-column prop="contact" label="联系人" width="100" />
        <el-table-column prop="phone" label="联系电话" width="120" />
        <el-table-column prop="email" label="邮箱" width="150" />
        <el-table-column prop="address" label="地址" min-width="200" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('mdm:customer:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:customer:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" />
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="客户编码" prop="code"><el-input v-model="formData.code" :disabled="!!formData.id" /></el-form-item>
        <el-form-item label="客户名称" prop="name"><el-input v-model="formData.name" /></el-form-item>
        <el-form-item label="客户类型"><el-input v-model="formData.type" /></el-form-item>
        <el-form-item label="联系人"><el-input v-model="formData.contact" /></el-form-item>
        <el-form-item label="联系电话"><el-input v-model="formData.phone" /></el-form-item>
        <el-form-item label="邮箱"><el-input v-model="formData.email" /></el-form-item>
        <el-form-item label="地址"><el-input v-model="formData.address" type="textarea" /></el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status"><el-radio :value="1">启用</el-radio><el-radio :value="2">禁用</el-radio></el-radio-group>
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
import { getCustomerList, createCustomer, updateCustomer, deleteCustomer } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const searchForm = reactive({ code: '', name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, code: '', name: '', type: '', contact: '', phone: '', email: '', address: '', status: 1 })
const rules: FormRules = { code: [{ required: true, message: '请输入客户编码', trigger: 'blur' }], name: [{ required: true, message: '请输入客户名称', trigger: 'blur' }] }
const dialogTitle = computed(() => formData.id ? '编辑客户' : '新增客户')

const loadData = async () => { loading.value = true; try { const res = await getCustomerList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize }); tableData.value = res.data.list || []; pagination.total = res.data.total || 0 } finally { loading.value = false } }
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.code = ''; searchForm.name = ''; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, code: '', name: '', type: '', contact: '', phone: '', email: '', address: '', status: 1 }); dialogVisible.value = true }
const handleEdit = (row: any) => { Object.assign(formData, row); dialogVisible.value = true }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该客户吗？', '提示', { type: 'warning' })
    await deleteCustomer(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}
const handleSubmit = async () => { if (!formRef.value) return; await formRef.value.validate(); submitLoading.value = true; try { formData.id ? await updateCustomer(formData.id, formData) : await createCustomer(formData); ElMessage.success(formData.id ? '更新成功' : '创建成功'); dialogVisible.value = false; loadData() } finally { submitLoading.value = false } }
onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.customer-list { .search-card, .toolbar-card { margin-bottom: 16px; } .toolbar-card :deep(.el-card__body) { padding: 12px 16px; } .pagination { margin-top: 16px; display: flex; justify-content: flex-end; } }
</style>
