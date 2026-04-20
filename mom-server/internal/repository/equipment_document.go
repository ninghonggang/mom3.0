package repository

import (
	"context"

	"mom-server/internal/model"
	"gorm.io/gorm"
)

// EquipmentDocumentRepository 设备文档仓储
type EquipmentDocumentRepository struct {
	db *gorm.DB
}

func NewEquipmentDocumentRepository(db *gorm.DB) *EquipmentDocumentRepository {
	return &EquipmentDocumentRepository{db: db}
}

func (r *EquipmentDocumentRepository) List(ctx context.Context, tenantID int64, query map[string]any) ([]model.EquipmentDocument, int64, error) {
	var list []model.EquipmentDocument
	var total int64

	q := r.db.WithContext(ctx).Model(&model.EquipmentDocument{}).Where("tenant_id = ?", tenantID)

	if equipmentID, ok := query["equipment_id"]; ok && equipmentID.(uint) > 0 {
		q = q.Where("equipment_id = ?", equipmentID)
	}
	if docType, ok := query["doc_type"]; ok && docType.(string) != "" {
		q = q.Where("doc_type = ?", docType)
	}
	if status, ok := query["status"]; ok && status.(int) > 0 {
		q = q.Where("status = ?", status)
	}

	q.Count(&total)
	err := q.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *EquipmentDocumentRepository) GetByID(ctx context.Context, id uint) (*model.EquipmentDocument, error) {
	var item model.EquipmentDocument
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *EquipmentDocumentRepository) Create(ctx context.Context, item *model.EquipmentDocument) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *EquipmentDocumentRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.EquipmentDocument{}).Where("id = ?", id).Updates(updates).Error
}

func (r *EquipmentDocumentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.EquipmentDocument{}).Error
}

func (r *EquipmentDocumentRepository) ListByEquipmentID(ctx context.Context, equipmentID int64) ([]model.EquipmentDocument, error) {
	var list []model.EquipmentDocument
	err := r.db.WithContext(ctx).Where("equipment_id = ?", equipmentID).Order("id DESC").Find(&list).Error
	return list, err
}
