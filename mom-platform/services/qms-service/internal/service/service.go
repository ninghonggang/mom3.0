package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

	"mom-platform/services/qms-service/internal/model"
	"mom-platform/services/qms-service/internal/repository"
)

// Sentinel errors returned by the service layer. Use errors.Is to check.
var (
	ErrNotFound            = errors.New("qms: not found")
	ErrInvalidInput        = errors.New("qms: invalid input")
	ErrInvalidTransition   = errors.New("qms: invalid state transition")
	ErrInspectionSheetNotFailed = errors.New("qms: inspection sheet is not FAILED")
)

// AQLConfig holds Acceptable Quality Level thresholds.
// When defectCount/sampleSize exceeds AQLThreshold the sheet auto-fails.
type AQLConfig struct {
	AQLThreshold float64 // e.g. 0.01 means 1% defect rate fails
}

// Service implements the QMS business-logic layer.
type Service struct {
	repo repository.Repository
	log  *zap.Logger
	aql  AQLConfig
}

// NewService creates a new Service.
func NewService(repo repository.Repository, log *zap.Logger, aql AQLConfig) *Service {
	return &Service{repo: repo, log: log, aql: aql}
}

// ==============================================================================
// InspectionSheet
// ==============================================================================

var validInspectionTypes = map[string]bool{
	"INCOMING":     true,
	"IN_PROCESS":   true,
	"FINAL":        true,
	"FIRST_ARTICLE": true,
}

// CreateInspectionSheet validates the type, sets status to PENDING,
// and generates a unique sheet number.
func (s *Service) CreateInspectionSheet(ctx context.Context, sheet *model.InspectionSheet) (*model.InspectionSheet, error) {
	if sheet.TenantID == "" {
		return nil, fmt.Errorf("%w: tenant_id is required", ErrInvalidInput)
	}
	if !validInspectionTypes[sheet.InspectionType] {
		return nil, fmt.Errorf("%w: invalid inspection_type %q", ErrInvalidInput, sheet.InspectionType)
	}
	if sheet.SampleSize <= 0 {
		return nil, fmt.Errorf("%w: sample_size must be positive", ErrInvalidInput)
	}

	sheet.Status = model.SheetStatusPending
	sheet.SheetNo = s.generateSheetNo(sheet.InspectionType)

	if err := s.repo.CreateInspectionSheet(ctx, sheet); err != nil {
		return nil, fmt.Errorf("create inspection sheet: %w", err)
	}

	s.log.Info("inspection sheet created",
		zap.Uint("id", sheet.ID),
		zap.String("sheet_no", sheet.SheetNo),
		zap.String("type", sheet.InspectionType),
	)
	return sheet, nil
}

// UpdateInspectionSheet validates state transitions before persisting.
// Allowed transitions:
//   PENDING    -> IN_PROGRESS
//   IN_PROGRESS -> PASSED | FAILED | WAIVED
//   PENDING     -> CANCELLED
//   IN_PROGRESS -> CANCELLED
func (s *Service) UpdateInspectionSheet(ctx context.Context, id uint, updates *model.InspectionSheet) (*model.InspectionSheet, error) {
	existing, err := s.repo.GetInspectionSheet(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get inspection sheet %d: %w", id, err)
	}
	if existing == nil {
		return nil, ErrNotFound
	}

	if updates.Status != "" && updates.Status != existing.Status {
		if !isValidSheetTransition(existing.Status, updates.Status) {
			return nil, fmt.Errorf("%w: %s -> %s", ErrInvalidTransition, existing.Status, updates.Status)
		}
		existing.Status = updates.Status
	}

	// Apply updatable fields if provided.
	if updates.InspectorID != "" {
		existing.InspectorID = updates.InspectorID
	}
	if updates.DefectCount != existing.DefectCount {
		existing.DefectCount = updates.DefectCount
	}
	if updates.InspectedAt != nil {
		existing.InspectedAt = updates.InspectedAt
	}
	if updates.SampleSize != 0 {
		existing.SampleSize = updates.SampleSize
	}
	if updates.MaterialID != "" {
		existing.MaterialID = updates.MaterialID
	}
	if updates.BatchID != "" {
		existing.BatchID = updates.BatchID
	}

	// AQL判定: if defect rate exceeds threshold, force FAILED (only when in progress).
	if existing.Status == model.SheetStatusInProgress && existing.SampleSize > 0 {
		if s.aql.AQLThreshold > 0 {
			rate := float64(existing.DefectCount) / float64(existing.SampleSize)
			if rate > s.aql.AQLThreshold {
				existing.Status = model.SheetStatusFailed
				s.log.Info("AQL auto-fail triggered",
					zap.Uint("sheet_id", id),
					zap.Float64("defect_rate", rate),
					zap.Float64("threshold", s.aql.AQLThreshold),
				)
			}
		}
	}

	if err := s.repo.UpdateInspectionSheet(ctx, existing); err != nil {
		return nil, fmt.Errorf("update inspection sheet %d: %w", id, err)
	}

	s.log.Info("inspection sheet updated",
		zap.Uint("id", id),
		zap.String("status", existing.Status),
	)
	return existing, nil
}

