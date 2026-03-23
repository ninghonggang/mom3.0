package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"mom-server/internal/dto"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type BOMService struct {
	bomRepo    *repository.BOMRepository
	bomItemRepo *repository.BOMItemRepository
}

func NewBOMService(bomRepo *repository.BOMRepository, bomItemRepo *repository.BOMItemRepository) *BOMService {
	return &BOMService{bomRepo: bomRepo, bomItemRepo: bomItemRepo}
}

func (s *BOMService) List(ctx context.Context) ([]model.MdmBOM, int64, error) {
	return s.bomRepo.List(ctx, 0)
}

func (s *BOMService) GetByID(ctx context.Context, id uint) (*model.MdmBOM, error) {
	return s.bomRepo.GetByID(ctx, id)
}

// GetWithItems 获取BOM及明细
func (s *BOMService) GetWithItems(ctx context.Context, id uint) (*BOMWithItems, error) {
	bom, err := s.bomRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	items, err := s.bomItemRepo.ListByBOMID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &BOMWithItems{BOM: bom, Items: items}, nil
}

// BOMWithItems BOM及明细结构
type BOMWithItems struct {
	BOM   *model.MdmBOM         `json:"bom"`
	Items []model.MdmBOMItem    `json:"items"`
}

func (s *BOMService) Create(ctx context.Context, req *dto.BOMCreateReq) error {
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

	bom := &model.MdmBOM{
		BOMCode:      req.BOMCode,
		BOMName:      req.BOMName,
		MaterialID:   req.MaterialID,
		MaterialCode: req.MaterialCode,
		MaterialName: req.MaterialName,
		Version:      req.Version,
		Status:       req.Status,
		EffDate:      effDate,
		ExpDate:      expDate,
	}
	if req.Remark != "" {
		bom.Remark = &req.Remark
	}
	if bom.Status == "" {
		bom.Status = "DRAFT"
	}

	// 创建BOM
	if err := s.bomRepo.Create(ctx, bom); err != nil {
		return err
	}

	// 创建BOM明细
	for _, item := range req.Items {
		bomItem := &model.MdmBOMItem{
			BOMID:          bom.ID,
			LineNo:         item.LineNo,
			MaterialID:     item.MaterialID,
			MaterialCode:   item.MaterialCode,
			MaterialName:   item.MaterialName,
			Quantity:       item.Quantity,
			Unit:          item.Unit,
			ScrapRate:     item.ScrapRate,
			IsAlternative: item.IsAlternative,
		}
		if item.SubstituteGroup != "" {
			bomItem.SubstituteGroup = &item.SubstituteGroup
		}
		if err := s.bomItemRepo.Create(ctx, bomItem); err != nil {
			return err
		}
	}

	return nil
}

func (s *BOMService) Update(ctx context.Context, id uint, req *dto.BOMCreateReq) error {
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

	updates := map[string]interface{}{
		"bom_name":      req.BOMName,
		"material_id":    req.MaterialID,
		"material_code":  req.MaterialCode,
		"material_name":  req.MaterialName,
		"version":        req.Version,
		"eff_date":       effDate,
		"exp_date":       expDate,
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	if err := s.bomRepo.Update(ctx, id, updates); err != nil {
		return err
	}

	// 更新明细：先删后插
	if err := s.bomItemRepo.DeleteByBOMID(ctx, id); err != nil {
		return err
	}
	for _, item := range req.Items {
		bomItem := &model.MdmBOMItem{
			BOMID:          int64(id),
			LineNo:         item.LineNo,
			MaterialID:     item.MaterialID,
			MaterialCode:   item.MaterialCode,
			MaterialName:   item.MaterialName,
			Quantity:       item.Quantity,
			Unit:          item.Unit,
			ScrapRate:     item.ScrapRate,
			IsAlternative: item.IsAlternative,
		}
		if item.SubstituteGroup != "" {
			bomItem.SubstituteGroup = &item.SubstituteGroup
		}
		if err := s.bomItemRepo.Create(ctx, bomItem); err != nil {
			return err
		}
	}

	return nil
}

func (s *BOMService) Delete(ctx context.Context, id uint) error {
	// 先删明细
	if err := s.bomItemRepo.DeleteByBOMID(ctx, id); err != nil {
		return err
	}
	return s.bomRepo.Delete(ctx, id)
}

func (s *BOMService) UpdateStatus(ctx context.Context, id uint, status string) error {
	return s.bomRepo.Update(ctx, id, map[string]interface{}{"status": status})
}

// Copy 复制BOM（带版本递增）
func (s *BOMService) Copy(ctx context.Context, id uint) (*model.MdmBOM, error) {
	original, err := s.bomRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 版本递增
	newVersion := fmt.Sprintf("V%d", len(original.Version)+1)

	newBOM := &model.MdmBOM{
		BOMCode:      original.BOMCode + "-COPY",
		BOMName:      original.BOMName + "(复制)",
		MaterialID:   original.MaterialID,
		MaterialCode: original.MaterialCode,
		MaterialName: original.MaterialName,
		Version:      newVersion,
		Status:       "DRAFT",
		EffDate:      original.EffDate,
		ExpDate:      original.ExpDate,
	}

	if err := s.bomRepo.Create(ctx, newBOM); err != nil {
		return nil, err
	}

	// 复制明细
	items, err := s.bomItemRepo.ListByBOMID(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		newItem := &model.MdmBOMItem{
			BOMID:          newBOM.ID,
			LineNo:         item.LineNo,
			MaterialID:     item.MaterialID,
			MaterialCode:   item.MaterialCode,
			MaterialName:   item.MaterialName,
			Quantity:       item.Quantity,
			Unit:          item.Unit,
			ScrapRate:     item.ScrapRate,
			SubstituteGroup: item.SubstituteGroup,
			IsAlternative: item.IsAlternative,
		}
		if err := s.bomItemRepo.Create(ctx, newItem); err != nil {
			return nil, err
		}
	}

	return newBOM, nil
}

// ValidateBOM 验证BOM合法性（闭环检查）
func (s *BOMService) ValidateBOM(ctx context.Context, id uint) error {
	bom, err := s.bomRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if bom.Status != "ACTIVE" {
		return errors.New("BOM状态必须为生效才能使用")
	}
	items, err := s.bomItemRepo.ListByBOMID(ctx, id)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return errors.New("BOM明细不能为空")
	}
	for _, item := range items {
		if item.Quantity <= 0 {
			return fmt.Errorf("物料%s用量必须大于0", item.MaterialCode)
		}
	}
	return nil
}
