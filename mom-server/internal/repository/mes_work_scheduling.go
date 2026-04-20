package repository

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"mom-server/internal/model"
)

// MesWorkSchedulingRepository 工单排程主表仓储
type MesWorkSchedulingRepository struct {
	db *gorm.DB
}

func NewMesWorkSchedulingRepository(db *gorm.DB) *MesWorkSchedulingRepository {
	return &MesWorkSchedulingRepository{db: db}
}

func (r *MesWorkSchedulingRepository) Create(ctx context.Context, m *model.MesWorkScheduling) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *MesWorkSchedulingRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesWorkScheduling{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MesWorkSchedulingRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.MesWorkScheduling{}, id).Error
}

func (r *MesWorkSchedulingRepository) GetByID(ctx context.Context, id int64) (*model.MesWorkScheduling, error) {
	var m model.MesWorkScheduling
	err := r.db.WithContext(ctx).First(&m, id).Error
	return &m, err
}

func (r *MesWorkSchedulingRepository) Page(ctx context.Context, tenantID int64, req *model.WorkSchedulePageVO) ([]model.MesWorkScheduling, int64, error) {
	var list []model.MesWorkScheduling
	var total int64

	q := r.db.WithContext(ctx).Model(&model.MesWorkScheduling{}).Where("tenant_id = ?", tenantID)

	if req.SchedulingCode != "" {
		q = q.Where("scheduling_code LIKE ?", "%"+req.SchedulingCode+"%")
	}
	if req.ProductCode != "" {
		q = q.Where("product_code LIKE ?", "%"+req.ProductCode+"%")
	}
	if req.Status != "" {
		q = q.Where("status = ?", req.Status)
	}
	if req.PlanDate != "" {
		q = q.Where("plan_date = ?", req.PlanDate)
	}
	if req.WorkshopCode != "" {
		q = q.Where("workshop_code = ?", req.WorkshopCode)
	}
	if req.LineCode != "" {
		q = q.Where("line_code = ?", req.LineCode)
	}

	q.Count(&total)

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	err := q.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&list).Error
	return list, total, err
}

// GenerateSchedulingCode 生成排程编码 WS-YYYYMMDD-NNN
func (r *MesWorkSchedulingRepository) GenerateSchedulingCode(ctx context.Context, tenantID int64) (string, error) {
	dateStr := time.Now().Format("20060102")
	var count int64
	r.db.WithContext(ctx).Model(&model.MesWorkScheduling{}).
		Where("tenant_id = ? AND scheduling_code LIKE ?", tenantID, fmt.Sprintf("WS-%s-%%", dateStr)).
		Count(&count)
	return fmt.Sprintf("WS-%s-%03d", dateStr, count+1), nil
}
