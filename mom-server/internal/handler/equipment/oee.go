package equipment

import (
	"strconv"

	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type OEEHandler struct {
	svc *service.OEEService
}

func NewOEEHandler(svc *service.OEEService) *OEEHandler {
	return &OEEHandler{svc: svc}
}

// List OEE列表
func (h *OEEHandler) List(c *gin.Context) {
	params := make(map[string]interface{})

	// 获取查询参数
	if equipID := c.Query("equipment_id"); equipID != "" {
		if id, err := strconv.ParseInt(equipID, 10, 64); err == nil && id > 0 {
			params["equipment_id"] = id
		}
	}
	if startDate := c.Query("start_date"); startDate != "" {
		params["start_date"] = startDate
	}
	if endDate := c.Query("end_date"); endDate != "" {
		params["end_date"] = endDate
	}
	if workshopID := c.Query("workshop_id"); workshopID != "" {
		if id, err := strconv.ParseInt(workshopID, 10, 64); err == nil && id > 0 {
			params["workshop_id"] = id
		}
	}
	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	params["page"] = page
	params["page_size"] = pageSize

	list, total, err := h.svc.List(c.Request.Context(), params)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
		"page":  page,
		"page_size": pageSize,
	})
}

// Get OEE详情
func (h *OEEHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	oee, err := h.svc.GetByID(c.Request.Context(), int64(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, oee)
}

// Calculate 计算OEE
func (h *OEEHandler) Calculate(c *gin.Context) {
	var req service.OECalculateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	oee, err := h.svc.Calculate(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, oee)
}

// Chart 获取图表数据
func (h *OEEHandler) Chart(c *gin.Context) {
	params := make(map[string]interface{})

	if equipID := c.Query("equipment_id"); equipID != "" {
		if id, err := strconv.ParseInt(equipID, 10, 64); err == nil && id > 0 {
			params["equipment_id"] = id
		}
	}
	if startDate := c.Query("start_date"); startDate != "" {
		params["start_date"] = startDate
	}
	if endDate := c.Query("end_date"); endDate != "" {
		params["end_date"] = endDate
	}

	data, err := h.svc.GetChartData(c.Request.Context(), params)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": data})
}

// Delete 删除OEE记录
func (h *OEEHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), int64(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
