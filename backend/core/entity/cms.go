package entity

// Status CMS
const (
	SlideStatusDraft     = "Draft"
	SlideStatusPublished = "Published"
	ArticleStatusDraft   = "Draft"
	ArticleStatusPublished = "Published"
	LinkStatusAktif      = "Aktif"
	LinkStatusNonaktif   = "Nonaktif"
)

// Slide for CMS Slider (JSON tags for API response)
type Slide struct {
	ID        string  `json:"id"`
	ImageURL  string  `json:"image_url"`
	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	LinkURL   string  `json:"link_url"`
	CTALabel  string  `json:"cta_label"`
	SortOrder int     `json:"sort_order"`
	Status    string  `json:"status"`
	DateStart string  `json:"date_start,omitempty"`
	DateEnd   string  `json:"date_end,omitempty"`
	SatkerID  *string `json:"satker_id,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}

// GetSlideListRequest query
type GetSlideListRequest struct {
	Status    string `form:"status"`
	SatkerID  string `form:"satker_id"`
	Page      int64  `form:"page"`
	PageSize  int64  `form:"page_size"`
	SortBy    string `form:"sort_by"`
	SortOrder string `form:"sort_order"`
}

// GetSlideListResponse response
type GetSlideListResponse struct {
	Items     []SlideListItem `json:"items"`
	TotalPage int64           `json:"total_page"`
	TotalData int64           `json:"total_data"`
	Page      int64           `json:"page"`
	PageSize  int64           `json:"page_size"`
}

// SlideListItem item
type SlideListItem struct {
	ID        string  `json:"id"`
	ImageURL  string  `json:"image_url"`
	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	LinkURL   string  `json:"link_url"`
	CTALabel  string  `json:"cta_label"`
	SortOrder int     `json:"sort_order"`
	Status    string  `json:"status"`
	DateStart string  `json:"date_start,omitempty"`
	DateEnd   string  `json:"date_end,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
}

// CreateSlideRequest body
type CreateSlideRequest struct {
	ImageURL   string  `json:"image_url" binding:"required"`
	Title      string  `json:"title"`
	Subtitle   string  `json:"subtitle"`
	LinkURL    string  `json:"link_url"`
	CTALabel   string  `json:"cta_label"`
	SortOrder  int     `json:"sort_order"`
	Status     string  `json:"status"`
	DateStart  string  `json:"tanggal_mulai_tampil"`
	DateEnd    string  `json:"tanggal_selesai_tampil"`
	SatkerID   *string `json:"satker_id"`
}

// UpdateSlideRequest body
type UpdateSlideRequest struct {
	ImageURL  string  `json:"image_url"`
	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	LinkURL   string  `json:"link_url"`
	CTALabel  string  `json:"cta_label"`
	SortOrder *int    `json:"sort_order"`
	Status    string  `json:"status"`
	DateStart string  `json:"tanggal_mulai_tampil"`
	DateEnd   string  `json:"tanggal_selesai_tampil"`
}

