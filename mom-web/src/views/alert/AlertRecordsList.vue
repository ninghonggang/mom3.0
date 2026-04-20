<template>
  <div class="alert-records-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="告警编号">
          <el-input v-model="searchForm.alert_no" placeholder="请输入告警编号" clearable />
        </el-form-item>
        <el-form-item label="告警类型">
          <el-select v-model="searchForm.alert_type" placeholder="请选择" clearable>
            <el-option label="设备告警" value="EQUIPMENT" />
            <el-option label="生产告警" value="PRODUCTION" />
            <el-option label="质量告警" value="QUALITY" />
            <el-option label="库存告警" value="INVENTORY" />
            <el-option label="系统告警" value="SYSTEM" />
          </el-select>
        </el-form-item>
        <el-form-item label="严重级别">
          <el-select v-model="searchForm.severity_level" placeholder="请选择" clearable>
            <el-option label="Critical" value="CRITICAL" />
            <el-option label="High" value="HIGH" />
            <el-option label="Medium" value="MEDIUM" />
            <el-option label="Low" value="LOW" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="触发" value="TRIGGERED" />
            <el-option label="已确认" value="ACKNOWLEDGED" />
            <el-option label="已解决" value="RESOLVED" />
            <el-option label="已关闭" value="CLOSED" />
          </el-select>
        </el-form-item>
        <el-form-item label="来源模块">
          <el-input v-model="searchForm.source_module" placeholder="请输入来源模块" clearable />
        </el-form-item>
        <el-form-item label="时间范围">
          <el-date-picker
            v-model="dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="alert_no" label="告警编号" width="160" />
        <el-table-column prop="title" label="告警标题" min-width="200" />
        <el-table-column prop="alert_type" label="类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getAlertTypeText(row.alert_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="severity_level" label="级别" width="90">
          <template #default="{ row }">
            <el-tag :type="getSeverityType(row.severity_level)">{{ row.severity_level }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="source_module" label="来源模块" width="110" />
        <el-table-column prop="source_no" label="来源单号" width="140" />
        <el-table-column prop="trigger_time" label="触发时间" width="160">
          <template #default="{ row }">
            {{ formatDateTime(row.trigger_time) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="acknowledged_by_name" label="确认人" width="90" />
        <el-table-column prop="resolved_by_name" label="解决人" width="90" />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('alert:records:ack') && row.status === 'TRIGGERED'" @click="handleAcknowledge(row)">确认</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('alert:records:resolve') && row.status === 'ACKNOWLEDGED'" @click="handleResolve(row)">解决</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('alert:records:close') && row.status === 'RESOLVED'" @click="handleClose(row)">关闭</el-button>
            <el-button link size="small" @click="handleView(row)">详情</el-button>
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

    <!-- 确认对话框 -->
    <el-dialog v-model="ackDialogVisible" title="确认告警" width="500px">
      <el-form ref="ackFormRef" :model="ackForm" label-width="80px">
        <el-form-item label="确认人">
          <el-input v-model="ackForm.user_name" disabled />
        </el-form-item>
        <el-form-item label="确认备注">
          <el-input v-model="ackForm.remark" type="textarea" :rows="3" placeholder="请输入确认备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="ackDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleAckSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 解决对话框 -->
    <el-dialog v-model="resolveDialogVisible" title="解决告警" width="500px">
      <el-form ref="resolveFormRef" :model="resolveForm" label-width="80px">
        <el-form-item label="解决人">
          <el-input v-model="resolveForm.user_name" disabled />
        </el-form-item>
        <el-form-item label="解决备注">
          <el-input v-model="resolveForm.remark" type="textarea" :rows="3" placeholder="请输入解决备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="resolveDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleResolveSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="告警详情" width="700px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="告警编号">{{ currentRecord?.alert_no }}</el-descriptions-item>
        <el-descriptions-item label="告警类型">{{ getAlertTypeText(currentRecord?.alert_type || '') }}</el-descriptions-item>
        <el-descriptions-item label="严重级别">
          <el-tag :type="getSeverityType(currentRecord?.severity_level || '')">{{ currentRecord?.severity_level }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(currentRecord?.status || '')">{{ getStatusText(currentRecord?.status || '') }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="告警标题" :span="2">{{ currentRecord?.title }}</el-descriptions-item>
        <el-descriptions-item label="告警内容" :span="2">{{ currentRecord?.content }}</el-descriptions-item>
        <el-descriptions-item label="来源模块">{{ currentRecord?.source_module }}</el-descriptions-item>
        <el-descriptions-item label="来源单号">{{ currentRecord?.source_no }}</el-descriptions-item>
        <el-descriptions-item label="触发时间">{{ formatDateTime(currentRecord?.trigger_time) }}</el-descriptions-item>
        <el-descriptions-item label="确认人">{{ currentRecord?.acknowledged_by_name || '-' }}</el-descriptions-item>
        <el-descriptions-item label="确认时间">{{ currentRecord?.acknowledged_time ? formatDateTime(currentRecord?.acknowledged_time) : '-' }}</el-descriptions-item>
        <el-descriptions-item label="确认备注" :span="2">{{ currentRecord?.acknowledged_remark || '-' }}</el-descriptions-item>
        <el-descriptions-item label="解决人">{{ currentRecord?.resolved_by_name || '-' }}</el-descriptions-item>
        <el-descriptions-item label="解决时间">{{ currentRecord?.resolved_time ? formatDateTime(currentRecord?.resolved_time) : '-' }}</el-descriptions-item>
        <el-descriptions-item label="解决备注" :span="2">{{ currentRecord?.resolution_remark || '-' }}</el-descriptions-item>
        <el-descriptions-item label="关闭时间">{{ currentRecord?.closed_time ? formatDateTime(currentRecord?.closed_time) : '-' }}</el-descriptions-item>
        <el-descriptions-item label="升级次数">{{ currentRecord?.escalation_count || 0 }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance } from 'element-plus'
import { useAuthStore } from '@/stores/auth'
import {
  getAlertRecordsList,
  acknowledgeAlertRecord,
  resolveAlertRecord,
  closeAlertRecord
} from '@/api/alert'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const currentRecord = ref<any>(null)
const dateRange = ref<[string, string] | null>(null)

const searchForm = reactive({
  alert_no: '',
  alert_type: '',
  severity_level: '',
  status: '',
  source_module: '',
  start_date: '',
  end_date: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

// 确认
const ackDialogVisible = ref(false)
const ackFormRef = ref<FormInstance>()
const ackForm = reactive({
  user_id: 0,
  user_name: '',
  remark: ''
})
const currentAckId = ref<number>(0)

// 解决
const resolveDialogVisible = ref(false)
const resolveFormRef = ref<FormInstance>()
const resolveForm = reactive({
  user_id: 0,
  user_name: '',
  remark: ''
})
const currentResolveId = ref<number>(0)

// 详情
const detailDialogVisible = ref(false)

const getAlertTypeText = (type: string) => {
  const map: Record<string, string> = {
    EQUIPMENT: '设备告警',
    PRODUCTION: '生产告警',
    QUALITY: '质量告警',
    INVENTORY: '库存告警',
    SYSTEM: '系统告警'
  }
  return map[type] || type
}

const getSeverityType = (level: string) => {
  const map: Record<string, string> = {
    CRITICAL: 'danger',
    HIGH: 'warning',
    MEDIUM: 'info',
    LOW: 'success'
  }
  return map[level] || 'info'
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    TRIGGERED: '触发',
    ACKNOWLEDGED: '已确认',
    RESOLVED: '已解决',
    CLOSED: '已关闭'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    TRIGGERED: 'danger',
    ACKNOWLEDGED: 'warning',
    RESOLVED: 'success',
    CLOSED: 'info'
  }
  return map[status] || 'info'
}

const formatDateTime = (datetime: string) => {
  if (!datetime) return '-'
  const d = new Date(datetime)
  return d.toLocaleString()
}

const loadData = async () => {
  loading.value = true
  try {
    if (dateRange.value) {
      searchForm.start_date = dateRange.value[0]
      searchForm.end_date = dateRange.value[1]
    } else {
      searchForm.start_date = ''
      searchForm.end_date = ''
    }
    const params = {
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getAlertRecordsList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }

const handleReset = () => {
  searchForm.alert_no = ''
  searchForm.alert_type = ''
  searchForm.severity_level = ''
  searchForm.status = ''
  searchForm.source_module = ''
  dateRange.value = null
  searchForm.start_date = ''
  searchForm.end_date = ''
  handleSearch()
}

const handleAcknowledge = (row: any) => {
  currentAckId.value = row.id
  ackForm.user_id = 0
  ackForm.user_name = ''
  ackForm.remark = ''
  ackDialogVisible.value = true
}

const handleAckSubmit = async () => {
  try {
    await acknowledgeAlertRecord(currentAckId.value, ackForm)
    ElMessage.success('确认成功')
    ackDialogVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('确认失败')
  }
}

const handleResolve = (row: any) => {
  currentResolveId.value = row.id
  resolveForm.user_id = 0
  resolveForm.user_name = ''
  resolveForm.remark = ''
  resolveDialogVisible.value = true
}

const handleResolveSubmit = async () => {
  try {
    await resolveAlertRecord(currentResolveId.value, resolveForm)
    ElMessage.success('解决成功')
    resolveDialogVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error('解决失败')
  }
}

const handleClose = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定关闭该告警吗？', '提示', { type: 'warning' })
    await closeAlertRecord(row.id)
    ElMessage.success('关闭成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleView = (row: any) => {
  currentRecord.value = row
  detailDialogVisible.value = true
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.alert-records-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
