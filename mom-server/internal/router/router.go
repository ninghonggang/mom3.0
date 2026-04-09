package router

import (
	"github.com/gin-gonic/gin"
	"mom-server/internal/handler/andon"
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
	"mom-server/internal/handler/ai"
	"mom-server/internal/handler/container"
	"mom-server/internal/handler/wms"
	"mom-server/internal/middleware"
	"mom-server/internal/pkg/jwt"
)

// Router 全局路由
type Router struct {
	engine              *gin.Engine
	jwtUtil            *jwt.JWT
	userHandler         *system.UserHandler
	authHandler         *system.AuthHandler
	loginLogHandler     *system.LoginLogHandler
	roleHandler         *system.RoleHandler
	menuHandler         *system.MenuHandler
	deptHandler         *system.DeptHandler
	dictHandler         *system.DictHandler
	postHandler         *system.PostHandler
	tenantHandler       *system.TenantHandler
	importHandler       *system.ImportHandler
	warehouseHandler    *wms.WarehouseHandler
	salesOrderHandler   *production.SalesOrderHandler
	reportHandler       *production.ReportHandler
	dispatchHandler     *production.DispatchHandler
	mpsHandler         *aps.MPSHandler
	mrpHandler         *aps.MRPHandler
	scheduleHandler     *aps.ScheduleHandler
	workCenterHandler  *aps.WorkCenterHandler
	traceHandler       *trace.TraceHandler
	energyHandler      *trace.EnergyHandler
	equipmentHandler  *equipment.EquipmentHandler
	checkHandler       *equipment.EquipmentCheckHandler
	maintHandler       *equipment.EquipmentMaintenanceHandler
	repairHandler      *equipment.EquipmentRepairHandler
	sparePartHandler   *equipment.SparePartHandler
	lineHandler        *business.ProductionLineHandler
	workstationHandler *business.WorkstationHandler
	shiftHandler       *business.ShiftHandler
	bomHandler         *mdm.BOMHandler
	opHandler          *mdm.OperationHandler
	mdmShiftHandler    *mdm.ShiftHandler
	productionOrderHandler *production.ProductionOrderHandler
	iqcHandler        *quality.IQCHandler
	ipqcHandler       *quality.IPQCHandler
	fqcHandler        *quality.FQCHandler
	oqcHandler        *quality.OQCHandler
	defectCodeHandler *quality.DefectCodeHandler
	defectRecordHandler *quality.DefectRecordHandler
	ncrHandler        *quality.NCRHandler
	spcHandler        *quality.SPCHandler
	supplierHandler    *supplier.SupplierHandler
	materialHandler    *mdm.MaterialHandler
	materialCategoryHandler *mdm.MaterialCategoryHandler
	customerHandler   *mdm.CustomerHandler
	workshopHandler    *mdm.WorkshopHandler
	operLogHandler     *system.OperLogHandler
	oeeHandler       *equipment.OEEHandler
	teepDataHandler   *equipment.TEEPDataHandler
	moldHandler       *equipment.MoldHandler
	moldMaintenanceHandler *equipment.MoldMaintenanceHandler
	moldRepairHandler *equipment.MoldRepairHandler
	gaugeHandler     *equipment.GaugeHandler
	gaugeCalibrationHandler *equipment.GaugeCalibrationHandler
	firstLastInspectHandler *production.FirstLastInspectHandler
	packageHandler     *production.PackageHandler
	dcHandler          *dc.DataCollectionHandler
	electronicSOPHandler *production.ElectronicSOPHandler
	codeRuleHandler    *production.CodeRuleHandler
	flowCardHandler   *production.FlowCardHandler
	noticeHandler     *system.NoticeHandler
	printTemplateHandler *system.PrintTemplateHandler
	capacityAnalysisHandler *aps.CapacityAnalysisHandler
	deliveryRateHandler *aps.DeliveryRateHandler
	changeoverMatrixHandler *aps.ChangeoverMatrixHandler
	rollingScheduleHandler *aps.RollingScheduleHandler
	jitDemandHandler *aps.JITDemandHandler
	transferOrderHandler *wms.TransferOrderHandler
	stockCheckHandler *wms.StockCheckHandler
	sideLocationHandler *wms.SideLocationHandler
	kanbanPullHandler *wms.KanbanPullHandler
	containerHandler  *container.ContainerHandler
	aiConfigHandler  *ai.AIConfigHandler
	aiChatHandler    *ai.AIChatHandler
	andonCallHandler   *andon.CallHandler
	andonRuleHandler  *andon.RuleHandler
	workshopConfigHandler *mdm.WorkshopConfigHandler
	workingCalendarHandler *mdm.WorkingCalendarHandler
}

