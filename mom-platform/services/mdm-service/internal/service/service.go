package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"mom-platform/services/mdm-service/internal/model"
	"mom-platform/services/mdm-service/internal/repository"
)

var (
	ErrDuplicateCode    = errors.New("duplicate code")
	ErrInvalidStatus    = errors.New("invalid status transition")
	ErrWorkshopNotFound = errors.New("workshop not found")
	ErrLineNotFound     = errors.New("production line not found")
)

type MDMService struct {
	logger          *zap.Logger
	db              *gorm.DB
	materialRepo    repository.MaterialRepository
	bomRepo         repository.BomRepository
	workshopRepo    repository.WorkshopRepository
	lineRepo        repository.ProductionLineRepository
	workstationRepo repository.WorkstationRepository
	customerRepo    repository.CustomerRepository
	supplierRepo    repository.SupplierRepository
	publisher       *eventbus.EventPublisher
}

func NewMDMService(
	logger *zap.Logger,
	db *gorm.DB,
	mr repository.MaterialRepository,
	br repository.BomRepository,
	wr repository.WorkshopRepository,
	lr repository.ProductionLineRepository,
	wsr repository.WorkstationRepository,
	cr repository.CustomerRepository,
	sr repository.SupplierRepository,
) *MDMService {
	return &MDMService{
		logger:          logger,
		db:              db,
		materialRepo:    mr,
		bomRepo:         br,
		workshopRepo:    wr,
		lineRepo:        lr,
		workstationRepo: wsr,
		customerRepo:    cr,
		supplierRepo:    sr,
	}
}

// SetPublisher injects an EventPublisher for domain event publishing.
func (s *MDMService) SetPublisher(p *eventbus.EventPublisher) {
	s.publisher = p
}

// ------------------- dbTransaction -------------------

func (s *MDMService) dbTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create new service with tx-pooled repositories
		txSvc := &MDMService{
			logger:          s.logger,
			db:              tx,
			materialRepo:    repository.NewMaterialRepo(tx),
			bomRepo:         repository.NewBomRepo(tx),
			workshopRepo:    repository.NewWorkshopRepo(tx),
			lineRepo:        repository.NewProductionLineRepo(tx),
			workstationRepo: repository.NewWorkstationRepo(tx),
			customerRepo:    repository.NewCustomerRepo(tx),
			supplierRepo:    repository.NewSupplierRepo(tx),
			publisher:       s.publisher,
		}
		return fn(context.WithValue(ctx, struct{}{}, txSvc))
	})
}

// ------------------- Material -------------------

func (s *MDMService) CreateMaterial(ctx context.Context, m *model.Material) error {
	existing, err := s.materialRepo.GetByCode(ctx, m.TenantID, m.MaterialCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: material_code=%s", ErrDuplicateCode, m.MaterialCode)
	}
	if err := s.materialRepo.Create(ctx, m); err != nil {
		return err
	}
	if s.publisher != nil {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMDMMaterialCreated, m)
	}
	return nil
}

func (s *MDMService) UpdateMaterial(ctx context.Context, m *model.Material) error {
	existing, err := s.materialRepo.GetByID(ctx, m.ID)
	if err != nil {
		return err
	}
	if existing.MaterialCode != m.MaterialCode {
		dup, err := s.materialRepo.GetByCode(ctx, m.TenantID, m.MaterialCode)
		if err != nil && !errors.Is(err, repository.ErrNotFound) {
			return err
		}
		if dup != nil {
			return fmt.Errorf("%w: material_code=%s", ErrDuplicateCode, m.MaterialCode)
		}
	}
	if err := s.materialRepo.Update(ctx, m); err != nil {
		return err
	}
	if s.publisher != nil && m.Status == "OBSOLETE" && existing.Status != "OBSOLETE" {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMDMMaterialObsolete, m)
	}
	return nil
}

func (s *MDMService) DeleteMaterial(ctx context.Context, id uint64) error {
	return s.materialRepo.Delete(ctx, id)
}

func (s *MDMService) GetMaterial(ctx context.Context, id uint64) (*model.Material, error) {
	return s.materialRepo.GetByID(ctx, id)
}

func (s *MDMService) ListMaterials(ctx context.Context, tenantID string, offset, limit int) ([]model.Material, int64, error) {
	return s.materialRepo.List(ctx, tenantID, offset, limit)
}

// ------------------- BOM -------------------

