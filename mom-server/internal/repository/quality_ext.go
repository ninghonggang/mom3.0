package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ========== QRCI Repository ==========

type QRCIRepository struct {
	db *gorm.DB
}

func NewQRCIRepository(db *gorm.DB) *QRCIRepository {
	return &QRCIRepository{db: db}
}

func (r *QRCIRepository) Create(ctx context.Context, m *model.QRCI) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *QRCIRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.QRCI{}).Where("id = ?", id).Updates(updates).Error
}

func (r *QRCIRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("qrci_id = ?", id).Delete(&model.QRCI5Why{})
		tx.Where("qrci_id = ?", id).Delete(&model.QRCIAction{})
		tx.Where("qrci_id = ?", id).Delete(&model.QRCIVerification{})
		return tx.Delete(&model.QRCI{}, id).Error
	})
}

func (r *QRCIRepository) GetByID(ctx context.Context, id uint) (*model.QRCI, error) {
	var m model.QRCI
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *QRCIRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.QRCI, int64, error) {
	var list []model.QRCI
	var total int64
	query := r.db.WithContext(ctx).Model(&model.QRCI{})
	for k, v := range filters {
		if v != "" {
			query = query.Where(k+" = ?", v)
		}
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit).Order("id desc")
	}
	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

type QRCI5WhyRepository struct {
	db *gorm.DB
}

func NewQRCI5WhyRepository(db *gorm.DB) *QRCI5WhyRepository {
	return &QRCI5WhyRepository{db: db}
}

func (r *QRCI5WhyRepository) Create(ctx context.Context, m *model.QRCI5Why) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *QRCI5WhyRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.QRCI5Why{}).Where("id = ?", id).Updates(updates).Error
}

func (r *QRCI5WhyRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.QRCI5Why{}, id).Error
}

func (r *QRCI5WhyRepository) ListByQRCI(ctx context.Context, qrciID uint) ([]model.QRCI5Why, error) {
	var list []model.QRCI5Why
	err := r.db.WithContext(ctx).Where("qrci_id = ?", qrciID).Order("why_level").Find(&list).Error
	return list, err
}

type QRCIActionRepository struct {
	db *gorm.DB
}

func NewQRCIActionRepository(db *gorm.DB) *QRCIActionRepository {
	return &QRCIActionRepository{db: db}
}

func (r *QRCIActionRepository) Create(ctx context.Context, m *model.QRCIAction) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *QRCIActionRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.QRCIAction{}).Where("id = ?", id).Updates(updates).Error
}

func (r *QRCIActionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.QRCIAction{}, id).Error
}

func (r *QRCIActionRepository) GetByID(ctx context.Context, id uint) (*model.QRCIAction, error) {
	var m model.QRCIAction
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *QRCIActionRepository) ListByQRCI(ctx context.Context, qrciID uint) ([]model.QRCIAction, error) {
	var list []model.QRCIAction
	err := r.db.WithContext(ctx).Where("qrci_id = ?", qrciID).Order("id").Find(&list).Error
	return list, err
}

type QRCIVerificationRepository struct {
	db *gorm.DB
}

func NewQRCIVerificationRepository(db *gorm.DB) *QRCIVerificationRepository {
	return &QRCIVerificationRepository{db: db}
}

func (r *QRCIVerificationRepository) Create(ctx context.Context, m *model.QRCIVerification) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *QRCIVerificationRepository) ListByQRCI(ctx context.Context, qrciID uint) ([]model.QRCIVerification, error) {
	var list []model.QRCIVerification
	err := r.db.WithContext(ctx).Where("qrci_id = ?", qrciID).Order("id desc").Find(&list).Error
	return list, err
}

// ========== LPA Repository ==========

type LPAStandardRepository struct {
	db *gorm.DB
}

func NewLPAStandardRepository(db *gorm.DB) *LPAStandardRepository {
	return &LPAStandardRepository{db: db}
}

func (r *LPAStandardRepository) Create(ctx context.Context, m *model.LPAStandard) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *LPAStandardRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LPAStandard{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LPAStandardRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tx.Where("standard_id = ?", id).Delete(&model.LPAQuestion{})
		return tx.Delete(&model.LPAStandard{}, id).Error
	})
}

func (r *LPAStandardRepository) GetByID(ctx context.Context, id uint) (*model.LPAStandard, error) {
	var m model.LPAStandard
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *LPAStandardRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.LPAStandard, int64, error) {
	var list []model.LPAStandard
	var total int64
	query := r.db.WithContext(ctx).Model(&model.LPAStandard{})
	for k, v := range filters {
		if v != "" {
			query = query.Where(k+" = ?", v)
		}
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit).Order("id desc")
	}
	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

type LPAQuestionRepository struct {
	db *gorm.DB
}

func NewLPAQuestionRepository(db *gorm.DB) *LPAQuestionRepository {
	return &LPAQuestionRepository{db: db}
}

func (r *LPAQuestionRepository) Create(ctx context.Context, m *model.LPAQuestion) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *LPAQuestionRepository) CreateBatch(ctx context.Context, items []model.LPAQuestion) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *LPAQuestionRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LPAQuestion{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LPAQuestionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.LPAQuestion{}, id).Error
}

func (r *LPAQuestionRepository) ListByStandard(ctx context.Context, standardID uint) ([]model.LPAQuestion, error) {
	var list []model.LPAQuestion
	err := r.db.WithContext(ctx).Where("standard_id = ?", standardID).Order("sort_order").Find(&list).Error
	return list, err
}

type LPARecordRepository struct {
	db *gorm.DB
}

func NewLPARecordRepository(db *gorm.DB) *LPARecordRepository {
	return &LPARecordRepository{db: db}
}

func (r *LPARecordRepository) Create(ctx context.Context, m *model.LPARecord) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *LPARecordRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LPARecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LPARecordRepository) GetByID(ctx context.Context, id uint) (*model.LPARecord, error) {
	var m model.LPARecord
	err := r.db.WithContext(ctx).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *LPARecordRepository) List(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]model.LPARecord, int64, error) {
	var list []model.LPARecord
	var total int64
	query := r.db.WithContext(ctx).Model(&model.LPARecord{})
	for k, v := range filters {
		if v != "" {
			query = query.Where(k+" = ?", v)
		}
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit > 0 {
		query = query.Offset(offset).Limit(limit).Order("id desc")
	}
	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

type LPARecordItemRepository struct {
	db *gorm.DB
}

func NewLPARecordItemRepository(db *gorm.DB) *LPARecordItemRepository {
	return &LPARecordItemRepository{db: db}
}

func (r *LPARecordItemRepository) CreateBatch(ctx context.Context, items []model.LPARecordItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}

func (r *LPARecordItemRepository) ListByRecord(ctx context.Context, recordID uint) ([]model.LPARecordItem, error) {
	var list []model.LPARecordItem
	err := r.db.WithContext(ctx).Where("record_id = ?", recordID).Find(&list).Error
	return list, err
}
