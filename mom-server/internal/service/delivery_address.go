package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// DeliveryAddressService 收货地址服务
type DeliveryAddressService struct {
	repo *repository.DeliveryAddressRepository
}

func NewDeliveryAddressService(repo *repository.DeliveryAddressRepository) *DeliveryAddressService {
	return &DeliveryAddressService{repo: repo}
}

func (s *DeliveryAddressService) List(ctx context.Context, query model.DeliveryAddressQuery) ([]model.DeliveryAddress, int64, error) {
	return s.repo.List(ctx, 0, query)
}

func (s *DeliveryAddressService) GetByID(ctx context.Context, id uint64) (*model.DeliveryAddress, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DeliveryAddressService) Create(ctx context.Context, req *model.DeliveryAddressCreateRequest) error {
	address := &model.DeliveryAddress{
		CustomerID:    req.CustomerID,
		AddressName:   req.AddressName,
		ContactPerson: req.ContactPerson,
		ContactPhone:  req.ContactPhone,
		Province:      req.Province,
		City:          req.City,
		District:      req.District,
		AddressDetail: req.AddressDetail,
		IsDefault:     req.IsDefault,
		IsActive:      req.IsActive,
	}
	if req.IsDefault {
		// 取消其他默认地址
		s.repo.SetDefault(ctx, req.CustomerID, 0)
	}
	return s.repo.Create(ctx, address)
}

func (s *DeliveryAddressService) Update(ctx context.Context, id uint64, req *model.DeliveryAddressUpdateRequest) error {
	updates := map[string]interface{}{}
	if req.AddressName != "" {
		updates["address_name"] = req.AddressName
	}
	if req.ContactPerson != "" {
		updates["contact_person"] = req.ContactPerson
	}
	if req.ContactPhone != "" {
		updates["contact_phone"] = req.ContactPhone
	}
	if req.Province != "" {
		updates["province"] = req.Province
	}
	if req.City != "" {
		updates["city"] = req.City
	}
	if req.District != "" {
		updates["district"] = req.District
	}
	if req.AddressDetail != "" {
		updates["address_detail"] = req.AddressDetail
	}
	if req.IsDefault {
		// 获取地址信息以获得customer_id
		addr, _ := s.repo.GetByID(ctx, id)
		if addr != nil {
			s.repo.SetDefault(ctx, addr.CustomerID, id)
		}
	}

	return s.repo.Update(ctx, id, updates)
}

func (s *DeliveryAddressService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *DeliveryAddressService) GetByCustomer(ctx context.Context, customerID uint64) ([]model.DeliveryAddress, error) {
	return s.repo.GetByCustomer(ctx, customerID)
}

func (s *DeliveryAddressService) SetDefault(ctx context.Context, customerID uint64, addressID uint64) error {
	return s.repo.SetDefault(ctx, customerID, addressID)
}
