package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"time"
)

type TraceService struct {
	traceRepo *repository.TraceRepository
}

func NewTraceService(traceRepo *repository.TraceRepository) *TraceService {
	return &TraceService{traceRepo: traceRepo}
}

func (s *TraceService) TraceBySerial(ctx context.Context, serialNumber string) (*model.SerialNumber, []model.TraceRecord, error) {
	sn, err := s.traceRepo.GetBySerialNumber(ctx, serialNumber)
	if err != nil {
		return nil, nil, err
	}
	records, err := s.traceRepo.GetTraceRecordsBySerial(ctx, serialNumber)
	if err != nil {
		return sn, nil, nil
	}
	return sn, records, nil
}

func (s *TraceService) TraceByBatch(ctx context.Context, batchNo string) ([]model.SerialNumber, error) {
	return s.traceRepo.GetByBatchNo(ctx, batchNo)
}

func (s *TraceService) TraceByOrder(ctx context.Context, orderID int64) ([]model.SerialNumber, error) {
	return s.traceRepo.GetByOrderID(ctx, orderID)
}

// ForwardTrace 正向追溯 - 从原材料到成品
func (s *TraceService) ForwardTrace(ctx context.Context, serialNumber string) ([]model.TraceRecord, error) {
	return s.traceRepo.GetForwardTrace(ctx, serialNumber)
}

// BackwardTrace 反向追溯 - 从成品到原材料
func (s *TraceService) BackwardTrace(ctx context.Context, serialNumber string) ([]model.TraceRecord, error) {
	return s.traceRepo.GetBackwardTrace(ctx, serialNumber)
}

type EnergyService struct {
	energyRepo *repository.EnergyRepository
}

func NewEnergyService(energyRepo *repository.EnergyRepository) *EnergyService {
	return &EnergyService{energyRepo: energyRepo}
}

func (s *EnergyService) List(ctx context.Context, energyType string, startDate, endDate time.Time) ([]model.EnergyRecord, int64, error) {
	return s.energyRepo.List(ctx, 0, energyType, startDate, endDate)
}

func (s *EnergyService) GetStats(ctx context.Context, startDate, endDate time.Time) (map[string]interface{}, error) {
	return s.energyRepo.GetStats(ctx, 0, startDate, endDate)
}

func (s *EnergyService) GetTrend(ctx context.Context, energyType string, startDate, endDate time.Time) ([]map[string]interface{}, error) {
	return s.energyRepo.GetTrend(ctx, 0, energyType, startDate, endDate)
}

func (s *EnergyService) Create(ctx context.Context, record *model.EnergyRecord) error {
	return s.energyRepo.Create(ctx, record)
}
