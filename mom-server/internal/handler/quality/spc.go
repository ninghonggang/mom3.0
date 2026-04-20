package quality

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SPCHandler struct {
	service *service.SPCDataService
}

func NewSPCHandler(s *service.SPCDataService) *SPCHandler {
	return &SPCHandler{service: s}
}

func (h *SPCHandler) List(c *gin.Context) {
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

func (h *SPCHandler) Get(c *gin.Context) {
	id := c.Param("id")
	spcData, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, spcData)
}

func (h *SPCHandler) Create(c *gin.Context) {
	var req model.SPCData
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

func (h *SPCHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.SPCData
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

func (h *SPCHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SPCHandler) GetChartData(c *gin.Context) {
	query := service.SPCChartQuery{}

	if equipmentID := c.Query("equipment_id"); equipmentID != "" {
		id, _ := strconv.ParseInt(equipmentID, 10, 64)
		query.EquipmentID = id
	}
	if processID := c.Query("process_id"); processID != "" {
		id, _ := strconv.ParseInt(processID, 10, 64)
		query.ProcessID = id
	}
	if stationID := c.Query("station_id"); stationID != "" {
		id, _ := strconv.ParseInt(stationID, 10, 64)
		query.StationID = id
	}
	if checkItem := c.Query("check_item"); checkItem != "" {
		query.CheckItem = checkItem
	}
	if limit := c.Query("limit"); limit != "" {
		l, _ := strconv.Atoi(limit)
		query.Limit = l
	}

	data, err := h.service.GetChartData(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list": data,
	})
}

func (h *SPCHandler) GetStats(c *gin.Context) {
	query := service.SPCChartQuery{}

	if equipmentID := c.Query("equipment_id"); equipmentID != "" {
		id, _ := strconv.ParseInt(equipmentID, 10, 64)
		query.EquipmentID = id
	}
	if processID := c.Query("process_id"); processID != "" {
		id, _ := strconv.ParseInt(processID, 10, 64)
		query.ProcessID = id
	}
	if stationID := c.Query("station_id"); stationID != "" {
		id, _ := strconv.ParseInt(stationID, 10, 64)
		query.StationID = id
	}
	if checkItem := c.Query("check_item"); checkItem != "" {
		query.CheckItem = checkItem
	}
	if limit := c.Query("limit"); limit != "" {
		l, _ := strconv.Atoi(limit)
		query.Limit = l
	} else {
		query.Limit = 100
	}

	data, err := h.service.GetChartData(c.Request.Context(), query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// 提取数值计算CPK
	var values []float64
	var usl, lsl float64 = 0, 0
	hasUSL, hasLSL := false, false

	for _, d := range data {
		values = append(values, d.CheckValue)
		if d.USL != nil && !hasUSL {
			usl = *d.USL
			hasUSL = true
		}
		if d.LSL != nil && !hasLSL {
			lsl = *d.LSL
			hasLSL = true
		}
	}

	if len(values) > 0 && hasUSL && hasLSL {
		stats := service.CalculateCPK(values, usl, lsl)
		response.Success(c, stats)
	} else {
		response.Success(c, service.SPCStats{Count: len(values)})
	}
}

// GetCapability 获取CP/CPK能力分析
func (h *SPCHandler) GetCapability(c *gin.Context) {
	configIDStr := c.Param("configId")
	configID, err := strconv.ParseInt(configIDStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的configId")
		return
	}

	checkItem := c.Query("check_item")

	stats, err := h.service.GetCapability(c.Request.Context(), configID, checkItem)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, stats)
}
