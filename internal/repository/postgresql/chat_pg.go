package postgresql

import (
	"errors"
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type ChatPg struct {
	db *gorm.DB
}

func NewChatPg(db *gorm.DB) *ChatPg {
	return &ChatPg{db: db}
}

// CreateChatRepository func create a new chat
func (pg *ChatPg) CreateChatRepository(chat models.Chats) (string, error) {
	// begin transaction
	tx := pg.db.Begin()
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
	// end transaction if return (id, err == nil)
}

// ExistenceChatName func check chat name and users in database
func (pg *ChatPg) ExistenceChatName(chat models.Chats) (bool, bool, error) {
	rawChat := pg.db.Table("chats").Limit(1).Where("name = ?", chat.Name).Find(&chat)
	if rawChat.Error != nil && rawChat.Error != gorm.ErrRecordNotFound {
		return false, false, errors.New("database error")
	}
	if countChat := rawChat.RowsAffected; countChat != 0 {
		return true, false, nil
	}
	rawUsers := pg.db.Table("users").Select("id").Find(&chat, chat.Users)
	if rawUsers.Error != nil && rawUsers.Error != gorm.ErrRecordNotFound {
		return false, false, errors.New("database error")
	}
	if countUsers := rawUsers.RowsAffected; countUsers != int64(len(chat.Users)) {
		return false, true, nil
	}
	return false, false, nil
}

func (pg *ChatPg) GetChatRepository(userId string) ([]*models.Chats, error) {
	panic("implement me")
}
