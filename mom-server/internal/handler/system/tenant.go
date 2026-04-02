package system

import (
	"strconv"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type TenantHandler struct {
	tenantSvc *service.TenantService
}

func NewTenantHandler(tenantSvc *service.TenantService) *TenantHandler {
	return &TenantHandler{tenantSvc: tenantSvc}
}

func (h *TenantHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	tenantName := c.Query("tenant_name")
	tenantKey := c.Query("tenant_key")
	status := c.Query("status")

	list, total, err := h.tenantSvc.GetList(c.Request.Context(), tenantName, tenantKey, status, page, pageSize)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

func (h *TenantHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	tenant, err := h.tenantSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, tenant)
}

func (h *TenantHandler) Create(c *gin.Context) {
	var req struct {
		TenantName    string  `json:"tenant_name" binding:"required"`
		TenantKey     string  `json:"tenant_key" binding:"required"`
		Province      *string `json:"province"`
		City          *string `json:"city"`
		District      *string `json:"district"`
		Address       *string `json:"address"`
		Manager       *string `json:"manager"`
		ContactName   *string `json:"contact_name"`
		ContactPhone  *string `json:"contact_phone"`
		ContactEmail  *string `json:"contact_email"`
		FactoryType   *string `json:"factory_type"`
		EmployeeCount *int    `json:"employee_count"`
		AreaSize      *float64 `json:"area_size"`
		AnnualCapacity *float64 `json:"annual_capacity"`
		Status        int     `json:"status"`
		ExpireTime    *string `json:"expire_time"`
		Remark        *string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenant := &model.Tenant{
		TenantName:    req.TenantName,
		TenantKey:    req.TenantKey,
		Province:     req.Province,
		City:         req.City,
		District:     req.District,
		Address:      req.Address,
		Manager:      req.Manager,
		ContactName:  req.ContactName,
		ContactPhone: req.ContactPhone,
		ContactEmail: req.ContactEmail,
		FactoryType:  req.FactoryType,
		Status:       1,
		Remark:       req.Remark,
	}

	if req.EmployeeCount != nil {
		tenant.EmployeeCount = req.EmployeeCount
	}
	if req.AreaSize != nil {
		tenant.AreaSize = req.AreaSize
	}
	if req.AnnualCapacity != nil {
		tenant.AnnualCapacity = req.AnnualCapacity
	}
	if req.Status != 0 {
		tenant.Status = req.Status
	}
	if req.ExpireTime != nil && *req.ExpireTime != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", *req.ExpireTime)
		tenant.ExpireTime = &t
	}

	if err := h.tenantSvc.Create(c.Request.Context(), tenant); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, tenant)
}

func (h *TenantHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.tenantSvc.Update(c.Request.Context(), id, req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *TenantHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.tenantSvc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
