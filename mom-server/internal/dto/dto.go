package dto

// 统一请求/响应 DTO

// PageRequest 分页请求
type PageRequest struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

func (r *PageRequest) GetPage() int {
	if r.Page <= 0 {
		return 1
	}
	return r.Page
}

func (r *PageRequest) GetPageSize() int {
	if r.PageSize <= 0 || r.PageSize > 100 {
		return 20
	}
	return r.PageSize
}

func (r *PageRequest) GetOffset() int {
	return (r.GetPage() - 1) * r.GetPageSize()
}

// PageData 分页数据
type PageData struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Captcha  string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	User         *UserDTO `json:"user"`
}

// RefreshTokenRequest 刷新Token请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserDTO 用户DTO
type UserDTO struct {
	ID        int64   `json:"id"`
	Username  string  `json:"username"`
	Nickname  string  `json:"nickname"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Avatar    *string `json:"avatar"`
	DeptID    *int64  `json:"dept_id"`
	Status    int     `json:"status"`
	Roles     []string `json:"roles"`
	Perms     []string `json:"perms"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string  `json:"username" binding:"required"`
	Nickname string  `json:"nickname"`
	Password string  `json:"password" binding:"required"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	DeptID   *int64  `json:"dept_id"`
	Status   int     `json:"status"`
	RoleIDs  []int64 `json:"role_ids"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Nickname string  `json:"nickname"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	DeptID   *int64  `json:"dept_id"`
	Status   int     `json:"status"`
	RoleIDs  []int64 `json:"role_ids"`
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// CreateRoleRequest 创建角色请求
type CreateRoleRequest struct {
	RoleName   string  `json:"role_name" binding:"required"`
	RoleKey    string  `json:"role_key" binding:"required"`
	RoleSort   int     `json:"role_sort"`
	DataScope int     `json:"data_scope"`
	Status    int     `json:"status"`
	Remark    *string `json:"remark"`
	MenuIDs   []int64 `json:"menu_ids"`
}

// UpdateRoleRequest 更新角色请求
type UpdateRoleRequest struct {
	RoleName   string  `json:"role_name"`
	RoleSort   int     `json:"role_sort"`
	DataScope int     `json:"data_scope"`
	Status    int     `json:"status"`
	Remark    *string `json:"remark"`
	MenuIDs   []int64 `json:"menu_ids"`
}

// CreateMenuRequest 创建菜单请求
type CreateMenuRequest struct {
	ParentID   int64   `json:"parent_id"`
	MenuName   string  `json:"menu_name" binding:"required"`
	MenuType   string  `json:"menu_type" binding:"required"`
	Path       string  `json:"path"`
	Component  *string `json:"component"`
	Perms      string  `json:"perms"`
	Icon       *string `json:"icon"`
	MenuSort   int     `json:"menu_sort"`
	Visible    int     `json:"visible"`
	Status     int     `json:"status"`
	IsFrame    int     `json:"is_frame"`
	IsCache    int     `json:"is_cache"`
}

// UpdateMenuRequest 更新菜单请求
type UpdateMenuRequest struct {
	ParentID   int64   `json:"parent_id"`
	MenuName   string  `json:"menu_name"`
	MenuType   string  `json:"menu_type"`
	Path       string  `json:"path"`
	Component  *string `json:"component"`
	Perms      string  `json:"perms"`
	Icon       *string `json:"icon"`
	MenuSort   int     `json:"menu_sort"`
	Visible    int     `json:"visible"`
	Status     int     `json:"status"`
	IsFrame    int     `json:"is_frame"`
	IsCache    int     `json:"is_cache"`
}

// CreateDeptRequest 创建部门请求
type CreateDeptRequest struct {
	ParentID  int64   `json:"parent_id"`
	DeptName  string  `json:"dept_name" binding:"required"`
	DeptSort  int     `json:"dept_sort"`
	Leader    *string `json:"leader"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Status    int     `json:"status"`
}

// UpdateDeptRequest 更新部门请求
type UpdateDeptRequest struct {
	ParentID  int64   `json:"parent_id"`
	DeptName  string  `json:"dept_name"`
	DeptSort  int     `json:"dept_sort"`
	Leader    *string `json:"leader"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Status    int     `json:"status"`
}

// CreateTenantRequest 创建租户请求
type CreateTenantRequest struct {
	TenantName string     `json:"tenant_name" binding:"required"`
	TenantKey  string     `json:"tenant_key" binding:"required"`
	Contact    *string    `json:"contact"`
	Phone      *string    `json:"phone"`
	Email      *string    `json:"email"`
	Status     int        `json:"status"`
	ExpireTime *string    `json:"expire_time"`
}

// UpdateTenantRequest 更新租户请求
type UpdateTenantRequest struct {
	TenantName string     `json:"tenant_name"`
	Contact    *string    `json:"contact"`
	Phone      *string    `json:"phone"`
	Email      *string    `json:"email"`
	Status     int        `json:"status"`
	ExpireTime *string    `json:"expire_time"`
}

// DictTypeRequest 字典类型请求
type DictTypeRequest struct {
	DictName string  `json:"dict_name" binding:"required"`
	DictType string  `json:"dict_type" binding:"required"`
	Status   int     `json:"status"`
	Remark   *string `json:"remark"`
}

// DictDataRequest 字典数据请求
type DictDataRequest struct {
	DictSort  int     `json:"dict_sort"`
	DictLabel string  `json:"dict_label" binding:"required"`
	DictValue string  `json:"dict_value" binding:"required"`
	DictType  string  `json:"dict_type" binding:"required"`
	DictKey   string  `json:"dict_key"`
	CssClass *string `json:"css_class"`
	ListClass *string `json:"list_class"`
	IsDefault int     `json:"is_default"`
	Status    int     `json:"status"`
	Remark    *string `json:"remark"`
}

// RoleListReq 角色列表请求
type RoleListReq struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	RoleName string `form:"role_name"`
}

// MaterialListReq 物料列表请求
type MaterialListReq struct {
	Page          int    `form:"page"`
	PageSize      int    `form:"page_size"`
	MaterialCode  string `form:"material_code"`
	MaterialName  string `form:"material_name"`
}

// ProductionOrderListReq 生产工单列表请求
type ProductionOrderListReq struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	OrderNo  string `form:"order_no"`
	Status   int    `form:"status"`
}

// EquipmentListReq 设备列表请求
type EquipmentListReq struct {
	Page           int    `form:"page"`
	PageSize       int    `form:"page_size"`
	EquipmentCode  string `form:"equipment_code"`
	EquipmentName  string `form:"equipment_name"`
}

// WarehouseListReq 仓库列表请求
type WarehouseListReq struct {
	Page          int    `form:"page"`
	PageSize      int    `form:"page_size"`
	WarehouseCode string `form:"warehouse_code"`
	WarehouseName string `form:"warehouse_name"`
}

// WorkshopListReq 车间列表请求
type WorkshopListReq struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

// PostListReq 岗位列表请求
type PostListReq struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	PostName string `form:"post_name"`
}
