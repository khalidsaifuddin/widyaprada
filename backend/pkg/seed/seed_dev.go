package seed

import (
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/pkg/auth"
	articlerepo "github.com/ProjectWidyaprada/backend/repository/article-repo"
	examrepo "github.com/ProjectWidyaprada/backend/repository/exam-repo"
	journalrepo "github.com/ProjectWidyaprada/backend/repository/journal-repo"
	linkrepo "github.com/ProjectWidyaprada/backend/repository/link-repo"
	packagerepo "github.com/ProjectWidyaprada/backend/repository/package-repo"
	questionrepo "github.com/ProjectWidyaprada/backend/repository/question-repo"
	sliderepo "github.com/ProjectWidyaprada/backend/repository/slide-repo"
	userrepo "github.com/ProjectWidyaprada/backend/repository/user-repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DevPassword password untuk semua user dummy (development only)
const DevPassword = "Password123!"

// SeedDevData seeds dummy data untuk development: users (login), berita, bank soal, paket, ujian, slide, jurnal, tautan.
// Idempotent: skip jika data sudah ada. Jalankan hanya saat ENVIRONMENT=development.
func SeedDevData(db *gorm.DB, environment string) error {
	if strings.ToLower(environment) != "development" {
		return nil
	}
	if err := seedDevUsers(db); err != nil {
		return err
	}
	if err := seedDevArticles(db); err != nil {
		return err
	}
	if err := seedDevQuestions(db); err != nil {
		return err
	}
	if err := seedDevPackages(db); err != nil {
		return err
	}
	if err := seedDevExams(db); err != nil {
		return err
	}
	if err := seedDevSlides(db); err != nil {
		return err
	}
	if err := seedDevJournals(db); err != nil {
		return err
	}
	if err := seedDevLinks(db); err != nil {
		return err
	}
	return nil
}

func seedDevUsers(db *gorm.DB) error {
	hash, err := auth.HashPassword(DevPassword)
	if err != nil {
		return err
	}
	now := time.Now().UTC()

	// Cari role IDs
	var superAdminRole, pesertaRole userrepo.Role
	if err := db.Where("UPPER(code) = ?", "SUPER_ADMIN").First(&superAdminRole).Error; err != nil {
		return err
	}
	if err := db.Where("UPPER(code) = ?", "PESERTA").First(&pesertaRole).Error; err != nil {
		return err
	}

	users := []struct {
		email    string
		name     string
		username string
		roleID   string
	}{
		{"admin@example.com", "Admin Demo", "admin", superAdminRole.ID},
		{"peserta@example.com", "Peserta Demo", "peserta", pesertaRole.ID},
	}

	for _, u := range users {
		var count int64
		db.Model(&userrepo.User{}).Where("LOWER(email) = ?", strings.ToLower(u.email)).Count(&count)
		if count > 0 {
			continue
		}
		user := userrepo.User{
			ID:           uuid.New().String(),
			Name:         u.name,
			Email:        u.email,
			Username:     u.username,
			PasswordHash: hash,
			IsActive:     true,
			CreatedAt:    &now,
			UpdatedAt:    &now,
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}
		if err := db.Create(&userrepo.UserRole{UserID: user.ID, RoleID: u.roleID, CreatedAt: &now}).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedDevArticles(db *gorm.DB) error {
	var count int64
	db.Model(&articlerepo.Article{}).Count(&count)
	if count > 0 {
		return nil
	}
	now := time.Now().UTC()
	pub := now.Add(-24 * time.Hour)

	// URL gambar placeholder untuk thumbnail & galeri
	img1 := "https://placehold.co/800x450/1e3a5f/ffffff?text=Pengumuman+Uji+Kompetensi"
	img2 := "https://placehold.co/800x450/022c55/ffffff?text=Workshop+PMP"
	img3 := "https://placehold.co/800x450/033259/ffffff?text=Pedoman+CBT"
	img4 := "https://placehold.co/800x450/044a7d/ffffff?text=Galeri+Berita"

	articles := []articlerepo.Article{
		{
			ID:           uuid.New().String(),
			Title:        "Pengumuman Uji Kompetensi Widyaprada 2025",
			Slug:         "pengumuman-uji-kompetensi-2025",
			Content:      "Pendaftaran uji kompetensi jabatan fungsional Widyaprada dibuka mulai 1 Maret 2025. Silakan daftar melalui portal ini.",
			Excerpt:      "Pendaftaran uji kompetensi dibuka.",
			ThumbnailURL: img1,
			GalleryURLs:  []string{img1, img2, img3},
			PublishedAt:  &pub,
			Status:       "Published",
			AuthorName:   "Tim Widyaprada",
			Category:     "Pengumuman",
			CreatedAt:    &now,
			UpdatedAt:    &now,
		},
		{
			ID:           uuid.New().String(),
			Title:        "Workshop Penjaminan Mutu Pendidikan Dasar",
			Slug:         "workshop-penjaminan-mutu-2025",
			Content:      "Workshop penjaminan mutu pendidikan dasar akan diselenggarakan pada bulan April 2025. Peserta diharapkan mendaftar sebelum tanggal 15 Maret.",
			Excerpt:      "Workshop penjaminan mutu pendidikan.",
			ThumbnailURL: img2,
			GalleryURLs:  []string{img2, img3, img4, img1},
			PublishedAt:  &pub,
			Status:       "Published",
			AuthorName:   "Admin",
			Category:     "Kegiatan",
			CreatedAt:    &now,
			UpdatedAt:    &now,
		},
		{
			ID:           uuid.New().String(),
			Title:        "Pedoman Pelaksanaan CBT Widyaprada",
			Slug:         "pedoman-cbt-widyaprada",
			Content:      "Berikut pedoman pelaksanaan Computer Based Test (CBT) untuk uji kompetensi Widyaprada. Pastikan perangkat dan koneksi internet memadai.",
			Excerpt:      "Pedoman CBT Widyaprada.",
			ThumbnailURL: img3,
			GalleryURLs:  []string{img3, img4},
			PublishedAt:  &pub,
			Status:       "Published",
			AuthorName:   "Tim Ujikom",
			Category:     "Dokumen",
			CreatedAt:    &now,
			UpdatedAt:    &now,
		},
	}
	return db.Create(&articles).Error
}

func seedDevQuestions(db *gorm.DB) error {
	var count int64
	db.Model(&questionrepo.Question{}).Count(&count)
	if count > 0 {
		return nil
	}
	var cat questionrepo.QuestionCategory
	if err := db.Where("UPPER(code) = ?", "UMUM").First(&cat).Error; err != nil {
		return err
	}
	now := time.Now().UTC()

	q1 := questionrepo.Question{
		ID:                 uuid.New().String(),
		Code:               "UMUM-001",
		Type:               "PG",
		CategoryID:         cat.ID,
		Difficulty:         "mudah",
		QuestionText:       "Apa kepanjangan dari PMP?",
		AnswerKey:          "A",
		Weight:             1,
		Status:             "Aktif",
		VerificationStatus: "Sudah",
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}
	if err := db.Create(&q1).Error; err != nil {
		return err
	}
	db.Create([]questionrepo.QuestionOption{
		{ID: uuid.New().String(), QuestionID: q1.ID, OptionKey: "A", OptionText: "Penjaminan Mutu Pendidikan", IsCorrect: true, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q1.ID, OptionKey: "B", OptionText: "Pendidikan Mutu Peserta", IsCorrect: false, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q1.ID, OptionKey: "C", OptionText: "Pemantauan Mutu Pendidikan", IsCorrect: false, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q1.ID, OptionKey: "D", OptionText: "Penilaian Mutu Pendidikan", IsCorrect: false, CreatedAt: &now},
	})

	q2 := questionrepo.Question{
		ID:                 uuid.New().String(),
		Code:               "UMUM-002",
		Type:               "PG",
		CategoryID:         cat.ID,
		Difficulty:         "sedang",
		QuestionText:       "Siapa yang bertugas sebagai widyaprada?",
		AnswerKey:          "B",
		Weight:             1,
		Status:             "Aktif",
		VerificationStatus: "Sudah",
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}
	if err := db.Create(&q2).Error; err != nil {
		return err
	}
	db.Create([]questionrepo.QuestionOption{
		{ID: uuid.New().String(), QuestionID: q2.ID, OptionKey: "A", OptionText: "Tenaga administratif sekolah", IsCorrect: false, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q2.ID, OptionKey: "B", OptionText: "Pejabat fungsional di bidang penjaminan mutu pendidikan", IsCorrect: true, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q2.ID, OptionKey: "C", OptionText: "Guru pengajar", IsCorrect: false, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q2.ID, OptionKey: "D", OptionText: "Kepala sekolah", IsCorrect: false, CreatedAt: &now},
	})

	q3 := questionrepo.Question{
		ID:                 uuid.New().String(),
		Code:               "UMUM-003",
		Type:               "PG",
		CategoryID:         cat.ID,
		Difficulty:         "mudah",
		QuestionText:       "Standar Nasional Pendidikan mencakup berapa standar?",
		AnswerKey:          "C",
		Weight:             1,
		Status:             "Aktif",
		VerificationStatus: "Sudah",
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}
	if err := db.Create(&q3).Error; err != nil {
		return err
	}
	db.Create([]questionrepo.QuestionOption{
		{ID: uuid.New().String(), QuestionID: q3.ID, OptionKey: "A", OptionText: "6 standar", IsCorrect: false, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q3.ID, OptionKey: "B", OptionText: "7 standar", IsCorrect: false, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q3.ID, OptionKey: "C", OptionText: "8 standar", IsCorrect: true, CreatedAt: &now},
		{ID: uuid.New().String(), QuestionID: q3.ID, OptionKey: "D", OptionText: "9 standar", IsCorrect: false, CreatedAt: &now},
	})

	return nil
}

func seedDevPackages(db *gorm.DB) error {
	var count int64
	db.Model(&packagerepo.QuestionPackage{}).Count(&count)
	if count > 0 {
		return nil
	}
	var questions []questionrepo.Question
	if err := db.Where("status = ?", "Aktif").Limit(3).Find(&questions).Error; err != nil || len(questions) < 2 {
		return nil
	}
	now := time.Now().UTC()
	pkg := packagerepo.QuestionPackage{
		ID:                 uuid.New().String(),
		Code:               "PAKET-001",
		Name:               "Paket Soal Umum Widyaprada",
		Description:        "Paket soal untuk uji kompetensi dasar widyaprada.",
		Status:             "Aktif",
		VerificationStatus: "Sudah",
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}
	if err := db.Create(&pkg).Error; err != nil {
		return err
	}
	for i, q := range questions {
		db.Create(&packagerepo.PackageQuestionItem{PackageID: pkg.ID, QuestionID: q.ID, SortOrder: i + 1, CreatedAt: &now})
	}
	return nil
}

func seedDevExams(db *gorm.DB) error {
	var count int64
	db.Model(&examrepo.Exam{}).Count(&count)
	if count > 0 {
		return nil
	}
	var pkg packagerepo.QuestionPackage
	if err := db.Where("code = ?", "PAKET-001").First(&pkg).Error; err != nil {
		return nil
	}
	now := time.Now().UTC()
	start := now.Add(24 * time.Hour)
	end := now.Add(7 * 24 * time.Hour)
	exam := examrepo.Exam{
		ID:                   uuid.New().String(),
		Code:                 "UJIKOM-001",
		Name:                 "Uji Kompetensi Widyaprada Dasar 2025",
		JadwalMulai:          &start,
		JadwalSelesai:        &end,
		DurasiMenit:          60,
		Status:               "Diterbitkan",
		VerificationStatus:   "Sudah",
		ShuffleQuestions:     true,
		TampilkanLeaderboard: true,
		CreatedAt:            &now,
		UpdatedAt:            &now,
	}
	if err := db.Create(&exam).Error; err != nil {
		return err
	}
	db.Create(&examrepo.ExamContent{ExamID: exam.ID, SourceType: "package", SourceID: pkg.ID, SortOrder: 1, CreatedAt: &now})

	var peserta userrepo.User
	if err := db.Where("LOWER(email) = ?", "peserta@example.com").First(&peserta).Error; err == nil {
		db.Create(&examrepo.ExamParticipant{ExamID: exam.ID, UserID: peserta.ID, CreatedAt: &now})
	}
	return nil
}

func seedDevSlides(db *gorm.DB) error {
	var count int64
	db.Model(&sliderepo.Slide{}).Count(&count)
	if count > 0 {
		return nil
	}
	now := time.Now().UTC()
	slides := []sliderepo.Slide{
		{ID: uuid.New().String(), ImageURL: "https://placehold.co/1200x400/e2e8f0/64748b?text=Selamat+Datang+Portal+Widyaprada", Title: "Selamat Datang", Subtitle: "Portal Uji Kompetensi Widyaprada", LinkURL: "/beranda", CTALabel: "Mulai", SortOrder: 1, Status: "Published", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), ImageURL: "https://placehold.co/1200x400/cbd5e1/475569?text=Uji+Kompetensi+2025", Title: "Uji Kompetensi 2025", Subtitle: "Daftar sekarang untuk uji kompetensi jabatan fungsional Widyaprada", LinkURL: "/wpujikom/assignment", CTALabel: "Daftar", SortOrder: 2, Status: "Published", CreatedAt: &now, UpdatedAt: &now},
	}
	return db.Create(&slides).Error
}

func seedDevJournals(db *gorm.DB) error {
	var count int64
	db.Model(&journalrepo.Journal{}).Count(&count)
	if count > 0 {
		return nil
	}
	now := time.Now().UTC()
	pub := now.Add(-48 * time.Hour)
	journals := []journalrepo.Journal{
		{ID: uuid.New().String(), Title: "Implementasi Penjaminan Mutu di Sekolah Dasar", Author: "Dr. Ahmad Sugiarto", Abstract: "Studi mengenai implementasi sistem penjaminan mutu pendidikan di sekolah dasar.", Content: "", PublishedAt: &pub, Status: "Published", Category: "Pendidikan Dasar", Year: 2025, CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Title: "Peran Widyaprada dalam Peningkatan Mutu", Author: "Prof. Siti Rahayu", Abstract: "Membahas peran pejabat fungsional widyaprada dalam peningkatan mutu pendidikan.", Content: "", PublishedAt: &pub, Status: "Published", Category: "Penjaminan Mutu", Year: 2025, CreatedAt: &now, UpdatedAt: &now},
	}
	return db.Create(&journals).Error
}

func seedDevLinks(db *gorm.DB) error {
	var count int64
	db.Model(&linkrepo.Link{}).Count(&count)
	if count > 0 {
		return nil
	}
	now := time.Now().UTC()
	links := []linkrepo.Link{
		{ID: uuid.New().String(), Title: "Ditjen PAUD Dikdas Dikmen", URL: "https://pauddikdasmen.kemdikbud.go.id", Description: "Situs resmi Ditjen PAUD Dikdas Dikmen", SortOrder: 1, Status: "Aktif", OpenInNewTab: true, CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Title: "Pusat Asesmen Pendidikan", URL: "https://pusmenjar.kemdikbud.go.id", Description: "Pusat Asesmen dan Pembelajaran", SortOrder: 2, Status: "Aktif", OpenInNewTab: true, CreatedAt: &now, UpdatedAt: &now},
	}
	return db.Create(&links).Error
}
