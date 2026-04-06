package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type DCDataPointRepository struct {
	db *gorm.DB
}

func NewDCDataPointRepository(db *gorm.DB) *DCDataPointRepository {
	return &DCDataPointRepository{db: db}
}

func (r *DCDataPointRepository) List(ctx context.Context, tenantID int64, query string) ([]model.DCDataPoint, int64, error) {
	var list []model.DCDataPoint
	var total int64
	q := r.db.WithContext(ctx).Model(&model.DCDataPoint{}).Where("tenant_id = ?", tenantID)
	if query != "" {
		q = q.Where("point_code LIKE ? OR point_name LIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *DCDataPointRepository) GetByID(ctx context.Context, id uint) (*model.DCDataPoint, error) {
	var item model.DCDataPoint
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *DCDataPointRepository) Create(ctx context.Context, item *model.DCDataPoint) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *DCDataPointRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DCDataPoint{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DCDataPointRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.DCDataPoint{}).Error
}

type DCScanLogRepository struct {
	db *gorm.DB
}

func NewDCScanLogRepository(db *gorm.DB) *DCScanLogRepository {
	return &DCScanLogRepository{db: db}
}

func (r *DCScanLogRepository) List(ctx context.Context, tenantID int64, req *DCScanLogQuery) ([]model.DCScanLog, int64, error) {
	var list []model.DCScanLog
	var total int64
	q := r.db.WithContext(ctx).Model(&model.DCScanLog{}).Where("tenant_id = ?", tenantID)
	if req.ScanType != "" {
		q = q.Where("scan_type = ?", req.ScanType)
	}
	if req.Status != "" {
		q = q.Where("status = ?", req.Status)
	}
	if req.StartTime != "" {
		q = q.Where("scan_time >= ?", req.StartTime)
	}
	if req.EndTime != "" {
		q = q.Where("scan_time <= ?", req.EndTime)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Limit(req.Limit).Offset(req.Offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *DCScanLogRepository) Create(ctx context.Context, log *model.DCScanLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *DCScanLogRepository) GetByID(ctx context.Context, id uint) (*model.DCScanLog, error) {
	var item model.DCScanLog
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

type DCScanLogQuery struct {
	ScanType  string
	Status    string
	StartTime string
	EndTime   string
	Limit     int
	Offset    int
}

type DCCollectRecordRepository struct {
	db *gorm.DB
}

func NewDCCollectRecordRepository(db *gorm.DB) *DCCollectRecordRepository {
	return &DCCollectRecordRepository{db: db}
}

func (r *DCCollectRecordRepository) List(ctx context.Context, tenantID int64, pointID int64, startTime, endTime string, limit, offset int) ([]model.DCCollectRecord, int64, error) {
	var list []model.DCCollectRecord
	var total int64
	q := r.db.WithContext(ctx).Model(&model.DCCollectRecord{}).Where("tenant_id = ?", tenantID)
	if pointID > 0 {
		q = q.Where("point_id = ?", pointID)
	}
	if startTime != "" {
		q = q.Where("collect_time >= ?", startTime)
	}
	if endTime != "" {
		q = q.Where("collect_time <= ?", endTime)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *DCCollectRecordRepository) Create(ctx context.Context, record *model.DCCollectRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}
