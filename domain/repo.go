package domain

type ScheduleRepo interface {
	CreateSchedule(schedule *Schedule) (*Schedule, error)
}

type RoundRepo interface {
	CreateRound(round *Round) (*Round, error)
}

type AccountRepo interface{}

type HistoryRepo interface{}

type ParticipantRepo interface{}
