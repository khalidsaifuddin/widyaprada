package example

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type ExampleUsecase interface {
	GetExampleList(ctx context.Context, request entity.GetExampleListRequest) (entity.GetExampleListResponse, error)
	GetExampleDetail(ctx context.Context, id int64) (entity.Example, error)
}

type exampleUsecase struct {
	exampleRepo repository.ExampleRepo
}

func NewExampleUsecase(exampleRepo repository.ExampleRepo) ExampleUsecase {
	return &exampleUsecase{exampleRepo: exampleRepo}
}
