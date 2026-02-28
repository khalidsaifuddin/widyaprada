package entity

// QuestionCategory kategori soal (domain)
type QuestionCategory struct {
	ID   string
	Code string
	Name string
}

// QuestionDetail detail soal (domain, untuk repo)
type QuestionDetail struct {
	ID                 string
	Code               string
	Type               string
	CategoryID         string
	Difficulty         string
	QuestionText       string
	AnswerKey          string
	Weight             float64
	Status             string
	VerificationStatus string
	DeletedReason      string
	CreatedAt          string
	UpdatedAt          string
}

// Question type constants
const (
	QuestionTypePG         = "PG"         // Pilihan Ganda
	QuestionTypeBenarSalah = "BENAR_SALAH"
	QuestionTypeEssay      = "ESSAY"
)

// Question status
const (
	QuestionStatusDraft  = "Draft"
	QuestionStatusAktif  = "Aktif"
	QuestionVerifBelum   = "Belum"
	QuestionVerifSudah   = "Sudah"
)

// GetQuestionListRequest untuk GET /api/v1/questions
type GetQuestionListRequest struct {
	Q                  string `form:"q"`
	Tipe               string `form:"tipe"`
	KategoriID         string `form:"kategori_id"`
	Status             string `form:"status"`              // Draft, Aktif, all
	StatusVerifikasi   string `form:"status_verifikasi"`   // Belum, Sudah, all
	Page               int64  `form:"page"`
	PageSize           int64  `form:"page_size"`
	SortBy             string `form:"sort_by"`
	SortOrder          string `form:"sort_order"`          // asc, desc
}

// GetQuestionListResponse response list questions
type GetQuestionListResponse struct {
	Items     []QuestionListItem `json:"items"`
	TotalPage int64              `json:"total_page"`
	TotalData int64              `json:"total_data"`
	Page      int64              `json:"page"`
	PageSize  int64              `json:"page_size"`
}

// QuestionListItem item dalam list
type QuestionListItem struct {
	ID                 string  `json:"id"`
	Code               string  `json:"code"`
	Type               string  `json:"type"`
	CategoryID         string  `json:"category_id"`
	CategoryName       string  `json:"category_name,omitempty"`
	Difficulty         string  `json:"difficulty"`
	QuestionText       string  `json:"question_text"`
	AnswerKey          string  `json:"answer_key"`
	Weight             float64 `json:"weight"`
	Status             string  `json:"status"`
	VerificationStatus string  `json:"verification_status"`
	CreatedAt          string  `json:"created_at,omitempty"`
}

// QuestionDetailResponse untuk GET /api/v1/questions/:id
type QuestionDetailResponse struct {
	ID                 string           `json:"id"`
	Code               string           `json:"code"`
	Type               string           `json:"type"`
	CategoryID         string           `json:"category_id"`
	CategoryName       string           `json:"category_name,omitempty"`
	Difficulty         string           `json:"difficulty"`
	QuestionText       string           `json:"question_text"`
	AnswerKey          string           `json:"answer_key"`
	Weight             float64          `json:"weight"`
	Status             string           `json:"status"`
	VerificationStatus string           `json:"verification_status"`
	Options            []QuestionOption `json:"options"`
	CreatedAt          string           `json:"created_at,omitempty"`
	UpdatedAt          string           `json:"updated_at,omitempty"`
}

// QuestionOption opsi jawaban
type QuestionOption struct {
	ID          string `json:"id"`
	QuestionID  string `json:"question_id"`
	OptionKey   string `json:"option_key"`
	OptionText  string `json:"option_text"`
	IsCorrect   bool   `json:"is_correct"`
}

// CreateQuestionRequest untuk POST /api/v1/questions
type CreateQuestionRequest struct {
	Code         string           `json:"code" binding:"required"`
	Type         string           `json:"type" binding:"required"`
	CategoryID   string           `json:"category_id"`
	Difficulty   string           `json:"difficulty"`
	QuestionText string           `json:"question_text" binding:"required"`
	AnswerKey    string           `json:"answer_key"`
	Weight       float64          `json:"weight"`
	Status       string           `json:"status"`
	Options      []QuestionOptionInput `json:"options"`
}

// QuestionOptionInput input opsi saat create/update
type QuestionOptionInput struct {
	OptionKey  string `json:"option_key"`
	OptionText string `json:"option_text"`
	IsCorrect  bool   `json:"is_correct"`
}

// UpdateQuestionRequest untuk PUT /api/v1/questions/:id
type UpdateQuestionRequest struct {
	Code         string                `json:"code"`
	Type         string                `json:"type"`
	CategoryID   string                `json:"category_id"`
	Difficulty   string                `json:"difficulty"`
	QuestionText string                `json:"question_text"`
	AnswerKey    string                `json:"answer_key"`
	Weight       float64               `json:"weight"`
	Status       string                `json:"status"`
	Options      []QuestionOptionInput `json:"options"`
}

// DeleteQuestionRequest untuk DELETE /api/v1/questions/:id
type DeleteQuestionRequest struct {
	Reason string `json:"reason" binding:"required"`
}