func isValidSheetTransition(from, to string) bool {
	transitions := map[string]map[string]bool{
		model.SheetStatusPending: {
			model.SheetStatusInProgress: true,
			model.SheetStatusCancelled:  true,
		},
		model.SheetStatusInProgress: {
			model.SheetStatusPassed:    true,
			model.SheetStatusFailed:    true,
			model.SheetStatusWaived:    true,
			model.SheetStatusCancelled: true,
		},
	}
	allowed, ok := transitions[from]
	if !ok {
		return false
	}
	return allowed[to]
}

// GetInspectionSheet retrieves a sheet by ID.
func (s *Service) GetInspectionSheet(ctx context.Context, id uint) (*model.InspectionSheet, error) {
	sheet, err := s.repo.GetInspectionSheet(ctx, id)
	if err != nil {
		return nil, err
	}
	if sheet == nil {
		return nil, ErrNotFound
	}
	return sheet, nil
}

// ListInspectionSheets lists sheets with pagination and filtering.
func (s *Service) ListInspectionSheets(ctx context.Context, page, pageSize int, filter map[string]interface{}) (*repository.PageResult[model.InspectionSheet], error) {
	return s.repo.ListInspectionSheets(ctx, repository.PageQuery{Page: page, PageSize: pageSize, Filter: filter})
}

// DeleteInspectionSheet soft-deletes a sheet.
func (s *Service) DeleteInspectionSheet(ctx context.Context, id uint) error {
	return s.repo.DeleteInspectionSheet(ctx, id)
}

// ==============================================================================
// InspectionResult
// ==============================================================================

// RecordInspectionResult creates a result record, auto-determining pass/fail
// based on the characteristic specification limits (USL/LSL).
func (s *Service) RecordInspectionResult(ctx context.Context, sheetID uint, charID uint, value string) (*model.InspectionResult, error) {
	// Validate the sheet exists.
	sheet, err := s.repo.GetInspectionSheet(ctx, sheetID)
	if err != nil {
		return nil, fmt.Errorf("get inspection sheet: %w", err)
	}
	if sheet == nil {
		return nil, fmt.Errorf("%w: inspection sheet %d not found", ErrNotFound, sheetID)
	}

	// Validate the characteristic exists.
	char, err := s.repo.GetInspectionCharacteristic(ctx, charID)
	if err != nil {
		return nil, fmt.Errorf("get inspection characteristic: %w", err)
	}
	if char == nil {
		return nil, fmt.Errorf("%w: characteristic %d not found", ErrNotFound, charID)
	}

	pass := s.evaluatePass(char, value)

	result := &model.InspectionResult{
		SheetID: sheetID,
		CharID:  charID,
		Value:   value,
		Pass:    pass,
	}

	if err := s.repo.CreateInspectionResult(ctx, result); err != nil {
		return nil, fmt.Errorf("create inspection result: %w", err)
	}

	// If the sheet is still PENDING, auto-advance to IN_PROGRESS.
	if sheet.Status == model.SheetStatusPending {
		sheet.Status = model.SheetStatusInProgress
		if sheet.InspectedAt == nil {
			now := time.Now()
			sheet.InspectedAt = &now
		}
		_ = s.repo.UpdateInspectionSheet(ctx, sheet)
	}

	s.log.Info("inspection result recorded",
		zap.Uint("result_id", result.ID),
		zap.Uint("sheet_id", sheetID),
		zap.Uint("char_id", charID),
		zap.Bool("pass", pass),
	)
	return result, nil
}

