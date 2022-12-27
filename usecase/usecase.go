package usecase

import (
	"github.com/yoteerakrit/assessment/domain"
	"github.com/yoteerakrit/assessment/pkg/merror"
)

type UseCase interface {
	CreateSchedule(input *CreateScheduleInput) merror.Error
}

type useCase struct {
	scheduleRepo domain.ScheduleRepo
}

func New(scheudleRepo domain.ScheduleRepo) *useCase {
	return &useCase{scheduleRepo: scheudleRepo}
}
