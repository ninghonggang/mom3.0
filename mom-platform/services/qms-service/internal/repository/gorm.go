package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"mom-platform/services/qms-service/internal/model"
)

// GormRepository implements Repository using GORM.
type GormRepository struct {
	db *gorm.DB
}

// NewGormRepository creates a new GormRepository.
func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

// DB exposes the underlying *gorm.DB (used by the service for transactions).
func (r *GormRepository) DB() *gorm.DB { return r.db }

// AutoMigrate runs GORM auto-migration for all QMS models.
func (r *GormRepository) AutoMigrate() error {
	return r.db.AutoMigrate(
		&model.InspectionSheet{},
		&model.InspectionCharacteristic{},
		&model.InspectionPlan{},
		&model.InspectionResult{},
		&model.Ncr{},
		&model.NcrAction{},
		&model.DefectCode{},
		&model.SpcData{},
	)
}

// applyPage normalises page query and returns offset/limit.
func applyPage(q PageQuery) (int, int) {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 1 {
		q.PageSize = 20
	}
	if q.PageSize > 200 {
		q.PageSize = 200
	}
	return (q.Page - 1) * q.PageSize, q.PageSize
}

// applyFilter adds WHERE conditions from the filter map.
func applyFilter(tx *gorm.DB, filter map[string]interface{}) *gorm.DB {
	for k, v := range filter {
		tx = tx.Where(k, v)
	}
	return tx
}

// =============================== InspectionSheet ===============================

func (r *GormRepository) CreateInspectionSheet(ctx context.Context, s *model.InspectionSheet) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *GormRepository) GetInspectionSheet(ctx context.Context, id uint) (*model.InspectionSheet, error) {
	var s model.InspectionSheet
	if err := r.db.WithContext(ctx).First(&s, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (r *GormRepository) UpdateInspectionSheet(ctx context.Context, s *model.InspectionSheet) error {
	return r.db.WithContext(ctx).Save(s).Error
}

func (r *GormRepository) DeleteInspectionSheet(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.InspectionSheet{}, id).Error
}

func (r *GormRepository) ListInspectionSheets(ctx context.Context, q PageQuery) (*PageResult[model.InspectionSheet], error) {
	var items []model.InspectionSheet
	var total int64

	tx := r.db.WithContext(ctx).Model(&model.InspectionSheet{})
	tx = applyFilter(tx, q.Filter)
	if err := tx.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("count inspection sheets: %w", err)
	}

	offset, limit := applyPage(q)
	if err := tx.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list inspection sheets: %w", err)
	}
	return &PageResult[model.InspectionSheet]{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize}, nil
}

// =============================== InspectionCharacteristic ===============================

func (r *GormRepository) CreateInspectionCharacteristic(ctx context.Context, c *model.InspectionCharacteristic) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *GormRepository) GetInspectionCharacteristic(ctx context.Context, id uint) (*model.InspectionCharacteristic, error) {
	var c model.InspectionCharacteristic
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *GormRepository) UpdateInspectionCharacteristic(ctx context.Context, c *model.InspectionCharacteristic) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *GormRepository) DeleteInspectionCharacteristic(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.InspectionCharacteristic{}, id).Error
}

func (r *GormRepository) ListInspectionCharacteristics(ctx context.Context, q PageQuery) (*PageResult[model.InspectionCharacteristic], error) {
	var items []model.InspectionCharacteristic
	var total int64

	tx := r.db.WithContext(ctx).Model(&model.InspectionCharacteristic{})
	tx = applyFilter(tx, q.Filter)
	if err := tx.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("count inspection characteristics: %w", err)
	}

	offset, limit := applyPage(q)
	if err := tx.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list inspection characteristics: %w", err)
	}
	return &PageResult[model.InspectionCharacteristic]{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize}, nil
}

// =============================== InspectionPlan ===============================

func (r *GormRepository) CreateInspectionPlan(ctx context.Context, p *model.InspectionPlan) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *GormRepository) GetInspectionPlan(ctx context.Context, id uint) (*model.InspectionPlan, error) {
	var p model.InspectionPlan
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *GormRepository) UpdateInspectionPlan(ctx context.Context, p *model.InspectionPlan) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *GormRepository) DeleteInspectionPlan(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.InspectionPlan{}, id).Error
}

func (r *GormRepository) ListInspectionPlans(ctx context.Context, q PageQuery) (*PageResult[model.InspectionPlan], error) {
	var items []model.InspectionPlan
	var total int64

	tx := r.db.WithContext(ctx).Model(&model.InspectionPlan{})
	tx = applyFilter(tx, q.Filter)
	if err := tx.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("count inspection plans: %w", err)
	}

	offset, limit := applyPage(q)
	if err := tx.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list inspection plans: %w", err)
	}
	return &PageResult[model.InspectionPlan]{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize}, nil
}

// =============================== InspectionResult ===============================

func (r *GormRepository) CreateInspectionResult(ctx context.Context, res *model.InspectionResult) error {
	return r.db.WithContext(ctx).Create(res).Error
}

