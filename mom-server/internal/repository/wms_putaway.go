package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type PutawayJobRepository struct {
	db *gorm.DB
}

func NewPutawayJobRepository(db *gorm.DB) *PutawayJobRepository {
	return &PutawayJobRepository{db: db}
}

func (r *PutawayJobRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.WMSPutawayJob, int64, error) {
	var list []model.WMSPutawayJob
	var total int64

	q := r.db.WithContext(ctx).Model(&model.WMSPutawayJob{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"].(string); ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if sourceType, ok := query["source_type"].(string); ok && sourceType != "" {
		q = q.Where("source_type = ?", sourceType)
	}
	if sourceNo, ok := query["source_no"].(string); ok && sourceNo != "" {
		q = q.Where("source_no LIKE ?", "%"+sourceNo+"%")
	}
	if warehouseID, ok := query["warehouse_id"].(int64); ok && warehouseID > 0 {
		q = q.Where("warehouse_id = ?", warehouseID)
	}

	err := q.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}

	offset := (page - 1) * limit
	err = q.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func (r *PutawayJobRepository) GetByID(ctx context.Context, id uint) (*model.WMSPutawayJob, error) {
	var job model.WMSPutawayJob
	err := r.db.WithContext(ctx).First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *PutawayJobRepository) GetByPutawayNo(ctx context.Context, putawayNo string) (*model.WMSPutawayJob, error) {
	var job model.WMSPutawayJob
	err := r.db.WithContext(ctx).Where("putaway_no = ?", putawayNo).First(&job).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *PutawayJobRepository) Create(ctx context.Context, job *model.WMSPutawayJob) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *PutawayJobRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WMSPutawayJob{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PutawayJobRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除明细
		if err := tx.Where("putaway_job_id = ?", id).Delete(&model.WMSPutawayRecord{}).Error; err != nil {
			return err
		}
		// 再删除主单
		return tx.Delete(&model.WMSPutawayJob{}, id).Error
	})
}

// PutawayRecordRepository 上架明细仓库
type PutawayRecordRepository struct {
	db *gorm.DB
}

func NewPutawayRecordRepository(db *gorm.DB) *PutawayRecordRepository {
	return &PutawayRecordRepository{db: db}
}

func (r *PutawayRecordRepository) ListByJobID(ctx context.Context, jobID uint) ([]model.WMSPutawayRecord, error) {
	var list []model.WMSPutawayRecord
	err := r.db.WithContext(ctx).Where("putaway_job_id = ?", jobID).Order("id ASC").Find(&list).Error
	return list, err
}

func (r *PutawayRecordRepository) GetByID(ctx context.Context, id uint) (*model.WMSPutawayRecord, error) {
	var record model.WMSPutawayRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *PutawayRecordRepository) Create(ctx context.Context, record *model.WMSPutawayRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *PutawayRecordRepository) CreateBatch(ctx context.Context, records []model.WMSPutawayRecord) error {
	return r.db.WithContext(ctx).Create(&records).Error
}

func (r *PutawayRecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WMSPutawayRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PutawayRecordRepository) UpdateByJobID(ctx context.Context, jobID uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WMSPutawayRecord{}).Where("putaway_job_id = ?", jobID).Updates(updates).Error
}

func (r *PutawayRecordRepository) DeleteByJobID(ctx context.Context, jobID uint) error {
	return r.db.WithContext(ctx).Where("putaway_job_id = ?", jobID).Delete(&model.WMSPutawayRecord{}).Error
}
