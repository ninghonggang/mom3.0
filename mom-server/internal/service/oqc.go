package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type OQCService struct {
	repo *repository.OQCRepository
}

func NewOQCService(repo *repository.OQCRepository) *OQCService {
	return &OQCService{repo: repo}
}

func (s *OQCService) List(ctx context.Context) ([]model.OQC, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *OQCService) GetByID(ctx context.Context, id string) (*model.OQC, error) {
	var oqcID uint
	_, err := fmt.Sscanf(id, "%d", &oqcID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, oqcID)
}

func (s *OQCService) Create(ctx context.Context, oqc *model.OQC) error {
	oqc.TenantID = 1
	return s.repo.Create(ctx, oqc)
}

func (s *OQCService) Update(ctx context.Context, id string, oqc *model.OQC) error {
	var oqcID uint
	_, err := fmt.Sscanf(id, "%d", &oqcID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"shipping_no":   oqc.ShippingNo,
		"customer_id":   oqc.CustomerID,
		"customer_name": oqc.CustomerName,
		"quantity":      oqc.Quantity,
		"result":        oqc.Result,
		"remark":        oqc.Remark,
	}
	return s.repo.Update(ctx, oqcID, updates)
}

func (s *OQCService) Delete(ctx context.Context, id string) error {
	var oqcID uint
	_, err := fmt.Sscanf(id, "%d", &oqcID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, oqcID)
}
