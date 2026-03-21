package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type IQCRepository struct {
	db *gorm.DB
}

func NewIQCRepository(db *gorm.DB) *IQCRepository {
	return &IQCRepository{db: db}
}

func (r *IQCRepository) List(ctx context.Context, tenantID int64) ([]model.IQC, int64, error) {
	var list []model.IQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.IQC{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *IQCRepository) GetByID(ctx context.Context, id uint) (*model.IQC, error) {
	var iqc model.IQC
	err := r.db.WithContext(ctx).First(&iqc, id).Error
	return &iqc, err
}

func (r *IQCRepository) Create(ctx context.Context, iqc *model.IQC) error {
	return r.db.WithContext(ctx).Create(iqc).Error
}

func (r *IQCRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.IQC{}).Where("id = ?", id).Updates(updates).Error
}

func (r *IQCRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.IQC{}, id).Error
}

type IPQCRepository struct {
	db *gorm.DB
}

func NewIPQCRepository(db *gorm.DB) *IPQCRepository {
	return &IPQCRepository{db: db}
}

func (r *IPQCRepository) List(ctx context.Context, tenantID int64) ([]model.IPQC, int64, error) {
	var list []model.IPQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.IPQC{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *IPQCRepository) Create(ctx context.Context, ipqc *model.IPQC) error {
	return r.db.WithContext(ctx).Create(ipqc).Error
}

type FQCRepository struct {
	db *gorm.DB
}

func NewFQCRepository(db *gorm.DB) *FQCRepository {
	return &FQCRepository{db: db}
}

func (r *FQCRepository) List(ctx context.Context, tenantID int64) ([]model.FQC, int64, error) {
	var list []model.FQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.FQC{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *FQCRepository) Create(ctx context.Context, fqc *model.FQC) error {
	return r.db.WithContext(ctx).Create(fqc).Error
}

type OQCRepository struct {
	db *gorm.DB
}

func NewOQCRepository(db *gorm.DB) *OQCRepository {
	return &OQCRepository{db: db}
}

func (r *OQCRepository) List(ctx context.Context, tenantID int64) ([]model.OQC, int64, error) {
	var list []model.OQC
	var total int64

	err := r.db.WithContext(ctx).Model(&model.OQC{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *OQCRepository) Create(ctx context.Context, oqc *model.OQC) error {
	return r.db.WithContext(ctx).Create(oqc).Error
}

type DefectRecordRepository struct {
	db *gorm.DB
}

func NewDefectRecordRepository(db *gorm.DB) *DefectRecordRepository {
	return &DefectRecordRepository{db: db}
}

func (r *DefectRecordRepository) List(ctx context.Context, tenantID int64) ([]model.DefectRecord, int64, error) {
	var list []model.DefectRecord
	var total int64

	err := r.db.WithContext(ctx).Model(&model.DefectRecord{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *DefectRecordRepository) Create(ctx context.Context, record *model.DefectRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *DefectRecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DefectRecord{}).Where("id = ?", id).Updates(updates).Error
}

type NCRRepository struct {
	db *gorm.DB
}

func NewNCRRepository(db *gorm.DB) *NCRRepository {
	return &NCRRepository{db: db}
}

func (r *NCRRepository) List(ctx context.Context, tenantID int64) ([]model.NCR, int64, error) {
	var list []model.NCR
	var total int64

	err := r.db.WithContext(ctx).Model(&model.NCR{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *NCRRepository) Create(ctx context.Context, ncr *model.NCR) error {
	return r.db.WithContext(ctx).Create(ncr).Error
}

func (r *NCRRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.NCR{}).Where("id = ?", id).Updates(updates).Error
}
