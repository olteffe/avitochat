package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type ChatRep struct {
	DB *gorm.DB
}

func (rep *ChatRep) CreateUserRepository() {

}

// CreateChatRepository func create new chat
func (rep *ChatRep) CreateChatRepository(chat models.ChatForm) (string, error) {

	result := rep.DB.Select("name", "users").Create(&chat)
	if err := result.Error; err != nil {
		return "", errors.New("cannot create chat. database error: " + err.Error())
	}
	return chat.Id.String(), nil
}

// ExistenceChatName func check chat name and users in database
func (rep *ChatRep) ExistenceChatName(chat models.ChatForm) (bool, bool, error) {
	count := int64(0)
	err := rep.DB.Table("chats").
		Select("chatName").
		Where("name=?", chat.Name).
		Limit(1).
		Count(&count).
		Error
	if err != nil {
		exists := count > 0
		return exists, false, err
	}
	// parsing string to uuid TODO

	var onlineUsers []uuid.UUID
	err = rep.DB.Table("users").
		Select("id").
		Where(chat.Users).
		Find(&onlineUsers).
		Error
	if err != nil {
		if len(chat.Users) == len(onlineUsers) {
			return true, false, err
		}
	}
	return true, false, err
}

func (rep *ChatRep) GetChatRepository() {

}

func (rep *ChatRep) GetMessagesRepository() {

}

func (rep *ChatRep) SendMessageRepository() {

}
