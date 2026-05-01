<template>
  <div class="productfamily-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="产品族编码">
          <el-input v-model="searchForm.code" placeholder="请输入编码" clearable />
        </el-form-item>
        <el-form-item label="产品族名称">
          <el-input v-model="searchForm.name" placeholder="请输入名称" clearable />
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
        <el-table-column prop="code" label="产品族编码" width="140" />
        <el-table-column prop="name" label="产品族名称" min-width="150" />
        <el-table-column prop="description" label="描述" min-width="180" show-overflow-tooltip />
        <el-table-column prop="product_count" label="产品数量" width="100" align="center" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="row.status === 'ACTIVE' ? 'success' : 'info'">
              {{ row.status === 'ACTIVE' ? '激活' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160" />
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
        <el-form-item label="产品族编码" prop="code">
          <el-input v-model="form.code" placeholder="请输入编码" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="产品族名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入描述" />
        </el-form-item>
        <el-form-item label="产品数量" prop="product_count">
          <el-input-number v-model="form.product_count" :min="0" />
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

interface ProductFamily {
  id: number
  code: string
  name: string
  description: string
  product_count: number
  status: string
  created_at: string
}

// Mock data store
let mockId = 100
const mockData: ProductFamily[] = [
  { id: 1, code: 'PF-001', name: '电子产品族', description: '涵盖所有消费电子产品线', product_count: 56, status: 'ACTIVE', created_at: '2025-01-15 10:30:00' },
  { id: 2, code: 'PF-002', name: '机械产品族', description: '工业机械及配件', product_count: 32, status: 'ACTIVE', created_at: '2025-01-16 14:20:00' },
  { id: 3, code: 'PF-003', name: '化工产品族', description: '化工原料及成品', product_count: 18, status: 'INACTIVE', created_at: '2025-01-17 09:00:00' },
  { id: 4, code: 'PF-004', name: '食品产品族', description: '食品加工与包装', product_count: 45, status: 'ACTIVE', created_at: '2025-01-18 11:45:00' },
  { id: 5, code: 'PF-005', name: '纺织产品族', description: '纺织品与服装', product_count: 27, status: 'ACTIVE', created_at: '2025-01-19 16:00:00' },
  { id: 6, code: 'PF-006', name: '汽车配件族', description: '汽车零配件生产', product_count: 63, status: 'ACTIVE', created_at: '2025-01-20 08:30:00' },
  { id: 7, code: 'PF-007', name: '医药产品族', description: '药品及医疗器械', product_count: 14, status: 'INACTIVE', created_at: '2025-01-21 10:00:00' },
  { id: 8, code: 'PF-008', name: '家具产品族', description: '家具制造与销售', product_count: 39, status: 'ACTIVE', created_at: '2025-01-22 13:20:00' },
  { id: 9, code: 'PF-009', name: '塑胶产品族', description: '塑胶制品加工', product_count: 22, status: 'ACTIVE', created_at: '2025-01-23 15:40:00' },
  { id: 10, code: 'PF-010', name: '五金产品族', description: '五金工具与配件', product_count: 48, status: 'ACTIVE', created_at: '2025-01-24 09:15:00' },
]

const loading = ref(false)
const tableData = ref<ProductFamily[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('新增产品族')
const isEdit = ref(false)
const formRef = ref()

const searchForm = reactive({ code: '', name: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const form = reactive({
  id: 0,
  code: '',
  name: '',
  description: '',
  product_count: 0,
  status: 'ACTIVE'
})

const rules = {
  code: [{ required: true, message: '请输入产品族编码', trigger: 'blur' }],
  name: [{ required: true, message: '请输入产品族名称', trigger: 'blur' }]
}

const loadData = () => {
  loading.value = true
  try {
    let filtered = mockData.filter(item => {
      if (searchForm.code && !item.code.toLowerCase().includes(searchForm.code.toLowerCase())) return false
      if (searchForm.name && !item.name.toLowerCase().includes(searchForm.name)) return false
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
const handleReset = () => { searchForm.code = ''; searchForm.name = ''; searchForm.status = ''; handleSearch() }

const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增产品族'
  Object.assign(form, { id: 0, code: '', name: '', description: '', product_count: 0, status: 'ACTIVE' })
  dialogVisible.value = true
}

const handleEdit = (row: ProductFamily) => {
  isEdit.value = true
  dialogTitle.value = '编辑产品族'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleSubmit = () => {
  formRef.value?.validate((valid: boolean) => {
    if (!valid) return
    if (isEdit.value) {
      const idx = mockData.findIndex(item => item.id === form.id)
      if (idx !== -1) {
        const now = new Date().toLocaleString('zh-CN')
        mockData[idx] = { ...form, created_at: mockData[idx].created_at }
        ElMessage.success('更新成功')
      }
    } else {
      const now = new Date().toLocaleString('zh-CN')
      mockData.unshift({ ...form, id: ++mockId, created_at: now })
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  })
}

const handleDelete = async (row: ProductFamily) => {
  try {
    await ElMessageBox.confirm('确定删除该产品族吗？', '提示', { type: 'warning' })
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
.productfamily-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