// Article for CMS Berita (JSON tags for API response)
type Article struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Slug         string   `json:"slug"`
	Content      string   `json:"content"`
	Excerpt      string   `json:"excerpt"`
	ThumbnailURL string   `json:"thumbnail_url,omitempty"`
	GalleryURLs  []string `json:"gallery_urls,omitempty"`
	PublishedAt  string   `json:"published_at,omitempty"`
	Status       string   `json:"status"`
	AuthorName   string   `json:"author_name,omitempty"`
	Category     string   `json:"category,omitempty"`
	SatkerID     *string  `json:"satker_id,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	UpdatedAt    string   `json:"updated_at,omitempty"`
}

// GetArticleListRequest query
type GetArticleListRequest struct {
	Q          string `form:"q"`
	Category   string `form:"kategori"`
	Status     string `form:"status"`
	SatkerID   string `form:"satker_id"`
	Page       int64  `form:"page"`
	PageSize   int64  `form:"page_size"`
	SortBy     string `form:"sort"`
	SortOrder  string `form:"sort_order"`
}

// GetArticleListResponse response
type GetArticleListResponse struct {
	Items     []ArticleListItem `json:"items"`
	TotalPage int64             `json:"total_page"`
	TotalData int64             `json:"total_data"`
	Page      int64             `json:"page"`
	PageSize  int64             `json:"page_size"`
}

// ArticleListItem item
type ArticleListItem struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Excerpt      string `json:"excerpt"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	PublishedAt  string `json:"published_at,omitempty"`
	Status       string `json:"status"`
	AuthorName   string `json:"author_name,omitempty"`
	Category     string `json:"category,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

// CreateArticleRequest body
type CreateArticleRequest struct {
	Title       string   `json:"judul" binding:"required"`
	Slug        string   `json:"slug"`
	Content     string   `json:"konten"`
	Excerpt     string   `json:"ringkasan"`
	Thumbnail   string   `json:"thumbnail"`
	GalleryURLs []string `json:"gallery_urls"`
	PublishedAt string   `json:"tanggal_publikasi"`
	Status      string   `json:"status"`
	AuthorName  string   `json:"penulis"`
	Category    string   `json:"kategori"`
	SatkerID    *string  `json:"satker_id"`
}

// ArticleDetailResponse untuk GET /api/v1/berita/:slug (publik)
type ArticleDetailResponse struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Slug         string   `json:"slug"`
	Content      string   `json:"content"`
	Excerpt      string   `json:"excerpt"`
	ThumbnailURL string   `json:"thumbnail_url,omitempty"`
	GalleryURLs  []string `json:"gallery_urls,omitempty"`
	PublishedAt  string   `json:"published_at,omitempty"`
	AuthorName   string   `json:"author_name,omitempty"`
	Category     string   `json:"category,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
}

// GetBeritaListResponse response GET /api/v1/berita (publik)
type GetBeritaListResponse struct {
	Items     []ArticlePublicItem `json:"items"`
	TotalPage int64               `json:"total_page"`
	TotalData int64               `json:"total_data"`
	Page      int64               `json:"page"`
	PageSize  int64               `json:"page_size"`
}

// UpdateArticleRequest body
type UpdateArticleRequest struct {
	Title       string   `json:"judul"`
	Slug        string   `json:"slug"`
	Content     string   `json:"konten"`
	Excerpt     string   `json:"ringkasan"`
	Thumbnail   string   `json:"thumbnail"`
	GalleryURLs []string `json:"gallery_urls"`
	PublishedAt string   `json:"tanggal_publikasi"`
	Status      string   `json:"status"`
	AuthorName  string   `json:"penulis"`
	Category    string   `json:"kategori"`
}

// Link for CMS Tautan (JSON tags for API response)
type Link struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	URL          string  `json:"url"`
	Description  string  `json:"description"`
	SortOrder    int     `json:"sort_order"`
	Status       string  `json:"status"`
	OpenInNewTab bool    `json:"buka_di_tab_baru"`
	SatkerID     *string `json:"satker_id,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

// GetLinkListRequest query
type GetLinkListRequest struct {
	Status   string `form:"status"`
	SatkerID string `form:"satker_id"`
	Page     int64  `form:"page"`
	PageSize int64  `form:"page_size"`
	SortBy   string `form:"sort_by"`
	SortOrder string `form:"sort_order"`
}

// GetLinkListResponse response
type GetLinkListResponse struct {
	Items     []LinkListItem `json:"items"`
	TotalPage int64          `json:"total_page"`
	TotalData int64          `json:"total_data"`
	Page      int64          `json:"page"`
	PageSize  int64          `json:"page_size"`
}

// LinkListItem item
type LinkListItem struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Description  string `json:"description"`
	SortOrder    int    `json:"sort_order"`
	Status       string `json:"status"`
	OpenInNewTab bool   `json:"buka_di_tab_baru"`
	CreatedAt    string `json:"created_at,omitempty"`
}

// CreateLinkRequest body
type CreateLinkRequest struct {
	Title        string  `json:"judul" binding:"required"`
	URL          string  `json:"url" binding:"required"`
	Description  string  `json:"deskripsi"`
	SortOrder    int     `json:"urutan"`
	Status       string  `json:"status"`
	OpenInNewTab bool    `json:"buka_di_tab_baru"`
	SatkerID     *string `json:"satker_id"`
}

// UpdateLinkRequest body
type UpdateLinkRequest struct {
	Title        string `json:"judul"`
	URL          string `json:"url"`
	Description  string `json:"deskripsi"`
	SortOrder    *int   `json:"urutan"`
	Status       string `json:"status"`
	OpenInNewTab *bool  `json:"buka_di_tab_baru"`
}
