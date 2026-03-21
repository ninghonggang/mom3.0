<template>
  <div class="trace-query">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="追溯类型">
          <el-radio-group v-model="searchForm.traceType">
            <el-radio value="serial">序列号</el-radio>
            <el-radio value="batch">批次号</el-radio>
            <el-radio value="order">订单号</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="追溯码">
          <el-input v-model="searchForm.traceCode" placeholder="请输入追溯码" clearable style="width: 200px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>查询
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card v-if="traceData">
      <template #header>
        <div class="card-header">
          <span>追溯结果</span>
        </div>
      </template>

      <el-descriptions :column="2" border>
        <el-descriptions-item label="产品名称">{{ traceData.product_name }}</el-descriptions-item>
        <el-descriptions-item label="产品编码">{{ traceData.product_code }}</el-descriptions-item>
        <el-descriptions-item label="批次号">{{ traceData.batch_no }}</el-descriptions-item>
        <el-descriptions-item label="序列号">{{ traceData.serial_number }}</el-descriptions-item>
        <el-descriptions-item label="生产日期">{{ traceData.produce_date }}</el-descriptions-item>
        <el-descriptions-item label="生产车间">{{ traceData.workshop }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>生产过程</el-divider>

      <el-timeline>
        <el-timeline-item
          v-for="(item, index) in traceData.processes"
          :key="index"
          :timestamp="item.timestamp"
          :type="item.status === 'success' ? 'success' : 'primary'"
          placement="top"
        >
          <el-card>
            <h4>{{ item.process_name }}</h4>
            <p>工序: {{ item.process }}</p>
            <p>操作人: {{ item.operator }}</p>
            <p>工位: {{ item.station }}</p>
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </el-card>

    <el-empty v-else-if="searched" description="未找到追溯信息" />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { traceBySerial, traceByBatch, traceByOrder } from '@/api/trace'

const searched = ref(false)
const traceData = ref<any>(null)

const searchForm = reactive({
  traceType: 'serial',
  traceCode: ''
})

const handleSearch = async () => {
  if (!searchForm.traceCode) return

  searched.value = true
  traceData.value = null

  try {
    let res
    if (searchForm.traceType === 'serial') {
      res = await traceBySerial(searchForm.traceCode)
    } else if (searchForm.traceType === 'batch') {
      res = await traceByBatch(searchForm.traceCode)
    } else {
      res = await traceByOrder(parseInt(searchForm.traceCode))
    }
    traceData.value = res.data
  } catch (e) {
    console.error(e)
  }
}
</script>

<style scoped lang="scss">
.trace-query {
  .search-card { margin-bottom: 16px; }
  .card-header { display: flex; justify-content: space-between; align-items: center; }
}
</style>
