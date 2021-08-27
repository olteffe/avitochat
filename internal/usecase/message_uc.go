package usecase

import (
	"github.com/olteffe/avitochat/internal/repository"
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

func (uc *MessageUseCase) GetMessagesUseCase() {
	panic("implement me")
}

func (uc *MessageUseCase) SendMessageUseCase() {
	panic("implement me")
}
