<template>
  <div class="schedule-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="计划编号">
          <el-input v-model="searchForm.plan_no" placeholder="请输入计划编号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待执行" :value="1" />
            <el-option label="执行中" :value="2" />
            <el-option label="已完成" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:schedule:add')" @click="dialogVisible = true; formData = {}">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('aps:schedule:gantt')" @click="showGanttDialog = true" :disabled="!selectedPlan">
        <el-icon><Histogram /></el-icon>甘特图
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" @row-click="handleRowClick">
        <el-table-column prop="plan_no" label="计划编号" width="140" />
        <el-table-column prop="plan_type" label="计划类型" width="90" />
        <el-table-column prop="algorithm" label="算法" width="90" />
        <el-table-column prop="start_date" label="开始日期" width="110" />
        <el-table-column prop="end_date" label="结束日期" width="110" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_by" label="创建人" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="210" fixed="right">
          <template #default="{ row }">
            <el-button link type="success" size="small" v-if="hasPermission('aps:schedule:execute') && row.status === 1" @click.stop="handleExecute(row)">执行</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('aps:schedule:results')" @click.stop="handleResults(row)">结果</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:schedule:delete')" @click.stop="handleDelete(row)">删除</el-button>
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

    <!-- 新增对话框 -->
    <el-dialog v-model="dialogVisible" title="新增排程计划" width="500px">
      <el-form :model="formData" label-width="100px">
        <el-form-item label="计划编号">
          <el-input v-model="formData.plan_no" placeholder="自动生成" disabled />
        </el-form-item>
        <el-form-item label="计划类型" required>
          <el-select v-model="formData.plan_type" placeholder="请选择" style="width: 100%">
            <el-option label="粗排" value="粗排" />
            <el-option label="细排" value="细排" />
          </el-select>
        </el-form-item>
        <el-form-item label="算法">
          <el-select v-model="formData.algorithm" placeholder="请选择" style="width: 100%">
            <el-option label="遗传算法" value="遗传" />
            <el-option label="粒子群" value="粒子群" />
            <el-option label="启发式" value="启发式" />
          </el-select>
        </el-form-item>
        <el-form-item label="开始日期">
          <el-date-picker v-model="formData.start_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
        </el-form-item>
        <el-form-item label="结束日期">
          <el-date-picker v-model="formData.end_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
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

    <!-- 甘特图对话框 - 重新设计 -->
    <el-dialog v-model="showGanttDialog" title="排程甘特图" width="95%" top="2vh">
      <div class="gantt-toolbar">
        <el-button type="primary" @click="loadGanttData" :loading="ganttLoading">
          <el-icon><Refresh /></el-icon>刷新
        </el-button>
        <el-select v-model="ganttZoom" size="small" style="width: 120px">
          <el-option label="小时" :value="1" />
          <el-option label="班次" :value="2" />
          <el-option label="天" :value="3" />
        </el-select>
        <el-select v-model="ganttViewMode" size="small" style="width: 120px">
          <el-option label="按产线" value="line" />
          <el-option label="按工位" value="station" />
          <el-option label="按订单" value="order" />
        </el-select>
        <div class="gantt-legend">
          <span class="legend-item"><span class="dot waiting"></span>待执行</span>
          <span class="legend-item"><span class="dot running"></span>执行中</span>
          <span class="legend-item"><span class="dot completed"></span>已完成</span>
          <span class="legend-item"><span class="dot delayed"></span>延期</span>
        </div>
      </div>

      <div class="gantt-wrapper" v-loading="ganttLoading">
        <!-- 时间刻度头部 -->
        <div class="gantt-header">
          <div class="gantt-y-header">
            <div class="header-cell">工序/产线</div>
          </div>
          <div class="gantt-time-header" :style="{ width: ganttContentWidth + 'px' }">
            <div class="time-scale" v-if="ganttZoom === 1">
              <div v-for="hour in 24" :key="hour" class="time-cell hour">
                {{ (hour - 1) }}:00
              </div>
            </div>
            <div class="time-scale" v-else-if="ganttZoom === 2">
              <div v-for="shift in 3" :key="shift" class="time-cell shift">
                {{ shift === 1 ? '早班' : shift === 2 ? '中班' : '晚班' }}
              </div>
            </div>
            <div class="time-scale" v-else>
              <div v-for="day in dateRange" :key="day" class="time-cell day">
                {{ day }}
              </div>
            </div>
          </div>
        </div>

        <!-- 甘特图主体 -->
        <div class="gantt-body">
          <div class="gantt-y-axis">
            <div v-for="row in ganttRows" :key="row.id" class="y-cell" :style="{ height: rowHeight + 'px' }">
              <div class="y-cell-label">{{ row.name }}</div>
              <div class="y-cell-sub">{{ row.type }}</div>
            </div>
          </div>
          <div class="gantt-chart-area" :style="{ width: ganttContentWidth + 'px' }" ref="chartAreaRef">
            <div class="gantt-grid">
              <!-- 网格线 -->
              <div class="grid-lines" v-if="ganttZoom === 1">
                <div v-for="hour in 24" :key="hour" class="grid-line vertical" :style="{ left: (hour - 1) * hourWidth + 'px' }" />
              </div>
              <div class="grid-lines" v-else-if="ganttZoom === 2">
                <div v-for="shift in 3" :key="shift" class="grid-line vertical" :style="{ left: (shift - 1) * shiftWidth + 'px' }" />
              </div>
              <div class="grid-lines" v-else>
                <div v-for="day in dateRange.length" :key="day" class="grid-line vertical" :style="{ left: (day - 1) * dayWidth + 'px' }" />
              </div>
            </div>
            <!-- 甘特条 -->
            <div class="gantt-tasks">
              <div v-for="task in ganttTasks" :key="task.id" class="task-row" :style="{ top: task.top + 'px', height: rowHeight + 'px' }">
                <div
                  class="task-bar"
                  :class="task.statusClass"
                  :style="{ left: task.left + 'px', width: task.width + 'px' }"
                  :title="task.tooltip"
                  @mousedown="startDrag($event, task)"
                >
                  <span class="task-name">{{ task.name }}</span>
                  <span class="task-time">{{ task.timeText }}</span>
                </div>
              </div>
            </div>
            <!-- 拖拽预览 -->
            <div v-if="dragging" class="drag-preview" :style="{ left: dragLeft + 'px', width: dragWidth + 'px', top: dragTop + 'px', height: rowHeight + 'px' }" />
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getScheduleList, createSchedule, executeSchedule, deleteSchedule, dragUpdateSchedule, getScheduleResults } from '@/api/aps'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const showGanttDialog = ref(false)
const ganttLoading = ref(false)
const tableData = ref<any[]>([])
const formData = ref<any>({})
const selectedPlan = ref<any>(null)

