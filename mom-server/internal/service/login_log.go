package service

import (
	"context"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type LoginLogService struct {
	loginLogRepo *repository.LoginLogRepository
}

func NewLoginLogService(loginLogRepo *repository.LoginLogRepository) *LoginLogService {
	return &LoginLogService{loginLogRepo: loginLogRepo}
}

func (s *LoginLogService) RecordLogin(ctx context.Context, tenantID int64, username, ip, location, browser, os string, status int, msg string) error {
	loginLog := &model.LoginLog{
		TenantID: tenantID,
		Username: username,
		IP:       ip,
		Browser:  browser,
		OS:       os,
		Status:   status,
		LoginTime: time.Now(),
	}
	if location != "" {
		loginLog.LoginLocation = &location
	}
	if msg != "" {
		loginLog.Msg = &msg
	}
	return s.loginLogRepo.Create(ctx, loginLog)
}

func (s *LoginLogService) GetList(ctx context.Context, tenantID int64, username, status, ip string, page, pageSize int) ([]model.LoginLog, int64, error) {
	return s.loginLogRepo.FindByPage(ctx, tenantID, username, status, ip, page, pageSize)
}

func (s *LoginLogService) DeleteClean(ctx context.Context, days int) error {
	return s.loginLogRepo.DeleteClean(ctx, days)
}
