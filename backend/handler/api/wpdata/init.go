package wpdata

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	wpdata_usecase "github.com/ProjectWidyaprada/backend/core/usecase/wpdata"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

type WPDataHTTPHandler interface {
	GetWPDataList(c *gin.Context)
	GetWPDataDetail(c *gin.Context)
	CreateWPData(c *gin.Context)
	UpdateWPData(c *gin.Context)
	DeleteWPData(c *gin.Context)
	GetCalonPesertaList(c *gin.Context)
	GetCalonPesertaDetail(c *gin.Context)
	VerifyCalonPeserta(c *gin.Context)
}

type wpdataHTTPHandler struct {
	wpdataUsecase        wpdata_usecase.WPDataUsecase
	calonPesertaUsecase  wpdata_usecase.CalonPesertaUsecase
	userRepo             repository.UserRepo
}

func NewWPDataHTTPHandler(wpdataUsecase wpdata_usecase.WPDataUsecase, calonPesertaUsecase wpdata_usecase.CalonPesertaUsecase, userRepo repository.UserRepo) WPDataHTTPHandler {
	return &wpdataHTTPHandler{
		wpdataUsecase:       wpdataUsecase,
		calonPesertaUsecase: calonPesertaUsecase,
		userRepo:            userRepo,
	}
}

func (h *wpdataHTTPHandler) getActor(c *gin.Context) *wpdata_usecase.ActorContext {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		return nil
	}
	actor := &wpdata_usecase.ActorContext{
		UserID:    claims.UserID,
		RoleCodes: claims.Roles,
	}
	isSuperAdmin := false
	for _, r := range claims.Roles {
		if r == "SUPER_ADMIN" {
			isSuperAdmin = true
			break
		}
	}
	if !isSuperAdmin {
		u, err := h.userRepo.GetByID(c.Request.Context(), claims.UserID)
		if err == nil && u != nil && u.SatkerID != nil {
			actor.SatkerID = u.SatkerID
		}
	}
	return actor
}
