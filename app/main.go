package main

import (
	"log"
	"net/http"

	"github.com/yoteerakrit/assessment/app/api"
	"github.com/yoteerakrit/assessment/pkg/merror"
	"github.com/yoteerakrit/assessment/repo"
	"github.com/yoteerakrit/assessment/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	mysql := repo.NewMySQL()

	// inject
	scheduleRepo := repo.NewSchedule(mysql)
	uc := usecase.New(scheduleRepo)
	scheduleHandler := api.NewHandler(uc)

	e := api.New(scheduleHandler)

	e.HTTPErrorHandler = customHTTPErrorHandler

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var (
		m   *merror.MError
		msg string
	)

	if me, ok := err.(*merror.MError); ok {
		m = me
	} else if he, ok := err.(*echo.HTTPError); ok {
		m = merror.ErrorType{
			Status: he.Code,
			Code:   "",
			Msg:    he.Error(),
			DevMsg: he.Error(),
		}.New().(*merror.MError)
	} else {
		m = merror.InternalError().New().(*merror.MError)
	}

	msg = m.Msg

	c.JSON(m.Status, msg)
}
