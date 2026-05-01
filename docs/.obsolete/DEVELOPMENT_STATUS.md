# MOM3.0 后端 Handler 开发状态报告

**生成日期**: 2026-04-23
**分析方法**: 静态代码分析 router.go + handler 目录
**覆盖范围**: `/data/mom3.0/mom-server/internal/`

---

## 一、总体概览

| 分类 | 数量 |
|------|------|
| Handler 文件总数 | ~107 |
| Router 中注册的 Handler 字段 | 178 |
| Handler 代码存在但路由未注册 | 约 22 个 Handler |

---

## 二、模块→Handler→路由 映射表

### 2.1 系统管理 (system) ✅ 全部完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| system/auth.go | AuthHandler | Login, Logout, GetUserInfo, ChangePassword, RefreshToken |
| system/user.go | UserHandler | GetList, GetByID, Create, Update, Delete, ResetPassword, AssignRoles |
| system/role.go | RoleHandler | List, Get, Create, Update, Delete, GetMenus, AssignMenus, GetPerms, AssignPerms |
| system/menu.go | MenuHandler | List, Tree, Get, Create, Update, Delete |
| system/dept.go | DeptHandler | List, Tree, Get, Create, Update, Delete |
| system/dict.go | DictHandler | ListType, GetType, CreateType, UpdateType, DeleteType, GetData |
| system/post.go | PostHandler | List, Get, Create, Update, Delete |
| system/tenant.go | TenantHandler | List, Get, Create, Update, Delete |
| system/oper_log.go | OperLogHandler | List |
| system/login_log.go | LoginLogHandler | List, Clean |
| system/notice.go | NoticeHandler | List, Get, Create, Update, Delete, Publish, GetMyNotices |
| system/print_template.go | PrintTemplateHandler | List, Get, Create, Update, Delete |
| system/import.go | ImportHandler | ImportMaterials, GetImportTask, DownloadTemplate, GetImportTaskResult, UploadFile, DoImport, ImportBOM, DownloadBOMTemplate |
| system/upload.go | UploadHandler | Upload, Health |

### 2.2 主数据管理 (mdm) ✅ 全部完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| mdm/material.go | MaterialHandler | List, Get, Create, Update, Delete |
| mdm/material_category.go | MaterialCategoryHandler | List, Tree, Get, Create, Update, Delete |
| mdm/bom.go | BOMHandler | List, Get, GetWithItems, Create, Update, Delete, UpdateStatus, CopyBOM, DownloadBOMTemplate, ImportBOM |
| mdm/operation.go | OperationHandler | List, Get, Create, Update, Delete |
| mdm/shift.go | ShiftHandler | List, Get, Create, Update, Delete |
| mdm/customer.go | CustomerHandler | List, Get, Create, Update, Delete |
| mdm/workshop_config.go | WorkshopConfigHandler | GetConfig, UpdateConfig |
| mdm/workshop.go | WorkshopHandler | List, Get, Create, Update, Delete |
| mdm/partner_ext.go | ContactHandler | List, Get, Create, Update, Delete |
| mdm/partner_ext.go | BankAccountHandler | List, Get, Create, Update, Delete |
| mdm/partner_ext.go | AttachmentHandler | List, Get, Create, Update, Delete |
| business/business.go | ProductionLineHandler | List, Create, Update, Delete |
| business/business.go | WorkstationHandler | List, Create, Update, Delete |
| business/business.go | ShiftHandler | List, Create, Update, Delete |
| business/equipment_org.go | EquipmentOrgHandler | List, Get, SyncFromMasterData, Update, Delete, GetFactoryList, GetFactory, CreateFactory, UpdateFactory, DeleteFactory |
| mdm/delivery_address.go | DeliveryAddressHandler | ⚠️ 需确认代码存在性 |

