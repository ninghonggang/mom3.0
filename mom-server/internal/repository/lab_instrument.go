package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// LabInstrumentRepository 实验室仪器仓储
type LabInstrumentRepository struct {
	db *gorm.DB
}

func NewLabInstrumentRepository(db *gorm.DB) *LabInstrumentRepository {
	return &LabInstrumentRepository{db: db}
}

func (r *LabInstrumentRepository) List(ctx context.Context, query *model.LabInstrumentQuery) ([]model.LabInstrument, int64, error) {
	var list []model.LabInstrument
	var total int64

	db := r.db.WithContext(ctx).Model(&model.LabInstrument{})
	if query.InstrumentCode != "" {
		db = db.Where("instrument_code LIKE ?", "%"+query.InstrumentCode+"%")
	}
	if query.InstrumentName != "" {
		db = db.Where("instrument_name LIKE ?", "%"+query.InstrumentName+"%")
	}
	if query.InstrumentType != "" {
		db = db.Where("instrument_type = ?", query.InstrumentType)
	}
	if query.CalibrationStatus != "" {
		db = db.Where("calibration_status = ?", query.CalibrationStatus)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	err = db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	return list, total, err
}

func (r *LabInstrumentRepository) GetByID(ctx context.Context, id uint64) (*model.LabInstrument, error) {
	var instrument model.LabInstrument
	err := r.db.WithContext(ctx).First(&instrument, id).Error
	if err != nil {
		return nil, err
	}
	return &instrument, nil
}

func (r *LabInstrumentRepository) GetByCode(ctx context.Context, code string) (*model.LabInstrument, error) {
	var instrument model.LabInstrument
	err := r.db.WithContext(ctx).Where("instrument_code = ?", code).First(&instrument).Error
	if err != nil {
		return nil, err
	}
	return &instrument, nil
}

func (r *LabInstrumentRepository) Create(ctx context.Context, instrument *model.LabInstrument) error {
	return r.db.WithContext(ctx).Create(instrument).Error
}

func (r *LabInstrumentRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LabInstrument{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LabInstrumentRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.LabInstrument{}, id).Error
}

// LabCalibrationRepository 校准记录仓储
type LabCalibrationRepository struct {
	db *gorm.DB
}

func NewLabCalibrationRepository(db *gorm.DB) *LabCalibrationRepository {
	return &LabCalibrationRepository{db: db}
}

func (r *LabCalibrationRepository) ListByInstrumentID(ctx context.Context, instrumentID uint64, query *model.LabCalibrationQuery) ([]model.LabCalibration, int64, error) {
	var list []model.LabCalibration
	var total int64

	db := r.db.WithContext(ctx).Model(&model.LabCalibration{}).Where("instrument_id = ?", instrumentID)

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	err = db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	return list, total, err
}

func (r *LabCalibrationRepository) GetByID(ctx context.Context, id uint64) (*model.LabCalibration, error) {
	var calibration model.LabCalibration
	err := r.db.WithContext(ctx).First(&calibration, id).Error
	if err != nil {
		return nil, err
	}
	return &calibration, nil
}

func (r *LabCalibrationRepository) Create(ctx context.Context, calibration *model.LabCalibration) error {
	return r.db.WithContext(ctx).Create(calibration).Error
}

func (r *LabCalibrationRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.LabCalibration{}).Where("id = ?", id).Updates(updates).Error
}

func (r *LabCalibrationRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.LabCalibration{}, id).Error
}

func (r *LabCalibrationRepository) DeleteByInstrumentID(ctx context.Context, instrumentID uint64) error {
	return r.db.WithContext(ctx).Where("instrument_id = ?", instrumentID).Delete(&model.LabCalibration{}).Error
}