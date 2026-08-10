package repository

import (
	"context"

	"mom-platform/services/trace-service/internal/model"
)

type Repository interface {
	// TraceRecord
	CreateTraceRecord(ctx context.Context, record *model.TraceRecord) error
	GetTraceRecord(ctx context.Context, id uint) (*model.TraceRecord, error)
	GetTraceRecordBySerialNo(ctx context.Context, serialNo string) (*model.TraceRecord, error)
	GetTraceRecordByBatchNo(ctx context.Context, batchNo string) ([]*model.TraceRecord, error)
	UpdateTraceRecord(ctx context.Context, record *model.TraceRecord) error
	ListTraceRecords(ctx context.Context, offset, limit int, traceType, materialID string, beginTime, endTime int64) ([]*model.TraceRecord, int64, error)

	// TraceLink
	CreateTraceLink(ctx context.Context, link *model.TraceLink) error
	GetTraceLinksByTraceID(ctx context.Context, traceID uint) ([]*model.TraceLink, error)
	GetTraceLinksByParentTraceID(ctx context.Context, parentTraceID uint) ([]*model.TraceLink, error)
	GetAllTraceLinks(ctx context.Context) ([]*model.TraceLink, error)

	// SerialNumber
	CreateSerialNumber(ctx context.Context, sn *model.SerialNumber) error
	BatchCreateSerialNumbers(ctx context.Context, sns []*model.SerialNumber) error
	GetSerialNumber(ctx context.Context, serialNo string) (*model.SerialNumber, error)
	// MaxSerialSeq 返回以 base 开头的序列号中最大的流水号（不存在时返回 0）。
	// 用于同一前缀+日期下多次生成时续号，避免 serial_no 唯一索引冲突。
	MaxSerialSeq(ctx context.Context, base string) (int, error)

	// DataPoint
	CreateDataPoint(ctx context.Context, dp *model.DataPoint) error
	GetDataPoint(ctx context.Context, id uint) (*model.DataPoint, error)
	GetDataPointByCode(ctx context.Context, code string) (*model.DataPoint, error)
	ListDataPoints(ctx context.Context, equipmentID int64, statusFilter string, offset, limit int) ([]*model.DataPoint, int64, error)

	// CollectRecord
	CreateCollectRecord(ctx context.Context, cr *model.CollectRecord) error
	ListCollectRecords(ctx context.Context, dataPointID int64, beginTime, endTime int64, offset, limit int) ([]*model.CollectRecord, int64, error)

	// ScanLog
	CreateScanLog(ctx context.Context, sl *model.ScanLog) error
}
