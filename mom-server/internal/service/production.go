package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type ProductionOrderService struct {
	repo            *repository.ProductionOrderRepository
	changeLogSvc    *ProductionOrderChangeLogService
}

func NewProductionOrderService(repo *repository.ProductionOrderRepository, changeLogSvc *ProductionOrderChangeLogService) *ProductionOrderService {
	return &ProductionOrderService{repo: repo, changeLogSvc: changeLogSvc}
}

func (s *ProductionOrderService) List(ctx context.Context) ([]model.ProductionOrder, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *ProductionOrderService) GetByID(ctx context.Context, id string) (*model.ProductionOrder, error) {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, orderID)
}

func (s *ProductionOrderService) Create(ctx context.Context, order *model.ProductionOrder) error {
	return s.repo.Create(ctx, order)
}

func (s *ProductionOrderService) Update(ctx context.Context, id string, order *model.ProductionOrder, changedBy string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}

	// 获取旧数据用于记录变更
	oldOrder, err := s.repo.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{}

	// 检查并记录数量变更
	if oldOrder.Quantity != order.Quantity {
		s.changeLogSvc.RecordChange(ctx, int64(orderID), order.OrderNo, ChangeTypeQuantityChange, oldOrder.Quantity, order.Quantity, "", changedBy)
		updates["quantity"] = order.Quantity
	}

	// 检查并记录优先级变更
	if oldOrder.Priority != order.Priority {
		s.changeLogSvc.RecordChange(ctx, int64(orderID), order.OrderNo, ChangeTypePriorityChange, oldOrder.Priority, order.Priority, "", changedBy)
		updates["priority"] = order.Priority
	}

	// 检查并记录产线变更
	if oldOrder.LineID != order.LineID {
		s.changeLogSvc.RecordChange(ctx, int64(orderID), order.OrderNo, ChangeTypeLineChange, oldOrder.LineID, order.LineID, "", changedBy)
		updates["line_id"] = order.LineID
	}

	// 检查并记录状态变更
	if oldOrder.Status != order.Status {
		s.changeLogSvc.RecordChange(ctx, int64(orderID), order.OrderNo, ChangeTypeStatusChange, oldOrder.Status, order.Status, "", changedBy)
		updates["status"] = order.Status
	}

	// 更新其他字段
	if order.MaterialID != 0 {
		updates["material_id"] = order.MaterialID
	}
	if order.WorkshopID != 0 {
		updates["workshop_id"] = order.WorkshopID
	}
	updates["order_no"] = order.OrderNo

	if len(updates) > 0 {
		return s.repo.Update(ctx, orderID, updates)
	}
	return nil
}

func (s *ProductionOrderService) Delete(ctx context.Context, id string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, orderID)
}

func (s *ProductionOrderService) Start(ctx context.Context, id string, changedBy string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}

	// 获取旧数据
	oldOrder, err := s.repo.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	// 记录状态变更
	s.changeLogSvc.RecordChange(ctx, int64(orderID), oldOrder.OrderNo, ChangeTypeStatusChange, oldOrder.Status, 2, "开始生产", changedBy)

	return s.repo.Update(ctx, orderID, map[string]interface{}{
		"status": 2, // 生产中
	})
}

func (s *ProductionOrderService) Complete(ctx context.Context, id string, changedBy string) error {
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		return err
	}

	// 获取旧数据
	oldOrder, err := s.repo.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	// 记录状态变更
	s.changeLogSvc.RecordChange(ctx, int64(orderID), oldOrder.OrderNo, ChangeTypeStatusChange, oldOrder.Status, 3, "完成生产", changedBy)

	return s.repo.Update(ctx, orderID, map[string]interface{}{
		"status": 3, // 已完成
	})
}

// ========== 看板服务 ==========

// KanbanService 看板服务
type KanbanService struct {
	productionOrderRepo *repository.ProductionOrderRepository
	productionLineRepo *repository.ProductionLineRepository
	productionReportRepo *repository.ProductionReportRepository
}

func NewKanbanService(
	productionOrderRepo *repository.ProductionOrderRepository,
	productionLineRepo *repository.ProductionLineRepository,
	productionReportRepo *repository.ProductionReportRepository,
) *KanbanService {
	return &KanbanService{
		productionOrderRepo: productionOrderRepo,
		productionLineRepo: productionLineRepo,
		productionReportRepo: productionReportRepo,
	}
}

// KanbanDashboard 看板数据结构
type KanbanDashboard struct {
	ProductionLines []LineStatus `json:"production_lines"`
	OutputStats     OutputStats  `json:"output_stats"`
	OrderProgress   OrderProgress `json:"order_progress"`
	HourlyOutput   []HourlyOutput `json:"hourly_output"`
}

