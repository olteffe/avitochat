package usecase

import (
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
)

//go:generate mockery --dir . --name Chat --output ./mocks
type Chat interface {
	CreateChatUseCase(chat *models.Chats) (string, error)
	GetChatUseCase(userId string) ([]*models.Chats, error)
}

//go:generate mockery --dir . --name User --output ./mocks
type User interface {
	CreateUserUseCase(user *models.Users) (string, error)
}

//go:generate mockery --dir . --name Message --output ./mocks
type Message interface {
	SendMessageUseCase(message *models.Messages) (string, error)
	GetMessagesUseCase(message *models.Messages) ([]*models.Messages, error)
}

type UseCase struct {
	User
	Chat
	Message
}

func NewService(repos *repository.Repository) *UseCase {
	return &UseCase{
		User:    NewUserUseCase(repos.User),
		Chat:    NewChatUseCase(repos.Chat),
		Message: NewMessageUseCase(repos.Message),
	}
}
