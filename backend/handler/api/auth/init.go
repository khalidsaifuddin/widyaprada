package auth

import (
	auth_usecase "github.com/ProjectWidyaprada/backend/core/usecase/auth"
	forgotpassword_usecase "github.com/ProjectWidyaprada/backend/core/usecase/forgotpassword"
	registration_usecase "github.com/ProjectWidyaprada/backend/core/usecase/registration"
	"github.com/gin-gonic/gin"
)

type AuthHTTPHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type authHTTPHandler struct {
	authUsecase           auth_usecase.AuthUsecase
	registrationUsecase   registration_usecase.RegistrationUsecase
	forgotPasswordUsecase forgotpassword_usecase.ForgotPasswordUsecase
}

func NewAuthHTTPHandler(
	authUsecase auth_usecase.AuthUsecase,
	registrationUsecase registration_usecase.RegistrationUsecase,
	forgotPasswordUsecase forgotpassword_usecase.ForgotPasswordUsecase,
) AuthHTTPHandler {
	return &authHTTPHandler{
		authUsecase:           authUsecase,
		registrationUsecase:   registrationUsecase,
		forgotPasswordUsecase: forgotPasswordUsecase,
	}
}
