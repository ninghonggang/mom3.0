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

// 设备保养
export const getEquipmentMaintenanceList = (params?: any) => {
  return request.get('/equipment/maintenance/list', { params })
}

export const createEquipmentMaintenance = (data: any) => {
  return request.post('/equipment/maintenance', data)
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

// OEE
export const getEquipmentOEE = (params: any) => {
  return request.get('/equipment/oee', { params })
}

export const getOEERealTime = (equipmentId: number) => {
  return request.get(`/equipment/${equipmentId}/oee/realtime`)
}
