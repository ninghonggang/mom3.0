package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mom-server/internal/config"
	"mom-server/internal/handler/aps"
	"mom-server/internal/handler/business"
	"mom-server/internal/handler/dc"
	"mom-server/internal/handler/equipment"
	"mom-server/internal/handler/mdm"
	"mom-server/internal/handler/production"
	"mom-server/internal/handler/quality"
	"mom-server/internal/handler/supplier"
	"mom-server/internal/handler/system"
	"mom-server/internal/handler/trace"
	"mom-server/internal/handler/wms"
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

	// 自动迁移（分批执行，便于错误排查）
	log.Println("开始数据库迁移...")
	
	// 第1批：系统基础表（14个）
	log.Println("迁移第1批：系统基础表")
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
		log.Fatalf("数据库迁移失败[第1批-系统基础表]: %v", err)
	}
	
	// 第2批：仓储管理表（7个）
	log.Println("迁移第2批：仓储管理表")
	if err := db.AutoMigrate(
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
	); err != nil {
		log.Fatalf("数据库迁移失败[第2批-仓储管理表]: %v", err)
	}
	
	// 第3批：生产执行表（6个）
	log.Println("迁移第3批：生产执行表")
	if err := db.AutoMigrate(
		&model.SalesOrder{},
		&model.SalesOrderItem{},
		&model.ProductionReport{},
		&model.Dispatch{},
		&model.ProductionOrder{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第3批-生产执行表]: %v", err)
	}
	
	// 第4批：APS计划表（8个）
	log.Println("迁移第4批：APS计划表")
	if err := db.AutoMigrate(
		&model.MPS{},
		&model.MRP{},
		&model.MRPItem{},
		&model.SchedulePlan{},
		&model.ScheduleResult{},
		&model.Resource{},
		&model.WorkCenter{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第4批-APS计划表]: %v", err)
	}
	
	// 第5批：追溯管理表（5个）
	log.Println("迁移第5批：追溯管理表")
	if err := db.AutoMigrate(
		&model.SerialNumber{},
		&model.TraceRecord{},
		&model.AndonCall{},
		&model.DataCollection{},
		&model.EnergyRecord{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第5批-追溯管理表]: %v", err)
	}
	
	// 第6批：主数据管理表（跳过已有数据的表避免迁移错误）
	log.Println("迁移第6批：主数据管理表")
	if err := db.AutoMigrate(
		&model.Material{},
		&model.MaterialCategory{},
		// BOM/BOMItem/Process/Route/RouteOperation等表已有数据,跳过迁移
		&model.Workshop{},
		&model.ProductionLine{},
		&model.Workstation{},
		&model.Shift{},
		&model.MdmShift{},
		&model.Supplier{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第6批-主数据管理表]: %v", err)
	}

	// 第7批：质量管理表
	log.Println("迁移第7批：质量管理表")
	if err := db.AutoMigrate(
		&model.IQC{},
		&model.IQCItem{},
		&model.IPQC{},
		&model.FQC{},
		&model.OQC{},
		&model.DefectCode{},
		&model.DefectRecord{},
		&model.NCR{},
		&model.SPCData{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第7批-质量管理表]: %v", err)
	}

	// 第8批：设备OEE表
	log.Println("迁移第8批：设备OEE表")
	if err := db.AutoMigrate(
		&model.OEE{},
		&model.OEEEvent{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第8批-设备OEE表]: %v", err)
	}

	// 第9批：首末件检验表
	log.Println("迁移第9批：首末件检验表")
	if err := db.AutoMigrate(
		&model.MesFirstLastInspect{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第9批-首末件检验表]: %v", err)
	}

	// 第10批：包装条码表
	log.Println("迁移第10批：包装条码表")
	if err := db.AutoMigrate(
		&model.MesPackage{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第10批-包装条码表]: %v", err)
	}

	// 第11批：数据采集表
	log.Println("迁移第11批：数据采集表")
	if err := db.AutoMigrate(
		&model.DCDataPoint{},
		&model.DCCollectRecord{},
		&model.DCScanLog{},
	); err != nil {
		log.Fatalf("数据库迁移失败[第11批-数据采集表]: %v", err)
	}

	log.Println("数据库迁移完成")

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
	tenantRepo := repository.NewTenantRepository(db)
	roleMenuRepo := repository.NewRoleMenuRepository(db)
	warehouseRepo := repository.NewWarehouseRepository(db)
	locationRepo := repository.NewLocationRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	receiveOrderRepo := repository.NewReceiveOrderRepository(db)
	receiveOrderItemRepo := repository.NewReceiveOrderItemRepository(db)
	deliveryOrderRepo := repository.NewDeliveryOrderRepository(db)
	deliveryOrderItemRepo := repository.NewDeliveryOrderItemRepository(db)
	salesOrderRepo := repository.NewSalesOrderRepository(db)
	reportRepo := repository.NewProductionReportRepository(db)
	productionOrderChangeLogRepo := repository.NewProductionOrderChangeLogRepository(db)
	dispatchRepo := repository.NewDispatchRepository(db)
	mpsRepo := repository.NewMPSRepository(db)
	mrpRepo := repository.NewMRPRepository(db)
	scheduleRepo := repository.NewScheduleRepository(db)
	workCenterRepo := repository.NewWorkCenterRepository(db)
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
	materialRepo := repository.NewMaterialRepository(db)
	materialCategoryRepo := repository.NewMaterialCategoryRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	workshopRepo := repository.NewWorkshopRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	shiftRepo := repository.NewShiftRepository(db)
	bomRepo := repository.NewBOMRepository(db)
	bomItemRepo := repository.NewBOMItemRepository(db)
	opRepo := repository.NewOperationRepository(db)
	mdmShiftRepo := repository.NewMdmShiftRepository(db)
	iqcRepo := repository.NewIQCRepository(db)
	ipqcRepo := repository.NewIPQCRepository(db)
	fqcRepo := repository.NewFQCRepository(db)
	oqcRepo := repository.NewOQCRepository(db)
	defectCodeRepo := repository.NewDefectCodeRepository(db)
	defectRecordRepo := repository.NewDefectRecordRepository(db)
	ncrRepo := repository.NewNCRRepository(db)
	spcRepo := repository.NewSPCDataRepository(db)
	operLogRepo := repository.NewOperLogRepository(db)
	loginLogRepo := repository.NewLoginLogRepository(db)
	oeeRepo := repository.NewOEERepository(db)
	oeeEventRepo := repository.NewOEEEventRepository(db)
	firstLastInspectRepo := repository.NewFirstLastInspectRepository(db)
	packageRepo := repository.NewPackageRepository(db)
	dcDataPointRepo := repository.NewDCDataPointRepository(db)
	dcScanLogRepo := repository.NewDCScanLogRepository(db)
	dcCollectRecordRepo := repository.NewDCCollectRecordRepository(db)
	importTaskRepo := repository.NewImportTaskRepository(db)

	// 初始化服务层
	userSvc := service.NewUserService(userRepo, roleRepo, menuRepo, roleMenuRepo)
	materialSvc := service.NewMaterialService(materialRepo)
	materialCategorySvc := service.NewMaterialCategoryService(materialCategoryRepo)
	customerSvc := service.NewCustomerService(customerRepo)
	workshopSvc := service.NewWorkshopService(workshopRepo)
	roleSvc := service.NewRoleService(roleRepo, menuRepo, roleMenuRepo)
	menuSvc := service.NewMenuService(menuRepo)
	deptSvc := service.NewDeptService(deptRepo)
	dictSvc := service.NewDictService(dictTypeRepo, dictDataRepo)
	postSvc := service.NewPostService(postRepo)
	tenantSvc := service.NewTenantService(tenantRepo)
	warehouseSvc := service.NewWarehouseService(db, warehouseRepo, locationRepo, inventoryRepo, receiveOrderRepo, receiveOrderItemRepo, deliveryOrderRepo, deliveryOrderItemRepo)
	salesOrderSvc := service.NewSalesOrderService(salesOrderRepo)
	reportSvc := service.NewProductionReportService(reportRepo)
	dispatchSvc := service.NewDispatchService(dispatchRepo)
	mpsSvc := service.NewMPSService(mpsRepo)
	mrpSvc := service.NewMRPService(mrpRepo, inventoryRepo, bomRepo, bomItemRepo, mpsRepo)
	scheduleSvc := service.NewScheduleService(scheduleRepo, productionRepo, lineRepo)
	workCenterSvc := service.NewWorkCenterService(workCenterRepo)
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
	bomSvc := service.NewBOMService(bomRepo, bomItemRepo)
	opSvc := service.NewOperationService(opRepo)
	mdmShiftSvc := service.NewMdmShiftService(mdmShiftRepo)
	iqcSvc := service.NewIQCService(iqcRepo)
	ipqcSvc := service.NewIPQCService(ipqcRepo)
	fqcSvc := service.NewFQCService(fqcRepo)
	oqcSvc := service.NewOQCService(oqcRepo)
	defectCodeSvc := service.NewDefectCodeService(defectCodeRepo)
	defectRecordSvc := service.NewDefectRecordService(defectRecordRepo)
	ncrSvc := service.NewNCRService(ncrRepo)
	spcSvc := service.NewSPCDataService(spcRepo)
	supplierSvc := service.NewSupplierService(supplierRepo)
	operLogSvc := service.NewOperLogService(operLogRepo)
	loginLogSvc := service.NewLoginLogService(loginLogRepo)
	productionOrderChangeLogSvc := service.NewProductionOrderChangeLogService(productionOrderChangeLogRepo)
	productionOrderSvc := service.NewProductionOrderService(productionRepo, productionOrderChangeLogSvc)
	oeeSvc := service.NewOEEService(oeeRepo, oeeEventRepo)
	firstLastInspectSvc := service.NewFirstLastInspectService(firstLastInspectRepo)
	packageSvc := service.NewPackageService(packageRepo)
	dcSvc := service.NewDCService(dcDataPointRepo, dcScanLogRepo, dcCollectRecordRepo)
	importSvc := service.NewImportService(importTaskRepo, materialRepo)

	// 初始化处理器层
	authHandler := system.NewAuthHandler(userSvc, jwtUtil, loginLogSvc)
	userHandler := system.NewUserHandler(userSvc)
	roleHandler := system.NewRoleHandler(roleSvc)
	menuHandler := system.NewMenuHandler(menuSvc)
	deptHandler := system.NewDeptHandler(deptSvc)
	dictHandler := system.NewDictHandler(dictSvc)
	postHandler := system.NewPostHandler(postSvc)
	tenantHandler := system.NewTenantHandler(tenantSvc)
	warehouseHandler := wms.NewWarehouseHandler(warehouseSvc)
	salesOrderHandler := production.NewSalesOrderHandler(salesOrderSvc)
	reportHandler := production.NewReportHandler(reportSvc)
	dispatchHandler := production.NewDispatchHandler(dispatchSvc)
	apsMPSHandler := aps.NewMPSHandler(mpsSvc)
	apsMRPHandler := aps.NewMRPHandler(mrpSvc)
	apsScheduleHandler := aps.NewScheduleHandler(scheduleSvc)
	workCenterHandler := aps.NewWorkCenterHandler(workCenterSvc)
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
	bomHandler := mdm.NewBOMHandler(bomSvc)
	opHandler := mdm.NewOperationHandler(opSvc)
	mdmShiftHandler := mdm.NewShiftHandler(mdmShiftSvc)
	productionOrderHandler := production.NewProductionOrderHandler(productionOrderSvc)
	iqcHandler := quality.NewIQCHandler(iqcSvc)
	ipqcHandler := quality.NewIPQCHandler(ipqcSvc)
	fqcHandler := quality.NewFQCHandler(fqcSvc)
	oqcHandler := quality.NewOQCHandler(oqcSvc)
	defectCodeHandler := quality.NewDefectCodeHandler(defectCodeSvc)
	defectRecordHandler := quality.NewDefectRecordHandler(defectRecordSvc)
	ncrHandler := quality.NewNCRHandler(ncrSvc)
	spcHandler := quality.NewSPCHandler(spcSvc)
	supplierHandler := supplier.NewSupplierHandler(supplierSvc)
	materialHandler := mdm.NewMaterialHandler(materialSvc)
	materialCategoryHandler := mdm.NewMaterialCategoryHandler(materialCategorySvc)
	customerHandler := mdm.NewCustomerHandler(customerSvc)
	workshopHandler := mdm.NewWorkshopHandler(workshopSvc)
	operLogHandler := system.NewOperLogHandler(operLogSvc)
	loginLogHandler := system.NewLoginLogHandler(loginLogSvc)
	importHandler := system.NewImportHandler(importSvc)
	oeeHandler := equipment.NewOEEHandler(oeeSvc)
	firstLastInspectHandler := production.NewFirstLastInspectHandler(firstLastInspectSvc)
	packageHandler := production.NewPackageHandler(packageSvc)
	dcHandler := dc.NewDataCollectionHandler(dcSvc)

	// 初始化路由
	gin.SetMode(cfg.Server.Mode)
	engine := gin.Default()
	router.New(jwtUtil, userHandler, authHandler, loginLogHandler, roleHandler, menuHandler, deptHandler, dictHandler, postHandler, tenantHandler, warehouseHandler, salesOrderHandler, reportHandler, dispatchHandler, apsMPSHandler, apsMRPHandler, apsScheduleHandler, workCenterHandler, traceHandler, andonHandler, energyHandler, checkHandler, maintHandler, repairHandler, sparePartHandler, lineHandler, workstationHandler, shiftHandler, bomHandler, opHandler, mdmShiftHandler, productionOrderHandler, iqcHandler, ipqcHandler, fqcHandler, oqcHandler, defectCodeHandler, defectRecordHandler, ncrHandler, spcHandler, supplierHandler, materialHandler, materialCategoryHandler, customerHandler, workshopHandler, operLogHandler, oeeHandler, importHandler, firstLastInspectHandler, packageHandler, dcHandler).Init(engine)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("服务器启动: http://localhost%s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
