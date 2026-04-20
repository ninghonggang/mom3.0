package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// DeliveryAddressRepository 收货地址仓库
type DeliveryAddressRepository struct {
	db *gorm.DB
}

func NewDeliveryAddressRepository(db *gorm.DB) *DeliveryAddressRepository {
	return &DeliveryAddressRepository{db: db}
}

func (r *DeliveryAddressRepository) List(ctx context.Context, tenantID int64, query model.DeliveryAddressQuery) ([]model.DeliveryAddress, int64, error) {
	var list []model.DeliveryAddress
	var total int64

	db := r.db.WithContext(ctx).Model(&model.DeliveryAddress{})
	if query.CustomerID > 0 {
		db = db.Where("customer_id = ?", query.CustomerID)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Order("is_default DESC, id DESC").Find(&list).Error
	return list, total, err
}

func (r *DeliveryAddressRepository) GetByID(ctx context.Context, id uint64) (*model.DeliveryAddress, error) {
	var address model.DeliveryAddress
	err := r.db.WithContext(ctx).First(&address, id).Error
	return &address, err
}

func (r *DeliveryAddressRepository) Create(ctx context.Context, address *model.DeliveryAddress) error {
	return r.db.WithContext(ctx).Create(address).Error
}

func (r *DeliveryAddressRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DeliveryAddress{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DeliveryAddressRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.DeliveryAddress{}, id).Error
}

func (r *DeliveryAddressRepository) GetByCustomer(ctx context.Context, customerID uint64) ([]model.DeliveryAddress, error) {
	var list []model.DeliveryAddress
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Order("is_default DESC, id DESC").Find(&list).Error
	return list, err
}

func (r *DeliveryAddressRepository) SetDefault(ctx context.Context, customerID uint64, addressID uint64) error {
	tx := r.db.WithContext(ctx)
	// 取消该客户的所有默认地址
	if err := tx.Model(&model.DeliveryAddress{}).Where("customer_id = ?", customerID).Update("is_default", false).Error; err != nil {
		return err
	}
	// 设置新的默认地址
	return tx.Model(&model.DeliveryAddress{}).Where("id = ?", addressID).Update("is_default", true).Error
}
