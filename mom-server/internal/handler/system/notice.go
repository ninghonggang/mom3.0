package system

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoticeHandler struct {
	service *service.NoticeService
}

func NewNoticeHandler(s *service.NoticeService) *NoticeHandler {
	return &NoticeHandler{service: s}
}

func (h *NoticeHandler) List(c *gin.Context) {
	query := c.Query("query")
	list, total, err := h.service.List(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *NoticeHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	notice, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, notice)
}

func (h *NoticeHandler) Create(c *gin.Context) {
	var req model.Notice
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *NoticeHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updates map[string]any
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.Update(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *NoticeHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *NoticeHandler) Publish(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.Publish(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *NoticeHandler) GetMyNotices(c *gin.Context) {
	userID := middleware.GetUserID(c)
	list, err := h.service.GetMyNotices(c.Request.Context(), userID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list})
}
