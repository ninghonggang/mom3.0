import request from '@/utils/request'

// 仓库
export const getWarehouseList = (params?: any) => {
  return request.get('/wms/warehouse/list', { params })
}

export const createWarehouse = (data: any) => {
  return request.post('/wms/warehouse', data)
}

export const updateWarehouse = (id: number, data: any) => {
  return request.put(`/wms/warehouse/${id}`, data)
}

export const deleteWarehouse = (id: number) => {
  return request.delete(`/wms/warehouse/${id}`)
}

// 库位
export const getLocationList = (params?: any) => {
  return request.get('/wms/location/list', { params })
}

export const createLocation = (data: any) => {
  return request.post('/wms/location', data)
}

export const updateLocation = (id: number, data: any) => {
  return request.put(`/wms/location/${id}`, data)
}

export const deleteLocation = (id: number) => {
  return request.delete(`/wms/location/${id}`)
}

// 库存
export const getInventoryList = (params?: any) => {
  return request.get('/wms/inventory/list', { params })
}

export const getInventoryByMaterial = (materialId: number) => {
  return request.get(`/wms/inventory/material/${materialId}`)
}

export const adjustInventory = (data: any) => {
  return request.post('/wms/inventory/adjust', data)
}

// 收货单
export const getReceiveOrderList = (params?: any) => {
  return request.get('/wms/receive/list', { params })
}

export const createReceiveOrder = (data: any) => {
  return request.post('/wms/receive', data)
}

export const receiveConfirm = (id: number) => {
  return request.put(`/wms/receive/${id}/confirm`)
}

// 发货单
export const getDeliveryOrderList = (params?: any) => {
  return request.get('/wms/delivery/list', { params })
}

export const createDeliveryOrder = (data: any) => {
  return request.post('/wms/delivery', data)
}

export const deliveryConfirm = (id: number) => {
  return request.put(`/wms/delivery/${id}/confirm`)
}

// 盘点
export const getStockCheckList = (params?: any) => {
  return request.get('/wms/stock-check/list', { params })
}

export const createStockCheck = (data: any) => {
  return request.post('/wms/stock-check', data)
}

export const submitStockCheck = (id: number, data: any) => {
  return request.put(`/wms/stock-check/${id}/submit`, data)
}
