package service

import (
	"context"
	"errors"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// MesJobReportLogService 报工记录服务
type MesJobReportLogService struct {
	repo *repository.MesJobReportLogRepository
}

func NewMesJobReportLogService(repo *repository.MesJobReportLogRepository) *MesJobReportLogService {
	return &MesJobReportLogService{repo: repo}
}

// Create 创建报工记录
func (s *MesJobReportLogService) Create(ctx context.Context, tenantID int64, userID int64, req *model.MesJobReportLogCreateReqVO) (*model.MesJobReportLog, error) {
	reportType := req.ReportType
	if reportType == "" {
		reportType = "NORMAL"
	}

	status := req.Status
	if status == "" {
		status = "PENDING"
	}

	var reportTime time.Time
	if req.ReportTime != "" {
		var err error
		reportTime, err = time.Parse("2006-01-02 15:04:05", req.ReportTime)
		if err != nil {
			reportTime, err = time.Parse("2006-01-02", req.ReportTime)
			if err != nil {
				return nil, errors.New("报工时间格式错误，请使用 YYYY-MM-DD 或 YYYY-MM-DD HH:mm:ss")
			}
		}
	} else {
		reportTime = time.Now()
	}

	m := &model.MesJobReportLog{
		TenantID:      tenantID,
		WorkOrderId:   req.WorkOrderId,
		WorkOrderCode: req.WorkOrderCode,
		ProcessCode:   req.ProcessCode,
		ProcessName:   req.ProcessName,
		ReportType:    reportType,
		Quantity:      req.Quantity,
		ReportTime:    reportTime,
		ReporterId:    req.ReporterId,
		ReporterName:  req.ReporterName,
		Remark:        req.Remark,
		Status:        status,
	}

	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

// Get 获取报工记录
func (s *MesJobReportLogService) Get(ctx context.Context, id uint64) (*model.MesJobReportLog, error) {
	m, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("报工记录不存在")
	}
	return m, nil
}

// Page 分页查询报工记录
func (s *MesJobReportLogService) Page(ctx context.Context, tenantID int64, query *model.MesJobReportLogQueryVO) ([]model.MesJobReportLog, int64, error) {
	return s.repo.Page(ctx, tenantID, query)
}

// Senior 高级搜索报工记录
func (s *MesJobReportLogService) Senior(ctx context.Context, tenantID int64, conditions []map[string]interface{}) ([]model.MesJobReportLog, int64, error) {
	return s.repo.Senior(ctx, tenantID, conditions)
}
