package repository

import (
	"errors"
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type ChatRep struct {
	DB *gorm.DB
}

func (rep *ChatRep) GetMessagesRepository() {
	panic("implement me")
}

func (rep *ChatRep) SendMessageRepository() {
	panic("implement me")
}
