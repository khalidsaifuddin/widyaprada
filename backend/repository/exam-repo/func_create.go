package examrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (r *examRepo) Create(ctx context.Context, e *entity.ExamDetail) error {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	var jadwalMulai, jadwalSelesai *time.Time
	if e.JadwalMulai != "" {
		if t, err := time.Parse(time.RFC3339, e.JadwalMulai); err == nil {
			t := t.UTC()
			jadwalMulai = &t
		}
	}
	if e.JadwalSelesai != "" {
		if t, err := time.Parse(time.RFC3339, e.JadwalSelesai); err == nil {
			t := t.UTC()
			jadwalSelesai = &t
		}
	}
	now := time.Now().UTC()
	dto := Exam{
		ID:                   e.ID,
		Code:                 e.Code,
		Name:                 e.Name,
		JadwalMulai:          jadwalMulai,
		JadwalSelesai:        jadwalSelesai,
		DurasiMenit:          e.DurasiMenit,
		Status:               e.Status,
		VerificationStatus:   e.VerificationStatus,
		ShuffleQuestions:     e.ShuffleQuestions,
		TampilkanLeaderboard: e.TampilkanLeaderboard,
		CreatedAt:            &now,
		UpdatedAt:            &now,
	}
	if dto.Status == "" {
		dto.Status = entity.ExamStatusDraft
	}
	if dto.VerificationStatus == "" {
		dto.VerificationStatus = entity.QuestionVerifBelum
	}
	return r.db.WithContext(ctx).Create(&dto).Error
}
