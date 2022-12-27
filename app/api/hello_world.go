package api

import (
	"fmt"
	"net/http"

	"github.com/yoteerakrit/assessment/domain"

	"github.com/labstack/echo/v4"
)

func (h *schduleHandler) Hello(c echo.Context) error {
	var schedule domain.Schedule
	if err := c.Bind(&schedule); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(schedule)

	return c.JSON(http.StatusOK, schedule)
}
