package repository

import (
	"context"

	"mom-platform/services/aps-service/internal/model"
)

type MpsPlanRepository interface {
	Create(ctx context.Context, m *model.MpsPlan) error
	Update(ctx context.Context, m *model.MpsPlan) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.MpsPlan, error)
	GetByPlanNo(ctx context.Context, tenantID, planNo string) (*model.MpsPlan, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.MpsPlan, int64, error)
	UpdateStatus(ctx context.Context, id uint64, status string) error
}

type MrpPlanRepository interface {
	Create(ctx context.Context, m *model.MrpPlan) error
	CreateBatch(ctx context.Context, plans []model.MrpPlan) error
	Update(ctx context.Context, m *model.MrpPlan) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.MrpPlan, error)
	GetByPlanNo(ctx context.Context, tenantID, planNo string) (*model.MrpPlan, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.MrpPlan, int64, error)
	ListByMpsID(ctx context.Context, mpsID uint64) ([]model.MrpPlan, error)
	DeleteByMpsID(ctx context.Context, mpsID uint64) error
}

type WorkCenterRepository interface {
	Create(ctx context.Context, wc *model.WorkCenter) error
	Update(ctx context.Context, wc *model.WorkCenter) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.WorkCenter, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.WorkCenter, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.WorkCenter, int64, error)
	ListByWorkshop(ctx context.Context, tenantID string, workshopID uint64) ([]model.WorkCenter, error)
}

type ScheduleJobRepository interface {
	Create(ctx context.Context, job *model.ScheduleJob) error
	Update(ctx context.Context, job *model.ScheduleJob) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.ScheduleJob, error)
	GetByJobNo(ctx context.Context, tenantID, jobNo string) (*model.ScheduleJob, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.ScheduleJob, int64, error)
	ListByWorkCenter(ctx context.Context, workCenterID uint64, startTime, endTime interface{}) ([]model.ScheduleJob, error)
}

type ChangeoverRepository interface {
	Create(ctx context.Context, c *model.Changeover) error
	Update(ctx context.Context, c *model.Changeover) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Changeover, error)
	GetChangeoverTime(ctx context.Context, fromMaterialID, toMaterialID uint64) (float64, error)
}
