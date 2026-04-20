package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type NCRService struct {
	repo *repository.NCRRepository
}

func NewNCRService(repo *repository.NCRRepository) *NCRService {
	return &NCRService{repo: repo}
}

func (s *NCRService) List(ctx context.Context) ([]model.NCR, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *NCRService) GetByID(ctx context.Context, id string) (*model.NCR, error) {
	var ncrID uint
	_, err := fmt.Sscanf(id, "%d", &ncrID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, ncrID)
}

func (s *NCRService) Create(ctx context.Context, ncr *model.NCR) error {
	ncr.TenantID = 1
	return s.repo.Create(ctx, ncr)
}

func (s *NCRService) Update(ctx context.Context, id string, ncr *model.NCR) error {
	var ncrID uint
	_, err := fmt.Sscanf(id, "%d", &ncrID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"source_type":        ncr.SourceType,
		"issue_desc":        ncr.IssueDesc,
		"root_cause":        ncr.RootCause,
		"corrective_action": ncr.CorrectiveAction,
		"preventive_action": ncr.PreventiveAction,
		"verify_result":     ncr.VerifyResult,
		"verify_user_id":    ncr.VerifyUserID,
		"status":            ncr.Status,
	}
	return s.repo.Update(ctx, ncrID, updates)
}

func (s *NCRService) Delete(ctx context.Context, id string) error {
	var ncrID uint
	_, err := fmt.Sscanf(id, "%d", &ncrID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, ncrID)
}

// Resolve NCR解决 - 更新处理结果和状态
func (s *NCRService) Resolve(ctx context.Context, id string, rootCause, correctiveAction, preventiveAction, verifyResult string, verifyUserID int64) error {
	var ncrID uint
	_, err := fmt.Sscanf(id, "%d", &ncrID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, ncrID, map[string]interface{}{
		"root_cause":         rootCause,
		"corrective_action":  correctiveAction,
		"preventive_action":  preventiveAction,
		"verify_result":      verifyResult,
		"verify_user_id":     verifyUserID,
		"status":             3, // 3=已完成
	})
}

// Assign NCR指派 - 指派处理人
func (s *NCRService) Assign(ctx context.Context, id string, handleUserID int64) error {
	var ncrID uint
	_, err := fmt.Sscanf(id, "%d", &ncrID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, ncrID, map[string]interface{}{
		"handle_user_id": handleUserID,
		"status":         2, // 2=处理中
	})
}

// Close NCR关闭 - 关闭 NCR 单据
func (s *NCRService) Close(ctx context.Context, id string) error {
	var ncrID uint
	_, err := fmt.Sscanf(id, "%d", &ncrID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, ncrID, map[string]interface{}{
		"status": 4, // 4=已关闭
	})
}
