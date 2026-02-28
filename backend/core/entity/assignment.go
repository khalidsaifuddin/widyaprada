package entity

import "errors"

var (
	ErrAssignmentForbidden   = errors.New("akses ditolak")
	ErrLeaderboardPrivate    = errors.New("leaderboard bersifat privat")
	ErrApplyAlreadyExists    = errors.New("sudah pernah apply untuk jenis ujikom ini")
	ErrApplyNotFound         = errors.New("pendaftaran tidak ditemukan")
	ErrDokumenPersyaratanReq = errors.New("dokumen persyaratan wajib belum lengkap")
)

// Jenis Ujikom
const (
	JenisUjikomPerpindahanJabatan = "perpindahan_jabatan"
	JenisUjikomKenaikanTingkat    = "kenaikan_tingkat"
)

// Status pendaftaran
const (
	ApplyStatusMenungguVerifikasi = "ujikom_menunggu_verifikasi"
	ApplyStatusLolos              = "ujikom_lolos"
	ApplyStatusTidakLolos         = "ujikom_tidak_lolos"
)

// Status pengerjaan assignment
const (
	AssignmentStatusBelumDikerjakan = "belum_dikerjakan"
	AssignmentStatusSudahDikerjakan = "sudah_dikerjakan"
)

// DokumenPersyaratanItem item dokumen persyaratan (ref)
type DokumenPersyaratanItem struct {
	ID              string
	Kode            string
	Nama            string
	Urutan          int
	TipeInput       string // file | text_portofolio | text_essay
	Batasan         string
	Deskripsi       string
	UntukJenisUjikom string
}

// ListDokumenPersyaratanResponse response GET /api/v1/ujikom/dokumen-persyaratan
type ListDokumenPersyaratanResponse struct {
	Items []DokumenPersyaratanItem `json:"items"`
}

// ApplyUjikomRequest untuk POST /api/v1/ujikom/apply (multipart/JSON)
type ApplyUjikomRequest struct {
	JenisUjikom string `form:"jenis_ujikom" json:"jenis_ujikom" binding:"required"`
	// Documents passed separately as ApplyUjikomDocumentInput slice
}

// ApplyUjikomDocumentInput dokumen dalam apply
type ApplyUjikomDocumentInput struct {
	DocumentType   string `json:"document_type"`
	FilePath       string `json:"file_path"`
	PortofolioText string `json:"portofolio_text"`
}

// ApplyStatusResponse response GET /api/v1/ujikom/apply/status
type ApplyStatusResponse struct {
	Status       string `json:"status"`        // menunggu_verifikasi | lolos | tidak_lolos
	StatusKode   string `json:"status_kode"`   // ujikom_menunggu_verifikasi, etc
	CatatanTolak string `json:"catatan_tolak,omitempty"`
	JenisUjikom  string `json:"jenis_ujikom"`
	AppliedAt    string `json:"applied_at,omitempty"`
}

// AssignmentListItem item Tugas Saya
type AssignmentListItem struct {
	ExamID            string  `json:"exam_id"`
	ExamName          string  `json:"exam_name"`
	Deadline          string  `json:"deadline"` // jadwal_selesai
	Status            string  `json:"status"`   // belum_dikerjakan | sudah_dikerjakan
	Score             *float64 `json:"score,omitempty"`
	CanViewLeaderboard bool   `json:"can_view_leaderboard"`
}

// GetAssignmentListRequest query GET /api/v1/assignments
type GetAssignmentListRequest struct {
	Status   string `form:"status"`
	SortBy   string `form:"sort_by"`
	SortOrder string `form:"sort_order"`
	Page     int64  `form:"page"`
	PageSize int64  `form:"page_size"`
}

// GetAssignmentListResponse response list assignments
type GetAssignmentListResponse struct {
	Items     []AssignmentListItem `json:"items"`
	TotalPage int64                `json:"total_page"`
	TotalData int64                `json:"total_data"`
	Page      int64                `json:"page"`
	PageSize  int64                `json:"page_size"`
}

// AssignmentResultResponse response GET /api/v1/assignments/:examId/result
type AssignmentResultResponse struct {
	ExamID     string   `json:"exam_id"`
	ExamName   string   `json:"exam_name"`
	Score      *float64 `json:"score,omitempty"`
	SubmittedAt string  `json:"submitted_at,omitempty"`
}

// LeaderboardItem item leaderboard
type LeaderboardItem struct {
	Rank   int     `json:"rank"`
	UserID string  `json:"user_id"`
	Name   string  `json:"name"`
	Score  float64 `json:"score"`
}

// GetLeaderboardResponse response GET /api/v1/assignments/:examId/leaderboard
type GetLeaderboardResponse struct {
	ExamID   string            `json:"exam_id"`
	ExamName string            `json:"exam_name"`
	Items    []LeaderboardItem `json:"items"`
}
