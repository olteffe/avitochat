package models

import (
	"time"
)

// HelloWorld is a sample data structure to make sure each endpoint return something.todo-temp
type HelloWorld struct {
	Message string `json:"message"`
}

type ChatForm struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Users     []string  `json:"users,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreateChat struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

type MessageForm struct {
	Id        string    `json:"id,omitempty"`
	Author    string    `json:"author,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type SendMessage struct {
	Chat   string `json:"chat"`
	Author string `json:"author"`
	Text   string `json:"text"`
}
