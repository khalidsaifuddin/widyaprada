package ujikomrepo

import (
	"time"

	"gorm.io/gorm"
)

// UjikomApplication DTO
type UjikomApplication struct {
	ID           string         `gorm:"column:id;primaryKey;type:uuid"`
	UserID       string         `gorm:"column:user_id;type:uuid;not null;index"`
	JenisUjikom  string         `gorm:"column:jenis_ujikom;size:50;not null;index"`
	StatusKode   string         `gorm:"column:status_kode;size:50;not null;index"`
	CatatanTolak string         `gorm:"column:catatan_tolak"`
	CreatedAt    *time.Time     `gorm:"column:created_at"`
	UpdatedAt    *time.Time     `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (UjikomApplication) TableName() string {
	return "ujikom_application"
}

// UjikomApplicationDocument DTO
type UjikomApplicationDocument struct {
	ID                 string         `gorm:"column:id;primaryKey;type:uuid"`
	UjikomApplicationID string        `gorm:"column:ujikom_application_id;type:uuid;not null;index"`
	DocumentType       string         `gorm:"column:document_type;size:100;not null;index"`
	FilePath           string         `gorm:"column:file_path;size:500"`
	PortofolioText     string         `gorm:"column:portofolio_text;type:text"`
	CreatedAt          *time.Time     `gorm:"column:created_at"`
	UpdatedAt          *time.Time     `gorm:"column:updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (UjikomApplicationDocument) TableName() string {
	return "ujikom_application_document"
}
