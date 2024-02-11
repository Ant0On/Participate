package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	OfferID  uint
	Messages []Message
}

func (c *Chat) Save() error {
	if err := DB.Create(&c).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (c *Chat) Update() error {
	if err := DB.Save(&c).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func GetChat(ID string) (*Chat, error) {
	var chat *Chat
	if err := DB.Where("id = ?", ID).First(&chat).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}

	return chat, nil
}

func GetChatByOfferId(offerID string) (*Chat, error) {
	var chat *Chat
	if err := DB.Where("offer_id = ?", offerID).First(&chat).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}

	return chat, nil
}
