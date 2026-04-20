package repository

import (
	"context"
	"fmt"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

// BpmTaskTransferRepository 任务转移记录仓库
type BpmTaskTransferRepository struct {
	db *gorm.DB
}

func NewBpmTaskTransferRepository(db *gorm.DB) *BpmTaskTransferRepository {
	return &BpmTaskTransferRepository{db: db}
}

// Create 创建转移记录
func (r *BpmTaskTransferRepository) Create(ctx context.Context, t *model.BpmTaskTransfer) error {
	return r.db.WithContext(ctx).Create(t).Error
}

// ListByTaskID 按任务ID获取转移历史
func (r *BpmTaskTransferRepository) ListByTaskID(ctx context.Context, taskID string) ([]model.BpmTaskTransfer, error) {
	var list []model.BpmTaskTransfer
	err := r.db.WithContext(ctx).
		Where("task_id = ?", taskID).
		Order("created_at ASC").
		Find(&list).Error
	return list, err
}

// GetTaskInstance 获取任务实例（用于转移时查询当前assignee）
func (r *BpmTaskTransferRepository) GetTaskInstance(ctx context.Context, taskID string) (*model.TaskInstance, error) {
	var task model.TaskInstance
	err := r.db.WithContext(ctx).Where("task_no = ?", taskID).First(&task).Error
	if err != nil {
		var id uint
		if _, scanErr := fmt.Sscanf(taskID, "%d", &id); scanErr == nil {
			err = r.db.WithContext(ctx).First(&task, id).Error
		}
	}
	return &task, err
}

// UpdateTaskAssignee 更新任务的当前处理人
func (r *BpmTaskTransferRepository) UpdateTaskAssignee(ctx context.Context, taskID string, toUserID int64, toUserName string) error {
	updates := map[string]interface{}{
		"assignee_id":   toUserID,
		"assignee_name": toUserName,
	}
	q := r.db.WithContext(ctx).Model(&model.TaskInstance{}).Where("task_no = ?", taskID).Updates(updates)
	if q.Error != nil {
		return q.Error
	}
	if q.RowsAffected == 0 {
		var id uint
		if _, scanErr := fmt.Sscanf(taskID, "%d", &id); scanErr == nil {
			return r.db.WithContext(ctx).Model(&model.TaskInstance{}).Where("id = ?", id).Updates(updates).Error
		}
	}
	return nil
}
