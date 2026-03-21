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

type AndonService struct {
	andonRepo *repository.AndonRepository
}

func NewAndonService(andonRepo *repository.AndonRepository) *AndonService {
	return &AndonService{andonRepo: andonRepo}
}

func (s *AndonService) List(ctx context.Context, status int, callNo string) ([]model.AndonCall, int64, error) {
	return s.andonRepo.List(ctx, 0, status, callNo)
}

func (s *AndonService) Create(ctx context.Context, call *model.AndonCall) error {
	return s.andonRepo.Create(ctx, call)
}

func (s *AndonService) Response(ctx context.Context, id uint) error {
	return s.andonRepo.Update(ctx, id, map[string]interface{}{
		"status":          2,
		"response_time": time.Now(),
	})
}

func (s *AndonService) Resolve(ctx context.Context, id uint) error {
	return s.andonRepo.Update(ctx, id, map[string]interface{}{
		"status":        3,
		"resolve_time": time.Now(),
	})
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
