package cbt

import (
	cbt_usecase "github.com/ProjectWidyaprada/backend/core/usecase/cbt"
)

type cbtHTTPHandler struct {
	cbtUsecase cbt_usecase.CBTUsecase
}

func NewCBTHTTPHandler(cbtUsecase cbt_usecase.CBTUsecase) *cbtHTTPHandler {
	return &cbtHTTPHandler{cbtUsecase: cbtUsecase}
}
