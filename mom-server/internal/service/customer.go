package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) List(ctx context.Context) ([]model.Customer, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *CustomerService) GetByID(ctx context.Context, id string) (*model.Customer, error) {
	var customerID uint
	_, err := fmt.Sscanf(id, "%d", &customerID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, customerID)
}

func (s *CustomerService) Create(ctx context.Context, customer *model.Customer) error {
	return s.repo.Create(ctx, customer)
}

func (s *CustomerService) Update(ctx context.Context, id string, customer *model.Customer) error {
	var customerID uint
	_, err := fmt.Sscanf(id, "%d", &customerID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, customerID, map[string]interface{}{
		"code":    customer.Code,
		"name":    customer.Name,
		"type":    customer.Type,
		"contact": customer.Contact,
		"phone":   customer.Phone,
		"email":   customer.Email,
		"address": customer.Address,
		"status":  customer.Status,
	})
}

func (s *CustomerService) Delete(ctx context.Context, id string) error {
	var customerID uint
	_, err := fmt.Sscanf(id, "%d", &customerID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, customerID)
}
