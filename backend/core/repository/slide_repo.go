package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type SlideRepo interface {
	List(ctx context.Context, req entity.GetSlideListRequest, satkerFilter *string) (*entity.GetSlideListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.Slide, error)
	Create(ctx context.Context, s *entity.Slide) error
	Update(ctx context.Context, s *entity.Slide) error
	Delete(ctx context.Context, id string) error
	ListPublishedForLanding(ctx context.Context, limit int) ([]entity.SlidePublicItem, error)
}
