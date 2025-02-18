package core

import (
	"time"

	"gorm.io/gorm"
)

const pageLimit int = 10

type Message struct {
	gorm.Model
	UserID uint
	Text   string
}

func CreateMessage(userId uint, text string) (*Message, error) {
	m := &Message{UserID: userId, Text: text}
	err := DB.Save(&m).Error
	return m, err
}

func (m *Message) Update(text string) error {
	m.Text = text
	return DB.Save(&m).Error
}

func GetLatestMessages(messages *[]Message) error {
	return DB.Limit(pageLimit).Order("created_at").Find(&messages).Error
}

func GetMessagePage(since time.Time, messages *[]Message) error {
	return DB.Limit(pageLimit).Order("created_at").Where("created_at < ?", since).Find(&messages).Error
}

func (m *Message) Delete() error {
	return DB.Delete(&m).Error
}
