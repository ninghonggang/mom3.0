import request from '@/utils/request'

// 追溯查询
export const traceBySerial = (serialNumber: string) => {
  return request.get('/trace/serial', { params: { serial_number: serialNumber } })
}

export const traceByBatch = (batchNo: string) => {
  return request.get('/trace/batch', { params: { batch_no: batchNo } })
}

export const traceByOrder = (orderId: number) => {
  return request.get(`/trace/order/${orderId}`)
}

// 序列号管理
export const getSerialNumberList = (params?: any) => {
  return request.get('/trace/serial/list', { params })
}

export const createSerialNumber = (data: any) => {
  return request.post('/trace/serial', data)
}

// 安东呼叫
export const getAndonCallList = (params?: any) => {
  return request.get('/andon/calls/list', { params })
}

export const getAndonCallDetail = (id: number) => {
  return request.get(`/andon/calls/${id}`)
}

export const createAndonCall = (data: any) => {
  return request.post('/andon/calls', data)
}

export const responseAndonCall = (id: number, data?: any) => {
  return request.put(`/andon/calls/${id}/respond`, data)
}

export const resolveAndonCall = (id: number, data?: any) => {
  return request.put(`/andon/calls/${id}/resolve`, data)
}

export const escalateAndonCall = (id: number, data?: any) => {
  return request.put(`/andon/calls/${id}/escalate`, data)
}

export const getAndonStats = (params?: any) => {
  return request.get('/andon/statistics', { params })
}

// 升级规则
export const getEscalationRuleList = (params?: any) => {
  return request.get('/andon/escalation-rules/list', { params })
}

export const getEscalationRuleDetail = (id: number) => {
  return request.get(`/andon/escalation-rules/${id}`)
}

export const createEscalationRule = (data: any) => {
  return request.post('/andon/escalation-rules', data)
}

export const updateEscalationRule = (id: number, data: any) => {
  return request.put(`/andon/escalation-rules/${id}`, data)
}

export const deleteEscalationRule = (id: number) => {
  return request.delete(`/andon/escalation-rules/${id}`)
}

// 数据采集
export const getDataCollectionList = (params?: any) => {
  return request.get('/datacollect/list', { params })
}

export const collectData = (data: any) => {
  return request.post('/datacollect', data)
}

export const getRealTimeData = (equipmentId: number) => {
  return request.get(`/datacollect/equipment/${equipmentId}/realtime`)
}

// 能源管理
export const getEnergyRecordList = (params?: any) => {
  return request.get('/energy/record/list', { params })
}

export const createEnergyRecord = (data: any) => {
  return request.post('/energy/record', data)
}

export const getEnergyStats = (params: any) => {
  return request.get('/energy/stats', { params })
}

export const getEnergyTrend = (params: any) => {
  return request.get('/energy/trend', { params })
}

// 物料追溯
export const getTraceList = (params?: any) => {
  return request.get('/trace/material/list', { params })
}

// 事件日志
export const eventApi = {
  list: (params?: any) => {
    return request.get('/event/log/list', { params })
  },
  get: (id: number) => {
    return request.get(`/event/log/${id}`)
  }
}
