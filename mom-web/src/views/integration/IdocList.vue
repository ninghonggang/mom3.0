<template>
  <div class="idoc-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="IDOC类型">
          <el-select v-model="searchForm.idoc_type" placeholder="请选择" clearable style="width: 120px">
            <el-option label="MATMAS" value="MATMAS" />
            <el-option label="ORDERS" value="ORDERS" />
            <el-option label="DESADV" value="DESADV" />
            <el-option label="DELFOR" value="DELFOR" />
          </el-select>
        </el-form-item>
        <el-form-item label="方向">
          <el-select v-model="searchForm.direction" placeholder="请选择" clearable style="width: 100px">
            <el-option label="接收" :value="1" />
            <el-option label="发送" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable style="width: 100px">
            <el-option label="新建" :value="1" />
            <el-option label="处理中" :value="2" />
            <el-option label="成功" :value="3" />
            <el-option label="失败" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="伙伴编号">
          <el-input v-model="searchForm.partner_no" placeholder="请输入" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleReceive">
        <el-icon><Download /></el-icon>接收IDOC
      </el-button>
      <el-button type="success" @click="handleSend">
        <el-icon><Upload /></el-icon>发送IDOC
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="idoc_number" label="IDOC编号" width="200" />
        <el-table-column prop="idoc_type" label="类型" width="100" />
        <el-table-column prop="direction" label="方向" width="80">
          <template #default="{ row }">
            <el-tag :type="row.direction === 1 ? 'primary' : 'success'">
              {{ row.direction === 1 ? '接收' : '发送' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.status === 1">新建</el-tag>
            <el-tag v-else-if="row.status === 2" type="warning">处理中</el-tag>
            <el-tag v-else-if="row.status === 3" type="success">成功</el-tag>
            <el-tag v-else type="danger">失败</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="partner_no" label="伙伴编号" width="120" />
        <el-table-column prop="message_type" label="消息类型" width="100" />
        <el-table-column prop="reference_no" label="参考编号" width="150" />
        <el-table-column prop="retry_count" label="重试次数" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="160">
          <template #default="{ row }">
            {{ row.created_at ? row.created_at.slice(0, 19) : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleView(row)">查看</el-button>
            <el-button link type="warning" size="small" @click="handleRetry(row)" v-if="row.status === 4">重试</el-button>
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

    <!-- 查看对话框 -->
    <el-dialog v-model="viewVisible" title="IDOC详情" width="800px">
      <el-descriptions :column="2" border v-if="currentRecord">
        <el-descriptions-item label="IDOC编号">{{ currentRecord.idoc_number }}</el-descriptions-item>
        <el-descriptions-item label="类型">{{ currentRecord.idoc_type }}</el-descriptions-item>
        <el-descriptions-item label="方向">{{ currentRecord.direction === 1 ? '接收' : '发送' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag v-if="currentRecord.status === 1">新建</el-tag>
          <el-tag v-else-if="currentRecord.status === 2" type="warning">处理中</el-tag>
          <el-tag v-else-if="currentRecord.status === 3" type="success">成功</el-tag>
          <el-tag v-else type="danger">失败</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="伙伴类型">{{ currentRecord.partner_type }}</el-descriptions-item>
        <el-descriptions-item label="伙伴编号">{{ currentRecord.partner_no }}</el-descriptions-item>
        <el-descriptions-item label="消息类型">{{ currentRecord.message_type }}</el-descriptions-item>
        <el-descriptions-item label="参考编号">{{ currentRecord.reference_no }}</el-descriptions-item>
        <el-descriptions-item label="重试次数">{{ currentRecord.retry_count }}</el-descriptions-item>
        <el-descriptions-item label="错误信息" :span="2">{{ currentRecord.error_message || '-' }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ currentRecord.created_at }}</el-descriptions-item>
        <el-descriptions-item label="处理时间">{{ currentRecord.processed_at || '-' }}</el-descriptions-item>
      </el-descriptions>
      <el-divider />
      <h4>原始内容</h4>
      <pre class="raw-content">{{ currentRecord?.raw_content }}</pre>
    </el-dialog>

    <!-- 接收对话框 -->
    <el-dialog v-model="receiveVisible" title="接收IDOC" width="600px">
      <el-form :model="receiveForm" label-width="100px">
        <el-form-item label="IDOC类型" required>
          <el-select v-model="receiveForm.idoc_type" placeholder="请选择" style="width: 100%">
            <el-option label="MATMAS" value="MATMAS" />
            <el-option label="ORDERS" value="ORDERS" />
            <el-option label="DESADV" value="DESADV" />
            <el-option label="DELFOR" value="DELFOR" />
          </el-select>
        </el-form-item>
        <el-form-item label="伙伴类型">
          <el-input v-model="receiveForm.partner_type" placeholder="如: WE/LI/KU" />
        </el-form-item>
        <el-form-item label="伙伴编号">
          <el-input v-model="receiveForm.partner_no" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="消息类型">
          <el-input v-model="receiveForm.message_type" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="参考编号">
          <el-input v-model="receiveForm.reference_no" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="原始内容" required>
          <el-input v-model="receiveForm.raw_content" type="textarea" rows="6" placeholder="请输入IDOC JSON内容" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="receiveVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleReceiveSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 发送对话框 -->
    <el-dialog v-model="sendVisible" title="发送IDOC" width="600px">
      <el-form :model="sendForm" label-width="100px">
        <el-form-item label="IDOC类型" required>
          <el-select v-model="sendForm.idoc_type" placeholder="请选择" style="width: 100%">
            <el-option label="MATMAS" value="MATMAS" />
            <el-option label="ORDERS" value="ORDERS" />
            <el-option label="DESADV" value="DESADV" />
            <el-option label="DELFOR" value="DELFOR" />
          </el-select>
        </el-form-item>
        <el-form-item label="目标类型">
          <el-input v-model="sendForm.target_type" placeholder="如: WE/LI/KU" />
        </el-form-item>
        <el-form-item label="目标编号">
          <el-input v-model="sendForm.target_no" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="消息类型">
          <el-input v-model="sendForm.message_type" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据内容" required>
          <el-input v-model="sendForm.dataJson" type="textarea" rows="6" placeholder="请输入JSON数据" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="sendVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSendSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getIdocList, getIdoc, receiveIdoc, sendIdoc, retryIdoc } from '@/api/idoc'

const loading = ref(false)
const saveLoading = ref(false)
const viewVisible = ref(false)
const receiveVisible = ref(false)
const sendVisible = ref(false)
const tableData = ref<any[]>([])
const currentRecord = ref<any>(null)

const searchForm = reactive({
  idoc_type: '',
  direction: null as number | null,
  status: null as number | null,
  partner_no: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const receiveForm = reactive({
  idoc_type: '',
  partner_type: '',
  partner_no: '',
  message_type: '',
  reference_no: '',
  raw_content: ''
})

const sendForm = reactive({
  idoc_type: '',
  target_type: '',
  target_no: '',
  message_type: '',
  dataJson: ''
})

const loadData = async () => {
  loading.value = true
  try {
    const res = await getIdocList({
      page: pagination.page,
      page_size: pagination.pageSize,
      idoc_type: searchForm.idoc_type,
      direction: searchForm.direction,
      status: searchForm.status,
      partner_no: searchForm.partner_no
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.idoc_type = ''
  searchForm.direction = null
  searchForm.status = null
  searchForm.partner_no = ''
  handleSearch()
}

const handleView = async (row: any) => {
  try {
    const res = await getIdoc(row.id)
    currentRecord.value = res.data
    viewVisible.value = true
  } catch (e) {
    console.error(e)
  }
}

const handleReceive = () => {
  receiveForm.idoc_type = ''
  receiveForm.partner_type = ''
  receiveForm.partner_no = ''
  receiveForm.message_type = ''
  receiveForm.reference_no = ''
  receiveForm.raw_content = ''
  receiveVisible.value = true
}

const handleReceiveSubmit = async () => {
  if (!receiveForm.idoc_type || !receiveForm.raw_content) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    await receiveIdoc(receiveForm)
    ElMessage.success('IDOC接收成功')
    receiveVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleSend = () => {
  sendForm.idoc_type = ''
  sendForm.target_type = ''
  sendForm.target_no = ''
  sendForm.message_type = ''
  sendForm.dataJson = ''
  sendVisible.value = true
}

const handleSendSubmit = async () => {
  if (!sendForm.idoc_type || !sendForm.dataJson) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    let data
    try {
      data = JSON.parse(sendForm.dataJson)
    } catch {
      ElMessage.error('JSON格式错误')
      return
    }
    await sendIdoc({
      idoc_type: sendForm.idoc_type,
      target_type: sendForm.target_type,
      target_no: sendForm.target_no,
      message_type: sendForm.message_type,
      data
    })
    ElMessage.success('IDOC发送成功')
    sendVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleRetry = async (row: any) => {
  try {
    await retryIdoc(row.id)
    ElMessage.success('重试成功')
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '重试失败')
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.idoc-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .raw-content {
    background: #f5f5f5;
    padding: 12px;
    border-radius: 4px;
    max-height: 300px;
    overflow: auto;
    font-size: 12px;
    white-space: pre-wrap;
    word-break: break-all;
  }
}
</style>