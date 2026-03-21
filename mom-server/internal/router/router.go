package router

import (
	"github.com/gin-gonic/gin"
	"mom-server/internal/handler/system"
	"mom-server/internal/middleware"
)

// Router 全局路由
type Router struct {
	engine      *gin.Engine
	userHandler *system.UserHandler
	authHandler *system.AuthHandler
	roleHandler *system.RoleHandler
	menuHandler *system.MenuHandler
	deptHandler *system.DeptHandler
	dictHandler *system.DictHandler
	postHandler *system.PostHandler
}

// New 创建路由
func New(
	userHandler *system.UserHandler,
	authHandler *system.AuthHandler,
	roleHandler *system.RoleHandler,
	menuHandler *system.MenuHandler,
	deptHandler *system.DeptHandler,
	dictHandler *system.DictHandler,
	postHandler *system.PostHandler,
) *Router {
	return &Router{
		userHandler:  userHandler,
		authHandler:  authHandler,
		roleHandler:  roleHandler,
		menuHandler:  menuHandler,
		deptHandler:  deptHandler,
		dictHandler:  dictHandler,
		postHandler:  postHandler,
	}
}

// Init 初始化路由
func (r *Router) Init(engine *gin.Engine) {
	r.engine = engine

	// 中间件
	r.engine.Use(middleware.CORS())
	r.engine.Use(middleware.Recovery())
	r.engine.Use(middleware.Logger())

	// 公开路由
	public := r.engine.Group("/api/v1")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/login", r.authHandler.Login)
			auth.POST("/logout", r.authHandler.Logout)
			auth.POST("/refresh", r.authHandler.RefreshToken)
		}
	}

	// 需要认证的路由
	protected := r.engine.Group("/api/v1")
	protected.Use(middleware.JWTAuth(nil))
	{
		// 认证相关
		auth := protected.Group("/auth")
		{
			auth.GET("/info", r.authHandler.GetUserInfo)
			auth.PUT("/password", r.authHandler.ChangePassword)
		}

		// 系统管理
		system := protected.Group("/system")
		{
			// 用户管理
			user := system.Group("/user")
			{
				user.GET("/list", r.userHandler.GetList)
				user.GET("/:id", r.userHandler.GetByID)
				user.POST("", r.userHandler.Create)
				user.PUT("/:id", r.userHandler.Update)
				user.DELETE("/:id", r.userHandler.Delete)
				user.PUT("/:id/password", r.userHandler.ResetPassword)
			}

			// 角色管理
			role := system.Group("/role")
			{
				role.GET("/list", r.roleHandler.List)
				role.GET("/:id", r.roleHandler.Get)
				role.POST("", r.roleHandler.Create)
				role.PUT("/:id", r.roleHandler.Update)
				role.DELETE("/:id", r.roleHandler.Delete)
				role.GET("/:id/menus", r.roleHandler.GetMenus)
				role.PUT("/:id/menus", r.roleHandler.AssignMenus)
			}

			// 菜单管理
			menu := system.Group("/menu")
			{
				menu.GET("/list", r.menuHandler.List)
				menu.GET("/tree", r.menuHandler.Tree)
				menu.GET("/:id", r.menuHandler.Get)
				menu.POST("", r.menuHandler.Create)
				menu.PUT("/:id", r.menuHandler.Update)
				menu.DELETE("/:id", r.menuHandler.Delete)
			}

			// 部门管理
			dept := system.Group("/dept")
			{
				dept.GET("/list", r.deptHandler.List)
				dept.GET("/tree", r.deptHandler.Tree)
				dept.GET("/:id", r.deptHandler.Get)
				dept.POST("", r.deptHandler.Create)
				dept.PUT("/:id", r.deptHandler.Update)
				dept.DELETE("/:id", r.deptHandler.Delete)
			}

			// 字典管理
			dict := system.Group("/dict")
			{
				dictType := dict.Group("/type")
				{
					dictType.GET("/list", r.dictHandler.ListType)
					dictType.GET("/:id", r.dictHandler.GetType)
					dictType.POST("", r.dictHandler.CreateType)
					dictType.PUT("/:id", r.dictHandler.UpdateType)
					dictType.DELETE("/:id", r.dictHandler.DeleteType)
				}
				dict.GET("/:dictType/data", r.dictHandler.GetData)
			}

			// 岗位管理
			post := system.Group("/post")
			{
				post.GET("/list", r.postHandler.List)
				post.GET("/:id", r.postHandler.Get)
				post.POST("", r.postHandler.Create)
				post.PUT("/:id", r.postHandler.Update)
				post.DELETE("/:id", r.postHandler.Delete)
			}
		}

		// TODO: 其他模块路由...
	}
}

// SetJWT 设置JWT中间件
func (r *Router) SetJWT(jwtFunc func() gin.HandlerFunc) {
	protected := r.engine.Group("/api/v1")
	protected.Use(jwtFunc())
}
