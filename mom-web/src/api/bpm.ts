import request from '@/utils/request'

// ==================== 流程模型 ====================
export const getProcessModelList = (params?: any) =>
  request.get('/bpm/process/list', { params })

export const getProcessModel = (id: number) =>
  request.get(`/bpm/process/${id}`)

export const createProcessModel = (data: any) =>
  request.post('/bpm/process', data)

export const updateProcessModel = (id: number, data: any) =>
  request.put(`/bpm/process/${id}`, data)

export const deleteProcessModel = (id: number) =>
  request.delete(`/bpm/process/${id}`)

export const publishProcessModel = (id: number) =>
  request.post(`/bpm/process/${id}/publish`)

// ==================== 节点定义 ====================
export const getNodeList = (modelId: number, params?: any) =>
  request.get(`/bpm/node/list?process_model_id=${modelId}`, { params })

export const createNode = (modelId: number, data: any) =>
  request.post(`/bpm/node`, { ...data, process_model_id: modelId })

export const updateNode = (id: number, data: any) =>
  request.put(`/bpm/node/${id}`, data)

export const deleteNode = (id: number) =>
  request.delete(`/bpm/node/${id}`)

// ==================== 连线定义 ====================
export const getFlowList = (modelId: number, params?: any) =>
  request.get(`/bpm/flow/list?process_model_id=${modelId}`, { params })

export const createFlow = (modelId: number, data: any) =>
  request.post(`/bpm/flow`, { ...data, process_model_id: modelId })

export const updateFlow = (id: number, data: any) =>
  request.put(`/bpm/flow/${id}`, data)

export const deleteFlow = (id: number) =>
  request.delete(`/bpm/flow/${id}`)

// ==================== 表单定义 ====================
export const getFormDefinitionList = (params?: any) =>
  request.get('/bpm/form/list', { params })

export const getFormDefinition = (id: number) =>
  request.get(`/bpm/form/${id}`)

export const createFormDefinition = (data: any) =>
  request.post('/bpm/form', data)

export const updateFormDefinition = (id: number, data: any) =>
  request.put(`/bpm/form/${id}`, data)

export const deleteFormDefinition = (id: number) =>
  request.delete(`/bpm/form/${id}`)

// ==================== 表单字段 ====================
export const getFormFieldList = (formId: number, params?: any) =>
  request.get(`/bpm/field/list?form_definition_id=${formId}`, { params })

export const createFormField = (formId: number, data: any) =>
  request.post(`/bpm/field`, { ...data, form_definition_id: formId })

export const updateFormField = (id: number, data: any) =>
  request.put(`/bpm/field/${id}`, data)

export const deleteFormField = (id: number) =>
  request.delete(`/bpm/field/${id}`)

// ==================== 流程实例 ====================
export const getProcessInstanceList = (params?: any) =>
  request.get('/bpm/instance/list', { params })

export const getProcessInstance = (id: number) =>
  request.get(`/bpm/instance/${id}`)

export const createProcessInstance = (data: any) =>
  request.post('/bpm/instance/start', data)

export const cancelProcessInstance = (id: number) =>
  request.post(`/bpm/instance/${id}/cancel`)

export const terminateProcessInstance = (id: number) =>
  request.post(`/bpm/instance/${id}/terminate`)

// ==================== 任务实例 ====================
export const getTaskList = (params?: any) =>
  request.get('/bpm/instance/task/list', { params })

export const getTask = (id: number) =>
  request.get(`/bpm/instance/task/${id}`)

export const approveTask = (id: number, data?: any) =>
  request.post(`/bpm/instance/task/${id}/approve`, data)

export const rejectTask = (id: number, data?: any) =>
  request.post(`/bpm/instance/task/${id}/reject`, data)

// ==================== 委托记录 ====================
export const getDelegateList = (params?: any) =>
  request.get('/bpm/delegate/list', { params })

export const createDelegate = (data: any) =>
  request.post('/bpm/delegate', data)

export const updateDelegate = (id: number, data: any) =>
  request.put(`/bpm/delegate/${id}`, data)

export const deleteDelegate = (id: number) =>
  request.delete(`/bpm/delegate/${id}`)

// ==================== 审批记录 ====================
export const getApprovalRecordList = (taskId: number, params?: any) =>
  request.get(`/bpm/instance/approve/records`, { params })
