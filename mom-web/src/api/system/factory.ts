import request from '@/utils/request'

// 工厂管理
export const getFactoryList = (params?: any) => {
  return request.get('/system/factory/list', { params })
}

export const getFactory = (id: number) => {
  return request.get(`/system/factory/${id}`)
}

export const createFactory = (data: any) => {
  return request.post('/system/factory', data)
}

export const updateFactory = (id: number, data: any) => {
  return request.put(`/system/factory/${id}`, data)
}

export const deleteFactory = (id: number) => {
  return request.delete(`/system/factory/${id}`)
}

export const setDefaultFactory = (id: number) => {
  return request.put(`/system/factory/default/${id}`)
}

export const getCurrentFactory = () => {
  return request.get('/system/factory/current')
}

export const setCurrentFactory = (factoryId: number) => {
  return request.put('/system/factory/current', { factory_id: factoryId })
}