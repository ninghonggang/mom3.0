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
