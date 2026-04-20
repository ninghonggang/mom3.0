package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type PurchasePlanHandler struct {
	planService *service.ScpPurchasePlanService
}

func NewPurchasePlanHandler(s *service.ScpPurchasePlanService) *PurchasePlanHandler {
	return &PurchasePlanHandler{planService: s}
}

// List 查询采购计划列表
func (h *PurchasePlanHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if planType := c.Query("planType"); planType != "" {
		query["plan_type"] = planType
	}
	if planYear := c.Query("planYear"); planYear != "" {
		query["plan_year"] = planYear
	}
	if planMonth := c.Query("planMonth"); planMonth != "" {
		query["plan_month"] = planMonth
	}
	if page := c.DefaultQuery("page", "1"); page != "" {
		query["page"] = 1
	}

	list, total, err := h.planService.ListPurchasePlans(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// Get 获取采购计划详情
func (h *PurchasePlanHandler) Get(c *gin.Context) {
	id := c.Param("id")
	plan, err := h.planService.GetPurchasePlan(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, plan)
}

// Create 创建采购计划
func (h *PurchasePlanHandler) Create(c *gin.Context) {
	var req model.ScpPurchasePlanCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	plan, err := h.planService.CreatePurchasePlan(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, plan)
}

// Update 更新采购计划
func (h *PurchasePlanHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.ScpPurchasePlanUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	plan, err := h.planService.UpdatePurchasePlan(c.Request.Context(), tenantID, id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, plan)
}

// Delete 删除采购计划
func (h *PurchasePlanHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.planService.DeletePurchasePlan(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Confirm 确认采购计划
func (h *PurchasePlanHandler) Confirm(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)
	if err := h.planService.ConfirmPurchasePlan(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Publish 发布采购计划
func (h *PurchasePlanHandler) Publish(c *gin.Context) {
	id := c.Param("id")
	userID := middleware.GetUserID(c)
	if err := h.planService.PublishPurchasePlan(c.Request.Context(), id, userID); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Close 关闭采购计划
func (h *PurchasePlanHandler) Close(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		CloseReason string `json:"closeReason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "请提供关闭原因")
		return
	}
	userID := middleware.GetUserID(c)
	if err := h.planService.ClosePurchasePlan(c.Request.Context(), id, userID, req.CloseReason); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetItems 获取采购计划明细
func (h *PurchasePlanHandler) GetItems(c *gin.Context) {
	id := c.Param("id")
	items, err := h.planService.GetPurchasePlanItems(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": items,
	})
}
