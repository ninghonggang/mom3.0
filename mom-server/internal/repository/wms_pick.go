package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// WMSPickJobRepository 拣货作业仓储
type WMSPickJobRepository struct {
	db *gorm.DB
}

func NewWMSPickJobRepository(db *gorm.DB) *WMSPickJobRepository {
	return &WMSPickJobRepository{db: db}
}

func (r *WMSPickJobRepository) List(ctx context.Context, tenantID int64, query string) ([]model.WMSPickJob, int64, error) {
	var list []model.WMSPickJob
	var total int64

	q := r.db.WithContext(ctx).Model(&model.WMSPickJob{})
	if tenantID > 0 {
		q = q.Where("tenant_id = ?", tenantID)
	}
	if query != "" {
		q = q.Where("pick_no LIKE ? OR source_no LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	err := q.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = q.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *WMSPickJobRepository) GetByID(ctx context.Context, id uint) (*model.WMSPickJob, error) {
	var job model.WMSPickJob
	err := r.db.WithContext(ctx).First(&job, id).Error
	return &job, err
}

func (r *WMSPickJobRepository) GetByPickNo(ctx context.Context, pickNo string) (*model.WMSPickJob, error) {
	var job model.WMSPickJob
	err := r.db.WithContext(ctx).Where("pick_no = ?", pickNo).First(&job).Error
	return &job, err
}

func (r *WMSPickJobRepository) Create(ctx context.Context, job *model.WMSPickJob) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *WMSPickJobRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WMSPickJob{}).Where("id = ?", id).Updates(updates).Error
}

func (r *WMSPickJobRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.WMSPickJob{}, id).Error
}

// WMSPickRecordRepository 拣货明细仓储
type WMSPickRecordRepository struct {
	db *gorm.DB
}

func NewWMSPickRecordRepository(db *gorm.DB) *WMSPickRecordRepository {
	return &WMSPickRecordRepository{db: db}
}

func (r *WMSPickRecordRepository) ListByPickJobID(ctx context.Context, pickJobID int64) ([]model.WMSPickRecord, error) {
	var list []model.WMSPickRecord
	err := r.db.WithContext(ctx).Where("pick_job_id = ?", pickJobID).Order("id ASC").Find(&list).Error
	return list, err
}

func (r *WMSPickRecordRepository) GetByID(ctx context.Context, id uint) (*model.WMSPickRecord, error) {
	var record model.WMSPickRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	return &record, err
}

func (r *WMSPickRecordRepository) Create(ctx context.Context, record *model.WMSPickRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *WMSPickRecordRepository) CreateBatch(ctx context.Context, records []model.WMSPickRecord) error {
	return r.db.WithContext(ctx).Create(&records).Error
}

func (r *WMSPickRecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WMSPickRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *WMSPickRecordRepository) UpdateByPickJobID(ctx context.Context, pickJobID int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WMSPickRecord{}).Where("pick_job_id = ?", pickJobID).Updates(updates).Error
}
