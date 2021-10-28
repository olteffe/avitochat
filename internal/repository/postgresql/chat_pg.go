package postgresql

import (
	"errors"
	"fmt"
	mError "github.com/olteffe/avitochat/internal/message_error"
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
func (pg *ChatPg) CreateChatRepository(chat *models.Chats) (string, error) {
	// begin transaction
	tx := pg.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return "", fmt.Errorf("CreateChatRepository: %w", mError.ErrDB)
	}
	createChat := tx.Table("chats").Omit("Users").Create(&chat)
	if createChat.Error != nil {
		tx.Rollback()
		return "", fmt.Errorf("CreateChatRepository: %w", mError.ErrDB)
	}
	for _, userID := range chat.Users {
		tx.Exec("INSERT INTO onlines (chat_id, user_id) VALUES (?, ?)", chat.ID, userID)
	}
	return chat.ID.String(), tx.Commit().Error
	// end transaction if return (id, err == nil)
}

// ExistenceChatName func check chat name and users in database
func (pg *ChatPg) ExistenceChatName(chat *models.Chats) error {
	rawChat := pg.db.Table("chats").Limit(1).Where("name = ?", chat.Name).First(&chat)
	if err := rawChat.Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("ExistenceChatName: %w", mError.ErrDB)
	}
	if countChat := rawChat.RowsAffected; countChat != 0 {
		return fmt.Errorf("ExistenceChatName: %w", mError.ErrChatInvalid)
	}
	var lenSliceID []models.Users
	rawUsers := pg.db.Table("users").Select("id").Where(chat.Users).Find(&lenSliceID)
	if rawUsers.Error != nil && !errors.Is(rawUsers.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("ExistenceChatName: %w", mError.ErrDB)
	}
	if countUsers := rawUsers.RowsAffected; countUsers != int64(len(chat.Users)) {
		return fmt.Errorf("ExistenceChatName: %w", mError.ErrUserInvalid)
	}
	return nil
}

func (pg *ChatPg) ExistenceUser(userId string) error {
	var tempUser models.Users
	userIDExist := pg.db.Table("users").Where("id = ?", userId).Limit(1).First(&tempUser)
	if userIDExist.Error != nil {
		if errors.Is(userIDExist.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ExistenceUser: %w", mError.ErrUserIdInvalid)
		}
		return fmt.Errorf("ExistenceUser: %w", mError.ErrDB)
	}
	return nil
}

// GetChatRepository - Get user chats
func (pg *ChatPg) GetChatRepository(userID string) ([]*models.Chats, error) {
	var allChats []*models.Chats
	query := `
		SELECT temp.id, temp.name, temp.created_at
		FROM
		(
			SELECT chats.id, chats.name, chats.created_at,
				(
					SELECT GREATEST(chats.created_at,
						(
							SELECT created_at
							FROM messages
							WHERE chat_id = chats.id
							ORDER BY created_at DESC
							LIMIT 1
						)
					)
				) AS last_message
			FROM chats JOIN onlines ON chats.id = onlines.chat_id
			WHERE onlines.user_id = ?
			ORDER BY last_message
		) AS temp`
	chats := pg.db.Raw(query, userID).Scan(&allChats)
	if chats.Error != nil {
		return nil, fmt.Errorf("GetChatRepository: %w", mError.ErrDB)
	}
	for _, chat := range allChats { // Don't do it, Dudley!(c)
		err := pg.db.Table("onlines").Select("user_id").Where("chat_id = ?", chat.ID).
			Pluck("user_id", &chat.Users).Error
		if err != nil {
			return nil, fmt.Errorf("GetChatRepository: %w", mError.ErrDB)
		}
	}
	return allChats, nil
}
