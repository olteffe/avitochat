package models

import (
	"time"

	"github.com/google/uuid"
)

type Chats struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"column:id"`
	Name      string    `json:"name,omitempty" gorm:"column:name"`
	Users     []string  `json:"users,omitempty" gorm:"-"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

type Messages struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Chat      string    `json:"chat,omitempty"`
	Author    string    `json:"author,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Users struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Online struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ChatID uuid.UUID `gorm:"column:chat_id"`
	UserID uuid.UUID `gorm:"column:user_id"`
}
