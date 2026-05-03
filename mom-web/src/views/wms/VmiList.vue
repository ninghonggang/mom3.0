<template>
  <div class="vmi-list">
    <el-tabs v-model="activeTab">
      <el-tab-pane label="VMI供应商" name="vendor">
        <el-card class="search-card">
          <el-form :model="vendorSearchForm" inline>
            <el-form-item label="供应商编码">
              <el-input v-model="vendorSearchForm.vendor_code" placeholder="请输入" clearable />
            </el-form-item>
            <el-form-item label="供应商名称">
              <el-input v-model="vendorSearchForm.vendor_name" placeholder="请输入" clearable />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="loadVendorData">查询</el-button>
              <el-button @click="vendorReset">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>

        <el-card class="toolbar-card">
          <el-button type="primary" @click="handleCreateVendor">
            <el-icon><Plus /></el-icon>新增供应商
          </el-button>
        </el-card>

        <el-table v-loading="vendorLoading" :data="vendorTableData">
          <el-table-column prop="vendor_code" label="供应商编码" width="120" />
          <el-table-column prop="vendor_name" label="供应商名称" min-width="150" />
          <el-table-column prop="warehouse_name" label="仓库" width="100" />
          <el-table-column prop="contact" label="联系人" width="100" />
          <el-table-column prop="phone" label="电话" width="120" />
          <el-table-column prop="min_stock" label="最小库存" width="100" />
          <el-table-column prop="max_stock" label="最大库存" width="100" />
          <el-table-column prop="is_active" label="状态" width="80">
            <template #default="{ row }">
              <el-tag :type="row.is_active === 1 ? 'success' : 'info'">
                {{ row.is_active === 1 ? '启用' : '停用' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right">
            <template #default="{ row }">
              <el-button link type="primary" size="small" @click="handleEditVendor(row)">编辑</el-button>
              <el-button link type="danger" size="small" @click="handleDeleteVendor(row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <div class="pagination">
          <el-pagination
            v-model:current-page="vendorPagination.page"
            v-model:page-size="vendorPagination.pageSize"
            :total="vendorPagination.total"
            :page-sizes="[10, 20, 50]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="loadVendorData"
            @current-change="loadVendorData"
          />
        </div>
      </el-tab-pane>

      <el-tab-pane label="VMI物料" name="material">
        <el-card class="search-card">
          <el-form :model="materialSearchForm" inline>
            <el-form-item label="物料编码">
              <el-input v-model="materialSearchForm.material_code" placeholder="请输入" clearable />
            </el-form-item>
            <el-form-item label="物料名称">
              <el-input v-model="materialSearchForm.material_name" placeholder="请输入" clearable />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="loadMaterialData">查询</el-button>
              <el-button @click="materialReset">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>

        <el-table v-loading="materialLoading" :data="materialTableData">
          <el-table-column prop="material_code" label="物料编码" width="120" />
          <el-table-column prop="material_name" label="物料名称" min-width="150" />
          <el-table-column prop="unit" label="单位" width="80" />
          <el-table-column prop="current_stock" label="当前库存" width="100" />
          <el-table-column prop="available_stock" label="可用库存" width="100" />
          <el-table-column prop="consume_qty" label="累计消耗" width="100" />
          <el-table-column prop="min_stock" label="最小库存" width="100" />
          <el-table-column prop="max_stock" label="最大库存" width="100" />
          <el-table-column label="操作" width="150" fixed="right">
            <template #default="{ row }">
              <el-button link type="primary" size="small" @click="handleConsume(row)">消耗</el-button>
              <el-button link type="success" size="small" @click="handleReplenish(row)">补货</el-button>
            </template>
          </el-table-column>
        </el-table>

        <div class="pagination">
          <el-pagination
            v-model:current-page="materialPagination.page"
            v-model:page-size="materialPagination.pageSize"
            :total="materialPagination.total"
            :page-sizes="[10, 20, 50]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="loadMaterialData"
            @current-change="loadMaterialData"
          />
        </div>
      </el-tab-pane>

      <el-tab-pane label="事务记录" name="transaction">
        <el-card class="search-card">
          <el-form :model="txSearchForm" inline>
            <el-form-item label="事务类型">
              <el-select v-model="txSearchForm.transaction_type" placeholder="请选择" clearable style="width: 100px">
                <el-option label="入库" :value="1" />
                <el-option label="消耗" :value="2" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="loadTxData">查询</el-button>
              <el-button @click="txReset">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>

        <el-table v-loading="txLoading" :data="txTableData">
          <el-table-column prop="transaction_type" label="类型" width="80">
            <template #default="{ row }">
              <el-tag :type="row.transaction_type === 1 ? 'success' : 'warning'">
                {{ row.transaction_type === 1 ? '入库' : '消耗' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="material_code" label="物料编码" width="120" />
          <el-table-column prop="material_name" label="物料名称" min-width="150" />
          <el-table-column prop="quantity" label="数量" width="80" />
          <el-table-column prop="before_qty" label="变动前" width="80" />
          <el-table-column prop="after_qty" label="变动后" width="80" />
          <el-table-column prop="operator_name" label="操作人" width="100" />
          <el-table-column prop="created_at" label="时间" width="160">
            <template #default="{ row }">
              {{ row.created_at ? row.created_at.slice(0, 19) : '-' }}
            </template>
          </el-table-column>
        </el-table>

        <div class="pagination">
          <el-pagination
            v-model:current-page="txPagination.page"
            v-model:page-size="txPagination.pageSize"
            :total="txPagination.total"
            :page-sizes="[10, 20, 50]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="loadTxData"
            @current-change="loadTxData"
          />
        </div>
      </el-tab-pane>
    </el-tabs>

    <!-- 供应商对话框 -->
    <el-dialog v-model="vendorDialogVisible" :title="vendorDialogTitle" width="500px">
      <el-form :model="vendorForm" label-width="100px">
        <el-form-item label="供应商" required>
          <el-select v-model="vendorForm.vendor_id" placeholder="请选择供应商" style="width: 100%">
            <el-option v-for="v in supplierList" :key="v.id" :label="v.supplier_name" :value="v.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="仓库" required>
          <el-select v-model="vendorForm.warehouse_id" placeholder="请选择仓库" style="width: 100%">
            <el-option v-for="w in warehouseList" :key="w.id" :label="w.warehouse_name" :value="w.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人">
          <el-input v-model="vendorForm.contact" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="电话">
          <el-input v-model="vendorForm.phone" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最小库存">
          <el-input-number v-model="vendorForm.min_stock" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="最大库存">
          <el-input-number v-model="vendorForm.max_stock" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="补货周期(天)">
          <el-input-number v-model="vendorForm.replenish_cycle" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="vendorForm.remarks" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="vendorDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleVendorSave">保存</el-button>
      </template>
    </el-dialog>

    <!-- 消耗对话框 -->
    <el-dialog v-model="consumeDialogVisible" title="VMI消耗" width="400px">
      <el-form :model="consumeForm" label-width="80px">
        <el-form-item label="物料">
          <span>{{ consumeForm.material_name }}</span>
        </el-form-item>
        <el-form-item label="当前库存">
          <span>{{ consumeForm.current_stock }}</span>
        </el-form-item>
        <el-form-item label="消耗数量" required>
          <el-input-number v-model="consumeForm.quantity" :min="1" style="width: 100%" />
        </el-form-item>
        <el-form-item label="参考单号">
          <el-input v-model="consumeForm.reference_no" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="consumeForm.remarks" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="consumeDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleConsumeSubmit">确认消耗</el-button>
      </template>
    </el-dialog>

    <!-- 补货对话框 -->
    <el-dialog v-model="replenishDialogVisible" title="VMI补货" width="400px">
      <el-form :model="replenishForm" label-width="80px">
        <el-form-item label="物料">
          <span>{{ replenishForm.material_name }}</span>
        </el-form-item>
        <el-form-item label="当前库存">
          <span>{{ replenishForm.current_stock }}</span>
        </el-form-item>
        <el-form-item label="补货数量" required>
          <el-input-number v-model="replenishForm.quantity" :min="1" style="width: 100%" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="replenishForm.remarks" type="textarea" rows="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="replenishDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleReplenishSubmit">确认补货</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getVmiVendorList, createVmiVendor, updateVmiVendor, deleteVmiVendor, getVmiMaterialList, getVmiTransactionList, consumeVmi, replenishVmi } from '@/api/vmi'
import { getSupplierList } from '@/api/supplier'
import { getWarehouseList } from '@/api/warehouse'

const activeTab = ref('vendor')

// Vendor
const vendorLoading = ref(false)
const vendorDialogVisible = ref(false)
const vendorDialogTitle = ref('新增供应商')
const vendorTableData = ref<any[]>([])
const supplierList = ref<any[]>([])
const warehouseList = ref<any[]>([])

const vendorSearchForm = reactive({ vendor_code: '', vendor_name: '' })
const vendorPagination = reactive({ page: 1, pageSize: 20, total: 0 })

const vendorForm = reactive<any>({
  id: null,
  vendor_id: null,
  warehouse_id: null,
  contact: '',
  phone: '',
  min_stock: 0,
  max_stock: 0,
  replenish_cycle: 7,
  remarks: ''
})

// Material
const materialLoading = ref(false)
const materialTableData = ref<any[]>([])
const materialSearchForm = reactive({ material_code: '', material_name: '' })
const materialPagination = reactive({ page: 1, pageSize: 20, total: 0 })

// Transaction
const txLoading = ref(false)
const txTableData = ref<any[]>([])
const txSearchForm = reactive({ transaction_type: null as number | null })
const txPagination = reactive({ page: 1, pageSize: 20, total: 0 })

// Dialogs
const consumeDialogVisible = ref(false)
const replenishDialogVisible = ref(false)
const consumeForm = reactive<any>({ id: null, material_name: '', current_stock: 0, quantity: 1, reference_no: '', remarks: '' })
const replenishForm = reactive<any>({ id: null, material_name: '', current_stock: 0, quantity: 1, remarks: '' })

const loadVendorData = async () => {
  vendorLoading.value = true
  try {
    const res = await getVmiVendorList({
      page: vendorPagination.page,
      page_size: vendorPagination.pageSize,
      vendor_code: vendorSearchForm.vendor_code,
      vendor_name: vendorSearchForm.vendor_name
    })
    vendorTableData.value = res.data.list || []
    vendorPagination.total = res.data.total || 0
  } finally {
    vendorLoading.value = false
  }
}

const loadMaterialData = async () => {
  materialLoading.value = true
  try {
    const res = await getVmiMaterialList({
      page: materialPagination.page,
      page_size: materialPagination.pageSize,
      material_code: materialSearchForm.material_code,
      material_name: materialSearchForm.material_name
    })
    materialTableData.value = res.data.list || []
    materialPagination.total = res.data.total || 0
  } finally {
    materialLoading.value = false
  }
}

const loadTxData = async () => {
  txLoading.value = true
  try {
    const res = await getVmiTransactionList({
      page: txPagination.page,
      page_size: txPagination.pageSize,
      transaction_type: txSearchForm.transaction_type
    })
    txTableData.value = res.data.list || []
    txPagination.total = res.data.total || 0
  } finally {
    txLoading.value = false
  }
}

const loadSuppliers = async () => {
  const res = await getSupplierList({ page: 1, page_size: 500 })
  supplierList.value = res.data.list || []
}

const loadWarehouses = async () => {
  const res = await getWarehouseList({ page: 1, page_size: 500 })
  warehouseList.value = res.data.list || []
}

const vendorReset = () => {
  vendorSearchForm.vendor_code = ''
  vendorSearchForm.vendor_name = ''
  loadVendorData()
}

const materialReset = () => {
  materialSearchForm.material_code = ''
  materialSearchForm.material_name = ''
  loadMaterialData()
}

const txReset = () => {
  txSearchForm.transaction_type = null
  loadTxData()
}

const handleCreateVendor = () => {
  vendorDialogTitle.value = '新增供应商'
  vendorForm.id = null
  vendorForm.vendor_id = null
  vendorForm.warehouse_id = null
  vendorForm.contact = ''
  vendorForm.phone = ''
  vendorForm.min_stock = 0
  vendorForm.max_stock = 0
  vendorForm.replenish_cycle = 7
  vendorForm.remarks = ''
  vendorDialogVisible.value = true
}

const handleEditVendor = (row: any) => {
  vendorDialogTitle.value = '编辑供应商'
  Object.assign(vendorForm, row)
  vendorDialogVisible.value = true
}

const handleVendorSave = async () => {
  if (!vendorForm.vendor_id || !vendorForm.warehouse_id) {
    ElMessage.warning('请填写必填项')
    return
  }
  if (vendorForm.id) {
    await updateVmiVendor(vendorForm.id, vendorForm)
  } else {
    await createVmiVendor(vendorForm)
  }
  ElMessage.success('保存成功')
  vendorDialogVisible.value = false
  loadVendorData()
}

const handleDeleteVendor = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定删除吗？', '提示', { type: 'warning' })
    await deleteVmiVendor(row.id)
    ElMessage.success('删除成功')
    loadVendorData()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

const handleConsume = (row: any) => {
  consumeForm.id = row.id
  consumeForm.material_name = row.material_name
  consumeForm.current_stock = row.current_stock
  consumeForm.quantity = 1
  consumeForm.reference_no = ''
  consumeForm.remarks = ''
  consumeDialogVisible.value = true
}

const handleConsumeSubmit = async () => {
  try {
    await consumeVmi({ vendor_id: 0, material_id: consumeForm.id, quantity: consumeForm.quantity, reference_no: consumeForm.reference_no, remarks: consumeForm.remarks })
    ElMessage.success('消耗成功')
    consumeDialogVisible.value = false
    loadMaterialData()
    loadTxData()
  } catch (e: any) {
    ElMessage.error(e.message || '操作失败')
  }
}

const handleReplenish = (row: any) => {
  replenishForm.id = row.id
  replenishForm.material_name = row.material_name
  replenishForm.current_stock = row.current_stock
  replenishForm.quantity = 1
  replenishForm.remarks = ''
  replenishDialogVisible.value = true
}

const handleReplenishSubmit = async () => {
  try {
    await replenishVmi({ vendor_id: 0, material_id: replenishForm.id, quantity: replenishForm.quantity, remarks: replenishForm.remarks })
    ElMessage.success('补货成功')
    replenishDialogVisible.value = false
    loadMaterialData()
    loadTxData()
  } catch (e: any) {
    ElMessage.error(e.message || '操作失败')
  }
}

onMounted(() => {
  loadVendorData()
  loadMaterialData()
  loadTxData()
  loadSuppliers()
  loadWarehouses()
})
</script>

<style scoped lang="scss">
.vmi-list {
  .search-card, .toolbar-card { margin-bottom: 16px; }
  .toolbar-card :deep(.el-card__body) { padding: 12px 16px; display: flex; gap: 12px; }
  .pagination { margin-top: 16px; display: flex; justify-content: flex-end; }
}
</style>