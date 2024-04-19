package main

import (
	"log"

	"github.com/harryhcs/conservio_web/handler"
	"github.com/hasura/go-graphql-client"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := envs["GQL_API"]

	app := echo.New()
	app.Static("/static", "static")

	client := graphql.NewClient(api, nil)

	paymentStatusHandler := handler.PayStatusHandler{
		Client: client,
	}

	app.GET("/booking/verify/:payID", paymentStatusHandler.HandleShowPayStatus)

	app.Start(":3030")
}
