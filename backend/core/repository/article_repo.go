package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type ArticleRepo interface {
	List(ctx context.Context, req entity.GetArticleListRequest, satkerFilter *string) (*entity.GetArticleListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.Article, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Article, error)
	SlugExists(ctx context.Context, slug, excludeID string) (bool, error)
	Create(ctx context.Context, a *entity.Article) (string, error)
	Update(ctx context.Context, a *entity.Article) error
	Delete(ctx context.Context, id string) error
	ListPublished(ctx context.Context, req entity.GetArticleListRequest, limit int) ([]entity.ArticlePublicItem, int64, error)
}