// evaluatePass determines pass/fail for a numeric value against USL/LSL.
// For non-numeric data types (BOOLEAN/TEXT) it checks for empty/failure markers.
func (s *Service) evaluatePass(char *model.InspectionCharacteristic, value string) bool {
	switch strings.ToUpper(char.DataType) {
	case "NUMERIC", "":
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			// Non-numeric value for numeric char -> fail.
			return false
		}
		if char.LSL != 0 && v < char.LSL {
			return false
		}
		if char.USL != 0 && v > char.USL {
			return false
		}
		return true
	case "BOOLEAN":
		// "true"/"1"/"pass" -> pass, anything else -> fail.
		v := strings.ToLower(strings.TrimSpace(value))
		return v == "true" || v == "1" || v == "pass"
	case "TEXT":
		// Non-empty text passes.
		return strings.TrimSpace(value) != ""
	default:
		return false
	}
}

// ListInspectionResults lists results for a sheet.
func (s *Service) ListInspectionResults(ctx context.Context, sheetID uint) ([]model.InspectionResult, error) {
	return s.repo.ListResultsBySheet(ctx, sheetID)
}

// ==============================================================================
// Ncr (Non-Conformance Report)
// ==============================================================================

// CreateNcr validates that the referenced inspection sheet exists and is FAILED,
// then creates the NCR with status OPEN.
func (s *Service) CreateNcr(ctx context.Context, ncr *model.Ncr) (*model.Ncr, error) {
	if ncr.TenantID == "" {
		return nil, fmt.Errorf("%w: tenant_id is required", ErrInvalidInput)
	}
	if ncr.InspectionSheetID == 0 {
		return nil, fmt.Errorf("%w: inspection_sheet_id is required", ErrInvalidInput)
	}
	if ncr.Quantity <= 0 {
		return nil, fmt.Errorf("%w: quantity must be positive", ErrInvalidInput)
	}

	sheet, err := s.repo.GetInspectionSheet(ctx, ncr.InspectionSheetID)
	if err != nil {
		return nil, fmt.Errorf("get inspection sheet: %w", err)
	}
	if sheet == nil {
		return nil, fmt.Errorf("%w: inspection sheet %d not found", ErrNotFound, ncr.InspectionSheetID)
	}
	if sheet.Status != model.SheetStatusFailed {
		return nil, fmt.Errorf("%w: inspection sheet %d status is %s, must be %s",
			ErrInspectionSheetNotFailed, ncr.InspectionSheetID, sheet.Status, model.SheetStatusFailed)
	}

	// Populate material/batch from the sheet if not set.
	if ncr.MaterialID == "" {
		ncr.MaterialID = sheet.MaterialID
	}
	if ncr.BatchID == "" {
		ncr.BatchID = sheet.BatchID
	}

	ncr.NcrNo = s.generateNcrNo()
	ncr.Status = model.NcrStatusOpen

	if err := s.repo.CreateNcr(ctx, ncr); err != nil {
		return nil, fmt.Errorf("create ncr: %w", err)
	}

	s.log.Info("ncr created",
		zap.Uint("id", ncr.ID),
		zap.String("ncr_no", ncr.NcrNo),
		zap.Uint("sheet_id", ncr.InspectionSheetID),
	)
	return ncr, nil
}

// UpdateNcr validates NCR state transitions.
// Allowed transitions:
//   OPEN          -> INVESTIGATING
//   INVESTIGATING -> DISPOSITIONED
//   DISPOSITIONED -> VERIFIED
//   VERIFIED      -> CLOSED
//   CLOSED        -> REOPENED
//   OPEN..DISPOSITIONED -> CANCELLED
func (s *Service) UpdateNcr(ctx context.Context, id uint, updates *model.Ncr) (*model.Ncr, error) {
	existing, err := s.repo.GetNcr(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get ncr %d: %w", id, err)
	}
	if existing == nil {
		return nil, ErrNotFound
	}

	if updates.Status != "" && updates.Status != existing.Status {
		if !isValidNcrTransition(existing.Status, updates.Status) {
			return nil, fmt.Errorf("%w: %s -> %s", ErrInvalidTransition, existing.Status, updates.Status)
		}
		existing.Status = updates.Status
	}

	if updates.Severity != "" {
		existing.Severity = updates.Severity
	}
	if updates.Quantity > 0 {
		existing.Quantity = updates.Quantity
	}

	if err := s.repo.UpdateNcr(ctx, existing); err != nil {
		return nil, fmt.Errorf("update ncr %d: %w", id, err)
	}

	s.log.Info("ncr updated", zap.Uint("id", id), zap.String("status", existing.Status))
	return existing, nil
}

