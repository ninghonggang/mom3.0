package service

import (
	"context"
	"errors"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type CustomerCreditService struct {
	repo *repository.CustomerCreditRepository
}

func NewCustomerCreditService(repo *repository.CustomerCreditRepository) *CustomerCreditService {
	return &CustomerCreditService{repo: repo}
}

func (s *CustomerCreditService) List(ctx context.Context, tenantID int64, query *model.CustomerCreditQuery) ([]model.CustomerCredit, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *CustomerCreditService) GetByID(ctx context.Context, id int64) (*model.CustomerCredit, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CustomerCreditService) GetByCustomerID(ctx context.Context, customerID int64) (*model.CustomerCredit, error) {
	return s.repo.GetByCustomerID(ctx, customerID)
}

func (s *CustomerCreditService) Create(ctx context.Context, tenantID int64, req *model.CustomerCreditCreateReq) (*model.CustomerCredit, error) {
	// Check if already exists
	existing, err := s.repo.GetByCustomerID(ctx, req.CustomerID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("customer credit record already exists")
	}

	credit := &model.CustomerCredit{
		CustomerID:     req.CustomerID,
		CustomerCode:   req.CustomerCode,
		CustomerName:   req.CustomerName,
		CreditLimit:    req.CreditLimit,
		UsedCredit:     0,
		AvailableCredit: req.CreditLimit,
		CreditLevel:    req.CreditLevel,
		PaymentDays:    req.PaymentDays,
		RiskLevel:      req.RiskLevel,
		AlertThreshold: req.AlertThreshold,
		Status:         1,
		Remarks:        req.Remarks,
	}

	if err := s.repo.Create(ctx, tenantID, credit); err != nil {
		return nil, err
	}
	return credit, nil
}

func (s *CustomerCreditService) Update(ctx context.Context, id int64, req *model.CustomerCreditCreateReq) error {
	credit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if credit == nil {
		return errors.New("customer credit not found")
	}

	credit.CreditLimit = req.CreditLimit
	credit.CreditLevel = req.CreditLevel
	credit.PaymentDays = req.PaymentDays
	credit.RiskLevel = req.RiskLevel
	credit.AlertThreshold = req.AlertThreshold
	credit.Remarks = req.Remarks
	// Recalculate available credit
	credit.AvailableCredit = credit.CreditLimit - credit.UsedCredit

	return s.repo.Update(ctx, id, credit)
}

func (s *CustomerCreditService) UpdateUsedCredit(ctx context.Context, id int64, usedCredit float64) error {
	credit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if credit == nil {
		return errors.New("customer credit not found")
	}

	credit.UsedCredit = usedCredit
	credit.AvailableCredit = credit.CreditLimit - usedCredit

	return s.repo.UpdateUsedCredit(ctx, id, usedCredit, credit.AvailableCredit)
}

func (s *CustomerCreditService) SetBlacklist(ctx context.Context, id int64, blacklist bool) error {
	credit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if credit == nil {
		return errors.New("customer credit not found")
	}

	if blacklist {
		credit.Blacklist = 1
	} else {
		credit.Blacklist = 0
	}

	return s.repo.Update(ctx, id, credit)
}

func (s *CustomerCreditService) Freeze(ctx context.Context, id int64) error {
	credit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if credit == nil {
		return errors.New("customer credit not found")
	}

	credit.Status = 0 // 冻结
	return s.repo.Update(ctx, id, credit)
}

func (s *CustomerCreditService) Unfreeze(ctx context.Context, id int64) error {
	credit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if credit == nil {
		return errors.New("customer credit not found")
	}

	credit.Status = 1 // 解冻
	return s.repo.Update(ctx, id, credit)
}

func (s *CustomerCreditService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}