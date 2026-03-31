package router

import (
	"github.com/gin-gonic/gin"
	"mom-server/internal/handler/aps"
	"mom-server/internal/handler/business"
	"mom-server/internal/handler/equipment"
	"mom-server/internal/handler/mdm"
	"mom-server/internal/handler/production"
	"mom-server/internal/handler/quality"
	"mom-server/internal/handler/supplier"
	"mom-server/internal/handler/system"
	"mom-server/internal/handler/trace"
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
	roleHandler         *system.RoleHandler
	menuHandler         *system.MenuHandler
	deptHandler         *system.DeptHandler
	dictHandler         *system.DictHandler
	postHandler         *system.PostHandler
	warehouseHandler    *wms.WarehouseHandler
	salesOrderHandler   *production.SalesOrderHandler
	reportHandler       *production.ReportHandler
	dispatchHandler     *production.DispatchHandler
	mpsHandler         *aps.MPSHandler
	mrpHandler         *aps.MRPHandler
	scheduleHandler     *aps.ScheduleHandler
	traceHandler       *trace.TraceHandler
	andonHandler       *trace.AndonHandler
	energyHandler      *trace.EnergyHandler
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
}

// New 创建路由
func New(
	jwtUtil *jwt.JWT,
	userHandler *system.UserHandler,
	authHandler *system.AuthHandler,
	roleHandler *system.RoleHandler,
	menuHandler *system.MenuHandler,
	deptHandler *system.DeptHandler,
	dictHandler *system.DictHandler,
	postHandler *system.PostHandler,
	warehouseHandler *wms.WarehouseHandler,
	salesOrderHandler *production.SalesOrderHandler,
	reportHandler *production.ReportHandler,
	dispatchHandler *production.DispatchHandler,
	mpsHandler *aps.MPSHandler,
	mrpHandler *aps.MRPHandler,
	scheduleHandler *aps.ScheduleHandler,
	traceHandler *trace.TraceHandler,
	andonHandler *trace.AndonHandler,
	energyHandler *trace.EnergyHandler,
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
) *Router {
	return &Router{
		jwtUtil:             jwtUtil,
		userHandler:         userHandler,
		authHandler:         authHandler,
		roleHandler:         roleHandler,
		menuHandler:         menuHandler,
		deptHandler:         deptHandler,
		dictHandler:         dictHandler,
		postHandler:         postHandler,
		warehouseHandler:    warehouseHandler,
		salesOrderHandler:   salesOrderHandler,
		reportHandler:       reportHandler,
		dispatchHandler:     dispatchHandler,
		mpsHandler:          mpsHandler,
		mrpHandler:          mrpHandler,
		scheduleHandler:      scheduleHandler,
		traceHandler:        traceHandler,
		andonHandler:        andonHandler,
		energyHandler:       energyHandler,
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
				schedule.GET("/:id/results", r.scheduleHandler.GetResults)
				schedule.DELETE("/:id", r.scheduleHandler.Delete)
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
		}

		// 追溯管理
		trace := protected.Group("/trace")
		{
			trace.GET("/serial", r.traceHandler.TraceBySerial)
			trace.GET("/batch", r.traceHandler.TraceByBatch)
			trace.GET("/order/:id", r.traceHandler.TraceByOrder)
		}

		// 安东呼叫
		andon := protected.Group("/andon")
		{
			call := andon.Group("/call")
			{
				call.GET("/list", r.andonHandler.List)
				call.POST("", r.andonHandler.Create)
				call.PUT("/:id/response", r.andonHandler.Response)
				call.PUT("/:id/resolve", r.andonHandler.Resolve)
			}
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

		// 班次 (旧版本)
		shift := protected.Group("/mdm/shift")
		{
			shift.GET("/list", r.shiftHandler.List)
			shift.POST("", r.shiftHandler.Create)
			shift.PUT("/:id", r.shiftHandler.Update)
			shift.DELETE("/:id", r.shiftHandler.Delete)
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

		// TODO: 其他模块路由...
	}
}

// SetJWT 设置JWT中间件
func (r *Router) SetJWT(jwtFunc func() gin.HandlerFunc) {
	protected := r.engine.Group("/api/v1")
	protected.Use(jwtFunc())
}
