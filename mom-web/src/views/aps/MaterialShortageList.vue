<template>
  <div class="materialshortage-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="物料编码">
          <el-input v-model="searchForm.material_code" placeholder="请输入编码" clearable />
        </el-form-item>
        <el-form-item label="物料名称">
          <el-input v-model="searchForm.material_name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="紧缺" value="SHORTAGE" />
            <el-option label="正常" value="NORMAL" />
            <el-option label="过剩" value="SURPLUS" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="material_code" label="物料编码" width="130" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="required_qty" label="需求数量" width="100" align="right" />
        <el-table-column prop="available_qty" label="可用数量" width="100" align="right" />
        <el-table-column prop="shortage_qty" label="缺口数量" width="100" align="right">
          <template #default="{ row }">
            <span :class="row.shortage_qty > 0 ? 'text-danger' : ''">{{ row.shortage_qty }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="shortage_ratio" label="缺口比例(%)" width="110" align="right">
          <template #default="{ row }">
            <span :class="row.shortage_ratio > 0 ? 'text-danger' : ''">{{ row.shortage_ratio.toFixed(1) }}%</span>
          </template>
        </el-table-column>
        <el-table-column prop="plan_date" label="计划日期" width="120" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新增/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="物料编码" prop="material_code">
          <el-input v-model="form.material_code" placeholder="请输入编码" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="物料名称" prop="material_name">
          <el-input v-model="form.material_name" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="需求数量" prop="required_qty">
          <el-input-number v-model="form.required_qty" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="可用数量" prop="available_qty">
          <el-input-number v-model="form.available_qty" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="计划日期" prop="plan_date">
          <el-date-picker v-model="form.plan_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="form.status" placeholder="请选择">
            <el-option label="紧缺" value="SHORTAGE" />
            <el-option label="正常" value="NORMAL" />
            <el-option label="过剩" value="SURPLUS" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'

interface MaterialShortage {
  id: number
  material_code: string
  material_name: string
  required_qty: number
  available_qty: number
  shortage_qty: number
  shortage_ratio: number
  plan_date: string
  status: string
}

let mockId = 100
const mockData: MaterialShortage[] = [
  { id: 1, material_code: 'MAT-1001', material_name: '钢板 Q235', required_qty: 5000, available_qty: 3200, shortage_qty: 1800, shortage_ratio: 36.0, plan_date: '2025-02-01', status: 'SHORTAGE' },
  { id: 2, material_code: 'MAT-1002', material_name: '铝合金板材', required_qty: 3000, available_qty: 3000, shortage_qty: 0, shortage_ratio: 0.0, plan_date: '2025-02-01', status: 'NORMAL' },
  { id: 3, material_code: 'MAT-1003', material_name: '铜线 2.5mm', required_qty: 1500, available_qty: 800, shortage_qty: 700, shortage_ratio: 46.7, plan_date: '2025-02-03', status: 'SHORTAGE' },
  { id: 4, material_code: 'MAT-1004', material_name: '塑料粒子 ABS', required_qty: 8000, available_qty: 9500, shortage_qty: 0, shortage_ratio: 0.0, plan_date: '2025-02-02', status: 'SURPLUS' },
  { id: 5, material_code: 'MAT-1005', material_name: '轴承 6205', required_qty: 2000, available_qty: 1500, shortage_qty: 500, shortage_ratio: 25.0, plan_date: '2025-02-04', status: 'SHORTAGE' },
  { id: 6, material_code: 'MAT-1006', material_name: '电机 Y2-132', required_qty: 300, available_qty: 300, shortage_qty: 0, shortage_ratio: 0.0, plan_date: '2025-02-05', status: 'NORMAL' },
  { id: 7, material_code: 'MAT-1007', material_name: '减速机 WPO50', required_qty: 150, available_qty: 80, shortage_qty: 70, shortage_ratio: 46.7, plan_date: '2025-02-06', status: 'SHORTAGE' },
  { id: 8, material_code: 'MAT-1008', material_name: '液压油 HLP46', required_qty: 1000, available_qty: 1200, shortage_qty: 0, shortage_ratio: 0.0, plan_date: '2025-02-07', status: 'SURPLUS' },
  { id: 9, material_code: 'MAT-1009', material_name: '密封圈 NBR70', required_qty: 5000, available_qty: 4200, shortage_qty: 800, shortage_ratio: 16.0, plan_date: '2025-02-08', status: 'SHORTAGE' },
  { id: 10, material_code: 'MAT-1010', material_name: '电线 RV2.5', required_qty: 10000, available_qty: 10000, shortage_qty: 0, shortage_ratio: 0.0, plan_date: '2025-02-09', status: 'NORMAL' },
  { id: 11, material_code: 'MAT-1011', material_name: '不锈钢管 304', required_qty: 2500, available_qty: 1800, shortage_qty: 700, shortage_ratio: 28.0, plan_date: '2025-02-10', status: 'SHORTAGE' },
  { id: 12, material_code: 'MAT-1012', material_name: '焊接材料 J507', required_qty: 3000, available_qty: 3000, shortage_qty: 0, shortage_ratio: 0.0, plan_date: '2025-02-11', status: 'NORMAL' },
]

const loading = ref(false)
const tableData = ref<MaterialShortage[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('新增缺料分析')
const isEdit = ref(false)
const formRef = ref()

const searchForm = reactive({ material_code: '', material_name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const form = reactive({
  id: 0,
  material_code: '',
  material_name: '',
  required_qty: 0,
  available_qty: 0,
  shortage_qty: 0,
  shortage_ratio: 0,
  plan_date: '',
  status: 'SHORTAGE'
})

const rules = {
  material_code: [{ required: true, message: '请输入物料编码', trigger: 'blur' }],
  material_name: [{ required: true, message: '请输入物料名称', trigger: 'blur' }],
  required_qty: [{ required: true, message: '请输入需求数量', trigger: 'blur' }],
  available_qty: [{ required: true, message: '请输入可用数量', trigger: 'blur' }]
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = { SHORTAGE: 'danger', NORMAL: 'success', SURPLUS: 'warning' }
  return map[status] || 'info'
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = { SHORTAGE: '紧缺', NORMAL: '正常', SURPLUS: '过剩' }
  return map[status] || status
}

const calcShortage = (row: typeof form) => {
  row.shortage_qty = Math.max(0, row.required_qty - row.available_qty)
  row.shortage_ratio = row.required_qty > 0 ? (row.shortage_qty / row.required_qty) * 100 : 0
}

const loadData = () => {
  loading.value = true
  try {
    let filtered = mockData.filter(item => {
      if (searchForm.material_code && !item.material_code.toLowerCase().includes(searchForm.material_code.toLowerCase())) return false
      if (searchForm.material_name && !item.material_name.includes(searchForm.material_name)) return false
      if (searchForm.status && item.status !== searchForm.status) return false
      return true
    })
    pagination.total = filtered.length
    const start = (pagination.page - 1) * pagination.pageSize
    tableData.value = filtered.slice(start, start + pagination.pageSize)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.material_code = ''; searchForm.material_name = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增缺料分析'
  Object.assign(form, { id: 0, material_code: '', material_name: '', required_qty: 0, available_qty: 0, shortage_qty: 0, shortage_ratio: 0, plan_date: '', status: 'SHORTAGE' })
  dialogVisible.value = true
}

const handleEdit = (row: MaterialShortage) => {
  isEdit.value = true
  dialogTitle.value = '编辑缺料分析'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleSubmit = () => {
  formRef.value?.validate((valid: boolean) => {
    if (!valid) return
    calcShortage(form)
    if (isEdit.value) {
      const idx = mockData.findIndex(item => item.id === form.id)
      if (idx !== -1) {
        mockData[idx] = { ...form }
        ElMessage.success('更新成功')
      }
    } else {
      mockData.unshift({ ...form, id: ++mockId })
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  })
}

const handleDelete = async (row: MaterialShortage) => {
  try {
    await ElMessageBox.confirm('确定删除该缺料分析记录吗？', '提示', { type: 'warning' })
    const idx = mockData.findIndex(item => item.id === row.id)
    if (idx !== -1) mockData.splice(idx, 1)
    ElMessage.success('删除成功')
    loadData()
  } catch {
    // cancel
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.materialshortage-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .text-danger { color: #f56c6c; font-weight: 600; }
}
</style>
