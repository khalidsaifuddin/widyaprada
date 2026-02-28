package berita

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type BeritaUsecase interface {
	ListPublished(ctx context.Context, req entity.GetArticleListRequest) (*entity.GetBeritaListResponse, error)
	GetBySlug(ctx context.Context, slug string) (*entity.ArticleDetailResponse, error)
}

type beritaUsecase struct {
	articleRepo repository.ArticleRepo
}

func NewBeritaUsecase(articleRepo repository.ArticleRepo) BeritaUsecase {
	return &beritaUsecase{articleRepo: articleRepo}
}
