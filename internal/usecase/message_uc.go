package usecase

import (
	"fmt"
	"github.com/google/uuid"
	mErr "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
	"time"
)

type MessageUseCase struct {
	repo repository.Message
}

func NewMessageUseCase(repo repository.Message) *MessageUseCase {
	return &MessageUseCase{
		repo: repo,
	}
}

func (uc *MessageUseCase) GetMessagesUseCase(message *models.Messages) ([]*models.Messages, error) {
	// simple input validation
	if _, err := uuid.Parse(message.Chat); err != nil {
		return nil, fmt.Errorf("GetMessagesUseCase: %w", mErr.ErrChatIdInvalid)
	}
	// db validation
	if err := uc.repo.ExistenceChat(message); err != nil {
		return nil, fmt.Errorf("GetMessagesUseCase: %w", err)
	}
	allMessages, err := uc.repo.GetMessagesRepository(message)
	if err != nil {
		return nil, fmt.Errorf("GetMessagesUseCase: %w", err)
	}
	return allMessages, nil
}

// SendMessageUseCase func send new message
func (uc *MessageUseCase) SendMessageUseCase(message *models.Messages) (string, error) {
	// simple input validation
	if _, err := uuid.Parse(message.Chat); err != nil {
		return "", fmt.Errorf("SendMessageUseCase: %w", mErr.ErrChatIdInvalid)
	}
	if _, err := uuid.Parse(message.Author); err != nil {
		return "", fmt.Errorf("SendMessageUseCase: %w", mErr.ErrUserIdInvalid)
	}
	// db validation
	if err := uc.repo.ExistenceChat(message); err != nil {
		return "", fmt.Errorf("SendMessageUseCase: %w", err)
	}
	if err := uc.repo.ExistenceAuthor(message); err != nil {
		return "", fmt.Errorf("SendMessageUseCase: %w", err)
	}
	message.ID = uuid.New()
	message.CreatedAt = time.Now()
	return uc.repo.SendMessageRepository(message)
}
