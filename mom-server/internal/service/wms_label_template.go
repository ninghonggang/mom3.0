package service

import (
	"context"
	"fmt"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

// WmsLabelTemplateService 标签模板服务层
type WmsLabelTemplateService struct {
	repo *repository.WmsLabelTemplateRepository
}

// NewWmsLabelTemplateService 创建标签模板服务实例
func NewWmsLabelTemplateService(repo *repository.WmsLabelTemplateRepository) *WmsLabelTemplateService {
	return &WmsLabelTemplateService{repo: repo}
}

// List 获取标签模板列表
func (s *WmsLabelTemplateService) List(ctx context.Context, query *model.WmsLabelTemplateQueryVO) ([]model.WmsLabelTemplate, int64, error) {
	// 设置默认分页
	if query != nil {
		if query.Page <= 0 {
			query.Page = 1
		}
		if query.PageSize <= 0 {
			query.PageSize = 20
		}
	}
	return s.repo.List(ctx, query)
}

// Get 获取标签模板详情
func (s *WmsLabelTemplateService) Get(ctx context.Context, id string) (*model.WmsLabelTemplate, error) {
	var templateID uint64
	_, err := fmt.Sscanf(id, "%d", &templateID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, templateID)
}

// Create 创建标签模板
func (s *WmsLabelTemplateService) Create(ctx context.Context, req *model.WmsLabelTemplateCreateReqVO) error {
	template := &model.WmsLabelTemplate{
		TemplateCode: req.TemplateCode,
		TemplateName: req.TemplateName,
		TemplateType: req.TemplateType,
		Width:        req.Width,
		Height:       req.Height,
		Content:      req.Content,
		Status:       "ACTIVE",
	}
	return s.repo.Create(ctx, template)
}

// Update 更新标签模板
func (s *WmsLabelTemplateService) Update(ctx context.Context, id string, req *model.WmsLabelTemplateUpdateReqVO) error {
	var templateID uint64
	_, err := fmt.Sscanf(id, "%d", &templateID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"template_name": req.TemplateName,
		"template_type": req.TemplateType,
		"width":         req.Width,
		"height":        req.Height,
		"content":       req.Content,
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	return s.repo.Update(ctx, templateID, updates)
}

// Delete 删除标签模板
func (s *WmsLabelTemplateService) Delete(ctx context.Context, id string) error {
	var templateID uint64
	_, err := fmt.Sscanf(id, "%d", &templateID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, templateID)
}
