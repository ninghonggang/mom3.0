package handler

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ninghonggang/mom-platform/gen/common"
	pb "github.com/ninghonggang/mom-platform/gen/mdm"

	"mom-platform/services/mdm-service/internal/model"
	"mom-platform/services/mdm-service/internal/repository"
	"mom-platform/services/mdm-service/internal/service"
)

// =========================================================================
// Proto <-> Internal Model converters
// =========================================================================

func protoToMaterial(pb *pb.Material) *model.Material {
	m := &model.Material{}
	if pb.Base != nil {
		m.ID = uint64(pb.Base.Id)
		m.TenantID = strconv.FormatInt(pb.Base.TenantId, 10)
	}
	m.MaterialCode = pb.MaterialCode
	m.MaterialName = pb.MaterialName
	m.Spec = pb.Spec
	m.Unit = pb.UnitName
	m.Status = pb.Status.String()
	return m
}

func materialToProto(m *model.Material) *pb.Material {
	tenantID, _ := strconv.ParseInt(m.TenantID, 10, 64)
	return &pb.Material{
		Base: &common.BaseModel{
			Id:       int64(m.ID),
			TenantId: tenantID,
		},
		MaterialCode: m.MaterialCode,
		MaterialName: m.MaterialName,
		Spec:         m.Spec,
		UnitName:     m.Unit,
	}
}

func protoToBom(pb *pb.BOMHeader) (*model.Bom, []model.BomItem) {
	bom := &model.Bom{}
	if pb.Base != nil {
		bom.ID = uint64(pb.Base.Id)
		bom.TenantID = strconv.FormatInt(pb.Base.TenantId, 10)
	}
	bom.BomCode = pb.BomCode
	bom.MaterialID = uint64(pb.ProductId)
	bom.Version = pb.Version
	if pb.EffectiveDate != nil {
		bom.EffectiveDate = pb.EffectiveDate.AsTime()
	}
	bom.Status = pb.Status.String()

	var items []model.BomItem
	for _, line := range pb.Lines {
		if line == nil {
			continue
		}
		item := model.BomItem{
			ChildMaterialID: uint64(line.ComponentId),
			Quantity:        line.Quantity,
			ScrapRate:       line.LossRate,
		}
		items = append(items, item)
	}
	return bom, items
}

func bomToProto(b *model.Bom) *pb.BOMHeader {
	tenantID, _ := strconv.ParseInt(b.TenantID, 10, 64)
	h := &pb.BOMHeader{
		Base: &common.BaseModel{
			Id:       int64(b.ID),
			TenantId: tenantID,
		},
		BomCode:   b.BomCode,
		ProductId: int64(b.MaterialID),
		Version:   b.Version,
	}
	for i, item := range b.Items {
		h.Lines = append(h.Lines, &pb.BOMLine{
			ComponentId: int64(item.ChildMaterialID),
			Quantity:    item.Quantity,
			LossRate:    item.ScrapRate,
			Level:       int32(i + 1),
		})
	}
	return h
}

func protoToWorkshop(pb *pb.Workshop) *model.Workshop {
	w := &model.Workshop{}
	if pb.Base != nil {
		w.ID = uint64(pb.Base.Id)
		w.TenantID = strconv.FormatInt(pb.Base.TenantId, 10)
	}
	w.WorkshopCode = pb.WorkshopCode
	w.WorkshopName = pb.WorkshopName
	w.FactoryID = uint64(pb.FactoryId)
	return w
}

func workshopToProto(w *model.Workshop) *pb.Workshop {
	tenantID, _ := strconv.ParseInt(w.TenantID, 10, 64)
	return &pb.Workshop{
		Base: &common.BaseModel{
			Id:       int64(w.ID),
			TenantId: tenantID,
		},
		WorkshopCode: w.WorkshopCode,
		WorkshopName: w.WorkshopName,
		FactoryId:    int64(w.FactoryID),
	}
}

func protoToProdLine(pb *pb.ProductionLine) *model.ProductionLine {
	pl := &model.ProductionLine{}
	if pb.Base != nil {
		pl.ID = uint64(pb.Base.Id)
		pl.TenantID = strconv.FormatInt(pb.Base.TenantId, 10)
	}
	pl.LineCode = pb.LineCode
	pl.LineName = pb.LineName
	pl.WorkshopID = uint64(pb.WorkshopId)
	return pl
}

