package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ContactRepository 联系人仓库
type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) List(ctx context.Context, tenantID int64, query model.ContactQuery) ([]model.Contact, int64, error) {
	var list []model.Contact
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Contact{})
	if query.OwnerType != "" {
		db = db.Where("owner_type = ?", query.OwnerType)
	}
	if query.OwnerID > 0 {
		db = db.Where("owner_id = ?", query.OwnerID)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *ContactRepository) GetByID(ctx context.Context, id uint64) (*model.Contact, error) {
	var contact model.Contact
	err := r.db.WithContext(ctx).First(&contact, id).Error
	return &contact, err
}

func (r *ContactRepository) Create(ctx context.Context, contact *model.Contact) error {
	return r.db.WithContext(ctx).Create(contact).Error
}

func (r *ContactRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Contact{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ContactRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Contact{}, id).Error
}

func (r *ContactRepository) GetByOwner(ctx context.Context, ownerType string, ownerID uint64) ([]model.Contact, error) {
	var list []model.Contact
	err := r.db.WithContext(ctx).Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).Order("is_primary DESC, id DESC").Find(&list).Error
	return list, err
}

// BankAccountRepository 银行账户仓库
type BankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *BankAccountRepository {
	return &BankAccountRepository{db: db}
}

func (r *BankAccountRepository) List(ctx context.Context, tenantID int64, query model.BankAccountQuery) ([]model.BankAccount, int64, error) {
	var list []model.BankAccount
	var total int64

	db := r.db.WithContext(ctx).Model(&model.BankAccount{})
	if query.OwnerType != "" {
		db = db.Where("owner_type = ?", query.OwnerType)
	}
	if query.OwnerID > 0 {
		db = db.Where("owner_id = ?", query.OwnerID)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *BankAccountRepository) GetByID(ctx context.Context, id uint64) (*model.BankAccount, error) {
	var account model.BankAccount
	err := r.db.WithContext(ctx).First(&account, id).Error
	return &account, err
}

func (r *BankAccountRepository) Create(ctx context.Context, account *model.BankAccount) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *BankAccountRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.BankAccount{}).Where("id = ?", id).Updates(updates).Error
}

func (r *BankAccountRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.BankAccount{}, id).Error
}

func (r *BankAccountRepository) GetByOwner(ctx context.Context, ownerType string, ownerID uint64) ([]model.BankAccount, error) {
	var list []model.BankAccount
	err := r.db.WithContext(ctx).Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).Order("is_primary DESC, id DESC").Find(&list).Error
	return list, err
}

// AttachmentRepository 附件仓库
type AttachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) *AttachmentRepository {
	return &AttachmentRepository{db: db}
}

func (r *AttachmentRepository) List(ctx context.Context, tenantID int64, query model.AttachmentQuery) ([]model.Attachment, int64, error) {
	var list []model.Attachment
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Attachment{})
	if query.OwnerType != "" {
		db = db.Where("owner_type = ?", query.OwnerType)
	}
	if query.OwnerID > 0 {
		db = db.Where("owner_id = ?", query.OwnerID)
	}
	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *AttachmentRepository) GetByID(ctx context.Context, id uint64) (*model.Attachment, error) {
	var attachment model.Attachment
	err := r.db.WithContext(ctx).First(&attachment, id).Error
	return &attachment, err
}

func (r *AttachmentRepository) Create(ctx context.Context, attachment *model.Attachment) error {
	return r.db.WithContext(ctx).Create(attachment).Error
}

func (r *AttachmentRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Attachment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AttachmentRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Attachment{}, id).Error
}

func (r *AttachmentRepository) GetByOwner(ctx context.Context, ownerType string, ownerID uint64) ([]model.Attachment, error) {
	var list []model.Attachment
	err := r.db.WithContext(ctx).Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).Order("id DESC").Find(&list).Error
	return list, err
}
