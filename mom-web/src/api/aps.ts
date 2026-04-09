import request from '@/utils/request'

// MPS
export const getMPSList = (params?: any) => {
  return request.get('/aps/mps/list', { params })
}

export const createMPS = (data: any) => {
  return request.post('/aps/mps', data)
}

export const updateMPS = (id: number, data: any) => {
  return request.put(`/aps/mps/${id}`, data)
}

// MRP
export const getMRPList = (params?: any) => {
  return request.get('/aps/mrp/list', { params })
}

export const getMRP = (id: number) => {
  return request.get(`/aps/mrp/${id}`)
}

export const createMRP = (data: any) => {
  return request.post('/aps/mrp', data)
}

export const updateMRP = (id: number, data: any) => {
  return request.put(`/aps/mrp/${id}`, data)
}

export const deleteMRP = (id: number) => {
  return request.delete(`/aps/mrp/${id}`)
}

export const calculateMRP = (id: number) => {
  return request.put(`/aps/mrp/${id}/calculate`)
}

export const runMRP = (data: { id: number; plan_month?: string }) => {
  return request.post('/aps/mrp/run', data)
}

export const getMRPResults = (id: number) => {
  return request.get(`/aps/mrp/${id}/results`)
}

export const getMRPShortage = (id: number) => {
  return request.get(`/aps/mrp/${id}/shortage`)
}

export const getMRPPurchaseSuggestion = (id: number) => {
  return request.get(`/aps/mrp/${id}/purchase-suggestion`)
}

// 排程
export const getSchedulePlanList = (params?: any) => {
  return request.get('/aps/schedule/list', { params })
}

export const getScheduleList = getSchedulePlanList

export const createSchedulePlan = (data: any) => {
  return request.post('/aps/schedule', data)
}

export const createSchedule = createSchedulePlan

export const runSchedule = (id: number) => {
  return request.put(`/aps/schedule/${id}/execute`)
}

export const executeSchedule = runSchedule

export const deleteSchedule = (id: number) => {
  return request.delete(`/aps/schedule/${id}`)
}

export const getScheduleGantt = (id: number) => {
  return request.get(`/aps/schedule/${id}/gantt`)
}

// 工作中心（APS模块）
export const getWorkCenterList = (params?: any) => {
  return request.get('/aps/workcenter/list', { params })
}

export const getWorkCenter = (id: number) => {
  return request.get(`/aps/workcenter/${id}`)
}

export const createWorkCenter = (data: any) => {
  return request.post('/aps/workcenter', data)
}

export const updateWorkCenter = (id: number, data: any) => {
  return request.put(`/aps/workcenter/${id}`, data)
}

export const deleteWorkCenter = (id: number) => {
  return request.delete(`/aps/workcenter/${id}`)
}

// 资源（APS模块）
export const getResourceList = (params?: any) => {
  return request.get('/aps/resource/list', { params })
}

export const createResource = (data: any) => {
  return request.post('/aps/resource', data)
}

// 甘特图
export const getGanttData = (params: any) => {
  return request.get('/aps/gantt', { params })
}

// 拖拽更新排程结果
export const dragUpdateSchedule = (data: {
  result_id: number
  line_id: number
  station_id: number
  plan_start_time: number
  plan_end_time: number
}) => {
  return request.put('/aps/schedule/drag-update', data)
}

// 获取排程结果
export const getScheduleResults = (id: number) => {
  return request.get(`/aps/schedule/${id}/results`)
}
