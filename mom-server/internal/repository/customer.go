package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) List(ctx context.Context, tenantID int64) ([]model.Customer, int64, error) {
	var list []model.Customer
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Customer{}).Where("tenant_id = ?", tenantID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *CustomerRepository) GetByID(ctx context.Context, id uint) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).First(&customer, id).Error
	return &customer, err
}

func (r *CustomerRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND code = ?", tenantID, code).First(&customer).Error
	return &customer, err
}

func (r *CustomerRepository) Create(ctx context.Context, customer *model.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

func (r *CustomerRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Customer{}).Where("id = ?", id).Updates(updates).Error
}

func (r *CustomerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Customer{}, id).Error
}
