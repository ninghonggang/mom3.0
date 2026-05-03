<template>
  <div class="equipment-bom-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="设备">
          <el-select v-model="searchForm.equipment_id" placeholder="请选择设备" clearable filterable style="width: 200px">
            <el-option v-for="eq in equipmentList" :key="eq.id" :label="eq.equipment_name" :value="eq.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="物料编码">
          <el-input v-model="searchForm.material_code" placeholder="请输入物料编码" clearable />
        </el-form-item>
        <el-form-item label="物料名称">
          <el-input v-model="searchForm.material_name" placeholder="请输入物料名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon>新增BOM
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="equipment_name" label="设备名称" width="150" />
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="quantity" label="标准用量" width="100" />
        <el-table-column prop="unit" label="单位" width="80" />
        <el-table-column prop="position" label="安装位置" width="120" />
        <el-table-column prop="replace_cycle" label="更换周期(天)" width="110" />
        <el-table-column prop="is_critical" label="关键备件" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.is_critical === 1" type="danger">关键</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
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
      <el-form :model="formData" label-width="100px" :rules="formRules" ref="formRef">
        <el-form-item label="设备" prop="equipment_id">
          <el-select v-model="formData.equipment_id" placeholder="请选择设备" filterable style="width: 100%" @change="handleEquipmentChange">
            <el-option v-for="eq in equipmentList" :key="eq.id" :label="eq.equipment_name" :value="eq.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="物料" prop="material_id">
          <el-select v-model="formData.material_id" placeholder="请选择物料" filterable style="width: 100%" @change="handleMaterialChange">
            <el-option v-for="m in materialList" :key="m.id" :label="`${m.material_code} - ${m.material_name}`" :value="m.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="标准用量">
          <el-input-number v-model="formData.quantity" :min="0" :precision="2" style="width: 100%" />
        </el-form-item>
        <el-form-item label="单位">
          <el-input v-model="formData.unit" placeholder="请输入单位" />
        </el-form-item>
        <el-form-item label="安装位置">
          <el-input v-model="formData.position" placeholder="请输入安装位置" />
        </el-form-item>
        <el-form-item label="更换周期(天)">
          <el-input-number v-model="formData.replace_cycle" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="关键备件">
          <el-switch v-model="formData.is_critical" :true-value="1" :false-value="0" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
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
import { ElMessage, ElMessageBox } from 'element-plus'
import { getEquipmentBomList, createEquipmentBom, updateEquipmentBom, deleteEquipmentBom, getEquipmentBomByEquipment } from '@/api/equipment_bom'
import { getEquipmentList } from '@/api/equipment'
import { getMaterialList } from '@/api/mdm'

const loading = ref(false)
const saveLoading = ref(false)
const dialogVisible = ref(false)
const tableData = ref<any[]>([])
const equipmentList = ref<any[]>([])
const materialList = ref<any[]>([])
const dialogTitle = ref('新增BOM')

const searchForm = reactive({
  equipment_id: null as number | null,
  material_code: '',
  material_name: ''
})

const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

const formData = reactive<any>({
  equipment_id: null,
  equipment_code: '',
  equipment_name: '',
  material_id: null,
  material_code: '',
  material_name: '',
  quantity: 1,
  unit: '',
  position: '',
  replace_cycle: 0,
  is_critical: 0,
  status: 1,
  remark: ''
})

const formRules = {
  equipment_id: [{ required: true, message: '请选择设备', trigger: 'change' }],
  material_id: [{ required: true, message: '请选择物料', trigger: 'change' }]
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getEquipmentBomList({
      page: pagination.page,
      page_size: pagination.pageSize,
      equipment_id: searchForm.equipment_id,
      material_code: searchForm.material_code,
      material_name: searchForm.material_name
    })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const loadEquipment = async () => {
  try {
    const res = await getEquipmentList({ page: 1, page_size: 500 })
    equipmentList.value = res.data.list || []
  } catch (e) {
    console.error(e)
  }
}

const loadMaterials = async () => {
  try {
    const res = await getMaterialList({ page: 1, page_size: 500 })
    materialList.value = res.data.list || []
  } catch (e) {
    console.error(e)
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.equipment_id = null; searchForm.material_code = ''; searchForm.material_name = ''; handleSearch() }

const handleEquipmentChange = (id: number) => {
  const eq = equipmentList.value.find(e => e.id === id)
  if (eq) {
    formData.equipment_code = eq.equipment_code
    formData.equipment_name = eq.equipment_name
  }
}

const handleMaterialChange = (id: number) => {
  const m = materialList.value.find(mat => mat.id === id)
  if (m) {
    formData.material_code = m.material_code
    formData.material_name = m.material_name
    if (!formData.unit) formData.unit = m.unit || ''
  }
}

const handleCreate = () => {
  dialogTitle.value = '新增BOM'
  formData.equipment_id = null
  formData.equipment_code = ''
  formData.equipment_name = ''
  formData.material_id = null
  formData.material_code = ''
  formData.material_name = ''
  formData.quantity = 1
  formData.unit = ''
  formData.position = ''
  formData.replace_cycle = 0
  formData.is_critical = 0
  formData.status = 1
  formData.remark = ''
  dialogVisible.value = true
}

const handleEdit = (row: any) => {
  dialogTitle.value = '编辑BOM'
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleSave = async () => {
  if (!formData.equipment_id || !formData.material_id) {
    ElMessage.warning('请填写必填项')
    return
  }
  saveLoading.value = true
  try {
    if (formData.id) {
      await updateEquipmentBom(formData.id, formData)
    } else {
      await createEquipmentBom(formData)
    }
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadData()
  } finally {
    saveLoading.value = false
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定删除该BOM记录吗？`, '提示', { type: 'warning' })
    await deleteEquipmentBom(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

onMounted(() => { loadData(); loadEquipment(); loadMaterials() })
</script>

<style scoped lang="scss">
.equipment-bom-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>