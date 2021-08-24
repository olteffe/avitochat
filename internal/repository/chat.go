package repository

import (
	"errors"
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type ChatRep struct {
	DB *gorm.DB
}

// CreateChatRepository func create a new chat
func (rep *ChatRep) CreateChatRepository(chat models.Chats) (string, error) {
	// begin transaction
	tx := rep.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return "", err
	}
	createChat := tx.Table("chats").Omit("Users").Create(&chat)
	if createChat.Error != nil {
		tx.Rollback()
		return "", createChat.Error
	}
	for _, userID := range chat.Users {
		tx.Exec("INSERT INTO online (chat_id, user_id) VALUES (?, ?)", chat.ID, userID)
	}
	return chat.ID.String(), tx.Commit().Error
	// end transaction
}

// ExistenceChatName func check chat name and users in database
func (rep *ChatRep) ExistenceChatName(chat models.Chats) (bool, bool, error) {
	rawChat := rep.DB.Table("chats").Limit(1).Where("name = ?", chat.Name).Find(&chat)
	if rawChat.Error != nil && rawChat.Error != gorm.ErrRecordNotFound {
		return false, false, errors.New("database error")
	}
	if countChat := rawChat.RowsAffected; countChat != 0 {
		return true, false, nil
	}
	rawUsers := rep.DB.Table("users").Select("id").Find(&chat, chat.Users)
	if rawUsers.Error != nil && rawUsers.Error != gorm.ErrRecordNotFound {
		return false, false, errors.New("database error")
	}
	if countUsers := rawUsers.RowsAffected; countUsers != int64(len(chat.Users)) {
		return false, true, nil
	}
	return false, false, nil
}

func (rep *ChatRep) GetChatRepository() {
	panic("implement me")
}
