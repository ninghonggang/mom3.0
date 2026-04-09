<template>
  <div class="order-change-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="工单号">
          <el-input v-model="searchForm.order_no" placeholder="请输入工单号" clearable />
        </el-form-item>
        <el-form-item label="变更类型">
          <el-select v-model="searchForm.change_type" placeholder="请选择" clearable>
            <el-option label="数量变更" value="QUANTITY_CHANGE" />
            <el-option label="日期变更" value="DATE_CHANGE" />
            <el-option label="优先级变更" value="PRIORITY_CHANGE" />
            <el-option label="产线变更" value="LINE_CHANGE" />
            <el-option label="状态变更" value="STATUS_CHANGE" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="order_no" label="工单号" width="140" />
        <el-table-column prop="change_type" label="变更类型" width="120">
          <template #default="{ row }">
            <el-tag :type="getChangeTypeTag(row.change_type)">{{ getChangeTypeText(row.change_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="old_value" label="旧值" min-width="120" />
        <el-table-column prop="new_value" label="新值" min-width="120" />
        <el-table-column prop="change_reason" label="变更原因" width="150" />
        <el-table-column prop="changed_by" label="操作人" width="100" />
        <el-table-column prop="created_at" label="变更时间" width="160" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('production:orderchange:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="warning" size="small" v-if="hasPermission('production:orderchange:approve')" @click="handleApprove(row)">审批</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('production:orderchange:delete')" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="detailVisible" title="变更详情" width="500px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="工单号">{{ detailData.order_no }}</el-descriptions-item>
        <el-descriptions-item label="变更类型">
          <el-tag :type="getChangeTypeTag(detailData.change_type)">{{ getChangeTypeText(detailData.change_type) }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="旧值" :span="2">{{ detailData.old_value }}</el-descriptions-item>
        <el-descriptions-item label="新值" :span="2">{{ detailData.new_value }}</el-descriptions-item>
        <el-descriptions-item label="变更原因" :span="2">{{ detailData.change_reason || '-' }}</el-descriptions-item>
        <el-descriptions-item label="操作人">{{ detailData.changed_by }}</el-descriptions-item>
        <el-descriptions-item label="变更时间">{{ detailData.created_at }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getOrderChangeList } from '@/api/production'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const detailVisible = ref(false)
const detailData = ref<any>({})

const searchForm = reactive({ order_no: '', change_type: '' })

const getChangeTypeText = (type: string) => {
  const map: Record<string, string> = {
    'QUANTITY_CHANGE': '数量变更',
    'DATE_CHANGE': '日期变更',
    'PRIORITY_CHANGE': '优先级变更',
    'LINE_CHANGE': '产线变更',
    'STATUS_CHANGE': '状态变更'
  }
  return map[type] || type
}

const getChangeTypeTag = (type: string) => {
  const map: Record<string, string> = {
    'QUANTITY_CHANGE': 'warning',
    'DATE_CHANGE': 'info',
    'PRIORITY_CHANGE': 'danger',
    'LINE_CHANGE': 'success',
    'STATUS_CHANGE': 'primary'
  }
  return map[type] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getOrderChangeList(searchForm)
    tableData.value = res.data.list || []
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { loadData() }
const handleReset = () => { searchForm.order_no = ''; searchForm.change_type = ''; loadData() }

const handleDetail = (row: any) => {
  detailData.value = row
  detailVisible.value = true
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.order-change-list {
  .search-card { margin-bottom: 16px; }
}
</style>
