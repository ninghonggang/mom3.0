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

export const deleteInventory = (id: number) => {
  return request.delete(`/wms/inventory/${id}`)
}

// 收货单
export const getReceiveOrderList = (params?: any) => {
  return request.get('/wms/receive/list', { params })
}

export const getReceiveOrderById = (id: number) => {
  return request.get(`/wms/receive/${id}`)
}

export const createReceiveOrder = (data: any) => {
  return request.post('/wms/receive', data)
}

export const updateReceiveOrder = (id: number, data: any) => {
  return request.put(`/wms/receive/${id}`, data)
}

export const deleteReceiveOrder = (id: number) => {
  return request.delete(`/wms/receive/${id}`)
}

export const receiveConfirm = (id: number) => {
  return request.put(`/wms/receive/${id}/confirm`)
}

// 发货单
export const getDeliveryOrderList = (params?: any) => {
  return request.get('/wms/delivery/list', { params })
}

export const getDeliveryOrderById = (id: number) => {
  return request.get(`/wms/delivery/${id}`)
}

export const createDeliveryOrder = (data: any) => {
  return request.post('/wms/delivery', data)
}

export const updateDeliveryOrder = (id: number, data: any) => {
  return request.put(`/wms/delivery/${id}`, data)
}

export const deleteDeliveryOrder = (id: number) => {
  return request.delete(`/wms/delivery/${id}`)
}

export const deliveryConfirm = (id: number) => {
  return request.put(`/wms/delivery/${id}/confirm`)
}

// 盘点
export const getStockCheckList = (params?: any) => {
  return request.get('/wms/stock-check/list', { params })
}

export const getStockCheckById = (id: number) => {
  return request.get(`/wms/stock-check/${id}`)
}

export const createStockCheck = (data: any) => {
  return request.post('/wms/stock-check', data)
}

export const updateStockCheck = (id: number, data: any) => {
  return request.put(`/wms/stock-check/${id}`, data)
}

export const deleteStockCheck = (id: number) => {
  return request.delete(`/wms/stock-check/${id}`)
}

export const submitStockCheck = (id: number) => {
  return request.post(`/wms/stock-check/${id}/submit`)
}

export const startStockCheck = (id: number) => {
  return request.post(`/wms/stock-check/${id}/start`)
}

export const completeStockCheck = (id: number) => {
  return request.post(`/wms/stock-check/${id}/complete`)
}

export const approveStockCheck = (id: number, data: any) => {
  return request.post(`/wms/stock-check/${id}/approve`, data)
}

export const addStockCheckItem = (data: any) => {
  return request.post('/wms/stock-check/item', data)
}

export const updateStockCheckItem = (id: number, data: any) => {
  return request.put(`/wms/stock-check/item/${id}`, data)
}

export const countStockCheckItem = (checkId: number, itemId: number, data: any) => {
  return request.post(`/wms/stock-check/${checkId}/items/${itemId}/count`, data)
}

export const handleStockCheckVariance = (checkId: number, itemId: number, data: any) => {
  return request.post(`/wms/stock-check/${checkId}/items/${itemId}/handle`, data)
}

export const recountStockCheckItem = (checkId: number, itemId: number, data: any) => {
  return request.post(`/wms/stock-check/${checkId}/items/${itemId}/recount`, data)
}

export const getStockCheckVariance = (id: number) => {
  return request.get(`/wms/stock-check/${id}/variance`)
}

// 数据采集点
export const getDataPointList = (params?: any) => {
  return request.get('/dc/data-point/list', { params })
}

export const getDataPoint = (id: number) => {
  return request.get(`/dc/data-point/${id}`)
}

export const createDataPoint = (data: any) => {
  return request.post('/dc/data-point', data)
}

export const updateDataPoint = (id: number, data: any) => {
  return request.put(`/dc/data-point/${id}`, data)
}

export const deleteDataPoint = (id: number) => {
  return request.delete(`/dc/data-point/${id}`)
}

// 扫描记录
export const getScanLogList = (params?: any) => {
  return request.get('/dc/scan-log/list', { params })
}

export const createScanLog = (data: any) => {
  return request.post('/dc/scan-log/scan', data)
}

// 采集记录
export const getCollectRecordList = (params?: any) => {
  return request.get('/dc/collect-record/list', { params })
}

// ========== 调拨管理 ==========
export const getTransferOrderList = (params?: any) => {
  return request.get('/wms/transfer/list', { params })
}

export const getTransferOrderById = (id: number) => {
  return request.get(`/wms/transfer/${id}`)
}

export const createTransferOrder = (data: any) => {
  return request.post('/wms/transfer', data)
}

export const updateTransferOrder = (id: number, data: any) => {
  return request.put(`/wms/transfer/${id}`, data)
}

export const deleteTransferOrder = (id: number) => {
  return request.delete(`/wms/transfer/${id}`)
}

export const addTransferOrderItem = (data: any) => {
  return request.post('/wms/transfer/item', data)
}

export const submitTransferOrder = (id: number) => {
  return request.post(`/wms/transfer/${id}/submit`)
}

export const approveTransferOrder = (id: number, data: any) => {
  return request.post(`/wms/transfer/${id}/approve`, data)
}

export const startTransferOrder = (id: number) => {
  return request.post(`/wms/transfer/${id}/start`)
}

export const shipTransferOrder = (id: number, data: any) => {
  return request.post(`/wms/transfer/${id}/ship`, data)
}

export const receiveTransferOrder = (id: number, data: any) => {
  return request.post(`/wms/transfer/${id}/receive`, data)
}

export const completeTransferOrder = (id: number) => {
  return request.post(`/wms/transfer/${id}/complete`)
}

export const cancelTransferOrder = (id: number, reason?: string) => {
  return request.post(`/wms/transfer/${id}/cancel`, { reason })
}

export const getTransferOrderTrace = (id: number) => {
  return request.get(`/wms/transfer/${id}/trace`)
}

// 入库管理
export const getInboundList = (params?: any) => {
  return request.get('/wms/inbound/list', { params })
}

export const createInbound = (data: any) => {
  return request.post('/wms/inbound', data)
}

export const updateInbound = (id: number, data: any) => {
  return request.put(`/wms/inbound/${id}`, data)
}

export const deleteInbound = (id: number) => {
  return request.delete(`/wms/inbound/${id}`)
}

// 出库管理
export const getOutboundList = (params?: any) => {
  return request.get('/wms/outbound/list', { params })
}

export const createOutbound = (data: any) => {
  return request.post('/wms/outbound', data)
}

export const updateOutbound = (id: number, data: any) => {
  return request.put(`/wms/outbound/${id}`, data)
}

export const deleteOutbound = (id: number) => {
  return request.delete(`/wms/outbound/${id}`)
}
