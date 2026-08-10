package repository

import (
	"context"

	"github.com/ninghonggang/mom-platform/services/mes-service/internal/model"
)

// OrderRepository — 工单仓储接口
type OrderRepository interface {
	GetByID(ctx context.Context, id int64) (*model.ProductionOrder, error)
	List(ctx context.Context, filter OrderFilter) ([]model.ProductionOrder, int64, error)
	Create(ctx context.Context, order *model.ProductionOrder) error
	UpdateStatus(ctx context.Context, id int64, status model.OrderStatus) error
	Update(ctx context.Context, order *model.ProductionOrder) error
}

type OrderFilter struct {
	TenantID   int64
	Keyword    string
	WorkshopID int64
	LineID     int64
	Status     *model.OrderStatus
	DateFrom   string
	DateTo     string
	Page       int
	PageSize   int
}

// JobReportRepository — 报工仓储接口
type JobReportRepository interface {
	GetByID(ctx context.Context, id int64) (*model.MobileJobReport, error)
	List(ctx context.Context, filter ReportFilter) ([]model.MobileJobReport, int64, error)
	Create(ctx context.Context, report *model.MobileJobReport) error
	UpdateStatus(ctx context.Context, id int64, status model.ReportStatus) error
}

type ReportFilter struct {
	TenantID   int64
	OrderID    int64
	EmployeeID int64
	DateFrom   string
	DateTo     string
	Page       int
	PageSize   int
}

// DispatchRepository — 派工仓储接口
type DispatchRepository interface {
	List(ctx context.Context, filter DispatchFilter) ([]model.Dispatch, error)
	CreateBatch(ctx context.Context, dispatches []model.Dispatch) error
}

type DispatchFilter struct {
	TenantID      int64
	OrderID       int64
	LineID        int64
	WorkstationID int64
}

// CompleteRepository — 完工入库仓储接口
type CompleteRepository interface {
	Create(ctx context.Context, c *model.ProductionComplete) error
	GetByID(ctx context.Context, id int64) (*model.ProductionComplete, error)
	ListByOrder(ctx context.Context, orderID int64) ([]model.ProductionComplete, error)
	SumQtyByOrder(ctx context.Context, orderID int64) (float64, error)
}
