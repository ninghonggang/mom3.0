package production

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type CostHandler struct {
	svc *service.ProductionCostService
}

func NewCostHandler(svc *service.ProductionCostService) *CostHandler {
	return &CostHandler{svc: svc}
}

// Create 创建成本记录
// POST /production/cost
func (h *CostHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.ProductionCostCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	cost, err := h.svc.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, cost)
}

// List 获取成本列表
// GET /production/cost/list
func (h *CostHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	orderID, _ := strconv.ParseInt(c.Query("order_id"), 10, 64)
	orderNo := c.Query("order_no")
	costType := c.Query("cost_type")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	req := &model.ProductionCostQuery{
		OrderID:   orderID,
		OrderNo:   orderNo,
		CostType:  costType,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
		PageSize:  pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetSummary 获取工单成本汇总
// GET /production/cost/summary?order_id=xxx
func (h *CostHandler) GetSummary(c *gin.Context) {
	orderIDStr := c.Query("order_id")
	if orderIDStr == "" {
		response.BadRequest(c, "order_id is required")
		return
	}
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid order_id")
		return
	}
	summary, err := h.svc.GetSummary(c.Request.Context(), orderID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, summary)
}

// Delete 删除成本记录
// DELETE /production/cost/:id
func (h *CostHandler) Delete(c *gin.Context) {
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