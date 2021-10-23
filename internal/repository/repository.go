package repository

import (
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository/postgresql"
	"gorm.io/gorm"
)

//go:generate mockgen -destination ./mocks/chat.go -package mock_repository github.com/olteffe/avitochat/internal/repository Chat
type Chat interface {
	CreateChatRepository(chat *models.Chats) (string, error)
	ExistenceChatName(chat *models.Chats) error
	GetChatRepository(userId string) ([]*models.Chats, error)
	ExistenceUser(userId string) error
}

//go:generate mockgen -destination ./mocks/user.go -package mock_repository github.com/olteffe/avitochat/internal/repository User
type User interface {
	CreateUserRepository(user *models.Users) (string, error)
	ExistenceUser(user *models.Users) error
}

//go:generate mockgen -destination ./mocks/message.go -package mock_repository github.com/olteffe/avitochat/internal/repository Message
type Message interface {
	GetMessagesRepository(message *models.Messages) ([]*models.Messages, error)
	ExistenceChat(message *models.Messages) error
	ExistenceAuthor(message *models.Messages) error
	SendMessageRepository(message *models.Messages) (string, error)
}

type Repository struct {
	User
	Chat
	Message
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    postgresql.NewUserPg(db),
		Chat:    postgresql.NewChatPg(db),
		Message: postgresql.NewMessagePg(db),
	}
}
