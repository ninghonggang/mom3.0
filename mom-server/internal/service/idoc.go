package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type IdocService struct {
	repo     *repository.IdocRepository
	sender   *IdocSender
}

func NewIdocService(repo *repository.IdocRepository) *IdocService {
	svc := &IdocService{repo: repo}
	svc.sender = NewIdocSender(svc)
	return svc
}

func (s *IdocService) List(ctx context.Context, tenantID int64, query *model.IdocQuery) ([]model.IdocRecord, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *IdocService) GetByID(ctx context.Context, id int64) (*model.IdocRecord, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *IdocService) Receive(ctx context.Context, tenantID int64, req *model.IdocReceiveReq) (*model.IdocRecord, error) {
	// Generate IDOC number
	idocNumber := s.repo.GenerateIdocNumber(req.IdocType)

	record := &model.IdocRecord{
		TenantID:    tenantID,
		IdocNumber: idocNumber,
		IdocType:   req.IdocType,
		Direction:  1, // 接收
		Status:     1,  // 新建
		PartnerType: req.PartnerType,
		PartnerNo:   req.PartnerNo,
		MessageType: req.MessageType,
		ReferenceNo: req.ReferenceNo,
		RawContent:  req.RawContent,
	}

	// Parse the IDOC content
	parsedData, parseErr := s.parseIdocContent(req.IdocType, req.RawContent)
	if parseErr != nil {
		record.Status = 4 // 失败
		record.ErrorMessage = parseErr.Error()
	} else {
		record.ParsedData = parsedData
	}

	if err := s.repo.Create(ctx, record); err != nil {
		return nil, err
	}

	// Process asynchronously
	go s.processReceive(record.ID)

	return record, nil
}

func (s *IdocService) processReceive(id int64) {
	ctx := context.Background()
	record, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return
	}

	// Update status to processing
	s.repo.UpdateStatus(ctx, id, 2, "")

	// Get config for this IDOC type
	config, err := s.repo.GetConfigByType(ctx, record.IdocType)
	if err != nil {
		s.repo.UpdateStatus(ctx, id, 4, "IDOC类型未配置: "+record.IdocType)
		return
	}

	// Apply mapping rules and process
	err = s.applyMappingAndProcess(record, config)
	if err != nil {
		s.repo.UpdateStatus(ctx, id, 4, err.Error())
		return
	}

	// Update status to success
	now := time.Now()
	record.Status = 3
	record.ProcessedAt = &now
	s.repo.Update(ctx, record)
}

func (s *IdocService) parseIdocContent(idocType, rawContent string) (string, error) {
	// Basic JSON validation for parsed data
	var data interface{}
	if err := json.Unmarshal([]byte(rawContent), &data); err != nil {
		// If not valid JSON, wrap it
		return fmt.Sprintf(`{"raw": %s}`, rawContent), nil
	}
	return rawContent, nil
}

func (s *IdocService) applyMappingAndProcess(record *model.IdocRecord, config *model.IdocTypeConfig) error {
	// This is a placeholder for actual IDOC processing logic
	// In a real implementation, you would:
	// 1. Parse the IDOC segments
	// 2. Apply mapping rules to convert to internal format
	// 3. Call appropriate business logic
	// 4. Create/update records in the system

	switch record.IdocType {
	case "MATMAS": // Material Master
		// Process material master data
	case "ORDERS": // Sales Order
		// Process sales order
	case "DESADV": // Delivery
		// Process delivery
	default:
		// Unknown type, just acknowledge
	}
	return nil
}

func (s *IdocService) Send(ctx context.Context, tenantID int64, req *model.IdocSendReq) (*model.IdocRecord, error) {
	// Get config
	config, err := s.repo.GetConfigByType(ctx, req.IdocType)
	if err != nil {
		return nil, errors.New("IDOC类型未配置: " + req.IdocType)
	}

	// Generate IDOC number
	idocNumber := s.repo.GenerateIdocNumber(req.IdocType)

	// Marshal data to JSON
	dataBytes, _ := json.Marshal(req.Data)
	rawContent := string(dataBytes)

	record := &model.IdocRecord{
		TenantID:    tenantID,
		IdocNumber: idocNumber,
		IdocType:   req.IdocType,
		Direction:  2, // 发送
		Status:     1, // 新建
		PartnerType: req.TargetType,
		PartnerNo:   req.TargetNo,
		MessageType: req.MessageType,
		ReferenceNo: req.IdocType + "-" + time.Now().Format("20060102"),
		RawContent:  rawContent,
	}

	if err := s.repo.Create(ctx, record); err != nil {
		return nil, err
	}

	// Send asynchronously
	go s.processSend(record.ID, config)

	return record, nil
}

func (s *IdocService) processSend(id int64, config *model.IdocTypeConfig) {
	ctx := context.Background()
	record, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return
	}

	// Update status to processing
	s.repo.UpdateStatus(ctx, id, 2, "")

	// Send the IDOC
	err = s.sender.Send(record, config)
	if err != nil {
		s.repo.UpdateStatus(ctx, id, 4, err.Error())
		return
	}

	// Update status to success
	now := time.Now()
	record.Status = 3
	record.ProcessedAt = &now
	s.repo.Update(ctx, record)
}

func (s *IdocService) Retry(ctx context.Context, id int64) error {
	record, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if record.Status != 4 {
		return errors.New("只能重试失败的IDOC")
	}

	// Increment retry count
	record.RetryCount++

	if record.Direction == 1 {
		go s.processReceive(record.ID)
	} else {
		config, _ := s.repo.GetConfigByType(ctx, record.IdocType)
		go s.processSend(record.ID, config)
	}

	return nil
}

func (s *IdocService) ListConfigs(ctx context.Context, tenantID int64) ([]model.IdocTypeConfig, error) {
	return s.repo.ListConfigs(ctx, tenantID)
}

type IdocSender struct {
	svc *IdocService
}

func NewIdocSender(svc *IdocService) *IdocSender {
	return &IdocSender{svc: svc}
}

func (s *IdocSender) Send(record *model.IdocRecord, config *model.IdocTypeConfig) error {
	// This is a placeholder for actual HTTP/FILE/RabbitMQ sending
	// In a real implementation, you would send to the configured endpoint
	return nil
}