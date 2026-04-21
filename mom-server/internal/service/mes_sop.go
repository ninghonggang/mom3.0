package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"os"
	"path/filepath"
	"time"
)

// MesSopService MES SOP管理服务
type MesSopService struct {
	repo *repository.ElectronicSOPRepository
}

// NewMesSopService 创建SOP服务
func NewMesSopService(repo *repository.ElectronicSOPRepository) *MesSopService {
	return &MesSopService{repo: repo}
}

// UploadReq SOP上传请求
type UploadReq struct {
	File           *multipart.FileHeader
	SopName        string
	ProcessRouteId int64
	WorkOrderId    int64
	Version        string
	Uploader       string
}

// Upload 上传SOP文档
func (s *MesSopService) Upload(ctx context.Context, tenantID int64, req *UploadReq) (*model.ElectronicSOP, error) {
	// 创建上传目录
	uploadDir := filepath.Join("uploads", "sop")
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("创建上传目录失败: %w", err)
	}

	// 生成文件名: {tenantID}_{timestamp}_{originalname}
	timestamp := time.Now().UnixMilli()
	originalName := req.File.Filename
	newFileName := fmt.Sprintf("%d_%d_%s", tenantID, timestamp, originalName)
	filePath := filepath.Join(uploadDir, newFileName)

	// 保存文件
	src, err := req.File.Open()
	if err != nil {
		return nil, fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return nil, fmt.Errorf("保存文件失败: %w", err)
	}

	// 创建SOP记录
	sop := &model.ElectronicSOP{
		TenantID:      tenantID,
		SopName:       req.SopName,
		ContentType:   "PDF",
		ContentURL:    filePath,
		Version:       req.Version,
		ProcessID:     req.ProcessRouteId,
		WorkstationID: req.WorkOrderId,
		Status:        2, // 已发布
	}

	if err := s.repo.Create(ctx, sop); err != nil {
		// 删除已保存的文件
		os.Remove(filePath)
		return nil, fmt.Errorf("保存SOP记录失败: %w", err)
	}

	return sop, nil
}

// Download 获取SOP文件路径
func (s *MesSopService) Download(ctx context.Context, id uint) (string, string, error) {
	sop, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return "", "", fmt.Errorf("获取SOP记录失败: %w", err)
	}
	return sop.ContentURL, sop.SopName, nil
}

// GetByID 获取SOP详情
func (s *MesSopService) GetByID(ctx context.Context, id uint) (*model.ElectronicSOP, error) {
	return s.repo.GetByID(ctx, id)
}

// Delete 删除SOP文档
func (s *MesSopService) Delete(ctx context.Context, id uint) error {
	sop, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("获取SOP记录失败: %w", err)
	}

	// 删除文件
	if sop.ContentURL != "" {
		os.Remove(sop.ContentURL)
	}

	return s.repo.Delete(ctx, id)
}

// GetByProcessRoute 获取工艺路线关联的SOP列表
func (s *MesSopService) GetByProcessRoute(ctx context.Context, tenantID int64, processRouteId int64) ([]model.ElectronicSOP, error) {
	var list []model.ElectronicSOP
	db := s.repo.GetDB().WithContext(ctx).Model(&model.ElectronicSOP{}).Where("tenant_id = ? AND process_id = ?", tenantID, processRouteId)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, err
}

// GetByWorkOrder 获取工单关联的SOP
func (s *MesSopService) GetByWorkOrder(ctx context.Context, tenantID int64, workOrderId int64) ([]model.ElectronicSOP, error) {
	var list []model.ElectronicSOP
	db := s.repo.GetDB().WithContext(ctx).Model(&model.ElectronicSOP{}).Where("tenant_id = ? AND workstation_id = ?", tenantID, workOrderId)
	err := db.Order("created_at DESC").Find(&list).Error
	return list, err
}

// ListAll 获取所有SOP列表
func (s *MesSopService) ListAll(ctx context.Context, tenantID int64, query string) ([]model.ElectronicSOP, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

// GetDB 获取数据库连接（供外部使用）
func (s *MesSopService) GetDB() *repository.ElectronicSOPRepository {
	return s.repo
}

// ========== VO结构 ==========

// SopUploadReqVO SOP上传请求VO
type SopUploadReqVO struct {
	File           *multipart.FileHeader `form:"file" binding:"required"`
	SopName        string                `form:"sopName" binding:"required"`
	ProcessRouteId int64                 `form:"processRouteId"`
	WorkOrderId    int64                 `form:"workOrderId"`
	Version        string                `form:"version"`
}

// SopRespVO SOP响应VO
type SopRespVO struct {
	Id              int64  `json:"id"`
	SopName         string `json:"sopName"`
	FileName        string `json:"fileName"`
	FileSize        int64  `json:"fileSize"`
	FileUrl         string `json:"fileUrl"`
	ProcessRouteId  int64  `json:"processRouteId"`
	WorkOrderId     int64  `json:"workOrderId"`
	Version         string `json:"version"`
	UploadTime      string `json:"uploadTime"`
	Uploader        string `json:"uploader"`
}

// ToRespVO 转换为响应VO
func (s *MesSopService) ToRespVO(sop *model.ElectronicSOP) SopRespVO {
	return SopRespVO{
		Id:              int64(sop.ID),
		SopName:         sop.SopName,
		FileName:        filepath.Base(sop.ContentURL),
		FileSize:        0,
		FileUrl:         sop.ContentURL,
		ProcessRouteId:  sop.ProcessID,
		WorkOrderId:     sop.WorkstationID,
		Version:         sop.Version,
		UploadTime:      sop.CreatedAt.Format("2006-01-02 15:04:05"),
		Uploader:        sop.ApprovedBy,
	}
}
