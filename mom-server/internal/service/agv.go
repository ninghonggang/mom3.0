package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// AGVService AGV服务
type AGVService struct {
	taskRepo      *repository.AGVTaskRepository
	deviceRepo    *repository.AGVDeviceRepository
	locationRepo  *repository.AGVLocationMappingRepository
	httpClient    *HTTPClient
	agvBaseURL   string
	agvAPIKey    string
}

// NewAGVService 创建AGV服务
func NewAGVService(
	taskRepo *repository.AGVTaskRepository,
	deviceRepo *repository.AGVDeviceRepository,
	locationRepo *repository.AGVLocationMappingRepository,
) *AGVService {
	return &AGVService{
		taskRepo:     taskRepo,
		deviceRepo:   deviceRepo,
		locationRepo: locationRepo,
		httpClient:   NewHTTPClient(30 * time.Second),
	}
}

// ConfigAGV 配置AGV系统连接
func (s *AGVService) ConfigAGV(baseURL, apiKey string) {
	s.agvBaseURL = baseURL
	s.agvAPIKey = apiKey
}

// CreateDeliveryTask 创建配送任务
func (s *AGVService) CreateDeliveryTask(ctx context.Context, req *model.CreateAGVTaskRequest) (*model.AGVTask, error) {
	// 生成任务编号
	taskNo, err := s.taskRepo.GenerateTaskNo(ctx)
	if err != nil {
		return nil, fmt.Errorf("生成任务编号失败: %w", err)
	}

	extData := ""
	if req.RelatedOrderNo != "" {
		data := map[string]string{
			"orderNo": req.RelatedOrderNo,
			"type":    req.RelatedOrderType,
		}
		bytes, _ := json.Marshal(data)
		extData = string(bytes)
	}

	task := &model.AGVTask{
		TenantID:         req.TenantID,
		TaskNo:          taskNo,
		TaskType:        model.AGVTaskType(req.TaskType),
		Status:          model.AGVTaskStatusPending,
		Priority:        req.Priority,
		SourceLocationID: req.SourceLocationID,
		SourceLocation:   req.SourceLocation,
		TargetLocationID: req.TargetLocationID,
		TargetLocation:   req.TargetLocation,
		MaterialID:       req.MaterialID,
		MaterialCode:     req.MaterialCode,
		MaterialName:     req.MaterialName,
		Quantity:         req.Quantity,
		Unit:             req.Unit,
		RelatedOrderNo:   req.RelatedOrderNo,
		RelatedOrderType: req.RelatedOrderType,
		ExtData:         extData,
	}

	if err := s.taskRepo.Create(ctx, task); err != nil {
		return nil, fmt.Errorf("创建AGV任务失败: %w", err)
	}

	// 如果配置了AGV系统地址，尝试自动分配
	if s.agvBaseURL != "" {
		go s.autoAssignTask(context.Background(), task)
	}

	return task, nil
}

// autoAssignTask 自动分配任务到AGV
func (s *AGVService) autoAssignTask(ctx context.Context, task *model.AGVTask) {
	// 查询可用AGV
	agvs, err := s.deviceRepo.ListAvailable(ctx, task.TenantID)
	if err != nil || len(agvs) == 0 {
		log.Printf("[AGV] 未找到可用AGV设备，任务 %s 等待分配", task.TaskNo)
		return
	}

	// 选择电量最充足的AGV
	agv := agvs[0]
	if err := s.taskRepo.AssignAGV(ctx, task.ID, agv.AGVCode, agv.AGVName); err != nil {
		log.Printf("[AGV] 分配AGV失败: %v", err)
		return
	}

	// 调用AGV系统创建任务
	if s.agvBaseURL != "" {
		if err := s.callAGVCreateTask(ctx, &agv, task); err != nil {
			log.Printf("[AGV] 调用AGV系统失败: %v", err)
		}
	}

	log.Printf("[AGV] 任务 %s 已分配给AGV %s", task.TaskNo, agv.AGVCode)
}

