package main

import (
	"chat-asisten/config"
	"chat-asisten/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome to RESTful API Services")
	})
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
