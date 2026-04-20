package repository

import (
	"context"
	"fmt"
	"time"

	"mom-server/internal/model"

	"gorm.io/gorm"
)

// AGVTaskRepository AGV任务仓库
type AGVTaskRepository struct {
	db *gorm.DB
}

func NewAGVTaskRepository(db *gorm.DB) *AGVTaskRepository {
	return &AGVTaskRepository{db: db}
}

// Create 创建AGV任务
func (r *AGVTaskRepository) Create(ctx context.Context, task *model.AGVTask) error {
	return r.db.WithContext(ctx).Create(task).Error
}

// GetByID 根据ID查询
func (r *AGVTaskRepository) GetByID(ctx context.Context, id int64) (*model.AGVTask, error) {
	var task model.AGVTask
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetByTaskNo 根据任务编号查询
func (r *AGVTaskRepository) GetByTaskNo(ctx context.Context, taskNo string) (*model.AGVTask, error) {
	var task model.AGVTask
	err := r.db.WithContext(ctx).Where("task_no = ?", taskNo).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Update 更新AGV任务
func (r *AGVTaskRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AGVTask{}).Where("id = ?", id).Updates(updates).Error
}

// UpdateStatus 更新任务状态
func (r *AGVTaskRepository) UpdateStatus(ctx context.Context, id int64, status model.AGVTaskStatus) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == model.AGVTaskStatusInProgress {
		now := time.Now()
		updates["started_at"] = &now
	}
	if status == model.AGVTaskStatusCompleted {
		now := time.Now()
		updates["completed_at"] = &now
	}
	return r.db.WithContext(ctx).Model(&model.AGVTask{}).Where("id = ?", id).Updates(updates).Error
}

// AssignAGV 分配AGV
func (r *AGVTaskRepository) AssignAGV(ctx context.Context, id int64, agvCode, agvName string) error {
	return r.db.WithContext(ctx).Model(&model.AGVTask{}).Where("id = ?", id).Updates(map[string]interface{}{
		"assigned_agv_code": agvCode,
		"assigned_agv_name": agvName,
		"status":            model.AGVTaskStatusAssigned,
	}).Error
}

// List 查询任务列表
func (r *AGVTaskRepository) List(ctx context.Context, q *model.AGVTaskQuery) ([]model.AGVTask, int64, error) {
	var tasks []model.AGVTask
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AGVTask{})

	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.TaskNo != "" {
		query = query.Where("task_no LIKE ?", "%"+q.TaskNo+"%")
	}
	if q.TaskType != "" {
		query = query.Where("task_type = ?", q.TaskType)
	}
	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}
	if q.AGVCode != "" {
		query = query.Where("assigned_agv_code = ?", q.AGVCode)
	}
	if q.StartDate != "" {
		query = query.Where("created_at >= ?", q.StartDate)
	}
	if q.EndDate != "" {
		query = query.Where("created_at <= ?", q.EndDate+" 23:59:59")
	}

	// 统计
	query.Count(&total)

	// 分页
	page := q.Page
	if page < 1 {
		page = 1
	}
	pageSize := q.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	err := query.Order("priority DESC, created_at DESC").Offset(offset).Limit(pageSize).Find(&tasks).Error
	return tasks, total, err
}

// ListByStatus 按状态查询
func (r *AGVTaskRepository) ListByStatus(ctx context.Context, tenantID int64, status model.AGVTaskStatus) ([]model.AGVTask, error) {
	var tasks []model.AGVTask
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND status = ?", tenantID, status).
		Order("priority DESC, created_at ASC").Find(&tasks).Error
	return tasks, err
}

// ListByAGV 按AGV查询任务
func (r *AGVTaskRepository) ListByAGV(ctx context.Context, agvCode string, statuses []model.AGVTaskStatus) ([]model.AGVTask, error) {
	var tasks []model.AGVTask
	query := r.db.WithContext(ctx).Where("assigned_agv_code = ?", agvCode)
	if len(statuses) > 0 {
		query = query.Where("status IN ?", statuses)
	}
	err := query.Order("priority DESC, created_at ASC").Find(&tasks).Error
	return tasks, err
}

// GetPendingTasks 获取待处理任务
func (r *AGVTaskRepository) GetPendingTasks(ctx context.Context, tenantID int64, limit int) ([]model.AGVTask, error) {
	var tasks []model.AGVTask
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND status IN ?", tenantID, []model.AGVTaskStatus{model.AGVTaskStatusPending, model.AGVTaskStatusAssigned}).
		Order("priority DESC, created_at ASC").
		Limit(limit).
		Find(&tasks).Error
	return tasks, err
}

// AGVDeviceRepository AGV设备仓库
type AGVDeviceRepository struct {
	db *gorm.DB
}

func NewAGVDeviceRepository(db *gorm.DB) *AGVDeviceRepository {
	return &AGVDeviceRepository{db: db}
}

// Create 创建设备
func (r *AGVDeviceRepository) Create(ctx context.Context, device *model.AGVDevice) error {
	return r.db.WithContext(ctx).Create(device).Error
}

