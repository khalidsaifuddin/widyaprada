package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type ExampleRepo interface {
	GetExampleList(ctx context.Context, request entity.GetExampleListRequest) (entity.GetExampleListResponse, error)
	GetExampleDetail(ctx context.Context, id int64) (entity.Example, error)
}
