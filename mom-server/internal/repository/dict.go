package repository

import (
	"context"
	"mom-server/internal/model"

	"gorm.io/gorm"
)

type DictTypeRepository struct {
	db *gorm.DB
}

func NewDictTypeRepository(db *gorm.DB) *DictTypeRepository {
	return &DictTypeRepository{db: db}
}

func (r *DictTypeRepository) List(ctx context.Context, tenantID int64) ([]model.DictType, int64, error) {
	var list []model.DictType
	var total int64

	err := r.db.WithContext(ctx).Model(&model.DictType{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (r *DictTypeRepository) GetByID(ctx context.Context, id uint) (*model.DictType, error) {
	var dict model.DictType
	err := r.db.WithContext(ctx).First(&dict, id).Error
	return &dict, err
}

func (r *DictTypeRepository) Create(ctx context.Context, dict *model.DictType) error {
	return r.db.WithContext(ctx).Create(dict).Error
}

func (r *DictTypeRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DictType{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DictTypeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DictType{}, id).Error
}

type DictDataRepository struct {
	db *gorm.DB
}

func NewDictDataRepository(db *gorm.DB) *DictDataRepository {
	return &DictDataRepository{db: db}
}

func (r *DictDataRepository) GetByType(ctx context.Context, dictType string) ([]model.DictData, error) {
	var data []model.DictData
	err := r.db.WithContext(ctx).Where("dict_type = ?", dictType).Order("dict_sort ASC").Find(&data).Error
	return data, err
}

func (r *DictDataRepository) Create(ctx context.Context, data *model.DictData) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *DictDataRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.DictData{}).Where("id = ?", id).Updates(updates).Error
}

func (r *DictDataRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.DictData{}, id).Error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) List(ctx context.Context, tenantID int64) ([]model.Post, int64, error) {
	var list []model.Post
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Order("post_sort ASC").Find(&list).Error
	return list, total, err
}

func (r *PostRepository) GetByID(ctx context.Context, id uint) (*model.Post, error) {
	var post model.Post
	err := r.db.WithContext(ctx).First(&post, id).Error
	return &post, err
}

func (r *PostRepository) Create(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *PostRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.Post{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PostRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Post{}, id).Error
}
