package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// ========== MesProcessRepository 工艺路线仓储 ==========

type MesProcessRepository struct {
	db *gorm.DB
}

func NewMesProcessRepository(db *gorm.DB) *MesProcessRepository {
	return &MesProcessRepository{db: db}
}

func (r *MesProcessRepository) List(ctx context.Context, tenantID int64, query map[string]interface{}) ([]model.MesProcess, int64, error) {
	var list []model.MesProcess
	var total int64

	q := r.db.WithContext(ctx).Model(&model.MesProcess{}).Where("tenant_id = ?", tenantID)

	if status, ok := query["status"]; ok && status != "" {
		q = q.Where("status = ?", status)
	}
	if processCode, ok := query["process_code"]; ok && processCode.(string) != "" {
		q = q.Where("process_code LIKE ?", "%"+processCode.(string)+"%")
	}
	if processName, ok := query["process_name"].(string); ok && processName != "" {
		q = q.Where("process_name LIKE ?", "%"+processName+"%")
	}
	if materialID, ok := query["material_id"]; ok && materialID.(int64) > 0 {
		q = q.Where("material_id = ?", materialID)
	}
	if workshopID, ok := query["workshop_id"]; ok && workshopID.(int64) > 0 {
		q = q.Where("workshop_id = ?", workshopID)
	}

	q.Count(&total)

	page := 1
	limit := 20
	if p, ok := query["page"].(int); ok && p > 0 {
		page = p
	}
	if l, ok := query["limit"].(int); ok && l > 0 {
		limit = l
	}
	q = q.Offset((page - 1) * limit).Limit(limit).Order("id DESC")

	err := q.Find(&list).Error
	return list, total, err
}

func (r *MesProcessRepository) GetByID(ctx context.Context, id uint) (*model.MesProcess, error) {
	var process model.MesProcess
	err := r.db.WithContext(ctx).First(&process, id).Error
	if err != nil {
		return nil, err
	}
	return &process, nil
}

func (r *MesProcessRepository) GetByCode(ctx context.Context, tenantID int64, code string) (*model.MesProcess, error) {
	var process model.MesProcess
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND process_code = ?", tenantID, code).First(&process).Error
	return &process, err
}

func (r *MesProcessRepository) GetByMaterialID(ctx context.Context, tenantID int64, materialID int64) ([]model.MesProcess, error) {
	var list []model.MesProcess
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND material_id = ? AND status = 'ACTIVE' AND is_current = 1", tenantID, materialID).Find(&list).Error
	return list, err
}

func (r *MesProcessRepository) Create(ctx context.Context, process *model.MesProcess) error {
	return r.db.WithContext(ctx).Create(process).Error
}

func (r *MesProcessRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesProcess{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MesProcessRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MesProcess{}, id).Error
}

// GenerateProcessCode 生成工艺路线编号
func (r *MesProcessRepository) GenerateProcessCode(ctx context.Context, tenantID int64) (string, error) {
	var count int64
	r.db.WithContext(ctx).Model(&model.MesProcess{}).Where("tenant_id = ?", tenantID).Count(&count)
	return fmt.Sprintf("PR-%s-%04d", "MES", count+1), nil
}

// ========== MesProcessOperationRepository 工艺路线工序仓储 ==========

type MesProcessOperationRepository struct {
	db *gorm.DB
}

func NewMesProcessOperationRepository(db *gorm.DB) *MesProcessOperationRepository {
	return &MesProcessOperationRepository{db: db}
}

func (r *MesProcessOperationRepository) ListByProcessID(ctx context.Context, processID int64) ([]model.MesProcessOperation, error) {
	var operations []model.MesProcessOperation
	err := r.db.WithContext(ctx).Where("process_id = ?", processID).Order("line_no").Find(&operations).Error
	return operations, err
}

func (r *MesProcessOperationRepository) Create(ctx context.Context, op *model.MesProcessOperation) error {
	return r.db.WithContext(ctx).Create(op).Error
}

func (r *MesProcessOperationRepository) CreateBatch(ctx context.Context, operations []model.MesProcessOperation) error {
	return r.db.WithContext(ctx).Create(&operations).Error
}

func (r *MesProcessOperationRepository) DeleteByProcessID(ctx context.Context, processID int64) error {
	return r.db.WithContext(ctx).Where("process_id = ?", processID).Delete(&model.MesProcessOperation{}).Error
}

func (r *MesProcessOperationRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesProcessOperation{}).Where("id = ?", id).Updates(updates).Error
}
