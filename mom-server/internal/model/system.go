package model

import (
	"encoding/json"
	"time"
)

// ImportTask 导入任务
type ImportTask struct {
	ID            int64           `json:"id" gorm:"primaryKey"`
	TenantID      int64           `json:"tenant_id" gorm:"index;default:1"`
	TaskNo        string          `json:"task_no" gorm:"size:50;uniqueIndex;not null"`
	ImportType    string          `json:"import_type" gorm:"size:20;not null"`
	FileName      string          `json:"file_name" gorm:"size:200"`
	FilePath      string          `json:"file_path" gorm:"size:500"`
	TotalRows     int             `json:"total_rows" gorm:"default:0"`
	SuccessRows   int             `json:"success_rows" gorm:"default:0"`
	FailRows      int             `json:"fail_rows" gorm:"default:0"`
	FailDataJSON  json.RawMessage `json:"fail_data_json" gorm:"type:jsonb"`
	Status        string          `json:"status" gorm:"size:20;default:PENDING"`
	CreatedBy     string          `json:"created_by" gorm:"size:50"`
	CreatedAt     time.Time       `json:"created_at" gorm:"autoCreateTime"`
	CompletedAt   *time.Time      `json:"completed_at"`
}

func (ImportTask) TableName() string {
	return "sys_import_task"
}

// ImportStatus 导入状态常量
const (
	ImportStatusPending   = "PENDING"
	ImportStatusProcessing = "PROCESSING"
	ImportStatusSuccess    = "SUCCESS"
	ImportStatusFail       = "FAIL"
)

// ImportType 导入类型常量
const (
	ImportTypeMaterial = "MATERIAL"
	ImportTypeBom       = "BOM"
)
