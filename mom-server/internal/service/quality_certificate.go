package service

import (
	"context"
	"errors"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type QualityCertificateService struct {
	repo *repository.QualityCertificateRepository
}

func NewQualityCertificateService(repo *repository.QualityCertificateRepository) *QualityCertificateService {
	return &QualityCertificateService{repo: repo}
}

func (s *QualityCertificateService) List(ctx context.Context, tenantID int64, query *model.QualityCertificateQuery) ([]model.QualityCertificate, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *QualityCertificateService) GetByID(ctx context.Context, id int64) (*model.QualityCertificate, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *QualityCertificateService) Create(ctx context.Context, tenantID int64, req *model.QualityCertificateCreateReq) (*model.QualityCertificate, error) {
	certCode, err := s.repo.GenerateCertCode(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	cert := &model.QualityCertificate{
		CertCode:    certCode,
		CertType:    req.CertType,
		OrderID:    req.OrderID,
		OrderCode:  req.OrderCode,
		ProductID:  req.ProductID,
		ProductCode: req.ProductCode,
		ProductName: req.ProductName,
		BatchNo:    req.BatchNo,
		Quantity:   req.Quantity,
		Unit:       req.Unit,
		Inspector:  req.Inspector,
		Result:     req.Result,
		Status:     1,
		Remarks:    req.Remarks,
		Attachments: req.Attachments,
	}

	if req.InspectDate != "" {
		if t, err := time.Parse("2006-01-02", req.InspectDate); err == nil {
			cert.InspectDate = t
		}
	}
	if req.IssueDate != "" {
		if t, err := time.Parse("2006-01-02", req.IssueDate); err == nil {
			cert.IssueDate = t
		}
	}
	if req.ExpiryDate != "" {
		if t, err := time.Parse("2006-01-02", req.ExpiryDate); err == nil {
			cert.ExpiryDate = t
		}
	}

	if err := s.repo.Create(ctx, tenantID, cert); err != nil {
		return nil, err
	}
	return cert, nil
}

func (s *QualityCertificateService) Update(ctx context.Context, id int64, req *model.QualityCertificateCreateReq) error {
	cert, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if cert == nil {
		return errors.New("certificate not found")
	}

	cert.CertType = req.CertType
	cert.OrderID = req.OrderID
	cert.OrderCode = req.OrderCode
	cert.ProductID = req.ProductID
	cert.ProductCode = req.ProductCode
	cert.ProductName = req.ProductName
	cert.BatchNo = req.BatchNo
	cert.Quantity = req.Quantity
	cert.Unit = req.Unit
	cert.Inspector = req.Inspector
	cert.Result = req.Result
	cert.Remarks = req.Remarks
	cert.Attachments = req.Attachments

	if req.InspectDate != "" {
		if t, err := time.Parse("2006-01-02", req.InspectDate); err == nil {
			cert.InspectDate = t
		}
	}
	if req.IssueDate != "" {
		if t, err := time.Parse("2006-01-02", req.IssueDate); err == nil {
			cert.IssueDate = t
		}
	}
	if req.ExpiryDate != "" {
		if t, err := time.Parse("2006-01-02", req.ExpiryDate); err == nil {
			cert.ExpiryDate = t
		}
	}

	return s.repo.Update(ctx, id, cert)
}

func (s *QualityCertificateService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}