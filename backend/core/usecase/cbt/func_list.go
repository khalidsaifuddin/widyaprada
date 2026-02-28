package cbt

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *cbtUsecase) ListUjianTersedia(ctx context.Context, userID string) (*entity.CBTListExamsResponse, error) {
	items, err := u.cbtRepo.ListExamsTersedia(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &entity.CBTListExamsResponse{Items: items}, nil
}
