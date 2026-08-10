package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"mom-platform/services/trace-service/internal/model"
	"mom-platform/services/trace-service/internal/repository"
)

type TraceService struct {
	repo   repository.Repository
	redis  *redis.Client
	logger *zap.Logger
}

func NewTraceService(repo repository.Repository, redis *redis.Client, logger *zap.Logger) *TraceService {
	return &TraceService{repo: repo, redis: redis, logger: logger}
}

func (s *TraceService) GetRepo() repository.Repository {
	return s.repo
}

// CreateTraceRecord 创建追溯节点并链接到父节点
func (s *TraceService) CreateTraceRecord(ctx context.Context, tenantID, traceType, serialNo, batchNo, materialID, orderID, parentSerialNo, parentBatchNo string) (*model.TraceRecord, error) {
	now := time.Now()
	traceNo := fmt.Sprintf("TR-%s-%d", strings.ReplaceAll(tenantID, "-", ""), now.UnixNano())

	record := &model.TraceRecord{
		TenantID:          tenantID,
		TraceNo:           traceNo,
		TraceType:         traceType,
		SerialNo:          serialNo,
		BatchNo:           batchNo,
		MaterialID:        materialID,
		ProductionOrderID: orderID,
		Status:            "ACTIVE",
		TraceAt:           now,
		CreatedAt:         now,
	}

	if err := s.repo.CreateTraceRecord(ctx, record); err != nil {
		return nil, fmt.Errorf("create trace record: %w", err)
	}

	// Link to parent by serial_no or batch_no
	var parentTrace *model.TraceRecord
	if parentSerialNo != "" {
		parentTrace, _ = s.repo.GetTraceRecordBySerialNo(ctx, parentSerialNo)
	} else if parentBatchNo != "" {
		records, _ := s.repo.GetTraceRecordByBatchNo(ctx, parentBatchNo)
		if len(records) > 0 {
			parentTrace = records[0]
		}
	}

	if parentTrace != nil {
		link := &model.TraceLink{
			TraceID:       record.ID,
			ParentTraceID: &parentTrace.ID,
			LinkType:      "MATERIAL",
			Level:         1,
		}
		if err := s.repo.CreateTraceLink(ctx, link); err != nil {
			s.logger.Warn("failed to create trace link", zap.Error(err))
		}
		s.logger.Info("trace linked to parent",
			zap.String("trace_no", traceNo),
			zap.String("parent_trace_no", parentTrace.TraceNo),
		)
	}

	s.logger.Info("trace record created", zap.String("trace_no", traceNo))
	return record, nil
}

// ForwardTrace 前向追溯：从物料向前到成品 (BFS)
func (s *TraceService) ForwardTrace(ctx context.Context, startTraceID uint) ([]*model.TraceRecord, error) {
	visits := map[uint]bool{}
	queue := []uint{startTraceID}
	var result []*model.TraceRecord

	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]

		if visits[currentID] {
			continue
		}
		visits[currentID] = true

		record, err := s.repo.GetTraceRecord(ctx, currentID)
		if err != nil {
			s.logger.Warn("failed to get trace record", zap.Uint("id", currentID), zap.Error(err))
			continue
		}
		result = append(result, record)

		// 找到以此节点为 parent 的子链接
		links, err := s.repo.GetTraceLinksByParentTraceID(ctx, currentID)
		if err != nil {
			continue
		}
		for _, link := range links {
			if !visits[link.TraceID] {
				queue = append(queue, link.TraceID)
			}
		}
	}

	return result, nil
}

// BackwardTrace 后向追溯：从成品回溯到原料
func (s *TraceService) BackwardTrace(ctx context.Context, startTraceID uint) ([]*model.TraceRecord, error) {
	visits := map[uint]bool{}
	queue := []uint{startTraceID}
	var result []*model.TraceRecord

	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]

		if visits[currentID] {
			continue
		}
		visits[currentID] = true

		record, err := s.repo.GetTraceRecord(ctx, currentID)
		if err != nil {
			s.logger.Warn("failed to get trace record", zap.Uint("id", currentID), zap.Error(err))
			continue
		}
		result = append(result, record)

		// 找到以此节点为 trace 的链接，追溯其 parent
		links, err := s.repo.GetTraceLinksByTraceID(ctx, currentID)
		if err != nil {
			continue
		}
		for _, link := range links {
			if link.ParentTraceID != nil && !visits[*link.ParentTraceID] {
				queue = append(queue, *link.ParentTraceID)
			}
		}
	}

	return result, nil
}

// ListTraces 分页列表追溯记录
func (s *TraceService) ListTraces(ctx context.Context, page, pageSize int, traceType, materialID string, beginTime, endTime int64) ([]*model.TraceRecord, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListTraceRecords(ctx, offset, pageSize, traceType, materialID, beginTime, endTime)
}

// GenerateSerials 批量生成序列号
func (s *TraceService) GenerateSerials(ctx context.Context, prefix string, count int, materialID, orderID, batchNo string) ([]*model.SerialNumber, error) {
	now := time.Now()
	sns := make([]*model.SerialNumber, 0, count)

	// 同一 prefix + 日期可能被多次调用，必须从已有的最大流水号续号，
	// 否则会撞上 idx_serial_numbers_serial_no 唯一索引。
	base := fmt.Sprintf("%s-%s-", prefix, now.Format("20060102"))
	startSeq, err := s.repo.MaxSerialSeq(ctx, base)
	if err != nil {
		return nil, fmt.Errorf("query max serial seq for %q: %w", base, err)
	}

	for i := 1; i <= count; i++ {
		serialNo := fmt.Sprintf("%s%06d", base, startSeq+i)

		sns = append(sns, &model.SerialNumber{
			SerialNo:          serialNo,
			MaterialID:        materialID,
			ProductionOrderID: orderID,
			BatchNo:           batchNo,
			Status:            "UNUSED",
			CreatedAt:         now,
		})
	}

	if err := s.repo.BatchCreateSerialNumbers(ctx, sns); err != nil {
		return nil, fmt.Errorf("batch create serial numbers: %w", err)
	}

	s.logger.Info("generated serial numbers",
		zap.String("batch_no", batchNo),
		zap.Int("count", count),
	)
	return sns, nil
}

