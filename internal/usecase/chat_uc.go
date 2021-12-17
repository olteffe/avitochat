package usecase

import (
	"fmt"
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
		return "", fmt.Errorf("CreateChatUseCase: %w", mError.ErrChatInvalid)
	}
	if len(chat.Users) < 2 {
		return "", fmt.Errorf("CreateChatUseCase: %w", mError.ErrCountUsers)
	}
	// and additionally injection protect in func ExistenceChatName
	for _, id := range chat.Users {
		if _, err := uuid.Parse(id); err != nil {
			return "", fmt.Errorf("CreateChatUseCase: %w", mError.ErrUserInvalid)
		}
		if err := uc.repo.ExistenceUser(id); err != nil {
			return "", fmt.Errorf("CreateChatUseCase: %w", mError.ErrUserInvalid)
		}
	}

	// uniqueness check
	err := uc.repo.ExistenceChatName(chat)
	if err != nil {
		return "", fmt.Errorf("ExistenceChatName: %w", err)
	}
	// generate default values
	chat.ID = uuid.New()
	chat.CreatedAt = time.Now()
	return uc.repo.CreateChatRepository(chat)
}

// GetChatUseCase - Get user chats
func (uc *ChatUseCase) GetChatUseCase(userID string) ([]*models.Chats, error) {
	if _, err := uuid.Parse(userID); err != nil {
		return nil, fmt.Errorf("GetChatUseCase: Parse: %w", mError.ErrUserInvalid)
	}
	err := uc.repo.ExistenceUser(userID)
	if err != nil {
		return nil, fmt.Errorf("ExistenceUser: %w", err)
	}
	allChats, err := uc.repo.GetChatRepository(userID)
	if err != nil {
		return nil, fmt.Errorf("GetChatUseCase: %w", err)
	}
	return allChats, nil
}
