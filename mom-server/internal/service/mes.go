package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// ========== OrderMonthService 月计划服务 ==========

type OrderMonthService struct {
	repo       *repository.OrderMonthRepository
	itemRepo   *repository.OrderMonthItemRepository
	auditRepo  *repository.OrderMonthAuditRepository
	orderDayRepo     *repository.OrderDayRepository
	orderDayItemRepo *repository.OrderDayItemRepository
}

func NewOrderMonthService(
	repo *repository.OrderMonthRepository,
	itemRepo *repository.OrderMonthItemRepository,
	auditRepo *repository.OrderMonthAuditRepository,
	orderDayRepo *repository.OrderDayRepository,
	orderDayItemRepo *repository.OrderDayItemRepository,
) *OrderMonthService {
	return &OrderMonthService{repo: repo, itemRepo: itemRepo, auditRepo: auditRepo, orderDayRepo: orderDayRepo, orderDayItemRepo: orderDayItemRepo}
}

func (s *OrderMonthService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.OrderMonth, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *OrderMonthService) Get(ctx context.Context, id int64) (*model.OrderMonth, error) {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	// 加载明细
	items, err := s.itemRepo.ListByMonthPlanID(ctx, id)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return order, nil
}

func (s *OrderMonthService) Create(ctx context.Context, tenantID int64, req *model.OrderMonthCreate, username string) (*model.OrderMonth, error) {
	// 生成单号
	monthPlanNo, err := s.repo.GenerateMonthPlanNo(ctx, tenantID, req.PlanMonth)
	if err != nil {
		return nil, fmt.Errorf("生成单号失败: %w", err)
	}

	// 计算汇总
	var totalPlanQty float64
	for _, item := range req.Items {
		totalPlanQty += item.PlanQty
	}

	order := &model.OrderMonth{
		TenantID:        tenantID,
		MonthPlanNo:     monthPlanNo,
		PlanMonth:       req.PlanMonth,
		Title:           req.Title,
		SourceType:      req.SourceType,
		SourceNo:        &req.SourceNo,
		WorkshopID:      req.WorkshopID,
		ApprovalStatus:  "DRAFT",
		TotalPlanQty:    totalPlanQty,
		TotalProductCount: len(req.Items),
		CreatedBy:       &username,
		Remark:          &req.Remark,
	}

	if err := s.repo.Create(ctx, order); err != nil {
		return nil, fmt.Errorf("创建月计划失败: %w", err)
	}

	// 创建明细
	for i, item := range req.Items {
		deliveryDate, _ := time.Parse("2006-01-02", item.DeliveryDate)
		monthItem := &model.OrderMonthItem{
			MonthPlanID:   order.ID,
			LineNo:        i + 1,
			ProductID:     item.ProductID,
			ProductCode:   &item.ProductCode,
			ProductName:   &item.ProductName,
			Specification: &item.Specification,
			Unit:          &item.Unit,
			PlanQty:       item.PlanQty,
			DeliveryDate:  &deliveryDate,
			Priority:      item.Priority,
			Remark:        &item.Remark,
			TenantID:      tenantID,
		}
		if err := s.itemRepo.Create(ctx, monthItem); err != nil {
			return nil, fmt.Errorf("创建月计划明细失败: %w", err)
		}
	}

	// 记录审核日志
	audit := &model.OrderMonthAudit{
		MonthPlanID:    order.ID,
		ApprovalStatus:  "CREATE",
		ApproverName:    &username,
		ApprovalTime:    toPtr(time.Now()),
		Comment:         &req.Remark,
		TenantID:       tenantID,
	}
	s.auditRepo.Create(ctx, audit)

	return s.Get(ctx, order.ID)
}

