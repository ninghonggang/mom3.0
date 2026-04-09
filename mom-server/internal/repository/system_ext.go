package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type PrintTemplateRepository struct {
	db *gorm.DB
}

func NewPrintTemplateRepository(db *gorm.DB) *PrintTemplateRepository {
	return &PrintTemplateRepository{db: db}
}

func (r *PrintTemplateRepository) List(ctx context.Context, tenantID int64) ([]model.PrintTemplate, error) {
	var list []model.PrintTemplate
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *PrintTemplateRepository) GetByID(ctx context.Context, id uint) (*model.PrintTemplate, error) {
	var t model.PrintTemplate
	err := r.db.WithContext(ctx).First(&t, id).Error
	return &t, err
}

func (r *PrintTemplateRepository) Create(ctx context.Context, t *model.PrintTemplate) error {
	return r.db.WithContext(ctx).Create(t).Error
}

func (r *PrintTemplateRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.PrintTemplate{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PrintTemplateRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.PrintTemplate{}, id).Error
}

type NoticeRepository struct {
	db *gorm.DB
}

func NewNoticeRepository(db *gorm.DB) *NoticeRepository {
	return &NoticeRepository{db: db}
}

func (r *NoticeRepository) List(ctx context.Context, tenantID int64, query string) ([]model.Notice, int64, error) {
	var list []model.Notice
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Notice{}).Where("tenant_id = ?", tenantID)

	if query != "" {
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	db.Count(&total)
	err := db.Order("is_top DESC, publish_time DESC").Find(&list).Error
	return list, total, err
}

func (r *NoticeRepository) GetByID(ctx context.Context, id uint) (*model.Notice, error) {
	var n model.Notice
	err := r.db.WithContext(ctx).First(&n, id).Error
	return &n, err
}

func (r *NoticeRepository) Create(ctx context.Context, n *model.Notice) error {
	return r.db.WithContext(ctx).Create(n).Error
}

func (r *NoticeRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.Notice{}).Where("id = ?", id).Updates(updates).Error
}

func (r *NoticeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Notice{}, id).Error
}

func (r *NoticeRepository) IncrementViewCount(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.Notice{}).Where("id = ?", id).Update("view_count", gorm.Expr("view_count + 1")).Error
}

func (r *NoticeRepository) MarkAsRead(ctx context.Context, noticeID, userID int64) error {
	record := model.NoticeReadRecord{
		TenantID: 1,
		NoticeID: noticeID,
		UserID:   userID,
	}
	return r.db.WithContext(ctx).Create(&record).Error
}

func (r *NoticeRepository) GetMyNotices(ctx context.Context, tenantID, userID int64) ([]model.Notice, error) {
	var list []model.Notice
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND status = 2 AND (target_type = 'ALL' OR target_ids LIKE ?)",
		tenantID, "%,"+fmt.Sprintf("%d", userID)+",%").
		Order("is_top DESC, publish_time DESC").Find(&list).Error
	return list, err
}
