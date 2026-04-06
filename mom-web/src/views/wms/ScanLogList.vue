<template>
  <div class="scan-log-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="扫描类型">
          <el-select v-model="searchForm.scan_type" placeholder="请选择" clearable>
            <el-option label="物料" value="MATERIAL" />
            <el-option label="产品" value="PRODUCT" />
            <el-option label="工装" value="TOOL" />
            <el-option label="容器" value="CONTAINER" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="成功" value="SUCCESS" />
            <el-option label="失败" value="FAILED" />
            <el-option label="重复" value="DUPLICATE" />
          </el-select>
        </el-form-item>
        <el-form-item label="时间范围">
          <el-date-picker v-model="searchForm.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="scan_time" label="扫描时间" width="160" />
        <el-table-column prop="scan_type" label="扫描类型" width="100">
          <template #default="{ row }">
            <el-tag>{{ getScanTypeText(row.scan_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="scan_code" label="扫描编码" min-width="180" show-overflow-tooltip />
        <el-table-column prop="scan_device" label="设备" width="100" />
        <el-table-column prop="business_type" label="业务类型" width="120" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 'SUCCESS' ? 'success' : row.status === 'FAILED' ? 'danger' : 'warning'">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="fail_reason" label="失败原因" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleDetail(row)">详情</el-button>
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getScanLogList } from '@/api/wms'

const loading = ref(false)
const tableData = ref<any[]>([])
const searchForm = reactive({ scan_type: '', status: '', dateRange: [] as string[] })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getScanTypeText = (val: string) => ({ MATERIAL: '物料', PRODUCT: '产品', TOOL: '工装', CONTAINER: '容器' }[val] || val)
const getStatusText = (val: string) => ({ SUCCESS: '成功', FAILED: '失败', DUPLICATE: '重复' }[val] || val)

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.scan_type) params.scan_type = searchForm.scan_type
    if (searchForm.status) params.status = searchForm.status
    if (searchForm.dateRange && searchForm.dateRange.length === 2) {
      params.start_time = searchForm.dateRange[0]
      params.end_time = searchForm.dateRange[1]
    }
    const res = await getScanLogList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.scan_type = ''; searchForm.status = ''; searchForm.dateRange = []; handleSearch() }
const handleDetail = (row: any) => { ElMessage.info('查看详情: ' + row.scan_code) }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.scan-log-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
