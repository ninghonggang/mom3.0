package system

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type OperLogHandler struct {
	operLogSvc *service.OperLogService
}

func NewOperLogHandler(operLogSvc *service.OperLogService) *OperLogHandler {
	return &OperLogHandler{operLogSvc: operLogSvc}
}

func (h *OperLogHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	title := c.Query("title")
	operName := c.Query("oper_name")
	businessType := c.Query("business_type")
	status := c.Query("status")
	tenantID := middleware.GetTenantID(c)

	list, total, err := h.operLogSvc.GetList(c.Request.Context(), tenantID, title, operName, businessType, status, page, pageSize)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}
