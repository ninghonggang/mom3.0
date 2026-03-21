package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type IQCService struct {
	repo *repository.IQCRepository
}

func NewIQCService(repo *repository.IQCRepository) *IQCService {
	return &IQCService{repo: repo}
}

func (s *IQCService) List(ctx context.Context) ([]model.IQC, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *IQCService) GetByID(ctx context.Context, id string) (*model.IQC, error) {
	var iqcID uint
	_, err := fmt.Sscanf(id, "%d", &iqcID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, iqcID)
}

func (s *IQCService) Create(ctx context.Context, iqc *model.IQC) error {
	return s.repo.Create(ctx, iqc)
}

func (s *IQCService) Update(ctx context.Context, id string, iqc *model.IQC) error {
	var iqcID uint
	_, err := fmt.Sscanf(id, "%d", &iqcID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, iqcID, map[string]interface{}{
		"result": iqc.Result,
	})
}

func (s *IQCService) Delete(ctx context.Context, id string) error {
	var iqcID uint
	_, err := fmt.Sscanf(id, "%d", &iqcID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, iqcID)
}
