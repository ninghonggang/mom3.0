package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"mom-platform/services/aps-service/internal/model"
)

var ErrNotFound = errors.New("record not found")

// -- MPS Plan --

type gormMpsPlanRepo struct{ db *gorm.DB }

func NewMpsPlanRepo(db *gorm.DB) MpsPlanRepository { return &gormMpsPlanRepo{db: db} }

func (r *gormMpsPlanRepo) Create(ctx context.Context, m *model.MpsPlan) error {
	return r.db.WithContext(ctx).Create(m).Error
}
func (r *gormMpsPlanRepo) Update(ctx context.Context, m *model.MpsPlan) error {
	return r.db.WithContext(ctx).Save(m).Error
}
func (r *gormMpsPlanRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.MpsPlan{}, id).Error
}
func (r *gormMpsPlanRepo) GetByID(ctx context.Context, id uint64) (*model.MpsPlan, error) {
	var m model.MpsPlan
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
func (r *gormMpsPlanRepo) GetByPlanNo(ctx context.Context, tenantID, planNo string) (*model.MpsPlan, error) {
	var m model.MpsPlan
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND plan_no = ?", tenantID, planNo).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
func (r *gormMpsPlanRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.MpsPlan, int64, error) {
	var list []model.MpsPlan
	var total int64
	query := r.db.WithContext(ctx).Model(&model.MpsPlan{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *gormMpsPlanRepo) UpdateStatus(ctx context.Context, id uint64, status string) error {
	return r.db.WithContext(ctx).Model(&model.MpsPlan{}).Where("id = ?", id).Update("status", status).Error
}

// -- MRP Plan --

type gormMrpPlanRepo struct{ db *gorm.DB }

func NewMrpPlanRepo(db *gorm.DB) MrpPlanRepository { return &gormMrpPlanRepo{db: db} }

func (r *gormMrpPlanRepo) Create(ctx context.Context, m *model.MrpPlan) error {
	return r.db.WithContext(ctx).Create(m).Error
}
func (r *gormMrpPlanRepo) CreateBatch(ctx context.Context, plans []model.MrpPlan) error {
	return r.db.WithContext(ctx).CreateInBatches(plans, 100).Error
}
func (r *gormMrpPlanRepo) Update(ctx context.Context, m *model.MrpPlan) error {
	return r.db.WithContext(ctx).Save(m).Error
}
func (r *gormMrpPlanRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.MrpPlan{}, id).Error
}
func (r *gormMrpPlanRepo) GetByID(ctx context.Context, id uint64) (*model.MrpPlan, error) {
	var m model.MrpPlan
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
func (r *gormMrpPlanRepo) GetByPlanNo(ctx context.Context, tenantID, planNo string) (*model.MrpPlan, error) {
	var m model.MrpPlan
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND plan_no = ?", tenantID, planNo).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
func (r *gormMrpPlanRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.MrpPlan, int64, error) {
	var list []model.MrpPlan
	var total int64
	query := r.db.WithContext(ctx).Model(&model.MrpPlan{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *gormMrpPlanRepo) ListByMpsID(ctx context.Context, mpsID uint64) ([]model.MrpPlan, error) {
	var list []model.MrpPlan
	if err := r.db.WithContext(ctx).Where("mps_id = ?", mpsID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (r *gormMrpPlanRepo) DeleteByMpsID(ctx context.Context, mpsID uint64) error {
	return r.db.WithContext(ctx).Where("mps_id = ?", mpsID).Delete(&model.MrpPlan{}).Error
}

// -- Work Center --

type gormWorkCenterRepo struct{ db *gorm.DB }

func NewWorkCenterRepo(db *gorm.DB) WorkCenterRepository { return &gormWorkCenterRepo{db: db} }

func (r *gormWorkCenterRepo) Create(ctx context.Context, wc *model.WorkCenter) error {
	return r.db.WithContext(ctx).Create(wc).Error
}
func (r *gormWorkCenterRepo) Update(ctx context.Context, wc *model.WorkCenter) error {
	return r.db.WithContext(ctx).Save(wc).Error
}
func (r *gormWorkCenterRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.WorkCenter{}, id).Error
}
func (r *gormWorkCenterRepo) GetByID(ctx context.Context, id uint64) (*model.WorkCenter, error) {
	var wc model.WorkCenter
	if err := r.db.WithContext(ctx).First(&wc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &wc, nil
}
func (r *gormWorkCenterRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.WorkCenter, error) {
	var wc model.WorkCenter
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND center_code = ?", tenantID, code).First(&wc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &wc, nil
}
func (r *gormWorkCenterRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.WorkCenter, int64, error) {
	var list []model.WorkCenter
	var total int64
	query := r.db.WithContext(ctx).Model(&model.WorkCenter{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *gormWorkCenterRepo) ListByWorkshop(ctx context.Context, tenantID string, workshopID uint64) ([]model.WorkCenter, error) {
	var list []model.WorkCenter
	q := r.db.WithContext(ctx)
	if tenantID != "" && tenantID != "0" {
		q = q.Where("tenant_id = ?", tenantID)
	}
	if workshopID > 0 {
		q = q.Where("workshop_id = ?", workshopID)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// -- Schedule Job --

type gormScheduleJobRepo struct{ db *gorm.DB }

func NewScheduleJobRepo(db *gorm.DB) ScheduleJobRepository { return &gormScheduleJobRepo{db: db} }

func (r *gormScheduleJobRepo) Create(ctx context.Context, job *model.ScheduleJob) error {
	return r.db.WithContext(ctx).Create(job).Error
}
func (r *gormScheduleJobRepo) Update(ctx context.Context, job *model.ScheduleJob) error {
	return r.db.WithContext(ctx).Save(job).Error
}
func (r *gormScheduleJobRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.ScheduleJob{}, id).Error
}
func (r *gormScheduleJobRepo) GetByID(ctx context.Context, id uint64) (*model.ScheduleJob, error) {
	var job model.ScheduleJob
	if err := r.db.WithContext(ctx).First(&job, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &job, nil
}
func (r *gormScheduleJobRepo) GetByJobNo(ctx context.Context, tenantID, jobNo string) (*model.ScheduleJob, error) {
	var job model.ScheduleJob
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND job_no = ?", tenantID, jobNo).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &job, nil
}
func (r *gormScheduleJobRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.ScheduleJob, int64, error) {
	var list []model.ScheduleJob
	var total int64
	query := r.db.WithContext(ctx).Model(&model.ScheduleJob{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 必须显式排序：否则 PostgreSQL 不保证 OFFSET/LIMIT across pages 的顺序稳定，
	// 会出现翻页时记录重复或漏读。id DESC 让最新排程排在首页。
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *gormScheduleJobRepo) ListByWorkCenter(ctx context.Context, workCenterID uint64, startTime, endTime interface{}) ([]model.ScheduleJob, error) {
	var list []model.ScheduleJob
	query := r.db.WithContext(ctx).Where("work_center_id = ?", workCenterID)
	if startTime != nil {
		query = query.Where("planned_start >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("planned_end <= ?", endTime)
	}
	if err := query.Order("planned_start ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// -- Changeover --

type gormChangeoverRepo struct{ db *gorm.DB }

func NewChangeoverRepo(db *gorm.DB) ChangeoverRepository { return &gormChangeoverRepo{db: db} }

func (r *gormChangeoverRepo) Create(ctx context.Context, c *model.Changeover) error {
	return r.db.WithContext(ctx).Create(c).Error
}
func (r *gormChangeoverRepo) Update(ctx context.Context, c *model.Changeover) error {
	return r.db.WithContext(ctx).Save(c).Error
}
func (r *gormChangeoverRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Changeover{}, id).Error
}
func (r *gormChangeoverRepo) GetByID(ctx context.Context, id uint64) (*model.Changeover, error) {
	var c model.Changeover
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &c, nil
}
func (r *gormChangeoverRepo) GetChangeoverTime(ctx context.Context, fromMaterialID, toMaterialID uint64) (float64, error) {
	var c model.Changeover
	if err := r.db.WithContext(ctx).
		Where("from_material_id = ? AND to_material_id = ?", fromMaterialID, toMaterialID).
		First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrNotFound
		}
		return 0, err
	}
	return c.ChangeoverTimeMinutes, nil
}
