package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	mError "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
)

type ChatUseCase struct {
	repo repository.Chat
}

func NewChatUseCase(repo repository.Chat) *ChatUseCase {
	return &ChatUseCase{repo: repo}
}

// CreateChatUseCase func create new chat
func (uc *ChatUseCase) CreateChatUseCase(chat *models.Chats) (string, error) {
	// simple validator for name and users.
	if chat.Name == "" || len(chat.Name) > 50 {
		return "", mError.ErrChatInvalid
	}
	if len(chat.Users) < 2 {
		return "", mError.ErrCountUsers
	}
	// uniqueness check
	err := uc.repo.ExistenceChatName(chat)
	if err != nil {
		return "", err
	}
	// generate default values
	chat.ID = uuid.New()
	chat.CreatedAt = time.Now()
	return uc.repo.CreateChatRepository(chat)
}

// GetChatUseCase - Get user chats
func (uc *ChatUseCase) GetChatUseCase(userID string) ([]*models.Chats, error) {
	if _, err := uuid.Parse(userID); err != nil {
		return nil, errors.New("invalid user ID")
	}
	allChats, err := uc.repo.GetChatRepository(userID)
	if err != nil {
		return nil, err
	}
	return allChats, nil
}
