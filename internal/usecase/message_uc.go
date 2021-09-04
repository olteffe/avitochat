package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
	"time"
)

type MessageUseCase struct {
	repo     repository.Message
	userRepo repository.User
	chatRepo repository.Chat
}

func NewMessageUseCase(repo repository.Message, userRepo repository.User, chatRepo repository.Chat) *MessageUseCase {
	return &MessageUseCase{
		repo:     repo,
		userRepo: userRepo,
		chatRepo: chatRepo,
	}
}

func (uc *MessageUseCase) GetMessagesUseCase(chatID string) ([]*models.Messages, error) {
	allMessages, err := uc.repo.GetMessagesRepository(chatID)
	if err != nil {
		return nil, err
	}
	return allMessages, nil
}

// SendMessageUseCase func send new message
func (uc *MessageUseCase) SendMessageUseCase(message *models.Messages) (string, error) {
	// simple input validation
	if _, err := uuid.Parse(message.Chat); err != nil {
		return "invalid chat or author ID", err
	}
	if _, err := uuid.Parse(message.Author); err != nil {
		return "invalid chat or author ID", err
	}
	// db validation
	chatAuthorNotExist, err := uc.repo.ExistenceChatAuthor(message)
	if err != nil {
		return "", errors.New("database error")
	}
	if chatAuthorNotExist {
		return "", errors.New("chat or author not exist")
	}
	message.ID = uuid.New()
	message.CreatedAt = time.Now()
	return uc.repo.SendMessageRepository(message)
}
