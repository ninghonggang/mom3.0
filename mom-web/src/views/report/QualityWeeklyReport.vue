<template>
  <div class="quality-weekly-report">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="报表周">
          <el-date-picker v-model="searchForm.report_week" type="week" placeholder="选择周" value-format="YYYY-WW" clearable />
        </el-form-item>
        <el-form-item label="车间">
          <el-input v-model="searchForm.workshop_name" placeholder="请输入车间名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('report:qualityWeekly:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button type="success" v-if="hasPermission('report:qualityWeekly:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="report_week" label="报表周" width="120" />
        <el-table-column prop="workshop_id" label="车间ID" width="100" />
        <el-table-column prop="inspection_count" label="检验数量" width="100" />
        <el-table-column prop="qualified_count" label="合格数量" width="100" />
        <el-table-column prop="quality_rate" label="质量率" width="100">
          <template #default="{ row }">
            {{ row.quality_rate }}%
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('report:qualityWeekly:detail')" @click="handleDetail(row)">明细</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('report:qualityWeekly:edit')" @click="handleEdit(row)">编辑</el-button>
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
import request from '@/utils/request'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ report_week: '', workshop_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/report/quality-weekly/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.report_week = ''; searchForm.workshop_name = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增质量周报') }
const handleExport = () => { ElMessage.info('导出数据') }
const handleDetail = (row: any) => { ElMessage.info('查看明细') }
const handleEdit = (row: any) => { ElMessage.info('编辑') }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.quality-weekly-report {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
