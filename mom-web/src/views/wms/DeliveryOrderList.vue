<template>
  <div class="delivery-order-list">
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="发货单号">
          <el-input v-model="searchForm.delivery_no" placeholder="请输入发货单号" clearable />
        </el-form-item>
        <el-form-item label="客户">
          <el-input v-model="searchForm.customer_name" placeholder="请输入客户名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="待发货" :value="1" />
            <el-option label="发货中" :value="2" />
            <el-option label="已完成" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="toolbar-card">
      <el-button type="primary" v-if="hasPermission('wms:deliveryorder:add')" @click="handleAdd">
        <el-icon><Plus /></el-icon>新增
      </el-button>
    </el-card>

    <el-card>
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="delivery_no" label="发货单号" width="150" />
        <el-table-column prop="customer_name" label="客户" min-width="150" />
        <el-table-column prop="warehouse_id" label="仓库ID" width="100" />
        <el-table-column prop="delivery_date" label="发货日期" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" v-if="hasPermission('wms:deliveryorder:view')" @click="handleView(row)">详情</el-button>
            <el-button link type="primary" v-if="hasPermission('wms:deliveryorder:edit')" @click="handleEdit(row)" :disabled="row.status === 3">编辑</el-button>
            <el-button link type="success" v-if="hasPermission('wms:deliveryorder:confirm')" @click="handleConfirm(row)" :disabled="row.status !== 1">确认</el-button>
            <el-button link type="danger" v-if="hasPermission('wms:deliveryorder:delete')" @click="handleDelete(row)" :disabled="row.status === 3">删除</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="客户ID" prop="customer_id">
          <el-input-number v-model="formData.customer_id" :min="1" />
        </el-form-item>
        <el-form-item label="客户名称" prop="customer_name">
          <el-input v-model="formData.customer_name" />
        </el-form-item>
        <el-form-item label="仓库" prop="warehouse_id">
          <el-input-number v-model="formData.warehouse_id" :min="1" />
        </el-form-item>
        <el-form-item label="发货日期" prop="delivery_date">
          <el-date-picker v-model="formData.delivery_date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" type="textarea" :rows="2" />
        </el-form-item>

        <el-divider>发货明细</el-divider>
        <div v-for="(item, index) in formData.items" :key="index" class="item-row">
          <el-form-item label="物料ID" prop="material_id">
            <el-input-number v-model="item.material_id" :min="1" @change="handleMaterialChange(index)" />
          </el-form-item>
          <el-form-item label="物料编码">
            <el-input v-model="item.material_code" />
          </el-form-item>
          <el-form-item label="物料名称">
            <el-input v-model="item.material_name" />
          </el-form-item>
          <el-form-item label="数量">
            <el-input-number v-model="item.quantity" :min="0" />
          </el-form-item>
          <el-form-item label="单位">
            <el-input v-model="item.unit" />
          </el-form-item>
          <el-form-item label="批次号">
            <el-input v-model="item.batch_no" />
          </el-form-item>
          <el-button type="danger" link @click="removeItem(index)">删除</el-button>
        </div>
        <el-button type="primary" link @click="addItem">+ 添加物料</el-button>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="发货单详情" width="800px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="发货单号">{{ detailData.delivery_no }}</el-descriptions-item>
        <el-descriptions-item label="客户">{{ detailData.customer_name }}</el-descriptions-item>
        <el-descriptions-item label="仓库ID">{{ detailData.warehouse_id }}</el-descriptions-item>
        <el-descriptions-item label="发货日期">{{ detailData.delivery_date }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ getStatusText(detailData.status) }}</el-descriptions-item>
        <el-descriptions-item label="备注">{{ detailData.remark }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>发货明细</el-divider>
      <el-table :data="detailData.items" border>
        <el-table-column prop="material_code" label="物料编码" width="120" />
        <el-table-column prop="material_name" label="物料名称" min-width="150" />
        <el-table-column prop="quantity" label="数量" width="100" />
        <el-table-column prop="shipped_qty" label="已发货" width="100" />
        <el-table-column prop="unit" label="单位" width="80" />
        <el-table-column prop="batch_no" label="批次号" width="120" />
      </el-table>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import { getDeliveryOrderList, createDeliveryOrder, updateDeliveryOrder, deleteDeliveryOrder, deliveryConfirm, getDeliveryOrderById } from '@/api/wms'
import { useAuthStore } from '@/stores/auth'

const { hasPermission } = useAuthStore()

const loading = ref(false)
const tableData = ref<any[]>([])
const dialogVisible = ref(false)
const detailVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref<FormInstance>()

const searchForm = reactive({ delivery_no: '', customer_name: '', status: undefined as number | undefined })
const pagination = reactive({ page: 1, pageSize: 20, total: 0 })
const detailData = reactive<any>({ items: [] })

interface OrderItem {
  material_id: number
  material_code: string
  material_name: string
  quantity: number
  unit: string
  batch_no: string
}

const formData = reactive<any>({
  id: 0,
  customer_id: 0,
  customer_name: '',
  warehouse_id: 0,
  delivery_date: '',
  remark: '',
  items: [] as OrderItem[]
})

const rules: FormRules = {
  customer_id: [{ required: true, message: '请输入客户ID', trigger: 'blur' }],
  warehouse_id: [{ required: true, message: '请输入仓库ID', trigger: 'blur' }]
}

const dialogTitle = computed(() => formData.id ? '编辑发货单' : '新增发货单')

const getStatusText = (status: number) => {
  const map: Record<number, string> = { 1: '待发货', 2: '发货中', 3: '已完成' }
  return map[status] || '未知'
}

const getStatusType = (status: number) => {
  const map: Record<number, string> = { 1: 'warning', 2: 'primary', 3: 'success' }
  return map[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await getDeliveryOrderList({ ...searchForm, page: pagination.page, page_size: pagination.pageSize })
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { pagination.page = 1; loadData() }
const handleReset = () => { searchForm.delivery_no = ''; searchForm.customer_name = ''; searchForm.status = undefined; handleSearch() }

const handleAdd = () => {
  Object.assign(formData, {
    id: 0, customer_id: 0, customer_name: '', warehouse_id: 0, delivery_date: '', remark: '', items: []
  })
  dialogVisible.value = true
}

const handleEdit = async (row: any) => {
  try {
    const res = await getDeliveryOrderById(row.id)
    const data = res.data
    Object.assign(formData, {
      id: data.order.id,
      customer_id: data.order.customer_id,
      customer_name: data.order.customer_name,
      warehouse_id: data.order.warehouse_id,
      delivery_date: data.order.delivery_date,
      remark: data.order.remark || '',
      items: data.items.map((item: any) => ({
        material_id: item.material_id,
        material_code: item.material_code,
        material_name: item.material_name,
        quantity: item.quantity,
        unit: item.unit,
        batch_no: item.batch_no || ''
      }))
    })
    dialogVisible.value = true
  } catch (error) {
    ElMessage.error('获取发货单详情失败')
  }
}

const handleView = async (row: any) => {
  try {
    const res = await getDeliveryOrderById(row.id)
    Object.assign(detailData, res.data)
    detailVisible.value = true
  } catch (error) {
    ElMessage.error('获取发货单详情失败')
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除该发货单吗？', '提示', { type: 'warning' })
    await deleteDeliveryOrder(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    ElMessage.error('删除失败')
  }
}

const handleConfirm = async (row: any) => {
  await ElMessageBox.confirm('确定确认发货单吗？', '提示', { type: 'warning' })
  try {
    await deliveryConfirm(row.id)
    ElMessage.success('确认成功')
    loadData()
  } catch (error) {
    ElMessage.error('确认失败')
  }
}

const addItem = () => {
  formData.items.push({ material_id: 0, material_code: '', material_name: '', quantity: 0, unit: '', batch_no: '' })
}

const removeItem = (index: number) => {
  formData.items.splice(index, 1)
}

const handleMaterialChange = (index: number) => {
  // Material change handler - could be used to fetch material details
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitLoading.value = true
  try {
    if (formData.id) {
      await updateDeliveryOrder(formData.id, formData)
    } else {
      await createDeliveryOrder(formData)
    }
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadData()
  } catch (error) {
    ElMessage.error(formData.id ? '更新失败' : '创建失败')
  } finally {
    submitLoading.value = false
  }
}

onMounted(() => { loadData() })
</script>

<style scoped lang="scss">
.delivery-order-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
  .item-row {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    align-items: center;
    margin-bottom: 12px;
    padding: 12px;
    background: #f5f7fa;
    border-radius: 4px;
  }
}
</style>
