package dokumenpersyaratanrepo

import (
	"time"

	"gorm.io/gorm"
)

// DokumenPersyaratanUjikom DTO (ref table)
type DokumenPersyaratanUjikom struct {
	ID               string         `gorm:"column:id;primaryKey;type:uuid"`
	Kode             string         `gorm:"column:kode;size:50;not null;index"`
	Nama             string         `gorm:"column:nama;type:text;not null"`
	Urutan           int            `gorm:"column:urutan;not null;default:0"`
	TipeInput        string         `gorm:"column:tipe_input;size:20;not null;default:file"`
	Batasan          string         `gorm:"column:batasan;type:text"`
	Deskripsi        string         `gorm:"column:deskripsi;type:text"`
	UntukJenisUjikom string         `gorm:"column:untuk_jenis_ujikom;size:50;index"`
	CreatedAt        *time.Time     `gorm:"column:created_at"`
	UpdatedAt        *time.Time     `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (DokumenPersyaratanUjikom) TableName() string {
	return "dokumen_persyaratan_ujikom"
}
