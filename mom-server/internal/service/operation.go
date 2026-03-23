package service

import (
	"context"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type OperationService struct {
	opRepo *repository.OperationRepository
}

func NewOperationService(opRepo *repository.OperationRepository) *OperationService {
	return &OperationService{opRepo: opRepo}
}

func (s *OperationService) List(ctx context.Context) ([]model.MdmOperation, int64, error) {
	return s.opRepo.List(ctx, 0)
}

func (s *OperationService) GetByID(ctx context.Context, id uint) (*model.MdmOperation, error) {
	return s.opRepo.GetByID(ctx, id)
}

func (s *OperationService) Create(ctx context.Context, op *model.MdmOperation) error {
	return s.opRepo.Create(ctx, op)
}

func (s *OperationService) Update(ctx context.Context, id uint, op *model.MdmOperation) error {
	updates := map[string]interface{}{
		"operation_name":     op.OperationName,
		"workcenter_id":      op.WorkcenterID,
		"workcenter_name":    op.WorkcenterName,
		"standard_worktime":  op.StandardWorktime,
		"quality_std":        op.QualityStd,
		"is_key_process":     op.IsKeyProcess,
		"is_qc_point":        op.IsQCPoint,
		"sequence":           op.Sequence,
	}
	if op.Remark != nil {
		updates["remark"] = op.Remark
	}
	return s.opRepo.Update(ctx, id, updates)
}

func (s *OperationService) Delete(ctx context.Context, id uint) error {
	return s.opRepo.Delete(ctx, id)
}
