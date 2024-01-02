package routes

import (
	"chat-asisten/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) (bool, error) {
			return origin == "http://localhost:3000", nil
		},
	}))


	e.GET("/chat-asisten", controller.GetChatAsistenController)
	e.POST("/chat-asisten", controller.CreateChatAsistenController)

}