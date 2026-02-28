package cbt

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *cbtUsecase) GetRiwayatHasil(ctx context.Context, userID string) (*entity.CBTHistoryResponse, error) {
	items, err := u.cbtRepo.ListHistory(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &entity.CBTHistoryResponse{Items: items}, nil
}