func (s *MDMService) CreateBom(ctx context.Context, b *model.Bom, items []model.BomItem) error {
	existing, err := s.bomRepo.GetByCode(ctx, b.TenantID, b.BomCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: bom_code=%s", ErrDuplicateCode, b.BomCode)
	}

	return s.dbTransaction(ctx, func(txCtx context.Context) error {
		if err := s.bomRepo.Create(txCtx, b); err != nil {
			return err
		}
		if len(items) > 0 {
			for i := range items {
				items[i].BomID = b.ID
			}
			if err := s.bomRepo.CreateItems(txCtx, items); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *MDMService) UpdateBom(ctx context.Context, b *model.Bom, items []model.BomItem) error {
	return s.dbTransaction(ctx, func(txCtx context.Context) error {
		if err := s.bomRepo.Update(txCtx, b); err != nil {
			return err
		}
		if err := s.bomRepo.DeleteItems(txCtx, b.ID); err != nil {
			return err
		}
		if len(items) > 0 {
			for i := range items {
				items[i].BomID = b.ID
			}
			return s.bomRepo.CreateItems(txCtx, items)
		}
		return nil
	})
}

var validStatusTransitions = map[string][]string{
	"DRAFT":    {"ACTIVE", "OBSOLETE"},
	"ACTIVE":   {"OBSOLETE"},
	"OBSOLETE": {},
}

func (s *MDMService) UpdateBomStatus(ctx context.Context, id uint64, newStatus string) error {
	bom, err := s.bomRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	allowed := validStatusTransitions[bom.Status]
	valid := false
	for _, st := range allowed {
		if st == newStatus {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("%w: cannot transition from %s to %s", ErrInvalidStatus, bom.Status, newStatus)
	}

	if err := s.bomRepo.UpdateStatus(ctx, id, newStatus); err != nil {
		return err
	}
	if s.publisher != nil && newStatus == "ACTIVE" {
		bom.Status = newStatus
		_ = s.publisher.Publish(ctx, eventbus.SubjectMDMBomActivated, bom)
	}
	return nil
}

func (s *MDMService) GetBom(ctx context.Context, id uint64) (*model.Bom, error) {
	return s.bomRepo.GetByID(ctx, id)
}

func (s *MDMService) ListBoms(ctx context.Context, tenantID string, offset, limit int) ([]model.Bom, int64, error) {
	return s.bomRepo.List(ctx, tenantID, offset, limit)
}

// BomExplosionResult holds the result of a BOM explosion
type BomExplosionResult struct {
	MaterialID       uint64  `json:"material_id"`
	MaterialCode     string  `json:"material_code"`
	Quantity         float64 `json:"quantity"`
	Unit             string  `json:"unit"`
	Level            int     `json:"level"`
	ParentMaterialID uint64  `json:"parent_material_id,omitempty"`
}

// ExplodeBom recursively explodes a BOM with quantity multiplication
func (s *MDMService) ExplodeBom(ctx context.Context, tenantID string, materialID uint64, baseQty float64) ([]BomExplosionResult, error) {
	var result []BomExplosionResult

	_, err := s.materialRepo.GetByID(ctx, materialID)
	if err != nil {
		return nil, fmt.Errorf("material %d: %w", materialID, err)
	}

	if err := s.explodeRecursive(ctx, tenantID, materialID, baseQty, 0, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *MDMService) explodeRecursive(ctx context.Context, tenantID string, materialID uint64, qty float64, level int, result *[]BomExplosionResult) error {
	if level > 50 {
		return errors.New("BOM explosion exceeded maximum recursion depth")
	}

	bom, err := s.bomRepo.GetActiveBom(ctx, tenantID, materialID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil // leaf material, no child components
		}
		return err
	}

	for _, item := range bom.Items {
		childQty := qty * item.Quantity * (1 + item.ScrapRate/100.0)

		childMaterial, err := s.materialRepo.GetByID(ctx, item.ChildMaterialID)
		if err != nil {
			s.logger.Warn("child material not found, skipping",
				zap.Uint64("material_id", item.ChildMaterialID),
				zap.Error(err))
			continue
		}

		*result = append(*result, BomExplosionResult{
			MaterialID:       item.ChildMaterialID,
			MaterialCode:     childMaterial.MaterialCode,
			Quantity:         childQty,
			Unit:             item.Unit,
			Level:            level + 1,
			ParentMaterialID: materialID,
		})

		if err := s.explodeRecursive(ctx, tenantID, item.ChildMaterialID, childQty, level+1, result); err != nil {
			return err
		}
	}

	return nil
}

// ------------------- Workshop -------------------

func (s *MDMService) CreateWorkshop(ctx context.Context, w *model.Workshop) error {
	existing, err := s.workshopRepo.GetByCode(ctx, w.TenantID, w.WorkshopCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: workshop_code=%s", ErrDuplicateCode, w.WorkshopCode)
	}
	return s.workshopRepo.Create(ctx, w)
}

func (s *MDMService) UpdateWorkshop(ctx context.Context, w *model.Workshop) error {
	return s.workshopRepo.Update(ctx, w)
}

func (s *MDMService) DeleteWorkshop(ctx context.Context, id uint64) error {
	return s.workshopRepo.Delete(ctx, id)
}

func (s *MDMService) GetWorkshop(ctx context.Context, id uint64) (*model.Workshop, error) {
	return s.workshopRepo.GetByID(ctx, id)
}

func (s *MDMService) ListWorkshops(ctx context.Context, tenantID string, offset, limit int) ([]model.Workshop, int64, error) {
	return s.workshopRepo.List(ctx, tenantID, offset, limit)
}

// ------------------- Production Line -------------------

func (s *MDMService) CreateProductionLine(ctx context.Context, pl *model.ProductionLine) error {
	_, err := s.workshopRepo.GetByID(ctx, pl.WorkshopID)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrWorkshopNotFound, err)
	}
	existing, err := s.lineRepo.GetByCode(ctx, pl.TenantID, pl.LineCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: line_code=%s", ErrDuplicateCode, pl.LineCode)
	}
	return s.lineRepo.Create(ctx, pl)
}

func (s *MDMService) UpdateProductionLine(ctx context.Context, pl *model.ProductionLine) error {
	return s.lineRepo.Update(ctx, pl)
}

func (s *MDMService) DeleteProductionLine(ctx context.Context, id uint64) error {
	return s.lineRepo.Delete(ctx, id)
}

func (s *MDMService) GetProductionLine(ctx context.Context, id uint64) (*model.ProductionLine, error) {
	return s.lineRepo.GetByID(ctx, id)
}

func (s *MDMService) ListProductionLines(ctx context.Context, tenantID string, offset, limit int) ([]model.ProductionLine, int64, error) {
	return s.lineRepo.List(ctx, tenantID, offset, limit)
}

// ------------------- Workstation -------------------

func (s *MDMService) CreateWorkstation(ctx context.Context, ws *model.Workstation) error {
	if ws.LineID > 0 {
		_, err := s.lineRepo.GetByID(ctx, ws.LineID)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrLineNotFound, err)
		}
	}
	existing, err := s.workstationRepo.GetByCode(ctx, ws.TenantID, ws.WorkstationCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: workstation_code=%s", ErrDuplicateCode, ws.WorkstationCode)
	}
	return s.workstationRepo.Create(ctx, ws)
}