### 2.3 生产执行 (production) ✅ 主体完整

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| production/sales.go | SalesOrderHandler | List, Get, Create, Update, Delete, Confirm | ✅ 完整 |
| production/report.go | ReportHandler | List, Create | ✅ 完整 |
| production/report.go | DispatchHandler | List, Create, Update, Start, Complete | ✅ 完整 |
| production/order.go | ProductionOrderHandler | List, Get, Create, Update, Delete, Start, Complete | ✅ 完整 |
| production/first_last_inspect.go | FirstLastInspectHandler | List, Get, Create, Update, Delete, ListOverdue | ✅ 完整 |
| production/package.go | PackageHandler | List, Get, Create, AddItem, Seal, Delete | ✅ 完整 |
| production/production_complete.go | ProductionCompleteHandler | List, Get, Create, SubmitForInspect, Qualify, StockIn, ListStockIn | ✅ 完整 |
| production/production_issue.go | ProductionIssueHandler | List, Get, Create, Update, Delete, Submit, StartPick, ConfirmPick, Issue, Cancel | ✅ 完整 |
| production/production_return.go | ProductionReturnHandler | List, Get, Create, Update, Delete, Submit, Approve, StartReturn, ConfirmReturn, Cancel | ✅ 完整 |
| production/electronic_sop.go | ElectronicSOPHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| production/electronic_sop.go | CodeRuleHandler | List, Get, Create, Update, Delete, GenerateCode | ✅ 完整 |
| production/electronic_sop.go | FlowCardHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| production/kanban.go | KanbanHandler | GetDashboard | ⚠️ 仅有1个方法，无路由注册 |

### 2.4 MES 生产执行扩展

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| mes/mes.go | MesHandler | ListMonthPlans, GetMonthPlan, CreateMonthPlan, UpdateMonthPlan, DeleteMonthPlan, SubmitMonthPlan, ApproveMonthPlan, ReleaseMonthPlan, CloseMonthPlan, CancelMonthPlan, GetMonthPlanAudits, ListDayPlans, GetDayPlan, CreateDayPlan, UpdateDayPlan, DeleteDayPlan, PublishDayPlan, CompleteDayPlan, TerminateDayPlan, KitCheckDayPlan, DecomposeMonthPlan, GetDayPlansByMonth | ✅ 完整 |
| mes/team.go | TeamHandler | List, Get, Create, Update, Delete, ListMembers, AddMember, UpdateMember, RemoveMember, ListShifts, CreateShift, UpdateShift, DeleteShift | ✅ 完整 |
| mes/process.go | ProcessHandler | List, Get, GetByMaterial, Create, Update, Delete, UpdateStatus, Copy, Validate | ✅ 完整 |
| mes/offline.go | OfflineHandler | List, Get, Create, Update, Delete, Handle, GetItems | ✅ 完整（路由已注册） |
| mes/sop.go | SopHandler | Upload, GetByWorkOrder, ListByProcessRoute, Get, Delete, Download, List | ✅ 完整 |
| mes/complete_inspect.go | CompleteInspectHandler | GetConfig, GetOrderDayBom, GetOrderDayBomPage, GetOrderDayWorkerPage, GetOrderDayEquipmentPage, GetOrderDayEquipment, GetOrderDayWorker, Update | ✅ 完整 |
| mes/job_report.go | JobReportHandler | Create, Get, Page, Senior | ✅ 完整 |
| mes/person_skill.go | PersonSkillHandler | ListPersonSkills, GetPersonSkill, CreatePersonSkill, UpdatePersonSkill, DeletePersonSkill, GetPersonSkillDetail, EvaluateSkill, GetPersonCapability | ✅ 完整 |
| mes/work_scheduling.go | WorkSchedulingHandler | Create, Update, Delete, Get, Page, CreateDetail, UpdateDetail, DeleteDetail, GetDetail, PageDetail, ListDetail, StartDetail, PauseDetail, ResumeDetail, CompleteDetail, ReportDetail, BindEquipment, BindWorker | ✅ 完整 |
| mes/temp_bom.go | TempBOMHandler | Create, Update, Delete, ListByOrderDayItem, Get, Approve | ❌ 路由未注册 |
| mes/temp_route.go | TempRouteHandler | Create, Update, Delete, ListByOrderDay, Get, Approve | ❌ 路由未注册 |

