package aps

import (
	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WorkOrderHandler struct {
	db *gorm.DB
}

func NewWorkOrderHandler(db *gorm.DB) *WorkOrderHandler {
	return &WorkOrderHandler{db: db}
}

func (h *WorkOrderHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	response.Success(c, gin.H{"list": []interface{}{}, "total": 0})
}

func (h *WorkOrderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}
	response.Success(c, gin.H{})
}

func (h *WorkOrderHandler) Create(c *gin.Context) {
	var req struct {
		WorkOrderNo string `json:"work_order_no"`
		TenantID    int64  `json:"tenant_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	response.Success(c, gin.H{"tenant_id": tenantID})
}

func (h *WorkOrderHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, gin.H{})
}

func (h *WorkOrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}
	response.Success(c, nil)
}
