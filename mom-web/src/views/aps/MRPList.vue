<template>
  <div class="mrp-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="MRP编号">
          <el-input v-model="searchForm.mrp_no" placeholder="请输入MRP编号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待计算" :value="1" />
            <el-option label="计算中" :value="2" />
            <el-option label="已完成" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="dialogVisible = true; formData = {}">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="mrp_no" label="MRP编号" width="140" />
        <el-table-column prop="mrp_type" label="类型" width="80" />
        <el-table-column prop="plan_date" label="计划日期" width="110" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="120" show-overflow-tooltip />
        <el-table-column prop="created_by" label="创建人" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="success" size="small" @click="handleCalculate(row)" v-if="row.status === 1">计算</el-button>
            <el-button link type="primary" size="small" @click="handleDetail(row)">明细</el-button>
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

    <!-- 新增对话框 -->
    <el-dialog v-model="dialogVisible" title="新增MRP" width="500px">
      <el-form :model="formData" label-width="100px">
        <el-form-item label="MRP编号">
          <el-input v-model="formData.mrp_no" placeholder="自动生成" disabled />
        </el-form-item>
        <el-form-item label="类型" required>
          <el-select v-model="formData.mrp_type" placeholder="请选择" style="width: 100%">
            <el-option label="MPS" value="MPS" />
            <el-option label="MRP" value="MRP" />
          </el-select>
        </el-form-item>
        <el-form-item label="计划日期">
          <el-date-picker v-model="formData.plan_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saveLoading" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getMRPList, calculateMRP } from '@/api/aps'

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const tableData = ref<any[]>([])
const formData = ref<any>({})

const searchForm = reactive({ mrp_no: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待计算', 2: '计算中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getMRPList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.mrp_no = ''; searchForm.status = ''; handleSearch() }

const handleSave = async () => {
  ElMessage.info('MRP创建需要后端支持明细录入')
  dialogVisible.value = false
}

const handleCalculate = async (row: any) => {
  await calculateMRP(row.id)
  ElMessage.success('计算完成')
  loadData()
}

const handleDetail = (row: any) => { ElMessage.info('查看MRP明细') }

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.mrp-list {
  .search-card { margin-bottom: 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
