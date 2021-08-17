package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatForm struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Users     []string  `json:"users,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type MessageForm struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Chat      string    `json:"chat,omitempty"`
	Author    string    `json:"author,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserForm struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
