package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type ImportTaskRepository struct {
	db *gorm.DB
}

func NewImportTaskRepository(db *gorm.DB) *ImportTaskRepository {
	return &ImportTaskRepository{db: db}
}

func (r *ImportTaskRepository) Create(ctx context.Context, task *model.ImportTask) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *ImportTaskRepository) GetByID(ctx context.Context, id uint) (*model.ImportTask, error) {
	var task model.ImportTask
	err := r.db.WithContext(ctx).First(&task, id).Error
	return &task, err
}

func (r *ImportTaskRepository) GetByTaskNo(ctx context.Context, taskNo string) (*model.ImportTask, error) {
	var task model.ImportTask
	err := r.db.WithContext(ctx).Where("task_no = ?", taskNo).First(&task).Error
	return &task, err
}

func (r *ImportTaskRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.ImportTask{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ImportTaskRepository) UpdateStatus(ctx context.Context, id uint, status string, totalRows, successRows, failRows int) error {
	updates := map[string]interface{}{
		"status":       status,
		"total_rows":   totalRows,
		"success_rows": successRows,
		"fail_rows":    failRows,
	}
	if status == model.ImportStatusSuccess || status == model.ImportStatusFail {
		updates["completed_at"] = gorm.Expr("NOW()")
	}
	return r.db.WithContext(ctx).Model(&model.ImportTask{}).Where("id = ?", id).Updates(updates).Error
}
