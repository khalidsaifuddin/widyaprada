package dokumenpersyaratanrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *dokumenPersyaratanRepo) ListByJenisUjikom(ctx context.Context, jenisUjikom string) ([]entity.DokumenPersyaratanItem, error) {
	var rows []DokumenPersyaratanUjikom
	q := r.db.WithContext(ctx).Order("urutan ASC")
	if jenisUjikom != "" {
		q = q.Where("untuk_jenis_ujikom = ? OR untuk_jenis_ujikom IS NULL OR untuk_jenis_ujikom = ''", jenisUjikom)
	}
	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}

	// Filter: untuk jenis khusus, include only where untuk_jenis_ujikom matches or is null
	filtered := make([]entity.DokumenPersyaratanItem, 0, len(rows))
	for i := range rows {
		row := &rows[i]
		if jenisUjikom != "" && row.UntukJenisUjikom != "" && row.UntukJenisUjikom != jenisUjikom {
			continue
		}
		filtered = append(filtered, entity.DokumenPersyaratanItem{
			ID:               row.ID,
			Kode:             row.Kode,
			Nama:             row.Nama,
			Urutan:           row.Urutan,
			TipeInput:        row.TipeInput,
			Batasan:          row.Batasan,
			Deskripsi:        row.Deskripsi,
			UntukJenisUjikom: row.UntukJenisUjikom,
		})
	}
	return filtered, nil
}
