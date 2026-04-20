import request from '@/utils/request'

// 采购订单
export const getPurchaseOrderList = (params?: any) => {
  return request.get('/scp/purchase-orders/list', { params })
}

export const getPurchaseOrderById = (id: number) => {
  return request.get(`/scp/purchase-orders/${id}`)
}

export const createPurchaseOrder = (data: any) => {
  return request.post('/scp/purchase-orders', data)
}

export const updatePurchaseOrder = (id: number, data: any) => {
  return request.put(`/scp/purchase-orders/${id}`, data)
}

export const deletePurchaseOrder = (id: number) => {
  return request.delete(`/scp/purchase-orders/${id}`)
}

export const submitPurchaseOrder = (id: number) => {
  return request.post(`/scp/purchase-orders/${id}/submit`)
}

export const approvePurchaseOrder = (id: number) => {
  return request.post(`/scp/purchase-orders/${id}/approve`)
}

export const rejectPurchaseOrder = (id: number) => {
  return request.post(`/scp/purchase-orders/${id}/reject`)
}

export const issuePurchaseOrder = (id: number) => {
  return request.post(`/scp/purchase-orders/${id}/issue`)
}

export const closePurchaseOrder = (id: number, closeReason: string) => {
  return request.post(`/scp/purchase-orders/${id}/close`, { close_reason: closeReason })
}

export const cancelPurchaseOrder = (id: number) => {
  return request.post(`/scp/purchase-orders/${id}/cancel`)
}

// 按行项目收货
export const receivePurchaseOrder = (itemId: number, receivedQty: number, batchNo?: string) => {
  return request.post('/scp/purchase-orders/receive', { item_id: itemId, received_qty: receivedQty, batch_no: batchNo })
}

// 供应商列表
export const getSupplierList = (params?: any) => {
  return request.get('/scp/supplier/list', { params })
}

export const getSupplierById = (id: number) => {
  return request.get(`/scp/supplier/${id}`)
}

export const createSupplier = (data: any) => {
  return request.post('/scp/supplier', data)
}

export const updateSupplier = (id: number, data: any) => {
  return request.put(`/scp/supplier/${id}`, data)
}

export const deleteSupplier = (id: number) => {
  return request.delete(`/scp/supplier/${id}`)
}

// 物料列表
export const getMaterialList = (params?: any) => {
  return request.get('/mdm/material/list', { params })
}

// ============ RFQ询价单 ============
export const getRFQList = (params?: any) => {
  return request.get('/scp/rfq/list', { params })
}

export const getRFQById = (id: number) => {
  return request.get(`/scp/rfq/${id}`)
}

export const createRFQ = (data: any) => {
  return request.post('/scp/rfq', data)
}

export const updateRFQ = (id: number, data: any) => {
  return request.put(`/scp/rfq/${id}`, data)
}

export const deleteRFQ = (id: number) => {
  return request.delete(`/scp/rfq/${id}`)
}

export const publishRFQ = (id: number) => {
  return request.put(`/scp/rfq/${id}/publish`)
}

export const closeRFQ = (id: number) => {
  return request.put(`/scp/rfq/${id}/close`)
}

export const getRFQQuotes = (id: number) => {
  return request.get(`/scp/rfq/${id}/quotes`)
}

export const awardRFQ = (id: number, data: any) => {
  return request.put(`/scp/rfq/${id}/award`, data)
}

// ============ 供应商报价 ============
export const getSupplierQuoteList = (params?: any) => {
  return request.get('/scp/supplier-quotes/list', { params })
}

// ============ 供应商绩效 ============
export const getSupplierKPIList = (params?: any) => {
  return request.get('/scp/supplier-kpi/list', { params })
}

export const getSupplierKPIMonthly = (supplierId: number) => {
  return request.get(`/scp/supplier-kpi/${supplierId}/monthly`)
}

export const createSupplierKPI = (data: any) => {
  return request.post('/scp/supplier-kpi', data)
}

export const getSupplierKPIRanking = () => {
  return request.get('/scp/supplier-kpi/ranking')
}

// ============ 销售订单 ============
export const getSCPSalesOrderList = (params?: any) => {
  return request.get('/scp/sales-orders/list', { params })
}

export const getSCPSalesOrder = (id: number) => {
  return request.get(`/scp/sales-orders/${id}`)
}

export const createSCPSalesOrder = (data: any) => {
  return request.post('/scp/sales-orders', data)
}

export const updateSCPSalesOrder = (id: number, data: any) => {
  return request.put(`/scp/sales-orders/${id}`, data)
}

export const deleteSCPSalesOrder = (id: number) => {
  return request.delete(`/scp/sales-orders/${id}`)
}

export const submitSCPSalesOrder = (id: number) => {
  return request.post(`/scp/sales-orders/${id}/submit`)
}

export const approveSCPSalesOrder = (id: number) => {
  return request.post(`/scp/sales-orders/${id}/approve`)
}

export const rejectSCPSalesOrder = (id: number) => {
  return request.post(`/scp/sales-orders/${id}/reject`)
}

export const confirmSCPSalesOrder = (id: number) => {
  return request.post(`/scp/sales-orders/${id}/confirm`)
}

export const closeSCPSalesOrder = (id: number) => {
  return request.post(`/scp/sales-orders/${id}/close`)
}

export const cancelSCPSalesOrder = (id: number) => {
  return request.post(`/scp/sales-orders/${id}/cancel`)
}

// ============ 客户询价 ============
export const getCustomerInquiryList = (params?: any) => {
  return request.get('/scp/customer-inquiry/list', { params })
}

export const getCustomerInquiryById = (id: number) => {
  return request.get(`/scp/customer-inquiry/${id}`)
}

export const createCustomerInquiry = (data: any) => {
  return request.post('/scp/customer-inquiry', data)
}

export const updateCustomerInquiry = (id: number, data: any) => {
  return request.put(`/scp/customer-inquiry/${id}`, data)
}

export const deleteCustomerInquiry = (id: number) => {
  return request.delete(`/scp/customer-inquiry/${id}`)
}

export const sendCustomerInquiry = (id: number) => {
  return request.post(`/scp/customer-inquiry/${id}/send`)
}

export const quoteCustomerInquiry = (id: number, quotedAmount: number) => {
  return request.post(`/scp/customer-inquiry/${id}/quote`, { quoted_amount: quotedAmount })
}

export const winCustomerInquiry = (id: number, supplierId: number) => {
  return request.post(`/scp/customer-inquiry/${id}/win`, { supplier_id: supplierId })
}

export const loseCustomerInquiry = (id: number) => {
  return request.post(`/scp/customer-inquiry/${id}/lose`)
}

export const cancelCustomerInquiry = (id: number) => {
  return request.post(`/scp/customer-inquiry/${id}/cancel`)
}