// New 创建路由
func New(
	jwtUtil *jwt.JWT,
	userHandler *system.UserHandler,
	authHandler *system.AuthHandler,
	loginLogHandler *system.LoginLogHandler,
	roleHandler *system.RoleHandler,
	menuHandler *system.MenuHandler,
	deptHandler *system.DeptHandler,
	dictHandler *system.DictHandler,
	postHandler *system.PostHandler,
	tenantHandler *system.TenantHandler,
	warehouseHandler *wms.WarehouseHandler,
	salesOrderHandler *production.SalesOrderHandler,
	reportHandler *production.ReportHandler,
	dispatchHandler *production.DispatchHandler,
	mpsHandler *aps.MPSHandler,
	mrpHandler *aps.MRPHandler,
	scheduleHandler *aps.ScheduleHandler,
	workCenterHandler *aps.WorkCenterHandler,
	traceHandler *trace.TraceHandler,
	energyHandler *trace.EnergyHandler,
	equipmentHandler *equipment.EquipmentHandler,
	checkHandler *equipment.EquipmentCheckHandler,
	maintHandler *equipment.EquipmentMaintenanceHandler,
	repairHandler *equipment.EquipmentRepairHandler,
	sparePartHandler *equipment.SparePartHandler,
	lineHandler *business.ProductionLineHandler,
	workstationHandler *business.WorkstationHandler,
	shiftHandler *business.ShiftHandler,
	bomHandler *mdm.BOMHandler,
	opHandler *mdm.OperationHandler,
	mdmShiftHandler *mdm.ShiftHandler,
	productionOrderHandler *production.ProductionOrderHandler,
	iqcHandler *quality.IQCHandler,
	ipqcHandler *quality.IPQCHandler,
	fqcHandler *quality.FQCHandler,
	oqcHandler *quality.OQCHandler,
	defectCodeHandler *quality.DefectCodeHandler,
	defectRecordHandler *quality.DefectRecordHandler,
	ncrHandler *quality.NCRHandler,
	spcHandler *quality.SPCHandler,
	supplierHandler *supplier.SupplierHandler,
	materialHandler *mdm.MaterialHandler,
	materialCategoryHandler *mdm.MaterialCategoryHandler,
	customerHandler *mdm.CustomerHandler,
	workshopHandler *mdm.WorkshopHandler,
	operLogHandler *system.OperLogHandler,
	oeeHandler *equipment.OEEHandler,
	teepDataHandler *equipment.TEEPDataHandler,
	moldHandler *equipment.MoldHandler,
	moldMaintenanceHandler *equipment.MoldMaintenanceHandler,
	moldRepairHandler *equipment.MoldRepairHandler,
	gaugeHandler *equipment.GaugeHandler,
	gaugeCalibrationHandler *equipment.GaugeCalibrationHandler,
	importHandler *system.ImportHandler,
	firstLastInspectHandler *production.FirstLastInspectHandler,
	packageHandler *production.PackageHandler,
	dcHandler *dc.DataCollectionHandler,
	electronicSOPHandler *production.ElectronicSOPHandler,
	codeRuleHandler *production.CodeRuleHandler,
	flowCardHandler *production.FlowCardHandler,
	noticeHandler *system.NoticeHandler,
	printTemplateHandler *system.PrintTemplateHandler,
	capacityAnalysisHandler *aps.CapacityAnalysisHandler,
	deliveryRateHandler *aps.DeliveryRateHandler,
	changeoverMatrixHandler *aps.ChangeoverMatrixHandler,
	rollingScheduleHandler *aps.RollingScheduleHandler,
	jitDemandHandler *aps.JITDemandHandler,
	transferOrderHandler *wms.TransferOrderHandler,
	stockCheckHandler *wms.StockCheckHandler,
	sideLocationHandler *wms.SideLocationHandler,
	kanbanPullHandler *wms.KanbanPullHandler,
	containerHandler *container.ContainerHandler,
	aiConfigHandler *ai.AIConfigHandler,
	aiChatHandler *ai.AIChatHandler,
	andonCallHandler *andon.CallHandler,
	andonRuleHandler *andon.RuleHandler,
	workshopConfigHandler *mdm.WorkshopConfigHandler,
	workingCalendarHandler *mdm.WorkingCalendarHandler,
) *Router {
	return &Router{
		jwtUtil:             jwtUtil,
		userHandler:         userHandler,
		authHandler:         authHandler,
		loginLogHandler:     loginLogHandler,
		roleHandler:         roleHandler,
		menuHandler:         menuHandler,
		deptHandler:         deptHandler,
		dictHandler:         dictHandler,
		postHandler:         postHandler,
		tenantHandler:       tenantHandler,
		warehouseHandler:    warehouseHandler,
		salesOrderHandler:   salesOrderHandler,
		reportHandler:       reportHandler,
		dispatchHandler:     dispatchHandler,
		mpsHandler:          mpsHandler,
		mrpHandler:          mrpHandler,
		scheduleHandler:      scheduleHandler,
		workCenterHandler:   workCenterHandler,
		traceHandler:        traceHandler,
		energyHandler:       energyHandler,
		equipmentHandler:    equipmentHandler,
		checkHandler:        checkHandler,
		maintHandler:        maintHandler,
		repairHandler:       repairHandler,
		sparePartHandler:    sparePartHandler,
		lineHandler:         lineHandler,
		workstationHandler:  workstationHandler,
		shiftHandler:        shiftHandler,
		bomHandler:         bomHandler,
		opHandler:          opHandler,
		mdmShiftHandler:    mdmShiftHandler,
		productionOrderHandler: productionOrderHandler,
		iqcHandler:            iqcHandler,
		ipqcHandler:           ipqcHandler,
		fqcHandler:            fqcHandler,
		oqcHandler:            oqcHandler,
		defectCodeHandler:     defectCodeHandler,
		defectRecordHandler:   defectRecordHandler,
		ncrHandler:            ncrHandler,
		spcHandler:            spcHandler,
		supplierHandler:       supplierHandler,
		materialHandler:       materialHandler,
		materialCategoryHandler: materialCategoryHandler,
		customerHandler:     customerHandler,
		workshopHandler:       workshopHandler,
		operLogHandler:         operLogHandler,
		oeeHandler:            oeeHandler,
			teepDataHandler:       teepDataHandler,
			moldHandler:           moldHandler,
			moldMaintenanceHandler: moldMaintenanceHandler,
			moldRepairHandler:     moldRepairHandler,
			gaugeHandler:          gaugeHandler,
			gaugeCalibrationHandler: gaugeCalibrationHandler,
		importHandler:         importHandler,
			firstLastInspectHandler: firstLastInspectHandler,
			packageHandler:          packageHandler,
			dcHandler:               dcHandler,
			electronicSOPHandler:    electronicSOPHandler,
			codeRuleHandler:          codeRuleHandler,
			flowCardHandler:          flowCardHandler,
			noticeHandler:           noticeHandler,
			printTemplateHandler:     printTemplateHandler,
			capacityAnalysisHandler:  capacityAnalysisHandler,
			deliveryRateHandler:      deliveryRateHandler,
			changeoverMatrixHandler:  changeoverMatrixHandler,
			rollingScheduleHandler:   rollingScheduleHandler,
			jitDemandHandler:          jitDemandHandler,
		containerHandler:           containerHandler,
		aiConfigHandler:            aiConfigHandler,
		aiChatHandler:              aiChatHandler,
		andonCallHandler:          andonCallHandler,
		andonRuleHandler:         andonRuleHandler,
		workshopConfigHandler:     workshopConfigHandler,
		workingCalendarHandler:   workingCalendarHandler,
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
	protected.Use(middleware.JWTAuth(r.jwtUtil))
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
				user.PUT("/:id/roles", r.userHandler.AssignRoles)
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
				role.GET("/:id/perms", r.roleHandler.GetPerms)
				role.PUT("/:id/perms", r.roleHandler.AssignPerms)
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

			// 租户管理
			tenant := system.Group("/tenant")
			{
				tenant.GET("/list", r.tenantHandler.List)
				tenant.GET("/:id", r.tenantHandler.Get)
				tenant.POST("", r.tenantHandler.Create)
				tenant.PUT("/:id", r.tenantHandler.Update)
				tenant.DELETE("/:id", r.tenantHandler.Delete)
			}

			// 操作日志
			operLog := system.Group("/operlog")
			{
				operLog.GET("/list", r.operLogHandler.List)
			}

			// 登录日志
			loginLog := system.Group("/loginlog")
			{
				loginLog.GET("/list", r.loginLogHandler.List)
				loginLog.DELETE("/clean", r.loginLogHandler.Clean)
			}
		}

		// 生产执行
		salesOrder := protected.Group("/production/sales-order")
		{
			salesOrder.GET("/list", r.salesOrderHandler.List)
			salesOrder.GET("/:id", r.salesOrderHandler.Get)
			salesOrder.POST("", r.salesOrderHandler.Create)
			salesOrder.PUT("/:id", r.salesOrderHandler.Update)
			salesOrder.DELETE("/:id", r.salesOrderHandler.Delete)
			salesOrder.PUT("/:id/confirm", r.salesOrderHandler.Confirm)
		}

		// 生产报工
		report := protected.Group("/production/report")
		{
			report.GET("/list", r.reportHandler.List)
			report.POST("", r.reportHandler.Create)
		}

		// 派工
		dispatch := protected.Group("/production/dispatch")
		{
			dispatch.GET("/list", r.dispatchHandler.List)
			dispatch.POST("", r.dispatchHandler.Create)
			dispatch.PUT("/:id", r.dispatchHandler.Update)
			dispatch.PUT("/:id/start", r.dispatchHandler.Start)
			dispatch.PUT("/:id/complete", r.dispatchHandler.Complete)
		}

		// 生产工单
		order := protected.Group("/production/order")
		{
			order.GET("/list", r.productionOrderHandler.List)
			order.GET("/:id", r.productionOrderHandler.Get)
			order.POST("", r.productionOrderHandler.Create)
			order.PUT("/:id", r.productionOrderHandler.Update)
			order.DELETE("/:id", r.productionOrderHandler.Delete)
			order.PUT("/:id/start", r.productionOrderHandler.Start)
			order.PUT("/:id/complete", r.productionOrderHandler.Complete)
		}

		// 包装条码
		packages := protected.Group("/production/packages")
		{
			packages.GET("/list", r.packageHandler.List)
			packages.GET("/:id", r.packageHandler.Get)
			packages.POST("/create", r.packageHandler.Create)
			packages.POST("/add-item", r.packageHandler.AddItem)
			packages.POST("/seal", r.packageHandler.Seal)
			packages.DELETE("/:id", r.packageHandler.Delete)
		}

		// IQC入库检验
		iqc := protected.Group("/quality/iqc")
		{
			iqc.GET("/list", r.iqcHandler.List)
			iqc.GET("/:id", r.iqcHandler.Get)
			iqc.POST("", r.iqcHandler.Create)
			iqc.PUT("/:id", r.iqcHandler.Update)
			iqc.DELETE("/:id", r.iqcHandler.Delete)
		}

		// IPQC过程检验
		ipqc := protected.Group("/quality/ipqc")
		{
			ipqc.GET("/list", r.ipqcHandler.List)
			ipqc.GET("/:id", r.ipqcHandler.Get)
			ipqc.POST("", r.ipqcHandler.Create)
			ipqc.PUT("/:id", r.ipqcHandler.Update)
			ipqc.DELETE("/:id", r.ipqcHandler.Delete)
		}

		// FQC最终检验
		fqc := protected.Group("/quality/fqc")
		{
			fqc.GET("/list", r.fqcHandler.List)
			fqc.GET("/:id", r.fqcHandler.Get)
			fqc.POST("", r.fqcHandler.Create)
			fqc.PUT("/:id", r.fqcHandler.Update)
			fqc.DELETE("/:id", r.fqcHandler.Delete)
		}

		// OQC出货检验
		oqc := protected.Group("/quality/oqc")
		{
			oqc.GET("/list", r.oqcHandler.List)
			oqc.GET("/:id", r.oqcHandler.Get)
			oqc.POST("", r.oqcHandler.Create)
			oqc.PUT("/:id", r.oqcHandler.Update)
			oqc.DELETE("/:id", r.oqcHandler.Delete)
		}

		// DefectCode不良品代码
		defectCode := protected.Group("/quality/defect-code")
		{
			defectCode.GET("/list", r.defectCodeHandler.List)
			defectCode.GET("/:id", r.defectCodeHandler.Get)
			defectCode.POST("", r.defectCodeHandler.Create)
			defectCode.PUT("/:id", r.defectCodeHandler.Update)
			defectCode.DELETE("/:id", r.defectCodeHandler.Delete)
		}

		// DefectRecord不良品记录
		defect := protected.Group("/quality/defect")
		{
			defect.GET("/list", r.defectRecordHandler.List)
			defect.GET("/:id", r.defectRecordHandler.Get)
			defect.POST("", r.defectRecordHandler.Create)
			defect.PUT("/:id", r.defectRecordHandler.Update)
			defect.DELETE("/:id", r.defectRecordHandler.Delete)
			defect.PUT("/:id/handle", r.defectRecordHandler.Handle)
		}

		// NCR不良品处理单
		ncr := protected.Group("/quality/ncr")
		{
			ncr.GET("/list", r.ncrHandler.List)
			ncr.GET("/:id", r.ncrHandler.Get)
			ncr.POST("", r.ncrHandler.Create)
			ncr.PUT("/:id", r.ncrHandler.Update)
			ncr.DELETE("/:id", r.ncrHandler.Delete)
		}

		// SPC数据
		spc := protected.Group("/quality/spc")
		{
			spc.POST("/data", r.spcHandler.Create)
			spc.GET("/chart", r.spcHandler.GetChartData)
			spc.GET("/stats", r.spcHandler.GetStats)
			spc.GET("/list", r.spcHandler.List)
			spc.GET("/:id", r.spcHandler.Get)
			spc.PUT("/:id", r.spcHandler.Update)
			spc.DELETE("/:id", r.spcHandler.Delete)
		}

		// APS计划
		aps := protected.Group("/aps")
		{
			mps := aps.Group("/mps")
			{
				mps.GET("/list", r.mpsHandler.List)
				mps.GET("/:id", r.mpsHandler.Get)
				mps.POST("", r.mpsHandler.Create)
				mps.PUT("/:id", r.mpsHandler.Update)
				mps.DELETE("/:id", r.mpsHandler.Delete)
				mps.PUT("/:id/submit", r.mpsHandler.Submit)
			}
			mrp := aps.Group("/mrp")
			{
				mrp.GET("/list", r.mrpHandler.List)
				mrp.PUT("/:id/calculate", r.mrpHandler.Calculate)
			}
			schedule := aps.Group("/schedule")
			{
				schedule.GET("/list", r.scheduleHandler.List)
				schedule.POST("", r.scheduleHandler.Create)
				schedule.PUT("/:id/execute", r.scheduleHandler.Execute)
				schedule.POST("/execute-constrained", r.scheduleHandler.ExecuteConstrained)
				schedule.GET("/:id/results", r.scheduleHandler.GetResults)
				schedule.DELETE("/:id", r.scheduleHandler.Delete)
				schedule.PUT("/drag-update", r.scheduleHandler.DragUpdate)
				schedule.GET("/suggestions/:plan_id", r.scheduleHandler.GetSuggestions)
			}
			workCenter := aps.Group("/work-center")
			{
				workCenter.GET("/list", r.workCenterHandler.List)
				workCenter.GET("/:id", r.workCenterHandler.Get)
				workCenter.POST("", r.workCenterHandler.Create)
				workCenter.PUT("/:id", r.workCenterHandler.Update)
				workCenter.DELETE("/:id", r.workCenterHandler.Delete)
				workCenter.GET("/by-workshop", r.workCenterHandler.ListByWorkshop)
			}
		}

		// 仓储管理
		wms := protected.Group("/wms")
		{
			warehouse := wms.Group("/warehouse")
			{
				warehouse.GET("/list", r.warehouseHandler.ListWarehouse)
				warehouse.GET("/:id", r.warehouseHandler.ListWarehouse)
				warehouse.POST("", r.warehouseHandler.CreateWarehouse)
				warehouse.PUT("/:id", r.warehouseHandler.UpdateWarehouse)
				warehouse.DELETE("/:id", r.warehouseHandler.DeleteWarehouse)
			}
			location := wms.Group("/location")
			{
				location.GET("/list", r.warehouseHandler.ListLocation)
				location.GET("/:id", r.warehouseHandler.GetLocation)
				location.POST("", r.warehouseHandler.CreateLocation)
				location.PUT("/:id", r.warehouseHandler.UpdateLocation)
				location.DELETE("/:id", r.warehouseHandler.DeleteLocation)
			}
			inventory := wms.Group("/inventory")
			{
				inventory.GET("/list", r.warehouseHandler.ListInventory)
				inventory.GET("/:id", r.warehouseHandler.GetInventory)
				inventory.POST("", r.warehouseHandler.CreateInventory)
				inventory.PUT("/:id", r.warehouseHandler.UpdateInventory)
				inventory.DELETE("/:id", r.warehouseHandler.DeleteInventory)
			}

			// 调拨管理
			transfer := wms.Group("/transfer")
			{
				transfer.GET("/list", r.transferOrderHandler.List)
				transfer.GET("/:id", r.transferOrderHandler.Get)
				transfer.POST("", r.transferOrderHandler.Create)
				transfer.PUT("/:id", r.transferOrderHandler.Update)
				transfer.DELETE("/:id", r.transferOrderHandler.Delete)
				transfer.POST("/item", r.transferOrderHandler.AddItem)
			}

			// 盘点管理
			stockcheck := wms.Group("/stock-check")
			{
				stockcheck.GET("/list", r.stockCheckHandler.List)
				stockcheck.GET("/:id", r.stockCheckHandler.Get)
				stockcheck.POST("", r.stockCheckHandler.Create)
				stockcheck.PUT("/:id", r.stockCheckHandler.Update)
				stockcheck.POST("/item", r.stockCheckHandler.AddItem)
				stockcheck.PUT("/item/:id", r.stockCheckHandler.UpdateItem)
			}

			// 线边库位
			sideloc := wms.Group("/side-location")
			{
				sideloc.GET("/list", r.sideLocationHandler.List)
				sideloc.GET("/:id", r.sideLocationHandler.Get)
				sideloc.POST("", r.sideLocationHandler.Create)
				sideloc.PUT("/:id", r.sideLocationHandler.Update)
				sideloc.DELETE("/:id", r.sideLocationHandler.Delete)
			}

			// 看板拉动
			kanban := wms.Group("/kanban")
			{
				kanban.GET("/list", r.kanbanPullHandler.List)
				kanban.GET("/:id", r.kanbanPullHandler.Get)
				kanban.POST("", r.kanbanPullHandler.Create)
				kanban.PUT("/:id", r.kanbanPullHandler.Update)
				kanban.DELETE("/:id", r.kanbanPullHandler.Delete)
			}
		}

		// 追溯管理
		trace := protected.Group("/trace")
		{
			trace.GET("/serial", r.traceHandler.TraceBySerial)
			trace.GET("/batch", r.traceHandler.TraceByBatch)
			trace.GET("/order/:id", r.traceHandler.TraceByOrder)
			trace.GET("/forward", r.traceHandler.ForwardTrace)
			trace.GET("/backward", r.traceHandler.BackwardTrace)
		}

		// 安东呼叫
		andon := protected.Group("/andon")
		{
			call := andon.Group("/calls")
			{
				call.GET("/list", r.andonCallHandler.List)
				call.GET("/:id", r.andonCallHandler.Get)
				call.POST("", r.andonCallHandler.Create)
				call.PUT("/:id/respond", r.andonCallHandler.Respond)
				call.PUT("/:id/resolve", r.andonCallHandler.Resolve)
				call.PUT("/:id/escalate", r.andonCallHandler.Escalate)
			}

			// 升级规则管理
			rules := andon.Group("/escalation-rules")
			{
				rules.GET("/list", r.andonRuleHandler.List)
				rules.GET("/:id", r.andonRuleHandler.Get)
				rules.POST("", r.andonRuleHandler.Create)
				rules.PUT("/:id", r.andonRuleHandler.Update)
				rules.DELETE("/:id", r.andonRuleHandler.Delete)
			}

			// 统计分析
			andon.GET("/statistics", r.andonCallHandler.GetStatistics)
		}

		// 能源管理
		energy := protected.Group("/energy")
		{
			record := energy.Group("/record")
			{
				record.GET("/list", r.energyHandler.List)
				record.POST("", r.energyHandler.Create)
			}
			energy.GET("/stats", r.energyHandler.GetStats)
			energy.GET("/trend", r.energyHandler.GetTrend)
		}

		// 设备管理
		equipment := protected.Group("/equipment")
		{
			equipment.GET("/list", r.equipmentHandler.List)
			equipment.GET("/:id", r.equipmentHandler.Get)
			equipment.POST("", r.equipmentHandler.Create)
			equipment.PUT("/:id", r.equipmentHandler.Update)
			equipment.DELETE("/:id", r.equipmentHandler.Delete)
			equipment.GET("/status", r.equipmentHandler.Status)
		}

		// 设备点检
		protected.Group("/equipment/check").GET("/list", r.checkHandler.List)

		// 设备保养
		protected.Group("/equipment/maintenance").GET("/list", r.maintHandler.List)

		// 设备维修
		equipmentRepair := protected.Group("/equipment/repair")
		{
			equipmentRepair.GET("/list", r.repairHandler.List)
			equipmentRepair.POST("", r.repairHandler.Create)
			equipmentRepair.PUT("/:id/start", r.repairHandler.Start)
			equipmentRepair.PUT("/:id/complete", r.repairHandler.Complete)
		}

		// 备件
		protected.Group("/equipment/spare").GET("/list", r.sparePartHandler.List)

		// OEE分析
		oee := protected.Group("/equipment/oee")
		{
			oee.GET("/list", r.oeeHandler.List)
			oee.GET("/:id", r.oeeHandler.Get)
			oee.POST("/calculate", r.oeeHandler.Calculate)
			oee.GET("/chart", r.oeeHandler.Chart)
			oee.DELETE("/:id", r.oeeHandler.Delete)
		}

		// TEEP分析
		teep := protected.Group("/equipment/teep")
		{
			teep.GET("/list", r.teepDataHandler.List)
			teep.GET("/:id", r.teepDataHandler.Get)
			teep.POST("", r.teepDataHandler.Create)
			teep.PUT("/:id", r.teepDataHandler.Update)
			teep.DELETE("/:id", r.teepDataHandler.Delete)
		}

		// 模具管理
		mold := protected.Group("/equipment/mold")
		{
			mold.GET("/list", r.moldHandler.List)
			mold.GET("/:id", r.moldHandler.Get)
			mold.POST("", r.moldHandler.Create)
			mold.PUT("/:id", r.moldHandler.Update)
			mold.DELETE("/:id", r.moldHandler.Delete)
			mold.GET("/maintenance/list", r.moldMaintenanceHandler.List)
			mold.POST("/maintenance", r.moldMaintenanceHandler.Create)
			mold.GET("/repair/list", r.moldRepairHandler.List)
			mold.POST("/repair", r.moldRepairHandler.Create)
		}

		// 量检具管理
		gauge := protected.Group("/equipment/gauge")
		{
			gauge.GET("/list", r.gaugeHandler.List)
			gauge.GET("/:id", r.gaugeHandler.Get)
			gauge.POST("", r.gaugeHandler.Create)
			gauge.PUT("/:id", r.gaugeHandler.Update)
			gauge.DELETE("/:id", r.gaugeHandler.Delete)
			gauge.GET("/calibration/list", r.gaugeCalibrationHandler.List)
			gauge.POST("/calibration", r.gaugeCalibrationHandler.Create)
		}

		// 生产线
		line := protected.Group("/mdm/line")
		{
			line.GET("/list", r.lineHandler.List)
			line.POST("", r.lineHandler.Create)
			line.PUT("/:id", r.lineHandler.Update)
			line.DELETE("/:id", r.lineHandler.Delete)
		}

		// 工位
		workstation := protected.Group("/mdm/workstation")
		{
			workstation.GET("/list", r.workstationHandler.List)
			workstation.POST("", r.workstationHandler.Create)
			workstation.PUT("/:id", r.workstationHandler.Update)
			workstation.DELETE("/:id", r.workstationHandler.Delete)
		}

		// MDM 物料管理
		material := protected.Group("/mdm/material")
		{
			material.GET("/list", r.materialHandler.List)
			material.GET("/:id", r.materialHandler.Get)
			material.POST("", r.materialHandler.Create)
			material.PUT("/:id", r.materialHandler.Update)
			material.DELETE("/:id", r.materialHandler.Delete)
		}

		// MDM 物料分类管理
		materialCategory := protected.Group("/mdm/material-category")
		{
			materialCategory.GET("/list", r.materialCategoryHandler.List)
			materialCategory.GET("/tree", r.materialCategoryHandler.Tree)
			materialCategory.GET("/:id", r.materialCategoryHandler.Get)
			materialCategory.POST("", r.materialCategoryHandler.Create)
			materialCategory.PUT("/:id", r.materialCategoryHandler.Update)
			materialCategory.DELETE("/:id", r.materialCategoryHandler.Delete)
		}

		// MDM 客户管理
		customer := protected.Group("/mdm/customer")
		{
			customer.GET("/list", r.customerHandler.List)
			customer.GET("/:id", r.customerHandler.Get)
			customer.POST("", r.customerHandler.Create)
			customer.PUT("/:id", r.customerHandler.Update)
			customer.DELETE("/:id", r.customerHandler.Delete)
		}

		// MDM 车间管理
		workshop := protected.Group("/mdm/workshop")
		{
			workshop.GET("/list", r.workshopHandler.List)
			workshop.GET("/:id", r.workshopHandler.Get)
			workshop.POST("", r.workshopHandler.Create)
			workshop.PUT("/:id", r.workshopHandler.Update)
			workshop.DELETE("/:id", r.workshopHandler.Delete)
		}

		// MDM 车间配置
		workshopConfig := protected.Group("/mdm/workshop-config")
		{
			workshopConfig.GET("/:workshop_id", r.workshopConfigHandler.GetConfig)
			workshopConfig.PUT("/:workshop_id", r.workshopConfigHandler.UpdateConfig)
		}

		// APS 工厂日历
		calendar := protected.Group("/aps/calendar")
		{
			calendar.GET("", r.workingCalendarHandler.GetCalendars)
			calendar.POST("", r.workingCalendarHandler.CreateCalendar)
			calendar.PUT("/:id", r.workingCalendarHandler.UpdateCalendar)
			calendar.DELETE("/:id", r.workingCalendarHandler.DeleteCalendar)
		}

		// MDM BOM管理
		bom := protected.Group("/mdm/bom")
		{
			bom.GET("/list", r.bomHandler.List)
			bom.GET("/:id", r.bomHandler.Get)
			bom.GET("/:id/items", r.bomHandler.GetWithItems)
			bom.POST("", r.bomHandler.Create)
			bom.PUT("/:id", r.bomHandler.Update)
			bom.DELETE("/:id", r.bomHandler.Delete)
			bom.PUT("/:id/status", r.bomHandler.UpdateStatus)
			bom.POST("/:id/copy", r.bomHandler.CopyBOM)
			bom.GET("/template", r.importHandler.DownloadBOMTemplate)
			bom.POST("/import", r.importHandler.ImportBOM)
		}

		// MDM 工序管理
		operation := protected.Group("/mdm/operation")
		{
			operation.GET("/list", r.opHandler.List)
			operation.GET("/:id", r.opHandler.Get)
			operation.POST("", r.opHandler.Create)
			operation.PUT("/:id", r.opHandler.Update)
			operation.DELETE("/:id", r.opHandler.Delete)
		}

		// MDM 班次管理
		mdmShift := protected.Group("/mdm/mdm-shift")
		{
			mdmShift.GET("/list", r.mdmShiftHandler.List)
			mdmShift.GET("/:id", r.mdmShiftHandler.Get)
			mdmShift.POST("", r.mdmShiftHandler.Create)
			mdmShift.PUT("/:id", r.mdmShiftHandler.Update)
			mdmShift.DELETE("/:id", r.mdmShiftHandler.Delete)
		}

		// 供应商管理
		supplier := protected.Group("/mdm/supplier")
		{
			supplier.GET("/list", r.supplierHandler.List)
			supplier.GET("/:id", r.supplierHandler.Get)
			supplier.POST("", r.supplierHandler.Create)
			supplier.PUT("/:id", r.supplierHandler.Update)
			supplier.DELETE("/:id", r.supplierHandler.Delete)
		}

		// 首末件检验
		firstLastInspect := protected.Group("/production/first-last-inspect")
		{
			firstLastInspect.GET("/list", r.firstLastInspectHandler.List)
			firstLastInspect.GET("/:id", r.firstLastInspectHandler.Get)
			firstLastInspect.GET("/overdue", r.firstLastInspectHandler.ListOverdue)
			firstLastInspect.POST("", r.firstLastInspectHandler.Create)
			firstLastInspect.PUT("/:id", r.firstLastInspectHandler.Update)
			firstLastInspect.DELETE("/:id", r.firstLastInspectHandler.Delete)
		}

		// 数据采集
		dc := protected.Group("/dc")
		{
			dataPoint := dc.Group("/data-point")
			{
				dataPoint.GET("/list", r.dcHandler.ListDataPoint)
				dataPoint.GET("/:id", r.dcHandler.GetDataPoint)
				dataPoint.POST("", r.dcHandler.CreateDataPoint)
				dataPoint.PUT("/:id", r.dcHandler.UpdateDataPoint)
				dataPoint.DELETE("/:id", r.dcHandler.DeleteDataPoint)
			}
			scanLog := dc.Group("/scan-log")
			{
				scanLog.GET("/list", r.dcHandler.ListScanLog)
				scanLog.POST("/scan", r.dcHandler.CreateScanLog)
			}
			collect := dc.Group("/collect-record")
			{
				collect.GET("/list", r.dcHandler.ListCollectRecord)
			}
		}

		// TODO: 其他模块路由...

		// 电子SOP
		electronicSOP := protected.Group("/production/electronic-sop")
		{
			electronicSOP.GET("/list", r.electronicSOPHandler.List)
			electronicSOP.GET("/:id", r.electronicSOPHandler.Get)
			electronicSOP.POST("", r.electronicSOPHandler.Create)
			electronicSOP.PUT("/:id", r.electronicSOPHandler.Update)
			electronicSOP.DELETE("/:id", r.electronicSOPHandler.Delete)
		}

		// 编码规则
		codeRule := protected.Group("/production/code-rule")
		{
			codeRule.GET("/list", r.codeRuleHandler.List)
			codeRule.GET("/:id", r.codeRuleHandler.Get)
			codeRule.POST("", r.codeRuleHandler.Create)
			codeRule.PUT("/:id", r.codeRuleHandler.Update)
			codeRule.DELETE("/:id", r.codeRuleHandler.Delete)
			codeRule.GET("/generate", r.codeRuleHandler.GenerateCode)
		}

		// 生产指示单
		flowCard := protected.Group("/production/flow-card")
		{
			flowCard.GET("/list", r.flowCardHandler.List)
			flowCard.GET("/:id", r.flowCardHandler.Get)
			flowCard.POST("", r.flowCardHandler.Create)
			flowCard.PUT("/:id", r.flowCardHandler.Update)
			flowCard.DELETE("/:id", r.flowCardHandler.Delete)
		}

		// 通知公告
		notice := protected.Group("/system/notice")
		{
			notice.GET("/list", r.noticeHandler.List)
			notice.GET("/:id", r.noticeHandler.Get)
			notice.POST("", r.noticeHandler.Create)
			notice.PUT("/:id", r.noticeHandler.Update)
			notice.DELETE("/:id", r.noticeHandler.Delete)
			notice.PUT("/:id/publish", r.noticeHandler.Publish)
			notice.GET("/my", r.noticeHandler.GetMyNotices)
		}

		// 打印模板
		printTemplate := protected.Group("/system/print-template")
		{
			printTemplate.GET("/list", r.printTemplateHandler.List)
			printTemplate.GET("/:id", r.printTemplateHandler.Get)
			printTemplate.POST("", r.printTemplateHandler.Create)
			printTemplate.PUT("/:id", r.printTemplateHandler.Update)
			printTemplate.DELETE("/:id", r.printTemplateHandler.Delete)
		}

		// APS产能分析
		capacity := protected.Group("/aps/capacity")
		{
			capacity.GET("/list", r.capacityAnalysisHandler.List)
			capacity.GET("/:id", r.capacityAnalysisHandler.Get)
			capacity.POST("", r.capacityAnalysisHandler.Create)
			capacity.PUT("/:id", r.capacityAnalysisHandler.Update)
			capacity.GET("/stats", r.capacityAnalysisHandler.GetStats)
		}

		// 交付率
		deliveryRate := protected.Group("/aps/delivery-rate")
		{
			deliveryRate.GET("/list", r.deliveryRateHandler.List)
			deliveryRate.GET("/:id", r.deliveryRateHandler.Get)
			deliveryRate.POST("", r.deliveryRateHandler.Create)
			deliveryRate.PUT("/:id", r.deliveryRateHandler.Update)
			deliveryRate.DELETE("/:id", r.deliveryRateHandler.Delete)
		}

		// 换型矩阵
		changeover := protected.Group("/aps/changeover")
		{
			changeover.GET("/list", r.changeoverMatrixHandler.List)
			changeover.GET("/:id", r.changeoverMatrixHandler.Get)
			changeover.POST("", r.changeoverMatrixHandler.Create)
			changeover.PUT("/:id", r.changeoverMatrixHandler.Update)
			changeover.DELETE("/:id", r.changeoverMatrixHandler.Delete)
		}

		// 滚动排程
		rolling := protected.Group("/aps/rolling")
		{
			rolling.GET("/list", r.rollingScheduleHandler.List)
			rolling.GET("/:id", r.rollingScheduleHandler.Get)
			rolling.POST("", r.rollingScheduleHandler.Create)
			rolling.PUT("/:id", r.rollingScheduleHandler.Update)
			rolling.DELETE("/:id", r.rollingScheduleHandler.Delete)
		}

		// JIT需求
		jit := protected.Group("/aps/jit")
		{
			jit.GET("/list", r.jitDemandHandler.List)
			jit.GET("/:id", r.jitDemandHandler.Get)
			jit.POST("", r.jitDemandHandler.Create)
			jit.PUT("/:id", r.jitDemandHandler.Update)
			jit.DELETE("/:id", r.jitDemandHandler.Delete)
		}

		// 器具管理
		container := protected.Group("/containers")
		{
			container.GET("/list", r.containerHandler.List)
			container.GET("/:id", r.containerHandler.Get)
			container.POST("", r.containerHandler.Create)
			container.PUT("/:id", r.containerHandler.Update)
			container.DELETE("/:id", r.containerHandler.Delete)
			container.POST("/:id/in", r.containerHandler.In)
			container.POST("/:id/out", r.containerHandler.Out)
			container.POST("/:id/return", r.containerHandler.Return)
			container.POST("/:id/transfer", r.containerHandler.Transfer)
			container.POST("/:id/clean", r.containerHandler.Clean)
			container.GET("/:id/movements", r.containerHandler.Movements)
		}

		// AI Chat
		ai := protected.Group("/ai")
		{
			ai.GET("/config", r.aiConfigHandler.GetConfig)
			ai.PUT("/config", r.aiConfigHandler.UpdateConfig)
			ai.POST("/config/test", r.aiConfigHandler.TestConfig)
			ai.GET("/schema", r.aiConfigHandler.GetSchema)

			chat := ai.Group("/chat")
			{
				chat.GET("/conversations", r.aiChatHandler.ListConversations)
				chat.GET("/conversations/:session_id", r.aiChatHandler.GetConversation)
				chat.DELETE("/conversations/:session_id", r.aiChatHandler.DeleteConversation)
				chat.POST("/send", r.aiChatHandler.SendMessage)
				chat.POST("/execute", r.aiChatHandler.ExecuteOperation)
			}
		}
	}
}

// SetJWT 设置JWT中间件
func (r *Router) SetJWT(jwtFunc func() gin.HandlerFunc) {
	protected := r.engine.Group("/api/v1")
	protected.Use(jwtFunc())
}
