package postgresql

import (
	"errors"
	mErr "github.com/olteffe/avitochat/internal/message_error"
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
func (pg *MessagePg) GetMessagesRepository(message *models.Messages) ([]*models.Messages, error) {
	var allMessages []*models.Messages
	messages := pg.db.Table("messages").Where("chat_id = ?", message.Chat).
		Order("created_at desc").Scan(&allMessages)
	if messages.Error != nil {
		return nil, mErr.ErrDB
	}
	return allMessages, nil
}

// SendMessageRepository func send message in chat
func (pg *MessagePg) SendMessageRepository(message *models.Messages) (string, error) {
	createMessage := pg.db.Table("messages").FirstOrCreate(&message)
	if createMessage.Error != nil {
		return "", mErr.ErrDB
	}
	return message.ID.String(), nil
}

// ExistenceChat func check chat ID in database
func (pg *MessagePg) ExistenceChat(message *models.Messages) error {
	var tempChat models.Chats
	chatIdNotExist := pg.db.Table("chats").Where("id = ?", message.Chat).Limit(1).First(&tempChat)
	if chatIdNotExist.Error != nil {
		if errors.Is(chatIdNotExist.Error, gorm.ErrRecordNotFound) {
			return mErr.ErrChatIdInvalid
		}
		return mErr.ErrDB
	}
	return nil
}

// ExistenceAuthor func check author ID in database
func (pg *MessagePg) ExistenceAuthor(message *models.Messages) error {
	var tempUser models.Users
	authorIDExist := pg.db.Table("users").Where("id = ?", message.Author).Limit(1).First(&tempUser)
	if authorIDExist.Error != nil {
		if errors.Is(authorIDExist.Error, gorm.ErrRecordNotFound) {
			return mErr.ErrUserIdInvalid
		}
		return mErr.ErrDB
	}
	return nil
}
