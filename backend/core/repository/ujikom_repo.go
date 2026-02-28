package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// UjikomRepo repository untuk apply pendaftaran ujikom
type UjikomRepo interface {
	// CreateApplication membuat pendaftaran baru
	CreateApplication(ctx context.Context, userID, jenisUjikom, statusID string, docs []entity.ApplyUjikomDocumentInput) error
	// GetLatestByUser jenis ujikom -> application terbaru (untuk status)
	GetLatestApplicationByUser(ctx context.Context, userID string) (*entity.ApplyStatusResponse, error)
	// HasPendingOrApproved cek apakah user sudah apply (menunggu/lolos) untuk jenis ini
	HasPendingOrApproved(ctx context.Context, userID, jenisUjikom string) (bool, error)

	// Calon peserta (SDD_Manajemen_Data_WP)
	ListCalonPeserta(ctx context.Context, req entity.GetCalonPesertaListRequest, satkerFilter *string) (*entity.GetCalonPesertaListResponse, error)
	GetCalonPesertaByID(ctx context.Context, id string) (*entity.CalonPesertaDetailResponse, error)
	UpdateApplicationStatus(ctx context.Context, id, statusKode, catatanTolak string) error
}
