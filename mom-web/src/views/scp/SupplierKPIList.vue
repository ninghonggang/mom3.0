<template>
  <div class="supplier-kpi-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="供应商名称">
          <el-input v-model="searchForm.supplier_name" placeholder="请输入供应商名称" clearable />
        </el-form-item>
        <el-form-item label="评估月份">
          <el-date-picker v-model="searchForm.evaluation_month" type="month" value-format="YYYY-MM" placeholder="选择月份" clearable />
        </el-form-item>
        <el-form-item label="评级">
          <el-select v-model="searchForm.grade" placeholder="请选择" clearable>
            <el-option label="A (90-100)" value="A" />
            <el-option label="B (80-89)" value="B" />
            <el-option label="C (70-79)" value="C" />
            <el-option label="D (60-69)" value="D" />
            <el-option label="E (<60)" value="E" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="info" @click="handleRanking">供应商排名</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="supplier_code" label="供应商编码" width="120" />
        <el-table-column prop="supplier_name" label="供应商名称" min-width="180" />
        <el-table-column prop="evaluation_month" label="评估月份" width="110" />
        <el-table-column prop="on_time_rate" label="准时率(%)" width="110" align="right">
          <template #default="{ row }">
            <span :class="getRateClass(row.on_time_rate)">{{ row.on_time_rate != null ? Number(row.on_time_rate).toFixed(1) + '%' : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="quality_rate" label="合格率(%)" width="110" align="right">
          <template #default="{ row }">
            <span :class="getRateClass(row.quality_rate)">{{ row.quality_rate != null ? Number(row.quality_rate).toFixed(1) + '%' : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="response_rate" label="响应率(%)" width="110" align="right">
          <template #default="{ row }">
            <span :class="getRateClass(row.response_rate)">{{ row.response_rate != null ? Number(row.response_rate).toFixed(1) + '%' : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="total_score" label="综合得分" width="100" align="right">
          <template #default="{ row }">
            <span class="score-text">{{ row.total_score != null ? Number(row.total_score).toFixed(1) : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="grade" label="评级" width="80" align="center">
          <template #default="{ row }">
            <el-tag size="small" :type="getGradeType(row.grade)">{{ row.grade || '-' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="rank" label="排名" width="70" align="center" />
        <el-table-column prop="delivery_count" label="交货次数" width="90" align="right" />
        <el-table-column prop="quality_count" label="检验次数" width="90" align="right" />
        <el-table-column prop="created_at" label="评估时间" width="160" />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleViewDetail(row)">查看详情</el-button>
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

    <!-- 绩效详情Dialog -->
    <el-dialog v-model="detailDialogVisible" title="供应商绩效详情" width="900px">
      <div v-if="detailData.supplier" class="detail-header">
        <el-descriptions :column="3" border size="small">
          <el-descriptions-item label="供应商编码">{{ detailData.supplier.supplier_code }}</el-descriptions-item>
          <el-descriptions-item label="供应商名称">{{ detailData.supplier.supplier_name }}</el-descriptions-item>
          <el-descriptions-item label="评估月份">{{ detailData.supplier.evaluation_month }}</el-descriptions-item>
          <el-descriptions-item label="综合得分">{{ detailData.supplier.total_score }}</el-descriptions-item>
          <el-descriptions-item label="评级">
            <el-tag size="small" :type="getGradeType(detailData.supplier.grade)">{{ detailData.supplier.grade }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="评估人">{{ detailData.supplier.evaluated_by || '-' }}</el-descriptions-item>
        </el-descriptions>
      </div>

      <el-tabs v-model="detailTab">
        <el-tab-pane label="评分明细" name="score">
          <el-descriptions :column="2" border size="small">
            <el-descriptions-item label="准时率得分(权重40%)">
              {{ detailData.on_time_rate != null ? Number(detailData.on_time_rate).toFixed(1) : '-' }}
            </el-descriptions-item>
            <el-descriptions-item label="来料合格率得分(权重40%)">
              {{ detailData.quality_rate != null ? Number(detailData.quality_rate).toFixed(1) : '-' }}
            </el-descriptions-item>
            <el-descriptions-item label="响应率得分(权重20%)">
              {{ detailData.response_rate != null ? Number(detailData.response_rate).toFixed(1) : '-' }}
            </el-descriptions-item>
            <el-descriptions-item label="综合得分">
              <span class="score-text">{{ detailData.supplier?.total_score || '-' }}</span>
            </el-descriptions-item>
          </el-descriptions>
        </el-tab-pane>

        <el-tab-pane label="交货记录" name="delivery">
          <el-table :data="detailData.delivery_records || []" border size="small">
            <el-table-column prop="po_no" label="采购单号" width="140" />
            <el-table-column prop="delivery_date" label="交货日期" width="120" />
            <el-table-column prop="planned_date" label="计划日期" width="120" />
            <el-table-column prop="material_code" label="物料编码" width="110" />
            <el-table-column prop="delivery_qty" label="交货数量" width="90" align="right" />
            <el-table-column prop="on_time" label="是否准时" width="90" align="center">
              <template #default="{ row }">
                <el-tag size="small" :type="row.on_time ? 'success' : 'danger'">{{ row.on_time ? '是' : '否' }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="quality_status" label="质量状态" width="100" align="center">
              <template #default="{ row }">
                <el-tag size="small" :type="row.quality_status === 'PASS' ? 'success' : 'danger'">{{ row.quality_status === 'PASS' ? '合格' : '不合格' }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="remark" label="备注" min-width="120" />
          </el-table>
        </el-tab-pane>

        <el-tab-pane label="质量记录" name="quality">
          <el-table :data="detailData.quality_records || []" border size="small">
            <el-table-column prop="iqc_no" label="IQC单号" width="130" />
            <el-table-column prop="material_code" label="物料编码" width="110" />
            <el-table-column prop="inspect_date" label="检验日期" width="120" />
            <el-table-column prop="inspect_qty" label="检验数量" width="90" align="right" />
            <el-table-column prop="qualified_qty" label="合格数量" width="90" align="right" />
            <el-table-column prop="qualified_rate" label="合格率" width="90" align="right">
              <template #default="{ row }">
                {{ row.qualified_rate != null ? Number(row.qualified_rate).toFixed(1) + '%' : '-' }}
              </template>
            </el-table-column>
            <el-table-column prop="defect_count" label="不良数量" width="90" align="right" />
            <el-table-column prop="result" label="结果" width="80" align="center">
              <template #default="{ row }">
                <el-tag size="small" :type="row.result === 'PASS' ? 'success' : 'danger'">{{ row.result === 'PASS' ? '通过' : '不通过' }}</el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>

        <el-tab-pane label="月度趋势" name="trend">
          <div v-if="trendData.length > 0" class="trend-chart">
            <el-table :data="trendData" border size="small">
              <el-table-column prop="evaluation_month" label="月份" width="110" align="center" />
              <el-table-column prop="on_time_rate" label="准时率(%)" width="110" align="right">
                <template #default="{ row }">
                  {{ row.on_time_rate != null ? Number(row.on_time_rate).toFixed(1) + '%' : '-' }}
                </template>
              </el-table-column>
              <el-table-column prop="quality_rate" label="合格率(%)" width="110" align="right">
                <template #default="{ row }">
                  {{ row.quality_rate != null ? Number(row.quality_rate).toFixed(1) + '%' : '-' }}
                </template>
              </el-table-column>
              <el-table-column prop="total_score" label="综合得分" width="100" align="right">
                <template #default="{ row }">
                  {{ row.total_score != null ? Number(row.total_score).toFixed(1) : '-' }}
                </template>
              </el-table-column>
              <el-table-column prop="grade" label="评级" width="80" align="center">
                <template #default="{ row }">
                  <el-tag size="small" :type="getGradeType(row.grade)">{{ row.grade || '-' }}</el-tag>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <el-empty v-else description="暂无趋势数据" />
        </el-tab-pane>
      </el-tabs>
      <template #footer>
        <el-button @click="detailDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>

    <!-- 排名Dialog -->
    <el-dialog v-model="rankingDialogVisible" title="供应商绩效排名" width="700px">
      <el-table :data="rankingData">
        <el-table-column prop="rank" label="排名" width="70" align="center" />
        <el-table-column prop="supplier_code" label="供应商编码" width="120" />
        <el-table-column prop="supplier_name" label="供应商名称" min-width="160" />
        <el-table-column prop="total_score" label="综合得分" width="100" align="right">
          <template #default="{ row }">
            <span class="score-text">{{ row.total_score != null ? Number(row.total_score).toFixed(1) : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="grade" label="评级" width="80" align="center">
          <template #default="{ row }">
            <el-tag size="small" :type="getGradeType(row.grade)">{{ row.grade || '-' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="on_time_rate" label="准时率(%)" width="110" align="right">
          <template #default="{ row }">
            {{ row.on_time_rate != null ? Number(row.on_time_rate).toFixed(1) + '%' : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="quality_rate" label="合格率(%)" width="110" align="right">
          <template #default="{ row }">
            {{ row.quality_rate != null ? Number(row.quality_rate).toFixed(1) + '%' : '-' }}
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <el-button @click="rankingDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { getSupplierKPIList, getSupplierKPIMonthly, getSupplierKPIRanking } from '@/api/scp'

const loading = ref(false)
const tableData = ref<any[]>([])
const detailDialogVisible = ref(false)
const rankingDialogVisible = ref(false)
const rankingLoading = ref(false)
const detailTab = ref('score')

const searchForm = reactive({ supplier_name: '', evaluation_month: '', grade: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const detailData = reactive<any>({ supplier: null, delivery_records: [], quality_records: [] })
const trendData = ref<any[]>([])
const rankingData = ref<any[]>([])

const getRateClass = (rate: number | null) => {
  if (rate == null) return ''
  if (rate >= 95) return 'rate-excellent'
  if (rate >= 85) return 'rate-good'
  if (rate >= 70) return 'rate-warning'
  return 'rate-danger'
}

const getGradeType = (grade: string) => {
  const map: Record<string, string> = { A: 'success', B: 'primary', C: 'warning', D: 'danger', E: 'danger' }
  return map[grade] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getSupplierKPIList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.supplier_name = ''; searchForm.evaluation_month = ''; searchForm.grade = ''; handleSearch() }

const handleViewDetail = async (row: any) => {
  detailDialogVisible.value = true
  detailTab.value = 'score'
  Object.assign(detailData, { supplier: row, delivery_records: [], quality_records: [] })
  // 加载月度趋势
  if (row.supplier_id) {
    try {
      const res = await getSupplierKPIMonthly(row.supplier_id)
      trendData.value = res.data || []
    } catch {
      trendData.value = []
    }
  }
}

const handleRanking = async () => {
  rankingDialogVisible.value = true
  rankingLoading.value = true
  try {
    const res = await getSupplierKPIRanking()
    rankingData.value = res.data || []
  } finally {
    rankingLoading.value = false
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.supplier-kpi-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .detail-header { margin-bottom: 16px; }
  .score-text { color: #409eff; font-weight: 600; font-size: 15px; }
  .rate-excellent { color: #67c23a; font-weight: 500; }
  .rate-good { color: #409eff; }
  .rate-warning { color: #e6a23c; }
  .rate-danger { color: #f56c6c; }
}
</style>
