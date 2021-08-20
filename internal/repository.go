package internal

import "github.com/olteffe/avitochat/internal/models"

type ChatRepository interface {
	CreateUserRepository()
	CreateChatRepository(chat models.ChatForm) (string, error)
	ExistenceChatName(chat models.ChatForm) (bool, bool, error)
	GetChatRepository()
	GetMessagesRepository()
	SendMessageRepository()
}
