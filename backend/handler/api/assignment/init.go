package assignment

import (
	assignment_usecase "github.com/ProjectWidyaprada/backend/core/usecase/assignment"
)

type assignmentHTTPHandler struct {
	assignmentUsecase assignment_usecase.AssignmentUsecase
}

func NewAssignmentHTTPHandler(assignmentUsecase assignment_usecase.AssignmentUsecase) *assignmentHTTPHandler {
	return &assignmentHTTPHandler{assignmentUsecase: assignmentUsecase}
}
