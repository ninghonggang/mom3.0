package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"mom-platform/services/mdm-service/internal/model"
)

var ErrNotFound = errors.New("record not found")

type gormMaterialRepo struct{ db *gorm.DB }

func NewMaterialRepo(db *gorm.DB) MaterialRepository { return &gormMaterialRepo{db: db} }

func (r *gormMaterialRepo) Create(ctx context.Context, m *model.Material) error {
	return r.db.WithContext(ctx).Create(m).Error
}
func (r *gormMaterialRepo) Update(ctx context.Context, m *model.Material) error {
	return r.db.WithContext(ctx).Save(m).Error
}
func (r *gormMaterialRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Material{}, id).Error
}
func (r *gormMaterialRepo) GetByID(ctx context.Context, id uint64) (*model.Material, error) {
	var m model.Material
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
func (r *gormMaterialRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.Material, error) {
	var m model.Material
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND material_code = ?", tenantID, code).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
func (r *gormMaterialRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.Material, int64, error) {
	var list []model.Material
	var total int64
	query := r.db.WithContext(ctx).Model(&model.Material{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// -- BOM --

type gormBomRepo struct{ db *gorm.DB }

func NewBomRepo(db *gorm.DB) BomRepository { return &gormBomRepo{db: db} }

func (r *gormBomRepo) Create(ctx context.Context, b *model.Bom) error {
	return r.db.WithContext(ctx).Create(b).Error
}
func (r *gormBomRepo) Update(ctx context.Context, b *model.Bom) error {
	return r.db.WithContext(ctx).Save(b).Error
}
func (r *gormBomRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Bom{}, id).Error
}
func (r *gormBomRepo) GetByID(ctx context.Context, id uint64) (*model.Bom, error) {
	var b model.Bom
	if err := r.db.WithContext(ctx).Preload("Items").First(&b, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &b, nil
}
func (r *gormBomRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.Bom, error) {
	var b model.Bom
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND bom_code = ?", tenantID, code).First(&b).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &b, nil
}
func (r *gormBomRepo) GetActiveBom(ctx context.Context, tenantID string, materialID uint64) (*model.Bom, error) {
	var b model.Bom
	if err := r.db.WithContext(ctx).Preload("Items").
		Where("tenant_id = ? AND material_id = ? AND status = ?", tenantID, materialID, "ACTIVE").
		First(&b).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &b, nil
}
func (r *gormBomRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.Bom, int64, error) {
	var list []model.Bom
	var total int64
	query := r.db.WithContext(ctx).Model(&model.Bom{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *gormBomRepo) UpdateStatus(ctx context.Context, id uint64, status string) error {
	return r.db.WithContext(ctx).Model(&model.Bom{}).Where("id = ?", id).Update("status", status).Error
}
func (r *gormBomRepo) DeleteItems(ctx context.Context, bomID uint64) error {
	return r.db.WithContext(ctx).Where("bom_id = ?", bomID).Delete(&model.BomItem{}).Error
}
func (r *gormBomRepo) CreateItems(ctx context.Context, items []model.BomItem) error {
	return r.db.WithContext(ctx).Create(&items).Error
}
func (r *gormBomRepo) GetItems(ctx context.Context, bomID uint64) ([]model.BomItem, error) {
	var items []model.BomItem
	if err := r.db.WithContext(ctx).Where("bom_id = ?", bomID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// -- Workshop --

type gormWorkshopRepo struct{ db *gorm.DB }

func NewWorkshopRepo(db *gorm.DB) WorkshopRepository { return &gormWorkshopRepo{db: db} }

func (r *gormWorkshopRepo) Create(ctx context.Context, w *model.Workshop) error { return r.db.WithContext(ctx).Create(w).Error }
func (r *gormWorkshopRepo) Update(ctx context.Context, w *model.Workshop) error { return r.db.WithContext(ctx).Save(w).Error }
func (r *gormWorkshopRepo) Delete(ctx context.Context, id uint64) error { return r.db.WithContext(ctx).Delete(&model.Workshop{}, id).Error }
func (r *gormWorkshopRepo) GetByID(ctx context.Context, id uint64) (*model.Workshop, error) {
	var w model.Workshop
	if err := r.db.WithContext(ctx).First(&w, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &w, nil
}
func (r *gormWorkshopRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.Workshop, error) {
	var w model.Workshop
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND workshop_code = ?", tenantID, code).First(&w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &w, nil
}
func (r *gormWorkshopRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.Workshop, int64, error) {
	var list []model.Workshop
	var total int64
	query := r.db.WithContext(ctx).Model(&model.Workshop{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// -- Production Line --

type gormProductionLineRepo struct{ db *gorm.DB }

func NewProductionLineRepo(db *gorm.DB) ProductionLineRepository { return &gormProductionLineRepo{db: db} }

func (r *gormProductionLineRepo) Create(ctx context.Context, pl *model.ProductionLine) error {
	return r.db.WithContext(ctx).Create(pl).Error
}
func (r *gormProductionLineRepo) Update(ctx context.Context, pl *model.ProductionLine) error {
	return r.db.WithContext(ctx).Save(pl).Error
}
func (r *gormProductionLineRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.ProductionLine{}, id).Error
}
func (r *gormProductionLineRepo) GetByID(ctx context.Context, id uint64) (*model.ProductionLine, error) {
	var pl model.ProductionLine
	if err := r.db.WithContext(ctx).First(&pl, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &pl, nil
}
func (r *gormProductionLineRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.ProductionLine, error) {
	var pl model.ProductionLine
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND line_code = ?", tenantID, code).First(&pl).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &pl, nil
}
func (r *gormProductionLineRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.ProductionLine, int64, error) {
	var list []model.ProductionLine
	var total int64
	query := r.db.WithContext(ctx).Model(&model.ProductionLine{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *gormProductionLineRepo) ListByWorkshop(ctx context.Context, tenantID string, workshopID uint64) ([]model.ProductionLine, error) {
	var list []model.ProductionLine
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND workshop_id = ?", tenantID, workshopID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// -- Workstation --

type gormWorkstationRepo struct{ db *gorm.DB }

func NewWorkstationRepo(db *gorm.DB) WorkstationRepository { return &gormWorkstationRepo{db: db} }

func (r *gormWorkstationRepo) Create(ctx context.Context, ws *model.Workstation) error {
	return r.db.WithContext(ctx).Create(ws).Error
}
func (r *gormWorkstationRepo) Update(ctx context.Context, ws *model.Workstation) error {
	return r.db.WithContext(ctx).Save(ws).Error
}
func (r *gormWorkstationRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.Workstation{}, id).Error
}
func (r *gormWorkstationRepo) GetByID(ctx context.Context, id uint64) (*model.Workstation, error) {
	var ws model.Workstation
	if err := r.db.WithContext(ctx).First(&ws, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &ws, nil
}
func (r *gormWorkstationRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.Workstation, error) {
	var ws model.Workstation
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND workstation_code = ?", tenantID, code).First(&ws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &ws, nil
}
func (r *gormWorkstationRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.Workstation, int64, error) {
	var list []model.Workstation
	var total int64
	query := r.db.WithContext(ctx).Model(&model.Workstation{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// -- Customer --

type gormCustomerRepo struct{ db *gorm.DB }

func NewCustomerRepo(db *gorm.DB) CustomerRepository { return &gormCustomerRepo{db: db} }

func (r *gormCustomerRepo) Create(ctx context.Context, c *model.Customer) error { return r.db.WithContext(ctx).Create(c).Error }
func (r *gormCustomerRepo) Update(ctx context.Context, c *model.Customer) error { return r.db.WithContext(ctx).Save(c).Error }
func (r *gormCustomerRepo) Delete(ctx context.Context, id uint64) error     { return r.db.WithContext(ctx).Delete(&model.Customer{}, id).Error }
func (r *gormCustomerRepo) GetByID(ctx context.Context, id uint64) (*model.Customer, error) {
	var c model.Customer
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &c, nil
}
func (r *gormCustomerRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.Customer, error) {
	var c model.Customer
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND customer_code = ?", tenantID, code).First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &c, nil
}
func (r *gormCustomerRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.Customer, int64, error) {
	var list []model.Customer
	var total int64
	query := r.db.WithContext(ctx).Model(&model.Customer{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// -- Supplier --

type gormSupplierRepo struct{ db *gorm.DB }

func NewSupplierRepo(db *gorm.DB) SupplierRepository { return &gormSupplierRepo{db: db} }

func (r *gormSupplierRepo) Create(ctx context.Context, s *model.Supplier) error { return r.db.WithContext(ctx).Create(s).Error }
func (r *gormSupplierRepo) Update(ctx context.Context, s *model.Supplier) error { return r.db.WithContext(ctx).Save(s).Error }
func (r *gormSupplierRepo) Delete(ctx context.Context, id uint64) error     { return r.db.WithContext(ctx).Delete(&model.Supplier{}, id).Error }
func (r *gormSupplierRepo) GetByID(ctx context.Context, id uint64) (*model.Supplier, error) {
	var s model.Supplier
	if err := r.db.WithContext(ctx).First(&s, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &s, nil
}
func (r *gormSupplierRepo) GetByCode(ctx context.Context, tenantID, code string) (*model.Supplier, error) {
	var s model.Supplier
	if err := r.db.WithContext(ctx).Where("tenant_id = ? AND supplier_code = ?", tenantID, code).First(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &s, nil
}
func (r *gormSupplierRepo) List(ctx context.Context, tenantID string, offset, limit int) ([]model.Supplier, int64, error) {
	var list []model.Supplier
	var total int64
	query := r.db.WithContext(ctx).Model(&model.Supplier{})
	if tenantID != "" && tenantID != "0" {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
