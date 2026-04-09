package service

import (
	"context"
	"errors"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"strconv"
)

type WorkshopService struct {
	repo *repository.WorkshopRepository
}

func NewWorkshopService(repo *repository.WorkshopRepository) *WorkshopService {
	return &WorkshopService{repo: repo}
}

func (s *WorkshopService) List(ctx context.Context) ([]model.Workshop, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *WorkshopService) GetByID(ctx context.Context, id string) (*model.Workshop, error) {
	workshopID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("invalid workshop id")
	}
	return s.repo.GetByID(ctx, uint(workshopID))
}

func (s *WorkshopService) Create(ctx context.Context, workshop *model.Workshop) error {
	return s.repo.Create(ctx, workshop)
}

func (s *WorkshopService) Update(ctx context.Context, id string, workshop *model.Workshop) error {
	workshopID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid workshop id")
	}
	return s.repo.Update(ctx, uint(workshopID), map[string]interface{}{
		"workshop_name": workshop.WorkshopName,
		"workshop_code": workshop.WorkshopCode,
		"workshop_type": workshop.WorkshopType,
		"manager":       workshop.Manager,
		"phone":         workshop.Phone,
		"address":       workshop.Address,
		"status":        workshop.Status,
	})
}

func (s *WorkshopService) Delete(ctx context.Context, id string) error {
	workshopID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid workshop id")
	}
	return s.repo.Delete(ctx, uint(workshopID))
}
