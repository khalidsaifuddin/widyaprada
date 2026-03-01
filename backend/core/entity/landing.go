package entity

// LandingHomeResponse GET /api/v1/landing/home
type LandingHomeResponse struct {
	Slider  []SlidePublicItem    `json:"slider"`
	Berita  []ArticlePublicItem  `json:"berita"`
	Tautan  []LinkPublicItem     `json:"tautan"`
	Jurnal  []JurnalPublicItem   `json:"jurnal"`
}

// SlidePublicItem slide untuk publik (Published, dalam range tanggal)
type SlidePublicItem struct {
	ID        string `json:"id"`
	ImageURL  string `json:"image_url"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	LinkURL   string `json:"link_url"`
	CTALabel  string `json:"cta_label"`
	SortOrder int    `json:"sort_order"`
}

// ArticlePublicItem berita untuk publik
type ArticlePublicItem struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Slug         string   `json:"slug"`
	Excerpt      string   `json:"excerpt"`
	ThumbnailURL string   `json:"thumbnail_url,omitempty"`
	GalleryURLs  []string `json:"gallery_urls,omitempty"`
	PublishedAt  string   `json:"published_at,omitempty"`
}

// LinkPublicItem tautan untuk publik (status Aktif)
type LinkPublicItem struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Description  string `json:"description"`
	SortOrder    int    `json:"sort_order"`
	OpenInNewTab bool   `json:"buka_di_tab_baru"`
}

// JurnalPublicItem jurnal untuk publik (Published)
type JurnalPublicItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Abstract    string `json:"abstract"`
	PublishedAt string `json:"published_at,omitempty"`
}

// BerandaPengumumanResponse GET /api/v1/beranda/pengumuman (auth Peserta)
type BerandaPengumumanResponse struct {
	HasilSeleksiAdmin  *ApplyStatusResponse `json:"hasil_seleksi_administrasi,omitempty"`
	InfoJadwalUjikom   []JadwalUjikomItem  `json:"info_jadwal_ujikom"`
	CanStartUjikom     bool                 `json:"can_start_ujikom"`
	ExamsTersedia      []CBTExamTersediaItem `json:"exams_tersedia,omitempty"`
}

// JadwalUjikomItem info jadwal
type JadwalUjikomItem struct {
	ExamID       string `json:"exam_id"`
	ExamName     string `json:"exam_name"`
	JadwalMulai  string `json:"jadwal_mulai"`
	JadwalSelesai string `json:"jadwal_selesai"`
}
