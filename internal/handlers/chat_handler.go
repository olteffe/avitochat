package handlers

import (
	"errors"
	mError "github.com/olteffe/avitochat/internal/message_error"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/models"
)

// user used for GetChatHandler
type user struct {
	ID string `json:"id"`
}

// initChatRoutes - Unites paths
func (h *Handler) initChatRoutes(api *echo.Group) {
	chats := api.Group("/chats")
	{
		chats.POST("/add", h.CreateChatHandler)
		chats.POST("/get", h.GetChatHandler)
	}
}

// CreateChatHandler - Create a chat between users
func (h *Handler) CreateChatHandler(ctx echo.Context) error {
	var input struct {
		Name  string   `json:"name"`
		Users []string `json:"users"`
	}
	if err := ctx.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	chat := &models.Chats{
		Name:  input.Name,
		Users: input.Users,
	}
	chatID, err := h.useCases.Chat.CreateChatUseCase(chat)
	if err != nil {
		switch {
		case errors.Is(err, mError.ErrUserInvalid):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mError.ErrChatInvalid):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mError.ErrCountUsers):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mError.ErrDB):
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, struct {
		Id string `json:"id"`
	}{
		chatID,
	})
}

// GetChatHandler - Get a list of user's chats
func (h *Handler) GetChatHandler(ctx echo.Context) error {
	var userID user
	if err := ctx.Bind(&userID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	allChats, err := h.useCases.Chat.GetChatUseCase(userID.ID)
	if err != nil {
		switch {
		case errors.Is(err, mError.ErrUserInvalid):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mError.ErrUserIdInvalid):
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		case errors.Is(err, mError.ErrDB):
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, allChats)
}
