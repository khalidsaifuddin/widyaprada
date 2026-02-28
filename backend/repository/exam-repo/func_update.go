package examrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *examRepo) Update(ctx context.Context, e *entity.ExamDetail) error {
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
	return r.db.WithContext(ctx).Model(&Exam{}).Where("id = ?", e.ID).Updates(map[string]interface{}{
		"code":                  e.Code,
		"name":                  e.Name,
		"jadwal_mulai":          jadwalMulai,
		"jadwal_selesai":        jadwalSelesai,
		"durasi_menit":          e.DurasiMenit,
		"status":                e.Status,
		"shuffle_questions":     e.ShuffleQuestions,
		"tampilkan_leaderboard": e.TampilkanLeaderboard,
		"updated_at":            now,
	}).Error
}
