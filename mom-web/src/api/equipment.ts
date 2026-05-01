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

// 点检标准
export const getInspectionTemplateList = (params?: any) => request.get('/equipment/inspection/templates/list', { params })
export const getInspectionTemplate = (id: number) => request.get(`/equipment/inspection/templates/${id}`)
export const createInspectionTemplate = (data: any) => request.post('/equipment/inspection/templates', data)
export const updateInspectionTemplate = (id: number, data: any) => request.put(`/equipment/inspection/templates/${id}`, data)
export const deleteInspectionTemplate = (id: number) => request.delete(`/equipment/inspection/templates/${id}`)

// 点检计划
export const getInspectionPlanList = (params?: any) => request.get('/equipment/inspection/plan/list', { params })
export const getInspectionPlan = (id: number) => request.get(`/equipment/inspection/plan/${id}`)
export const createInspectionPlan = (data: any) => request.post('/equipment/inspection/plan', data)
export const updateInspectionPlan = (id: number, data: any) => request.put(`/equipment/inspection/plan/${id}`, data)
export const assignInspectionPlan = (id: number, data: any) => request.put(`/equipment/inspection/plan/${id}/assign`, data)
export const cancelInspectionPlan = (id: number) => request.delete(`/equipment/inspection/plan/${id}`)

// 点检记录
export const getInspectionRecordList = (params?: any) => request.get('/equipment/inspection/records/list', { params })
export const getInspectionRecord = (id: number) => request.get(`/equipment/inspection/records/${id}`)
export const startInspection = (data: any) => request.post('/equipment/inspection/records', data)
export const completeInspection = (id: number, data: any) => request.put(`/equipment/inspection/records/${id}/complete`, data)
export const deleteInspectionRecord = (id: number) => request.delete(`/equipment/inspection/records/${id}`)

// 点检异常
export const getInspectionDefectList = (params?: any) => request.get('/equipment/inspection/defects/list', { params })
export const getInspectionDefect = (id: number) => request.get(`/equipment/inspection/defects/${id}`)
export const createInspectionDefect = (data: any) => request.post('/equipment/inspection/defects', data)
export const assignInspectionDefect = (id: number, data: any) => request.put(`/equipment/inspection/defects/${id}/assign`, data)
export const resolveInspectionDefect = (id: number, data: any) => request.put(`/equipment/inspection/defects/${id}/resolve`, data)

// 量检具管理
export const getGaugeList = (params?: any) => {
  return request.get('/equipment/gauge/list', { params })
}

export const getGaugeById = (id: number) => {
  return request.get(`/equipment/gauge/${id}`)
}

export const createGauge = (data: any) => {
  return request.post('/equipment/gauge', data)
}

export const updateGauge = (id: number, data: any) => {
  return request.put(`/equipment/gauge/${id}`, data)
}

export const deleteGauge = (id: number) => {
  return request.delete(`/equipment/gauge/${id}`)
}
