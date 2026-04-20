package model

import (
	"time"
)

// VisualInspectionTaskType 视觉检测任务类型
type VisualInspectionTaskType string

const (
	TaskTypeDefectDetection VisualInspectionTaskType = "DEFECT_DETECTION" // 缺陷检测
	TaskTypeClassification  VisualInspectionTaskType = "CLASSIFICATION"    // 分类
	TaskTypeMeasurement     VisualInspectionTaskType = "MEASUREMENT"       // 测量
	TaskTypeGhostDetection  VisualInspectionTaskType = "GHOST_DETECTION"  // 鬼影检测
)

// VisualInspectionStatus 视觉检测状态
type VisualInspectionStatus string

const (
	InspectionStatusPending    VisualInspectionStatus = "PENDING"    // 待处理
	InspectionStatusProcessing VisualInspectionStatus = "PROCESSING" // 处理中
	InspectionStatusCompleted  VisualInspectionStatus = "COMPLETED"  // 已完成
	InspectionStatusFailed     VisualInspectionStatus = "FAILED"     // 失败
)

// VisualInspectionPriority 优先级
type VisualInspectionPriority string

const (
	PriorityLow    VisualInspectionPriority = "LOW"
	PriorityNormal VisualInspectionPriority = "NORMAL"
	PriorityHigh   VisualInspectionPriority = "HIGH"
	PriorityUrgent VisualInspectionPriority = "URGENT"
)

// VIDetectionResult 视觉检测结果
type VIDetectionResult string

const (
	VIDetectionPass VIDetectionResult = "PASS" // 通过
	VIDetectionFail VIDetectionResult = "FAIL" // 失败
)

// ManualReviewResult 人工复核结果
type ManualReviewResult string

const (
	ManualReviewConfirmed ManualReviewResult = "CONFIRMED" // 确认
	ManualReviewRejected  ManualReviewResult = "REJECTED"  // 驳回
)

// VisualInspectionTask 视觉检测任务表
type VisualInspectionTask struct {
	ID                 uint                      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
	TaskNo             string                    `json:"task_no" gorm:"size:50;uniqueIndex;not null"` // 任务编号
	TaskType           VisualInspectionTaskType  `json:"task_type" gorm:"size:30;not null"`            // 任务类型
	ProductID          int64                     `json:"product_id"`                                      // 产品ID
	ProductCode        string                    `json:"product_code" gorm:"size:50"`                    // 产品编码
	ProductName        string                    `json:"product_name" gorm:"size:100"`                  // 产品名称
	ProductionOrderID  *int64                    `json:"production_order_id"`                             // 生产工单ID
	WorkshopID         *int64                    `json:"workshop_id"`                                     // 车间ID
	ImageURL           string                    `json:"image_url" gorm:"size:500"`                      // 检测图片URL
	ImageHash          *string                   `json:"image_hash" gorm:"size:64"`                      // 图片hash
	DetectionStandard  string                    `json:"detection_standard" gorm:"type:json"`           // 检测标准JSON
	AIModelVersion     *string                   `json:"ai_model_version" gorm:"size:50"`                // AI模型版本
	Status             VisualInspectionStatus    `json:"status" gorm:"size:20;default:PENDING"`          // 状态
	Priority           VisualInspectionPriority  `json:"priority" gorm:"size:10;default:NORMAL"`        // 优先级
	RequestedBy        string                    `json:"requested_by" gorm:"size:50"`                   // 请求人
	RequestedAt        *time.Time                `json:"requested_at"`                                   // 请求时间
	CompletedAt        *time.Time                `json:"completed_at"`                                   // 完成时间
	Remark             *string                   `json:"remark" gorm:"size:500"`                         // 备注
	TenantID           int64                     `json:"tenant_id" gorm:"index;not null"`                // 租户ID
	CreatedBy          *int64                    `json:"created_by"`                                     // 创建人
}

func (VisualInspectionTask) TableName() string {
	return "visual_inspection_task"
}

// VisualInspectionResult 视觉检测结果表
type VisualInspectionResult struct {
	ID                uint               `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	TaskID            uint               `json:"task_id" gorm:"index;not null"`                         // 任务ID
	DetectionTime     time.Time          `json:"detection_time"`                                        // 检测时间
	Result            VIDetectionResult `json:"result" gorm:"size:10;not null"`                        // 检测结果: PASS/FAIL
	Confidence        float64            `json:"confidence" gorm:"type:decimal(5,4)"`                   // 置信度 0-1
	DefectType        *string            `json:"defect_type" gorm:"size:50"`                            // 缺陷类型
	DefectLocation    string             `json:"defect_location" gorm:"type:json"`                      // 缺陷位置JSON: {x,y,width,height}
	DefectImageURL    *string            `json:"defect_image_url" gorm:"size:500"`                      // 缺陷图片URL
	AIAnalysis         string             `json:"ai_analysis" gorm:"type:json"`                        // AI分析结果JSON
	ManualReviewResult *ManualReviewResult `json:"manual_review_result" gorm:"size:20"`                  // 人工复核结果: CONFIRMED/REJECTED
	ManualReviewBy     *int64             `json:"manual_review_by"`                                     // 复核人
	ManualReviewAt     *time.Time         `json:"manual_review_at"`                                     // 复核时间
	Remark            *string            `json:"remark" gorm:"size:500"`                               // 备注
	TenantID          int64              `json:"tenant_id" gorm:"index;not null"`                       // 租户ID
}

func (VisualInspectionResult) TableName() string {
	return "visual_inspection_result"
}

// CreateVisualInspectionTaskRequest 创建视觉检测任务请求
type CreateVisualInspectionTaskRequest struct {
	TaskType          VisualInspectionTaskType `json:"task_type" binding:"required"`
	ProductID         int64                    `json:"product_id"`
	ProductCode       string                   `json:"product_code"`
	ProductName       string                   `json:"product_name"`
	ProductionOrderID *int64                   `json:"production_order_id"`
	WorkshopID        *int64                   `json:"workshop_id"`
	ImageURL          string                   `json:"image_url" binding:"required"`
	ImageHash         *string                  `json:"image_hash"`
	DetectionStandard string                   `json:"detection_standard"`
	AIModelVersion    *string                  `json:"ai_model_version"`
	Priority          VisualInspectionPriority `json:"priority"`
	RequestedBy       string                   `json:"requested_by"`
	Remark            *string                 `json:"remark"`
}

// ManualReviewRequest 人工复核请求
type ManualReviewRequest struct {
	Result ManualReviewResult `json:"result" binding:"required"` // CONFIRMED/REJECTED
	Remark *string            `json:"remark"`
}

// VisualInspectionStats 视觉检测统计数据
type VisualInspectionStats struct {
	TotalTasks     int64   `json:"total_tasks"`      // 总任务数
	PendingTasks   int64   `json:"pending_tasks"`    // 待处理
	ProcessingTasks int64   `json:"processing_tasks"` // 处理中
	CompletedTasks int64   `json:"completed_tasks"`  // 已完成
	FailedTasks    int64   `json:"failed_tasks"`     // 失败
	PassRate       float64 `json:"pass_rate"`        // 通过率
	AvgConfidence  float64 `json:"avg_confidence"`   // 平均置信度
	TodayTasks     int64   `json:"today_tasks"`      // 今日任务数
	TodayDefects   int64   `json:"today_defects"`    // 今日检出缺陷数
}
