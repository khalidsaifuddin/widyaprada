package user

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	usermanagement_usecase "github.com/ProjectWidyaprada/backend/core/usecase/usermanagement"
	"github.com/gin-gonic/gin"
)

type UserHTTPHandler interface {
	GetUserList(c *gin.Context)
	GetUserDetail(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHTTPHandler struct {
	userUsecase usermanagement_usecase.UserUsecase
	userRepo    repository.UserRepo
}

func NewUserHTTPHandler(userUsecase usermanagement_usecase.UserUsecase, userRepo repository.UserRepo) UserHTTPHandler {
	return &userHTTPHandler{
		userUsecase: userUsecase,
		userRepo:    userRepo,
	}
}