// callAGVCreateTask 调用AGV系统创建任务
func (s *AGVService) callAGVCreateTask(ctx context.Context, agv *model.AGVDevice, task *model.AGVTask) error {
	if s.agvBaseURL == "" {
		return nil
	}

	url := s.agvBaseURL + "/api/v1/agv/task/create"
	body := map[string]any{
		"task_type":      task.TaskType,
		"pickup_point":   task.SourceLocation,
		"delivery_point":  task.TargetLocation,
		"material_code":   task.MaterialCode,
		"material_name":   task.MaterialName,
		"qty":            task.Quantity,
		"unit":           task.Unit,
		"priority":        task.Priority,
		"agv_code":       agv.AGVCode,
		"callback_url":   "", // TODO: 配置回调地址
	}

	bytes, _ := json.Marshal(body)
	resp, err := s.httpClient.Post(url, bytes)
	if err != nil {
		return fmt.Errorf("调用AGV系统失败: %w", err)
	}

	log.Printf("[AGV] AGV系统响应: %s", string(resp))
	return nil
}

// CancelTask 取消任务
func (s *AGVService) CancelTask(ctx context.Context, taskID int64) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	// 只有PENDING和ASSIGNED状态可以取消
	if task.Status != model.AGVTaskStatusPending && task.Status != model.AGVTaskStatusAssigned {
		return fmt.Errorf("当前状态 %s 不允许取消", task.Status)
	}

	// 调用AGV系统取消任务
	if s.agvBaseURL != "" && task.AssignedAGVCode != "" {
		if err := s.callAGVCancelTask(ctx, task.AssignedAGVCode, task.TaskNo); err != nil {
			log.Printf("[AGV] 取消AGV任务失败: %v", err)
		}
	}

	return s.taskRepo.UpdateStatus(ctx, taskID, model.AGVTaskStatusCancelled)
}

// callAGVCancelTask 调用AGV系统取消任务
func (s *AGVService) callAGVCancelTask(ctx context.Context, agvCode, taskNo string) error {
	if s.agvBaseURL == "" {
		return nil
	}

	url := fmt.Sprintf("%s/api/v1/agv/task/%s/cancel", s.agvBaseURL, taskNo)
	_, err := s.httpClient.Post(url, nil)
	return err
}

// GetTaskStatus 获取任务状态
func (s *AGVService) GetTaskStatus(ctx context.Context, taskID int64) (*model.AGVTask, error) {
	return s.taskRepo.GetByID(ctx, taskID)
}

// GetTaskByNo 根据编号获取任务
func (s *AGVService) GetTaskByNo(ctx context.Context, taskNo string) (*model.AGVTask, error) {
	return s.taskRepo.GetByTaskNo(ctx, taskNo)
}

// ListTasks 查询任务列表
func (s *AGVService) ListTasks(ctx context.Context, q *model.AGVTaskQuery) ([]model.AGVTask, int64, error) {
	return s.taskRepo.List(ctx, q)
}

// UpdateTaskStatus 更新任务状态（供回调使用）
func (s *AGVService) UpdateTaskStatus(ctx context.Context, taskNo string, status string, errorMsg string) error {
	task, err := s.taskRepo.GetByTaskNo(ctx, taskNo)
	if err != nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	updates := map[string]interface{}{
		"status": status,
	}
	if errorMsg != "" {
		updates["error_message"] = errorMsg
	}
	if status == string(model.AGVTaskStatusCompleted) {
		now := time.Now()
		updates["completed_at"] = &now
	}

	return s.taskRepo.Update(ctx, task.ID, updates)
}

// CompleteTask 完成AGV任务
func (s *AGVService) CompleteTask(ctx context.Context, taskID int64) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	if task.Status != model.AGVTaskStatusInProgress {
		return fmt.Errorf("当前状态 %s 不允许完成", task.Status)
	}

	return s.taskRepo.UpdateStatus(ctx, taskID, model.AGVTaskStatusCompleted)
}

// StartTask 开始AGV任务
func (s *AGVService) StartTask(ctx context.Context, taskID int64) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	if task.Status != model.AGVTaskStatusAssigned {
		return fmt.Errorf("当前状态 %s 不允许开始", task.Status)
	}

	return s.taskRepo.UpdateStatus(ctx, taskID, model.AGVTaskStatusInProgress)
}

