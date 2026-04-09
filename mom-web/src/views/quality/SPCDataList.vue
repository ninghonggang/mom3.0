<template>
  <div class="spc-data-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工序">
          <el-input v-model="searchForm.process_name" placeholder="请输入工序" clearable />
        </el-form-item>
        <el-form-item label="检测项目">
          <el-input v-model="searchForm.check_item" placeholder="请输入项目" clearable />
        </el-form-item>
        <el-form-item label="时间范围">
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
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <div class="chart-container">
        <div ref="chartRef" class="spc-chart"></div>
      </div>
    </el-card>

    <el-card class="table-card">
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="process_name" label="工序" width="100" />
        <el-table-column prop="check_item" label="检测项目" width="120" />
        <el-table-column prop="check_value" label="检测值" width="100">
          <template #default="{ row }">
            <span :class="getValueClass(row)">{{ row.check_value }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="usl" label="规格上限" width="100" />
        <el-table-column prop="lsl" label="规格下限" width="100" />
        <el-table-column prop="ucl" label="控制上限" width="100" />
        <el-table-column prop="cl" label="中心值" width="100" />
        <el-table-column prop="lcl" label="控制下限" width="100" />
        <el-table-column prop="check_time" label="检测时间" width="160" />
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getSPCData, getSPCChart } from '@/api/quality'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const chartRef = ref<HTMLElement>()
const dateRange = ref<string[]>([])

const searchForm = reactive({ process_name: '', check_item: '', start_date: '', end_date: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const handleDateChange = (val: string[]) => {
  if (val && val.length === 2) {
    searchForm.start_date = val[0]
    searchForm.end_date = val[1]
  } else {
    searchForm.start_date = ''
    searchForm.end_date = ''
  }
}

const getValueClass = (row: any) => {
  if (row.check_value > row.ucl || row.check_value < row.lcl) {
    return 'value-out-of-control'
  }
  if (row.check_value > row.usl || row.check_value < row.lsl) {
    return 'value-out-of-spec'
  }
  return ''
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getSPCData({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  searchForm.process_name = ''
  searchForm.check_item = ''
  searchForm.start_date = ''
  searchForm.end_date = ''
  dateRange.value = []
  handleSearch()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.spc-data-list {
  .search-card { margin-bottom: 16px; }
  .table-card { margin-top: 16px; }
  .chart-container {
    height: 300px;
    .spc-chart { width: 100%; height: 100%; }
  }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .value-out-of-control { color: #f56c6c; font-weight: bold; }
  .value-out-of-spec { color: #e6a23c; font-weight: bold; }
}
</style>
