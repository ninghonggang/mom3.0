<template>
  <div class="dict-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="字典名称">
          <el-input v-model="searchForm.dict_name" placeholder="请输入字典名称" clearable />
        </el-form-item>
        <el-form-item label="字典类型">
          <el-input v-model="searchForm.dict_type" placeholder="请输入字典类型" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('system:dict:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="dict_name" label="字典名称" min-width="150" />
        <el-table-column prop="dict_type" label="字典类型" min-width="150" />
        <el-table-column prop="remark" label="备注" min-width="200" />
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleData(row)">字典数据</el-button>
            <el-button link type="primary" v-if="hasPermission('system:dict:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('system:dict:delete')" @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="字典名称" prop="dict_name">
          <el-input v-model="formData.dict_name" />
        </el-form-item>
        <el-form-item label="字典类型" prop="dict_type">
          <el-input v-model="formData.dict_type" :disabled="!!formData.id" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="dataDialogVisible" title="字典数据" width="800px">
      <div class="mb-4">
        <el-button type="primary" size="small" v-if="hasPermission('system:dict:edit')" @click="handleAddData">
          <el-icon><Plus /></el-icon>新增
        </el-button>
      </div>
      <el-table :data="dictDataList">
        <el-table-column prop="dict_label" label="标签" min-width="120" />
        <el-table-column prop="dict_value" label="值" min-width="120" />
        <el-table-column prop="sort" label="排序" width="80" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="{ row, $index }">
            <el-button link type="primary" size="small" @click="handleEditData(row, $index)">编辑</el-button>
            <el-button link type="danger" size="small" @click="handleDeleteData($index)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <el-button @click="dataDialogVisible = false">关闭</el-button>
        <el-button type="primary" @click="handleSaveData">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getDictTypeList, createDictType, updateDictType, deleteDictType, getDictDataList } from '@/api/system'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const dataDialogVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const currentDictType = ref('')
const dictDataList = ref<any[]>([])

const searchForm = reactive({ dict_name: '', dict_type: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const formData = reactive({ id: 0, dict_name: '', dict_type: '', remark: '' })

const rules: FormRules = {
  dict_name: [{ required: true, message: '请输入字典名称', trigger: 'blur' }],
  dict_type: [{ required: true, message: '请输入字典类型', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑字典' : '新增字典')

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDictTypeList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.dict_name = ''; searchForm.dict_type = ''; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, { id: 0, dict_name: '', dict_type: '', remark: '' })
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该字典吗？', '提示', { type: 'warning' })
    await deleteDictType(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    formData.id ? await updateDictType(formData.id, formData) : await createDictType(formData)
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } finally { submitLoading.value = false }
}

const handleData = async (row: any) => {
  currentDictType.value = row.dict_type
  const res = await getDictDataList(row.dict_type)
  dictDataList.value = res.data || []
  dataDialogVisible.value = true
}

const handleAddData = () => {
  dictDataList.value.push({ dict_label: '', dict_value: '', sort: 0, status: 1 })
}

const handleEditData = (row: any, _index: number) => {
  Object.assign(row, row)
}

const handleDeleteData = (index: number) => {
  dictDataList.value.splice(index, 1)
}

const handleSaveData = async () => {
  ElMessage.success('保存成功')
  dataDialogVisible.value = false
  loadData()
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.dict-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .mb-4 { margin-bottom: 16px; }
}
</style>
