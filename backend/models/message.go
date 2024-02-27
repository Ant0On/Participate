package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID  uint   `json:"user_id" binding:"required"`
	Email   string `gorm:"size:100" json:"email" binding:"required"`
	Content string `gorm:"not null" json:"content" binding:"required"`
	ChatID  uint   `json:"chat_id" binding:"required"`
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
