package api

import (
	"github.com/yoteerakrit/assessment/usecase"
)

type schduleHandler struct {
	uc usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *schduleHandler {
	return &schduleHandler{uc: uc}
}
