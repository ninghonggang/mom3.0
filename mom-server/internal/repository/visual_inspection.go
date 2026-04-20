package repository

import (
	"context"
	"mom-server/internal/model"
	"time"

	"gorm.io/gorm"
)

// VisualInspectionRepository 视觉检测仓储
type VisualInspectionRepository struct {
	db *gorm.DB
}

// NewVisualInspectionRepository 创建视觉检测仓储
func NewVisualInspectionRepository(db *gorm.DB) *VisualInspectionRepository {
	return &VisualInspectionRepository{db: db}
}

// CreateTask 创建检测任务
func (r *VisualInspectionRepository) CreateTask(ctx context.Context, task *model.VisualInspectionTask) error {
	return r.db.WithContext(ctx).Create(task).Error
}

// UpdateTask 更新任务
func (r *VisualInspectionRepository) UpdateTask(ctx context.Context, task *model.VisualInspectionTask) error {
	return r.db.WithContext(ctx).Save(task).Error
}

// GetTaskByID 根据ID获取任务
func (r *VisualInspectionRepository) GetTaskByID(ctx context.Context, id uint) (*model.VisualInspectionTask, error) {
	var task model.VisualInspectionTask
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetTaskByTaskNo 根据任务编号获取任务
func (r *VisualInspectionRepository) GetTaskByTaskNo(ctx context.Context, taskNo string) (*model.VisualInspectionTask, error) {
	var task model.VisualInspectionTask
	err := r.db.WithContext(ctx).Where("task_no = ?", taskNo).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// ListTasks 查询任务列表（分页）
func (r *VisualInspectionRepository) ListTasks(ctx context.Context, tenantID int64, taskType string, status string, productID int64, page, pageSize int) ([]model.VisualInspectionTask, int64, error) {
	var tasks []model.VisualInspectionTask
	var total int64

	query := r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("tenant_id = ?", tenantID)

	if taskType != "" {
		query = query.Where("task_type = ?", taskType)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if productID > 0 {
		query = query.Where("product_id = ?", productID)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

// DeleteTask 删除任务
func (r *VisualInspectionRepository) DeleteTask(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.VisualInspectionTask{}, id).Error
}

// UpdateTaskStatus 更新任务状态
func (r *VisualInspectionRepository) UpdateTaskStatus(ctx context.Context, id uint, status model.VisualInspectionStatus) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == model.InspectionStatusCompleted || status == model.InspectionStatusFailed {
		now := time.Now()
		updates["completed_at"] = &now
	}
	return r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("id = ?", id).Updates(updates).Error
}

// CreateResult 创建检测结果
func (r *VisualInspectionRepository) CreateResult(ctx context.Context, result *model.VisualInspectionResult) error {
	return r.db.WithContext(ctx).Create(result).Error
}

// UpdateResult 更新检测结果
func (r *VisualInspectionRepository) UpdateResult(ctx context.Context, result *model.VisualInspectionResult) error {
	return r.db.WithContext(ctx).Save(result).Error
}

// GetResultByTaskID 根据任务ID获取检测结果
func (r *VisualInspectionRepository) GetResultByTaskID(ctx context.Context, taskID uint) (*model.VisualInspectionResult, error) {
	var result model.VisualInspectionResult
	err := r.db.WithContext(ctx).Where("task_id = ?", taskID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ManualReview 人工复核
func (r *VisualInspectionRepository) ManualReview(ctx context.Context, taskID uint, reviewerID int64, reviewResult model.ManualReviewResult, remark *string) error {
	now := time.Now()
	updates := map[string]interface{}{
		"manual_review_result": reviewResult,
		"manual_review_by":     reviewerID,
		"manual_review_at":     &now,
	}
	if remark != nil {
		updates["remark"] = *remark
	}

	// 更新结果表
	if err := r.db.WithContext(ctx).Model(&model.VisualInspectionResult{}).Where("task_id = ?", taskID).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

// GetStats 获取统计数据
func (r *VisualInspectionRepository) GetStats(ctx context.Context, tenantID int64) (*model.VisualInspectionStats, error) {
	stats := &model.VisualInspectionStats{}

	// 总任务数
	r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("tenant_id = ?", tenantID).Count(&stats.TotalTasks)

	// 待处理
	r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("tenant_id = ? AND status = ?", tenantID, model.InspectionStatusPending).Count(&stats.PendingTasks)

	// 处理中
	r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("tenant_id = ? AND status = ?", tenantID, model.InspectionStatusProcessing).Count(&stats.ProcessingTasks)

	// 已完成
	r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("tenant_id = ? AND status = ?", tenantID, model.InspectionStatusCompleted).Count(&stats.CompletedTasks)

	// 失败
	r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("tenant_id = ? AND status = ?", tenantID, model.InspectionStatusFailed).Count(&stats.FailedTasks)

	// 今日任务数
	today := time.Now().Truncate(24 * time.Hour)
	r.db.WithContext(ctx).Model(&model.VisualInspectionTask{}).Where("tenant_id = ? AND created_at >= ?", tenantID, today).Count(&stats.TodayTasks)

	// 今日检出缺陷数（结果为FAIL的）
	var defectCount int64
	r.db.WithContext(ctx).Model(&model.VisualInspectionResult{}).
		Joins("JOIN visual_inspection_task ON visual_inspection_result.task_id = visual_inspection_task.id").
		Where("visual_inspection_task.tenant_id = ? AND visual_inspection_result.detection_time >= ? AND visual_inspection_result.result = ?", tenantID, today, model.VIDetectionFail).
		Count(&defectCount)
	stats.TodayDefects = defectCount

	// 通过率（已完成的检测中PASS的比例）
	var totalCompleted, passCount int64
	if err := r.db.WithContext(ctx).Model(&model.VisualInspectionResult{}).
		Joins("JOIN visual_inspection_task ON visual_inspection_result.task_id = visual_inspection_task.id").
		Where("visual_inspection_task.tenant_id = ?", tenantID).
		Count(&totalCompleted).Error; err != nil {
		return stats, err
	}
	if totalCompleted > 0 {
		r.db.WithContext(ctx).Model(&model.VisualInspectionResult{}).
			Joins("JOIN visual_inspection_task ON visual_inspection_result.task_id = visual_inspection_task.id").
			Where("visual_inspection_task.tenant_id = ? AND visual_inspection_result.result = ?", tenantID, model.VIDetectionPass).
			Count(&passCount)
		stats.PassRate = float64(passCount) / float64(totalCompleted) * 100
	}

	// 平均置信度
	var avgConf float64
	r.db.WithContext(ctx).Model(&model.VisualInspectionResult{}).
		Select("COALESCE(AVG(confidence), 0)").
		Joins("JOIN visual_inspection_task ON visual_inspection_result.task_id = visual_inspection_task.id").
		Where("visual_inspection_task.tenant_id = ?", tenantID).
		Scan(&avgConf)
	stats.AvgConfidence = avgConf

	return stats, nil
}
