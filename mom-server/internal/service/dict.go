package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type DictService struct {
	dictTypeRepo *repository.DictTypeRepository
	dictDataRepo *repository.DictDataRepository
}

func NewDictService(dtr *repository.DictTypeRepository, ddr *repository.DictDataRepository) *DictService {
	return &DictService{
		dictTypeRepo: dtr,
		dictDataRepo: ddr,
	}
}

func (s *DictService) ListType(ctx context.Context) ([]model.DictType, int64, error) {
	return s.dictTypeRepo.List(ctx, 0)
}

func (s *DictService) GetTypeByID(ctx context.Context, id string) (*model.DictType, error) {
	var dictID uint
	_, err := fmt.Sscanf(id, "%d", &dictID)
	if err != nil {
		return nil, err
	}
	return s.dictTypeRepo.GetByID(ctx, dictID)
}

func (s *DictService) CreateType(ctx context.Context, dict *model.DictType) error {
	return s.dictTypeRepo.Create(ctx, dict)
}

func (s *DictService) UpdateType(ctx context.Context, id string, dict *model.DictType) error {
	var dictID uint
	_, err := fmt.Sscanf(id, "%d", &dictID)
	if err != nil {
		return err
	}
	return s.dictTypeRepo.Update(ctx, dictID, map[string]interface{}{
		"dict_name": dict.DictName,
		"dict_type": dict.DictType,
		"status":    dict.Status,
		"remark":    dict.Remark,
	})
}

func (s *DictService) DeleteType(ctx context.Context, id string) error {
	var dictID uint
	_, err := fmt.Sscanf(id, "%d", &dictID)
	if err != nil {
		return err
	}
	return s.dictTypeRepo.Delete(ctx, dictID)
}

func (s *DictService) GetDataByType(ctx context.Context, dictType string) ([]model.DictData, error) {
	return s.dictDataRepo.GetByType(ctx, dictType)
}
