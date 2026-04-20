package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// PersonSkillRepository 人员技能仓储
type PersonSkillRepository struct {
	db *gorm.DB
}

func NewPersonSkillRepository(db *gorm.DB) *PersonSkillRepository {
	return &PersonSkillRepository{db: db}
}

// List 获取人员技能列表（分页）
func (r *PersonSkillRepository) List(ctx context.Context, tenantID int64, personID int64, workshopID int64, skillLevel string, page, pageSize int) ([]model.PersonSkill, int64, error) {
	var list []model.PersonSkill
	var total int64

	query := r.db.WithContext(ctx).Model(&model.PersonSkill{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}
	if personID > 0 {
		query = query.Where("person_id = ?", personID)
	}
	if workshopID > 0 {
		query = query.Where("workshop_id = ?", workshopID)
	}
	if skillLevel != "" {
		query = query.Where("skill_level = ?", skillLevel)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// GetByID 获取单个人员技能
func (r *PersonSkillRepository) GetByID(ctx context.Context, id uint) (*model.PersonSkill, error) {
	var skill model.PersonSkill
	err := r.db.WithContext(ctx).First(&skill, id).Error
	return &skill, err
}

// GetByPersonID 获取人员的所有技能
func (r *PersonSkillRepository) GetByPersonID(ctx context.Context, personID int64) ([]model.PersonSkill, error) {
	var list []model.PersonSkill
	err := r.db.WithContext(ctx).Where("person_id = ?", personID).Order("id DESC").Find(&list).Error
	return list, err
}

// Create 创建人员技能
func (r *PersonSkillRepository) Create(ctx context.Context, skill *model.PersonSkill) error {
	return r.db.WithContext(ctx).Create(skill).Error
}

// Update 更新人员技能
func (r *PersonSkillRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.PersonSkill{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除人员技能
func (r *PersonSkillRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.PersonSkill{}, id).Error
}

// GetByPersonAndType 获取人员指定类型的技能
func (r *PersonSkillRepository) GetByPersonAndType(ctx context.Context, personID int64, skillType string) (*model.PersonSkill, error) {
	var skill model.PersonSkill
	err := r.db.WithContext(ctx).Where("person_id = ? AND skill_level = ?", personID, skillType).First(&skill).Error
	return &skill, err
}

// PersonSkillScoreRepository 人员技能评分仓储
type PersonSkillScoreRepository struct {
	db *gorm.DB
}

func NewPersonSkillScoreRepository(db *gorm.DB) *PersonSkillScoreRepository {
	return &PersonSkillScoreRepository{db: db}
}

// ListByPersonSkillID 获取技能的所有评分
func (r *PersonSkillScoreRepository) ListByPersonSkillID(ctx context.Context, personSkillID int64) ([]model.PersonSkillScore, error) {
	var list []model.PersonSkillScore
	err := r.db.WithContext(ctx).Where("person_skill_id = ?", personSkillID).Order("evaluated_at DESC").Find(&list).Error
	return list, err
}

// Create 创建评分
func (r *PersonSkillScoreRepository) Create(ctx context.Context, score *model.PersonSkillScore) error {
	return r.db.WithContext(ctx).Create(score).Error
}

// GetAverageScore 获取平均分
func (r *PersonSkillScoreRepository) GetAverageScore(ctx context.Context, personSkillID int64) (float64, error) {
	var result struct {
		Avg float64
	}
	err := r.db.WithContext(ctx).Model(&model.PersonSkillScore{}).
		Where("person_skill_id = ?", personSkillID).
		Select("AVG(score) as avg").Scan(&result).Error
	return result.Avg, err
}

// DeleteByPersonSkillID 删除技能的所有评分
func (r *PersonSkillScoreRepository) DeleteByPersonSkillID(ctx context.Context, personSkillID uint) error {
	return r.db.WithContext(ctx).Where("person_skill_id = ?", personSkillID).Delete(&model.PersonSkillScore{}).Error
}