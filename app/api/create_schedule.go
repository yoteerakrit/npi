package api

import (
	"net/http"
	"time"

	"github.com/yoteerakrit/assessment/usecase"

	"github.com/labstack/echo/v4"
)

type CreateScheduleRequest struct {
	Description string                `json:"description" validate:"required"`
	Round       []*CreateRoundRequest `json:"round"`
}

type CreateRoundRequest struct {
	Description string    `json:"description" validate:"required"`
	Date        time.Time `json:"date"`
}

func (hdl *schduleHandler) CreateSchedule(c echo.Context) error {
	var req CreateScheduleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	round := make([]*usecase.CreateRoundInput, 0, len(req.Round))
	for _, v := range req.Round {
		round = append(round, &usecase.CreateRoundInput{
			Description: v.Description,
			Date:        v.Date,
		})
	}

	err := hdl.uc.CreateSchedule(&usecase.CreateScheduleInput{
		Description: req.Description,
		Round:       round,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, req.Round)
}
