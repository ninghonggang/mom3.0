package production

import (
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductionOrderChangeLogHandler struct {
	service *service.ProductionOrderChangeLogService
}

func NewProductionOrderChangeLogHandler(s *service.ProductionOrderChangeLogService) *ProductionOrderChangeLogHandler {
	return &ProductionOrderChangeLogHandler{service: s}
}

// List 获取变更历史列表
func (h *ProductionOrderChangeLogHandler) List(c *gin.Context) {
	orderIDStr := c.Query("orderId")
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil || orderID == 0 {
		response.ErrorMsg(c, "invalid orderId")
		return
	}

	list, err := h.service.GetOrderChanges(c.Request.Context(), orderID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  list,
		"total": len(list),
	})
}
