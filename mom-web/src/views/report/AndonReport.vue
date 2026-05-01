<template>
  <div class="andon-report">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="报表日期">
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
        <el-form-item label="车间">
          <el-input v-model="searchForm.workshop_name" placeholder="车间名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="showGenerateDialog">生成报表</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" stripe>
        <el-table-column prop="report_date" label="报表日期" width="120" />
        <el-table-column prop="workshop_name" label="车间" width="120" />
        <el-table-column prop="line_name" label="产线" width="120" />
        <el-table-column prop="station_name" label="工位" width="100" />
        <el-table-column prop="total_call_count" label="呼叫总数" width="90">
          <template #default="{ row }">
            <span class="stat-num danger">{{ row.total_call_count }}</span>
          </template>
        </el-table-column>
        <el-table-column label="类型分布" min-width="180">
          <template #default="{ row }">
            <span class="type-tag"><el-tag size="small" type="danger">设备{{ row.equipment_call_count || 0 }}</el-tag></span>
            <span class="type-tag"><el-tag size="small" type="warning">物料{{ row.material_call_count || 0 }}</el-tag></span>
            <span class="type-tag"><el-tag size="small" type="info">质量{{ row.quality_call_count || 0 }}</el-tag></span>
            <span class="type-tag"><el-tag size="small">其他{{ row.other_call_count || 0 }}</el-tag></span>
          </template>
        </el-table-column>
        <el-table-column prop="avg_response_time" label="平均响应(min)" width="130">
          <template #default="{ row }">{{ row.avg_response_time ? row.avg_response_time.toFixed(1) : '-' }}</template>
        </el-table-column>
        <el-table-column prop="avg_resolve_time" label="平均解决(min)" width="130">
          <template #default="{ row }">{{ row.avg_resolve_time ? row.avg_resolve_time.toFixed(1) : '-' }}</template>
        </el-table-column>
        <el-table-column prop="unresolved_count" label="未解决" width="90">
          <template #default="{ row }">
            <el-tag type="danger" v-if="row.unresolved_count > 0">{{ row.unresolved_count }}</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleDetail(row)">查看</el-button>
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
    <el-dialog v-model="detailVisible" title="报表详情" width="700px" destroy-on-close>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="报表日期">{{ detailData.report_date }}</el-descriptions-item>
        <el-descriptions-item label="车间">{{ detailData.workshop_name }}</el-descriptions-item>
        <el-descriptions-item label="产线">{{ detailData.line_name }}</el-descriptions-item>
        <el-descriptions-item label="工位">{{ detailData.station_name }}</el-descriptions-item>
        <el-descriptions-item label="呼叫总数"><span class="stat-num danger">{{ detailData.total_call_count }}</span></el-descriptions-item>
        <el-descriptions-item label="设备呼叫">{{ detailData.equipment_call_count }}</el-descriptions-item>
        <el-descriptions-item label="物料呼叫">{{ detailData.material_call_count }}</el-descriptions-item>
        <el-descriptions-item label="质量呼叫">{{ detailData.quality_call_count }}</el-descriptions-item>
        <el-descriptions-item label="其他呼叫">{{ detailData.other_call_count }}</el-descriptions-item>
        <el-descriptions-item label="平均响应时间">{{ detailData.avg_response_time ? detailData.avg_response_time.toFixed(1) + '分钟' : '-' }}</el-descriptions-item>
        <el-descriptions-item label="平均解决时间">{{ detailData.avg_resolve_time ? detailData.avg_resolve_time.toFixed(1) + '分钟' : '-' }}</el-descriptions-item>
        <el-descriptions-item label="未解决数">{{ detailData.unresolved_count || 0 }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>

    <!-- 生成报表弹窗 -->
    <el-dialog v-model="generateVisible" title="生成安东报表" width="500px" destroy-on-close>
      <el-form ref="genFormRef" :model="generateForm" :rules="genRules" label-width="110px">
        <el-form-item label="报表日期" prop="report_date">
          <el-date-picker
            v-model="generateForm.report_date"
            type="date"
            placeholder="选择日期"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item label="车间ID" prop="workshop_id">
          <el-input-number v-model="generateForm.workshop_id" :min="1" />
        </el-form-item>
        <el-form-item label="车间名称">
          <el-input v-model="generateForm.workshop_name" placeholder="车间名称" />
        </el-form-item>
        <el-form-item label="产线ID" prop="line_id">
          <el-input-number v-model="generateForm.line_id" :min="1" />
        </el-form-item>
        <el-form-item label="产线名称">
          <el-input v-model="generateForm.line_name" placeholder="产线名称" />
        </el-form-item>
        <el-form-item label="工位ID" prop="station_id">
          <el-input-number v-model="generateForm.station_id" :min="1" />
        </el-form-item>
        <el-form-item label="工位名称">
          <el-input v-model="generateForm.station_name" placeholder="工位名称" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="generateVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleGenerate">生成</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, FormInstance, FormRules } from 'element-plus'
import request from '@/utils/request'

const loading = ref(false)
const tableData = ref<any[]>([])
const detailVisible = ref(false)
const generateVisible = ref(false)
const submitLoading = ref(false)
const genFormRef = ref<FormInstance>()
const detailData = ref<any>({})
const dateRange = ref<[string, string] | null>(null)
const searchForm = reactive({ start_date: '', end_date: '', workshop_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const generateForm = reactive({
  report_date: '', workshop_id: 1, workshop_name: '', line_id: 1, line_name: '', station_id: 1, station_name: ''
})
const genRules: FormRules = {
  report_date: [{ required: true, message: '请选择报表日期', trigger: 'change' }],
  workshop_id: [{ required: true, message: '请输入车间ID', trigger: 'blur' }],
  line_id: [{ required: true, message: '请输入产线ID', trigger: 'blur' }],
  station_id: [{ required: true, message: '请输入工位ID', trigger: 'blur' }]
}

const handleDateChange = (val: [string, string] | null) => {
  if (val) { searchForm.start_date = val[0]; searchForm.end_date = val[1] }
  else { searchForm.start_date = ''; searchForm.end_date = '' }
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.start_date) params.start_date = searchForm.start_date
    if (searchForm.end_date) params.end_date = searchForm.end_date
    if (searchForm.workshop_name) params.workshop_name = searchForm.workshop_name
    const res = await request.get('/report/andon/list', { params })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } catch (e: any) {
    ElMessage.error(e.message || '加载失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => {
  Object.assign(searchForm, { start_date: '', end_date: '', workshop_name: '' })
  dateRange.value = null
  handleSearch()
}

const handleDetail = async (row: any) => {
  try {
    const res = await request.get(`/report/andon/${row.id}`)
    detailData.value = res.data || res
    detailVisible.value = true
  } catch (e: any) {
    ElMessage.error(e.message || '加载详情失败')
  }
}

const showGenerateDialog = () => {
  const today = new Date().toISOString().slice(0, 10)
  Object.assign(generateForm, {
    report_date: today, workshop_id: 1, workshop_name: '', line_id: 1, line_name: '', station_id: 1, station_name: ''
  })
  generateVisible.value = true
}
const handleGenerate = async () => {
  const valid = await genFormRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    await request.post('/report/andon/generate', generateForm)
    ElMessage.success('报表生成成功')
    generateVisible.value = false
    loadData()
  } catch (e: any) {
    ElMessage.error(e.message || '生成失败')
  } finally {
    submitLoading.value = false
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.andon-report {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .type-tag { margin-right: 6px; }
  .stat-num {
    font-weight: bold;
    &.danger { color: #f56c6c; }
    &.warning { color: #e6a23c; }
  }
}
</style>
