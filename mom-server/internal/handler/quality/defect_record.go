package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type DefectRecordHandler struct {
	service *service.DefectRecordService
}

func NewDefectRecordHandler(s *service.DefectRecordService) *DefectRecordHandler {
	return &DefectRecordHandler{service: s}
}

func (h *DefectRecordHandler) List(c *gin.Context) {
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

func (h *DefectRecordHandler) Get(c *gin.Context) {
	id := c.Param("id")
	record, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, record)
}

func (h *DefectRecordHandler) Create(c *gin.Context) {
	var req model.DefectRecord
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

func (h *DefectRecordHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.DefectRecord
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

func (h *DefectRecordHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

type HandleRequest struct {
	HandleMethod int   `json:"handle_method"`
	HandleUserID int64 `json:"handle_user_id"`
}

func (h *DefectRecordHandler) Handle(c *gin.Context) {
	id := c.Param("id")
	var req HandleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Handle(c.Request.Context(), id, req.HandleMethod, req.HandleUserID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
