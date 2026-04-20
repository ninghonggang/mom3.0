package bpm

import (
	"mom-server/internal/middleware"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

// TaskTransferHandler 任务转移/候选人处理器
type TaskTransferHandler struct {
	svc *service.BpmTaskTransferService
}

func NewTaskTransferHandler(svc *service.BpmTaskTransferService) *TaskTransferHandler {
	return &TaskTransferHandler{svc: svc}
}

// TransferTask POST /bpm/task/transfer
// 请求体: {taskId string, toUserId uint64, reason string}
func (h *TaskTransferHandler) TransferTask(c *gin.Context) {
	var req service.TransferTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	operatorID := middleware.GetUserID(c)

	transfer, err := h.svc.TransferTask(c.Request.Context(), tenantID, &req, operatorID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, transfer)
}

// GetTransferHistory GET /bpm/task/transfer/history/:taskId
func (h *TaskTransferHandler) GetTransferHistory(c *gin.Context) {
	taskID := c.Param("taskId")
	if taskID == "" {
		response.BadRequest(c, "taskId is required")
		return
	}

	history, err := h.svc.GetTransferHistory(c.Request.Context(), taskID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": history})
}

// GetTaskCandidates GET /bpm/task/candidate/:taskId
func (h *TaskTransferHandler) GetTaskCandidates(c *gin.Context) {
	taskID := c.Param("taskId")
	if taskID == "" {
		response.BadRequest(c, "taskId is required")
		return
	}

	candidates, err := h.svc.GetTaskCandidates(c.Request.Context(), taskID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": candidates})
}

// GetTaskCandidateGroups GET /bpm/task/candidate-group/:taskId
func (h *TaskTransferHandler) GetTaskCandidateGroups(c *gin.Context) {
	taskID := c.Param("taskId")
	if taskID == "" {
		response.BadRequest(c, "taskId is required")
		return
	}

	groups, err := h.svc.GetTaskCandidateGroups(c.Request.Context(), taskID)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": groups})
}

// AssignTask POST /bpm/task/assign
// 指定候选人执行任务
func (h *TaskTransferHandler) AssignTask(c *gin.Context) {
	var req service.AssignTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.AssignTask(c.Request.Context(), &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}
