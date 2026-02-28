package example

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *exampleUsecase) GetExampleList(ctx context.Context, request entity.GetExampleListRequest) (entity.GetExampleListResponse, error) {
	return u.exampleRepo.GetExampleList(ctx, request)
}
