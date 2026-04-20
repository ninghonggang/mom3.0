<template>
  <div class="delivery-address-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="客户ID">
          <el-input-number v-model="searchForm.customer_id" :min="1" placeholder="请输入客户ID" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mdm:delivery_address:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="address_name" label="地址别名" width="120" />
        <el-table-column prop="contact_person" label="联系人" width="100" />
        <el-table-column prop="contact_phone" label="联系电话" width="120" />
        <el-table-column prop="province" label="省" width="80" />
        <el-table-column prop="city" label="市" width="80" />
        <el-table-column prop="district" label="区" width="80" />
        <el-table-column prop="address_detail" label="详细地址" min-width="200" show-overflow-tooltip />
        <el-table-column prop="is_default" label="默认" width="60">
          <template #default="{ row }">
            <el-tag :type="row.is_default ? 'success' : 'info'" size="small">{{ row.is_default ? '是' : '否' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="is_active" label="状态" width="60">
          <template #default="{ row }">
            <el-tag :type="row.is_active ? 'success' : 'danger'">{{ row.is_active ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('mdm:delivery_address:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="primary" v-if="!row.is_default && hasPermission('mdm:delivery_address:set_default')" @click="handleSetDefault(row)">设为默认</el-button>
            <el-button link type="danger" v-if="hasPermission('mdm:delivery_address:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" :total="pagination.total" :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" @size-change="loadData" @current-change="loadData" />
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="客户ID" prop="customer_id">
          <el-input-number v-model="formData.customer_id" :min="1" />
        </el-form-item>
        <el-form-item label="地址别名"><el-input v-model="formData.address_name" /></el-form-item>
        <el-form-item label="联系人"><el-input v-model="formData.contact_person" /></el-form-item>
        <el-form-item label="联系电话"><el-input v-model="formData.contact_phone" /></el-form-item>
        <el-form-item label="省"><el-input v-model="formData.province" /></el-form-item>
        <el-form-item label="市"><el-input v-model="formData.city" /></el-form-item>
        <el-form-item label="区"><el-input v-model="formData.district" /></el-form-item>
        <el-form-item label="详细地址"><el-input v-model="formData.address_detail" type="textarea" /></el-form-item>
        <el-form-item label="默认地址">
          <el-switch v-model="formData.is_default" />
        </el-form-item>
        <el-form-item label="启用状态">
          <el-switch v-model="formData.is_active" />
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
import { getDeliveryAddressList, getDeliveryAddressById, createDeliveryAddress, updateDeliveryAddress, deleteDeliveryAddress } from '@/api/mdm'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const searchForm = reactive({ customer_id: undefined as number | undefined })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, customer_id: 0, address_name: '', contact_person: '', contact_phone: '', province: '', city: '', district: '', address_detail: '', is_default: false, is_active: true })
const rules: FormRules = { customer_id: [{ required: true, message: '请输入客户ID', trigger: 'blur' }] }
const dialogTitle = computed(() => formData.id ? '编辑收货地址' : '新增收货地址')

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.customer_id) {
      params.customer_id = searchForm.customer_id
    }
    const res = await getDeliveryAddressList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}
const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.customer_id = undefined; handleSearch() }
const handleAdd = () => { Object.assign(formData, { id: 0, customer_id: 0, address_name: '', contact_person: '', contact_phone: '', province: '', city: '', district: '', address_detail: '', is_default: false, is_active: true }); dialogVisible.value = true }
const handleEdit = async (row: any) => {
  try {
    const res = await getDeliveryAddressById(row.id)
    Object.assign(formData, res.data)
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取收货地址详情失败')
  }
}
const handleSetDefault = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定将该地址设为默认收货地址吗？', '提示', { type: 'warning' })
    await updateDeliveryAddress(row.id, { ...row, is_default: true })
    ElMessage.success('设置成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该收货地址吗？', '提示', { type: 'warning' })
    await deleteDeliveryAddress(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (formData.id) {
      await updateDeliveryAddress(formData.id, formData)
    } else {
      await createDeliveryAddress(formData)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally {
    submitLoading.value = false
  }
}
onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.delivery-address-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
