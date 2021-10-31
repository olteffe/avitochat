package models

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID string `json:"id"`
}

type Chats struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Users     []string  `json:"users" gorm:"column:users"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type Messages struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Chat      string    `json:"chat" gorm:"column:chat_id"`
	Author    string    `json:"author" gorm:"column:author_id"`
	Text      string    `json:"text" gorm:"column:text"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type Users struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `json:"author" gorm:"column:username"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type Onlines struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ChatID uuid.UUID `gorm:"column:chat_id"`
	UserID uuid.UUID `gorm:"column:user_id"`
}
