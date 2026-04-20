package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type EquipmentDocumentService struct {
	repo *repository.EquipmentDocumentRepository
}

func NewEquipmentDocumentService(repo *repository.EquipmentDocumentRepository) *EquipmentDocumentService {
	return &EquipmentDocumentService{repo: repo}
}

func (s *EquipmentDocumentService) List(ctx context.Context, tenantID int64, query map[string]any) ([]model.EquipmentDocument, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *EquipmentDocumentService) GetByID(ctx context.Context, id uint) (*model.EquipmentDocument, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EquipmentDocumentService) Create(ctx context.Context, tenantID int64, req *model.EquipmentDocumentCreate, username string) (*model.EquipmentDocument, error) {
	doc := &model.EquipmentDocument{
		TenantID:       tenantID,
		EquipmentID:   req.EquipmentID,
		DocType:       req.DocType,
		DocName:       req.DocName,
		DocCode:       req.DocCode,
		FileName:      req.FileName,
		FilePath:      req.FilePath,
		FileSize:      req.FileSize,
		FileType:      req.FileType,
		FileURL:       req.FileURL,
		Version:       req.Version,
		EffectiveDate: req.EffectiveDate,
		ExpiryDate:    req.ExpiryDate,
		Description:   req.Description,
		UploadTime:    func() *time.Time { t := time.Now(); return &t }(),
		Status:        1,
	}
	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *EquipmentDocumentService) Update(ctx context.Context, id uint, req *model.EquipmentDocumentUpdate) error {
	updates := make(map[string]any)
	if req.DocType != nil {
		updates["doc_type"] = *req.DocType
	}
	if req.DocName != nil {
		updates["doc_name"] = *req.DocName
	}
	if req.DocCode != nil {
		updates["doc_code"] = *req.DocCode
	}
	if req.FileName != nil {
		updates["file_name"] = *req.FileName
	}
	if req.FilePath != nil {
		updates["file_path"] = *req.FilePath
	}
	if req.FileSize != nil {
		updates["file_size"] = *req.FileSize
	}
	if req.FileType != nil {
		updates["file_type"] = *req.FileType
	}
	if req.FileURL != nil {
		updates["file_url"] = *req.FileURL
	}
	if req.Version != nil {
		updates["version"] = *req.Version
	}
	if req.EffectiveDate != nil {
		updates["effective_date"] = *req.EffectiveDate
	}
	if req.ExpiryDate != nil {
		updates["expiry_date"] = *req.ExpiryDate
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	return s.repo.Update(ctx, id, updates)
}

func (s *EquipmentDocumentService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *EquipmentDocumentService) ListByEquipmentID(ctx context.Context, equipmentID int64) ([]model.EquipmentDocument, error) {
	return s.repo.ListByEquipmentID(ctx, equipmentID)
}
