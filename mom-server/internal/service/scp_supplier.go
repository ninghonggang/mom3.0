package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ScpSupplierService struct {
	contactRepo *repository.ScpSupplierContactRepository
	bankRepo    *repository.ScpSupplierBankRepository
}

func NewScpSupplierService(
	contactRepo *repository.ScpSupplierContactRepository,
	bankRepo *repository.ScpSupplierBankRepository,
) *ScpSupplierService {
	return &ScpSupplierService{
		contactRepo: contactRepo,
		bankRepo:    bankRepo,
	}
}

// ==================== 供应商联系人 ====================

// ListSupplierContacts 查询供应商联系人列表
func (s *ScpSupplierService) ListSupplierContacts(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpSupplierContact, int64, error) {
	return s.contactRepo.List(ctx, tenantID, query)
}

// GetSupplierContact 获取供应商联系人详情
func (s *ScpSupplierService) GetSupplierContact(ctx context.Context, id string) (*model.ScpSupplierContact, error) {
	var contactID uint
	_, err := fmt.Sscanf(id, "%d", &contactID)
	if err != nil {
		return nil, err
	}
	return s.contactRepo.GetByID(ctx, contactID)
}

// ListContactsBySupplier 查询指定供应商的联系人
func (s *ScpSupplierService) ListContactsBySupplier(ctx context.Context, supplierID int64) ([]model.ScpSupplierContact, error) {
	return s.contactRepo.ListBySupplier(ctx, supplierID)
}

// CreateSupplierContact 创建供应商联系人
func (s *ScpSupplierService) CreateSupplierContact(ctx context.Context, tenantID int64, req *model.ScpSupplierContactCreateReqVO) (*model.ScpSupplierContact, error) {
	contact := &model.ScpSupplierContact{
		TenantID:     tenantID,
		SupplierID:   req.SupplierID,
		SupplierCode: req.SupplierCode,
		SupplierName: req.SupplierName,
		Name:         req.Name,
		Gender:       req.Gender,
		Department:   req.Department,
		Position:     req.Position,
		Phone:        req.Phone,
		Mobile:       req.Mobile,
		Email:        req.Email,
		Wechat:       req.Wechat,
		QQ:           req.QQ,
		IsPrimary:    req.IsPrimary,
		IsActive:     true,
		Remark:       req.Remark,
	}

	if err := s.contactRepo.Create(ctx, contact); err != nil {
		return nil, err
	}

	return contact, nil
}

// UpdateSupplierContact 更新供应商联系人
func (s *ScpSupplierService) UpdateSupplierContact(ctx context.Context, id string, req *model.ScpSupplierContactUpdateReqVO) error {
	var contactID uint
	_, err := fmt.Sscanf(id, "%d", &contactID)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Gender != "" {
		updates["gender"] = req.Gender
	}
	if req.Department != "" {
		updates["department"] = req.Department
	}
	if req.Position != "" {
		updates["position"] = req.Position
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Mobile != "" {
		updates["mobile"] = req.Mobile
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Wechat != "" {
		updates["wechat"] = req.Wechat
	}
	if req.QQ != "" {
		updates["qq"] = req.QQ
	}
	updates["is_primary"] = req.IsPrimary
	updates["is_active"] = req.IsActive
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.contactRepo.Update(ctx, contactID, updates)
}

// DeleteSupplierContact 删除供应商联系人
func (s *ScpSupplierService) DeleteSupplierContact(ctx context.Context, id string) error {
	var contactID uint
	_, err := fmt.Sscanf(id, "%d", &contactID)
	if err != nil {
		return err
	}
	return s.contactRepo.Delete(ctx, contactID)
}

// ==================== 供应商银行账户 ====================

// ListSupplierBanks 查询供应商银行账户列表
func (s *ScpSupplierService) ListSupplierBanks(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.ScpSupplierBank, int64, error) {
	return s.bankRepo.List(ctx, tenantID, query)
}

// GetSupplierBank 获取供应商银行账户详情
func (s *ScpSupplierService) GetSupplierBank(ctx context.Context, id string) (*model.ScpSupplierBank, error) {
	var bankID uint
	_, err := fmt.Sscanf(id, "%d", &bankID)
	if err != nil {
		return nil, err
	}
	return s.bankRepo.GetByID(ctx, bankID)
}

// ListBanksBySupplier 查询指定供应商的银行账户
func (s *ScpSupplierService) ListBanksBySupplier(ctx context.Context, supplierID int64) ([]model.ScpSupplierBank, error) {
	return s.bankRepo.ListBySupplier(ctx, supplierID)
}

// CreateSupplierBank 创建供应商银行账户
func (s *ScpSupplierService) CreateSupplierBank(ctx context.Context, tenantID int64, req *model.ScpSupplierBankCreateReqVO) (*model.ScpSupplierBank, error) {
	bank := &model.ScpSupplierBank{
		TenantID:     tenantID,
		SupplierID:   req.SupplierID,
		SupplierCode: req.SupplierCode,
		SupplierName: req.SupplierName,
		BankName:     req.BankName,
		BankCode:     req.BankCode,
		BranchName:   req.BranchName,
		BranchCode:   req.BranchCode,
		AccountType:  req.AccountType,
		AccountNo:    req.AccountNo,
		AccountName:  req.AccountName,
		Currency:     req.Currency,
		IsPrimary:    req.IsPrimary,
		IsActive:     true,
		Status:       "ACTIVE",
		Remark:       req.Remark,
	}

	if bank.Currency == "" {
		bank.Currency = "CNY"
	}
	if bank.AccountType == "" {
		bank.AccountType = "GENERAL"
	}

	if err := s.bankRepo.Create(ctx, bank); err != nil {
		return nil, err
	}

	return bank, nil
}

// UpdateSupplierBank 更新供应商银行账户
func (s *ScpSupplierService) UpdateSupplierBank(ctx context.Context, id string, req *model.ScpSupplierBankUpdateReqVO) error {
	var bankID uint
	_, err := fmt.Sscanf(id, "%d", &bankID)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{}
	if req.BankName != "" {
		updates["bank_name"] = req.BankName
	}
	if req.BankCode != "" {
		updates["bank_code"] = req.BankCode
	}
	if req.BranchName != "" {
		updates["branch_name"] = req.BranchName
	}
	if req.BranchCode != "" {
		updates["branch_code"] = req.BranchCode
	}
	if req.AccountType != "" {
		updates["account_type"] = req.AccountType
	}
	if req.AccountNo != "" {
		updates["account_no"] = req.AccountNo
	}
	if req.AccountName != "" {
		updates["account_name"] = req.AccountName
	}
	if req.Currency != "" {
		updates["currency"] = req.Currency
	}
	updates["is_primary"] = req.IsPrimary
	updates["is_active"] = req.IsActive
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.bankRepo.Update(ctx, bankID, updates)
}

// DeleteSupplierBank 删除供应商银行账户
func (s *ScpSupplierService) DeleteSupplierBank(ctx context.Context, id string) error {
	var bankID uint
	_, err := fmt.Sscanf(id, "%d", &bankID)
	if err != nil {
		return err
	}
	return s.bankRepo.Delete(ctx, bankID)
}
