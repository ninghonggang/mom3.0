package service

import (
	"context"
	"fmt"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) List(ctx context.Context) ([]model.Post, int64, error) {
	return s.repo.List(ctx, 0)
}

func (s *PostService) GetByID(ctx context.Context, id string) (*model.Post, error) {
	var postID uint
	_, err := fmt.Sscanf(id, "%d", &postID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, postID)
}

func (s *PostService) Create(ctx context.Context, post *model.Post) error {
	return s.repo.Create(ctx, post)
}

func (s *PostService) Update(ctx context.Context, id string, post *model.Post) error {
	var postID uint
	_, err := fmt.Sscanf(id, "%d", &postID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, postID, map[string]interface{}{
		"post_name": post.PostName,
		"post_code": post.PostCode,
		"post_sort": post.PostSort,
		"status":    post.Status,
		"remark":    post.Remark,
	})
}

func (s *PostService) Delete(ctx context.Context, id string) error {
	var postID uint
	_, err := fmt.Sscanf(id, "%d", &postID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, postID)
}
