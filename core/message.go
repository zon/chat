package core

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const pageLimit int = 10

type Message struct {
	gorm.Model
	UserID  uint
	Content string
}

func CreateMessage(userId uint, content string) (*Message, error) {
	m := &Message{UserID: userId, Content: content}
	err := DB.Save(&m).Error
	return m, err
}

func GetLatestMessages(messages *[]Message) error {
	return DB.Limit(pageLimit).Order("created_at").Find(&messages).Error
}

func GetMessagePage(since time.Time, messages *[]Message) error {
	return DB.Limit(pageLimit).Order("created_at").Where("created_at < ?", since).Find(&messages).Error
}

func (m *Message) HtmlID() string {
	return fmt.Sprintf("msg-%d", m.ID)
}

func (m *Message) IsUpdated() bool {
	return m.UpdatedAt.After(m.CreatedAt)
}

func (m *Message) Update(content string) error {
	m.Content = content
	return DB.Save(&m).Error
}

func (m *Message) Delete() error {
	return DB.Delete(&m).Error
}
