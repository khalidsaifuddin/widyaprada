package wpdatarepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *wpdataRepo) List(ctx context.Context, req entity.GetWPDataListRequest, satkerFilter *string) (*entity.GetWPDataListResponse, error) {
	var items []WidyapradaData
	db := r.db.WithContext(ctx).Model(&WidyapradaData{}).Where("deleted_at IS NULL")

	if req.Q != "" {
		q := "%" + req.Q + "%"
		db = db.Where("(nip ILIKE ? OR nama_lengkap ILIKE ? OR email ILIKE ?)", q, q, q)
	}
	if req.SatkerID != "" {
		db = db.Where("satker_id = ?", req.SatkerID)
	}
	if satkerFilter != nil && *satkerFilter != "" {
		db = db.Where("satker_id = ?", *satkerFilter)
	}
	if req.UnitKerja != "" {
		db = db.Where("unit_kerja ILIKE ?", "%"+req.UnitKerja+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	page, pageSize := req.Page, req.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	sortBy := req.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	sortOrder := req.SortOrder
	if sortOrder == "" {
		sortOrder = "desc"
	}
	order := sortBy + " " + sortOrder

	offset := (page - 1) * pageSize
	if err := db.Order(order).Offset(int(offset)).Limit(int(pageSize)).Find(&items).Error; err != nil {
		return nil, err
	}

	list := make([]entity.WPDataListItem, len(items))
	for i := range items {
		list[i] = items[i].ToListItem()
	}

	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}

	return &entity.GetWPDataListResponse{
		Items:     list,
		TotalPage: totalPage,
		TotalData: total,
		Page:      page,
		PageSize:  pageSize,
	}, nil
}

func (r *wpdataRepo) GetByID(ctx context.Context, id string) (*entity.WidyapradaData, error) {
	var w WidyapradaData
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&w).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toWPDataEntity(&w), nil
}

func (r *wpdataRepo) GetByNIP(ctx context.Context, nip, excludeID string) (*entity.WidyapradaData, error) {
	var w WidyapradaData
	db := r.db.WithContext(ctx).Where("nip = ? AND deleted_at IS NULL", nip)
	if excludeID != "" {
		db = db.Where("id != ?", excludeID)
	}
	err := db.First(&w).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toWPDataEntity(&w), nil
}

func toWPDataEntity(w *WidyapradaData) *entity.WidyapradaData {
	createdAt, updatedAt, deletedAt := "", "", ""
	if w.CreatedAt != nil {
		createdAt = w.CreatedAt.UTC().Format(time.RFC3339)
	}
	if w.UpdatedAt != nil {
		updatedAt = w.UpdatedAt.UTC().Format(time.RFC3339)
	}
	return &entity.WidyapradaData{
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
		UserID:                   w.UserID,
		DeletedReason:            w.DeletedReason,
		CreatedAt:                createdAt,
		UpdatedAt:                updatedAt,
		DeletedAt:                deletedAt,
	}
}

func (r *wpdataRepo) Create(ctx context.Context, w *entity.WidyapradaData) (string, error) {
	now := time.Now().UTC()
	status := w.Status
	if status == "" {
		status = entity.WPDataStatusAktif
	}
	dto := WidyapradaData{
		ID:                       uuid.New().String(),
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
		Status:                   status,
		Keterangan:               w.Keterangan,
		UserID:                   w.UserID,
		CreatedAt:                &now,
		UpdatedAt:                &now,
	}
	if err := r.db.WithContext(ctx).Create(&dto).Error; err != nil {
		return "", err
	}
	return dto.ID, nil
}

func (r *wpdataRepo) Update(ctx context.Context, w *entity.WidyapradaData) error {
	upd := map[string]interface{}{
		"nip":                       w.NIP,
		"nama_lengkap":              w.NamaLengkap,
		"jenis_kelamin":             w.JenisKelamin,
		"golongan_ruang":            w.GolonganRuang,
		"pangkat":                   w.Pangkat,
		"jenjang_jabatan_fungsional": w.JenjangJabatanFungsional,
		"satker_id":                 w.SatkerID,
		"unit_kerja":                w.UnitKerja,
		"pendidikan_terakhir":       w.PendidikanTerakhir,
		"tmt_golongan":              w.TMTGolongan,
		"tmt_jabatan_fungsional":    w.TMTJabatanFungsional,
		"no_sk_pengangkatan":        w.NoSK,
		"no_hp":                     w.NoHP,
		"email":                     w.Email,
		"alamat":                    w.Alamat,
		"status":                    w.Status,
		"keterangan":                w.Keterangan,
		"updated_at":                time.Now().UTC(),
	}
	return r.db.WithContext(ctx).Model(&WidyapradaData{}).Where("id = ?", w.ID).Updates(upd).Error
}

func (r *wpdataRepo) Delete(ctx context.Context, id, reason string) error {
	return r.db.WithContext(ctx).Model(&WidyapradaData{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at":     time.Now().UTC(),
			"deleted_reason": reason,
			"updated_at":     time.Now().UTC(),
		}).Error
}
