package questionrepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// QuestionCategory DTO untuk question_categories
type QuestionCategory struct {
	ID        string     `gorm:"column:id;primaryKey;type:uuid"`
	Code      string     `gorm:"column:code;size:50;uniqueIndex;not null"`
	Name      string     `gorm:"column:name;size:255;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (QuestionCategory) TableName() string {
	return "question_categories"
}

func (c *QuestionCategory) ToEntity() entity.QuestionCategory {
	return entity.QuestionCategory{
		ID:   c.ID,
		Code: c.Code,
		Name: c.Name,
	}
}

// Question DTO untuk questions
type Question struct {
	ID                 string         `gorm:"column:id;primaryKey;type:uuid"`
	Code               string         `gorm:"column:code;size:100;uniqueIndex;not null"`
	Type               string         `gorm:"column:type;size:50;not null"`
	CategoryID         *string        `gorm:"column:category_id;type:uuid"`
	Difficulty         string         `gorm:"column:difficulty;size:20"`
	QuestionText       string         `gorm:"column:question_text;type:text;not null"`
	AnswerKey          string         `gorm:"column:answer_key;size:10"`
	Weight             float64        `gorm:"column:weight;default:1"`
	Status             string         `gorm:"column:status;size:20;default:Draft"`
	VerificationStatus string         `gorm:"column:verification_status;size:20;default:Belum"`
	DeletedReason      string         `gorm:"column:deleted_reason"`
	CreatedAt          *time.Time     `gorm:"column:created_at"`
	UpdatedAt          *time.Time     `gorm:"column:updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Question) TableName() string {
	return "questions"
}

func (q *Question) ToEntity() entity.QuestionDetail {
	createdAt, updatedAt := "", ""
	if q.CreatedAt != nil {
		createdAt = q.CreatedAt.UTC().Format(time.RFC3339)
	}
	if q.UpdatedAt != nil {
		updatedAt = q.UpdatedAt.UTC().Format(time.RFC3339)
	}
	catID := ""
	if q.CategoryID != nil {
		catID = *q.CategoryID
	}
	return entity.QuestionDetail{
		ID:                 q.ID,
		Code:               q.Code,
		Type:               q.Type,
		CategoryID:         catID,
		Difficulty:         q.Difficulty,
		QuestionText:       q.QuestionText,
		AnswerKey:          q.AnswerKey,
		Weight:             q.Weight,
		Status:             q.Status,
		VerificationStatus: q.VerificationStatus,
		DeletedReason:      q.DeletedReason,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}

func (Question) FromEntity(e entity.QuestionDetail) Question {
	var createdAt, updatedAt *time.Time
	if e.CreatedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.CreatedAt); err == nil {
			t := t.UTC()
			createdAt = &t
		}
	}
	if e.UpdatedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.UpdatedAt); err == nil {
			t := t.UTC()
			updatedAt = &t
		}
	}
	var catID *string
	if e.CategoryID != "" {
		s := e.CategoryID
		catID = &s
	}
	return Question{
		ID:                 e.ID,
		Code:               e.Code,
		Type:               e.Type,
		CategoryID:         catID,
		Difficulty:         e.Difficulty,
		QuestionText:       e.QuestionText,
		AnswerKey:          e.AnswerKey,
		Weight:             e.Weight,
		Status:             e.Status,
		VerificationStatus: e.VerificationStatus,
		DeletedReason:      e.DeletedReason,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}

// QuestionOption DTO untuk question_options
type QuestionOption struct {
	ID           string     `gorm:"column:id;primaryKey;type:uuid"`
	QuestionID   string     `gorm:"column:question_id;type:uuid;not null;index"`
	OptionKey    string     `gorm:"column:option_key;size:10;not null"`
	OptionText   string     `gorm:"column:option_text;type:text;not null"`
	IsCorrect    bool       `gorm:"column:is_correct;default:false"`
	OptionWeight float64    `gorm:"column:option_weight;default:1"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
}

func (QuestionOption) TableName() string {
	return "question_options"
}

func (o *QuestionOption) ToEntity() entity.QuestionOption {
	w := o.OptionWeight
	if w <= 0 {
		w = 1
	}
	return entity.QuestionOption{
		ID:           o.ID,
		QuestionID:   o.QuestionID,
		OptionKey:    o.OptionKey,
		OptionText:   o.OptionText,
		IsCorrect:    o.IsCorrect,
		OptionWeight: w,
	}
}
