<template>
  <div class="page-container">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="配置编码">
          <el-input v-model="searchForm.config_code" placeholder="请输入配置编码" clearable />
        </el-form-item>
        <el-form-item label="配置名称">
          <el-input v-model="searchForm.config_name" placeholder="请输入配置名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.is_enabled" placeholder="请选择" clearable>
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
      <el-button type="primary" v-if="hasPermission('aps:rolling-config:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="config_code" label="配置编码" width="140" />
        <el-table-column prop="config_name" label="配置名称" min-width="150" />
        <el-table-column prop="config_type" label="配置类型" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.config_type === 'DAILY'" type="success">每日</el-tag>
            <el-tag v-else-if="row.config_type === 'HOURLY'" type="warning">小时</el-tag>
            <el-tag v-else type="info">事件</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="trigger_type" label="触发方式" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.trigger_type === 'CRON'" type="info">Cron</el-tag>
            <el-tag v-else-if="row.trigger_type === 'MANUAL'" type="warning">手动</el-tag>
            <el-tag v-else type="info">事件</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="scheduling_algorithm" label="算法" width="120" />
        <el-table-column prop="horizon_days" label="视野(天)" width="80" />
        <el-table-column prop="lead_time_days" label="提前期(天)" width="90" />
        <el-table-column prop="is_enabled" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.is_enabled === 1 ? 'success' : 'danger'">
              {{ row.is_enabled === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="last_execute_time" label="最后执行" width="160" />
        <el-table-column prop="auto_execute" label="自动执行" width="90">
          <template #default="{ row }">
            <el-tag :type="row.auto_execute === 1 ? 'success' : 'info'" size="small">
              {{ row.auto_execute === 1 ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:rolling-config:edit')" @click.stop="handleEdit(row)">编辑</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('aps:rolling-config:execute')" @click.stop="handleExecute(row)">执行</el-button>
            <el-button link type="warning" size="small" v-if="row.is_enabled === 1 && hasPermission('aps:rolling-config:disable')" @click.stop="handleDisable(row)">禁用</el-button>
            <el-button link type="success" size="small" v-if="row.is_enabled === 0 && hasPermission('aps:rolling-config:enable')" @click.stop="handleEnable(row)">启用</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:rolling-config:delete')" @click.stop="handleDelete(row)">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form :model="formData" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="配置编码" required>
              <el-input v-model="formData.config_code" placeholder="请输入" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="配置名称" required>
              <el-input v-model="formData.config_name" placeholder="请输入" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="配置类型">
              <el-select v-model="formData.config_type" placeholder="请选择" style="width: 100%">
                <el-option label="每日" value="DAILY" />
                <el-option label="小时" value="HOURLY" />
                <el-option label="事件" value="EVENT" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="触发方式">
              <el-select v-model="formData.trigger_type" placeholder="请选择" style="width: 100%">
                <el-option label="Cron" value="CRON" />
                <el-option label="手动" value="MANUAL" />
                <el-option label="事件" value="EVENT" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="Cron表达式">
              <el-input v-model="formData.trigger_cron" placeholder="如: 0 0 * * * *" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排程视野(天)">
              <el-input-number v-model="formData.horizon_days" :min="1" :max="30" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="提前下单期(天)">
              <el-input-number v-model="formData.lead_time_days" :min="0" :max="30" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排程算法">
              <el-select v-model="formData.scheduling_algorithm" placeholder="请选择" style="width: 100%">
                <el-option label="FIFO" value="FIFO" />
                <el-option label="EDD" value="EDD" />
                <el-option label="SPT" value="SPT" />
                <el-option label="LPT" value="LPT" />
                <el-option label="JIT优先" value="JIT_FIRST" />
                <el-option label="关键比" value="CR" />
                <el-option label="产品族" value="FAMILY" />
                <el-option label="瓶颈" value="BOTTLENECK" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="排程方向">
              <el-select v-model="formData.direction" placeholder="请选择" style="width: 100%">
                <el-option label="正向" value="FORWARD" />
                <el-option label="反向" value="BACKWARD" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="优化目标">
              <el-select v-model="formData.optimize_target" placeholder="请选择" style="width: 100%">
                <el-option label="交付" value="DELIVERY" />
                <el-option label="利用率" value="UTILITY" />
                <el-option label="均衡" value="EQUILIBRIUM" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="换型时间占比上限(%)">
              <el-input-number v-model="formData.max_changeover_pct" :min="0" :max="100" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="最低资源利用率(%)">
              <el-input-number v-model="formData.min_resource_utilization" :min="0" :max="100" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="允许加班">
              <el-switch v-model="formData.allow_overtime" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="产品族聚类">
              <el-switch v-model="formData.family_grouping" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="自动执行">
              <el-switch v-model="formData.auto_execute" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="启用">
              <el-switch v-model="formData.is_enabled" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="锁定任务处理">
          <el-radio-group v-model="formData.locked_tasks_handling">
            <el-radio label="RETAIN">保留</el-radio>
            <el-radio label="ADJUST">调整</el-radio>
            <el-radio label="IGNORE">忽略</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getRollingConfigList,
  createRollingConfig,
  updateRollingConfig,
  deleteRollingConfig,
  enableRollingConfig,
  disableRollingConfig,
  executeRollingConfig
} from '@/api/aps'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const dialogTitle = ref('新增配置')
const tableData = ref<any[]>([])
const isEdit = ref(false)
const currentId = ref<number | null>(null)

const searchForm = reactive({
  config_code: '',
  config_name: '',
  is_enabled: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = ref<any>({
  config_code: '',
  config_name: '',
  config_type: 'DAILY',
  trigger_type: 'MANUAL',
  trigger_cron: '',
  horizon_days: 7,
  lead_time_days: 3,
  scheduling_algorithm: 'FIFO',
  direction: 'FORWARD',
  optimize_target: 'DELIVERY',
  max_changeover_pct: 30,
  min_resource_utilization: 70,
  allow_overtime: 0,
  family_grouping: 1,
  locked_tasks_handling: 'RETAIN',
  auto_execute: 0,
  is_enabled: 1,
  remark: ''
})

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getRollingConfigList(params)
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
  searchForm.config_code = ''
  searchForm.config_name = ''
  searchForm.is_enabled = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  currentId.value = null
  dialogTitle.value = '新增配置'
  formData.value = {
    config_code: '',
    config_name: '',
    config_type: 'DAILY',
    trigger_type: 'MANUAL',
    trigger_cron: '',
    horizon_days: 7,
    lead_time_days: 3,
    scheduling_algorithm: 'FIFO',
    direction: 'FORWARD',
    optimize_target: 'DELIVERY',
    max_changeover_pct: 30,
    min_resource_utilization: 70,
    allow_overtime: 0,
    family_grouping: 1,
    locked_tasks_handling: 'RETAIN',
    auto_execute: 0,
    is_enabled: 1,
    remark: ''
  }
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  currentId.value = row.id
  dialogTitle.value = '编辑配置'
  formData.value = { ...row }
  dialogVisible.value = true
}

const handleSave = async () => {
  if (!formData.value.config_code || !formData.value.config_name) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    if (isEdit.value && currentId.value) {
      await updateRollingConfig(currentId.value, formData.value)
      ElMessage.success('更新成功')
    } else {
      await createRollingConfig(formData.value)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleExecute = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定执行该滚动排程配置吗？', '提示', { type: 'info' })
    await executeRollingConfig(row.id)
    ElMessage.success('执行已触发')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error('执行失败')
    }
  }
}

const handleEnable = async (row: any) => {
  try {
    await enableRollingConfig(row.id)
    ElMessage.success('启用成功')
    loadData()
  } catch {
    ElMessage.error('启用失败')
  }
}

const handleDisable = async (row: any) => {
  try {
    await disableRollingConfig(row.id)
    ElMessage.success('禁用成功')
    loadData()
  } catch {
    ElMessage.error('禁用失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该配置吗？', '提示', { type: 'warning' })
    await deleteRollingConfig(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped lang="scss">
.page-container {
  height: 100%;
}

.search-card {
  margin-bottom: 16px;
}

.toolbar-card {
  margin-bottom: 16px;

  :deep(.el-card__body) {
    padding: 12px 16px;
    display: flex;
    gap: 12px;
  }
}

.pagination {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}
</style>