### 2.5 APS 计划排程

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| aps/mps.go | MPSHandler | List, Get, Create, Update, Delete, Submit | ✅ 完整 |
| aps/mps.go | MRPHandler | List, Get, Create, Update, Delete, Calculate, Run, GetResults, GetShortage, GetPurchaseSuggestion | ✅ 完整 |
| aps/mps.go | ScheduleHandler | List, Create, Execute, GetResults, Delete, DragUpdate, ExecuteConstrained, GetSuggestions | ✅ 完整 |
| aps/work_center.go | WorkCenterHandler | List, Get, Create, Update, Delete, ListByWorkshop | ✅ 完整 |
| aps/aps_ext.go | CapacityAnalysisHandler | List, Get, Create, Update, GetStats | ✅ 完整 |
| aps/aps_ext.go | DeliveryRateHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| aps/aps_ext.go | ChangeoverMatrixHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| aps/aps_ext.go | RollingScheduleHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| aps/aps_ext.go | JITDemandHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| aps/rolling.go | RollingConfigHandler | List, Get, Create, Update, Delete, Enable, Disable, Execute | ❌ 路由未注册 |
| aps/rolling.go | ScheduleResultHandler | List, Get, Execute, GetWarnings | ❌ 路由未注册 |
| aps/rolling.go | DeliveryAnalysisHandler | List, Get, Create, GetDailyStats, GetWeeklyStats, GetMonthlyStats | ❌ 路由未注册 |
| aps/rolling.go | DeliveryWarningHandler | List, Get, Acknowledge, Mitigate | ❌ 路由未注册 |
| aps/rolling.go | MaterialShortageHandler | List, Get, Create, GetDailyStats | ❌ 路由未注册 |
| aps/rolling.go | ShortageRuleHandler | List, Get, Create, Update, Delete | ❌ 路由未注册 |
| aps/rolling.go | APSShiftHandler | List, Get, Create, Update, Delete | ❌ 路由未注册 |
| aps/rolling.go | APSChangeoverHandler | List, Get, Create, Update, Delete, GetByProducts | ❌ 路由未注册 |
| aps/rolling.go | ProductFamilyHandler | List, Get, Create, Update, Delete | ❌ 路由未注册 |
| aps/scheduling.go | SchedulingHandler | ExecuteScheduling, GetResults, GetGanttData, UpdateTaskTime, GetLoadData, Optimize, List, Create, Delete, Get | ❌ 路由未注册 |

### 2.6 WMS 仓储管理 ✅ 主体完整

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| wms/warehouse.go | WarehouseHandler | ListWarehouse, CreateWarehouse, UpdateWarehouse, DeleteWarehouse, ListLocation, GetLocation, CreateLocation, UpdateLocation, DeleteLocation, ListInventory, GetInventory, CreateInventory, UpdateInventory, DeleteInventory, ListReceiveOrder, GetReceiveOrder, CreateReceiveOrder, UpdateReceiveOrder, DeleteReceiveOrder, ConfirmReceiveOrder, ListDeliveryOrder, GetDeliveryOrder, CreateDeliveryOrder, UpdateDeliveryOrder, DeleteDeliveryOrder, ConfirmDeliveryOrder | ✅ 完整 |
| wms/wms_ext.go | TransferOrderHandler | List, Get, Create, Update, Delete, AddItem, Submit, Approve, Start, Ship, Receive, Complete, Cancel, GetTrace | ✅ 完整（之前报错页面已修复） |
| wms/wms_ext.go | StockCheckHandler | List, Get, Create, Update, Delete, AddItem, UpdateItem, Submit, Start, Complete, Approve, CountItem, HandleVariance, Recount, GetVariance | ✅ 完整（之前报错页面已修复） |
| wms/wms_ext.go | SideLocationHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| wms/wms_ext.go | KanbanPullHandler | List, Get, Create, Update, Delete, Trigger | ✅ 完整 |
| wms/area.go | WmsAreaHandler | Create, Update, Delete, Get, Page, Tree, ListByWarehouse | ✅ 完整 |
| wms/item.go | WMSItemHandler | List, Get, Search, Create, Update, Delete, ListByMaterial, Senior | ✅ 完整 |
| wms/label_template.go | WmsLabelTemplateHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| wms/pick.go | WMSPickHandler | Create, Assign, List, Get, Start, Complete, Cancel | ❌ 路由未注册 |
| wms/putaway.go | PutawayHandler | Create, Assign, List, Get, Start, Complete, Cancel | ❌ 路由未注册 |
| wms/purchase_return.go | PurchaseReturnHandler | ListPurchaseReturns, GetPurchaseReturn, CreatePurchaseReturn, UpdatePurchaseReturn, DeletePurchaseReturn, SubmitPurchaseReturn, ApprovePurchaseReturn, StartReturnPurchaseReturn, ConfirmPurchaseReturn, CancelPurchaseReturn | ✅ 完整 |
| wms/sales_return.go | SalesReturnHandler | ListSalesReturns, GetSalesReturn, CreateSalesReturn, UpdateSalesReturn, DeleteSalesReturn, SubmitSalesReturn, ApproveSalesReturn, StartReturnSalesReturn, ConfirmSalesReturn, CancelSalesReturn | ✅ 完整 |
| wms/strategy.go | WmsStrategyHandler | List, Get, Create, Update, Delete | ✅ 完整 |

