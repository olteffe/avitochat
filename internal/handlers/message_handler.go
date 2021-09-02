package handlers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/models"

	//"github.com/olteffe/avitochat/internal"
	"net/http"
)

// chat input data for GetMessagesHandler
type chat struct {
	ID string `json:"id"`
}

// messageInput input data for SendMessageHandler
type messageInput struct {
	ChatId   string `json:"chat"`
	AuthorId string `json:"author"`
	Text     string `json:"text"`
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
	if _, err := uuid.Parse(chatID.ID); err != nil {
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
	var input messageInput
	if err := ctx.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// simple input validation
	if _, err := uuid.Parse(input.ChatId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid chat or author ID")
	}
	if _, err := uuid.Parse(input.AuthorId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid chat or author ID")
	}
	message := &models.Messages{
		Chat:   input.ChatId,
		Author: input.AuthorId,
		Text:   input.Text,
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
