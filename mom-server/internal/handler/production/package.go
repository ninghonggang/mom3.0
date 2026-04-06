package production

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"mom-server/internal/service"
	"mom-server/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type PackageHandler struct {
	service *service.PackageService
}

func NewPackageHandler(s *service.PackageService) *PackageHandler {
	return &PackageHandler{service: s}
}

func (h *PackageHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}
	req := &repository.PackageQuery{
		PackageNo:   c.Query("package_no"),
		Status:      c.Query("status"),
		PackageType: c.Query("package_type"),
		Limit:       pageSize,
		Offset:      (page - 1) * pageSize,
	}
	list, total, err := h.service.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

func (h *PackageHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	item, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}

// CreatePackage 创建箱
type CreatePackageReq struct {
	ProductID         int64  `json:"product_id" binding:"required"`
	ProductCode       string `json:"product_code"`
	ProductionOrderID int64  `json:"production_order_id"`
	PackageType       string `json:"package_type"` // SMALL_BOX/BIG_BOX/PALLET
	Qty               int    `json:"qty"`
	CustomerID        int64  `json:"customer_id"`
	ContainerID       int64  `json:"container_id"`
}

func (h *PackageHandler) Create(c *gin.Context) {
	var req CreatePackageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	pkg := &model.MesPackage{
		TenantID:    tenantID,
		ProductID:   req.ProductID,
		ProductCode: req.ProductCode,
	}
	if req.ProductionOrderID > 0 {
		pkg.ProductionOrderID = &req.ProductionOrderID
	}
	pkg.PackageType = req.PackageType
	pkg.Qty = req.Qty
	if req.CustomerID > 0 {
		pkg.CustomerID = &req.CustomerID
	}
	if req.ContainerID > 0 {
		pkg.ContainerID = &req.ContainerID
	}
	if err := h.service.CreatePackage(c.Request.Context(), pkg); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, pkg)
}

// AddItem 添加序列号到箱
type AddItemReq struct {
	PackageNo string `json:"package_no" binding:"required"`
	SerialNo  string `json:"serial_no" binding:"required"`
}

func (h *PackageHandler) AddItem(c *gin.Context) {
	var req AddItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	pkg, err := h.service.GetByPackageNo(c.Request.Context(), req.PackageNo)
	if err != nil {
		response.ErrorMsg(c, "箱不存在")
		return
	}
	if err := h.service.AddSerialNo(c.Request.Context(), pkg.ID, req.SerialNo); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "添加成功"})
}

// Seal 封箱
type SealReq struct {
	PackageNo string `json:"package_no" binding:"required"`
	SealBy    string `json:"seal_by"`
}

func (h *PackageHandler) Seal(c *gin.Context) {
	var req SealReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	pkg, err := h.service.GetByPackageNo(c.Request.Context(), req.PackageNo)
	if err != nil {
		response.ErrorMsg(c, "箱不存在")
		return
	}
	if err := h.service.SealPackage(c.Request.Context(), pkg.ID, req.SealBy); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "封箱成功"})
}

func (h *PackageHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}
