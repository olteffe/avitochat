package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/models"
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
	var user models.Users
	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userID, err := h.useCases.CreateUserUseCase(&user)
	if err != nil {
		if err == errors.New("invalid username") {
			return echo.NewHTTPError(http.StatusConflict, "Username already used")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, struct {
		ID string `json:"id"`
	}{
		userID,
	})
}
