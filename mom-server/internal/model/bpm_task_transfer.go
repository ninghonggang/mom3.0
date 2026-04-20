package model

import "time"

// BpmTaskTransfer 任务转移记录
type BpmTaskTransfer struct {
	ID             uint64    `json:"id" gorm:"primaryKey"`
	TenantID       int64     `json:"tenant_id" gorm:"index"`
	TaskID         string    `json:"task_id" gorm:"size:64;index"`
	FromUserID     uint64    `json:"from_user_id" gorm:"index"`
	FromUserName   string    `json:"from_user_name" gorm:"size:100"`
	ToUserID       uint64    `json:"to_user_id" gorm:"index"`
	ToUserName     string    `json:"to_user_name" gorm:"size:100"`
	TransferReason string    `json:"transfer_reason" gorm:"size:500"`
	TransferTime   time.Time `json:"transfer_time"`
	OperatorID     uint64    `json:"operator_id"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (BpmTaskTransfer) TableName() string {
	return "bpm_task_transfer"
}

// TaskCandidateUser 任务候选人（查询结果DTO，不对应独立表，从TaskInstance派生）
type TaskCandidateUser struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	TaskID   string `json:"task_id"`
}

// TaskCandidateGroup 任务候选组（查询结果DTO）
type TaskCandidateGroup struct {
	GroupID   int64               `json:"group_id"`
	GroupName string              `json:"group_name"`
	Members   []TaskCandidateUser `json:"members"`
}