// LineStatus 产线状态
type LineStatus struct {
	LineID       int64   `json:"line_id"`
	LineName     string  `json:"line_name"`
	WorkshopName string  `json:"workshop_name"`
	Status       string  `json:"status"` // running/idle/fault
	StatusText   string  `json:"status_text"`
	Output       float64 `json:"output"`
	TargetOutput float64 `json:"target_output"`
	Completion   float64 `json:"completion"`
}

// OutputStats 产量统计
type OutputStats struct {
	TodayOutput   float64 `json:"today_output"`
	WeekOutput    float64 `json:"week_output"`
	MonthOutput   float64 `json:"month_output"`
	QualifiedRate float64 `json:"qualified_rate"`
}

// OrderProgress 工单进度
type OrderProgress struct {
	Total     int `json:"total"`
	Pending   int `json:"pending"`
	InProcess int `json:"in_process"`
	Completed int `json:"completed"`
	Cancelled int `json:"cancelled"`
}

// HourlyOutput 每小时产量
type HourlyOutput struct {
	Hour   string  `json:"hour"`
	Output float64 `json:"output"`
}

func (s *KanbanService) GetDashboardData(ctx context.Context) (*KanbanDashboard, error) {
	dashboard := &KanbanDashboard{
		ProductionLines: make([]LineStatus, 0),
		OutputStats:     OutputStats{},
		OrderProgress:   OrderProgress{},
		HourlyOutput:    make([]HourlyOutput, 0),
	}

	// 获取产线列表
	lines, _, err := s.productionLineRepo.List(ctx, 0)
	if err != nil {
		return nil, err
	}

	// 获取工单列表
	orders, _, err := s.productionOrderRepo.List(ctx, 0)
	if err != nil {
		return nil, err
	}

	// 统计数据
	orderProgress := OrderProgress{Total: len(orders)}
	outputStats := OutputStats{}
	hourlyOutput := make([]HourlyOutput, 0)

	// 模拟当前时间的小时产量数据
	now := time.Now()
	r := rand.New(rand.NewSource(int64(now.Hour())))
	for i := 0; i <= now.Hour(); i++ {
		hour := fmt.Sprintf("%02d:00", i)
		baseOutput := float64(80 + r.Intn(40))
		hourlyOutput = append(hourlyOutput, HourlyOutput{
			Hour:   hour,
			Output: baseOutput,
		})
		outputStats.TodayOutput += baseOutput
	}

	// 计算周和月产量（模拟数据）
	outputStats.WeekOutput = outputStats.TodayOutput * float64(7-now.Weekday()) / float64(now.Weekday()+1)
	outputStats.MonthOutput = outputStats.TodayOutput * 30 / float64(now.Day())
	outputStats.QualifiedRate = 95.5 + r.Float64()*3

	// 产线状态
	statuses := []string{"running", "idle", "fault"}
	statusTexts := map[string]string{"running": "运行中", "idle": "待机", "fault": "故障"}
	workshops := []string{"一车间", "二车间", "三车间"}

	for i, line := range lines {
		status := statuses[i%3]
		targetOutput := 1000.0
		actualOutput := targetOutput * (0.6 + r.Float64()*0.4)

		dashboard.ProductionLines = append(dashboard.ProductionLines, LineStatus{
			LineID:       line.ID,
			LineName:     line.LineName,
			WorkshopName: workshops[i%len(workshops)],
			Status:       status,
			StatusText:   statusTexts[status],
			Output:       actualOutput,
			TargetOutput: targetOutput,
			Completion:   actualOutput / targetOutput * 100,
		})
	}

	// 如果没有产线数据，添加模拟数据
	if len(dashboard.ProductionLines) == 0 {
		lineNames := []string{"A线", "B线", "C线", "D线"}
		for i, name := range lineNames {
			status := statuses[i%3]
			targetOutput := 1000.0
			actualOutput := targetOutput * (0.6 + r.Float64()*0.4)
			dashboard.ProductionLines = append(dashboard.ProductionLines, LineStatus{
				LineID:       int64(i + 1),
				LineName:     name,
				WorkshopName: workshops[i%len(workshops)],
				Status:       status,
				StatusText:   statusTexts[status],
				Output:       actualOutput,
				TargetOutput: targetOutput,
				Completion:   actualOutput / targetOutput * 100,
			})
		}
	}

	// 工单进度统计
	for _, order := range orders {
		switch order.Status {
		case 1:
			orderProgress.Pending++
		case 2:
			orderProgress.InProcess++
		case 3:
			orderProgress.Completed++
		case 4:
			orderProgress.Cancelled++
		}
	}

	// 如果没有工单数据，添加模拟数据
	if orderProgress.Total == 0 {
		orderProgress = OrderProgress{
			Total:     12,
			Pending:   3,
			InProcess: 5,
			Completed: 3,
			Cancelled: 1,
		}
	}

	dashboard.OutputStats = outputStats
	dashboard.OrderProgress = orderProgress
	dashboard.HourlyOutput = hourlyOutput

	return dashboard, nil
}
