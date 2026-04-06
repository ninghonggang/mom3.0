package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type FirstLastInspectService struct {
	repo *repository.FirstLastInspectRepository
}

func NewFirstLastInspectService(repo *repository.FirstLastInspectRepository) *FirstLastInspectService {
	return &FirstLastInspectService{repo: repo}
}

func (s *FirstLastInspectService) List(ctx context.Context, tenantID int64, req *repository.FirstLastInspectQuery) ([]model.MesFirstLastInspect, int64, error) {
	return s.repo.List(ctx, tenantID, req)
}

func (s *FirstLastInspectService) GetByID(ctx context.Context, id uint) (*model.MesFirstLastInspect, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FirstLastInspectService) Create(ctx context.Context, item *model.MesFirstLastInspect) error {
	if item.TenantID == 0 {
		item.TenantID = 1
	}
	item.InspectNo = fmt.Sprintf("INS-%d-%d", time.Now().Unix(), item.ProductionOrderID%10000)
	return s.repo.Create(ctx, item)
}

func (s *FirstLastInspectService) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return s.repo.Update(ctx, id, updates)
}

func (s *FirstLastInspectService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *FirstLastInspectService) ListOverdue(ctx context.Context, tenantID int64) ([]model.MesFirstLastInspect, error) {
	return s.repo.ListOverdue(ctx, tenantID)
}

// CheckOverdueInspect 定时任务：检查超期未检
func (s *FirstLastInspectService) CheckOverdueInspect() {
	// TODO: 实现超期提醒逻辑
}