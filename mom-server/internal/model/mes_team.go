package model

import (
	"time"
)

// MesTeam 班组
type MesTeam struct {
	BaseModel
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
	TeamCode    string  `json:"team_code" gorm:"size:50;not null;uniqueIndex:idx_tenant_team_code"`
	TeamName    string  `json:"team_name" gorm:"size:100;not null"`
	WorkshopID  int64   `json:"workshop_id" gorm:"not null"` // 车间ID
	ShiftID     *int64  `json:"shift_id"`                    // 班次ID
	LeaderID    *int64  `json:"leader_id"`                   // 组长ID
	Phone       *string `json:"phone" gorm:"size:20"`        // 联系电话
	Remark      *string `json:"remark" gorm:"size:500"`
	Status      int     `json:"status" gorm:"default:1"` // 1启用 0禁用
}

func (MesTeam) TableName() string {
	return "mes_team"
}

// MesTeamMember 班组成员
type MesTeamMember struct {
	BaseModel
	TenantID  int64  `json:"tenant_id" gorm:"index;not null"`
	TeamID    int64  `json:"team_id" gorm:"not null;index"`
	UserID    int64  `json:"user_id" gorm:"not null"`
	UserName  string `json:"user_name" gorm:"size:50"`
	Role      string `json:"role" gorm:"size:20"` // 角色：组长/组员
	JoinDate  *time.Time `json:"join_date"`
	Remark    *string `json:"remark" gorm:"size:500"`
	Status    int    `json:"status" gorm:"default:1"`
}

func (MesTeamMember) TableName() string {
	return "mes_team_member"
}

// MesTeamShift 班组排班
type MesTeamShift struct {
	BaseModel
	TenantID    int64     `json:"tenant_id" gorm:"index;not null"`
	TeamID      int64     `json:"team_id" gorm:"not null;index"`
	ShiftID     int64     `json:"shift_id" gorm:"not null"`
	ShiftDate   time.Time `json:"shift_date" gorm:"type:date;not null;index"` // 排班日期
	StartTime   string    `json:"start_time" gorm:"size:10"`                  // HH:mm
	EndTime     string    `json:"end_time" gorm:"size:10"`                    // HH:mm
	LeaderID    *int64    `json:"leader_id"`                                  // 当班组长
	Remark      *string   `json:"remark" gorm:"size:500"`
}

func (MesTeamShift) TableName() string {
	return "mes_team_shift"
}
