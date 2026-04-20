import request from '@/utils/request'

export function getProductionIssueList(params: any) {
  return request.get('/api/production-issue/production-issues', { params })
}

export function getProductionIssue(id: number) {
  return request.get(`/api/production-issue/production-issues/${id}`)
}

export function createProductionIssue(data: any) {
  return request.post('/api/production-issue/production-issues', data)
}

export function updateProductionIssue(id: number, data: any) {
  return request.put(`/api/production-issue/production-issues/${id}`, data)
}

export function deleteProductionIssue(id: number) {
  return request.delete(`/api/production-issue/production-issues/${id}`)
}

export function submitProductionIssue(id: number, data: any) {
  return request.post(`/api/production-issue/production-issues/${id}/submit`, data)
}

export function startPickProductionIssue(id: number) {
  return request.post(`/api/production-issue/production-issues/${id}/start-pick`)
}

export function confirmPickProductionIssue(id: number, data: any) {
  return request.post(`/api/production-issue/production-issues/${id}/confirm-pick`, data)
}

export function issueProductionIssue(id: number) {
  return request.post(`/api/production-issue/production-issues/${id}/issue`)
}

export function cancelProductionIssue(id: number) {
  return request.post(`/api/production-issue/production-issues/${id}/cancel`)
}
