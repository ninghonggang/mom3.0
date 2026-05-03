package repository

import (
	"context"
	"fmt"
	"strings"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

type QualityCertificateRepository struct {
	db *gorm.DB
}

func NewQualityCertificateRepository(db *gorm.DB) *QualityCertificateRepository {
	return &QualityCertificateRepository{db: db}
}

func (r *QualityCertificateRepository) List(ctx context.Context, tenantID int64, query *model.QualityCertificateQuery) ([]model.QualityCertificate, int64, error) {
	var list []model.QualityCertificate
	var total int64

	db := r.db.WithContext(ctx).Model(&model.QualityCertificate{}).Where("tenant_id = ?", tenantID)

	if query.OrderCode != "" {
		db = db.Where("order_code LIKE ?", "%"+query.OrderCode+"%")
	}
	if query.ProductCode != "" {
		db = db.Where("product_code LIKE ?", "%"+query.ProductCode+"%")
	}
	if query.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+query.ProductName+"%")
	}
	if query.BatchNo != "" {
		db = db.Where("batch_no LIKE ?", "%"+query.BatchNo+"%")
	}
	if query.CertType != "" {
		db = db.Where("cert_type = ?", query.CertType)
	}
	if query.Result > 0 {
		db = db.Where("result = ?", query.Result)
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}
	if query.StartDate != "" {
		db = db.Where("inspect_date >= ?", query.StartDate)
	}
	if query.EndDate != "" {
		db = db.Where("inspect_date <= ?", query.EndDate)
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

func (r *QualityCertificateRepository) GetByID(ctx context.Context, id int64) (*model.QualityCertificate, error) {
	var cert model.QualityCertificate
	if err := r.db.WithContext(ctx).First(&cert, id).Error; err != nil {
		return nil, err
	}
	return &cert, nil
}

func (r *QualityCertificateRepository) Create(ctx context.Context, tenantID int64, cert *model.QualityCertificate) error {
	cert.TenantID = tenantID
	return r.db.WithContext(ctx).Create(cert).Error
}

func (r *QualityCertificateRepository) Update(ctx context.Context, id int64, cert *model.QualityCertificate) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Updates(cert).Error
}

func (r *QualityCertificateRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.QualityCertificate{}, id).Error
}

func (r *QualityCertificateRepository) GenerateCertCode(ctx context.Context, tenantID int64) (string, error) {
	var maxCode string
	r.db.WithContext(ctx).Model(&model.QualityCertificate{}).
		Where("cert_code LIKE ?", "COC-"+strings.Repeat("0", 8)).
		Order("id DESC").Select("cert_code").Scan(&maxCode)
	if maxCode == "" {
		return "COC-00000001", nil
	}
	var num int
	if _, err := fmt.Sscanf(maxCode, "COC-%d", &num); err != nil {
		return "", err
	}
	return fmt.Sprintf("COC-%08d", num+1), nil
}