package dashboard

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type DashboardUsecase interface {
	GetAssignments(ctx context.Context, userID string, limit, page int64) (*entity.DashboardAssignmentsResponse, error)
	GetMyJournals(ctx context.Context, userID string, limit, page int64) (*entity.DashboardJournalsResponse, error)
}

type dashboardUsecase struct {
	assignmentRepo repository.AssignmentRepo
	journalRepo    repository.JournalRepo
}

func NewDashboardUsecase(assignmentRepo repository.AssignmentRepo, journalRepo repository.JournalRepo) DashboardUsecase {
	return &dashboardUsecase{
		assignmentRepo: assignmentRepo,
		journalRepo:    journalRepo,
	}
}
