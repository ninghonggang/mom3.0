<template>
  <el-dialog
    v-model="visible"
    :title="dialogTitle"
    width="900px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="BOM编码" prop="bom_code">
            <el-input v-model="formData.bom_code" :disabled="mode === 'edit'" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="BOM名称" prop="bom_name">
            <el-input v-model="formData.bom_name" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="产品编码" prop="material_code">
            <el-input v-model="formData.material_code" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="产品名称">
            <el-input v-model="formData.material_name" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="版本">
            <el-input v-model="formData.version" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="状态">
            <el-select v-model="formData.status">
              <el-option label="草稿" value="DRAFT" />
              <el-option label="生效" value="ACTIVE" />
              <el-option label="失效" value="EXPIRED" />
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="生效日期">
            <el-date-picker v-model="formData.eff_date" type="date" value-format="YYYY-MM-DD" style="width: 100%" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="失效日期">
            <el-date-picker v-model="formData.exp_date" type="date" value-format="YYYY-MM-DD" style="width: 100%" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="备注">
        <el-input v-model="formData.remark" type="textarea" :rows="2" />
      </el-form-item>

      <!-- BOM明细 -->
      <el-divider content-position="left">BOM明细</el-divider>
      <el-button type="primary" size="small" @click="handleAddItem" style="margin-bottom: 10px">
        <el-icon><Plus /></el-icon>添加物料
      </el-button>
      <el-table :data="formData.items" border size="small">
        <el-table-column label="行号" width="60">
          <template #default="{ row, $index }">
            <span>{{ $index + 1 }}</span>
          </template>
        </el-table-column>
        <el-table-column label="物料编码" width="150">
          <template #default="{ row, $index }">
            <el-input v-model="row.material_code" size="small" />
          </template>
        </el-table-column>
        <el-table-column label="物料名称" width="180">
          <template #default="{ row }">
            <el-input v-model="row.material_name" size="small" />
          </template>
        </el-table-column>
        <el-table-column label="用量" width="100">
          <template #default="{ row }">
            <el-input-number v-model="row.quantity" :min="0" :precision="4" size="small" style="width: 80px" />
          </template>
        </el-table-column>
        <el-table-column label="单位" width="80">
          <template #default="{ row }">
            <el-input v-model="row.unit" size="small" />
          </template>
        </el-table-column>
        <el-table-column label="损耗率%" width="100">
          <template #default="{ row }">
            <el-input-number v-model="row.scrap_rate" :min="0" :max="100" :precision="2" size="small" style="width: 80px" />
          </template>
        </el-table-column>
        <el-table-column label="替代组" width="100">
          <template #default="{ row }">
            <el-input v-model="row.substitute_group" size="small" />
          </template>
        </el-table-column>
        <el-table-column label="替代料" width="80">
          <template #default="{ row }">
            <el-checkbox v-model="row.is_alternative" :true-value="1" :false-value="0" />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="80">
          <template #default="{ $index }">
            <el-button link type="danger" size="small" @click="handleRemoveItem($index)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-form>

    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { createBOM, updateBOM, getBOMList } from '@/api/mdm'

const props = defineProps<{
  modelValue: boolean
  bomId: number | null
  mode: 'create' | 'edit'
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', val: boolean): void
  (e: 'refresh'): void
}>()

const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const dialogTitle = computed(() => props.mode === 'create' ? '新增BOM' : '编辑BOM')

const formRef = ref()
const submitLoading = ref(false)

const formData = reactive({
  bom_code: '',
  bom_name: '',
  material_code: '',
  material_name: '',
  version: 'V1',
  status: 'DRAFT',
  eff_date: '',
  exp_date: '',
  remark: '',
  items: [] as any[]
})

const rules = {
  bom_code: [{ required: true, message: '请输入BOM编码', trigger: 'blur' }],
  bom_name: [{ required: true, message: '请输入BOM名称', trigger: 'blur' }]
}

watch(() => props.modelValue, async (val) => {
  if (val) {
    if (props.mode === 'edit' && props.bomId) {
      await loadBOM()
    } else {
      resetForm()
    }
  }
})

const loadBOM = async () => {
  try {
    const res = await getBOMList({ id: props.bomId })
    const list = res.data.list || []
    if (list.length > 0) {
      const bom = list[0]
      Object.assign(formData, {
        bom_code: bom.bom_code,
        bom_name: bom.bom_name,
        material_code: bom.material_code,
        material_name: bom.material_name,
        version: bom.version,
        status: bom.status,
        eff_date: bom.eff_date || '',
        exp_date: bom.exp_date || '',
        remark: bom.remark || ''
      })
      // 加载明细
      formData.items = bom.items || []
    }
  } catch (error) {
    ElMessage.error('加载BOM失败')
  }
}

const resetForm = () => {
  formData.bom_code = ''
  formData.bom_name = ''
  formData.material_code = ''
  formData.material_name = ''
  formData.version = 'V1'
  formData.status = 'DRAFT'
  formData.eff_date = ''
  formData.exp_date = ''
  formData.remark = ''
  formData.items = []
}

const handleAddItem = () => {
  formData.items.push({
    material_code: '',
    material_name: '',
    quantity: 1,
    unit: '',
    scrap_rate: 0,
    substitute_group: '',
    is_alternative: 0
  })
}

const handleRemoveItem = (index: number) => {
  formData.items.splice(index, 1)
}

const handleClose = () => {
  visible.value = false
}

const handleSubmit = async () => {
  await formRef.value?.validate()
  submitLoading.value = true
  try {
    if (props.mode === 'create') {
      await createBOM(formData)
      ElMessage.success('创建成功')
    } else {
      await updateBOM(props.bomId!, formData)
      ElMessage.success('更新成功')
    }
    emit('refresh')
    handleClose()
  } catch (error: any) {
    ElMessage.error(error.message || '操作失败')
  } finally {
    submitLoading.value = false
  }
}
</script>