// GetByID 根据ID查询
func (r *AGVDeviceRepository) GetByID(ctx context.Context, id int64) (*model.AGVDevice, error) {
	var device model.AGVDevice
	err := r.db.WithContext(ctx).First(&device, id).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

// GetByCode 根据编号查询
func (r *AGVDeviceRepository) GetByCode(ctx context.Context, agvCode string) (*model.AGVDevice, error) {
	var device model.AGVDevice
	err := r.db.WithContext(ctx).Where("agv_code = ?", agvCode).First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

// Update 更新设备
func (r *AGVDeviceRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AGVDevice{}).Where("id = ?", id).Updates(updates).Error
}

// UpdateStatus 更新设备状态
func (r *AGVDeviceRepository) UpdateStatus(ctx context.Context, agvCode string, status model.AGVDeviceStatus) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == model.AGVDeviceStatusOnline || status == model.AGVDeviceStatusOffline {
		now := time.Now()
		updates["last_heartbeat"] = &now
	}
	return r.db.WithContext(ctx).Model(&model.AGVDevice{}).Where("agv_code = ?", agvCode).Updates(updates).Error
}

// List 查询设备列表
func (r *AGVDeviceRepository) List(ctx context.Context, q *model.AGVDeviceQuery) ([]model.AGVDevice, int64, error) {
	var devices []model.AGVDevice
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AGVDevice{})

	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.AGVCode != "" {
		query = query.Where("agv_code LIKE ?", "%"+q.AGVCode+"%")
	}
	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}

	query.Count(&total)

	page := q.Page
	if page < 1 {
		page = 1
	}
	pageSize := q.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	err := query.Order("agv_code ASC").Offset(offset).Limit(pageSize).Find(&devices).Error
	return devices, total, err
}

// ListAvailable 获取可用AGV列表
func (r *AGVDeviceRepository) ListAvailable(ctx context.Context, tenantID int64) ([]model.AGVDevice, error) {
	var devices []model.AGVDevice
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND status = ?", tenantID, model.AGVDeviceStatusOnline).
		Where("battery_level > 20").
		Order("battery_level DESC").Find(&devices).Error
	return devices, err
}

// UpdateHeartbeat 更新心跳
func (r *AGVDeviceRepository) UpdateHeartbeat(ctx context.Context, agvCode string, batteryLevel float64) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.AGVDevice{}).Where("agv_code = ?", agvCode).
		Updates(map[string]interface{}{
			"status":          model.AGVDeviceStatusOnline,
			"battery_level":    batteryLevel,
			"last_heartbeat":   &now,
		}).Error
}

// AGVLocationMappingRepository AGV库位映射仓库
type AGVLocationMappingRepository struct {
	db *gorm.DB
}

func NewAGVLocationMappingRepository(db *gorm.DB) *AGVLocationMappingRepository {
	return &AGVLocationMappingRepository{db: db}
}

// Create 创建映射
func (r *AGVLocationMappingRepository) Create(ctx context.Context, mapping *model.AGVLocationMapping) error {
	return r.db.WithContext(ctx).Create(mapping).Error
}

// GetByID 根据ID查询
func (r *AGVLocationMappingRepository) GetByID(ctx context.Context, id int64) (*model.AGVLocationMapping, error) {
	var mapping model.AGVLocationMapping
	err := r.db.WithContext(ctx).First(&mapping, id).Error
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

// GetByLocationCode 根据库位编码查询
func (r *AGVLocationMappingRepository) GetByLocationCode(ctx context.Context, locationCode string) (*model.AGVLocationMapping, error) {
	var mapping model.AGVLocationMapping
	err := r.db.WithContext(ctx).Where("location_code = ?", locationCode).First(&mapping).Error
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

// Update 更新映射
func (r *AGVLocationMappingRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AGVLocationMapping{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除映射
func (r *AGVLocationMappingRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.AGVLocationMapping{}, id).Error
}

// List 查询映射列表
func (r *AGVLocationMappingRepository) List(ctx context.Context, q *model.AGVLocationQuery) ([]model.AGVLocationMapping, int64, error) {
	var mappings []model.AGVLocationMapping
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AGVLocationMapping{})

	if q.TenantID > 0 {
		query = query.Where("tenant_id = ?", q.TenantID)
	}
	if q.LocationCode != "" {
		query = query.Where("location_code LIKE ?", "%"+q.LocationCode+"%")
	}
	if q.LocationType != "" {
		query = query.Where("location_type = ?", q.LocationType)
	}
	if q.Enabled != nil {
		query = query.Where("enabled = ?", *q.Enabled)
	}

	query.Count(&total)

	page := q.Page
	if page < 1 {
		page = 1
	}
	pageSize := q.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	err := query.Order("priority DESC, location_code ASC").Offset(offset).Limit(pageSize).Find(&mappings).Error
	return mappings, total, err
}

// ListByType 按类型查询
func (r *AGVLocationMappingRepository) ListByType(ctx context.Context, tenantID int64, locationType model.AGVLocationType) ([]model.AGVLocationMapping, error) {
	var mappings []model.AGVLocationMapping
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND location_type = ? AND enabled = ?", tenantID, locationType, true).
		Order("priority DESC, location_code ASC").Find(&mappings).Error
	return mappings, err
}

// GetByAGVLocationCode 根据AGV定位编码查询
func (r *AGVLocationMappingRepository) GetByAGVLocationCode(ctx context.Context, agvLocationCode string) (*model.AGVLocationMapping, error) {
	var mapping model.AGVLocationMapping
	err := r.db.WithContext(ctx).Where("agv_location_code = ?", agvLocationCode).First(&mapping).Error
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

// GenerateTaskNo 生成任务编号
func (r *AGVTaskRepository) GenerateTaskNo(ctx context.Context) (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("AGV-%s-", now.Format("20060102"))
	var count int64
	r.db.WithContext(ctx).Model(&model.AGVTask{}).Where("task_no LIKE ?", prefix+"%").Count(&count)
	return fmt.Sprintf("%s%04d", prefix, count+1), nil
}
