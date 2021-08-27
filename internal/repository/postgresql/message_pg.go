package postgresql

import (
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type MessagePg struct {
	db *gorm.DB
}

func NewMessagePg(db *gorm.DB) *MessagePg {
	return &MessagePg{db: db}
}

func (pg *MessagePg) GetMessagesRepository(chatId string) ([]*models.Messages, error) {
	panic("implement me")
}

func (pg *MessagePg) SendMessageRepository(message *models.Messages) (string, error) {
	panic("implement me")
}
