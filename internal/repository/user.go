package repository

import (
	"errors"
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type ChatRep struct {
	DB *gorm.DB
}

func (rep *ChatRep) CreateUserRepository() {
	panic("implement me")
}