func (s *OrderMonthService) Update(ctx context.Context, id int64, req *model.OrderMonthUpdate, username string) (*model.OrderMonth, error) {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	// 仅草稿状态可更新
	if order.ApprovalStatus != "DRAFT" {
		return nil, errors.New("仅草稿状态可更新")
	}

	// 更新主表
	updates := map[string]interface{}{}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	// 重新计算汇总
	var totalPlanQty float64
	for _, item := range req.Items {
		totalPlanQty += item.PlanQty
	}
	updates["total_plan_qty"] = totalPlanQty
	updates["total_product_count"] = len(req.Items)

	if err := s.repo.Update(ctx, uint(id), updates); err != nil {
		return nil, fmt.Errorf("更新月计划失败: %w", err)
	}

	// 删除旧明细，创建新明细
	if err := s.itemRepo.DeleteByMonthPlanID(ctx, id); err != nil {
		return nil, fmt.Errorf("删除旧明细失败: %w", err)
	}

	for i, item := range req.Items {
		deliveryDate, _ := time.Parse("2006-01-02", item.DeliveryDate)
		monthItem := &model.OrderMonthItem{
			MonthPlanID:   int64(id),
			LineNo:        i + 1,
			ProductID:     item.ProductID,
			ProductCode:   &item.ProductCode,
			ProductName:   &item.ProductName,
			Specification: &item.Specification,
			Unit:          &item.Unit,
			PlanQty:       item.PlanQty,
			DeliveryDate:  &deliveryDate,
			Priority:      item.Priority,
			Remark:        &item.Remark,
			TenantID:      order.TenantID,
		}
		if err := s.itemRepo.Create(ctx, monthItem); err != nil {
			return nil, fmt.Errorf("创建月计划明细失败: %w", err)
		}
	}

	return s.Get(ctx, id)
}

func (s *OrderMonthService) Delete(ctx context.Context, id int64) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	// 仅草稿状态可删除
	if order.ApprovalStatus != "DRAFT" {
		return errors.New("仅草稿状态可删除")
	}

	// 删除明细
	if err := s.itemRepo.DeleteByMonthPlanID(ctx, id); err != nil {
		return err
	}

	return s.repo.Delete(ctx, uint(id))
}

func (s *OrderMonthService) Submit(ctx context.Context, id int64, userID int64, username string) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.ApprovalStatus != "DRAFT" {
		return errors.New("仅草稿状态可提交")
	}

	now := time.Now()
	if err := s.repo.Update(ctx, uint(id), map[string]interface{}{
		"approval_status": "SUBMITTED",
		"submitted_by":    userID,
		"submitted_at":    now,
	}); err != nil {
		return err
	}

	// 记录审核日志
	audit := &model.OrderMonthAudit{
		MonthPlanID:    order.ID,
		ApprovalStatus: "SUBMIT",
		ApproverID:     &userID,
		ApproverName:   &username,
		ApprovalTime:   &now,
		TenantID:       order.TenantID,
	}
	return s.auditRepo.Create(ctx, audit)
}

func (s *OrderMonthService) Approve(ctx context.Context, id int64, userID int64, username string, comment string) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.ApprovalStatus != "SUBMITTED" {
		return errors.New("仅已提交状态可审核")
	}

	now := time.Now()
	if err := s.repo.Update(ctx, uint(id), map[string]interface{}{
		"approval_status": "APPROVED",
		"approved_by":     userID,
		"approved_at":     now,
	}); err != nil {
		return err
	}

	audit := &model.OrderMonthAudit{
		MonthPlanID:    order.ID,
		ApprovalStatus: "APPROVE",
		ApproverID:     &userID,
		ApproverName:   &username,
		ApprovalTime:   &now,
		Comment:        &comment,
		TenantID:       order.TenantID,
	}
	return s.auditRepo.Create(ctx, audit)
}

func (s *OrderMonthService) Release(ctx context.Context, id int64, userID int64, username string) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.ApprovalStatus != "APPROVED" {
		return errors.New("仅审核通过状态可下达")
	}

	now := time.Now()
	if err := s.repo.Update(ctx, uint(id), map[string]interface{}{
		"approval_status": "RELEASED",
		"released_by":      userID,
		"released_at":      now,
	}); err != nil {
		return err
	}

	audit := &model.OrderMonthAudit{
		MonthPlanID:    order.ID,
		ApprovalStatus: "RELEASE",
		ApproverID:     &userID,
		ApproverName:   &username,
		ApprovalTime:   &now,
		TenantID:       order.TenantID,
	}
	return s.auditRepo.Create(ctx, audit)
}

