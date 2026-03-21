# 图表组件规范

> 本文件适用于 MOM3.0 Vue3 前端项目，使用 ECharts 5.x。

## 1. 图表设计原则

### 1.1 选择正确的图表类型

| 数据关系 | 推荐图表 | 说明 |
|----------|----------|------|
| 构成比例 | 饼图、环形图、堆叠柱状图 | 展示部分与整体关系 |
| 趋势变化 | 折线图、面积图 | 展示随时间变化 |
| 比较大小 | 柱状图、条形图 | 展示不同类别对比 |
| 分布情况 | 直方图、散点图 | 展示数据分布 |
| 流程/关系 | 桑基图、关系图 | 展示流向/关系 |

### 1.2 配色规范

```javascript
// MES 主题色板
const CHART_COLORS = [
  '#409EFF', // 主色蓝
  '#67C23A', // 成功绿
  '#E6A23C', // 警告橙
  '#F56C6C', // 危险红
  '#909399', // 信息灰
  '#9B59B6', // 紫色
  '#3498DB', // 浅蓝
  '#1ABC9C', // 青色
]

// 设备状态色
const STATUS_COLORS = {
  running: '#67C23A',
  idle: '#E6A23C',
  fault: '#F56C6C',
  maintenance: '#909399',
  offline: '#C0C4CC'
}
```

## 2. 通用图表组件

### 2.1 折线图

```vue
<template>
  <BaseChart :option="chartOption" :loading="loading" />
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  data: { type: Array, default: () => [] },
  loading: { type: Boolean, default: false }
})

const chartOption = computed(() => ({
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: ['产量', '合格率']
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: props.data.map(d => d.date)
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      name: '产量',
      type: 'line',
      data: props.data.map(d => d.output),
      smooth: true,
      areaStyle: {
        color: 'rgba(64, 158, 255, 0.2)'
      }
    },
    {
      name: '合格率',
      type: 'line',
      yAxisIndex: 1,
      data: props.data.map(d => d.rate),
      smooth: true
    }
  ]
}))
</script>
```

### 2.2 柱状图

```vue
<template>
  <BaseChart :option="chartOption" />
</template>

<script setup>
const chartOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    axisPointer: { type: 'shadow' }
  },
  legend: {
    data: ['计划', '实际']
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    data: props.data.map(d => d.name)
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      name: '计划',
      type: 'bar',
      data: props.data.map(d => d.plan),
      itemStyle: { color: '#409EFF' }
    },
    {
      name: '实际',
      type: 'bar',
      data: props.data.map(d => d.actual),
      itemStyle: { color: '#67C23A' }
    }
  ]
}))
</script>
```

### 2.3 饼图

```vue
<template>
  <BaseChart :option="chartOption" />
</template>

<script setup>
const chartOption = computed(() => ({
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)'
  },
  legend: {
    orient: 'vertical',
    left: 'left'
  },
  series: [
    {
      type: 'pie',
      radius: ['40%', '70%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 10,
        borderColor: '#fff',
        borderWidth: 2
      },
      label: {
        show: false,
        position: 'center'
      },
      emphasis: {
        label: {
          show: true,
          fontSize: 20,
          fontWeight: 'bold'
        }
      },
      data: props.data
    }
  ]
}))
</script>
```

### 2.4 仪表盘

```vue
<template>
  <BaseChart :option="chartOption" />
</template>

<script setup>
const props = defineProps({
  value: { type: Number, required: true },
  max: { type: Number, default: 100 },
  title: { type: String, default: '' }
})

const chartOption = computed(() => ({
  series: [
    {
      type: 'gauge',
      startAngle: 180,
      endAngle: 0,
      min: 0,
      max: props.max,
      splitNumber: 8,
      axisLine: {
        lineStyle: {
          width: 6,
          color: [
            [0.3, '#67C23A'],
            [0.7, '#409EFF'],
            [1, '#F56C6C']
          ]
        }
      },
      pointer: {
        icon: 'path://M12.8,0.7l12,40.1H0.7L12.8,0.7z',
        length: '12%',
        width: 20,
        offsetCenter: [0, '-60%'],
        itemStyle: { color: 'auto' }
      },
      axisTick: { length: 12, lineStyle: { color: 'auto', width: 2 } },
      splitLine: { length: 20, lineStyle: { color: 'auto', width: 5 } },
      axisLabel: {
        color: '#464646',
        fontSize: 12,
        distance: -60
      },
      title: {
        offsetCenter: [0, '-10%'],
        fontSize: 16
      },
      detail: {
        fontSize: 30,
        offsetCenter: [0, '0%'],
        valueAnimation: true,
        formatter: '{value}%',
        color: 'auto'
      },
      data: [{ value: props.value, name: props.title }]
    }
  ]
}))
</script>
```

## 3. 制造业专用图表

### 3.1 设备状态分布

```vue
<template>
  <div class="equipment-status-chart">
    <el-row :gutter="20">
      <el-col :span="8">
        <BaseChart :option="statusPieOption" :height="300" />
      </el-col>
      <el-col :span="16">
        <BaseChart :option="trendLineOption" :height="300" />
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
// 设备状态数据
const statusData = [
  { value: 85, name: '运行中', itemStyle: { color: '#67C23A' } },
  { value: 10, name: '待机', itemStyle: { color: '#E6A23C' } },
  { value: 3, name: '故障', itemStyle: { color: '#F56C6C' } },
  { value: 2, name: '维修', itemStyle: { color: '#909399' } }
]
</script>
```

### 3.2 OEE 趋势图

