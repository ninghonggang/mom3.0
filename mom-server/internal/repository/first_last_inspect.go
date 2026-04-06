package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type FirstLastInspectRepository struct {
	db *gorm.DB
}

func NewFirstLastInspectRepository(db *gorm.DB) *FirstLastInspectRepository {
	return &FirstLastInspectRepository{db: db}
}

func (r *FirstLastInspectRepository) List(ctx context.Context, tenantID int64, req *FirstLastInspectQuery) ([]model.MesFirstLastInspect, int64, error) {
	var list []model.MesFirstLastInspect
	var total int64
	q := r.db.WithContext(ctx).Model(&model.MesFirstLastInspect{}).Where("tenant_id = ?", tenantID)
	if req.OrderNo != "" {
		q = q.Where("production_order_id IN (SELECT id FROM pro_production_order WHERE order_no LIKE ?)", "%"+req.OrderNo+"%")
	}
	if req.InspectType != "" {
		q = q.Where("inspect_type = ?", req.InspectType)
	}
	if req.Status != "" {
		q = q.Where("status = ?", req.Status)
	}
	if req.StartTime != "" {
		q = q.Where("created_at >= ?", req.StartTime)
	}
	if req.EndTime != "" {
		q = q.Where("created_at <= ?", req.EndTime)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Limit(req.Limit).Offset(req.Offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *FirstLastInspectRepository) GetByID(ctx context.Context, id uint) (*model.MesFirstLastInspect, error) {
	var item model.MesFirstLastInspect
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *FirstLastInspectRepository) Create(ctx context.Context, item *model.MesFirstLastInspect) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *FirstLastInspectRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesFirstLastInspect{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FirstLastInspectRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.MesFirstLastInspect{}).Error
}

func (r *FirstLastInspectRepository) ListOverdue(ctx context.Context, tenantID int64) ([]model.MesFirstLastInspect, error) {
	var list []model.MesFirstLastInspect
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND status = ? AND created_at < ?", tenantID, "PENDING", "NOW() - INTERVAL '1 day'").
		Order("id DESC").
		Find(&list).Error
	return list, err
}

type FirstLastInspectQuery struct {
	OrderNo     string
	InspectType string
	Status      string
	StartTime   string
	EndTime     string
	Limit       int
	Offset      int
}