package api

import (
	"github.com/labstack/echo/v4"
)

func New(hdl *schduleHandler) *echo.Echo {
	e := echo.New()

	npi := e.Group("/npi")

	npi.POST("/create", hdl.CreateSchedule)
	npi.GET("", hdl.Hello)

	return e
}
