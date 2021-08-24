package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal"
	"net/http"
)

type UserHandler struct {
	ChatUC internal.ChatUseCase
}

// CreateUserHandler - Create new user
func (h *ChatHandler) CreateUserHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "implement me")
}
