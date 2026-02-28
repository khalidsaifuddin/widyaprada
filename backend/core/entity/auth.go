package entity

// User untuk domain/repository (internal)
type User struct {
	ID           string
	Name         string
	Email        string
	Username     string
	NIP          string
	PasswordHash string
	SatkerID     *string
	IsActive     bool
	CreatedAt    string
	UpdatedAt    string
}

// LoginRequest adalah request untuk POST /api/v1/auth/login
type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"` // email atau username
	Password   string `json:"password" binding:"required"`
}

// LoginResponse sesuai SDD: access_token, token_type, expires_in, user
type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	TokenType   string       `json:"token_type"`
	ExpiresIn   int          `json:"expires_in"`
	User        UserResponse `json:"user"`
}

// UserResponse untuk response login (user object)
type UserResponse struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Username        string     `json:"username"`
	Roles           []RoleInfo `json:"roles"`
	DefaultHomePath string     `json:"default_home_path"`
}

// RoleInfo untuk roles di response
type RoleInfo struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// RegisterRequest untuk POST /api/v1/auth/register
type RegisterRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	NIP   string `json:"nip"`
}

// RegisterResponse response sukses registrasi
type RegisterResponse struct {
	Message string `json:"message"`
}

// LogoutResponse response sukses logout
type LogoutResponse struct {
	Message string `json:"message"`
}

// ForgotPasswordRequest untuk POST /api/v1/auth/forgot-password
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

// ForgotPasswordResponse response sukses forgot password
type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

// ResetPasswordRequest untuk POST /api/v1/auth/reset-password
type ResetPasswordRequest struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

// ResetPasswordResponse response sukses reset password
type ResetPasswordResponse struct {
	Message string `json:"message"`
}
