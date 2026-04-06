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

type FirstLastInspectHandler struct {
	service *service.FirstLastInspectService
}

func NewFirstLastInspectHandler(s *service.FirstLastInspectService) *FirstLastInspectHandler {
	return &FirstLastInspectHandler{service: s}
}

func (h *FirstLastInspectHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}
	req := &repository.FirstLastInspectQuery{
		OrderNo:     c.Query("order_no"),
		InspectType: c.Query("inspect_type"),
		Status:      c.Query("status"),
		StartTime:   c.Query("start_time"),
		EndTime:     c.Query("end_time"),
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

func (h *FirstLastInspectHandler) Get(c *gin.Context) {
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

func (h *FirstLastInspectHandler) Create(c *gin.Context) {
	var item model.MesFirstLastInspect
	if err := c.ShouldBindJSON(&item); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	item.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &item); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}

func (h *FirstLastInspectHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "更新成功"})
}

func (h *FirstLastInspectHandler) Delete(c *gin.Context) {
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

func (h *FirstLastInspectHandler) ListOverdue(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	list, err := h.service.ListOverdue(c.Request.Context(), tenantID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": len(list)})
}