### 2.7 质量管理 (quality)

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| quality/iqc.go | IQCHandler | List, Get, Create, Update, Delete, Inspect | ✅ 完整 |
| quality/ipqc.go | IPQCHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| quality/fqc.go | FQCHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| quality/oqc.go | OQCHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| quality/defect_code.go | DefectCodeHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| quality/defect_record.go | DefectRecordHandler | List, Get, Create, Update, Delete, Handle | ✅ 完整 |
| quality/ncr.go | NCRHandler | List, Get, Create, Update, Delete, Resolve, Assign, Close | ✅ 完整 |
| quality/spc.go | SPCHandler | List, Get, Create, Update, Delete, GetChartData, GetStats, GetCapability | ✅ 完整 |
| quality/lab.go | LabSampleHandler | List, Get, Create, Update, Delete, SubmitForInspection | ✅ 完整 |
| quality/lab.go | LabTestItemHandler | ListBySampleID, Create, Update, Delete | ✅ 完整 |
| quality/lab.go | LabReportHandler | List, Get, Create, Update, Approve, Delete | ✅ 完整 |
| quality/lab_instrument.go | LabInstrumentHandler | ListLabInstruments, GetLabInstrument, CreateLabInstrument, UpdateLabInstrument, DeleteLabInstrument, GetLabInstrumentCalibrations, RecordCalibration | ✅ 完整 |
| quality/inspection_feature.go | InspectionFeatureHandler | ListInspectionFeatures, GetInspectionFeature, CreateInspectionFeature, UpdateInspectionFeature, DeleteInspectionFeature, BatchCreateInspectionFeature, GetFeaturesByProduct | ✅ 完整 |
| quality/aql.go | AQLHandler | ListAQLLevels, GetAQLLevel, CreateAQLLevel, UpdateAQLLevel, DeleteAQLLevel, ListAQLTableRows, CreateAQLTableRow, CalculateSampleSize, ListSamplingPlans, GetSamplingPlan, CreateSamplingPlan, UpdateSamplingPlan, DeleteSamplingPlan | ✅ 完整 |
| quality/sampling.go | QMSSamplingHandler | CreatePlan, UpdatePlan, DeletePlan, ListPlan, GetPlan, UpdateRules, Calculate, CreateRecord, ListRecord | ✅ 完整 |
| quality/inspection.go | InspectionHandler | ListPlans, GetPlanByID, CreatePlan, UpdatePlan, DeletePlan, CalculateSampleSize, SeedAQLData | ❌ 路由未注册 |
| quality/lpa.go | LPAHandler | ListStandards, GetStandard, CreateStandard, UpdateStandard, DeleteStandard, ListQuestions, AddQuestion, ListRecords, GetRecord, CreateRecord, VerifyRecord | ❌ 路由未注册 |
| quality/qrci.go | QRCIHandler | List, Get, Create, Update, Close, Delete, List5Why, Add5Why, ListActions, AddAction, UpdateAction, AddVerification | ❌ 路由未注册 |

### 2.8 设备管理 (equipment)

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| equipment/equipment.go | EquipmentHandler | List, Get, Create, Update, Delete, Status | ✅ 完整 |
| equipment/inspection.go | EquipmentCheckHandler | List | ✅ 完整 |
| equipment/tool_container.go | EquipmentMaintenanceHandler | List | ✅ 完整 |
| equipment/tool_container.go | EquipmentRepairHandler | List, Create, Start, Complete | ✅ 完整 |
| equipment/tool_container.go | SparePartHandler | List | ✅ 完整 |
| equipment/oee.go | OEEHandler | List, Get, Calculate, Chart, Delete | ✅ 完整 |
| equipment/equipment_ext.go | TEEPDataHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| equipment/tool_container.go | MoldHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| equipment/tool_container.go | MoldMaintenanceHandler | List, Create | ✅ 完整 |
| equipment/tool_container.go | MoldRepairHandler | List, Create | ✅ 完整 |
| equipment/tool_container.go | GaugeHandler | List, Get, Create, Update, Delete | ✅ 完整 |
| equipment/tool_container.go | GaugeCalibrationHandler | List, Create | ✅ 完整 |
| equipment/equipment_part.go | EquipmentPartHandler | List, Get, Create, Update, Delete, ListByEquipment | ✅ 完整 |
| equipment/equipment_document.go | EquipmentDocumentHandler | List, Get, Create, Update, Delete, ListByEquipment | ✅ 完整 |

