import request from '@/utils/request'

// IDOC管理
export const getIdocList = (params?: any) => {
  return request.get('/integration/idoc/page', { params })
}

export const getIdoc = (id: number) => {
  return request.get(`/integration/idoc/${id}`)
}

export const receiveIdoc = (data: any) => {
  return request.post('/integration/idoc/receive', data)
}

export const sendIdoc = (data: any) => {
  return request.post('/integration/idoc/send', data)
}

export const retryIdoc = (id: number) => {
  return request.post(`/integration/idoc/${id}/retry`)
}

export const getIdocConfigs = () => {
  return request.get('/integration/idoc/configs')
}