<template>
  <div class="inspection-record-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="设备编码">
          <el-input v-model="searchForm.equipment_code" placeholder="请输入编码" clearable />
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.equipment_name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="点检结果">
          <el-select v-model="searchForm.result" placeholder="请选择" clearable>
            <el-option label="合格" value="OK" />
            <el-option label="异常" value="NG" />
          </el-select>
        </el-form-item>
        <el-form-item label="点检人">
          <el-input v-model="searchForm.inspector" placeholder="请输入点检人" clearable />
        </el-form-item>
        <el-form-item label="点检日期">
          <el-date-picker
            v-model="searchForm.date_range"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('equipment:inspection:record:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增点检
      </el-button>
      <el-button type="success" v-if="hasPermission('equipment:inspection:record:export')" @click="handleExport">
        <el-icon><Download /></el-icon>导出
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="record_code" label="点检记录编号" width="160" />
        <el-table-column prop="equipment_code" label="设备编码" width="120" />
        <el-table-column prop="equipment_name" label="设备名称" min-width="150" />
        <el-table-column prop="inspection_template_name" label="点检模板" width="140" />
        <el-table-column prop="plan_code" label="点检计划编号" width="160" />
        <el-table-column prop="inspector" label="点检人" width="100" />
        <el-table-column prop="inspection_time" label="点检时间" width="160" />
        <el-table-column prop="result" label="点检结果" width="100">
          <template #default="{ row }">
            <el-tag :type="row.result === 'OK' ? 'success' : 'danger'">
              {{ row.result === 'OK' ? '合格' : '异常' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="备注" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" v-if="hasPermission('equipment:inspection:record:view')" @click="handleView(row)">查看</el-button>
            <el-button link type="success" size="small" v-if="hasPermission('equipment:inspection:record:complete') && row.status === 'IN_PROGRESS'" @click="handleComplete(row)">完成</el-button>
            <el-button link type="danger" size="small" v-if="hasPermission('equipment:inspection:record:delete')" @click="handleDelete(row)">删除</el-button>
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

    <!-- 新增点检对话框 -->
    <el-dialog v-model="addDialogVisible" title="新增点检" width="600px">
      <el-form ref="addFormRef" :model="addForm" :rules="addRules" label-width="120px">
        <el-form-item label="设备" prop="equipment_id">
          <el-select v-model="addForm.equipment_id" placeholder="请选择设备" filterable @change="handleEquipmentChange">
            <el-option
              v-for="item in equipmentList"
              :key="item.id"
              :label="`${item.equipment_code} - ${item.equipment_name}`"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="点检模板" prop="template_id">
          <el-select v-model="addForm.template_id" placeholder="请选择点检模板" filterable>
            <el-option
              v-for="item in templateList"
              :key="item.id"
              :label="item.template_name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="点检人" prop="inspector">
          <el-input v-model="addForm.inspector" placeholder="请输入点检人" />
        </el-form-item>
        <el-form-item label="备注" prop="description">
          <el-input v-model="addForm.description" type="textarea" :rows="3" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="addDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleStartInspection">确定</el-button>
      </template>
    </el-dialog>

    <!-- 完成点检对话框 -->
    <el-dialog v-model="completeDialogVisible" title="完成点检" width="700px">
      <el-form ref="completeFormRef" :model="completeForm" :rules="completeRules" label-width="100px">
        <el-form-item label="点检结果" prop="result">
          <el-radio-group v-model="completeForm.result">
            <el-radio label="OK">合格</el-radio>
            <el-radio label="NG">异常</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="实际点检时间" prop="inspection_time">
          <el-date-picker
            v-model="completeForm.inspection_time"
            type="datetime"
            placeholder="请选择点检时间"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="异常描述" prop="defect_description" v-if="completeForm.result === 'NG'">
          <el-input v-model="completeForm.defect_description" type="textarea" :rows="3" placeholder="请输入异常描述" />
        </el-form-item>
        <el-form-item label="处理措施" prop="remedial_action" v-if="completeForm.result === 'NG'">
          <el-input v-model="completeForm.remedial_action" type="textarea" :rows="2" placeholder="请输入处理措施" />
        </el-form-item>
        <el-form-item label="备注" prop="description">
          <el-input v-model="completeForm.description" type="textarea" :rows="2" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="completeDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmitComplete">确定</el-button>
      </template>
    </el-dialog>

    <!-- 查看详情对话框 -->
    <el-dialog v-model="viewDialogVisible" title="点检记录详情" width="700px">
      <el-descriptions :column="2" border v-if="currentRow">
        <el-descriptions-item label="点检记录编号">{{ currentRow.record_code }}</el-descriptions-item>
        <el-descriptions-item label="点检结果">
          <el-tag :type="currentRow.result === 'OK' ? 'success' : 'danger'">
            {{ currentRow.result === 'OK' ? '合格' : '异常' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="设备编码">{{ currentRow.equipment_code }}</el-descriptions-item>
        <el-descriptions-item label="设备名称">{{ currentRow.equipment_name }}</el-descriptions-item>
        <el-descriptions-item label="点检模板">{{ currentRow.inspection_template_name }}</el-descriptions-item>
        <el-descriptions-item label="点检计划编号">{{ currentRow.plan_code }}</el-descriptions-item>
        <el-descriptions-item label="点检人">{{ currentRow.inspector }}</el-descriptions-item>
        <el-descriptions-item label="点检时间">{{ currentRow.inspection_time }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ getStatusText(currentRow.status) }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ currentRow.create_time }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ currentRow.description || '-' }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="viewDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Download } from '@element-plus/icons-vue'
import { getInspectionRecordList, startInspection, completeInspection, deleteInspectionRecord, getInspectionTemplateList } from '@/api/equipment'
import { getEquipmentList } from '@/api/equipment'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const currentRow = ref<any>(null)

// 对话框状态
const addDialogVisible = ref(false)
const completeDialogVisible = ref(false)
const viewDialogVisible = ref(false)
const addFormRef = ref()
const completeFormRef = ref()

// 下拉列表数据
const equipmentList = ref<any[]>([])
const templateList = ref<any[]>([])

// 搜索表单
const searchForm = reactive({
  equipment_code: '',
  equipment_name: '',
  result: '',
  inspector: '',
  date_range: []
})

// 分页
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

// 新增表单
const addForm = reactive({
  equipment_id: null as number | null,
  template_id: null as number | null,
  inspector: '',
  description: ''
})

// 完成表单
const completeForm = reactive({
  result: 'OK',
  inspection_time: '',
  defect_description: '',
  remedial_action: '',
  description: ''
})

const currentRecordId = ref<number | null>(null)

// 表单验证规则
const addRules = {
  equipment_id: [{ required: true, message: '请选择设备', trigger: 'change' }],
  template_id: [{ required: true, message: '请选择点检模板', trigger: 'change' }],
  inspector: [{ required: true, message: '请输入点检人', trigger: 'blur' }]
}

const completeRules = {
  result: [{ required: true, message: '请选择点检结果', trigger: 'change' }],
  inspection_time: [{ required: true, message: '请选择点检时间', trigger: 'change' }]
}

const getStatusType = (status: string) => {
  const map: Record<string, any> = {
    PENDING: 'info',
    IN_PROGRESS: 'warning',
    COMPLETED: 'success',
    CANCELLED: 'danger'
  }
  return map[status] || 'info'
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    PENDING: '待点检',
    IN_PROGRESS: '点检中',
    COMPLETED: '已完成',
    CANCELLED: '已取消'
  }
  return map[status] || status
}

const loadEquipmentList = async () => {
  try {
    const res = await getEquipmentList({ page: 1, page_size: 1000 })
    equipmentList.value = res.data.list || []
  } catch (error: any) {
    ElMessage.error(error.message || '加载设备列表失败')
  }
}

const loadTemplateList = async () => {
  try {
    const res = await getInspectionTemplateList({ page: 1, page_size: 1000 })
    templateList.value = res.data.list || []
  } catch (error: any) {
    ElMessage.error(error.message || '加载点检模板列表失败')
  }
}

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.pageSize
    }
    if (searchForm.equipment_code) params.equipment_code = searchForm.equipment_code
    if (searchForm.equipment_name) params.equipment_name = searchForm.equipment_name
    if (searchForm.result) params.result = searchForm.result
    if (searchForm.inspector) params.inspector = searchForm.inspector
    if (searchForm.date_range && searchForm.date_range.length === 2) {
      params.start_date = searchForm.date_range[0]
      params.end_date = searchForm.date_range[1]
    }
    const res = await getInspectionRecordList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  searchForm.equipment_code = ''
  searchForm.equipment_name = ''
  searchForm.result = ''
  searchForm.inspector = ''
  searchForm.date_range = []
  handleSearch()
}