// AssignTaskToAGV 分配任务给AGV
func (s *AGVService) AssignTaskToAGV(ctx context.Context, taskID int64, agvCode string) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	if task.Status != model.AGVTaskStatusPending {
		return fmt.Errorf("当前状态 %s 不允许分配", task.Status)
	}

	agv, err := s.deviceRepo.GetByCode(ctx, agvCode)
	if err != nil {
		return fmt.Errorf("AGV设备不存在: %w", err)
	}

	if agv.Status != model.AGVDeviceStatusOnline {
		return fmt.Errorf("AGV设备 %s 当前不在线", agvCode)
	}

	// 调用AGV系统
	if s.agvBaseURL != "" {
		if err := s.callAGVCreateTask(ctx, agv, task); err != nil {
			return fmt.Errorf("调用AGV系统失败: %w", err)
		}
	}

	return s.taskRepo.AssignAGV(ctx, taskID, agv.AGVCode, agv.AGVName)
}

// GetAvailableAGVs 获取可用AGV列表
func (s *AGVService) GetAvailableAGVs(ctx context.Context, tenantID int64) ([]model.AGVDevice, error) {
	return s.deviceRepo.ListAvailable(ctx, tenantID)
}

// ============ AGV设备管理 ============

// CreateDevice 创建设备
func (s *AGVService) CreateDevice(ctx context.Context, req *model.CreateAGVDeviceRequest) (*model.AGVDevice, error) {
	device := &model.AGVDevice{
		TenantID: req.TenantID,
		AGVCode:  req.AGVCode,
		AGVName:  req.AGVName,
		AGVType:  req.AGVType,
		Status:   model.AGVDeviceStatusOffline,
		MaxLoad:  req.MaxLoad,
		Remark:   req.Remark,
	}

	if err := s.deviceRepo.Create(ctx, device); err != nil {
		return nil, fmt.Errorf("创建设备失败: %w", err)
	}
	return device, nil
}

// GetDevice 获取设备
func (s *AGVService) GetDevice(ctx context.Context, id int64) (*model.AGVDevice, error) {
	return s.deviceRepo.GetByID(ctx, id)
}

// ListDevices 查询设备列表
func (s *AGVService) ListDevices(ctx context.Context, q *model.AGVDeviceQuery) ([]model.AGVDevice, int64, error) {
	return s.deviceRepo.List(ctx, q)
}

// UpdateDeviceStatus 更新设备状态（供心跳回调使用）
func (s *AGVService) UpdateDeviceStatus(ctx context.Context, agvCode string, status string, batteryLevel float64, currentLocation string) error {
	// 更新状态
	s.deviceRepo.UpdateStatus(ctx, agvCode, model.AGVDeviceStatus(status))
	// 更新电量
	if batteryLevel > 0 {
		s.deviceRepo.UpdateHeartbeat(ctx, agvCode, batteryLevel)
	}
	// 更新位置
	if currentLocation != "" {
		s.deviceRepo.Update(ctx, 0, map[string]interface{}{
			"current_location": currentLocation,
		})
	}
	return nil
}

// RegisterAGVHeartbeat AGV心跳注册
func (s *AGVService) RegisterAGVHeartbeat(ctx context.Context, agvCode string, batteryLevel float64, currentLocation string) error {
	device, err := s.deviceRepo.GetByCode(ctx, agvCode)
	if err != nil {
		// 设备不存在，自动注册
		if s.agvBaseURL != "" {
			// 尝试从AGV系统获取设备信息
			info, err := s.fetchAGVInfoFromSystem(ctx, agvCode)
			if err != nil {
				return fmt.Errorf("AGV设备未注册且无法获取信息: %w", err)
			}
			device = info
		} else {
			return fmt.Errorf("AGV设备未注册: %s", agvCode)
		}
	}

	updates := map[string]interface{}{
		"status":           model.AGVDeviceStatusOnline,
		"battery_level":    batteryLevel,
		"last_heartbeat":   time.Now(),
	}
	if currentLocation != "" {
		updates["current_location"] = currentLocation
	}

	return s.deviceRepo.Update(ctx, device.ID, updates)
}

