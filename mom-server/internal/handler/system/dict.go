package system

import (
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type DictHandler struct {
	dictService *service.DictService
}

func NewDictHandler(ds *service.DictService) *DictHandler {
	return &DictHandler{dictService: ds}
}

func (h *DictHandler) ListType(c *gin.Context) {
	list, total, err := h.dictService.ListType(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *DictHandler) GetType(c *gin.Context) {
	id := c.Param("id")
	dict, err := h.dictService.GetTypeByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, dict)
}

func (h *DictHandler) CreateType(c *gin.Context) {
	var req model.DictType
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.dictService.CreateType(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *DictHandler) UpdateType(c *gin.Context) {
	id := c.Param("id")
	var req model.DictType
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.dictService.UpdateType(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *DictHandler) DeleteType(c *gin.Context) {
	id := c.Param("id")
	err := h.dictService.DeleteType(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *DictHandler) GetData(c *gin.Context) {
	dictType := c.Param("dictType")
	data, err := h.dictService.GetDataByType(c.Request.Context(), dictType)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, data)
}
