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

func (uc ChatUC) GetMessagesUseCase() {
	panic("implement me")
}

func (uc ChatUC) SendMessageUseCase() {
	panic("implement me")
}
