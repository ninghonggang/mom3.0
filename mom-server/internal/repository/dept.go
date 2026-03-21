package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type DeptRepository struct {
	db *gorm.DB
}

func NewDeptRepository(db *gorm.DB) *DeptRepository {
	return &DeptRepository{db: db}
}

func (r *DeptRepository) List(ctx context.Context, tenantID int64) ([]model.Dept, error) {
	var depts []model.Dept
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("dept_sort ASC").Find(&depts).Error
	return depts, err
}

func (r *DeptRepository) Tree(ctx context.Context, tenantID int64) ([]model.Dept, error) {
	depts, err := r.List(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	return r.BuildTree(depts), nil
}

func (r *DeptRepository) BuildTree(depts []model.Dept) []model.Dept {
	var result []model.Dept
	deptMap := make(map[int64]*model.Dept)

	for i := range depts {
		deptMap[depts[i].ID] = &depts[i]
	}

	for i := range depts {
		if depts[i].ParentID == 0 {
			result = append(result, depts[i])
		} else {
			if parent, ok := deptMap[depts[i].ParentID]; ok {
				parent.Children = append(parent.Children, depts[i])
			}
		}
	}

	return result
}

func (r *DeptRepository) GetByID(ctx context.Context, id uint) (*model.Dept, error) {
	var dept model.Dept
	err := r.db.WithContext(ctx).First(&dept, id).Error
	return &dept, err
}

func (r *DeptRepository) Create(ctx context.Context, dept *model.Dept) error {
	return r.db.WithContext(ctx).Create(dept).Error
}

func (r *DeptRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Dept{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DeptRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Dept{}, id).Error
}

// Add Children field to Dept model for tree structure
func init() {
	// This is handled in the model definition
}
