package middleware

import (
	"net/http"
	"strings"

	"github.com/ProjectWidyaprada/backend/config"
	assignment_api "github.com/ProjectWidyaprada/backend/handler/api/assignment"
	auth_api "github.com/ProjectWidyaprada/backend/handler/api/auth"
	beranda_api "github.com/ProjectWidyaprada/backend/handler/api/beranda"
	berita_api "github.com/ProjectWidyaprada/backend/handler/api/berita"
	cbt_api "github.com/ProjectWidyaprada/backend/handler/api/cbt"
	cms_api "github.com/ProjectWidyaprada/backend/handler/api/cms"
	dashboard_api "github.com/ProjectWidyaprada/backend/handler/api/dashboard"
	example_api "github.com/ProjectWidyaprada/backend/handler/api/example"
	exam_api "github.com/ProjectWidyaprada/backend/handler/api/exam"
	landing_api "github.com/ProjectWidyaprada/backend/handler/api/landing"
	jurnal_api "github.com/ProjectWidyaprada/backend/handler/api/jurnal"
	question_api "github.com/ProjectWidyaprada/backend/handler/api/question"
	upload_api "github.com/ProjectWidyaprada/backend/handler/api/upload"
	questionpackage_api "github.com/ProjectWidyaprada/backend/handler/api/questionpackage"
	rbac_api "github.com/ProjectWidyaprada/backend/handler/api/rbac"
	user_api "github.com/ProjectWidyaprada/backend/handler/api/user"
	wpdata_api "github.com/ProjectWidyaprada/backend/handler/api/wpdata"
	assignment_usecase "github.com/ProjectWidyaprada/backend/core/usecase/assignment"
	auth_usecase "github.com/ProjectWidyaprada/backend/core/usecase/auth"
	beranda_usecase "github.com/ProjectWidyaprada/backend/core/usecase/beranda"
	berita_usecase "github.com/ProjectWidyaprada/backend/core/usecase/berita"
	banksoal_usecase "github.com/ProjectWidyaprada/backend/core/usecase/banksoal"
	cbt_usecase "github.com/ProjectWidyaprada/backend/core/usecase/cbt"
	cms_usecase "github.com/ProjectWidyaprada/backend/core/usecase/cms"
	dashboard_usecase "github.com/ProjectWidyaprada/backend/core/usecase/dashboard"
	example_usecase "github.com/ProjectWidyaprada/backend/core/usecase/example"
	exam_usecase "github.com/ProjectWidyaprada/backend/core/usecase/exam"
	forgotpassword_usecase "github.com/ProjectWidyaprada/backend/core/usecase/forgotpassword"
	jurnal_usecase "github.com/ProjectWidyaprada/backend/core/usecase/jurnal"
	landing_usecase "github.com/ProjectWidyaprada/backend/core/usecase/landing"
	paketsoal_usecase "github.com/ProjectWidyaprada/backend/core/usecase/paketsoal"
	rbac_usecase "github.com/ProjectWidyaprada/backend/core/usecase/rbac"
	registration_usecase "github.com/ProjectWidyaprada/backend/core/usecase/registration"
	usermanagement_usecase "github.com/ProjectWidyaprada/backend/core/usecase/usermanagement"
	wpdata_usecase "github.com/ProjectWidyaprada/backend/core/usecase/wpdata"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/email"
	"github.com/ProjectWidyaprada/backend/pkg/seed"
	assignmentrepo "github.com/ProjectWidyaprada/backend/repository/assignment-repo"
	articlerepo "github.com/ProjectWidyaprada/backend/repository/article-repo"
	cbtrepo "github.com/ProjectWidyaprada/backend/repository/cbt-repo"
	dokumenpersyaratanrepo "github.com/ProjectWidyaprada/backend/repository/dokumen-persyaratan-repo"
	examplerepo "github.com/ProjectWidyaprada/backend/repository/example-repo"
	examrepo "github.com/ProjectWidyaprada/backend/repository/exam-repo"
	journalrepo "github.com/ProjectWidyaprada/backend/repository/journal-repo"
	linkrepo "github.com/ProjectWidyaprada/backend/repository/link-repo"
	packagerepo "github.com/ProjectWidyaprada/backend/repository/package-repo"
	passwordresettokenrepo "github.com/ProjectWidyaprada/backend/repository/password-reset-token-repo"
	questionrepo "github.com/ProjectWidyaprada/backend/repository/question-repo"
	rbacrepo "github.com/ProjectWidyaprada/backend/repository/rbac-repo"
	sliderepo "github.com/ProjectWidyaprada/backend/repository/slide-repo"
	ujikomrepo "github.com/ProjectWidyaprada/backend/repository/ujikom-repo"
	userrepo "github.com/ProjectWidyaprada/backend/repository/user-repo"
	wpdatarepo "github.com/ProjectWidyaprada/backend/repository/wpdata-repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (*gin.Engine, interface{}) {
	if strings.EqualFold(cfg.Environment, "production") {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(InputValidationMiddleware())
	if cfg.IsEnableSentry {
		router.Use(SentryRecoveryMiddleware())
		router.Use(SentryMiddleware())
	}
	router.MaxMultipartMemory = 16 << 20
	router.Use(func(c *gin.Context) {
		maxBytes := int64(1 << 20)
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			maxBytes = 16 << 20
		}
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		c.Next()
	})

	// Auto-migrate: for SQLite always; for Postgres also when development (untuk seed dummy data)
	if cfg.DBType != "postgres" || strings.EqualFold(cfg.Environment, "development") {
		_ = db.AutoMigrate(&examplerepo.Example{})
		_ = db.AutoMigrate(&questionrepo.QuestionCategory{}, &questionrepo.Question{}, &questionrepo.QuestionOption{})
		_ = db.AutoMigrate(&packagerepo.QuestionPackage{}, &packagerepo.PackageQuestionItem{})
		_ = db.AutoMigrate(&examrepo.Exam{}, &examrepo.ExamContent{}, &examrepo.ExamParticipant{}, &examrepo.ExamAttempt{}, &examrepo.ExamAnswer{}, &examrepo.ExamAttemptQuestion{})
		_ = db.AutoMigrate(&ujikomrepo.UjikomApplication{}, &ujikomrepo.UjikomApplicationDocument{})
		_ = db.AutoMigrate(&dokumenpersyaratanrepo.DokumenPersyaratanUjikom{})
		_ = db.AutoMigrate(&userrepo.User{}, &userrepo.Role{}, &userrepo.UserRole{})
		_ = db.AutoMigrate(&sliderepo.Slide{})
		_ = db.AutoMigrate(&articlerepo.Article{})
		_ = db.AutoMigrate(&linkrepo.Link{})
		_ = db.AutoMigrate(&journalrepo.Journal{})
		_ = db.AutoMigrate(&wpdatarepo.WidyapradaData{})
		_ = db.AutoMigrate(&rbacrepo.Permission{}, &rbacrepo.RolePermission{})
		_ = db.AutoMigrate(&passwordresettokenrepo.PasswordResetToken{})
	}
	// Seeds (idempotent) for both SQLite and Postgres
	_ = dokumenpersyaratanrepo.SeedDokumenPersyaratan(db)
	_ = questionrepo.SeedDefaultCategories(db)
	_ = userrepo.SeedDefaultRoles(db)
	_ = seed.SeedDevData(db, cfg.Environment)

	// Auth feature (SDD Auth Login + Registrasi + Logout + Lupa Password)
	userRepo := userrepo.NewUserRepo(db)
	tokenBlacklist := auth.NewMemoryBlacklist()
	authUsecase := auth_usecase.NewAuthUsecase(userRepo, tokenBlacklist, cfg)
	emailService := email.NewLogEmailService()
	registrationUsecase := registration_usecase.NewRegistrationUsecase(userRepo, emailService)
	passwordResetTokenRepo := passwordresettokenrepo.NewPasswordResetTokenRepo(db)
	forgotPasswordUsecase := forgotpassword_usecase.NewForgotPasswordUsecase(userRepo, passwordResetTokenRepo, emailService, cfg)
	authHandler := auth_api.NewAuthHTTPHandler(authUsecase, registrationUsecase, forgotPasswordUsecase)

	// User Management (SDD_Auth_Manajemen_Pengguna)
	userUsecase := usermanagement_usecase.NewUserUsecase(userRepo)
	userHandler := user_api.NewUserHTTPHandler(userUsecase, userRepo)

	// RBAC (SDD_RBAC) - Super Admin only
	rbacRepo := rbacrepo.NewRBACRepo(db)
	roleUsecase := rbac_usecase.NewRoleUsecase(rbacRepo)
	permissionUsecase := rbac_usecase.NewPermissionUsecase(rbacRepo)
	rbacHandler := rbac_api.NewRBACHTTPHandler(roleUsecase, permissionUsecase)

	// Example feature (hexagon: repo -> usecase -> handler)
	exampleRepo := examplerepo.NewExampleRepo(db)
	exampleUsecase := example_usecase.NewExampleUsecase(exampleRepo)
	exampleHandler := example_api.NewExampleHTTPHandler(exampleUsecase)

	// Bank Soal (SDD_Bank_Soal)
	questionRepo := questionrepo.NewQuestionRepo(db)
	bankSoalUsecase := banksoal_usecase.NewBankSoalUsecase(questionRepo)
	questionHandler := question_api.NewQuestionHTTPHandler(bankSoalUsecase)

	packageRepo := packagerepo.NewPackageRepo(db)
	paketSoalUsecase := paketsoal_usecase.NewPaketSoalUsecase(packageRepo, questionRepo)
	packageHandler := questionpackage_api.NewQuestionPackageHTTPHandler(paketSoalUsecase)

	examRepo := examrepo.NewExamRepo(db)
	examUsecase := exam_usecase.NewExamUsecase(examRepo, userRepo)
	examHandler := exam_api.NewExamHTTPHandler(examUsecase)

	assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
	dokumenRepo := dokumenpersyaratanrepo.NewDokumenPersyaratanRepo(db)
	ujikomRepo := ujikomrepo.NewUjikomRepo(db)
	assignmentUsecase := assignment_usecase.NewAssignmentUsecase(dokumenRepo, ujikomRepo, assignmentRepo)
	assignmentHandler := assignment_api.NewAssignmentHTTPHandler(assignmentUsecase)

	cbtRepo := cbtrepo.NewCBTRepo(db)
	cbtUsecase := cbt_usecase.NewCBTUsecase(cbtRepo, assignmentRepo)
	cbtHandler := cbt_api.NewCBTHTTPHandler(cbtUsecase)

	// Landing, Berita, Jurnal, Beranda, CMS, WP Data
	slideRepo := sliderepo.NewSlideRepo(db)
	articleRepo := articlerepo.NewArticleRepo(db)
	linkRepo := linkrepo.NewLinkRepo(db)
	journalRepo := journalrepo.NewJournalRepo(db)

	landingUsecase := landing_usecase.NewLandingUsecase(slideRepo, articleRepo, linkRepo, journalRepo)
	landingHandler := landing_api.NewLandingHTTPHandler(landingUsecase)

	berandaUsecase := beranda_usecase.NewBerandaUsecase(ujikomRepo, cbtRepo)
	berandaHandler := beranda_api.NewBerandaHTTPHandler(berandaUsecase)

	beritaUsecase := berita_usecase.NewBeritaUsecase(articleRepo)
	beritaHandler := berita_api.NewBeritaHTTPHandler(beritaUsecase)

	jurnalUsecase := jurnal_usecase.NewJurnalUsecase(journalRepo)
	jurnalHandler := jurnal_api.NewJurnalHTTPHandler(jurnalUsecase)

	sliderUsecase := cms_usecase.NewSliderUsecase(slideRepo)
	beritaCMSUsecase := cms_usecase.NewBeritaCMSUsecase(articleRepo)
	tautanUsecase := cms_usecase.NewTautanUsecase(linkRepo)
	cmsHandler := cms_api.NewCMSHTTPHandler(sliderUsecase, beritaCMSUsecase, tautanUsecase, userRepo)

	wpdataRepo := wpdatarepo.NewWPDataRepo(db)
	wpdataUsecase := wpdata_usecase.NewWPDataUsecase(wpdataRepo)
	calonPesertaUsecase := wpdata_usecase.NewCalonPesertaUsecase(ujikomRepo)
	wpdataHandler := wpdata_api.NewWPDataHTTPHandler(wpdataUsecase, calonPesertaUsecase, userRepo)

	dashboardUsecase := dashboard_usecase.NewDashboardUsecase(assignmentRepo, journalRepo)
	dashboardHandler := dashboard_api.NewDashboardHTTPHandler(dashboardUsecase)

	apiGroup := router.Group("/api")
	v1Group := apiGroup.Group("/v1")

	authGroup := v1Group.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/logout", AuthRequired(cfg, tokenBlacklist), authHandler.Logout)
	authGroup.POST("/forgot-password", authHandler.ForgotPassword)
	authGroup.POST("/reset-password", authHandler.ResetPassword)

	examplesGroup := v1Group.Group("/examples")
	examplesGroup.GET("", exampleHandler.GetExampleList)
	examplesGroup.GET("/:id", exampleHandler.GetExampleDetail)

	// User Management - requires auth
	usersGroup := v1Group.Group("/users")
	usersGroup.Use(AuthRequired(cfg, tokenBlacklist))
	usersGroup.GET("", userHandler.GetUserList)
	usersGroup.GET("/:id", userHandler.GetUserDetail)
	usersGroup.POST("", userHandler.CreateUser)
	usersGroup.PUT("/:id", userHandler.UpdateUser)
	usersGroup.DELETE("/:id", userHandler.DeleteUser)

	// RBAC - requires auth + Super Admin
	rbacGroup := v1Group.Group("/rbac")
	rbacGroup.Use(AuthRequired(cfg, tokenBlacklist))
	rbacGroup.Use(RequireSuperAdmin())
	rbacGroup.GET("/roles", rbacHandler.GetRoleList)
	rbacGroup.GET("/roles/:id", rbacHandler.GetRoleDetail)
	rbacGroup.POST("/roles", rbacHandler.CreateRole)
	rbacGroup.PUT("/roles/:id", rbacHandler.UpdateRole)
	rbacGroup.DELETE("/roles/:id", rbacHandler.DeleteRole)
	rbacGroup.GET("/permissions", rbacHandler.GetPermissionList)
	rbacGroup.GET("/permissions/:id", rbacHandler.GetPermissionDetail)
	rbacGroup.POST("/permissions", rbacHandler.CreatePermission)
	rbacGroup.PUT("/permissions/:id", rbacHandler.UpdatePermission)
	rbacGroup.DELETE("/permissions/:id", rbacHandler.DeletePermission)

	// Questions (Bank Soal) - GET: Admin Uji Kompetensi or Super Admin; POST/PUT/DELETE: Super Admin; Verify/Unverify: Verifikator or Super Admin
	questionsGroup := v1Group.Group("/questions")
	questionsGroup.Use(AuthRequired(cfg, tokenBlacklist))
	questionsGroup.GET("", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), questionHandler.GetQuestionList)
	questionsGroup.GET("/categories", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), questionHandler.GetCategories)
	questionsGroup.GET("/:id", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), questionHandler.GetQuestionDetail)
	questionsGroup.POST("", RequireSuperAdmin(), questionHandler.CreateQuestion)
	questionsGroup.PUT("/:id", RequireSuperAdmin(), questionHandler.UpdateQuestion)
	questionsGroup.DELETE("/:id", RequireSuperAdmin(), questionHandler.DeleteQuestion)
	questionsGroup.POST("/:id/verify", RequireAnyRole("SUPER_ADMIN", "VERIFIKATOR"), questionHandler.VerifyQuestion)
	questionsGroup.POST("/:id/unverify", RequireAnyRole("SUPER_ADMIN", "VERIFIKATOR"), questionHandler.UnverifyQuestion)

	// Question Packages (Paket Soal) - CRUD: Admin Uji Kompetensi or Super Admin; Verify/Unverify: Verifikator or Super Admin
	packagesGroup := v1Group.Group("/question-packages")
	packagesGroup.Use(AuthRequired(cfg, tokenBlacklist))
	packagesGroup.GET("", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), packageHandler.GetPackageList)
	packagesGroup.GET("/:id", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), packageHandler.GetPackageDetail)
	packagesGroup.POST("", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), packageHandler.CreatePackage)
	packagesGroup.PUT("/:id", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), packageHandler.UpdatePackage)
	packagesGroup.DELETE("/:id", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), packageHandler.DeletePackage)
	packagesGroup.POST("/:id/verify", RequireAnyRole("SUPER_ADMIN", "VERIFIKATOR"), packageHandler.VerifyPackage)
	packagesGroup.POST("/:id/unverify", RequireAnyRole("SUPER_ADMIN", "VERIFIKATOR"), packageHandler.UnverifyPackage)

	// Exams (Manajemen Uji Kompetensi) - CRUD, Publish: Admin Uji Kompetensi or Super Admin; Verify/Unverify: Verifikator or Super Admin
	examsGroup := v1Group.Group("/exams")
	examsGroup.Use(AuthRequired(cfg, tokenBlacklist))
	examsGroup.GET("", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), examHandler.GetExamList)
	examsGroup.GET("/:id", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), examHandler.GetExamDetail)
	examsGroup.POST("", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), examHandler.CreateExam)
	examsGroup.PUT("/:id", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), examHandler.UpdateExam)
	examsGroup.DELETE("/:id", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), examHandler.DeleteExam)
	examsGroup.POST("/:id/publish", RequireAnyRole("SUPER_ADMIN", "ADMIN_UJIKOM"), examHandler.PublishExam)
	examsGroup.POST("/:id/verify", RequireAnyRole("SUPER_ADMIN", "VERIFIKATOR"), examHandler.VerifyExam)
	examsGroup.POST("/:id/unverify", RequireAnyRole("SUPER_ADMIN", "VERIFIKATOR"), examHandler.UnverifyExam)

	// Ujikom (Apply, Dokumen Persyaratan) - PESERTA
	ujikomGroup := v1Group.Group("/ujikom")
	ujikomGroup.Use(AuthRequired(cfg, tokenBlacklist))
	ujikomGroup.GET("/dokumen-persyaratan", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), assignmentHandler.GetDokumenPersyaratan)
	ujikomGroup.POST("/apply", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), assignmentHandler.ApplyUjikom)
	ujikomGroup.GET("/apply/status", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), assignmentHandler.GetApplyStatus)

	// Assignments (Tugas Saya, Hasil, Leaderboard) - PESERTA
	assignmentsGroup := v1Group.Group("/assignments")
	assignmentsGroup.Use(AuthRequired(cfg, tokenBlacklist))
	assignmentsGroup.GET("", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), assignmentHandler.GetAssignmentList)
	assignmentsGroup.GET("/:examId/result", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), assignmentHandler.GetAssignmentResult)
	assignmentsGroup.GET("/:examId/leaderboard", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), assignmentHandler.GetLeaderboard)

	// CBT (Computer-Based Test) - PESERTA
	cbtGroup := v1Group.Group("/cbt")
	cbtGroup.Use(AuthRequired(cfg, tokenBlacklist))
	cbtGroup.GET("/exams", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), cbtHandler.GetCBTExams)
	cbtGroup.POST("/exams/:id/start", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), cbtHandler.StartCBTExam)
	cbtGroup.GET("/attempts/:attemptId/questions", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), cbtHandler.GetCBTQuestions)
	cbtGroup.GET("/attempts/:attemptId/questions/:num", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), cbtHandler.GetCBTQuestionByNum)
	cbtGroup.POST("/attempts/:attemptId/answers", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), cbtHandler.SaveCBTAnswer)
	cbtGroup.POST("/attempts/:attemptId/submit", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), cbtHandler.SubmitCBTExam)
	cbtGroup.GET("/history", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), cbtHandler.GetCBTHistory)

	// Dashboard Widyaprada
	dashboardGroup := v1Group.Group("/dashboard")
	dashboardGroup.Use(AuthRequired(cfg, tokenBlacklist))
	dashboardGroup.Use(RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM", "ADMIN_SATKER"))
	dashboardGroup.GET("/assignments", dashboardHandler.GetAssignments)
	dashboardGroup.GET("/journals", dashboardHandler.GetMyJournals)

	// Landing (publik)
	landingGroup := v1Group.Group("/landing")
	landingGroup.GET("/home", landingHandler.GetHome)

	// Beranda pengumuman (auth Peserta)
	berandaGroup := v1Group.Group("/beranda")
	berandaGroup.Use(AuthRequired(cfg, tokenBlacklist))
	berandaGroup.GET("/pengumuman", RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM"), berandaHandler.GetPengumuman)

	// Berita (publik)
	beritaGroup := v1Group.Group("/berita")
	beritaGroup.GET("", beritaHandler.GetList)
	beritaGroup.GET("/:slug", beritaHandler.GetBySlug)

	// Jurnal (publik)
	jurnalGroup := v1Group.Group("/jurnal")
	jurnalGroup.GET("", jurnalHandler.GetList)
	jurnalGroup.GET("/:id", jurnalHandler.GetByID)

	// WPJurnal (manajemen jurnal - auth)
	wpjurnalGroup := v1Group.Group("/wpjurnal")
	wpjurnalGroup.Use(AuthRequired(cfg, tokenBlacklist))
	wpjurnalGroup.Use(RequireAnyRole("PESERTA", "SUPER_ADMIN", "ADMIN_UJIKOM", "ADMIN_SATKER"))
	uploadHandler := upload_api.NewUploadHTTPHandler()
	wpjurnalGroup.GET("/:id", jurnalHandler.GetByIDForOwner)
	wpjurnalGroup.POST("/upload-pdf", uploadHandler.UploadJournalPDF)
	wpjurnalGroup.POST("", jurnalHandler.Create)
	wpjurnalGroup.PUT("/:id", jurnalHandler.Update)

	// CMS Slider, Berita, Tautan (auth: Admin Satker or Super Admin)
	cmsGroup := v1Group.Group("/cms")
	cmsGroup.Use(AuthRequired(cfg, tokenBlacklist))
	cmsGroup.Use(RequireAnyRole("SUPER_ADMIN", "ADMIN_SATKER"))
	cmsGroup.POST("/upload-image", uploadHandler.UploadImage)
	sliderGroup := cmsGroup.Group("/slider")
	sliderGroup.GET("", cmsHandler.GetSlideList)
	sliderGroup.POST("", cmsHandler.CreateSlide)
	sliderGroup.GET("/:id", cmsHandler.GetSlideDetail)
	sliderGroup.PUT("/:id", cmsHandler.UpdateSlide)
	sliderGroup.DELETE("/:id", cmsHandler.DeleteSlide)
	beritaCMSGroup := cmsGroup.Group("/berita")
	beritaCMSGroup.GET("", cmsHandler.GetArticleList)
	beritaCMSGroup.POST("", cmsHandler.CreateArticle)
	beritaCMSGroup.GET("/:id", cmsHandler.GetArticleDetail)
	beritaCMSGroup.PUT("/:id", cmsHandler.UpdateArticle)
	beritaCMSGroup.DELETE("/:id", cmsHandler.DeleteArticle)
	tautanGroup := cmsGroup.Group("/tautan")
	tautanGroup.GET("", cmsHandler.GetLinkList)
	tautanGroup.POST("", cmsHandler.CreateLink)
	tautanGroup.GET("/:id", cmsHandler.GetLinkDetail)
	tautanGroup.PUT("/:id", cmsHandler.UpdateLink)
	tautanGroup.DELETE("/:id", cmsHandler.DeleteLink)

	// WP Data & Calon Peserta (auth: Admin Satker or Super Admin)
	wpdataGroup := v1Group.Group("/wp-data")
	wpdataGroup.Use(AuthRequired(cfg, tokenBlacklist))
	wpdataGroup.Use(RequireAnyRole("SUPER_ADMIN", "ADMIN_SATKER", "ADMIN_UJIKOM"))
	wpdataGroup.GET("/calon-peserta", wpdataHandler.GetCalonPesertaList)
	wpdataGroup.GET("/calon-peserta/:id", wpdataHandler.GetCalonPesertaDetail)
	wpdataGroup.POST("/calon-peserta/:id/verify", wpdataHandler.VerifyCalonPeserta)
	wpdataGroup.GET("", wpdataHandler.GetWPDataList)
	wpdataGroup.POST("", wpdataHandler.CreateWPData)
	wpdataGroup.GET("/:id", wpdataHandler.GetWPDataDetail)
	wpdataGroup.PUT("/:id", wpdataHandler.UpdateWPData)
	wpdataGroup.DELETE("/:id", wpdataHandler.DeleteWPData)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "404", "message": "Page not found"})
	})

	uploadDir := cfg.UploadDir
	if uploadDir == "" {
		uploadDir = "uploads"
	}
	router.Static("/uploads", uploadDir)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Widyaprada Backend API",
			"version": "1.0",
			"endpoints": gin.H{
				"health":  "/_health",
				"swagger": "/swagger/index.html",
				"api":     "/api/v1",
			},
		})
	})

	router.GET("/_health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	router.GET("/_db-test", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Database OK", "db_type": cfg.DBType})
	})

	return router, nil
}
