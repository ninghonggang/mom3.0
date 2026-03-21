import request from '@/utils/request'

// IQC检验
export const getIQCList = (params?: any) => {
  return request.get('/quality/iqc/list', { params })
}

export const createIQC = (data: any) => {
  return request.post('/quality/iqc', data)
}

export const updateIQC = (id: number, data: any) => {
  return request.put(`/quality/iqc/${id}`, data)
}

// IPQC检验
export const getIPQCList = (params?: any) => {
  return request.get('/quality/ipqc/list', { params })
}

export const createIPQC = (data: any) => {
  return request.post('/quality/ipqc', data)
}

// FQC检验
export const getFQCList = (params?: any) => {
  return request.get('/quality/fqc/list', { params })
}

export const createFQC = (data: any) => {
  return request.post('/quality/fqc', data)
}

// OQC检验
export const getOQCList = (params?: any) => {
  return request.get('/quality/oqc/list', { params })
}

export const createOQC = (data: any) => {
  return request.post('/quality/oqc', data)
}

// 不良品记录
export const getDefectRecordList = (params?: any) => {
  return request.get('/quality/defect/list', { params })
}

export const createDefectRecord = (data: any) => {
  return request.post('/quality/defect', data)
}

export const handleDefect = (id: number, data: any) => {
  return request.put(`/quality/defect/${id}/handle`, data)
}

// NCR
export const getNCRList = (params?: any) => {
  return request.get('/quality/ncr/list', { params })
}

export const createNCR = (data: any) => {
  return request.post('/quality/ncr', data)
}

export const updateNCR = (id: number, data: any) => {
  return request.put(`/quality/ncr/${id}`, data)
}

// SPC数据
export const getSPCData = (params: any) => {
  return request.get('/quality/spc/data', { params })
}

export const getSPCChart = (params: any) => {
  return request.get('/quality/spc/chart', { params })
}

// 不良品代码
export const getDefectCodeList = (params?: any) => {
  return request.get('/quality/defect-code/list', { params })
}

export const createDefectCode = (data: any) => {
  return request.post('/quality/defect-code', data)
}

export const updateDefectCode = (id: number, data: any) => {
  return request.put(`/quality/defect-code/${id}`, data)
}

export const deleteDefectCode = (id: number) => {
  return request.delete(`/quality/defect-code/${id}`)
}
