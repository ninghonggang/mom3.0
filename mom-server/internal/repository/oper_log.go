package repository

import (
	"context"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

type OperLogRepository struct {
	db *gorm.DB
}

func NewOperLogRepository(db *gorm.DB) *OperLogRepository {
	return &OperLogRepository{db: db}
}

func (r *OperLogRepository) Create(ctx context.Context, log *model.OperLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *OperLogRepository) FindByPage(ctx context.Context, tenantID int64, title, operName, businessType, status string, page, pageSize int) ([]model.OperLog, int64, error) {
	var logs []model.OperLog
	var total int64

	query := r.db.WithContext(ctx).Model(&model.OperLog{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if operName != "" {
		query = query.Where("oper_name LIKE ?", "%"+operName+"%")
	}
	if businessType != "" {
		query = query.Where("business_type = ?", businessType)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&logs).Error
	return logs, total, err
}

func (r *OperLogRepository) DeleteClean(ctx context.Context, days int) error {
	cutoff := time.Now().AddDate(0, 0, -days)
	return r.db.WithContext(ctx).Where("oper_time < ?", cutoff).Delete(&model.OperLog{}).Error
}
