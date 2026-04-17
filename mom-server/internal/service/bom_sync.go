package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mom-server/internal/model"
	"mom-server/internal/repository"
	"net/http"
	"time"
)

// BomSyncService BOM金蝶同步服务
type BomSyncService struct {
	bomRepo     *repository.BOMRepository
	bomItemRepo *repository.BOMItemRepository
	materialRepo *repository.MaterialRepository
	httpClient  *HTTPClient
	syncURL     string
	appID       string
	appSecret   string
}

// HTTPClient 简化的HTTP客户端
type HTTPClient struct {
	timeout time.Duration
}

func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{timeout: timeout}
}

func (c *HTTPClient) Get(url string) ([]byte, error) {
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

func (c *HTTPClient) Post(url string, body []byte) ([]byte, error) {
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

// KingdeeBOMResponse 金蝶BOM响应
type KingdeeBOMResponse struct {
	ResultCode int             `json:"ResultCode"`
	ResultMsg  string         `json:"ResultMsg"`
	Data       json.RawMessage `json:"Data"`
}

// KingdeeBOMItem 金蝶BOM物料项
type KingdeeBOMItem struct {
	MaterialCode   string  `json:"MaterialCode"`
	MaterialName   string  `json:"MaterialName"`
	ParentCode     string  `json:"ParentCode"`
	ChildCode      string  `json:"ChildCode"`
	ChildName      string  `json:"ChildName"`
	Qty            float64 `json:"Qty"`
	Unit           string  `json:"Unit"`
	ErpBomCode     string  `json:"ErpBomCode"`
	EffectDate     string  `json:"EffectDate"`
	ExpireDate     string  `json:"ExpireDate"`
}

func NewBomSyncService(
	bomRepo *repository.BOMRepository,
	bomItemRepo *repository.BOMItemRepository,
	materialRepo *repository.MaterialRepository,
) *BomSyncService {
	return &BomSyncService{
		bomRepo:     bomRepo,
		bomItemRepo: bomItemRepo,
		materialRepo: materialRepo,
		httpClient:  NewHTTPClient(30 * time.Second),
	}
}

// SyncConfig 同步配置
type SyncConfig struct {
	SyncURL     string // 金蝶API地址
	AppID       string // 应用ID
	AppSecret   string // 应用密钥
	SyncType    string // FULL增量/全部
	LastSyncTime string // 上次同步时间
}

// StartSync 开始同步
func (s *BomSyncService) StartSync(ctx context.Context, config *SyncConfig) error {
	log.Printf("[BOM同步] 开始同步任务...")

	// 1. 获取金蝶BOM数据
	bomData, err := s.fetchBOMFromKingdee(ctx, config)
	if err != nil {
		return fmt.Errorf("获取金蝶BOM数据失败: %w", err)
	}

	// 2. 解析并转换数据
	items, err := s.parseBOMData(bomData)
	if err != nil {
		return fmt.Errorf("解析BOM数据失败: %w", err)
	}

	// 3. 保存到本地数据库
	savedCount, err := s.saveBOMItems(ctx, items)
	if err != nil {
		return fmt.Errorf("保存BOM数据失败: %w", err)
	}

	log.Printf("[BOM同步] 同步完成，新增/更新 %d 条BOM", savedCount)
	return nil
}

// fetchBOMFromKingdee 从金蝶获取BOM数据
func (s *BomSyncService) fetchBOMFromKingdee(ctx context.Context, config *SyncConfig) ([]byte, error) {
	url := fmt.Sprintf("%s/api/bom/list?app_id=%s&app_secret=%s&sync_type=%s&last_time=%s",
		config.SyncURL, config.AppID, config.AppSecret, config.SyncType, config.LastSyncTime)

	data, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	var resp KingdeeBOMResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("解析金蝶响应失败: %w", err)
	}

	if resp.ResultCode != 0 {
		return nil, fmt.Errorf("金蝶API错误: %s", resp.ResultMsg)
	}

	return resp.Data, nil
}

// parseBOMData 解析BOM数据
func (s *BomSyncService) parseBOMData(data []byte) ([]KingdeeBOMItem, error) {
	var items []KingdeeBOMItem
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// saveBOMItems 保存BOM项到数据库
func (s *BomSyncService) saveBOMItems(ctx context.Context, items []KingdeeBOMItem) (int, error) {
	savedCount := 0
	now := time.Now()

	for _, item := range items {
		// 查找或创建物料
		material, err := s.materialRepo.GetByCode(ctx, item.ChildCode)
		if err != nil {
			log.Printf("[BOM同步] 跳过物料 %s: 未找到", item.ChildCode)
			continue
		}

		// 查找父物料
		parentMaterial, err := s.materialRepo.GetByCode(ctx, item.ParentCode)
		if err != nil {
			log.Printf("[BOM同步] 跳过父物料 %s: 未找到", item.ParentCode)
			continue
		}

		// 查找是否已存在BOM
		existingBOM, _ := s.bomRepo.GetByMaterialID(ctx, parentMaterial.ID)
		if existingBOM != nil {
			// 更新现有BOM
			updates := map[string]any{
				"erp_bom_code":     item.ErpBomCode,
				"erp_sync_time":    now,
				"erp_sync_status":  "SYNCED",
			}
			s.bomRepo.Update(ctx, uint(existingBOM.ID), updates)
		} else {
			// 创建新BOM
			bom := &model.MdmBOM{
				TenantID:      1,
				BOMCode:      fmt.Sprintf("BOM%d", time.Now().UnixNano()),
				BOMName:      fmt.Sprintf("BOM-%s", parentMaterial.MaterialCode),
				MaterialID:    parentMaterial.ID,
				MaterialCode: parentMaterial.MaterialCode,
				MaterialName: parentMaterial.MaterialName,
				Version:       "1.0",
				Status:        "ACTIVE",
				ErpBomCode:    item.ErpBomCode,
				ErpSyncTime:   &now,
				ErpSyncStatus: "SYNCED",
				IsCurrent:     1,
			}
			s.bomRepo.Create(ctx, bom)

			// 创建BOM项
			bomItem := &model.MdmBOMItem{
				TenantID:       1,
				BOMID:         int64(bom.ID),
				MaterialID:    material.ID,
				MaterialCode:  material.MaterialCode,
				MaterialName:  material.MaterialName,
				Quantity:      item.Qty,
				Unit:          item.Unit,
				LineNo:        1,
			}
			s.bomItemRepo.Create(ctx, bomItem)
			savedCount++
		}
	}

	return savedCount, nil
}

// StartSyncWorker 启动同步Worker (后台定时任务)
func (s *BomSyncService) StartSyncWorker(ctx context.Context, interval time.Duration, config *SyncConfig) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 立即执行一次
	s.doSync(ctx, config)

	for {
		select {
		case <-ctx.Done():
			log.Printf("[BOM同步] 同步Worker已停止")
			return
		case <-ticker.C:
			s.doSync(ctx, config)
		}
	}
}

func (s *BomSyncService) doSync(ctx context.Context, config *SyncConfig) {
	if err := s.StartSync(ctx, config); err != nil {
		log.Printf("[BOM同步] 同步失败: %v", err)
	}
}
