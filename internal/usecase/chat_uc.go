package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
)

type ChatUseCase struct {
	repo     repository.Chat
	userRepo repository.User
}

func NewChatUseCase(repo repository.Chat, userRepo repository.User) *ChatUseCase {
	return &ChatUseCase{repo: repo, userRepo: userRepo}
}

// CreateChatUseCase func create new chat
func (uc *ChatUseCase) CreateChatUseCase(chat models.Chats) (string, error) {
	// simple validator for name and users.
	if chat.Name == "" || len(chat.Name) > 50 {
		return "", errors.New("invalid chat name")
	}
	if len(chat.Users) < 2 {
		return "", errors.New("two or more users required")
	}
	// uniqueness check
	existChatName, wrongUsers, err := uc.repo.ExistenceChatName(chat)
	if err != nil {
		return "", errors.New("cannot validate chat name and users. database error")
	}
	if existChatName {
		return "", errors.New("a chat already exists")
	}
	if wrongUsers {
		return "", errors.New("one or more users do not exist")
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
