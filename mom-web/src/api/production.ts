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

// 生产看板
export const getKanbanDashboard = () => {
  return request.get('/production/kanban/dashboard')
}

// 工单变更记录
export const getOrderChangeList = (params?: any) => {
  return request.get('/production/order-change/list', { params })
}

// 包装条码
export const getPackageList = (params?: any) => {
  return request.get('/production/packages/list', { params })
}

export const getPackage = (id: number) => {
  return request.get(`/production/packages/${id}`)
}

export const createPackage = (data: any) => {
  return request.post('/production/packages/create', data)
}

export const addPackageItem = (data: any) => {
  return request.post('/production/packages/add-item', data)
}

export const sealPackage = (data: any) => {
  return request.post('/production/packages/seal', data)
}

export const deletePackage = (id: number) => {
  return request.delete(`/production/packages/${id}`)
}

// 首末件检验
export const getFirstLastInspectList = (params?: any) => {
  return request.get('/production/first-last-inspect/list', { params })
}

export const getFirstLastInspect = (id: number) => {
  return request.get(`/production/first-last-inspect/${id}`)
}

export const createFirstLastInspect = (data: any) => {
  return request.post('/production/first-last-inspect', data)
}

export const updateFirstLastInspect = (id: number, data: any) => {
  return request.put(`/production/first-last-inspect/${id}`, data)
}

export const deleteFirstLastInspect = (id: number) => {
  return request.delete(`/production/first-last-inspect/${id}`)
}

export const getFirstLastInspectOverdue = () => {
  return request.get('/production/first-last-inspect/overdue')
}

// 电子SOP
export const getElectronicSOPList = (params?: any) => {
  return request.get('/production/electronic-sop/list', { params })
}

export const getElectronicSOP = (id: number) => {
  return request.get(`/production/electronic-sop/${id}`)
}

export const createElectronicSOP = (data: any) => {
  return request.post('/production/electronic-sop', data)
}

export const updateElectronicSOP = (id: number, data: any) => {
  return request.put(`/production/electronic-sop/${id}`, data)
}

export const deleteElectronicSOP = (id: number) => {
  return request.delete(`/production/electronic-sop/${id}`)
}

// 编码规则
export const getCodeRuleList = (params?: any) => {
  return request.get('/production/code-rule/list', { params })
}

export const getCodeRule = (id: number) => {
  return request.get(`/production/code-rule/${id}`)
}

export const createCodeRule = (data: any) => {
  return request.post('/production/code-rule', data)
}

export const updateCodeRule = (id: number, data: any) => {
  return request.put(`/production/code-rule/${id}`, data)
}

export const deleteCodeRule = (id: number) => {
  return request.delete(`/production/code-rule/${id}`)
}

export const generateCode = (ruleCode: string) => {
  return request.get('/production/code-rule/generate', { params: { rule_code: ruleCode } })
}

// 生产指示单
export const getFlowCardList = (params?: any) => {
  return request.get('/production/flow-card/list', { params })
}

export const getFlowCard = (id: number) => {
  return request.get(`/production/flow-card/${id}`)
}

export const createFlowCard = (data: any) => {
  return request.post('/production/flow-card', data)
}

export const updateFlowCard = (id: number, data: any) => {
  return request.put(`/production/flow-card/${id}`, data)
}

export const deleteFlowCard = (id: number) => {
  return request.delete(`/production/flow-card/${id}`)
}
