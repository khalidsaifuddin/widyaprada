package assignment

import (
	"context"
	"errors"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *assignmentUsecase) ApplyUjikom(ctx context.Context, userID string, req entity.ApplyUjikomRequest, docs []entity.ApplyUjikomDocumentInput) error {
	if req.JenisUjikom != entity.JenisUjikomPerpindahanJabatan && req.JenisUjikom != entity.JenisUjikomKenaikanTingkat {
		return errors.New("jenis_ujikom harus perpindahan_jabatan atau kenaikan_tingkat")
	}

	ok, err := u.ujikomRepo.HasPendingOrApproved(ctx, userID, req.JenisUjikom)
	if err != nil {
		return err
	}
	if ok {
		return entity.ErrApplyAlreadyExists
	}

	// Validate required docs from ref
	required, err := u.dokumenRepo.ListByJenisUjikom(ctx, req.JenisUjikom)
	if err != nil {
		return err
	}
	docMap := make(map[string]entity.ApplyUjikomDocumentInput)
	for _, d := range docs {
		docMap[d.DocumentType] = d
	}
	for _, r := range required {
		d, ok := docMap[r.Kode]
		if !ok {
			return entity.ErrDokumenPersyaratanReq
		}
		switch r.TipeInput {
		case "file":
			if d.FilePath == "" {
				return entity.ErrDokumenPersyaratanReq
			}
		case "text_portofolio", "text_essay":
			if d.PortofolioText == "" {
				return entity.ErrDokumenPersyaratanReq
			}
		}
	}

	return u.ujikomRepo.CreateApplication(ctx, userID, req.JenisUjikom, entity.ApplyStatusMenungguVerifikasi, docs)
}
