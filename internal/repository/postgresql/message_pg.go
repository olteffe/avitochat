package postgresql

import (
	"errors"
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type MessagePg struct {
	db *gorm.DB
}

func NewMessagePg(db *gorm.DB) *MessagePg {
	return &MessagePg{db: db}
}

// GetMessagesRepository func Get all chat messages
func (pg *MessagePg) GetMessagesRepository(chatID string) ([]*models.Messages, error) {
	var allMessages []*models.Messages
	messages := pg.db.Table("messages").Where("chat_id = ?", chatID).
		Order("created_at desc").Scan(&allMessages)
	if messages.Error != nil {
		return nil, messages.Error
	}
	return allMessages, nil
}

// SendMessageRepository func send message in chat
func (pg *MessagePg) SendMessageRepository(message *models.Messages) (string, error) {
	createMessage := pg.db.Table("message").FirstOrCreate(&message)
	if createMessage.Error != nil {
		return "", createMessage.Error
	}
	return message.ID.String(), nil
}

// ExistenceChatAuthor func check chatID and authorID in database
func (pg *MessagePg) ExistenceChatAuthor(message *models.Messages) (bool, error) {
	chatIdNotExist := pg.db.Table("chats").Limit(1).First(&message.Chat, "id = ?", message.Chat)
	if chatIdNotExist.Error != nil {
		if errors.Is(chatIdNotExist.Error, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return true, chatIdNotExist.Error
	}
	authorIDExist := pg.db.Table("users").Limit(1).First(&message.Author, "id = ?", message.Author)
	if authorIDExist.Error != nil {
		if errors.Is(authorIDExist.Error, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return true, authorIDExist.Error
	}
	return false, nil
}
