package model

import (
	"time"
)

// SkillLevel 技能等级
type SkillLevel string

const (
	SkillLevelJunior      SkillLevel = "JUNIOR"      // 初级
	SkillLevelIntermediate SkillLevel = "INTERMEDIATE" // 中级
	SkillLevelSenior      SkillLevel = "SENIOR"      // 高级
	SkillLevelMaster      SkillLevel = "MASTER"      // 技师/大师
)

// PersonSkillStatus 人员技能状态
type PersonSkillStatus string

const (
	PersonSkillStatusActive   PersonSkillStatus = "ACTIVE"   // 正常
	PersonSkillStatusInactive PersonSkillStatus = "INACTIVE" // 禁用
	PersonSkillStatusExpired  PersonSkillStatus = "EXPIRED"  // 已过期
)

// SkillType 技能类型
type SkillType string

const (
	SkillTypeAssembly   SkillType = "ASSEMBLY"   // 装配
	SkillTypeInspection SkillType = "INSPECTION" // 检验
	SkillTypePacking    SkillType = "PACKING"    // 包装
	SkillTypeOther      SkillType = "OTHER"      // 其他
)

// PersonSkill 人员技能表
type PersonSkill struct {
	BaseModel
	TenantID       int64            `json:"tenant_id" gorm:"index;not null"`
	PersonID      int64            `json:"person_id" gorm:"not null;index"`     // 人员ID
	PersonCode    string           `json:"person_code" gorm:"size:50;not null"`  // 工号
	PersonName    string           `json:"person_name" gorm:"size:100;not null"` // 姓名
	WorkshopID    int64            `json:"workshop_id" gorm:"not null;index"`  // 车间ID
	WorkstationID  *int64           `json:"workstation_id" gorm:"index"`        // 工位ID
	SkillLevel    SkillLevel       `json:"skill_level" gorm:"size:20;not null"` // 技能等级
	CertifiedDate time.Time        `json:"certified_date" gorm:"type:date"`      // 认证日期
	ExpiryDate    *time.Time       `json:"expiry_date" gorm:"type:date"`         // 有效期
	Status        PersonSkillStatus `json:"status" gorm:"size:20;default:ACTIVE"` // 状态
	CreatedBy     *int64           `json:"created_by"`
}

func (PersonSkill) TableName() string {
	return "mes_person_skill"
}

// PersonSkillScore 人员技能评分表
type PersonSkillScore struct {
	BaseModel
	TenantID      int64     `json:"tenant_id" gorm:"index;not null"`
	PersonSkillID int64     `json:"person_skill_id" gorm:"not null;index"` // 人员技能ID
	SkillType    SkillType `json:"skill_type" gorm:"size:20;not null"`     // 技能类型
	Score        float64   `json:"score" gorm:"type:decimal(5,2)"`          // 评分 0-100
	EvaluatedBy  *int64    `json:"evaluated_by"`                           // 评估人
	EvaluatorName *string  `json:"evaluator_name" gorm:"size:50"`          // 评估人姓名
	EvaluatedAt  time.Time `json:"evaluated_at" gorm:"type:timestamp"`     // 评估时间
	Remark       *string   `json:"remark" gorm:"size:500"`                  // 备注
}

func (PersonSkillScore) TableName() string {
	return "mes_person_skill_score"
}

// PersonSkillDetail 人员技能详情（包含评分列表）
type PersonSkillDetail struct {
	PersonSkill
	Scores []PersonSkillScore `json:"scores"`
}