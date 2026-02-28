package entity

// PackageDetail domain model paket soal
type PackageDetail struct {
	ID                 string
	Code               string
	Name               string
	Description        string
	Status             string
	VerificationStatus string
	DeletedReason      string
	CreatedAt          string
	UpdatedAt          string
}

// PackageQuestionItem item soal dalam paket
type PackageQuestionItem struct {
	PackageID  string
	QuestionID string
	SortOrder  int
}

// GetPackageListRequest untuk GET /api/v1/question-packages
type GetPackageListRequest struct {
	Q                string `form:"q"`
	Status           string `form:"status"`
	StatusVerifikasi string `form:"status_verifikasi"`
	Page             int64  `form:"page"`
	PageSize         int64  `form:"page_size"`
	SortBy           string `form:"sort_by"`
	SortOrder        string `form:"sort_order"`
}

// GetPackageListResponse response list packages
type GetPackageListResponse struct {
	Items     []PackageListItem `json:"items"`
	TotalPage int64             `json:"total_page"`
	TotalData int64             `json:"total_data"`
	Page      int64             `json:"page"`
	PageSize  int64             `json:"page_size"`
}

// PackageListItem item dalam list
type PackageListItem struct {
	ID                 string `json:"id"`
	Code               string `json:"code"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Status             string `json:"status"`
	VerificationStatus string `json:"verification_status"`
	QuestionCount      int    `json:"question_count"`
	CreatedAt          string `json:"created_at,omitempty"`
}

// PackageDetailResponse untuk GET /api/v1/question-packages/:id
type PackageDetailResponse struct {
	ID                 string                 `json:"id"`
	Code               string                 `json:"code"`
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	Status             string                 `json:"status"`
	VerificationStatus string                 `json:"verification_status"`
	Questions          []PackageQuestionInfo  `json:"questions"`
	CreatedAt          string                 `json:"created_at,omitempty"`
	UpdatedAt          string                 `json:"updated_at,omitempty"`
}

// PackageQuestionInfo soal dalam paket (detail)
type PackageQuestionInfo struct {
	QuestionID   string `json:"question_id"`
	QuestionCode string `json:"question_code"`
	SortOrder    int    `json:"sort_order"`
}

// CreatePackageRequest untuk POST /api/v1/question-packages
type CreatePackageRequest struct {
	Code        string   `json:"code" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	QuestionIDs []string `json:"question_ids" binding:"required"`
}

// UpdatePackageRequest untuk PUT /api/v1/question-packages/:id
type UpdatePackageRequest struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	QuestionIDs []string `json:"question_ids"`
}

// DeletePackageRequest untuk DELETE /api/v1/question-packages/:id
type DeletePackageRequest struct {
	Reason string `json:"reason" binding:"required"`
}
