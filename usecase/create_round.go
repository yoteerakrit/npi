package usecase

import (
	"github.com/yoteerakrit/assessment/domain"
	"github.com/yoteerakrit/assessment/pkg/merror"
)

func (u *useCase) CreateRound(round *domain.Round) (*domain.Round, error) {

	u.scheduleRepo.CreateSchedule(&domain.Schedule{})

	return nil, merror.New("test")
}
