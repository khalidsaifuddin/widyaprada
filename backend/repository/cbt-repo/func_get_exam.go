package cbtrepo

import (
	"context"
	"errors"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
	"gorm.io/gorm"
)

func (r *cbtRepo) GetExamByIDForInstructions(ctx context.Context, examID, userID string) (*entity.CBTExamDetailResponse, error) {
	var exam examrepo.Exam
	if err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", examID).First(&exam).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrExamNotFound
		}
		return nil, err
	}

	// Cek peserta
	var participantCount int64
	if err := r.db.WithContext(ctx).Table("exam_participants").
		Where("exam_id = ? AND user_id = ?", examID, userID).
		Count(&participantCount).Error; err != nil {
		return nil, err
	}
	if participantCount == 0 {
		return nil, entity.ErrCBTNotParticipant
	}

	jadwalMulai, jadwalSelesai := "", ""
	if exam.JadwalMulai != nil {
		jadwalMulai = exam.JadwalMulai.UTC().Format(time.RFC3339)
	}
	if exam.JadwalSelesai != nil {
		jadwalSelesai = exam.JadwalSelesai.UTC().Format(time.RFC3339)
	}

	resp := &entity.CBTExamDetailResponse{
		ID:            exam.ID,
		Code:          exam.Code,
		Name:          exam.Name,
		JadwalMulai:   jadwalMulai,
		JadwalSelesai: jadwalSelesai,
		DurasiMenit:   exam.DurasiMenit,
	}

	now := time.Now().UTC()

	if exam.Status != entity.ExamStatusDiterbitkan {
		resp.DapatMulai = false
		resp.Alasan = "Ujian belum diterbitkan."
		return resp, nil
	}

	if exam.JadwalMulai != nil && now.Before(*exam.JadwalMulai) {
		resp.DapatMulai = false
		resp.Alasan = "Ujian belum dimulai."
		return resp, nil
	}

	if exam.JadwalSelesai != nil && now.After(*exam.JadwalSelesai) {
		resp.DapatMulai = false
		resp.Alasan = "Ujian sudah selesai."
		return resp, nil
	}

	// Cek sudah submit?
	var submittedCount int64
	if err := r.db.WithContext(ctx).Model(&examrepo.ExamAttempt{}).
		Where("exam_id = ? AND user_id = ? AND submitted_at IS NOT NULL", examID, userID).
		Count(&submittedCount).Error; err != nil {
		return nil, err
	}
	if submittedCount > 0 {
		resp.DapatMulai = false
		resp.Alasan = "Anda sudah menyelesaikan ujian ini."
		return resp, nil
	}

	resp.DapatMulai = true
	return resp, nil
}
