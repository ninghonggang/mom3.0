package scp

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierKPIHandler struct {
	scpService *service.SCPService
}

func NewSupplierKPIHandler(s *service.SCPService) *SupplierKPIHandler {
	return &SupplierKPIHandler{scpService: s}
}

func (h *SupplierKPIHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query["supplier_id"] = supplierID
	}
	if month := c.Query("evaluation_month"); month != "" {
		query["evaluation_month"] = month
	}

	list, total, err := h.scpService.ListSupplierKPI(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *SupplierKPIHandler) GetByMonthly(c *gin.Context) {
	supplierID := c.Param("supplierId")
	month := c.Query("month")

	supplierIDInt, _ := strconv.ParseInt(supplierID, 10, 64)
	kpi, err := h.scpService.GetSupplierKPIByMonthly(c.Request.Context(), supplierIDInt, month)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, kpi)
}

func (h *SupplierKPIHandler) Create(c *gin.Context) {
	var req model.SupplierKPI
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.scpService.CreateSupplierKPI(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SupplierKPIHandler) GetRanking(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	month := c.DefaultQuery("month", "")

	ranking, err := h.scpService.GetSupplierRanking(c.Request.Context(), tenantID, month)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": ranking,
	})
}

func (h *SupplierKPIHandler) GetPurchaseInfo(c *gin.Context) {
	supplierID := c.Param("supplierId")

	info, err := h.scpService.GetSupplierPurchaseInfo(c.Request.Context(), supplierID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, info)
}

func (h *SupplierKPIHandler) CreatePurchaseInfo(c *gin.Context) {
	var req model.SupplierPurchaseInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.scpService.CreateSupplierPurchaseInfo(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *SupplierKPIHandler) UpdatePurchaseInfo(c *gin.Context) {
	supplierID := c.Param("supplierId")
	var req model.SupplierPurchaseInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.scpService.UpdateSupplierPurchaseInfo(c.Request.Context(), supplierID, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}
