package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// WmsLabelTemplateRepository 标签模板仓储层
type WmsLabelTemplateRepository struct {
	db *gorm.DB
}

// NewWmsLabelTemplateRepository 创建标签模板仓储实例
func NewWmsLabelTemplateRepository(db *gorm.DB) *WmsLabelTemplateRepository {
	return &WmsLabelTemplateRepository{db: db}
}

// List 获取标签模板列表（分页）
func (r *WmsLabelTemplateRepository) List(ctx context.Context, query *model.WmsLabelTemplateQueryVO) ([]model.WmsLabelTemplate, int64, error) {
	var list []model.WmsLabelTemplate
	var total int64

	queryDB := r.db.WithContext(ctx).Model(&model.WmsLabelTemplate{})

	if query != nil && query.Keyword != "" {
		queryDB = queryDB.Where("template_code LIKE ? OR template_name LIKE ?",
			"%"+query.Keyword+"%", "%"+query.Keyword+"%")
	}
	if query != nil && query.TemplateType != "" {
		queryDB = queryDB.Where("template_type = ?", query.TemplateType)
	}
	if query != nil && query.Status != "" {
		queryDB = queryDB.Where("status = ?", query.Status)
	}

	err := queryDB.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页
	if query != nil && query.Page > 0 && query.PageSize > 0 {
		offset := (query.Page - 1) * query.PageSize
		queryDB = queryDB.Offset(offset).Limit(query.PageSize)
	}

	err = queryDB.Order("id DESC").Find(&list).Error
	return list, total, err
}

// GetByID 根据ID获取标签模板
func (r *WmsLabelTemplateRepository) GetByID(ctx context.Context, id uint64) (*model.WmsLabelTemplate, error) {
	var template model.WmsLabelTemplate
	err := r.db.WithContext(ctx).First(&template, id).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// GetByCode 根据编码获取标签模板
func (r *WmsLabelTemplateRepository) GetByCode(ctx context.Context, templateCode string) (*model.WmsLabelTemplate, error) {
	var template model.WmsLabelTemplate
	err := r.db.WithContext(ctx).Where("template_code = ?", templateCode).First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// Create 创建标签模板
func (r *WmsLabelTemplateRepository) Create(ctx context.Context, template *model.WmsLabelTemplate) error {
	return r.db.WithContext(ctx).Create(template).Error
}

// Update 更新标签模板
func (r *WmsLabelTemplateRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.WmsLabelTemplate{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除标签模板
func (r *WmsLabelTemplateRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.WmsLabelTemplate{}, id).Error
}