### 2.9 EAM 设备资产维修

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| eam/downtime.go | EquipmentDowntimeHandler | List, Get, Create, Update, Delete, StartDowntime, EndDowntime | ✅ 完整 |
| eam/spare.go | SpareHandler | List, Get, Create, Update, Delete, Input, Output, Transactions | ✅ 完整 |
| eam/repair_job.go | EamRepairJobHandler | 待确认方法 | ✅ 已注册 |
| eam/inspection.go | EAMInspectionHandler | 待确认 | ❌ 需检查代码和路由 |
| business/equipment_org.go | EquipmentOrgHandler | List, Get, SyncFromMasterData, Update, Delete, GetFactoryList, GetFactory, CreateFactory, UpdateFactory, DeleteFactory | ✅ 完整 |

### 2.10 追溯管理 (trace)

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| trace/trace.go | TraceHandler | TraceBySerial, TraceByBatch, TraceByOrder, ForwardTrace, BackwardTrace | ✅ 完整 |
| trace/trace.go | AndonHandler | List, Create, Response, Resolve | ⚠️ Handler 存在但无独立路由（与 andon 模块重复） |
| trace/trace.go | EnergyHandler | List, GetStats, GetTrend, Create | ✅ 完整 |

### 2.11 安灯管理 (andon) ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| andon/call.go | CallHandler | List, Get, Create, Respond, Resolve, Escalate, GetStatistics |
| andon/rule.go | RuleHandler | List, Get, Create, Update, Delete |

### 2.12 SCP 供应链

| Handler 文件 | Handler 名称 | 已注册方法 | 状态 |
|-------------|-------------|-----------|------|
| scp/rfq.go | RFQHandler | List, Get, Create, Update, Delete, Publish, Close, GetQuotes, Award | ✅ 完整 |
| scp/purchase_order.go | PurchaseOrderHandler | List, Get, Create, Update, Delete, Submit, Approve, Reject, Issue, Close, Cancel, Receive | ✅ 完整 |
| scp/sales_order.go | SalesOrderHandler | List, Get, Create, Update, Delete, Submit, Approve, Reject, Confirm, Close, Cancel | ✅ 完整 |
| scp/supplier_quote.go | SupplierQuoteHandler | List, Get, Create, GetQuotes, Award | ✅ 完整 |
| scp/kpi.go | SupplierKPIHandler | List, GetByMonthly, Create, GetRanking | ✅ 完整 |
| scp/customer_inquiry.go | CustomerInquiryHandler | List, Get, Create, Update, Delete, Send, Quote, Win, Lose, Cancel | ✅ 完整 |
| scp/purchase_plan.go | PurchasePlanHandler | List, Get, Create, Update, Delete, GetItems, Confirm, Publish, Close | ✅ 完整 |
| scp/supplier.go | SupplierExtHandler | ListContacts, GetContact, ListContactsBySupplier, CreateContact, UpdateContact, DeleteContact, ListBanks, GetBank, ListBanksBySupplier, CreateBank, UpdateBank, DeleteBank | ✅ 完整 |
| scp/qad.go | QadHandler | Sync, GetSyncStatus, GetSyncLog, Confirm, Delivery | ✅ 完整 |
| scp/mrs.go | MRSHandler | 待确认 | ❌ 需检查 |

### 2.13 供应商管理 ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| supplier/supplier.go | SupplierHandler | List, Get, Create, Update, Delete |
| supplier/supplier_material.go | SupplierMaterialHandler | List, Get, Create, Update, Delete, ListBySupplier, ListByMaterial, SetPreferred |
| supplier_asn/supplier_asn.go | SupplierASNHandler | List, Get, GetByNo, Create, Update, Delete, Submit, Confirm, StartReceiving, CompleteReceiving, Cancel, AddItem |

### 2.14 数据采集 (dc) ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| dc/data_collection.go | DataCollectionHandler | ListDataPoint, GetDataPoint, CreateDataPoint, UpdateDataPoint, DeleteDataPoint, ListScanLog, CreateScanLog, ListCollectRecord |

