package repository

import (
	"context"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

type LoginLogRepository struct {
	db *gorm.DB
}

func NewLoginLogRepository(db *gorm.DB) *LoginLogRepository {
	return &LoginLogRepository{db: db}
}

func (r *LoginLogRepository) Create(ctx context.Context, log *model.LoginLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *LoginLogRepository) FindByPage(ctx context.Context, tenantID int64, username, status, ip string, page, pageSize int) ([]model.LoginLog, int64, error) {
	var logs []model.LoginLog
	var total int64

	query := r.db.WithContext(ctx).Model(&model.LoginLog{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if ip != "" {
		query = query.Where("ip LIKE ?", "%"+ip+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&logs).Error
	return logs, total, err
}

func (r *LoginLogRepository) DeleteClean(ctx context.Context, days int) error {
	cutoff := time.Now().AddDate(0, 0, -days)
	return r.db.WithContext(ctx).Where("login_time < ?", cutoff).Delete(&model.LoginLog{}).Error
}
