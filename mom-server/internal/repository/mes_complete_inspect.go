package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// CompleteInspectRepository 齐套检查仓储
type CompleteInspectRepository struct {
	db *gorm.DB
}

func NewCompleteInspectRepository(db *gorm.DB) *CompleteInspectRepository {
	return &CompleteInspectRepository{db: db}
}

// Create 创建齐套检查记录
func (r *CompleteInspectRepository) Create(ctx context.Context, inspect *model.MesCompleteInspect) error {
	return r.db.WithContext(ctx).Create(inspect).Error
}

// GetByID 根据ID获取齐套检查记录
func (r *CompleteInspectRepository) GetByID(ctx context.Context, id uint) (*model.MesCompleteInspect, error) {
	var inspect model.MesCompleteInspect
	if err := r.db.WithContext(ctx).First(&inspect, id).Error; err != nil {
		return nil, err
	}
	return &inspect, nil
}

// GetByOrderDayID 根据日计划ID获取齐套检查记录
func (r *CompleteInspectRepository) GetByOrderDayID(ctx context.Context, orderDayID int64) (*model.MesCompleteInspect, error) {
	var inspect model.MesCompleteInspect
	if err := r.db.WithContext(ctx).Where("order_day_id = ?", orderDayID).First(&inspect).Error; err != nil {
		return nil, err
	}
	return &inspect, nil
}

// Update 更新齐套检查记录
func (r *CompleteInspectRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesCompleteInspect{}).Where("id = ?", id).Updates(updates).Error
}

// UpdateByOrderDayID 根据日计划ID更新齐套检查记录
func (r *CompleteInspectRepository) UpdateByOrderDayID(ctx context.Context, orderDayID int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesCompleteInspect{}).Where("order_day_id = ?", orderDayID).Updates(updates).Error
}

// Delete 删除齐套检查记录
func (r *CompleteInspectRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MesCompleteInspect{}, id).Error
}

