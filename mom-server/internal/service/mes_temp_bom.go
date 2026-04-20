package service

import (
	"context"
	"errors"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// TempBOMService 临时替代BOM服务
type TempBOMService struct {
	repo *repository.TempBOMRepository
}

func NewTempBOMService(repo *repository.TempBOMRepository) *TempBOMService {
	return &TempBOMService{repo: repo}
}

// Create 创建临时替代BOM
func (s *TempBOMService) Create(ctx context.Context, tenantID int64, req *model.TempBOMCreate, username string) (*model.TempBOM, error) {
	tempBOM := &model.TempBOM{
		TenantID:       tenantID,
		OrderDayID:     req.OrderDayID,
		OrderDayItemID: req.OrderDayItemID,
		OriginalBOMID:  req.OriginalBOMID,
		TempBOMName:    req.TempBOMName,
		BOMContent:     req.BOMContent,
		Reason:         req.Reason,
		Status:         0,
		Creator:        username,
	}

	if err := s.repo.Create(ctx, tempBOM); err != nil {
		return nil, fmt.Errorf("创建临时替代BOM失败: %w", err)
	}

	return s.repo.GetByID(ctx, uint(tempBOM.ID))
}

// GetByID 获取临时替代BOM
func (s *TempBOMService) GetByID(ctx context.Context, id int64) (*model.TempBOM, error) {
	return s.repo.GetByID(ctx, uint(id))
}

// Update 更新临时替代BOM
func (s *TempBOMService) Update(ctx context.Context, id int64, req *model.TempBOMUpdate, username string) (*model.TempBOM, error) {
	tempBOM, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	if tempBOM.Status != 0 {
		return nil, errors.New("仅待审核状态可更新")
	}

	updates := map[string]interface{}{}
	if req.TempBOMName != "" {
		updates["temp_bom_name"] = req.TempBOMName
	}
	if req.BOMContent != "" {
		updates["bom_content"] = req.BOMContent
	}
	if req.Reason != "" {
		updates["reason"] = req.Reason
	}

	if err := s.repo.Update(ctx, uint(id), updates); err != nil {
		return nil, fmt.Errorf("更新临时替代BOM失败: %w", err)
	}

	return s.repo.GetByID(ctx, uint(id))
}

// Delete 删除临时替代BOM
func (s *TempBOMService) Delete(ctx context.Context, id int64) error {
	tempBOM, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if tempBOM.Status != 0 {
		return errors.New("仅待审核状态可删除")
	}

	return s.repo.Delete(ctx, uint(id))
}

// ListByOrderDayItemID 根据日计划明细项ID查询临时替代BOM列表
func (s *TempBOMService) ListByOrderDayItemID(ctx context.Context, orderDayItemID int64) ([]model.TempBOM, error) {
	return s.repo.ListByOrderDayItemID(ctx, orderDayItemID)
}

// Approve 审核临时替代BOM
func (s *TempBOMService) Approve(ctx context.Context, id int64, status int, username string) error {
	tempBOM, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if tempBOM.Status != 0 {
		return errors.New("仅待审核状态可审核")
	}

	if status != 1 && status != 2 {
		return errors.New("审核状态无效，1=批准，2=拒绝")
	}

	return s.repo.Update(ctx, uint(id), map[string]interface{}{
		"status": status,
	})
}
