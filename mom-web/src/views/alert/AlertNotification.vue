<template>
  <div class="alert-notification">
    <!-- 统计卡片 -->
    <el-row :gutter="16" class="stats-row">
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-value">{{ statistics.total_count || 0 }}</div>
            <div class="stat-label">总通知数</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-value stat-sent">{{ statistics.sent_count || 0 }}</div>
            <div class="stat-label">已发送</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-value stat-pending">{{ statistics.pending_count || 0 }}</div>
            <div class="stat-label">待发送</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-value stat-failed">{{ statistics.failed_count || 0 }}</div>
            <div class="stat-label">发送失败</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 搜索卡片 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="渠道类型">
          <el-select v-model="searchForm.channel_type" placeholder="请选择" clearable>
            <el-option label="站内信" value="IN_SITE" />
            <el-option label="飞书" value="FEISHU" />
            <el-option label="企业微信" value="WECOM" />
            <el-option label="邮件" value="EMAIL" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待发送" value="PENDING" />
            <el-option label="已发送" value="SENT" />
            <el-option label="已失败" value="FAILED" />
            <el-option label="已读" value="READ" />
          </el-select>
        </el-form-item>
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="标题/内容/编号" clearable />
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

    <!-- 工具栏 -->
    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('alert:notification:send')" @click="handleSend">
        <el-icon><Plus /></el-icon>发送通知
      </el-button>
      <el-button v-if="hasPermission('alert:notification:refresh')" @click="loadStatistics">
        <el-icon><Refresh /></el-icon>刷新统计
      </el-button>
    </el-card>

    <!-- 数据表格 -->
    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="notification_no" label="通知编号" width="180" />
        <el-table-column prop="title" label="通知标题" min-width="200" show-overflow-tooltip />
        <el-table-column prop="channel_type" label="渠道" width="100">
          <template #default="{ row }">
            <el-tag>{{ getChannelTypeText(row.channel_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="receiver_name" label="接收人" width="120" />
        <el-table-column prop="receiver_value" label="接收值" width="150" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sent_time" label="发送时间" width="160">
          <template #default="{ row }">
            {{ row.sent_time ? formatDateTime(row.sent_time) : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="retry_count" label="重试次数" width="90" />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleView(row)">详情</el-button>
            <el-button link type="danger" size="small" v-if="row.status === 'FAILED' && hasPermission('alert:notification:retry')" @click="handleRetry(row)">重试</el-button>
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

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="通知详情" width="600px">
      <el-descriptions :column="2" border v-if="currentNotification">
        <el-descriptions-item label="通知编号" :span="2">{{ currentNotification.notification_no }}</el-descriptions-item>
        <el-descriptions-item label="通知标题" :span="2">{{ currentNotification.title }}</el-descriptions-item>
        <el-descriptions-item label="通知内容" :span="2">
          <div class="content-text">{{ currentNotification.content }}</div>
        </el-descriptions-item>
        <el-descriptions-item label="渠道类型">
          <el-tag>{{ getChannelTypeText(currentNotification.channel_type) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(currentNotification.status)">{{ getStatusText(currentNotification.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="接收人">{{ currentNotification.receiver_name || '-' }}</el-descriptions-item>
        <el-descriptions-item label="接收值">{{ currentNotification.receiver_value || '-' }}</el-descriptions-item>
        <el-descriptions-item label="关联告警">{{ currentNotification.alert_no || '-' }}</el-descriptions-item>
        <el-descriptions-item label="发送时间">
          {{ currentNotification.sent_time ? formatDateTime(currentNotification.sent_time) : '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="阅读时间">
          {{ currentNotification.read_time ? formatDateTime(currentNotification.read_time) : '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="重试次数">{{ currentNotification.retry_count || 0 }}</el-descriptions-item>
        <el-descriptions-item label="错误信息" :span="2" v-if="currentNotification.error_msg">
          <div class="error-text">{{ currentNotification.error_msg }}</div>
        </el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>

    <!-- 发送通知对话框 -->
    <el-dialog v-model="sendDialogVisible" title="发送通知" width="600px">
      <el-form ref="sendFormRef" :model="sendForm" :rules="sendRules" label-width="100px">
        <el-form-item label="渠道类型" prop="channel_type">
          <el-select v-model="sendForm.channel_type" placeholder="请选择渠道类型">
            <el-option label="站内信" value="IN_SITE" />
            <el-option label="飞书" value="FEISHU" />
            <el-option label="企业微信" value="WECOM" />
            <el-option label="邮件" value="EMAIL" />
          </el-select>
        </el-form-item>
        <el-form-item label="通知标题" prop="title">
          <el-input v-model="sendForm.title" placeholder="请输入通知标题" />
        </el-form-item>
        <el-form-item label="通知内容" prop="content">
          <el-input v-model="sendForm.content" type="textarea" :rows="4" placeholder="请输入通知内容" />
        </el-form-item>
        <el-form-item label="接收人类型" prop="receiver_type">
          <el-select v-model="sendForm.receiver_type" placeholder="请选择">
            <el-option label="用户" value="USER" />
            <el-option label="角色" value="ROLE" />
            <el-option label="部门" value="DEPT" />
            <el-option label="全部" value="ALL" />
          </el-select>
        </el-form-item>
        <el-form-item label="接收人ID">
          <el-input-number v-model="sendForm.receiver_id" :min="0" placeholder="接收人ID" />
        </el-form-item>
        <el-form-item label="接收人姓名">
          <el-input v-model="sendForm.receiver_name" placeholder="接收人姓名" />
        </el-form-item>
        <el-form-item label="接收值">
          <el-input v-model="sendForm.receiver_value" placeholder="邮箱地址/手机号/飞书/企微账号" />
        </el-form-item>
        <el-form-item label="关联告警">
          <el-input v-model="sendForm.alert_no" placeholder="关联告警编号（可选）" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="sendDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="sendLoading" @click="handleSendSubmit">发送</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, FormInstance, FormRules } from 'element-plus'
import { useAuthStore } from '@/stores/auth'
import {
  getNotificationList,
  getNotification,
  getNotificationStatistics,
  markNotificationAsRead,
  sendNotification
} from '@/api/alert'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const currentNotification = ref<any>(null)
const detailVisible = ref(false)
const sendDialogVisible = ref(false)
const sendLoading = ref(false)
const sendFormRef = ref<FormInstance>()
const dateRange = ref<[string, string] | null>(null)

const statistics = reactive({
  total_count: 0,
  sent_count: 0,
  pending_count: 0,
  failed_count: 0
})

const searchForm = reactive({
  channel_type: '',
  status: '',
  keyword: '',
  start_date: '',
  end_date: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const sendForm = reactive<any>({
  channel_type: 'IN_SITE',
  title: '',
  content: '',
  receiver_type: 'USER',
  receiver_id: null,
  receiver_name: '',
  receiver_value: '',
  alert_no: ''
})

const sendRules: FormRules = {
  channel_type: [{ required: true, message: '请选择渠道类型', trigger: 'change' }],
  title: [{ required: true, message: '请输入通知标题', trigger: 'blur' }],
  content: [{ required: true, message: '请输入通知内容', trigger: 'blur' }],
  receiver_type: [{ required: true, message: '请选择接收人类型', trigger: 'change' }]
}

const getChannelTypeText = (type: string) => {
  const map: Record<string, string> = {
    IN_SITE: '站内信',
    FEISHU: '飞书',
    WECOM: '企业微信',
    EMAIL: '邮件'
  }
  return map[type] || type
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    PENDING: '待发送',
    SENT: '已发送',
    FAILED: '已失败',
    READ: '已读'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    PENDING: 'warning',
    SENT: 'success',
    FAILED: 'danger',
    READ: 'info'
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
    const res = await getNotificationList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadStatistics = async () => {
  try {
    const res = await getNotificationStatistics()
    Object.assign(statistics, res.data || {})
  } catch (error) {
    // ignore
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }

const handleReset = () => {
  searchForm.channel_type = ''
  searchForm.status = ''
  searchForm.keyword = ''
  dateRange.value = null
  searchForm.start_date = ''
  searchForm.end_date = ''
  handleSearch()
}

const handleView = async (row: any) => {
  try {
    const res = await getNotification(row.id)
    currentNotification.value = res.data
    // 标记已读
    if (row.status === 'SENT') {
      markNotificationAsRead(row.id).catch(() => {})
    }
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取详情失败')
  }
}

const handleSend = () => {
  Object.assign(sendForm, {
    channel_type: 'IN_SITE',
    title: '',
    content: '',
    receiver_type: 'USER',
    receiver_id: null,
    receiver_name: '',
    receiver_value: '',
    alert_no: ''
  })
  sendDialogVisible.value = true
}

const handleSendSubmit = async () => {
  if (!sendFormRef.value) return
  await sendFormRef.value.validate()
  sendLoading.value = true
  try {
    await sendNotification(sendForm)
    ElMessage.success('发送成功')
    sendDialogVisible.value = false
    loadData()
    loadStatistics()
  } finally {
    sendLoading.value = false
  }
}

const handleRetry = async (row: any) => {
  try {
    await sendNotification({
      channel_type: row.channel_type,
      title: row.title,
      content: row.content,
      receiver_type: row.receiver_type,
      receiver_id: row.receiver_id,
      receiver_name: row.receiver_name,
      receiver_value: row.receiver_value,
      alert_no: row.alert_no
    })
    ElMessage.success('重试成功')
    loadData()
    loadStatistics()
  } catch (error) {
    ElMessage.error('重试失败')
  }
}

onMounted(() => {
  loadData()
  loadStatistics()
})
</script>

<style scoped lang="scss">
.alert-notification {
  .stats-row {
    margin-bottom: 16px;
  }
  .stat-card {
    .stat-content {
      text-align: center;
      padding: 8px 0;
    }
    .stat-value {
      font-size: 28px;
      font-weight: bold;
      color: #409eff;
      &.stat-sent { color: #67c23a; }
      &.stat-pending { color: #e6a23c; }
      &.stat-failed { color: #f56c6c; }
    }
    .stat-label {
      font-size: 14px;
      color: #909399;
      margin-top: 4px;
    }
  }
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
  .content-text {
    max-height: 150px;
    overflow-y: auto;
    white-space: pre-wrap;
  }
  .error-text {
    color: #f56c6c;
    white-space: pre-wrap;
  }
}
</style>
