package service

import (
	"context"
	"errors"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// MesProcessService 工艺路线服务
type MesProcessService struct {
	processRepo *repository.MesProcessRepository
	opRepo      *repository.MesProcessOperationRepository
}

func NewMesProcessService(processRepo *repository.MesProcessRepository, opRepo *repository.MesProcessOperationRepository) *MesProcessService {
	return &MesProcessService{processRepo: processRepo, opRepo: opRepo}
}

// List 获取工艺路线列表
func (s *MesProcessService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.MesProcess, int64, error) {
	return s.processRepo.List(ctx, tenantID, query)
}

// GetByID 获取工艺路线详情
func (s *MesProcessService) GetByID(ctx context.Context, id uint) (*model.MesProcess, error) {
	process, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// 加载工序明细
	operations, err := s.opRepo.ListByProcessID(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	process.Operations = operations
	return process, nil
}

// GetByMaterialID 获取产品的有效工艺路线
func (s *MesProcessService) GetByMaterialID(ctx context.Context, tenantID int64, materialID int64) ([]model.MesProcess, error) {
	return s.processRepo.GetByMaterialID(ctx, tenantID, materialID)
}

// Create 创建工艺路线
func (s *MesProcessService) Create(ctx context.Context, tenantID int64, req *model.MesProcessCreate) (*model.MesProcess, error) {
	// 检查编号是否已存在
	existing, _ := s.processRepo.GetByCode(ctx, tenantID, req.ProcessCode)
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("工艺路线编号已存在")
	}

	// 解析日期
	var effDate, expDate *time.Time
	if req.EffDate != "" {
		t, err := time.Parse("2006-01-02", req.EffDate)
		if err == nil {
			effDate = &t
		}
	}
	if req.ExpDate != "" {
		t, err := time.Parse("2006-01-02", req.ExpDate)
		if err == nil {
			expDate = &t
		}
	}

	process := &model.MesProcess{
		TenantID:     tenantID,
		ProcessCode:  req.ProcessCode,
		ProcessName:  req.ProcessName,
		Version:     req.Version,
		Status:      req.Status,
		EffDate:     effDate,
		ExpDate:     expDate,
	}
	if req.MaterialID != nil {
		process.MaterialID = req.MaterialID
	}
	if req.MaterialCode != "" {
		process.MaterialCode = &req.MaterialCode
	}
	if req.MaterialName != "" {
		process.MaterialName = &req.MaterialName
	}
	if req.Remark != "" {
		process.Remark = &req.Remark
	}
	if process.Status == "" {
		process.Status = "DRAFT"
	}

	// 创建工艺路线
	if err := s.processRepo.Create(ctx, process); err != nil {
		return nil, err
	}

	// 创建工序明细
	for _, op := range req.Operations {
		operation := &model.MesProcessOperation{
			TenantID:          tenantID,
			ProcessID:         process.ID,
			OperationID:       op.OperationID,
			OperationCode:     op.OperationCode,
			OperationName:     op.OperationName,
			LineNo:            op.LineNo,
			StandardWorktime:  op.StandardWorktime,
			WorkcenterID:      op.WorkcenterID,
			RequiredCapacity:  op.RequiredCapacity,
			MinWorkers:        op.MinWorkers,
			MaxWorkers:        op.MaxWorkers,
			IsKeyProcess:      op.IsKeyProcess,
			IsQCPoint:         op.IsQCPoint,
			Status:            "ACTIVE",
		}
		if op.WorkcenterName != "" {
			operation.WorkcenterName = &op.WorkcenterName
		}
		if op.QualityStd != "" {
			operation.QualityStd = &op.QualityStd
		}
		if op.Remark != "" {
			operation.Remark = &op.Remark
		}
		if err := s.opRepo.Create(ctx, operation); err != nil {
			return nil, err
		}
	}

	return s.GetByID(ctx, uint(process.ID))
}

// Update 更新工艺路线
func (s *MesProcessService) Update(ctx context.Context, id uint, req *model.MesProcessUpdate) (*model.MesProcess, error) {
	process, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 仅草稿状态可更新
	if process.Status != "DRAFT" {
		return nil, errors.New("仅草稿状态可更新")
	}

	// 解析日期
	var effDate, expDate *time.Time
	if req.EffDate != "" {
		t, err := time.Parse("2006-01-02", req.EffDate)
		if err == nil {
			effDate = &t
		}
	}
	if req.ExpDate != "" {
		t, err := time.Parse("2006-01-02", req.ExpDate)
		if err == nil {
			expDate = &t
		}
	}

	updates := map[string]interface{}{}
	if req.ProcessName != "" {
		updates["process_name"] = req.ProcessName
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if effDate != nil {
		updates["eff_date"] = effDate
	}
	if expDate != nil {
		updates["exp_date"] = expDate
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	if err := s.processRepo.Update(ctx, id, updates); err != nil {
		return nil, err
	}

	// 更新工序明细：先删后插
	if err := s.opRepo.DeleteByProcessID(ctx, int64(id)); err != nil {
		return nil, err
	}

	for _, op := range req.Operations {
		operation := &model.MesProcessOperation{
			TenantID:          process.TenantID,
			ProcessID:        int64(id),
			OperationID:      op.OperationID,
			OperationCode:    op.OperationCode,
			OperationName:    op.OperationName,
			LineNo:           op.LineNo,
			StandardWorktime: op.StandardWorktime,
			WorkcenterID:     op.WorkcenterID,
			RequiredCapacity: op.RequiredCapacity,
			MinWorkers:       op.MinWorkers,
			MaxWorkers:       op.MaxWorkers,
			IsKeyProcess:     op.IsKeyProcess,
			IsQCPoint:        op.IsQCPoint,
			Status:           "ACTIVE",
		}
		if op.WorkcenterName != "" {
			operation.WorkcenterName = &op.WorkcenterName
		}
		if op.QualityStd != "" {
			operation.QualityStd = &op.QualityStd
		}
		if op.Remark != "" {
			operation.Remark = &op.Remark
		}
		if err := s.opRepo.Create(ctx, operation); err != nil {
			return nil, err
		}
	}

	return s.GetByID(ctx, id)
}

// Delete 删除工艺路线
func (s *MesProcessService) Delete(ctx context.Context, id uint) error {
	process, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 仅草稿状态可删除
	if process.Status != "DRAFT" {
		return errors.New("仅草稿状态可删除")
	}

	// 删除工序明细
	if err := s.opRepo.DeleteByProcessID(ctx, int64(id)); err != nil {
		return err
	}

	return s.processRepo.Delete(ctx, id)
}

// UpdateStatus 更新工艺路线状态
func (s *MesProcessService) UpdateStatus(ctx context.Context, id uint, status string) error {
	return s.processRepo.Update(ctx, id, map[string]interface{}{"status": status})
}

// Copy 复制工艺路线（带版本递增）
func (s *MesProcessService) Copy(ctx context.Context, id uint) (*model.MesProcess, error) {
	original, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 生成新编号
	newCode := original.ProcessCode + "-COPY"

	// 检查新编号是否已存在
	existing, _ := s.processRepo.GetByCode(ctx, original.TenantID, newCode)
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("复制后的工艺路线编号已存在")
	}

	// 版本递增
	newVersion := "V1"

	newProcess := &model.MesProcess{
		TenantID:     original.TenantID,
		ProcessCode:  newCode,
		ProcessName:  original.ProcessName + "(复制)",
		MaterialID:   original.MaterialID,
		MaterialCode: original.MaterialCode,
		MaterialName: original.MaterialName,
		Version:      newVersion,
		Status:       "DRAFT",
		EffDate:      original.EffDate,
		ExpDate:      original.ExpDate,
	}

	if err := s.processRepo.Create(ctx, newProcess); err != nil {
		return nil, err
	}

	// 复制工序明细
	operations, err := s.opRepo.ListByProcessID(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	for _, op := range operations {
		newOp := &model.MesProcessOperation{
			TenantID:          original.TenantID,
			ProcessID:         newProcess.ID,
			OperationID:       op.OperationID,
			OperationCode:     op.OperationCode,
			OperationName:     op.OperationName,
			LineNo:            op.LineNo,
			StandardWorktime:  op.StandardWorktime,
			WorkcenterID:      op.WorkcenterID,
			WorkcenterName:    op.WorkcenterName,
			RequiredCapacity:  op.RequiredCapacity,
			MinWorkers:        op.MinWorkers,
			MaxWorkers:        op.MaxWorkers,
			IsKeyProcess:      op.IsKeyProcess,
			IsQCPoint:         op.IsQCPoint,
			QualityStd:        op.QualityStd,
			Status:            "ACTIVE",
		}
		if err := s.opRepo.Create(ctx, newOp); err != nil {
			return nil, err
		}
	}

	return s.GetByID(ctx, uint(newProcess.ID))
}

// ValidateProcess 验证工艺路线合法性
func (s *MesProcessService) ValidateProcess(ctx context.Context, id uint) error {
	process, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if process.Status != "ACTIVE" {
		return errors.New("工艺路线状态必须为生效才能使用")
	}
	operations, err := s.opRepo.ListByProcessID(ctx, int64(id))
	if err != nil {
		return err
	}
	if len(operations) == 0 {
		return errors.New("工艺路线工序不能为空")
	}
	// 检查工序顺序是否连续
	for i, op := range operations {
		if op.LineNo != i+1 {
			return errors.New("工序顺序号必须连续")
		}
	}
	return nil
}

// CheckOperationCapacity 检查工序能力是否满足
func (s *MesProcessService) CheckOperationCapacity(ctx context.Context, processID int64, operationLineNo int, requiredHours float64) (bool, error) {
	operations, err := s.opRepo.ListByProcessID(ctx, processID)
	if err != nil {
		return false, err
	}
	for _, op := range operations {
		if op.LineNo == operationLineNo {
			return op.RequiredCapacity >= requiredHours, nil
		}
	}
	return false, errors.New("未找到指定工序")
}
