package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type SPCDataService struct {
	repo *repository.SPCDataRepository
}

func NewSPCDataService(repo *repository.SPCDataRepository) *SPCDataService {
	return &SPCDataService{repo: repo}
}

func (s *SPCDataService) List(ctx context.Context) ([]model.SPCData, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *SPCDataService) GetByID(ctx context.Context, id string) (*model.SPCData, error) {
	var spcID uint
	_, err := fmt.Sscanf(id, "%d", &spcID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, spcID)
}

func (s *SPCDataService) Create(ctx context.Context, spcData *model.SPCData) error {
	spcData.TenantID = 1
	return s.repo.Create(ctx, spcData)
}

func (s *SPCDataService) Update(ctx context.Context, id string, spcData *model.SPCData) error {
	var spcID uint
	_, err := fmt.Sscanf(id, "%d", &spcID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"check_value": spcData.CheckValue,
		"usl":         spcData.USL,
		"lsl":         spcData.LSL,
		"cl":          spcData.CL,
		"ucl":         spcData.UCL,
		"lcl":         spcData.LCL,
	}
	return s.repo.Update(ctx, spcID, updates)
}

func (s *SPCDataService) Delete(ctx context.Context, id string) error {
	var spcID uint
	_, err := fmt.Sscanf(id, "%d", &spcID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, spcID)
}

type SPCChartQuery struct {
	EquipmentID int64
	ProcessID   int64
	StationID   int64
	CheckItem   string
	Limit       int
}

func (s *SPCDataService) GetChartData(ctx context.Context, query SPCChartQuery) ([]model.SPCData, error) {
	if query.Limit <= 0 {
		query.Limit = 100
	}
	return s.repo.GetChartData(ctx, 1, query.EquipmentID, query.ProcessID, query.StationID, query.CheckItem, query.Limit)
}
