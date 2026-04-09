import request from '@/utils/request'

// IQC检验
export const getIQCList = (params?: any) => request.get('/quality/iqc/list', { params })
export const createIQC = (data: any) => request.post('/quality/iqc', data)
export const updateIQC = (id: number, data: any) => request.put(`/quality/iqc/${id}`, data)
export const deleteIQC = (id: number) => request.delete(`/quality/iqc/${id}`)

// IPQC检验
export const getIPQCList = (params?: any) => request.get('/quality/ipqc/list', { params })
export const createIPQC = (data: any) => request.post('/quality/ipqc', data)
export const updateIPQC = (id: number, data: any) => request.put(`/quality/ipqc/${id}`, data)
export const deleteIPQC = (id: number) => request.delete(`/quality/ipqc/${id}`)

// FQC检验
export const getFQCList = (params?: any) => request.get('/quality/fqc/list', { params })
export const createFQC = (data: any) => request.post('/quality/fqc', data)
export const updateFQC = (id: number, data: any) => request.put(`/quality/fqc/${id}`, data)
export const deleteFQC = (id: number) => request.delete(`/quality/fqc/${id}`)

// OQC检验
export const getOQCList = (params?: any) => request.get('/quality/oqc/list', { params })
export const createOQC = (data: any) => request.post('/quality/oqc', data)
export const updateOQC = (id: number, data: any) => request.put(`/quality/oqc/${id}`, data)
export const deleteOQC = (id: number) => request.delete(`/quality/oqc/${id}`)

// 不良品记录
export const getDefectRecordList = (params?: any) => request.get('/quality/defect/list', { params })
export const createDefectRecord = (data: any) => request.post('/quality/defect', data)
export const updateDefectRecord = (id: number, data: any) => request.put(`/quality/defect/${id}`, data)
export const handleDefect = (id: number, data: any) => request.put(`/quality/defect/${id}/handle`, data)
export const deleteDefectRecord = (id: number) => request.delete(`/quality/defect/${id}`)

// NCR
export const getNCRList = (params?: any) => request.get('/quality/ncr/list', { params })
export const createNCR = (data: any) => request.post('/quality/ncr', data)
export const updateNCR = (id: number, data: any) => request.put(`/quality/ncr/${id}`, data)
export const deleteNCR = (id: number) => request.delete(`/quality/ncr/${id}`)

// SPC数据
export const getSPCData = (params: any) => request.get('/quality/spc/list', { params })
export const getSPCChart = (params: any) => request.get('/quality/spc/chart', { params })

// 不良品代码
export const getDefectCodeList = (params?: any) => request.get('/quality/defect-code/list', { params })
export const createDefectCode = (data: any) => request.post('/quality/defect-code', data)
export const updateDefectCode = (id: number, data: any) => request.put(`/quality/defect-code/${id}`, data)
export const deleteDefectCode = (id: number) => request.delete(`/quality/defect-code/${id}`)
