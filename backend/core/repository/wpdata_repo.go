package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type WPDataRepo interface {
	List(ctx context.Context, req entity.GetWPDataListRequest, satkerFilter *string) (*entity.GetWPDataListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.WidyapradaData, error)
	GetByNIP(ctx context.Context, nip, excludeID string) (*entity.WidyapradaData, error)
	Create(ctx context.Context, w *entity.WidyapradaData) (string, error)
	Update(ctx context.Context, w *entity.WidyapradaData) error
	Delete(ctx context.Context, id, reason string) error
}
