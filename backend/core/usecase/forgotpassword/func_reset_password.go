package forgotpassword

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
)

// ResetPassword mengatur ulang password dengan token
func (u *forgotPasswordUsecase) ResetPassword(ctx context.Context, token, password, passwordConfirm string) (*entity.ResetPasswordResponse, error) {
	token = strings.TrimSpace(token)
	password = strings.TrimSpace(password)
	passwordConfirm = strings.TrimSpace(passwordConfirm)

	if token == "" {
		return nil, entity.ErrResetTokenInvalid
	}
	if len(password) < 8 {
		return nil, entity.ErrPasswordTooShort
	}
	if password != passwordConfirm {
		return nil, entity.ErrPasswordMismatch
	}

	tokenHash := auth.HashResetToken(token)
	tokenRecord, err := u.tokenRepo.FindByTokenHash(ctx, tokenHash)
	if err != nil {
		return nil, err
	}
	if tokenRecord == nil {
		return nil, entity.ErrResetTokenInvalid
	}

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.UpdatePassword(ctx, tokenRecord.UserID, passwordHash); err != nil {
		return nil, err
	}

	if err := u.tokenRepo.Delete(ctx, tokenRecord.ID); err != nil {
		// Log, tapi tetap return sukses
	}

	return &entity.ResetPasswordResponse{
		Message: "Kata sandi berhasil diubah. Silakan masuk dengan kata sandi baru.",
	}, nil
}
