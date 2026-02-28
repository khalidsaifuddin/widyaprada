package cms

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	cms_usecase "github.com/ProjectWidyaprada/backend/core/usecase/cms"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type CMSHTTPHandler interface {
	// Slider
	GetSlideList(c *gin.Context)
	GetSlideDetail(c *gin.Context)
	CreateSlide(c *gin.Context)
	UpdateSlide(c *gin.Context)
	DeleteSlide(c *gin.Context)
	// Berita
	GetArticleList(c *gin.Context)
	GetArticleDetail(c *gin.Context)
	CreateArticle(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
	// Tautan
	GetLinkList(c *gin.Context)
	GetLinkDetail(c *gin.Context)
	CreateLink(c *gin.Context)
	UpdateLink(c *gin.Context)
	DeleteLink(c *gin.Context)
}

type cmsHTTPHandler struct {
	sliderUsecase   cms_usecase.SliderUsecase
	beritaUsecase   cms_usecase.BeritaCMSUsecase
	tautanUsecase   cms_usecase.TautanUsecase
	userRepo        repository.UserRepo
}

func NewCMSHTTPHandler(sliderUsecase cms_usecase.SliderUsecase, beritaUsecase cms_usecase.BeritaCMSUsecase, tautanUsecase cms_usecase.TautanUsecase, userRepo repository.UserRepo) CMSHTTPHandler {
	return &cmsHTTPHandler{
		sliderUsecase: sliderUsecase,
		beritaUsecase: beritaUsecase,
		tautanUsecase: tautanUsecase,
		userRepo:      userRepo,
	}
}

func (h *cmsHTTPHandler) getActor(c *gin.Context) *cms_usecase.ActorContext {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		return nil
	}
	actor := &cms_usecase.ActorContext{
		UserID:    claims.UserID,
		RoleCodes: claims.Roles,
	}
	if !actor.IsSuperAdmin() {
		u, err := h.userRepo.GetByID(c.Request.Context(), claims.UserID)
		if err == nil && u != nil && u.SatkerID != nil {
			actor.SatkerID = u.SatkerID
		}
	}
	return actor
}
