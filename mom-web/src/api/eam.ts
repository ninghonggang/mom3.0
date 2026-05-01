import request from '@/utils/request'

// 设备资产
export const getAssetList = (params?: any) => {
  return request.get('/eam/asset/list', { params })
}

export const createAsset = (data: any) => {
  return request.post('/eam/asset', data)
}

export const updateAsset = (id: number, data: any) => {
  return request.put(`/eam/asset/${id}`, data)
}

export const deleteAsset = (id: number) => {
  return request.delete(`/eam/asset/${id}`)
}
