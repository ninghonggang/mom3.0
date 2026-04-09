<template>
  <div class="workcenter-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工作中心编码">
          <el-input v-model="searchForm.work_center_code" placeholder="请输入编码" clearable />
        </el-form-item>
        <el-form-item label="工作中心名称">
          <el-input v-model="searchForm.work_center_name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="激活" value="ACTIVE" />
            <el-option label="停用" value="INACTIVE" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:workcenter:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="work_center_code" label="工作中心编码" width="140" />
        <el-table-column prop="work_center_name" label="工作中心名称" min-width="150" />
        <el-table-column prop="work_center_type" label="类型" width="100">
          <template #default="{ row }">
            {{ getTypeText(row.work_center_type) }}
          </template>
        </el-table-column>
        <el-table-column prop="workshop_id" label="所属车间" width="100" />
        <el-table-column prop="standard_capacity" label="标准产能" width="100" />
        <el-table-column prop="max_capacity" label="最大产能" width="100" />
        <el-table-column prop="efficiency_factor" label="效率因子(%)" width="110" />
        <el-table-column prop="utilization_target" label="利用率目标(%)" width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="row.status === 'ACTIVE' ? 'success' : 'info'">
              {{ row.status === 'ACTIVE' ? '激活' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:workcenter:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:workcenter:delete')" @click="handleDelete(row)">删除</el-button>
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
      <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="工作中心编码" prop="work_center_code">
          <el-input v-model="form.work_center_code" placeholder="请输入编码" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="工作中心名称" prop="work_center_name">
          <el-input v-model="form.work_center_name" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="工作中心类型" prop="work_center_type">
          <el-select v-model="form.work_center_type" placeholder="请选择类型">
            <el-option label="产线" value="LINE" />
            <el-option label="工作中心" value="WORKCENTER" />
            <el-option label="设备" value="EQUIPMENT" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属车间" prop="workshop_id">
          <el-input-number v-model="form.workshop_id" :min="0" />
        </el-form-item>
        <el-form-item label="产能单位" prop="capacity_unit">
          <el-select v-model="form.capacity_unit" placeholder="请选择单位">
            <el-option label="小时" value="HOUR" />
            <el-option label="分钟" value="MINUTE" />
            <el-option label="件" value="PCS" />
          </el-select>
        </el-form-item>
        <el-form-item label="标准产能" prop="standard_capacity">
          <el-input-number v-model="form.standard_capacity" :min="0" :precision="3" />
        </el-form-item>
        <el-form-item label="最大产能" prop="max_capacity">
          <el-input-number v-model="form.max_capacity" :min="0" :precision="3" />
        </el-form-item>
        <el-form-item label="效率因子(%)" prop="efficiency_factor">
          <el-input-number v-model="form.efficiency_factor" :min="0" :max="200" :precision="2" />
        </el-form-item>
        <el-form-item label="利用率目标(%)" prop="utilization_target">
          <el-input-number v-model="form.utilization_target" :min="0" :max="100" :precision="2" />
        </el-form-item>
        <el-form-item label="准备时间(分钟)" prop="setup_time">
          <el-input-number v-model="form.setup_time" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio label="ACTIVE">激活</el-radio>
            <el-radio label="INACTIVE">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getWorkCenterList, createWorkCenter, updateWorkCenter, deleteWorkCenter } from '@/api/aps'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('新增工作中心')
const isEdit = ref(false)
const formRef = ref()

const searchForm = reactive({ work_center_code: '', work_center_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const form = reactive({
  id: 0,
  work_center_code: '',
  work_center_name: '',
  work_center_type: 'LINE',
  workshop_id: 0,
  capacity_unit: 'HOUR',
  standard_capacity: 1,
  max_capacity: 1,
  efficiency_factor: 100,
  utilization_target: 85,
  setup_time: 0,
  status: 'ACTIVE',
  description: ''
})

const rules = {
  work_center_code: [{ required: true, message: '请输入工作中心编码', trigger: 'blur' }],
  work_center_name: [{ required: true, message: '请输入工作中心名称', trigger: 'blur' }]
}

const getTypeText = (type: string) => {
  const map: Record<string, string> = { LINE: '产线', WORKCENTER: '工作中心', EQUIPMENT: '设备' }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getWorkCenterList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.work_center_code = ''; searchForm.work_center_name = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增工作中心'
  Object.assign(form, {
    id: 0,
    work_center_code: '',
    work_center_name: '',
    work_center_type: 'LINE',
    workshop_id: 0,
    capacity_unit: 'HOUR',
    standard_capacity: 1,
    max_capacity: 1,
    efficiency_factor: 100,
    utilization_target: 85,
    setup_time: 0,
    status: 'ACTIVE',
    description: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑工作中心'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateWorkCenter(form.id, form)
      ElMessage.success('更新成功')
    } else {
      await createWorkCenter(form)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } catch (error: any) {
    ElMessage.error(error.message || '操作失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该工作中心吗？', '提示', { type: 'warning' })
    await deleteWorkCenter(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.workcenter-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
