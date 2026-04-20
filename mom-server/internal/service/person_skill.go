package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// PersonSkillService 人员技能服务
type PersonSkillService struct {
	skillRepo  *repository.PersonSkillRepository
	scoreRepo  *repository.PersonSkillScoreRepository
}

func NewPersonSkillService(skillRepo *repository.PersonSkillRepository, scoreRepo *repository.PersonSkillScoreRepository) *PersonSkillService {
	return &PersonSkillService{
		skillRepo: skillRepo,
		scoreRepo: scoreRepo,
	}
}

// List 获取人员技能列表
func (s *PersonSkillService) List(ctx context.Context, tenantID int64, personID int64, workshopID int64, skillLevel string, page, pageSize int) ([]model.PersonSkill, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	return s.skillRepo.List(ctx, tenantID, personID, workshopID, skillLevel, page, pageSize)
}

// GetByID 获取单个人员技能
func (s *PersonSkillService) GetByID(ctx context.Context, id uint) (*model.PersonSkill, error) {
	return s.skillRepo.GetByID(ctx, id)
}

// GetByPersonID 获取人员的所有技能
func (s *PersonSkillService) GetByPersonID(ctx context.Context, personID int64) ([]model.PersonSkill, error) {
	return s.skillRepo.GetByPersonID(ctx, personID)
}

// Create 创建人员技能
func (s *PersonSkillService) Create(ctx context.Context, skill *model.PersonSkill) error {
	return s.skillRepo.Create(ctx, skill)
}

// Update 更新人员技能
func (s *PersonSkillService) Update(ctx context.Context, id uint, skill *model.PersonSkill) error {
	updates := map[string]interface{}{
		"person_code":     skill.PersonCode,
		"person_name":     skill.PersonName,
		"workshop_id":     skill.WorkshopID,
		"workstation_id":  skill.WorkstationID,
		"skill_level":     skill.SkillLevel,
		"certified_date": skill.CertifiedDate,
		"expiry_date":    skill.ExpiryDate,
		"status":         skill.Status,
	}
	return s.skillRepo.Update(ctx, id, updates)
}

// Delete 删除人员技能
func (s *PersonSkillService) Delete(ctx context.Context, id uint) error {
	// 删除关联的评分记录
	if err := s.scoreRepo.DeleteByPersonSkillID(ctx, uint(id)); err != nil {
		return err
	}
	return s.skillRepo.Delete(ctx, id)
}

// GetDetail 获取人员技能详情（包含评分列表）
func (s *PersonSkillService) GetDetail(ctx context.Context, id uint) (*model.PersonSkillDetail, error) {
	skill, err := s.skillRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	scores, err := s.scoreRepo.ListByPersonSkillID(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return &model.PersonSkillDetail{
		PersonSkill: *skill,
		Scores:      scores,
	}, nil
}

// EvaluateSkill 添加技能评分
func (s *PersonSkillService) EvaluateSkill(ctx context.Context, score *model.PersonSkillScore) error {
	return s.scoreRepo.Create(ctx, score)
}

// GetPersonCapability 获取人员能力报告
func (s *PersonSkillService) GetPersonCapability(ctx context.Context, personID int64) (map[string]interface{}, error) {
	skills, err := s.skillRepo.GetByPersonID(ctx, personID)
	if err != nil {
		return nil, err
	}

	var capabilityList []map[string]interface{}
	for _, skill := range skills {
		scores, err := s.scoreRepo.ListByPersonSkillID(ctx, skill.ID)
		if err != nil {
			continue
		}

		var totalScore float64
		for _, s := range scores {
			totalScore += s.Score
		}
		avgScore := float64(0)
		if len(scores) > 0 {
			avgScore = totalScore / float64(len(scores))
		}

		capabilityList = append(capabilityList, map[string]interface{}{
			"skill_id":      skill.ID,
			"skill_level":   skill.SkillLevel,
			"status":        skill.Status,
			"certified_at":  skill.CertifiedDate,
			"expiry_at":     skill.ExpiryDate,
			"avg_score":     avgScore,
			"eval_count":    len(scores),
			"recent_scores": scores,
		})
	}

	return map[string]interface{}{
		"person_id":   personID,
		"skill_count": len(skills),
		"capabilities": capabilityList,
	}, nil
}