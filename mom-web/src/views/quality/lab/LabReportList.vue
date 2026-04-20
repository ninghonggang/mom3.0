<template>
  <div class="lab-report-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="报告编号">
          <el-input v-model="searchForm.report_no" placeholder="请输入报告编号" clearable />
        </el-form-item>
        <el-form-item label="检验类型">
          <el-select v-model="searchForm.inspection_type" placeholder="请选择" clearable>
            <el-option label="来料检验" value="IQC" />
            <el-option label="过程检验" value="IPQC" />
            <el-option label="成品检验" value="FQC" />
            <el-option label="出货检验" value="OQC" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('quality:lab:report:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>生成报告
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="report_no" label="报告编号" width="150" />
        <el-table-column prop="sample_no" label="样品编号" width="150" />
        <el-table-column prop="inspection_type" label="检验类型" width="100" />
        <el-table-column prop="conclusion" label="检验结论" width="120">
          <template #default="{ row }">
            <el-tag :type="row.conclusion === 'PASS' ? 'success' : 'danger'">{{ row.conclusion }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="inspector" label="检验员" width="100" />
        <el-table-column prop="inspect_time" label="检验时间" width="160" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleView(row)">查看</el-button>
            <el-button link type="success" v-if="hasPermission('quality:lab:report:print')" @click="handlePrint(row)">打印</el-button>
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

const searchForm = reactive({
  report_no: '',
  inspection_type: ''
})
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.report_no) params.report_no = searchForm.report_no
    if (searchForm.inspection_type) params.inspection_type = searchForm.inspection_type

    const res = await request.get('/quality/lab/reports/list', { params })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.report_no = ''
  searchForm.inspection_type = ''
  handleSearch()
}

const handleAdd = () => {
  ElMessage.info('生成报告功能开发中')
}

const handleView = (row: any) => {
  ElMessage.info('查看功能开发中')
}

const handlePrint = (row: any) => {
  ElMessage.info('打印功能开发中')
}

onMounted(() => {
  loadData()
})
</script>
