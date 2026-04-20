package bpm

import (
	"strconv"

	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// BpmInstanceApiHandler 跨模块流程实例API处理器
type BpmInstanceApiHandler struct {
	instanceApi *service.BpmProcessInstanceApi
	messageSvc  *service.BpmMessageService
}

// NewBpmInstanceApiHandler creates a new BPM instance API handler
func NewBpmInstanceApiHandler(instanceApi *service.BpmProcessInstanceApi, messageSvc *service.BpmMessageService) *BpmInstanceApiHandler {
	return &BpmInstanceApiHandler{
		instanceApi: instanceApi,
		messageSvc:  messageSvc,
	}
}

// StartProcessInstance 启动流程实例
// POST /bpm/instance-api/start
func (h *BpmInstanceApiHandler) StartProcessInstance(c *gin.Context) {
	var req service.StartProcessReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.ProcessDefKey == "" {
		response.BadRequest(c, "processDefKey is required")
		return
	}

	instance, err := h.instanceApi.StartProcessInstance(c.Request.Context(), &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	// Send task created message if there are initial tasks
	response.Success(c, instance)
}

// CompleteTask 完成任务
// POST /bpm/instance-api/complete
func (h *BpmInstanceApiHandler) CompleteTask(c *gin.Context) {
	var req service.CompleteTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.TaskID <= 0 {
		response.BadRequest(c, "taskId is required")
		return
	}

	if err := h.instanceApi.CompleteTask(c.Request.Context(), req.TaskID, req.Variables); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Task completed successfully"})
}

// GetProcessInstance 获取流程实例
// GET /bpm/instance-api/:id
func (h *BpmInstanceApiHandler) GetProcessInstance(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "instance id is required")
		return
	}

	instanceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid instance id")
		return
	}

	instance, err := h.instanceApi.GetProcessInstance(c.Request.Context(), instanceID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}

	response.Success(c, instance)
}

// SendTaskCreatedMessage 发送任务创建消息（内部使用）
func (h *BpmInstanceApiHandler) SendTaskCreatedMessage(c *gin.Context, taskID int64, taskName string, assignee string) {
	if err := h.messageSvc.SendTaskCreatedMessage(c.Request.Context(), taskID, taskName, assignee); err != nil {
		// Log error but don't fail the request
		c.Error(err)
	}
}

// SendProcessCompletedMessage 发送流程完成消息（内部使用）
func (h *BpmInstanceApiHandler) SendProcessCompletedMessage(c *gin.Context, processInstanceID int64, businessKey string) {
	if err := h.messageSvc.SendProcessCompletedMessage(c.Request.Context(), processInstanceID, businessKey); err != nil {
		// Log error but don't fail the request
		c.Error(err)
	}
}
