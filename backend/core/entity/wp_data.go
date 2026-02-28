package entity

const (
	WPDataStatusAktif    = "Aktif"
	WPDataStatusNonaktif = "Nonaktif"
)

// WidyapradaData data WP
type WidyapradaData struct {
	ID                        string
	NIP                       string
	NamaLengkap               string
	JenisKelamin              string
	GolonganRuang             string
	Pangkat                   string
	JenjangJabatanFungsional  string
	SatkerID                  string
	UnitKerja                 string
	PendidikanTerakhir        string
	TMTGolongan               string
	TMTJabatanFungsional      string
	NoSK                      string
	NoHP                      string
	Email                     string
	Alamat                    string
	Status                    string
	Keterangan                string
	UserID                    *string
	CreatedAt                 string
	UpdatedAt                 string
	DeletedAt                 string
	DeletedReason             string
}

// GetWPDataListRequest query
type GetWPDataListRequest struct {
	Q          string `form:"q"`
	SatkerID   string `form:"satker_id"`
	UnitKerja  string `form:"unit_kerja"`
	Status     string `form:"status"`
	Page       int64  `form:"page"`
	PageSize   int64  `form:"page_size"`
	SortBy     string `form:"sort_by"`
	SortOrder  string `form:"sort_order"`
}

// GetWPDataListResponse response
type GetWPDataListResponse struct {
	Items     []WPDataListItem `json:"items"`
	TotalPage int64            `json:"total_page"`
	TotalData int64            `json:"total_data"`
	Page      int64            `json:"page"`
	PageSize  int64            `json:"page_size"`
}

// WPDataListItem item
type WPDataListItem struct {
	ID               string  `json:"id"`
	NIP              string  `json:"nip"`
	NamaLengkap      string  `json:"nama_lengkap"`
	JenisKelamin     string  `json:"jenis_kelamin,omitempty"`
	GolonganRuang    string  `json:"golongan_ruang,omitempty"`
	Pangkat          string  `json:"pangkat,omitempty"`
	SatkerID         string  `json:"satker_id"`
	UnitKerja        string  `json:"unit_kerja,omitempty"`
	Status           string  `json:"status"`
	PendidikanTerakhir string `json:"pendidikan_terakhir,omitempty"`
	CreatedAt        string  `json:"created_at,omitempty"`
}

// WPDataDetailResponse detail
type WPDataDetailResponse struct {
	ID                       string  `json:"id"`
	NIP                      string  `json:"nip"`
	NamaLengkap              string  `json:"nama_lengkap"`
	JenisKelamin             string  `json:"jenis_kelamin,omitempty"`
	GolonganRuang            string  `json:"golongan_ruang,omitempty"`
	Pangkat                  string  `json:"pangkat,omitempty"`
	JenjangJabatanFungsional string  `json:"jenjang_jabatan_fungsional,omitempty"`
	SatkerID                 string  `json:"satker_id"`
	UnitKerja                string  `json:"unit_kerja,omitempty"`
	PendidikanTerakhir       string  `json:"pendidikan_terakhir,omitempty"`
	TMTGolongan              string  `json:"tmt_golongan,omitempty"`
	TMTJabatanFungsional     string  `json:"tmt_jabatan_fungsional,omitempty"`
	NoSK                     string  `json:"no_sk,omitempty"`
	NoHP                     string  `json:"no_hp,omitempty"`
	Email                    string  `json:"email,omitempty"`
	Alamat                   string  `json:"alamat,omitempty"`
	Status                   string  `json:"status"`
	Keterangan               string  `json:"keterangan,omitempty"`
	CreatedAt                string  `json:"created_at,omitempty"`
	UpdatedAt                string  `json:"updated_at,omitempty"`
}

