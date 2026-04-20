package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type NCRHandler struct {
	service *service.NCRService
}

func NewNCRHandler(s *service.NCRService) *NCRHandler {
	return &NCRHandler{service: s}
}

func (h *NCRHandler) List(c *gin.Context) {
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

func (h *NCRHandler) Get(c *gin.Context) {
	id := c.Param("id")
	ncr, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ncr)
}

func (h *NCRHandler) Create(c *gin.Context) {
	var req model.NCR
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

func (h *NCRHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.NCR
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

func (h *NCRHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// NCRResolveReq NCR解决请求
type NCRResolveReq struct {
	RootCause         string `json:"rootCause"`
	CorrectiveAction  string `json:"correctiveAction"`
	PreventiveAction  string `json:"preventiveAction"`
	VerifyResult      string `json:"verifyResult"`
	VerifyUserID      int64  `json:"verifyUserId"`
}

// Resolve NCR解决
func (h *NCRHandler) Resolve(c *gin.Context) {
	id := c.Param("id")
	var req NCRResolveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Resolve(c.Request.Context(), id, req.RootCause, req.CorrectiveAction, req.PreventiveAction, req.VerifyResult, req.VerifyUserID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// NCRAssignReq NCR指派请求
type NCRAssignReq struct {
	HandleUserID int64 `json:"handleUserId" binding:"required"`
}

// Assign NCR指派
func (h *NCRHandler) Assign(c *gin.Context) {
	id := c.Param("id")
	var req NCRAssignReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Assign(c.Request.Context(), id, req.HandleUserID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Close NCR关闭
func (h *NCRHandler) Close(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Close(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
