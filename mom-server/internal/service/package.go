package service

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type PackageService struct {
	repo *repository.PackageRepository
}

func NewPackageService(repo *repository.PackageRepository) *PackageService {
	return &PackageService{repo: repo}
}

func (s *PackageService) List(ctx context.Context, tenantID int64, req *repository.PackageQuery) ([]model.MesPackage, int64, error) {
	return s.repo.List(ctx, tenantID, req)
}

func (s *PackageService) GetByID(ctx context.Context, id uint) (*model.MesPackage, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PackageService) GetByPackageNo(ctx context.Context, packageNo string) (*model.MesPackage, error) {
	return s.repo.GetByPackageNo(ctx, packageNo)
}

// CreatePackage 创建箱
func (s *PackageService) CreatePackage(ctx context.Context, pkg *model.MesPackage) error {
	if pkg.TenantID == 0 {
		pkg.TenantID = 1
	}
	// 自动生成箱条码
	pkg.PackageNo = fmt.Sprintf("PKG-%d-%d", time.Now().Unix(), pkg.ProductID%10000)
	pkg.Status = "OPEN"
	return s.repo.Create(ctx, pkg)
}

// AddSerialNo 添加序列号到箱
func (s *PackageService) AddSerialNo(ctx context.Context, pkgID uint, serialNo string) error {
	pkg, err := s.repo.GetByID(ctx, pkgID)
	if err != nil {
		return err
	}
	if pkg.Status != "OPEN" {
		return fmt.Errorf("箱已封箱，无法添加")
	}
	// 检查是否重复
	for _, s := range pkg.SerialNos {
		if s == serialNo {
			return fmt.Errorf("序列号已存在")
		}
	}
	pkg.SerialNos = append(pkg.SerialNos, serialNo)
	pkg.Qty = len(pkg.SerialNos)
	return s.repo.UpdateSerialNos(ctx, pkgID, pkg.SerialNos)
}

// SealPackage 封箱
func (s *PackageService) SealPackage(ctx context.Context, pkgID uint, sealBy string) error {
	pkg, err := s.repo.GetByID(ctx, pkgID)
	if err != nil {
		return err
	}
	if pkg.Status != "OPEN" {
		return fmt.Errorf("箱状态不是OPEN")
	}
	now := time.Now()
	return s.repo.Update(ctx, pkgID, map[string]interface{}{
		"status":    "SEALED",
		"seal_time": now,
		"seal_by":   sealBy,
	})
}

// Delete 删除
func (s *PackageService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