// CreateWPDataRequest body
type CreateWPDataRequest struct {
	NIP                      string `json:"nip" binding:"required"`
	NamaLengkap              string `json:"nama_lengkap" binding:"required"`
	JenisKelamin             string `json:"jenis_kelamin"`
	GolonganRuang            string `json:"golongan_ruang"`
	Pangkat                  string `json:"pangkat"`
	JenjangJabatanFungsional string `json:"jenjang_jabatan_fungsional"`
	SatkerID                 string `json:"satker_id" binding:"required"`
	UnitKerja                string `json:"unit_kerja"`
	PendidikanTerakhir       string `json:"pendidikan_terakhir"`
	TMTGolongan              string `json:"tmt_golongan"`
	TMTJabatanFungsional     string `json:"tmt_jabatan_fungsional"`
	NoSK                     string `json:"no_sk"`
	NoHP                     string `json:"no_hp"`
	Email                    string `json:"email"`
	Alamat                   string `json:"alamat"`
	Status                   string `json:"status"`
	Keterangan               string `json:"keterangan"`
}

// UpdateWPDataRequest body
type UpdateWPDataRequest struct {
	NIP                      string `json:"nip"`
	NamaLengkap              string `json:"nama_lengkap"`
	JenisKelamin             string `json:"jenis_kelamin"`
	GolonganRuang            string `json:"golongan_ruang"`
	Pangkat                  string `json:"pangkat"`
	JenjangJabatanFungsional string `json:"jenjang_jabatan_fungsional"`
	SatkerID                 string `json:"satker_id"`
	UnitKerja                string `json:"unit_kerja"`
	PendidikanTerakhir       string `json:"pendidikan_terakhir"`
	TMTGolongan              string `json:"tmt_golongan"`
	TMTJabatanFungsional     string `json:"tmt_jabatan_fungsional"`
	NoSK                     string `json:"no_sk"`
	NoHP                     string `json:"no_hp"`
	Email                    string `json:"email"`
	Alamat                   string `json:"alamat"`
	Status                   string `json:"status"`
	Keterangan               string `json:"keterangan"`
}

// GetCalonPesertaListRequest query calon peserta
type GetCalonPesertaListRequest struct {
	Q                 string `form:"q"`
	StatusVerifikasi  string `form:"status_verifikasi"`
	Page              int64  `form:"page"`
	PageSize          int64  `form:"page_size"`
}

// GetCalonPesertaListResponse response list calon peserta
type GetCalonPesertaListResponse struct {
	Items     []CalonPesertaListItem `json:"items"`
	TotalPage int64                  `json:"total_page"`
	TotalData int64                  `json:"total_data"`
	Page      int64                  `json:"page"`
	PageSize  int64                  `json:"page_size"`
}

// CalonPesertaListItem item
type CalonPesertaListItem struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	UserName        string `json:"user_name"`
	UserEmail       string `json:"user_email"`
	JenisUjikom     string `json:"jenis_ujikom"`
	StatusKode      string `json:"status_kode"`
	CatatanTolak    string `json:"catatan_tolak,omitempty"`
	AppliedAt       string `json:"applied_at"`
}

// CalonPesertaDetailResponse detail calon peserta + dokumen
type CalonPesertaDetailResponse struct {
	ID              string                     `json:"id"`
	UserID          string                     `json:"user_id"`
	UserName        string                     `json:"user_name"`
	UserEmail       string                     `json:"user_email"`
	JenisUjikom     string                     `json:"jenis_ujikom"`
	StatusKode      string                     `json:"status_kode"`
	CatatanTolak    string                     `json:"catatan_tolak,omitempty"`
	AppliedAt       string                     `json:"applied_at"`
	Documents       []CalonPesertaDocumentItem `json:"documents"`
}

// CalonPesertaDocumentItem dokumen
type CalonPesertaDocumentItem struct {
	ID           string `json:"id"`
	DocumentType string `json:"document_type"`
	FilePath     string `json:"file_path,omitempty"`
	PortofolioText string `json:"portofolio_text,omitempty"`
}

// VerifyCalonPesertaRequest body
type VerifyCalonPesertaRequest struct {
	Approved bool   `json:"approved" binding:"required"`
	Catatan  string `json:"catatan"`
}