const handleAdd = () => {
  addForm.equipment_id = null
  addForm.template_id = null
  addForm.inspector = ''
  addForm.description = ''
  addDialogVisible.value = true
}

const handleEquipmentChange = (val: number) => {
  const equipment = equipmentList.value.find(e => e.id === val)
  if (equipment) {
    // 可以根据设备自动筛选关联的点检模板
  }
}

const handleStartInspection = async () => {
  try {
    await addFormRef.value.validate()
    await startInspection(addForm)
    ElMessage.success('点检任务已创建')
    addDialogVisible.value = false
    loadData()
  } catch (error: any) {
    if (error !== false) {
      ElMessage.error(error.message || '操作失败')
    }
  }
}

const handleComplete = (row: any) => {
  currentRecordId.value = row.id
  completeForm.result = 'OK'
  completeForm.inspection_time = ''
  completeForm.defect_description = ''
  completeForm.remedial_action = ''
  completeForm.description = ''
  completeDialogVisible.value = true
}

const handleSubmitComplete = async () => {
  try {
    await completeFormRef.value.validate()
    await completeInspection(currentRecordId.value!, completeForm)
    ElMessage.success('点检已完成')
    completeDialogVisible.value = false
    loadData()
  } catch (error: any) {
    if (error !== false) {
      ElMessage.error(error.message || '操作失败')
    }
  }
}

const handleView = (row: any) => {
  currentRow.value = row
  viewDialogVisible.value = true
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该点检记录吗？', '提示', { type: 'warning' })
    await deleteInspectionRecord(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

const handleExport = () => {
  ElMessage.info('导出功能开发中')
}

onMounted(() => {
  loadData()
  loadEquipmentList()
  loadTemplateList()
})
</script>

<style scoped lang="scss">
.inspection-record-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>
