<template>
  <div class="gauge-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="量检具编码">
          <el-input v-model="searchForm.gauge_code" placeholder="请输入编码" clearable />
        </el-form-item>
        <el-form-item label="量检具名称">
          <el-input v-model="searchForm.gauge_name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="规格型号">
          <el-input v-model="searchForm.specification" placeholder="请输入规格型号" clearable />
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
      <el-button type="primary" v-if="hasPermission('equipment:gauge:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="gauge_code" label="量检具编码" width="140" />
        <el-table-column prop="gauge_name" label="量检具名称" min-width="150" />
        <el-table-column prop="specification" label="规格型号" width="120" />
        <el-table-column prop="gauge_type" label="类型" width="100">
          <template #default="{ row }">
            {{ getTypeText(row.gauge_type) }}
          </template>
        </el-table-column>
        <el-table-column prop="measurement_range" label="测量范围" width="120">
          <template #default="{ row }">
            {{ row.measurement_range || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="precision_level" label="精度等级" width="100">
          <template #default="{ row }">
            {{ row.precision_level || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="location" label="存放位置" width="120">
          <template #default="{ row }">
            {{ row.location || '-' }}
          </template>
        </el-table-column>
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
            <el-button link type="primary" size="small" v-if="hasPermission('equipment:gauge:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('equipment:gauge:delete')" @click="handleDelete(row)">删除</el-button>
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
        <el-form-item label="量检具编码" prop="gauge_code">
          <el-input v-model="form.gauge_code" placeholder="请输入编码" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="量检具名称" prop="gauge_name">
          <el-input v-model="form.gauge_name" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="规格型号" prop="specification">
          <el-input v-model="form.specification" placeholder="请输入规格型号" />
        </el-form-item>
        <el-form-item label="量检具类型" prop="gauge_type">
          <el-select v-model="form.gauge_type" placeholder="请选择类型">
            <el-option label="卡尺" value="CALIPER" />
            <el-option label="千分尺" value="MICROMETER" />
            <el-option label="百分表" value="DIAL_INDICATOR" />
            <el-option label="高度尺" value="HEIGHT_GAUGE" />
            <el-option label="量块" value="GAUGE_BLOCK" />
            <el-option label="塞尺" value="FEELER_GAUGE" />
            <el-option label="角度尺" value="ANGLE_GAUGE" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item label="测量范围" prop="measurement_range">
          <el-input v-model="form.measurement_range" placeholder="请输入测量范围" />
        </el-form-item>
        <el-form-item label="精度等级" prop="precision_level">
          <el-input v-model="form.precision_level" placeholder="请输入精度等级" />
        </el-form-item>
        <el-form-item label="存放位置" prop="location">
          <el-input v-model="form.location" placeholder="请输入存放位置" />
        </el-form-item>
        <el-form-item label="管理周期(天)" prop="management_cycle">
          <el-input-number v-model="form.management_cycle" :min="0" />
        </el-form-item>
        <el-form-item label="检定周期(天)" prop="verification_cycle">
          <el-input-number v-model="form.verification_cycle" :min="0" />
        </el-form-item>
        <el-form-item label="最近检定日期" prop="last_verification_date">
          <el-date-picker v-model="form.last_verification_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
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
import { getGaugeList, createGauge, updateGauge, deleteGauge } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('新增量检具')
const isEdit = ref(false)
const formRef = ref()

const searchForm = reactive({
  gauge_code: '',
  gauge_name: '',
  specification: '',
  status: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const form = reactive({
  id: 0,
  gauge_code: '',
  gauge_name: '',
  specification: '',
  gauge_type: 'CALIPER',
  measurement_range: '',
  precision_level: '',
  location: '',
  management_cycle: 30,
  verification_cycle: 90,
  last_verification_date: '',
  status: 'ACTIVE',
  description: ''
})

const rules = {
  gauge_code: [{ required: true, message: '请输入量检具编码', trigger: 'blur' }],
  gauge_name: [{ required: true, message: '请输入量检具名称', trigger: 'blur' }]
}

const getTypeText = (type: string) => {
  const map: Record<string, string> = {
    CALIPER: '卡尺',
    MICROMETER: '千分尺',
    DIAL_INDICATOR: '百分表',
    HEIGHT_GAUGE: '高度尺',
    GAUGE_BLOCK: '量块',
    FEELER_GAUGE: '塞尺',
    ANGLE_GAUGE: '角度尺',
    OTHER: '其他'
  }
  return map[type] || type
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getGaugeList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.gauge_code = ''
  searchForm.gauge_name = ''
  searchForm.specification = ''
  searchForm.status = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增量检具'
  Object.assign(form, {
    id: 0,
    gauge_code: '',
    gauge_name: '',
    specification: '',
    gauge_type: 'CALIPER',
    measurement_range: '',
    precision_level: '',
    location: '',
    management_cycle: 30,
    verification_cycle: 90,
    last_verification_date: '',
    status: 'ACTIVE',
    description: ''
  })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  dialogTitle.value = '编辑量检具'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateGauge(form.id, form)
      ElMessage.success('更新成功')
    } else {
      await createGauge(form)
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
    await ElMessageBox.confirm('确定删除该量检具吗？', '提示', { type: 'warning' })
    await deleteGauge(row.id)
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
.gauge-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
