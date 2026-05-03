import request from '@/utils/request'

// 客户信用管理
export const getCustomerCreditList = (params?: any) => {
  return request.get('/scp/customer-credit/page', { params })
}

export const getCustomerCredit = (id: number) => {
  return request.get(`/scp/customer-credit/${id}`)
}

export const getCustomerCreditByCustomer = (customerId: number) => {
  return request.get(`/scp/customer-credit/customer/${customerId}`)
}

export const createCustomerCredit = (data: any) => {
  return request.post('/scp/customer-credit', data)
}

export const updateCustomerCredit = (id: number, data: any) => {
  return request.put(`/scp/customer-credit/${id}`, data)
}

export const updateCustomerCreditUsedCredit = (id: number, usedCredit: number) => {
  return request.put(`/scp/customer-credit/${id}/used-credit`, { used_credit: usedCredit })
}

export const setCustomerCreditBlacklist = (id: number, blacklist: boolean) => {
  return request.put(`/scp/customer-credit/${id}/blacklist`, { blacklist })
}

export const freezeCustomerCredit = (id: number) => {
  return request.put(`/scp/customer-credit/${id}/freeze`)
}

export const unfreezeCustomerCredit = (id: number) => {
  return request.put(`/scp/customer-credit/${id}/unfreeze`)
}

export const deleteCustomerCredit = (id: number) => {
  return request.delete(`/scp/customer-credit/${id}`)
}