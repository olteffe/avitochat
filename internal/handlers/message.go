package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal"
	"net/http"
)

type MessageHandler struct {
	ChatUC internal.ChatUseCase
}

// GetMessagesHandler - Get all chat messages
func (h *MessageHandler) GetMessagesHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "implement me")
}

// SendMessageHandler - Send a user message
func (h *MessageHandler) SendMessageHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "implement me")
}
