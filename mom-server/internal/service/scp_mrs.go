package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ScpMRSService struct {
	mrsRepo *repository.ScpMRSRepository
}

func NewScpMRSService(mrsRepo *repository.ScpMRSRepository) *ScpMRSService {
	return &ScpMRSService{mrsRepo: mrsRepo}
}

// ListMRS 查询MRS列表
func (s *ScpMRSService) ListMRS(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpMRS, int64, error) {
	return s.mrsRepo.List(ctx, tenantID, query)
}

// GetMRS 获取MRS详情
func (s *ScpMRSService) GetMRS(ctx context.Context, id string) (*model.ScpMRS, error) {
	var mrsID uint
	_, err := fmt.Sscanf(id, "%d", &mrsID)
	if err != nil {
		return nil, err
	}
	return s.mrsRepo.GetByID(ctx, mrsID)
}

// CreateMRS 创建MRS
func (s *ScpMRSService) CreateMRS(ctx context.Context, tenantID int64, req *model.ScpMRSCreateReqVO) (*model.ScpMRS, error) {
	// 生成MRS编号
	mrsNo := generateMrsNo(tenantID)

	mrs := &model.ScpMRS{
		TenantID:   tenantID,
		MrsNo:      mrsNo,
		PlanMonth:  req.PlanMonth,
		SourceType: req.SourceType,
		SourceNo:   req.SourceNo,
		Remark:     req.Remark,
		Status:     "DRAFT",
	}

	// 处理明细
	for i, itemReq := range req.Items {
		item := model.ScpMRSItem{
			MrsNo:        mrsNo,
			MaterialID:   itemReq.MaterialID,
			MaterialCode: itemReq.MaterialCode,
			MaterialName: itemReq.MaterialName,
			Spec:         itemReq.Spec,
			Unit:         itemReq.Unit,
			ReqQty:       itemReq.ReqQty,
			OnHandQty:    itemReq.OnHandQty,
			ShortQty:     itemReq.ShortQty,
			SupplierID:   itemReq.SupplierID,
			SupplierName: itemReq.SupplierName,
			Status:       "PENDING",
		}
		if itemReq.PromiseDate != "" {
			t, _ := time.Parse("2006-01-02", itemReq.PromiseDate)
			item.PromiseDate = &t
		}
		mrs.Items = append(mrs.Items, item)
		mrs.TotalQty += itemReq.ReqQty
		mrs.TotalItems = i + 1
	}

	if err := s.mrsRepo.CreateWithItems(ctx, mrs); err != nil {
		return nil, err
	}

	return mrs, nil
}

// UpdateMRS 更新MRS
func (s *ScpMRSService) UpdateMRS(ctx context.Context, tenantID int64, id string, req *model.ScpMRSUpdateReqVO) (*model.ScpMRS, error) {
	var mrsID uint
	_, err := fmt.Sscanf(id, "%d", &mrsID)
	if err != nil {
		return nil, err
	}

	// 获取现有MRS
	existing, err := s.mrsRepo.GetByID(ctx, mrsID)
	if err != nil {
		return nil, err
	}

	if existing.Status != "DRAFT" {
		return nil, fmt.Errorf("只有草稿状态的MRS可以更新")
	}

	mrs := &model.ScpMRS{
		PlanMonth:  req.PlanMonth,
		SourceType: req.SourceType,
		SourceNo:   req.SourceNo,
		Remark:     req.Remark,
	}

	// 处理明细
	var totalQty float64
	for _, itemReq := range req.Items {
		item := model.ScpMRSItem{
			MrsID:        int64(mrsID),
			MrsNo:        existing.MrsNo,
			MaterialID:   itemReq.MaterialID,
			MaterialCode: itemReq.MaterialCode,
			MaterialName: itemReq.MaterialName,
			Spec:         itemReq.Spec,
			Unit:         itemReq.Unit,
			ReqQty:       itemReq.ReqQty,
			OnHandQty:    itemReq.OnHandQty,
			ShortQty:     itemReq.ShortQty,
			SupplierID:   itemReq.SupplierID,
			SupplierName: itemReq.SupplierName,
			Status:       "PENDING",
		}
		if itemReq.PromiseDate != "" {
			t, _ := time.Parse("2006-01-02", itemReq.PromiseDate)
			item.PromiseDate = &t
		}
		mrs.Items = append(mrs.Items, item)
		totalQty += itemReq.ReqQty
	}
	mrs.TotalItems = len(req.Items)
	mrs.TotalQty = totalQty

	if err := s.mrsRepo.UpdateWithItems(ctx, mrsID, mrs, nil); err != nil {
		return nil, err
	}

	return s.mrsRepo.GetByID(ctx, mrsID)
}