func isValidNcrTransition(from, to string) bool {
	transitions := map[string]map[string]bool{
		model.NcrStatusOpen: {
			model.NcrStatusInvestigating: true,
			model.NcrStatusCancelled:     true,
		},
		model.NcrStatusInvestigating: {
			model.NcrStatusDispositioned: true,
			model.NcrStatusCancelled:     true,
		},
		model.NcrStatusDispositioned: {
			model.NcrStatusVerified:  true,
			model.NcrStatusCancelled: true,
		},
		model.NcrStatusVerified: {
			model.NcrStatusClosed: true,
		},
		model.NcrStatusClosed: {
			model.NcrStatusReopened: true,
		},
		model.NcrStatusReopened: {
			model.NcrStatusInvestigating: true,
		},
	}
	allowed, ok := transitions[from]
	if !ok {
		return false
	}
	return allowed[to]
}

// GetNcr retrieves an NCR by ID.
func (s *Service) GetNcr(ctx context.Context, id uint) (*model.Ncr, error) {
	ncr, err := s.repo.GetNcr(ctx, id)
	if err != nil {
		return nil, err
	}
	if ncr == nil {
		return nil, ErrNotFound
	}
	return ncr, nil
}

// ListNcrs lists NCRs with pagination and filtering.
func (s *Service) ListNcrs(ctx context.Context, page, pageSize int, filter map[string]interface{}) (*repository.PageResult[model.Ncr], error) {
	return s.repo.ListNcrs(ctx, repository.PageQuery{Page: page, PageSize: pageSize, Filter: filter})
}

// AddNcrAction adds a disposition action to an NCR.
func (s *Service) AddNcrAction(ctx context.Context, action *model.NcrAction) (*model.NcrAction, error) {
	if action.NcrID == 0 {
		return nil, fmt.Errorf("%w: ncr_id is required", ErrInvalidInput)
	}
	if action.ActionType == "" {
		return nil, fmt.Errorf("%w: action_type is required", ErrInvalidInput)
	}

	ncr, err := s.repo.GetNcr(ctx, action.NcrID)
	if err != nil {
		return nil, fmt.Errorf("get ncr: %w", err)
	}
	if ncr == nil {
		return nil, fmt.Errorf("%w: ncr %d not found", ErrNotFound, action.NcrID)
	}

	if err := s.repo.CreateNcrAction(ctx, action); err != nil {
		return nil, fmt.Errorf("create ncr action: %w", err)
	}

	s.log.Info("ncr action added",
		zap.Uint("action_id", action.ID),
		zap.Uint("ncr_id", action.NcrID),
		zap.String("action_type", action.ActionType),
	)
	return action, nil
}

// ListNcrActions lists actions for an NCR.
func (s *Service) ListNcrActions(ctx context.Context, ncrID uint) ([]model.NcrAction, error) {
	return s.repo.ListNcrActions(ctx, ncrID)
}

// ==============================================================================
// InspectionCharacteristic
// ==============================================================================

func (s *Service) CreateInspectionCharacteristic(ctx context.Context, c *model.InspectionCharacteristic) error {
	if c.CharCode == "" || c.CharName == "" {
		return fmt.Errorf("%w: char_code and char_name are required", ErrInvalidInput)
	}
	if c.DataType == "" {
		c.DataType = "NUMERIC"
	}
	return s.repo.CreateInspectionCharacteristic(ctx, c)
}

func (s *Service) GetInspectionCharacteristic(ctx context.Context, id uint) (*model.InspectionCharacteristic, error) {
	c, err := s.repo.GetInspectionCharacteristic(ctx, id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, ErrNotFound
	}
	return c, nil
}

