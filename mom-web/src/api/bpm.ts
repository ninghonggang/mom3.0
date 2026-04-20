import request from '@/utils/request'

// ==================== 流程模型 ====================
export const getProcessModelList = (params?: any) =>
  request.get('/bpm/models', { params })

export const getProcessModel = (id: number) =>
  request.get(`/bpm/models/${id}`)

export const createProcessModel = (data: any) =>
  request.post('/bpm/models', data)

export const updateProcessModel = (id: number, data: any) =>
  request.put(`/bpm/models/${id}`, data)

export const deleteProcessModel = (id: number) =>
  request.delete(`/bpm/models/${id}`)

export const publishProcessModel = (id: number) =>
  request.post(`/bpm/models/${id}/publish`)

// ==================== 节点定义 ====================
export const getNodeList = (modelId: number, params?: any) =>
  request.get(`/bpm/models/${modelId}/nodes`, { params })

export const createNode = (modelId: number, data: any) =>
  request.post(`/bpm/models/${modelId}/nodes`, data)

export const updateNode = (id: number, data: any) =>
  request.put(`/bpm/nodes/${id}`, data)

export const deleteNode = (id: number) =>
  request.delete(`/bpm/nodes/${id}`)

// ==================== 连线定义 ====================
export const getFlowList = (modelId: number, params?: any) =>
  request.get(`/bpm/models/${modelId}/flows`, { params })

export const createFlow = (modelId: number, data: any) =>
  request.post(`/bpm/models/${modelId}/flows`, data)

export const updateFlow = (id: number, data: any) =>
  request.put(`/bpm/flows/${id}`, data)

export const deleteFlow = (id: number) =>
  request.delete(`/bpm/flows/${id}`)

// ==================== 表单定义 ====================
export const getFormDefinitionList = (params?: any) =>
  request.get('/bpm/forms', { params })

export const getFormDefinition = (id: number) =>
  request.get(`/bpm/forms/${id}`)

export const createFormDefinition = (data: any) =>
  request.post('/bpm/forms', data)

export const updateFormDefinition = (id: number, data: any) =>
  request.put(`/bpm/forms/${id}`, data)

export const deleteFormDefinition = (id: number) =>
  request.delete(`/bpm/forms/${id}`)

// ==================== 表单字段 ====================
export const getFormFieldList = (formId: number, params?: any) =>
  request.get(`/bpm/forms/${formId}/fields`, { params })

export const createFormField = (formId: number, data: any) =>
  request.post(`/bpm/forms/${formId}/fields`, data)

export const updateFormField = (id: number, data: any) =>
  request.put(`/bpm/fields/${id}`, data)

export const deleteFormField = (id: number) =>
  request.delete(`/bpm/fields/${id}`)

// ==================== 流程实例 ====================
export const getProcessInstanceList = (params?: any) =>
  request.get('/bpm/instances', { params })

export const getProcessInstance = (id: number) =>
  request.get(`/bpm/instances/${id}`)

export const createProcessInstance = (data: any) =>
  request.post('/bpm/instances', data)

export const cancelProcessInstance = (id: number) =>
  request.post(`/bpm/instances/${id}/cancel`)

export const terminateProcessInstance = (id: number) =>
  request.post(`/bpm/instances/${id}/terminate`)

// ==================== 任务实例 ====================
export const getTaskList = (params?: any) =>
  request.get('/bpm/tasks', { params })

export const getTask = (id: number) =>
  request.get(`/bpm/tasks/${id}`)

export const approveTask = (id: number, data?: any) =>
  request.post(`/bpm/tasks/${id}/approve`, data)

export const rejectTask = (id: number, data?: any) =>
  request.post(`/bpm/tasks/${id}/reject`, data)

// ==================== 委托记录 ====================
export const getDelegateList = (params?: any) =>
  request.get('/bpm/delegates', { params })

export const createDelegate = (data: any) =>
  request.post('/bpm/delegates', data)

export const updateDelegate = (id: number, data: any) =>
  request.put(`/bpm/delegates/${id}`, data)

export const deleteDelegate = (id: number) =>
  request.delete(`/bpm/delegates/${id}`)

// ==================== 审批记录 ====================
export const getApprovalRecordList = (taskId: number, params?: any) =>
  request.get(`/bpm/tasks/${taskId}/records`, { params })
