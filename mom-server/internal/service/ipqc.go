package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type IPQCService struct {
	repo *repository.IPQCRepository
}

func NewIPQCService(repo *repository.IPQCRepository) *IPQCService {
	return &IPQCService{repo: repo}
}

func (s *IPQCService) List(ctx context.Context) ([]model.IPQC, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *IPQCService) GetByID(ctx context.Context, id string) (*model.IPQC, error) {
	var ipqcID uint
	_, err := fmt.Sscanf(id, "%d", &ipqcID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, ipqcID)
}

func (s *IPQCService) Create(ctx context.Context, ipqc *model.IPQC) error {
	ipqc.TenantID = 1
	return s.repo.Create(ctx, ipqc)
}

func (s *IPQCService) Update(ctx context.Context, id string, ipqc *model.IPQC) error {
	var ipqcID uint
	_, err := fmt.Sscanf(id, "%d", &ipqcID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"process_id":    ipqc.ProcessID,
		"process_name":  ipqc.ProcessName,
		"quantity":      ipqc.Quantity,
		"sample_size":   ipqc.SampleSize,
		"result":        ipqc.Result,
		"remark":        ipqc.Remark,
	}
	return s.repo.Update(ctx, ipqcID, updates)
}

func (s *IPQCService) Delete(ctx context.Context, id string) error {
	var ipqcID uint
	_, err := fmt.Sscanf(id, "%d", &ipqcID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, ipqcID)
}
