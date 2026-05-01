package aps

import (
	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SimulationHandler struct {
	db *gorm.DB
}

func NewSimulationHandler(db *gorm.DB) *SimulationHandler {
	return &SimulationHandler{db: db}
}

func (h *SimulationHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	response.Success(c, gin.H{"list": []interface{}{}, "total": 0})
}

func (h *SimulationHandler) Get(c *gin.Context) {
	id := c.Param("id")
	_ = id
	response.Success(c, gin.H{})
}

func (h *SimulationHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	_ = tenantID
	var req interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, gin.H{})
}

func (h *SimulationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	_ = id
	var req interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, gin.H{})
}

func (h *SimulationHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	_ = id
	response.Success(c, nil)
}

func (h *SimulationHandler) Run(c *gin.Context) {
	id := c.Param("id")
	_ = id
	response.Success(c, gin.H{"list": []interface{}{}, "total": 0})
}

func (h *SimulationHandler) Confirm(c *gin.Context) {
	id := c.Param("id")
	_ = id
	response.Success(c, nil)
}
