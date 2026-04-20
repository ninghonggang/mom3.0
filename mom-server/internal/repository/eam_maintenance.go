package repository

import (
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// EquipmentSpareRepository 备件Repository
type EquipmentSpareRepository struct {
	db *gorm.DB
}

// NewEquipmentSpareRepository 创建备件Repository
func NewEquipmentSpareRepository(db *gorm.DB) *EquipmentSpareRepository {
	return &EquipmentSpareRepository{db: db}
}

// Create 创建备件
func (r *EquipmentSpareRepository) Create(spare *model.EquipmentSpare) error {
	return r.db.Create(spare).Error
}

// Update 更新备件
func (r *EquipmentSpareRepository) Update(spare *model.EquipmentSpare) error {
	return r.db.Save(spare).Error
}

// Delete 删除备件
func (r *EquipmentSpareRepository) Delete(id uint) error {
	return r.db.Delete(&model.EquipmentSpare{}, id).Error
}

// GetByID 根据ID获取备件
func (r *EquipmentSpareRepository) GetByID(id uint) (*model.EquipmentSpare, error) {
	var spare model.EquipmentSpare
	err := r.db.First(&spare, id).Error
	if err != nil {
		return nil, err
	}
	return &spare, nil
}

// GetByCode 根据编码获取备件
func (r *EquipmentSpareRepository) GetByCode(code string) (*model.EquipmentSpare, error) {
	var spare model.EquipmentSpare
	err := r.db.Where("spare_code = ?", code).First(&spare).Error
	if err != nil {
		return nil, err
	}
	return &spare, nil
}

// ListPage 分页查询备件
func (r *EquipmentSpareRepository) ListPage(tenantID int64, query map[string]interface{}, page, pageSize int) ([]model.EquipmentSpare, int64, error) {
	var spares []model.EquipmentSpare
	var total int64

	db := r.db.Model(&model.EquipmentSpare{}).Where("tenant_id = ?", tenantID)

	if keyword, ok := query["keyword"]; ok && keyword != "" {
		db = db.Where("spare_code LIKE ? OR spare_name LIKE ?", "%"+keyword.(string)+"%", "%"+keyword.(string)+"%")
	}
	if category, ok := query["category"]; ok && category != "" {
		db = db.Where("category = ?", category)
	}
	if status, ok := query["status"]; ok && status != "" {
		db = db.Where("status = ?", status)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&spares).Error
	return spares, total, err
}

// ListAll 获取所有备件
func (r *EquipmentSpareRepository) ListAll(tenantID int64) ([]model.EquipmentSpare, error) {
	var spares []model.EquipmentSpare
	err := r.db.Where("tenant_id = ?", tenantID).Order("id DESC").Find(&spares).Error
	return spares, err
}

// UpdateQuantity 更新库存
func (r *EquipmentSpareRepository) UpdateQuantity(id uint, quantity float64) error {
	return r.db.Model(&model.EquipmentSpare{}).Where("id = ?", id).Update("quantity", quantity).Error
}

// EquipmentSpareTransactionRepository 备件事务Repository
type EquipmentSpareTransactionRepository struct {
	db *gorm.DB
}

// NewEquipmentSpareTransactionRepository 创建备件事务Repository
func NewEquipmentSpareTransactionRepository(db *gorm.DB) *EquipmentSpareTransactionRepository {
	return &EquipmentSpareTransactionRepository{db: db}
}

// Create 创建备件事务
func (r *EquipmentSpareTransactionRepository) Create(tx *model.EquipmentSpareTransaction) error {
	return r.db.Create(tx).Error
}

// GetBySpareID 根据备件ID查询事务
func (r *EquipmentSpareTransactionRepository) GetBySpareID(spareID uint) ([]model.EquipmentSpareTransaction, error) {
	var txs []model.EquipmentSpareTransaction
	err := r.db.Where("spare_id = ?", spareID).Order("id DESC").Find(&txs).Error
	return txs, err
}

// ListPage 分页查询事务
func (r *EquipmentSpareTransactionRepository) ListPage(tenantID uint, query map[string]interface{}, page, pageSize int) ([]model.EquipmentSpareTransaction, int64, error) {
	var txs []model.EquipmentSpareTransaction
	var total int64

	db := r.db.Model(&model.EquipmentSpareTransaction{}).Where("tenant_id = ?", tenantID)

	if spareID, ok := query["spare_id"]; ok {
		db = db.Where("spare_id = ?", spareID)
	}
	if txType, ok := query["transaction_type"]; ok && txType != "" {
		db = db.Where("transaction_type = ?", txType)
	}
	if orderNo, ok := query["order_no"]; ok && orderNo != "" {
		db = db.Where("order_no = ?", orderNo)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&txs).Error
	return txs, total, err
}