func (s *OrderMonthService) Close(ctx context.Context, id int64, userID int64, username string) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.ApprovalStatus != "RELEASED" && order.ApprovalStatus != "APPROVED" {
		return errors.New("只有已下达或审核通过状态可关闭")
	}

	now := time.Now()
	if err := s.repo.Update(ctx, uint(id), map[string]interface{}{
		"approval_status": "CLOSED",
	}); err != nil {
		return err
	}

	audit := &model.OrderMonthAudit{
		MonthPlanID:    order.ID,
		ApprovalStatus: "CLOSE",
		ApproverID:     &userID,
		ApproverName:   &username,
		ApprovalTime:   &now,
		TenantID:       order.TenantID,
	}
	return s.auditRepo.Create(ctx, audit)
}

func (s *OrderMonthService) Cancel(ctx context.Context, id int64, userID int64, username string, comment string) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.ApprovalStatus == "CLOSED" || order.ApprovalStatus == "CANCELLED" {
		return errors.New("已关闭或已取消的月计划不能取消")
	}

	now := time.Now()
	if err := s.repo.Update(ctx, uint(id), map[string]interface{}{
		"approval_status": "CANCELLED",
	}); err != nil {
		return err
	}

	audit := &model.OrderMonthAudit{
		MonthPlanID:    order.ID,
		ApprovalStatus: "CANCEL",
		ApproverID:     &userID,
		ApproverName:   &username,
		ApprovalTime:   &now,
		Comment:        &comment,
		TenantID:       order.TenantID,
	}
	return s.auditRepo.Create(ctx, audit)
}

func (s *OrderMonthService) GetAudits(ctx context.Context, monthPlanID int64) ([]model.OrderMonthAudit, error) {
	return s.auditRepo.ListByMonthPlanID(ctx, monthPlanID)
}

// Decompose 将月计划分解为日计划
func (s *OrderMonthService) Decompose(ctx context.Context, monthPlanID int64, username string) ([]model.OrderDay, error) {
	monthPlan, err := s.repo.GetByID(ctx, uint(monthPlanID))
	if err != nil {
		return nil, fmt.Errorf("获取月计划失败: %w", err)
	}

	if monthPlan.ApprovalStatus != "RELEASED" {
		return nil, errors.New("仅已下达状态的月计划可分解")
	}

	items, err := s.itemRepo.ListByMonthPlanID(ctx, monthPlanID)
	if err != nil {
		return nil, fmt.Errorf("获取月计划明细失败: %w", err)
	}
	if len(items) == 0 {
		return nil, errors.New("月计划没有明细，无法分解")
	}

	planMonth := monthPlan.PlanMonth
	startDate, err := time.Parse("2006-01", planMonth)
	if err != nil {
		return nil, fmt.Errorf("解析计划月份失败: %w", err)
	}
	endDate := startDate.AddDate(0, 1, -1)

	existingDays, err := s.orderDayRepo.GetOrderDaysByMonthID(ctx, monthPlanID)
	if err != nil {
		return nil, fmt.Errorf("获取已有日计划失败: %w", err)
	}
	existingDates := make(map[string]bool)
	for _, d := range existingDays {
		existingDates[d.PlanDate.Format("2006-01-02")] = true
	}

	var workDays int
	current := startDate
	for !current.After(endDate) {
		if current.Weekday() != time.Saturday && current.Weekday() != time.Sunday {
			workDays++
		}
		current = current.AddDate(0, 0, 1)
	}

	if workDays == 0 {
		return nil, errors.New("计划月份没有工作日")
	}

	dayAvgQty := monthPlan.TotalPlanQty / float64(workDays)

	var createdDays []model.OrderDay
	current = startDate
	for !current.After(endDate) {
		if current.Weekday() == time.Saturday || current.Weekday() == time.Sunday {
			current = current.AddDate(0, 0, 1)
			continue
		}

		dateStr := current.Format("2006-01-02")
		if existingDates[dateStr] {
			current = current.AddDate(0, 0, 1)
			continue
		}

		dayPlanNo, err := s.orderDayRepo.GenerateDayPlanNo(ctx, monthPlan.TenantID, current)
		if err != nil {
			return nil, fmt.Errorf("生成日计划单号失败: %w", err)
		}

		dayPlan := &model.OrderDay{
			TenantID:     monthPlan.TenantID,
			DayPlanNo:   dayPlanNo,
			PlanDate:    current,
			MonthPlanID: &monthPlan.ID,
			MonthPlanNo: &monthPlan.MonthPlanNo,
			WorkshopID:  monthPlan.WorkshopID,
			WorkshopName: monthPlan.WorkshopName,
			ShiftType:   "ALL",
			Status:      "DRAFT",
			KitStatus:   "PENDING",
			TotalPlanQty: dayAvgQty,
			TotalProductCount: len(items),
			CreatedBy:   &username,
		}
		if err := s.orderDayRepo.Create(ctx, dayPlan); err != nil {
			return nil, fmt.Errorf("创建日计划失败: %w", err)
		}

		for i, item := range items {
			itemDayQty := item.PlanQty / float64(workDays)
			dayItem := &model.OrderDayItem{
				DayPlanID:       dayPlan.ID,
				LineNo:          i + 1,
				ProductID:       item.ProductID,
				ProductCode:     item.ProductCode,
				ProductName:     item.ProductName,
				Specification:   item.Specification,
				Unit:            item.Unit,
				PlanQty:         itemDayQty,
				ProductionMode:  "BATCH",
				MonthPlanItemID: &item.ID,
				ItemStatus:      "PENDING",
				KitStatus:       "PENDING",
				Priority:        item.Priority,
				Remark:          item.Remark,
				TenantID:        monthPlan.TenantID,
			}
			_ = s.orderDayItemRepo.Create(ctx, dayItem)
		}

		createdDays = append(createdDays, *dayPlan)
		current = current.AddDate(0, 0, 1)
	}

	return createdDays, nil
}

