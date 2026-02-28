package wpdata

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *wpdataUsecase) getSatkerFilter(actor *ActorContext) *string {
	if actor == nil || actor.IsSuperAdmin() {
		return nil
	}
	return actor.SatkerID
}

func (u *wpdataUsecase) List(ctx context.Context, req entity.GetWPDataListRequest, actor *ActorContext) (*entity.GetWPDataListResponse, error) {
	return u.wpdataRepo.List(ctx, req, u.getSatkerFilter(actor))
}

func (u *wpdataUsecase) Get(ctx context.Context, id string, actor *ActorContext) (*entity.WPDataDetailResponse, error) {
	w, err := u.wpdataRepo.GetByID(ctx, id)
	if err != nil || w == nil {
		return nil, entity.ErrRecordNotFound
	}
	return &entity.WPDataDetailResponse{
		ID:                       w.ID,
		NIP:                      w.NIP,
		NamaLengkap:              w.NamaLengkap,
		JenisKelamin:             w.JenisKelamin,
		GolonganRuang:            w.GolonganRuang,
		Pangkat:                  w.Pangkat,
		JenjangJabatanFungsional: w.JenjangJabatanFungsional,
		SatkerID:                 w.SatkerID,
		UnitKerja:                w.UnitKerja,
		PendidikanTerakhir:       w.PendidikanTerakhir,
		TMTGolongan:              w.TMTGolongan,
		TMTJabatanFungsional:     w.TMTJabatanFungsional,
		NoSK:                     w.NoSK,
		NoHP:                     w.NoHP,
		Email:                    w.Email,
		Alamat:                   w.Alamat,
		Status:                   w.Status,
		Keterangan:               w.Keterangan,
		CreatedAt:                w.CreatedAt,
		UpdatedAt:                w.UpdatedAt,
	}, nil
}

func (u *wpdataUsecase) Create(ctx context.Context, req entity.CreateWPDataRequest, actor *ActorContext) (*entity.WPDataDetailResponse, error) {
	existing, _ := u.wpdataRepo.GetByNIP(ctx, req.NIP, "")
	if existing != nil {
		return nil, entity.ErrNIPExists
	}
	status := req.Status
	if status == "" {
		status = entity.WPDataStatusAktif
	}
	w := &entity.WidyapradaData{
		NIP:                      req.NIP,
		NamaLengkap:              req.NamaLengkap,
		JenisKelamin:             req.JenisKelamin,
		GolonganRuang:            req.GolonganRuang,
		Pangkat:                  req.Pangkat,
		JenjangJabatanFungsional: req.JenjangJabatanFungsional,
		SatkerID:                 req.SatkerID,
		UnitKerja:                req.UnitKerja,
		PendidikanTerakhir:       req.PendidikanTerakhir,
		TMTGolongan:              req.TMTGolongan,
		TMTJabatanFungsional:     req.TMTJabatanFungsional,
		NoSK:                     req.NoSK,
		NoHP:                     req.NoHP,
		Email:                    req.Email,
		Alamat:                   req.Alamat,
		Status:                   status,
		Keterangan:               req.Keterangan,
	}
	id, err := u.wpdataRepo.Create(ctx, w)
	if err != nil {
		return nil, err
	}
	return u.Get(ctx, id, actor)
}

func (u *wpdataUsecase) Update(ctx context.Context, id string, req entity.UpdateWPDataRequest, actor *ActorContext) (*entity.WPDataDetailResponse, error) {
	existing, err := u.wpdataRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return nil, entity.ErrRecordNotFound
	}
	if req.NIP != "" {
		dup, _ := u.wpdataRepo.GetByNIP(ctx, req.NIP, id)
		if dup != nil {
			return nil, entity.ErrNIPExists
		}
		existing.NIP = req.NIP
	}
	if req.NamaLengkap != "" {
		existing.NamaLengkap = req.NamaLengkap
	}
	if req.JenisKelamin != "" {
		existing.JenisKelamin = req.JenisKelamin
	}
	if req.GolonganRuang != "" {
		existing.GolonganRuang = req.GolonganRuang
	}
	if req.Pangkat != "" {
		existing.Pangkat = req.Pangkat
	}
	if req.JenjangJabatanFungsional != "" {
		existing.JenjangJabatanFungsional = req.JenjangJabatanFungsional
	}
	if req.SatkerID != "" {
		existing.SatkerID = req.SatkerID
	}
	if req.UnitKerja != "" {
		existing.UnitKerja = req.UnitKerja
	}
	if req.PendidikanTerakhir != "" {
		existing.PendidikanTerakhir = req.PendidikanTerakhir
	}
	if req.TMTGolongan != "" {
		existing.TMTGolongan = req.TMTGolongan
	}
	if req.TMTJabatanFungsional != "" {
		existing.TMTJabatanFungsional = req.TMTJabatanFungsional
	}
	if req.NoSK != "" {
		existing.NoSK = req.NoSK
	}
	if req.NoHP != "" {
		existing.NoHP = req.NoHP
	}
	if req.Email != "" {
		existing.Email = req.Email
	}
	if req.Alamat != "" {
		existing.Alamat = req.Alamat
	}
	if req.Status != "" {
		existing.Status = req.Status
	}
	if req.Keterangan != "" {
		existing.Keterangan = req.Keterangan
	}
	if err := u.wpdataRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return u.Get(ctx, id, actor)
}

func (u *wpdataUsecase) Delete(ctx context.Context, id string, reason string, actor *ActorContext) error {
	existing, err := u.wpdataRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return entity.ErrRecordNotFound
	}
	if reason == "" {
		return entity.ErrInvalidData
	}
	return u.wpdataRepo.Delete(ctx, id, reason)
}
