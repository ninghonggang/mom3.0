import request from '@/utils/request'

// 物料相关
export const getMaterialList = (params?: any) => {
  return request.get('/mdm/material/list', { params })
}

export const getMaterialById = (id: number) => {
  return request.get(`/mdm/material/${id}`)
}

export const createMaterial = (data: any) => {
  return request.post('/mdm/material', data)
}

export const updateMaterial = (id: number, data: any) => {
  return request.put(`/mdm/material/${id}`, data)
}

export const deleteMaterial = (id: number) => {
  return request.delete(`/mdm/material/${id}`)
}

// BOM相关
export const getBOMList = (params?: any) => {
  return request.get('/mdm/bom/list', { params })
}

export const getBOMTree = (productId: number) => {
  return request.get(`/mdm/bom/${productId}/tree`)
}

export const createBOM = (data: any) => {
  return request.post('/mdm/bom', data)
}

export const updateBOM = (id: number, data: any) => {
  return request.put(`/mdm/bom/${id}`, data)
}

export const deleteBOM = (id: number) => {
  return request.delete(`/mdm/bom/${id}`)
}

// 工艺路线
export const getProcessList = (params?: any) => {
  return request.get('/mdm/process/list', { params })
}

export const createProcess = (data: any) => {
  return request.post('/mdm/process', data)
}

export const updateProcess = (id: number, data: any) => {
  return request.put(`/mdm/process/${id}`, data)
}

export const deleteProcess = (id: number) => {
  return request.delete(`/mdm/process/${id}`)
}

// 车间
export const getWorkshopList = (params?: any) => {
  return request.get('/mdm/workshop/list', { params })
}

export const createWorkshop = (data: any) => {
  return request.post('/mdm/workshop', data)
}

export const updateWorkshop = (id: number, data: any) => {
  return request.put(`/mdm/workshop/${id}`, data)
}

export const deleteWorkshop = (id: number) => {
  return request.delete(`/mdm/workshop/${id}`)
}

// 生产线
export const getProductionLineList = (params?: any) => {
  return request.get('/mdm/line/list', { params })
}

export const createProductionLine = (data: any) => {
  return request.post('/mdm/line', data)
}

export const updateProductionLine = (id: number, data: any) => {
  return request.put(`/mdm/line/${id}`, data)
}

export const deleteProductionLine = (id: number) => {
  return request.delete(`/mdm/line/${id}`)
}

// 工位
export const getWorkstationList = (params?: any) => {
  return request.get('/mdm/workstation/list', { params })
}

export const createWorkstation = (data: any) => {
  return request.post('/mdm/workstation', data)
}

export const updateWorkstation = (id: number, data: any) => {
  return request.put(`/mdm/workstation/${id}`, data)
}

export const deleteWorkstation = (id: number) => {
  return request.delete(`/mdm/workstation/${id}`)
}

// 班次
export const getShiftList = (params?: any) => {
  return request.get('/mes/shift/list', { params })
}

export const createShift = (data: any) => {
  return request.post('/mes/shift', data)
}

export const updateShift = (id: number, data: any) => {
  return request.put(`/mes/shift/${id}`, data)
}

export const deleteShift = (id: number) => {
  return request.delete(`/mes/shift/${id}`)
}

// MDM 工序
export const getOperationList = (params?: any) => {
  return request.get('/mdm/operation/list', { params })
}

export const getOperationById = (id: number) => {
  return request.get(`/mdm/operation/${id}`)
}

export const createOperation = (data: any) => {
  return request.post('/mdm/operation', data)
}

export const updateOperation = (id: number, data: any) => {
  return request.put(`/mdm/operation/${id}`, data)
}

export const deleteOperation = (id: number) => {
  return request.delete(`/mdm/operation/${id}`)
}

// MDM 班次
export const getMdmShiftList = (params?: any) => {
  return request.get('/mdm/mdm-shift/list', { params })
}

export const getMdmShiftById = (id: number) => {
  return request.get(`/mdm/mdm-shift/${id}`)
}

export const createMdmShift = (data: any) => {
  return request.post('/mdm/mdm-shift', data)
}

export const updateMdmShift = (id: number, data: any) => {
  return request.put(`/mdm/mdm-shift/${id}`, data)
}

export const deleteMdmShift = (id: number) => {
  return request.delete(`/mdm/mdm-shift/${id}`)
}

// 客户管理
export const getCustomerList = (params?: any) => {
  return request.get('/mdm/customer/list', { params })
}

