package wpdatarepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// WidyapradaData DTO
type WidyapradaData struct {
	ID                       string         `gorm:"column:id;primaryKey;type:uuid"`
	NIP                      string         `gorm:"column:nip;size:50;uniqueIndex;not null"`
	NamaLengkap              string         `gorm:"column:nama_lengkap;size:255;not null"`
	JenisKelamin             string         `gorm:"column:jenis_kelamin;size:20"`
	GolonganRuang            string         `gorm:"column:golongan_ruang;size:50"`
	Pangkat                  string         `gorm:"column:pangkat;size:100"`
	JenjangJabatanFungsional string         `gorm:"column:jenjang_jabatan_fungsional;size:100"`
	SatkerID                 string         `gorm:"column:satker_id;type:uuid;not null;index"`
	UnitKerja                string         `gorm:"column:unit_kerja;size:255"`
	PendidikanTerakhir       string         `gorm:"column:pendidikan_terakhir;size:100"`
	TMTGolongan              string         `gorm:"column:tmt_golongan;size:50"`
	TMTJabatanFungsional     string         `gorm:"column:tmt_jabatan_fungsional;size:50"`
	NoSK                     string         `gorm:"column:no_sk_pengangkatan;size:100"`
	NoHP                     string         `gorm:"column:no_hp;size:50"`
	Email                    string         `gorm:"column:email;size:255"`
	Alamat                   string         `gorm:"column:alamat;type:text"`
	Status                   string         `gorm:"column:status;size:20;default:Aktif"`
	Keterangan               string         `gorm:"column:keterangan;type:text"`
	UserID                   *string        `gorm:"column:user_id;type:uuid"`
	DeletedReason            string         `gorm:"column:deleted_reason"`
	CreatedAt                *time.Time     `gorm:"column:created_at"`
	UpdatedAt                *time.Time     `gorm:"column:updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (WidyapradaData) TableName() string {
	return "widyaprada_data"
}

func (w *WidyapradaData) ToListItem() entity.WPDataListItem {
	createdAt := ""
	if w.CreatedAt != nil {
		createdAt = w.CreatedAt.UTC().Format(time.RFC3339)
	}
	return entity.WPDataListItem{
		ID:                 w.ID,
		NIP:                w.NIP,
		NamaLengkap:        w.NamaLengkap,
		JenisKelamin:       w.JenisKelamin,
		GolonganRuang:      w.GolonganRuang,
		Pangkat:            w.Pangkat,
		SatkerID:           w.SatkerID,
		UnitKerja:          w.UnitKerja,
		Status:             w.Status,
		PendidikanTerakhir: w.PendidikanTerakhir,
		CreatedAt:          createdAt,
	}
}

func (w *WidyapradaData) ToDetail() entity.WPDataDetailResponse {
	createdAt, updatedAt := "", ""
	if w.CreatedAt != nil {
		createdAt = w.CreatedAt.UTC().Format(time.RFC3339)
	}
	if w.UpdatedAt != nil {
		updatedAt = w.UpdatedAt.UTC().Format(time.RFC3339)
	}
	return entity.WPDataDetailResponse{
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
		CreatedAt:                createdAt,
		UpdatedAt:                updatedAt,
	}
}
