package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ContainerLifecycleRepository 容器生命周期仓储
type ContainerLifecycleRepository struct {
	db *gorm.DB
}

func NewContainerLifecycleRepository(db *gorm.DB) *ContainerLifecycleRepository {
	return &ContainerLifecycleRepository{db: db}
}

// List 分页查询容器生命周期记录
func (r *ContainerLifecycleRepository) List(ctx context.Context, tenantID int64, query *model.ContainerLifecycleQuery) ([]model.ContainerLifecycle, int64, error) {
	var list []model.ContainerLifecycle
	var total int64

	db := r.db.WithContext(ctx).Model(&model.ContainerLifecycle{}).Where("tenant_id = ?", tenantID)

	if query != nil {
		if query.ContainerID > 0 {
			db = db.Where("container_id = ?", query.ContainerID)
		}
		if query.EventType != "" {
			db = db.Where("event_type = ?", query.EventType)
		}
		if query.StartDate != nil {
			db = db.Where("event_date >= ?", query.StartDate)
		}
		if query.EndDate != nil {
			db = db.Where("event_date <= ?", query.EndDate)
		}
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	page := 1
	pageSize := 20
	if query != nil {
		if query.Page > 0 {
			page = query.Page
		}
		if query.PageSize > 0 {
			pageSize = query.PageSize
		}
	}

	err = db.Order("event_date DESC, id DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// GetByID 根据ID查询
func (r *ContainerLifecycleRepository) GetByID(ctx context.Context, id uint) (*model.ContainerLifecycle, error) {
	var lifecycle model.ContainerLifecycle
	err := r.db.WithContext(ctx).First(&lifecycle, id).Error
	if err != nil {
		return nil, err
	}
	return &lifecycle, nil
}

// GetByContainerID 查询某个容器的所有生命周期记录
func (r *ContainerLifecycleRepository) GetByContainerID(ctx context.Context, containerID int64) ([]model.ContainerLifecycle, error) {
	var list []model.ContainerLifecycle
	err := r.db.WithContext(ctx).Where("container_id = ?", containerID).
		Order("event_date DESC, id DESC").Find(&list).Error
	return list, err
}

// Create 创建生命周期记录
func (r *ContainerLifecycleRepository) Create(ctx context.Context, lifecycle *model.ContainerLifecycle) error {
	return r.db.WithContext(ctx).Create(lifecycle).Error
}

// GetLatestByContainerID 获取容器最新状态记录
func (r *ContainerLifecycleRepository) GetLatestByContainerID(ctx context.Context, containerID int64) (*model.ContainerLifecycle, error) {
	var lifecycle model.ContainerLifecycle
	err := r.db.WithContext(ctx).Where("container_id = ?", containerID).
		Order("event_date DESC, id DESC").First(&lifecycle).Error
	if err != nil {
		return nil, err
	}
	return &lifecycle, nil
}

// ContainerMaintenanceRepository 容器维修记录仓储
type ContainerMaintenanceRepository struct {
	db *gorm.DB
}

func NewContainerMaintenanceRepository(db *gorm.DB) *ContainerMaintenanceRepository {
	return &ContainerMaintenanceRepository{db: db}
}

// List 分页查询容器维修记录
func (r *ContainerMaintenanceRepository) List(ctx context.Context, tenantID int64, query *model.ContainerMaintenanceQuery) ([]model.ContainerMaintenance, int64, error) {
	var list []model.ContainerMaintenance
	var total int64

	db := r.db.WithContext(ctx).Model(&model.ContainerMaintenance{}).Where("tenant_id = ?", tenantID)

	if query != nil {
		if query.ContainerID > 0 {
			db = db.Where("container_id = ?", query.ContainerID)
		}
		if query.ContainerCode != "" {
			db = db.Where("container_code LIKE ?", "%"+query.ContainerCode+"%")
		}
		if query.MaintenanceType != "" {
			db = db.Where("maintenance_type = ?", query.MaintenanceType)
		}
		if query.Status != "" {
			db = db.Where("status = ?", query.Status)
		}
		if query.StartDate != nil {
			db = db.Where("maintenance_date >= ?", query.StartDate)
		}
		if query.EndDate != nil {
			db = db.Where("maintenance_date <= ?", query.EndDate)
		}
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	page := 1
	pageSize := 20
	if query != nil {
		if query.Page > 0 {
			page = query.Page
		}
		if query.PageSize > 0 {
			pageSize = query.PageSize
		}
	}

	err = db.Order("maintenance_date DESC, id DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// GetByID 根据ID查询
func (r *ContainerMaintenanceRepository) GetByID(ctx context.Context, id uint) (*model.ContainerMaintenance, error) {
	var maintenance model.ContainerMaintenance
	err := r.db.WithContext(ctx).First(&maintenance, id).Error
	if err != nil {
		return nil, err
	}
	return &maintenance, nil
}

// GetByContainerID 查询某个容器的所有维修记录
func (r *ContainerMaintenanceRepository) GetByContainerID(ctx context.Context, containerID int64) ([]model.ContainerMaintenance, error) {
	var list []model.ContainerMaintenance
	err := r.db.WithContext(ctx).Where("container_id = ?", containerID).
		Order("maintenance_date DESC, id DESC").Find(&list).Error
	return list, err
}

// Create 创建维修记录
func (r *ContainerMaintenanceRepository) Create(ctx context.Context, maintenance *model.ContainerMaintenance) error {
	return r.db.WithContext(ctx).Create(maintenance).Error
}

// Update 更新维修记录
func (r *ContainerMaintenanceRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ContainerMaintenance{}).Where("id = ?", id).Updates(updates).Error
}

// GetByIDOnly 根据ID查询（不区分租户，用于内部校验）
func (r *ContainerMaintenanceRepository) GetByIDOnly(ctx context.Context, id uint) (*model.ContainerMaintenance, error) {
	var maintenance model.ContainerMaintenance
	err := r.db.WithContext(ctx).First(&maintenance, id).Error
	if err != nil {
		return nil, err
	}
	return &maintenance, nil
}