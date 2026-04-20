package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ContactService 联系人服务
type ContactService struct {
	repo *repository.ContactRepository
}

func NewContactService(repo *repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) List(ctx context.Context, query model.ContactQuery) ([]model.Contact, int64, error) {
	return s.repo.List(ctx, 0, query)
}

func (s *ContactService) GetByID(ctx context.Context, id uint64) (*model.Contact, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ContactService) Create(ctx context.Context, req *model.ContactCreateRequest) error {
	contact := &model.Contact{
		OwnerType:  req.OwnerType,
		OwnerID:    req.OwnerID,
		Name:       req.Name,
		Gender:     req.Gender,
		Phone:      req.Phone,
		Mobile:     req.Mobile,
		Email:      req.Email,
		Department: req.Department,
		Position:   req.Position,
		IsPrimary:  req.IsPrimary,
		Remark:     req.Remark,
	}
	return s.repo.Create(ctx, contact)
}

func (s *ContactService) Update(ctx context.Context, id uint64, req *model.ContactUpdateRequest) error {
	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Gender != "" {
		updates["gender"] = req.Gender
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
	if req.Department != "" {
		updates["department"] = req.Department
	}
	if req.Position != "" {
		updates["position"] = req.Position
	}
	updates["is_primary"] = req.IsPrimary
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *ContactService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *ContactService) GetByOwner(ctx context.Context, ownerType string, ownerID uint64) ([]model.Contact, error) {
	return s.repo.GetByOwner(ctx, ownerType, ownerID)
}

// BankAccountService 银行账户服务
type BankAccountService struct {
	repo *repository.BankAccountRepository
}

func NewBankAccountService(repo *repository.BankAccountRepository) *BankAccountService {
	return &BankAccountService{repo: repo}
}

func (s *BankAccountService) List(ctx context.Context, query model.BankAccountQuery) ([]model.BankAccount, int64, error) {
	return s.repo.List(ctx, 0, query)
}

func (s *BankAccountService) GetByID(ctx context.Context, id uint64) (*model.BankAccount, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BankAccountService) Create(ctx context.Context, req *model.BankAccountCreateRequest) error {
	account := &model.BankAccount{
		OwnerType:   req.OwnerType,
		OwnerID:     req.OwnerID,
		BankName:    req.BankName,
		BankAccount: req.BankAccount,
		AccountName: req.AccountName,
		BranchName:  req.BranchName,
		Currency:    req.Currency,
		IsPrimary:   req.IsPrimary,
		Status:      req.Status,
		Remark:      req.Remark,
	}
	if account.Currency == "" {
		account.Currency = "CNY"
	}
	if account.Status == "" {
		account.Status = "ACTIVE"
	}
	return s.repo.Create(ctx, account)
}

func (s *BankAccountService) Update(ctx context.Context, id uint64, req *model.BankAccountUpdateRequest) error {
	updates := map[string]interface{}{}
	if req.BankName != "" {
		updates["bank_name"] = req.BankName
	}
	if req.BankAccount != "" {
		updates["bank_account"] = req.BankAccount
	}
	if req.AccountName != "" {
		updates["account_name"] = req.AccountName
	}
	if req.BranchName != "" {
		updates["branch_name"] = req.BranchName
	}
	if req.Currency != "" {
		updates["currency"] = req.Currency
	}
	updates["is_primary"] = req.IsPrimary
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *BankAccountService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *BankAccountService) GetByOwner(ctx context.Context, ownerType string, ownerID uint64) ([]model.BankAccount, error) {
	return s.repo.GetByOwner(ctx, ownerType, ownerID)
}

// AttachmentService 附件服务
type AttachmentService struct {
	repo *repository.AttachmentRepository
}

func NewAttachmentService(repo *repository.AttachmentRepository) *AttachmentService {
	return &AttachmentService{repo: repo}
}

func (s *AttachmentService) List(ctx context.Context, query model.AttachmentQuery) ([]model.Attachment, int64, error) {
	return s.repo.List(ctx, 0, query)
}

func (s *AttachmentService) GetByID(ctx context.Context, id uint64) (*model.Attachment, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AttachmentService) Create(ctx context.Context, req *model.AttachmentCreateRequest) error {
	attachment := &model.Attachment{
		OwnerType:   req.OwnerType,
		OwnerID:     req.OwnerID,
		FileName:    req.FileName,
		FilePath:    req.FilePath,
		FileSize:    req.FileSize,
		FileType:    req.FileType,
		Category:    req.Category,
		Description: req.Description,
	}
	return s.repo.Create(ctx, attachment)
}

func (s *AttachmentService) Update(ctx context.Context, id uint64, req *model.AttachmentUpdateRequest) error {
	updates := map[string]interface{}{}
	if req.FileName != "" {
		updates["file_name"] = req.FileName
	}
	if req.FilePath != "" {
		updates["file_path"] = req.FilePath
	}
	if req.FileSize > 0 {
		updates["file_size"] = req.FileSize
	}
	if req.FileType != "" {
		updates["file_type"] = req.FileType
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *AttachmentService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *AttachmentService) GetByOwner(ctx context.Context, ownerType string, ownerID uint64) ([]model.Attachment, error) {
	return s.repo.GetByOwner(ctx, ownerType, ownerID)
}
