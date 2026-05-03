import request from '@/utils/request'

// VMI库存管理
export const getVmiVendorList = (params?: any) => {
  return request.get('/wms/vmi/vendor/page', { params })
}

export const getVmiVendor = (id: number) => {
  return request.get(`/wms/vmi/vendor/${id}`)
}

export const createVmiVendor = (data: any) => {
  return request.post('/wms/vmi/vendor', data)
}

export const updateVmiVendor = (id: number, data: any) => {
  return request.put(`/wms/vmi/vendor/${id}`, data)
}

export const deleteVmiVendor = (id: number) => {
  return request.delete(`/wms/vmi/vendor/${id}`)
}

export const getVmiMaterialList = (params?: any) => {
  return request.get('/wms/vmi/material/page', { params })
}

export const getVmiMaterial = (id: number) => {
  return request.get(`/wms/vmi/material/${id}`)
}

export const getVmiTransactionList = (params?: any) => {
  return request.get('/wms/vmi/transaction/page', { params })
}

export const consumeVmi = (data: any) => {
  return request.post('/wms/vmi/consume', data)
}

export const replenishVmi = (data: any) => {
  return request.post('/wms/vmi/replenish', data)
}