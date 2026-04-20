package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// IntegrationERPSyncLogRepository ERP同步日志仓库
type IntegrationERPSyncLogRepository struct {
	db *gorm.DB
}

func NewIntegrationERPSyncLogRepository(db *gorm.DB) *IntegrationERPSyncLogRepository {
	return &IntegrationERPSyncLogRepository{db: db}
}

// Create 创建同步日志
func (r *IntegrationERPSyncLogRepository) Create(ctx context.Context, log *model.IntegrationERPSyncLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

// GetByID 根据ID查询
func (r *IntegrationERPSyncLogRepository) GetByID(ctx context.Context, id int64) (*model.IntegrationERPSyncLog, error) {
	var log model.IntegrationERPSyncLog
	err := r.db.WithContext(ctx).First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// Update 更新同步日志
func (r *IntegrationERPSyncLogRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.IntegrationERPSyncLog{}).Where("id = ?", id).Updates(updates).Error
}

// List 查询同步日志列表
func (r *IntegrationERPSyncLogRepository) List(ctx context.Context, q *model.ERPSyncLogQuery) ([]model.IntegrationERPSyncLog, int64, error) {
	var list []model.IntegrationERPSyncLog
	var total int64

	query := r.db.WithContext(ctx).Model(&model.IntegrationERPSyncLog{})

	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.SyncType != "" {
		query = query.Where("sync_type = ?", q.SyncType)
	}
	if q.Direction != "" {
		query = query.Where("direction = ?", q.Direction)
	}
	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}
	if q.StartDate != "" {
		query = query.Where("created_at >= ?", q.StartDate)
	}
	if q.EndDate != "" {
		query = query.Where("created_at <= ?", q.EndDate+" 23:59:59")
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

// GetByERPBillNo 根据ERP单据号查询
func (r *IntegrationERPSyncLogRepository) GetByERPBillNo(ctx context.Context, syncType, erpBillNo string) (*model.IntegrationERPSyncLog, error) {
	var log model.IntegrationERPSyncLog
	err := r.db.WithContext(ctx).Where("sync_type = ? AND erp_bill_no = ?", syncType, erpBillNo).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// IntegrationERPMappingRepository ERP字段映射仓库
type IntegrationERPMappingRepository struct {
	db *gorm.DB
}

func NewIntegrationERPMappingRepository(db *gorm.DB) *IntegrationERPMappingRepository {
	return &IntegrationERPMappingRepository{db: db}
}

// Create 创建映射
func (r *IntegrationERPMappingRepository) Create(ctx context.Context, mapping *model.IntegrationERPMapping) error {
	return r.db.WithContext(ctx).Create(mapping).Error
}

// GetByID 根据ID查询
func (r *IntegrationERPMappingRepository) GetByID(ctx context.Context, id int64) (*model.IntegrationERPMapping, error) {
	var mapping model.IntegrationERPMapping
	err := r.db.WithContext(ctx).First(&mapping, id).Error
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

// Delete 删除映射
func (r *IntegrationERPMappingRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.IntegrationERPMapping{}, id).Error
}

// ListByTable 按表名查询映射
func (r *IntegrationERPMappingRepository) ListByTable(ctx context.Context, tenantID int64, erpTableName string) ([]model.IntegrationERPMapping, error) {
	var list []model.IntegrationERPMapping
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND erp_table_name = ?", tenantID, erpTableName).Find(&list).Error
	return list, err
}

// List 查询映射列表
func (r *IntegrationERPMappingRepository) List(ctx context.Context, tenantID int64) ([]model.IntegrationERPMapping, error) {
	var list []model.IntegrationERPMapping
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&list).Error
	return list, err
}
