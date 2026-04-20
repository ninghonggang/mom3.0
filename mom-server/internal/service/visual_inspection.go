package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// VisualInspectionService 视觉检测服务
type VisualInspectionService struct {
	repo     *repository.VisualInspectionRepository
	aiExecutor *AIExecutor
}

// NewVisualInspectionService 创建视觉检测服务
func NewVisualInspectionService(repo *repository.VisualInspectionRepository, aiExecutor *AIExecutor) *VisualInspectionService {
	return &VisualInspectionService{
		repo:     repo,
		aiExecutor: aiExecutor,
	}
}

// generateTaskNo 生成任务编号
func (s *VisualInspectionService) generateTaskNo() string {
	now := time.Now()
	return fmt.Sprintf("VI%s%s%04d", now.Format("20060102"), now.Format("150405"), rand.Intn(10000))
}

// CreateTask 创建检测任务
func (s *VisualInspectionService) CreateTask(ctx context.Context, req *model.CreateVisualInspectionTaskRequest, tenantID, userID int64) (*model.VisualInspectionTask, error) {
	// 生成任务编号
	taskNo := s.generateTaskNo()

	// 设置默认值
	priority := req.Priority
	if priority == "" {
		priority = model.PriorityNormal
	}

	now := time.Now()
	task := &model.VisualInspectionTask{
		TaskNo:            taskNo,
		TaskType:          req.TaskType,
		ProductID:         req.ProductID,
		ProductCode:       req.ProductCode,
		ProductName:       req.ProductName,
		ProductionOrderID: req.ProductionOrderID,
		WorkshopID:        req.WorkshopID,
		ImageURL:          req.ImageURL,
		ImageHash:         req.ImageHash,
		DetectionStandard: req.DetectionStandard,
		AIModelVersion:    req.AIModelVersion,
		Status:            model.InspectionStatusPending,
		Priority:          priority,
		RequestedBy:       req.RequestedBy,
		RequestedAt:       &now,
		TenantID:          tenantID,
		CreatedBy:         &userID,
		Remark:            req.Remark,
	}

	if err := s.repo.CreateTask(ctx, task); err != nil {
		return nil, fmt.Errorf("创建检测任务失败: %w", err)
	}

	// 异步触发AI检测（这里简化处理，实际应该发送到队列）
	go s.triggerInspection(task.ID, tenantID)

	return task, nil
}

// triggerInspection 触发AI检测
func (s *VisualInspectionService) triggerInspection(taskID uint, tenantID int64) {
	ctx := context.Background()

	// 更新状态为处理中
	s.repo.UpdateTaskStatus(ctx, taskID, model.InspectionStatusProcessing)

	// 获取任务信息
	task, err := s.repo.GetTaskByID(ctx, taskID)
	if err != nil {
		s.repo.UpdateTaskStatus(ctx, taskID, model.InspectionStatusFailed)
		return
	}

	// 模拟AI检测过程（实际应该调用AI模型服务）
	result := s.performAIInspection(task)

	// 保存检测结果
	if err := s.repo.CreateResult(ctx, result); err != nil {
		s.repo.UpdateTaskStatus(ctx, taskID, model.InspectionStatusFailed)
		return
	}

	// 更新任务状态
	s.repo.UpdateTaskStatus(ctx, taskID, model.InspectionStatusCompleted)
}

// performAIInspection 执行AI检测（模拟实现）
func (s *VisualInspectionService) performAIInspection(task *model.VisualInspectionTask) *model.VisualInspectionResult {
	// 模拟AI检测结果
	// 实际实现中应该调用AI模型服务进行真实检测
	now := time.Now()

	// 随机生成检测结果（简化模拟）
	result := &model.VisualInspectionResult{
		TaskID:        task.ID,
		DetectionTime: now,
		Confidence:    0.85 + rand.Float64()*0.15, // 0.85-1.0
		AIAnalysis:    fmt.Sprintf(`{"model":"%s","detected_objects":3,"processing_time":"120ms"}`, "v1.0.0"),
		TenantID:      task.TenantID,
	}

	// 根据任务类型设置不同结果
	switch task.TaskType {
	case model.TaskTypeDefectDetection:
		// 缺陷检测：90%通过率
		if rand.Float64() < 0.9 {
			result.Result = model.VIDetectionPass
			result.DefectType = nil
		} else {
			result.Result = model.VIDetectionFail
			defectTypes := []string{"scratch", "dent", "contamination", "misalignment"}
			result.DefectType = &defectTypes[rand.Intn(len(defectTypes))]
			result.DefectLocation = fmt.Sprintf(`{"x":%d,"y":%d,"width":%d,"height":%d}`, rand.Intn(100), rand.Intn(100), rand.Intn(50)+20, rand.Intn(50)+20)
		}
	case model.TaskTypeClassification:
		// 分类：95%通过率
		if rand.Float64() < 0.95 {
			result.Result = model.VIDetectionPass
		} else {
			result.Result = model.VIDetectionFail
			result.DefectType = viStringPtr("misclassification")
		}
	case model.TaskTypeMeasurement:
		// 测量：92%通过率
		if rand.Float64() < 0.92 {
			result.Result = model.VIDetectionPass
		} else {
			result.Result = model.VIDetectionFail
			result.DefectType = viStringPtr("out_of_tolerance")
		}
	case model.TaskTypeGhostDetection:
		// 鬼影检测：88%通过率
		if rand.Float64() < 0.88 {
			result.Result = model.VIDetectionPass
		} else {
			result.Result = model.VIDetectionFail
			result.DefectType = viStringPtr("ghost_image")
		}
	default:
		result.Result = model.VIDetectionPass
	}

	// 解析检测标准获取分析详情
	if task.DetectionStandard != "" {
		var standard map[string]interface{}
		if err := json.Unmarshal([]byte(task.DetectionStandard), &standard); err == nil {
			if analysis, ok := standard["analysis_prompt"]; ok {
				result.AIAnalysis = fmt.Sprintf(`{"model":"v1.0.0","prompt":"%v","detected_objects":3,"processing_time":"120ms"}`, analysis)
			}
		}
	}

	return result
}

// viStringPtr 字符串指针辅助函数
func viStringPtr(s string) *string {
	return &s
}

// GetTask 获取任务详情
func (s *VisualInspectionService) GetTask(ctx context.Context, id uint) (*model.VisualInspectionTask, error) {
	return s.repo.GetTaskByID(ctx, id)
}

// ListTasks 查询任务列表
func (s *VisualInspectionService) ListTasks(ctx context.Context, tenantID int64, taskType, status string, productID int64, page, pageSize int) ([]model.VisualInspectionTask, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return s.repo.ListTasks(ctx, tenantID, taskType, status, productID, page, pageSize)
}

// DeleteTask 删除任务
func (s *VisualInspectionService) DeleteTask(ctx context.Context, id uint) error {
	return s.repo.DeleteTask(ctx, id)
}

// GetResult 获取检测结果
func (s *VisualInspectionService) GetResult(ctx context.Context, taskID uint) (*model.VisualInspectionResult, error) {
	return s.repo.GetResultByTaskID(ctx, taskID)
}

// ManualReview 人工复核
func (s *VisualInspectionService) ManualReview(ctx context.Context, taskID uint, reviewerID int64, req *model.ManualReviewRequest) error {
	return s.repo.ManualReview(ctx, taskID, reviewerID, req.Result, req.Remark)
}

// GetStats 获取统计数据
func (s *VisualInspectionService) GetStats(ctx context.Context, tenantID int64) (*model.VisualInspectionStats, error) {
	return s.repo.GetStats(ctx, tenantID)
}
