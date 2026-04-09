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

// GetForwardTrace 正向追溯 - 从当前工序往前追溯所有工序
func (r *TraceRepository) GetForwardTrace(ctx context.Context, serialNumber string) ([]model.TraceRecord, error) {
	var records []model.TraceRecord
	// 按时间顺序获取所有追溯记录
	err := r.db.WithContext(ctx).Where("serial_number = ?", serialNumber).Order("operate_time ASC").Find(&records).Error
	return records, err
}

// GetBackwardTrace 反向追溯 - 从当前工序往后追溯所有工序
func (r *TraceRepository) GetBackwardTrace(ctx context.Context, serialNumber string) ([]model.TraceRecord, error) {
	var records []model.TraceRecord
	// 按时间倒序获取所有追溯记录
	err := r.db.WithContext(ctx).Where("serial_number = ?", serialNumber).Order("operate_time DESC").Find(&records).Error
	return records, err
}
