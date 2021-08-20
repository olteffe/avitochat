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

func (uc *ChatUC) CreateUserUseCase() {
	panic("implement me")
}

// CreateChatUseCase func
func (uc *ChatUC) CreateChatUseCase(chat models.ChatForm) (string, error) {
	// simple validator for name and users. verification will be performed before inserting into the DB
	if chat.Name == "" || len(chat.Name) > 50 {
		return "", errors.New("invalid chat name")
	}
	if len(chat.Users) < 2 {
		return "", errors.New("two or more users required")
	}
	// uniqueness check
	existChatName, wrongUsers, err := uc.ChatRep.ExistenceChatName(chat)
	if err != nil {
		return "", errors.New("cannot validate chat name and users. database error: " + err.Error())
	}
	if existChatName {
		return "", errors.New("a chat already exists")
	}
	if wrongUsers {
		return "", errors.New("one or more users do not exist")
	}

	// generate default values
	chat.Id = uuid.New()
	chat.CreatedAt = time.Now()
	return uc.ChatRep.CreateChatRepository(chat)
}

func (uc ChatUC) GetChatUseCase() {
	panic("implement me")
}

func (uc ChatUC) GetMessagesUseCase() {
	panic("implement me")
}

func (uc ChatUC) SendMessageUseCase() {
	panic("implement me")
}

func () CreateUserUseCase() {
	panic("implement me")
}

func () CreateChatUseCase() {
	panic("implement me")
}

func () GetChatUseCase() {
	panic("implement me")
}

func () GetMessagesUseCase() {
	panic("implement me")
}

func () SendMessageUseCase() {
	panic("implement me")
}
