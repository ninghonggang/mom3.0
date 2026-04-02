package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 基础模型
type BaseModel struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TenantModel 带租户的模型
type TenantModel struct {
	TenantID int64 `json:"tenant_id" gorm:"index;not null"`
}

// User 用户
type User struct {
	BaseModel
	TenantModel
	Username  string     `json:"username" gorm:"size:50;not null"`
	Nickname  string     `json:"nickname" gorm:"size:50"`
	Password  string     `json:"-" gorm:"size:200;not null"`
	Email     *string    `json:"email" gorm:"size:100"`
	Phone     *string    `json:"phone" gorm:"size:20"`
	Avatar    *string    `json:"avatar" gorm:"size:500"`
	DeptID    *int64     `json:"dept_id" gorm:"index"`
	Status    int        `json:"status" gorm:"default:1"` // 1正常 0停用
	LoginIP   *string    `json:"login_ip" gorm:"size:128"`
	LoginDate *time.Time `json:"login_date"`
}

func (User) TableName() string {
	return "sys_user"
}

// Role 角色
type Role struct {
	BaseModel
	TenantModel
	RoleName   string  `json:"role_name" gorm:"size:50;not null"`
	RoleKey    string  `json:"role_key" gorm:"size:100;not null"`
	RoleSort   int     `json:"role_sort" gorm:"default:0"`
	DataScope int     `json:"data_scope" gorm:"default:1"` // 1全部 2本部门 3本人
	Status    int     `json:"status" gorm:"default:1"`
	Remark    *string `json:"remark" gorm:"size:500"`
}

func (Role) TableName() string {
	return "sys_role"
}

// Menu 菜单
type Menu struct {
	BaseModel
	TenantModel
	ParentID   int64      `json:"parent_id" gorm:"default:0;index"`
	MenuName   string     `json:"menu_name" gorm:"size:50;not null"`
	MenuType   string     `json:"menu_type" gorm:"size:1"` // M目录 C菜单 F按钮 L链接
	Path       string     `json:"path" gorm:"size:200"`
	Component  *string    `json:"component" gorm:"size:200"`
	Perms      string     `json:"perms" gorm:"size:200"`
	Icon       *string    `json:"icon" gorm:"size:100"`
	Sort       int        `json:"sort" gorm:"default:0"`
	Visible    int        `json:"visible" gorm:"default:1"`
	Status     int        `json:"status" gorm:"default:1"`
	IsFrame    int        `json:"is_frame" gorm:"default:1"`
	IsCache    int        `json:"is_cache" gorm:"default:0"`
	Children   []Menu     `json:"children" gorm:"-"`
}

func (Menu) TableName() string {
	return "sys_menu"
}

// Dept 部门
type Dept struct {
	BaseModel
	TenantModel
	ParentID   int64   `json:"parent_id" gorm:"default:0;index"`
	DeptName   string  `json:"dept_name" gorm:"size:50;not null"`
	DeptCode   string  `json:"dept_code" gorm:"size:50"`
	DeptSort   int     `json:"dept_sort" gorm:"default:0"`
	Leader     *string `json:"leader" gorm:"size:50"`
	Phone      *string `json:"phone" gorm:"size:20"`
	Email      *string `json:"email" gorm:"size:100"`
	Status     int     `json:"status" gorm:"default:1"`
	Children   []Dept  `json:"children" gorm:"-"`
}

func (Dept) TableName() string {
	return "sys_dept"
}

// Post 岗位
type Post struct {
	BaseModel
	TenantModel
	PostCode   string  `json:"post_code" gorm:"size:50;not null"`
	PostName   string  `json:"post_name" gorm:"size:100;not null"`
	PostSort   int     `json:"post_sort" gorm:"default:0"`
	Status     int     `json:"status" gorm:"default:1"`
	Remark    *string `json:"remark" gorm:"size:500"`
}

func (Post) TableName() string {
	return "sys_post"
}

// DictType 字典类型
type DictType struct {
	BaseModel
	DictName  string  `json:"dict_name" gorm:"size:100;not null"`
	DictType  string  `json:"dict_type" gorm:"size:100;not null;uniqueIndex"`
	Status    int     `json:"status" gorm:"default:1"`
	Remark    *string `json:"remark" gorm:"size:500"`
}

func (DictType) TableName() string {
	return "sys_dict_type"
}

