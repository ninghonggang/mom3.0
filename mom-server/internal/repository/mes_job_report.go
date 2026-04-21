package repository

import (
	"context"

	"gorm.io/gorm"
	"mom-server/internal/model"
)

// MesJobReportLogRepository 报工记录仓储
type MesJobReportLogRepository struct {
	db *gorm.DB
}

func NewMesJobReportLogRepository(db *gorm.DB) *MesJobReportLogRepository {
	return &MesJobReportLogRepository{db: db}
}

func (r *MesJobReportLogRepository) Create(ctx context.Context, log *model.MesJobReportLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *MesJobReportLogRepository) GetByID(ctx context.Context, id uint64) (*model.MesJobReportLog, error) {
	var log model.MesJobReportLog
	err := r.db.WithContext(ctx).First(&log, id).Error
	return &log, err
}

func (r *MesJobReportLogRepository) Page(ctx context.Context, tenantID int64, query *model.MesJobReportLogQueryVO) ([]model.MesJobReportLog, int64, error) {
	var list []model.MesJobReportLog
	var total int64

	q := r.db.WithContext(ctx).Model(&model.MesJobReportLog{}).Where("tenant_id = ?", tenantID)

	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"
		q = q.Where("work_order_code LIKE ? OR process_name LIKE ?", keyword, keyword)
	}
	if query.WorkOrderId > 0 {
		q = q.Where("work_order_id = ?", query.WorkOrderId)
	}
	if query.ProcessCode != "" {
		q = q.Where("process_code = ?", query.ProcessCode)
	}
	if query.ReporterId > 0 {
		q = q.Where("reporter_id = ?", query.ReporterId)
	}
	if query.StartDate != "" {
		q = q.Where("report_time >= ?", query.StartDate)
	}
	if query.EndDate != "" {
		q = q.Where("report_time <= ?", query.EndDate)
	}

	q.Count(&total)

	page := query.Page
	if page <= 0 {
		page = 1
	}
	pageSize := query.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	err := q.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *MesJobReportLogRepository) Senior(ctx context.Context, tenantID int64, conditions []map[string]interface{}) ([]model.MesJobReportLog, int64, error) {
	var list []model.MesJobReportLog
	var total int64

	q := r.db.WithContext(ctx).Model(&model.MesJobReportLog{}).Where("tenant_id = ?", tenantID)

	for _, cond := range conditions {
		field, ok := cond["field"].(string)
		if !ok {
			continue
		}
		operator, _ := cond["operator"].(string)
		value := cond["value"]

		switch operator {
		case "eq":
			q = q.Where(field+" = ?", value)
		case "ne":
			q = q.Where(field+" != ?", value)
		case "like":
			q = q.Where(field+" LIKE ?", "%"+value.(string)+"%")
		case "gt":
			q = q.Where(field+" > ?", value)
		case "ge":
			q = q.Where(field+" >= ?", value)
		case "lt":
			q = q.Where(field+" < ?", value)
		case "le":
			q = q.Where(field+" <= ?", value)
		case "between":
			if arr, ok := value.([]interface{}); ok && len(arr) == 2 {
				q = q.Where(field+" BETWEEN ? AND ?", arr[0], arr[1])
			}
		case "in":
			if arr, ok := value.([]interface{}); ok {
				q = q.Where(field+" IN ?", arr)
			}
		}
	}

	q.Count(&total)

	page := 1
	pageSize := 20

	err := q.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&list).Error
	return list, total, err
}
