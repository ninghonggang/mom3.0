package errno

// 错误码规范
// 格式：AABBCC
// AA: 模块编号
//   - 01: 系统通用错误
//   - 02: 认证授权
//   - 03: 用户管理
//   - 04: 角色管理
//   - 05: 菜单管理
//   - 06: 部门管理
//   - 07: 字典管理
//   - 08: 生产执行
//   - 09: APS计划
//   - 10: 仓储管理
//   - 11: 质量管理
//   - 12: 设备管理
//   - 13: 追溯管理
// BB: 错误类型
//   - 00: 通用错误
//   - 01: 参数错误
//   - 02: 未找到
//   - 03: 权限不足
//   - 04: 重复/冲突
//   - 05: 状态异常
// CC: 具体错误序号（01-99）

const (
	// 系统通用错误 (01)
	Success           = 0     // 成功
	ErrInternalServer = 10100 // 服务器内部错误
	ErrParamInvalid   = 10101 // 参数错误
	ErrUnauthorized   = 10201 // 未认证
	ErrForbidden      = 10203 // 权限不足
	ErrNotFound       = 10202 // 资源不存在
	ErrTooManyRequest = 10301 // 请求过于频繁

	// 认证授权 (02)
	ErrAuthInvalidUsername = 20101 // 用户名或密码错误
	ErrAuthUserDisabled    = 20105 // 用户已被禁用
	ErrAuthTokenInvalid    = 20201 // Token无效
	ErrAuthTokenExpired    = 20202 // Token已过期
	ErrAuthTokenMalformed  = 20203 // Token格式错误

	// 用户管理 (03)
	ErrUserNotFound = 30202 // 用户不存在
	ErrUserExists   = 30401 // 用户名已存在
	ErrUserCreate   = 30100 // 创建用户失败
	ErrUserUpdate   = 30101 // 更新用户失败
	ErrUserDelete   = 30102 // 删除用户失败

	// 角色管理 (04)
	ErrRoleNotFound = 40202 // 角色不存在
	ErrRoleExists   = 40401 // 角色已存在
	ErrRoleAssign   = 40101 // 角色分配失败

	// 菜单管理 (05)
	ErrMenuNotFound = 50202 // 菜单不存在

	// 生产执行 (08)
	ErrSalesOrderNotFound = 80202 // 销售订单不存在

	// APS计划 (09)
	ErrMPSNotFound = 90202 // MPS计划不存在
	ErrMRPNotFound = 91202 // MRP计划不存在
)

// 错误消息映射
var errorMessages = map[int]string{
	Success:                "操作成功",
	ErrInternalServer:      "服务器内部错误",
	ErrParamInvalid:        "参数错误",
	ErrUnauthorized:        "未认证或认证失败",
	ErrForbidden:           "权限不足",
	ErrNotFound:            "资源不存在",
	ErrTooManyRequest:      "请求过于频繁，请稍后再试",
	ErrAuthInvalidUsername: "用户名或密码错误",
	ErrAuthUserDisabled:    "用户已被禁用",
	ErrAuthTokenInvalid:    "Token无效",
	ErrAuthTokenExpired:    "Token已过期",
	ErrAuthTokenMalformed:  "Token格式错误",
	ErrUserNotFound:        "用户不存在",
	ErrUserExists:          "用户名已存在",
	ErrUserCreate:          "创建用户失败",
	ErrUserUpdate:          "更新用户失败",
	ErrUserDelete:          "删除用户失败",
	ErrRoleNotFound:        "角色不存在",
	ErrRoleExists:          "角色已存在",
	ErrRoleAssign:          "角色分配失败",
	ErrMenuNotFound:        "菜单不存在",
	ErrSalesOrderNotFound:  "销售订单不存在",
	ErrMPSNotFound:         "MPS计划不存在",
	ErrMRPNotFound:         "MRP计划不存在",
}

// GetErrorMessage 获取错误消息
func GetErrorMessage(code int) string {
	if msg, ok := errorMessages[code]; ok {
		return msg
	}
	return "未知错误"
}
