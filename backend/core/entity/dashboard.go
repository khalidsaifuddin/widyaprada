package entity

// DashboardAssignmentItem item GET /api/v1/dashboard/assignments
type DashboardAssignmentItem struct {
	ID                string   `json:"id"`
	ExamName          string   `json:"exam_name"`
	Deadline          string   `json:"deadline"`
	Status            string   `json:"status"` // belum_dikerjakan | sudah_dikerjakan
	Score             *float64 `json:"score,omitempty"`
	CanStart          bool     `json:"can_start"`
	CanViewResult     bool     `json:"can_view_result"`
	CanViewLeaderboard bool    `json:"can_view_leaderboard"`
}

// DashboardAssignmentsResponse GET /api/v1/dashboard/assignments
type DashboardAssignmentsResponse struct {
	Data []DashboardAssignmentItem `json:"data"`
	Meta struct {
		Total int64 `json:"total"`
	} `json:"meta"`
}

// DashboardJournalItem item GET /api/v1/dashboard/journals
type DashboardJournalItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	SubmittedAt string `json:"submitted_at"`
	Status      string `json:"status"` // Draft|Menunggu Verifikasi|Diverifikasi|Ditolak|Published
}

// DashboardJournalsResponse GET /api/v1/dashboard/journals
type DashboardJournalsResponse struct {
	Data []DashboardJournalItem `json:"data"`
	Meta struct {
		Total int64 `json:"total"`
	} `json:"meta"`
}
