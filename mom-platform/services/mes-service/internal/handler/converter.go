package handler

import (
	"time"

	"github.com/ninghonggang/mom-platform/gen/common"
	mes "github.com/ninghonggang/mom-platform/gen/mes"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// --- Timestamp helpers ---

func timeToProto(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func timeFromProto(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

// --- BaseModel helper ---

func baseModelToProto(id, tenantId int64, createdAt, updatedAt time.Time, deletedAt gorm.DeletedAt) *common.BaseModel {
	base := &common.BaseModel{
		Id:        id,
		TenantId:  tenantId,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}
	if deletedAt.Valid {
		base.DeletedAt = timestamppb.New(deletedAt.Time)
	}
	return base
}

// --- Pagination helpers ---

func paginationFromProto(p *common.Pagination) (page, pageSize int) {
	if p != nil {
		return int(p.Page), int(p.PageSize)
	}
	return 1, 20
}

func paginationToProto(total int64, page, pageSize int) *common.Pagination {
	t := int32(total)
	var totalPages int32
	if pageSize > 0 {
		totalPages = t / int32(pageSize)
		if t%int32(pageSize) != 0 {
			totalPages++
		}
	}
	return &common.Pagination{
		Page:       int32(page),
		PageSize:   int32(pageSize),
		Total:      t,
		TotalPages: totalPages,
	}
}

// --- ProductionOrder conversions ---

func orderToProto(o *model.ProductionOrder) *mes.ProductionOrder {
	if o == nil {
		return nil
	}
	return &mes.ProductionOrder{
		Base:            baseModelToProto(o.ID, o.TenantID, o.CreatedAt, o.UpdatedAt, o.DeletedAt),
		OrderNo:         o.OrderNo,
		SalesOrderNo:    o.SalesOrderNo,
		MaterialId:      o.MaterialID,
		MaterialCode:    o.MaterialCode,
		MaterialName:    o.MaterialName,
		MaterialSpec:    o.MaterialSpec,
		Quantity:        o.Quantity,
		CompletedQty:    o.CompletedQty,
		RejectedQty:     o.RejectedQty,
		WorkshopId:      o.WorkshopID,
		WorkshopName:    o.WorkshopName,
		LineId:          o.LineID,
		LineName:        o.LineName,
		Status:          common.ProductionOrderStatus(o.Status),
		Priority:        common.Priority(o.Priority),
		PlanStartDate:   timeToProto(o.PlanStartDate),
		PlanEndDate:     timeToProto(o.PlanEndDate),
		ActualStartDate: timeToProto(o.ActualStartDate),
		ActualEndDate:   timeToProto(o.ActualEndDate),
		Remark:          o.Remark,
	}
}

func orderFromProto(o *mes.ProductionOrder) *model.ProductionOrder {
	if o == nil {
		return nil
	}
	m := &model.ProductionOrder{
		OrderNo:         o.OrderNo,
		SalesOrderNo:    o.SalesOrderNo,
		MaterialID:      o.MaterialId,
		MaterialCode:    o.MaterialCode,
		MaterialName:    o.MaterialName,
		MaterialSpec:    o.MaterialSpec,
		Quantity:        o.Quantity,
		CompletedQty:    o.CompletedQty,
		RejectedQty:     o.RejectedQty,
		WorkshopID:      o.WorkshopId,
		WorkshopName:    o.WorkshopName,
		LineID:          o.LineId,
		LineName:        o.LineName,
		Status:          model.OrderStatus(o.Status),
		Priority:        model.Priority(o.Priority),
		PlanStartDate:   timeFromProto(o.PlanStartDate),
		PlanEndDate:     timeFromProto(o.PlanEndDate),
		ActualStartDate: timeFromProto(o.ActualStartDate),
		ActualEndDate:   timeFromProto(o.ActualEndDate),
		Remark:          o.Remark,
	}
	if o.Base != nil {
		m.ID = o.Base.Id
		m.TenantID = o.Base.TenantId
	}
	return m
}

// --- Dispatch conversions ---

func dispatchToProto(d *model.Dispatch) *mes.Dispatch {
	if d == nil {
		return nil
	}
	return &mes.Dispatch{
		Base:          baseModelToProto(d.ID, d.TenantID, d.CreatedAt, d.UpdatedAt, d.DeletedAt),
		OrderId:       d.OrderID,
		OrderNo:       d.OrderNo,
		LineId:        d.LineID,
		WorkstationId: d.WorkstationID,
		ProcessId:     d.ProcessID,
		OperationId:   d.OperationID,
		OperationName: d.OperationName,
		EmployeeId:    d.EmployeeID,
		EmployeeName:  d.EmployeeName,
		PlannedQty:    d.PlannedQty,
		CompletedQty:  d.CompletedQty,
		PlanStartTime: timeToProto(d.PlanStartTime),
		PlanEndTime:   timeToProto(d.PlanEndTime),
		Status:        common.DispatchStatus(d.Status),
	}
}

func dispatchFromProto(d *mes.Dispatch) *model.Dispatch {
	if d == nil {
		return nil
	}
	m := &model.Dispatch{
		OrderID:       d.OrderId,
		OrderNo:       d.OrderNo,
		LineID:        d.LineId,
		WorkstationID: d.WorkstationId,
		ProcessID:     d.ProcessId,
		OperationID:   d.OperationId,
		OperationName: d.OperationName,
		EmployeeID:    d.EmployeeId,
		EmployeeName:  d.EmployeeName,
		PlannedQty:    d.PlannedQty,
		CompletedQty:  d.CompletedQty,
		PlanStartTime: timeFromProto(d.PlanStartTime),
		PlanEndTime:   timeFromProto(d.PlanEndTime),
		Status:        model.DispatchStatus(d.Status),
	}
	if d.Base != nil {
		m.ID = d.Base.Id
		m.TenantID = d.Base.TenantId
	}
	return m
}

// --- MobileJobReport conversions ---

func reportToProto(r *model.MobileJobReport) *mes.MobileJobReport {
	if r == nil {
		return nil
	}
	return &mes.MobileJobReport{
		Base:          baseModelToProto(r.ID, r.TenantID, r.CreatedAt, r.UpdatedAt, r.DeletedAt),
		OrderId:       r.OrderID,
		OrderNo:       r.OrderNo,
		ProcessId:     r.ProcessID,
		OperationId:   r.OperationID,
		OperationName: r.OperationName,
		EmployeeId:    r.EmployeeID,
		EmployeeName:  r.EmployeeName,
		WorkstationId: r.WorkstationID,
		ReportedQty:   r.ReportedQty,
		QualifiedQty:  r.QualifiedQty,
		DefectiveQty:  r.DefectiveQty,
		WorkMinutes:   int32(r.WorkMinutes),
		ReportType:    common.ReportType(r.ReportType),
		Status:        common.ReportStatus(r.Status),
		StartTime:     timeToProto(r.StartTime),
		EndTime:       timeToProto(r.EndTime),
		DefectCodes:   r.DefectCodes,
		Remark:        r.Remark,
	}
}

func reportFromProto(r *mes.MobileJobReport) *model.MobileJobReport {
	if r == nil {
		return nil
	}
	m := &model.MobileJobReport{
		OrderID:       r.OrderId,
		OrderNo:       r.OrderNo,
		ProcessID:     r.ProcessId,
		OperationID:   r.OperationId,
		OperationName: r.OperationName,
		EmployeeID:    r.EmployeeId,
		EmployeeName:  r.EmployeeName,
		WorkstationID: r.WorkstationId,
		ReportedQty:   r.ReportedQty,
		QualifiedQty:  r.QualifiedQty,
		DefectiveQty:  r.DefectiveQty,
		WorkMinutes:   int(r.WorkMinutes),
		ReportType:    model.ReportType(r.ReportType),
		Status:        model.ReportStatus(r.Status),
		StartTime:     timeFromProto(r.StartTime),
		EndTime:       timeFromProto(r.EndTime),
		DefectCodes:   r.DefectCodes,
		Remark:        r.Remark,
	}
	if r.Base != nil {
		m.ID = r.Base.Id
		m.TenantID = r.Base.TenantId
	}
	return m
}
