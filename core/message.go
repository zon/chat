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
	User    User
	Content string
}

func CreateMessage(user User, content string) (*Message, error) {
	m := &Message{User: user, Content: content}
	err := DB.Save(&m).Error
	return m, err
}

func GetLatestMessages(messages *[]Message) error {
	return DB.Joins("User").Limit(pageLimit).Order("created_at desc").Find(&messages).Error
}

func GetMessagesBefore(since time.Time, messages *[]Message) error {
	return DB.Joins("User").Limit(pageLimit).Order("created_at desc").Where("created_at < ?", since).Find(&messages).Error
}

func GetMessagesAfter(since time.Time, messages *[]Message) error {
	return DB.Joins("User").Order("created_at desc").Where("created_at > ?", since).Find(&messages).Error
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