// CreateDataPoint 注册数据采集点
func (s *TraceService) CreateDataPoint(ctx context.Context, tenantID, pointCode, pointName, equipmentID, dataType string, upperLimit, lowerLimit *float64, collectIntervalSeconds int) (*model.DataPoint, error) {
	dp := &model.DataPoint{
		TenantID:               tenantID,
		PointCode:              pointCode,
		PointName:              pointName,
		EquipmentID:            equipmentID,
		DataType:               dataType,
		UpperLimit:             upperLimit,
		LowerLimit:             lowerLimit,
		CollectIntervalSeconds: collectIntervalSeconds,
		Status:                 "ACTIVE",
	}

	if err := s.repo.CreateDataPoint(ctx, dp); err != nil {
		return nil, fmt.Errorf("create data point: %w", err)
	}

	s.logger.Info("data point created", zap.String("point_code", pointCode))
	return dp, nil
}

// CollectData 采集数据，校验上下限，判定质量
func (s *TraceService) CollectData(ctx context.Context, tenantID string, dataPointID uint, rawValue string) (*model.CollectRecord, error) {
	dp, err := s.repo.GetDataPoint(ctx, dataPointID)
	if err != nil {
		return nil, fmt.Errorf("get data point: %w", err)
	}

	quality := "GOOD"

	if dp.DataType == "NUMBER" {
		val, parseErr := strconv.ParseFloat(rawValue, 64)
		if parseErr == nil {
			if dp.UpperLimit != nil && val > *dp.UpperLimit {
				quality = "BAD"
				s.logger.Warn("value exceeds upper limit",
					zap.String("point_code", dp.PointCode),
					zap.Float64("value", val),
					zap.Float64("upper_limit", *dp.UpperLimit),
				)
			} else if dp.LowerLimit != nil && val < *dp.LowerLimit {
				quality = "BAD"
				s.logger.Warn("value below lower limit",
					zap.String("point_code", dp.PointCode),
					zap.Float64("value", val),
					zap.Float64("lower_limit", *dp.LowerLimit),
				)
			}
		} else {
			quality = "UNCERTAIN"
		}
	}

	now := time.Now()
	cr := &model.CollectRecord{
		TenantID:    tenantID,
		DataPointID: dataPointID,
		Value:       rawValue,
		Quality:     quality,
		CollectedAt: now,
	}

	if err := s.repo.CreateCollectRecord(ctx, cr); err != nil {
		return nil, fmt.Errorf("create collect record: %w", err)
	}

	if quality == "BAD" {
		s.logger.Warn("threshold alert triggered",
			zap.String("point_code", dp.PointCode),
			zap.String("value", rawValue),
		)
	}

	return cr, nil
}

// ListCollectRecords 分页列表采集记录
func (s *TraceService) ListCollectRecords(ctx context.Context, dataPointID int64, beginTime, endTime int64, page, pageSize int) ([]*model.CollectRecord, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListCollectRecords(ctx, dataPointID, beginTime, endTime, offset, pageSize)
}

// ListDataPoints 分页查询采集点，可按设备与状态过滤。
func (s *TraceService) ListDataPoints(ctx context.Context, equipmentID int64, statusFilter string, page, pageSize int) ([]*model.DataPoint, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.ListDataPoints(ctx, equipmentID, statusFilter, offset, pageSize)
}

// CreateScanLog 记录扫码日志，匹配 serial_no 自动创建追溯记录
func (s *TraceService) CreateScanLog(ctx context.Context, tenantID, scanCode, scanType, operatorID, equipmentID, workstationID string) (*model.ScanLog, *model.TraceRecord, error) {
	now := time.Now()
	sl := &model.ScanLog{
		TenantID:      tenantID,
		ScanCode:      scanCode,
		ScanType:      scanType,
		OperatorID:    operatorID,
		EquipmentID:   equipmentID,
		WorkstationID: workstationID,
		ScanTime:      now,
	}

	// 尝试通过 serial_no 匹配已有 TraceRecord
	traceRec, err := s.repo.GetTraceRecordBySerialNo(ctx, scanCode)
	if err == nil && traceRec != nil {
		sl.TraceID = &traceRec.ID
	} else {
		// 未找到，自动创建新的 TraceRecord
		tn := fmt.Sprintf("TR-SCAN-%d", now.UnixNano())
		newTR := &model.TraceRecord{
			TenantID:  tenantID,
			TraceNo:   tn,
			TraceType: "SERIAL",
			SerialNo:  scanCode,
			Status:    "PENDING",
			TraceAt:   now,
			CreatedAt: now,
		}
		if err := s.repo.CreateTraceRecord(ctx, newTR); err != nil {
			return nil, nil, fmt.Errorf("create trace record from scan: %w", err)
		}
		traceRec = newTR
		sl.TraceID = &newTR.ID
	}

	if err := s.repo.CreateScanLog(ctx, sl); err != nil {
		return nil, nil, fmt.Errorf("create scan log: %w", err)
	}

	s.logger.Info("scan log created",
		zap.String("scan_code", scanCode),
		zap.String("scan_type", scanType),
	)

	return sl, traceRec, nil
}
