<template>
  <div class="first-last-inspect-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="检验类型">
          <el-select v-model="searchForm.inspect_type" placeholder="请选择" clearable>
            <el-option label="首件" value="FIRST" />
            <el-option label="末件" value="LAST" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="待检" value="PENDING" />
            <el-option label="已完成" value="COMPLETED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('production:first-last-inspect:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增检验
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="inspect_no" label="检验单号" min-width="150" />
        <el-table-column prop="inspect_type" label="类型" width="80">
          <template #default="{ row }">
            <el-tag :type="row.inspect_type === 'FIRST' ? 'success' : 'warning'">
              {{ row.inspect_type === 'FIRST' ? '首件' : '末件' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="production_order_id" label="工单ID" width="100" />
        <el-table-column prop="process_id" label="工序ID" width="80" />
        <el-table-column prop="workstation_id" label="工位ID" width="80" />
        <el-table-column prop="inspector_name" label="检验员" width="100" />
        <el-table-column prop="overall_result" label="结果" width="80">
          <template #default="{ row }">
            <el-tag :type="row.overall_result === 'OK' ? 'success' : row.overall_result === 'NG' ? 'danger' : 'info'">
              {{ row.overall_result === 'OK' ? '合格' : row.overall_result === 'NG' ? '不合格' : '待检' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'COMPLETED' ? 'success' : 'info'">
              {{ row.status === 'COMPLETED' ? '已完成' : '待检' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('production:first-last-inspect:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" v-if="hasPermission('production:first-last-inspect:delete')" @click="handleDelete(row)">删除</el-button>
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getFirstLastInspectList, deleteFirstLastInspect } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const searchForm = reactive({ order_no: '', inspect_type: '', status: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const loadData = async () => {
  loading.value = true
  try {
    const params: any = { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    Object.keys(params).forEach(k => params[k] === '' && delete params[k])
    const res = await getFirstLastInspectList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.order_no = ''; searchForm.inspect_type = ''; searchForm.status = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增检验功能') }
const handleEdit = (row: any) => { ElMessage.info('编辑检验: ' + row.inspect_no) }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该检验单吗？', '提示', { type: 'warning' })
    await deleteFirstLastInspect(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.first-last-inspect-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>