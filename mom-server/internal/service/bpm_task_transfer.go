package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// BpmTaskTransferService 任务转移/候选人服务
type BpmTaskTransferService struct {
	transferRepo *repository.BpmTaskTransferRepository
	userRepo     *repository.UserRepository
}

func NewBpmTaskTransferService(
	transferRepo *repository.BpmTaskTransferRepository,
	userRepo *repository.UserRepository,
) *BpmTaskTransferService {
	return &BpmTaskTransferService{
		transferRepo: transferRepo,
		userRepo:     userRepo,
	}
}

// TransferTaskReq 转移任务请求
type TransferTaskReq struct {
	TaskID   string `json:"taskId"`
	ToUserID uint64 `json:"toUserId"`
	Reason   string `json:"reason"`
}

// TransferTask 转移任务：记录转移历史 + 更新任务assignee
func (s *BpmTaskTransferService) TransferTask(ctx context.Context, tenantID int64, req *TransferTaskReq, operatorID int64) (*model.BpmTaskTransfer, error) {
	if req.TaskID == "" {
		return nil, fmt.Errorf("taskId is required")
	}
	if req.ToUserID == 0 {
		return nil, fmt.Errorf("toUserId is required")
	}

	// 查询目标用户信息
	toUser, err := s.userRepo.FindByID(ctx, int64(req.ToUserID))
	if err != nil {
		return nil, fmt.Errorf("目标用户不存在: %w", err)
	}
	toUserName := toUser.Username

	// 查询当前任务信息（获取当前assignee作为fromUser）
	task, err := s.transferRepo.GetTaskInstance(ctx, req.TaskID)
	if err != nil {
		return nil, fmt.Errorf("任务不存在: %w", err)
	}

	fromUserID := uint64(0)
	fromUserName := ""
	if task.AssigneeID != nil {
		fromUserID = uint64(*task.AssigneeID)
	}
	if task.AssigneeName != nil {
		fromUserName = *task.AssigneeName
	}

	// 更新任务assignee
	if err := s.transferRepo.UpdateTaskAssignee(ctx, req.TaskID, int64(req.ToUserID), toUserName); err != nil {
		return nil, fmt.Errorf("更新任务处理人失败: %w", err)
	}

	// 构建并保存转移记录
	transfer := &model.BpmTaskTransfer{
		TenantID:       tenantID,
		TaskID:         req.TaskID,
		FromUserID:     fromUserID,
		FromUserName:   fromUserName,
		ToUserID:       req.ToUserID,
		ToUserName:     toUserName,
		TransferReason: req.Reason,
		TransferTime:   time.Now(),
		OperatorID:     uint64(operatorID),
	}
	if err := s.transferRepo.Create(ctx, transfer); err != nil {
		return nil, fmt.Errorf("保存转移记录失败: %w", err)
	}

	return transfer, nil
}

// GetTransferHistory 获取任务转移历史
func (s *BpmTaskTransferService) GetTransferHistory(ctx context.Context, taskID string) ([]model.BpmTaskTransfer, error) {
	if taskID == "" {
		return nil, fmt.Errorf("taskId is required")
	}
	return s.transferRepo.ListByTaskID(ctx, taskID)
}

// GetTaskCandidates 获取任务候选人列表（从TaskInstance.AssigneeList解析）
func (s *BpmTaskTransferService) GetTaskCandidates(ctx context.Context, taskID string) ([]model.TaskCandidateUser, error) {
	task, err := s.transferRepo.GetTaskInstance(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("任务不存在: %w", err)
	}

	var candidates []model.TaskCandidateUser

	// AssigneeList 格式: [{"id":1,"name":"张三"}, ...]
	if len(task.AssigneeList) > 0 {
		var raw []struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		}
		if jsonErr := json.Unmarshal(task.AssigneeList, &raw); jsonErr == nil {
			for _, item := range raw {
				candidates = append(candidates, model.TaskCandidateUser{
					UserID:   item.ID,
					UserName: item.Name,
					TaskID:   taskID,
				})
			}
		}
	}

	// fallback：返回当前处理人
	if len(candidates) == 0 && task.AssigneeID != nil {
		name := ""
		if task.AssigneeName != nil {
			name = *task.AssigneeName
		}
		candidates = append(candidates, model.TaskCandidateUser{
			UserID:   *task.AssigneeID,
			UserName: name,
			TaskID:   taskID,
		})
	}

	return candidates, nil
}

// GetTaskCandidateGroups 获取任务候选组列表
func (s *BpmTaskTransferService) GetTaskCandidateGroups(ctx context.Context, taskID string) ([]model.TaskCandidateGroup, error) {
	task, err := s.transferRepo.GetTaskInstance(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("任务不存在: %w", err)
	}

	var groups []model.TaskCandidateGroup
	if len(task.AssigneeList) > 0 {
		var raw []struct {
			GroupID   int64  `json:"groupId"`
			GroupName string `json:"groupName"`
			Members   []struct {
				ID   int64  `json:"id"`
				Name string `json:"name"`
			} `json:"members"`
		}
		if jsonErr := json.Unmarshal(task.AssigneeList, &raw); jsonErr == nil {
			for _, g := range raw {
				if g.GroupID == 0 {
					continue
				}
				group := model.TaskCandidateGroup{
					GroupID:   g.GroupID,
					GroupName: g.GroupName,
				}
				for _, m := range g.Members {
					group.Members = append(group.Members, model.TaskCandidateUser{
						UserID:   m.ID,
						UserName: m.Name,
						TaskID:   taskID,
					})
				}
				groups = append(groups, group)
			}
		}
	}

	return groups, nil
}

// AssignTaskReq 指定候选人执行任务请求
type AssignTaskReq struct {
	TaskID   string `json:"taskId"`
	UserID   uint64 `json:"userId"`
	UserName string `json:"userName"`
}

// AssignTask 指定候选人执行任务（更新assignee，不记录转移历史）
func (s *BpmTaskTransferService) AssignTask(ctx context.Context, req *AssignTaskReq) error {
	if req.TaskID == "" {
		return fmt.Errorf("taskId is required")
	}
	if req.UserID == 0 {
		return fmt.Errorf("userId is required")
	}

	userName := req.UserName
	if userName == "" {
		if user, findErr := s.userRepo.FindByID(ctx, int64(req.UserID)); findErr == nil {
			userName = user.Username
		}
	}

	return s.transferRepo.UpdateTaskAssignee(ctx, req.TaskID, int64(req.UserID), userName)
}
