package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type TraceRepository struct {
	db *gorm.DB
}

func NewTraceRepository(db *gorm.DB) *TraceRepository {
	return &TraceRepository{db: db}
}

func (r *TraceRepository) GetBySerialNumber(ctx context.Context, serialNumber string) (*model.SerialNumber, error) {
	var sn model.SerialNumber
	err := r.db.WithContext(ctx).Where("serial_number = ?", serialNumber).First(&sn).Error
	return &sn, err
}

func (r *TraceRepository) GetByBatchNo(ctx context.Context, batchNo string) ([]model.SerialNumber, error) {
	var list []model.SerialNumber
	err := r.db.WithContext(ctx).Where("batch_no = ?", batchNo).Find(&list).Error
	return list, err
}

func (r *TraceRepository) GetTraceRecordsBySerial(ctx context.Context, serialNumber string) ([]model.TraceRecord, error) {
	var records []model.TraceRecord
	err := r.db.WithContext(ctx).Where("serial_number = ?", serialNumber).Order("operate_time ASC").Find(&records).Error
	return records, err
}

func (r *TraceRepository) GetByOrderID(ctx context.Context, orderID int64) ([]model.SerialNumber, error) {
	var list []model.SerialNumber
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&list).Error
	return list, err
}

type AndonRepository struct {
	db *gorm.DB
}

func NewAndonRepository(db *gorm.DB) *AndonRepository {
	return &AndonRepository{db: db}
}

func (r *AndonRepository) List(ctx context.Context, tenantID int64, status int, callNo string) ([]model.AndonCall, int64, error) {
	var list []model.AndonCall
	var total int64
	query := r.db.WithContext(ctx).Model(&model.AndonCall{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if callNo != "" {
		query = query.Where("call_no LIKE ?", "%"+callNo+"%")
	}
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *AndonRepository) GetByID(ctx context.Context, id uint) (*model.AndonCall, error) {
	var call model.AndonCall
	err := r.db.WithContext(ctx).First(&call, id).Error
	return &call, err
}

func (r *AndonRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.AndonCall{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AndonRepository) Create(ctx context.Context, call *model.AndonCall) error {
	return r.db.WithContext(ctx).Create(call).Error
}
