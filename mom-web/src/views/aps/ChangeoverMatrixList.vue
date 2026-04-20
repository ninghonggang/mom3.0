<template>
  <div class="page-container">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="起始产品">
          <el-input v-model="searchForm.from_product_code" placeholder="起始产品编码" clearable />
        </el-form-item>
        <el-form-item label="目标产品">
          <el-input v-model="searchForm.to_product_code" placeholder="目标产品编码" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('aps:changeover-matrix:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
      <el-button @click="loadData">
        <el-icon><Refresh /></el-icon>刷新
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="from_product_code" label="起始产品编码" width="130" />
        <el-table-column prop="from_product_name" label="起始产品名称" min-width="150" />
        <el-table-column prop="to_product_code" label="目标产品编码" width="130" />
        <el-table-column prop="to_product_name" label="目标产品名称" min-width="150" />
        <el-table-column prop="changeover_time" label="换型时间(分钟)" width="120" />
        <el-table-column prop="setup_time" label="准备时间(分钟)" width="120" />
        <el-table-column prop="clean_time" label="清洁时间(分钟)" width="120" />
        <el-table-column prop="changeover_type" label="换型类型" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.changeover_type === 'CLEAN'" type="warning">清洁</el-tag>
            <el-tag v-else-if="row.changeover_type === 'RETOOL'" type="danger">换模</el-tag>
            <el-tag v-else-if="row.changeover_type === 'ADJUST'" type="info">调整</el-tag>
            <el-tag v-else type="info">其他</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="特殊换型" width="120">
          <template #default="{ row }">
            <span v-if="row.is_color_change === 1" class="tag-warning">颜色换型</span>
            <span v-if="row.is_material_change === 1" class="tag-info">物料换型</span>
            <span v-if="row.is_color_change !== 1 && row.is_material_change !== 1">-</span>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="120" show-overflow-tooltip />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('aps:changeover-matrix:edit')" @click.stop="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('aps:changeover-matrix:delete')" @click.stop="handleDelete(row)">删除</el-button>
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
      <el-form :model="formData" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="起始产品ID">
              <el-input-number v-model="formData.from_product_id" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="起始产品编码" required>
              <el-input v-model="formData.from_product_code" placeholder="请输入" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="起始产品名称">
          <el-input v-model="formData.from_product_name" placeholder="请输入" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="目标产品ID">
              <el-input-number v-model="formData.to_product_id" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目标产品编码" required>
              <el-input v-model="formData.to_product_code" placeholder="请输入" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="目标产品名称">
          <el-input v-model="formData.to_product_name" placeholder="请输入" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="换型时间(分钟)">
              <el-input-number v-model="formData.changeover_time" :min="0" :precision="2" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="准备时间(分钟)">
              <el-input-number v-model="formData.setup_time" :min="0" :precision="2" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="清洁时间(分钟)">
              <el-input-number v-model="formData.clean_time" :min="0" :precision="2" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="换型类型">
          <el-select v-model="formData.changeover_type" placeholder="请选择" style="width: 100%">
            <el-option label="清洁" value="CLEAN" />
            <el-option label="换模" value="RETOOL" />
            <el-option label="调整" value="ADJUST" />
            <el-option label="其他" value="OTHER" />
          </el-select>
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="颜色换型">
              <el-switch v-model="formData.is_color_change" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="物料换型">
              <el-switch v-model="formData.is_material_change" :active-value="1" :inactive-value="0" />
            </el-form-item>
          </el-col>
        </el-row>
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
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getChangeoverMatrixList,
  createChangeoverMatrix,
  updateChangeoverMatrix,
  deleteChangeoverMatrix
} from '@/api/aps'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const dialogTitle = ref('新增换型矩阵')
const tableData = ref<any[]>([])
const isEdit = ref(false)
const currentId = ref<number | null>(null)

const searchForm = reactive({
  from_product_code: '',
  to_product_code: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = ref<any>({
  from_product_id: 0,
  from_product_code: '',
  from_product_name: '',
  to_product_id: 0,
  to_product_code: '',
  to_product_name: '',
  changeover_time: 0,
  setup_time: 0,
  clean_time: 0,
  changeover_type: 'CLEAN',
  is_color_change: 0,
  is_material_change: 0,
  remark: ''
})

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      ...searchForm,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getChangeoverMatrixList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.from_product_code = ''
  searchForm.to_product_code = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  currentId.value = null
  dialogTitle.value = '新增换型矩阵'
  formData.value = {
    from_product_id: 0,
    from_product_code: '',
    from_product_name: '',
    to_product_id: 0,
    to_product_code: '',
    to_product_name: '',
    changeover_time: 0,
    setup_time: 0,
    clean_time: 0,
    changeover_type: 'CLEAN',
    is_color_change: 0,
    is_material_change: 0,
    remark: ''
  }
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  isEdit.value = true
  currentId.value = row.id
  dialogTitle.value = '编辑换型矩阵'
  formData.value = { ...row }
  dialogVisible.value = true
}

const handleSave = async () => {
  if (!formData.value.from_product_code || !formData.value.to_product_code) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    if (isEdit.value && currentId.value) {
      await updateChangeoverMatrix(currentId.value, formData.value)
      ElMessage.success('更新成功')
    } else {
      await createChangeoverMatrix(formData.value)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该换型矩阵记录吗？', '提示', { type: 'warning' })
    await deleteChangeoverMatrix(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped lang="scss">
.page-container {
  height: 100%;
}

.search-card {
  margin-bottom: 16px;
}

.toolbar-card {
  margin-bottom: 16px;

  :deep(.el-card__body) {
    padding: 12px 16px;
    display: flex;
    gap: 12px;
  }
}

.pagination {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}

.tag-warning {
  color: #faad14;
  font-size: 12px;
}

.tag-info {
  color: #1890ff;
  font-size: 12px;
  margin-left: 4px;
}
</style>
