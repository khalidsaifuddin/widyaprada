package berita

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *beritaUsecase) ListPublished(ctx context.Context, req entity.GetArticleListRequest) (*entity.GetBeritaListResponse, error) {
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	items, total, err := u.articleRepo.ListPublished(ctx, req, int(pageSize))
	if err != nil {
		return nil, err
	}
	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}
	page := req.Page
	if page < 1 {
		page = 1
	}
	return &entity.GetBeritaListResponse{
		Items:     items,
		TotalPage: totalPage,
		TotalData: total,
		Page:      page,
		PageSize:  pageSize,
	}, nil
}
