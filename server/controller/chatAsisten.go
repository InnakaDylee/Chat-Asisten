package controller

import (
	"chat-asisten/config"
	"chat-asisten/model"
	"chat-asisten/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetChatAsistenController(ctx echo.Context) error {
	var message []model.Message
	if err := config.DB.Find(&message).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get message",
		"data":    message,
	})
}

func CreateChatAsistenController(ctx echo.Context) error {
	message := model.Message{}
	err := ctx.Bind(&message)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
		})
	}

	if message.Content == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "content is required",
		})
	}

	answer, err := service.ResolveComplaint(ctx, message.Content)
	convertAnswer := model.Message{
		Content: answer,
		Role:    "asisten",
	}

	if err := config.DB.Save(&message).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gbsa gorm? 1",
		})
	}

	if err := config.DB.Save(&convertAnswer).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gbsa gorm? 2",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create message",
		"data":    convertAnswer,
	})
}
