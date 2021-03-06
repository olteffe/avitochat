package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	mError "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"net/http"
)

func (h *Handler) initUserRoutes(api *echo.Group) {
	users := api.Group("/users")
	{
		users.POST("/add", h.CreateUserHandler)
	}
}

// CreateUserHandler - Create new user
func (h *Handler) CreateUserHandler(ctx echo.Context) error {
	var input struct {
		Username string `json:"username"`
	}
	if err := ctx.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	user := &models.Users{
		Username: input.Username,
	}
	userID, err := h.useCases.User.CreateUserUseCase(user)
	if err != nil {
		switch {
		case errors.Is(err, mError.ErrUserInvalid):
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		case errors.Is(err, mError.ErrUserAlreadyUsed):
			return echo.NewHTTPError(http.StatusConflict, "Username already used")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
	}
	return ctx.JSON(http.StatusCreated, struct {
		ID string `json:"id"`
	}{
		userID,
	})
}
