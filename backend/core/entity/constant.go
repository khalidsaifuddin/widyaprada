package entity

import (
	"errors"
	"fmt"
	"regexp"
)

// Generic
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateKey   = errors.New("duplicate key violation")
	ErrInvalidData    = errors.New("invalid data")
	ErrDatabaseError  = errors.New("database error")
)

// Auth
var (
	ErrInvalidCredentials = errors.New("email/username atau kata sandi salah")
	ErrAccountInactive    = errors.New("akun ini tidak aktif")
	ErrInvalidToken       = errors.New("invalid token")
)

// Forgot Password
var (
	ErrResetTokenInvalid   = errors.New("token tidak valid atau kadaluarsa")
	ErrPasswordMismatch    = errors.New("konfirmasi kata sandi tidak cocok")
	ErrPasswordTooShort    = errors.New("kata sandi minimal 8 karakter")
)

// Registration
var (
	ErrEmailAlreadyRegistered = errors.New("email ini sudah terdaftar")
	ErrInvalidEmailFormat     = errors.New("format email tidak valid")
	ErrValidation             = errors.New("validasi gagal")
	ErrRoleNotFound           = errors.New("role default tidak ditemukan")
)

// User Management (SDD_Auth_Manajemen_Pengguna)
var (
	ErrEmailExists         = errors.New("email sudah digunakan")
	ErrUsernameExists      = errors.New("username sudah digunakan")
	ErrUserNotFound        = errors.New("pengguna tidak ditemukan")
	ErrCannotDeleteSelf    = errors.New("tidak dapat menghapus akun sendiri")
	ErrDeleteReasonRequired = errors.New("alasan penghapusan wajib diisi")
	ErrAtLeastOneRole      = errors.New("minimal satu role wajib dipilih")
)

// RBAC (SDD_RBAC)
var (
	ErrRoleInUse       = errors.New("role ini masih digunakan oleh pengguna")
	ErrPermissionInUse = errors.New("permission ini masih di-assign ke role")
)

// Bank Soal (SDD_Bank_Soal)
var (
	ErrQuestionNotFound      = errors.New("soal tidak ditemukan")
	ErrQuestionCodeExists    = errors.New("kode soal sudah digunakan")
	ErrQuestionInUseByPaket  = errors.New("soal ini digunakan oleh paket ujian")
	ErrQuestionDeleteReason  = errors.New("alasan penghapusan wajib diisi")
	ErrQuestionOptionsRequired = errors.New("opsi dan kunci jawaban wajib untuk PG, MRA, dan Benar-Salah")
	ErrCategoryNotFound      = errors.New("kategori tidak ditemukan")
)

// Paket Soal (SDD_Paket_Soal)
var (
	ErrPackageNotFound       = errors.New("paket soal tidak ditemukan")
	ErrPackageCodeExists     = errors.New("kode paket sudah digunakan")
	ErrPackageMinOneQuestion = errors.New("minimal 1 soal dalam paket")
	ErrPackageDeleteReason   = errors.New("alasan penghapusan wajib diisi")
	ErrPackageInUseByExam    = errors.New("paket ini digunakan oleh ujian")
)

// Manajemen Data WP (SDD_Manajemen_Data_WP)
var (
	ErrNIPExists      = errors.New("NIP sudah digunakan")
	ErrWPDataNotFound = errors.New("data widyaprada tidak ditemukan")
)

// Manajemen Uji Kompetensi (SDD_Manajemen_Uji_Kompetensi)
var (
	ErrExamNotFound         = errors.New("ujian tidak ditemukan")
	ErrExamCodeExists       = errors.New("kode ujian sudah digunakan")
	ErrExamMinContent       = errors.New("minimal 1 soal atau paket dalam ujian")
	ErrExamMinParticipant   = errors.New("minimal 1 peserta dalam ujian")
	ErrExamDeleteReason     = errors.New("alasan penghapusan wajib diisi")
	ErrExamNotDraft         = errors.New("hanya ujian status Draft yang dapat diedit")
	ErrExamAlreadyPublished = errors.New("ujian sudah diterbitkan")
)

// Constants
const (
	DefaultRoleCodePeserta = "PESERTA"
	BcryptCost             = 10
)

// EmailRegex untuk validasi format email
var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type RecordNotFoundError struct {
	Message string
	Err     error
}

func (e *RecordNotFoundError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Err.Error()
}

func (e *RecordNotFoundError) Unwrap() error {
	return e.Err
}

func WrapRecordNotFound(message string) error {
	return &RecordNotFoundError{Message: message, Err: ErrRecordNotFound}
}

func WrapRecordNotFoundf(format string, args ...interface{}) error {
	return &RecordNotFoundError{Message: fmt.Sprintf(format, args...), Err: ErrRecordNotFound}
}

func IsRecordNotFound(err error) bool {
	var recordNotFoundErr *RecordNotFoundError
	if errors.As(err, &recordNotFoundErr) {
		return true
	}
	return errors.Is(err, ErrRecordNotFound)
}
