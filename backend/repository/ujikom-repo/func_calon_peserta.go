package ujikomrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

type calonPesertaRow struct {
	ID           string
	UserID       string
	UserName     string
	UserEmail    string
	JenisUjikom  string
	StatusKode   string
	CatatanTolak string
	AppliedAt    *time.Time
}

func (r *ujikomRepo) ListCalonPeserta(ctx context.Context, req entity.GetCalonPesertaListRequest, satkerFilter *string) (*entity.GetCalonPesertaListResponse, error) {
	db := r.db.WithContext(ctx).Table("ujikom_application").
		Select("ujikom_application.id, ujikom_application.user_id, users.name as user_name, users.email as user_email, ujikom_application.jenis_ujikom, ujikom_application.status_kode, ujikom_application.catatan_tolak, ujikom_application.created_at as applied_at").
		Joins("LEFT JOIN users ON users.id = ujikom_application.user_id").
		Where("ujikom_application.deleted_at IS NULL")

	if req.StatusVerifikasi != "" {
		db = db.Where("ujikom_application.status_kode = ?", req.StatusVerifikasi)
	}
	if req.Q != "" {
		q := "%" + req.Q + "%"
		db = db.Where("(users.name ILIKE ? OR users.email ILIKE ?)", q, q)
	}
	if satkerFilter != nil && *satkerFilter != "" {
		db = db.Where("users.satker_id = ?", *satkerFilter)
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
	offset := (page - 1) * pageSize

	var rows []calonPesertaRow
	if err := db.Order("ujikom_application.created_at DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}

	items := make([]entity.CalonPesertaListItem, len(rows))
	for i := range rows {
		appliedAt := ""
		if rows[i].AppliedAt != nil {
			appliedAt = rows[i].AppliedAt.UTC().Format(time.RFC3339)
		}
		items[i] = entity.CalonPesertaListItem{
			ID:           rows[i].ID,
			UserID:       rows[i].UserID,
			UserName:     rows[i].UserName,
			UserEmail:    rows[i].UserEmail,
			JenisUjikom:  rows[i].JenisUjikom,
			StatusKode:   rows[i].StatusKode,
			CatatanTolak: rows[i].CatatanTolak,
			AppliedAt:    appliedAt,
		}
	}

	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}

	return &entity.GetCalonPesertaListResponse{
		Items:     items,
		TotalPage: totalPage,
		TotalData: total,
		Page:      page,
		PageSize:  pageSize,
	}, nil
}

func (r *ujikomRepo) GetCalonPesertaByID(ctx context.Context, id string) (*entity.CalonPesertaDetailResponse, error) {
	var app UjikomApplication
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&app).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var userName, userEmail string
	_ = r.db.WithContext(ctx).Table("users").Where("id = ?", app.UserID).Select("name, email").Row().Scan(&userName, &userEmail)

	appliedAt := ""
	if app.CreatedAt != nil {
		appliedAt = app.CreatedAt.UTC().Format(time.RFC3339)
	}

	var docs []UjikomApplicationDocument
	_ = r.db.WithContext(ctx).Where("ujikom_application_id = ?", id).Find(&docs).Error

	docItems := make([]entity.CalonPesertaDocumentItem, len(docs))
	for i := range docs {
		docItems[i] = entity.CalonPesertaDocumentItem{
			ID:             docs[i].ID,
			DocumentType:   docs[i].DocumentType,
			FilePath:       docs[i].FilePath,
			PortofolioText: docs[i].PortofolioText,
		}
	}

	return &entity.CalonPesertaDetailResponse{
		ID:           app.ID,
		UserID:       app.UserID,
		UserName:     userName,
		UserEmail:    userEmail,
		JenisUjikom:  app.JenisUjikom,
		StatusKode:   app.StatusKode,
		CatatanTolak: app.CatatanTolak,
		AppliedAt:    appliedAt,
		Documents:    docItems,
	}, nil
}

func (r *ujikomRepo) UpdateApplicationStatus(ctx context.Context, id, statusKode, catatanTolak string) error {
	upd := map[string]interface{}{
		"status_kode":   statusKode,
		"catatan_tolak": catatanTolak,
		"updated_at":    time.Now().UTC(),
	}
	return r.db.WithContext(ctx).Model(&UjikomApplication{}).Where("id = ?", id).Updates(upd).Error
}
