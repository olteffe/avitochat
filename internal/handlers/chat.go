package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/olteffe/avitochat/internal"
	"github.com/olteffe/avitochat/internal/models"
)

type ChatHandler struct {
	ChatUC internal.ChatUseCase
}

// CreateChatHandler - Create a chat between users
func (h *ChatHandler) CreateChatHandler(ctx echo.Context) error {
	var chat models.Chats
	if err := ctx.Bind(&chat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	chatID, err := h.ChatUC.CreateChatUseCase(chat)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, struct {
		Id string `json:"id"`
	}{
		chatID,
	})
}

// GetChatHandler - Get all user chats
func (h *ChatHandler) GetChatHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "implement me")
}
