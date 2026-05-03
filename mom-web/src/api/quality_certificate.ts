import request from '@/utils/request'

// 质量证书管理
export const getQualityCertificateList = (params?: any) => {
  return request.get('/quality/certificate/page', { params })
}

export const getQualityCertificate = (id: number) => {
  return request.get(`/quality/certificate/${id}`)
}

export const createQualityCertificate = (data: any) => {
  return request.post('/quality/certificate', data)
}

export const updateQualityCertificate = (id: number, data: any) => {
  return request.put(`/quality/certificate/${id}`, data)
}

export const deleteQualityCertificate = (id: number) => {
  return request.delete(`/quality/certificate/${id}`)
}