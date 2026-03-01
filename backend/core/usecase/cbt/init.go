package cbt

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

// CBTUsecase interface
type CBTUsecase interface {
	ListUjianTersedia(ctx context.Context, userID string) (*entity.CBTListExamsResponse, error)
	GetExamDetail(ctx context.Context, userID, examID string) (*entity.CBTExamDetailResponse, error)
	MulaiUjian(ctx context.Context, userID, examID string) (*entity.CBTStartResponse, error)
	GetSoal(ctx context.Context, userID, attemptID string) (*entity.CBTListQuestionsResponse, error)
	GetSoalByNomor(ctx context.Context, userID, attemptID string, num int) (*entity.CBTQuestionItem, error)
	SimpanJawaban(ctx context.Context, userID, attemptID string, req entity.CBTSaveAnswerRequest) error
	SubmitUjian(ctx context.Context, userID, attemptID string) (*entity.CBTSubmitResponse, error)
	GetRiwayatHasil(ctx context.Context, userID string) (*entity.CBTHistoryResponse, error)
}

type cbtUsecase struct {
	cbtRepo     repository.CBTRepo
	assignRepo  repository.AssignmentRepo
}

func NewCBTUsecase(cbtRepo repository.CBTRepo, assignRepo repository.AssignmentRepo) CBTUsecase {
	return &cbtUsecase{
		cbtRepo:    cbtRepo,
		assignRepo: assignRepo,
	}
}
