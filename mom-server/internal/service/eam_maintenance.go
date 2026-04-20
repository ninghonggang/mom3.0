package service

import (
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// EquipmentSpareService 备件Service
type EquipmentSpareService struct {
	spareRepo *repository.EquipmentSpareRepository
	txRepo    *repository.EquipmentSpareTransactionRepository
}

// NewEquipmentSpareService 创建备件Service
func NewEquipmentSpareService(spareRepo *repository.EquipmentSpareRepository, txRepo *repository.EquipmentSpareTransactionRepository) *EquipmentSpareService {
	return &EquipmentSpareService{
		spareRepo: spareRepo,
		txRepo:    txRepo,
	}
}

// Create 创建备件
func (s *EquipmentSpareService) Create(tenantID int64, req *model.SpareCreateReq) error {
	spare := &model.EquipmentSpare{
		TenantID:      tenantID,
		SpareCode:     req.SpareCode,
		SpareName:     req.SpareName,
		Category:      req.Category,
		Specification: req.Specification,
		Unit:          req.Unit,
		Quantity:      req.Quantity,
		MinQuantity:   req.MinQuantity,
		MaxQuantity:   req.MaxQuantity,
		Location:      req.Location,
		UnitPrice:     req.UnitPrice,
		Status:        "AVAILABLE",
		Remark:        req.Remark,
	}
	return s.spareRepo.Create(spare)
}

// Update 更新备件
func (s *EquipmentSpareService) Update(tenantID int64, req *model.SpareUpdateReq) error {
	spare, err := s.spareRepo.GetByID(req.ID)
	if err != nil {
		return err
	}
	if spare.TenantID != tenantID {
		return nil
	}
	spare.SpareName = req.SpareName
	spare.Category = req.Category
	spare.Specification = req.Specification
	spare.Unit = req.Unit
	spare.MinQuantity = req.MinQuantity
	spare.MaxQuantity = req.MaxQuantity
	spare.Location = req.Location
	spare.UnitPrice = req.UnitPrice
	spare.Remark = req.Remark
	return s.spareRepo.Update(spare)
}

// Delete 删除备件
func (s *EquipmentSpareService) Delete(tenantID int64, id uint) error {
	spare, err := s.spareRepo.GetByID(id)
	if err != nil {
		return err
	}
	if spare.TenantID != tenantID {
		return nil
	}
	return s.spareRepo.Delete(id)
}

// GetByID 获取备件
func (s *EquipmentSpareService) GetByID(tenantID int64, id uint) (*model.EquipmentSpare, error) {
	spare, err := s.spareRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if spare.TenantID != tenantID {
		return nil, nil
	}
	return spare, nil
}

// ListPage 分页查询
func (s *EquipmentSpareService) ListPage(tenantID int64, query map[string]interface{}, page, pageSize int) ([]model.EquipmentSpare, int64, error) {
	return s.spareRepo.ListPage(tenantID, query, page, pageSize)
}

// ListAll 获取所有
func (s *EquipmentSpareService) ListAll(tenantID int64) ([]model.EquipmentSpare, error) {
	return s.spareRepo.ListAll(tenantID)
}

// In 入库
func (s *EquipmentSpareService) In(tenantID int64, handlerID uint, handlerName string, req *model.SpareTransactionReq) error {
	spare, err := s.spareRepo.GetByID(req.SpareID)
	if err != nil {
		return err
	}
	if spare.TenantID != tenantID {
		return nil
	}

	beforeQty := spare.Quantity
	spare.Quantity += req.Quantity
	afterQty := spare.Quantity

	err = s.spareRepo.Update(spare)
	if err != nil {
		return err
	}

	tx := &model.EquipmentSpareTransaction{
		TenantID:        tenantID,
		SpareID:         req.SpareID,
		SpareCode:       spare.SpareCode,
		TransactionType: "IN",
		Quantity:        req.Quantity,
		BeforeQty:       beforeQty,
		AfterQty:        afterQty,
		OrderNo:         req.OrderNo,
		HandlerID:       handlerID,
		HandlerName:     handlerName,
		Remark:          req.Remark,
	}
	return s.txRepo.Create(tx)
}

// Out 出库
func (s *EquipmentSpareService) Out(tenantID int64, handlerID uint, handlerName string, req *model.SpareTransactionReq) error {
	spare, err := s.spareRepo.GetByID(req.SpareID)
	if err != nil {
		return err
	}
	if spare.TenantID != tenantID {
		return nil
	}
	if spare.Quantity < req.Quantity {
		return nil
	}

	beforeQty := spare.Quantity
	spare.Quantity -= req.Quantity
	afterQty := spare.Quantity

	err = s.spareRepo.Update(spare)
	if err != nil {
		return err
	}

	tx := &model.EquipmentSpareTransaction{
		TenantID:        tenantID,
		SpareID:         req.SpareID,
		SpareCode:       spare.SpareCode,
		TransactionType: "OUT",
		Quantity:        req.Quantity,
		BeforeQty:       beforeQty,
		AfterQty:        afterQty,
		OrderNo:         req.OrderNo,
		HandlerID:       handlerID,
		HandlerName:     handlerName,
		Remark:          req.Remark,
	}
	return s.txRepo.Create(tx)
}

// GetTransactions 获取事务记录
func (s *EquipmentSpareService) GetTransactions(tenantID int64, spareID uint) ([]model.EquipmentSpareTransaction, error) {
	return s.txRepo.GetBySpareID(spareID)
}