// fetchAGVInfoFromSystem 从AGV系统获取设备信息
func (s *AGVService) fetchAGVInfoFromSystem(ctx context.Context, agvCode string) (*model.AGVDevice, error) {
	if s.agvBaseURL == "" {
		return nil, fmt.Errorf("AGV系统未配置")
	}

	url := fmt.Sprintf("%s/api/v1/agv/device/%s/info", s.agvBaseURL, agvCode)
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			AGVCode  string  `json:"agv_code"`
			AGVName  string  `json:"agv_name"`
			AGVType  string  `json:"agv_type"`
			MaxLoad  float64 `json:"max_load"`
		} `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, fmt.Errorf("获取AGV信息失败: %s", result.Message)
	}

	device := &model.AGVDevice{
		TenantID: 1,
		AGVCode:  result.Data.AGVCode,
		AGVName:  result.Data.AGVName,
		AGVType:  result.Data.AGVType,
		MaxLoad:  result.Data.MaxLoad,
		Status:   model.AGVDeviceStatusOnline,
	}

	if err := s.deviceRepo.Create(ctx, device); err != nil {
		return device, nil // 返回已存在的设备
	}

	return device, nil
}

// ============ AGV库位映射管理 ============

// CreateLocation 创建库位映射
func (s *AGVService) CreateLocation(ctx context.Context, req *model.CreateAGVLocationRequest) (*model.AGVLocationMapping, error) {
	mapping := &model.AGVLocationMapping{
		TenantID:     req.TenantID,
		LocationCode: req.LocationCode,
		LocationName: req.LocationName,
		LocationType: model.AGVLocationType(req.LocationType),
		AGVLocationCode: req.AGVLocationCode,
		XCoord:       req.XCoord,
		YCoord:       req.YCoord,
		Priority:     req.Priority,
		Enabled:      req.Enabled,
	}

	if err := s.locationRepo.Create(ctx, mapping); err != nil {
		return nil, fmt.Errorf("创建库位映射失败: %w", err)
	}
	return mapping, nil
}

// GetLocation 获取库位映射
func (s *AGVService) GetLocation(ctx context.Context, id int64) (*model.AGVLocationMapping, error) {
	return s.locationRepo.GetByID(ctx, id)
}

// ListLocations 查询库位映射列表
func (s *AGVService) ListLocations(ctx context.Context, q *model.AGVLocationQuery) ([]model.AGVLocationMapping, int64, error) {
	return s.locationRepo.List(ctx, q)
}

// UpdateLocation 更新库位映射
func (s *AGVService) UpdateLocation(ctx context.Context, id int64, updates map[string]interface{}) error {
	return s.locationRepo.Update(ctx, id, updates)
}

// DeleteLocation 删除库位映射
func (s *AGVService) DeleteLocation(ctx context.Context, id int64) error {
	return s.locationRepo.Delete(ctx, id)
}

// GetLocationByCode 根据库位编码获取映射
func (s *AGVService) GetLocationByCode(ctx context.Context, locationCode string) (*model.AGVLocationMapping, error) {
	return s.locationRepo.GetByLocationCode(ctx, locationCode)
}

// AGVCallbackResult AGV回调结果
type AGVCallbackResult struct {
	TaskNo   string `json:"task_no"`
	AGVCode  string `json:"agv_code"`
	Status   string `json:"status"` // SUBMITTED/IN_PROGRESS/COMPLETED/EXCEPTION
	Position string `json:"position"`
	Message  string `json:"message"`
}

// HandleAGVCallback 处理AGV回调
func (s *AGVService) HandleAGVCallback(ctx context.Context, result *AGVCallbackResult) error {
	task, err := s.taskRepo.GetByTaskNo(ctx, result.TaskNo)
	if err != nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	updates := map[string]interface{}{}

	switch result.Status {
	case "SUBMITTED":
		updates["status"] = model.AGVTaskStatusAssigned
	case "IN_PROGRESS":
		updates["status"] = model.AGVTaskStatusInProgress
		now := time.Now()
		updates["started_at"] = &now
	case "COMPLETED":
		updates["status"] = model.AGVTaskStatusCompleted
		now := time.Now()
		updates["completed_at"] = &now
	case "EXCEPTION":
		updates["status"] = model.AGVTaskStatusException
		updates["error_message"] = result.Message
	default:
		return fmt.Errorf("未知状态: %s", result.Status)
	}

	return s.taskRepo.Update(ctx, task.ID, updates)
}
