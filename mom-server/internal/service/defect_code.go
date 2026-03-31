package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type DefectCodeService struct {
	repo *repository.DefectCodeRepository
}

func NewDefectCodeService(repo *repository.DefectCodeRepository) *DefectCodeService {
	return &DefectCodeService{repo: repo}
}

func (s *DefectCodeService) List(ctx context.Context) ([]model.DefectCode, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *DefectCodeService) GetByID(ctx context.Context, id string) (*model.DefectCode, error) {
	var defectCodeID uint
	_, err := fmt.Sscanf(id, "%d", &defectCodeID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, defectCodeID)
}

func (s *DefectCodeService) Create(ctx context.Context, defectCode *model.DefectCode) error {
	defectCode.TenantID = 1
	return s.repo.Create(ctx, defectCode)
}

func (s *DefectCodeService) Update(ctx context.Context, id string, defectCode *model.DefectCode) error {
	var defectCodeID uint
	_, err := fmt.Sscanf(id, "%d", &defectCodeID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"defect_code": defectCode.DefectCode,
		"defect_name": defectCode.DefectName,
		"defect_type": defectCode.DefectType,
		"severity":    defectCode.Severity,
		"status":      defectCode.Status,
	}
	return s.repo.Update(ctx, defectCodeID, updates)
}

func (s *DefectCodeService) Delete(ctx context.Context, id string) error {
	var defectCodeID uint
	_, err := fmt.Sscanf(id, "%d", &defectCodeID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, defectCodeID)
}
