package repository

import (
	"context"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

// EquipmentDowntimeRepository 设备停机仓储
type EquipmentDowntimeRepository struct {
	db *gorm.DB
}

// NewEquipmentDowntimeRepository 创建设备停机仓储
func NewEquipmentDowntimeRepository(db *gorm.DB) *EquipmentDowntimeRepository {
	return &EquipmentDowntimeRepository{db: db}
}

// List 获取设备停机列表
func (r *EquipmentDowntimeRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.EquipmentDowntime, int64, error) {
	var list []model.EquipmentDowntime
	var total int64

	q := r.db.WithContext(ctx).Model(&model.EquipmentDowntime{})
	if tenantID > 0 {
		q = q.Where("tenant_id = ?", tenantID)
	}
	if query != nil {
		if equipID, ok := query["equipment_id"].(int64); ok && equipID > 0 {
			q = q.Where("equipment_id = ?", equipID)
		}
		if equipCode, ok := query["equipment_code"].(string); ok && equipCode != "" {
			q = q.Where("equipment_code LIKE ?", "%"+equipCode+"%")
		}
		if dtType, ok := query["downtime_type"].(string); ok && dtType != "" {
			q = q.Where("downtime_type = ?", dtType)
		}
		if status, ok := query["status"].(string); ok && status != "" {
			q = q.Where("status = ?", status)
		}
		if startTime, ok := query["start_time"].(string); ok && startTime != "" {
			q = q.Where("start_time >= ?", startTime)
		}
		if endTime, ok := query["end_time"].(string); ok && endTime != "" {
			q = q.Where("start_time <= ?", endTime)
		}
	}

	err := q.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	page := 1
	pageSize := 20
	if query != nil {
		if p, ok := query["page"].(int); ok && p > 0 {
			page = p
		}
		if ps, ok := query["page_size"].(int); ok && ps > 0 {
			pageSize = ps
		}
	}
	offset := (page - 1) * pageSize

	err = q.Order("id DESC").Offset(offset).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// GetByID 根据ID获取设备停机记录
func (r *EquipmentDowntimeRepository) GetByID(ctx context.Context, id int64) (*model.EquipmentDowntime, error) {
	var downtime model.EquipmentDowntime
	err := r.db.WithContext(ctx).First(&downtime, id).Error
	return &downtime, err
}

// Create 创建设备停机记录
func (r *EquipmentDowntimeRepository) Create(ctx context.Context, d *model.EquipmentDowntime) error {
	return r.db.WithContext(ctx).Create(d).Error
}

// Update 更新设备停机记录
func (r *EquipmentDowntimeRepository) Update(ctx context.Context, d *model.EquipmentDowntime) error {
	return r.db.WithContext(ctx).Save(d).Error
}

// Delete 删除设备停机记录
func (r *EquipmentDowntimeRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.EquipmentDowntime{}, id).Error
}

// UpdateStatus 更新停机状态
func (r *EquipmentDowntimeRepository) UpdateStatus(ctx context.Context, id int64, status string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == "CLOSED" {
		now := time.Now()
		updates["end_time"] = &now
	}
	return r.db.WithContext(ctx).Model(&model.EquipmentDowntime{}).Where("id = ?", id).Updates(updates).Error
}
