<template>
  <div class="tenant-list">
    <!-- 搜索区域 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工厂名称">
          <el-input v-model="searchForm.tenant_name" placeholder="请输入工厂名称" clearable />
        </el-form-item>
        <el-form-item label="工厂代码">
          <el-input v-model="searchForm.tenant_key" placeholder="请输入工厂代码" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="正常" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 工具栏 -->
    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="danger" :disabled="!selectedRows.length" @click="handleBatchDelete">
        <el-icon><Delete /></el-icon>批量删除
      </el-button>
    </el-card>

    <!-- 表格 -->
    <el-card>
      <el-table
        v-loading="loading"
        :data="tableData"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="tenant_name" label="工厂名称" min-width="150" />
        <el-table-column prop="tenant_key" label="工厂代码" width="120" />
        <el-table-column prop="province" label="省份" width="100" />
        <el-table-column prop="city" label="城市" width="100" />
        <el-table-column prop="manager" label="负责人" width="100" />
        <el-table-column prop="employee_count" label="员工人数" width="100" />
        <el-table-column prop="factory_type" label="工厂类型" width="100" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="loadData"
          @current-change="loadData"
        />
      </div>
    </el-card>

    <!-- 编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="800px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="工厂名称" prop="tenant_name">
              <el-input v-model="formData.tenant_name" placeholder="请输入工厂名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="工厂代码" prop="tenant_key">
              <el-input v-model="formData.tenant_key" placeholder="请输入工厂代码(唯一)" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="省份">
              <el-input v-model="formData.province" placeholder="省份" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="城市">
              <el-input v-model="formData.city" placeholder="城市" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="区县">
              <el-input v-model="formData.district" placeholder="区县" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="详细地址">
          <el-input v-model="formData.address" placeholder="请输入详细地址" />
        </el-form-item>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="负责人/厂长">
              <el-input v-model="formData.manager" placeholder="负责人" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="联系人">
              <el-input v-model="formData.contact_name" placeholder="联系人" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="联系电话">
              <el-input v-model="formData.contact_phone" placeholder="电话" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="联系邮箱">
              <el-input v-model="formData.contact_email" placeholder="邮箱" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="工厂类型">
              <el-select v-model="formData.factory_type" placeholder="请选择工厂类型" clearable>
                <el-option label="离散制造" value="discrete" />
                <el-option label="流程制造" value="process" />
                <el-option label="混合制造" value="mixed" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="员工人数">
              <el-input-number v-model="formData.employee_count" :min="0" placeholder="人数" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="占地面积(m²)">
              <el-input-number v-model="formData.area_size" :min="0" :precision="2" placeholder="面积" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="年产能">
              <el-input-number v-model="formData.annual_capacity" :min="0" :precision="2" placeholder="产能" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="授权过期时间">
              <el-date-picker
                v-model="formData.expire_time"
                type="datetime"
                placeholder="选择日期时间"
                value-format="YYYY-MM-DD HH:mm:ss"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态">
              <el-radio-group v-model="formData.status">
                <el-radio :value="1">正常</el-radio>
                <el-radio :value="0">禁用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="备注说明" />
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
import { getTenantList, createTenant, updateTenant, deleteTenant } from '@/api/system'

const loading = ref(false)
const tableData = ref<any[]>([])
const selectedRows = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({
  tenant_name: '',
  tenant_key: '',
  status: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

const formData = reactive({
  id: 0,
  tenant_name: '',
  tenant_key: '',
  province: '',
  city: '',
  district: '',
  address: '',
  manager: '',
  contact_name: '',
  contact_phone: '',
  contact_email: '',
  factory_type: '',
  employee_count: null as number | null,
  area_size: null as number | null,
  annual_capacity: null as number | null,
  expire_time: '',
  status: 1,
  remark: ''
})

const rules: FormRules = {
  tenant_name: [{ required: true, message: '请输入工厂名称', trigger: 'blur' }],
  tenant_key: [{ required: true, message: '请输入工厂代码', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑工厂' : '新增工厂')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getTenantList({
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.tenant_name = ''
  searchForm.tenant_key = ''
  searchForm.status = ''
  handleSearch()
}

const handleSelectionChange = (rows: any[]) => {
  selectedRows.value = rows
}

const handleAdd = () => {
  Object.assign(formData, {
    id: 0,
    tenant_name: '',
    tenant_key: '',
    province: '',
    city: '',
    district: '',
    address: '',
    manager: '',
    contact_name: '',
    contact_phone: '',
    contact_email: '',
    factory_type: '',
    employee_count: null,
    area_size: null,
    annual_capacity: null,
    expire_time: '',
    status: 1,
    remark: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  await ElMessageBox.confirm('确定要删除该工厂吗？', '提示', { type: 'warning' })
  await deleteTenant(row.id)
  ElMessage.success('删除成功')
  loadData()
}

const handleBatchDelete = async () => {
  const ids = selectedRows.value.map(r => r.id)
  await ElMessageBox.confirm(`确定要删除选中的 ${ids.length} 个工厂吗？`, '提示', { type: 'warning' })
  ElMessage.success('删除成功')
  loadData()
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()

  submitLoading.value = true
  try {
    if (formData.id) {
      await updateTenant(formData.id, formData)
      ElMessage.success('更新成功')
    } else {
      await createTenant(formData)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } finally {
    submitLoading.value = false
  }
}

onMounted(() => {
  loadData()
})
</script>

<script lang="ts">
export default { name: 'TenantList' }
</script>

<style scoped lang="scss">
.tenant-list {
  .search-card, .toolbar-card {
    margin-bottom: 16px;
  }

  .toolbar-card :deep(.el-card__body) {
    padding: 12px 16px;
  }

  .pagination {
    margin-top: 16px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>
