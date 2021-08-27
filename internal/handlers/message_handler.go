package handlers

import (
	"github.com/labstack/echo/v4"
	//"github.com/olteffe/avitochat/internal"
	"net/http"
)

func (h *Handler) initMessageRoutes(api *echo.Group) {
	messages := api.Group("/messages")
	{
		messages.POST("/add", h.GetMessagesHandler)
		messages.POST("/get", h.SendMessageHandler)
	}
}

// GetMessagesHandler - Get all chat messages
func (h *Handler) GetMessagesHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "implement me")
}

// SendMessageHandler - Send a user message
func (h *Handler) SendMessageHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "implement me")
}