func (r *GormRepository) GetInspectionResult(ctx context.Context, id uint) (*model.InspectionResult, error) {
	var res model.InspectionResult
	if err := r.db.WithContext(ctx).First(&res, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}

func (r *GormRepository) ListInspectionResults(ctx context.Context, q PageQuery) (*PageResult[model.InspectionResult], error) {
	var items []model.InspectionResult
	var total int64

	tx := r.db.WithContext(ctx).Model(&model.InspectionResult{})
	tx = applyFilter(tx, q.Filter)
	if err := tx.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("count inspection results: %w", err)
	}

	offset, limit := applyPage(q)
	if err := tx.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list inspection results: %w", err)
	}
	return &PageResult[model.InspectionResult]{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize}, nil
}

func (r *GormRepository) ListResultsBySheet(ctx context.Context, sheetID uint) ([]model.InspectionResult, error) {
	var items []model.InspectionResult
	if err := r.db.WithContext(ctx).Where("sheet_id = ?", sheetID).Order("id ASC").Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list results by sheet %d: %w", sheetID, err)
	}
	return items, nil
}

// =============================== Ncr ===============================

func (r *GormRepository) CreateNcr(ctx context.Context, n *model.Ncr) error {
	return r.db.WithContext(ctx).Create(n).Error
}

func (r *GormRepository) GetNcr(ctx context.Context, id uint) (*model.Ncr, error) {
	var n model.Ncr
	if err := r.db.WithContext(ctx).First(&n, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &n, nil
}

func (r *GormRepository) UpdateNcr(ctx context.Context, n *model.Ncr) error {
	return r.db.WithContext(ctx).Save(n).Error
}

func (r *GormRepository) DeleteNcr(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Ncr{}, id).Error
}

func (r *GormRepository) ListNcrs(ctx context.Context, q PageQuery) (*PageResult[model.Ncr], error) {
	var items []model.Ncr
	var total int64

	tx := r.db.WithContext(ctx).Model(&model.Ncr{})
	tx = applyFilter(tx, q.Filter)
	if err := tx.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("count ncrs: %w", err)
	}

	offset, limit := applyPage(q)
	if err := tx.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list ncrs: %w", err)
	}
	return &PageResult[model.Ncr]{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize}, nil
}

// =============================== NcrAction ===============================

func (r *GormRepository) CreateNcrAction(ctx context.Context, a *model.NcrAction) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *GormRepository) GetNcrAction(ctx context.Context, id uint) (*model.NcrAction, error) {
	var a model.NcrAction
	if err := r.db.WithContext(ctx).First(&a, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *GormRepository) ListNcrActions(ctx context.Context, ncrID uint) ([]model.NcrAction, error) {
	var items []model.NcrAction
	if err := r.db.WithContext(ctx).Where("ncr_id = ?", ncrID).Order("id ASC").Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list ncr actions for ncr %d: %w", ncrID, err)
	}
	return items, nil
}

// =============================== DefectCode ===============================

func (r *GormRepository) CreateDefectCode(ctx context.Context, d *model.DefectCode) error {
	return r.db.WithContext(ctx).Create(d).Error
}

func (r *GormRepository) GetDefectCode(ctx context.Context, id uint) (*model.DefectCode, error) {
	var d model.DefectCode
	if err := r.db.WithContext(ctx).First(&d, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func (r *GormRepository) UpdateDefectCode(ctx context.Context, d *model.DefectCode) error {
	return r.db.WithContext(ctx).Save(d).Error
}

func (r *GormRepository) DeleteDefectCode(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DefectCode{}, id).Error
}

func (r *GormRepository) ListDefectCodes(ctx context.Context, q PageQuery) (*PageResult[model.DefectCode], error) {
	var items []model.DefectCode
	var total int64

	tx := r.db.WithContext(ctx).Model(&model.DefectCode{})
	tx = applyFilter(tx, q.Filter)
	if err := tx.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("count defect codes: %w", err)
	}

	offset, limit := applyPage(q)
	if err := tx.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list defect codes: %w", err)
	}
	return &PageResult[model.DefectCode]{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize}, nil
}

// =============================== SpcData ===============================

func (r *GormRepository) CreateSpcData(ctx context.Context, s *model.SpcData) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *GormRepository) GetSpcData(ctx context.Context, id uint) (*model.SpcData, error) {
	var s model.SpcData
	if err := r.db.WithContext(ctx).First(&s, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (r *GormRepository) ListSpcData(ctx context.Context, q PageQuery) (*PageResult[model.SpcData], error) {
	var items []model.SpcData
	var total int64

	tx := r.db.WithContext(ctx).Model(&model.SpcData{})
	tx = applyFilter(tx, q.Filter)
	if err := tx.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("count spc data: %w", err)
	}

	offset, limit := applyPage(q)
	if err := tx.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list spc data: %w", err)
	}
	return &PageResult[model.SpcData]{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize}, nil
}

func (r *GormRepository) ListSpcDataByChar(ctx context.Context, charID uint) ([]model.SpcData, error) {
	var items []model.SpcData
	if err := r.db.WithContext(ctx).Where("char_id = ?", charID).Order("sample_time ASC").Find(&items).Error; err != nil {
		return nil, fmt.Errorf("list spc data for char %d: %w", charID, err)
	}
	return items, nil
}