// GetConfigByCode 根据配置码获取配置信息
func (r *CompleteInspectRepository) GetConfigByCode(ctx context.Context, configCode string) (*model.MesConfigInfo, error) {
	var config model.MesConfigInfo
	// 从sys_dict_data表获取配置，dict_type='mes_config'
	err := r.db.WithContext(ctx).Model(&model.DictData{}).
		Select("id", "dict_label as config_code", "dict_label as config_name", "dict_value as config_value", "status").
		Where("dict_type = ? AND dict_key = ?", "mes_config", configCode).
		First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// ListOrderDayBom 查询日计划BOM信息
func (r *CompleteInspectRepository) ListOrderDayBom(ctx context.Context, orderDayID int64) ([]model.MesOrderDayBomRespVO, error) {
	var result []model.MesOrderDayBomRespVO
	// 实际应关联mes_order_day_item和物料表，此处返回空列表供后续扩展
	err := r.db.WithContext(ctx).
		Model(&model.OrderDayItem{}).
		Select(`item.id, item.day_plan_id as order_day_id, od.day_plan_no,
			item.product_id, p.product_code, p.product_name,
			0 as material_id, '' as material_code, '' as material_name,
			'' as specification, '' as unit, 0 as required_qty, 0 as available_qty, 0 as shortage_qty,
			item.kit_status, 0 as warehouse_id, '' as warehouse_name`).
		Joins("LEFT JOIN mes_order_day od ON item.day_plan_id = od.id").
		Joins("LEFT JOIN mdm_product p ON item.product_id = p.id").
		Where("item.day_plan_id = ?", orderDayID).
		Scan(&result).Error
	return result, err
}

// ListOrderDayBomPage 分页查询日计划BOM信息
func (r *CompleteInspectRepository) ListOrderDayBomPage(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.MesOrderDayBomRespVO, int64, error) {
	var result []model.MesOrderDayBomRespVO
	var total int64

	db := r.db.WithContext(ctx).Model(&model.OrderDayItem{}).
		Select(`item.id, item.day_plan_id as order_day_id, od.day_plan_no,
			item.product_id, p.product_code, p.product_name,
			0 as material_id, '' as material_code, '' as material_name,
			'' as specification, '' as unit, 0 as required_qty, 0 as available_qty, 0 as shortage_qty,
			item.kit_status, 0 as warehouse_id, '' as warehouse_name`).
		Joins("LEFT JOIN mes_order_day od ON item.day_plan_id = od.id").
		Joins("LEFT JOIN mdm_product p ON item.product_id = p.id").
		Where("od.tenant_id = ?", tenantID)

	if orderDayID, ok := query["order_day_id"].(int64); ok && orderDayID > 0 {
		db = db.Where("item.day_plan_id = ?", orderDayID)
	}
	if orderDayNo, ok := query["order_day_no"].(string); ok && orderDayNo != "" {
		db = db.Where("od.day_plan_no LIKE ?", "%"+orderDayNo+"%")
	}

	db.Count(&total)

	page := 1
	pageSize := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if ps, ok := query["page_size"].(int); ok && ps > 0 {
		pageSize = ps
	}

	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Order("item.id DESC").Scan(&result).Error
	return result, total, err
}

// ListOrderDayWorker 查询日计划人员信息
func (r *CompleteInspectRepository) ListOrderDayWorker(ctx context.Context, orderDayID int64) ([]model.MesOrderDayWorkerRespVO, error) {
	var result []model.MesOrderDayWorkerRespVO
	// 实际应关联工单人员表，此处返回空列表供后续扩展
	err := r.db.WithContext(ctx).
		Model(&model.OrderDayItem{}).
		Select(`0 as id, item.day_plan_id as order_day_id, od.day_plan_no,
			item.product_id, p.product_code, p.product_name,
			COALESCE(item.process_route_id, 0) as process_route_id, '' as process_route_name,
			0 as worker_id, '' as worker_code, '' as worker_name,
			0 as team_id, '' as team_name, '' as shift_type,
			item.kit_status`).
		Joins("LEFT JOIN mes_order_day od ON item.day_plan_id = od.id").
		Joins("LEFT JOIN mdm_product p ON item.product_id = p.id").
		Where("item.day_plan_id = ?", orderDayID).
		Scan(&result).Error
	return result, err
}

// ListOrderDayWorkerPage 分页查询日计划人员信息
func (r *CompleteInspectRepository) ListOrderDayWorkerPage(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.MesOrderDayWorkerRespVO, int64, error) {
	var result []model.MesOrderDayWorkerRespVO
	var total int64

	db := r.db.WithContext(ctx).Model(&model.OrderDayItem{}).
		Select(`0 as id, item.day_plan_id as order_day_id, od.day_plan_no,
			item.product_id, p.product_code, p.product_name,
			COALESCE(item.process_route_id, 0) as process_route_id, '' as process_route_name,
			0 as worker_id, '' as worker_code, '' as worker_name,
			0 as team_id, '' as team_name, '' as shift_type,
			item.kit_status`).
		Joins("LEFT JOIN mes_order_day od ON item.day_plan_id = od.id").
		Joins("LEFT JOIN mdm_product p ON item.product_id = p.id").
		Where("od.tenant_id = ?", tenantID)

	if orderDayID, ok := query["order_day_id"].(int64); ok && orderDayID > 0 {
		db = db.Where("item.day_plan_id = ?", orderDayID)
	}
	if orderDayNo, ok := query["order_day_no"].(string); ok && orderDayNo != "" {
		db = db.Where("od.day_plan_no LIKE ?", "%"+orderDayNo+"%")
	}

	db.Count(&total)

	page := 1
	pageSize := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if ps, ok := query["page_size"].(int); ok && ps > 0 {
		pageSize = ps
	}

	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Order("item.id DESC").Scan(&result).Error
	return result, total, err
}

// ListOrderDayEquipment 查询日计划设备信息
func (r *CompleteInspectRepository) ListOrderDayEquipment(ctx context.Context, orderDayID int64) ([]model.MesOrderDayEquipmentRespVO, error) {
	var result []model.MesOrderDayEquipmentRespVO
	// 实际应关联工单设备表，此处返回空列表供后续扩展
	err := r.db.WithContext(ctx).
		Model(&model.OrderDayItem{}).
		Select(`0 as id, item.day_plan_id as order_day_id, od.day_plan_no,
			item.product_id, p.product_code, p.product_name,
			COALESCE(item.process_route_id, 0) as process_route_id, '' as process_route_name,
			0 as equipment_id, '' as equipment_code, '' as equipment_name,
			0 as workstation_id, '' as workstation_name,
			'' as status, item.kit_status`).
		Joins("LEFT JOIN mes_order_day od ON item.day_plan_id = od.id").
		Joins("LEFT JOIN mdm_product p ON item.product_id = p.id").
		Where("item.day_plan_id = ?", orderDayID).
		Scan(&result).Error
	return result, err
}

// ListOrderDayEquipmentPage 分页查询日计划设备信息
func (r *CompleteInspectRepository) ListOrderDayEquipmentPage(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.MesOrderDayEquipmentRespVO, int64, error) {
	var result []model.MesOrderDayEquipmentRespVO
	var total int64

	db := r.db.WithContext(ctx).Model(&model.OrderDayItem{}).
		Select(`0 as id, item.day_plan_id as order_day_id, od.day_plan_no,
			item.product_id, p.product_code, p.product_name,
			COALESCE(item.process_route_id, 0) as process_route_id, '' as process_route_name,
			0 as equipment_id, '' as equipment_code, '' as equipment_name,
			0 as workstation_id, '' as workstation_name,
			'' as status, item.kit_status`).
		Joins("LEFT JOIN mes_order_day od ON item.day_plan_id = od.id").
		Joins("LEFT JOIN mdm_product p ON item.product_id = p.id").
		Where("od.tenant_id = ?", tenantID)

	if orderDayID, ok := query["order_day_id"].(int64); ok && orderDayID > 0 {
		db = db.Where("item.day_plan_id = ?", orderDayID)
	}
	if orderDayNo, ok := query["order_day_no"].(string); ok && orderDayNo != "" {
		db = db.Where("od.day_plan_no LIKE ?", "%"+orderDayNo+"%")
	}

	db.Count(&total)

	page := 1
	pageSize := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if ps, ok := query["page_size"].(int); ok && ps > 0 {
		pageSize = ps
	}

	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Order("item.id DESC").Scan(&result).Error
	return result, total, err
}
