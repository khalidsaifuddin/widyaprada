package entity

import "errors"

var (
	ErrCBTExamNotAvailable   = errors.New("ujian tidak tersedia")
	ErrCBTNotParticipant     = errors.New("anda bukan peserta ujian ini")
	ErrCBTAlreadyStarted     = errors.New("sudah pernah mulai ujian ini")
	ErrCBTAttemptNotFound    = errors.New("attempt tidak ditemukan")
	ErrCBTAttemptNotOwned    = errors.New("attempt bukan milik anda")
	ErrCBTAlreadySubmitted   = errors.New("ujian sudah disubmit")
	ErrCBTTimeUp             = errors.New("waktu ujian telah habis")
)

// CBTExamTersediaItem item daftar ujian tersedia
type CBTExamTersediaItem struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	JadwalMulai   string `json:"jadwal_mulai"`
	JadwalSelesai string `json:"jadwal_selesai"`
	DurasiMenit   int    `json:"durasi_menit"`
}

// CBTListExamsResponse response GET /api/v1/cbt/exams
type CBTListExamsResponse struct {
	Items []CBTExamTersediaItem `json:"items"`
}

// CBTStartResponse response POST /api/v1/cbt/exams/:id/start
type CBTStartResponse struct {
	AttemptID     string `json:"attempt_id"`
	ExamID        string `json:"exam_id"`
	StartedAt     string `json:"started_at"`
	DurasiMenit   int    `json:"durasi_menit"`
	JadwalSelesai string `json:"jadwal_selesai"`
}

// CBTQuestionItem soal untuk CBT (tanpa kunci jawaban)
type CBTQuestionItem struct {
	Num          int               `json:"num"`
	QuestionID   string            `json:"question_id"`
	Type         string            `json:"type"`
	QuestionText string            `json:"question_text"`
	Weight       float64           `json:"weight"`
	Options      []CBTQuestionOption `json:"options,omitempty"`
}

// CBTQuestionOption opsi untuk PG/B-S (tanpa is_correct)
type CBTQuestionOption struct {
	ID         string `json:"id"`
	OptionKey  string `json:"option_key"`
	OptionText string `json:"option_text"`
}

// CBTListQuestionsResponse response GET /api/v1/cbt/attempts/:attemptId/questions
type CBTListQuestionsResponse struct {
	AttemptID string            `json:"attempt_id"`
	Total     int               `json:"total"`
	Questions []CBTQuestionItem `json:"questions"`
}

// CBTSaveAnswerRequest untuk POST /api/v1/cbt/attempts/:attemptId/answers
type CBTSaveAnswerRequest struct {
	QuestionID  string `json:"question_id" binding:"required"`
	OptionID    string `json:"option_id"`    // PG, B-S
	AnswerText  string `json:"answer_text"`  // Essay
}

// CBTSaveAnswerResponse response
type CBTSaveAnswerResponse struct {
	Message string `json:"message"`
}

// CBTSubmitResponse response POST /api/v1/cbt/attempts/:attemptId/submit
type CBTSubmitResponse struct {
	AttemptID   string   `json:"attempt_id"`
	Score       *float64 `json:"score,omitempty"`
	SubmittedAt string   `json:"submitted_at"`
}

// CBTHistoryItem item riwayat ujian
type CBTHistoryItem struct {
	AttemptID   string   `json:"attempt_id"`
	ExamID      string   `json:"exam_id"`
	ExamName    string   `json:"exam_name"`
	StartedAt   string   `json:"started_at"`
	SubmittedAt string   `json:"submitted_at"`
	Score       *float64 `json:"score,omitempty"`
}

// CBTHistoryResponse response GET /api/v1/cbt/history
type CBTHistoryResponse struct {
	Items []CBTHistoryItem `json:"items"`
}
