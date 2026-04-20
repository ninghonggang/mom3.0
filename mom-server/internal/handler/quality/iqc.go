package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type IQCHandler struct {
	service *service.IQCService
}

func NewIQCHandler(s *service.IQCService) *IQCHandler {
	return &IQCHandler{service: s}
}

func (h *IQCHandler) List(c *gin.Context) {
	list, total, err := h.service.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *IQCHandler) Get(c *gin.Context) {
	id := c.Param("id")
	iqc, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, iqc)
}

func (h *IQCHandler) Create(c *gin.Context) {
	var req model.IQC
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// 设置默认租户ID
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	req.TenantID = tenantID
	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *IQCHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.IQC
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *IQCHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// IQCInspectReq IQC检验判定请求
type IQCInspectReq struct {
	Result int    `json:"result" binding:"required"` // 2=合格/3=不合格
	Remark string `json:"remark"`
}

// Inspect IQC检验判定
func (h *IQCHandler) Inspect(c *gin.Context) {
	id := c.Param("id")
	var req IQCInspectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Inspect(c.Request.Context(), id, req.Result, req.Remark)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
