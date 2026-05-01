package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mom-server/internal/config"
	"mom-server/internal/handler/andon"
	"mom-server/internal/handler/aps"
	"mom-server/internal/handler/business"
	"mom-server/internal/handler/bpm"
	"mom-server/internal/handler/dc"
	"mom-server/internal/handler/equipment"
	"mom-server/internal/handler/eam"
	"mom-server/internal/handler/fin"
	"mom-server/internal/handler/mdm"
	"mom-server/internal/handler/production"
	"mom-server/internal/handler/quality"
	"mom-server/internal/handler/scp"
	"mom-server/internal/handler/supplier"
	"mom-server/internal/handler/supplier_asn"
	"mom-server/internal/handler/system"
	"mom-server/internal/handler/trace"
	"mom-server/internal/handler/wms"
	"mom-server/internal/handler/ai"
	"mom-server/internal/handler/container"
	"mom-server/internal/handler/alert"
	"mom-server/internal/handler/mes"
	"mom-server/internal/handler/agv"
	"mom-server/internal/handler/erp_sync"
	"mom-server/internal/handler/integration"
	"mom-server/internal/handler/report"
	"mom-server/internal/model"
	"mom-server/internal/pkg/jwt"
	"mom-server/internal/repository"
	"mom-server/internal/router"
	"mom-server/internal/service"
)

// migrateBatch wraps GORM AutoMigrate with self-healing logic for known PostgreSQL
// migration errors: duplicate key violations (23505) and missing PK on FK refs (42830).
func migrateBatch(db *gorm.DB, name string, models ...interface{}) {
	var lastErr error
	for attempt := 1; attempt <= 3; attempt++ {
		lastErr = nil
		err := db.AutoMigrate(models...)
		if err == nil {
			return
		}

		errMsg := err.Error()

		// Type 1: unique constraint violation (23505) — could be duplicate rows OR pre-existing index
		if strings.Contains(errMsg, "23505") {
			// Check if this is a "could not create index" error (index already exists)
			idxNameMatch := regexp.MustCompile(`could not create (unique )?index "(\w+)"`).FindStringSubmatch(errMsg)
			if idxNameMatch != nil {
				idxName := idxNameMatch[2]
				var exists int
				db.Raw("SELECT 1 FROM pg_indexes WHERE indexname = ?", idxName).Scan(&exists)
				if exists == 1 {
					log.Printf("  ✓ 索引 %s 已存在，跳过", idxName)
					return // index exists, migration is effectively complete
				}
			}
			// Otherwise it's a row-level duplicate key — try to deduplicate
			log.Printf("  ⚠ %s: 唯一约束冲突，尝试清理重复行 (%d/%d)", name, attempt, 3)
			idxMatch := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`).FindStringSubmatch(errMsg)
			if len(idxMatch) == 3 {
				colVal := idxMatch[2]
				idxName := regexp.MustCompile(`Key \((\w+)\)=`).FindString(errMsg)
				if idxName != "" {
					idxName = strings.TrimSuffix(strings.TrimPrefix(idxName, "Key ("), "=")
					var idxDef string
					db.Raw("SELECT indexdef FROM pg_indexes WHERE indexname = ?", idxName).Scan(&idxDef)
					m := regexp.MustCompile(`ON (\w+) USING`).FindStringSubmatch(idxDef)
					if len(m) == 2 {
						tbl := m[1]
						var col string
						db.Raw("SELECT column_name FROM information_schema.key_column_usage WHERE constraint_name = ?", idxName).Scan(&col)
						db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE ctid NOT IN (SELECT MIN(ctid) FROM %s WHERE %s = $1 GROUP BY %s HAVING COUNT(*) > 1)`,
							tbl, tbl, col, col), colVal)
						log.Printf("  ✓ 清理重复行: %s.%s = %s", tbl, col, colVal)
						continue
					}
				}
			}
		}

		// Type 2: foreign key constraint violation (42830) — referenced table missing PK
		if strings.Contains(errMsg, "42830") {
			log.Printf("  ⚠ %s: 外键引用表缺主键，尝试修复 (%d/%d)", name, attempt, 3)
			m := regexp.MustCompile(`(\w+)\s+FOREIGN KEY\s+\((\w+)\)\s+REFERENCES\s+(\w+)`).FindStringSubmatch(errMsg)
			if len(m) == 4 {
				refTable := m[3]
				// Check if PK already exists before trying to add
				var pkExists int
				db.Raw("SELECT 1 FROM information_schema.table_constraints WHERE table_name = ? AND constraint_type = 'PRIMARY KEY'", refTable).Scan(&pkExists)
				if pkExists == 0 {
					db.Exec(fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY (id)", refTable))
					log.Printf("  ✓ 添加主键: %s.id", refTable)
				} else {
					// PK exists but no unique constraint — GORM may need explicit unique constraint on the PK column
					db.Exec(fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s_id_unique UNIQUE (id)", refTable, refTable))
					log.Printf("  ✓ 添加唯一约束: %s.id (for FK)", refTable)
				}
				continue
			}
		}

		log.Printf("  ⚠ %s: %v", name, err)
		lastErr = err
	}
	if lastErr != nil {
		if config.GlobalConfig != nil && config.GlobalConfig.Server.Env == "production" {
			log.Fatalf("  ✗ %s 迁移失败 (3次重试后): %v", name, lastErr)
		}
		log.Printf("  ✗ %s 迁移失败 (3次重试后): %v", name, lastErr)
	}
}

