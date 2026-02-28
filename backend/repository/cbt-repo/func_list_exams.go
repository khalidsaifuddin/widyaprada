package cbtrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
)

func (r *cbtRepo) ListExamsTersedia(ctx context.Context, userID string) ([]entity.CBTExamTersediaItem, error) {
	now := time.Now().UTC()

	var exams []examrepo.Exam
	err := r.db.WithContext(ctx).Table("exams").
		Select("exams.id, exams.code, exams.name, exams.jadwal_mulai, exams.jadwal_selesai, exams.durasi_menit").
		Joins("INNER JOIN exam_participants ep ON ep.exam_id = exams.id AND ep.user_id = ?", userID).
		Where("exams.deleted_at IS NULL").
		Where("exams.status = ?", entity.ExamStatusDiterbitkan).
		Where("exams.jadwal_mulai <= ?", now).
		Where("exams.jadwal_selesai >= ?", now).
		Where("NOT EXISTS (SELECT 1 FROM exam_attempts ea WHERE ea.exam_id = exams.id AND ea.user_id = ? AND ea.submitted_at IS NOT NULL)", userID).
		Find(&exams).Error
	if err != nil {
		return nil, err
	}

	items := make([]entity.CBTExamTersediaItem, len(exams))
	for i := range exams {
		e := &exams[i]
		jadwalMulai, jadwalSelesai := "", ""
		if e.JadwalMulai != nil {
			jadwalMulai = e.JadwalMulai.UTC().Format(time.RFC3339)
		}
		if e.JadwalSelesai != nil {
			jadwalSelesai = e.JadwalSelesai.UTC().Format(time.RFC3339)
		}
		items[i] = entity.CBTExamTersediaItem{
			ID:            e.ID,
			Code:          e.Code,
			Name:          e.Name,
			JadwalMulai:   jadwalMulai,
			JadwalSelesai: jadwalSelesai,
			DurasiMenit:   e.DurasiMenit,
		}
	}
	return items, nil
}
