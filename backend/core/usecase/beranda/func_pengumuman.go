package beranda

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *berandaUsecase) GetPengumuman(ctx context.Context, userID string) (*entity.BerandaPengumumanResponse, error) {
	resp := &entity.BerandaPengumumanResponse{
		InfoJadwalUjikom: []entity.JadwalUjikomItem{},
		ExamsTersedia:    []entity.CBTExamTersediaItem{},
	}

	// Hasil seleksi administrasi
	applyStatus, _ := u.ujikomRepo.GetLatestApplicationByUser(ctx, userID)
	if applyStatus != nil {
		resp.HasilSeleksiAdmin = applyStatus
	}

	// Exams tersedia (dalam jadwal, bisa mulai)
	exams, err := u.cbtRepo.ListExamsTersedia(ctx, userID)
	if err == nil {
		resp.ExamsTersedia = exams
		resp.CanStartUjikom = len(exams) > 0
		for _, e := range exams {
			resp.InfoJadwalUjikom = append(resp.InfoJadwalUjikom, entity.JadwalUjikomItem{
				ExamID:        e.ID,
				ExamName:      e.Name,
				JadwalMulai:   e.JadwalMulai,
				JadwalSelesai: e.JadwalSelesai,
			})
		}
	}

	return resp, nil
}
