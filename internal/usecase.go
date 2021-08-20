package internal

import (
	"github.com/olteffe/avitochat/internal/models"
)

type ChatUseCase interface {
	CreateUserUseCase()
	CreateChatUseCase(chat models.ChatForm) (string, error)
	GetChatUseCase()
	GetMessagesUseCase()
	SendMessageUseCase()
}
