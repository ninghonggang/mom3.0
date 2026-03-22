package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mom-server/internal/config"
	"mom-server/internal/handler/system"
	"mom-server/internal/handler/aps"
	"mom-server/internal/handler/business"
	"mom-server/internal/handler/equipment"
	"mom-server/internal/handler/production"
	"mom-server/internal/handler/trace"
	"mom-server/internal/handler/wms"
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
		&model.Warehouse{},
		&model.Location{},
		&model.Inventory{},
		&model.InventoryRecord{},
		&model.ReceiveOrder{},
		&model.ReceiveOrderItem{},
		&model.DeliveryOrder{},
		&model.DeliveryOrderItem{},
		&model.TransferOrder{},
		&model.StockCheck{},
		&model.SalesOrder{},
		&model.SalesOrderItem{},
		&model.ProductionReport{},
		&model.Dispatch{},
		&model.MPS{},
		&model.MRP{},
		&model.MRPItem{},
		&model.SchedulePlan{},
		&model.ScheduleResult{},
		&model.Resource{},
		&model.WorkCenter{},
		&model.SerialNumber{},
		&model.TraceRecord{},
		&model.AndonCall{},
		&model.DataCollection{},
		&model.EnergyRecord{},
	); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化JWT
	log.Printf("DEBUG: cfg.Server.JWT.Secret = '%s'", cfg.Server.JWT.Secret)
	log.Printf("DEBUG: cfg.Server.JWT.AccessTokenExpire = %v", cfg.Server.JWT.AccessTokenExpire)
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
	warehouseRepo := repository.NewWarehouseRepository(db)
	locationRepo := repository.NewLocationRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	salesOrderRepo := repository.NewSalesOrderRepository(db)
	reportRepo := repository.NewProductionReportRepository(db)
	dispatchRepo := repository.NewDispatchRepository(db)
	mpsRepo := repository.NewMPSRepository(db)
	mrpRepo := repository.NewMRPRepository(db)
	scheduleRepo := repository.NewScheduleRepository(db)
	productionRepo := repository.NewProductionOrderRepository(db)
	traceRepo := repository.NewTraceRepository(db)
	andonRepo := repository.NewAndonRepository(db)
	energyRepo := repository.NewEnergyRepository(db)
	checkRepo := repository.NewEquipmentCheckRepository(db)
	maintRepo := repository.NewEquipmentMaintenanceRepository(db)
	repairRepo := repository.NewEquipmentRepairRepository(db)
	sparePartRepo := repository.NewSparePartRepository(db)
	lineRepo := repository.NewProductionLineRepository(db)
	workstationRepo := repository.NewWorkstationRepository(db)
	shiftRepo := repository.NewShiftRepository(db)

	// 初始化服务层
	userSvc := service.NewUserService(userRepo, roleRepo)
	roleSvc := service.NewRoleService(roleRepo, menuRepo, roleMenuRepo)
	menuSvc := service.NewMenuService(menuRepo)
	deptSvc := service.NewDeptService(deptRepo)
	dictSvc := service.NewDictService(dictTypeRepo, dictDataRepo)
	postSvc := service.NewPostService(postRepo)
	warehouseSvc := service.NewWarehouseService(warehouseRepo, locationRepo, inventoryRepo)
	salesOrderSvc := service.NewSalesOrderService(salesOrderRepo)
	reportSvc := service.NewProductionReportService(reportRepo)
	dispatchSvc := service.NewDispatchService(dispatchRepo)
	mpsSvc := service.NewMPSService(mpsRepo)
	mrpSvc := service.NewMRPService(mrpRepo, inventoryRepo)
	scheduleSvc := service.NewScheduleService(scheduleRepo, productionRepo)
	traceSvc := service.NewTraceService(traceRepo)
	andonSvc := service.NewAndonService(andonRepo)
	energySvc := service.NewEnergyService(energyRepo)
	checkSvc := service.NewEquipmentCheckService(checkRepo)
	maintSvc := service.NewEquipmentMaintenanceService(maintRepo)
	repairSvc := service.NewEquipmentRepairService(repairRepo)
	sparePartSvc := service.NewSparePartService(sparePartRepo)
	lineSvc := service.NewProductionLineService(lineRepo)
	workstationSvc := service.NewWorkstationService(workstationRepo)
	shiftSvc := service.NewShiftService(shiftRepo)

	// 初始化处理器层
	authHandler := system.NewAuthHandler(userSvc, jwtUtil)
	userHandler := system.NewUserHandler(userSvc)
	roleHandler := system.NewRoleHandler(roleSvc)
	menuHandler := system.NewMenuHandler(menuSvc)
	deptHandler := system.NewDeptHandler(deptSvc)
	dictHandler := system.NewDictHandler(dictSvc)
	postHandler := system.NewPostHandler(postSvc)
	warehouseHandler := wms.NewWarehouseHandler(warehouseSvc)
	salesOrderHandler := production.NewSalesOrderHandler(salesOrderSvc)
	reportHandler := production.NewReportHandler(reportSvc)
	dispatchHandler := production.NewDispatchHandler(dispatchSvc)
	apsMPSHandler := aps.NewMPSHandler(mpsSvc)
	apsMRPHandler := aps.NewMRPHandler(mrpSvc)
	apsScheduleHandler := aps.NewScheduleHandler(scheduleSvc)
	traceHandler := trace.NewTraceHandler(traceSvc)
	andonHandler := trace.NewAndonHandler(andonSvc)
	energyHandler := trace.NewEnergyHandler(energySvc)
	checkHandler := equipment.NewEquipmentCheckHandler(checkSvc)
	maintHandler := equipment.NewEquipmentMaintenanceHandler(maintSvc)
	repairHandler := equipment.NewEquipmentRepairHandler(repairSvc)
	sparePartHandler := equipment.NewSparePartHandler(sparePartSvc)
	lineHandler := business.NewProductionLineHandler(lineSvc)
	workstationHandler := business.NewWorkstationHandler(workstationSvc)
	shiftHandler := business.NewShiftHandler(shiftSvc)

	// 初始化路由
	gin.SetMode(cfg.Server.Mode)
	engine := gin.Default()
	r := router.New(userHandler, authHandler, roleHandler, menuHandler, deptHandler, dictHandler, postHandler, warehouseHandler, salesOrderHandler, reportHandler, dispatchHandler, apsMPSHandler, apsMRPHandler, apsScheduleHandler, traceHandler, andonHandler, energyHandler, checkHandler, maintHandler, repairHandler, sparePartHandler, lineHandler, workstationHandler, shiftHandler)
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
