package entity

const (
	JurnalStatusDraft     = "Draft"
	JurnalStatusPublished = "Published"
)

// Jurnal untuk WPJurnal
type Jurnal struct {
	ID          string
	Title       string
	Author      string
	Abstract    string
	Content     string
	PublishedAt string
	Status      string
	Category    string
	Year        int
	UserID      *string // pemilik/submitter (untuk dashboard "my journals")
	CreatedAt   string
	UpdatedAt   string
}

// GetJurnalListRequest query GET /api/v1/jurnal
type GetJurnalListRequest struct {
	Q         string `form:"q"`
	Tahun     string `form:"tahun"`
	Kategori  string `form:"kategori"`
	Sort      string `form:"sort"`
	Page      int64  `form:"page"`
	PageSize  int64  `form:"page_size"`
	SortOrder string `form:"sort_order"`
}

// GetJurnalListResponse response
type GetJurnalListResponse struct {
	Items     []JurnalListItem `json:"items"`
	TotalPage int64            `json:"total_page"`
	TotalData int64            `json:"total_data"`
	Page      int64            `json:"page"`
	PageSize  int64            `json:"page_size"`
}

// JurnalListItem item
type JurnalListItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Abstract    string `json:"abstract"`
	PublishedAt string `json:"published_at,omitempty"`
	Year        int    `json:"year,omitempty"`
	Category    string `json:"category,omitempty"`
}

// JurnalDetailResponse GET /api/v1/jurnal/:id
type JurnalDetailResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Abstract    string `json:"abstract"`
	Content     string `json:"content"`
	PublishedAt string `json:"published_at,omitempty"`
	Status      string `json:"status"`
	Year        int    `json:"year,omitempty"`
	Category    string `json:"category,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}
