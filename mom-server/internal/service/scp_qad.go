package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ScpQadService QAD对接服务
type ScpQadService struct {
	syncLogRepo *repository.ScpQadSyncLogRepository
}

// NewScpQadService 创建QAD对接服务
func NewScpQadService(syncLogRepo *repository.ScpQadSyncLogRepository) *ScpQadService {
	return &ScpQadService{
		syncLogRepo: syncLogRepo,
	}
}

// SyncToQad 同步数据到QAD，记录同步请求并返回syncId
func (s *ScpQadService) SyncToQad(ctx context.Context, tenantID int64, req *model.QadSyncRequest) (*model.QadSyncResponse, error) {
	reqContent, _ := json.Marshal(req.Data)
	syncLog := &model.ScpQadSyncLog{
		TenantID:       tenantID,
		SyncType:       req.SyncType,
		SyncDirection:  model.QadSyncDirection.Upload,
		MomDocNo:       req.DocNo,
		Status:         model.QadSyncStatus.Pending,
		RequestContent: string(reqContent),
		SyncTime:       time.Now(),
	}

	if err := s.syncLogRepo.Create(ctx, syncLog); err != nil {
		return nil, fmt.Errorf("创建QAD同步记录失败: %w", err)
	}

	// TODO: 实际调用QAD系统接口推送数据
	// 当前以PENDING状态落库，由异步任务或QAD回调更新最终状态
	return &model.QadSyncResponse{
		SyncID:  syncLog.ID,
		Status:  model.QadSyncStatus.Pending,
		Message: "同步请求已提交",
	}, nil
}

// GetSyncStatus 查询单条同步状态
func (s *ScpQadService) GetSyncStatus(ctx context.Context, syncID uint64) (*model.ScpQadSyncLog, error) {
	return s.syncLogRepo.GetByID(ctx, syncID)
}

// GetSyncLogByDocNo 根据单据号查询所有同步日志
func (s *ScpQadService) GetSyncLogByDocNo(ctx context.Context, docNo string) ([]model.ScpQadSyncLog, error) {
	return s.syncLogRepo.ListByDocNo(ctx, docNo)
}

// HandleQadConfirm 处理QAD订单确认回调，更新本地订单状态
func (s *ScpQadService) HandleQadConfirm(ctx context.Context, tenantID int64, req *model.QadConfirmRequest) error {
	reqContent, _ := json.Marshal(req)
	syncLog := &model.ScpQadSyncLog{
		TenantID:       tenantID,
		SyncType:       model.QadSyncType.Order,
		SyncDirection:  model.QadSyncDirection.Download,
		QadDocNo:       req.QadDocNo,
		MomDocNo:       req.MomDocNo,
		Status:         model.QadSyncStatus.Success,
		RequestContent: string(reqContent),
		SyncTime:       time.Now(),
	}

	if err := s.syncLogRepo.Create(ctx, syncLog); err != nil {
		return fmt.Errorf("记录QAD确认回调失败: %w", err)
	}

	// TODO: 根据 req.Status (CONFIRMED/REJECTED) 更新本地采购订单或销售订单状态
	return nil
}

// HandleQadDelivery 处理QAD发货通知回调，生成收货记录
func (s *ScpQadService) HandleQadDelivery(ctx context.Context, tenantID int64, req *model.QadDeliveryRequest) error {
	reqContent, _ := json.Marshal(req)
	syncLog := &model.ScpQadSyncLog{
		TenantID:       tenantID,
		SyncType:       model.QadSyncType.Delivery,
		SyncDirection:  model.QadSyncDirection.Download,
		QadDocNo:       req.QadDocNo,
		MomDocNo:       req.MomDocNo,
		Status:         model.QadSyncStatus.Success,
		RequestContent: string(reqContent),
		SyncTime:       time.Now(),
	}

	if err := s.syncLogRepo.Create(ctx, syncLog); err != nil {
		return fmt.Errorf("记录QAD发货通知失败: %w", err)
	}

	// TODO: 根据发货信息生成WMS收货记录，更新库存
	return nil
}
