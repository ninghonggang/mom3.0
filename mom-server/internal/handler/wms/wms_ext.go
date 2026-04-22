package wms

import (
	"log"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== 调拨管理 ==========

type TransferOrderHandler struct {
	service *service.TransferOrderService
}

func NewTransferOrderHandler(s *service.TransferOrderService) *TransferOrderHandler {
	log.Printf("DEBUG NewTransferOrderHandler called with service pointer: %p", s)
	return &TransferOrderHandler{service: s}
}

func (h *TransferOrderHandler) List(c *gin.Context) {
	list, total, err := h.service.List(c.Request.Context(), c.Query("query"))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *TransferOrderHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	order, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	items, _ := h.service.GetItems(c.Request.Context(), int64(order.ID))
	response.Success(c, gin.H{"order": order, "items": items})
}

func (h *TransferOrderHandler) Create(c *gin.Context) {
	var req model.TransferOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *TransferOrderHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *TransferOrderHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *TransferOrderHandler) AddItem(c *gin.Context) {
	var req model.TransferOrderItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.AddItem(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

// Submit 提交调拨单
func (h *TransferOrderHandler) Submit(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Submit(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Approve 审批调拨单
func (h *TransferOrderHandler) Approve(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		Approved bool    `json:"approved"`
		Comment  *string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Approve(c.Request.Context(), uint(id), req.Approved, req.Comment); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Start 开始调拨
func (h *TransferOrderHandler) Start(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.StartTransfer(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Ship 发货确认
func (h *TransferOrderHandler) Ship(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		OperatorID   int64  `json:"operator_id"`
		OperatorName string `json:"operator_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Ship(c.Request.Context(), uint(id), req.OperatorID, req.OperatorName); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Receive 收货确认
func (h *TransferOrderHandler) Receive(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		OperatorID   int64  `json:"operator_id"`
		OperatorName string `json:"operator_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Receive(c.Request.Context(), uint(id), req.OperatorID, req.OperatorName); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Complete 完成调拨
func (h *TransferOrderHandler) Complete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Complete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Cancel 取消调拨
func (h *TransferOrderHandler) Cancel(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	reason := c.Query("reason")
	if err := h.service.Cancel(c.Request.Context(), uint(id), reason); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetTrace 获取调拨跟踪记录
func (h *TransferOrderHandler) GetTrace(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	traces, err := h.service.GetTraces(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"traces": traces})
}

// ========== 盘点管理 ==========

type StockCheckHandler struct {
	service *service.StockCheckService
}

func NewStockCheckHandler(s *service.StockCheckService) *StockCheckHandler {
	return &StockCheckHandler{service: s}
}

func (h *StockCheckHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *StockCheckHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	check, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	items, _ := h.service.GetItems(c.Request.Context(), int64(check.ID))
	response.Success(c, gin.H{"check": check, "items": items})
}

func (h *StockCheckHandler) Create(c *gin.Context) {
	var req model.StockCheck
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *StockCheckHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *StockCheckHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *StockCheckHandler) AddItem(c *gin.Context) {
	var req model.StockCheckItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.AddItem(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *StockCheckHandler) UpdateItem(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.UpdateItem(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Submit 提交盘点单
func (h *StockCheckHandler) Submit(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Submit(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Start 开始盘点
func (h *StockCheckHandler) Start(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.StartCheck(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Complete 完成盘点
func (h *StockCheckHandler) Complete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.CompleteCheck(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Approve 审核盘点差异
func (h *StockCheckHandler) Approve(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		Approved bool    `json:"approved"`
		Comment  *string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.ApproveCheck(c.Request.Context(), uint(id), req.Approved, req.Comment); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// CountItem 录入盘点数据
func (h *StockCheckHandler) CountItem(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	itemID, _ := strconv.ParseUint(c.Param("itemId"), 10, 64)
	var req struct {
		CountedQty float64 `json:"counted_qty"`
		CounterID  int64   `json:"counter_id"`
		CounterName string  `json:"counter_name"`
		Remark     *string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	now := time.Now()
	if err := h.service.CountItem(c.Request.Context(), uint(id), uint(itemID), req.CountedQty, req.CounterID, req.CounterName, &now); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// HandleVariance 处理差异
func (h *StockCheckHandler) HandleVariance(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	itemID, _ := strconv.ParseUint(c.Param("itemId"), 10, 64)
	var req struct {
		HandleMethod string  `json:"handle_method"` // ADJUST/WRITE_OFF/WRITE_IN
		HandleQty    float64 `json:"handle_qty"`
		HandlerID    int64   `json:"handler_id"`
		HandlerName  string  `json:"handler_name"`
		Remark       *string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.HandleVariance(c.Request.Context(), uint(id), uint(itemID), req.HandleMethod, req.HandleQty, req.HandlerID, req.HandlerName); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Recount 复盘
func (h *StockCheckHandler) Recount(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	itemID, _ := strconv.ParseUint(c.Param("itemId"), 10, 64)
	var req struct {
		RecountQty  float64 `json:"recount_qty"`
		RecountBy   int64   `json:"recount_by"`
		RecountName string  `json:"recount_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	now := time.Now()
	if err := h.service.Recount(c.Request.Context(), uint(id), uint(itemID), req.RecountQty, req.RecountBy, req.RecountName, &now); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetVariance 获取盘点差异列表
func (h *StockCheckHandler) GetVariance(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	variances, err := h.service.GetVariances(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"variances": variances})
}

// ========== 线边库位 ==========

type SideLocationHandler struct {
	service *service.SideLocationService
}

func NewSideLocationHandler(s *service.SideLocationService) *SideLocationHandler {
	return &SideLocationHandler{service: s}
}

func (h *SideLocationHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *SideLocationHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	loc, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, loc)
}

func (h *SideLocationHandler) Create(c *gin.Context) {
	var req model.SideLocation
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SideLocationHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SideLocationHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ========== 看板拉动 ==========

type KanbanPullHandler struct {
	service *service.KanbanPullService
}

func NewKanbanPullHandler(s *service.KanbanPullService) *KanbanPullHandler {
	return &KanbanPullHandler{service: s}
}

func (h *KanbanPullHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *KanbanPullHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	k, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, k)
}

func (h *KanbanPullHandler) Create(c *gin.Context) {
	var req model.KanbanPull
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *KanbanPullHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *KanbanPullHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Trigger 触发看板拉动
func (h *KanbanPullHandler) Trigger(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Trigger(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
