package repository

import (
	"context"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

type ElectronicSOPRepository struct {
	db *gorm.DB
}

func NewElectronicSOPRepository(db *gorm.DB) *ElectronicSOPRepository {
	return &ElectronicSOPRepository{db: db}
}

func (r *ElectronicSOPRepository) List(ctx context.Context, tenantID int64, query string) ([]model.ElectronicSOP, int64, error) {
	var list []model.ElectronicSOP
	var total int64

	db := r.db.WithContext(ctx).Model(&model.ElectronicSOP{}).Where("tenant_id = ?", tenantID)

	if query != "" {
		db = db.Where("sop_no LIKE ? OR sop_name LIKE ? OR material_code LIKE ? OR material_name LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%")
	}

	db.Count(&total)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *ElectronicSOPRepository) GetByID(ctx context.Context, id uint) (*model.ElectronicSOP, error) {
	var sop model.ElectronicSOP
	err := r.db.WithContext(ctx).First(&sop, id).Error
	return &sop, err
}

func (r *ElectronicSOPRepository) Create(ctx context.Context, sop *model.ElectronicSOP) error {
	return r.db.WithContext(ctx).Create(sop).Error
}

func (r *ElectronicSOPRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.ElectronicSOP{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ElectronicSOPRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.ElectronicSOP{}, id).Error
}

func (r *ElectronicSOPRepository) GetByMaterial(ctx context.Context, materialID int64) ([]model.ElectronicSOP, error) {
	var list []model.ElectronicSOP
	err := r.db.WithContext(ctx).Where("material_id = ? AND status = 2", materialID).Find(&list).Error
	return list, err
}

type CodeRuleRepository struct {
	db *gorm.DB
}

func NewCodeRuleRepository(db *gorm.DB) *CodeRuleRepository {
	return &CodeRuleRepository{db: db}
}

func (r *CodeRuleRepository) List(ctx context.Context, tenantID int64) ([]model.CodeRule, error) {
	var list []model.CodeRule
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *CodeRuleRepository) GetByID(ctx context.Context, id uint) (*model.CodeRule, error) {
	var rule model.CodeRule
	err := r.db.WithContext(ctx).First(&rule, id).Error
	return &rule, err
}

func (r *CodeRuleRepository) GetByCode(ctx context.Context, tenantID int64, ruleCode string) (*model.CodeRule, error) {
	var rule model.CodeRule
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND rule_code = ?", tenantID, ruleCode).First(&rule).Error
	return &rule, err
}

func (r *CodeRuleRepository) Create(ctx context.Context, rule *model.CodeRule) error {
	return r.db.WithContext(ctx).Create(rule).Error
}

func (r *CodeRuleRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.CodeRule{}).Where("id = ?", id).Updates(updates).Error
}

func (r *CodeRuleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.CodeRule{}, id).Error
}

func (r *CodeRuleRepository) UpdateSeq(ctx context.Context, id int64, seqValue int) error {
	return r.db.WithContext(ctx).Model(&model.CodeRule{}).Where("id = ?", id).Updates(map[string]any{"seq_current": seqValue, "last_gen_date": time.Now().Format("20060102")}).Error
}

func (r *CodeRuleRepository) CreateRecord(ctx context.Context, record *model.CodeRuleRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *CodeRuleRepository) GetRecordByCode(ctx context.Context, genCode string) (*model.CodeRuleRecord, error) {
	var record model.CodeRuleRecord
	err := r.db.WithContext(ctx).Where("gen_code = ?", genCode).First(&record).Error
	return &record, err
}

type FlowCardRepository struct {
	db *gorm.DB
}

func NewFlowCardRepository(db *gorm.DB) *FlowCardRepository {
	return &FlowCardRepository{db: db}
}

func (r *FlowCardRepository) List(ctx context.Context, tenantID int64, query string) ([]model.FlowCard, int64, error) {
	var list []model.FlowCard
	var total int64

	db := r.db.WithContext(ctx).Model(&model.FlowCard{}).Where("tenant_id = ?", tenantID)

	if query != "" {
		db = db.Where("card_no LIKE ? OR order_no LIKE ? OR material_code LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}

	db.Count(&total)
	err := db.Order("priority DESC, created_at DESC").Find(&list).Error
	return list, total, err
}

func (r *FlowCardRepository) GetByID(ctx context.Context, id uint) (*model.FlowCard, error) {
	var card model.FlowCard
	err := r.db.WithContext(ctx).First(&card, id).Error
	return &card, err
}

func (r *FlowCardRepository) Create(ctx context.Context, card *model.FlowCard) error {
	return r.db.WithContext(ctx).Create(card).Error
}

func (r *FlowCardRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.FlowCard{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FlowCardRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.FlowCard{}, id).Error
}

func (r *FlowCardRepository) GetDetails(ctx context.Context, cardID int64) ([]model.FlowCardDetail, error) {
	var details []model.FlowCardDetail
	err := r.db.WithContext(ctx).Where("card_id = ?", cardID).Order("step_no ASC").Find(&details).Error
	return details, err
}

func (r *FlowCardRepository) CreateDetail(ctx context.Context, detail *model.FlowCardDetail) error {
	return r.db.WithContext(ctx).Create(detail).Error
}

func (r *FlowCardRepository) CreateDetails(ctx context.Context, details []model.FlowCardDetail) error {
	return r.db.WithContext(ctx).Create(&details).Error
}
