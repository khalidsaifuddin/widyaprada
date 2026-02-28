package email

// SendPasswordEmailParams untuk kirim password ke email (one-time)
type SendPasswordEmailParams struct {
	To       string
	Name     string
	Password string
}

// SendPasswordResetLinkParams untuk kirim link reset password
type SendPasswordResetLinkParams struct {
	To       string
	Name     string
	ResetURL string
}

// EmailService interface untuk kirim email
type EmailService interface {
	SendPasswordEmail(params SendPasswordEmailParams) error
	SendPasswordResetLink(params SendPasswordResetLinkParams) error
}
