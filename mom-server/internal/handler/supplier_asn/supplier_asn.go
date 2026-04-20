package supplier_asn

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// SupplierASNHandler ASN接口处理
type SupplierASNHandler struct {
	svc *service.SupplierASNService
}

// NewSupplierASNHandler 创建ASN处理器
func NewSupplierASNHandler(svc *service.SupplierASNService) *SupplierASNHandler {
	return &SupplierASNHandler{svc: svc}
}

// List GET /supplier/asn/list - 查询ASN列表
func (h *SupplierASNHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	q := &model.SupplierASNQuery{
		TenantID:     tenantID,
		SupplierCode: c.Query("supplierCode"),
		Status:       c.Query("status"),
		StartDate:    c.Query("startDate"),
		EndDate:      c.Query("endDate"),
		Page:         page,
		PageSize:     pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), q)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Get GET /supplier/asn/:id - 获取ASN详情
func (h *SupplierASNHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	asn, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, asn)
}

// GetByNo GET /supplier/asn/no/:asnNo - 根据ASN编号获取
func (h *SupplierASNHandler) GetByNo(c *gin.Context) {
	asnNo := c.Param("asnNo")
	if asnNo == "" {
		response.BadRequest(c, "asnNo is required")
		return
	}

	asn, err := h.svc.GetByASNNo(c.Request.Context(), asnNo)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, asn)
}

// Create POST /supplier/asn - 创建ASN
func (h *SupplierASNHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	var req model.CreateSupplierASNRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.TenantID = tenantID

	asn, err := h.svc.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, asn)
}

// Update PUT /supplier/asn/:id - 更新ASN
func (h *SupplierASNHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.UpdateSupplierASNRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "更新成功"})
}

// Delete DELETE /supplier/asn/:id - 删除ASN
func (h *SupplierASNHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// Submit PUT /supplier/asn/:id/submit - 提交ASN
func (h *SupplierASNHandler) Submit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.Submit(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "提交成功"})
}

// Confirm PUT /supplier/asn/:id/confirm - 确认ASN
func (h *SupplierASNHandler) Confirm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.ConfirmASNRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Confirm(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "确认成功"})
}

// StartReceiving PUT /supplier/asn/:id/start-receiving - 开始收货
func (h *SupplierASNHandler) StartReceiving(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.StartReceiving(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "开始收货"})
}

// CompleteReceiving PUT /supplier/asn/:id/complete-receiving - 完成收货
func (h *SupplierASNHandler) CompleteReceiving(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.CompleteReceiving(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "收货完成"})
}

// Cancel PUT /supplier/asn/:id/cancel - 取消ASN
func (h *SupplierASNHandler) Cancel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.Cancel(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "取消成功"})
}

// AddItem POST /supplier/asn/:id/items - 添加ASN明细
func (h *SupplierASNHandler) AddItem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req model.CreateASNItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	item, err := h.svc.AddItem(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}
