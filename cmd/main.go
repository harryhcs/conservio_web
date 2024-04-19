package main

import (
	"github.com/harryhcs/conservio_web/handler"
	"github.com/hasura/go-graphql-client"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	app.Static("/static", "static")

	client := graphql.NewClient("https://staging.graph.conservio.com/graphql", nil)

	paymentStatusHandler := handler.PayStatusHandler{
		Client: client,
	}

	app.GET("/booking/verify/:payID", paymentStatusHandler.HandleShowPayStatus)

	app.Start(":3030")
}
