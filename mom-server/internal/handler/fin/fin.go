package fin

import (
	"fmt"
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

func toInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

type FinHandler struct {
	finSvc *service.FinService
}

func NewFinHandler(finSvc *service.FinService) *FinHandler {
	return &FinHandler{finSvc: finSvc}
}

// ==================== 采购结算 ====================

func (h *FinHandler) ListPurchaseSettlements(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query["supplier_id"] = toInt64(supplierID)
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if startDate := c.Query("start_date"); startDate != "" {
		query["start_date"] = startDate
	}
	if endDate := c.Query("end_date"); endDate != "" {
		query["end_date"] = endDate
	}
	if settlementNo := c.Query("settlement_no"); settlementNo != "" {
		query["settlement_no"] = settlementNo
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.finSvc.ListPurchaseSettlements(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *FinHandler) GetPurchaseSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	settlement, err := h.finSvc.GetPurchaseSettlement(c.Request.Context(), settlementID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, settlement)
}

func (h *FinHandler) CreatePurchaseSettlement(c *gin.Context) {
	var req model.PurchaseSettlementCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	username := middleware.GetUsername(c)

	settlement, err := h.finSvc.CreatePurchaseSettlement(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, settlement)
}

func (h *FinHandler) SubmitPurchaseSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	if err := h.finSvc.SubmitPurchaseSettlement(c.Request.Context(), settlementID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) ApprovePurchaseSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	userID := middleware.GetUserID(c)

	if err := h.finSvc.ApprovePurchaseSettlement(c.Request.Context(), settlementID, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) CancelPurchaseSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	if err := h.finSvc.CancelPurchaseSettlement(c.Request.Context(), settlementID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) DeletePurchaseSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	if err := h.finSvc.DeletePurchaseSettlement(c.Request.Context(), settlementID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 销售结算 ====================

func (h *FinHandler) ListSalesSettlements(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if customerID := c.Query("customer_id"); customerID != "" {
		query["customer_id"] = toInt64(customerID)
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if startDate := c.Query("start_date"); startDate != "" {
		query["start_date"] = startDate
	}
	if endDate := c.Query("end_date"); endDate != "" {
		query["end_date"] = endDate
	}
	if settlementNo := c.Query("settlement_no"); settlementNo != "" {
		query["settlement_no"] = settlementNo
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.finSvc.ListSalesSettlements(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *FinHandler) GetSalesSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	settlement, err := h.finSvc.GetSalesSettlement(c.Request.Context(), settlementID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, settlement)
}

func (h *FinHandler) CreateSalesSettlement(c *gin.Context) {
	var req model.SalesSettlementCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	username := middleware.GetUsername(c)

	settlement, err := h.finSvc.CreateSalesSettlement(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, settlement)
}

func (h *FinHandler) SubmitSalesSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	if err := h.finSvc.SubmitSalesSettlement(c.Request.Context(), settlementID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) ApproveSalesSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	userID := middleware.GetUserID(c)

	if err := h.finSvc.ApproveSalesSettlement(c.Request.Context(), settlementID, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) CancelSalesSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	if err := h.finSvc.CancelSalesSettlement(c.Request.Context(), settlementID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) DeleteSalesSettlement(c *gin.Context) {
	id := c.Param("id")
	var settlementID uint
	_, err := fmt.Sscanf(id, "%d", &settlementID)
	if err != nil {
		response.ErrorMsg(c, "invalid settlement id")
		return
	}

	if err := h.finSvc.DeleteSalesSettlement(c.Request.Context(), settlementID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 付款申请 ====================

func (h *FinHandler) ListPaymentRequests(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if requestType := c.Query("request_type"); requestType != "" {
		query["request_type"] = requestType
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if startDate := c.Query("start_date"); startDate != "" {
		query["start_date"] = startDate
	}
	if endDate := c.Query("end_date"); endDate != "" {
		query["end_date"] = endDate
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.finSvc.ListPaymentRequests(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *FinHandler) GetPaymentRequest(c *gin.Context) {
	id := c.Param("id")
	var requestID uint
	_, err := fmt.Sscanf(id, "%d", &requestID)
	if err != nil {
		response.ErrorMsg(c, "invalid request id")
		return
	}

	req, err := h.finSvc.GetPaymentRequest(c.Request.Context(), requestID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *FinHandler) CreatePaymentRequest(c *gin.Context) {
	var req model.PaymentRequestCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	username := middleware.GetUsername(c)

	paymentReq, err := h.finSvc.CreatePaymentRequest(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, paymentReq)
}

func (h *FinHandler) SubmitPaymentRequest(c *gin.Context) {
	id := c.Param("id")
	var requestID uint
	_, err := fmt.Sscanf(id, "%d", &requestID)
	if err != nil {
		response.ErrorMsg(c, "invalid request id")
		return
	}

	if err := h.finSvc.SubmitPaymentRequest(c.Request.Context(), requestID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) ApprovePaymentRequest(c *gin.Context) {
	id := c.Param("id")
	var requestID uint
	_, err := fmt.Sscanf(id, "%d", &requestID)
	if err != nil {
		response.ErrorMsg(c, "invalid request id")
		return
	}

	var reqBody struct {
		Comment string `json:"comment"`
	}
	c.ShouldBindJSON(&reqBody)

	userID := middleware.GetUserID(c)

	if err := h.finSvc.ApprovePaymentRequest(c.Request.Context(), requestID, userID, reqBody.Comment); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) RejectPaymentRequest(c *gin.Context) {
	id := c.Param("id")
	var requestID uint
	_, err := fmt.Sscanf(id, "%d", &requestID)
	if err != nil {
		response.ErrorMsg(c, "invalid request id")
		return
	}

	var reqBody struct {
		Comment string `json:"comment"`
	}
	c.ShouldBindJSON(&reqBody)

	userID := middleware.GetUserID(c)

	if err := h.finSvc.RejectPaymentRequest(c.Request.Context(), requestID, userID, reqBody.Comment); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) PayPaymentRequest(c *gin.Context) {
	id := c.Param("id")
	var requestID uint
	_, err := fmt.Sscanf(id, "%d", &requestID)
	if err != nil {
		response.ErrorMsg(c, "invalid request id")
		return
	}

	userID := middleware.GetUserID(c)

	if err := h.finSvc.PayPaymentRequest(c.Request.Context(), requestID, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *FinHandler) DeletePaymentRequest(c *gin.Context) {
	id := c.Param("id")
	var requestID uint
	_, err := fmt.Sscanf(id, "%d", &requestID)
	if err != nil {
		response.ErrorMsg(c, "invalid request id")
		return
	}

	if err := h.finSvc.DeletePaymentRequest(c.Request.Context(), requestID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 采购预付款 ====================

func (h *FinHandler) ListPurchaseAdvances(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query["supplier_id"] = toInt64(supplierID)
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.finSvc.ListPurchaseAdvances(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *FinHandler) CreatePurchaseAdvance(c *gin.Context) {
	var advance model.PurchaseAdvance
	if err := c.ShouldBindJSON(&advance); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	username := middleware.GetUsername(c)
	advance.TenantID = tenantID
	advance.CreatedBy = &username

	if err := h.finSvc.CreatePurchaseAdvance(c.Request.Context(), tenantID, &advance); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, advance)
}

// ==================== 销售收款 ====================

func (h *FinHandler) ListSalesReceipts(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if customerID := c.Query("customer_id"); customerID != "" {
		query["customer_id"] = toInt64(customerID)
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.finSvc.ListSalesReceipts(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *FinHandler) CreateSalesReceipt(c *gin.Context) {
	var receipt model.SalesReceipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	username := middleware.GetUsername(c)
	receipt.TenantID = tenantID
	receipt.CreatedBy = &username

	if err := h.finSvc.CreateSalesReceipt(c.Request.Context(), tenantID, &receipt); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, receipt)
}

// ==================== 供应商对账 ====================

func (h *FinHandler) ListSupplierStatements(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query["supplier_id"] = toInt64(supplierID)
	}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if page := c.Query("page"); page != "" {
		query["page"] = toInt(page)
	}
	if limit := c.Query("limit"); limit != "" {
		query["limit"] = toInt(limit)
	}

	list, total, err := h.finSvc.ListSupplierStatements(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *FinHandler) GetSupplierStatement(c *gin.Context) {
	id := c.Param("id")
	var statementID uint
	_, err := fmt.Sscanf(id, "%d", &statementID)
	if err != nil {
		response.ErrorMsg(c, "invalid statement id")
		return
	}

	statement, err := h.finSvc.GetSupplierStatement(c.Request.Context(), statementID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, statement)
}
