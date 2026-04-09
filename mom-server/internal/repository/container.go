package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type ContainerRepository struct {
	db *gorm.DB
}

func NewContainerRepository(db *gorm.DB) *ContainerRepository {
	return &ContainerRepository{db: db}
}

// List 分页查询器具
func (r *ContainerRepository) List(ctx context.Context, tenantID int64, params model.ContainerQueryParams) ([]model.ContainerMaster, int64, error) {
	var list []model.ContainerMaster
	var total int64

	query := r.db.WithContext(ctx).Model(&model.ContainerMaster{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}
	if params.ContainerType != "" {
		query = query.Where("container_type = ?", params.ContainerType)
	}
	if params.CustomerID > 0 {
		query = query.Where("customer_id = ?", params.CustomerID)
	}
	if params.Keyword != "" {
		query = query.Where("container_code LIKE ? OR container_name LIKE ? OR barcode LIKE ?",
			"%"+params.Keyword+"%", "%"+params.Keyword+"%", "%"+params.Keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	page := params.Page
	if page < 1 {
		page = 1
	}
	pageSize := params.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	err = query.Order("id DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// GetByID 根据ID查询
func (r *ContainerRepository) GetByID(ctx context.Context, id uint) (*model.ContainerMaster, error) {
	var container model.ContainerMaster
	err := r.db.WithContext(ctx).First(&container, id).Error
	if err != nil {
		return nil, err
	}
	return &container, nil
}

// GetByCode 根据器具编号查询
func (r *ContainerRepository) GetByCode(ctx context.Context, code string) (*model.ContainerMaster, error) {
	var container model.ContainerMaster
	err := r.db.WithContext(ctx).Where("container_code = ?", code).First(&container).Error
	if err != nil {
		return nil, err
	}
	return &container, nil
}

// GetByBarcode 根据条码查询
func (r *ContainerRepository) GetByBarcode(ctx context.Context, barcode string) (*model.ContainerMaster, error) {
	var container model.ContainerMaster
	err := r.db.WithContext(ctx).Where("barcode = ?", barcode).First(&container).Error
	if err != nil {
		return nil, err
	}
	return &container, nil
}

// Create 创建器具
func (r *ContainerRepository) Create(ctx context.Context, container *model.ContainerMaster) error {
	return r.db.WithContext(ctx).Create(container).Error
}

// Update 更新器具
func (r *ContainerRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ContainerMaster{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除器具
func (r *ContainerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ContainerMaster{}, id).Error
}

// HasOpenMovements 检查是否有未完结的流转记录
func (r *ContainerRepository) HasOpenMovements(ctx context.Context, containerID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.ContainerMovement{}).
		Where("container_id = ? AND movement_type IN ?", containerID, []string{"OUT", "TRANSFER"}).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateMovement 创建流转记录
func (r *ContainerRepository) CreateMovement(ctx context.Context, movement *model.ContainerMovement) error {
	return r.db.WithContext(ctx).Create(movement).Error
}

// ListMovements 查询流转记录
func (r *ContainerRepository) ListMovements(ctx context.Context, containerID uint, page, pageSize int) ([]model.ContainerMovement, int64, error) {
	var list []model.ContainerMovement
	var total int64

	query := r.db.WithContext(ctx).Model(&model.ContainerMovement{}).Where("container_id = ?", containerID)
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	err = query.Order("movement_time DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// IncrementTrips 增加使用次数
func (r *ContainerRepository) IncrementTrips(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.ContainerMaster{}).Where("id = ?", id).
		Update("total_trips", gorm.Expr("total_trips + ?", 1)).Error
}

// UpdateStatus 更新状态
func (r *ContainerRepository) UpdateStatus(ctx context.Context, id uint, status model.ContainerStatus) error {
	return r.db.WithContext(ctx).Model(&model.ContainerMaster{}).Where("id = ?", id).
		Update("status", status).Error
}

// UpdateLocation 更新位置
func (r *ContainerRepository) UpdateLocation(ctx context.Context, id uint, location string) error {
	return r.db.WithContext(ctx).Model(&model.ContainerMaster{}).Where("id = ?", id).
		Update("current_location", location).Error
}

// UpdateLastCleanDate 更新清洁日期
func (r *ContainerRepository) UpdateLastCleanDate(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.ContainerMaster{}).Where("id = ?", id).
		Update("last_clean_date", gorm.Expr("NOW()")).Error
}

// Transaction 事务处理
func (r *ContainerRepository) Transaction(ctx context.Context, fn func(*ContainerRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := &ContainerRepository{db: tx}
		return fn(repo)
	})
}

// GetDB 获取数据库实例
func (r *ContainerRepository) GetDB() *gorm.DB {
	return r.db
}

// ContainerMovementRepository 流转记录仓库
type ContainerMovementRepository struct {
	db *gorm.DB
}

func NewContainerMovementRepository(db *gorm.DB) *ContainerMovementRepository {
	return &ContainerMovementRepository{db: db}
}

// Create 创建流转记录
func (r *ContainerMovementRepository) Create(ctx context.Context, movement *model.ContainerMovement) error {
	return r.db.WithContext(ctx).Create(movement).Error
}

// ListByContainerID 查询器具流转记录
func (r *ContainerMovementRepository) ListByContainerID(ctx context.Context, containerID uint, page, pageSize int) ([]model.ContainerMovement, int64, error) {
	var list []model.ContainerMovement
	var total int64

	query := r.db.WithContext(ctx).Model(&model.ContainerMovement{}).Where("container_id = ?", containerID)
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	err = query.Order("movement_time DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// GetByID 根据ID查询
func (r *ContainerMovementRepository) GetByID(ctx context.Context, id uint) (*model.ContainerMovement, error) {
	var movement model.ContainerMovement
	err := r.db.WithContext(ctx).First(&movement, id).Error
	if err != nil {
		return nil, err
	}
	return &movement, nil
}

func parseUint(s string) uint {
	var v uint
	fmt.Sscanf(s, "%d", &v)
	return v
}
