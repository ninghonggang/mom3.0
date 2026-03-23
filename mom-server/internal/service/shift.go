package service

import (
	"context"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type MdmShiftService struct {
	shiftRepo *repository.MdmShiftRepository
}

func NewMdmShiftService(shiftRepo *repository.MdmShiftRepository) *MdmShiftService {
	return &MdmShiftService{shiftRepo: shiftRepo}
}

func (s *MdmShiftService) List(ctx context.Context) ([]model.MdmShift, int64, error) {
	return s.shiftRepo.List(ctx, 0)
}

func (s *MdmShiftService) GetByID(ctx context.Context, id uint) (*model.MdmShift, error) {
	return s.shiftRepo.GetByID(ctx, id)
}

func (s *MdmShiftService) Create(ctx context.Context, shift *model.MdmShift) error {
	return s.shiftRepo.Create(ctx, shift)
}

func (s *MdmShiftService) Update(ctx context.Context, id uint, shift *model.MdmShift) error {
	updates := map[string]interface{}{
		"shift_name": shift.ShiftName,
		"start_time": shift.StartTime,
		"end_time":   shift.EndTime,
		"work_hours": shift.WorkHours,
		"is_night":   shift.IsNight,
	}
	if shift.Remark != nil {
		updates["remark"] = shift.Remark
	}
	return s.shiftRepo.Update(ctx, id, updates)
}

func (s *MdmShiftService) Delete(ctx context.Context, id uint) error {
	return s.shiftRepo.Delete(ctx, id)
}
