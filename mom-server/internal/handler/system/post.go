package system

import (
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(ps *service.PostService) *PostHandler {
	return &PostHandler{postService: ps}
}

func (h *PostHandler) List(c *gin.Context) {
	list, total, err := h.postService.List(c.Request.Context())
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func (h *PostHandler) Get(c *gin.Context) {
	id := c.Param("id")
	post, err := h.postService.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, post)
}

func (h *PostHandler) Create(c *gin.Context) {
	var req model.Post
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

	err := h.postService.Create(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *PostHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.Post
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.postService.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, req)
}

func (h *PostHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.postService.Delete(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
