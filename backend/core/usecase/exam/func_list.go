package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *examUsecase) List(ctx context.Context, req entity.GetExamListRequest) (*entity.GetExamListResponse, error) {
	return u.examRepo.List(ctx, req)
}
