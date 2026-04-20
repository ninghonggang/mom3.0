<template>
  <div class="person-skill-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="人员姓名">
          <el-input v-model="searchForm.person_name" placeholder="请输入人员姓名" clearable />
        </el-form-item>
        <el-form-item label="技能等级">
          <el-select v-model="searchForm.skill_level" placeholder="请选择" clearable>
            <el-option label="初级" :value="1" />
            <el-option label="中级" :value="2" />
            <el-option label="高级" :value="3" />
            <el-option label="专家" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('mes:personSkill:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="person_name" label="人员姓名" min-width="120" />
        <el-table-column prop="skill_level" label="技能等级" width="100">
          <template #default="{ row }">
            <el-tag :type="getSkillType(row.skill_level)">{{ getSkillText(row.skill_level) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="qualified_operations" label="合格工序" min-width="200" show-overflow-tooltip />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('mes:personSkill:detail')" @click="handleDetail(row)">明细</el-button>
            <el-button link type="primary" size="small" v-if="hasPermission('mes:personSkill:edit')" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('mes:personSkill:delete')" @click="handleDelete(row)">删除</el-button>
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
import request from '@/utils/request'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])

const searchForm = reactive({ person_name: '', skill_level: '' })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const getSkillText = (level: number) => {
  const map: Record<number, string> = { 1: '初级', 2: '中级', 3: '高级', 4: '专家' }
  return map[level] || '未知'
}

const getSkillType = (level: number) => {
  const map: Record<number, string> = { 1: 'info', 2: 'warning', 3: 'primary', 4: 'success' }
  return map[level] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await request.get('/mes/person-skill/list', {
      params: { ...searchForm, page: pagination.page, page_size: pagination.pageSize }
    })
    tableData.value = res.data?.list || []
    pagination.total = res.data?.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.person_name = ''; searchForm.skill_level = ''; handleSearch() }
const handleAdd = () => { ElMessage.info('新增人员能力') }
const handleDetail = (row: any) => { ElMessage.info('查看明细') }
const handleEdit = (row: any) => { ElMessage.info('编辑') }
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该人员能力记录吗？', '提示', { type: 'warning' })
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    // user cancelled or API error
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.person-skill-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
