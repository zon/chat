package core

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const UserPath string = "/user"

type User struct {
	gorm.Model
	AuthID string `gorm:"uniqueIndex"`
	Name   string `gorm:"unique"`
	Ready  bool
}

func (u *User) Url() string {
	return UserUrl(u.ID)
}

func (u *User) Save() error {
	return DB.Save(&u).Error
}

func GetUserByAuthID(authID string) (*User, error) {
	var user *User
	err := DB.Limit(1).Find(&user, "auth_id = ?", authID).Error
	if err != nil {
		return user, err
	}
	if user == nil || user.ID == 0 {
		user = &User{
			AuthID: authID,
			Name:   RandomString(16),
		}
		err = user.Save()
	}
	return user, err
}

func GetUser(id uint) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsersAfter(since time.Time, users *[]User) error {
	return DB.Order("updated_at desc").Where("updated_at > ?", since).Find(&users).Error
}

func UserUrl(id uint) string {
	return fmt.Sprintf("%s/%d", UserPath, id)
}
