package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/olteffe/avitochat/internal"
	"github.com/olteffe/avitochat/internal/models"
)

type ChatHandler struct {
	ChatUC internal.ChatUseCase
}

// CreateChatHandler - Create a chat between users
func (h *ChatHandler) CreateChatHandler(ctx echo.Context) error {
	var chat models.ChatForm
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

// CreateUserHandler - Create new user
func (h *ChatHandler) CreateUserHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.)
}

// GetChatHandler - Get all user chats
func (h *ChatHandler) GetChatHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.)
}

// GetMessagesHandler - Get all chat messages
func (h *ChatHandler) GetMessagesHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.)
}

// SendMessageHandler - Send a user message
func (h *ChatHandler) SendMessageHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.)
}
