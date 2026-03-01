package entity

// Exam status constants
const (
	ExamStatusDraft       = "Draft"
	ExamStatusDiterbitkan = "Diterbitkan"
	ExamStatusBerlangsung = "Berlangsung"
	ExamStatusSelesai     = "Selesai"
)

// Exam content source type
const (
	ExamContentSourceQuestion = "question"
	ExamContentSourcePackage  = "package"
)

// ExamDetail domain model
type ExamDetail struct {
	ID                 string
	Code               string
	Name               string
	JadwalMulai        string
	JadwalSelesai      string
	DurasiMenit        int
	Status             string
	VerificationStatus string
	ShuffleQuestions   bool
	TampilkanLeaderboard bool
	DeletedReason      string
	CreatedAt          string
	UpdatedAt          string
}

// ExamContentItem konten ujian (soal atau paket)
type ExamContentItem struct {
	ExamID     string
	SourceType string // question | package
	SourceID   string
	SortOrder  int
}

// ExamParticipantItem peserta ujian
type ExamParticipantItem struct {
	ExamID  string
	UserID  string
}

// GetExamListRequest untuk GET /api/v1/exams
type GetExamListRequest struct {
	Q                string `form:"q"`
	Status           string `form:"status"`
	StatusVerifikasi string `form:"status_verifikasi"`
	Page             int64  `form:"page"`
	PageSize         int64  `form:"page_size"`
	SortBy           string `form:"sort_by"`
	SortOrder        string `form:"sort_order"`
}

// GetExamListResponse response list exams
type GetExamListResponse struct {
	Items     []ExamListItem `json:"items"`
	TotalPage int64          `json:"total_page"`
	TotalData int64          `json:"total_data"`
	Page      int64          `json:"page"`
	PageSize  int64          `json:"page_size"`
}

// ExamListItem item dalam list
type ExamListItem struct {
	ID                 string `json:"id"`
	Code               string `json:"code"`
	Name               string `json:"name"`
	JadwalMulai        string `json:"jadwal_mulai"`
	JadwalSelesai      string `json:"jadwal_selesai"`
	DurasiMenit        int    `json:"durasi_menit"`
	Status             string `json:"status"`
	VerificationStatus string `json:"verification_status"`
	ParticipantCount   int    `json:"participant_count"`
	CreatedAt          string `json:"created_at,omitempty"`
}

// ExamDetailResponse untuk GET /api/v1/exams/:id
type ExamDetailResponse struct {
	ID                   string              `json:"id"`
	Code                 string              `json:"code"`
	Name                 string              `json:"name"`
	JadwalMulai          string              `json:"jadwal_mulai"`
	JadwalSelesai        string              `json:"jadwal_selesai"`
	DurasiMenit          int                 `json:"durasi_menit"`
	Status               string              `json:"status"`
	VerificationStatus   string              `json:"verification_status"`
	ShuffleQuestions     bool                `json:"shuffle_questions"`
	TampilkanLeaderboard bool                `json:"tampilkan_leaderboard"`
	Contents             []ExamContentInfo   `json:"contents"`
	Participants         []ExamParticipantInfo `json:"participants"`
	CreatedAt            string              `json:"created_at,omitempty"`
	UpdatedAt            string              `json:"updated_at,omitempty"`
}

// ExamContentInfo konten dalam response
type ExamContentInfo struct {
	SourceType string `json:"source_type"`
	SourceID   string `json:"source_id"`
	SortOrder  int    `json:"sort_order"`
}

// ExamParticipantInfo peserta dalam response
type ExamParticipantInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name,omitempty"`
}

// ExamContentInput item untuk create/update (urutan preserved)
type ExamContentInput struct {
	SourceType string `json:"source_type"` // question | package
	SourceID   string `json:"source_id"`
}

// CreateExamRequest untuk POST /api/v1/exams
type CreateExamRequest struct {
	Code                 string              `json:"code" binding:"required"`
	Name                 string              `json:"name" binding:"required"`
	JadwalMulai          string              `json:"jadwal_mulai" binding:"required"`
	JadwalSelesai        string              `json:"jadwal_selesai" binding:"required"`
	DurasiMenit          int                 `json:"durasi_menit" binding:"required"`
	Contents             []ExamContentInput  `json:"contents"` // urutan: soal/paket dicampur
	QuestionIDs          []string            `json:"question_ids"` // fallback jika contents kosong
	PackageIDs           []string            `json:"package_ids"`
	ParticipantIDs       []string            `json:"participant_ids"`
	ShuffleQuestions     bool                `json:"shuffle_questions"`
	TampilkanLeaderboard bool                `json:"tampilkan_leaderboard"`
}

// UpdateExamRequest untuk PUT /api/v1/exams/:id
type UpdateExamRequest struct {
	Code                 string              `json:"code"`
	Name                 string              `json:"name"`
	JadwalMulai          string              `json:"jadwal_mulai"`
	JadwalSelesai        string              `json:"jadwal_selesai"`
	DurasiMenit          int                 `json:"durasi_menit"`
	Contents             []ExamContentInput  `json:"contents"`
	QuestionIDs          []string            `json:"question_ids"`
	PackageIDs           []string            `json:"package_ids"`
	ParticipantIDs       []string            `json:"participant_ids"`
	ShuffleQuestions     *bool               `json:"shuffle_questions"`
	TampilkanLeaderboard *bool               `json:"tampilkan_leaderboard"`
}

// DeleteExamRequest untuk DELETE /api/v1/exams/:id
type DeleteExamRequest struct {
	Reason string `json:"reason" binding:"required"`
}
