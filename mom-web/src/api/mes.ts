import request from '@/utils/request'

const API_BASE = '/mes'

// ========== 月计划 API ==========

// 获取月计划列表
export const getMonthPlanList = (params?: any) => {
  return request.get(`${API_BASE}/month-plans`, { params })
}

// 获取月计划详情
export const getMonthPlan = (id: number) => {
  return request.get(`${API_BASE}/month-plans/${id}`)
}

// 创建月计划
export const createMonthPlan = (data: any) => {
  return request.post(`${API_BASE}/month-plans`, data)
}

// 更新月计划
export const updateMonthPlan = (id: number, data: any) => {
  return request.put(`${API_BASE}/month-plans/${id}`, data)
}

// 删除月计划
export const deleteMonthPlan = (id: number) => {
  return request.delete(`${API_BASE}/month-plans/${id}`)
}

// 提交月计划
export const submitMonthPlan = (id: number) => {
  return request.post(`${API_BASE}/month-plans/${id}/submit`)
}

// 审核月计划
export const approveMonthPlan = (id: number, data?: { comment?: string }) => {
  return request.post(`${API_BASE}/month-plans/${id}/approve`, data)
}

// 下达月计划
export const releaseMonthPlan = (id: number) => {
  return request.post(`${API_BASE}/month-plans/${id}/release`)
}

// 关闭月计划
export const closeMonthPlan = (id: number) => {
  return request.post(`${API_BASE}/month-plans/${id}/close`)
}

// 取消月计划
export const cancelMonthPlan = (id: number, data?: { comment?: string }) => {
  return request.post(`${API_BASE}/month-plans/${id}/cancel`, data)
}

// 获取月计划审核记录
export const getMonthPlanAudits = (id: number) => {
  return request.get(`${API_BASE}/month-plans/${id}/audits`)
}

// ========== 日计划 API ==========

// 获取日计划列表
export const getDayPlanList = (params?: any) => {
  return request.get(`${API_BASE}/day-plans`, { params })
}

// 获取日计划详情
export const getDayPlan = (id: number) => {
  return request.get(`${API_BASE}/day-plans/${id}`)
}

// 创建日计划
export const createDayPlan = (data: any) => {
  return request.post(`${API_BASE}/day-plans`, data)
}

// 更新日计划
export const updateDayPlan = (id: number, data: any) => {
  return request.put(`${API_BASE}/day-plans/${id}`, data)
}

// 删除日计划
export const deleteDayPlan = (id: number) => {
  return request.delete(`${API_BASE}/day-plans/${id}`)
}

// 发布日计划
export const publishDayPlan = (id: number) => {
  return request.post(`${API_BASE}/day-plans/${id}/publish`)
}

// 完成日计划
export const completeDayPlan = (id: number) => {
  return request.post(`${API_BASE}/day-plans/${id}/complete`)
}

// 终止日计划
export const terminateDayPlan = (id: number) => {
  return request.post(`${API_BASE}/day-plans/${id}/terminate`)
}

// 齐套检查
export const kitCheckDayPlan = (id: number) => {
  return request.post(`${API_BASE}/day-plans/${id}/kit-check`)
}