// ========== OrderDayService 日计划服务 ==========

type OrderDayService struct {
	repo              *repository.OrderDayRepository
	itemRepo          *repository.OrderDayItemRepository
	workOrderMapRepo  *repository.OrderDayWorkOrderMapRepository
	productionRepo   *repository.ProductionOrderRepository
}

func NewOrderDayService(
	repo *repository.OrderDayRepository,
	itemRepo *repository.OrderDayItemRepository,
	workOrderMapRepo *repository.OrderDayWorkOrderMapRepository,
	productionRepo *repository.ProductionOrderRepository,
) *OrderDayService {
	return &OrderDayService{
		repo:             repo,
		itemRepo:         itemRepo,
		workOrderMapRepo: workOrderMapRepo,
		productionRepo:  productionRepo,
	}
}

func (s *OrderDayService) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.OrderDay, int64, error) {
	return s.repo.List(ctx, tenantID, query)
}

func (s *OrderDayService) Get(ctx context.Context, id int64) (*model.OrderDay, error) {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	// 加载明细
	items, err := s.itemRepo.ListByDayPlanID(ctx, id)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return order, nil
}

func (s *OrderDayService) Create(ctx context.Context, tenantID int64, req *model.OrderDayCreate, username string) (*model.OrderDay, error) {
	// 解析计划日期
	planDate, err := time.Parse("2006-01-02", req.PlanDate)
	if err != nil {
		return nil, fmt.Errorf("日期格式错误: %w", err)
	}

	// 生成单号
	dayPlanNo, err := s.repo.GenerateDayPlanNo(ctx, tenantID, planDate)
	if err != nil {
		return nil, fmt.Errorf("生成单号失败: %w", err)
	}

	// 计算汇总
	var totalPlanQty float64
	for _, item := range req.Items {
		totalPlanQty += item.PlanQty
	}

	order := &model.OrderDay{
		TenantID:        tenantID,
		DayPlanNo:      dayPlanNo,
		PlanDate:       planDate,
		MonthPlanID:    req.MonthPlanID,
		WorkshopID:     req.WorkshopID,
		ProductionLineID: req.ProductionLineID,
		ShiftType:      req.ShiftType,
		Status:         "DRAFT",
		KitStatus:      "PENDING",
		TotalPlanQty:   totalPlanQty,
		TotalProductCount: len(req.Items),
		CreatedBy:      &username,
		Remark:         &req.Remark,
	}

	if req.MonthPlanID != nil {
		// 获取月计划单号
		order.MonthPlanNo = toStrPtr(fmt.Sprintf("MP-%d", *req.MonthPlanID))
	}

	if err := s.repo.Create(ctx, order); err != nil {
		return nil, fmt.Errorf("创建日计划失败: %w", err)
	}

	// 创建明细
	for i, item := range req.Items {
		dayItem := &model.OrderDayItem{
			DayPlanID:       order.ID,
			LineNo:          i + 1,
			ProductID:       item.ProductID,
			ProductCode:     &item.ProductCode,
			ProductName:     &item.ProductName,
			Specification:  &item.Specification,
			Unit:            &item.Unit,
			PlanQty:         item.PlanQty,
			BOMID:           item.BOMID,
			ProcessRouteID:  item.ProcessRouteID,
			ProductionMode:  item.ProductionMode,
			MonthPlanItemID: item.MonthPlanItemID,
			ItemStatus:      "PENDING",
			KitStatus:       "PENDING",
			Priority:        item.Priority,
			Remark:          &item.Remark,
			TenantID:        tenantID,
		}
		if err := s.itemRepo.Create(ctx, dayItem); err != nil {
			return nil, fmt.Errorf("创建日计划明细失败: %w", err)
		}
	}

	return s.Get(ctx, order.ID)
}

