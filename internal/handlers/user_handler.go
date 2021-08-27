package handlers

import (
	"github.com/labstack/echo/v4"
	//"github.com/olteffe/avitochat/internal"
	"net/http"
)

func (h *Handler) initUserRoutes(api *echo.Group) {
	users := api.Group("/users")
	{
		users.POST("/add", h.CreateChatHandler)
	}
}

// CreateUserHandler - Create new user
func (h *Handler) CreateUserHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, "implement me")
}
