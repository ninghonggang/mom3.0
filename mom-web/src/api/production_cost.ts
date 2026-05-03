import request from '@/utils/request'

// 生产成本管理
export const getProductionCostList = (params?: any) => {
  return request.get('/production/cost/list', { params })
}

export const createProductionCost = (data: any) => {
  return request.post('/production/cost', data)
}

export const getProductionCostSummary = (orderId: number) => {
  return request.get('/production/cost/summary', { params: { order_id: orderId } })
}

export const deleteProductionCost = (id: number) => {
  return request.delete(`/production/cost/${id}`)
}