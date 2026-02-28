package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// DokumenPersyaratanRepo repository referensi dokumen persyaratan
type DokumenPersyaratanRepo interface {
	// ListByJenisUjikom daftar dokumen persyaratan (untuk_jenis_ujikom = jenis OR null)
	ListByJenisUjikom(ctx context.Context, jenisUjikom string) ([]entity.DokumenPersyaratanItem, error)
}
