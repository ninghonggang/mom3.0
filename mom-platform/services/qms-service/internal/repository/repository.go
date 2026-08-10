package repository

import (
	"context"

	"mom-platform/services/qms-service/internal/model"
)

// PageQuery holds pagination and filtering parameters for list queries.
type PageQuery struct {
	Page     int
	PageSize int
	Filter   map[string]interface{}
}

// PageResult wraps a list of results with pagination metadata.
type PageResult[T any] struct {
	Items    []T
	Total    int64
	Page     int
	PageSize int
}

// Repository defines the data-access interface for the QMS service.
type Repository interface {
	// --- InspectionSheet ---
	CreateInspectionSheet(ctx context.Context, sheet *model.InspectionSheet) error
	GetInspectionSheet(ctx context.Context, id uint) (*model.InspectionSheet, error)
	UpdateInspectionSheet(ctx context.Context, sheet *model.InspectionSheet) error
	DeleteInspectionSheet(ctx context.Context, id uint) error
	ListInspectionSheets(ctx context.Context, q PageQuery) (*PageResult[model.InspectionSheet], error)

	// --- InspectionCharacteristic ---
	CreateInspectionCharacteristic(ctx context.Context, c *model.InspectionCharacteristic) error
	GetInspectionCharacteristic(ctx context.Context, id uint) (*model.InspectionCharacteristic, error)
	UpdateInspectionCharacteristic(ctx context.Context, c *model.InspectionCharacteristic) error
	DeleteInspectionCharacteristic(ctx context.Context, id uint) error
	ListInspectionCharacteristics(ctx context.Context, q PageQuery) (*PageResult[model.InspectionCharacteristic], error)

	// --- InspectionPlan ---
	CreateInspectionPlan(ctx context.Context, p *model.InspectionPlan) error
	GetInspectionPlan(ctx context.Context, id uint) (*model.InspectionPlan, error)
	UpdateInspectionPlan(ctx context.Context, p *model.InspectionPlan) error
	DeleteInspectionPlan(ctx context.Context, id uint) error
	ListInspectionPlans(ctx context.Context, q PageQuery) (*PageResult[model.InspectionPlan], error)

	// --- InspectionResult ---
	CreateInspectionResult(ctx context.Context, r *model.InspectionResult) error
	GetInspectionResult(ctx context.Context, id uint) (*model.InspectionResult, error)
	ListInspectionResults(ctx context.Context, q PageQuery) (*PageResult[model.InspectionResult], error)
	ListResultsBySheet(ctx context.Context, sheetID uint) ([]model.InspectionResult, error)

	// --- Ncr ---
	CreateNcr(ctx context.Context, n *model.Ncr) error
	GetNcr(ctx context.Context, id uint) (*model.Ncr, error)
	UpdateNcr(ctx context.Context, n *model.Ncr) error
	DeleteNcr(ctx context.Context, id uint) error
	ListNcrs(ctx context.Context, q PageQuery) (*PageResult[model.Ncr], error)

	// --- NcrAction ---
	CreateNcrAction(ctx context.Context, a *model.NcrAction) error
	GetNcrAction(ctx context.Context, id uint) (*model.NcrAction, error)
	ListNcrActions(ctx context.Context, ncrID uint) ([]model.NcrAction, error)

	// --- DefectCode ---
	CreateDefectCode(ctx context.Context, d *model.DefectCode) error
	GetDefectCode(ctx context.Context, id uint) (*model.DefectCode, error)
	UpdateDefectCode(ctx context.Context, d *model.DefectCode) error
	DeleteDefectCode(ctx context.Context, id uint) error
	ListDefectCodes(ctx context.Context, q PageQuery) (*PageResult[model.DefectCode], error)

	// --- SpcData ---
	CreateSpcData(ctx context.Context, s *model.SpcData) error
	GetSpcData(ctx context.Context, id uint) (*model.SpcData, error)
	ListSpcData(ctx context.Context, q PageQuery) (*PageResult[model.SpcData], error)
	ListSpcDataByChar(ctx context.Context, charID uint) ([]model.SpcData, error)
}
