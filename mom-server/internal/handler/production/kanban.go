package production

import (
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type KanbanHandler struct {
	service *service.KanbanService
}

func NewKanbanHandler(s *service.KanbanService) *KanbanHandler {
	return &KanbanHandler{service: s}
}

// GetDashboard 获取看板数据
func (h *KanbanHandler) GetDashboard(c *gin.Context) {
	data, err := h.service.GetDashboardData(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, data)
}
