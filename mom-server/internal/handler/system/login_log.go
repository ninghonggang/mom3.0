package system

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type LoginLogHandler struct {
	loginLogSvc *service.LoginLogService
}

func NewLoginLogHandler(loginLogSvc *service.LoginLogService) *LoginLogHandler {
	return &LoginLogHandler{loginLogSvc: loginLogSvc}
}

func (h *LoginLogHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	username := c.Query("username")
	status := c.Query("status")
	ip := c.Query("ip")
	tenantID := middleware.GetTenantID(c)

	list, total, err := h.loginLogSvc.GetList(c.Request.Context(), tenantID, username, status, ip, page, pageSize)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

func (h *LoginLogHandler) Clean(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))
	if err := h.loginLogSvc.DeleteClean(c.Request.Context(), days); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
