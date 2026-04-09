package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type OEERepository struct {
	db *gorm.DB
}

func NewOEERepository(db *gorm.DB) *OEERepository {
	return &OEERepository{db: db}
}

func (r *OEERepository) List(ctx context.Context, tenantID int64, params map[string]interface{}) ([]model.OEE, int64, error) {
	var list []model.OEE
	var total int64

	query := r.db.WithContext(ctx).Model(&model.OEE{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}

	// 设备ID筛选
	if equipID, ok := params["equipment_id"]; ok && equipID.(int64) > 0 {
		query = query.Where("equipment_id = ?", equipID)
	}
	// 日期范围筛选
	if startDate, ok := params["start_date"]; ok && startDate.(string) != "" {
		query = query.Where("record_date >= ?", startDate)
	}
	if endDate, ok := params["end_date"]; ok && endDate.(string) != "" {
		query = query.Where("record_date <= ?", endDate)
	}
	// 车间ID筛选
	if workshopID, ok := params["workshop_id"]; ok && workshopID.(int64) > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("record_date DESC, id DESC").Find(&list).Error
	return list, total, err
}

func (r *OEERepository) GetByID(ctx context.Context, id int64) (*model.OEE, error) {
	var oee model.OEE
	err := r.db.WithContext(ctx).First(&oee, id).Error
	return &oee, err
}

func (r *OEERepository) Create(ctx context.Context, oee *model.OEE) error {
	return r.db.WithContext(ctx).Create(oee).Error
}

func (r *OEERepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.OEE{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OEERepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.OEE{}, id).Error
}

// GetChartData 获取图表数据
func (r *OEERepository) GetChartData(ctx context.Context, tenantID int64, params map[string]interface{}) ([]model.OEE, error) {
	var list []model.OEE
	query := r.db.WithContext(ctx).Model(&model.OEE{})

	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}

	if equipID, ok := params["equipment_id"]; ok && equipID.(int64) > 0 {
		query = query.Where("equipment_id = ?", equipID)
	}
	if startDate, ok := params["start_date"]; ok && startDate.(string) != "" {
		query = query.Where("record_date >= ?", startDate)
	}
	if endDate, ok := params["end_date"]; ok && endDate.(string) != "" {
		query = query.Where("record_date <= ?", endDate)
	}

	err := query.Order("record_date ASC").Find(&list).Error
	return list, err
}

// GetByEquipmentAndDate 获取某设备某日期的OEE记录
func (r *OEERepository) GetByEquipmentAndDate(ctx context.Context, equipmentID int64, recordDate string) (*model.OEE, error) {
	var oee model.OEE
	err := r.db.WithContext(ctx).Where("equipment_id = ? AND record_date = ?", equipmentID, recordDate).First(&oee).Error
	if err != nil {
		return nil, err
	}
	return &oee, nil
}

// OEEEventRepository OEE事件记录
type OEEEventRepository struct {
	db *gorm.DB
}

func NewOEEEventRepository(db *gorm.DB) *OEEEventRepository {
	return &OEEEventRepository{db: db}
}

func (r *OEEEventRepository) List(ctx context.Context, oeeID int64) ([]model.OEEEvent, error) {
	var list []model.OEEEvent
	err := r.db.WithContext(ctx).Where("oee_id = ?", oeeID).Order("start_time ASC").Find(&list).Error
	return list, err
}

func (r *OEEEventRepository) Create(ctx context.Context, event *model.OEEEvent) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *OEEEventRepository) DeleteByOEEID(ctx context.Context, oeeID int64) error {
	return r.db.WithContext(ctx).Where("oee_id = ?", oeeID).Delete(&model.OEEEvent{}).Error
}
