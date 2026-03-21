package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mom-server/internal/config"
	"mom-server/internal/handler/system"
	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/jwt"
	"mom-server/internal/repository"
	"mom-server/internal/router"
	"mom-server/internal/service"
)

func main() {
	// 加载配置
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移
	if err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Menu{},
		&model.Dept{},
		&model.Post{},
		&model.DictType{},
		&model.DictData{},
		&model.Tenant{},
		&model.OperLog{},
		&model.LoginLog{},
		&model.RoleMenu{},
		&model.UserRole{},
	); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化JWT
	jwtUtil := jwt.New(&cfg.Server.JWT)

	// 初始化仓储层
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	deptRepo := repository.NewDeptRepository(db)
	dictTypeRepo := repository.NewDictTypeRepository(db)
	dictDataRepo := repository.NewDictDataRepository(db)
	postRepo := repository.NewPostRepository(db)
	roleMenuRepo := repository.NewRoleMenuRepository(db)

	// 初始化服务层
	userSvc := service.NewUserService(userRepo, roleRepo)
	roleSvc := service.NewRoleService(roleRepo, menuRepo, roleMenuRepo)
	menuSvc := service.NewMenuService(menuRepo)
	deptSvc := service.NewDeptService(deptRepo)
	dictSvc := service.NewDictService(dictTypeRepo, dictDataRepo)
	postSvc := service.NewPostService(postRepo)

	// 初始化处理器层
	authHandler := system.NewAuthHandler(userSvc, jwtUtil)
	userHandler := system.NewUserHandler(userSvc)
	roleHandler := system.NewRoleHandler(roleSvc)
	menuHandler := system.NewMenuHandler(menuSvc)
	deptHandler := system.NewDeptHandler(deptSvc)
	dictHandler := system.NewDictHandler(dictSvc)
	postHandler := system.NewPostHandler(postSvc)

	// 初始化路由
	gin.SetMode(cfg.Server.Mode)
	engine := gin.Default()
	r := router.New(userHandler, authHandler, roleHandler, menuHandler, deptHandler, dictHandler, postHandler)
	r.Init(engine)

	// 设置JWT中间件
	r.SetJWT(func() gin.HandlerFunc {
		return middleware.JWTAuth(jwtUtil)
	})

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("服务器启动: http://localhost%s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
