<template>
  <div class="andon-call">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="呼叫编号">
          <el-input v-model="searchForm.call_no" placeholder="请输入呼叫编号" clearable />
        </el-form-item>
        <el-form-item label="呼叫类型">
          <el-select v-model="searchForm.andon_type" placeholder="请选择" clearable>
            <el-option label="设备故障" value="EQUIPMENT" />
            <el-option label="物料缺料" value="MATERIAL" />
            <el-option label="品质异常" value="QUALITY" />
            <el-option label="工艺异常" value="TECHNICAL" />
            <el-option label="工装夹具" value="TOOLING" />
            <el-option label="安全警示" value="SAFETY" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待响应" value="CALLING" />
            <el-option label="响应中" value="RESPONDED" />
            <el-option label="处理中" value="HANDLING" />
            <el-option label="已解决" value="RESOLVED" />
            <el-option label="已关闭" value="CLOSED" />
          </el-select>
        </el-form-item>
        <el-form-item label="日期范围">
          <el-date-picker
            v-model="dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
            @change="handleDateChange"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="handleCreate">发起呼叫</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" stripe>
        <el-table-column prop="call_no" label="呼叫编号" width="150" />
        <el-table-column prop="andon_type" label="呼叫类型" width="110">
          <template #default="{ row }">
            <el-tag :type="getCallTypeTag(row.andon_type)">{{ getCallTypeText(row.andon_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="call_level" label="等级" width="70">
          <template #default="{ row }">
            <el-tag type="danger" v-if="row.call_level > 1">L{{ row.call_level }}</el-tag>
            <span v-else>L1</span>
          </template>
        </el-table-column>
        <el-table-column prop="workshop_name" label="车间" width="100" />
        <el-table-column prop="production_line_name" label="产线" width="120" />
        <el-table-column prop="workstation_name" label="工位" width="100" />
        <el-table-column prop="call_by" label="呼叫人" width="80" />
        <el-table-column prop="call_time" label="呼叫时间" width="160" />
        <el-table-column prop="response_by" label="响应人" width="80" />
        <el-table-column prop="handle_by" label="处理人" width="80" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="is_escalated" label="升级" width="70">
          <template #default="{ row }">
            <el-tag type="warning" v-if="row.is_escalated === 1">已升级</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="220" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleDetail(row)">详情</el-button>
            <el-button link type="warning" size="small" @click="handleEscalate(row)" v-if="row.status !== 'RESOLVED' && row.status !== 'CLOSED'">升级</el-button>
            <el-button link type="success" size="small" @click="handleResponse(row)" v-if="row.status === 'CALLING'">响应</el-button>
            <el-button link type="primary" size="small" @click="handleResolve(row)" v-if="row.status === 'RESPONDED' || row.status === 'HANDLING'">解决</el-button>
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

    <!-- 详情弹窗 -->
    <el-dialog v-model="detailVisible" title="呼叫详情" width="700px" destroy-on-close>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="呼叫编号">{{ currentRow.call_no }}</el-descriptions-item>
        <el-descriptions-item label="呼叫类型">
          <el-tag :type="getCallTypeTag(currentRow.andon_type)">{{ getCallTypeText(currentRow.andon_type) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="车间">{{ currentRow.workshop_name }}</el-descriptions-item>
        <el-descriptions-item label="产线">{{ currentRow.production_line_name }}</el-descriptions-item>
        <el-descriptions-item label="工位">{{ currentRow.workstation_name }}</el-descriptions-item>
        <el-descriptions-item label="当前等级">
          <el-tag type="danger" v-if="currentRow.call_level > 1">L{{ currentRow.call_level }}</el-tag>
          <span v-else>L1</span>
        </el-descriptions-item>
        <el-descriptions-item label="呼叫人">{{ currentRow.call_by }}</el-descriptions-item>
        <el-descriptions-item label="呼叫时间">{{ currentRow.call_time }}</el-descriptions-item>
        <el-descriptions-item label="响应人">{{ currentRow.response_by || '-' }}</el-descriptions-item>
        <el-descriptions-item label="响应时间">{{ currentRow.response_time || '-' }}</el-descriptions-item>
        <el-descriptions-item label="处理人">{{ currentRow.handle_by || '-' }}</el-descriptions-item>
        <el-descriptions-item label="处理时间">{{ currentRow.handle_time || '-' }}</el-descriptions-item>
        <el-descriptions-item label="处理结果">{{ currentRow.handle_result || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(currentRow.status)">{{ getStatusText(currentRow.status) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="升级次数">{{ currentRow.escalation_count || 0 }}</el-descriptions-item>
        <el-descriptions-item label="是否升级">{{ currentRow.is_escalated === 1 ? '是' : '否' }}</el-descriptions-item>
        <el-descriptions-item label="描述" :span="2">{{ currentRow.description || '-' }}</el-descriptions-item>
        <el-descriptions-item label="处理备注" :span="2">{{ currentRow.handle_remark || '-' }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>

    <!-- 发起呼叫弹窗 -->
    <el-dialog v-model="createVisible" title="发起安东呼叫" width="600px" destroy-on-close>
      <el-form ref="createFormRef" :model="createForm" :rules="createRules" label-width="100px">
        <el-form-item label="呼叫类型" prop="andon_type">
          <el-select v-model="createForm.andon_type" placeholder="请选择呼叫类型">
            <el-option label="设备故障" value="EQUIPMENT" />
            <el-option label="物料缺料" value="MATERIAL" />
            <el-option label="品质异常" value="QUALITY" />
            <el-option label="工艺异常" value="TECHNICAL" />
            <el-option label="工装夹具" value="TOOLING" />
            <el-option label="安全警示" value="SAFETY" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-form-item label="车间" prop="workshop_id">
          <el-input-number v-model="createForm.workshop_id" :min="1" placeholder="车间ID" />
        </el-form-item>
        <el-form-item label="工位" prop="workstation_id">
          <el-input-number v-model="createForm.workstation_id" :min="1" placeholder="工位ID" />
        </el-form-item>
        <el-form-item label="产线">
          <el-input-number v-model="createForm.production_line_id" :min="1" placeholder="产线ID" />
        </el-form-item>
        <el-form-item label="优先级">
          <el-input-number v-model="createForm.priority" :min="1" :max="10" :value="5" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="createForm.description" type="textarea" :rows="3" placeholder="请输入问题描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleCreateSubmit">确认呼叫</el-button>
      </template>
    </el-dialog>

    <!-- 响应弹窗 -->
    <el-dialog v-model="responseVisible" title="响应呼叫" width="500px" destroy-on-close>
      <el-form :model="responseForm" label-width="100px">
        <el-form-item label="响应人">
          <el-input v-model="responseForm.response_by" placeholder="请输入响应人姓名" />
        </el-form-item>
        <el-form-item label="响应备注">
          <el-input v-model="responseForm.response_remark" type="textarea" :rows="2" placeholder="处理说明" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="responseVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleResponseSubmit">确认响应</el-button>
      </template>
    </el-dialog>

    <!-- 解决弹窗 -->
    <el-dialog v-model="resolveVisible" title="解决呼叫" width="500px" destroy-on-close>
      <el-form ref="resolveFormRef" :model="resolveForm" :rules="resolveRules" label-width="100px">
        <el-form-item label="处理人">
          <el-input v-model="resolveForm.handle_by" placeholder="请输入处理人姓名" />
        </el-form-item>
        <el-form-item label="处理结果" prop="handle_result">
          <el-radio-group v-model="resolveForm.handle_result">
            <el-radio value="RESOLVED">已解决</el-radio>
            <el-radio value="CARRY_OVER">遗留</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="处理备注">
          <el-input v-model="resolveForm.handle_remark" type="textarea" :rows="3" placeholder="处理说明" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="resolveVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleResolveSubmit">确认解决</el-button>
      </template>
    </el-dialog>

    <!-- 升级弹窗 -->
    <el-dialog v-model="escalateVisible" title="手动升级" width="500px" destroy-on-close>
      <el-form :model="escalateForm" label-width="110px">
        <el-form-item label="升级到等级">
          <el-select v-model="escalateForm.escalate_to_level" placeholder="请选择">
            <el-option label="L2" :value="2" />
            <el-option label="L3" :value="3" />
            <el-option label="L4" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="升级原因">
          <el-input v-model="escalateForm.escalation_reason" type="textarea" :rows="2" placeholder="请输入升级原因" />
        </el-form-item>
        <el-form-item label="触发人">
          <el-input v-model="escalateForm.trigger_user" placeholder="请输入触发人" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="escalateVisible = false">取消</el-button>
        <el-button type="warning" :loading="submitLoading" @click="handleEscalateSubmit">确认升级</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, FormInstance, FormRules } from 'element-plus'
import {
  getAndonCallList,
  createAndonCall,
  responseAndonCall,
  resolveAndonCall,
  escalateAndonCall
} from '@/api/trace'

const loading = ref(false)
const tableData = ref<any[]>([])
const currentRow = ref<any>({})
const detailVisible = ref(false)
const createVisible = ref(false)
const responseVisible = ref(false)
const resolveVisible = ref(false)
const escalateVisible = ref(false)
const submitLoading = ref(false)
const createFormRef = ref<FormInstance>()
const resolveFormRef = ref<FormInstance>()

const dateRange = ref<[string, string] | null>(null)
const searchForm = reactive({ call_no: '', andon_type: '', status: '', start_date: '', end_date: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const createForm = reactive({
  andon_type: '', workshop_id: 1, workstation_id: 1, production_line_id: 1, priority: 5, description: ''
})
const createRules: FormRules = {
  andon_type: [{ required: true, message: '请选择呼叫类型', trigger: 'change' }]
}
const responseForm = reactive({ response_by: '', response_remark: '' })
const resolveForm = reactive({ handle_by: '', handle_result: 'RESOLVED', handle_remark: '' })
const resolveRules: FormRules = {
  handle_result: [{ required: true, message: '请选择处理结果', trigger: 'change' }]
}
const escalateForm = reactive({ escalate_to_level: 2, escalation_reason: '', trigger_user: '' })

const getCallTypeText = (type: string) => {
  const map: Record<string, string> = {
    EQUIPMENT: '设备故障', MATERIAL: '物料缺料', QUALITY: '品质异常',
    TECHNICAL: '工艺异常', TOOLING: '工装夹具', SAFETY: '安全警示', OTHER: '其他'
  }
  return map[type] || type
}
const getCallTypeTag = (type: string) => {
  const map: Record<string, string> = {
    EQUIPMENT: 'danger', MATERIAL: 'warning', QUALITY: 'danger',
    TECHNICAL: 'info', TOOLING: 'info', SAFETY: 'warning', OTHER: 'info'
  }
  return map[type] || 'info'
}
const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    CALLING: '待响应', RESPONDED: '响应中', HANDLING: '处理中', RESOLVED: '已解决', CLOSED: '已关闭'
  }
  return map[status] || status
}
const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    CALLING: 'danger', RESPONDED: 'warning', HANDLING: 'warning', RESOLVED: 'success', CLOSED: 'info'
  }
  return map[status] || 'info'
}

const handleDateChange = (val: [string, string] | null) => {
  if (val) { searchForm.start_date = val[0]; searchForm.end_date = val[1] }
  else { searchForm.start_date = ''; searchForm.end_date = '' }
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => { if (!params[k]) delete params[k] })
    const res = await getAndonCallList(params)
    tableData.value = res.data?.list || res.data || []
    pagination.total = res.data?.total || 0
  } catch (e: any) {
    ElMessage.error(e.message || '加载失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  Object.assign(searchForm, { call_no: '', andon_type: '', status: '', start_date: '', end_date: '' })
  dateRange.value = null
  handleSearch()
}

const handleDetail = (row: any) => { currentRow.value = row; detailVisible.value = true }

const handleCreate = () => {
  Object.assign(createForm, { andon_type: '', workshop_id: 1, workstation_id: 1, production_line_id: 1, priority: 5, description: '' })
  createVisible.value = true
}
const handleCreateSubmit = async () => {
  const valid = await createFormRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    await createAndonCall(createForm)
    ElMessage.success('呼叫已发起')
    createVisible.value = false
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '发起失败')
  } finally {
    submitLoading.value = false
  }
}

const handleResponse = (row: any) => {
  currentRow.value = row
  Object.assign(responseForm, { response_by: '', response_remark: '' })
  responseVisible.value = true
}
const handleResponseSubmit = async () => {
  submitLoading.value = true
  try {
    await responseAndonCall(currentRow.value.id, responseForm)
    ElMessage.success('已响应')
    responseVisible.value = false
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '响应失败')
  } finally {
    submitLoading.value = false
  }
}

const handleResolve = (row: any) => {
  currentRow.value = row
  Object.assign(resolveForm, { handle_by: '', handle_result: 'RESOLVED', handle_remark: '' })
  resolveVisible.value = true
}
const handleResolveSubmit = async () => {
  const valid = await resolveFormRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    await resolveAndonCall(currentRow.value.id, resolveForm)
    ElMessage.success('已解决')
    resolveVisible.value = false
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '解决失败')
  } finally {
    submitLoading.value = false
  }
}

const handleEscalate = (row: any) => {
  currentRow.value = row
  Object.assign(escalateForm, { escalate_to_level: (row.call_level || 1) + 1, escalation_reason: '', trigger_user: '' })
  escalateVisible.value = true
}
const handleEscalateSubmit = async () => {
  submitLoading.value = true
  try {
    await escalateAndonCall(currentRow.value.id, escalateForm)
    ElMessage.success('已升级')
    escalateVisible.value = false
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '升级失败')
  } finally {
    submitLoading.value = false
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.andon-call {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
