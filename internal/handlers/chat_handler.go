package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/models"
)

func (h *Handler) initChatRoutes(api *echo.Group) {
	chats := api.Group("/chats")
	{
		chats.POST("/add", h.CreateChatHandler)
		chats.POST("/get", h.GetChatHandler)
	}
}

// CreateChatHandler - Create a chat between users
func (h *Handler) CreateChatHandler(ctx echo.Context) error {
	var chat models.Chats
	if err := ctx.Bind(&chat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	chatID, err := h.useCases.CreateChatUseCase(&chat)
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
func (h *Handler) GetChatHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "implement me")
}
