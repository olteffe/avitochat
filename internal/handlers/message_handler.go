package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	mErr "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"net/http"
)

// initMessageRoutes - Unites paths
func (h *Handler) initMessageRoutes(api *echo.Group) {
	messages := api.Group("/messages")
	{
		messages.POST("/add", h.SendMessageHandler)
		messages.POST("/get", h.GetMessagesHandler)
	}
}

// GetMessagesHandler - Get all chat messages
func (h *Handler) GetMessagesHandler(ctx echo.Context) error {
	var chatID struct {
		ID string `json:"chat"`
	}
	if err := ctx.Bind(&chatID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	message := &models.Messages{
		Chat: chatID.ID,
	}
	allMessages, err := h.useCases.Message.GetMessagesUseCase(message)
	if err != nil {
		switch {
		case errors.Is(err, mErr.ErrChatIdInvalid):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mErr.ErrUserOrChat):
			return echo.NewHTTPError(http.StatusNotFound, "Chat not found")
		case errors.Is(err, mErr.ErrDB):
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, allMessages)
}

// SendMessageHandler - Send a user message
func (h *Handler) SendMessageHandler(ctx echo.Context) error {
	var input struct {
		ChatId   string `json:"chat"`
		AuthorId string `json:"author"`
		Text     string `json:"text"`
	}
	if err := ctx.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	message := &models.Messages{
		Chat:   input.ChatId,
		Author: input.AuthorId,
		Text:   input.Text,
	}
	messageID, err := h.useCases.Message.SendMessageUseCase(message)
	if err != nil {
		switch {
		case errors.Is(err, mErr.ErrChatIdInvalid):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mErr.ErrUserIdInvalid):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mErr.ErrUserOrChat):
			return echo.NewHTTPError(http.StatusNotFound, "User or chat not found")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
	}
	return ctx.JSON(http.StatusCreated, struct {
		ID string `json:"id"`
	}{
		messageID,
	})
}