func productionLineToProto(pl *model.ProductionLine) *pb.ProductionLine {
	tenantID, _ := strconv.ParseInt(pl.TenantID, 10, 64)
	return &pb.ProductionLine{
		Base: &common.BaseModel{
			Id:       int64(pl.ID),
			TenantId: tenantID,
		},
		LineCode:   pl.LineCode,
		LineName:   pl.LineName,
		WorkshopId: int64(pl.WorkshopID),
	}
}

func protoToWorkstation(pb *pb.Workstation) *model.Workstation {
	ws := &model.Workstation{}
	if pb.Base != nil {
		ws.ID = uint64(pb.Base.Id)
		ws.TenantID = strconv.FormatInt(pb.Base.TenantId, 10)
	}
	ws.WorkstationCode = pb.StationCode
	ws.WorkstationName = pb.StationName
	ws.LineID = uint64(pb.LineId)
	return ws
}

func workstationToProto(ws *model.Workstation) *pb.Workstation {
	tenantID, _ := strconv.ParseInt(ws.TenantID, 10, 64)
	return &pb.Workstation{
		Base: &common.BaseModel{
			Id:       int64(ws.ID),
			TenantId: tenantID,
		},
		StationCode: ws.WorkstationCode,
		StationName: ws.WorkstationName,
		LineId:      int64(ws.LineID),
	}
}

func protoToCustomer(pb *pb.Customer) *model.Customer {
	c := &model.Customer{}
	if pb.Base != nil {
		c.ID = uint64(pb.Base.Id)
		c.TenantID = strconv.FormatInt(pb.Base.TenantId, 10)
	}
	c.CustomerCode = pb.CustomerCode
	c.CustomerName = pb.CustomerName
	c.ContactPerson = pb.ContactPerson
	c.ContactPhone = pb.ContactPhone
	c.Address = pb.Address
	return c
}

func customerToProto(c *model.Customer) *pb.Customer {
	tenantID, _ := strconv.ParseInt(c.TenantID, 10, 64)
	return &pb.Customer{
		Base: &common.BaseModel{
			Id:       int64(c.ID),
			TenantId: tenantID,
		},
		CustomerCode:  c.CustomerCode,
		CustomerName:  c.CustomerName,
		ContactPerson: c.ContactPerson,
		ContactPhone:  c.ContactPhone,
		Address:       c.Address,
	}
}

func protoToSupplier(pb *pb.Supplier) *model.Supplier {
	s := &model.Supplier{}
	if pb.Base != nil {
		s.ID = uint64(pb.Base.Id)
		s.TenantID = strconv.FormatInt(pb.Base.TenantId, 10)
	}
	s.SupplierCode = pb.SupplierCode
	s.SupplierName = pb.SupplierName
	s.ContactPerson = pb.ContactPerson
	s.ContactPhone = pb.ContactPhone
	s.Address = pb.Address
	return s
}

func supplierToProto(s *model.Supplier) *pb.Supplier {
	tenantID, _ := strconv.ParseInt(s.TenantID, 10, 64)
	return &pb.Supplier{
		Base: &common.BaseModel{
			Id:       int64(s.ID),
			TenantId: tenantID,
		},
		SupplierCode:  s.SupplierCode,
		SupplierName:  s.SupplierName,
		ContactPerson: s.ContactPerson,
		ContactPhone:  s.ContactPhone,
		Address:       s.Address,
	}
}

func paginationToProto(total int64, req *common.Pagination) *common.Pagination {
	if req == nil {
		req = &common.Pagination{Page: 1, PageSize: 20}
	}
	pb := &common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int32(total),
	}
	if pb.PageSize > 0 {
		pb.TotalPages = int32((total + int64(pb.PageSize) - 1) / int64(pb.PageSize))
	}
	return pb
}

func paginationToOffsetLimit(pb *common.Pagination) (int, int) {
	page := 1
	pageSize := 20
	if pb != nil {
		if pb.Page > 0 {
			page = int(pb.Page)
		}
		if pb.PageSize > 0 {
			pageSize = int(pb.PageSize)
		}
	}
	return (page - 1) * pageSize, pageSize
}

// =========================================================================
// MaterialService gRPC server implementation
// =========================================================================

type materialServer struct {
	pb.UnimplementedMaterialServiceServer
	logger  *zap.Logger
	service *service.MDMService
}

