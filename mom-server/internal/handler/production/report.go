package production

import (
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service *service.ProductionReportService
}

func NewReportHandler(s *service.ProductionReportService) *ReportHandler {
	return &ReportHandler{service: s}
}

func (h *ReportHandler) List(c *gin.Context) {
	list, total, err := h.service.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *ReportHandler) Create(c *gin.Context) {
	var req model.ProductionReport
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

type DispatchHandler struct {
	service *service.DispatchService
}

func NewDispatchHandler(s *service.DispatchService) *DispatchHandler {
	return &DispatchHandler{service: s}
}

func (h *DispatchHandler) List(c *gin.Context) {
	list, total, err := h.service.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

func (h *DispatchHandler) Create(c *gin.Context) {
	var req model.Dispatch
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, req)
}

func (h *DispatchHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Dispatch
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

func (h *DispatchHandler) Start(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Start(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *DispatchHandler) Complete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Complete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
