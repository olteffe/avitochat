package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatForm struct {
	Id        uuid.UUID `json:"id,omitempty" gorm:"column:id"`
	Name      string    `json:"name,omitempty" gorm:"column:name"`
	Users     []string  `json:"users,omitempty" gorm:"many2many:online"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

type MessageForm struct {
	Id        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Chat      string    `json:"chat,omitempty"`
	Author    string    `json:"author,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserForm struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
