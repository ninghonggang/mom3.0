<template>
  <div class="data-point-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="采集点编码">
          <el-input v-model="searchForm.query" placeholder="请输入采集点编码/名称" clearable />
        </el-form-item>
        <el-form-item label="协议">
          <el-select v-model="searchForm.protocol" placeholder="请选择协议" clearable>
            <el-option label="OPC-UA" value="OPCUA" />
            <el-option label="MQTT" value="MQTT" />
            <el-option label="Modbus" value="MODBUS" />
            <el-option label="手动" value="MANUAL" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('wms:data-point:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="point_code" label="采集点编码" min-width="120" />
        <el-table-column prop="point_name" label="采集点名称" min-width="150" />
        <el-table-column prop="device_id" label="设备ID" width="80" />
        <el-table-column prop="protocol" label="协议" width="100">
          <template #default="{ row }">
            <el-tag>{{ row.protocol }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="data_type" label="数据类型" width="100" />
        <el-table-column prop="address" label="地址" min-width="180" show-overflow-tooltip />
        <el-table-column prop="scan_rate" label="采集周期(ms)" width="110" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 'ACTIVE' ? 'success' : 'info'">
              {{ row.status === 'ACTIVE' ? '启用' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('wms:data-point:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('wms:data-point:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-form-item label="采集点编码" prop="point_code">
          <el-input v-model="formData.point_code" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="采集点名称" prop="point_name">
          <el-input v-model="formData.point_name" />
        </el-form-item>
        <el-form-item label="设备ID" prop="device_id">
          <el-input-number v-model="formData.device_id" :min="0" />
        </el-form-item>
        <el-form-item label="协议" prop="protocol">
          <el-select v-model="formData.protocol">
            <el-option label="OPC-UA" value="OPCUA" />
            <el-option label="MQTT" value="MQTT" />
            <el-option label="Modbus" value="MODBUS" />
            <el-option label="手动" value="MANUAL" />
          </el-select>
        </el-form-item>
        <el-form-item label="数据类型" prop="data_type">
          <el-select v-model="formData.data_type">
            <el-option label="模拟量" value="ANALOG" />
            <el-option label="数字量" value="DIGITAL" />
            <el-option label="文本" value="TEXT" />
            <el-option label="枚举" value="ENUM" />
          </el-select>
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model="formData.address" placeholder="OPC节点ID/MQTT Topic/Modbus地址" />
        </el-form-item>
        <el-form-item label="单位">
          <el-input v-model="formData.unit" />
        </el-form-item>
        <el-form-item label="采集周期(ms)">
          <el-input-number v-model="formData.scan_rate" :min="100" />
        </el-form-item>
        <el-form-item label="存储策略">
          <el-select v-model="formData.store_policy">
            <el-option label="全部存储" value="ALL" />
            <el-option label="变化存储" value="CHANGE" />
            <el-option label="周期存储" value="PERIOD" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="formData.status">
            <el-radio value="ACTIVE">启用</el-radio>
            <el-radio value="INACTIVE">停用</el-radio>
          </el-radio-group>
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
import { getDataPointList, createDataPoint, updateDataPoint, deleteDataPoint } from '@/api/wms'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ query: '', protocol: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({
  id: 0, point_code: '', point_name: '', device_id: 0, protocol: 'OPCUA',
  data_type: 'ANALOG', address: '', unit: '', scan_rate: 1000, store_policy: 'ALL', status: 'ACTIVE'
})

const rules: FormRules = {
  point_code: [{ required: true, message: '请输入采集点编码', trigger: 'blur' }],
  point_name: [{ required: true, message: '请输入采集点名称', trigger: 'blur' }],
  device_id: [{ required: true, message: '请输入设备ID', trigger: 'blur' }],
  protocol: [{ required: true, message: '请选择协议', trigger: 'change' }],
}

const dialogTitle = computed(() => formData.id ? '编辑采集点' : '新增采集点')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDataPointList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.query = ''; searchForm.protocol = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, point_code: '', point_name: '', device_id: 0, protocol: 'OPCUA',
    data_type: 'ANALOG', address: '', unit: '', scan_rate: 1000, store_policy: 'ALL', status: 'ACTIVE' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该采集点吗？', '提示', { type: 'warning' })
    await deleteDataPoint(row.id)
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
    formData.id ? await updateDataPoint(formData.id, formData) : await createDataPoint(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.data-point-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
