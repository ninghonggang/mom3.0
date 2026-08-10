package repository

import (
	"context"

	"mom-platform/services/mdm-service/internal/model"
)

type MaterialRepository interface {
	Create(ctx context.Context, m *model.Material) error
	Update(ctx context.Context, m *model.Material) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Material, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.Material, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.Material, int64, error)
}

type BomRepository interface {
	Create(ctx context.Context, b *model.Bom) error
	Update(ctx context.Context, b *model.Bom) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Bom, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.Bom, error)
	GetActiveBom(ctx context.Context, tenantID string, materialID uint64) (*model.Bom, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.Bom, int64, error)
	UpdateStatus(ctx context.Context, id uint64, status string) error
	DeleteItems(ctx context.Context, bomID uint64) error
	CreateItems(ctx context.Context, items []model.BomItem) error
	GetItems(ctx context.Context, bomID uint64) ([]model.BomItem, error)
}

type WorkshopRepository interface {
	Create(ctx context.Context, w *model.Workshop) error
	Update(ctx context.Context, w *model.Workshop) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Workshop, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.Workshop, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.Workshop, int64, error)
}

type ProductionLineRepository interface {
	Create(ctx context.Context, pl *model.ProductionLine) error
	Update(ctx context.Context, pl *model.ProductionLine) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.ProductionLine, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.ProductionLine, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.ProductionLine, int64, error)
	ListByWorkshop(ctx context.Context, tenantID string, workshopID uint64) ([]model.ProductionLine, error)
}

type WorkstationRepository interface {
	Create(ctx context.Context, ws *model.Workstation) error
	Update(ctx context.Context, ws *model.Workstation) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Workstation, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.Workstation, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.Workstation, int64, error)
}

type CustomerRepository interface {
	Create(ctx context.Context, c *model.Customer) error
	Update(ctx context.Context, c *model.Customer) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Customer, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.Customer, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.Customer, int64, error)
}

type SupplierRepository interface {
	Create(ctx context.Context, s *model.Supplier) error
	Update(ctx context.Context, s *model.Supplier) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Supplier, error)
	GetByCode(ctx context.Context, tenantID, code string) (*model.Supplier, error)
	List(ctx context.Context, tenantID string, offset, limit int) ([]model.Supplier, int64, error)
}
