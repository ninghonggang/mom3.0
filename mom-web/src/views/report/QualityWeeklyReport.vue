<template>
  <div class="quality-weekly-report">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="年度">
          <el-input-number v-model="searchForm.year" :min="2020" :max="2099" />
        </el-form-item>
        <el-form-item label="周次">
          <el-input-number v-model="searchForm.week" :min="1" :max="53" />
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
        <el-table-column prop="report_year" label="年度" width="80" />
        <el-table-column prop="report_week" label="周次" width="70" />
        <el-table-column label="日期范围" width="200">
          <template #default="{ row }">
            {{ row.start_date?.slice(0,10) || '-' }} ~ {{ row.end_date?.slice(0,10) || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="workshop_name" label="车间" width="120" />
        <el-table-column prop="total_inspection_qty" label="总检验数" width="90" />
        <el-table-column prop="qualified_qty" label="合格数" width="80" />
        <el-table-column prop="defect_qty" label="不良数" width="80">
          <template #default="{ row }">
            <span class="danger" v-if="row.defect_qty > 0">{{ row.defect_qty }}</span>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column label="合格率%" width="90">
          <template #default="{ row }">
            <span :class="getPassRateClass(row.pass_rate)">{{ row.pass_rate ? row.pass_rate.toFixed(1) + '%' : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column label="IQC" width="80">
          <template #default="{ row }">
            <el-tooltip content="IQC检验数" placement="top">{{ row.iqc_insp_qty || 0 }}</el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="IPQC" width="80">
          <template #default="{ row }">{{ row.ipqc_insp_qty || 0 }}</template>
        </el-table-column>
        <el-table-column label="FQC" width="80">
          <template #default="{ row }">{{ row.fqc_insp_qty || 0 }}</template>
        </el-table-column>
        <el-table-column label="OQC" width="80">
          <template #default="{ row }">{{ row.oqc_insp_qty || 0 }}</template>
        </el-table-column>
        <el-table-column prop="ncr_count" label="NCR数" width="80">
          <template #default="{ row }">
            <el-tag type="danger" size="small" v-if="row.ncr_count > 0">{{ row.ncr_count }}</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="customer_complaint_count" label="客诉数" width="80">
          <template #default="{ row }">
            <el-tag type="warning" size="small" v-if="row.customer_complaint_count > 0">{{ row.customer_complaint_count }}</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="80" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleDetail(row)">详情</el-button>
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
    <el-dialog v-model="detailVisible" title="质量周报详情" width="800px" destroy-on-close>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="年度">{{ detailData.report_year }}</el-descriptions-item>
        <el-descriptions-item label="周次">{{ detailData.report_week }}</el-descriptions-item>
        <el-descriptions-item label="开始日期">{{ detailData.start_date?.slice(0,10) }}</el-descriptions-item>
        <el-descriptions-item label="结束日期">{{ detailData.end_date?.slice(0,10) }}</el-descriptions-item>
        <el-descriptions-item label="车间" :span="2">{{ detailData.workshop_name }}</el-descriptions-item>
      </el-descriptions>
      <el-divider>汇总</el-divider>
      <el-descriptions :column="4" border size="small">
        <el-descriptions-item label="总检验数">{{ detailData.total_inspection_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格数">{{ detailData.qualified_qty }}</el-descriptions-item>
        <el-descriptions-item label="不良数">{{ detailData.defect_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格率%">{{ detailData.pass_rate ? detailData.pass_rate.toFixed(2) + '%' : '-' }}</el-descriptions-item>
      </el-descriptions>
      <el-divider>IQC 来料检验</el-divider>
      <el-descriptions :column="3" border size="small">
        <el-descriptions-item label="检验数">{{ detailData.iqc_insp_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格数">{{ detailData.iqc_qualified_qty }}</el-descriptions-item>
        <el-descriptions-item label="不良数">{{ detailData.iqc_defect_qty }}</el-descriptions-item>
      </el-descriptions>
      <el-divider>IPQC 过程检验</el-divider>
      <el-descriptions :column="3" border size="small">
        <el-descriptions-item label="检验数">{{ detailData.ipqc_insp_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格数">{{ detailData.ipqc_qualified_qty }}</el-descriptions-item>
        <el-descriptions-item label="不良数">{{ detailData.ipqc_defect_qty }}</el-descriptions-item>
      </el-descriptions>
      <el-divider>FQC 成品检验</el-divider>
      <el-descriptions :column="3" border size="small">
        <el-descriptions-item label="检验数">{{ detailData.fqc_insp_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格数">{{ detailData.fqc_qualified_qty }}</el-descriptions-item>
        <el-descriptions-item label="不良数">{{ detailData.fqc_defect_qty }}</el-descriptions-item>
      </el-descriptions>
      <el-divider>OQC 出货检验</el-divider>
      <el-descriptions :column="3" border size="small">
        <el-descriptions-item label="检验数">{{ detailData.oqc_insp_qty }}</el-descriptions-item>
        <el-descriptions-item label="合格数">{{ detailData.oqc_qualified_qty }}</el-descriptions-item>
        <el-descriptions-item label="不良数">{{ detailData.oqc_defect_qty }}</el-descriptions-item>
      </el-descriptions>
      <el-divider>质量事件</el-divider>
      <el-descriptions :column="2" border size="small">
        <el-descriptions-item label="NCR数量">
          <el-tag type="danger">{{ detailData.ncr_count || 0 }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="客诉数量">
          <el-tag type="warning">{{ detailData.customer_complaint_count || 0 }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ detailData.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>

    <!-- 生成报表弹窗 -->
    <el-dialog v-model="generateVisible" title="生成质量周报" width="500px" destroy-on-close>
      <el-form ref="genFormRef" :model="generateForm" :rules="genRules" label-width="100px">
        <el-form-item label="年度" prop="year">
          <el-input-number v-model="generateForm.year" :min="2020" :max="2099" />
        </el-form-item>
        <el-form-item label="周次" prop="week">
          <el-input-number v-model="generateForm.week" :min="1" :max="53" />
        </el-form-item>
        <el-form-item label="车间ID" prop="workshop_id">
          <el-input-number v-model="generateForm.workshop_id" :min="1" />
        </el-form-item>
        <el-form-item label="车间名称">
          <el-input v-model="generateForm.workshop_name" placeholder="车间名称" />
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
const searchForm = reactive({ year: new Date().getFullYear(), week: 1, workshop_name: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const generateForm = reactive({
  year: new Date().getFullYear(),
  week: getWeekNumber(new Date()),
  workshop_id: 1,
  workshop_name: ''
})
const genRules: FormRules = {
  year: [{ required: true, message: '请输入年度', trigger: 'blur' }],
  week: [{ required: true, message: '请输入周次', trigger: 'blur' }],
  workshop_id: [{ required: true, message: '请输入车间ID', trigger: 'blur' }]
}

function getWeekNumber(d: Date): number {
  const date = new Date(Date.UTC(d.getFullYear(), d.getMonth(), d.getDate()))
  const dayNum = date.getUTCDay() || 7
  date.setUTCDate(date.getUTCDate() + 4 - dayNum)
  const yearStart = new Date(Date.UTC(date.getUTCFullYear(), 0, 1))
  return Math.ceil((((date.getTime() - yearStart.getTime()) / 86400000) + 1) / 7)
}

const getPassRateClass = (rate: number) => {
  if (!rate) return ''
  if (rate >= 98) return 'rate-green'
  if (rate >= 95) return 'rate-orange'
  return 'rate-red'
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { page: pagination.page, page_size: pagination.pageSize }
    if (searchForm.year) params.year = searchForm.year
    if (searchForm.week) params.week = searchForm.week
    if (searchForm.workshop_name) params.workshop_name = searchForm.workshop_name
    const res = await request.get('/report/quality-weekly/list', { params })
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
  searchForm.year = new Date().getFullYear()
  searchForm.week = 1
  searchForm.workshop_name = ''
  handleSearch()
}

const handleDetail = async (row: any) => {
  try {
    const res = await request.get(`/report/quality-weekly/${row.id}`)
    detailData.value = res.data || res
    detailVisible.value = true
  } catch (e: any) {
    ElMessage.error(e.message || '加载详情失败')
  }
}

const showGenerateDialog = () => {
  Object.assign(generateForm, {
    year: new Date().getFullYear(),
    week: getWeekNumber(new Date()),
    workshop_id: 1,
    workshop_name: ''
  })
  generateVisible.value = true
}

const handleGenerate = async () => {
  const valid = await genFormRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    await request.post('/report/quality-weekly/generate', {
      year: generateForm.year,
      week: generateForm.week,
      workshop_id: generateForm.workshop_id,
      workshop_name: generateForm.workshop_name
    })
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
.quality-weekly-report {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .danger { color: #f56c6c; font-weight: bold; }
  .rate-green { color: #67c23a; font-weight: bold; }
  .rate-orange { color: #e6a23c; font-weight: bold; }
  .rate-red { color: #f56c6c; font-weight: bold; }
}
</style>
