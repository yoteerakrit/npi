package usecase

import (
	"time"

	"github.com/yoteerakrit/assessment/domain"
	"github.com/yoteerakrit/assessment/pkg/merror"
)

type CreateScheduleInput struct {
	Description string              `json:"description"`
	Round       []*CreateRoundInput `json:"round"`
}

type CreateRoundInput struct {
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (u *useCase) CreateSchedule(input *CreateScheduleInput) merror.Error {

	u.scheduleRepo.CreateSchedule(&domain.Schedule{})

	return merror.New("test")
}
