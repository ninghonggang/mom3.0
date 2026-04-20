package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ToolRepository 器具仓储
type ToolRepository struct {
	db *gorm.DB
}

func NewToolRepository(db *gorm.DB) *ToolRepository {
	return &ToolRepository{db: db}
}

func (r *ToolRepository) List(ctx context.Context, tenantID int64, query *model.ToolQuery) ([]model.Tool, int64, error) {
	var list []model.Tool
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Tool{}).Where("tenant_id = ?", tenantID)
	if query != nil {
		if query.Query != "" {
			db = db.Where("tool_code LIKE ? OR tool_name LIKE ?", "%"+query.Query+"%", "%"+query.Query+"%")
		}
		if query.ToolType != "" {
			db = db.Where("tool_type = ?", query.ToolType)
		}
		if query.Status != "" {
			db = db.Where("status = ?", query.Status)
		}
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	page := 1
	pageSize := 20
	if query != nil && query.Page > 0 {
		page = query.Page
	}
	if query != nil && query.PageSize > 0 {
		pageSize = query.PageSize
	}
	err = db.Order("id DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

func (r *ToolRepository) GetByID(ctx context.Context, id int64) (*model.Tool, error) {
	var tool model.Tool
	err := r.db.WithContext(ctx).First(&tool, id).Error
	return &tool, err
}

func (r *ToolRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.Tool, error) {
	var tool model.Tool
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND tool_code = ?", tenantID, code).First(&tool).Error
	return &tool, err
}

func (r *ToolRepository) Create(ctx context.Context, tool *model.Tool) error {
	return r.db.WithContext(ctx).Create(tool).Error
}

func (r *ToolRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Tool{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ToolRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Tool{}, id).Error
}

// ToolContainerRepository 容器仓储
type ToolContainerRepository struct {
	db *gorm.DB
}

func NewToolContainerRepository(db *gorm.DB) *ToolContainerRepository {
	return &ToolContainerRepository{db: db}
}

func (r *ToolContainerRepository) List(ctx context.Context, tenantID int64, query *model.ToolContainerQuery) ([]model.ToolContainer, int64, error) {
	var list []model.ToolContainer
	var total int64

	db := r.db.WithContext(ctx).Model(&model.ToolContainer{}).Where("tenant_id = ?", tenantID)
	if query != nil {
		if query.Query != "" {
			db = db.Where("container_code LIKE ? OR container_name LIKE ?", "%"+query.Query+"%", "%"+query.Query+"%")
		}
		if query.ContainerType != "" {
			db = db.Where("container_type = ?", query.ContainerType)
		}
		if query.Status != "" {
			db = db.Where("status = ?", query.Status)
		}
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	page := 1
	pageSize := 20
	if query != nil && query.Page > 0 {
		page = query.Page
	}
	if query != nil && query.PageSize > 0 {
		pageSize = query.PageSize
	}
	err = db.Order("id DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

func (r *ToolContainerRepository) GetByID(ctx context.Context, id int64) (*model.ToolContainer, error) {
	var container model.ToolContainer
	err := r.db.WithContext(ctx).First(&container, id).Error
	return &container, err
}

func (r *ToolContainerRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.ToolContainer, error) {
	var container model.ToolContainer
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND container_code = ?", tenantID, code).First(&container).Error
	return &container, err
}

func (r *ToolContainerRepository) Create(ctx context.Context, container *model.ToolContainer) error {
	return r.db.WithContext(ctx).Create(container).Error
}

func (r *ToolContainerRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ToolContainer{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ToolContainerRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.ToolContainer{}, id).Error
}

// ToolContainerBindingRepository 器具容器绑定仓储
type ToolContainerBindingRepository struct {
	db *gorm.DB
}

func NewToolContainerBindingRepository(db *gorm.DB) *ToolContainerBindingRepository {
	return &ToolContainerBindingRepository{db: db}
}

func (r *ToolContainerBindingRepository) List(ctx context.Context, tenantID int64, query *model.ToolContainerBindingQuery) ([]model.ToolContainerBinding, int64, error) {
	var list []model.ToolContainerBinding
	var total int64

	db := r.db.WithContext(ctx).Model(&model.ToolContainerBinding{}).Where("tenant_id = ?", tenantID)
	if query != nil {
		if query.Query != "" {
			db = db.Where("tool_code LIKE ? OR container_code LIKE ?", "%"+query.Query+"%", "%"+query.Query+"%")
		}
		if query.ToolID > 0 {
			db = db.Where("tool_id = ?", query.ToolID)
		}
		if query.ContainerID > 0 {
			db = db.Where("container_id = ?", query.ContainerID)
		}
		if query.Status != "" {
			db = db.Where("status = ?", query.Status)
		}
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	page := 1
	pageSize := 20
	if query != nil && query.Page > 0 {
		page = query.Page
	}
	if query != nil && query.PageSize > 0 {
		pageSize = query.PageSize
	}
	err = db.Order("id DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

func (r *ToolContainerBindingRepository) GetByID(ctx context.Context, id int64) (*model.ToolContainerBinding, error) {
	var binding model.ToolContainerBinding
	err := r.db.WithContext(ctx).First(&binding, id).Error
	return &binding, err
}

func (r *ToolContainerBindingRepository) GetActiveBinding(ctx context.Context, tenantID int64, toolID int64) (*model.ToolContainerBinding, error) {
	var binding model.ToolContainerBinding
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND tool_id = ? AND status = ?", tenantID, toolID, "BOUND").First(&binding).Error
	return &binding, err
}

func (r *ToolContainerBindingRepository) GetActiveBindingByContainer(ctx context.Context, tenantID int64, containerID int64) (*model.ToolContainerBinding, error) {
	var binding model.ToolContainerBinding
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND container_id = ? AND status = ?", tenantID, containerID, "BOUND").First(&binding).Error
	return &binding, err
}

func (r *ToolContainerBindingRepository) GetByToolID(ctx context.Context, tenantID int64, toolID int64) ([]model.ToolContainerBinding, error) {
	var list []model.ToolContainerBinding
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND tool_id = ?", tenantID, toolID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *ToolContainerBindingRepository) GetByContainerID(ctx context.Context, tenantID int64, containerID int64) ([]model.ToolContainerBinding, error) {
	var list []model.ToolContainerBinding
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND container_id = ?", tenantID, containerID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *ToolContainerBindingRepository) Create(ctx context.Context, binding *model.ToolContainerBinding) error {
	return r.db.WithContext(ctx).Create(binding).Error
}

func (r *ToolContainerBindingRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ToolContainerBinding{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ToolContainerBindingRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.ToolContainerBinding{}, id).Error
}
