package scp

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type SupplierExtHandler struct {
	supplierService *service.ScpSupplierService
}

func NewSupplierExtHandler(s *service.ScpSupplierService) *SupplierExtHandler {
	return &SupplierExtHandler{supplierService: s}
}

// ==================== 供应商联系人 ====================

// ListContacts 查询供应商联系人列表
func (h *SupplierExtHandler) ListContacts(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if supplierID := c.Query("supplierId"); supplierID != "" {
		query["supplier_id"] = supplierID
	}
	if page := c.DefaultQuery("page", "1"); page != "" {
		query["page"] = 1
	}

	list, total, err := h.supplierService.ListSupplierContacts(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetContact 获取供应商联系人详情
func (h *SupplierExtHandler) GetContact(c *gin.Context) {
	id := c.Param("id")
	contact, err := h.supplierService.GetSupplierContact(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, contact)
}

// ListContactsBySupplier 查询指定供应商的联系人
func (h *SupplierExtHandler) ListContactsBySupplier(c *gin.Context) {
	supplierID, err := strconv.ParseInt(c.Param("supplierId"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid supplier id")
		return
	}

	list, err := h.supplierService.ListContactsBySupplier(c.Request.Context(), supplierID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": list,
	})
}

// CreateContact 创建供应商联系人
func (h *SupplierExtHandler) CreateContact(c *gin.Context) {
	var req model.ScpSupplierContactCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	contact, err := h.supplierService.CreateSupplierContact(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, contact)
}

// UpdateContact 更新供应商联系人
func (h *SupplierExtHandler) UpdateContact(c *gin.Context) {
	id := c.Param("id")
	var req model.ScpSupplierContactUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.supplierService.UpdateSupplierContact(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteContact 删除供应商联系人
func (h *SupplierExtHandler) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	if err := h.supplierService.DeleteSupplierContact(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// ==================== 供应商银行账户 ====================

// ListBanks 查询供应商银行账户列表
func (h *SupplierExtHandler) ListBanks(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if supplierID := c.Query("supplierId"); supplierID != "" {
		query["supplier_id"] = supplierID
	}
	if page := c.DefaultQuery("page", "1"); page != "" {
		query["page"] = 1
	}

	list, total, err := h.supplierService.ListSupplierBanks(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetBank 获取供应商银行账户详情
func (h *SupplierExtHandler) GetBank(c *gin.Context) {
	id := c.Param("id")
	bank, err := h.supplierService.GetSupplierBank(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, bank)
}

// ListBanksBySupplier 查询指定供应商的银行账户
func (h *SupplierExtHandler) ListBanksBySupplier(c *gin.Context) {
	supplierID, err := strconv.ParseInt(c.Param("supplierId"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid supplier id")
		return
	}

	list, err := h.supplierService.ListBanksBySupplier(c.Request.Context(), supplierID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": list,
	})
}

// CreateBank 创建供应商银行账户
func (h *SupplierExtHandler) CreateBank(c *gin.Context) {
	var req model.ScpSupplierBankCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	bank, err := h.supplierService.CreateSupplierBank(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, bank)
}

// UpdateBank 更新供应商银行账户
func (h *SupplierExtHandler) UpdateBank(c *gin.Context) {
	id := c.Param("id")
	var req model.ScpSupplierBankUpdateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.supplierService.UpdateSupplierBank(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteBank 删除供应商银行账户
func (h *SupplierExtHandler) DeleteBank(c *gin.Context) {
	id := c.Param("id")
	if err := h.supplierService.DeleteSupplierBank(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
