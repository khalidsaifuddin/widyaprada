package email

import "log"

// LogEmailService implementasi EmailService yang hanya log ke stdout (dev/test)
// Untuk production, ganti dengan SMTP/SendGrid/dll
type LogEmailService struct{}

func NewLogEmailService() EmailService {
	return &LogEmailService{}
}

func (s *LogEmailService) SendPasswordEmail(params SendPasswordEmailParams) error {
	log.Printf("[EMAIL] To: %s | Name: %s | Password: %s (ini hanya log, tidak kirim email sungguhan)", params.To, params.Name, params.Password)
	return nil
}

func (s *LogEmailService) SendPasswordResetLink(params SendPasswordResetLinkParams) error {
	log.Printf("[EMAIL] To: %s | Name: %s | ResetURL: %s (ini hanya log, tidak kirim email sungguhan)", params.To, params.Name, params.ResetURL)
	return nil
}
