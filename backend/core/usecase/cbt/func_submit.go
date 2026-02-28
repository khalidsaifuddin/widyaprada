package cbt

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *cbtUsecase) SubmitUjian(ctx context.Context, userID, attemptID string) (*entity.CBTSubmitResponse, error) {
	_, attemptUserID, _, err := u.cbtRepo.GetAttemptByID(ctx, attemptID)
	if err != nil || attemptUserID == "" {
		return nil, entity.ErrCBTAttemptNotFound
	}
	if attemptUserID != userID {
		return nil, entity.ErrCBTAttemptNotOwned
	}

	return u.cbtRepo.SubmitAttempt(ctx, attemptID)
}
