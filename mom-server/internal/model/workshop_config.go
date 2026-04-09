package model

import "time"

// ========== 多车间管理模块 ==========

// WorkshopConfig 车间配置
type WorkshopConfig struct {
	BaseModel
	TenantID           int64   `json:"tenant_id" gorm:"index;not null"`
	WorkshopID        int64   `json:"workshop_id" gorm:"uniqueIndex;not null"` // 关联车间
	ErpPlantCode      string  `json:"erp_plant_code" gorm:"size:50"`           // 金蝶工厂编码
	MaxDevices        int     `json:"max_devices" gorm:"default:0"`           // 最大设备数
	MaxWorkers        int     `json:"max_workers" gorm:"default:0"`           // 最大人员数
	MaxCapacityPerDay  float64 `json:"max_capacity_per_day" gorm:"type:decimal(18,2)"` // 日最大产能
	TimeZone          string  `json:"time_zone" gorm:"size:20;default:'Asia/Shanghai'"`
	IsDefault         int     `json:"is_default" gorm:"default:0"`             // 是否默认车间
}

func (WorkshopConfig) TableName() string {
	return "mdm_workshop_config"
}

// ShiftTemplate 班次模板
type ShiftTemplate struct {
	Name        string `json:"name"`        // 班次名称
	Start      string `json:"start"`      // 开始时间 HH:MM
	End        string `json:"end"`        // 结束时间 HH:MM
	BreakStart string `json:"break_start"` // 休息开始时间
	BreakEnd   string `json:"break_end"`  // 休息结束时间
}

// WorkingCalendar 工厂日历
type WorkingCalendar struct {
	BaseModel
	TenantID         int64          `json:"tenant_id" gorm:"index;not null"`
	WorkshopID      int64          `json:"workshop_id" gorm:"index;not null"`
	CalendarName    string         `json:"calendar_name" gorm:"size:100"`           // 日历名称
	WorkDays        string         `json:"work_days" gorm:"type:jsonb"`            // 工作日 [1,2,3,4,5] 周一到周五
	Shifts          string         `json:"shifts" gorm:"type:jsonb"`              // 班次模板
	HolidayDates    string         `json:"holiday_dates" gorm:"type:jsonb"`        // 节假日日期
	SpecialWorkDates string        `json:"special_work_dates" gorm:"type:jsonb"`  // 特殊工作日
	EffectiveFrom   *time.Time    `json:"effective_from"`                        // 生效日期
	EffectiveTo     *time.Time     `json:"effective_to"`                          // 失效日期
	Status          int            `json:"status" gorm:"default:1"`               // 1启用/0禁用
}

func (WorkingCalendar) TableName() string {
	return "aps_working_calendar"
}

// WorkHoursPerDay 每日可用工时
func (c *WorkingCalendar) GetWorkHoursPerDay() float64 {
	// 解析班次计算每日可用工时
	// 简化：默认8小时工作制
	return 8.0
}