// DeleteMRS 删除MRS
func (s *ScpMRSService) DeleteMRS(ctx context.Context, id string) error {
	var mrsID uint
	_, err := fmt.Sscanf(id, "%d", &mrsID)
	if err != nil {
		return err
	}

	// 获取现有MRS
	existing, err := s.mrsRepo.GetByID(ctx, mrsID)
	if err != nil {
		return err
	}

	if existing.Status != "DRAFT" {
		return fmt.Errorf("只有草稿状态的MRS可以删除")
	}

	return s.mrsRepo.Delete(ctx, mrsID)
}

// PublishMRS 发布MRS
func (s *ScpMRSService) PublishMRS(ctx context.Context, id string, userID int64) error {
	var mrsID uint
	_, err := fmt.Sscanf(id, "%d", &mrsID)
	if err != nil {
		return err
	}

	existing, err := s.mrsRepo.GetByID(ctx, mrsID)
	if err != nil {
		return err
	}

	if existing.Status != "DRAFT" {
		return fmt.Errorf("只有草稿状态的MRS可以发布")
	}

	now := time.Now()
	return s.mrsRepo.Update(ctx, mrsID, map[string]interface{}{
		"status":       "PUBLISHED",
		"published_by": userID,
		"published_at":  now,
	})
}

// CloseMRS 关闭MRS
func (s *ScpMRSService) CloseMRS(ctx context.Context, id string) error {
	var mrsID uint
	_, err := fmt.Sscanf(id, "%d", &mrsID)
	if err != nil {
		return err
	}

	existing, err := s.mrsRepo.GetByID(ctx, mrsID)
	if err != nil {
		return err
	}

	if existing.Status != "PUBLISHED" {
		return fmt.Errorf("只有已发布状态的MRS可以关闭")
	}

	return s.mrsRepo.UpdateStatus(ctx, mrsID, "CLOSED")
}

// GetItems 获取MRS明细
func (s *ScpMRSService) GetItems(ctx context.Context, mrsID string) ([]model.ScpMRSItem, error) {
	var id uint
	_, err := fmt.Sscanf(mrsID, "%d", &id)
	if err != nil {
		return nil, err
	}
	return s.mrsRepo.GetItems(ctx, id)
}

// SyncMRS 从外部系统同步MRS
func (s *ScpMRSService) SyncMRS(ctx context.Context, tenantID int64, req *model.ScpMRSSyncReqVO) (*model.ScpMRS, error) {
	// 检查是否已存在
	existing, _ := s.mrsRepo.GetByNo(ctx, tenantID, req.SourceNo)
	if existing != nil {
		return nil, fmt.Errorf("MRS单号 %s 已存在", req.SourceNo)
	}

	mrsNo := generateMrsNo(tenantID)
	mrs := &model.ScpMRS{
		TenantID:   tenantID,
		MrsNo:      mrsNo,
		SourceType: "APS",
		SourceNo:   req.SourceNo,
		Status:     "DRAFT",
	}

	var totalQty float64
	for i, itemReq := range req.Data {
		item := model.ScpMRSItem{
			MrsNo:        mrsNo,
			MaterialCode: itemReq.MaterialCode,
			MaterialName: itemReq.MaterialName,
			Spec:         itemReq.Spec,
			Unit:         itemReq.Unit,
			ReqQty:       itemReq.ReqQty,
			OnHandQty:    itemReq.OnHandQty,
			ShortQty:     itemReq.ShortQty,
			SupplierName: itemReq.SupplierName,
			Status:       "PENDING",
		}
		if itemReq.PromiseDate != "" {
			t, _ := time.Parse("2006-01-02", itemReq.PromiseDate)
			item.PromiseDate = &t
		}
		mrs.Items = append(mrs.Items, item)
		totalQty += itemReq.ReqQty
		mrs.TotalItems = i + 1
	}
	mrs.TotalQty = totalQty

	if err := s.mrsRepo.CreateWithItems(ctx, mrs); err != nil {
		return nil, err
	}

	return mrs, nil
}

// generateMrsNo 生成MRS编号
func generateMrsNo(tenantID int64) string {
	return fmt.Sprintf("MRS-%s-%d", time.Now().Format("20060102"), tenantID)
}
