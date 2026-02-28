package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// AssignmentRepo repository untuk Tugas Saya, Hasil, Leaderboard
type AssignmentRepo interface {
	// ListMyAssignments list ujian yang user ikuti (exam_participants), dengan status & score dari exam_attempts
	ListMyAssignments(ctx context.Context, userID string, req entity.GetAssignmentListRequest) (*entity.GetAssignmentListResponse, error)
	// GetAttemptByExamUser mendapatkan attempt user untuk exam (untuk result)
	GetAttemptByExamUser(ctx context.Context, examID, userID string) (*entity.AssignmentResultResponse, error)
	// GetLeaderboard leaderboard exam (hanya jika tampilkan_leaderboard=true, urut score desc)
	GetLeaderboard(ctx context.Context, examID string) ([]entity.LeaderboardItem, error)
	// IsParticipant cek apakah user peserta exam
	IsParticipant(ctx context.Context, examID, userID string) (bool, error)
	// GetExamForLeaderboard mengambil nama exam dan tampilkan_leaderboard
	GetExamForLeaderboard(ctx context.Context, examID string) (name string, tampilkan bool, err error)

	// ListForDashboard list assignments untuk dashboard (exam status Diterbitkan/Berlangsung/Selesai)
	ListForDashboard(ctx context.Context, userID string, limit, page int64) ([]entity.DashboardAssignmentItem, int64, error)
}
