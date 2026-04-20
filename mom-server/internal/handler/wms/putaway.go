package wms

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type PutawayHandler struct {
	putawaySvc *service.PutawayService
}

func NewPutawayHandler(putawaySvc *service.PutawayService) *PutawayHandler {
	return &PutawayHandler{putawaySvc: putawaySvc}
}

// Create POST /wms/putaway/create
func (h *PutawayHandler) Create(c *gin.Context) {
	var req model.WMSPutawayJobCreateReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	username := middleware.GetUsername(c)

	ret, err := h.putawaySvc.Create(c.Request.Context(), tenantID, &req, username)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ret)
}

// Assign PUT /wms/putaway/assign
func (h *PutawayHandler) Assign(c *gin.Context) {
	var req model.WMSPutawayJobAssignReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.putawaySvc.Assign(c.Request.Context(), uint(req.ID), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// List GET /wms/putaway/list
func (h *PutawayHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}

	query := map[string]interface{}{}
	if status := c.Query("status"); status != "" {
		query["status"] = status
	}
	if sourceType := c.Query("source_type"); sourceType != "" {
		query["source_type"] = sourceType
	}
	if sourceNo := c.Query("source_no"); sourceNo != "" {
		query["source_no"] = sourceNo
	}
	if warehouseID := c.Query("warehouse_id"); warehouseID != "" {
		if id, err := strconv.ParseInt(warehouseID, 10, 64); err == nil {
			query["warehouse_id"] = id
		}
	}
	if page := c.Query("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			query["page"] = p
		}
	}
	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			query["limit"] = l
		}
	}

	list, total, err := h.putawaySvc.List(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total})
}

// Get GET /wms/putaway/:id
func (h *PutawayHandler) Get(c *gin.Context) {
	id := toUint64(c.Param("id"))
	ret, err := h.putawaySvc.GetWithRecords(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, ret)
}

// Start PUT /wms/putaway/:id/start
func (h *PutawayHandler) Start(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.putawaySvc.Start(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Complete PUT /wms/putaway/:id/complete
func (h *PutawayHandler) Complete(c *gin.Context) {
	id := toUint64(c.Param("id"))
	if err := h.putawaySvc.Complete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Cancel PUT /wms/putaway/cancel
func (h *PutawayHandler) Cancel(c *gin.Context) {
	var req model.WMSPutawayCancelReqVO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.putawaySvc.Cancel(c.Request.Context(), uint(req.ID), req.Reason); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