func (s *materialServer) GetMaterial(ctx context.Context, req *pb.GetMaterialRequest) (*pb.GetMaterialResponse, error) {
	m, err := s.service.GetMaterial(ctx, uint64(req.Id))
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, status.Error(codes.NotFound, "material not found")
		}
		s.logger.Error("GetMaterial failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetMaterialResponse{Material: materialToProto(m)}, nil
}

func (s *materialServer) ListMaterials(ctx context.Context, req *pb.ListMaterialsRequest) (*pb.ListMaterialsResponse, error) {
	tenantID := strconv.FormatInt(req.TenantId, 10)
	offset, limit := paginationToOffsetLimit(req.Pagination)
	list, total, err := s.service.ListMaterials(ctx, tenantID, offset, limit)
	if err != nil {
		s.logger.Error("ListMaterials failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	items := make([]*pb.Material, len(list))
	for i := range list {
		items[i] = materialToProto(&list[i])
	}
	return &pb.ListMaterialsResponse{
		Items:      items,
		Pagination: paginationToProto(total, req.Pagination),
	}, nil
}

func (s *materialServer) CreateMaterial(ctx context.Context, req *pb.CreateMaterialRequest) (*pb.CreateMaterialResponse, error) {
	m := protoToMaterial(req.Material)
	if err := s.service.CreateMaterial(ctx, m); err != nil {
		if err == service.ErrDuplicateCode {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		s.logger.Error("CreateMaterial failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateMaterialResponse{Material: materialToProto(m)}, nil
}

func (s *materialServer) UpdateMaterial(ctx context.Context, req *pb.UpdateMaterialRequest) (*pb.UpdateMaterialResponse, error) {
	m := protoToMaterial(req.Material)
	if err := s.service.UpdateMaterial(ctx, m); err != nil {
		if err == repository.ErrNotFound {
			return nil, status.Error(codes.NotFound, "material not found")
		}
		if err == service.ErrDuplicateCode {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		s.logger.Error("UpdateMaterial failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UpdateMaterialResponse{Material: materialToProto(m)}, nil
}

// =========================================================================
// BOMService gRPC server implementation
// =========================================================================

type bomServer struct {
	pb.UnimplementedBOMServiceServer
	logger  *zap.Logger
	service *service.MDMService
}

func (s *bomServer) GetBOM(ctx context.Context, req *pb.GetBOMRequest) (*pb.GetBOMResponse, error) {
	bom, err := s.service.GetBom(ctx, uint64(req.Id))
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, status.Error(codes.NotFound, "BOM not found")
		}
		s.logger.Error("GetBOM failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetBOMResponse{Bom: bomToProto(bom)}, nil
}

func (s *bomServer) ListBOMs(ctx context.Context, req *pb.ListBOMsRequest) (*pb.ListBOMsResponse, error) {
	tenantID := strconv.FormatInt(req.TenantId, 10)
	offset, limit := paginationToOffsetLimit(req.Pagination)
	list, total, err := s.service.ListBoms(ctx, tenantID, offset, limit)
	if err != nil {
		s.logger.Error("ListBOMs failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	items := make([]*pb.BOMHeader, len(list))
	for i := range list {
		items[i] = bomToProto(&list[i])
	}
	return &pb.ListBOMsResponse{
		Items:      items,
		Pagination: paginationToProto(total, req.Pagination),
	}, nil
}

func (s *bomServer) CreateBOM(ctx context.Context, req *pb.CreateBOMRequest) (*pb.CreateBOMResponse, error) {
	bom, items := protoToBom(req.Bom)
	if err := s.service.CreateBom(ctx, bom, items); err != nil {
		if err == service.ErrDuplicateCode {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		s.logger.Error("CreateBOM failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Re-fetch to get populated items
	bom, err := s.service.GetBom(ctx, bom.ID)
	if err != nil {
		s.logger.Error("GetBOM after create failed", zap.Error(err))
		return &pb.CreateBOMResponse{Bom: bomToProto(bom)}, nil
	}
	return &pb.CreateBOMResponse{Bom: bomToProto(bom)}, nil
}

// =========================================================================
// FactoryAssetService gRPC server implementation
// =========================================================================

type factoryAssetServer struct {
	pb.UnimplementedFactoryAssetServiceServer
	logger  *zap.Logger
	service *service.MDMService
}

func (s *factoryAssetServer) ListWorkshops(ctx context.Context, req *pb.ListWorkshopsRequest) (*pb.ListWorkshopsResponse, error) {
	tenantID := strconv.FormatInt(req.TenantId, 10)
	list, _, err := s.service.ListWorkshops(ctx, tenantID, 0, 1000)
	if err != nil {
		s.logger.Error("ListWorkshops failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	items := make([]*pb.Workshop, len(list))
	for i := range list {
		items[i] = workshopToProto(&list[i])
	}
	return &pb.ListWorkshopsResponse{Items: items}, nil
}

func (s *factoryAssetServer) ListProductionLines(ctx context.Context, req *pb.ListProductionLinesRequest) (*pb.ListProductionLinesResponse, error) {
	list, _, err := s.service.ListProductionLines(ctx, "", 0, 1000)
	if err != nil {
		s.logger.Error("ListProductionLines failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Filter by workshop if specified
	var filtered []model.ProductionLine
	for _, pl := range list {
		if req.WorkshopId == 0 || uint64(req.WorkshopId) == pl.WorkshopID {
			filtered = append(filtered, pl)
		}
	}
	items := make([]*pb.ProductionLine, len(filtered))
	for i := range filtered {
		items[i] = productionLineToProto(&filtered[i])
	}
	return &pb.ListProductionLinesResponse{Items: items}, nil
}

func (s *factoryAssetServer) ListWorkstations(ctx context.Context, req *pb.ListWorkstationsRequest) (*pb.ListWorkstationsResponse, error) {
	list, _, err := s.service.ListWorkstations(ctx, "", 0, 1000)
	if err != nil {
		s.logger.Error("ListWorkstations failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	var filtered []model.Workstation
	for _, ws := range list {
		if req.LineId == 0 || uint64(req.LineId) == ws.LineID {
			filtered = append(filtered, ws)
		}
	}
	items := make([]*pb.Workstation, len(filtered))
	for i := range filtered {
		items[i] = workstationToProto(&filtered[i])
	}
	return &pb.ListWorkstationsResponse{Items: items}, nil
}

// =========================================================================
// PartnerService gRPC server implementation
// =========================================================================

type partnerServer struct {
	pb.UnimplementedPartnerServiceServer
	logger  *zap.Logger
	service *service.MDMService
}

func (s *partnerServer) ListCustomers(ctx context.Context, req *pb.ListCustomersRequest) (*pb.ListCustomersResponse, error) {
	tenantID := strconv.FormatInt(req.TenantId, 10)
	offset, limit := paginationToOffsetLimit(req.Pagination)
	list, total, err := s.service.ListCustomers(ctx, tenantID, offset, limit)
	if err != nil {
		s.logger.Error("ListCustomers failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	items := make([]*pb.Customer, len(list))
	for i := range list {
		items[i] = customerToProto(&list[i])
	}
	return &pb.ListCustomersResponse{
		Items:      items,
		Pagination: paginationToProto(total, req.Pagination),
	}, nil
}

func (s *partnerServer) ListSuppliers(ctx context.Context, req *pb.ListSuppliersRequest) (*pb.ListSuppliersResponse, error) {
	tenantID := strconv.FormatInt(req.TenantId, 10)
	offset, limit := paginationToOffsetLimit(req.Pagination)
	list, total, err := s.service.ListSuppliers(ctx, tenantID, offset, limit)
	if err != nil {
		s.logger.Error("ListSuppliers failed", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	items := make([]*pb.Supplier, len(list))
	for i := range list {
		items[i] = supplierToProto(&list[i])
	}
	return &pb.ListSuppliersResponse{
		Items:      items,
		Pagination: paginationToProto(total, req.Pagination),
	}, nil
}

// =========================================================================
// Register all gRPC services
// =========================================================================

// RegisterGRPCServices registers all MDM gRPC service implementations on the
// provided gRPC server.
func RegisterGRPCServices(grpcSrv *grpc.Server, logger *zap.Logger, svc *service.MDMService) {
	pb.RegisterMaterialServiceServer(grpcSrv, &materialServer{logger: logger, service: svc})
	pb.RegisterBOMServiceServer(grpcSrv, &bomServer{logger: logger, service: svc})
	pb.RegisterFactoryAssetServiceServer(grpcSrv, &factoryAssetServer{logger: logger, service: svc})
	pb.RegisterPartnerServiceServer(grpcSrv, &partnerServer{logger: logger, service: svc})
}
