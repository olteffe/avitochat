package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/models"
)

// CreateUser - Create new user
func (c *Container) CreateUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// CreateChat - Create a chat between users
func (c *Container) CreateChat(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetChat - Get all user chats
func (c *Container) GetChat(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetMessages - Get all chat messages
func (c *Container) GetMessages(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// SendMessage - Send a user message
func (c *Container) SendMessage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
