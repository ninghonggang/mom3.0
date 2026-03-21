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
  return request.get('/mdm/shift/list', { params })
}

export const createShift = (data: any) => {
  return request.post('/mdm/shift', data)
}

export const updateShift = (id: number, data: any) => {
  return request.put(`/mdm/shift/${id}`, data)
}

export const deleteShift = (id: number) => {
  return request.delete(`/mdm/shift/${id}`)
}
