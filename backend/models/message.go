package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	CustomerID uint
	UserName   string
	Content    string `gorm:"not null" json:"content"`
	CreatedAt  time.Time
	ChatID     uint
}

func (m *Message) Save() error {
	if err := DB.Create(&m).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func GetAllMessages(chat *Chat) ([]Message, error) {
	var messages []Message

	if err := DB.Model(chat).Association("Messages").Find(&messages); err != nil {
		return nil, fmt.Errorf("DB.Model: %w", err)
	}
	return messages, nil
}
