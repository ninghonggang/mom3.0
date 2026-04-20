package router

import (
	"github.com/gin-gonic/gin"
	"mom-server/internal/handler/andon"
	"mom-server/internal/handler/aps"
	"mom-server/internal/handler/business"
	"mom-server/internal/handler/container"
	"mom-server/internal/handler/dc"
	"mom-server/internal/handler/equipment"
	"mom-server/internal/handler/alert"
	"mom-server/internal/handler/bpm"
	"mom-server/internal/handler/eam"
	"mom-server/internal/handler/fin"
	"mom-server/internal/handler/mdm"
	"mom-server/internal/handler/scp"
	"mom-server/internal/handler/production"
	"mom-server/internal/handler/quality"
"mom-server/internal/handler/supplier"
	"mom-server/internal/handler/supplier_asn"
	"mom-server/internal/handler/system"
	"mom-server/internal/handler/trace"
	"mom-server/internal/handler/ai"
	"mom-server/internal/handler/wms"
	"mom-server/internal/handler/agv"
	"mom-server/internal/handler/erp_sync"
	"mom-server/internal/handler/integration"
	"mom-server/internal/handler/mes"
	"mom-server/internal/handler/report"
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
	supplierHandler       *supplier.SupplierHandler
	supplierASNHandler    *supplier_asn.SupplierASNHandler
	materialHandler       *mdm.MaterialHandler
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
	finHandler             *fin.FinHandler
	equipmentPartHandler     *equipment.EquipmentPartHandler
	equipmentDocumentHandler *equipment.EquipmentDocumentHandler
	equipmentDowntimeHandler *eam.EquipmentDowntimeHandler
	spareHandler            *eam.SpareHandler
	alertHandler            *alert.AlertHandler
	bpmHandler              *bpm.BPMHandler
	bpmTaskMsgRuleHandler  *bpm.BpmTaskMessageRuleHandler
	bpmInstanceApiHandler  *bpm.BpmInstanceApiHandler
	bpmTaskTransferHandler *bpm.TaskTransferHandler
	rfqHandler              *scp.RFQHandler
	purchaseOrderHandler     *scp.PurchaseOrderHandler
	scpSalesOrderHandler     *scp.SalesOrderHandler
	supplierKPIHandler      *scp.SupplierKPIHandler
	supplierQuoteHandler    *scp.SupplierQuoteHandler
	customerInquiryHandler   *scp.CustomerInquiryHandler
	purchasePlanHandler     *scp.PurchasePlanHandler
	scpSupplierExtHandler    *scp.SupplierExtHandler
	qadHandler               *scp.QadHandler
	contactHandler                   *mdm.ContactHandler
	bankAccountHandler                *mdm.BankAccountHandler
	attachmentHandler                 *mdm.AttachmentHandler
	supplierMaterialHandler           *supplier.SupplierMaterialHandler
	containerLifecycleHandler         *container.ContainerLifecycleHandler
	visualInspectionHandler           *ai.VisualInspectionHandler
	labSampleHandler                 *quality.LabSampleHandler
	labTestItemHandler               *quality.LabTestItemHandler
	labReportHandler                 *quality.LabReportHandler
	labInstrumentHandler             *quality.LabInstrumentHandler
	inspectionFeatureHandler         *quality.InspectionFeatureHandler
	inspectionCharacteristicHandler  *quality.InspectionCharacteristicHandler
	aqlHandler                       *quality.AQLHandler
	qmsSamplingHandler               *quality.QMSSamplingHandler
	mesTeamHandler                   *mes.TeamHandler
	mesProcessHandler                *mes.ProcessHandler
	mesOfflineHandler                *mes.OfflineHandler
	productionIssueHandler           *production.ProductionIssueHandler
	productionReturnHandler          *production.ProductionReturnHandler
	productionCompleteHandler        *production.ProductionCompleteHandler
	purchaseReturnHandler            *wms.PurchaseReturnHandler
	salesReturnHandler               *wms.SalesReturnHandler
	mesHandler                       *mes.MesHandler
	personSkillHandler               *mes.PersonSkillHandler
	completeInspectHandler          *mes.CompleteInspectHandler
	productionDailyReportHandler     *report.ProductionDailyReportHandler
	qualityWeeklyReportHandler       *report.QualityWeeklyReportHandler
	oeeReportHandler                 *report.OEEReportHandler
	deliveryReportHandler            *report.DeliveryReportHandler
	andonReportHandler               *report.AndonReportHandler
	integrationHandler              *integration.IntegrationHandler
	agvHandler                    *agv.AGVHandler
	erpSyncHandler                *erp_sync.ERPSyncHandler
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
	supplierASNHandler *supplier_asn.SupplierASNHandler,
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
	finHandler *fin.FinHandler,
	equipmentPartHandler *equipment.EquipmentPartHandler,
	equipmentDocumentHandler *equipment.EquipmentDocumentHandler,
	equipmentDowntimeHandler *eam.EquipmentDowntimeHandler,
	spareHandler *eam.SpareHandler,
	alertHandler *alert.AlertHandler,
	bpmHandler *bpm.BPMHandler,
	bpmTaskMsgRuleHandler *bpm.BpmTaskMessageRuleHandler,
	bpmInstanceApiHandler *bpm.BpmInstanceApiHandler,
	bpmTaskTransferHandler *bpm.TaskTransferHandler,
	rfqHandler *scp.RFQHandler,
	purchaseOrderHandler *scp.PurchaseOrderHandler,
	scpSalesOrderHandler *scp.SalesOrderHandler,
	supplierKPIHandler *scp.SupplierKPIHandler,
	supplierQuoteHandler *scp.SupplierQuoteHandler,
	customerInquiryHandler *scp.CustomerInquiryHandler,
	purchasePlanHandler *scp.PurchasePlanHandler,
	scpSupplierExtHandler *scp.SupplierExtHandler,
	qadHandler *scp.QadHandler,
	contactHandler *mdm.ContactHandler,
	bankAccountHandler *mdm.BankAccountHandler,
	attachmentHandler *mdm.AttachmentHandler,
	supplierMaterialHandler *supplier.SupplierMaterialHandler,
	containerLifecycleHandler *container.ContainerLifecycleHandler,
	visualInspectionHandler *ai.VisualInspectionHandler,
	labSampleHandler *quality.LabSampleHandler,
	labTestItemHandler *quality.LabTestItemHandler,
	labReportHandler *quality.LabReportHandler,
	labInstrumentHandler *quality.LabInstrumentHandler,
	inspectionFeatureHandler *quality.InspectionFeatureHandler,
	inspectionCharacteristicHandler *quality.InspectionCharacteristicHandler,
aqlHandler *quality.AQLHandler,
	qmsSamplingHandler *quality.QMSSamplingHandler,
	mesTeamHandler *mes.TeamHandler,
	mesProcessHandler *mes.ProcessHandler,
	mesOfflineHandler *mes.OfflineHandler,
	productionIssueHandler *production.ProductionIssueHandler,
	productionReturnHandler *production.ProductionReturnHandler,
	productionCompleteHandler *production.ProductionCompleteHandler,
	purchaseReturnHandler *wms.PurchaseReturnHandler,
	salesReturnHandler *wms.SalesReturnHandler,
	mesHandler *mes.MesHandler,
	personSkillHandler *mes.PersonSkillHandler,
	completeInspectHandler *mes.CompleteInspectHandler,
	productionDailyReportHandler *report.ProductionDailyReportHandler,
	qualityWeeklyReportHandler *report.QualityWeeklyReportHandler,
	oeeReportHandler *report.OEEReportHandler,
	deliveryReportHandler *report.DeliveryReportHandler,
	andonReportHandler *report.AndonReportHandler,
	integrationHandler *integration.IntegrationHandler,
	agvHandler *agv.AGVHandler,
	erpSyncHandler *erp_sync.ERPSyncHandler,
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
		supplierASNHandler:    supplierASNHandler,
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
		finHandler:               finHandler,
		equipmentPartHandler:       equipmentPartHandler,
		equipmentDocumentHandler:   equipmentDocumentHandler,
		equipmentDowntimeHandler:   equipmentDowntimeHandler,
		alertHandler:              alertHandler,
			spareHandler:             spareHandler,
		bpmHandler:                bpmHandler,
		bpmTaskMsgRuleHandler:   bpmTaskMsgRuleHandler,
		bpmInstanceApiHandler:   bpmInstanceApiHandler,
		bpmTaskTransferHandler:  bpmTaskTransferHandler,
		rfqHandler:              rfqHandler,
		purchaseOrderHandler:       purchaseOrderHandler,
		scpSalesOrderHandler:      scpSalesOrderHandler,
		supplierKPIHandler:        supplierKPIHandler,
		supplierQuoteHandler:      supplierQuoteHandler,
		purchasePlanHandler:     purchasePlanHandler,
		scpSupplierExtHandler:    scpSupplierExtHandler,
		contactHandler:             contactHandler,
		bankAccountHandler:        bankAccountHandler,
		attachmentHandler:         attachmentHandler,
		supplierMaterialHandler:   supplierMaterialHandler,
		containerLifecycleHandler: containerLifecycleHandler,
		visualInspectionHandler:   visualInspectionHandler,
		labSampleHandler:          labSampleHandler,
		labTestItemHandler:        labTestItemHandler,
		labReportHandler:          labReportHandler,
		labInstrumentHandler:      labInstrumentHandler,
		inspectionFeatureHandler:  inspectionFeatureHandler,
		inspectionCharacteristicHandler: inspectionCharacteristicHandler,
		aqlHandler:               aqlHandler,
		mesTeamHandler:           mesTeamHandler,
		mesProcessHandler:        mesProcessHandler,
		mesOfflineHandler:        mesOfflineHandler,
		productionIssueHandler:   productionIssueHandler,
		productionReturnHandler:  productionReturnHandler,
		productionCompleteHandler: productionCompleteHandler,
		purchaseReturnHandler:    purchaseReturnHandler,
		salesReturnHandler:       salesReturnHandler,
		mesHandler:               mesHandler,
		personSkillHandler:       personSkillHandler,
		completeInspectHandler:          completeInspectHandler,
		productionDailyReportHandler:  productionDailyReportHandler,
		qualityWeeklyReportHandler:    qualityWeeklyReportHandler,
		oeeReportHandler:              oeeReportHandler,
		deliveryReportHandler:         deliveryReportHandler,
		andonReportHandler:            andonReportHandler,
		integrationHandler:           integrationHandler,
		agvHandler:                  agvHandler,
		erpSyncHandler:              erpSyncHandler,
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

		// 生产日报
		prodDaily := protected.Group("/report/production-daily")
		{
			prodDaily.GET("/list", r.productionDailyReportHandler.ListProductionDailyReports)
			prodDaily.GET("/:id", r.productionDailyReportHandler.GetProductionDailyReport)
			prodDaily.POST("/generate", r.productionDailyReportHandler.GenerateDailyReport)
			prodDaily.GET("/summary", r.productionDailyReportHandler.GetDailyReportSummary)
		}

		// 质量周报
		qualityWeekly := protected.Group("/report/quality-weekly")
		{
			qualityWeekly.GET("/list", r.qualityWeeklyReportHandler.ListQualityWeeklyReports)
			qualityWeekly.GET("/:id", r.qualityWeeklyReportHandler.GetQualityWeeklyReport)
			qualityWeekly.POST("/generate", r.qualityWeeklyReportHandler.GenerateWeeklyReport)
		}

		// OEE报表
		oeeReport := protected.Group("/report/oee")
		{
			oeeReport.GET("/list", r.oeeReportHandler.ListOEEReports)
			oeeReport.GET("/:id", r.oeeReportHandler.GetOEEReport)
			oeeReport.POST("/generate", r.oeeReportHandler.GenerateOEEReport)
		}

		// 交付率报表
		deliveryReport := protected.Group("/report/delivery")
		{
			deliveryReport.GET("/list", r.deliveryReportHandler.ListDeliveryReports)
			deliveryReport.GET("/:id", r.deliveryReportHandler.GetDeliveryReport)
			deliveryReport.POST("/generate", r.deliveryReportHandler.GenerateDeliveryReport)
		}

		// 安东报表
		andonReport := protected.Group("/report/andon")
		{
			andonReport.GET("/list", r.andonReportHandler.ListAndonReports)
			andonReport.GET("/:id", r.andonReportHandler.GetAndonReport)
			andonReport.POST("/generate", r.andonReportHandler.GenerateAndonReport)
		}

		// 系统集成 - 接口配置
		integration := protected.Group("/integration")
		{
			integration.GET("/interface-config/list", r.integrationHandler.ListConfigs)
			integration.GET("/interface-config/:id", r.integrationHandler.GetConfig)
			integration.POST("/interface-config", r.integrationHandler.CreateConfig)
			integration.PUT("/interface-config/:id", r.integrationHandler.UpdateConfig)
			integration.DELETE("/interface-config/:id", r.integrationHandler.DeleteConfig)
			integration.POST("/interface-config/:id/execute", r.integrationHandler.ExecuteConfig)
			integration.POST("/interface-config/:id/test", r.integrationHandler.TestConfig)

			// 字段映射
			integration.GET("/interface-config/:id/field-maps", r.integrationHandler.ListFieldMaps)
			integration.POST("/interface-config/:id/field-maps", r.integrationHandler.CreateFieldMap)
			integration.PUT("/field-map/:id", r.integrationHandler.UpdateFieldMap)
			integration.DELETE("/field-map/:id", r.integrationHandler.DeleteFieldMap)

			// 触发器
			integration.GET("/interface-config/:id/triggers", r.integrationHandler.ListTriggers)
			integration.POST("/interface-config/:id/triggers", r.integrationHandler.CreateTrigger)
			integration.PUT("/trigger/:id", r.integrationHandler.UpdateTrigger)
			integration.DELETE("/trigger/:id", r.integrationHandler.DeleteTrigger)

			// 执行日志
			integration.GET("/execution-log/list", r.integrationHandler.ListExecutionLogs)
			integration.GET("/execution-log/:id", r.integrationHandler.GetExecutionLog)

			// 枚举选项
			integration.GET("/options", r.integrationHandler.GetConstantOptions)

			// ERP同步
			integration.GET("/erp/sync-log/list", r.erpSyncHandler.ListSyncLogs)
			integration.GET("/erp/sync-log/:id", r.erpSyncHandler.GetSyncLog)
			integration.GET("/erp/status/:syncId", r.erpSyncHandler.GetSyncStatus)
			integration.POST("/erp/bom/sync", r.erpSyncHandler.SyncBOM)
			integration.POST("/erp/order/sync", r.erpSyncHandler.SyncProductionOrder)
			integration.POST("/erp/stock/sync", r.erpSyncHandler.SyncStock)
			integration.POST("/erp/report/push", r.erpSyncHandler.PushReport)
			integration.POST("/erp/stockin/push", r.erpSyncHandler.PushStockIn)
			integration.POST("/erp/quality/push", r.erpSyncHandler.PushQualityData)
		}

		// AGV管理
		agv := protected.Group("/agv")
		{
			// AGV任务
			agv.GET("/task/list", r.agvHandler.ListTasks)
			agv.GET("/task/:id", r.agvHandler.GetTask)
			agv.POST("/task", r.agvHandler.CreateTask)
			agv.PUT("/task/:id/cancel", r.agvHandler.CancelTask)
			agv.PUT("/task/:id/assign", r.agvHandler.AssignTask)
			agv.PUT("/task/:id/complete", r.agvHandler.CompleteTask)
			agv.PUT("/task/:id/start", r.agvHandler.StartTask)

			// AGV设备
			agv.GET("/device/list", r.agvHandler.ListDevices)
			agv.GET("/device/:id", r.agvHandler.GetDevice)
			agv.POST("/device", r.agvHandler.CreateDevice)
			agv.PUT("/device/:id/status", r.agvHandler.UpdateDeviceStatus)
			agv.GET("/device/available", r.agvHandler.GetAvailableAGVs)

			// AGV库位映射
			agv.GET("/location/list", r.agvHandler.ListLocations)
			agv.GET("/location/:id", r.agvHandler.GetLocation)
			agv.POST("/location", r.agvHandler.CreateLocation)
			agv.PUT("/location/:id", r.agvHandler.UpdateLocation)
			agv.DELETE("/location/:id", r.agvHandler.DeleteLocation)

			// AGV回调
			agv.POST("/callback/heartbeat", r.agvHandler.Heartbeat)
			agv.POST("/callback/task", r.agvHandler.TaskCallback)
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

		// 人员技能矩阵
		personSkill := protected.Group("/mes/person-skill")
		{
			personSkill.GET("/list", r.personSkillHandler.ListPersonSkills)
			personSkill.GET("/:id", r.personSkillHandler.GetPersonSkill)
			personSkill.POST("", r.personSkillHandler.CreatePersonSkill)
			personSkill.PUT("/:id", r.personSkillHandler.UpdatePersonSkill)
			personSkill.DELETE("/:id", r.personSkillHandler.DeletePersonSkill)
			personSkill.GET("/detail/:person_id", r.personSkillHandler.GetPersonSkillDetail)
			personSkill.POST("/evaluate", r.personSkillHandler.EvaluateSkill)
			personSkill.GET("/capability/:person_id", r.personSkillHandler.GetPersonCapability)
		}

		// MES班组管理
		team := protected.Group("/mes/team")
		{
			team.GET("/list", r.mesTeamHandler.List)
			team.GET("/:id", r.mesTeamHandler.Get)
			team.POST("", r.mesTeamHandler.Create)
			team.PUT("/:id", r.mesTeamHandler.Update)
			team.DELETE("/:id", r.mesTeamHandler.Delete)
			team.GET("/:id/members", r.mesTeamHandler.ListMembers)
			team.POST("/:id/members", r.mesTeamHandler.AddMember)
			team.PUT("/members/:member_id", r.mesTeamHandler.UpdateMember)
			team.DELETE("/members/:member_id", r.mesTeamHandler.RemoveMember)
			team.GET("/:id/shifts", r.mesTeamHandler.ListShifts)
			team.POST("/:id/shifts", r.mesTeamHandler.CreateShift)
			team.PUT("/shifts/:shift_id", r.mesTeamHandler.UpdateShift)
			team.DELETE("/shifts/:shift_id", r.mesTeamHandler.DeleteShift)
		}

		// MES工艺路线
		process := protected.Group("/mes/process-routes")
		{
			process.GET("/list", r.mesProcessHandler.List)
			process.GET("/:id", r.mesProcessHandler.Get)
			process.GET("/material/:material_code", r.mesProcessHandler.GetByMaterial)
			process.POST("", r.mesProcessHandler.Create)
			process.PUT("/:id", r.mesProcessHandler.Update)
			process.DELETE("/:id", r.mesProcessHandler.Delete)
			process.PUT("/:id/status", r.mesProcessHandler.UpdateStatus)
			process.POST("/:id/copy", r.mesProcessHandler.Copy)
			process.POST("/validate", r.mesProcessHandler.Validate)
		}

		// MES产品离线
		offline := protected.Group("/mes/offline")
		{
			offline.GET("/list", r.mesOfflineHandler.List)
			offline.GET("/:id", r.mesOfflineHandler.Get)
			offline.POST("", r.mesOfflineHandler.Create)
			offline.PUT("/:id", r.mesOfflineHandler.Update)
			offline.DELETE("/:id", r.mesOfflineHandler.Delete)
			offline.POST("/:id/handle", r.mesOfflineHandler.Handle)
			offline.GET("/:id/items", r.mesOfflineHandler.GetItems)
		}

		// MES齐套检查
		completeInspect := protected.Group("/mes/complete-inspect")
		{
			completeInspect.GET("/get", r.completeInspectHandler.GetConfig)
			completeInspect.POST("/get-orderDay-bom", r.completeInspectHandler.GetOrderDayBom)
			completeInspect.POST("/get-orderDay-bom-page", r.completeInspectHandler.GetOrderDayBomPage)
			completeInspect.POST("/get-orderDay-worker-page", r.completeInspectHandler.GetOrderDayWorkerPage)
			completeInspect.POST("/get-orderDay-equipment-page", r.completeInspectHandler.GetOrderDayEquipmentPage)
			completeInspect.POST("/get-orderDay-equipment", r.completeInspectHandler.GetOrderDayEquipment)
			completeInspect.POST("/get-orderDay-worker", r.completeInspectHandler.GetOrderDayWorker)
			completeInspect.POST("/update", r.completeInspectHandler.Update)
		}

		// 生产完工
		productionComplete := protected.Group("/production/complete")
		{
			productionComplete.GET("/list", r.productionCompleteHandler.List)
			productionComplete.GET("/:id", r.productionCompleteHandler.Get)
			productionComplete.POST("", r.productionCompleteHandler.Create)
			productionComplete.POST("/:id/submit-inspect", r.productionCompleteHandler.SubmitForInspect)
			productionComplete.POST("/:id/qualify", r.productionCompleteHandler.Qualify)
			productionComplete.POST("/:id/stock-in", r.productionCompleteHandler.StockIn)
			productionComplete.GET("/stock-in/list", r.productionCompleteHandler.ListStockIn)
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
			iqc.PUT("/:id/inspect", r.iqcHandler.Inspect) // 检验判定
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
			ncr.PUT("/:id/resolve", r.ncrHandler.Resolve) // NCR解决
			ncr.POST("/:id/assign", r.ncrHandler.Assign) // NCR指派
			ncr.POST("/:id/close", r.ncrHandler.Close)   // NCR关闭
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
			spc.GET("/capability/:configId", r.spcHandler.GetCapability) // CP/CPK分析
		}

		// 实验室仪器
		labInstrument := protected.Group("/quality/lab-instrument")
		{
			labInstrument.GET("/list", r.labInstrumentHandler.ListLabInstruments)
			labInstrument.GET("/:id", r.labInstrumentHandler.GetLabInstrument)
			labInstrument.POST("", r.labInstrumentHandler.CreateLabInstrument)
			labInstrument.PUT("/:id", r.labInstrumentHandler.UpdateLabInstrument)
			labInstrument.DELETE("/:id", r.labInstrumentHandler.DeleteLabInstrument)
			labInstrument.GET("/:id/calibrations", r.labInstrumentHandler.GetLabInstrumentCalibrations)
			labInstrument.POST("/:id/calibrate", r.labInstrumentHandler.RecordCalibration)
		}

		// 检验特性
		inspectionFeature := protected.Group("/quality/inspection-feature")
		{
			inspectionFeature.GET("/list", r.inspectionFeatureHandler.ListInspectionFeatures)
			inspectionFeature.GET("/:id", r.inspectionFeatureHandler.GetInspectionFeature)
			inspectionFeature.POST("", r.inspectionFeatureHandler.CreateInspectionFeature)
			inspectionFeature.PUT("/:id", r.inspectionFeatureHandler.UpdateInspectionFeature)
			inspectionFeature.DELETE("/:id", r.inspectionFeatureHandler.DeleteInspectionFeature)
			inspectionFeature.POST("/batch", r.inspectionFeatureHandler.BatchCreateInspectionFeature)
			inspectionFeature.GET("/product/:product_id", r.inspectionFeatureHandler.GetFeaturesByProduct)
		}

		// QMS抽样方案
		sampling := protected.Group("/qms/sampling")
		{
			sampling.POST("/plan/create", r.qmsSamplingHandler.CreatePlan)
			sampling.PUT("/plan/update", r.qmsSamplingHandler.UpdatePlan)
			sampling.DELETE("/plan/:id", r.qmsSamplingHandler.DeletePlan)
			sampling.GET("/plan/list", r.qmsSamplingHandler.ListPlan)
			sampling.GET("/plan/:id", r.qmsSamplingHandler.GetPlan)
			sampling.PUT("/plan/:id/rules", r.qmsSamplingHandler.UpdateRules)
			sampling.GET("/calculate", r.qmsSamplingHandler.Calculate)
			sampling.POST("/record", r.qmsSamplingHandler.CreateRecord)
			sampling.GET("/record/list", r.qmsSamplingHandler.ListRecord)
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

			// 采购退货
			purchaseReturn := wms.Group("/purchase-return")
			{
				purchaseReturn.GET("/list", r.purchaseReturnHandler.ListPurchaseReturns)
				purchaseReturn.GET("/:id", r.purchaseReturnHandler.GetPurchaseReturn)
				purchaseReturn.POST("", r.purchaseReturnHandler.CreatePurchaseReturn)
				purchaseReturn.PUT("/:id", r.purchaseReturnHandler.UpdatePurchaseReturn)
				purchaseReturn.DELETE("/:id", r.purchaseReturnHandler.DeletePurchaseReturn)
				purchaseReturn.POST("/:id/submit", r.purchaseReturnHandler.SubmitPurchaseReturn)
				purchaseReturn.POST("/:id/approve", r.purchaseReturnHandler.ApprovePurchaseReturn)
				purchaseReturn.POST("/:id/start-return", r.purchaseReturnHandler.StartReturnPurchaseReturn)
				purchaseReturn.POST("/:id/confirm", r.purchaseReturnHandler.ConfirmPurchaseReturn)
				purchaseReturn.POST("/:id/cancel", r.purchaseReturnHandler.CancelPurchaseReturn)
			}

			// 销售退货
			salesReturn := wms.Group("/sales-return")
			{
				salesReturn.GET("/list", r.salesReturnHandler.ListSalesReturns)
				salesReturn.GET("/:id", r.salesReturnHandler.GetSalesReturn)
				salesReturn.POST("", r.salesReturnHandler.CreateSalesReturn)
				salesReturn.PUT("/:id", r.salesReturnHandler.UpdateSalesReturn)
				salesReturn.DELETE("/:id", r.salesReturnHandler.DeleteSalesReturn)
				salesReturn.POST("/:id/submit", r.salesReturnHandler.SubmitSalesReturn)
				salesReturn.POST("/:id/approve", r.salesReturnHandler.ApproveSalesReturn)
				salesReturn.POST("/:id/start-return", r.salesReturnHandler.StartReturnSalesReturn)
				salesReturn.POST("/:id/confirm", r.salesReturnHandler.ConfirmSalesReturn)
				salesReturn.POST("/:id/cancel", r.salesReturnHandler.CancelSalesReturn)
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

		// 设备部件
		equipmentPart := protected.Group("/equipment/part")
		{
			equipmentPart.GET("/list", r.equipmentPartHandler.List)
			equipmentPart.GET("/:id", r.equipmentPartHandler.Get)
			equipmentPart.POST("", r.equipmentPartHandler.Create)
			equipmentPart.PUT("/:id", r.equipmentPartHandler.Update)
			equipmentPart.DELETE("/:id", r.equipmentPartHandler.Delete)
			equipmentPart.GET("/equipment/:equipment_id", r.equipmentPartHandler.ListByEquipment)
		}

		// 设备文档
		equipmentDocument := protected.Group("/equipment/document")
		{
			equipmentDocument.GET("/list", r.equipmentDocumentHandler.List)
			equipmentDocument.GET("/:id", r.equipmentDocumentHandler.Get)
			equipmentDocument.POST("", r.equipmentDocumentHandler.Create)
			equipmentDocument.PUT("/:id", r.equipmentDocumentHandler.Update)
			equipmentDocument.DELETE("/:id", r.equipmentDocumentHandler.Delete)
			equipmentDocument.GET("/equipment/:equipment_id", r.equipmentDocumentHandler.ListByEquipment)
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

		// SCP供应链供应商（别名，供前端调用）
		scpSupplier := protected.Group("/scp/supplier")
		{
			scpSupplier.GET("/list", r.supplierHandler.List)
			scpSupplier.GET("/:id", r.supplierHandler.Get)
		}

		// 供应商ASN管理
		asn := protected.Group("/supplier/asn")
		{
			asn.GET("/list", r.supplierASNHandler.List)
			asn.GET("/:id", r.supplierASNHandler.Get)
			asn.GET("/no/:asnNo", r.supplierASNHandler.GetByNo)
			asn.POST("", r.supplierASNHandler.Create)
			asn.PUT("/:id", r.supplierASNHandler.Update)
			asn.DELETE("/:id", r.supplierASNHandler.Delete)
			asn.PUT("/:id/submit", r.supplierASNHandler.Submit)
			asn.PUT("/:id/confirm", r.supplierASNHandler.Confirm)
			asn.PUT("/:id/start-receiving", r.supplierASNHandler.StartReceiving)
			asn.PUT("/:id/complete-receiving", r.supplierASNHandler.CompleteReceiving)
			asn.PUT("/:id/cancel", r.supplierASNHandler.Cancel)
			asn.POST("/:id/items", r.supplierASNHandler.AddItem)
		}

		// 供应商物料关联
		supplierMaterial := protected.Group("/mdm/supplier-material")
		{
			supplierMaterial.GET("/list", r.supplierMaterialHandler.List)
			supplierMaterial.GET("/:id", r.supplierMaterialHandler.Get)
			supplierMaterial.POST("", r.supplierMaterialHandler.Create)
			supplierMaterial.PUT("/:id", r.supplierMaterialHandler.Update)
			supplierMaterial.DELETE("/:id", r.supplierMaterialHandler.Delete)
			supplierMaterial.GET("/supplier/:supplier_id", r.supplierMaterialHandler.ListBySupplier)
			supplierMaterial.GET("/material/:material_id", r.supplierMaterialHandler.ListByMaterial)
			supplierMaterial.POST("/preferred", r.supplierMaterialHandler.SetPreferred)
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

		// 器具生命周期
		containerLifecycle := protected.Group("/containers/lifecycle")
		{
			containerLifecycle.GET("/list", r.containerLifecycleHandler.ListContainerLifecycles)
			containerLifecycle.GET("/:id", r.containerLifecycleHandler.GetContainerLifecycle)
			containerLifecycle.POST("/:id/init", r.containerLifecycleHandler.InitializeContainer)
			containerLifecycle.POST("/:id/maintain", r.containerLifecycleHandler.RecordMaintenance)
			containerLifecycle.POST("/:id/complete-maintain", r.containerLifecycleHandler.CompleteMaintenance)
			containerLifecycle.POST("/:id/retire", r.containerLifecycleHandler.RetireContainer)
			containerLifecycle.GET("/timeline/:id", r.containerLifecycleHandler.GetContainerTimeline)
			containerLifecycle.GET("/maintenances/:id", r.containerLifecycleHandler.ListContainerMaintenances)
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

			// AI视觉检测
			visualInspection := ai.Group("/visual-inspection")
			{
				visualInspection.GET("/list", r.visualInspectionHandler.ListVisualInspectionTasks)
				visualInspection.GET("/:id", r.visualInspectionHandler.GetVisualInspectionTask)
				visualInspection.POST("", r.visualInspectionHandler.CreateVisualInspectionTask)
				visualInspection.DELETE("/:id", r.visualInspectionHandler.DeleteVisualInspectionTask)
				visualInspection.GET("/:id/result", r.visualInspectionHandler.GetVisualInspectionResult)
				visualInspection.POST("/:id/manual-review", r.visualInspectionHandler.ManualReview)
				visualInspection.GET("/stats", r.visualInspectionHandler.GetVisualInspectionStats)
			}
		}

		// 财务结算
		fin := protected.Group("/fin")
		{
			// 采购结算
			purchaseSettlement := fin.Group("/purchase-settlement")
			{
				purchaseSettlement.GET("/list", r.finHandler.ListPurchaseSettlements)
				purchaseSettlement.GET("/:id", r.finHandler.GetPurchaseSettlement)
				purchaseSettlement.POST("", r.finHandler.CreatePurchaseSettlement)
				purchaseSettlement.PUT("/:id/submit", r.finHandler.SubmitPurchaseSettlement)
				purchaseSettlement.PUT("/:id/approve", r.finHandler.ApprovePurchaseSettlement)
				purchaseSettlement.PUT("/:id/cancel", r.finHandler.CancelPurchaseSettlement)
				purchaseSettlement.DELETE("/:id", r.finHandler.DeletePurchaseSettlement)
			}

			// 销售结算
			salesSettlement := fin.Group("/sales-settlement")
			{
				salesSettlement.GET("/list", r.finHandler.ListSalesSettlements)
				salesSettlement.GET("/:id", r.finHandler.GetSalesSettlement)
				salesSettlement.POST("", r.finHandler.CreateSalesSettlement)
				salesSettlement.PUT("/:id/submit", r.finHandler.SubmitSalesSettlement)
				salesSettlement.PUT("/:id/approve", r.finHandler.ApproveSalesSettlement)
				salesSettlement.PUT("/:id/cancel", r.finHandler.CancelSalesSettlement)
				salesSettlement.DELETE("/:id", r.finHandler.DeleteSalesSettlement)
			}

			// 付款申请
			paymentRequest := fin.Group("/payment-request")
			{
				paymentRequest.GET("/list", r.finHandler.ListPaymentRequests)
				paymentRequest.GET("/:id", r.finHandler.GetPaymentRequest)
				paymentRequest.POST("", r.finHandler.CreatePaymentRequest)
				paymentRequest.PUT("/:id/submit", r.finHandler.SubmitPaymentRequest)
				paymentRequest.PUT("/:id/approve", r.finHandler.ApprovePaymentRequest)
				paymentRequest.PUT("/:id/reject", r.finHandler.RejectPaymentRequest)
				paymentRequest.PUT("/:id/pay", r.finHandler.PayPaymentRequest)
				paymentRequest.DELETE("/:id", r.finHandler.DeletePaymentRequest)
			}

			// 采购预付款
			purchaseAdvance := fin.Group("/purchase-advance")
			{
				purchaseAdvance.GET("/list", r.finHandler.ListPurchaseAdvances)
				purchaseAdvance.POST("", r.finHandler.CreatePurchaseAdvance)
			}

			// 销售收款
			salesReceipt := fin.Group("/sales-receipt")
			{
				salesReceipt.GET("/list", r.finHandler.ListSalesReceipts)
				salesReceipt.POST("", r.finHandler.CreateSalesReceipt)
			}

			// 供应商对账
			supplierStatement := fin.Group("/supplier-statement")
			{
				supplierStatement.GET("/list", r.finHandler.ListSupplierStatements)
				supplierStatement.GET("/:id", r.finHandler.GetSupplierStatement)
			}

			// ========== SCP 采购管理 ==========
			// RFQ
			rfq := protected.Group("/scp/rfq")
			{
				rfq.GET("/list", r.rfqHandler.List)
				rfq.GET("/:id", r.rfqHandler.Get)
				rfq.POST("", r.rfqHandler.Create)
				rfq.PUT("/:id", r.rfqHandler.Update)
				rfq.DELETE("/:id", r.rfqHandler.Delete)
				rfq.POST("/:id/publish", r.rfqHandler.Publish)
				rfq.POST("/:id/close", r.rfqHandler.Close)
				rfq.GET("/:id/quotes", r.rfqHandler.GetQuotes)
				rfq.POST("/:id/award", r.rfqHandler.Award)
			}

			// 采购订单
			purchaseOrder := protected.Group("/scp/purchase-orders")
			{
				purchaseOrder.GET("/list", r.purchaseOrderHandler.List)
				purchaseOrder.GET("/:id", r.purchaseOrderHandler.Get)
				purchaseOrder.POST("", r.purchaseOrderHandler.Create)
				purchaseOrder.PUT("/:id", r.purchaseOrderHandler.Update)
				purchaseOrder.DELETE("/:id", r.purchaseOrderHandler.Delete)
				purchaseOrder.POST("/:id/submit", r.purchaseOrderHandler.Submit)
				purchaseOrder.POST("/:id/approve", r.purchaseOrderHandler.Approve)
				purchaseOrder.POST("/:id/reject", r.purchaseOrderHandler.Reject)
				purchaseOrder.POST("/:id/issue", r.purchaseOrderHandler.Issue)
				purchaseOrder.POST("/:id/close", r.purchaseOrderHandler.Close)
				purchaseOrder.POST("/:id/cancel", r.purchaseOrderHandler.Cancel)
				purchaseOrder.POST("/:id/receive", r.purchaseOrderHandler.Receive)
			}

			// 销售订单
			salesOrder := protected.Group("/scp/sales-orders")
			{
				salesOrder.GET("/list", r.scpSalesOrderHandler.List)
				salesOrder.GET("/:id", r.scpSalesOrderHandler.Get)
				salesOrder.POST("", r.scpSalesOrderHandler.Create)
				salesOrder.PUT("/:id", r.scpSalesOrderHandler.Update)
				salesOrder.DELETE("/:id", r.scpSalesOrderHandler.Delete)
				salesOrder.POST("/:id/submit", r.scpSalesOrderHandler.Submit)
				salesOrder.POST("/:id/approve", r.scpSalesOrderHandler.Approve)
				salesOrder.POST("/:id/reject", r.scpSalesOrderHandler.Reject)
				salesOrder.POST("/:id/confirm", r.scpSalesOrderHandler.Confirm)
				salesOrder.POST("/:id/close", r.scpSalesOrderHandler.Close)
				salesOrder.POST("/:id/cancel", r.scpSalesOrderHandler.Cancel)
			}

			// 供应商KPI
			supplierKPI := protected.Group("/scp/supplier-kpi")
			{
				supplierKPI.GET("/list", r.supplierKPIHandler.List)
				supplierKPI.GET("/monthly", r.supplierKPIHandler.GetByMonthly)
				supplierKPI.POST("", r.supplierKPIHandler.Create)
				supplierKPI.GET("/ranking", r.supplierKPIHandler.GetRanking)
			}

			// 供应商报价
			supplierQuote := protected.Group("/scp/supplier-quotes")
			{
				supplierQuote.GET("/list", r.supplierQuoteHandler.List)
				supplierQuote.GET("/:id", r.supplierQuoteHandler.Get)
				supplierQuote.POST("", r.supplierQuoteHandler.Create)
				supplierQuote.GET("/rfq/:rfqId/quotes", r.supplierQuoteHandler.GetQuotes)
				supplierQuote.POST("/rfq/:rfqId/award", r.supplierQuoteHandler.Award)
			}

			// 客户询价
			customerInquiry := protected.Group("/scp/customer-inquiry")
			{
				customerInquiry.GET("/list", r.customerInquiryHandler.List)
				customerInquiry.GET("/:id", r.customerInquiryHandler.Get)
				customerInquiry.POST("", r.customerInquiryHandler.Create)
				customerInquiry.PUT("/:id", r.customerInquiryHandler.Update)
				customerInquiry.DELETE("/:id", r.customerInquiryHandler.Delete)
				customerInquiry.POST("/:id/send", r.customerInquiryHandler.Send)
				customerInquiry.POST("/:id/quote", r.customerInquiryHandler.Quote)
				customerInquiry.POST("/:id/win", r.customerInquiryHandler.Win)
				customerInquiry.POST("/:id/lose", r.customerInquiryHandler.Lose)
				customerInquiry.POST("/:id/cancel", r.customerInquiryHandler.Cancel)
			}

			// 采购计划
			purchasePlan := protected.Group("/scp/purchase-plan")
			{
				purchasePlan.GET("/list", r.purchasePlanHandler.List)
				purchasePlan.GET("/:id", r.purchasePlanHandler.Get)
				purchasePlan.POST("/create", r.purchasePlanHandler.Create)
				purchasePlan.PUT("/update", r.purchasePlanHandler.Update)
				purchasePlan.DELETE("/delete", r.purchasePlanHandler.Delete)
				purchasePlan.GET("/:id/items", r.purchasePlanHandler.GetItems)
				purchasePlan.POST("/:id/confirm", r.purchasePlanHandler.Confirm)
				purchasePlan.POST("/:id/publish", r.purchasePlanHandler.Publish)
				purchasePlan.POST("/:id/close", r.purchasePlanHandler.Close)
			}

			// 供应商联系人
			supplierContact := protected.Group("/scp/supplier-contact")
			{
				supplierContact.GET("/list", r.scpSupplierExtHandler.ListContacts)
				supplierContact.GET("/:id", r.scpSupplierExtHandler.GetContact)
				supplierContact.POST("/create", r.scpSupplierExtHandler.CreateContact)
				supplierContact.PUT("/update", r.scpSupplierExtHandler.UpdateContact)
				supplierContact.DELETE("/delete", r.scpSupplierExtHandler.DeleteContact)
				supplierContact.GET("/supplier/:supplierId", r.scpSupplierExtHandler.ListContactsBySupplier)
			}

			// 供应商银行账户
			supplierBank := protected.Group("/scp/supplier-bank")
			{
				supplierBank.GET("/list", r.scpSupplierExtHandler.ListBanks)
				supplierBank.GET("/:id", r.scpSupplierExtHandler.GetBank)
				supplierBank.POST("/create", r.scpSupplierExtHandler.CreateBank)
				supplierBank.PUT("/update", r.scpSupplierExtHandler.UpdateBank)
				supplierBank.DELETE("/delete", r.scpSupplierExtHandler.DeleteBank)
				supplierBank.GET("/supplier/:supplierId", r.scpSupplierExtHandler.ListBanksBySupplier)
			}

			// ========== 设备停机 (EAM) ==========
			downtime := protected.Group("/eam/downtime")
			{
				downtime.GET("/list", r.equipmentDowntimeHandler.List)
				downtime.GET("/:id", r.equipmentDowntimeHandler.Get)
				downtime.POST("", r.equipmentDowntimeHandler.Create)
				downtime.PUT("/:id", r.equipmentDowntimeHandler.Update)
				downtime.DELETE("/:id", r.equipmentDowntimeHandler.Delete)
				downtime.POST("/:id/start", r.equipmentDowntimeHandler.StartDowntime)
				downtime.POST("/:id/end", r.equipmentDowntimeHandler.EndDowntime)
			}

			// ========== 备件管理 (EAM) ==========
			spare := protected.Group("/eam/spare")
			{
				spare.GET("/list", r.spareHandler.List)
				spare.GET("/:id", r.spareHandler.Get)
				spare.POST("", r.spareHandler.Create)
				spare.PUT("", r.spareHandler.Update)
				spare.DELETE("/:id", r.spareHandler.Delete)
				spare.POST("/input", r.spareHandler.Input)
				spare.POST("/output", r.spareHandler.Output)
				spare.GET("/transactions", r.spareHandler.Transactions)
			}

			// ========== Alert 告警管理 ==========
			alert := protected.Group("/alert")
			{
				alert.GET("/rule/list", r.alertHandler.ListRules)
				alert.GET("/rule/:id", r.alertHandler.GetRule)
				alert.POST("/rule", r.alertHandler.CreateRule)
				alert.PUT("/rule/:id", r.alertHandler.UpdateRule)
				alert.DELETE("/rule/:id", r.alertHandler.DeleteRule)
				alert.POST("/rule/:id/enable", r.alertHandler.EnableRule)
				alert.POST("/rule/:id/disable", r.alertHandler.DisableRule)
				alert.GET("/record/list", r.alertHandler.ListRecords)
				alert.GET("/record/:id", r.alertHandler.GetRecord)
				alert.POST("/record/:id/ack", r.alertHandler.AcknowledgeRecord)
				alert.POST("/record/:id/resolve", r.alertHandler.ResolveRecord)
				alert.POST("/record/:id/close", r.alertHandler.CloseRecord)
				alert.GET("/statistics", r.alertHandler.GetStatistics)
				alert.GET("/notify/logs", r.alertHandler.ListNotificationLogs)
				alert.GET("/escalation/list", r.alertHandler.ListEscalationRules)
				alert.POST("/escalation", r.alertHandler.CreateEscalationRule)
				alert.GET("/channel/list", r.alertHandler.ListChannels)
				alert.GET("/channel/:id", r.alertHandler.GetChannel)
				alert.POST("/channel", r.alertHandler.CreateChannel)
				alert.PUT("/channel/:id", r.alertHandler.UpdateChannel)
				alert.DELETE("/channel/:id", r.alertHandler.DeleteChannel)
				alert.POST("/channel/:id/enable", r.alertHandler.EnableChannel)
				alert.POST("/channel/:id/disable", r.alertHandler.DisableChannel)
				alert.POST("/send", r.alertHandler.SendNotification)
			}

			// ========== BPM 业务流程 ==========
			bpmProc := protected.Group("/bpm/process")
			{
				bpmProc.GET("/list", r.bpmHandler.ListProcessModels)
				bpmProc.GET("/:id", r.bpmHandler.GetProcessModel)
				bpmProc.POST("", r.bpmHandler.CreateProcessModel)
				bpmProc.PUT("/:id", r.bpmHandler.UpdateProcessModel)
				bpmProc.DELETE("/:id", r.bpmHandler.DeleteProcessModel)
				bpmProc.POST("/:id/publish", r.bpmHandler.PublishProcessModel)
			}
			bpmNode := protected.Group("/bpm/node")
			{
				bpmNode.GET("/list", r.bpmHandler.ListNodes)
				bpmNode.POST("", r.bpmHandler.CreateNode)
				bpmNode.PUT("/:id", r.bpmHandler.UpdateNode)
				bpmNode.DELETE("/:id", r.bpmHandler.DeleteNode)
			}
			bpmFlow := protected.Group("/bpm/flow")
			{
				bpmFlow.GET("/list", r.bpmHandler.ListFlows)
				bpmFlow.POST("", r.bpmHandler.CreateFlow)
				bpmFlow.PUT("/:id", r.bpmHandler.UpdateFlow)
				bpmFlow.DELETE("/:id", r.bpmHandler.DeleteFlow)
			}
			bpmForm := protected.Group("/bpm/form")
			{
				bpmForm.GET("/list", r.bpmHandler.ListFormDefinitions)
				bpmForm.GET("/:id", r.bpmHandler.GetFormDefinition)
				bpmForm.POST("", r.bpmHandler.CreateFormDefinition)
				bpmForm.PUT("/:id", r.bpmHandler.UpdateFormDefinition)
				bpmForm.DELETE("/:id", r.bpmHandler.DeleteFormDefinition)
			}
			bpmField := protected.Group("/bpm/field")
			{
				bpmField.GET("/list", r.bpmHandler.ListFormFields)
				bpmField.POST("", r.bpmHandler.CreateFormField)
				bpmField.PUT("/:id", r.bpmHandler.UpdateFormField)
				bpmField.DELETE("/:id", r.bpmHandler.DeleteFormField)
			}
			bpmInstance := protected.Group("/bpm/instance")
			{
				bpmInstance.GET("/list", r.bpmHandler.ListProcessInstances)
				bpmInstance.GET("/:id", r.bpmHandler.GetProcessInstance)
				bpmInstance.POST("/start", r.bpmHandler.CreateProcessInstance)
				bpmInstance.POST("/:id/cancel", r.bpmHandler.CancelProcessInstance)
				bpmInstance.POST("/:id/terminate", r.bpmHandler.TerminateProcessInstance)
				bpmInstance.GET("/task/list", r.bpmHandler.ListTasksByAssignee)
				bpmInstance.GET("/task/:id", r.bpmHandler.GetTask)
				bpmInstance.POST("/task/:id/approve", r.bpmHandler.ApproveTask)
				bpmInstance.POST("/task/:id/reject", r.bpmHandler.RejectTask)
				bpmInstance.GET("/approve/records", r.bpmHandler.ListApprovalRecords)
			}

			// BPM 委托
			bpmDelegate := protected.Group("/bpm/delegate")
			{
				bpmDelegate.GET("/list", r.bpmHandler.ListDelegates)
				bpmDelegate.POST("", r.bpmHandler.CreateDelegate)
				bpmDelegate.PUT("/:id", r.bpmHandler.UpdateDelegate)
				bpmDelegate.DELETE("/:id", r.bpmHandler.DeleteDelegate)
			}

			// BPM 跨模块API
			instanceApi := protected.Group("/bpm/instance-api")
			{
				instanceApi.POST("/start", r.bpmInstanceApiHandler.StartProcessInstance)
				instanceApi.POST("/complete", r.bpmInstanceApiHandler.CompleteTask)
				instanceApi.GET("/:id", r.bpmInstanceApiHandler.GetProcessInstance)
			}

			// BPM 任务转移/候选人
			bpmTask := protected.Group("/bpm/task")
			{
				bpmTask.POST("/transfer", r.bpmTaskTransferHandler.TransferTask)
				bpmTask.GET("/transfer/history/:taskId", r.bpmTaskTransferHandler.GetTransferHistory)
				bpmTask.GET("/candidate/:taskId", r.bpmTaskTransferHandler.GetTaskCandidates)
				bpmTask.GET("/candidate-group/:taskId", r.bpmTaskTransferHandler.GetTaskCandidateGroups)
				bpmTask.POST("/assign", r.bpmTaskTransferHandler.AssignTask)
			}

			// ========== MDM 合作伙伴扩展 ==========
			contact := protected.Group("/mdm/contact")
			{
				contact.GET("/list", r.contactHandler.List)
				contact.GET("/:id", r.contactHandler.Get)
				contact.POST("", r.contactHandler.Create)
				contact.PUT("/:id", r.contactHandler.Update)
				contact.DELETE("/:id", r.contactHandler.Delete)
			}
			bankAccount := protected.Group("/mdm/bank-account")
			{
				bankAccount.GET("/list", r.bankAccountHandler.List)
				bankAccount.GET("/:id", r.bankAccountHandler.Get)
				bankAccount.POST("", r.bankAccountHandler.Create)
				bankAccount.PUT("/:id", r.bankAccountHandler.Update)
				bankAccount.DELETE("/:id", r.bankAccountHandler.Delete)
			}
			attachment := protected.Group("/mdm/attachment")
			{
				attachment.GET("/list", r.attachmentHandler.List)
				attachment.GET("/:id", r.attachmentHandler.Get)
				attachment.POST("", r.attachmentHandler.Create)
				attachment.PUT("/:id", r.attachmentHandler.Update)
				attachment.DELETE("/:id", r.attachmentHandler.Delete)
			}
		}
	}
}

// SetJWT 设置JWT中间件
func (r *Router) SetJWT(jwtFunc func() gin.HandlerFunc) {
	protected := r.engine.Group("/api/v1")
	protected.Use(jwtFunc())
}
