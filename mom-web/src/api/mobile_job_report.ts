import request from '@/utils/request'

// 移动端报工
export const getMobileJobReportList = (params?: any) => {
  return request.get('/mes/mobile-job-report/page', { params })
}

export const getMobileJobReport = (id: number) => {
  return request.get(`/mes/mobile-job-report/${id}`)
}

export const createMobileJobReport = (data: any) => {
  return request.post('/mes/mobile-job-report', data)
}

export const confirmMobileJobReport = (id: number) => {
  return request.put(`/mes/mobile-job-report/${id}/confirm`)
}

export const auditMobileJobReport = (id: number) => {
  return request.put(`/mes/mobile-job-report/${id}/audit`)
}

export const getMobilePendingOrders = (params?: any) => {
  return request.get('/mes/mobile-job-report/pending-orders', { params })
}

export const deleteMobileJobReport = (id: number) => {
  return request.delete(`/mes/mobile-job-report/${id}`)
}