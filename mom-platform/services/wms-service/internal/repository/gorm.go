package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"mom-platform/services/wms-service/internal/model"
)

// txKey is the context key used to carry the current transaction.
type txKey struct{}

// txFromCtx extracts a transaction from the context; returns fallback db when absent.
func txFromCtx(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok && tx != nil {
		return tx
	}
	return db
}

// Ensure gormRepository implements Repository.
var _ Repository = (*gormRepository)(nil)

// gormRepository is the GORM-backed implementation of Repository.
type gormRepository struct {
	db *gorm.DB
}

// New creates a new GORM-backed Repository.
func New(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

// WithTx executes fn within a database transaction. The transaction is stored
// in the context so all repository calls inside fn participate in the same tx.
func (r *gormRepository) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txKey{}, tx)
		return fn(txCtx)
	})
}

// --- Warehouse ---

func (r *gormRepository) CreateWarehouse(ctx context.Context, w *model.Warehouse) error {
	return txFromCtx(ctx, r.db).Create(w).Error
}

func (r *gormRepository) GetWarehouse(ctx context.Context, id uint) (*model.Warehouse, error) {
	var w model.Warehouse
	if err := txFromCtx(ctx, r.db).First(&w, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &w, nil
}

func (r *gormRepository) GetWarehouseByCode(ctx context.Context, tenantID uint, code string) (*model.Warehouse, error) {
	var w model.Warehouse
	err := txFromCtx(ctx, r.db).
		Where("tenant_id = ? AND warehouse_code = ?", tenantID, code).
		First(&w).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &w, nil
}

func (r *gormRepository) ListWarehouses(ctx context.Context, tenantID uint) ([]*model.Warehouse, error) {
	var list []*model.Warehouse
	q := txFromCtx(ctx, r.db)
	if tenantID != 0 {
		q = q.Where("tenant_id = ?", tenantID)
	}
	err := q.Find(&list).Error
	return list, err
}

// --- Location ---

func (r *gormRepository) CreateLocation(ctx context.Context, l *model.Location) error {
	return txFromCtx(ctx, r.db).Create(l).Error
}

func (r *gormRepository) GetLocation(ctx context.Context, id uint) (*model.Location, error) {
	var l model.Location
	if err := txFromCtx(ctx, r.db).First(&l, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &l, nil
}

func (r *gormRepository) GetLocationByCode(ctx context.Context, warehouseID uint, code string) (*model.Location, error) {
	var l model.Location
	err := txFromCtx(ctx, r.db).
		Where("warehouse_id = ? AND location_code = ?", warehouseID, code).
		First(&l).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &l, nil
}

func (r *gormRepository) ListLocations(ctx context.Context, warehouseID, areaID uint) ([]*model.Location, error) {
	var list []*model.Location
	q := txFromCtx(ctx, r.db)
	if warehouseID > 0 {
		q = q.Where("warehouse_id = ?", warehouseID)
	}
	if areaID > 0 {
		q = q.Where("area_id = ?", areaID)
	}
	err := q.Find(&list).Error
	return list, err
}

func (r *gormRepository) UpdateLocationUsedCapacity(ctx context.Context, id uint, delta float64) error {
	result := txFromCtx(ctx, r.db).Model(&model.Location{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"used_capacity": gorm.Expr("used_capacity + ?", delta),
			"updated_at":    time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// --- InventoryBalance ---

func (r *gormRepository) CreateInventoryBalance(ctx context.Context, ib *model.InventoryBalance) error {
	return txFromCtx(ctx, r.db).Create(ib).Error
}

func (r *gormRepository) GetInventoryBalance(ctx context.Context, id uint) (*model.InventoryBalance, error) {
	var ib model.InventoryBalance
	if err := txFromCtx(ctx, r.db).First(&ib, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &ib, nil
}

func (r *gormRepository) GetInventoryBalanceForUpdate(ctx context.Context, id uint) (*model.InventoryBalance, error) {
	var ib model.InventoryBalance
	err := txFromCtx(ctx, r.db).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&ib, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &ib, nil
}

func (r *gormRepository) GetInventoryBalanceByMaterialLocation(
	ctx context.Context, tenantID, materialID, locationID uint, batchNo string,
) (*model.InventoryBalance, error) {
	var ib model.InventoryBalance
	query := txFromCtx(ctx, r.db).
		Where("tenant_id = ? AND material_id = ? AND location_id = ?", tenantID, materialID, locationID)
	if batchNo != "" {
		query = query.Where("batch_no = ?", batchNo)
	} else {
		query = query.Where("batch_no = '' OR batch_no IS NULL")
	}
	err := query.First(&ib).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &ib, nil
}

func (r *gormRepository) ListInventoryByMaterial(ctx context.Context, tenantID, materialID uint) ([]*model.InventoryBalance, error) {
	var list []*model.InventoryBalance
	q := txFromCtx(ctx, r.db)
	if tenantID != 0 {
		q = q.Where("tenant_id = ?", tenantID)
	}
	err := q.Where("material_id = ?", materialID).Find(&list).Error
	return list, err
}

func (r *gormRepository) ListBalances(ctx context.Context, tenantID, materialID uint, status model.InventoryStatus) ([]*model.InventoryBalance, error) {
	var list []*model.InventoryBalance
	q := txFromCtx(ctx, r.db)
	if tenantID > 0 {
		q = q.Where("tenant_id = ?", tenantID)
	}
	if materialID > 0 {
		q = q.Where("material_id = ?", materialID)
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}
	err := q.Find(&list).Error
	return list, err
}

// IncrementInventoryWithCost adds qty and recomputes unit_cost by moving
// weighted average: (oldQty*oldCost + qty*unitPrice) / (oldQty + qty).
//
// Postgres evaluates every SET expression against the pre-update row, so the
// cost formula safely references the old quantity/unit_cost in the same
// statement that increments quantity. A non-positive unitPrice means "price
// unknown" and leaves the existing cost untouched rather than diluting it to 0.
func (r *gormRepository) IncrementInventoryWithCost(ctx context.Context, id uint, qty, unitPrice float64) error {
	updates := map[string]interface{}{
		"quantity":      gorm.Expr("quantity + ?", qty),
		"available_qty": gorm.Expr("available_qty + ?", qty),
		"status":        gorm.Expr("CASE WHEN locked_qty <= 0 THEN ? ELSE status END", model.InventoryStatusNormal),
		"updated_at":    time.Now(),
	}
	if unitPrice > 0 {
		// 入库金额在 Go 侧先算好：SQL 里写成 "? * ?" 会让 Postgres 面对
		// unknown * unknown 而报 42725（operator is not unique）。
		inboundValue := qty * unitPrice
		updates["unit_cost"] = gorm.Expr(
			"CASE WHEN quantity + ? > 0 THEN (quantity * unit_cost + ?) / (quantity + ?) ELSE unit_cost END",
			qty, inboundValue, qty,
		)
	}

	result := txFromCtx(ctx, r.db).Session(&gorm.Session{SkipHooks: true}).
		Model(&model.InventoryBalance{}).
		Where("id = ?", id).
		Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// IncrementInventoryQuantity adds qty to both quantity and available_qty.
// Hooks are skipped because the invariant is maintained via SQL expressions.
func (r *gormRepository) IncrementInventoryQuantity(ctx context.Context, id uint, qty float64) error {
	result := txFromCtx(ctx, r.db).Session(&gorm.Session{SkipHooks: true}).
		Model(&model.InventoryBalance{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"quantity":      gorm.Expr("quantity + ?", qty),
			"available_qty": gorm.Expr("available_qty + ?", qty),
			"status":        gorm.Expr("CASE WHEN locked_qty <= 0 THEN ? ELSE status END", model.InventoryStatusNormal),
			"updated_at":    time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// LockInventory atomically decreases available_qty and increases locked_qty.
// Returns ErrInsufficientInventory when available_qty is insufficient.
func (r *gormRepository) LockInventory(ctx context.Context, id uint, qty float64) error {
	result := txFromCtx(ctx, r.db).Session(&gorm.Session{SkipHooks: true}).
		Model(&model.InventoryBalance{}).
		Where("id = ? AND available_qty >= ?", id, qty).
		Updates(map[string]interface{}{
			"locked_qty":    gorm.Expr("locked_qty + ?", qty),
			"available_qty": gorm.Expr("available_qty - ?", qty),
			"status":        gorm.Expr("CASE WHEN available_qty - ? <= 0 THEN ? ELSE status END",
				qty, model.InventoryStatusLocked),
			"updated_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrInsufficientInventory
	}
	return nil
}

// UnlockInventory atomically increases available_qty and decreases locked_qty.
// Returns ErrInsufficientLockedQty when locked_qty is insufficient.
func (r *gormRepository) UnlockInventory(ctx context.Context, id uint, qty float64) error {
	result := txFromCtx(ctx, r.db).Session(&gorm.Session{SkipHooks: true}).
		Model(&model.InventoryBalance{}).
		Where("id = ? AND locked_qty >= ?", id, qty).
		Updates(map[string]interface{}{
			"locked_qty":    gorm.Expr("locked_qty - ?", qty),
			"available_qty": gorm.Expr("available_qty + ?", qty),
			"status":        gorm.Expr("CASE WHEN locked_qty - ? <= 0 THEN ? ELSE status END",
				qty, model.InventoryStatusNormal),
			"updated_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrInsufficientLockedQty
	}
	return nil
}

// DeductLockedInventory removes qty from both quantity and locked_qty.
// Called on shipment to physically remove picked goods from inventory.
func (r *gormRepository) DeductLockedInventory(ctx context.Context, id uint, qty float64) error {
	result := txFromCtx(ctx, r.db).Session(&gorm.Session{SkipHooks: true}).
		Model(&model.InventoryBalance{}).
		Where("id = ? AND locked_qty >= ? AND quantity >= ?", id, qty, qty).
		Updates(map[string]interface{}{
			"quantity":     gorm.Expr("quantity - ?", qty),
			"locked_qty":   gorm.Expr("locked_qty - ?", qty),
			"status":       gorm.Expr("CASE WHEN locked_qty - ? <= 0 THEN ? ELSE status END",
				qty, model.InventoryStatusNormal),
			"updated_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrInsufficientLockedQty
	}
	return nil
}

// --- ReceiveOrder ---

func (r *gormRepository) CreateReceiveOrder(ctx context.Context, ro *model.ReceiveOrder) error {
	return txFromCtx(ctx, r.db).Create(ro).Error
}

func (r *gormRepository) GetReceiveOrder(ctx context.Context, id uint) (*model.ReceiveOrder, error) {
	var ro model.ReceiveOrder
	if err := txFromCtx(ctx, r.db).First(&ro, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &ro, nil
}

func (r *gormRepository) GetReceiveOrderByNo(ctx context.Context, tenantID uint, no string) (*model.ReceiveOrder, error) {
	var ro model.ReceiveOrder
	err := txFromCtx(ctx, r.db).
		Where("tenant_id = ? AND receive_no = ?", tenantID, no).
		First(&ro).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &ro, nil
}

func (r *gormRepository) UpdateReceiveOrder(ctx context.Context, ro *model.ReceiveOrder) error {
	return txFromCtx(ctx, r.db).Save(ro).Error
}

func (r *gormRepository) UpdateReceiveOrderStatus(ctx context.Context, id uint, status model.ReceiveOrderStatus) error {
	result := txFromCtx(ctx, r.db).Model(&model.ReceiveOrder{}).
		Where("id = ?", id).
		Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *gormRepository) ListReceiveOrders(ctx context.Context, status model.ReceiveOrderStatus, supplierID uint) ([]*model.ReceiveOrder, error) {
	var list []*model.ReceiveOrder
	q := txFromCtx(ctx, r.db)
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if supplierID > 0 {
		q = q.Where("supplier_id = ?", supplierID)
	}
	err := q.Find(&list).Error
	return list, err
}

// --- ReceiveOrderLine ---

func (r *gormRepository) CreateReceiveOrderLine(ctx context.Context, rol *model.ReceiveOrderLine) error {
	return txFromCtx(ctx, r.db).Create(rol).Error
}

func (r *gormRepository) GetReceiveOrderLine(ctx context.Context, id uint) (*model.ReceiveOrderLine, error) {
	var rol model.ReceiveOrderLine
	if err := txFromCtx(ctx, r.db).First(&rol, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &rol, nil
}

func (r *gormRepository) ListReceiveOrderLines(ctx context.Context, receiveOrderID uint) ([]*model.ReceiveOrderLine, error) {
	var list []*model.ReceiveOrderLine
	err := txFromCtx(ctx, r.db).Where("receive_order_id = ?", receiveOrderID).Find(&list).Error
	return list, err
}

func (r *gormRepository) UpdateReceiveOrderLineReceivedQty(ctx context.Context, id uint, receivedQty float64) error {
	result := txFromCtx(ctx, r.db).Model(&model.ReceiveOrderLine{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"received_qty": receivedQty,
			"updated_at":   time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// --- DeliveryOrder ---

func (r *gormRepository) CreateDeliveryOrder(ctx context.Context, do *model.DeliveryOrder) error {
	return txFromCtx(ctx, r.db).Create(do).Error
}

func (r *gormRepository) GetDeliveryOrder(ctx context.Context, id uint) (*model.DeliveryOrder, error) {
	var do model.DeliveryOrder
	if err := txFromCtx(ctx, r.db).First(&do, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &do, nil
}

func (r *gormRepository) GetDeliveryOrderByNo(ctx context.Context, tenantID uint, no string) (*model.DeliveryOrder, error) {
	var do model.DeliveryOrder
	err := txFromCtx(ctx, r.db).
		Where("tenant_id = ? AND delivery_no = ?", tenantID, no).
		First(&do).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &do, nil
}

func (r *gormRepository) UpdateDeliveryOrder(ctx context.Context, do *model.DeliveryOrder) error {
	return txFromCtx(ctx, r.db).Save(do).Error
}

func (r *gormRepository) UpdateDeliveryOrderStatus(ctx context.Context, id uint, status model.DeliveryOrderStatus) error {
	result := txFromCtx(ctx, r.db).Model(&model.DeliveryOrder{}).
		Where("id = ?", id).
		Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *gormRepository) ListDeliveryOrders(ctx context.Context, status model.DeliveryOrderStatus, customerID uint) ([]*model.DeliveryOrder, error) {
	var list []*model.DeliveryOrder
	q := txFromCtx(ctx, r.db)
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if customerID > 0 {
		q = q.Where("customer_id = ?", customerID)
	}
	err := q.Find(&list).Error
	return list, err
}

// --- DeliveryOrderLine ---

func (r *gormRepository) CreateDeliveryOrderLine(ctx context.Context, dol *model.DeliveryOrderLine) error {
	return txFromCtx(ctx, r.db).Create(dol).Error
}

func (r *gormRepository) GetDeliveryOrderLine(ctx context.Context, id uint) (*model.DeliveryOrderLine, error) {
	var dol model.DeliveryOrderLine
	if err := txFromCtx(ctx, r.db).First(&dol, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &dol, nil
}

func (r *gormRepository) ListDeliveryOrderLines(ctx context.Context, deliveryOrderID uint) ([]*model.DeliveryOrderLine, error) {
	var list []*model.DeliveryOrderLine
	err := txFromCtx(ctx, r.db).Where("delivery_order_id = ?", deliveryOrderID).Find(&list).Error
	return list, err
}

func (r *gormRepository) UpdateDeliveryOrderLinePickedQty(ctx context.Context, id uint, pickedQty float64) error {
	result := txFromCtx(ctx, r.db).Model(&model.DeliveryOrderLine{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"picked_qty":  pickedQty,
			"updated_at":  time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// --- PutawayRecord ---

func (r *gormRepository) CreatePutawayRecord(ctx context.Context, pr *model.PutawayRecord) error {
	return txFromCtx(ctx, r.db).Create(pr).Error
}

func (r *gormRepository) ListPutawayRecords(ctx context.Context, receiveOrderID uint) ([]*model.PutawayRecord, error) {
	var list []*model.PutawayRecord
	err := txFromCtx(ctx, r.db).Where("receive_order_id = ?", receiveOrderID).Find(&list).Error
	return list, err
}

// --- PickRecord ---

func (r *gormRepository) CreatePickRecord(ctx context.Context, pr *model.PickRecord) error {
	return txFromCtx(ctx, r.db).Create(pr).Error
}

func (r *gormRepository) ListPickRecords(ctx context.Context, deliveryOrderID uint) ([]*model.PickRecord, error) {
	var list []*model.PickRecord
	err := txFromCtx(ctx, r.db).Where("delivery_order_id = ?", deliveryOrderID).Find(&list).Error
	return list, err
}

// --- CountPlan ---

func (r *gormRepository) CreateCountPlan(ctx context.Context, cp *model.CountPlan) error {
	return txFromCtx(ctx, r.db).Create(cp).Error
}

func (r *gormRepository) GetCountPlan(ctx context.Context, id uint) (*model.CountPlan, error) {
	var cp model.CountPlan
	if err := txFromCtx(ctx, r.db).First(&cp, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &cp, nil
}

func (r *gormRepository) UpdateCountPlan(ctx context.Context, cp *model.CountPlan) error {
	return txFromCtx(ctx, r.db).Save(cp).Error
}

// --- CountRecord ---

func (r *gormRepository) CreateCountRecord(ctx context.Context, cr *model.CountRecord) error {
	cr.DiffQty = cr.ActualQty - cr.SystemQty
	return txFromCtx(ctx, r.db).Create(cr).Error
}

func (r *gormRepository) UpdateCountRecord(ctx context.Context, cr *model.CountRecord) error {
	cr.DiffQty = cr.ActualQty - cr.SystemQty
	return txFromCtx(ctx, r.db).Save(cr).Error
}

func (r *gormRepository) GetCountRecord(ctx context.Context, id uint) (*model.CountRecord, error) {
	var cr model.CountRecord
	if err := txFromCtx(ctx, r.db).First(&cr, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &cr, nil
}
