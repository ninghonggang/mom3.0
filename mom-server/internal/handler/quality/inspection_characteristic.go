package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type InspectionCharacteristicHandler struct {
	svc *service.InspectionCharacteristicService
}

func NewInspectionCharacteristicHandler(svc *service.InspectionCharacteristicService) *InspectionCharacteristicHandler {
	return &InspectionCharacteristicHandler{svc: svc}
}

func (h *InspectionCharacteristicHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	query := c.Query("query")
	list, total, err := h.svc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *InspectionCharacteristicHandler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}

func (h *InspectionCharacteristicHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.InspectionCharacteristic
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.TenantID = tenantID
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *InspectionCharacteristicHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.InspectionCharacteristic
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), id, map[string]any{
		"name":               req.Name,
		"type":               req.Type,
		"spec_lower":         req.SpecLower,
		"spec_upper":         req.SpecUpper,
		"usl":               req.USL,
		"lsl":               req.LSL,
		"target":             req.Target,
		"unit":               req.Unit,
		"aql":               req.AQL,
		"inspection_method":   req.InspectionMethod,
		"status":             req.Status,
	}); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *InspectionCharacteristicHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *InspectionCharacteristicHandler) RegisterRoutes(g *gin.RouterGroup) {
	char := g.Group("/characteristics")
	{
		char.GET("/list", h.List)
		char.GET("/:id", h.Get)
		char.POST("", h.Create)
		char.PUT("/:id", h.Update)
		char.DELETE("/:id", h.Delete)
	}
}
