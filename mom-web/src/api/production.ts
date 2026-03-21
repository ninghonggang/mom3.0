import request from '@/utils/request'

// 生产工单
export const getProductionOrderList = (params?: any) => {
  return request.get('/production/order/list', { params })
}

export const getProductionOrderById = (id: number) => {
  return request.get(`/production/order/${id}`)
}

export const createProductionOrder = (data: any) => {
  return request.post('/production/order', data)
}

export const updateProductionOrder = (id: number, data: any) => {
  return request.put(`/production/order/${id}`, data)
}

export const deleteProductionOrder = (id: number) => {
  return request.delete(`/production/order/${id}`)
}

export const startProductionOrder = (id: number) => {
  return request.put(`/production/order/${id}/start`)
}

export const completeProductionOrder = (id: number) => {
  return request.put(`/production/order/${id}/complete`)
}

export const cancelProductionOrder = (id: number) => {
  return request.put(`/production/order/${id}/cancel`)
}

// 生产报工
export const getProductionReportList = (params?: any) => {
  return request.get('/production/report/list', { params })
}

export const createProductionReport = (data: any) => {
  return request.post('/production/report', data)
}

// 派工
export const getDispatchList = (params?: any) => {
  return request.get('/production/dispatch/list', { params })
}

export const createDispatch = (data: any) => {
  return request.post('/production/dispatch', data)
}

export const updateDispatch = (id: number, data: any) => {
  return request.put(`/production/dispatch/${id}`, data)
}

export const startDispatch = (id: number) => {
  return request.put(`/production/dispatch/${id}/start`)
}

export const completeDispatch = (id: number) => {
  return request.put(`/production/dispatch/${id}/complete`)
}

// 销售订单
export const getSalesOrderList = (params?: any) => {
  return request.get('/production/sales-order/list', { params })
}

export const createSalesOrder = (data: any) => {
  return request.post('/production/sales-order', data)
}

export const updateSalesOrder = (id: number, data: any) => {
  return request.put(`/production/sales-order/${id}`, data)
}

export const deleteSalesOrder = (id: number) => {
  return request.delete(`/production/sales-order/${id}`)
}

export const confirmSalesOrder = (id: number) => {
  return request.put(`/production/sales-order/${id}/confirm`)
}
