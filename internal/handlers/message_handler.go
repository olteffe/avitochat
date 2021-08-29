package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/models"

	//"github.com/olteffe/avitochat/internal"
	"net/http"
)

// chat used for input data
type chat struct {
	ID string `json:"id"`
}

// initMessageRoutes - Unites paths
func (h *Handler) initMessageRoutes(api *echo.Group) {
	messages := api.Group("/messages")
	{
		messages.POST("/add", h.GetMessagesHandler)
		messages.POST("/get", h.SendMessageHandler)
	}
}

// GetMessagesHandler - Get all chat messages
func (h *Handler) GetMessagesHandler(ctx echo.Context) error {
	var chatID chat
	if err := ctx.Bind(&chatID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid chat ID")
	}
	allMessages, err := h.useCases.GetMessagesUseCase(chatID.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, allMessages)
}

// SendMessageHandler - Send a user message
func (h *Handler) SendMessageHandler(ctx echo.Context) error {
	var message models.Messages
	if err := ctx.Bind(&message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	messageID, err := h.useCases.SendMessageUseCase(message)
	if err != nil {
		if err == errors.New("user or chat not found") {
			return echo.NewHTTPError(http.StatusNotFound, "User or chat not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, struct {
		ID string `json:"id"`
	}{
		messageID,
	})
}
