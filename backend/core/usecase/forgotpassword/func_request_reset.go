package forgotpassword

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/email"
)

// RequestReset mengirim link reset password ke email (jika terdaftar)
func (u *forgotPasswordUsecase) RequestReset(ctx context.Context, emailAddr string) (*entity.ForgotPasswordResponse, error) {
	emailAddr = strings.TrimSpace(strings.ToLower(emailAddr))
	if emailAddr == "" || !entity.EmailRegex.MatchString(emailAddr) {
		return nil, entity.ErrInvalidEmailFormat
	}

	user, err := u.userRepo.FindByEmail(ctx, emailAddr)
	if err != nil {
		return nil, err
	}
	// Tetap return sukses meskipun email tidak terdaftar (security)
	if user == nil {
		return &entity.ForgotPasswordResponse{
			Message: "Jika email Anda terdaftar, Anda akan menerima tautan untuk mengatur ulang kata sandi. Periksa juga folder spam.",
		}, nil
	}

	plainToken, hashedToken, err := auth.GenerateResetToken()
	if err != nil {
		return nil, err
	}

	hours := u.cfg.ResetTokenExpiryHr
	if hours <= 0 {
		hours = 1
	}
	expiresAt := time.Now().UTC().Add(time.Duration(hours) * time.Hour)
	if err := u.tokenRepo.Create(ctx, user.ID, hashedToken, expiresAt); err != nil {
		return nil, err
	}

	resetURL := fmt.Sprintf("%s/auth/reset-password?token=%s", strings.TrimSuffix(u.cfg.FrontendURL, "/"), plainToken)
	if err := u.emailService.SendPasswordResetLink(email.SendPasswordResetLinkParams{
		To:       emailAddr,
		Name:     user.Name,
		ResetURL: resetURL,
	}); err != nil {
		// Log error, tetap return sukses
	}

	return &entity.ForgotPasswordResponse{
		Message: "Jika email Anda terdaftar, Anda akan menerima tautan untuk mengatur ulang kata sandi. Periksa juga folder spam.",
	}, nil
}
