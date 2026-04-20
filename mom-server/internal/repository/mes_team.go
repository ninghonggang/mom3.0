package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type MesTeamRepository struct {
	db *gorm.DB
}

func NewMesTeamRepository(db *gorm.DB) *MesTeamRepository {
	return &MesTeamRepository{db: db}
}

func (r *MesTeamRepository) List(ctx context.Context, tenantID int64) ([]model.MesTeam, int64, error) {
	var list []model.MesTeam
	var total int64

	query := r.db.WithContext(ctx).Model(&model.MesTeam{})
	if tenantID > 0 {
		query = query.Where("tenant_id = ?", tenantID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *MesTeamRepository) GetByID(ctx context.Context, id uint) (*model.MesTeam, error) {
	var team model.MesTeam
	err := r.db.WithContext(ctx).First(&team, id).Error
	return &team, err
}

func (r *MesTeamRepository) Create(ctx context.Context, team *model.MesTeam) error {
	return r.db.WithContext(ctx).Create(team).Error
}

func (r *MesTeamRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesTeam{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MesTeamRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MesTeam{}, id).Error
}

// MesTeamMember Repository
type MesTeamMemberRepository struct {
	db *gorm.DB
}

func NewMesTeamMemberRepository(db *gorm.DB) *MesTeamMemberRepository {
	return &MesTeamMemberRepository{db: db}
}

func (r *MesTeamMemberRepository) ListByTeamID(ctx context.Context, teamID int64) ([]model.MesTeamMember, error) {
	var list []model.MesTeamMember
	err := r.db.WithContext(ctx).Where("team_id = ?", teamID).Order("id ASC").Find(&list).Error
	return list, err
}

func (r *MesTeamMemberRepository) Create(ctx context.Context, member *model.MesTeamMember) error {
	return r.db.WithContext(ctx).Create(member).Error
}

func (r *MesTeamMemberRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesTeamMember{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MesTeamMemberRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MesTeamMember{}, id).Error
}

func (r *MesTeamMemberRepository) DeleteByTeamID(ctx context.Context, teamID uint) error {
	return r.db.WithContext(ctx).Where("team_id = ?", teamID).Delete(&model.MesTeamMember{}).Error
}

// MesTeamShift Repository
type MesTeamShiftRepository struct {
	db *gorm.DB
}

func NewMesTeamShiftRepository(db *gorm.DB) *MesTeamShiftRepository {
	return &MesTeamShiftRepository{db: db}
}

func (r *MesTeamShiftRepository) ListByTeamID(ctx context.Context, teamID int64) ([]model.MesTeamShift, error) {
	var list []model.MesTeamShift
	err := r.db.WithContext(ctx).Where("team_id = ?", teamID).Order("shift_date DESC").Find(&list).Error
	return list, err
}

func (r *MesTeamShiftRepository) ListByDate(ctx context.Context, shiftDate string) ([]model.MesTeamShift, error) {
	var list []model.MesTeamShift
	err := r.db.WithContext(ctx).Where("shift_date = ?", shiftDate).Find(&list).Error
	return list, err
}

func (r *MesTeamShiftRepository) Create(ctx context.Context, shift *model.MesTeamShift) error {
	return r.db.WithContext(ctx).Create(shift).Error
}

func (r *MesTeamShiftRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.MesTeamShift{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MesTeamShiftRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.MesTeamShift{}, id).Error
}