func (s *MDMService) UpdateWorkstation(ctx context.Context, ws *model.Workstation) error {
	return s.workstationRepo.Update(ctx, ws)
}

func (s *MDMService) DeleteWorkstation(ctx context.Context, id uint64) error {
	return s.workstationRepo.Delete(ctx, id)
}

func (s *MDMService) GetWorkstation(ctx context.Context, id uint64) (*model.Workstation, error) {
	return s.workstationRepo.GetByID(ctx, id)
}

func (s *MDMService) ListWorkstations(ctx context.Context, tenantID string, offset, limit int) ([]model.Workstation, int64, error) {
	return s.workstationRepo.List(ctx, tenantID, offset, limit)
}

// ------------------- Customer -------------------

func (s *MDMService) CreateCustomer(ctx context.Context, c *model.Customer) error {
	existing, err := s.customerRepo.GetByCode(ctx, c.TenantID, c.CustomerCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: customer_code=%s", ErrDuplicateCode, c.CustomerCode)
	}
	return s.customerRepo.Create(ctx, c)
}

func (s *MDMService) UpdateCustomer(ctx context.Context, c *model.Customer) error {
	return s.customerRepo.Update(ctx, c)
}

func (s *MDMService) DeleteCustomer(ctx context.Context, id uint64) error {
	return s.customerRepo.Delete(ctx, id)
}

func (s *MDMService) GetCustomer(ctx context.Context, id uint64) (*model.Customer, error) {
	return s.customerRepo.GetByID(ctx, id)
}

func (s *MDMService) ListCustomers(ctx context.Context, tenantID string, offset, limit int) ([]model.Customer, int64, error) {
	return s.customerRepo.List(ctx, tenantID, offset, limit)
}

// ------------------- Supplier -------------------

func (s *MDMService) CreateSupplier(ctx context.Context, sup *model.Supplier) error {
	existing, err := s.supplierRepo.GetByCode(ctx, sup.TenantID, sup.SupplierCode)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("%w: supplier_code=%s", ErrDuplicateCode, sup.SupplierCode)
	}
	if err := s.supplierRepo.Create(ctx, sup); err != nil {
		return err
	}
	if s.publisher != nil && sup.Status == "ACTIVE" {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMDMSupplierActivated, sup)
	}
	return nil
}

func (s *MDMService) UpdateSupplier(ctx context.Context, sup *model.Supplier) error {
	existing, err := s.supplierRepo.GetByID(ctx, sup.ID)
	if err != nil {
		return err
	}
	if err := s.supplierRepo.Update(ctx, sup); err != nil {
		return err
	}
	if s.publisher != nil && sup.Status == "BLACKLIST" && existing.Status != "BLACKLIST" {
		_ = s.publisher.Publish(ctx, eventbus.SubjectMDMSupplierBlacklist, sup)
	}
	return nil
}

func (s *MDMService) DeleteSupplier(ctx context.Context, id uint64) error {
	return s.supplierRepo.Delete(ctx, id)
}

func (s *MDMService) GetSupplier(ctx context.Context, id uint64) (*model.Supplier, error) {
	return s.supplierRepo.GetByID(ctx, id)
}

func (s *MDMService) ListSuppliers(ctx context.Context, tenantID string, offset, limit int) ([]model.Supplier, int64, error) {
	return s.supplierRepo.List(ctx, tenantID, offset, limit)
}
