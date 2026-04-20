package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ScpQadSyncLogRepository QAD同步日志仓库
type ScpQadSyncLogRepository struct {
	db *gorm.DB
}

func NewScpQadSyncLogRepository(db *gorm.DB) *ScpQadSyncLogRepository {
	return &ScpQadSyncLogRepository{db: db}
}

// Create 创建同步日志
func (r *ScpQadSyncLogRepository) Create(ctx context.Context, log *model.ScpQadSyncLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

// GetByID 根据ID查询
func (r *ScpQadSyncLogRepository) GetByID(ctx context.Context, id uint64) (*model.ScpQadSyncLog, error) {
	var log model.ScpQadSyncLog
	err := r.db.WithContext(ctx).First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// Update 更新同步日志
func (r *ScpQadSyncLogRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ScpQadSyncLog{}).Where("id = ?", id).Updates(updates).Error
}

// ListByDocNo 根据单据号查询日志
func (r *ScpQadSyncLogRepository) ListByDocNo(ctx context.Context, docNo string) ([]model.ScpQadSyncLog, error) {
	var list []model.ScpQadSyncLog
	err := r.db.WithContext(ctx).
		Where("qad_doc_no = ? OR mom_doc_no = ?", docNo, docNo).
		Order("created_at DESC").
		Find(&list).Error
	return list, err
}

// List 分页查询同步日志列表
func (r *ScpQadSyncLogRepository) List(ctx context.Context, q *model.QadSyncLogQuery) ([]model.ScpQadSyncLog, int64, error) {
	var list []model.ScpQadSyncLog
	var total int64

	query := r.db.WithContext(ctx).Model(&model.ScpQadSyncLog{})

	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.SyncType != "" {
		query = query.Where("sync_type = ?", q.SyncType)
	}
	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}
	if q.DocNo != "" {
		query = query.Where("qad_doc_no = ? OR mom_doc_no = ?", q.DocNo, q.DocNo)
	}

	query.Count(&total)

	page := q.Page
	if page < 1 {
		page = 1
	}
	pageSize := q.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&list).Error
	return list, total, err
}
