package usecase

import (
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
)

type Chat interface {
	CreateChatUseCase(chat *models.Chats) (string, error)
	GetChatUseCase(userId string) ([]*models.Chats, error)
}

type User interface {
	CreateUserUseCase(user *models.Users) (string, error)
}

type Message interface {
	GetMessagesUseCase(message *models.Messages) (string, error)
	SendMessageUseCase(chatId string) ([]*models.Messages, error)
}

type UseCase struct {
	User
	Chat
	Message
}

func NewService(repos *repository.Repository) *UseCase {
	return &UseCase{
		User:    NewUserUseCase(repos.User),
		Chat:    NewChatUseCase(repos.Chat, repos.User),
		Message: NewMessageUseCase(repos.Message, repos.User, repos.Chat),
	}
}