// 甘特图相关
const chartAreaRef = ref<HTMLElement>()
const ganttZoom = ref(1) // 1:小时, 2:班次, 3:天
const ganttViewMode = ref('line')
const ganttRows = ref<any[]>([])
const ganttTasks = ref<any[]>([])
const dragging = ref(false)
const dragLeft = ref(0)
const dragWidth = ref(0)
const dragTop = ref(0)
const dragTask = ref<any>(null)

const rowHeight = 50
const hourWidth = 60
const shiftWidth = 200
const dayWidth = 120

const ganttContentWidth = computed(() => {
  if (ganttZoom.value === 1) return 24 * hourWidth
  if (ganttZoom.value === 2) return 3 * shiftWidth
  return dateRange.value.length * dayWidth
})

const dateRange = computed(() => {
  if (!selectedPlan.value) return []
  const start = new Date(selectedPlan.value.start_date || Date.now())
  const end = new Date(selectedPlan.value.end_date || Date.now())
  const dates = []
  for (let d = new Date(start); d <= end; d.setDate(d.getDate() + 1)) {
    dates.push(d.toISOString().split('T')[0])
  }
  return dates
})

const searchForm = reactive({ plan_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待执行', 2: '执行中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getScheduleList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.plan_no = ''; searchForm.status = ''; handleSearch() }

const handleSave = async () => {
  saveLoading.value = true
  try {
    const data = { ...formData.value }
    if (data.start_date) data.start_date = data.start_date + 'T00:00:00Z'
    if (data.end_date) data.end_date = data.end_date + 'T00:00:00Z'
    await createSchedule(data)
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleExecute = async (row: any) => {
  await executeSchedule(row.id)
  ElMessage.success('执行完成')
  loadData()
}

const handleRowClick = (row: any) => {
  selectedPlan.value = row
}

const handleResults = async (row: any) => {
  selectedPlan.value = row
  showGanttDialog.value = true
  await loadGanttData()
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该计划吗？', '提示', { type: 'warning' })
    await deleteSchedule(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const loadGanttData = async () => {
  if (!selectedPlan.value) {
    ElMessage.warning('请先选择计划')
    return
  }
  ganttLoading.value = true
  ganttTasks.value = []
  ganttRows.value = []
  try {
    const res = await getScheduleResults(selectedPlan.value.id)
    const results = res.data || []

    // 按产线分组
    const lineMap = new Map<string, any>()
    results.forEach((item: any) => {
      const lineName = item.line_name || '产线1'
      if (!lineMap.has(lineName)) {
        lineMap.set(lineName, {
          id: lineName,
          name: lineName,
          type: '产线',
          tasks: []
        })
      }
      lineMap.get(lineName).tasks.push(item)
    })

    ganttRows.value = Array.from(lineMap.values())

    // 转换任务
    const startDate = new Date(selectedPlan.value.start_date || Date.now()).getTime()
    ganttTasks.value = results.map((item: any, idx: number) => {
      const planStart = new Date(item.plan_start_time || Date.now()).getTime()
      const planEnd = new Date(item.plan_end_time || Date.now()).getTime()
      const left = (planStart - startDate) / 3600000 * hourWidth
      const width = Math.max((planEnd - planStart) / 3600000 * hourWidth, 30)
      const rowIdx = Array.from(lineMap.keys()).indexOf(item.line_name || '产线1')

      let statusClass = 'waiting'
      if (item.status === 2) statusClass = 'running'
      else if (item.status === 3) statusClass = 'completed'
      else if (planEnd < Date.now()) statusClass = 'delayed'

      return {
        id: item.id,
        name: item.order_no || `任务${idx + 1}`,
        top: rowIdx * rowHeight,
        left,
        width,
        statusClass,
        tooltip: `${item.order_no}\n开始: ${item.plan_start_time}\n结束: ${item.plan_end_time}`,
        timeText: `${new Date(planStart).toLocaleTimeString()} - ${new Date(planEnd).toLocaleTimeString()}`,
        data: item
      }
    })

    ElMessage.success('甘特图数据加载成功')
  } catch (error) {
    ElMessage.error('加载甘特图数据失败')
  } finally {
    ganttLoading.value = false
  }
}

const startDrag = (e: MouseEvent, task: any) => {
  dragging.value = true
  dragTask.value = task
  dragLeft.value = task.left
  dragWidth.value = task.width
  dragTop.value = task.top

  const onMouseMove = (e: MouseEvent) => {
    if (!dragging.value) return
    const rect = chartAreaRef.value?.getBoundingClientRect()
    if (rect) {
      dragLeft.value = Math.max(0, e.clientX - rect.left - dragWidth.value / 2)
    }
  }

  const onMouseUp = async () => {
    if (dragging.value && dragTask.value) {
      const startDate = new Date(selectedPlan.value.start_date).getTime()
      const newStartTime = Math.round(startDate + dragLeft.value / hourWidth * 3600000)
      const newEndTime = newStartTime + dragWidth.value / hourWidth * 3600000

      try {
        await dragUpdateSchedule({
          result_id: dragTask.value.id,
          line_id: dragTask.value.data.line_id || 0,
          station_id: dragTask.value.data.station_id || 0,
          plan_start_time: newStartTime / 1000,
          plan_end_time: newEndTime / 1000,
        })
        ElMessage.success('调整成功')
        await loadGanttData()
      } catch {
        ElMessage.error('保存失败')
      }
    }
    dragging.value = false
    dragTask.value = null
    document.removeEventListener('mousemove', onMouseMove)
    document.removeEventListener('mouseup', onMouseUp)
  }

  document.addEventListener('mousemove', onMouseMove)
  document.addEventListener('mouseup', onMouseUp)
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.schedule-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}

.gantt-toolbar {
  margin-bottom: 16px;
  display: flex;
  gap: 12px;
  align-items: center;

  .gantt-legend {
    margin-left: auto;
    display: flex;
    gap: 16px;

    .legend-item {
      display: flex;
      align-items: center;
      gap: 6px;
      font-size: 12px;
      color: #666;

      .dot {
        width: 12px;
        height: 12px;
        border-radius: 2px;

        &.waiting { background: #91d5ff; }
        &.running { background: #1890ff; }
        &.completed { background: #52c41a; }
        &.delayed { background: #ff4d4f; }
      }
    }
  }
}

.gantt-wrapper {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  overflow: hidden;
}

.gantt-header {
  display: flex;
  border-bottom: 1px solid #e4e7ed;
  background: #fafafa;
  position: sticky;
  top: 0;
  z-index: 10;
}

.gantt-y-header {
  width: 150px;
  flex-shrink: 0;
  border-right: 1px solid #e4e7ed;

  .header-cell {
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 600;
    color: #333;
  }
}

.gantt-time-header {
  overflow-x: auto;

  .time-scale {
    display: flex;
    min-width: 100%;

    .time-cell {
      flex-shrink: 0;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 12px;
      color: #666;
      border-right: 1px solid #f0f0f0;

      &.hour { width: 60px; height: 40px; }
      &.shift { width: 200px; height: 40px; }
      &.day { width: 120px; height: 40px; }
    }
  }
}

.gantt-body {
  display: flex;
  max-height: 60vh;
  overflow-y: auto;
}

.gantt-y-axis {
  width: 150px;
  flex-shrink: 0;
  border-right: 1px solid #e4e7ed;

  .y-cell {
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 0 12px;
    border-bottom: 1px solid #f0f0f0;
    font-size: 13px;

    .y-cell-label { font-weight: 500; color: #333; }
    .y-cell-sub { font-size: 11px; color: #999; margin-top: 2px; }
  }
}

.gantt-chart-area {
  position: relative;
  min-width: 100%;
  overflow-x: auto;
}

.gantt-grid {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;

  .grid-lines {
    position: relative;
    height: 100%;

    .grid-line.vertical {
      position: absolute;
      top: 0;
      bottom: 0;
      border-left: 1px dashed #f0f0f0;
    }
  }
}

.gantt-tasks {
  position: relative;
  min-height: 100%;

  .task-row {
    position: absolute;
    left: 0;
    right: 0;
    display: flex;
    align-items: center;
    padding: 4px 0;

    .task-bar {
      position: absolute;
      height: 36px;
      border-radius: 4px;
      cursor: pointer;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 12px;
      color: #fff;
      overflow: hidden;
      transition: transform 0.1s, box-shadow 0.1s;

      &:hover {
        transform: scaleY(1.1);
        box-shadow: 0 2px 8px rgba(0,0,0,0.15);
      }

      &.waiting { background: linear-gradient(135deg, #91d5ff, #69c0ff); }
      &.running { background: linear-gradient(135deg, #1890ff, #096dd9); }
      &.completed { background: linear-gradient(135deg, #52c41a, #389e0d); }
      &.delayed { background: linear-gradient(135deg, #ff4d4f, #cf1322); }

      .task-name {
        font-weight: 500;
        margin-left: 8px;
        white-space: nowrap;
      }

      .task-time {
        font-size: 10px;
        margin-left: 4px;
        opacity: 0.9;
        white-space: nowrap;
      }
    }
  }
}

.drag-preview {
  position: absolute;
  background: rgba(24, 144, 255, 0.3);
  border: 2px dashed #1890ff;
  border-radius: 4px;
  pointer-events: none;
}
</style>
