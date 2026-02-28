package example

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *exampleUsecase) GetExampleDetail(ctx context.Context, id int64) (entity.Example, error) {
	return u.exampleRepo.GetExampleDetail(ctx, id)
}
