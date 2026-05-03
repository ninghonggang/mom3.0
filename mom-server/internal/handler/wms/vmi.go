package wms

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type VmiHandler struct {
	svc *service.VmiService
}

func NewVmiHandler(svc *service.VmiService) *VmiHandler {
	return &VmiHandler{svc: svc}
}

// Vendor List
// GET /wms/vmi/vendor/page
func (h *VmiHandler) ListVendors(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	vendorCode := c.Query("vendor_code")
	vendorName := c.Query("vendor_name")
	warehouseID, _ := strconv.ParseInt(c.Query("warehouse_id"), 10, 64)
	isActive, _ := strconv.Atoi(c.Query("is_active"))
	status, _ := strconv.Atoi(c.Query("status"))

	req := &model.VmiVendorQuery{
		VendorCode:  vendorCode,
		VendorName:  vendorName,
		WarehouseID: warehouseID,
		IsActive:    isActive,
		Status:      status,
		Page:        page,
		PageSize:    pageSize,
	}

	list, total, err := h.svc.ListVendors(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetVendor
// GET /wms/vmi/vendor/:id
func (h *VmiHandler) GetVendor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	vendor, err := h.svc.GetVendorByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, vendor)
}

// CreateVendor
// POST /wms/vmi/vendor
func (h *VmiHandler) CreateVendor(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.VmiVendorCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	vendor, err := h.svc.CreateVendor(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, vendor)
}

// UpdateVendor
// PUT /wms/vmi/vendor/:id
func (h *VmiHandler) UpdateVendor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.VmiVendorCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.UpdateVendor(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteVendor
// DELETE /wms/vmi/vendor/:id
func (h *VmiHandler) DeleteVendor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.DeleteVendor(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Material List
// GET /wms/vmi/material/page
func (h *VmiHandler) ListMaterials(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	vendorID, _ := strconv.ParseInt(c.Query("vendor_id"), 10, 64)
	materialCode := c.Query("material_code")
	materialName := c.Query("material_name")

	req := &model.VmiMaterialQuery{
		VendorID:     vendorID,
		MaterialCode: materialCode,
		MaterialName: materialName,
		Page:         page,
		PageSize:     pageSize,
	}

	list, total, err := h.svc.ListMaterials(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// GetMaterial
// GET /wms/vmi/material/:id
func (h *VmiHandler) GetMaterial(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	material, err := h.svc.GetMaterialByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, material)
}

// Transaction List
// GET /wms/vmi/transaction/page
func (h *VmiHandler) ListTransactions(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	vendorID, _ := strconv.ParseInt(c.Query("vendor_id"), 10, 64)
	materialID, _ := strconv.ParseInt(c.Query("material_id"), 10, 64)
	txType, _ := strconv.Atoi(c.Query("transaction_type"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	req := &model.VmiTransactionQuery{
		VendorID:        vendorID,
		MaterialID:      materialID,
		TransactionType: txType,
		StartDate:      startDate,
		EndDate:        endDate,
		Page:           page,
		PageSize:       pageSize,
	}

	list, total, err := h.svc.ListTransactions(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Consume
// POST /wms/vmi/consume
func (h *VmiHandler) Consume(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	userID := middleware.GetUserID(c)
	userName := middleware.GetUsername(c)

	var req model.VmiConsumeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Consume(c.Request.Context(), tenantID, userID, userName, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Replenish
// POST /wms/vmi/replenish
func (h *VmiHandler) Replenish(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	userID := middleware.GetUserID(c)
	userName := middleware.GetUsername(c)

	var req model.VmiReplenishReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Replenish(c.Request.Context(), tenantID, userID, userName, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}