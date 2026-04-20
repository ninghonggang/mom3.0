package service

import (
	"context"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type MesTeamService struct {
	teamRepo  *repository.MesTeamRepository
	memberRepo *repository.MesTeamMemberRepository
	shiftRepo  *repository.MesTeamShiftRepository
}

func NewMesTeamService(teamRepo *repository.MesTeamRepository, memberRepo *repository.MesTeamMemberRepository, shiftRepo *repository.MesTeamShiftRepository) *MesTeamService {
	return &MesTeamService{
		teamRepo:  teamRepo,
		memberRepo: memberRepo,
		shiftRepo:  shiftRepo,
	}
}

func (s *MesTeamService) List(ctx context.Context) ([]model.MesTeam, int64, error) {
	return s.teamRepo.List(ctx, 0)
}

func (s *MesTeamService) GetByID(ctx context.Context, id uint) (*model.MesTeam, error) {
	return s.teamRepo.GetByID(ctx, id)
}

func (s *MesTeamService) Create(ctx context.Context, team *model.MesTeam) error {
	return s.teamRepo.Create(ctx, team)
}

func (s *MesTeamService) Update(ctx context.Context, id uint, team *model.MesTeam) error {
	updates := map[string]interface{}{
		"team_name":  team.TeamName,
		"workshop_id": team.WorkshopID,
		"shift_id":   team.ShiftID,
		"leader_id":  team.LeaderID,
		"phone":      team.Phone,
		"status":     team.Status,
	}
	if team.Remark != nil {
		updates["remark"] = team.Remark
	}
	return s.teamRepo.Update(ctx, id, updates)
}

func (s *MesTeamService) Delete(ctx context.Context, id uint) error {
	// 删除成员
	if err := s.memberRepo.DeleteByTeamID(ctx, id); err != nil {
		return err
	}
	return s.teamRepo.Delete(ctx, id)
}

// 成员管理
func (s *MesTeamService) ListMembers(ctx context.Context, teamID int64) ([]model.MesTeamMember, error) {
	return s.memberRepo.ListByTeamID(ctx, teamID)
}

func (s *MesTeamService) AddMember(ctx context.Context, member *model.MesTeamMember) error {
	return s.memberRepo.Create(ctx, member)
}

func (s *MesTeamService) UpdateMember(ctx context.Context, id uint, member *model.MesTeamMember) error {
	updates := map[string]interface{}{
		"role":   member.Role,
		"status":  member.Status,
		"remark":  member.Remark,
	}
	return s.memberRepo.Update(ctx, id, updates)
}

func (s *MesTeamService) RemoveMember(ctx context.Context, id uint) error {
	return s.memberRepo.Delete(ctx, id)
}

// 排班管理
func (s *MesTeamService) ListShifts(ctx context.Context, teamID int64) ([]model.MesTeamShift, error) {
	return s.shiftRepo.ListByTeamID(ctx, teamID)
}

func (s *MesTeamService) CreateShift(ctx context.Context, shift *model.MesTeamShift) error {
	return s.shiftRepo.Create(ctx, shift)
}

func (s *MesTeamService) UpdateShift(ctx context.Context, id uint, shift *model.MesTeamShift) error {
	updates := map[string]interface{}{
		"shift_id":   shift.ShiftID,
		"shift_date": shift.ShiftDate,
		"start_time": shift.StartTime,
		"end_time":   shift.EndTime,
		"leader_id":  shift.LeaderID,
	}
	if shift.Remark != nil {
		updates["remark"] = shift.Remark
	}
	return s.shiftRepo.Update(ctx, id, updates)
}

func (s *MesTeamService) DeleteShift(ctx context.Context, id uint) error {
	return s.shiftRepo.Delete(ctx, id)
}
