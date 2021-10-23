package usecase

import (
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
)

//go:generate mockgen -destination ./mocks/chat.go -package mock_usecase github.com/olteffe/avitochat/internal/usecase Chat
type Chat interface {
	CreateChatUseCase(chat *models.Chats) (string, error)
	GetChatUseCase(userId string) ([]*models.Chats, error)
}

//go:generate mockgen -destination ./mocks/user.go -package mock_usecase github.com/olteffe/avitochat/internal/usecase User
type User interface {
	CreateUserUseCase(user *models.Users) (string, error)
}

//go:generate mockgen -destination ./mocks/message.go -package mock_usecase github.com/olteffe/avitochat/internal/usecase Message
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
