package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/usecase"
)

type Handler struct {
	useCases *usecase.UseCase
}

func NewHandler(useCases *usecase.UseCase) *Handler {
	return &Handler{
		useCases: useCases,
	}
}

func (h *Handler) Init(router *echo.Echo) {
	api := router.Group("")
	{
		h.initChatRoutes(api)
		h.initUserRoutes(api)
		h.initMessageRoutes(api)
	}
}