// DictData 字典数据
type DictData struct {
	BaseModel
	DictSort  int     `json:"dict_sort" gorm:"default:0"`
	DictLabel string   `json:"dict_label" gorm:"size:100;not null"`
	DictValue string   `json:"dict_value" gorm:"size:100;not null"`
	DictType  string   `json:"dict_type" gorm:"size:100;not null;index"`
	DictKey   string   `json:"dict_key" gorm:"size:100"`
	CssClass *string  `json:"css_class" gorm:"size:100"`
	ListClass *string `json:"list_class" gorm:"size:100"`
	IsDefault int     `json:"is_default" gorm:"default:0"`
	Status    int     `json:"status" gorm:"default:1"`
	Remark    *string `json:"remark" gorm:"size:500"`
}

func (DictData) TableName() string {
	return "sys_dict_data"
}

// Tenant 租户/工厂
type Tenant struct {
	BaseModel
	TenantName   string   `json:"tenant_name" gorm:"size:100;not null"`        // 工厂名称
	TenantKey    string   `json:"tenant_key" gorm:"size:50;not null;uniqueIndex"` // 工厂代码(唯一标识)

	// 地址信息
	Province     *string  `json:"province" gorm:"size:50"`    // 省份
	City        *string  `json:"city" gorm:"size:50"`        // 城市
	District    *string  `json:"district" gorm:"size:50"`    // 区县
	Address     *string  `json:"address" gorm:"size:255"`     // 详细地址

	// 联系方式
	Manager     *string  `json:"manager" gorm:"size:50"`      // 负责人/厂长
	ContactName *string  `json:"contact_name" gorm:"size:50"` // 联系人姓名
	ContactPhone *string `json:"contact_phone" gorm:"size:20"` // 联系电话
	ContactEmail *string `json:"contact_email" gorm:"size:100"` // 联系邮箱

	// 工厂属性
	FactoryType *string  `json:"factory_type" gorm:"size:50"` // 工厂类型(离散/流程/混合)
	EmployeeCount *int   `json:"employee_count"`             // 员工人数
	AreaSize    *float64 `json:"area_size"`                   // 占地面积(平方米)
	AnnualCapacity *float64 `json:"annual_capacity"`          // 年产能

	// 系统管理
	Status     int       `json:"status" gorm:"default:1"`     // 状态: 1正常 0禁用
	ExpireTime *time.Time `json:"expire_time"`               // 授权过期时间
	Remark     *string   `json:"remark" gorm:"type:text"`     // 备注说明
}

func (Tenant) TableName() string {
	return "sys_tenant"
}

// OperLog 操作日志
type OperLog struct {
	BaseModel
	TenantID    int64     `json:"tenant_id" gorm:"index"`
	Title       string    `json:"title" gorm:"size:200"`
	BusinessType *string `json:"business_type" gorm:"size:20"`
	Method      string    `json:"method" gorm:"size:100"`
	RequestMethod string  `json:"request_method" gorm:"size:10"`
	OperatorType *int    `json:"operator_type" gorm:"default:1"`
	OperName    *string  `json:"oper_name" gorm:"size:50"`
	DeptName    *string `json:"dept_name" gorm:"size:100"`
	OperURL     string    `json:"oper_url" gorm:"size:255"`
	OperIP      string    `json:"oper_ip" gorm:"size:50"`
	OperLocation *string `json:"oper_location" gorm:"size:255"`
	OperParam   *string  `json:"oper_param" gorm:"type:text"`
	JSONResult  *string  `json:"json_result" gorm:"type:text"`
	Status      int      `json:"status" gorm:"default:0"`
	ErrorMsg    *string `json:"error_msg" gorm:"type:text"`
	OperTime    time.Time `json:"oper_time"`
}

// LoginLog 登录日志
type LoginLog struct {
	BaseModel
	TenantID   int64     `json:"tenant_id" gorm:"index"`
	Username   string    `json:"username" gorm:"size:50"`
	IP        string    `json:"ip" gorm:"size:50"`
	LoginLocation *string `json:"login_location" gorm:"size:255"`
	Browser    string    `json:"browser" gorm:"size:50"`
	OS         string    `json:"os" gorm:"size:50"`
	Status     int       `json:"status" gorm:"default:0"`
	Msg        *string  `json:"msg" gorm:"size:255"`
	LoginTime  time.Time `json:"login_time"`
}

// RoleMenu 角色菜单关联
type RoleMenu struct {
	RoleID int64 `json:"role_id" gorm:"primaryKey"`
	MenuID int64 `json:"menu_id" gorm:"primaryKey"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}

// UserRole 用户角色关联
type UserRole struct {
	UserID int64 `json:"user_id" gorm:"primaryKey"`
	RoleID int64 `json:"role_id" gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}

// RolePerm 角色权限关联
type RolePerm struct {
	RoleID int64 `json:"role_id" gorm:"primaryKey"`
	Perm   string `json:"perm" gorm:"size:100;primaryKey"`
}

func (RolePerm) TableName() string {
	return "sys_role_perm"
}
