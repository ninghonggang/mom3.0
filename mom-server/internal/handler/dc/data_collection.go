package dc

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"mom-server/internal/service"
	"mom-server/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type DataCollectionHandler struct {
	service *service.DCService
}

func NewDataCollectionHandler(s *service.DCService) *DataCollectionHandler {
	return &DataCollectionHandler{service: s}
}

// DataPoint List
func (h *DataCollectionHandler) ListDataPoint(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	query := c.Query("query")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}

	list, total, err := h.service.ListDataPoints(c.Request.Context(), tenantID, query)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// DataPoint Get
func (h *DataCollectionHandler) GetDataPoint(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	item, err := h.service.GetDataPoint(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}

// DataPoint Create
func (h *DataCollectionHandler) CreateDataPoint(c *gin.Context) {
	var item model.DCDataPoint
	if err := c.ShouldBindJSON(&item); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	item.TenantID = tenantID
	if err := h.service.CreateDataPoint(c.Request.Context(), &item); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, item)
}

// DataPoint Update
func (h *DataCollectionHandler) UpdateDataPoint(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.service.UpdateDataPoint(c.Request.Context(), uint(id), updates); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "更新成功"})
}

// DataPoint Delete
func (h *DataCollectionHandler) DeleteDataPoint(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.service.DeleteDataPoint(c.Request.Context(), uint(id)); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ScanLog List
func (h *DataCollectionHandler) ListScanLog(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	req := &repository.DCScanLogQuery{
		ScanType:  c.Query("scan_type"),
		Status:    c.Query("status"),
		StartTime: c.Query("start_time"),
		EndTime:   c.Query("end_time"),
		Limit:     pageSize,
		Offset:    offset,
	}
	list, total, err := h.service.ListScanLogs(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// ScanLog Create (扫码入系统)
func (h *DataCollectionHandler) CreateScanLog(c *gin.Context) {
	var log model.DCScanLog
	if err := c.ShouldBindJSON(&log); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	log.TenantID = tenantID
	if err := h.service.CreateScanLog(c.Request.Context(), &log); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, log)
}

// CollectRecord List
func (h *DataCollectionHandler) ListCollectRecord(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	pointID, _ := strconv.ParseInt(c.Query("point_id"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	list, total, err := h.service.ListCollectRecords(
		c.Request.Context(), tenantID, pointID,
		c.Query("start_time"), c.Query("end_time"),
		pageSize, offset,
	)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}