// preFlightFixFK scans information_schema for all FK constraints in the database,
// finds the referenced table/column, and ensures a UNIQUE constraint exists on it.
// GORM's PostgreSQL driver requires UNIQUE (not just PK) when adding FK constraints.
// Also ensures all tables have primary keys (needed because backup SQL may lack PKs).
func preFlightFixFK(db *gorm.DB) {
	// Step 1: ensure every table has a primary key on (id)
	log.Println("  [pre-flight] 检查并修复缺主键的表...")
	var tables []string
	db.Raw(`SELECT table_name FROM information_schema.tables
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE'`).Scan(&tables)

	pkFixed := 0
	for _, tbl := range tables {
		// Skip junction/m2m tables that don't have an id column
		var idExists int
		db.Raw(`SELECT 1 FROM information_schema.columns WHERE table_name = $1 AND column_name = 'id' LIMIT 1`, tbl).Scan(&idExists)
		if idExists == 0 {
			continue
		}
		var pkCount int
		db.Raw(`SELECT COUNT(*) FROM information_schema.table_constraints
			WHERE table_name = $1 AND table_schema = 'public' AND constraint_type = 'PRIMARY KEY'`, tbl).Scan(&pkCount)
		if pkCount == 0 {
			db.Exec(fmt.Sprintf(`ALTER TABLE %s ADD PRIMARY KEY (id)`, tbl))
			log.Printf("    ✓ 添加主键: %s(id)", tbl)
			pkFixed++
		}
	}
	if pkFixed > 0 {
		log.Printf("  [pre-flight] 修复了 %d 个缺主键的表", pkFixed)
	} else {
		log.Printf("  [pre-flight] 所有表已有主键")
	}

	// Step 2: ensure referenced columns in FK relationships have UNIQUE constraints
	log.Println("  [pre-flight] 修复外键引用的唯一约束...")
	rows, err := db.Raw(`
		SELECT kcu.table_name, kcu.column_name, kcu.constraint_name,
		       ccu.table_name AS referenced_table, ccu.column_name AS referenced_column
		FROM information_schema.key_column_usage kcu
		JOIN information_schema.constraint_column_usage ccu
		  ON ccu.constraint_name = kcu.constraint_name
		JOIN information_schema.table_constraints tc
		  ON tc.constraint_name = kcu.constraint_name
		WHERE tc.constraint_type = 'FOREIGN KEY'
		  AND kcu.table_schema = 'public'
		  AND ccu.table_schema = 'public'
	`).Rows()
	if err != nil {
		log.Printf("  pre-flight FK scan failed: %v", err)
		return
	}
	defer rows.Close()

	fixed := 0
	for rows.Next() {
		var tbl, col, fkName, refTbl, refCol string
		rows.Scan(&tbl, &col, &fkName, &refTbl, &refCol)

		// Check if a UNIQUE or PK constraint already exists on referenced_column in refTbl
		var hasUnique int
		db.Raw(`
			SELECT 1 FROM information_schema.table_constraints tc
			JOIN information_schema.key_column_usage kcu
			  ON tc.constraint_name = kcu.constraint_name
			WHERE tc.table_name = ? AND tc.table_schema = 'public'
			  AND tc.constraint_type IN ('PRIMARY KEY','UNIQUE')
			  AND kcu.column_name = ?
			LIMIT 1
		`, refTbl, refCol).Scan(&hasUnique)

		if hasUnique == 0 {
			constraintName := fmt.Sprintf("preflight_uk_%s_%s", refTbl, refCol)
			db.Exec(fmt.Sprintf(
				"ALTER TABLE %s ADD CONSTRAINT %s UNIQUE (%s)",
				refTbl, constraintName, refCol,
			))
			log.Printf("  ✓ [pre-flight] 添加唯一约束: %s(%s) → %s.%s",
				constraintName, refCol, refTbl, refCol)
			fixed++
		}
	}
	if fixed > 0 {
		log.Printf("  pre-flight 修复了 %d 个唯一约束", fixed)
	} else {
		log.Printf("  pre-flight: 所有外键引用已满足唯一约束")
	}
}

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
	// Pre-flight: scan all tables that GORM will try to add FK constraints to,
	// and ensure the referenced columns have UNIQUE constraints (GORM PostgreSQL
	// driver requires unique, not just PK, on the referenced column).
	log.Println("Pre-flight: 修复所有外键引用的唯一约束...")
	preFlightFixFK(db)
	// Pre-flight 完成，开始正式迁移...
	log.Println("开始数据库迁移...")
	
	// 第1批：系统基础表（14个）
	log.Println("迁移第1批：系统基础表")
	migrateBatch(db, "第1批-系统基础表",
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
	)
	
	// 第2批：仓储管理表（7个）
	log.Println("迁移第2批：仓储管理表")
	// 修复：wms_transfer_order.status 从 bigint 改为 varchar(20)
	db.Exec("ALTER TABLE wms_transfer_order ALTER COLUMN status TYPE varchar(20) USING status::varchar(20)")
	migrateBatch(db, "第2批-仓储管理表",
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
		&model.SideLocation{},
		&model.KanbanPull{},
	)
	
	// 第3批：生产执行表（6个）
	log.Println("迁移第3批：生产执行表")
	migrateBatch(db, "第3批-生产执行表",
		&model.SalesOrder{},
		&model.SalesOrderItem{},
		&model.ProductionReport{},
		&model.Dispatch{},
		&model.ProductionOrder{},
	)

	// 第4批：APS计划表（8个）
	log.Println("迁移第4批：APS计划表")
	migrateBatch(db, "第4批-APS计划表",
		&model.MPS{},
		&model.MRP{},
		&model.MRPItem{},
		&model.SchedulePlan{},
		&model.ScheduleResult{},
		&model.Resource{},
		&model.WorkCenter{},
		&model.CapacityAnalysis{},
		&model.DeliveryRate{},
		&model.ChangeoverMatrix{},
		&model.RollingSchedule{},
		&model.JITDemand{},
		&model.WorkingCalendar{},
	)
	
	// 第5批：追溯管理表（5个）
	log.Println("迁移第5批：追溯管理表")
	migrateBatch(db, "第5批-追溯管理表",
		&model.SerialNumber{},
		&model.TraceRecord{},
		&model.AndonCall{},
		&model.DataCollection{},
		&model.EnergyRecord{},
	)

	// 第5.5批：安灯升级机制表
	log.Println("迁移第5.5批：安灯升级机制表")
	migrateBatch(db, "第5.5批-安灯升级机制表",
		&model.AndonEscalationRule{},
		&model.AndonEscalationLog{},
		&model.AndonNotificationLog{},
	)
	
	// 第6批：主数据管理表（跳过已有数据的表避免迁移错误）
	log.Println("迁移第6批：主数据管理表")
	migrateBatch(db, "第6批-主数据管理表",
		&model.Material{},
		&model.MaterialCategory{},
		// BOM/BOMItem/Process/Route/RouteOperation等表已有数据,跳过迁移
		&model.Workshop{},
		&model.ProductionLine{},
		&model.Workstation{},
		&model.Shift{},
		&model.MdmShift{},
		&model.Supplier{},
		&model.SupplierMaterial{},
		&model.WorkshopConfig{},
		&model.Contact{},
		&model.BankAccount{},
		&model.Attachment{},
	)

	// 第7批：质量管理表
	log.Println("迁移第7批：质量管理表")
	migrateBatch(db, "第7批-质量管理表",
		&model.IQC{},
		&model.IQCItem{},
		&model.IPQC{},
		&model.FQC{},
		&model.OQC{},
		&model.DefectCode{},
		&model.DefectRecord{},
		&model.NCR{},
		&model.SPCData{},
		&model.QRCI{},
	)

	// 第8批：设备OEE表
	log.Println("迁移第8批：设备OEE表")
	migrateBatch(db, "第8批-设备OEE表",
		&model.OEE{},
		&model.OEEEvent{},
		&model.TEEPData{},
		&model.Mold{},
		&model.MoldMaintenance{},
		&model.MoldRepair{},
		&model.Gauge{},
		&model.GaugeCalibration{},
	)

	// 第9批：首末件检验表
	log.Println("迁移第9批：首末件检验表")
	migrateBatch(db, "第9批-首末件检验表",
		&model.MesFirstLastInspect{},
	)

	// 第10批：包装条码表
	log.Println("迁移第10批：包装条码表")
	migrateBatch(db, "第10批-包装条码表",
		&model.MesPackage{},
	)

	// 第11批：数据采集表
	log.Println("迁移第11批：数据采集表")
	migrateBatch(db, "第11批-数据采集表",
		&model.DCDataPoint{},
		&model.DCCollectRecord{},
		&model.DCScanLog{},
	)

	// 第12批：电子SOP、编码规则、流程卡
	log.Println("迁移第12批：电子SOP、编码规则、流程卡")
	migrateBatch(db, "第12批",
		&model.ElectronicSOP{},
		&model.CodeRule{},
		&model.CodeRuleRecord{},
		&model.FlowCard{},
		&model.FlowCardDetail{},
		&model.PrintTemplate{},
		&model.Notice{},
		&model.NoticeReadRecord{},
	)

	// 第13批：器具管理表
	log.Println("迁移第13批：器具管理表")
	migrateBatch(db, "第13批-器具管理表",
		&model.ContainerMaster{},
		&model.ContainerMovement{},
	)

	// 第14批：AI聊天表 + 设备部件文档表
	log.Println("迁移第14批：AI聊天表 + 设备部件文档表")
	migrateBatch(db, "第14批",
		&model.AIConfig{},
		&model.ChatConversation{},
		&model.ChatMessage{},
		&model.EquipmentPart{},
		&model.EquipmentDocument{},
	)

	// 第28批：实验室检测管理表
	log.Println("迁移第28批：实验室检测管理表")
	migrateBatch(db, "第28批-实验室检测管理表",
		&model.LabSample{},
		&model.LabTestItem{},
		&model.LabReport{},
	)

	// 第29批：MES生产执行扩展表（班组/工艺路线/发料/退料/离线）
	log.Println("迁移第29批：MES生产执行扩展表")
	migrateBatch(db, "第29批-MES生产执行扩展表",
		&model.MesTeam{},
		&model.MesTeamMember{},
		&model.MesTeamShift{},
		&model.MesProcess{},
		&model.MesProcessOperation{},
		&model.ProductionIssue{},
		&model.ProductionIssueItem{},
		&model.ProductionReturn{},
		&model.ProductionReturnItem{},
		&model.ProductionOffline{},
		&model.ProductionOfflineItem{},
		&model.ProductionComplete{},
		&model.ProductionCompleteItem{},
		&model.ProductionStockIn{},
		&model.ProductionStockInItem{},
	)

	// 第30批：WMS采购退货和销售退货表
	log.Println("迁移第30批：WMS采购退货和销售退货表")
	migrateBatch(db, "第30批-WMS采购退货和销售退货表",
		&model.PurchaseReturn{},
		&model.PurchaseReturnItem{},
		&model.SalesReturn{},
		&model.SalesReturnItem{},
	)

	// 第31批：检验特性管理表
	log.Println("迁移第31批：检验特性管理表")
	migrateBatch(db, "第31批-检验特性管理表",
		&model.InspectionFeature{},
	)

	// 第32批：实验室仪器管理表
	log.Println("迁移第32批：实验室仪器管理表")
	migrateBatch(db, "第32批-实验室仪器管理表",
		&model.LabInstrument{},
		&model.LabCalibration{},
		&model.InspectionCharacteristic{},
		&model.AQLLevel{},
		&model.AQLTableRow{},
		&model.SamplingPlan{},
		&model.QMSSamplingPlan{},
		&model.QMSSamplingRule{},
		&model.QMSSamplingRecord{},
	)

	// 第33批：AI视觉检测表
	log.Println("迁移第33批：AI视觉检测表")
	migrateBatch(db, "第33批-AI视觉检测表",
		&model.VisualInspectionTask{},
		&model.VisualInspectionResult{},
	)

	// 第34批：容器生命周期管理表
	log.Println("迁移第34批：容器生命周期管理表")
	migrateBatch(db, "第34批-容器生命周期管理表",
		&model.ContainerLifecycle{},
		&model.ContainerMaintenance{},
	)

	// 第35批：人员能力矩阵表
	log.Println("迁移第35批：人员能力矩阵表")
	migrateBatch(db, "第35批-人员能力矩阵表",
		&model.PersonSkill{},
		&model.PersonSkillScore{},
	)

	// 第36批：报表表
	log.Println("迁移第36批：报表表")
	migrateBatch(db, "第36批-报表表",
		&model.ProductionDailyReport{},
		&model.QualityWeeklyReport{},
		&model.OEEReport{},
		&model.DeliveryReport{},
		&model.AndonReport{},
	)

	// 第37批：系统集成接口配置表
	log.Println("迁移第37批：系统集成接口配置表")
	migrateBatch(db, "第37批-系统集成接口配置表",
		&model.InterfaceConfig{},
		&model.InterfaceFieldMap{},
		&model.InterfaceTrigger{},
		&model.InterfaceExecutionLog{},
	)

	migrateBatch(db, "第38批-AGV表",
		&model.AGVTask{},
		&model.AGVDevice{},
		&model.AGVLocationMapping{},
	)

	migrateBatch(db, "第39批-供应商ASN表",
		&model.SupplierASN{},
		&model.SupplierASNItem{},
	)

	migrateBatch(db, "第40批-ERP同步表",
		&model.IntegrationERPSyncLog{},
		&model.IntegrationERPMapping{},
	)

	// 第41批：SCP供应链管理表
	log.Println("迁移第41批：SCP供应链管理表")
	migrateBatch(db, "第41批-SCP供应链管理表",
		&model.PurchaseOrder{},
		&model.PurchaseOrderItem{},
		&model.POChangeLog{},
		&model.RFQ{},
		&model.RFQItem{},
		&model.RFQInvite{},
		&model.SupplierQuote{},
		&model.QuoteItem{},
		&model.QuoteComparison{},
		&model.SCPSalesOrder{},
		&model.SCPSalesOrderItem{},
		&model.SOChangeLog{},
		&model.CustomerInquiry{},
		&model.InquiryItem{},
		&model.SupplierKPI{},
		&model.SupplierDeliveryRecord{},
		&model.SupplierQualityRecord{},
		&model.SupplierGradeStandard{},
		&model.SupplierPurchaseInfo{},
		&model.SupplierMaterial{},
		&model.ScpMRS{},
		&model.ScpMRSItem{},
		&model.ScpPurchasePlan{},
		&model.ScpPurchasePlanItem{},
		&model.ScpSupplierContact{},
		&model.ScpSupplierBank{},
	)

	// 第42批：EAM巡检管理表
	// 第43批：EAM维修工单/流程/标准
	log.Println("迁移第43批：EAM维修工单/流程/标准")
	migrateBatch(db, "第43批-EAM维修工单",
		&model.EamRepairJob{},
		&model.EamRepairFlow{},
		&model.EamRepairStd{},
		&model.EquipmentDowntime{},
	)

	// 第44批：MES工单排程表
	log.Println("迁移第44批：MES工单排程表")
	migrateBatch(db, "第44批-MES工单排程表",
		&model.MesWorkScheduling{},
		&model.MesWorkSchedulingDetail{},
	)

	// 第46批：MES报工记录表
	log.Println("迁移第46批：MES报工记录表")
	migrateBatch(db, "第46批-MES报工记录表",
		&model.MesJobReportLog{},
	)

	// 第45批：SCP QAD对接同步记录表
	log.Println("迁移第45批：SCP QAD同步记录表")
	migrateBatch(db, "第45批-SCP QAD同步表",
		&model.ScpQadSyncLog{},
	)

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
	escalationRuleRepo := repository.NewEscalationRuleRepository(db)
	escalationLogRepo := repository.NewEscalationLogRepository(db)
	notificationLogRepo := repository.NewNotificationLogRepository(db)
	energyRepo := repository.NewEnergyRepository(db)
	equipmentRepo := repository.NewEquipmentRepository(db)
	checkRepo := repository.NewEquipmentCheckRepository(db)
	maintRepo := repository.NewEquipmentMaintenanceRepository(db)
	repairRepo := repository.NewEquipmentRepairRepository(db)
	sparePartRepo := repository.NewSparePartRepository(db)
	equipmentPartRepo := repository.NewEquipmentPartRepository(db)
	equipmentDocumentRepo := repository.NewEquipmentDocumentRepository(db)
	equipmentDowntimeRepo := repository.NewEquipmentDowntimeRepository(db)
	equipmentDowntimeSvc := service.NewEquipmentDowntimeService(equipmentDowntimeRepo)

	spareRepo := repository.NewEquipmentSpareRepository(db)
	spareTxRepo := repository.NewEquipmentSpareTransactionRepository(db)
	spareSvc := service.NewEquipmentSpareService(spareRepo, spareTxRepo)
	spareHandler := eam.NewSpareHandler(spareSvc)

	// EAM 维修工单/流程/标准
	eamRepairJobRepo := repository.NewEamRepairJobRepository(db)
	eamRepairFlowRepo := repository.NewEamRepairFlowRepository(db)
	eamRepairStdRepo := repository.NewEamRepairStdRepository(db)
	eamRepairJobSvc := service.NewEamRepairJobService(eamRepairJobRepo, eamRepairFlowRepo, eamRepairStdRepo)
	eamRepairJobHandler := eam.NewEamRepairJobHandler(eamRepairJobSvc)

	lineRepo := repository.NewProductionLineRepository(db)
	workstationRepo := repository.NewWorkstationRepository(db)
	materialRepo := repository.NewMaterialRepository(db)
	materialCategoryRepo := repository.NewMaterialCategoryRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	workshopRepo := repository.NewWorkshopRepository(db)
	workshopConfigRepo := repository.NewWorkshopConfigRepository(db)
	workingCalendarRepo := repository.NewWorkingCalendarRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	supplierASNRepo := repository.NewSupplierASNRepository(db)
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
	teepDataRepo := repository.NewTEEPDataRepository(db)
	moldRepo := repository.NewMoldRepository(db)
	moldMaintenanceRepo := repository.NewMoldMaintenanceRepository(db)
	moldRepairRepo := repository.NewMoldRepairRepository(db)
	gaugeRepo := repository.NewGaugeRepository(db)
	gaugeCalibrationRepo := repository.NewGaugeCalibrationRepository(db)
	firstLastInspectRepo := repository.NewFirstLastInspectRepository(db)
	packageRepo := repository.NewPackageRepository(db)
	dcDataPointRepo := repository.NewDCDataPointRepository(db)
	dcScanLogRepo := repository.NewDCScanLogRepository(db)
	dcCollectRecordRepo := repository.NewDCCollectRecordRepository(db)
	importTaskRepo := repository.NewImportTaskRepository(db)
	electronicSOPRepo := repository.NewElectronicSOPRepository(db)
	codeRuleRepo := repository.NewCodeRuleRepository(db)
	flowCardRepo := repository.NewFlowCardRepository(db)
	inspectionTemplateRepo := repository.NewInspectionTemplateRepository(db)
	inspectionItemRepo := repository.NewInspectionItemRepository(db)
	inspectionPlanRepo := repository.NewInspectionPlanRepository(db)
	inspectionRecordRepo := repository.NewInspectionRecordRepository(db)
	inspectionResultRepo := repository.NewInspectionResultRepository(db)
	inspectionDefectRepo := repository.NewInspectionDefectRepository(db)
	inspectionRepo := repository.NewInspectionRepository(db)
	printTemplateRepo := repository.NewPrintTemplateRepository(db)
	noticeRepo := repository.NewNoticeRepository(db)
	capacityAnalysisRepo := repository.NewCapacityAnalysisRepository(db)
	deliveryRateRepo := repository.NewDeliveryRateRepository(db)
	changeoverMatrixRepo := repository.NewChangeoverMatrixRepository(db)
	rollingScheduleRepo := repository.NewRollingScheduleRepository(db)
	jitDemandRepo := repository.NewJITDemandRepository(db)
	transferOrderRepo := repository.NewTransferOrderRepository(db)
	transferOrderItemRepo := repository.NewTransferOrderItemRepository(db)
	stockCheckRepo := repository.NewStockCheckRepository(db)
	stockCheckItemRepo := repository.NewStockCheckItemRepository(db)
	sideLocationRepo := repository.NewSideLocationRepository(db)
	kanbanPullRepo := repository.NewKanbanPullRepository(db)
	containerRepo := repository.NewContainerRepository(db)
	containerMovementRepo := repository.NewContainerMovementRepository(db)
	containerLifecycleRepo := repository.NewContainerLifecycleRepository(db)
	containerMaintenanceRepo := repository.NewContainerMaintenanceRepository(db)
	labelTemplateRepo := repository.NewWmsLabelTemplateRepository(db)
	strategyRepo := repository.NewWmsStrategyRepository(db)
	areaRepo := repository.NewWmsAreaRepository(db)
	itemRepo := repository.NewWMSItemRepository(db)
	aiConfigRepo := repository.NewAIConfigRepository(db)
	aiConversationRepo := repository.NewConversationRepository(db)
	aiMessageRepo := repository.NewMessageRepository(db)
	labSampleRepo := repository.NewLabSampleRepository(db)
	labTestItemRepo := repository.NewLabTestItemRepository(db)
	labReportRepo := repository.NewLabReportRepository(db)
	labInstrumentRepo := repository.NewLabInstrumentRepository(db)
	labCalibrationRepo := repository.NewLabCalibrationRepository(db)
	visualInspectionRepo := repository.NewVisualInspectionRepository(db)
	mesWorkSchedulingRepo := repository.NewMesWorkSchedulingRepository(db)
	mesWorkSchedulingDetailRepo := repository.NewMesWorkSchedulingDetailRepository(db)
	mesJobReportLogRepo := repository.NewMesJobReportLogRepository(db)
	orderMonthRepo := repository.NewOrderMonthRepository(db)
	orderMonthItemRepo := repository.NewOrderMonthItemRepository(db)
	orderMonthAuditRepo := repository.NewOrderMonthAuditRepository(db)
	orderDayRepo := repository.NewOrderDayRepository(db)
	orderDayItemRepo := repository.NewOrderDayItemRepository(db)
	orderDayWorkOrderMapRepo := repository.NewOrderDayWorkOrderMapRepository(db)
	mesTeamRepo := repository.NewMesTeamRepository(db)
	mesTeamMemberRepo := repository.NewMesTeamMemberRepository(db)
	mesTeamShiftRepo := repository.NewMesTeamShiftRepository(db)
	mesProcessRepo := repository.NewMesProcessRepository(db)
	mesProcessOpRepo := repository.NewMesProcessOperationRepository(db)
	productionIssueRepo := repository.NewProductionIssueRepository(db)
	productionIssueItemRepo := repository.NewProductionIssueItemRepository(db)
	productionReturnRepo := repository.NewProductionReturnRepository(db)
	productionReturnItemRepo := repository.NewProductionReturnItemRepository(db)
	productionCompleteRepo := repository.NewProductionCompleteRepository(db)
	productionStockInRepo := repository.NewProductionStockInRepository(db)
	productionOfflineRepo := repository.NewProductionOfflineRepository(db)
	purchaseReturnRepo := repository.NewPurchaseReturnRepository(db)
	purchaseReturnItemRepo := repository.NewPurchaseReturnItemRepository(db)
	salesReturnRepo := repository.NewSalesReturnRepository(db)
	salesReturnItemRepo := repository.NewSalesReturnItemRepository(db)
	inspectionFeatureRepo := repository.NewInspectionFeatureRepository(db)
	inspectionCharacteristicRepo := repository.NewInspectionCharacteristicRepository(db)
	aqlLevelRepo := repository.NewAQLLevelRepository(db)
	aqlTableRowRepo := repository.NewAQLTableRowRepository(db)
	samplingPlanRepo := repository.NewSamplingPlanRepository(db)
	qmsSamplingPlanRepo := repository.NewQMSSamplingPlanRepository(db)
	qmsSamplingRuleRepo := repository.NewQMSSamplingRuleRepository(db)
	qmsSamplingRecordRepo := repository.NewQMSSamplingRecordRepository(db)
	personSkillRepo := repository.NewPersonSkillRepository(db)
	personSkillScoreRepo := repository.NewPersonSkillScoreRepository(db)
	completeInspectRepo := repository.NewCompleteInspectRepository(db)
	productionDailyReportRepo := repository.NewProductionDailyReportRepository(db)
	qualityWeeklyReportRepo := repository.NewQualityWeeklyReportRepository(db)
	oeeReportRepo := repository.NewOEEReportRepository(db)
	deliveryReportRepo := repository.NewDeliveryReportRepository(db)
	andonReportRepo := repository.NewAndonReportRepository(db)
	interfaceConfigRepo := repository.NewInterfaceConfigRepository(db)
	interfaceExecutionLogRepo := repository.NewInterfaceExecutionLogRepository(db)
	erpSyncLogRepo := repository.NewIntegrationERPSyncLogRepository(db)
	erpMappingRepo := repository.NewIntegrationERPMappingRepository(db)
	agvTaskRepo := repository.NewAGVTaskRepository(db)
	agvDeviceRepo := repository.NewAGVDeviceRepository(db)
	agvLocationRepo := repository.NewAGVLocationMappingRepository(db)

	// SCP Repositories
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(db)
	rfqRepo := repository.NewRFQRepository(db)
	supplierQuoteRepo := repository.NewSupplierQuoteRepository(db)
	scpSalesOrderRepo := repository.NewSCPSalesOrderRepository(db)
	inquiryRepo := repository.NewCustomerInquiryRepository(db)
	kpiRepo := repository.NewSupplierKPIRepository(db)
	purchaseInfoRepo := repository.NewSupplierPurchaseInfoRepository(db)
	poChangeRepo := repository.NewPOChangeLogRepository(db)
	purchasePlanRepo := repository.NewScpPurchasePlanRepository(db)
	scpSupplierContactRepo := repository.NewScpSupplierContactRepository(db)
	scpSupplierBankRepo := repository.NewScpSupplierBankRepository(db)
	scpQadSyncLogRepo := repository.NewScpQadSyncLogRepository(db)

	// FIN Repositories
	purchaseSettlementRepo := repository.NewPurchaseSettlementRepository(db)
	salesSettlementRepo := repository.NewSalesSettlementRepository(db)
	paymentRequestRepo := repository.NewPaymentRequestRepository(db)
	purchaseAdvanceRepo := repository.NewPurchaseAdvanceRepository(db)
	salesReceiptRepo := repository.NewSalesReceiptRepository(db)
	supplierStatementRepo := repository.NewSupplierStatementRepository(db)
	quoteCompRepo := repository.NewQuoteComparisonRepository(db)

	// Alert Repositories
	alertRuleRepo := repository.NewAlertRuleRepository(db)
	alertRecordRepo := repository.NewAlertRecordRepository(db)
	alertNotifyLogRepo := repository.NewAlertNotificationLogRepository(db)
	alertEscalationRepo := repository.NewAlertEscalationRuleRepository(db)
	alertChannelRepo := repository.NewNotificationChannelRepository(db)

	// BPM Repositories
	processModelRepo := repository.NewProcessModelRepository(db)
	nodeDefRepo := repository.NewNodeDefinitionRepository(db)
	sequenceFlowRepo := repository.NewSequenceFlowRepository(db)
	formDefRepo := repository.NewFormDefinitionRepository(db)
	formFieldRepo := repository.NewFormFieldRepository(db)
	processInstanceRepo := repository.NewProcessInstanceRepository(db)
	taskInstanceRepo := repository.NewTaskInstanceRepository(db)
	delegateRecordRepo := repository.NewDelegateRecordRepository(db)
	approvalRecordRepo := repository.NewApprovalRecordRepository(db)

	// MDM Partner Extension Repositories
	contactRepo := repository.NewContactRepository(db)
	bankAccountRepo := repository.NewBankAccountRepository(db)
	attachmentRepo := repository.NewAttachmentRepository(db)

	// 初始化服务层
	userSvc := service.NewUserService(userRepo, roleRepo, menuRepo, roleMenuRepo)
	materialSvc := service.NewMaterialService(materialRepo)
	materialCategorySvc := service.NewMaterialCategoryService(materialCategoryRepo)
	customerSvc := service.NewCustomerService(customerRepo)
	workshopSvc := service.NewWorkshopService(workshopRepo)
	workshopConfigSvc := service.NewWorkshopConfigService(workshopConfigRepo)
	workingCalendarSvc := service.NewWorkingCalendarService(workingCalendarRepo)
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
	andonSvc := service.NewAndonService(andonRepo, escalationRuleRepo, escalationLogRepo, notificationLogRepo, nil)
	escalationRuleSvc := service.NewEscalationRuleService(escalationRuleRepo, escalationLogRepo, notificationLogRepo)
	energySvc := service.NewEnergyService(energyRepo)
	equipmentSvc := service.NewEquipmentService(equipmentRepo)
	checkSvc := service.NewEquipmentCheckService(checkRepo)
	maintSvc := service.NewEquipmentMaintenanceService(maintRepo)
	repairSvc := service.NewEquipmentRepairService(repairRepo)
	sparePartSvc := service.NewSparePartService(sparePartRepo)
	equipmentPartSvc := service.NewEquipmentPartService(equipmentPartRepo)
	equipmentDocumentSvc := service.NewEquipmentDocumentService(equipmentDocumentRepo)
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
	supplierASNSvc := service.NewSupplierASNService(supplierASNRepo, supplierRepo, materialRepo)
	operLogSvc := service.NewOperLogService(operLogRepo)
	loginLogSvc := service.NewLoginLogService(loginLogRepo)
	productionOrderChangeLogSvc := service.NewProductionOrderChangeLogService(productionOrderChangeLogRepo)
	productionOrderSvc := service.NewProductionOrderService(productionRepo, productionOrderChangeLogSvc)
	oeeSvc := service.NewOEEService(oeeRepo, oeeEventRepo)
	teepDataSvc := service.NewTEEPDataService(teepDataRepo)
	moldSvc := service.NewMoldService(moldRepo)
	moldMaintenanceSvc := service.NewMoldMaintenanceService(moldMaintenanceRepo)
	moldRepairSvc := service.NewMoldRepairService(moldRepairRepo)
	gaugeSvc := service.NewGaugeService(gaugeRepo)
	gaugeCalibrationSvc := service.NewGaugeCalibrationService(gaugeCalibrationRepo)
	firstLastInspectSvc := service.NewFirstLastInspectService(firstLastInspectRepo)
	packageSvc := service.NewPackageService(packageRepo)
	kanbanSvc := service.NewKanbanService(productionRepo, lineRepo, reportRepo)
	dcSvc := service.NewDCService(dcDataPointRepo, dcScanLogRepo, dcCollectRecordRepo)
	importSvc := service.NewImportService(importTaskRepo, materialRepo, bomRepo, bomItemRepo)
	electronicSOPSvc := service.NewElectronicSOPService(electronicSOPRepo)
	mesSopSvc := service.NewMesSopService(electronicSOPRepo)
	codeRuleSvc := service.NewCodeRuleService(codeRuleRepo)
	inspectionTemplateSvc := service.NewInspectionTemplateService(inspectionTemplateRepo)
	inspectionItemSvc := service.NewInspectionItemService(inspectionItemRepo)
	inspectionPlanSvc := service.NewInspectionPlanService(inspectionPlanRepo, codeRuleSvc)
	inspectionRecordSvc := service.NewInspectionRecordService(inspectionRecordRepo, codeRuleSvc)
	inspectionResultSvc := service.NewInspectionResultService(inspectionResultRepo)
	inspectionDefectSvc := service.NewInspectionDefectService(inspectionDefectRepo, codeRuleSvc)
	inspectionSvc := service.NewInspectionService(inspectionRepo)
	labSampleSvc := service.NewLabSampleService(labSampleRepo, codeRuleSvc)
	labTestItemSvc := service.NewLabTestItemService(labTestItemRepo)
	labReportSvc := service.NewLabReportService(labReportRepo, codeRuleSvc)
	labInstrumentSvc := service.NewLabInstrumentService(labInstrumentRepo, labCalibrationRepo)
	orderMonthSvc := service.NewOrderMonthService(orderMonthRepo, orderMonthItemRepo, orderMonthAuditRepo, orderDayRepo, orderDayItemRepo)
	orderDaySvc := service.NewOrderDayService(orderDayRepo, orderDayItemRepo, orderDayWorkOrderMapRepo, productionRepo)
	mesWorkSchedulingSvc := service.NewMesWorkSchedulingService(mesWorkSchedulingRepo, mesWorkSchedulingDetailRepo)
	mesJobReportLogSvc := service.NewMesJobReportLogService(mesJobReportLogRepo)
	mesTeamSvc := service.NewMesTeamService(mesTeamRepo, mesTeamMemberRepo, mesTeamShiftRepo)
	mesProcessSvc := service.NewMesProcessService(mesProcessRepo, mesProcessOpRepo)
	productionIssueSvc := service.NewProductionIssueService(productionIssueRepo, productionIssueItemRepo, inventoryRepo)
	productionReturnSvc := service.NewProductionReturnService(productionReturnRepo, productionReturnItemRepo)
	productionCompleteSvc := service.NewProductionCompleteService(db, productionCompleteRepo, productionStockInRepo, productionRepo, inventoryRepo)
	productionOfflineSvc := service.NewProductionOfflineService(productionOfflineRepo)
	purchaseReturnSvc := service.NewPurchaseReturnService(purchaseReturnRepo, purchaseReturnItemRepo)
	salesReturnSvc := service.NewSalesReturnService(salesReturnRepo, salesReturnItemRepo)
	inspectionFeatureSvc := service.NewInspectionFeatureService(inspectionFeatureRepo)
	inspectionCharacteristicSvc := service.NewInspectionCharacteristicService(inspectionCharacteristicRepo)
	aqlSvc := service.NewAQLService(aqlLevelRepo, aqlTableRowRepo, samplingPlanRepo)
	qmsSamplingSvc := service.NewQMSSamplingService(qmsSamplingPlanRepo, qmsSamplingRuleRepo, qmsSamplingRecordRepo)
	flowCardSvc := service.NewFlowCardService(flowCardRepo)
	printTemplateSvc := service.NewPrintTemplateService(printTemplateRepo)
	noticeSvc := service.NewNoticeService(noticeRepo)
	capacityAnalysisSvc := service.NewCapacityAnalysisService(capacityAnalysisRepo)
	deliveryRateSvc := service.NewDeliveryRateService(deliveryRateRepo)
	changeoverMatrixSvc := service.NewChangeoverMatrixService(changeoverMatrixRepo)
	rollingScheduleSvc := service.NewRollingScheduleService(rollingScheduleRepo)
	jitDemandSvc := service.NewJITDemandService(jitDemandRepo)
	transferOrderSvc := service.NewTransferOrderService(transferOrderRepo, transferOrderItemRepo)
	stockCheckSvc := service.NewStockCheckService(stockCheckRepo, stockCheckItemRepo)
	sideLocationSvc := service.NewSideLocationService(sideLocationRepo)
	kanbanPullSvc := service.NewKanbanPullService(kanbanPullRepo)
	containerSvc := service.NewContainerService(containerRepo, containerMovementRepo)
	containerLifecycleSvc := service.NewContainerLifecycleService(containerLifecycleRepo, containerMaintenanceRepo, containerRepo)
	labelTemplateSvc := service.NewWmsLabelTemplateService(labelTemplateRepo)
	strategySvc := service.NewWmsStrategyService(strategyRepo)
	areaSvc := service.NewWmsAreaService(areaRepo)
	itemSvc := service.NewWMSItemService(itemRepo)
