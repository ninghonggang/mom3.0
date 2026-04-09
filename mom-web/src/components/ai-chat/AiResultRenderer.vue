<template>
  <div class="result-renderer">
    <div v-if="resultType === 'text'" class="text-result">
      {{ data }}
    </div>

    <div v-else-if="resultType === 'table'" class="table-result">
      <el-table :data="tableData" stripe border max-height="300" style="width: 100%">
        <el-table-column
          v-for="(col, index) in tableColumns"
          :key="index"
          :prop="col"
          :label="col"
          min-width="120"
        />
      </el-table>
    </div>

    <div v-else-if="resultType === 'chart'" class="chart-result">
      <div ref="chartRef" class="echarts-container"></div>
    </div>

    <div v-else-if="resultType === 'form'" class="form-result">
      <el-descriptions :column="2" border>
        <el-descriptions-item
          v-for="(value, key) in data"
          :key="String(key)"
          :label="String(key)"
        >
          {{ value }}
        </el-descriptions-item>
      </el-descriptions>
    </div>

    <div v-else-if="resultType === 'error'" class="error-result">
      <el-alert type="error" :closable="false">
        <template #title>
          {{ data?.error || data }}
        </template>
      </el-alert>
    </div>

    <div v-else-if="resultType === 'confirmation'" class="confirmation-result">
      <el-alert type="warning" :closable="false">
        <template #title>
          {{ operation?.natural_desc || '确认执行此操作？' }}
        </template>
      </el-alert>
      <div class="confirmation-actions">
        <el-button
          type="primary"
          size="small"
          :loading="loading"
          @click="handleConfirm"
        >
          确认
        </el-button>
        <el-button
          size="small"
          @click="handleCancel"
        >
          取消
        </el-button>
      </div>
    </div>

    <div v-else class="json-result">
      <pre>{{ JSON.stringify(data, null, 2) }}</pre>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted, watch } from 'vue'
import * as echarts from 'echarts'

interface Props {
  resultType?: string
  data?: unknown
  operation?: unknown
  messageId?: number
}

const props = defineProps<Props>()
const emit = defineEmits<{
  confirm: [messageId: number, confirmed: boolean]
  cancel: [messageId: number, confirmed: boolean]
}>()

const chartRef = ref<HTMLElement>()
const loading = ref(false)
let chartInstance: echarts.ECharts | null = null

const tableData = computed(() => {
  if (!props.data) return []
  if (Array.isArray(props.data)) return props.data
  if ((props.data as any).list) return (props.data as any).list
  return [props.data]
})

const tableColumns = computed(() => {
  if (tableData.value.length === 0) return []
  return Object.keys(tableData.value[0])
})

function handleConfirm() {
  loading.value = true
  emit('confirm', props.messageId!, true)
}

function handleCancel() {
  emit('cancel', props.messageId!, false)
}

function renderChart() {
  if (!chartRef.value || !props.data) return

  if (chartInstance) {
    chartInstance.dispose()
  }

  chartInstance = echarts.init(chartRef.value)
  const option = (props.data as any).option || props.data
  chartInstance.setOption(option)
}

onMounted(() => {
  if (props.resultType === 'chart') {
    renderChart()
  }
})

watch(() => props.resultType, (val) => {
  if (val === 'chart') {
    setTimeout(renderChart, 100)
  }
})
</script>

<style scoped>
.result-renderer {
  padding: 8px 0;
}

.table-result {
  max-height: 300px;
  overflow: auto;
}

.chart-result {
  height: 250px;
}

.echarts-container {
  width: 100%;
  height: 100%;
}

.confirmation-result {
  padding: 8px 0;
}

.confirmation-actions {
  margin-top: 12px;
  text-align: right;
}

.json-result pre {
  background: #f5f5f5;
  padding: 12px;
  border-radius: 4px;
  overflow: auto;
  font-size: 12px;
}
</style>
