package service

import (
	"context"
	"errors"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type VmiService struct {
	repo *repository.VmiRepository
}

func NewVmiService(repo *repository.VmiRepository) *VmiService {
	return &VmiService{repo: repo}
}

// Vendor operations
func (s *VmiService) ListVendors(ctx context.Context, tenantID int64, query *model.VmiVendorQuery) ([]model.VmiVendor, int64, error) {
	return s.repo.ListVendors(ctx, tenantID, query)
}

func (s *VmiService) GetVendorByID(ctx context.Context, id int64) (*model.VmiVendor, error) {
	return s.repo.GetVendorByID(ctx, id)
}

func (s *VmiService) CreateVendor(ctx context.Context, tenantID int64, req *model.VmiVendorCreateReq) (*model.VmiVendor, error) {
	vendor := &model.VmiVendor{
		TenantID:       tenantID,
		VendorID:      req.VendorID,
		VendorCode:    req.VendorCode,
		VendorName:    req.VendorName,
		WarehouseID:   req.WarehouseID,
		WarehouseName: req.WarehouseName,
		Contact:       req.Contact,
		Phone:         req.Phone,
		MinStock:      req.MinStock,
		MaxStock:      req.MaxStock,
		ReplenishCycle: req.ReplenishCycle,
		IsActive:      1,
		Status:        1,
		Remarks:       req.Remarks,
	}

	if err := s.repo.CreateVendor(ctx, vendor); err != nil {
		return nil, err
	}
	return vendor, nil
}

func (s *VmiService) UpdateVendor(ctx context.Context, id int64, req *model.VmiVendorCreateReq) error {
	vendor, err := s.repo.GetVendorByID(ctx, id)
	if err != nil {
		return err
	}
	if vendor == nil {
		return errors.New("vendor not found")
	}

	vendor.WarehouseID = req.WarehouseID
	vendor.WarehouseName = req.WarehouseName
	vendor.Contact = req.Contact
	vendor.Phone = req.Phone
	vendor.MinStock = req.MinStock
	vendor.MaxStock = req.MaxStock
	vendor.ReplenishCycle = req.ReplenishCycle
	vendor.Remarks = req.Remarks

	return s.repo.UpdateVendor(ctx, vendor)
}

func (s *VmiService) DeleteVendor(ctx context.Context, id int64) error {
	return s.repo.DeleteVendor(ctx, id)
}

// Material operations
func (s *VmiService) ListMaterials(ctx context.Context, tenantID int64, query *model.VmiMaterialQuery) ([]model.VmiMaterial, int64, error) {
	return s.repo.ListMaterials(ctx, tenantID, query)
}

func (s *VmiService) GetMaterialByID(ctx context.Context, id int64) (*model.VmiMaterial, error) {
	return s.repo.GetMaterialByID(ctx, id)
}

func (s *VmiService) CreateMaterial(ctx context.Context, tenantID int64, material *model.VmiMaterial) error {
	material.TenantID = tenantID
	material.CurrentStock = 0
	material.AvailableStock = 0
	material.ConsumeQty = 0
	return s.repo.CreateMaterial(ctx, material)
}

func (s *VmiService) UpdateMaterial(ctx context.Context, material *model.VmiMaterial) error {
	return s.repo.UpdateMaterial(ctx, material)
}

func (s *VmiService) DeleteMaterial(ctx context.Context, id int64) error {
	return s.repo.DeleteMaterial(ctx, id)
}

// Transaction operations
func (s *VmiService) ListTransactions(ctx context.Context, tenantID int64, query *model.VmiTransactionQuery) ([]model.VmiTransaction, int64, error) {
	return s.repo.ListTransactions(ctx, tenantID, query)
}

// Consume consumption - when material is consumed from VMI stock
func (s *VmiService) Consume(ctx context.Context, tenantID int64, operatorID int64, operatorName string, req *model.VmiConsumeReq) error {
	material, err := s.repo.GetMaterialByID(ctx, req.MaterialID)
	if err != nil {
		return err
	}
	if material == nil {
		return errors.New("material not found")
	}

	if material.CurrentStock < req.Quantity {
		return errors.New("insufficient stock")
	}

	// Update stock
	beforeQty := material.CurrentStock
	material.CurrentStock -= req.Quantity
	material.AvailableStock -= req.Quantity
	material.ConsumeQty += req.Quantity
	now := time.Now()
	material.LastConsumeDate = &now

	if err := s.repo.UpdateMaterialStock(ctx, req.MaterialID, material.CurrentStock, material.AvailableStock); err != nil {
		return err
	}

	// Record transaction
	tx := &model.VmiTransaction{
		TenantID:         tenantID,
		VendorID:        req.VendorID,
		MaterialID:      req.MaterialID,
		TransactionType: 2, // 消耗
		Quantity:        req.Quantity,
		BeforeQty:       beforeQty,
		AfterQty:        material.CurrentStock,
		ReferenceNo:     req.ReferenceNo,
		Remarks:         req.Remarks,
		OperatorID:      operatorID,
		OperatorName:    operatorName,
	}

	return s.repo.CreateTransaction(ctx, tx)
}

// Replenish - when supplier replenishes VMI stock
func (s *VmiService) Replenish(ctx context.Context, tenantID int64, operatorID int64, operatorName string, req *model.VmiReplenishReq) error {
	material, err := s.repo.GetMaterialByID(ctx, req.MaterialID)
	if err != nil {
		return err
	}
	if material == nil {
		return errors.New("material not found")
	}

	// Update stock
	beforeQty := material.CurrentStock
	material.CurrentStock += req.Quantity
	material.AvailableStock += req.Quantity

	if err := s.repo.UpdateMaterialStock(ctx, req.MaterialID, material.CurrentStock, material.AvailableStock); err != nil {
		return err
	}

	// Record transaction
	tx := &model.VmiTransaction{
		TenantID:         tenantID,
		VendorID:        req.VendorID,
		MaterialID:      req.MaterialID,
		TransactionType: 1, // 入库
		Quantity:        req.Quantity,
		BeforeQty:       beforeQty,
		AfterQty:        material.CurrentStock,
		Remarks:         req.Remarks,
		OperatorID:      operatorID,
		OperatorName:    operatorName,
	}

	return s.repo.CreateTransaction(ctx, tx)
}