### 2.15 容器管理 (container) ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| container/container.go | ContainerHandler | List, Get, Create, Update, Delete, In, Out, Return, Transfer, Clean, Movements |
| container/container_lifecycle.go | ContainerLifecycleHandler | ListContainerLifecycles, GetContainerLifecycle, InitializeContainer, RecordMaintenance, CompleteMaintenance, RetireContainer, GetContainerTimeline, ListContainerMaintenances |

### 2.16 AI / 视觉检测 ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| ai/ai_config.go | AIConfigHandler | GetConfig, UpdateConfig, TestConfig, GetSchema |
| ai/chat.go | AIChatHandler | SendMessage, ExecuteOperation, ListConversations, GetConversation, DeleteConversation |
| ai/visual_inspection.go | VisualInspectionHandler | ListVisualInspectionTasks, GetVisualInspectionTask, CreateVisualInspectionTask, DeleteVisualInspectionTask, GetVisualInspectionResult, ManualReview, GetVisualInspectionStats |

### 2.17 BPM 业务流程 ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| bpm/bpm.go | BPMHandler | ListProcessModels, GetProcessModel, CreateProcessModel, UpdateProcessModel, DeleteProcessModel, PublishProcessModel, ListNodes, CreateNode, UpdateNode, DeleteNode, ListFlows, CreateFlow, UpdateFlow, DeleteFlow, ListFormDefinitions, GetFormDefinition, CreateFormDefinition, UpdateFormDefinition, DeleteFormDefinition, ListFormFields, CreateFormField, UpdateFormField, DeleteFormField, ListProcessInstances, GetProcessInstance, CreateProcessInstance, CancelProcessInstance, TerminateProcessInstance, ListTasksByAssignee, GetTask, ApproveTask, RejectTask, ListDelegates, CreateDelegate, UpdateDelegate, DeleteDelegate, ListApprovalRecords |
| bpm/bpm_instance.go | BpmInstanceApiHandler | StartProcessInstance, CompleteTask, GetProcessInstance |
| bpm/task_message_rule.go | BpmTaskMessageRuleHandler | List, Get, Create, Update, Delete, Enable, Disable |
| bpm/task_transfer.go | TaskTransferHandler | TransferTask, GetTransferHistory, GetTaskCandidates, GetTaskCandidateGroups, AssignTask |

### 2.18 Alert 告警管理 ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| alert/alert.go | AlertHandler | ListRules, GetRule, CreateRule, UpdateRule, DeleteRule, EnableRule, DisableRule, ListRecords, GetRecord, AcknowledgeRecord, ResolveRecord, CloseRecord, GetStatistics, ListNotificationLogs, ListEscalationRules, CreateEscalationRule, ListChannels, GetChannel, CreateChannel, UpdateChannel, DeleteChannel, EnableChannel, DisableChannel, SendNotification |

### 2.19 报表 (report) ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| report/production_daily_report.go | ProductionDailyReportHandler | ListProductionDailyReports, GetProductionDailyReport, GenerateDailyReport, GetDailyReportSummary |
| report/quality_weekly_report.go | QualityWeeklyReportHandler | ListQualityWeeklyReports, GetQualityWeeklyReport, GenerateWeeklyReport |
| report/oee_report.go | OEEReportHandler | ListOEEReports, GetOEEReport, GenerateOEEReport |
| report/delivery_report.go | DeliveryReportHandler | ListDeliveryReports, GetDeliveryReport, GenerateDeliveryReport |
| report/andon_report.go | AndonReportHandler | ListAndonReports, GetAndonReport, GenerateAndonReport |

### 2.20 财务 (fin) ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| fin/fin.go | FinHandler | ListPurchaseSettlements, GetPurchaseSettlement, CreatePurchaseSettlement, SubmitPurchaseSettlement, ApprovePurchaseSettlement, CancelPurchaseSettlement, DeletePurchaseSettlements, ListSalesSettlements, GetSalesSettlement, CreateSalesSettlement, SubmitSalesSettlement, ApproveSalesSettlement, CancelSalesSettlement, DeleteSalesSettlements, ListPaymentRequests, GetPaymentRequest, CreatePaymentRequest, SubmitPaymentRequest, ApprovePaymentRequest, RejectPaymentRequest, PayPaymentRequest, DeletePaymentRequest, ListPurchaseAdvances, CreatePurchaseAdvance, ListSalesReceipts, CreateSalesReceipt, ListSupplierStatements, GetSupplierStatement |

