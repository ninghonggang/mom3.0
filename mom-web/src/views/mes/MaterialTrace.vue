<template>
  <div class="material-trace-container">
    <el-card class="trace-header">
      <template #header>
        <div class="card-header">
          <span>物料追溯查询</span>
          <el-button type="primary" @click="queryTrace">查询</el-button>
        </div>
      </template>
      <el-form :inline="true" :model="queryForm">
        <el-form-item label="物料编码">
          <el-input v-model="queryForm.materialCode" placeholder="请输入物料编码" clearable />
        </el-form-item>
        <el-form-item label="批次号">
          <el-input v-model="queryForm.batchNo" placeholder="请输入批次号" clearable />
        </el-form-item>
        <el-form-item label="时间范围">
          <el-date-picker
            v-model="queryForm.dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="trace-content" v-loading="loading">
      <el-tabs v-model="activeTab" type="border-card">
        <el-tab-pane label="追溯链条" name="chain">
          <el-timeline v-if="traceData.length > 0">
            <el-timeline-item
              v-for="(item, index) in traceData"
              :key="index"
              :timestamp="item.createTime"
              :color="getTimelineColor(item.eventType)"
              placement="top"
            >
              <el-card class="trace-card">
                <template #header>
                  <div class="trace-card-header">
                    <el-tag :type="getEventTypeTag(item.eventType)" size="small">
                      {{ getEventTypeName(item.eventType) }}
                    </el-tag>
                    <span class="trace-no">{{ item.sourceNo }}</span>
                  </div>
                </template>
                <div class="trace-card-body">
                  <el-descriptions :column="2" border size="small">
                    <el-descriptions-item label="模块">
                      {{ item.sourceModule }}
                    </el-descriptions-item>
                    <el-descriptions-item label="事件类型">
                      {{ item.eventType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="关联单号" v-if="item.refNo">
                      {{ item.refNo }}
                    </el-descriptions-item>
                    <el-descriptions-item label="操作人" v-if="item.operator">
                      {{ item.operator }}
                    </el-descriptions-item>
                    <el-descriptions-item label="状态" :span="2">
                      <el-tag :type="getStatusTag(item.status)" size="small">
                        {{ item.status }}
                      </el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item label="详情" :span="2" v-if="item.detail">
                      <div class="trace-detail">{{ item.detail }}</div>
                    </el-descriptions-item>
                  </el-descriptions>
                </div>
              </el-card>
            </el-timeline-item>
          </el-timeline>
          <el-empty v-else description="暂无追溯数据" />
        </el-tab-pane>

        <el-tab-pane label="事件汇总" name="summary">
          <el-table :data="eventSummary" border size="small" v-if="eventSummary.length > 0">
            <el-table-column prop="eventType" label="事件类型" width="200" />
            <el-table-column prop="count" label="事件数量" width="120" align="center" />
            <el-table-column prop="firstTime" label="首次发生" width="180" />
            <el-table-column prop="lastTime" label="最近发生" width="180" />
            <el-table-column label="占比" width="150">
              <template #default="{ row }">
                <el-progress :percentage="row.percentage" :color="getTimelineColor(row.eventType)" />
              </template>
            </el-table-column>
          </el-table>
          <el-empty v-else description="暂无事件汇总" />
        </el-tab-pane>

        <el-tab-pane label="闭环流程" name="闭环">
          <div class="loop-diagram">
            <div class="loop-node start">
              <div class="node-icon">📋</div>
              <div class="node-label">工单开工</div>
              <div class="node-event">MES_ORDER_STARTED</div>
            </div>
            <div class="loop-arrow">→</div>
            <div class="loop-node">
              <div class="node-icon">📦</div>
              <div class="node-label">配料单创建</div>
              <div class="node-event">WMS_PICKLIST_CREATED</div>
            </div>
            <div class="loop-arrow">→</div>
            <div class="loop-node">
              <div class="node-icon">🏭</div>
              <div class="node-label">工单报工</div>
              <div class="node-event">MES_ORDER_REPORTED</div>
            </div>
            <div class="loop-arrow">→</div>
            <div class="loop-node end">
              <div class="node-icon">📥</div>
              <div class="node-label">完工收货</div>
              <div class="node-event">WMS_GOODS_RECEIVED</div>
            </div>
          </div>
          <el-divider />
          <div class="loop-diagram">
            <div class="loop-node start">
              <div class="node-icon">🛒</div>
              <div class="node-label">采购收货</div>
              <div class="node-event">WMS_PURCHASE_RECEIVED</div>
            </div>
            <div class="loop-arrow">→</div>
            <div class="loop-node">
              <div class="node-icon">🔍</div>
              <div class="node-label">IQC检验</div>
              <div class="node-event">QMS_INSPECTION</div>
            </div>
            <div class="loop-arrow">→</div>
            <div class="loop-node end">
              <div class="node-icon">✅</div>
              <div class="node-label">检验完成</div>
              <div class="node-event">QMS_INSPECTION_COMPLETED</div>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>

    <el-card class="event-log" v-if="showEventLog">
      <template #header>
        <span>事件日志</span>
        <el-button text @click="showEventLog = false">收起</el-button>
      </template>
      <el-table :data="eventLogs" border size="small" max-height="300">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="eventType" label="事件类型" width="200" />
        <el-table-column prop="sourceModule" label="来源模块" width="100" />
        <el-table-column prop="sourceId" label="来源单号" width="150" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'PROCESSED' ? 'success' : row.status === 'FAILED' ? 'danger' : 'warning'" size="small">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="180" />
        <el-table-column prop="errorMsg" label="错误信息" show-overflow-tooltip />
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { getTraceList, eventApi } from '@/api/trace'

const loading = ref(false)
const activeTab = ref('chain')
const showEventLog = ref(false)

const queryForm = reactive({
  materialCode: '',
  batchNo: '',
  dateRange: []
})

const traceData = ref<any[]>([])
const eventLogs = ref<any[]>([])

// 事件类型映射
const eventTypeMap: Record<string, string> = {
  'MES_ORDER_STARTED': 'MES工单开工',
  'MES_ORDER_REPORTED': 'MES工单报工',
  'MES_ORDER_COMPLETED': 'MES工单完工',
  'WMS_PURCHASE_RECEIVED': 'WMS采购收货',
  'WMS_PICKLIST_CREATED': 'WMS配料单创建',
  'WMS_PICKLIST_COMPLETED': 'WMS配料完成',
  'WMS_GOODS_RECEIVED': 'WMS完工收货',
  'QMS_INSPECTION_COMPLETED': 'QMS检验完成',
  'QMS_INSPECTION_FAILED': 'QMS检验不合格'
}

const eventSummary = computed(() => {
  const summary: Record<string, any> = {}
  traceData.value.forEach(item => {
    if (!summary[item.eventType]) {
      summary[item.eventType] = {
        eventType: item.eventType,
        count: 0,
        firstTime: item.createTime,
        lastTime: item.createTime,
        percentage: 0
      }
    }
    summary[item.eventType].count++
    if (item.createTime < summary[item.eventType].firstTime) {
      summary[item.eventType].firstTime = item.createTime
    }
    if (item.createTime > summary[item.eventType].lastTime) {
      summary[item.eventType].lastTime = item.createTime
    }
  })
  const total = traceData.value.length
  return Object.values(summary).map((item: any) => ({
    ...item,
    percentage: total > 0 ? Math.round((item.count / total) * 100) : 0
  }))
})

function getEventTypeName(type_: string) {
  return eventTypeMap[type_] || type_
}

function getEventTypeTag(type_: string) {
  if (type_.startsWith('MES')) return 'primary'
  if (type_.startsWith('WMS')) return 'success'
  if (type_.startsWith('QMS')) return 'warning'
  return 'info'
}

function getTimelineColor(type_: string) {
  if (type_.startsWith('MES')) return '#409EFF'
  if (type_.startsWith('WMS')) return '#67C23A'
  if (type_.startsWith('QMS')) return '#E6A23C'
  return '#909399'
}

function getStatusTag(status: string) {
  if (status === 'PROCESSED') return 'success'
  if (status === 'FAILED') return 'danger'
  return 'warning'
}

async function queryTrace() {
  loading.value = true
  try {
    // TODO: 调用实际的追溯API
    const res = await getTraceList({
      materialCode: queryForm.materialCode,
      batchNo: queryForm.batchNo,
      startDate: queryForm.dateRange?.[0],
      endDate: queryForm.dateRange?.[1]
    })
    traceData.value = res.data || []
    ElMessage.success('查询成功')
  } catch (error: any) {
    ElMessage.error(error.message || '查询失败')
  } finally {
    loading.value = false
  }
}

async function loadEventLogs() {
  try {
    const res = await eventApi.list()
    eventLogs.value = res.data || []
    showEventLog.value = true
  } catch (error: any) {
    ElMessage.error(error.message || '加载事件日志失败')
  }
}

// 初始化
loadEventLogs()
</script>

<style scoped>
.material-trace-container {
  padding: 20px;
}

.trace-header {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.trace-content {
  margin-bottom: 20px;
}

.trace-card-header {
  display: flex;
  align-items: center;
  gap: 10px;
}

.trace-no {
  font-weight: bold;
  color: #409EFF;
}

.trace-card-body {
  padding: 10px 0;
}

.trace-detail {
  white-space: pre-wrap;
  font-size: 12px;
  color: #666;
}

.loop-diagram {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 30px 0;
  gap: 20px;
}

.loop-node {
  text-align: center;
  padding: 15px 20px;
  border: 2px solid #dcdfe6;
  border-radius: 8px;
  background: #f5f7fa;
  min-width: 120px;
}

.loop-node.start {
  border-color: #409EFF;
  background: #ecf5ff;
}

.loop-node.end {
  border-color: #67C23A;
  background: #f0f9eb;
}

.node-icon {
  font-size: 24px;
  margin-bottom: 5px;
}

.node-label {
  font-weight: bold;
  margin-bottom: 5px;
}

.node-event {
  font-size: 10px;
  color: #909399;
}

.loop-arrow {
  font-size: 24px;
  color: #909399;
}

.event-log {
  margin-top: 20px;
}
</style>
