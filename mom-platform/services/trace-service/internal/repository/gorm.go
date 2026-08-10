package repository

import (
	"context"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"mom-platform/services/trace-service/internal/model"
)

type gormRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewGormRepository(db *gorm.DB, logger *zap.Logger) Repository {
	return &gormRepo{db: db, logger: logger}
}

// TraceRecord
func (r *gormRepo) CreateTraceRecord(ctx context.Context, record *model.TraceRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *gormRepo) GetTraceRecord(ctx context.Context, id uint) (*model.TraceRecord, error) {
	var record model.TraceRecord
	err := r.db.WithContext(ctx).First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *gormRepo) GetTraceRecordBySerialNo(ctx context.Context, serialNo string) (*model.TraceRecord, error) {
	var record model.TraceRecord
	err := r.db.WithContext(ctx).Where("serial_no = ?", serialNo).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *gormRepo) GetTraceRecordByBatchNo(ctx context.Context, batchNo string) ([]*model.TraceRecord, error) {
	var records []*model.TraceRecord
	err := r.db.WithContext(ctx).Where("batch_no = ?", batchNo).Find(&records).Error
	return records, err
}

func (r *gormRepo) UpdateTraceRecord(ctx context.Context, record *model.TraceRecord) error {
	return r.db.WithContext(ctx).Save(record).Error
}

func (r *gormRepo) ListTraceRecords(ctx context.Context, offset, limit int, traceType, materialID string, beginTime, endTime int64) ([]*model.TraceRecord, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.TraceRecord{})

	if traceType != "" {
		q = q.Where("trace_type = ?", traceType)
	}
	if materialID != "" {
		q = q.Where("material_id = ?", materialID)
	}
	if beginTime > 0 {
		q = q.Where("trace_at >= ?", beginTime)
	}
	if endTime > 0 {
		q = q.Where("trace_at <= ?", endTime)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var records []*model.TraceRecord
	err := q.Offset(offset).Limit(limit).Order("trace_at DESC").Find(&records).Error
	return records, total, err
}

// TraceLink
func (r *gormRepo) CreateTraceLink(ctx context.Context, link *model.TraceLink) error {
	return r.db.WithContext(ctx).Create(link).Error
}

func (r *gormRepo) GetTraceLinksByTraceID(ctx context.Context, traceID uint) ([]*model.TraceLink, error) {
	var links []*model.TraceLink
	err := r.db.WithContext(ctx).Where("trace_id = ?", traceID).Find(&links).Error
	return links, err
}

func (r *gormRepo) GetTraceLinksByParentTraceID(ctx context.Context, parentTraceID uint) ([]*model.TraceLink, error) {
	var links []*model.TraceLink
	err := r.db.WithContext(ctx).Where("parent_trace_id = ?", parentTraceID).Find(&links).Error
	return links, err
}

func (r *gormRepo) GetAllTraceLinks(ctx context.Context) ([]*model.TraceLink, error) {
	var links []*model.TraceLink
	err := r.db.WithContext(ctx).Find(&links).Error
	return links, err
}

// SerialNumber
func (r *gormRepo) CreateSerialNumber(ctx context.Context, sn *model.SerialNumber) error {
	return r.db.WithContext(ctx).Create(sn).Error
}

func (r *gormRepo) BatchCreateSerialNumbers(ctx context.Context, sns []*model.SerialNumber) error {
	return r.db.WithContext(ctx).Create(&sns).Error
}

func (r *gormRepo) GetSerialNumber(ctx context.Context, serialNo string) (*model.SerialNumber, error) {
	var sn model.SerialNumber
	err := r.db.WithContext(ctx).Where("serial_no = ?", serialNo).First(&sn).Error
	if err != nil {
		return nil, err
	}
	return &sn, nil
}

// MaxSerialSeq 取 base 前缀下已存在的最大流水号。序列号格式为 <base><000001>，
// 流水号定长补零，因此字符串倒序的第一条即为最大值。
func (r *gormRepo) MaxSerialSeq(ctx context.Context, base string) (int, error) {
	var nos []string
	err := r.db.WithContext(ctx).
		Model(&model.SerialNumber{}).
		Where("serial_no LIKE ?", base+"%").
		Order("serial_no DESC").
		Limit(1).
		Pluck("serial_no", &nos).Error
	if err != nil {
		return 0, err
	}
	if len(nos) == 0 {
		return 0, nil
	}
	seq, convErr := strconv.Atoi(strings.TrimPrefix(nos[0], base))
	if convErr != nil {
		return 0, nil
	}
	return seq, nil
}

// DataPoint
func (r *gormRepo) CreateDataPoint(ctx context.Context, dp *model.DataPoint) error {
	return r.db.WithContext(ctx).Create(dp).Error
}

func (r *gormRepo) GetDataPoint(ctx context.Context, id uint) (*model.DataPoint, error) {
	var dp model.DataPoint
	err := r.db.WithContext(ctx).First(&dp, id).Error
	if err != nil {
		return nil, err
	}
	return &dp, nil
}

func (r *gormRepo) GetDataPointByCode(ctx context.Context, code string) (*model.DataPoint, error) {
	var dp model.DataPoint
	err := r.db.WithContext(ctx).Where("point_code = ?", code).First(&dp).Error
	if err != nil {
		return nil, err
	}
	return &dp, nil
}

// ListDataPoints 分页查询采集点，可按设备与状态过滤。
func (r *gormRepo) ListDataPoints(ctx context.Context, equipmentID int64, statusFilter string, offset, limit int) ([]*model.DataPoint, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.DataPoint{})

	if equipmentID > 0 {
		q = q.Where("equipment_id = ?", strconv.FormatInt(equipmentID, 10))
	}
	if statusFilter != "" {
		q = q.Where("status = ?", statusFilter)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var points []*model.DataPoint
	err := q.Offset(offset).Limit(limit).Order("id DESC").Find(&points).Error
	return points, total, err
}

// CollectRecord
func (r *gormRepo) CreateCollectRecord(ctx context.Context, cr *model.CollectRecord) error {
	return r.db.WithContext(ctx).Create(cr).Error
}

func (r *gormRepo) ListCollectRecords(ctx context.Context, dataPointID int64, beginTime, endTime int64, offset, limit int) ([]*model.CollectRecord, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.CollectRecord{})

	if dataPointID > 0 {
		q = q.Where("data_point_id = ?", dataPointID)
	}
	if beginTime > 0 {
		q = q.Where("collected_at >= ?", beginTime)
	}
	if endTime > 0 {
		q = q.Where("collected_at <= ?", endTime)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var records []*model.CollectRecord
	err := q.Offset(offset).Limit(limit).Order("collected_at DESC").Find(&records).Error
	return records, total, err
}

// ScanLog
func (r *gormRepo) CreateScanLog(ctx context.Context, sl *model.ScanLog) error {
	return r.db.WithContext(ctx).Create(sl).Error
}

// Ensure unused import compiles
var _ = strconv.FormatInt
