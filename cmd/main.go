package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	payStatusHandler := handler.payStatusHandler{}

	app.GET("/booking/verify/:payID", payStatusHandler.HandleShowPayStatus)

	app.Start(":3030")
}
