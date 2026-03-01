package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// CBTRepo repository untuk CBT
type CBTRepo interface {
	// ListExamsTersedia daftar ujian tersedia: Diterbitkan, jadwal aktif, user peserta, belum submit
	ListExamsTersedia(ctx context.Context, userID string) ([]entity.CBTExamTersediaItem, error)
	// GetExamByIDForInstructions detail ujian untuk halaman petunjuk
	GetExamByIDForInstructions(ctx context.Context, examID, userID string) (*entity.CBTExamDetailResponse, error)
	// CreateAttempt membuat attempt baru
	CreateAttempt(ctx context.Context, examID, userID string) (*entity.CBTStartResponse, error)
	// GetAttemptByID mendapatkan attempt by ID
	GetAttemptByID(ctx context.Context, attemptID string) (examID, userID string, submitted bool, err error)
	// GetQuestionIDsForAttempt daftar question IDs untuk attempt (expand contents, shuffle jika perlu)
	GetQuestionIDsForAttempt(ctx context.Context, attemptID string) ([]string, error)
	// GetQuestionForCBT detail soal untuk CBT (tanpa kunci)
	GetQuestionForCBT(ctx context.Context, questionID string) (*entity.CBTQuestionItem, error)
	// SaveAnswer simpan jawaban
	SaveAnswer(ctx context.Context, attemptID, questionID, answerValue string) error
	// GetAnswers map questionID -> answerValue
	GetAnswers(ctx context.Context, attemptID string) (map[string]string, error)
	// SubmitAttempt set submitted_at, hitung skor PG/B-S
	SubmitAttempt(ctx context.Context, attemptID string) (*entity.CBTSubmitResponse, error)
	// ListHistory riwayat attempt user
	ListHistory(ctx context.Context, userID string) ([]entity.CBTHistoryItem, error)
}