aiSvc := service.NewAIService(aiConfigRepo, aiConversationRepo, aiMessageRepo)
	personSkillSvc := service.NewPersonSkillService(personSkillRepo, personSkillScoreRepo)
	completeInspectSvc := service.NewCompleteInspectService(completeInspectRepo)
	aiExecutor := service.NewAIExecutor(productionRepo, materialRepo, warehouseRepo, inventoryRepo, iqcRepo, ipqcRepo, fqcRepo, oqcRepo, equipmentRepo, mpsRepo, scheduleRepo, userRepo, deptRepo, roleRepo)
	visualInspectionSvc := service.NewVisualInspectionService(visualInspectionRepo, aiExecutor)
	productionDailyReportSvc := service.NewProductionDailyReportService(productionDailyReportRepo)
	qualityWeeklyReportSvc := service.NewQualityWeeklyReportService(qualityWeeklyReportRepo)
	oeeReportSvc := service.NewOEEReportService(oeeReportRepo)
	deliveryReportSvc := service.NewDeliveryReportService(deliveryReportRepo)
	andonReportSvc := service.NewAndonReportService(andonReportRepo)
	integrationSvc := service.NewIntegrationService(interfaceConfigRepo, interfaceExecutionLogRepo)
	integrationExecutor := service.NewIntegrationExecutor(interfaceConfigRepo, interfaceExecutionLogRepo)
	agvSvc := service.NewAGVService(agvTaskRepo, agvDeviceRepo, agvLocationRepo)
	erpSyncSvc := service.NewERPSyncService(erpSyncLogRepo, erpMappingRepo, materialRepo, bomRepo, bomItemRepo, productionRepo)

	// 初始化事件总线
	service.InitEventBus()

	// 订阅WMS采购收货事件
	purchaseReceiveWMSHandler := service.NewPurchaseReceiveWMSHandler(db, receiveOrderRepo, receiveOrderItemRepo, warehouseRepo)
	purchaseReceiveWMSHandler.Subscribe()

	// 订阅供应商绩效计算事件
	deliveryRecordRepo := repository.NewSupplierDeliveryRecordRepository(db)
	supplierKPICalculator := service.NewSupplierKPICalculator(db, deliveryRecordRepo, kpiRepo, supplierRepo)
	supplierKPICalculator.Subscribe()

	// SCP Service
	scpSvc := service.NewSCPService(
		purchaseOrderRepo, rfqRepo, supplierQuoteRepo,
		scpSalesOrderRepo, inquiryRepo, kpiRepo,
		purchaseInfoRepo, poChangeRepo, quoteCompRepo,
	)
	purchasePlanSvc := service.NewScpPurchasePlanService(purchasePlanRepo)
	scpSupplierSvc := service.NewScpSupplierService(scpSupplierContactRepo, scpSupplierBankRepo)
	scpQadSvc := service.NewScpQadService(scpQadSyncLogRepo)

	// FIN Service
	finSvc := service.NewFinService(purchaseSettlementRepo, salesSettlementRepo, paymentRequestRepo, purchaseAdvanceRepo, salesReceiptRepo, supplierStatementRepo)

	// Alert Service
	alertSvc := service.NewAlertService(alertRuleRepo, alertRecordRepo, alertNotifyLogRepo, alertEscalationRepo, alertChannelRepo)

	// BPM Service
	bpmSvc := service.NewBPMService(processModelRepo, nodeDefRepo, sequenceFlowRepo, formDefRepo, formFieldRepo, processInstanceRepo, taskInstanceRepo, delegateRecordRepo, approvalRecordRepo, userRepo, roleRepo)

	// MDM Partner Extension Services
	contactSvc := service.NewContactService(contactRepo)
	bankAccountSvc := service.NewBankAccountService(bankAccountRepo)
	attachmentSvc := service.NewAttachmentService(attachmentRepo)

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
	wmsInboundHandler := wms.NewWMSInboundHandler(warehouseSvc)
	wmsOutboundHandler := wms.NewWMSOutboundHandler(warehouseSvc)
	salesOrderHandler := production.NewSalesOrderHandler(salesOrderSvc)
	reportHandler := production.NewReportHandler(reportSvc)
	dispatchHandler := production.NewDispatchHandler(dispatchSvc)
	apsMPSHandler := aps.NewMPSHandler(mpsSvc)
	apsMRPHandler := aps.NewMRPHandler(mrpSvc)
	apsScheduleHandler := aps.NewScheduleHandler(scheduleSvc)
	workCenterHandler := aps.NewWorkCenterHandler(workCenterSvc)
	traceHandler := trace.NewTraceHandler(traceSvc)
	andonCallHandler := andon.NewCallHandler(andonSvc, nil)
	andonRuleHandler := andon.NewRuleHandler(escalationRuleSvc)
	energyHandler := trace.NewEnergyHandler(energySvc)
	equipmentHandler := equipment.NewEquipmentHandler(equipmentSvc)
	eamAssetHandler := eam.NewAssetHandler(equipmentSvc)
	checkHandler := equipment.NewEquipmentCheckHandler(checkSvc)
	maintHandler := equipment.NewEquipmentMaintenanceHandler(maintSvc)
	repairHandler := equipment.NewEquipmentRepairHandler(repairSvc)
	sparePartHandler := equipment.NewSparePartHandler(sparePartSvc)
	equipmentPartHandler := equipment.NewEquipmentPartHandler(equipmentPartSvc)
	equipmentDocumentHandler := equipment.NewEquipmentDocumentHandler(equipmentDocumentSvc)
	equipmentDowntimeHandler := eam.NewEquipmentDowntimeHandler(equipmentDowntimeSvc)
	eamFactoryHandler := eam.NewEAMFactoryHandler(db)
	eamEquipmentOrgHandler := eam.NewEAMEquipmentOrgHandler(db)
	inspectionHandler := equipment.NewInspectionHandler(inspectionTemplateSvc, inspectionItemSvc, inspectionPlanSvc, inspectionRecordSvc, inspectionResultSvc, inspectionDefectSvc)
	dynamicRuleHandler := quality.NewDynamicRuleHandler(db)
	qualityInspPlanHandler := quality.NewQualityInspectionPlanHandler(inspectionSvc)
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
	supplierASNHandler := supplier_asn.NewSupplierASNHandler(supplierASNSvc)
	supplierMaterialRepo := repository.NewSupplierMaterialRepository(db)
	supplierMaterialSvc := service.NewSupplierMaterialService(supplierMaterialRepo)
	supplierMaterialHandler := supplier.NewSupplierMaterialHandler(supplierMaterialSvc)
	materialHandler := mdm.NewMaterialHandler(materialSvc)
	productUnitHandler := mdm.NewProductUnitHandler(materialSvc)
	materialCategoryHandler := mdm.NewMaterialCategoryHandler(materialCategorySvc)
	customerHandler := mdm.NewCustomerHandler(customerSvc)
	workshopHandler := mdm.NewWorkshopHandler(workshopSvc)
	workshopConfigHandler := mdm.NewWorkshopConfigHandler(workshopConfigSvc)
	workingCalendarHandler := mdm.NewWorkingCalendarHandler(workingCalendarSvc)
	operLogHandler := system.NewOperLogHandler(operLogSvc)
	loginLogHandler := system.NewLoginLogHandler(loginLogSvc)
	importHandler := system.NewImportHandler(importSvc)
	oeeHandler := equipment.NewOEEHandler(oeeSvc)
	teepDataHandler := equipment.NewTEEPDataHandler(teepDataSvc)
	moldHandler := equipment.NewMoldHandler(moldSvc)
	moldMaintenanceHandler := equipment.NewMoldMaintenanceHandler(moldMaintenanceSvc)
	moldRepairHandler := equipment.NewMoldRepairHandler(moldRepairSvc)
	gaugeHandler := equipment.NewGaugeHandler(gaugeSvc)
	gaugeCalibrationHandler := equipment.NewGaugeCalibrationHandler(gaugeCalibrationSvc)
	firstLastInspectHandler := production.NewFirstLastInspectHandler(firstLastInspectSvc)
	kanbanHandler := production.NewKanbanHandler(kanbanSvc)
	packageHandler := production.NewPackageHandler(packageSvc)
	dcHandler := dc.NewDataCollectionHandler(dcSvc)
	electronicSOPHandler := production.NewElectronicSOPHandler(electronicSOPSvc)
	codeRuleHandler := production.NewCodeRuleHandler(codeRuleSvc)
	flowCardHandler := production.NewFlowCardHandler(flowCardSvc)
	noticeHandler := system.NewNoticeHandler(noticeSvc)
	printTemplateHandler := system.NewPrintTemplateHandler(printTemplateSvc)
	capacityAnalysisHandler := aps.NewCapacityAnalysisHandler(capacityAnalysisSvc)
	deliveryRateHandler := aps.NewDeliveryRateHandler(deliveryRateSvc)
	changeoverMatrixHandler := aps.NewChangeoverMatrixHandler(changeoverMatrixSvc)
	rollingScheduleHandler := aps.NewRollingScheduleHandler(rollingScheduleSvc)
	jitDemandHandler := aps.NewJITDemandHandler(jitDemandSvc)
	simulationHandler := aps.NewSimulationHandler(db)
	workOrderHandler := aps.NewWorkOrderHandler(db)
	transferOrderHandler := wms.NewTransferOrderHandler(transferOrderSvc)
	stockCheckHandler := wms.NewStockCheckHandler(stockCheckSvc)
	sideLocationHandler := wms.NewSideLocationHandler(sideLocationSvc)
	kanbanPullHandler := wms.NewKanbanPullHandler(kanbanPullSvc)
	containerHandler := container.NewContainerHandler(containerSvc)
	containerLifecycleHandler := container.NewContainerLifecycleHandler(containerLifecycleSvc)
	aiConfigHandler := ai.NewAIConfigHandler(aiConfigRepo, aiSvc)
	aiChatHandler := ai.NewAIChatHandler(aiSvc, aiExecutor, aiConversationRepo, aiMessageRepo)
	visualInspectionHandler := ai.NewVisualInspectionHandler(visualInspectionSvc)
	labSampleHandler := quality.NewLabSampleHandler(labSampleSvc)
	labTestItemHandler := quality.NewLabTestItemHandler(labTestItemSvc)
	labReportHandler := quality.NewLabReportHandler(labReportSvc)
	labInstrumentHandler := quality.NewLabInstrumentHandler(labInstrumentSvc)
	inspectionFeatureHandler := quality.NewInspectionFeatureHandler(inspectionFeatureSvc)
	inspectionCharacteristicHandler := quality.NewInspectionCharacteristicHandler(inspectionCharacteristicSvc)
	aqlHandler := quality.NewAQLHandler(aqlSvc)
	qmsSamplingHandler := quality.NewQMSSamplingHandler(qmsSamplingSvc)

	// LPA (分层过程审核)
	lpaStandardRepo := repository.NewLPAStandardRepository(db)
	lpaQuestionRepo := repository.NewLPAQuestionRepository(db)
	lpaRecordRepo := repository.NewLPARecordRepository(db)
	lpaRecordItemRepo := repository.NewLPARecordItemRepository(db)
	lpaStandardSvc := service.NewLPAStandardService(lpaStandardRepo)
	lpaQuestionSvc := service.NewLPAQuestionService(lpaQuestionRepo)
	lpaRecordSvc := service.NewLPARecordService(lpaRecordRepo, codeRuleSvc)
	lpaRecordItemSvc := service.NewLPARecordItemService(lpaRecordItemRepo)
	lpaHandler := quality.NewLPAHandler(lpaStandardSvc, lpaQuestionSvc, lpaRecordSvc, lpaRecordItemSvc)

	// QRCI (质量改善)
	qrciRepo := repository.NewQRCIRepository(db)
	qrci5WhyRepo := repository.NewQRCI5WhyRepository(db)
	qrciActionRepo := repository.NewQRCIActionRepository(db)
	qrciVerificationRepo := repository.NewQRCIVerificationRepository(db)
	qrciSvc := service.NewQRCIService(qrciRepo, codeRuleSvc)
	qrci5WhySvc := service.NewQRCI5WhyService(qrci5WhyRepo)
	qrciActionSvc := service.NewQRCIActionService(qrciActionRepo)
	qrciVerificationSvc := service.NewQRCIVerificationService(qrciVerificationRepo)
	qrciHandler := quality.NewQRCIHandler(qrciSvc, qrci5WhySvc, qrciActionSvc, qrciVerificationSvc)

	mesTeamHandler := mes.NewTeamHandler(mesTeamSvc)
	mesProcessHandler := mes.NewProcessHandler(mesProcessSvc)
	mesOfflineHandler := mes.NewOfflineHandler(productionOfflineSvc)
	mesSopHandler := mes.NewSopHandler(mesSopSvc)
	productionIssueHandler := production.NewProductionIssueHandler(productionIssueSvc)
	productionReturnHandler := production.NewProductionReturnHandler(productionReturnSvc)
	productionCompleteHandler := production.NewProductionCompleteHandler(productionCompleteSvc)
	purchaseReturnHandler := wms.NewPurchaseReturnHandler(purchaseReturnSvc)
	salesReturnHandler := wms.NewSalesReturnHandler(salesReturnSvc)
	labelTemplateHandler := wms.NewWmsLabelTemplateHandler(labelTemplateSvc)
	strategyHandler := wms.NewWmsStrategyHandler(strategySvc)
	areaHandler := wms.NewWmsAreaHandler(areaSvc)
	wmsItemHandler := wms.NewWMSItemHandler(itemSvc)
	mesHandler := mes.NewMesHandler(orderMonthSvc, orderDaySvc)
	workSchedulingHandler := mes.NewWorkSchedulingHandler(mesWorkSchedulingSvc)
	jobReportHandler := mes.NewJobReportHandler(mesJobReportLogSvc)
	eamRepairJobHandler = eam.NewEamRepairJobHandler(eamRepairJobSvc)
	personSkillHandler := mes.NewPersonSkillHandler(personSkillSvc)
	completeInspectHandler := mes.NewCompleteInspectHandler(completeInspectSvc)
	productionDailyReportHandler := report.NewProductionDailyReportHandler(productionDailyReportSvc)
	qualityWeeklyReportHandler := report.NewQualityWeeklyReportHandler(qualityWeeklyReportSvc)
	oeeReportHandler := report.NewOEEReportHandler(oeeReportSvc)
	deliveryReportHandler := report.NewDeliveryReportHandler(deliveryReportSvc)
	andonReportHandler := report.NewAndonReportHandler(andonReportSvc)
	integrationHandler := integration.NewIntegrationHandler(integrationSvc, integrationExecutor)
	agvHandler := agv.NewAGVHandler(agvSvc)
	erpSyncHandler := erp_sync.NewERPSyncHandler(erpSyncSvc)

	// SCP Handlers
	rfqHandler := scp.NewRFQHandler(scpSvc)
	purchaseOrderHandler := scp.NewPurchaseOrderHandler(scpSvc)
	scpSalesOrderHandler := scp.NewSalesOrderHandler(scpSvc)
	supplierKPIHandler := scp.NewSupplierKPIHandler(scpSvc)
	supplierQuoteHandler := scp.NewSupplierQuoteHandler(scpSvc)
	customerInquiryHandler := scp.NewCustomerInquiryHandler(scpSvc)
	purchasePlanHandler := scp.NewPurchasePlanHandler(purchasePlanSvc)
	scpSupplierExtHandler := scp.NewSupplierExtHandler(scpSupplierSvc)
	qadHandler := scp.NewQadHandler(scpQadSvc)
	finHandler := fin.NewFinHandler(finSvc)

	// Alert Handler
	alertHandler := alert.NewAlertHandler(alertSvc)

	// BPM Handler
	bpmHandler := bpm.NewBPMHandler(bpmSvc)

	// BPM Task Message Rule Handler
	bpmTaskMsgRuleRepo := repository.NewBpmTaskMessageRuleRepository(db)
	bpmTaskMsgRuleSvc := service.NewBpmTaskMessageRuleService(bpmTaskMsgRuleRepo)
	bpmTaskMsgRuleHandler := bpm.NewBpmTaskMessageRuleHandler(bpmTaskMsgRuleSvc)

	// BPM Message Service and Instance API
	bpmMessageSvc := service.NewBpmMessageService(processInstanceRepo, taskInstanceRepo)
	bpmInstanceApi := service.NewBpmProcessInstanceApi(bpmSvc)
	bpmInstanceApiHandler := bpm.NewBpmInstanceApiHandler(bpmInstanceApi, bpmMessageSvc)

	// BPM Task Transfer Handler
	bpmTaskTransferRepo := repository.NewBpmTaskTransferRepository(db)
	bpmTaskTransferSvc := service.NewBpmTaskTransferService(bpmTaskTransferRepo, userRepo)
	bpmTaskTransferHandler := bpm.NewTaskTransferHandler(bpmTaskTransferSvc)

	// MDM Partner Extension Handlers
	contactHandler := mdm.NewContactHandler(contactSvc)
	bankAccountHandler := mdm.NewBankAccountHandler(bankAccountSvc)
	attachmentHandler := mdm.NewAttachmentHandler(attachmentSvc)

	// 初始化路由
	gin.SetMode(cfg.Server.Mode)
	engine := gin.Default()
	router.New(jwtUtil, userHandler, authHandler, loginLogHandler, roleHandler, menuHandler, deptHandler, dictHandler, postHandler, tenantHandler, importHandler, warehouseHandler, salesOrderHandler, reportHandler, dispatchHandler, apsMPSHandler, apsMRPHandler, apsScheduleHandler, workCenterHandler, traceHandler, energyHandler, equipmentHandler, checkHandler, maintHandler, repairHandler, sparePartHandler, lineHandler, workstationHandler, shiftHandler, bomHandler, opHandler, mdmShiftHandler, productionOrderHandler, iqcHandler, ipqcHandler, fqcHandler, oqcHandler, defectCodeHandler, defectRecordHandler, ncrHandler, spcHandler, supplierHandler, supplierASNHandler, materialHandler, materialCategoryHandler, customerHandler, workshopHandler, operLogHandler, oeeHandler, teepDataHandler, moldHandler, moldMaintenanceHandler, moldRepairHandler, gaugeHandler, gaugeCalibrationHandler, firstLastInspectHandler, kanbanHandler, packageHandler, dcHandler, electronicSOPHandler, codeRuleHandler, flowCardHandler, noticeHandler, printTemplateHandler, capacityAnalysisHandler, deliveryRateHandler, changeoverMatrixHandler, rollingScheduleHandler, jitDemandHandler, simulationHandler, workOrderHandler, transferOrderHandler, stockCheckHandler, sideLocationHandler, kanbanPullHandler, containerHandler, aiConfigHandler, aiChatHandler, andonCallHandler, andonRuleHandler, workshopConfigHandler, workingCalendarHandler, finHandler, equipmentPartHandler, equipmentDocumentHandler, equipmentDowntimeHandler, spareHandler, alertHandler, eamFactoryHandler, eamEquipmentOrgHandler, inspectionHandler, dynamicRuleHandler, qualityInspPlanHandler, bpmHandler, bpmTaskMsgRuleHandler, bpmInstanceApiHandler, bpmTaskTransferHandler, rfqHandler, purchaseOrderHandler, scpSalesOrderHandler, supplierKPIHandler, supplierQuoteHandler, customerInquiryHandler, purchasePlanHandler, scpSupplierExtHandler, qadHandler, contactHandler, bankAccountHandler, attachmentHandler, supplierMaterialHandler, containerLifecycleHandler, visualInspectionHandler, labSampleHandler, labTestItemHandler, labReportHandler, labInstrumentHandler, inspectionFeatureHandler, inspectionCharacteristicHandler, aqlHandler, qmsSamplingHandler, lpaHandler, qrciHandler, mesTeamHandler, mesProcessHandler, mesOfflineHandler, mesSopHandler, productionIssueHandler, productionReturnHandler, productionCompleteHandler, purchaseReturnHandler, salesReturnHandler, labelTemplateHandler, strategyHandler, areaHandler, wmsItemHandler, wmsInboundHandler, wmsOutboundHandler, productUnitHandler, eamAssetHandler, mesHandler, workSchedulingHandler, jobReportHandler, eamRepairJobHandler, personSkillHandler, completeInspectHandler, productionDailyReportHandler, qualityWeeklyReportHandler, oeeReportHandler, deliveryReportHandler, andonReportHandler, integrationHandler, agvHandler, erpSyncHandler).Init(engine)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("服务器启动: http://localhost%s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
