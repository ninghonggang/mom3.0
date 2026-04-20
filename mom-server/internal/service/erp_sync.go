package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ERPSyncService ERP同步服务
type ERPSyncService struct {
	syncLogRepo  *repository.IntegrationERPSyncLogRepository
	mappingRepo  *repository.IntegrationERPMappingRepository
	materialRepo *repository.MaterialRepository
	bomRepo      *repository.BOMRepository
	bomItemRepo  *repository.BOMItemRepository
	prodRepo     *repository.ProductionOrderRepository
	httpClient   *ERPHTTPClient
	erpBaseURL   string
	erpAPIKey    string
}

// ERPHTTPClient ERP HTTP客户端
type ERPHTTPClient struct {
	timeout time.Duration
}

func NewERPHTTPClient(timeout time.Duration) *ERPHTTPClient {
	return &ERPHTTPClient{timeout: timeout}
}

func (c *ERPHTTPClient) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP状态码异常: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func (c *ERPHTTPClient) Post(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP状态码异常: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// NewERPSyncService 创建ERP同步服务
func NewERPSyncService(
	syncLogRepo *repository.IntegrationERPSyncLogRepository,
	mappingRepo *repository.IntegrationERPMappingRepository,
	materialRepo *repository.MaterialRepository,
	bomRepo *repository.BOMRepository,
	bomItemRepo *repository.BOMItemRepository,
	prodRepo *repository.ProductionOrderRepository,
) *ERPSyncService {
	return &ERPSyncService{
		syncLogRepo:  syncLogRepo,
		mappingRepo: mappingRepo,
		materialRepo: materialRepo,
		bomRepo:     bomRepo,
		bomItemRepo: bomItemRepo,
		prodRepo:    prodRepo,
		httpClient:  NewERPHTTPClient(30 * time.Second),
	}
}

// ConfigERP 配置ERP连接
func (s *ERPSyncService) ConfigERP(baseURL, apiKey string) {
	s.erpBaseURL = baseURL
	s.erpAPIKey = apiKey
}

// ListSyncLogs 查询同步日志
func (s *ERPSyncService) ListSyncLogs(ctx context.Context, q *model.ERPSyncLogQuery) ([]model.IntegrationERPSyncLog, int64, error) {
	return s.syncLogRepo.List(ctx, q)
}

// GetSyncLog 获取同步日志
func (s *ERPSyncService) GetSyncLog(ctx context.Context, id int64) (*model.IntegrationERPSyncLog, error) {
	return s.syncLogRepo.GetByID(ctx, id)
}

// GetSyncStatus 获取同步状态
func (s *ERPSyncService) GetSyncStatus(ctx context.Context, id int64) (*model.IntegrationERPSyncLog, error) {
	return s.syncLogRepo.GetByID(ctx, id)
}

// createSyncLog 创建同步日志
func (s *ERPSyncService) createSyncLog(ctx context.Context, syncType, direction string, req interface{}, erpBillNo, mesBillNo string) (*model.IntegrationERPSyncLog, error) {
	reqBody, _ := json.Marshal(req)
	log := &model.IntegrationERPSyncLog{
		SyncType:    syncType,
		Direction:   direction,
		ERPBillNo:   erpBillNo,
		MESBillNo:   mesBillNo,
		RequestBody: string(reqBody),
		Status:      model.ERPSyncStatus.Pending,
		TenantID:    1,
	}
	if err := s.syncLogRepo.Create(ctx, log); err != nil {
		return nil, err
	}
	return log, nil
}

// updateSyncLog 更新同步日志
func (s *ERPSyncService) updateSyncLog(ctx context.Context, id int64, respBody string, status, errorMsg string) error {
	updates := map[string]interface{}{
		"response_body": respBody,
		"status":        status,
	}
	if errorMsg != "" {
		updates["error_msg"] = errorMsg
	}
	return s.syncLogRepo.Update(ctx, id, updates)
}

// SyncBOM BOM数据同步(ERP→MES)
func (s *ERPSyncService) SyncBOM(ctx context.Context, req *model.ERPSyncBOMRequest) (*model.ERPSyncResult, error) {
	// 创建同步日志
	syncLog, err := s.createSyncLog(ctx, model.ERPSyncType.BOM, model.ERPSyncDirection.Inbound, req, req.ERPBillNo, "")
	if err != nil {
		return nil, fmt.Errorf("创建同步日志失败: %w", err)
	}

	// 解析并保存BOM
	savedCount := 0
	for _, item := range req.Items {
		// 查找子物料
		childMat, err := s.materialRepo.GetByCode(ctx, item.ChildCode)
		if err != nil {
			continue
		}
		// 查找父物料
		parentMat, err := s.materialRepo.GetByCode(ctx, item.ParentCode)
		if err != nil {
			continue
		}

		now := time.Now()
		bom := &model.MdmBOM{
			TenantID:      1,
			BOMCode:      fmt.Sprintf("BOM-%s-%d", item.ParentCode, now.Unix()),
			BOMName:      fmt.Sprintf("BOM-%s", parentMat.MaterialCode),
			MaterialID:    parentMat.ID,
			MaterialCode: parentMat.MaterialCode,
			MaterialName: parentMat.MaterialName,
			Version:       "1.0",
			Status:        "ACTIVE",
			ErpBomCode:    req.ERPBillNo,
			ErpSyncTime:   &now,
			ErpSyncStatus: "SYNCED",
			IsCurrent:     1,
		}
		s.bomRepo.Create(ctx, bom)

		bomItem := &model.MdmBOMItem{
			TenantID:      1,
			BOMID:        int64(bom.ID),
			MaterialID:   childMat.ID,
			MaterialCode: childMat.MaterialCode,
			MaterialName: childMat.MaterialName,
			Quantity:      item.Qty,
			Unit:         item.Unit,
			LineNo:       savedCount + 1,
		}
		s.bomItemRepo.Create(ctx, bomItem)
		savedCount++
	}

	// 更新日志
	resp, _ := json.Marshal(map[string]interface{}{"savedCount": savedCount})
	s.updateSyncLog(ctx, syncLog.ID, string(resp), model.ERPSyncStatus.Success, "")

	return &model.ERPSyncResult{
		Success:   true,
		ERPBillNo: req.ERPBillNo,
		Message:   fmt.Sprintf("成功同步 %d 条BOM项", savedCount),
	}, nil
}

// SyncProductionOrder 生产订单同步(ERP→MES)
func (s *ERPSyncService) SyncProductionOrder(ctx context.Context, req *model.ERPSyncProductionOrderRequest) (*model.ERPSyncResult, error) {
	syncLog, err := s.createSyncLog(ctx, model.ERPSyncType.ProductionOrder, model.ERPSyncDirection.Inbound, req, req.ERPBillNo, req.OrderNo)
	if err != nil {
		return nil, fmt.Errorf("创建同步日志失败: %w", err)
	}

	// TODO: 调用生产订单服务创建/更新订单
	resp, _ := json.Marshal(map[string]interface{}{"orderNo": req.OrderNo})
	s.updateSyncLog(ctx, syncLog.ID, string(resp), model.ERPSyncStatus.Success, "")

	return &model.ERPSyncResult{
		Success:   true,
		ERPBillNo: req.ERPBillNo,
		MESBillNo: req.OrderNo,
		Message:   "生产订单同步成功",
	}, nil
}

// SyncStock 库存同步(ERP→MES)
func (s *ERPSyncService) SyncStock(ctx context.Context, req *model.ERPSyncStockRequest) (*model.ERPSyncResult, error) {
	syncLog, err := s.createSyncLog(ctx, model.ERPSyncType.Stock, model.ERPSyncDirection.Inbound, req, "", "")
	if err != nil {
		return nil, fmt.Errorf("创建同步日志失败: %w", err)
	}

	// TODO: 调用库存服务更新库存
	resp, _ := json.Marshal(map[string]interface{}{"materialCode": req.MaterialCode})
	s.updateSyncLog(ctx, syncLog.ID, string(resp), model.ERPSyncStatus.Success, "")

	return &model.ERPSyncResult{
		Success:   true,
		Message:   "库存同步成功",
	}, nil
}

// PushReport 报工回传(MES→ERP)
func (s *ERPSyncService) PushReport(ctx context.Context, req *model.ERPPushReportRequest) (*model.ERPSyncResult, error) {
	syncLog, err := s.createSyncLog(ctx, model.ERPSyncType.Report, model.ERPSyncDirection.Outbound, req, "", req.MESBillNo)
	if err != nil {
		return nil, fmt.Errorf("创建同步日志失败: %w", err)
	}

	// 调用ERP系统回传报工
	if s.erpBaseURL != "" {
		resp, err := s.callERPPush(ctx, "/api/v1/mes/report/push", req)
		if err != nil {
			s.updateSyncLog(ctx, syncLog.ID, "", model.ERPSyncStatus.Failed, err.Error())
			return &model.ERPSyncResult{
				Success:   false,
				MESBillNo: req.MESBillNo,
				Message:   err.Error(),
			}, nil
		}
		s.updateSyncLog(ctx, syncLog.ID, string(resp), model.ERPSyncStatus.Success, "")
	} else {
		s.updateSyncLog(ctx, syncLog.ID, "{}", model.ERPSyncStatus.Success, "ERP未配置")
	}

	return &model.ERPSyncResult{
		Success:   true,
		MESBillNo: req.MESBillNo,
		Message:   "报工回传成功",
	}, nil
}

// PushStockIn 入库通知回传(MES→ERP)
func (s *ERPSyncService) PushStockIn(ctx context.Context, req *model.ERPPushStockInRequest) (*model.ERPSyncResult, error) {
	syncLog, err := s.createSyncLog(ctx, model.ERPSyncType.StockIn, model.ERPSyncDirection.Outbound, req, "", req.MESBillNo)
	if err != nil {
		return nil, fmt.Errorf("创建同步日志失败: %w", err)
	}

	if s.erpBaseURL != "" {
		resp, err := s.callERPPush(ctx, "/api/v1/mes/stockin/push", req)
		if err != nil {
			s.updateSyncLog(ctx, syncLog.ID, "", model.ERPSyncStatus.Failed, err.Error())
			return &model.ERPSyncResult{
				Success:   false,
				MESBillNo: req.MESBillNo,
				Message:   err.Error(),
			}, nil
		}
		s.updateSyncLog(ctx, syncLog.ID, string(resp), model.ERPSyncStatus.Success, "")
	} else {
		s.updateSyncLog(ctx, syncLog.ID, "{}", model.ERPSyncStatus.Success, "ERP未配置")
	}

	return &model.ERPSyncResult{
		Success:   true,
		MESBillNo: req.MESBillNo,
		Message:   "入库通知回传成功",
	}, nil
}

// PushQualityData 质检数据回传(MES→ERP)
func (s *ERPSyncService) PushQualityData(ctx context.Context, req *model.ERPPushQualityRequest) (*model.ERPSyncResult, error) {
	syncLog, err := s.createSyncLog(ctx, model.ERPSyncType.Quality, model.ERPSyncDirection.Outbound, req, "", req.InspectNo)
	if err != nil {
		return nil, fmt.Errorf("创建同步日志失败: %w", err)
	}

	if s.erpBaseURL != "" {
		resp, err := s.callERPPush(ctx, "/api/v1/mes/quality/push", req)
		if err != nil {
			s.updateSyncLog(ctx, syncLog.ID, "", model.ERPSyncStatus.Failed, err.Error())
			return &model.ERPSyncResult{
				Success: false,
				Message: err.Error(),
			}, nil
		}
		s.updateSyncLog(ctx, syncLog.ID, string(resp), model.ERPSyncStatus.Success, "")
	} else {
		s.updateSyncLog(ctx, syncLog.ID, "{}", model.ERPSyncStatus.Success, "ERP未配置")
	}

	return &model.ERPSyncResult{
		Success:   true,
		Message:   "质检数据回传成功",
	}, nil
}

// callERPPush 调用ERP推送接口
func (s *ERPSyncService) callERPPush(ctx context.Context, path string, req interface{}) ([]byte, error) {
	if s.erpBaseURL == "" {
		return []byte("{}"), nil
	}

	url := s.erpBaseURL + path
	bytes, _ := json.Marshal(req)
	resp, err := s.httpClient.Post(url, bytes)
	if err != nil {
		return nil, fmt.Errorf("调用ERP接口失败: %w", err)
	}
	return resp, nil
}
