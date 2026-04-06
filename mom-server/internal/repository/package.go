package repository

import (
	"context"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type PackageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) *PackageRepository {
	return &PackageRepository{db: db}
}

func (r *PackageRepository) List(ctx context.Context, tenantID int64, req *PackageQuery) ([]model.MesPackage, int64, error) {
	var list []model.MesPackage
	var total int64
	q := r.db.WithContext(ctx).Model(&model.MesPackage{}).Where("tenant_id = ?", tenantID)
	if req.PackageNo != "" {
		q = q.Where("package_no LIKE ?", "%"+req.PackageNo+"%")
	}
	if req.Status != "" {
		q = q.Where("status = ?", req.Status)
	}
	if req.PackageType != "" {
		q = q.Where("package_type = ?", req.PackageType)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("id DESC").Limit(req.Limit).Offset(req.Offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *PackageRepository) GetByID(ctx context.Context, id uint) (*model.MesPackage, error) {
	var item model.MesPackage
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *PackageRepository) GetByPackageNo(ctx context.Context, packageNo string) (*model.MesPackage, error) {
	var item model.MesPackage
	if err := r.db.WithContext(ctx).Where("package_no = ?", packageNo).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *PackageRepository) Create(ctx context.Context, item *model.MesPackage) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *PackageRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesPackage{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PackageRepository) UpdateSerialNos(ctx context.Context, id uint, serialNos model.SerialNos) error {
	return r.db.WithContext(ctx).Model(&model.MesPackage{}).Where("id = ?", id).Update("serial_nos", serialNos).Error
}

func (r *PackageRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.MesPackage{}).Error
}

type PackageQuery struct {
	PackageNo   string
	Status      string
	PackageType string
	Limit       int
	Offset      int
}
