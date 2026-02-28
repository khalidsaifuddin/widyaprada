package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type LinkRepo interface {
	List(ctx context.Context, req entity.GetLinkListRequest, satkerFilter *string) (*entity.GetLinkListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.Link, error)
	Create(ctx context.Context, l *entity.Link) (string, error)
	Update(ctx context.Context, l *entity.Link) error
	Delete(ctx context.Context, id string) error
	ListActiveForLanding(ctx context.Context, limit int) ([]entity.LinkPublicItem, error)
}
