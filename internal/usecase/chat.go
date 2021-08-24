package usecase

import (
	"errors"
	"github.com/google/uuid"
	"time"

	"github.com/olteffe/avitochat/internal"
	"github.com/olteffe/avitochat/internal/models"
)

type ChatUC struct {
	ChatRep internal.ChatRepository
}

// CreateChatUseCase func
func (uc *ChatUC) CreateChatUseCase(chat models.Chats) (string, error) {
	// simple validator for name and users.
	if chat.Name == "" || len(chat.Name) > 50 {
		return "", errors.New("invalid chat name")
	}
	if len(chat.Users) < 2 {
		return "", errors.New("two or more users required")
	}
	// uniqueness check
	existChatName, wrongUsers, err := uc.ChatRep.ExistenceChatName(chat)
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
	return uc.ChatRep.CreateChatRepository(chat)
}

func (uc ChatUC) GetChatUseCase() {
	panic("implement me")
}
