import request from '@/utils/request'

// 设备台账
export const getEquipmentList = (params?: any) => {
  return request.get('/equipment/list', { params })
}

export const getEquipmentById = (id: number) => {
  return request.get(`/equipment/${id}`)
}

export const createEquipment = (data: any) => {
  return request.post('/equipment', data)
}

export const updateEquipment = (id: number, data: any) => {
  return request.put(`/equipment/${id}`, data)
}

export const deleteEquipment = (id: number) => {
  return request.delete(`/equipment/${id}`)
}

export const getEquipmentStatus = () => {
  return request.get('/equipment/status')
}

// 设备点检
export const getEquipmentCheckList = (params?: any) => {
  return request.get('/equipment/check/list', { params })
}

export const createEquipmentCheck = (data: any) => {
  return request.post('/equipment/check', data)
}

export const updateEquipmentCheck = (id: number, data: any) => {
  return request.put(`/equipment/check/${id}`, data)
}

export const deleteEquipmentCheck = (id: number) => {
  return request.delete(`/equipment/check/${id}`)
}

// 设备保养
export const getEquipmentMaintenanceList = (params?: any) => {
  return request.get('/equipment/maintenance/list', { params })
}

export const createEquipmentMaintenance = (data: any) => {
  return request.post('/equipment/maintenance', data)
}

export const updateEquipmentMaintenance = (id: number, data: any) => {
  return request.put(`/equipment/maintenance/${id}`, data)
}

export const deleteEquipmentMaintenance = (id: number) => {
  return request.delete(`/equipment/maintenance/${id}`)
}

// 设备维修
export const getEquipmentRepairList = (params?: any) => {
  return request.get('/equipment/repair/list', { params })
}

export const createEquipmentRepair = (data: any) => {
  return request.post('/equipment/repair', data)
}

export const startRepair = (id: number) => {
  return request.put(`/equipment/repair/${id}/start`)
}

export const completeRepair = (id: number) => {
  return request.put(`/equipment/repair/${id}/complete`)
}

export const updateEquipmentRepair = (id: number, data: any) => {
  return request.put(`/equipment/repair/${id}`, data)
}

export const deleteEquipmentRepair = (id: number) => {
  return request.delete(`/equipment/repair/${id}`)
}

// OEE设备综合效率
export const getOEEList = (params?: any) => {
  return request.get('/equipment/oee/list', { params })
}

export const getOEEById = (id: number) => {
  return request.get(`/equipment/oee/${id}`)
}

export const calculateOEE = (data: any) => {
  return request.post('/equipment/oee/calculate', data)
}

export const getOEEChart = (params?: any) => {
  return request.get('/equipment/oee/chart', { params })
}

export const deleteOEE = (id: number) => {
  return request.delete(`/equipment/oee/${id}`)
}

// 备件管理
export const getSparePartList = (params?: any) => {
  return request.get('/equipment/spare/list', { params })
}

export const createSparePart = (data: any) => {
  return request.post('/equipment/spare', data)
}

export const updateSparePart = (id: number, data: any) => {
  return request.put(`/equipment/spare/${id}`, data)
}

export const deleteSparePart = (id: number) => {
  return request.delete(`/equipment/spare/${id}`)
}
