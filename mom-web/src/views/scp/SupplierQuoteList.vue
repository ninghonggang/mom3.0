<template>
  <div class="supplier-quote-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="询价单号">
          <el-input v-model="searchForm.rfq_no" placeholder="请输入询价单号" clearable />
        </el-form-item>
        <el-form-item label="供应商名称">
          <el-input v-model="searchForm.supplier_name" placeholder="请输入供应商名称" clearable />
        </el-form-item>
        <el-form-item label="报价状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="已提交" value="SUBMITTED" />
            <el-option label="已撤回" value="WITHDRAWN" />
            <el-option label="已中标" value="AWARDED" />
            <el-option label="未中标" value="LOST" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData" @row-click="handleRowClick">
        <el-table-column prop="quote_no" label="报价单号" width="150" />
        <el-table-column prop="rfq_no" label="询价单号" width="150" />
        <el-table-column prop="supplier_name" label="供应商名称" min-width="180" />
        <el-table-column prop="contact_person" label="联系人" width="100" />
        <el-table-column prop="contact_phone" label="联系电话" width="130" />
        <el-table-column prop="quote_date" label="报价日期" width="120" />
        <el-table-column prop="total_amount" label="报价总额" width="120" align="right">
          <template #default="{ row }">
            <span class="amount-text">{{ row.total_amount ? Number(row.total_amount).toFixed(2) : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="valid_until" label="有效期至" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag size="small" :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160" />
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

    <!-- 报价详情 & 比价分析 -->
    <el-card v-if="selectedQuote" class="comparison-card">
      <template #header>
        <div class="card-header">
          <span>报价详情 - {{ selectedQuote.quote_no }}</span>
          <el-button size="small" @click="selectedQuote = null">关闭详情</el-button>
        </div>
      </template>

      <el-tabs v-model="activeTab">
        <el-tab-pane label="报价明细" name="items">
          <el-table :data="selectedQuote.items || []" border size="small">
            <el-table-column prop="material_code" label="物料编码" width="120" />
            <el-table-column prop="material_name" label="物料名称" min-width="150" />
            <el-table-column prop="specification" label="规格型号" width="140" />
            <el-table-column prop="unit" label="单位" width="70" align="center" />
            <el-table-column prop="quoted_qty" label="报价数量" width="100" align="right" />
            <el-table-column prop="unit_price" label="单价" width="100" align="right">
              <template #default="{ row }">
                {{ row.unit_price ? Number(row.unit_price).toFixed(4) : '-' }}
              </template>
            </el-table-column>
            <el-table-column prop="lead_time_days" label="交期(天)" width="90" align="center" />
            <el-table-column prop="remark" label="备注" min-width="120" />
          </el-table>
        </el-tab-pane>

        <el-tab-pane label="比价分析" name="compare">
          <el-alert v-if="comparisonData.length === 0" type="info" :closable="false">
            暂无其他供应商报价可对比
          </el-alert>
          <el-table v-else :data="comparisonData" border size="small">
            <el-table-column prop="supplier_name" label="供应商" min-width="150" />
            <el-table-column v-for="col in comparisonColumns" :key="col.prop" :prop="col.prop" :label="col.label" :width="col.width" align="right">
              <template #default="{ row }">
                <span :class="{ 'best-price': col.isBest && row[col.prop] === col.bestValue }">
                  {{ row[col.prop] !== null && row[col.prop] !== undefined ? (typeof row[col.prop] === 'number' ? Number(row[col.prop]).toFixed(4) : row[col.prop]) : '-' }}
                </span>
              </template>
            </el-table-column>
          </el-table>
          <div v-if="comparisonData.length > 0" class="comparison-legend">
            <span class="best-price-indicator">* 最低价</span>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { getSupplierQuoteList, getRFQQuotes } from '@/api/scp'

const loading = ref(false)
const tableData = ref<any[]>([])
const selectedQuote = ref<any>(null)
const activeTab = ref('items')
const comparisonData = ref<any[]>([])
const comparisonColumns = ref<any[]>([])

const searchForm = reactive({ rfq_no: '', supplier_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: string) => {
  const map: Record<string, string> = { SUBMITTED: '已提交', WITHDRAWN: '已撤回', AWARDED: '已中标', LOST: '未中标' }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = { SUBMITTED: 'success', WITHDRAWN: 'info', AWARDED: 'success', LOST: 'danger' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getSupplierQuoteList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.rfq_no = ''; searchForm.supplier_name = ''; searchForm.status = ''; handleSearch() }

const handleRowClick = async (row: any) => {
  selectedQuote.value = row
  activeTab.value = 'items'
  // 加载比价数据
  if (row.rfq_id) {
    try {
      const res = await getRFQQuotes(row.rfq_id)
      const quotes = res.data || []
      buildComparison(quotes, row)
    } catch {
      comparisonData.value = []
    }
  }
}

const buildComparison = (quotes: any[], currentQuote: any) => {
  if (quotes.length < 2) {
    comparisonData.value = []
    comparisonColumns.value = []
    return
  }

  // 收集所有物料行
  const materialMap = new Map<string, any>()
  quotes.forEach(q => {
    ;(q.items || []).forEach((item: any) => {
      if (!materialMap.has(item.material_code)) {
        materialMap.set(item.material_code, {
          material_code: item.material_code,
          material_name: item.material_name,
          specification: item.specification
        })
      }
    })
  })

  // 构建对比数据
  const rows = quotes.map(q => {
    const row: any = { supplier_name: q.supplier_name }
    ;(q.items || []).forEach((item: any) => {
      row[`price_${item.material_code}`] = item.unit_price
      row[`total_${item.material_code}`] = item.unit_price * item.quoted_qty
    })
    row.total_amount = q.total_amount
    return row
  })

  // 构建列
  const materials = Array.from(materialMap.values())
  const cols: any[] = []
  materials.forEach(m => {
    cols.push({ prop: `price_${m.material_code}`, label: `${m.material_code}单价`, width: 110, isBest: true, bestValue: null })
    cols.push({ prop: `total_${m.material_code}`, label: `${m.material_code}小计`, width: 110, isBest: true, bestValue: null })
  })
  cols.push({ prop: 'total_amount', label: '报价总额', width: 120, isBest: true, bestValue: null })

  // 计算最低价
  cols.forEach(col => {
    if (col.isBest) {
      const values = rows.map(r => r[col.prop]).filter(v => v !== null && v !== undefined)
      col.bestValue = values.length > 0 ? Math.min(...values) : null
    }
  })

  comparisonColumns.value = cols
  comparisonData.value = rows
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.supplier-quote-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .comparison-card { margin-top: 16px; }
  .card-header { display: flex; justify-content: space-between; align-items: center; }
  .amount-text { color: #409eff; font-weight: 500; }
  .best-price { color: #67c23a; font-weight: bold; }
  .comparison-legend { margin-top: 8px; font-size: 12px; color: #909399; }
  .best-price-indicator::before { content: ''; display: inline-block; width: 8px; height: 8px; background: #67c23a; border-radius: 50%; margin-right: 4px; }
}
</style>