export const getCustomerById = (id: number) => {
  return request.get(`/mdm/customer/${id}`)
}

export const createCustomer = (data: any) => {
  return request.post('/mdm/customer', data)
}

export const updateCustomer = (id: number, data: any) => {
  return request.put(`/mdm/customer/${id}`, data)
}

export const deleteCustomer = (id: number) => {
  return request.delete(`/mdm/customer/${id}`)
}

// 物料分类
export const getMaterialCategoryList = (params?: any) => {
  return request.get('/mdm/material-category/list', { params })
}

export const getMaterialCategoryTree = (params?: any) => {
  return request.get('/mdm/material-category/tree', { params })
}

export const getMaterialCategoryById = (id: number) => {
  return request.get(`/mdm/material-category/${id}`)
}

export const createMaterialCategory = (data: any) => {
  return request.post('/mdm/material-category', data)
}

export const updateMaterialCategory = (id: number, data: any) => {
  return request.put(`/mdm/material-category/${id}`, data)
}

export const deleteMaterialCategory = (id: number) => {
  return request.delete(`/mdm/material-category/${id}`)
}

// 物料导入
export const importMaterials = (data: FormData) => {
  return request.post('/mdm/material/import', data, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}

export const downloadMaterialTemplate = () => {
  return request.get('/mdm/material/template', { responseType: 'blob' })
}

export const getImportTask = (id: number) => {
  return request.get(`/system/import-task/${id}`)
}

export const getImportTaskResult = (id: number) => {
  return request.get(`/system/import-task/${id}/result`)
}

// BOM导入
export const importBOMs = (data: FormData) => {
  return request.post('/mdm/bom/import', data, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}

export const downloadBOMTemplate = () => {
  return request.get('/mdm/bom/template', { responseType: 'blob' })
}

// 联系人管理
export const getContactList = (params?: any) => {
  return request.get('/mdm/contact/list', { params })
}

export const getContactById = (id: number) => {
  return request.get(`/mdm/contact/${id}`)
}

export const createContact = (data: any) => {
  return request.post('/mdm/contact', data)
}

export const updateContact = (id: number, data: any) => {
  return request.put(`/mdm/contact/${id}`, data)
}

export const deleteContact = (id: number) => {
  return request.delete(`/mdm/contact/${id}`)
}

// 银行账户管理
export const getBankAccountList = (params?: any) => {
  return request.get('/mdm/bank-account/list', { params })
}

export const getBankAccountById = (id: number) => {
  return request.get(`/mdm/bank-account/${id}`)
}

export const createBankAccount = (data: any) => {
  return request.post('/mdm/bank-account', data)
}

export const updateBankAccount = (id: number, data: any) => {
  return request.put(`/mdm/bank-account/${id}`, data)
}

export const deleteBankAccount = (id: number) => {
  return request.delete(`/mdm/bank-account/${id}`)
}

// 附件管理
export const getAttachmentList = (params?: any) => {
  return request.get('/mdm/attachment/list', { params })
}

export const getAttachmentById = (id: number) => {
  return request.get(`/mdm/attachment/${id}`)
}

export const createAttachment = (data: any) => {
  return request.post('/mdm/attachment', data)
}

export const updateAttachment = (id: number, data: any) => {
  return request.put(`/mdm/attachment/${id}`, data)
}

export const deleteAttachment = (id: number) => {
  return request.delete(`/mdm/attachment/${id}`)
}

// 客户收货地址
export const getDeliveryAddressList = (params?: any) => {
  return request.get('/mdm/delivery-address/list', { params })
}

export const getDeliveryAddressById = (id: number) => {
  return request.get(`/mdm/delivery-address/${id}`)
}

export const createDeliveryAddress = (data: any) => {
  return request.post('/mdm/delivery-address', data)
}

export const updateDeliveryAddress = (id: number, data: any) => {
  return request.put(`/mdm/delivery-address/${id}`, data)
}

export const deleteDeliveryAddress = (id: number) => {
  return request.delete(`/mdm/delivery-address/${id}`)
}

// 计量单位
export const getProductUnitList = (params?: any) => {
  return request.get('/mdm/product-unit/list', { params })
}

export const createProductUnit = (data: any) => {
  return request.post('/mdm/product-unit', data)
}

export const updateProductUnit = (id: number, data: any) => {
  return request.put(`/mdm/product-unit/${id}`, data)
}

export const deleteProductUnit = (id: number) => {
  return request.delete(`/mdm/product-unit/${id}`)
}
