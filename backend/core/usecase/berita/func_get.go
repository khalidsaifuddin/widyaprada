package berita

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *beritaUsecase) GetBySlug(ctx context.Context, slug string) (*entity.ArticleDetailResponse, error) {
	a, err := u.articleRepo.GetBySlug(ctx, slug)
	if err != nil || a == nil {
		return nil, entity.ErrRecordNotFound
	}
	if a.Status != entity.ArticleStatusPublished {
		return nil, entity.ErrRecordNotFound
	}
	return &entity.ArticleDetailResponse{
		ID:           a.ID,
		Title:        a.Title,
		Slug:         a.Slug,
		Content:      a.Content,
		Excerpt:      a.Excerpt,
		ThumbnailURL: a.ThumbnailURL,
		GalleryURLs:  a.GalleryURLs,
		PublishedAt:  a.PublishedAt,
		AuthorName:   a.AuthorName,
		Category:     a.Category,
		CreatedAt:    a.CreatedAt,
	}, nil
}
