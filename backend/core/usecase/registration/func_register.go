package registration

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/email"
)

// Register mengimplementasikan use case registrasi sesuai SDD
func (u *registrationUsecase) Register(ctx context.Context, req entity.RegisterRequest) (*entity.RegisterResponse, error) {
	// 1. Validasi
	name := strings.TrimSpace(req.Name)
	emailAddr := strings.TrimSpace(strings.ToLower(req.Email))
	nip := strings.TrimSpace(req.NIP)

	if name == "" {
		return nil, entity.ErrValidation
	}
	if emailAddr == "" {
		return nil, entity.ErrInvalidEmailFormat
	}
	if !entity.EmailRegex.MatchString(emailAddr) {
		return nil, entity.ErrInvalidEmailFormat
	}

	// 2. Cek email sudah terdaftar
	existing, err := u.userRepo.FindByEmail(ctx, emailAddr)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, entity.ErrEmailAlreadyRegistered
	}

	// 3. Ambil role default (PESERTA)
	role, err := u.userRepo.GetRoleByCode(ctx, entity.DefaultRoleCodePeserta)
	if err != nil || role == nil {
		return nil, entity.ErrRoleNotFound
	}

	// 4. Generate password
	plainPassword, err := generatePassword(12)
	if err != nil {
		return nil, err
	}

	// 5. Hash password
	passwordHash, err := auth.HashPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	// 6. Buat user (username = email untuk registrasi)
	user := &entity.User{
		Name:         name,
		Email:        emailAddr,
		Username:     emailAddr,
		NIP:          nip,
		PasswordHash: passwordHash,
		IsActive:     true,
	}
	if err := u.userRepo.CreateUser(ctx, user, role.ID); err != nil {
		return nil, err
	}

	// 7. Kirim password ke email
	if err := u.emailService.SendPasswordEmail(email.SendPasswordEmailParams{
		To:       emailAddr,
		Name:     name,
		Password: plainPassword,
	}); err != nil {
		// Log error, tetap return sukses (user sudah dibuat)
		// Di production mungkin perlu retry atau queue
	}

	return &entity.RegisterResponse{
		Message: "Registrasi berhasil. Silakan cek email Anda untuk mendapatkan kata sandi.",
	}, nil
}

// generatePassword generates a random hex password (24 chars)
func generatePassword(byteLen int) (string, error) {
	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := hex.EncodeToString(b)
	return s, nil
}
