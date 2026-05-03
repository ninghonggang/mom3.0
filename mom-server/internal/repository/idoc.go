package repository

import (
	"context"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type IdocRepository struct {
	db *gorm.DB
}

func NewIdocRepository(db *gorm.DB) *IdocRepository {
	return &IdocRepository{db: db}
}

func (r *IdocRepository) List(ctx context.Context, tenantID int64, query *model.IdocQuery) ([]model.IdocRecord, int64, error) {
	var list []model.IdocRecord
	var total int64

	db := r.db.WithContext(ctx).Model(&model.IdocRecord{}).Where("tenant_id = ?", tenantID)

	if query.IdocType != "" {
		db = db.Where("idoc_type = ?", query.IdocType)
	}
	if query.Direction > 0 {
		db = db.Where("direction = ?", query.Direction)
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}
	if query.PartnerNo != "" {
		db = db.Where("partner_no LIKE ?", "%"+query.PartnerNo+"%")
	}
	if query.StartDate != "" {
		db = db.Where("created_at >= ?", query.StartDate)
	}
	if query.EndDate != "" {
		db = db.Where("created_at <= ?", query.EndDate)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}
	offset := (query.Page - 1) * query.PageSize

	if err := db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *IdocRepository) GetByID(ctx context.Context, id int64) (*model.IdocRecord, error) {
	var record model.IdocRecord
	if err := r.db.WithContext(ctx).First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *IdocRepository) GetByIdocNumber(ctx context.Context, idocNumber string) (*model.IdocRecord, error) {
	var record model.IdocRecord
	if err := r.db.WithContext(ctx).Where("idoc_number = ?", idocNumber).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *IdocRepository) Create(ctx context.Context, record *model.IdocRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *IdocRepository) Update(ctx context.Context, record *model.IdocRecord) error {
	return r.db.WithContext(ctx).Where("id = ?", record.ID).Updates(record).Error
}

func (r *IdocRepository) UpdateStatus(ctx context.Context, id int64, status int, errorMsg string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if errorMsg != "" {
		updates["error_message"] = errorMsg
	}
	return r.db.WithContext(ctx).Model(&model.IdocRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *IdocRepository) GetConfigByType(ctx context.Context, idocType string) (*model.IdocTypeConfig, error) {
	var config model.IdocTypeConfig
	if err := r.db.WithContext(ctx).Where("idoc_type = ? AND is_active = 1", idocType).First(&config).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *IdocRepository) ListConfigs(ctx context.Context, tenantID int64) ([]model.IdocTypeConfig, error) {
	var list []model.IdocTypeConfig
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND is_active = 1", tenantID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *IdocRepository) GenerateIdocNumber(idocType string) string {
	return idocType + "-" + time.Now().Format("20060102150405")
}