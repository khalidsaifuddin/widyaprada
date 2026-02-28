package cbt

import (
	"context"
	"errors"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *cbtUsecase) MulaiUjian(ctx context.Context, userID, examID string) (*entity.CBTStartResponse, error) {
	ok, err := u.assignRepo.IsParticipant(ctx, examID, userID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, entity.ErrCBTNotParticipant
	}

	resp, err := u.cbtRepo.CreateAttempt(ctx, examID, userID)
	if err != nil {
		if errors.Is(err, entity.ErrCBTAlreadyStarted) {
			return nil, err
		}
		return nil, err
	}
	return resp, nil
}
