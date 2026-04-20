package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type InspectionRepository struct {
	db *gorm.DB
}

func NewInspectionRepository(db *gorm.DB) *InspectionRepository {
	return &InspectionRepository{db: db}
}

// ListPlans 查询检验计划列表
func (r *InspectionRepository) ListPlans(ctx context.Context, tenantID uint64, query map[string]interface{}) ([]model.QualityInspectionPlan, int64, error) {
	var list []model.QualityInspectionPlan
	var total int64

	db := r.db.WithContext(ctx).Model(&model.QualityInspectionPlan{}).Where("tenant_id = ?", tenantID)

	if planCode, ok := query["plan_code"].(string); ok && planCode != "" {
		db = db.Where("plan_code LIKE ?", "%"+planCode+"%")
	}
	if planName, ok := query["plan_name"].(string); ok && planName != "" {
		db = db.Where("plan_name LIKE ?", "%"+planName+"%")
	}
	if inspectionType, ok := query["inspection_type"].(string); ok && inspectionType != "" {
		db = db.Where("inspection_type = ?", inspectionType)
	}
	if aqlLevel, ok := query["aql_level"].(string); ok && aqlLevel != "" {
		db = db.Where("aql_level = ?", aqlLevel)
	}
	if status, ok := query["status"].(string); ok && status != "" {
		db = db.Where("status = ?", status)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Order("id DESC").Find(&list).Error
	return list, total, err
}

// GetPlanByID 根据ID获取检验计划
func (r *InspectionRepository) GetPlanByID(ctx context.Context, id uint64) (*model.QualityInspectionPlan, error) {
	var plan model.QualityInspectionPlan
	err := r.db.WithContext(ctx).First(&plan, id).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

// CreatePlan 创建检验计划
func (r *InspectionRepository) CreatePlan(ctx context.Context, plan *model.QualityInspectionPlan) error {
	return r.db.WithContext(ctx).Create(plan).Error
}

// UpdatePlan 更新检验计划
func (r *InspectionRepository) UpdatePlan(ctx context.Context, plan *model.QualityInspectionPlan) error {
	return r.db.WithContext(ctx).Save(plan).Error
}

// DeletePlan 删除检验计划
func (r *InspectionRepository) DeletePlan(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.QualityInspectionPlan{}, id).Error
}

// GetAQLSampleSizes 根据AQL值、批量大小和检验水平获取抽样方案
func (r *InspectionRepository) GetAQLSampleSizes(ctx context.Context, tenantID uint64, aql string, batchSize int, level string) (*model.AQLSampleSize, error) {
	var result model.AQLSampleSize
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND aql_value = ? AND inspection_level = ? AND batch_size_min <= ? AND batch_size_max >= ?",
		tenantID, aql, level, batchSize, batchSize).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SeedData 批量插入AQL标准抽样数据
func (r *InspectionRepository) SeedData(ctx context.Context, data []model.AQLSampleSize) error {
	return r.db.WithContext(ctx).CreateInBatches(data, 100).Error
}

// GetAQLSampleSizeByCode 根据样本量字码和AQL值获取抽样方案
func (r *InspectionRepository) GetAQLSampleSizeByCode(ctx context.Context, tenantID uint64, code string, aql float64, level string) (*model.AQLSampleSize, error) {
	var result model.AQLSampleSize
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND sample_size_code = ? AND aql_value = ? AND inspection_level = ?",
		tenantID, code, aql, level).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CountAQLData 统计AQL数据条数
func (r *InspectionRepository) CountAQLData(ctx context.Context, tenantID uint64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.AQLSampleSize{}).Where("tenant_id = ?", tenantID).Count(&count).Error
	return count, err
}

// GetAQLSampleSizeByBatchSize 根据批量范围获取样本量字码
func (r *InspectionRepository) GetAQLSampleSizeByBatchSize(ctx context.Context, tenantID uint64, batchSize int, level string) (*model.AQLSampleSize, error) {
	var result model.AQLSampleSize
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND inspection_level = ? AND batch_size_min <= ? AND batch_size_max >= ?",
		tenantID, level, batchSize, batchSize).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GenerateSampleSizeCode 根据批量大小生成样本量字码
func GenerateSampleSizeCode(batchSize int) string {
	switch {
	case batchSize >= 2 && batchSize <= 8:
		return "D"
	case batchSize >= 9 && batchSize <= 15:
		return "E"
	case batchSize >= 16 && batchSize <= 25:
		return "F"
	case batchSize >= 26 && batchSize <= 50:
		return "G"
	case batchSize >= 51 && batchSize <= 90:
		return "H"
	case batchSize >= 91 && batchSize <= 150:
		return "J"
	case batchSize >= 151 && batchSize <= 280:
		return "K"
	case batchSize >= 281 && batchSize <= 500:
		return "L"
	case batchSize >= 501 && batchSize <= 1200:
		return "M"
	case batchSize >= 1201 && batchSize <= 3200:
		return "N"
	case batchSize >= 3201 && batchSize <= 10000:
		return "P"
	case batchSize >= 10001 && batchSize <= 35000:
		return "Q"
	case batchSize >= 35001 && batchSize <= 150000:
		return "R"
	case batchSize >= 150001 && batchSize <= 500000:
		return "S"
	case batchSize > 500000:
		return "T"
	default:
		return fmt.Sprintf("N/A")
	}
}
