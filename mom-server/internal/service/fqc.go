package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type FQCService struct {
	repo *repository.FQCRepository
}

func NewFQCService(repo *repository.FQCRepository) *FQCService {
	return &FQCService{repo: repo}
}

func (s *FQCService) List(ctx context.Context) ([]model.FQC, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *FQCService) GetByID(ctx context.Context, id string) (*model.FQC, error) {
	var fqcID uint
	_, err := fmt.Sscanf(id, "%d", &fqcID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, fqcID)
}

func (s *FQCService) Create(ctx context.Context, fqc *model.FQC) error {
	fqc.TenantID = 1
	return s.repo.Create(ctx, fqc)
}

func (s *FQCService) Update(ctx context.Context, id string, fqc *model.FQC) error {
	var fqcID uint
	_, err := fmt.Sscanf(id, "%d", &fqcID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"quantity":      fqc.Quantity,
		"sample_size":   fqc.SampleSize,
		"qualified_qty": fqc.QualifiedQty,
		"rejected_qty": fqc.RejectedQty,
		"result":        fqc.Result,
		"remark":        fqc.Remark,
	}
	return s.repo.Update(ctx, fqcID, updates)
}

func (s *FQCService) Delete(ctx context.Context, id string) error {
	var fqcID uint
	_, err := fmt.Sscanf(id, "%d", &fqcID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, fqcID)
}