func (s *OrderDayService) Update(ctx context.Context, id int64, req *model.OrderDayUpdate, username string) (*model.OrderDay, error) {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	// 仅草稿状态可更新
	if order.Status != "DRAFT" {
		return nil, errors.New("仅草稿状态可更新")
	}

	// 更新主表
	updates := map[string]interface{}{}
	if req.ProductionLineID != nil {
		updates["production_line_id"] = *req.ProductionLineID
	}
	if req.ShiftType != "" {
		updates["shift_type"] = req.ShiftType
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	// 重新计算汇总
	var totalPlanQty float64
	for _, item := range req.Items {
		totalPlanQty += item.PlanQty
	}
	updates["total_plan_qty"] = totalPlanQty
	updates["total_product_count"] = len(req.Items)

	if err := s.repo.Update(ctx, uint(id), updates); err != nil {
		return nil, fmt.Errorf("更新日计划失败: %w", err)
	}

	// 删除旧明细，创建新明细
	if err := s.itemRepo.DeleteByDayPlanID(ctx, id); err != nil {
		return nil, fmt.Errorf("删除旧明细失败: %w", err)
	}

	for i, item := range req.Items {
		dayItem := &model.OrderDayItem{
			DayPlanID:       int64(id),
			LineNo:          i + 1,
			ProductID:       item.ProductID,
			ProductCode:     &item.ProductCode,
			ProductName:     &item.ProductName,
			Specification:  &item.Specification,
			Unit:            &item.Unit,
			PlanQty:         item.PlanQty,
			BOMID:           item.BOMID,
			ProcessRouteID:  item.ProcessRouteID,
			ProductionMode:  item.ProductionMode,
			MonthPlanItemID: item.MonthPlanItemID,
			ItemStatus:      "PENDING",
			KitStatus:       "PENDING",
			Priority:        item.Priority,
			Remark:          &item.Remark,
			TenantID:        order.TenantID,
		}
		if err := s.itemRepo.Create(ctx, dayItem); err != nil {
			return nil, fmt.Errorf("创建日计划明细失败: %w", err)
		}
	}

	return s.Get(ctx, id)
}

func (s *OrderDayService) Delete(ctx context.Context, id int64) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	// 仅草稿状态可删除
	if order.Status != "DRAFT" {
		return errors.New("仅草稿状态可删除")
	}

	// 删除明细
	if err := s.itemRepo.DeleteByDayPlanID(ctx, id); err != nil {
		return err
	}

	return s.repo.Delete(ctx, uint(id))
}

