package repository

import (
	"context"

	"gorm.io/gorm"

	"mom-server/internal/model"
)

// EamRepairJobRepository 维修工单仓储
type EamRepairJobRepository struct {
	db *gorm.DB
}

func NewEamRepairJobRepository(db *gorm.DB) *EamRepairJobRepository {
	return &EamRepairJobRepository{db: db}
}

func (r *EamRepairJobRepository) Create(ctx context.Context, m *model.EamRepairJob) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EamRepairJobRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EamRepairJob{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EamRepairJobRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.EamRepairJob{}, id).Error
}

func (r *EamRepairJobRepository) GetByID(ctx context.Context, id int64) (*model.EamRepairJob, error) {
	var m model.EamRepairJob
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *EamRepairJobRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.EamRepairJob, int64, error) {
	var list []model.EamRepairJob
	var total int64
	query := r.db.WithContext(ctx).Model(&model.EamRepairJob{})
	if tenantID, ok := filters["tenant_id"].(int64); ok && tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if jobCode, ok := filters["job_code"].(string); ok && jobCode != "" {
		query = query.Where("job_code LIKE ?", "%"+jobCode+"%")
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if equipmentID, ok := filters["equipment_id"].(int64); ok && equipmentID > 0 {
		query = query.Where("equipment_id = ?", equipmentID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}
	if err := query.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// EamRepairFlowRepository 维修流程仓储
type EamRepairFlowRepository struct {
	db *gorm.DB
}

func NewEamRepairFlowRepository(db *gorm.DB) *EamRepairFlowRepository {
	return &EamRepairFlowRepository{db: db}
}

func (r *EamRepairFlowRepository) Create(ctx context.Context, m *model.EamRepairFlow) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EamRepairFlowRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EamRepairFlow{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EamRepairFlowRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.EamRepairFlow{}, id).Error
}

func (r *EamRepairFlowRepository) GetByID(ctx context.Context, id int64) (*model.EamRepairFlow, error) {
	var m model.EamRepairFlow
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *EamRepairFlowRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.EamRepairFlow, int64, error) {
	var list []model.EamRepairFlow
	var total int64
	query := r.db.WithContext(ctx).Model(&model.EamRepairFlow{})
	if tenantID, ok := filters["tenant_id"].(int64); ok && tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if flowCode, ok := filters["flow_code"].(string); ok && flowCode != "" {
		query = query.Where("flow_code LIKE ?", "%"+flowCode+"%")
	}
	if flowName, ok := filters["flow_name"].(string); ok && flowName != "" {
		query = query.Where("flow_name LIKE ?", "%"+flowName+"%")
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}
	if err := query.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// EamRepairStdRepository 维修标准仓储
type EamRepairStdRepository struct {
	db *gorm.DB
}

func NewEamRepairStdRepository(db *gorm.DB) *EamRepairStdRepository {
	return &EamRepairStdRepository{db: db}
}

func (r *EamRepairStdRepository) Create(ctx context.Context, m *model.EamRepairStd) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *EamRepairStdRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.EamRepairStd{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EamRepairStdRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.EamRepairStd{}, id).Error
}

func (r *EamRepairStdRepository) GetByID(ctx context.Context, id int64) (*model.EamRepairStd, error) {
	var m model.EamRepairStd
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *EamRepairStdRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.EamRepairStd, int64, error) {
	var list []model.EamRepairStd
	var total int64
	query := r.db.WithContext(ctx).Model(&model.EamRepairStd{})
	if tenantID, ok := filters["tenant_id"].(int64); ok && tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if stdCode, ok := filters["std_code"].(string); ok && stdCode != "" {
		query = query.Where("std_code LIKE ?", "%"+stdCode+"%")
	}
	if stdName, ok := filters["std_name"].(string); ok && stdName != "" {
		query = query.Where("std_name LIKE ?", "%"+stdName+"%")
	}
	if faultType, ok := filters["fault_type"].(string); ok && faultType != "" {
		query = query.Where("fault_type = ?", faultType)
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}
	if err := query.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