### 2.21 AGV ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| agv/agv.go | AGVHandler | ListTasks, GetTask, CreateTask, CancelTask, AssignTask, CompleteTask, StartTask, ListDevices, GetDevice, CreateDevice, UpdateDeviceStatus, ListLocations, GetLocation, CreateLocation, UpdateLocation, DeleteLocation, Heartbeat, TaskCallback, GetAvailableAGVs |

### 2.22 系统集成 (integration / erp_sync) ✅ 完整

| Handler 文件 | Handler 名称 | 已注册方法 |
|-------------|-------------|-----------|
| integration/integration.go | IntegrationHandler | ListConfigs, GetConfig, CreateConfig, UpdateConfig, DeleteConfig, ExecuteConfig, TestConfig, ListFieldMaps, CreateFieldMap, UpdateFieldMap, DeleteFieldMap, ListTriggers, CreateTrigger, UpdateTrigger, DeleteTrigger, ListExecutionLogs, GetExecutionLog, GetConstantOptions |
| erp_sync/erp_sync.go | ERPSyncHandler | ListSyncLogs, GetSyncLog, GetSyncStatus, SyncBOM, SyncProductionOrder, SyncStock, PushReport, PushStockIn, PushQualityData |

---

## 三、Handler 代码存在但路由未注册清单

| # | Handler 文件 | Handler 名称 | 未注册方法数 | 对应前端路由 |
|---|-------------|-------------|------------|-------------|
| 1 | mes/temp_bom.go | TempBOMHandler | 6 | /mes/temp-bom |
| 2 | mes/temp_route.go | TempRouteHandler | 6 | /mes/temp-route |
| 3 | aps/rolling.go | RollingConfigHandler | 8 | /aps/rolling-config |
| 4 | aps/rolling.go | ScheduleResultHandler | 4 | /aps/schedule-result |
| 5 | aps/rolling.go | DeliveryAnalysisHandler | 6 | /aps/delivery-analysis |
| 6 | aps/rolling.go | DeliveryWarningHandler | 4 | /aps/delivery-warning |
| 7 | aps/rolling.go | MaterialShortageHandler | 4 | /aps/material-shortage |
| 8 | aps/rolling.go | ShortageRuleHandler | 5 | /aps/shortage-rule |
| 9 | aps/rolling.go | APSShiftHandler | 5 | /aps/shift-config |
| 10 | aps/rolling.go | APSChangeoverHandler | 6 | /aps/changeover-matrix |
| 11 | aps/rolling.go | ProductFamilyHandler | 5 | /aps/product-family |
| 12 | aps/scheduling.go | SchedulingHandler | 10 | /aps/scheduling |
| 13 | production/kanban.go | KanbanHandler | 1 | /production/kanban |
| 14 | wms/pick.go | WMSPickHandler | 7 | /wms/pick |
| 15 | wms/putaway.go | PutawayHandler | 7 | /wms/putaway |
| 16 | quality/inspection.go | InspectionHandler | 7 | /quality/inspection-plan |
| 17 | quality/lpa.go | LPAHandler | 11 | /quality/lpa |
| 18 | quality/qrci.go | QRCIHandler | 12 | /quality/qrci |
| 19 | trace/trace.go | AndonHandler | 4 | (与 andon 模块重复) |
| 20 | mdm/delivery_address.go | DeliveryAddressHandler | 需确认 | /mdm/delivery-address |
| 21 | eam/inspection.go | EAMInspectionHandler | 需确认 | /eam/inspection |
| 22 | scp/mrs.go | MRSHandler | 需确认 | /scp/mrs |

---

## 四、关键发现

1. **之前报错的 3 个页面实际上路由已注册**：
   - `/wms/transfer` → TransferOrderHandler ✅ 已注册
   - `/wms/stock-check` → StockCheckHandler ✅ 已注册
   - `/mes/offline` → OfflineHandler ✅ 已注册
   - 注：如果之前测试报 500，可能是数据库表缺失或 Handler 代码问题，非路由问题

2. **APS 模块是路由未注册的重灾区**：9 个 Handler 存在但未注册路由

3. **质量模块的 LPA/QRCI/检验计划**：Handler 代码已有但路由未注册

4. **Fin、BPM、Report、AGV、Integration**：全部路由已注册，代码完整

5. **WMS 的 Pick（拣货）和 Putaway（上架）**：Handler 存在但路由未注册

---

*报告生成时间: 2026-04-23*
*数据来源: router.go + handler 目录静态分析*
