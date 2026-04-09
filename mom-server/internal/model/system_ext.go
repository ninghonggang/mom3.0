package model

import "time"

// PrintTemplate 打印模板
type PrintTemplate struct {
	BaseModel
	TenantID      int64   `json:"tenant_id" gorm:"index;not null"`
	TemplateCode  string  `json:"template_code" gorm:"size:50;not null"` // 模板编码
	TemplateName  string  `json:"template_name" gorm:"size:100;not null"` // 模板名称
	TemplateType  string  `json:"template_type" gorm:"size:30;not null"` // 模板类型: PRODUCTION_ORDER/PACKAGE/LABEL/REPORT
	EntityType    string  `json:"entity_type" gorm:"size:50"` // 适用单据类型
	Content       string  `json:"content" gorm:"type:text"` // 模板内容
	PaperType     string  `json:"paper_type" gorm:"size:20"` // 纸张类型
	PaperWidth    float64 `json:"paper_width"` // 纸张宽度(mm)
	PaperHeight   float64 `json:"paper_height"` // 纸张高度(mm)
	IsDefault     int     `json:"is_default" gorm:"default:0"` // 是否默认模板
	Status        int     `json:"status" gorm:"default:1"` // 1启用/0禁用
	Remark        string  `json:"remark" gorm:"size:500"` // 备注
}

func (PrintTemplate) TableName() string {
	return "sys_print_template"
}

// Notice 通知公告
type Notice struct {
	BaseModel
	TenantID       int64      `json:"tenant_id" gorm:"index;not null"`
	Title         string     `json:"title" gorm:"size:200;not null"` // 标题
	Content       string     `json:"content" gorm:"type:text"` // 内容
	NoticeType    string     `json:"notice_type" gorm:"size:20"` // 公告类型: SYSTEM/OPERATION/MAINTENANCE/OTHER
	Priority      int        `json:"priority" gorm:"default:1"` // 优先级: 1普通/2重要/3紧急
	PublishDept   string     `json:"publish_dept" gorm:"size:100"` // 发布部门
	PublisherID   int64      `json:"publisher_id"` // 发布人ID
	PublisherName string     `json:"publisher_name" gorm:"size:50"` // 发布人姓名
	PublishTime   *time.Time `json:"publish_time"` // 发布时间
	EffectTime    *time.Time `json:"effect_time"` // 生效时间
	ExpireTime    *time.Time `json:"expire_time"` // 失效时间
	TargetType    string     `json:"target_type" gorm:"size:20"` // 发布范围: ALL/DEPT/ROLE/USER
	TargetIds     string     `json:"target_ids" gorm:"size:500"` // 目标ID列表
	IsTop         int        `json:"is_top" gorm:"default:0"` // 是否置顶
	Status        int        `json:"status" gorm:"default:1"` // 1草稿/2已发布/3已撤回
	ViewCount     int        `json:"view_count" gorm:"default:0"` // 阅读次数
	Remark        string     `json:"remark" gorm:"size:500"` // 备注
}

func (Notice) TableName() string {
	return "sys_notice"
}

// NoticeReadRecord 公告阅读记录
type NoticeReadRecord struct {
	BaseModel
	TenantID  int64  `json:"tenant_id" gorm:"index;not null"`
	NoticeID  int64  `json:"notice_id" gorm:"index;not null"` // 公告ID
	UserID    int64  `json:"user_id" gorm:"index;not null"` // 用户ID
	UserName  string `json:"user_name" gorm:"size:50"` // 用户姓名
	ReadTime  string `json:"read_time" gorm:"size:30"` // 阅读时间
}

func (NoticeReadRecord) TableName() string {
	return "sys_notice_read_record"
}
