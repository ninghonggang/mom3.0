package repository

import (
	"context"
	"errors"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type CustomerCreditRepository struct {
	db *gorm.DB
}

func NewCustomerCreditRepository(db *gorm.DB) *CustomerCreditRepository {
	return &CustomerCreditRepository{db: db}
}

func (r *CustomerCreditRepository) List(ctx context.Context, tenantID int64, query *model.CustomerCreditQuery) ([]model.CustomerCredit, int64, error) {
	var list []model.CustomerCredit
	var total int64

	db := r.db.WithContext(ctx).Model(&model.CustomerCredit{}).Where("tenant_id = ?", tenantID)

	if query.CustomerCode != "" {
		db = db.Where("customer_code LIKE ?", "%"+query.CustomerCode+"%")
	}
	if query.CustomerName != "" {
		db = db.Where("customer_name LIKE ?", "%"+query.CustomerName+"%")
	}
	if query.CreditLevel != "" {
		db = db.Where("credit_level = ?", query.CreditLevel)
	}
	if query.RiskLevel > 0 {
		db = db.Where("risk_level = ?", query.RiskLevel)
	}
	if query.Blacklist > 0 {
		db = db.Where("blacklist = ?", query.Blacklist)
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}
	offset := (query.Page - 1) * query.PageSize

	if err := db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *CustomerCreditRepository) GetByID(ctx context.Context, id int64) (*model.CustomerCredit, error) {
	var credit model.CustomerCredit
	if err := r.db.WithContext(ctx).First(&credit, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &credit, nil
}

func (r *CustomerCreditRepository) GetByCustomerID(ctx context.Context, customerID int64) (*model.CustomerCredit, error) {
	var credit model.CustomerCredit
	if err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).First(&credit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &credit, nil
}

func (r *CustomerCreditRepository) Create(ctx context.Context, tenantID int64, credit *model.CustomerCredit) error {
	credit.TenantID = tenantID
	return r.db.WithContext(ctx).Create(credit).Error
}

func (r *CustomerCreditRepository) Update(ctx context.Context, id int64, credit *model.CustomerCredit) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Updates(credit).Error
}

func (r *CustomerCreditRepository) UpdateUsedCredit(ctx context.Context, id int64, usedCredit, availableCredit float64) error {
	return r.db.WithContext(ctx).Model(&model.CustomerCredit{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"used_credit":      usedCredit,
			"available_credit": availableCredit,
		}).Error
}

func (r *CustomerCreditRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.CustomerCredit{}, id).Error
}