package examrepo

import (
	"time"

	"gorm.io/gorm"
)

// Exam DTO
type Exam struct {
	ID                   string         `gorm:"column:id;primaryKey;type:uuid"`
	Code                 string         `gorm:"column:code;size:100;uniqueIndex;not null"`
	Name                 string         `gorm:"column:name;size:255;not null"`
	JadwalMulai          *time.Time     `gorm:"column:jadwal_mulai;not null"`
	JadwalSelesai        *time.Time     `gorm:"column:jadwal_selesai;not null"`
	DurasiMenit          int            `gorm:"column:durasi_menit;not null"`
	Status               string         `gorm:"column:status;size:20;default:Draft"`
	VerificationStatus   string         `gorm:"column:verification_status;size:20;default:Belum"`
	ShuffleQuestions     bool           `gorm:"column:shuffle_questions;default:true"`
	TampilkanLeaderboard bool           `gorm:"column:tampilkan_leaderboard;default:false"`
	DeletedReason        string         `gorm:"column:deleted_reason"`
	CreatedAt            *time.Time     `gorm:"column:created_at"`
	UpdatedAt            *time.Time     `gorm:"column:updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Exam) TableName() string {
	return "exams"
}

// ExamContent DTO
type ExamContent struct {
	ExamID     string     `gorm:"column:exam_id;primaryKey;type:uuid"`
	SourceType string     `gorm:"column:source_type;size:20;not null"`
	SourceID   string     `gorm:"column:source_id;primaryKey;type:uuid"`
	SortOrder  int        `gorm:"column:sort_order;not null"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
}

func (ExamContent) TableName() string {
	return "exam_contents"
}

// ExamParticipant DTO
type ExamParticipant struct {
	ExamID    string     `gorm:"column:exam_id;primaryKey;type:uuid"`
	UserID    string     `gorm:"column:user_id;primaryKey;type:uuid"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

func (ExamParticipant) TableName() string {
	return "exam_participants"
}

// ExamAttempt DTO - satu attempt per user per exam (CBT & Assignment result)
type ExamAttempt struct {
	ID          string     `gorm:"column:id;primaryKey;type:uuid"`
	ExamID      string     `gorm:"column:exam_id;type:uuid;not null;index"`
	UserID      string     `gorm:"column:user_id;type:uuid;not null;index"`
	StartedAt   *time.Time `gorm:"column:started_at"`
	SubmittedAt *time.Time `gorm:"column:submitted_at"`
	Score       *float64   `gorm:"column:score"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (ExamAttempt) TableName() string {
	return "exam_attempts"
}

// ExamAnswer DTO - jawaban per soal per attempt
type ExamAnswer struct {
	AttemptID   string     `gorm:"column:attempt_id;primaryKey;type:uuid"`
	QuestionID  string     `gorm:"column:question_id;primaryKey;type:uuid"`
	AnswerValue string     `gorm:"column:answer_value;type:text"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (ExamAnswer) TableName() string {
	return "exam_answers"
}

// ExamAttemptQuestion urutan soal untuk attempt (fix order saat mulai)
type ExamAttemptQuestion struct {
	AttemptID  string     `gorm:"column:attempt_id;primaryKey;type:uuid"`
	QuestionID string     `gorm:"column:question_id;primaryKey;type:uuid"`
	SortOrder  int        `gorm:"column:sort_order;not null"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
}

func (ExamAttemptQuestion) TableName() string {
	return "exam_attempt_questions"
}
