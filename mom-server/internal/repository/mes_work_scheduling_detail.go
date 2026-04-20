package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
	"mom-server/internal/model"
)

// MesWorkSchedulingDetailRepository 工单排程明细仓储
type MesWorkSchedulingDetailRepository struct {
	db *gorm.DB
}

func NewMesWorkSchedulingDetailRepository(db *gorm.DB) *MesWorkSchedulingDetailRepository {
	return &MesWorkSchedulingDetailRepository{db: db}
}

func (r *MesWorkSchedulingDetailRepository) Create(ctx context.Context, m *model.MesWorkSchedulingDetail) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *MesWorkSchedulingDetailRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesWorkSchedulingDetail{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MesWorkSchedulingDetailRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.MesWorkSchedulingDetail{}, id).Error
}

func (r *MesWorkSchedulingDetailRepository) GetByID(ctx context.Context, id int64) (*model.MesWorkSchedulingDetail, error) {
	var m model.MesWorkSchedulingDetail
	err := r.db.WithContext(ctx).First(&m, id).Error
	return &m, err
}

func (r *MesWorkSchedulingDetailRepository) ListBySchedulingID(ctx context.Context, schedulingID int64) ([]model.MesWorkSchedulingDetail, error) {
	var list []model.MesWorkSchedulingDetail
	err := r.db.WithContext(ctx).Where("scheduling_id = ?", schedulingID).Order("id ASC").Find(&list).Error
	return list, err
}

func (r *MesWorkSchedulingDetailRepository) Page(ctx context.Context, tenantID int64, req *model.MesWorkSchedulingDetailPageReqVO) ([]model.MesWorkSchedulingDetail, int64, error) {
	var list []model.MesWorkSchedulingDetail
	var total int64

	q := r.db.WithContext(ctx).Model(&model.MesWorkSchedulingDetail{}).Where("tenant_id = ?", tenantID)

	if req.SchedulingID > 0 {
		q = q.Where("scheduling_id = ?", req.SchedulingID)
	}
	if req.Status != "" {
		q = q.Where("status = ?", req.Status)
	}
	if req.WorkingNode != "" {
		q = q.Where("working_node = ?", req.WorkingNode)
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

	err := q.Offset((page-1)*pageSize).Limit(pageSize).Order("id ASC").Find(&list).Error
	return list, total, err
}

// UpdateStatus 更新工序状态及时间
func (r *MesWorkSchedulingDetailRepository) UpdateStatus(ctx context.Context, id int64, status string, updates map[string]interface{}) error {
	updates["status"] = status
	return r.db.WithContext(ctx).Model(&model.MesWorkSchedulingDetail{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteBySchedulingID 删除指定排程下所有明细
func (r *MesWorkSchedulingDetailRepository) DeleteBySchedulingID(ctx context.Context, schedulingID int64) error {
	return r.db.WithContext(ctx).Where("scheduling_id = ?", schedulingID).Delete(&model.MesWorkSchedulingDetail{}).Error
}

// helper — exported for use in service
func TimePtr(t time.Time) *time.Time {
	return &t
}
