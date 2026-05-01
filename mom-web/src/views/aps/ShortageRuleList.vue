<template>
  <div class="shortagerule-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="规则编码">
          <el-input v-model="searchForm.rule_code" placeholder="请输入编码" clearable />
        </el-form-item>
        <el-form-item label="规则名称">
          <el-input v-model="searchForm.rule_name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="物料类别">
          <el-select v-model="searchForm.material_category" placeholder="请选择" clearable>
            <el-option label="原材料" value="RAW" />
            <el-option label="半成品" value="SEMI" />
            <el-option label="成品" value="FINISHED" />
            <el-option label="包装" value="PACKAGING" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="激活" value="ACTIVE" />
            <el-option label="停用" value="INACTIVE" />
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
        <el-table-column prop="rule_code" label="规则编码" width="130" />
        <el-table-column prop="rule_name" label="规则名称" min-width="150" />
        <el-table-column prop="material_category" label="物料类别" width="100">
          <template #default="{ row }">
            {{ getCategoryText(row.material_category) }}
          </template>
        </el-table-column>
        <el-table-column prop="min_stock" label="最低库存" width="100" align="right" />
        <el-table-column prop="max_stock" label="最高库存" width="100" align="right" />
        <el-table-column prop="reorder_point" label="再订货点" width="100" align="right" />
        <el-table-column prop="supplier_lead_time" label="供应商交期(天)" width="120" align="right" />
        <el-table-column prop="priority" label="优先级" width="80" align="center">
          <template #default="{ row }">
            <el-tag size="small" :type="row.priority === 1 ? 'danger' : row.priority === 2 ? 'warning' : 'info'">
              {{ row.priority }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="row.status === 'ACTIVE' ? 'success' : 'info'">
              {{ row.status === 'ACTIVE' ? '激活' : '停用' }}
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
      <el-form ref="formRef" :model="form" :rules="rules" label-width="130px">
        <el-form-item label="规则编码" prop="rule_code">
          <el-input v-model="form.rule_code" placeholder="请输入编码" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="规则名称" prop="rule_name">
          <el-input v-model="form.rule_name" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="物料类别" prop="material_category">
          <el-select v-model="form.material_category" placeholder="请选择">
            <el-option label="原材料" value="RAW" />
            <el-option label="半成品" value="SEMI" />
            <el-option label="成品" value="FINISHED" />
            <el-option label="包装" value="PACKAGING" />
          </el-select>
        </el-form-item>
        <el-form-item label="最低库存" prop="min_stock">
          <el-input-number v-model="form.min_stock" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="最高库存" prop="max_stock">
          <el-input-number v-model="form.max_stock" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="再订货点" prop="reorder_point">
          <el-input-number v-model="form.reorder_point" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="供应商交期(天)" prop="supplier_lead_time">
          <el-input-number v-model="form.supplier_lead_time" :min="0" />
        </el-form-item>
        <el-form-item label="优先级" prop="priority">
          <el-input-number v-model="form.priority" :min="1" :max="10" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio label="ACTIVE">激活</el-radio>
            <el-radio label="INACTIVE">停用</el-radio>
          </el-radio-group>
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

interface ShortageRule {
  id: number
  rule_code: string
  rule_name: string
  material_category: string
  min_stock: number
  max_stock: number
  reorder_point: number
  supplier_lead_time: number
  priority: number
  status: string
}

let mockId = 100
const mockData: ShortageRule[] = [
  { id: 1, rule_code: 'SR-001', rule_name: '原材料紧急补货', material_category: 'RAW', min_stock: 1000, max_stock: 10000, reorder_point: 2000, supplier_lead_time: 7, priority: 1, status: 'ACTIVE' },
  { id: 2, rule_code: 'SR-002', rule_name: '半成品安全库存', material_category: 'SEMI', min_stock: 500, max_stock: 5000, reorder_point: 1000, supplier_lead_time: 14, priority: 2, status: 'ACTIVE' },
  { id: 3, rule_code: 'SR-003', rule_name: '成品缓冲策略', material_category: 'FINISHED', min_stock: 200, max_stock: 2000, reorder_point: 500, supplier_lead_time: 3, priority: 3, status: 'ACTIVE' },
  { id: 4, rule_code: 'SR-004', rule_name: '包装材料补货', material_category: 'PACKAGING', min_stock: 3000, max_stock: 20000, reorder_point: 5000, supplier_lead_time: 5, priority: 4, status: 'INACTIVE' },
  { id: 5, rule_code: 'SR-005', rule_name: '进口原材料策略', material_category: 'RAW', min_stock: 2000, max_stock: 15000, reorder_point: 5000, supplier_lead_time: 30, priority: 1, status: 'ACTIVE' },
  { id: 6, rule_code: 'SR-006', rule_name: '关键件备货', material_category: 'SEMI', min_stock: 100, max_stock: 1000, reorder_point: 300, supplier_lead_time: 10, priority: 1, status: 'ACTIVE' },
  { id: 7, rule_code: 'SR-007', rule_name: '季节性备货', material_category: 'FINISHED', min_stock: 500, max_stock: 5000, reorder_point: 1500, supplier_lead_time: 7, priority: 2, status: 'ACTIVE' },
  { id: 8, rule_code: 'SR-008', rule_name: '低值易耗品', material_category: 'PACKAGING', min_stock: 5000, max_stock: 30000, reorder_point: 8000, supplier_lead_time: 3, priority: 5, status: 'ACTIVE' },
  { id: 9, rule_code: 'SR-009', rule_name: '出口产品规则', material_category: 'FINISHED', min_stock: 100, max_stock: 3000, reorder_point: 400, supplier_lead_time: 5, priority: 2, status: 'INACTIVE' },
  { id: 10, rule_code: 'SR-010', rule_name: '化学品管控', material_category: 'RAW', min_stock: 200, max_stock: 2000, reorder_point: 500, supplier_lead_time: 21, priority: 1, status: 'ACTIVE' },
  { id: 11, rule_code: 'SR-011', rule_name: '电子元器件策略', material_category: 'SEMI', min_stock: 1000, max_stock: 8000, reorder_point: 2000, supplier_lead_time: 14, priority: 2, status: 'ACTIVE' },
  { id: 12, rule_code: 'SR-012', rule_name: '医疗器械备货', material_category: 'FINISHED', min_stock: 50, max_stock: 500, reorder_point: 100, supplier_lead_time: 15, priority: 1, status: 'ACTIVE' },
]

const loading = ref(false)
const tableData = ref<ShortageRule[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('新增缺料规则')
const isEdit = ref(false)
const formRef = ref()

const searchForm = reactive({ rule_code: '', rule_name: '', material_category: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const form = reactive({
  id: 0,
  rule_code: '',
  rule_name: '',
  material_category: 'RAW',
  min_stock: 0,
  max_stock: 0,
  reorder_point: 0,
  supplier_lead_time: 0,
  priority: 1,
  status: 'ACTIVE'
})

const rules = {
  rule_code: [{ required: true, message: '请输入规则编码', trigger: 'blur' }],
  rule_name: [{ required: true, message: '请输入规则名称', trigger: 'blur' }],
  material_category: [{ required: true, message: '请选择物料类别', trigger: 'change' }]
}

const getCategoryText = (cat: string) => {
  const map: Record<string, string> = { RAW: '原材料', SEMI: '半成品', FINISHED: '成品', PACKAGING: '包装' }
  return map[cat] || cat
}

const loadData = () => {
  loading.value = true
  try {
    let filtered = mockData.filter(item => {
      if (searchForm.rule_code && !item.rule_code.toLowerCase().includes(searchForm.rule_code.toLowerCase())) return false
      if (searchForm.rule_name && !item.rule_name.includes(searchForm.rule_name)) return false
      if (searchForm.material_category && item.material_category !== searchForm.material_category) return false
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
const handleReset = () => { searchForm.rule_code = ''; searchForm.rule_name = ''; searchForm.material_category = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增缺料规则'
  Object.assign(form, { id: 0, rule_code: '', rule_name: '', material_category: 'RAW', min_stock: 0, max_stock: 0, reorder_point: 0, supplier_lead_time: 0, priority: 1, status: 'ACTIVE' })
  dialogVisible.value = true
}

const handleEdit = (row: ShortageRule) => {
  isEdit.value = true
  dialogTitle.value = '编辑缺料规则'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleSubmit = () => {
  formRef.value?.validate((valid: boolean) => {
    if (!valid) return
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

const handleDelete = async (row: ShortageRule) => {
  try {
    await ElMessageBox.confirm('确定删除该缺料规则吗？', '提示', { type: 'warning' })
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
.shortagerule-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