func (s *OrderDayService) Publish(ctx context.Context, id int64, userID int64, username string) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.Status != "DRAFT" {
		return errors.New("仅草稿状态可发布")
	}

	if order.KitStatus != "READY" {
		return errors.New("齐套检查未通过，不能发布")
	}

	// 获取日计划明细
	items, err := s.itemRepo.ListByDayPlanID(ctx, id)
	if err != nil {
		return fmt.Errorf("获取日计划明细失败: %w", err)
	}

	now := time.Now()

	// 为每个明细行生成工单
	for _, item := range items {
		// 生成工单单号
		orderNo := fmt.Sprintf("WO-%s-%04d", order.PlanDate.Format("20060102"), item.LineNo)

		// 创建工单
		prodOrder := &model.ProductionOrder{
			TenantID:       order.TenantID,
			OrderNo:        orderNo,
			MaterialID:     item.ProductID,
			MaterialCode:   ptrStr(item.ProductCode),
			MaterialName:   ptrStr(item.ProductName),
			MaterialSpec:   item.Specification,
			Unit:           ptrStr(item.Unit),
			Quantity:       item.PlanQty,
			CompletedQty:   0,
			RejectedQty:    0,
			WorkshopID:     order.WorkshopID,
			WorkshopName:   order.WorkshopName,
			LineID:        0,
			RouteID:       0,
			BOMID:         0,
			PlanStartDate: &order.PlanDate,
			PlanEndDate:   &order.PlanDate,
			Status:        1, // 待生产
		}
		if order.ProductionLineID != nil {
			prodOrder.LineID = *order.ProductionLineID
		}
		if item.ProcessRouteID != nil {
			prodOrder.RouteID = *item.ProcessRouteID
		}
		if item.BOMID != nil {
			prodOrder.BOMID = *item.BOMID
		}
		if err := s.productionRepo.Create(ctx, prodOrder); err != nil {
			return fmt.Errorf("生成工单失败: %w", err)
		}

		// 记录日计划明细与工单的映射
		woMap := &model.OrderDayWorkOrderMap{
			DayPlanID:       order.ID,
			DayPlanItemID:   item.ID,
			WorkOrderID:     prodOrder.ID,
			WorkOrderNo:     orderNo,
			TenantID:        order.TenantID,
		}
		if err := s.workOrderMapRepo.Create(ctx, woMap); err != nil {
			return fmt.Errorf("记录工单映射失败: %w", err)
		}

		// 更新明细行的工单状态
		s.itemRepo.Update(ctx, uint(item.ID), map[string]interface{}{
			"item_status":    "PUBLISHED",
			"work_order_id":  prodOrder.ID,
			"work_order_no":  orderNo,
		})
	}

	// 更新日计划状态
	if err := s.repo.Update(ctx, uint(id), map[string]interface{}{
		"status":       "PUBLISHED",
		"published_at": now,
		"published_by": userID,
	}); err != nil {
		return err
	}

	return nil
}

func (s *OrderDayService) Complete(ctx context.Context, id int64) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.Status != "PUBLISHED" && order.Status != "IN_PRODUCTION" {
		return errors.New("只有已发布或生产中的日计划可完成")
	}

	return s.repo.Update(ctx, uint(id), map[string]interface{}{
		"status": "COMPLETED",
	})
}

func (s *OrderDayService) Terminate(ctx context.Context, id int64, userID int64, username string) error {
	order, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	if order.Status == "COMPLETED" || order.Status == "TERMINATED" {
		return errors.New("已完成的日计划不能终止")
	}

	return s.repo.Update(ctx, uint(id), map[string]interface{}{
		"status": "TERMINATED",
	})
}

func (s *OrderDayService) KitCheck(ctx context.Context, id int64, userID int64) error {
	_, err := s.repo.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	// TODO: 实际齐套检查逻辑
	// 这里简化处理，设置为READY
	now := time.Now()
	return s.repo.Update(ctx, uint(id), map[string]interface{}{
		"kit_status":   "READY",
		"kit_check_time": now,
		"kit_check_by": userID,
	})
}

func (s *OrderDayService) GetWorkOrderMaps(ctx context.Context, dayPlanID int64) ([]model.OrderDayWorkOrderMap, error) {
	return s.workOrderMapRepo.ListByDayPlanID(ctx, dayPlanID)
}

func (s *OrderDayService) GetByMonthPlanID(ctx context.Context, monthPlanID int64) ([]model.OrderDay, error) {
	return s.repo.GetOrderDaysByMonthID(ctx, monthPlanID)
}

// Helper function
func toPtr[T any](v T) *T {
	return &v
}

func toStrPtr(v string) *string {
	return &v
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
