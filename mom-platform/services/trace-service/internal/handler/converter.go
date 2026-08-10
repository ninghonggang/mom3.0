package handler

import (
	"fmt"
	"strconv"

	common "github.com/ninghonggang/mom-platform/gen/common"
	trace "github.com/ninghonggang/mom-platform/gen/trace"

	"mom-platform/services/trace-service/internal/model"
)

// --- model -> proto ---

func modelToProtoTraceRecord(m *model.TraceRecord) *trace.TraceRecord {
	return &trace.TraceRecord{
		Id:                int64(m.ID),
		TraceNo:           m.TraceNo,
		TraceType:         stringToTraceType(m.TraceType),
		SerialNo:          m.SerialNo,
		BatchNo:           m.BatchNo,
		Status:            stringToTraceRecordStatus(m.Status),
		TraceAt:           m.TraceAt.Unix(),
	}
}

func modelToProtoSerialNumber(m *model.SerialNumber) *trace.SerialNumber {
	return &trace.SerialNumber{
		Id:                int64(m.ID),
		SerialNo:          m.SerialNo,
		BatchNo:           m.BatchNo,
		Status:            m.Status,
		CreatedAt:         m.CreatedAt.Unix(),
	}
}

func modelToProtoDataPoint(m *model.DataPoint) *trace.DataPoint {
	return &trace.DataPoint{
		Id:                     int64(m.ID),
		TenantId:               parseInt64(m.TenantID),
		PointCode:              m.PointCode,
		PointName:              m.PointName,
		EquipmentId:            parseInt64(m.EquipmentID),
		DataType:               stringToDataType(m.DataType),
		UpperLimit:             float64PtrToString(m.UpperLimit),
		LowerLimit:             float64PtrToString(m.LowerLimit),
		CollectIntervalSeconds: int32(m.CollectIntervalSeconds),
		Status:                 stringToDataPointStatus(m.Status),
	}
}

// parseInt64 将数值型字符串字段转为 int64，非法或空值返回 0。
func parseInt64(s string) int64 {
	if s == "" {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func modelToProtoCollectRecord(m *model.CollectRecord) *trace.CollectRecord {
	return &trace.CollectRecord{
		Id:          int64(m.ID),
		DataPointId: int64(m.DataPointID),
		Value:       m.Value,
		Quality:     stringToDataQuality(m.Quality),
		CollectedAt: m.CollectedAt.Unix(),
	}
}

func modelToProtoScanLog(m *model.ScanLog) *trace.ScanLog {
	var traceID int64
	if m.TraceID != nil {
		traceID = int64(*m.TraceID)
	}
	return &trace.ScanLog{
		Id:          int64(m.ID),
		ScanCode:    m.ScanCode,
		ScanType:    m.ScanType,
		ScanTime:    m.ScanTime.Unix(),
		TraceId:     traceID,
	}
}

func protoPageResponse(page, pageSize, total, totalPages int32) *common.PageResponse {
	return &common.PageResponse{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
}

// --- string <-> proto enum helpers ---

func stringToTraceType(s string) trace.TraceType {
	switch s {
	case "SERIAL":
		return trace.TraceType_TRACE_TYPE_SERIAL
	case "BATCH":
		return trace.TraceType_TRACE_TYPE_BATCH
	case "ORDER":
		return trace.TraceType_TRACE_TYPE_ORDER
	case "MATERIAL":
		return trace.TraceType_TRACE_TYPE_MATERIAL
	default:
		return trace.TraceType_TRACE_TYPE_UNSPECIFIED
	}
}

func protoTraceTypeToString(t trace.TraceType) string {
	switch t {
	case trace.TraceType_TRACE_TYPE_SERIAL:
		return "SERIAL"
	case trace.TraceType_TRACE_TYPE_BATCH:
		return "BATCH"
	case trace.TraceType_TRACE_TYPE_ORDER:
		return "ORDER"
	case trace.TraceType_TRACE_TYPE_MATERIAL:
		return "MATERIAL"
	default:
		return ""
	}
}

func stringToTraceRecordStatus(s string) trace.TraceRecordStatus {
	switch s {
	case "PENDING":
		return trace.TraceRecordStatus_TRACE_STATUS_PENDING
	case "ACTIVE":
		return trace.TraceRecordStatus_TRACE_STATUS_ACTIVE
	case "BROKEN":
		return trace.TraceRecordStatus_TRACE_STATUS_BROKEN
	case "ARCHIVED":
		return trace.TraceRecordStatus_TRACE_STATUS_ARCHIVED
	default:
		return trace.TraceRecordStatus_TRACE_STATUS_UNSPECIFIED
	}
}

func stringToDataType(s string) trace.DataType {
	switch s {
	case "NUMBER":
		return trace.DataType_DATA_TYPE_NUMBER
	case "STRING":
		return trace.DataType_DATA_TYPE_STRING
	case "BOOLEAN":
		return trace.DataType_DATA_TYPE_BOOLEAN
	default:
		return trace.DataType_DATA_TYPE_UNSPECIFIED
	}
}

func protoDataTypeToString(d trace.DataType) string {
	switch d {
	case trace.DataType_DATA_TYPE_NUMBER:
		return "NUMBER"
	case trace.DataType_DATA_TYPE_STRING:
		return "STRING"
	case trace.DataType_DATA_TYPE_BOOLEAN:
		return "BOOLEAN"
	default:
		return ""
	}
}

func stringToDataPointStatus(s string) trace.DataPointStatus {
	switch s {
	case "ACTIVE":
		return trace.DataPointStatus_DATA_POINT_STATUS_ACTIVE
	case "PAUSED":
		return trace.DataPointStatus_DATA_POINT_STATUS_PAUSED
	case "ERROR":
		return trace.DataPointStatus_DATA_POINT_STATUS_ERROR
	default:
		return trace.DataPointStatus_DATA_POINT_STATUS_UNSPECIFIED
	}
}

// dataPointStatusToString 将 proto 枚举反向映射为 model 中的状态字符串；
// UNSPECIFIED 返回空串，表示"不按状态过滤"。
func dataPointStatusToString(s trace.DataPointStatus) string {
	switch s {
	case trace.DataPointStatus_DATA_POINT_STATUS_ACTIVE:
		return "ACTIVE"
	case trace.DataPointStatus_DATA_POINT_STATUS_PAUSED:
		return "PAUSED"
	case trace.DataPointStatus_DATA_POINT_STATUS_ERROR:
		return "ERROR"
	default:
		return ""
	}
}

func stringToDataQuality(s string) trace.DataQuality {
	switch s {
	case "GOOD":
		return trace.DataQuality_DATA_QUALITY_GOOD
	case "BAD":
		return trace.DataQuality_DATA_QUALITY_BAD
	case "UNCERTAIN":
		return trace.DataQuality_DATA_QUALITY_UNCERTAIN
	default:
		return trace.DataQuality_DATA_QUALITY_UNSPECIFIED
	}
}

func float64PtrToString(f *float64) string {
	if f == nil {
		return ""
	}
	return formatFloat(*f)
}

func formatFloat(f float64) string {
	s := fmt.Sprintf("%.6f", f)
	return s
}
