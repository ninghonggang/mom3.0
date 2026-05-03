package scp

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type CustomerCreditHandler struct {
	svc *service.CustomerCreditService
}

func NewCustomerCreditHandler(svc *service.CustomerCreditService) *CustomerCreditHandler {
	return &CustomerCreditHandler{svc: svc}
}

// List 获取客户信用列表(分页)
// GET /scp/customer-credit/page
func (h *CustomerCreditHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	customerCode := c.Query("customer_code")
	customerName := c.Query("customer_name")
	creditLevel := c.Query("credit_level")
	riskLevel, _ := strconv.Atoi(c.Query("risk_level"))
	blacklist, _ := strconv.Atoi(c.Query("blacklist"))
	status, _ := strconv.Atoi(c.Query("status"))

	req := &model.CustomerCreditQuery{
		CustomerCode: customerCode,
		CustomerName: customerName,
		CreditLevel:  creditLevel,
		RiskLevel:    riskLevel,
		Blacklist:    blacklist,
		Status:       status,
		Page:         page,
		PageSize:     pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Get 获取单个客户信用
// GET /scp/customer-credit/:id
func (h *CustomerCreditHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	credit, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, credit)
}

// GetByCustomer 获取客户信用ByCustomerID
// GET /scp/customer-credit/customer/:customerId
func (h *CustomerCreditHandler) GetByCustomer(c *gin.Context) {
	customerIDStr := c.Param("customerId")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid customer_id")
		return
	}
	credit, err := h.svc.GetByCustomerID(c.Request.Context(), customerID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, credit)
}

// Create 创建客户信用
// POST /scp/customer-credit
func (h *CustomerCreditHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.CustomerCreditCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	credit, err := h.svc.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, credit)
}

// Update 更新客户信用
// PUT /scp/customer-credit/:id
func (h *CustomerCreditHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.CustomerCreditCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// UpdateUsedCredit 更新已用额度
// PUT /scp/customer-credit/:id/used-credit
func (h *CustomerCreditHandler) UpdateUsedCredit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req struct {
		UsedCredit float64 `json:"used_credit"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.UpdateUsedCredit(c.Request.Context(), id, req.UsedCredit); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// SetBlacklist 设置黑名单
// PUT /scp/customer-credit/:id/blacklist
func (h *CustomerCreditHandler) SetBlacklist(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req struct {
		Blacklist bool `json:"blacklist"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.SetBlacklist(c.Request.Context(), id, req.Blacklist); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Freeze 冻结
// PUT /scp/customer-credit/:id/freeze
func (h *CustomerCreditHandler) Freeze(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Freeze(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Unfreeze 解冻
// PUT /scp/customer-credit/:id/unfreeze
func (h *CustomerCreditHandler) Unfreeze(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Unfreeze(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除客户信用
// DELETE /scp/customer-credit/:id
func (h *CustomerCreditHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}