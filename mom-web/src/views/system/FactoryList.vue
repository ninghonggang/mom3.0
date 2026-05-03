<template>
  <div class="factory-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工厂编码">
          <el-input v-model="searchForm.factory_code" placeholder="请输入工厂编码" clearable />
        </el-form-item>
        <el-form-item label="工厂名称">
          <el-input v-model="searchForm.factory_name" placeholder="请输入工厂名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
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
        <el-icon><Plus /></el-icon>新增工厂
      </el-button>
      <el-button type="success" @click="handleSwitchFactory" :disabled="!currentFactory">
        <el-icon><Switch /></el-icon>切换当前工厂: {{ currentFactory?.factory_name || '未选择' }}
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="factory_code" label="工厂编码" width="120" />
        <el-table-column prop="factory_name" label="工厂名称" min-width="150" />
        <el-table-column prop="province" label="省份" width="100" />
        <el-table-column prop="city" label="城市" width="100" />
        <el-table-column prop="manager" label="负责人" width="100" />
        <el-table-column prop="phone" label="联系电话" width="120" />
        <el-table-column prop="area_size" label="面积(m²)" width="100">
          <template #default="{ row }">
            {{ row.area_size || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="is_default" label="默认" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.is_default === 1" type="success">是</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="warning" size="small" v-if="row.is_default !== 1" @click="handleSetDefault(row)">设为默认</el-button>
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
        <el-form-item label="工厂编码" prop="factory_code">
          <el-input v-model="formData.factory_code" placeholder="请输入工厂编码" />
        </el-form-item>
        <el-form-item label="工厂名称" prop="factory_name">
          <el-input v-model="formData.factory_name" placeholder="请输入工厂名称" />
        </el-form-item>
        <el-form-item label="省份">
          <el-input v-model="formData.province" placeholder="请输入省份" />
        </el-form-item>
        <el-form-item label="城市">
          <el-input v-model="formData.city" placeholder="请输入城市" />
        </el-form-item>
        <el-form-item label="详细地址">
          <el-input v-model="formData.address" placeholder="请输入详细地址" />
        </el-form-item>
        <el-form-item label="负责人">
          <el-input v-model="formData.manager" placeholder="请输入负责人" />
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="formData.phone" placeholder="请输入联系电话" />
        </el-form-item>
        <el-form-item label="占地面积(m²)">
          <el-input-number v-model="formData.area_size" :min="0" :precision="2" style="width: 100%" />
        </el-form-item>
        <el-form-item label="设为默认">
          <el-switch v-model="formData.is_default" :true-value="1" :false-value="0" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>

    <!-- 切换工厂对话框 -->
    <el-dialog v-model="switchDialogVisible" title="切换当前工厂" width="500px">
      <el-radio-group v-model="selectedFactoryId" style="width: 100%">
        <el-radio v-for="f in tableData" :key="f.id" :value="f.id" style="display: block; margin-bottom: 10px">
          {{ f.factory_name }} ({{ f.factory_code }})
        </el-radio>
      </el-radio-group>
      <template #footer>
        <el-button @click="switchDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleConfirmSwitch">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getFactoryList, createFactory, updateFactory, deleteFactory, setDefaultFactory, getCurrentFactory, setCurrentFactory } from '@/api/system/factory'

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const switchDialogVisible = ref(false)
const tableData = ref<any[]>([])
const currentFactory = ref<any>(null)
const selectedFactoryId = ref<number | null>(null)
const dialogTitle = ref('新增工厂')

const searchForm = reactive({
  factory_code: '',
  factory_name: '',
  status: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  factory_code: '',
  factory_name: '',
  province: '',
  city: '',
  district: '',
  address: '',
  manager: '',
  phone: '',
  area_size: null,
  is_default: 0,
  status: 1
})

const formRules = {
  factory_code: [{ required: true, message: '请输入工厂编码', trigger: 'blur' }],
  factory_name: [{ required: true, message: '请输入工厂名称', trigger: 'blur' }]
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getFactoryList({
      page: pagination.page,
      page_size: pagination.pageSize,
      factory_code: searchForm.factory_code,
      factory_name: searchForm.factory_name,
      status: searchForm.status
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadCurrentFactory = async () => {
  try {
    const res = await getCurrentFactory()
    currentFactory.value = res.data
  } catch (e) {
    console.error(e)
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.factory_code = ''; searchForm.factory_name = ''; searchForm.status = ''; handleSearch() }

const handleCreate = () => {
  dialogTitle.value = '新增工厂'
  formData.factory_code = ''
  formData.factory_name = ''
  formData.province = ''
  formData.city = ''
  formData.address = ''
  formData.manager = ''
  formData.phone = ''
  formData.area_size = null
  formData.is_default = 0
  formData.status = 1
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  dialogTitle.value = '编辑工厂'
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleSave = async () => {
  if (!formData.factory_code || !formData.factory_name) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    if (formData.id) {
      await updateFactory(formData.id, formData)
    } else {
      await createFactory(formData)
    }
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleSetDefault = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定将"${row.factory_name}"设为默认工厂吗？`, '提示', { type: 'warning' })
    await setDefaultFactory(row.id)
    ElMessage.success('设置成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定删除"${row.factory_name}"吗？`, '提示', { type: 'warning' })
    await deleteFactory(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

const handleSwitchFactory = () => {
  selectedFactoryId.value = currentFactory.value?.id || null
  switchDialogVisible.value = true
}

const handleConfirmSwitch = async () => {
  if (!selectedFactoryId.value) {
    ElMessage.warning('请选择工厂')
    return
  }
  try {
    await setCurrentFactory(selectedFactoryId.value)
    ElMessage.success('切换成功')
    switchDialogVisible.value = false
    loadCurrentFactory()
  } catch (e) {
    console.error(e)
    ElMessage.error('切换失败')
  }
}

onMounted(() => { loadData(); loadCurrentFactory() })
</script>

<style scoped lang="scss">
.factory-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>