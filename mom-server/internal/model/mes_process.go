package model

import (
	"time"
)

// ========== MES工艺路线模块 ==========

// MesProcess 工艺路线主表
type MesProcess struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	ProcessCode   string     `json:"process_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_process_code"`
	ProcessName   string     `json:"process_name" gorm:"size:200;not null"`
	MaterialID    *int64     `json:"material_id"`    // 产品物料ID（可选，null表示通用工艺路线）
	MaterialCode  *string    `json:"material_code" gorm:"size:50"`
	MaterialName  *string    `json:"material_name" gorm:"size:100"`
	Version       string     `json:"version" gorm:"size:20"` // 版本号
	Status       string     `json:"status" gorm:"size:20;default:'DRAFT'"` // DRAFT草稿/ACTIVE生效/EXPIRED失效
	EffDate      *time.Time `json:"eff_date" gorm:"type:date"` // 生效日期
	ExpDate      *time.Time `json:"exp_date" gorm:"type:date"` // 失效日期
	Remark       *string    `json:"remark" gorm:"size:500"`
	IsCurrent    int        `json:"is_current" gorm:"default:1"` // 是否当前版本 0否 1是

	// 关联的工序明细
	Operations    []MesProcessOperation `json:"operations" gorm:"foreignKey:ProcessID"`
}

func (MesProcess) TableName() string {
	return "mes_process"
}

// MesProcessOperation 工艺路线工序表
type MesProcessOperation struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	ProcessID      int64   `json:"process_id" gorm:"index;not null"`
	OperationID    int64   `json:"operation_id"` // 关联的工序ID（来自mdm_operation）
	OperationCode  string  `json:"operation_code" gorm:"size:50"`
	OperationName  string  `json:"operation_name" gorm:"size:100"`
	LineNo         int     `json:"line_no" gorm:"not null"` // 工序顺序号
	StandardWorktime int   `json:"standard_worktime" gorm:"default:0"` // 标准工时(分钟)
	WorkcenterID   *int64  `json:"workcenter_id"`
	WorkcenterName *string `json:"workcenter_name" gorm:"size:100"`
	// 能力要求
	RequiredCapacity float64 `json:"required_capacity" gorm:"type:decimal(10,2);default:0"` // 所需产能(小时)
	MinWorkers       int     `json:"min_workers" gorm:"default:1"` // 最少人数
	MaxWorkers       int     `json:"max_workers" gorm:"default:10"` // 最多人数
	// 品质要求
	IsKeyProcess    int     `json:"is_key_process" gorm:"default:0"` // 是否关键工序
	IsQCPoint       int     `json:"is_qc_point" gorm:"default:0"` // 是否质检点
	QualityStd      *string `json:"quality_std" gorm:"size:500"` // 质量标准
	// 工序状态
	Status          string  `json:"status" gorm:"size:20;default:'ACTIVE'"` // ACTIVE/INACTIVE
	Remark          *string `json:"remark" gorm:"size:500"`
}

func (MesProcessOperation) TableName() string {
	return "mes_process_operation"
}

// ========== 请求/响应结构 ==========

// MesProcessCreate 工艺路线创建请求
type MesProcessCreate struct {
	ProcessCode  string                    `json:"process_code"`
	ProcessName  string                    `json:"process_name"`
	MaterialID   *int64                    `json:"material_id"`
	MaterialCode string                    `json:"material_code"`
	MaterialName string                    `json:"material_name"`
	Version      string                    `json:"version"`
	Status       string                    `json:"status"`
	EffDate      string                    `json:"eff_date"` // YYYY-MM-DD
	ExpDate      string                    `json:"exp_date"` // YYYY-MM-DD
	Remark       string                    `json:"remark"`
	Operations   []MesProcessOperationCreate `json:"operations"`
}

// MesProcessOperationCreate 工序创建请求
type MesProcessOperationCreate struct {
	OperationID      int64   `json:"operation_id"`
	OperationCode    string  `json:"operation_code"`
	OperationName    string  `json:"operation_name"`
	LineNo           int     `json:"line_no"`
	StandardWorktime int     `json:"standard_worktime"`
	WorkcenterID     *int64  `json:"workcenter_id"`
	WorkcenterName   string  `json:"workcenter_name"`
	RequiredCapacity float64 `json:"required_capacity"`
	MinWorkers       int     `json:"min_workers"`
	MaxWorkers       int     `json:"max_workers"`
	IsKeyProcess     int     `json:"is_key_process"`
	IsQCPoint        int     `json:"is_qc_point"`
	QualityStd       string  `json:"quality_std"`
	Remark           string  `json:"remark"`
}

// MesProcessUpdate 工艺路线更新请求
type MesProcessUpdate struct {
	ProcessName  string                      `json:"process_name"`
	Status       string                      `json:"status"`
	EffDate      string                      `json:"eff_date"`
	ExpDate      string                      `json:"exp_date"`
	Remark       string                      `json:"remark"`
	Operations   []MesProcessOperationCreate `json:"operations"`
}

// MesProcessWithOperations 工艺路线及工序完整结构
type MesProcessWithOperations struct {
	MesProcess
	Operations []MesProcessOperation `json:"operations"`
}
