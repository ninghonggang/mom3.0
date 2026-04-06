package service

import (
	"context"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type DCService struct {
	dataPointRepo *repository.DCDataPointRepository
	scanLogRepo   *repository.DCScanLogRepository
	collectRepo   *repository.DCCollectRecordRepository
}

func NewDCService(dp *repository.DCDataPointRepository, sl *repository.DCScanLogRepository, cr *repository.DCCollectRecordRepository) *DCService {
	return &DCService{
		dataPointRepo: dp,
		scanLogRepo:   sl,
		collectRepo:   cr,
	}
}

// DataPoint CRUD
func (s *DCService) ListDataPoints(ctx context.Context, tenantID int64, query string) ([]model.DCDataPoint, int64, error) {
	return s.dataPointRepo.List(ctx, tenantID, query)
}

func (s *DCService) GetDataPoint(ctx context.Context, id uint) (*model.DCDataPoint, error) {
	return s.dataPointRepo.GetByID(ctx, id)
}

func (s *DCService) CreateDataPoint(ctx context.Context, item *model.DCDataPoint) error {
	if item.TenantID == 0 {
		item.TenantID = 1
	}
	return s.dataPointRepo.Create(ctx, item)
}

func (s *DCService) UpdateDataPoint(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.dataPointRepo.Update(ctx, id, updates)
}

func (s *DCService) DeleteDataPoint(ctx context.Context, id uint) error {
	return s.dataPointRepo.Delete(ctx, id)
}

// ScanLog
func (s *DCService) ListScanLogs(ctx context.Context, tenantID int64, req *repository.DCScanLogQuery) ([]model.DCScanLog, int64, error) {
	return s.scanLogRepo.List(ctx, tenantID, req)
}

func (s *DCService) CreateScanLog(ctx context.Context, log *model.DCScanLog) error {
	if log.TenantID == 0 {
		log.TenantID = 1
	}
	return s.scanLogRepo.Create(ctx, log)
}

// CollectRecord
func (s *DCService) ListCollectRecords(ctx context.Context, tenantID int64, pointID int64, startTime, endTime string, limit, offset int) ([]model.DCCollectRecord, int64, error) {
	return s.collectRepo.List(ctx, tenantID, pointID, startTime, endTime, limit, offset)
}