func (s *Service) ListInspectionCharacteristics(ctx context.Context, page, pageSize int, filter map[string]interface{}) (*repository.PageResult[model.InspectionCharacteristic], error) {
	return s.repo.ListInspectionCharacteristics(ctx, repository.PageQuery{Page: page, PageSize: pageSize, Filter: filter})
}

// ==============================================================================
// InspectionPlan
// ==============================================================================

func (s *Service) CreateInspectionPlan(ctx context.Context, p *model.InspectionPlan) error {
	if p.SchemeCode == "" || p.SchemeName == "" {
		return fmt.Errorf("%w: scheme_code and scheme_name are required", ErrInvalidInput)
	}
	if p.Status == "" {
		p.Status = "ACTIVE"
	}
	return s.repo.CreateInspectionPlan(ctx, p)
}

func (s *Service) GetInspectionPlan(ctx context.Context, id uint) (*model.InspectionPlan, error) {
	p, err := s.repo.GetInspectionPlan(ctx, id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, ErrNotFound
	}
	return p, nil
}

func (s *Service) ListInspectionPlans(ctx context.Context, page, pageSize int, filter map[string]interface{}) (*repository.PageResult[model.InspectionPlan], error) {
	return s.repo.ListInspectionPlans(ctx, repository.PageQuery{Page: page, PageSize: pageSize, Filter: filter})
}

// ==============================================================================
// DefectCode
// ==============================================================================

func (s *Service) CreateDefectCode(ctx context.Context, d *model.DefectCode) error {
	if d.DefectCode == "" || d.DefectName == "" {
		return fmt.Errorf("%w: defect_code and defect_name are required", ErrInvalidInput)
	}
	return s.repo.CreateDefectCode(ctx, d)
}

func (s *Service) GetDefectCode(ctx context.Context, id uint) (*model.DefectCode, error) {
	d, err := s.repo.GetDefectCode(ctx, id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, ErrNotFound
	}
	return d, nil
}

func (s *Service) ListDefectCodes(ctx context.Context, page, pageSize int, filter map[string]interface{}) (*repository.PageResult[model.DefectCode], error) {
	return s.repo.ListDefectCodes(ctx, repository.PageQuery{Page: page, PageSize: pageSize, Filter: filter})
}

// ==============================================================================
// SpcData
// ==============================================================================

func (s *Service) RecordSpcData(ctx context.Context, data *model.SpcData) error {
	if data.CharID == 0 {
		return fmt.Errorf("%w: char_id is required", ErrInvalidInput)
	}
	if data.SampleTime.IsZero() {
		data.SampleTime = time.Now()
	}
	return s.repo.CreateSpcData(ctx, data)
}

// ComputeSpcStats computes X-bar (mean) and R (range) for a characteristic's SPC data.
func (s *Service) ComputeSpcStats(ctx context.Context, charID uint) (xbar, r float64, err error) {
	data, err := s.repo.ListSpcDataByChar(ctx, charID)
	if err != nil {
		return 0, 0, fmt.Errorf("list spc data: %w", err)
	}
	if len(data) == 0 {
		return 0, 0, nil
	}

	min := math.MaxFloat64
	max := -math.MaxFloat64
	sum := 0.0
	for _, d := range data {
		sum += d.SampleValue
		if d.SampleValue < min {
			min = d.SampleValue
		}
		if d.SampleValue > max {
			max = d.SampleValue
		}
	}
	xbar = sum / float64(len(data))
	r = max - min
	return xbar, r, nil
}

// ==============================================================================
// Helpers
// ==============================================================================

// generateSheetNo produces a unique sheet number: IS-{TYPE}-{YYYYMMDD}-{timestamp_ns}.
func (s *Service) generateSheetNo(inspType string) string {
	typePrefix := inspType
	if typePrefix == "" {
		typePrefix = "GEN"
	}
	return fmt.Sprintf("IS-%s-%s-%d", typePrefix, time.Now().Format("20060102"), time.Now().UnixNano())
}

// generateNcrNo produces a unique NCR number: NCR-{YYYYMMDD}-{timestamp_ns}.
func (s *Service) generateNcrNo() string {
	return fmt.Sprintf("NCR-%s-%d", time.Now().Format("20060102"), time.Now().UnixNano())
}