```vue
<template>
  <BaseChart :option="oeeOption" :height="350" />
</template>

<script setup>
const oeeOption = computed(() => ({
  title: { text: 'OEE 趋势', left: 'center' },
  tooltip: {
    trigger: 'axis',
    formatter: (params) => {
      const p = params[0]
      return `${p.name}<br/>
              OEE: ${p.value}%<br/>
              可用率: ${p.data.availability}%<br/>
              性能: ${p.data.performance}%<br/>
              质量: ${p.data.quality}%`
    }
  },
  legend: { data: ['OEE', '可用率', '性能', '质量'], bottom: 0 },
  grid: { left: '3%', right: '4%', bottom: '15%', containLabel: true },
  xAxis: { type: 'category', data: oeeData.map(d => d.date) },
  yAxis: [
    { type: 'value', name: 'OEE', max: 100, axisLabel: { formatter: '{value}%' } },
    { type: 'value', name: '其他', max: 100, axisLabel: { formatter: '{value}%' }, show: false }
  ],
  series: [
    { name: 'OEE', type: 'line', data: oeeData.map(d => d.oee), smooth: true },
    { name: '可用率', type: 'line', yAxisIndex: 1, data: oeeData.map(d => d.availability) },
    { name: '性能', type: 'line', yAxisIndex: 1, data: oeeData.map(d => d.performance) },
    { name: '质量', type: 'line', yAxisIndex: 1, data: oeeData.map(d => d.quality) }
  ]
}))
</script>
```

### 3.3 生产进度图

```vue
<template>
  <BaseChart :option="progressOption" />
</template>

<script setup>
const progressOption = computed(() => ({
  tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
  grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
  xAxis: { type: 'value', max: 'dataMax' },
  yAxis: {
    type: 'category',
    data: props.data.map(d => d.orderNo),
    inverse: true
  },
  series: [
    {
      type: 'bar',
      data: props.data.map(d => d.completed),
      label: { show: true, position: 'insideRight' },
      itemStyle: { color: '#67C23A' }
    },
    {
      type: 'bar',
      data: props.data.map(d => d.remaining),
      label: { show: true, position: 'insideRight' },
      itemStyle: { color: '#E6A23C' }
    },
    {
      type: 'bar',
      data: props.data.map(d => d.planned),
      label: { show: true, position: 'right' },
      itemStyle: { color: '#409EFF' }
    }
  ]
}))
</script>
```

### 3.4 甘特图 (APS)

```vue
<template>
  <div class="gantt-chart">
    <div class="gantt-header">
      <div class="gantt-row-title">工单</div>
      <div class="gantt-timeline">
        <div v-for="day in timeline" :key="day" class="gantt-day">{{ day }}</div>
      </div>
    </div>
    <div class="gantt-body">
      <div v-for="task in tasks" :key="task.id" class="gantt-row">
        <div class="gantt-row-title">{{ task.name }}</div>
        <div class="gantt-row-bars">
          <div
            class="gantt-bar"
            :style="getBarStyle(task)"
            :class="'status-' + task.status"
          >
            {{ task.progress }}%
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const getBarStyle = (task) => {
  const startDay = getDayIndex(task.startDate)
  const duration = task.days
  return {
    left: `${startDay * dayWidth}px`,
    width: `${duration * dayWidth}px`
  }
}
</script>
```

## 4. 图表交互

### 4.1 点击事件

```javascript
const chartRef = ref()

const handleChartClick = (params) => {
  if (params.componentType === 'series') {
    const data = params.data
    router.push(`/detail/${data.id}`)
  }
}

// 配置点击事件
const chartOption = computed(() => ({
  onClick: handleChartClick
}))
```

### 4.2 图例筛选

```javascript
const chartOption = computed(() => ({
  legend: {
    selected: {
      '产量': true,
      '合格率': false
    }
  }
}))
```

### 4.3 数据缩放

```javascript
const chartOption = computed(() => ({
  dataZoom: [
    { type: 'inside', start: 0, end: 100 },
    { type: 'slider', start: 0, end: 100 }
  ]
}))
```

## 5. 响应式图表

### 5.1 自动调整大小

```vue
<template>
  <BaseChart ref="chartRef" :option="option" autoresize />
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const chartRef = ref()
let resizeObserver = null

onMounted(() => {
  resizeObserver = new ResizeObserver(() => {
    chartRef.value?.resize()
  })
  resizeObserver.observe(chartRef.value?.$el)
})

onUnmounted(() => {
  resizeObserver?.disconnect()
})
</script>
```

## 6. 性能优化

### 6.1 大数据优化

```javascript
const chartOption = computed(() => ({
  // 关闭动画提升性能
  animation: false,

  // 使用采样
  sampling: 'lttb',

  // 渐进渲染
  progressive: 1000,

  // 关闭 hover 效果
  hoverAnimation: false
}))
```

### 6.2 懒加载

```vue
<script setup>
import { ref, onMounted } from 'vue'
import { useIntersectionObserver } from '@vueuse/core'

const chartRef = ref()
const isVisible = ref(false)

useIntersectionObserver(chartRef, ([{ isIntersecting }]) => {
  if (isIntersecting) {
    isVisible.value = true
  }
})
</script>

<template>
  <BaseChart v-if="isVisible" :option="option" />
  <div v-else class="chart-placeholder">加载中...</div>
</template>
```

## 7. 主题定制

### 7.1 浅色主题

```javascript
const lightTheme = {
  backgroundColor: '#ffffff',
  textColor: '#303133',
  axisLineColor: '#dcdfe6',
  splitLineColor: '#e4e7ed'
}
```

### 7.2 深色主题

```javascript
const darkTheme = {
  backgroundColor: '#1a1a1a',
  textColor: '#e5e5e5',
  axisLineColor: '#3a3a3a',
  splitLineColor: '#2a2a2a'
}
```
